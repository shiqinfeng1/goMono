package hour

import (
	"context"
	"time"
)

type CmdRepository interface {
	GetHour(ctx context.Context, hourTime time.Time) (*Hour, error)
	UpdateHour(
		ctx context.Context,
		hourTime time.Time,
		updateFn func(h *Hour) (*Hour, error),
	) error
}

type QueryRepository interface {
	AvailableHours(ctx context.Context, from time.Time, to time.Time) ([]Date, error)
}
