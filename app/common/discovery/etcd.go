package discovery

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/shiqinfeng1/goMono/app/common/client"
)

func EtcdDiscovery(endpoints []string) (registry.Discovery, error) {
	eclient, err := client.NewEtcd(endpoints)
	if err != nil {
		return nil, err
	}
	return etcd.New(eclient), nil
}
