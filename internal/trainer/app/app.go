package app

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/shiqinfeng1/goMono/internal/common/metrics"
	"github.com/shiqinfeng1/goMono/internal/trainer/app/command"
	"github.com/shiqinfeng1/goMono/internal/trainer/app/query"
	"github.com/shiqinfeng1/goMono/internal/trainer/domain/hour"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CancelTraining       command.CancelTrainingHandler
	ScheduleTraining     command.ScheduleTrainingHandler
	MakeHoursAvailable   command.MakeHoursAvailableHandler
	MakeHoursUnavailable command.MakeHoursUnavailableHandler
}

type Queries struct {
	HourAvailability      query.HourAvailabilityHandler
	TrainerAvailableHours query.AvailableHoursHandler
}

func NewApplication(logger log.Logger, hourCmdRepo hour.CmdRepo, hourQueryRepo query.QueryRepository) Application {
	return newApplication(logger, hourCmdRepo, hourQueryRepo)

}

func newApplication(logger log.Logger, hourCmdRepo hour.CmdRepo, hourQueryRepo query.QueryRepository) Application {

	metricsClient := metrics.NoOp{}

	return Application{
		Commands: Commands{
			CancelTraining:       command.NewCancelTrainingHandler(hourCmdRepo, logger, metricsClient),
			ScheduleTraining:     command.NewScheduleTrainingHandler(hourCmdRepo, logger, metricsClient),
			MakeHoursAvailable:   command.NewMakeHoursAvailableHandler(hourCmdRepo, logger, metricsClient),
			MakeHoursUnavailable: command.NewMakeHoursUnavailableHandler(hourCmdRepo, logger, metricsClient),
		},
		Queries: Queries{
			HourAvailability:      query.NewHourAvailabilityHandler(hourCmdRepo, logger, metricsClient),
			TrainerAvailableHours: query.NewAvailableHoursHandler(hourQueryRepo, logger, metricsClient),
		},
	}
}
