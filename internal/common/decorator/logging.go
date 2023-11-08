package decorator

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type commandLoggingDecorator[C any] struct {
	base   CommandHandler[C]
	logger log.Logger
}

func (d commandLoggingDecorator[C]) Handle(ctx context.Context, cmd C) (err error) {
	handlerType := generateActionName(cmd)
	logger := log.NewHelper(log.With(d.logger, "command", handlerType, "command_body", fmt.Sprintf("%#v", cmd)))
	logger.Debug("Executing command")
	defer func() {
		if err == nil {
			logger.Info("Command executed successfully")
		} else {
			logger.Errorf("Failed to execute command: %v", err)
		}
	}()

	return d.base.Handle(ctx, cmd)
}

type queryLoggingDecorator[C any, R any] struct {
	base   QueryHandler[C, R]
	logger log.Logger
}

func (d queryLoggingDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err error) {
	handlerType := generateActionName(cmd)
	logger := log.NewHelper(log.With(d.logger, "query", handlerType, "query_body", fmt.Sprintf("%#v", cmd)))
	logger.Debug("Executing query")
	defer func() {
		if err == nil {
			logger.Info("Query executed successfully")
		} else {
			logger.Errorf("Failed to execute query:%v", err)
		}
	}()

	return d.base.Handle(ctx, cmd)
}
