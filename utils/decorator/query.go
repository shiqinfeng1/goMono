package decorator

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
)

func ApplyQueryDecorators[H any, R any](handler QueryHandler[H, R], logger log.Logger) QueryHandler[H, R] {
	return queryLoggingDecorator[H, R]{
		base: handler,
		logger: log.With(logger,
			"trace.id", tracing.TraceID(),
			"span.id", tracing.SpanID(),
		),
	}
}

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, q Q) (R, error)
}
