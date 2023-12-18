package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kratos/gateway/client"
	"github.com/go-kratos/gateway/config"
	"github.com/go-kratos/gateway/middleware"
	"github.com/go-kratos/gateway/proxy"
	"github.com/go-kratos/gateway/proxy/debug"
	"github.com/go-kratos/gateway/server"
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
	ctrlService string
	proxyAddrs  = newSliceVar(":8080")
	proxyConfig string
	withDebug   bool
)

type sliceVar struct {
	val        []string
	defaultVal []string
}

func newSliceVar(defaultVal ...string) sliceVar {
	return sliceVar{defaultVal: defaultVal}
}
func (s *sliceVar) Get() []string {
	if len(s.val) <= 0 {
		return s.defaultVal
	}
	return s.val
}
func (s *sliceVar) Set(val string) error {
	s.val = append(s.val, val)
	return nil
}
func (s *sliceVar) String() string { return fmt.Sprintf("%+v", *s) }

func init() {
	rand.Seed(uint64(time.Now().Nanosecond()))

	// flag.BoolVar(&withDebug, "debug", false, "enable debug handlers")
	// flag.Var(&proxyAddrs, "addr", "proxy address, eg: -addr 0.0.0.0:8080")
	// flag.StringVar(&proxyConfig, "conf", "config.yaml", "config path, eg: -conf config.yaml")
	// 动态更新配置。key：需要监听的字段；value：配置变化后的处理函数
	onChanges := map[string]func(key string, value kcfg.Value){
		"log.level": func(key string, value kcfg.Value) {
			_ = key
			lvl, _ := value.String()
			log.SetLevel(lvl) // 动态更新level等级
		},
		// todo： 这里添加需要监听的字段，及处理函数
	}
	config.Bootstrap(
		[]string{"public.yaml", "trainer.yaml"}, // 指定要加载的配置文件
		[]interface{}{&pubCfg, &srvCfg},
		onChanges,
	)
}

func main() {
	flag.Parse()
	clientFactory := client.NewFactory(discovery.MustNacosDiscovery())
	p, err := proxy.New(clientFactory, middleware.Create)
	if err != nil {
		log.Fatalf("failed to new proxy: %v", err)
	}
	circuitbreaker.Init(clientFactory)

	ctx := context.Background()
	confLoader, err := config.NewFileLoader(proxyConfig)
	if err != nil {
		log.Fatalf("failed to create config file loader: %v", err)
	}
	defer confLoader.Close()
	bc, err := confLoader.Load(context.Background())
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if err := p.Update(bc); err != nil {
		log.Fatalf("failed to update service config: %v", err)
	}
	reloader := func() error {
		bc, err := confLoader.Load(context.Background())
		if err != nil {
			log.Errorf("failed to load config: %v", err)
			return err
		}
		if err := p.Update(bc); err != nil {
			log.Errorf("failed to update service config: %v", err)
			return err
		}
		log.Infof("config reloaded")
		return nil
	}
	confLoader.Watch(reloader)

	var serverHandler http.Handler = p
	if withDebug {
		debug.Register("proxy", p)
		debug.Register("config", confLoader)
		// if ctrlLoader != nil {
		// 	debug.Register("ctrl", ctrlLoader)
		// }
		serverHandler = debug.MashupWithDebugHandler(p)
	}
	servers := make([]transport.Server, 0, len(proxyAddrs.Get()))
	for _, addr := range proxyAddrs.Get() {
		servers = append(servers, server.NewProxy(serverHandler, addr))
	}
	app := kratos.New(
		kratos.Name(bc.Name),
		kratos.Context(ctx),
		kratos.Server(
			servers...,
		),
	)
	if err := app.Run(); err != nil {
		log.Errorf("failed to run servers: %v", err)
	}
}
