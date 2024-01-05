// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/shiqinfeng1/goMono/app/biz-trainer/internal/adapters"
	"github.com/shiqinfeng1/goMono/app/biz-trainer/internal/application"
	conf2 "github.com/shiqinfeng1/goMono/app/biz-trainer/internal/conf"
	"github.com/shiqinfeng1/goMono/app/biz-trainer/internal/ports"
	"github.com/shiqinfeng1/goMono/app/biz-trainer/internal/service"
	"github.com/shiqinfeng1/goMono/app/common/conf"
	"github.com/shiqinfeng1/goMono/app/common/log"
	"github.com/shiqinfeng1/goMono/app/common/registrar"
	"github.com/shiqinfeng1/goMono/app/common/types"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(contextContext context.Context, srvInfo *types.SrvInfo, discovery *conf.Discovery, confLog *conf.Log, adapter *conf.Adapter, grpc *conf2.GRPC, register *conf2.Register) (*kratos.App, func(), error) {
	logger := log.New(srvInfo, confLog)
	registryRegistrar := registrar.MustNacosRegistrar(discovery)
	cmdRepo := adapters.NewHourRepo(adapter, logger)
	queryRepository := adapters.NewDatesMysqlRepo(adapter, logger)
	applicationApplication := application.NewApplication(logger, cmdRepo, queryRepository)
	grpcService := service.NewGrpcService(applicationApplication)
	server := ports.NewGRPCServer(grpc, logger, grpcService)
	app := newApp(register, logger, registryRegistrar, server)
	return app, func() {
	}, nil
}
