package decorator

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
)

// 命令的装饰器。最外层是日志，其次是测量，最后才是业务命令
func ApplyCommandDecorators[H any](handler CommandHandler[H], logger log.Logger) CommandHandler[H] {
	return commandLoggingDecorator[H]{
		base: handler,
		logger: log.With(logger,
			"trace.id", tracing.TraceID(),
			"span.id", tracing.SpanID(),
		),
	}
}

// 根据C定义不同的命令接口
type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}

func generateActionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}
