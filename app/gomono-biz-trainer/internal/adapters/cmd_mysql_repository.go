package adapters

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/shiqinfeng1/goMono/app/gomono-biz-trainer/internal/domain/hour"
	"github.com/shiqinfeng1/goMono/utils/conf"
	"go.uber.org/multierr"
)

var (
	ErrNotExistHour   = func(err error) error { return fmt.Errorf("unable to get hour from db:%w", err) }
	ErrStartTxnFail   = func(err error) error { return fmt.Errorf("unable to start transaction:%w", err) }
	ErrUpsertHourFail = func(err error) error { return fmt.Errorf("unable to upsert hour:%w", err) }
)

// 数据库表项对应的结构体
type mysqlHour struct {
	ID           string    `db:"id"`
	Hour         time.Time `db:"hour"`
	Availability string    `db:"availability"`
}

// HourRepo 实现领域层repo的数据结构
type HourRepo struct {
	db          *sqlx.DB // 这里演示使用sqlx操作数据库
	log         *log.Helper
	hourFactory hour.Factory
}

// NewTrainingRepo 实现领域层定义的repo，参数通过wire初始化自动注入
func NewHourRepo(cfg *conf.Adapter, logger log.Logger) hour.CmdRepo {
	// 使用sqlx连接到数据库
	db, err := sqlx.Connect(cfg.Database.Driver, cfg.Database.Source)
	if err != nil {
		return nil
	}
	return &HourRepo{
		log: log.NewHelper(log.With(logger, "module", "trainer/adapters/cmd")), // 区分各模块的log
		db:  db,
		hourFactory: hour.MustNewFactory(hour.FactoryConfig{ // 该hourFactory 用来生成领域层的Hour实体结构
			MaxWeeksInTheFutureToSet: 500,
			MinUtcHour:               0,
			MaxUtcHour:               24,
		}),
	}
}

// sqlContextGetter 封装了一个能同时用于事务和普通连接的查询接口
type sqlContextGetter interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

// GetHour 查询数据库，并返回领域实体
func (m HourRepo) GetHour(ctx context.Context, time time.Time) (*hour.Hour, error) {
	return m.getOrCreateHour(ctx, m.db, time, false)
}

// 查询数据并返回领域层的数据结构
func (m HourRepo) getOrCreateHour(
	ctx context.Context,
	db sqlContextGetter,
	hourTime time.Time,
	forUpdate bool,
) (*hour.Hour, error) {
	dbHour := mysqlHour{}

	query := "SELECT * FROM `hours` WHERE `hour` = ?"
	if forUpdate { // 用于事务时的查询
		query += " FOR UPDATE"
	}

	err := db.GetContext(ctx, &dbHour, query, hourTime.UTC())
	if errors.Is(err, sql.ErrNoRows) {
		// 如果不存在，也返回一个新的Hour给领域层处理
		return m.hourFactory.NewNotAvailableHour(hourTime)
	} else if err != nil { // 其他错误
		return nil, ErrNotExistHour(err)
	}
	// 解析数据，获取可用性，解析的实现由领域层实现
	availability, err := hour.NewAvailabilityFromString(dbHour.Availability)
	if err != nil {
		return nil, err
	}
	// 工厂负责组装领域层需要的Hour实体，工厂也是由领域层实现
	domainHour, err := m.hourFactory.UnmarshalHourFromDatabase(dbHour.Hour.Local(), availability)
	if err != nil {
		return nil, err
	}

	return domainHour, nil
}

const mySQLDeadlockErrorCode = 1213

// UpdateHour 更新数据
func (m HourRepo) UpdateHour(
	ctx context.Context,
	hourTime time.Time,
	updateFn func(h *hour.Hour) (*hour.Hour, error),
) error {
	for {
		err := m.updateHour(ctx, hourTime, updateFn)
		// 在死锁错误时，需要再次重试
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
	// 开始一个事务
	tx, err := m.db.Beginx()
	if err != nil {
		return ErrStartTxnFail(err)
	}
	// 定义结束事务的处理，如果根据err进行提交事务或回滚，并返回合并的err
	defer func() {
		err = m.finishTransaction(err, tx)
	}()
	// 查询或创建一个新的领域实体，注意: update标记为true
	existingHour, err := m.getOrCreateHour(ctx, tx, hourTime, true)
	if err != nil {
		return err
	}
	// 更新数据，更新的逻辑由外面的app层决定，输入旧数据，输出新的数据
	updatedHour, err := updateFn(existingHour)
	if err != nil {
		return err
	}
	// 持久化更新后的数据
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
		return ErrUpsertHourFail(err)
	}

	return nil
}

// 如果有err，就回滚；如果没有，则提交事务
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
