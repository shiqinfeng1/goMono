package cmd

import (
	"context"
	"net/url"
	"os"

	"github.com/go-kratos/kratos/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/shiqinfeng1/goMono/app/gomono-bff/internal/conf"
	"github.com/shiqinfeng1/goMono/utils/types"
)

// go build -ldflags "-X cmd.Version=x.y.z"
var (
	Name    = "bff"         // Name is the name of the compiled software.
	Version string          // Version is the version of the compiled software.
	ID, _   = os.Hostname() // 主机信息
	svcInfo = &types.SrvInfo{
		ID:      ID,
		Name:    Name,
		Version: Version,
	}
)

func init() {
	Bootstrap()
}

func newApp(register *conf.Register, logger klog.Logger, regstr registry.Registrar, hs *http.Server) *kratos.App {
	if len(register.Endpoints) == 0 {
		panic("bff register.Endpoints is nil")
	}
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
			hs,
		),
		kratos.Registrar(regstr),
		kratos.Endpoint(&url.URL{Scheme: "http", Host: register.Endpoints[0]}), //  指定服务地址，该地址会提交给注册中心，如果不指定，那么将注册容器内部地址，导致外部无法访问本服务
	)
}

var Main = gcmd.Command{
	Name:  "main",
	Usage: "main <sub-command>",
	Brief: "this is main command, please specify a sub command",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		app, cleanup, err := wireApp(
			ctx,
			svcInfo,
			pubCfg.Discovery,
			pubCfg.Log,
			pubCfg.Trace,
			bffCfg.Http,
			bffCfg.Auth,
			bffCfg.Register,
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
