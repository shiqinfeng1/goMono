package cmd

import (
	gcfg "github.com/go-kratos/gateway/api/gateway/config/v1"
	kcfg "github.com/go-kratos/kratos/v2/config"
	klog "github.com/go-kratos/kratos/v2/log"
	cconf "github.com/shiqinfeng1/goMono/utils/conf"
	"github.com/shiqinfeng1/goMono/utils/log"
)

var (
	pubCfg     cconf.Public
	gatewayCfg gcfg.Gateway
)

func Bootstrap() {
	// 动态更新配置。key：需要监听的字段；value：配置变化后的处理函数
	onChanges := map[string]func(key string, value kcfg.Value){
		"public.log.level": func(key string, value kcfg.Value) {
			_ = key
			lvl, _ := value.String()
			log.SetLevel(lvl) // 动态更新level等级
		},
		"middlewares": func(key string, value kcfg.Value) {
			_ = key
			var middl []*gcfg.Middleware
			value.Scan(&middl)
			gatewayCfg.Middlewares = middl
			if err := p.Update(&gatewayCfg); err != nil {
				klog.Fatalf("failed to update middleware config: %v", err)
			}
		},
		"endpoints": func(key string, value kcfg.Value) {
			_ = key
			var edps []*gcfg.Endpoint
			value.Scan(&edps)
			gatewayCfg.Endpoints = edps
			if err := p.Update(&gatewayCfg); err != nil {
				klog.Fatalf("failed to update middleware config: %v", err)
			}
		},
	}
	scan := []cconf.ScanTarget{
		{
			File:   "gateway.yaml",
			Field:  "",
			Target: &gatewayCfg,
		},
		{
			File:   "gateway.yaml",
			Field:  "public",
			Target: &pubCfg,
		},
	}
	cconf.Bootstrap(scan, onChanges)
}
