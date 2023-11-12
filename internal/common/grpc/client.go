package client

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	trainerApi "github.com/shiqinfeng1/goMono/api/trainer/v1"
	userApi "github.com/shiqinfeng1/goMono/api/users/v1"
	etcdClient "go.etcd.io/etcd/client/v3"
	ggrpc "google.golang.org/grpc"
)

func newConnWithEtcd(endpoints []string, srvName string) (*ggrpc.ClientConn, func() error, error) {
	eclient, err := etcdClient.New(etcdClient.Config{
		Endpoints:   endpoints, //[]string{"127.0.0.1:2379"},
		DialTimeout: time.Second,
		DialOptions: []ggrpc.DialOption{ggrpc.WithBlock()},
	})
	if err != nil {
		return nil, func() error { return nil }, err
	}
	r := etcd.New(eclient)
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+srvName),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		return nil, func() error { return nil }, err
	}
	return conn, func() error { return conn.Close() }, nil
}

// trainer服务的rpc客户端的封装，便于其他微服务调用
func NewTrainerClient(endpoints []string) (client trainerApi.TrainerServiceClient, close func() error, err error) {
	conn, close, err := newConnWithEtcd(endpoints, "trainer")
	if err != nil {
		return nil, func() error { return nil }, err
	}
	return trainerApi.NewTrainerServiceClient(conn), func() error { conn.Close(); return close() }, nil
}

// users服务的grpc客户端
func NewUsersClient(endpoints []string) (client userApi.UsersServiceClient, close func() error, err error) {
	conn, close, err := newConnWithEtcd(endpoints, "user")
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return userApi.NewUsersServiceClient(conn), func() error { conn.Close(); return close() }, nil
}
