package adapters

import (
	"context"
	"fmt"

	userApi "github.com/shiqinfeng1/goMono/api/user/v1"
	"github.com/shiqinfeng1/goMono/internal/common/client"
	"github.com/shiqinfeng1/goMono/internal/common/config"
	"github.com/shiqinfeng1/goMono/internal/common/discovery"
)

type UserGrpc struct {
	endpoints []string
	client    userApi.UserServiceClient
	close     func() error
}

func NewUserGrpc(dis *config.Discovery) *UserGrpc {
	return &UserGrpc{
		endpoints: dis.Endpoints,
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
		conn, err := client.NewGrpcConn(dis, "user")
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
