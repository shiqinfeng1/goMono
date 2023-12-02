package adapters

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/shiqinfeng1/goMono/app/common/config"
	"github.com/shiqinfeng1/goMono/app/trainer/internal/application/query"
	"github.com/shiqinfeng1/goMono/app/trainer/internal/domain/hour"
)

type DateModel struct {
	Date         time.Time   `db:"date"`
	HasFreeHours bool        `db:"has_free_hours"`
	Hours        []HourModel `db:"hours"`
}

type HourModel struct {
	Available            bool      `db:"available"`
	HasTrainingScheduled bool      `db:"has_training_scheduled"`
	Hour                 time.Time `db:"hour"`
}

type DatesMysqlRepo struct {
	db          *sqlx.DB
	hourFactory hour.Factory
	log         *log.Helper
}

func NewDatesMysqlRepo(pubCfg *config.Adapter, logger log.Logger) query.QueryRepository {
	db, err := sqlx.Connect(pubCfg.Database.Driver, pubCfg.Database.Source)
	if err != nil {
		return nil
	}
	return &DatesMysqlRepo{
		log: log.NewHelper(log.With(logger, "module", "trainer/adapters/query")),
		db:  db,
		hourFactory: hour.MustNewFactory(hour.FactoryConfig{
			MaxWeeksInTheFutureToSet: 100,
			MinUtcHour:               0,
			MaxUtcHour:               24,
		}),
	}
}

// 查询可用的Hour
func (d DatesMysqlRepo) AvailableHours(ctx context.Context, from time.Time, to time.Time) (query.Date, error) {
	date := DateModel{}
	// todo：query date from db，then convert date to Date
	qd := dateModelToApp(date)
	return qd, nil
}

func dateModelToApp(dm DateModel) query.Date {
	var hours []query.Hour
	for _, h := range dm.Hours {
		hours = append(hours, query.Hour{
			Available:            h.Available,
			HasTrainingScheduled: h.HasTrainingScheduled,
			Hour:                 h.Hour,
		})
	}
	return query.Date{
		Date:         dm.Date,
		HasFreeHours: dm.HasFreeHours,
		Hours:        hours,
	}
}
