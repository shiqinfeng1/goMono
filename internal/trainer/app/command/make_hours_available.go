package command

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shiqinfeng1/goMono/internal/common/decorator"
	"github.com/shiqinfeng1/goMono/internal/common/errors"
	"github.com/shiqinfeng1/goMono/internal/trainer/domain/hour"
)

type MakeHoursAvailable struct {
	Hours []time.Time
}

type MakeHoursAvailableHandler decorator.CommandHandler[MakeHoursAvailable]

type makeHoursAvailableHandler struct {
	hourRepo hour.CmdRepository
}

func NewMakeHoursAvailableHandler(
	hourRepo hour.CmdRepository,
	logger log.Logger,
	metricsClient decorator.MetricsClient,
) MakeHoursAvailableHandler {
	if hourRepo == nil {
		panic("hourRepo is nil")
	}

	return decorator.ApplyCommandDecorators[MakeHoursAvailable](
		makeHoursAvailableHandler{hourRepo: hourRepo},
		logger,
		metricsClient,
	)
}

func (c makeHoursAvailableHandler) Handle(ctx context.Context, cmd MakeHoursAvailable) error {
	for _, hourToUpdate := range cmd.Hours {
		if err := c.hourRepo.UpdateHour(ctx, hourToUpdate, func(h *hour.Hour) (*hour.Hour, error) {
			if err := h.MakeAvailable(); err != nil {
				return nil, err
			}
			return h, nil
		}); err != nil {
			return errors.NewSlugError(err.Error(), "unable-to-update-availability")
		}
	}

	return nil
}
