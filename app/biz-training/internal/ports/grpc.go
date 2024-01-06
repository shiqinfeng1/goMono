package ports

import (
	"context"

	v1 "github.com/shiqinfeng1/goMono/api/training/v1"
	conf "github.com/shiqinfeng1/goMono/app/biz-training/internal/conf"
	"github.com/shiqinfeng1/goMono/app/biz-training/internal/service"
	"github.com/shiqinfeng1/goMono/app/common/client"
	"github.com/shiqinfeng1/goMono/app/common/trace"

	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	kmetrics "github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	cconf "github.com/shiqinfeng1/goMono/app/common/conf"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(ctx context.Context, c *conf.GRPC, logger log.Logger, tr *cconf.Trace, training service.GrpcService) *grpc.Server {
	trace.NewTrace(ctx, tr)
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			kmetrics.Server(
				kmetrics.WithSeconds(prom.NewHistogram(client.MetricsSeconds)),
				kmetrics.WithRequests(prom.NewCounter(client.MetricsRequests)),
			),
			logging.Server(log.With(logger,
				"layer", "ports",
				"trace.id", tracing.TraceID(),
				"span.id", tracing.SpanID(),
			)),
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
	v1.RegisterTrainingServiceServer(srv, training)
	return srv
}
