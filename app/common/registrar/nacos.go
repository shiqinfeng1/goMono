package registrar

import (
	"fmt"

	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/shiqinfeng1/goMono/app/common/client"
	"github.com/shiqinfeng1/goMono/app/common/conf"
)

func MustNacosRegistrar(dis *conf.Discovery) registry.Registrar {
	if len(dis.Endpoints) == 0 {
		panic("no such discovery config endpoint")
	}
	c, err := client.NewNacosNamingClient(dis.Endpoints[0])
	if err != nil {
		panic(err)
	}
	srvs, err := c.GetAllServicesInfo(vo.GetAllServiceInfoParam{})
	fmt.Println("all services in nacos:", srvs, err)
	return nacos.New(c)
}
