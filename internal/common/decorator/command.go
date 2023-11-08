package decorator

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
)

func ApplyCommandDecorators[H any](handler CommandHandler[H], logger log.Logger, metricsClient MetricsClient) CommandHandler[H] {
	return commandLoggingDecorator[H]{
		base: commandMetricsDecorator[H]{
			base:   handler,
			client: metricsClient,
		},
		logger: logger,
	}
}

type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}

func generateActionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}
