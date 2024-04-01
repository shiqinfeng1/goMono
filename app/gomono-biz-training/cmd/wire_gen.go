// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/shiqinfeng1/goMono/app/gomono-biz-training/internal/adapters"
	"github.com/shiqinfeng1/goMono/app/gomono-biz-training/internal/application"
	conf2 "github.com/shiqinfeng1/goMono/app/gomono-biz-training/internal/conf"
	"github.com/shiqinfeng1/goMono/app/gomono-biz-training/internal/ports"
	"github.com/shiqinfeng1/goMono/app/gomono-biz-training/internal/service"
	"github.com/shiqinfeng1/goMono/utils/conf"
	"github.com/shiqinfeng1/goMono/utils/log"
	"github.com/shiqinfeng1/goMono/utils/registrar"
	"github.com/shiqinfeng1/goMono/utils/types"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(contextContext context.Context, srvInfo *types.SrvInfo, discovery *conf.Discovery, confLog *conf.Log, adapter *conf.Adapter, trace *conf.Trace, grpc *conf2.GRPC, register *conf2.Register) (*kratos.App, func(), error) {
	logger := log.New(srvInfo, confLog)
	registryRegistrar := registrar.MustNacosRegistrar(discovery)
	repository := adapters.NewTrainingRepo(adapter, logger)
	trainerGrpc := adapters.NewTrainerGrpc(discovery, logger)
	userGrpc := adapters.NewUserGrpc(discovery, logger)
	applicationApplication := application.NewApplication(logger, repository, trainerGrpc, userGrpc)
	grpcService := service.NewGrpcService(applicationApplication)
	server := ports.NewGRPCServer(contextContext, srvInfo, grpc, logger, trace, grpcService)
	app := newApp(register, logger, registryRegistrar, server)
	return app, func() {
	}, nil
}
