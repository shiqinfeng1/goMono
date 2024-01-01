// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	log2 "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
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
func wireApp(contextContext context.Context, srvInfo *types.SrvInfo, discovery *conf.Discovery, confLog *conf.Log, adapter *conf.Adapter, grpc *conf2.GRPC, auth *conf2.Auth) (*kratos.App, func(), error) {
	logger := log.New(srvInfo, confLog)
	registryRegistrar := registrar.MustNacosRegistrar(discovery)
	cmdRepo := adapters.NewHourRepo(adapter, logger)
	queryRepository := adapters.NewDatesMysqlRepo(adapter, logger)
	applicationApplication := application.NewApplication(logger, cmdRepo, queryRepository)
	grpcService := service.NewGrpcService(applicationApplication)
	server := ports.NewGRPCServer(grpc, grpcService)
	app := newApp(logger, registryRegistrar, server)
	return app, func() {
	}, nil
}

// wire.go:

func newApp(logger log2.Logger, regstr registry.Registrar, gs *grpc.Server) *kratos.App {
	return kratos.New(kratos.ID(ID), kratos.Name(Name), kratos.Version(Version), kratos.Metadata(map[string]string{}), kratos.Logger(logger), kratos.Server(
		gs,
	), kratos.Registrar(regstr),
	)
}
