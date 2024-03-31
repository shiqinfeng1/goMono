package discovery

import (
	"errors"

	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/shiqinfeng1/goMono/utils/client"
)

var ErrServiceKindInvaild = errors.New("service kind is invalid")

func MustNacosDiscovery(endpoint string, kind string) registry.Discovery {
	c, err := client.NewNacosNamingClient(endpoint)
	if err != nil {
		panic(err)
	}
	if kind != "grpc" && kind != "http" {
		panic(ErrServiceKindInvaild)
	}
	return nacos.New(c, nacos.WithDefaultKind(kind))
}
