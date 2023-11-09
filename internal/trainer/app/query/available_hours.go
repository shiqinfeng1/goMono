package query

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/shiqinfeng1/goMono/internal/common/decorator"
	"github.com/shiqinfeng1/goMono/internal/common/errors"
	"github.com/shiqinfeng1/goMono/internal/trainer/adapters"
	"github.com/shiqinfeng1/goMono/internal/trainer/domain/hour"
)

type AvailableHours struct {
	From time.Time
	To   time.Time
}

type AvailableHoursHandler decorator.QueryHandler[AvailableHours, []Date]

type availableHoursHandler struct {
	repo hour.QueryRepository
}

func NewAvailableHoursHandler(
	repo hour.QueryRepository,
	logger log.Logger,
	metricsClient decorator.MetricsClient,
) AvailableHoursHandler {
	return decorator.ApplyQueryDecorators[AvailableHours, []Date](
		availableHoursHandler{repo: repo},
		logger,
		metricsClient,
	)
}

// setDefaultAvailability adds missing hours to Date model if they were not set
func setDefaultAvailability(date Date) Date {

	return date
}

func (h availableHoursHandler) Handle(ctx context.Context, query AvailableHours) (d []Date, err error) {
	var dates []Date
	if query.From.After(query.To) {
		return nil, errors.NewIncorrectInputError("date-from-after-date-to", "Date from after date to")
	}
	date := adapters.DateModel{}
	if err := doc.DataTo(&date); err != nil {
		return nil, err
	}
	dates = append(dates, dateModelToApp(date))

	return h.repo.AvailableHours(ctx, query.From, query.To)
}
