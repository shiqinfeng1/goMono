package ports

import (
	v1 "github.com/shiqinfeng1/goMono/api/trainer/v1"
	conf "github.com/shiqinfeng1/goMono/internal/common/config/trainer"
	"github.com/shiqinfeng1/goMono/internal/trainer/service"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.GRPC, trainer service.GrpcService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Network != "" {
		opts = append(opts, grpc.Network(c.Network))
	}
	if c.Addr != "" {
		opts = append(opts, grpc.Address(c.Addr))
	}
	if c.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterTrainerServiceServer(srv, trainer)
	return srv
}
