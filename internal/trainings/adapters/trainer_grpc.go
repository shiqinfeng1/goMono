package adapters

import (
	"context"
	"time"

	"github.com/pkg/errors"
	trainer "github.com/shiqinfeng1/goMono/api/trainer/v1"
	grpcClient "github.com/shiqinfeng1/goMono/internal/common/grpc"
	"github.com/shiqinfeng1/goMono/internal/trainings/conf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewTrainerClient
func NewTrainerClient(c *conf.Adapter) (trainer.TrainerServiceClient, func(), error) {
	client, close, err := grpcClient.NewTrainerClient(c.TrainerGrpc.GetAddr())
	return client, func() { _ = close() }, err
}

type TrainerGrpc struct {
	client trainer.TrainerServiceClient
}

func NewTrainerGrpc(client trainer.TrainerServiceClient) TrainerGrpc {
	return TrainerGrpc{client: client}
}

func (s TrainerGrpc) ScheduleTraining(ctx context.Context, trainingTime time.Time) error {
	_, err := s.client.ScheduleTraining(ctx, &trainer.UpdateHourRequest{
		Time: timestamppb.New(trainingTime),
	})

	return err
}

func (s TrainerGrpc) CancelTraining(ctx context.Context, trainingTime time.Time) error {
	_, err := s.client.CancelTraining(ctx, &trainer.UpdateHourRequest{
		Time: timestamppb.New(trainingTime),
	})

	return err
}

func (s TrainerGrpc) MoveTraining(
	ctx context.Context,
	newTime time.Time,
	originalTrainingTime time.Time,
) error {
	err := s.ScheduleTraining(ctx, newTime)
	if err != nil {
		return errors.Wrap(err, "unable to schedule training")
	}

	err = s.CancelTraining(ctx, originalTrainingTime)
	if err != nil {
		return errors.Wrap(err, "unable to cancel training")
	}

	return nil
}
