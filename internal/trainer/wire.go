//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/shiqinfeng1/goMono/internal/common/config"
	"github.com/shiqinfeng1/goMono/internal/common/config/trainer"
	"github.com/shiqinfeng1/goMono/internal/common/log"
	"github.com/shiqinfeng1/goMono/internal/common/registrar"
	"github.com/shiqinfeng1/goMono/internal/common/trace"
	"github.com/shiqinfeng1/goMono/internal/common/types"
	"github.com/shiqinfeng1/goMono/internal/trainer/adapters"
	"github.com/shiqinfeng1/goMono/internal/trainer/application"
	"github.com/shiqinfeng1/goMono/internal/trainer/ports"
	"github.com/shiqinfeng1/goMono/internal/trainer/service"
)

// wireApp init kratos application.
func wireApp(
	context.Context,
	*types.SrvInfo,
	*config.Discovery,
	*config.Log,
	*config.Trace,
	*config.Adapter,
	*trainer.HTTP,
	*trainer.GRPC,
	*trainer.Auth) (*kratos.App, func(), error) {
	panic(wire.Build(
		log.ProviderSet,
		trace.ProviderSet,
		adapters.ProviderSet,
		application.ProviderSet,
		ports.ProviderSet,
		service.ProviderSet,
		registrar.ProviderSet,
		newApp))
}
