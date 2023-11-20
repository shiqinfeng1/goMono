package trace

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func New(ctx context.Context, svcName, endpoint string) *trace.TracerProvider {
	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithEndpoint(endpoint),
	)
	exp, err := otlptrace.New(ctx, client)
	if err != nil {
		panic(err)
	}
	return trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(svcName),
		)),
	)
}
