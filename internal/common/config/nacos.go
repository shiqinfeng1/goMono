package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	nacos "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/natefinch/lumberjack.v2"

	cfg "github.com/go-kratos/kratos/v2/config"
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

var (
	ErrInvalidAddr  = errors.New("invalid env of NACOS_HOST or NACOS_PORT")
	ErrCreateClient = func(err error) error { return fmt.Errorf("create nacos client:%w", err) }
)

func MustNewNacosSource(group, cfgFile string) cfg.Source {
	host := os.Getenv("NACOS_HOST")
	port, err := strconv.Atoi(os.Getenv("NACOS_PORT"))
	if host == "" || err != nil || port == 0 {
		panic(ErrInvalidAddr)
	}
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(host, uint64(port)),
	}

	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(ErrCreateClient(err))
	}
	return nacos.NewConfigSource(client, nacos.WithGroup(group), nacos.WithDataID(cfgFile))
}

// 配置中心的地址需通过环境变量配置，其他配置从配置找你想你获取
func NewNacosSource(group, cfgFile string) (cfg.Source, error) {
	host := os.Getenv("NACOS_HOST")
	port, err := strconv.Atoi(os.Getenv("NACOS_PORT"))
	if host == "" || err != nil || port == 0 {
		return nil, ErrInvalidAddr
	}
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(host, uint64(port)),
	}

	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		return nil, ErrCreateClient(err)
	}
	return nacos.NewConfigSource(client, nacos.WithGroup(group), nacos.WithDataID(cfgFile)), nil
}
