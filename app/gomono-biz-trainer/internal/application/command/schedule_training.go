package command

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/shiqinfeng1/goMono/app/gomono-biz-trainer/api/v1"
	"github.com/shiqinfeng1/goMono/app/gomono-biz-trainer/internal/domain/hour"
	"github.com/shiqinfeng1/goMono/utils/decorator"
)

type ScheduleTraining struct {
	Hour time.Time
}

type ScheduleTrainingHandler decorator.CommandHandler[ScheduleTraining]

type scheduleTrainingHandler struct {
	hourRepo hour.CmdRepo
}

func NewScheduleTrainingHandler(
	hourRepo hour.CmdRepo,
	logger log.Logger,
) ScheduleTrainingHandler {
	if hourRepo == nil {
		panic("nil hourRepo")
	}

	return decorator.ApplyCommandDecorators[ScheduleTraining](
		scheduleTrainingHandler{hourRepo: hourRepo},
		logger,
	)
}

func (h scheduleTrainingHandler) Handle(ctx context.Context, cmd ScheduleTraining) error {
	if err := h.hourRepo.UpdateHour(ctx, cmd.Hour, func(hr *hour.Hour) (*hour.Hour, error) {
		if err := hr.ScheduleTraining(); err != nil {
			return nil, err
		}
		return hr, nil
	}); err != nil {
		return v1.ErrorUpdateAvailabilityFail("unable to schedule").WithCause(err)
	}
	return nil
}
