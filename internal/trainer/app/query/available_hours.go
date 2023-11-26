package query

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/shiqinfeng1/goMono/api/trainer/v1"
	"github.com/shiqinfeng1/goMono/internal/common/decorator"
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
		return Date{}, v1.ErrorIncorrectInput("date from after date to")
	}

	return h.repo.AvailableHours(ctx, query.From, query.To)
}
