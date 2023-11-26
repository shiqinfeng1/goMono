package command

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/shiqinfeng1/goMono/api/trainer/v1"
	"github.com/shiqinfeng1/goMono/internal/common/decorator"
	"github.com/shiqinfeng1/goMono/internal/trainer/domain/hour"
)

type MakeHoursAvailable struct {
	Hours []time.Time
}

type MakeHoursAvailableHandler decorator.CommandHandler[MakeHoursAvailable]

type makeHoursAvailableHandler struct {
	hourRepo hour.CmdRepo
}

func NewMakeHoursAvailableHandler(
	hourRepo hour.CmdRepo,
	logger log.Logger,
) MakeHoursAvailableHandler {
	if hourRepo == nil {
		panic("hourRepo is nil")
	}

	return decorator.ApplyCommandDecorators[MakeHoursAvailable](
		makeHoursAvailableHandler{hourRepo: hourRepo},
		logger,
	)
}

func (h makeHoursAvailableHandler) Handle(ctx context.Context, cmd MakeHoursAvailable) error {
	for _, hourToUpdate := range cmd.Hours {
		if err := h.hourRepo.UpdateHour(ctx, hourToUpdate, func(hr *hour.Hour) (*hour.Hour, error) {
			if err := hr.MakeAvailable(); err != nil {
				return nil, err
			}
			return hr, nil
		}); err != nil {
			return v1.ErrorUpdateAvailabilityFail("unable to enable availability").WithCause(err)
		}
	}
	return nil
}
