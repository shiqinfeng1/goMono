package main

import (
	"context"
	"os"

	"github.com/go-kratos/kratos/v2"
	kcfg "github.com/go-kratos/kratos/v2/config"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"
	conf "github.com/shiqinfeng1/goMono/app/biz-training/internal/conf"
	"github.com/shiqinfeng1/goMono/app/common/config"
	"github.com/shiqinfeng1/goMono/app/common/log"
	"github.com/shiqinfeng1/goMono/app/common/types"
	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Name    = "trainer"     // Name is the name of the compiled software.
	Version string          // Version is the version of the compiled software.
	ID, _   = os.Hostname() // 主机信息
	srvCfg  conf.Server     // 应用配置参数
	pubCfg  config.Public
)

func init() {
	onChanges := map[string]func(key string, value kcfg.Value){
		"log.level": func(key string, value kcfg.Value) {
			_ = key
			lvl, _ := value.String()
			log.SetLevel(lvl) // 动态更新level等级
		},
	}
	config.Bootstrap(
		[]string{"public.yaml", "training.yaml"},
		[]interface{}{&pubCfg, &srvCfg},
		onChanges,
	)
}

func newApp(logger klog.Logger, regstr registry.Registrar, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(ID),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
		),
		kratos.Registrar(regstr),
	)
}

func main() {

	// 	docker run -d --name jaeger \
	//   -e COLLECTOR_OTLP_ENABLED=true \
	//   -p 16686:16686 \
	//   -p 4317:4317 \
	//   -p 4318:4318 \
	//   jaegertracing/all-in-one:latest
	ctx := context.Background()
	svcInfo := &types.SrvInfo{
		ID:      ID,
		Name:    Name,
		Version: Version,
	}
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
}