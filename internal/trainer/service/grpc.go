package service

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	v1 "github.com/shiqinfeng1/goMono/api/trainer/v1"
	"github.com/shiqinfeng1/goMono/internal/trainer/app"
	"github.com/shiqinfeng1/goMono/internal/trainer/app/command"
	"github.com/shiqinfeng1/goMono/internal/trainer/app/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}

func (g GrpcServer) MakeHourAvailable(ctx context.Context, request *v1.UpdateHourRequest) (*empty.Empty, error) {
	trainingTime := protoTimestampToTime(request.Time)

	if err := g.app.Commands.MakeHoursAvailable.Handle(ctx, command.MakeHoursAvailable{Hours: []time.Time{trainingTime}}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcServer) ScheduleTraining(ctx context.Context, request *v1.UpdateHourRequest) (*empty.Empty, error) {
	trainingTime := protoTimestampToTime(request.Time)

	if err := g.app.Commands.ScheduleTraining.Handle(ctx, command.ScheduleTraining{Hour: trainingTime}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcServer) CancelTraining(ctx context.Context, request *v1.UpdateHourRequest) (*empty.Empty, error) {
	trainingTime := protoTimestampToTime(request.Time)

	if err := g.app.Commands.CancelTraining.Handle(ctx, command.CancelTraining{Hour: trainingTime}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcServer) IsHourAvailable(ctx context.Context, request *v1.IsHourAvailableRequest) (*v1.IsHourAvailableResponse, error) {
	trainingTime := protoTimestampToTime(request.Time)

	isAvailable, err := g.app.Queries.HourAvailability.Handle(ctx, query.HourAvailability{Hour: trainingTime})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &trainer.IsHourAvailableResponse{IsAvailable: isAvailable}, nil
}

func protoTimestampToTime(timestamp *timestamp.Timestamp) time.Time {
	return timestamp.AsTime().UTC().Truncate(time.Hour)
}
