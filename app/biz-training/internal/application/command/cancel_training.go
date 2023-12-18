package command

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"github.com/shiqinfeng1/goMono/app/biz-training/internal/domain/training"
	"github.com/shiqinfeng1/goMono/app/common/decorator"
)

type CancelTraining struct {
	TrainingUUID string
	User         training.User
}

type CancelTrainingHandler decorator.CommandHandler[CancelTraining]

type cancelTrainingHandler struct {
	repo           training.Repository
	userService    UserService
	trainerService TrainerService
}

func NewCancelTrainingHandler(
	repo training.Repository,
	userService UserService,
	trainerService TrainerService,
	logger log.Logger,
) decorator.CommandHandler[CancelTraining] {
	if repo == nil {
		panic("nil repo")
	}
	if userService == nil {
		panic("nil user service")
	}
	if trainerService == nil {
		panic("nil trainer service")
	}

	return decorator.ApplyCommandDecorators[CancelTraining](
		cancelTrainingHandler{repo: repo, userService: userService, trainerService: trainerService},
		logger,
	)
}

func (h cancelTrainingHandler) Handle(ctx context.Context, cmd CancelTraining) (err error) {

	return h.repo.UpdateTraining(
		ctx,
		cmd.TrainingUUID,
		cmd.User,
		func(ctx context.Context, tr *training.Training) (*training.Training, error) {
			if err := tr.Cancel(); err != nil {
				return nil, err
			}

			if balanceDelta := training.CancelBalanceDelta(*tr, cmd.User.Type()); balanceDelta != 0 {
				err := h.userService.UpdateTrainingBalance(ctx, tr.UserUUID(), balanceDelta)
				if err != nil {
					return nil, errors.Wrap(err, "unable to change training balance")
				}
			}

			if err := h.trainerService.CancelTraining(ctx, tr.Time()); err != nil {
				return nil, errors.Wrap(err, "unable to cancel training")
			}

			return tr, nil
		},
	)
}
