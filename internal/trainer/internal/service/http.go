package service

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/shiqinfeng1/goMono/api/trainer/v1"
	"github.com/shiqinfeng1/goMono/internal/common/auth"
	"github.com/shiqinfeng1/goMono/internal/trainer/internal/application"
	"github.com/shiqinfeng1/goMono/internal/trainer/internal/application/command"
	"github.com/shiqinfeng1/goMono/internal/trainer/internal/application/query"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	dateModels, err := h.app.Queries.TrainerAvailableHours.Handle(ctx, query.AvailableHours{
		From: req.DateFrom.AsTime(),
		To:   req.DateTo.AsTime(),
	})
	if err != nil {
		return nil, err
	}

	dates := dateModelsToResponse(dateModels)
	return &v1.GetTrainerAvailableHoursRespone{
		Dates: dates,
	}, nil
}

func dateModelsToResponse(model query.Date) []*v1.GetTrainerAvailableHoursRespone_Date {
	var dates []*v1.GetTrainerAvailableHoursRespone_Date

	var hours []*v1.GetTrainerAvailableHoursRespone_Hour
	for _, h := range model.Hours {
		hours = append(hours, &v1.GetTrainerAvailableHoursRespone_Hour{
			Available:            h.Available,
			HasTrainingScheduled: h.HasTrainingScheduled,
			Hour:                 timestamppb.New(h.Hour),
		})
	}

	dates = append(dates, &v1.GetTrainerAvailableHoursRespone_Date{
		Date:         timestamppb.New(model.Date),
		HasFreeHours: model.HasFreeHours,
		Hours:        hours,
	})

	return dates
}

func (h HttpService) MakeHourAvailable(ctx context.Context, req *v1.MakeHourAvailableRequest) (*emptypb.Empty, error) {
	user, err := auth.UserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if user.Role != "trainer" {
		return nil, fmt.Errorf("invalid-role:%v", user.Role)
	}

	hourUpdate := make([]time.Time, 0)
	for _, hour := range req.Time {
		hourUpdate = append(hourUpdate, hour.AsTime())
	}

	err = h.app.Commands.MakeHoursAvailable.Handle(ctx, command.MakeHoursAvailable{Hours: hourUpdate})
	if err != nil {
		return nil, err
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

	hourUpdate := make([]time.Time, 0)
	for _, hour := range req.Time {
		hourUpdate = append(hourUpdate, hour.AsTime())
	}

	err = h.app.Commands.MakeHoursUnavailable.Handle(ctx, command.MakeHoursUnavailable{Hours: hourUpdate})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
