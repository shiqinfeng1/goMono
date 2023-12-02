package application

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/shiqinfeng1/goMono/internal/trainer/application/command"
	"github.com/shiqinfeng1/goMono/internal/trainer/application/query"
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

	return Application{
		Commands: Commands{
			CancelTraining:       command.NewCancelTrainingHandler(hourCmdRepo, logger),
			ScheduleTraining:     command.NewScheduleTrainingHandler(hourCmdRepo, logger),
			MakeHoursAvailable:   command.NewMakeHoursAvailableHandler(hourCmdRepo, logger),
			MakeHoursUnavailable: command.NewMakeHoursUnavailableHandler(hourCmdRepo, logger),
		},
		Queries: Queries{
			HourAvailability:      query.NewHourAvailabilityHandler(hourCmdRepo, logger),
			TrainerAvailableHours: query.NewAvailableHoursHandler(hourQueryRepo, logger),
		},
	}
}
