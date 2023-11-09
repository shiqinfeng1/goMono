package command

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shiqinfeng1/goMono/internal/common/decorator"
	"github.com/shiqinfeng1/goMono/internal/common/errors"
	"github.com/shiqinfeng1/goMono/internal/trainer/domain/hour"
)

type CancelTraining struct {
	Hour time.Time
}

type CancelTrainingHandler decorator.CommandHandler[CancelTraining]

type cancelTrainingHandler struct {
	hourRepo hour.CmdRepository
}

func NewCancelTrainingHandler(
	hourRepo hour.CmdRepository,
	logger log.Logger,
	metricsClient decorator.MetricsClient,
) CancelTrainingHandler {
	if hourRepo == nil {
		panic("nil hourRepo")
	}

	return decorator.ApplyCommandDecorators[CancelTraining](
		cancelTrainingHandler{hourRepo: hourRepo},
		logger,
		metricsClient,
	)
}

func (h cancelTrainingHandler) Handle(ctx context.Context, cmd CancelTraining) error {
	if err := h.hourRepo.UpdateHour(ctx, cmd.Hour, func(h *hour.Hour) (*hour.Hour, error) {
		if err := h.CancelTraining(); err != nil {
			return nil, err
		}
		return h, nil
	}); err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-update-availability")
	}

	return nil
}
