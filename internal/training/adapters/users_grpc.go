package adapters

import (
	"context"

	userApi "github.com/shiqinfeng1/goMono/api/user/v1"
)

type UserGrpc struct {
	client userApi.UserServiceClient
}

func NewUserGrpc(client userApi.UserServiceClient) UserGrpc {
	return UserGrpc{client: client}
}

func (s UserGrpc) UpdateTrainingBalance(ctx context.Context, userID string, amountChange int) error {
	_, err := s.client.UpdateTrainingBalance(ctx, &userApi.UpdateTrainingBalanceRequest{
		UserId:       userID,
		AmountChange: int64(amountChange),
	})

	return err
}
