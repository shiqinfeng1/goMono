// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/shiqinfeng1/goMono/internal/trainings/adapters"
	"github.com/shiqinfeng1/goMono/internal/trainings/app"
	"github.com/shiqinfeng1/goMono/internal/trainings/conf"
	"github.com/shiqinfeng1/goMono/internal/trainings/ports"
	"github.com/shiqinfeng1/goMono/internal/trainings/service"
	"go.opentelemetry.io/otel/sdk/trace"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(server *conf.Server, tracerProvider *trace.TracerProvider, adapter *conf.Adapter, auth *conf.Auth, logger log.Logger) (*kratos.App, func(), error) {
	repository := adapters.NewTrainingRepo(adapter, logger)
	trainerServiceClient, cleanup, err := adapters.NewTrainerClient(adapter)
	if err != nil {
		return nil, nil, err
	}
	usersServiceClient, cleanup2, err := adapters.NewUsersClient(adapter, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	application := app.NewApplication(logger, repository, trainerServiceClient, usersServiceClient)
	httpService := service.NewHttpService(application)
	httpServer := ports.NewHTTPServer(server, auth, logger, tracerProvider, httpService)
	kratosApp := newApp(logger, httpServer)
	return kratosApp, func() {
		cleanup2()
		cleanup()
	}, nil
}
