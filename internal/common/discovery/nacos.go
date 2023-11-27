package discovery

import (
	"os"
	"strconv"

	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/natefinch/lumberjack.v2"
)

var cc = constant.ClientConfig{
	TimeoutMs:           5000,
	NotLoadCacheAtStart: true,
	LogDir:              "/tmp/nacos/log",
	CacheDir:            "/tmp/nacos/cache",
	LogRollingConfig: &lumberjack.Logger{
		MaxAge: 7,
	},
	LogLevel: "debug",
}

func MustNacosRegistrar() registry.Registrar {
	return nacosRegistry()
}
func MustNacosDiscovery() registry.Discovery {
	return nacosRegistry()
}
func nacosRegistry() *nacos.Registry {
	host := os.Getenv("NACOS_HOST")
	port, err := strconv.Atoi(os.Getenv("NACOS_PORT"))
	if host == "" || err != nil || port == 0 {
		panic("ErrInvalidAddr")
	}
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(host, uint64(port)),
	}

	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic("ErrCreateClient(err)")
	}
	return nacos.New(client)
}
