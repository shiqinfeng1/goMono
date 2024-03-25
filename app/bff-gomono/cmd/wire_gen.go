// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/shiqinfeng1/goMono/app/bff-gomono/internal/adapters"
	"github.com/shiqinfeng1/goMono/app/bff-gomono/internal/application"
	conf2 "github.com/shiqinfeng1/goMono/app/bff-gomono/internal/conf"
	"github.com/shiqinfeng1/goMono/app/bff-gomono/internal/ports"
	"github.com/shiqinfeng1/goMono/app/bff-gomono/internal/service"
	"github.com/shiqinfeng1/goMono/app/common/conf"
	"github.com/shiqinfeng1/goMono/app/common/log"
	"github.com/shiqinfeng1/goMono/app/common/registrar"
	"github.com/shiqinfeng1/goMono/app/common/types"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(contextContext context.Context, srvInfo *types.SrvInfo, discovery *conf.Discovery, confLog *conf.Log, adapter *conf.Adapter, trace *conf.Trace, http *conf2.HTTP, auth *conf2.Auth, register *conf2.Register) (*kratos.App, func(), error) {
	logger := log.New(srvInfo, confLog)
	registryRegistrar := registrar.MustNacosRegistrar(discovery)
	trainerGrpc := adapters.NewTrainerGrpc(discovery, logger)
	userGrpc := adapters.NewUserGrpc(discovery, logger)
	applicationApplication := application.NewApplication(logger, trainerGrpc, userGrpc)
	httpService := service.NewHttpService(applicationApplication)
	server := ports.NewHTTPServer(contextContext, srvInfo, http, auth, trace, logger, httpService)
	app := newApp(register, logger, registryRegistrar, server)
	return app, func() {
	}, nil
}
