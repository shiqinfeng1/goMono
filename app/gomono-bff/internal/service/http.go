package service

import (
	"context"
	"fmt"

	v1 "github.com/shiqinfeng1/goMono/app/gomono-bff/api/v1"
	"github.com/shiqinfeng1/goMono/app/gomono-bff/internal/application"
	trainingV1 "github.com/shiqinfeng1/goMono/app/gomono-biz-training/api/v1"
	"github.com/shiqinfeng1/goMono/utils/auth"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HttpService struct {
	app application.Application
}

func NewHttpService(application application.Application) *HttpService {
	return &HttpService{
		app: application,
	}
}

func (h HttpService) GetTrainerAvailableHours(ctx context.Context, req *v1.GetTrainerAvailableHoursRequest) (*v1.GetTrainerAvailableHoursRespone, error) {
	return &v1.GetTrainerAvailableHoursRespone{}, nil
}

func (h HttpService) MakeHourAvailable(ctx context.Context, req *v1.MakeHourAvailableRequest) (*emptypb.Empty, error) {
	user, err := auth.UserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if user.Role != "trainer" {
		return nil, fmt.Errorf("invalid-role:%v", user.Role)
	}

	return &emptypb.Empty{}, nil
}

func (h HttpService) MakeHourUnavailable(ctx context.Context, req *v1.MakeHourUnavailableRequest) (*emptypb.Empty, error) {
	user, err := auth.UserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if user.Role != "trainer" {
		return nil, fmt.Errorf("invalid-role:%v", user.Role)
	}

	return &emptypb.Empty{}, nil
}

func (h HttpService) GetTraining(ctx context.Context, epy *emptypb.Empty) (*trainingV1.GetTrainingResponse, error) {
	return &trainingV1.GetTrainingResponse{}, nil
}

func (h HttpService) CreateTraining(ctx context.Context, req *trainingV1.CreateTrainingRequest) (*trainingV1.CreateTrainingResponse, error) {
	return &trainingV1.CreateTrainingResponse{}, nil
}

func (h HttpService) CancelTraining(ctx context.Context, req *trainingV1.CancelTrainingRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (h HttpService) RescheduleTraining(ctx context.Context, req *trainingV1.RescheduleTrainingRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (h HttpService) ApproveRescheduleTraining(ctx context.Context, req *trainingV1.ApproveRescheduleTrainingRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (h HttpService) RequestRescheduleTraining(ctx context.Context, req *trainingV1.RequestRescheduleTrainingRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (h HttpService) RejectRescheduleTraining(ctx context.Context, req *trainingV1.RejectRescheduleTrainingRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
