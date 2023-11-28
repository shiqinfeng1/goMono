package registrar

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/shiqinfeng1/goMono/internal/common/client"
)

func MustNacosRegistrar(endpoint string) registry.Registrar {
	c, err := client.NewNacosNamingClient(endpoint)
	if err != nil {
		panic(err)
	}
	return nacos.New(c)
}
