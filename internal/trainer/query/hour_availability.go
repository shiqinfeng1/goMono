package query

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shiqinfeng1/goMono/internal/common/decorator"
	"github.com/shiqinfeng1/goMono/internal/trainer/domain/hour"
)

type HourAvailability struct {
	Hour time.Time
}

type HourAvailabilityHandler decorator.QueryHandler[HourAvailability, bool]

type hourAvailabilityHandler struct {
	hourRepo hour.Repository
}

func NewHourAvailabilityHandler(
	hourRepo hour.Repository,
	logger log.Logger,
	metricsClient decorator.MetricsClient,
) HourAvailabilityHandler {
	if hourRepo == nil {
		panic("nil hourRepo")
	}

	return decorator.ApplyQueryDecorators[HourAvailability, bool](
		hourAvailabilityHandler{hourRepo: hourRepo},
		logger,
		metricsClient,
	)
}

func (h hourAvailabilityHandler) Handle(ctx context.Context, query HourAvailability) (bool, error) {
	hour, err := h.hourRepo.GetHour(ctx, query.Hour)
	if err != nil {
		return false, err
	}

	return hour.IsAvailable(), nil
}
