package cmd

import (
	kcfg "github.com/go-kratos/kratos/v2/config"
	"github.com/shiqinfeng1/goMono/app/gomono-biz-training/internal/conf"
	cconf "github.com/shiqinfeng1/goMono/utils/conf"
	"github.com/shiqinfeng1/goMono/utils/log"
)

var (
	srvCfg conf.Server // 应用配置参数
	pubCfg cconf.Public
)

func Bootstrap() {
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
