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
	"github.com/shiqinfeng1/goMono/app/common/types"

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
	"github.com/shiqinfeng1/goMono/app/common/log"
	"golang.org/x/exp/rand"
)

var (
	ID, _      = os.Hostname() // 主机信息
	gatewayCfg gcfg.Gateway
	pubCfg     cconf.Public
	p          *proxy.Proxy
)

func init() {
	rand.Seed(uint64(time.Now().Nanosecond()))
	// 动态更新配置。key：需要监听的字段；value：配置变化后的处理函数
	onChanges := map[string]func(key string, value kcfg.Value){
		"public.log.level": func(key string, value kcfg.Value) {
			_ = key
			lvl, _ := value.String()
			log.SetLevel(lvl) // 动态更新level等级
		},
		"middlewares": func(key string, value kcfg.Value) {
			_ = key
			var middl []*gcfg.Middleware
			value.Scan(&middl)
			gatewayCfg.Middlewares = middl
			if err := p.Update(&gatewayCfg); err != nil {
				klog.Fatalf("failed to update middleware config: %v", err)
			}
		},
		"endpoints": func(key string, value kcfg.Value) {
			_ = key
			var edps []*gcfg.Endpoint
			value.Scan(&edps)
			gatewayCfg.Endpoints = edps
			if err := p.Update(&gatewayCfg); err != nil {
				klog.Fatalf("failed to update middleware config: %v", err)
			}
		},
	}
	scan := []cconf.ScanTarget{
		{
			File:   "gateway.yaml",
			Field:  "",
			Target: &gatewayCfg,
		},
		{
			File:   "public.yaml",
			Field:  "public",
			Target: &pubCfg,
		},
	}
	cconf.Bootstrap(scan, onChanges)
}

func main() {
	ctx := context.Background()
	// 获取环境变量配置
	DEBUG := os.Getenv("DEBUG")
	PROXYADDRS := os.Getenv("PROXYADDRS")
	withDebug, _ := strconv.ParseBool(DEBUG)
	proxyAddrs := strings.Split(PROXYADDRS, ",")
	// 日志接口
	logger := log.New(&types.SrvInfo{
		ID:      ID,
		Name:    gatewayCfg.Name,
		Version: gatewayCfg.Version,
	}, pubCfg.Log)

	klog.SetLogger(logger)
	// 实例化一个main函数使用的log
	l := klog.NewHelper(klog.With(logger, "scope", "main"))
	// util.SetFilterNetNumberAndMask("192.168.68.0/24")
	clientFactory := client.NewFactory(discovery.MustNacosDiscovery(pubCfg.Discovery.Endpoints[0], "http"))
	var err error
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
		// kratos.Endpoint(&url.URL{Scheme: "http", Host: pubCfg.GatewayRegister.Endpoints[0]}), //  指定服务地址，该地址会提交给注册中心，如果不指定，那么将注册容器内部地址，导致外部无法访问
	)
	if err := app.Run(); err != nil {
		l.Errorf("failed to run servers: %v", err)
	}
}
