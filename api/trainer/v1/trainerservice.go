package service

import (
	"context"

	pb "github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer"
)

type TrainerServiceService struct {
	pb.UnimplementedTrainerServiceServer
}

func NewTrainerServiceService() *TrainerServiceService {
	return &TrainerServiceService{}
}

func (s *TrainerServiceService) IsHourAvailable(ctx context.Context, req *pb.IsHourAvailableRequest) (*pb.IsHourAvailableResponse, error) {
	return &pb.IsHourAvailableResponse{}, nil
}
func (s *TrainerServiceService) ScheduleTraining(ctx context.Context, req *pb.UpdateHourRequest) (*pb.google_protobuf_Empty, error) {
	return &pb.google_protobuf_Empty{}, nil
}
func (s *TrainerServiceService) CancelTraining(ctx context.Context, req *pb.UpdateHourRequest) (*pb.google_protobuf_Empty, error) {
	return &pb.google_protobuf_Empty{}, nil
}
func (s *TrainerServiceService) MakeHourAvailable(ctx context.Context, req *pb.UpdateHourRequest) (*pb.google_protobuf_Empty, error) {
	return &pb.google_protobuf_Empty{}, nil
}
