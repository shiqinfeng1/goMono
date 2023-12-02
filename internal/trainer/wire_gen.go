// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/shiqinfeng1/goMono/internal/common/config"
	"github.com/shiqinfeng1/goMono/internal/common/config/trainer"
	"github.com/shiqinfeng1/goMono/internal/common/log"
	"github.com/shiqinfeng1/goMono/internal/common/registrar"
	"github.com/shiqinfeng1/goMono/internal/common/trace"
	"github.com/shiqinfeng1/goMono/internal/common/types"
	"github.com/shiqinfeng1/goMono/internal/trainer/adapters"
	"github.com/shiqinfeng1/goMono/internal/trainer/application"
	"github.com/shiqinfeng1/goMono/internal/trainer/ports"
	"github.com/shiqinfeng1/goMono/internal/trainer/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(contextContext context.Context, srvInfo *types.SrvInfo, discovery *config.Discovery, configLog *config.Log, configTrace *config.Trace, adapter *config.Adapter, http *trainer.HTTP, grpc *trainer.GRPC, auth *trainer.Auth) (*kratos.App, func(), error) {
	logger := log.New(srvInfo, configLog)
	registryRegistrar := registrar.MustNacosRegistrar(discovery)
	cmdRepo := adapters.NewHourRepo(adapter, logger)
	queryRepository := adapters.NewDatesMysqlRepo(adapter, logger)
	applicationApplication := application.NewApplication(logger, cmdRepo, queryRepository)
	grpcService := service.NewGrpcService(applicationApplication)
	server := ports.NewGRPCServer(grpc, grpcService)
	tracerProvider := trace.New(contextContext, srvInfo, configTrace)
	httpService := service.NewHttpService(applicationApplication)
	httpServer := ports.NewHTTPServer(http, auth, logger, tracerProvider, httpService)
	app := newApp(logger, registryRegistrar, server, httpServer)
	return app, func() {
	}, nil
}
