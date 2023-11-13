package app

import (
	"github.com/go-kratos/kratos/v2/log"
	trainer "github.com/shiqinfeng1/goMono/api/trainer/v1"
	user "github.com/shiqinfeng1/goMono/api/user/v1"
	"github.com/shiqinfeng1/goMono/internal/common/metrics"
	"github.com/shiqinfeng1/goMono/internal/training/adapters"
	"github.com/shiqinfeng1/goMono/internal/training/app/command"
	"github.com/shiqinfeng1/goMono/internal/training/app/query"
	"github.com/shiqinfeng1/goMono/internal/training/domain/training"
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
	AllTraining     query.AllTrainingHandler
	TrainingForUser query.TrainingForUserHandler
}

func NewApplication(logger log.Logger, repo training.Repository, trainerClient trainer.TrainerServiceClient, userClient user.UserServiceClient) Application {
	trainerGrpc := adapters.NewTrainerGrpc(trainerClient)
	userGrpc := adapters.NewUserGrpc(userClient)
	return newApplication(logger, repo, trainerGrpc, userGrpc)

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
		// todo: 读写分离后，最好不要使用同一个repo
		Queries: Queries{
			AllTraining:     query.NewAllTrainingHandler(repo, logger, metricsClient),
			TrainingForUser: query.NewTrainingForUserHandler(repo, logger, metricsClient),
		},
	}
}