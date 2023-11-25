package grpc

import (
	"time"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	etcdClient "go.etcd.io/etcd/client/v3"
	ggrpc "google.golang.org/grpc"
)

func EtcdDiscovery(endpoints []string) (registry.Discovery, error) {
	eclient, err := etcdClient.New(etcdClient.Config{
		Endpoints:   endpoints, //[]string{"127.0.0.1:2379"},
		DialTimeout: time.Second,
		DialOptions: []ggrpc.DialOption{ggrpc.WithBlock()},
	})
	if err != nil {
		return nil, err
	}
	return etcd.New(eclient), nil
}
