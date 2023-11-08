package query

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/shiqinfeng1/goMono/internal/common/decorator"
	"github.com/shiqinfeng1/goMono/internal/common/errors"
)

type AvailableHours struct {
	From time.Time
	To   time.Time
}

type AvailableHoursHandler decorator.QueryHandler[AvailableHours, []Date]

type AvailableHoursReadModel interface {
	AvailableHours(ctx context.Context, from time.Time, to time.Time) ([]Date, error)
}

type availableHoursHandler struct {
	readModel AvailableHoursReadModel
}

func NewAvailableHoursHandler(
	readModel AvailableHoursReadModel,
	logger log.Logger,
	metricsClient decorator.MetricsClient,
) AvailableHoursHandler {
	return decorator.ApplyQueryDecorators[AvailableHours, []Date](
		availableHoursHandler{readModel: readModel},
		logger,
		metricsClient,
	)
}

func (h availableHoursHandler) Handle(ctx context.Context, query AvailableHours) (d []Date, err error) {
	if query.From.After(query.To) {
		return nil, errors.NewIncorrectInputError("date-from-after-date-to", "Date from after date to")
	}

	return h.readModel.AvailableHours(ctx, query.From, query.To)
}
