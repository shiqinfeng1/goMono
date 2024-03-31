package cmd

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/gateway/client"
	"github.com/go-kratos/gateway/middleware"
	"github.com/go-kratos/gateway/proxy"
	"github.com/go-kratos/gateway/proxy/debug"
	"github.com/go-kratos/gateway/server"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/shiqinfeng1/goMono/utils/discovery"
	"github.com/shiqinfeng1/goMono/utils/types"

	_ "net/http/pprof"

	_ "github.com/go-kratos/gateway/middleware/bbr"
	"github.com/go-kratos/gateway/middleware/circuitbreaker"
	_ "github.com/go-kratos/gateway/middleware/cors"
	_ "github.com/go-kratos/gateway/middleware/logging"
	_ "github.com/go-kratos/gateway/middleware/rewrite"
	_ "github.com/go-kratos/gateway/middleware/tracing"
	_ "github.com/go-kratos/gateway/middleware/transcoder"
	_ "go.uber.org/automaxprocs"

	"github.com/go-kratos/kratos/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/shiqinfeng1/goMono/utils/log"
	"golang.org/x/exp/rand"
)

// go build -ldflags "-X cmd.Version=x.y.z"
var (
	ID, _ = os.Hostname() // 主机信息
	p     *proxy.Proxy
)

func init() {
	rand.Seed(uint64(time.Now().Nanosecond()))
	Bootstrap()
}

var Main = gcmd.Command{
	Name:  "main",
	Usage: "main <sub-command>",
	Brief: "this is main command, please specify a sub command",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		// 获取环境变量配置
		withDebug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
		proxyAddrs := strings.Split(os.Getenv("PROXYADDRS"), ",") // 代理地址
		// 日志接口
		logger := log.New(&types.SrvInfo{
			ID:      ID,
			Name:    gatewayCfg.Name,
			Version: gatewayCfg.Version,
		}, pubCfg.Log)

		klog.SetLogger(logger)
		// 实例化一个main函数使用的log
		l := klog.NewHelper(klog.With(logger, "scope", "main"))
		clientFactory := client.NewFactory(discovery.MustNacosDiscovery(pubCfg.Discovery.Endpoints[0], "http"))

		p, err = proxy.New(clientFactory, middleware.Create)
		if err != nil {
			l.Fatalf("failed to new proxy: %v", err)
		}
		circuitbreaker.Init(clientFactory)

		if err := p.Update(&gatewayCfg); err != nil {
			l.Fatalf("failed to update service config: %v", err)
		}
		var serverHandler http.Handler = p
		if withDebug {
			debug.Register("proxy", p)
			serverHandler = debug.MashupWithDebugHandler(p)
		}
		servers := make([]transport.Server, 0, len(proxyAddrs))
		for _, addr := range proxyAddrs {
			servers = append(servers, server.NewProxy(serverHandler, addr))
		}
		app := kratos.New(
			kratos.Name(gatewayCfg.Name),
			kratos.Version(gatewayCfg.Version),
			kratos.Context(ctx),
			kratos.Server(
				servers...,
			),
			kratos.Metadata(map[string]string{}),
			kratos.Logger(logger),
		)
		if err := app.Run(); err != nil {
			l.Errorf("failed to run servers: %v", err)
		}
		return nil
	},
}
