package config

import (
	"log"
	"os"
	"strconv"

	nacos "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/natefinch/lumberjack.v2"

	cfg "github.com/go-kratos/kratos/v2/config"
)

func NewNacosSource(group, cfgFile string) cfg.Source {
	host := os.Getenv("NACOS_HOST")
	port, err := strconv.Atoi(os.Getenv("NACOS_PORT"))
	if host == "" || err != nil || port == 0 {
		log.Fatalln("invalid env of NACOS_HOST or NACOS_PORT")
	}
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(host, uint64(port)),
	}

	cc := constant.ClientConfig{
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogRollingConfig: &lumberjack.Logger{
			MaxAge: 7,
		},
		LogLevel: "debug",
	}

	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	return nacos.NewConfigSource(client, nacos.WithGroup(group), nacos.WithDataID(cfgFile))

}
