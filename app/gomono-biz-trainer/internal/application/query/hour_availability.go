package query

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/shiqinfeng1/goMono/app/gomono-biz-trainer/api/v1"
	"github.com/shiqinfeng1/goMono/app/gomono-biz-trainer/internal/domain/hour"
	"github.com/shiqinfeng1/goMono/utils/decorator"
)

type HourAvailability struct {
	Hour time.Time
}

type HourAvailabilityHandler decorator.QueryHandler[HourAvailability, bool]

type hourAvailabilityHandler struct {
	hourRepo hour.CmdRepo
}

func NewHourAvailabilityHandler(
	hourRepo hour.CmdRepo,
	logger log.Logger,
) HourAvailabilityHandler {
	if hourRepo == nil {
		panic("nil hourRepo")
	}

	return decorator.ApplyQueryDecorators[HourAvailability, bool](
		hourAvailabilityHandler{hourRepo: hourRepo},
		logger,
	)
}

func (h hourAvailabilityHandler) Handle(ctx context.Context, query HourAvailability) (bool, error) {
	hour, err := h.hourRepo.GetHour(ctx, query.Hour)
	if err != nil {
		return false, v1.ErrorQueryFail("query availability fail by %v", query.Hour).WithCause(err)
	}

	return hour.IsAvailable(), nil
}
