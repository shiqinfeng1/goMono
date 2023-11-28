package client

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
)

func NewOtlptraceExporter(ctx context.Context, endpoint string) (*otlptrace.Exporter, error) {
	c := otlptracegrpc.NewClient(
		otlptracegrpc.WithEndpoint(endpoint),
	)
	exp, err := otlptrace.New(ctx, c)
	if err != nil {
		return nil, err
	}
	return exp, nil
}
