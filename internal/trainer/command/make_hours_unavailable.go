package command

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shiqinfeng1/goMono/internal/common/decorator"
	"github.com/shiqinfeng1/goMono/internal/common/errors"
	"github.com/shiqinfeng1/goMono/internal/trainer/domain/hour"
)

type MakeHoursUnavailable struct {
	Hours []time.Time
}

type MakeHoursUnavailableHandler decorator.CommandHandler[MakeHoursUnavailable]

type makeHoursUnavailableHandler struct {
	hourRepo hour.CmdRepository
}

func NewMakeHoursUnavailableHandler(
	hourRepo hour.CmdRepository,
	logger log.Logger,
	metricsClient decorator.MetricsClient,
) MakeHoursUnavailableHandler {
	if hourRepo == nil {
		panic("hourRepo is nil")
	}

	return decorator.ApplyCommandDecorators[MakeHoursUnavailable](
		makeHoursUnavailableHandler{hourRepo: hourRepo},
		logger,
		metricsClient,
	)
}

func (c makeHoursUnavailableHandler) Handle(ctx context.Context, cmd MakeHoursUnavailable) error {
	for _, hourToUpdate := range cmd.Hours {
		if err := c.hourRepo.UpdateHour(ctx, hourToUpdate, func(h *hour.Hour) (*hour.Hour, error) {
			if err := h.MakeNotAvailable(); err != nil {
				return nil, err
			}
			return h, nil
		}); err != nil {
			return errors.NewSlugError(err.Error(), "unable-to-update-availability")
		}
	}

	return nil
}
