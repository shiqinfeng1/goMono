package command

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shiqinfeng1/goMono/internal/common/decorator"
	"github.com/shiqinfeng1/goMono/internal/training/domain/training"
)

type RejectTrainingReschedule struct {
	TrainingUUID string
	User         training.User
}

type RejectTrainingRescheduleHandler decorator.CommandHandler[RejectTrainingReschedule]

type rejectTrainingRescheduleHandler struct {
	repo training.Repository
}

func NewRejectTrainingRescheduleHandler(
	repo training.Repository,
	logger log.Logger,
	metricsClient decorator.MetricsClient,
) RejectTrainingRescheduleHandler {
	if repo == nil {
		panic("nil repo service")
	}

	return decorator.ApplyCommandDecorators[RejectTrainingReschedule](
		rejectTrainingRescheduleHandler{repo: repo},
		logger,
		metricsClient,
	)
}

func (h rejectTrainingRescheduleHandler) Handle(ctx context.Context, cmd RejectTrainingReschedule) (err error) {

	return h.repo.UpdateTraining(
		ctx,
		cmd.TrainingUUID,
		cmd.User,
		func(ctx context.Context, tr *training.Training) (*training.Training, error) {
			if err := tr.RejectReschedule(); err != nil {
				return nil, err
			}

			return tr, nil
		},
	)
}
