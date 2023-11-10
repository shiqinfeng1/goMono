// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/shiqinfeng1/goMono/internal/trainer/adapters"
	"github.com/shiqinfeng1/goMono/internal/trainer/app"
	"github.com/shiqinfeng1/goMono/internal/trainer/conf"
	"github.com/shiqinfeng1/goMono/internal/trainer/ports"
	"github.com/shiqinfeng1/goMono/internal/trainer/service"
	"go.opentelemetry.io/otel/sdk/trace"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(server *conf.Server, tracerProvider *trace.TracerProvider, adapter *conf.Adapter, auth *conf.Auth, logger log.Logger) (*kratos.App, func(), error) {
	cmdRepository := adapters.NewHourRepo(adapter, logger)
	queryRepository := adapters.NewDatesMysqlRepo(adapter, logger)
	application := app.NewApplication(logger, cmdRepository, queryRepository)
	grpcService := service.NewGrpcService(application)
	grpcServer := ports.NewGRPCServer(server, grpcService)
	httpService := service.NewHttpService(application)
	httpServer := ports.NewHTTPServer(server, auth, logger, tracerProvider, httpService)
	kratosApp := newApp(logger, grpcServer, httpServer)
	return kratosApp, func() {
	}, nil
}
