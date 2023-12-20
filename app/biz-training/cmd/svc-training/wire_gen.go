// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/shiqinfeng1/goMono/app/biz-training/internal/adapters"
	"github.com/shiqinfeng1/goMono/app/biz-training/internal/application"
	"github.com/shiqinfeng1/goMono/app/biz-training/internal/conf"
	"github.com/shiqinfeng1/goMono/app/biz-training/internal/ports"
	"github.com/shiqinfeng1/goMono/app/biz-training/internal/service"
	"github.com/shiqinfeng1/goMono/app/common/config"
	"github.com/shiqinfeng1/goMono/app/common/log"
	"github.com/shiqinfeng1/goMono/app/common/registrar"
	"github.com/shiqinfeng1/goMono/app/common/types"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(contextContext context.Context, srvInfo *types.SrvInfo, discovery *config.Discovery, configLog *config.Log, adapter *config.Adapter, http *conf.HTTP, auth *conf.Auth) (*kratos.App, func(), error) {
	logger := log.New(srvInfo, configLog)
	registryRegistrar := registrar.MustNacosRegistrar(discovery)
	repository := adapters.NewTrainingRepo(adapter, logger)
	trainerGrpc := adapters.NewTrainerGrpc(discovery)
	userGrpc := adapters.NewUserGrpc(discovery)
	applicationApplication := application.NewApplication(logger, repository, trainerGrpc, userGrpc)
	httpService := service.NewHttpService(applicationApplication)
	server := ports.NewHTTPServer(http, auth, logger, httpService)
	app := newApp(logger, registryRegistrar, server)
	return app, func() {
	}, nil
}