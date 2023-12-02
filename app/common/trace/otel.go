package trace

import (
	"context"

	"github.com/shiqinfeng1/goMono/app/common/client"
	"github.com/shiqinfeng1/goMono/app/common/config"
	"github.com/shiqinfeng1/goMono/app/common/types"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func New(ctx context.Context, svcInfo *types.SrvInfo, tr *config.Trace) *trace.TracerProvider {
	exp, err := client.NewOtlptraceExporter(ctx, tr.Endpoint)
	if err != nil {
		return nil
	}
	return trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(svcInfo.Name),
		)),
	)
}
