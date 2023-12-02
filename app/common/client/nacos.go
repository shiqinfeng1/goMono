package client

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
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
var (
	ErrInvalidAddr  = errors.New("invalid env of NACOS_HOST or NACOS_PORT")
	ErrCreateClient = func(err error) error { return fmt.Errorf("create nacos client:%w", err) }
)

func NewNacosConfigClient() (config_client.IConfigClient, error) {
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
	return client, nil
}

func NewNacosNamingClient(endpoint string) (naming_client.INamingClient, error) {
	addr := strings.Split(endpoint, ":")
	if len(addr) != 2 {
		return nil, ErrInvalidAddr
	}
	host := addr[0]
	port, err := strconv.Atoi(addr[1])
	if err != nil {
		return nil, ErrCreateClient(err)
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
		return nil, ErrCreateClient(err)
	}
	return client, nil
}
