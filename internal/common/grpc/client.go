package grpc

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	trainerApi "github.com/shiqinfeng1/goMono/api/trainer/v1"
	trainingApi "github.com/shiqinfeng1/goMono/api/training/v1"
	userApi "github.com/shiqinfeng1/goMono/api/user/v1"
	etcdClient "go.etcd.io/etcd/client/v3"
	ggrpc "google.golang.org/grpc"
)

// 微服务内部之间的通信，无需安全选项
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
func NewTrainerClient(endpoints []string) (client trainerApi.TrainerServiceClient, close func(), err error) {
	conn, closeConn, err := newConnWithEtcd(endpoints, "trainer")
	if err != nil {
		return nil, func() {}, err
	}
	client = trainerApi.NewTrainerServiceClient(conn)
	close = func() {
		_ = conn.Close()
		_ = closeConn()
	}
	return
}

// user服务的grpc客户端
func NewUserClient(endpoints []string) (client userApi.UserServiceClient, close func(), err error) {
	conn, closeConn, err := newConnWithEtcd(endpoints, "user")
	if err != nil {
		return nil, func() {}, err
	}
	client = userApi.NewUserServiceClient(conn)
	close = func() {
		_ = conn.Close()
		_ = closeConn()
	}
	return
}

// user服务的grpc客户端
func NewTrainingClient(endpoints []string) (client trainingApi.TrainingServiceClient, close func(), err error) {
	conn, closeConn, err := newConnWithEtcd(endpoints, "training")
	if err != nil {
		return nil, func() {}, err
	}
	client = trainingApi.NewTrainingServiceClient(conn)
	close = func() {
		_ = conn.Close()
		_ = closeConn()
	}
	return
}
