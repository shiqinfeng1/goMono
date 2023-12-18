package command

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/shiqinfeng1/goMono/api/trainer/v1"
	"github.com/shiqinfeng1/goMono/app/biz-trainer/internal/domain/hour"
	"github.com/shiqinfeng1/goMono/app/common/decorator"
)

type MakeHoursUnavailable struct {
	Hours []time.Time
}

type MakeHoursUnavailableHandler decorator.CommandHandler[MakeHoursUnavailable]

type makeHoursUnavailableHandler struct {
	hourRepo hour.CmdRepo
}

func NewMakeHoursUnavailableHandler(
	hourRepo hour.CmdRepo,
	logger log.Logger,
) MakeHoursUnavailableHandler {
	if hourRepo == nil {
		panic("hourRepo is nil")
	}

	return decorator.ApplyCommandDecorators[MakeHoursUnavailable](
		makeHoursUnavailableHandler{hourRepo: hourRepo},
		logger,
	)
}

func (h makeHoursUnavailableHandler) Handle(ctx context.Context, cmd MakeHoursUnavailable) error {
	for _, hourToUpdate := range cmd.Hours {
		if err := h.hourRepo.UpdateHour(ctx, hourToUpdate, func(hr *hour.Hour) (*hour.Hour, error) {
			if err := hr.MakeNotAvailable(); err != nil {
				return nil, err
			}
			return hr, nil
		}); err != nil {
			return v1.ErrorUpdateAvailabilityFail("unable to disbale availability").WithCause(err)
		}
	}
	return nil
}
