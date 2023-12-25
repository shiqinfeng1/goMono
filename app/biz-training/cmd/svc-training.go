package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Service = &gcmd.Command{
		Name:        "svc",
		Brief:       "start training grpc service",
		Description: "entry to start a training grpc service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			app, cleanup, err := wireApp(
				ctx,
				svcInfo,
				pubCfg.Discovery,
				pubCfg.Log,
				pubCfg.Adapter,
				srvCfg.Http,
				srvCfg.Auth,
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