//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package cmd

import (
	"context"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/shiqinfeng1/goMono/app/gomono-bff/internal/adapters"
	"github.com/shiqinfeng1/goMono/app/gomono-bff/internal/application"
	"github.com/shiqinfeng1/goMono/app/gomono-bff/internal/conf"
	"github.com/shiqinfeng1/goMono/app/gomono-bff/internal/ports"
	"github.com/shiqinfeng1/goMono/app/gomono-bff/internal/service"
	cconf "github.com/shiqinfeng1/goMono/utils/conf"
	"github.com/shiqinfeng1/goMono/utils/log"
	"github.com/shiqinfeng1/goMono/utils/registrar"
	"github.com/shiqinfeng1/goMono/utils/types"
)

// wireApp init kratos application.
func wireApp(
	context.Context,
	*types.SrvInfo,
	*cconf.Discovery,
	*cconf.Log,
	*cconf.Trace,
	*conf.HTTP,
	*conf.Auth,
	*conf.Register,
) (*kratos.App, func(), error) {
	panic(wire.Build(
		log.ProviderSet,
		adapters.ProviderSet,
		application.ProviderSet,
		ports.ProviderSet,
		service.ProviderSet,
		registrar.ProviderSet,
		newApp))
}
