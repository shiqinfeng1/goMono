package cmd

import (
	kcfg "github.com/go-kratos/kratos/v2/config"
	"github.com/shiqinfeng1/goMono/app/gomono-bff/internal/conf"
	cconf "github.com/shiqinfeng1/goMono/utils/conf"
	"github.com/shiqinfeng1/goMono/utils/log"
)

var (
	bffCfg conf.Server // 应用配置参数
	pubCfg cconf.Public
)

func Bootstrap() {
	// 动态更新配置
	onChanges := cconf.CfgOnChanges{
		"public.log.level": func(key string, value kcfg.Value) {
			_ = key
			lvl, _ := value.String()
			log.SetLevel(lvl) // 动态更新level等级
		},
		// todo： 这里添加需要监听的字段，及处理函数
	}
	scan := []cconf.ScanTarget{
		{
			File:   "bff.yaml",
			Field:  "public",
			Target: &pubCfg,
		},
		{
			File:   "bff.yaml",
			Field:  "server",
			Target: &bffCfg,
		},
	}
	cconf.Bootstrap(scan, onChanges)
}
