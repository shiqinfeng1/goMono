package app

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/shiqinfeng1/goMono/internal/common/metrics"
	"github.com/shiqinfeng1/goMono/internal/trainer/app/command"
	"github.com/shiqinfeng1/goMono/internal/trainer/app/query"
	"github.com/shiqinfeng1/goMono/internal/trainer/domain/hour"
	"github.com/shiqinfeng1/goMono/internal/trainings/app"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CancelTraining   command.CancelTrainingHandler
	ScheduleTraining command.ScheduleTrainingHandler

	MakeHoursAvailable   command.MakeHoursAvailableHandler
	MakeHoursUnavailable command.MakeHoursUnavailableHandler
}

type Queries struct {
	HourAvailability      query.HourAvailabilityHandler
	TrainerAvailableHours query.AvailableHoursHandler
}

func NewApplication(ctx context.Context) app.Application {
	firestoreClient, err := firestore.NewClient(ctx, os.Getenv("GCP_PROJECT"))
	if err != nil {
		panic(err)
	}

	factoryConfig := hour.FactoryConfig{
		MaxWeeksInTheFutureToSet: 6,
		MinUtcHour:               12,
		MaxUtcHour:               20,
	}

	datesRepository := adapters.NewDatesFirestoreRepository(firestoreClient, factoryConfig)

	hourFactory, err := hour.NewFactory(factoryConfig)
	if err != nil {
		panic(err)
	}

	hourRepository := adapters.NewFirestoreHourRepository(firestoreClient, hourFactory)

	logger := log.NewStdLogger(io.stderr)
	metricsClient := metrics.NoOp{}

	return app.Application{
		Commands: app.Commands{
			CancelTraining:       command.NewCancelTrainingHandler(hourRepository, logger, metricsClient),
			ScheduleTraining:     command.NewScheduleTrainingHandler(hourRepository, logger, metricsClient),
			MakeHoursAvailable:   command.NewMakeHoursAvailableHandler(hourRepository, logger, metricsClient),
			MakeHoursUnavailable: command.NewMakeHoursUnavailableHandler(hourRepository, logger, metricsClient),
		},
		Queries: app.Queries{
			HourAvailability:      query.NewHourAvailabilityHandler(hourRepository, logger, metricsClient),
			TrainerAvailableHours: query.NewAvailableHoursHandler(datesRepository, logger, metricsClient),
		},
	}
}
