package adapters

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	userApi "github.com/shiqinfeng1/goMono/api/user/v1"
	"github.com/shiqinfeng1/goMono/app/common/client"
	"github.com/shiqinfeng1/goMono/app/common/conf"
	"github.com/shiqinfeng1/goMono/app/common/discovery"
)

type UserGrpc struct {
	logger    log.Logger
	endpoints []string
	client    userApi.UserServiceClient
	close     func() error
}

func NewUserGrpc(dis *conf.Discovery, logger log.Logger) *UserGrpc {
	return &UserGrpc{
		endpoints: dis.Endpoints,
		logger: log.With(logger,
			"trace.id", tracing.TraceID(),
			"span.id", tracing.SpanID(),
		),
	}
}
func (s UserGrpc) Close() {
	if s.close != nil {
		s.close()
	}
}
func (s *UserGrpc) getClient() userApi.UserServiceClient {
	once.Do(func() {
		dis, err := discovery.EtcdDiscovery(s.endpoints)
		if err != nil {
			panic(fmt.Errorf("invalid discovery %v: %w", s.endpoints, err))
		}
		conn, err := client.NewGrpcConn(dis, s.logger, "user")
		if err != nil {
			panic(fmt.Errorf("invalid trainer client from %v: %w", s.endpoints, err))
		}
		s.client = userApi.NewUserServiceClient(conn)
		s.close = conn.Close
	})
	return s.client
}
func (s UserGrpc) UpdateTrainingBalance(ctx context.Context, userID string, amountChange int) error {
	_, err := s.getClient().UpdateTrainingBalance(ctx, &userApi.UpdateTrainingBalanceRequest{
		UserId:       userID,
		AmountChange: int64(amountChange),
	})

	return err
}
