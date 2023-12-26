package cmd

import (
	"os"

	kcfg "github.com/go-kratos/kratos/v2/config"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/shiqinfeng1/goMono/app/biz-trainer/internal/conf"
	pubcfg "github.com/shiqinfeng1/goMono/app/common/conf"
	"github.com/shiqinfeng1/goMono/app/common/log"
	"github.com/shiqinfeng1/goMono/app/common/types"
)

// go build -ldflags "-X cmd.Version=x.y.z"
var (
	Name    = "trainer"     // Name is the name of the compiled software.
	Version string          // Version is the version of the compiled software.
	ID, _   = os.Hostname() // 主机信息
)

var svcInfo = &types.SrvInfo{
	ID:      ID,
	Name:    Name,
	Version: Version,
}
var (
	srvCfg conf.Server // 应用配置参数
	pubCfg pubcfg.Public
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
	pubcfg.Bootstrap(
		[]string{"public.yaml", "trainer.yaml"}, // 指定要加载的配置文件
		[]interface{}{&pubCfg, &srvCfg},
		onChanges,
	)
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main <sub-command>",
		Brief: "this is main command, please specify a sub command",
	}
)
