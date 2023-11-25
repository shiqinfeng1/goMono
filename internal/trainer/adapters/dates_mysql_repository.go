package adapters

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/shiqinfeng1/goMono/internal/common/config"
	"github.com/shiqinfeng1/goMono/internal/trainer/app/query"
	"github.com/shiqinfeng1/goMono/internal/trainer/domain/hour"
)

type DateModel struct {
	Date         time.Time   `firestore:"Date"`
	HasFreeHours bool        `firestore:"HasFreeHours"`
	Hours        []HourModel `firestore:"Hours"`
}

type HourModel struct {
	Available            bool      `firestore:"Available"`
	HasTrainingScheduled bool      `firestore:"HasTrainingScheduled"`
	Hour                 time.Time `firestore:"Hour"`
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
		log: log.NewHelper(logger),
		db:  db,
		hourFactory: hour.MustNewFactory(hour.FactoryConfig{
			MaxWeeksInTheFutureToSet: 100,
			MinUtcHour:               0,
			MaxUtcHour:               24,
		}),
	}
}

func (d DatesMysqlRepo) AvailableHours(ctx context.Context, from time.Time, to time.Time) (query.Date, error) {
	date := DateModel{}
	// date to Date
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
