package main

import (
	"context"
	"os"

	"github.com/go-kratos/kratos/v2"
	kcfg "github.com/go-kratos/kratos/v2/config"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/shiqinfeng1/goMono/internal/common/config"
	conf "github.com/shiqinfeng1/goMono/internal/common/config/trainer"
	"github.com/shiqinfeng1/goMono/internal/common/discovery"
	"github.com/shiqinfeng1/goMono/internal/common/log"
	"github.com/shiqinfeng1/goMono/internal/common/types"

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
		// 这里添加需要监听的字段，及处理函数
	}
	config.Bootstrap(
		[]string{"public.yaml", "trainer.yaml"},
		[]interface{}{&pubCfg, &srvCfg},
		onChanges,
	)
}

func newApp(logger klog.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(ID),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(discovery.MustNacosRegistrar()),
	)
}

func main() {
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
		pubCfg.Trace,
		pubCfg.Adapter,
		srvCfg.Http,
		srvCfg.Grpc,
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
