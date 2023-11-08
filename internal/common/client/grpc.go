package client

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
	trainer "github.com/shiqinfeng1/goMono/api/trainer/v1"
	users "github.com/shiqinfeng1/goMono/api/users/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// trainer服务的frpc客户端的封装，便于其他服务调用
func NewTrainerClient(grpcAddr string) (client trainer.TrainerServiceClient, close func() error, err error) {
	if grpcAddr == "" {
		grpcAddr = os.Getenv("TRAINER_GRPC_ADDR")
	}
	if grpcAddr == "" {
		return nil, func() error { return nil }, errors.New("empty grpcAddr && env TRAINER_GRPC_ADDR")
	}

	opts, err := grpcDialOpts(grpcAddr)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	conn, err := grpc.Dial(grpcAddr, opts...)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return trainer.NewTrainerServiceClient(conn), conn.Close, nil
}

// 等待trainer服务启动完成
func WaitForTrainerService(timeout time.Duration) bool {
	return waitForPort(os.Getenv("TRAINER_GRPC_ADDR"), timeout)
}

// users服务的grpc客户端
func NewUsersClient(grpcAddr string) (client users.UsersServiceClient, close func() error, err error) {
	if grpcAddr == "" {
		grpcAddr = os.Getenv("USERS_GRPC_ADDR")
	}
	if grpcAddr == "" {
		return nil, func() error { return nil }, errors.New("empty grpcAddr && env USERS_GRPC_ADDR")
	}

	opts, err := grpcDialOpts(grpcAddr)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	conn, err := grpc.Dial(grpcAddr, opts...)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return users.NewUsersServiceClient(conn), conn.Close, nil
}

// 等待users服务启动完成
func WaitForUsersService(timeout time.Duration) bool {
	return waitForPort(os.Getenv("USERS_GRPC_ADDR"), timeout)
}

// 配置grpc连接的参数选项
func grpcDialOpts(grpcAddr string) ([]grpc.DialOption, error) {
	if noTLS, _ := strconv.ParseBool(os.Getenv("GRPC_NO_TLS")); noTLS {
		return []grpc.DialOption{}, nil
	}
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		return nil, errors.Wrap(err, "cannot load root CA cert")
	}
	creds := credentials.NewTLS(&tls.Config{
		RootCAs:    systemRoots,
		MinVersion: tls.VersionTLS12,
	})

	return []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMetadataServerToken(grpcAddr)),
	}, nil
}
