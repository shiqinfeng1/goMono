package client

import (
	"context"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	ggrpc "google.golang.org/grpc"
)

// 微服务内部之间的通信，无需安全选项
func NewGrpcConn(discovery registry.Discovery, srvName string) (*ggrpc.ClientConn, error) {
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
