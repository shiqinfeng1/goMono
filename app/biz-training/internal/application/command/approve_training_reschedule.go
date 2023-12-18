package command

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shiqinfeng1/goMono/app/biz-training/internal/domain/training"
	"github.com/shiqinfeng1/goMono/app/common/decorator"
)

type ApproveTrainingReschedule struct {
	TrainingUUID string
	User         training.User
}

type ApproveTrainingRescheduleHandler decorator.CommandHandler[ApproveTrainingReschedule]

type approveTrainingRescheduleHandler struct {
	repo           training.Repository
	userService    UserService
	trainerService TrainerService
}

func NewApproveTrainingRescheduleHandler(
	repo training.Repository,
	userService UserService,
	trainerService TrainerService,
	logger log.Logger,
) decorator.CommandHandler[ApproveTrainingReschedule] {
	if repo == nil {
		panic("nil repo")
	}
	if userService == nil {
		panic("nil userService")
	}
	if trainerService == nil {
		panic("nil trainerService")
	}

	return decorator.ApplyCommandDecorators[ApproveTrainingReschedule](
		approveTrainingRescheduleHandler{repo, userService, trainerService},
		logger,
	)
}

func (h approveTrainingRescheduleHandler) Handle(ctx context.Context, cmd ApproveTrainingReschedule) (err error) {

	return h.repo.UpdateTraining(
		ctx,
		cmd.TrainingUUID,
		cmd.User,
		func(ctx context.Context, tr *training.Training) (*training.Training, error) {
			originalTrainingTime := tr.Time()

			if err := tr.ApproveReschedule(cmd.User.Type()); err != nil {
				return nil, err
			}

			err := h.trainerService.MoveTraining(ctx, tr.Time(), originalTrainingTime)
			if err != nil {
				return nil, err
			}

			return tr, nil
		},
	)
}
