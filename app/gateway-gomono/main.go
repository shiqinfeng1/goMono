package main

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	gcfg "github.com/go-kratos/gateway/api/gateway/config/v1"
	"github.com/go-kratos/gateway/client"
	"github.com/go-kratos/gateway/middleware"
	"github.com/go-kratos/gateway/proxy"
	"github.com/go-kratos/gateway/proxy/debug"
	"github.com/go-kratos/gateway/server"
	kcfg "github.com/go-kratos/kratos/v2/config"
	cconf "github.com/shiqinfeng1/goMono/app/common/conf"
	"github.com/shiqinfeng1/goMono/app/common/discovery"

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
	gatewayCfg gcfg.Gateway
	p          *proxy.Proxy
)

func init() {
	rand.Seed(uint64(time.Now().Nanosecond()))
	// 动态更新配置。key：需要监听的字段；value：配置变化后的处理函数
	onChanges := map[string]func(key string, value kcfg.Value){
		"middlewares": func(key string, value kcfg.Value) {
			_ = key
			var middl []*gcfg.Middleware
			value.Scan(&middl)
			gatewayCfg.Middlewares = middl
			if err := p.Update(&gatewayCfg); err != nil {
				log.Fatalf("failed to update middleware config: %v", err)
			}
		},
		"endpoints": func(key string, value kcfg.Value) {
			_ = key
			var edps []*gcfg.Endpoint
			value.Scan(&edps)
			gatewayCfg.Endpoints = edps
			if err := p.Update(&gatewayCfg); err != nil {
				log.Fatalf("failed to update middleware config: %v", err)
			}
		},
	}
	cconf.Bootstrap(
		[]string{"gateway.yaml"}, // 指定要加载的配置文件
		[]interface{}{&gatewayCfg},
		onChanges,
	)
}

func main() {
	ctx := context.Background()
	DEBUG := os.Getenv("DEBUG")
	PROXYADDRS := os.Getenv("PROXYADDRS")
	withDebug, _ := strconv.ParseBool(DEBUG)
	proxyAddrs := strings.Split(PROXYADDRS, ",")
	nacos_host := os.Getenv("NACOS_HOST")
	nacos_port := os.Getenv("NACOS_PORT")
	clientFactory := client.NewFactory(discovery.MustNacosDiscovery(nacos_host + ":" + nacos_port))
	var err error
	p, err = proxy.New(clientFactory, middleware.Create)
	if err != nil {
		log.Fatalf("failed to new proxy: %v", err)
	}
	circuitbreaker.Init(clientFactory)

	if err := p.Update(&gatewayCfg); err != nil {
		log.Fatalf("failed to update service config: %v", err)
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
		kratos.Context(ctx),
		kratos.Server(
			servers...,
		),
	)
	if err := app.Run(); err != nil {
		log.Errorf("failed to run servers: %v", err)
	}
}
