package service

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/google/uuid"
	v1 "github.com/shiqinfeng1/goMono/api/training/v1"
	"github.com/shiqinfeng1/goMono/internal/common/auth"
	"github.com/shiqinfeng1/goMono/internal/training/internal/application"
	"github.com/shiqinfeng1/goMono/internal/training/internal/application/command"
	"github.com/shiqinfeng1/goMono/internal/training/internal/application/query"
	"github.com/shiqinfeng1/goMono/internal/training/internal/domain/training"
)

type HttpService struct {
	app application.Application
}

func NewHttpService(app application.Application) *HttpService {
	return &HttpService{app}
}

func (h HttpService) GetTraining(ctx context.Context, req *emptypb.Empty) (*v1.GetTrainingResponse, error) {
	user, err := auth.UserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	var appTraining []query.Training

	if user.Role == "trainer" {
		appTraining, err = h.app.Queries.AllTraining.Handle(ctx, query.AllTraining{})
	} else {
		appTraining, err = h.app.Queries.TrainingForUser.Handle(ctx, query.TrainingForUser{User: user})
	}

	if err != nil {
		return nil, err
	}

	training := appTrainingToResponse(appTraining)
	return &v1.GetTrainingResponse{
		Training: training,
	}, nil
}

func (h HttpService) CreateTraining(ctx context.Context, req *v1.CreateTrainingRequest) (*v1.CreateTrainingResponse, error) {

	user, err := auth.UserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if user.Role != "attendee" {
		return nil, fmt.Errorf("role(%v) has no auth", user.Role)
	}

	cmd := command.ScheduleTraining{
		TrainingUUID: uuid.New().String(),
		UserUUID:     user.UUID,
		UserName:     user.DisplayName,
		TrainingTime: req.Time.AsTime(),
		Notes:        req.Notes,
	}
	err = h.app.Commands.ScheduleTraining.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return &v1.CreateTrainingResponse{
		TrainingUuid: cmd.TrainingUUID,
	}, nil
}

func (h HttpService) CancelTraining(ctx context.Context, req *v1.CancelTrainingRequest) (*emptypb.Empty, error) {
	user, err := newDomainUserFromAuthUser(ctx)
	if err != nil {
		return nil, nil
	}

	err = h.app.Commands.CancelTraining.Handle(ctx, command.CancelTraining{
		TrainingUUID: req.TrainingUuid,
		User:         user,
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h HttpService) RescheduleTraining(ctx context.Context, req *v1.RescheduleTrainingRequest) (*emptypb.Empty, error) {

	user, err := newDomainUserFromAuthUser(ctx)
	if err != nil {
		return nil, err
	}

	err = h.app.Commands.RescheduleTraining.Handle(ctx, command.RescheduleTraining{
		User:         user,
		TrainingUUID: req.TrainingUuid,
		NewTime:      req.Time.AsTime(),
		NewNotes:     req.Notes,
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h HttpService) RequestRescheduleTraining(ctx context.Context, req *v1.RequestRescheduleTrainingRequest) (*emptypb.Empty, error) {

	user, err := newDomainUserFromAuthUser(ctx)
	if err != nil {
		return nil, err
	}

	err = h.app.Commands.RequestTrainingReschedule.Handle(ctx, command.RequestTrainingReschedule{
		User:         user,
		TrainingUUID: req.TrainingUuid,
		NewTime:      req.Time.AsTime(),
		NewNotes:     req.Notes,
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h HttpService) ApproveRescheduleTraining(ctx context.Context, req *v1.ApproveRescheduleTrainingRequest) (*emptypb.Empty, error) {
	user, err := newDomainUserFromAuthUser(ctx)
	if err != nil {
		return nil, err
	}
	err = h.app.Commands.ApproveTrainingReschedule.Handle(ctx, command.ApproveTrainingReschedule{
		User:         user,
		TrainingUUID: req.TrainingUuid,
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h HttpService) RejectRescheduleTraining(ctx context.Context, req *v1.RejectRescheduleTrainingRequest) (*emptypb.Empty, error) {
	user, err := newDomainUserFromAuthUser(ctx)
	if err != nil {
		return nil, err
	}

	err = h.app.Commands.RejectTrainingReschedule.Handle(ctx, command.RejectTrainingReschedule{
		User:         user,
		TrainingUUID: req.TrainingUuid,
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func appTrainingToResponse(appTraining []query.Training) []*v1.GetTrainingResponse_Training {
	var training []*v1.GetTrainingResponse_Training
	for _, tm := range appTraining {
		t := &v1.GetTrainingResponse_Training{
			CanBeCancelled:     tm.CanBeCancelled,
			MoveProposedBy:     *tm.MoveProposedBy,
			MoveRequiresAccept: tm.CanBeCancelled,
			Notes:              tm.Notes,
			ProposedTime:       timestamppb.New(*tm.ProposedTime),
			Time:               timestamppb.New(tm.Time),
			User:               tm.User,
			UserUuid:           tm.UserUUID,
			Uuid:               tm.UUID,
		}

		training = append(training, t)
	}

	return training
}

func newDomainUserFromAuthUser(ctx context.Context) (training.User, error) {
	user, err := auth.UserFromCtx(ctx)
	if err != nil {
		return training.User{}, err
	}

	userType, err := training.NewUserTypeFromString(user.Role)
	if err != nil {
		return training.User{}, err
	}

	return training.NewUser(user.UUID, userType)
}
