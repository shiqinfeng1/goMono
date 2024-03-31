package registrar

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/shiqinfeng1/goMono/utils/client"
	"github.com/shiqinfeng1/goMono/utils/conf"
)

func MustNacosRegistrar(dis *conf.Discovery) registry.Registrar {
	if len(dis.Endpoints) == 0 {
		panic("no such discovery config endpoint")
	}
	c, err := client.NewNacosNamingClient(dis.Endpoints[0])
	if err != nil {
		panic(err)
	}
	return nacos.New(c)
}
