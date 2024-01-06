package application

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/shiqinfeng1/goMono/app/bff-gomono/internal/adapters"
)

type Application struct {
	logger      log.Logger
	trainerGrpc *adapters.TrainerGrpc
	userGrpc    *adapters.UserGrpc
}

func NewApplication(logger log.Logger, trainerGrpc *adapters.TrainerGrpc, userGrpc *adapters.UserGrpc) Application {
	return Application{
		logger: log.With(logger,
			"layer", "app",
			"trace.id", tracing.TraceID(),
			"span.id", tracing.SpanID(),
		),
		trainerGrpc: trainerGrpc,
		userGrpc:    userGrpc,
	}
}
