package conf

import (
	nacos "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	"github.com/shiqinfeng1/goMono/utils/client"

	kcfg "github.com/go-kratos/kratos/v2/config"
)

func MustNewNacosSource(group, cfgFile string) kcfg.Source {
	c, err := client.NewNacosConfigClient()
	if err != nil {
		panic(err)
	}
	return nacos.NewConfigSource(c, nacos.WithGroup(group), nacos.WithDataID(cfgFile))
}

// 配置中心的地址需通过环境变量配置，其他配置从配置文件获取
func NewNacosSource(group, cfgFile string) (kcfg.Source, error) {
	c, err := client.NewNacosConfigClient()
	if err != nil {
		return nil, err
	}
	return nacos.NewConfigSource(c, nacos.WithGroup(group), nacos.WithDataID(cfgFile)), nil
}
