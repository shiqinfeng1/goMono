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
	c1 := config.Bootstrap("trainer.yaml")
	if err := c1.Scan(&srvCfg); err != nil {
		panic(err)
	}
	c2 := config.Bootstrap("public.yaml")
	if err := c2.Scan(&pubCfg); err != nil {
		panic(err)
	}

	// 这里添加监听需要动态更新的字段
	if err := c2.Watch("log.level", func(key string, value kcfg.Value) {
		if key == "log.level" {
			lvl, _ := value.String()
			log.SetLevel(lvl)
		}
	}); err != nil {
		panic(err)
	}

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
