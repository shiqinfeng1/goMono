package discovery

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/shiqinfeng1/goMono/app/common/client"
)

func MustNacosDiscovery(endpoint string) registry.Discovery {
	c, err := client.NewNacosNamingClient(endpoint)
	if err != nil {
		panic(err)
	}
	return nacos.New(c)
}
