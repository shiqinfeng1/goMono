package cmd

import (
	"context"
	"net/url"

	"github.com/go-kratos/kratos/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/shiqinfeng1/goMono/app/biz-training/internal/conf"
	"github.com/shiqinfeng1/goMono/app/common/trace"
)

func newApp(register *conf.Register, logger klog.Logger, regstr registry.Registrar, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID(ID),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(klog.With(logger,
			"layer", "service",
			"trace.id", tracing.TraceID(),
			"span.id", tracing.SpanID(),
		)),
		kratos.Server(
			gs,
		),
		kratos.Registrar(regstr),
		kratos.Endpoint(&url.URL{Scheme: "http", Host: register.Endpoints[0]}), //  指定服务地址，该地址会提交给注册中心，如果不指定，那么将注册容器内部地址，导致外部无法访问
	)
}

var (
	Service = &gcmd.Command{
		Name:        "svc",
		Brief:       "start training grpc service",
		Description: "entry to start a training grpc service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			trace.NewTrace(ctx, pubCfg.Trace)
			app, cleanup, err := wireApp(
				ctx,
				svcInfo,
				pubCfg.Discovery,
				pubCfg.Log,
				pubCfg.Adapter,
				pubCfg.Trace,
				srvCfg.Grpc,
				srvCfg.Register,
			)
			if err != nil {
				panic(err)
			}
			defer cleanup()

			// start and wait for stop signal
			if err := app.Run(); err != nil {
				panic(err)
			}
			return nil
		},
	}
)
