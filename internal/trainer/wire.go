//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/shiqinfeng1/goMono/internal/common/grpc"
	"github.com/shiqinfeng1/goMono/internal/common/log"
	"github.com/shiqinfeng1/goMono/internal/common/trace"
	"github.com/shiqinfeng1/goMono/internal/trainer/adapters"
	"github.com/shiqinfeng1/goMono/internal/trainer/app"
	"github.com/shiqinfeng1/goMono/internal/trainer/conf"
	"github.com/shiqinfeng1/goMono/internal/trainer/ports"
	"github.com/shiqinfeng1/goMono/internal/trainer/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(svcID,svcName, svcVer string,*conf.Server, *conf.Adapter, *conf.Auth, *conf.Log, *conf.Trace) (*kratos.App, func(), error) {
	panic(wire.Build(grpc.ProviderSet, log.ProviderSet, trace.ProviderSet, app.ProviderSet, ports.ProviderSet, adapters.ProviderSet, service.ProviderSet, newApp))
}
