package application

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/shiqinfeng1/goMono/app/bff-gomono/internal/adapters"
)

type Application struct {
	logger      log.Logger
	trainerGrpc *adapters.TrainerGrpc
	userGrpc    *adapters.UserGrpc
}

func NewApplication(logger log.Logger, trainerGrpc *adapters.TrainerGrpc, userGrpc *adapters.UserGrpc) Application {
	return Application{
		logger:      logger,
		trainerGrpc: trainerGrpc,
		userGrpc:    userGrpc,
	}
}
