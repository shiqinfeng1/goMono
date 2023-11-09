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

type AvailableHoursHandler decorator.QueryHandler[AvailableHours, Date]

type availableHoursHandler struct {
	repo QueryRepository
}

func NewAvailableHoursHandler(
	repo QueryRepository,
	logger log.Logger,
	metricsClient decorator.MetricsClient,
) AvailableHoursHandler {
	return decorator.ApplyQueryDecorators[AvailableHours, Date](
		availableHoursHandler{repo: repo},
		logger,
		metricsClient,
	)
}

// setDefaultAvailability adds missing hours to Date model if they were not set
func setDefaultAvailability(date Date) Date {

	return date
}

type QueryRepository interface {
	AvailableHours(ctx context.Context, from time.Time, to time.Time) (Date, error)
}

func (h availableHoursHandler) Handle(ctx context.Context, query AvailableHours) (d Date, err error) {

	if query.From.After(query.To) {
		return Date{}, errors.NewIncorrectInputError("date-from-after-date-to", "Date from after date to")
	}

	return h.repo.AvailableHours(ctx, query.From, query.To)
}
