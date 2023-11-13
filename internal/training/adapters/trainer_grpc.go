package adapters

import (
	"context"
	"time"

	"github.com/pkg/errors"
	trainerApi "github.com/shiqinfeng1/goMono/api/trainer/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TrainerGrpc struct {
	client trainerApi.TrainerServiceClient
}

func NewTrainerGrpc(client trainerApi.TrainerServiceClient) TrainerGrpc {
	return TrainerGrpc{client: client}
}

func (s TrainerGrpc) ScheduleTraining(ctx context.Context, trainingTime time.Time) error {
	_, err := s.client.ScheduleTraining(ctx, &trainerApi.UpdateHourRequest{
		Time: timestamppb.New(trainingTime),
	})

	return err
}

func (s TrainerGrpc) CancelTraining(ctx context.Context, trainingTime time.Time) error {
	_, err := s.client.CancelTraining(ctx, &trainerApi.UpdateHourRequest{
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
