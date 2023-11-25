package grpc

import (
	"context"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	trainerApi "github.com/shiqinfeng1/goMono/api/trainer/v1"
	trainingApi "github.com/shiqinfeng1/goMono/api/training/v1"
	userApi "github.com/shiqinfeng1/goMono/api/user/v1"

	ggrpc "google.golang.org/grpc"
)

// 微服务内部之间的通信，无需安全选项
func NewGrpcClientConn(discovery registry.Discovery, srvName string) (*ggrpc.ClientConn, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+srvName),
		grpc.WithDiscovery(discovery),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// trainer服务的rpc客户端的封装，便于其他微服务调用
func NewTrainerClient(discovery registry.Discovery) (client trainerApi.TrainerServiceClient, close func(), err error) {
	conn, err := NewGrpcClientConn(discovery, "trainer")
	if err != nil {
		return nil, func() {}, err
	}
	client = trainerApi.NewTrainerServiceClient(conn)
	close = func() {
		_ = conn.Close()
	}
	return
}

// user服务的grpc客户端
func NewUserClient(discovery registry.Discovery) (client userApi.UserServiceClient, close func(), err error) {
	conn, err := NewGrpcClientConn(discovery, "user")
	if err != nil {
		return nil, func() {}, err
	}
	client = userApi.NewUserServiceClient(conn)
	close = func() {
		_ = conn.Close()
	}
	return
}

// user服务的grpc客户端
func NewTrainingClient(discovery registry.Discovery) (client trainingApi.TrainingServiceClient, close func(), err error) {
	conn, err := NewGrpcClientConn(discovery, "training")
	if err != nil {
		return nil, func() {}, err
	}
	client = trainingApi.NewTrainingServiceClient(conn)
	close = func() {
		_ = conn.Close()
	}
	return
}
