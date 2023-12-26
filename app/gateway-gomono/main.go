package main

import (
	"context"
	"flag"
	"net/http"
	"time"

	"github.com/go-kratos/gateway/client"
	"github.com/go-kratos/gateway/middleware"
	"github.com/go-kratos/gateway/proxy"
	"github.com/go-kratos/gateway/proxy/debug"
	"github.com/go-kratos/gateway/server"
	kcfg "github.com/go-kratos/kratos/v2/config"
	pubcfg "github.com/shiqinfeng1/goMono/app/common/conf"
	"github.com/shiqinfeng1/goMono/app/common/discovery"
	"github.com/shiqinfeng1/goMono/app/gateway-gomono/internal/conf"

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
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"golang.org/x/exp/rand"
)

var (
	gatewayCfg conf.Proxy
)

func init() {
	rand.Seed(uint64(time.Now().Nanosecond()))
	// 动态更新配置。key：需要监听的字段；value：配置变化后的处理函数
	onChanges := map[string]func(key string, value kcfg.Value){
		// todo： 这里添加需要监听的字段，及处理函数
	}

	pubcfg.Bootstrap(
		[]string{"gateway.yaml"}, // 指定要加载的配置文件
		[]interface{}{&gatewayCfg},
		onChanges,
	)
}

func main() {
	ctx := context.Background()
	flag.Parse()
	clientFactory := client.NewFactory(discovery.MustNacosDiscovery(""))
	p, err := proxy.New(clientFactory, middleware.Create)
	if err != nil {
		log.Fatalf("failed to new proxy: %v", err)
	}
	circuitbreaker.Init(clientFactory)

	var serverHandler http.Handler = p
	if gatewayCfg.Debug {
		debug.Register("proxy", p)
		// if ctrlLoader != nil {
		// 	debug.Register("ctrl", ctrlLoader)
		// }
		serverHandler = debug.MashupWithDebugHandler(p)
	}
	servers := make([]transport.Server, 0, len(gatewayCfg.Addrs))
	for _, addr := range gatewayCfg.Addrs {
		servers = append(servers, server.NewProxy(serverHandler, addr))
	}
	app := kratos.New(
		kratos.Name("gateway-gomono"),
		kratos.Context(ctx),
		kratos.Server(
			servers...,
		),
	)
	if err := app.Run(); err != nil {
		log.Errorf("failed to run servers: %v", err)
	}
}
