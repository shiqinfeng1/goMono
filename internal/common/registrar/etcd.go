package registrar

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/shiqinfeng1/goMono/internal/common/client"
	"github.com/shiqinfeng1/goMono/internal/common/config"
)

func EtcdRegistrar(dis *config.Discovery) (registry.Registrar, error) {
	if len(dis.Endpoints) == 0 {
		panic("no such discovery config endpoint")
	}
	eclient, err := client.NewEtcd(dis.Endpoints)
	if err != nil {
		return nil, err
	}
	return etcd.New(eclient), nil
}
