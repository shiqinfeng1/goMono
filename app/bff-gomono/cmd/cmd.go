package cmd

import (
	"context"
	"os"

	kcfg "github.com/go-kratos/kratos/v2/config"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/shiqinfeng1/goMono/app/bff-gomono/internal/conf"
	cconf "github.com/shiqinfeng1/goMono/app/common/conf"
	"github.com/shiqinfeng1/goMono/app/common/log"
	"github.com/shiqinfeng1/goMono/app/common/types"
)

// go build -ldflags "-X cmd.Version=x.y.z"
var (
	Name    = "bff"         // Name is the name of the compiled software.
	Version string          // Version is the version of the compiled software.
	ID, _   = os.Hostname() // 主机信息
)

var svcInfo = &types.SrvInfo{
	ID:      ID,
	Name:    Name,
	Version: Version,
}
var (
	bffCfg conf.Server  // 应用配置参数
	pubCfg cconf.Public // 应用配置参数
)

func init() {
	// 动态更新配置。key：需要监听的字段；value：配置变化后的处理函数
	onChanges := map[string]func(key string, value kcfg.Value){
		"log.level": func(key string, value kcfg.Value) {
			_ = key
			lvl, _ := value.String()
			log.SetLevel(lvl) // 动态更新level等级
		},
		// todo： 这里添加需要监听的字段，及处理函数
	}
	cconf.Bootstrap(
		[]string{"bff.yaml", "public.yaml"}, // 指定要加载的配置文件
		[]interface{}{&bffCfg, &pubCfg},
		onChanges,
	)
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main <sub-command>",
		Brief: "this is main command, please specify a sub command",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			app, cleanup, err := wireApp(
				ctx,
				svcInfo,
				pubCfg.Discovery,
				pubCfg.Log,
				pubCfg.Adapter,
				bffCfg.Http,
				bffCfg.Auth,
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
