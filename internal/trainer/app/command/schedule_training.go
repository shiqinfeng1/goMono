package command

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/shiqinfeng1/goMono/api/trainer/v1"
	"github.com/shiqinfeng1/goMono/internal/common/decorator"
	"github.com/shiqinfeng1/goMono/internal/trainer/domain/hour"
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
	metricsClient decorator.MetricsClient,
) ScheduleTrainingHandler {
	if hourRepo == nil {
		panic("nil hourRepo")
	}

	return decorator.ApplyCommandDecorators[ScheduleTraining](
		scheduleTrainingHandler{hourRepo: hourRepo},
		logger,
		metricsClient,
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
