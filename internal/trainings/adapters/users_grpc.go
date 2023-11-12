package adapters

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	users "github.com/shiqinfeng1/goMono/api/users/v1"
	grpcClient "github.com/shiqinfeng1/goMono/internal/common/grpc"
	"github.com/shiqinfeng1/goMono/internal/trainings/conf"
)

// NewTrainerGrpc
func NewUsersClient(c *conf.Adapter, logger log.Logger) (users.UsersServiceClient, func(), error) {
	client, close, err := grpcClient.NewUsersClient(c.UserGrpc.GetAddr())
	return client, func() { _ = close() }, err
}

type UsersGrpc struct {
	client users.UsersServiceClient
}

func NewUsersGrpc(client users.UsersServiceClient) UsersGrpc {
	return UsersGrpc{client: client}
}

func (s UsersGrpc) UpdateTrainingBalance(ctx context.Context, userID string, amountChange int) error {
	_, err := s.client.UpdateTrainingBalance(ctx, &users.UpdateTrainingBalanceRequest{
		UserId:       userID,
		AmountChange: int64(amountChange),
	})

	return err
}
