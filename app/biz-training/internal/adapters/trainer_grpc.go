package adapters

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	trainerApi "github.com/shiqinfeng1/goMono/api/trainer/v1"
	"github.com/shiqinfeng1/goMono/app/common/client"
	"github.com/shiqinfeng1/goMono/app/common/conf"
	"github.com/shiqinfeng1/goMono/app/common/discovery"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	once sync.Once
)

type TrainerGrpc struct {
	logger    log.Logger
	endpoints []string
	client    trainerApi.TrainerServiceClient
	close     func() error
}

func NewTrainerGrpc(dis *conf.Discovery, logger log.Logger) *TrainerGrpc {
	return &TrainerGrpc{
		endpoints: dis.Endpoints,
		logger:    logger,
	}
}
func (s TrainerGrpc) Close() {
	if s.close != nil {
		s.close()
	}
}
func (s *TrainerGrpc) getClient() trainerApi.TrainerServiceClient {
	once.Do(func() {
		dis, err := discovery.EtcdDiscovery(s.endpoints)
		if err != nil {
			panic(fmt.Errorf("invalid discovery %v: %w", s.endpoints, err))
		}
		conn, err := client.NewGrpcConn(dis, s.logger, "trainer")
		if err != nil {
			panic(fmt.Errorf("invalid trainer client from %v: %w", s.endpoints, err))
		}
		s.client = trainerApi.NewTrainerServiceClient(conn)
		s.close = conn.Close
	})
	return s.client
}

func (s TrainerGrpc) ScheduleTraining(ctx context.Context, trainingTime time.Time) error {
	_, err := s.getClient().ScheduleTraining(ctx, &trainerApi.UpdateHourRequest{
		Time: timestamppb.New(trainingTime),
	})

	return err
}

func (s TrainerGrpc) CancelTraining(ctx context.Context, trainingTime time.Time) error {
	_, err := s.getClient().CancelTraining(ctx, &trainerApi.UpdateHourRequest{
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
