package cmd

import (
	"context"
	"net/url"

	"github.com/go-kratos/kratos/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/shiqinfeng1/goMono/app/biz-trainer/internal/conf"
)

func newApp(register *conf.Register, logger klog.Logger, regstr registry.Registrar, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID(ID),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
		// kratos.Registrar(regstr),
		kratos.Endpoint(&url.URL{Scheme: "http", Host: register.Endpoints[0]}),
	)
}

var (
	Service = &gcmd.Command{
		Name:        "svc",
		Brief:       "start trainer grpc service",
		Description: "entry to start a trainer grpc service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			app, cleanup, err := wireApp(
				ctx,
				svcInfo,
				pubCfg.Discovery,
				pubCfg.Log,
				pubCfg.Adapter,
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
