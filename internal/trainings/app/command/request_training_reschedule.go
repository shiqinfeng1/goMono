package command

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shiqinfeng1/goMono/internal/common/decorator"
	"github.com/shiqinfeng1/goMono/internal/common/logs"
	"github.com/shiqinfeng1/goMono/internal/trainings/domain/training"
)

type RequestTrainingReschedule struct {
	TrainingUUID string
	NewTime      time.Time

	User training.User

	NewNotes string
}

type RequestTrainingRescheduleHandler decorator.CommandHandler[RequestTrainingReschedule]

type requestTrainingRescheduleHandler struct {
	repo training.Repository
}

func NewRequestTrainingRescheduleHandler(
	repo training.Repository,
	logger log.Logger,
	metricsClient decorator.MetricsClient,
) RequestTrainingRescheduleHandler {
	if repo == nil {
		panic("nil repo service")
	}

	return decorator.ApplyCommandDecorators[RequestTrainingReschedule](
		requestTrainingRescheduleHandler{repo: repo},
		logger,
		metricsClient,
	)
}

func (h requestTrainingRescheduleHandler) Handle(ctx context.Context, cmd RequestTrainingReschedule) (err error) {
	defer func() {
		logs.LogCommandExecution("RequestTrainingReschedule", cmd, err)
	}()

	return h.repo.UpdateTraining(
		ctx,
		cmd.TrainingUUID,
		cmd.User,
		func(ctx context.Context, tr *training.Training) (*training.Training, error) {
			if err := tr.UpdateNotes(cmd.NewNotes); err != nil {
				return nil, err
			}

			tr.ProposeReschedule(cmd.NewTime, cmd.User.Type())

			return tr, nil
		},
	)
}
