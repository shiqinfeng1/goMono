//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/shiqinfeng1/goMono/internal/trainings/adapters"
	"github.com/shiqinfeng1/goMono/internal/trainings/app"
	"github.com/shiqinfeng1/goMono/internal/trainings/conf"
	"github.com/shiqinfeng1/goMono/internal/trainings/ports"
	"github.com/shiqinfeng1/goMono/internal/trainings/service"
	"go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *trace.TracerProvider, *conf.Adapter, *conf.Auth, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(app.ProviderSet, ports.ProviderSet, adapters.ProviderSet, service.ProviderSet, newApp))
}
