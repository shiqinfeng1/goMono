//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package cmd

import (
	"context"

	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
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

func newApp(logger klog.Logger, regstr registry.Registrar, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(ID),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
		),
		kratos.Registrar(regstr),
	)
}

// wireApp init kratos application.
func wireApp(
	context.Context,
	*types.SrvInfo,
	*config.Discovery,
	*config.Log,
	*config.Adapter,
	*conf.HTTP,
	*conf.Auth) (*kratos.App, func(), error) {
	panic(wire.Build(
		log.ProviderSet,
		adapters.ProviderSet,
		application.ProviderSet,
		ports.ProviderSet,
		service.ProviderSet,
		registrar.ProviderSet,
		newApp))
}