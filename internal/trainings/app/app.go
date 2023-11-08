package app

import (
	"github.com/go-kratos/kratos/v2/log"
	trainer "github.com/shiqinfeng1/goMono/api/trainer/v1"
	users "github.com/shiqinfeng1/goMono/api/users/v1"
	"github.com/shiqinfeng1/goMono/internal/common/metrics"
	"github.com/shiqinfeng1/goMono/internal/trainings/adapters"
	"github.com/shiqinfeng1/goMono/internal/trainings/app/command"
	"github.com/shiqinfeng1/goMono/internal/trainings/app/query"
	"github.com/shiqinfeng1/goMono/internal/trainings/domain/training"
)

// ProviderSet is service providers.

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	ApproveTrainingReschedule command.ApproveTrainingRescheduleHandler
	CancelTraining            command.CancelTrainingHandler
	RejectTrainingReschedule  command.RejectTrainingRescheduleHandler
	RescheduleTraining        command.RescheduleTrainingHandler
	RequestTrainingReschedule command.RequestTrainingRescheduleHandler
	ScheduleTraining          command.ScheduleTrainingHandler
}

type Queries struct {
	AllTrainings     query.AllTrainingsHandler
	TrainingsForUser query.TrainingsForUserHandler
}

func NewApplication(logger log.Logger, repo training.Repository, trainerClient trainer.TrainerServiceClient, usersClient users.UsersServiceClient) Application {
	trainerGrpc := adapters.NewTrainerGrpc(trainerClient)
	usersGrpc := adapters.NewUsersGrpc(usersClient)
	return newApplication(logger, repo, trainerGrpc, usersGrpc)

}

// 用于组件测试
func NewComponentTestApplication(logger log.Logger, repo training.Repository) Application {
	return newApplication(logger, repo, TrainerServiceMock{}, UserServiceMock{})
}
func newApplication(logger log.Logger, repo training.Repository, trainerService command.TrainerService, userService command.UserService) Application {

	metricsClient := metrics.NoOp{}

	return Application{
		Commands: Commands{
			ApproveTrainingReschedule: command.NewApproveTrainingRescheduleHandler(repo, userService, trainerService, logger, metricsClient),
			CancelTraining:            command.NewCancelTrainingHandler(repo, userService, trainerService, logger, metricsClient),
			RejectTrainingReschedule:  command.NewRejectTrainingRescheduleHandler(repo, logger, metricsClient),
			RescheduleTraining:        command.NewRescheduleTrainingHandler(repo, userService, trainerService, logger, metricsClient),
			RequestTrainingReschedule: command.NewRequestTrainingRescheduleHandler(repo, logger, metricsClient),
			ScheduleTraining:          command.NewScheduleTrainingHandler(repo, userService, trainerService, logger, metricsClient),
		},
		Queries: Queries{
			AllTrainings:     query.NewAllTrainingsHandler(repo, logger, metricsClient),
			TrainingsForUser: query.NewTrainingsForUserHandler(repo, logger, metricsClient),
		},
	}
}
