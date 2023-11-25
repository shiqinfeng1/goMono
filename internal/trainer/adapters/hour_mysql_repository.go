package adapters

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/shiqinfeng1/goMono/internal/common/config"
	"github.com/shiqinfeng1/goMono/internal/trainer/domain/hour"
	"go.uber.org/multierr"
)

type mysqlHour struct {
	ID           string    `db:"id"`
	Hour         time.Time `db:"hour"`
	Availability string    `db:"availability"`
}

type HourRepo struct {
	db          *sqlx.DB
	log         *log.Helper
	hourFactory hour.Factory
}

// NewTrainingRepo .
func NewHourRepo(adpr *config.Adapter, logger log.Logger) hour.CmdRepo {
	db, err := sqlx.Connect(adpr.Database.Driver, adpr.Database.Source)
	if err != nil {
		return nil
	}
	return &HourRepo{
		log: log.NewHelper(logger),
		db:  db,
		hourFactory: hour.MustNewFactory(hour.FactoryConfig{
			// 500 weeks gives us enough entropy to avoid duplicated dates
			// (even if duplicate dates should be not a problem)
			MaxWeeksInTheFutureToSet: 500,
			MinUtcHour:               0,
			MaxUtcHour:               24,
		}),
	}
}

// sqlContextGetter is an interface provided both by transaction and standard db connection
type sqlContextGetter interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

func (m HourRepo) GetHour(ctx context.Context, time time.Time) (*hour.Hour, error) {
	return m.getOrCreateHour(ctx, m.db, time, false)
}

func (m HourRepo) getOrCreateHour(
	ctx context.Context,
	db sqlContextGetter,
	hourTime time.Time,
	forUpdate bool,
) (*hour.Hour, error) {
	dbHour := mysqlHour{}

	query := "SELECT * FROM `hours` WHERE `hour` = ?"
	if forUpdate {
		query += " FOR UPDATE"
	}

	err := db.GetContext(ctx, &dbHour, query, hourTime.UTC())
	if errors.Is(err, sql.ErrNoRows) {
		// in reality this date exists, even if it's not persisted
		return m.hourFactory.NewNotAvailableHour(hourTime)
	} else if err != nil {
		return nil, errors.Wrap(err, "unable to get hour from db")
	}

	availability, err := hour.NewAvailabilityFromString(dbHour.Availability)
	if err != nil {
		return nil, err
	}

	domainHour, err := m.hourFactory.UnmarshalHourFromDatabase(dbHour.Hour.Local(), availability)
	if err != nil {
		return nil, err
	}

	return domainHour, nil
}

const mySQLDeadlockErrorCode = 1213

func (m HourRepo) UpdateHour(
	ctx context.Context,
	hourTime time.Time,
	updateFn func(h *hour.Hour) (*hour.Hour, error),
) error {
	for {
		err := m.updateHour(ctx, hourTime, updateFn)

		if val, ok := errors.Cause(err).(*mysql.MySQLError); ok && val.Number == mySQLDeadlockErrorCode {
			continue
		}

		return err
	}
}

func (m HourRepo) updateHour(
	ctx context.Context,
	hourTime time.Time,
	updateFn func(h *hour.Hour) (*hour.Hour, error),
) (err error) {
	tx, err := m.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "unable to start transaction")
	}

	// Defer is executed on function just before exit.
	// With defer, we are always sure that we will close our transaction properly.
	defer func() {
		// In `UpdateHour` we are using named return - `(err error)`.
		// Thanks to that, that can check if function exits with error.
		//
		// Even if function exits without error, commit still can return error.
		// In that case we can override nil to err `err = m.finish...`.
		err = m.finishTransaction(err, tx)
	}()

	existingHour, err := m.getOrCreateHour(ctx, tx, hourTime, true)
	if err != nil {
		return err
	}

	updatedHour, err := updateFn(existingHour)
	if err != nil {
		return err
	}

	if err := m.upsertHour(tx, updatedHour); err != nil {
		return err
	}

	return nil
}

// upsertHour updates hour if hour already exists in the database.
// If your doesn't exists, it's inserted.
func (m HourRepo) upsertHour(tx *sqlx.Tx, hourToUpdate *hour.Hour) error {
	updatedDbHour := mysqlHour{
		Hour:         hourToUpdate.Time().UTC(),
		Availability: hourToUpdate.Availability().String(),
	}

	_, err := tx.NamedExec(
		`INSERT INTO 
			hours (hour, availability) 
		VALUES 
			(:hour, :availability)
		ON DUPLICATE KEY UPDATE 
			availability = :availability`,
		updatedDbHour,
	)
	if err != nil {
		return errors.Wrap(err, "unable to upsert hour")
	}

	return nil
}

// finishTransaction rollbacks transaction if error is provided.
// If err is nil transaction is committed.
//
// If the rollback fails, we are using multierr library to add error about rollback failure.
// If the commit fails, commit error is returned.
func (m HourRepo) finishTransaction(err error, tx *sqlx.Tx) error {
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return multierr.Combine(err, rollbackErr)
		}

		return err
	} else {
		if commitErr := tx.Commit(); commitErr != nil {
			return errors.Wrap(err, "failed to commit tx")
		}

		return nil
	}
}
