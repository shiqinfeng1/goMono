package cmd

import (
	"os"

	kcfg "github.com/go-kratos/kratos/v2/config"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/shiqinfeng1/goMono/app/biz-training/internal/conf"
	cconf "github.com/shiqinfeng1/goMono/app/common/conf"
	"github.com/shiqinfeng1/goMono/app/common/log"
	"github.com/shiqinfeng1/goMono/app/common/types"
)

// go build -ldflags "-X cmd.Version=x.y.z"
var (
	Name    = "training"    // Name is the name of the compiled software.
	Version string          // Version is the version of the compiled software.
	ID, _   = os.Hostname() // 主机信息
	srvCfg  conf.Server     // 应用配置参数
	pubCfg  cconf.Public
)

var svcInfo = &types.SrvInfo{
	ID:      ID,
	Name:    Name,
	Version: Version,
}

func init() {
	onChanges := cconf.CfgOnChanges{
		"public.log.level": func(key string, value kcfg.Value) {
			_ = key
			lvl, _ := value.String()
			log.SetLevel(lvl) // 动态更新level等级
		},
	}
	scan := []cconf.ScanTarget{
		{
			File:   "public.yaml",
			Field:  "public",
			Target: &pubCfg,
		},
		{
			File:   "training.yaml",
			Field:  "server",
			Target: &srvCfg,
		},
	}
	cconf.Bootstrap(scan, onChanges)
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main <sub-command>",
		Brief: "this is main command, please specify a sub command",
	}
)
