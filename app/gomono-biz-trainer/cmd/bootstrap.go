package cmd

import (
	kcfg "github.com/go-kratos/kratos/v2/config"
	"github.com/shiqinfeng1/goMono/app/gomono-biz-trainer/internal/conf"
	cconf "github.com/shiqinfeng1/goMono/utils/conf"
	"github.com/shiqinfeng1/goMono/utils/log"
)

var (
	srvCfg conf.Server // 应用配置参数
	pubCfg cconf.Public
)

func Bootstrap() {
	// 动态更新配置。key：需要监听的字段；value：配置变化后的处理函数
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
			File:   "trainer.yaml",
			Field:  "public",
			Target: &pubCfg,
		},
		{
			File:   "trainer.yaml",
			Field:  "server",
			Target: &srvCfg,
		},
	}
	cconf.Bootstrap(scan, onChanges)
}
