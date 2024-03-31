package trace

import (
	"context"

	"github.com/shiqinfeng1/goMono/utils/conf"
	"github.com/shiqinfeng1/goMono/utils/types"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func New(ctx context.Context, srvInfo *types.SrvInfo, url *conf.Trace) error {
	// 创建 Jaeger exporter
	exp, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint(url.Endpoint), otlptracehttp.WithInsecure())
	if err != nil {
		return err
	}
	tp := tracesdk.NewTracerProvider(
		// 将基于父span的采样率设置为100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// 始终确保在生产中批量处理
		tracesdk.WithBatcher(exp),
		// 在资源中记录有关此应用程序的信息
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(srvInfo.Name),
			semconv.ServiceVersionKey.String(srvInfo.Version),
			attribute.String("exporter", "otlptrace-http"),
			attribute.String("service.id(hostname)", srvInfo.ID),
		)),
	)
	otel.SetTracerProvider(tp)
	return nil
}
