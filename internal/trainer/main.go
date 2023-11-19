package main

import (
	"context"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	cfgsource "github.com/shiqinfeng1/goMono/internal/common/config"
	"github.com/shiqinfeng1/goMono/internal/common/log"
	"github.com/shiqinfeng1/goMono/internal/trainer/conf"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Name    = "trainer"     // Name is the name of the compiled software.
	Version string          // Version is the version of the compiled software.
	ID, _   = os.Hostname() // 主机信息
	BootCfg conf.Bootstrap  // 应用配置参数
	Logger  klog.Logger
)

func init() {
	// 读入配置
	c := cfgsource.Bootstrap("trainer.yaml")
	if err := c.Scan(&BootCfg); err != nil {
		panic(err)
	}
	Logger = log.NewKLogger(ID, Name, Version, BootCfg.Log.Level)

	// 这里添加监听需哟啊动态更新的字段
	if err := c.Watch("log.level", func(key string, value config.Value) {
		if key == "log.level" {
			lvl, _ := value.String()
			log.SetLevel(lvl)
		}
	}); err != nil {
		panic(err)
	}

}

func newApp(logger klog.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(ID),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {

	ctx := context.Background()
	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithEndpoint(BootCfg.Trace.Endpoint),
	)
	exp, err := otlptrace.New(ctx, client)
	if err != nil {
		panic(err)
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(Name),
		)),
	)
	app, cleanup, err := wireApp(BootCfg.Server, tp, BootCfg.Adapter, BootCfg.Auth, Logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
