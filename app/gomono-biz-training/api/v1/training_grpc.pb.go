// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: gomono-biz-training/api/v1/training.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TrainingService_GetTraining_FullMethodName               = "/training.TrainingService/GetTraining"
	TrainingService_CreateTraining_FullMethodName            = "/training.TrainingService/CreateTraining"
	TrainingService_CancelTraining_FullMethodName            = "/training.TrainingService/CancelTraining"
	TrainingService_RescheduleTraining_FullMethodName        = "/training.TrainingService/RescheduleTraining"
	TrainingService_ApproveRescheduleTraining_FullMethodName = "/training.TrainingService/ApproveRescheduleTraining"
	TrainingService_RequestRescheduleTraining_FullMethodName = "/training.TrainingService/RequestRescheduleTraining"
	TrainingService_RejectRescheduleTraining_FullMethodName  = "/training.TrainingService/RejectRescheduleTraining"
)

// TrainingServiceClient is the client API for TrainingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrainingServiceClient interface {
	GetTraining(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTrainingResponse, error)
	CreateTraining(ctx context.Context, in *CreateTrainingRequest, opts ...grpc.CallOption) (*CreateTrainingResponse, error)
	CancelTraining(ctx context.Context, in *CancelTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RescheduleTraining(ctx context.Context, in *RescheduleTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ApproveRescheduleTraining(ctx context.Context, in *ApproveRescheduleTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RequestRescheduleTraining(ctx context.Context, in *RequestRescheduleTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RejectRescheduleTraining(ctx context.Context, in *RejectRescheduleTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type trainingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrainingServiceClient(cc grpc.ClientConnInterface) TrainingServiceClient {
	return &trainingServiceClient{cc}
}

func (c *trainingServiceClient) GetTraining(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTrainingResponse, error) {
	out := new(GetTrainingResponse)
	err := c.cc.Invoke(ctx, TrainingService_GetTraining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingServiceClient) CreateTraining(ctx context.Context, in *CreateTrainingRequest, opts ...grpc.CallOption) (*CreateTrainingResponse, error) {
	out := new(CreateTrainingResponse)
	err := c.cc.Invoke(ctx, TrainingService_CreateTraining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingServiceClient) CancelTraining(ctx context.Context, in *CancelTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TrainingService_CancelTraining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingServiceClient) RescheduleTraining(ctx context.Context, in *RescheduleTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TrainingService_RescheduleTraining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingServiceClient) ApproveRescheduleTraining(ctx context.Context, in *ApproveRescheduleTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TrainingService_ApproveRescheduleTraining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingServiceClient) RequestRescheduleTraining(ctx context.Context, in *RequestRescheduleTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TrainingService_RequestRescheduleTraining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingServiceClient) RejectRescheduleTraining(ctx context.Context, in *RejectRescheduleTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TrainingService_RejectRescheduleTraining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrainingServiceServer is the server API for TrainingService service.
// All implementations must embed UnimplementedTrainingServiceServer
// for forward compatibility
type TrainingServiceServer interface {
	GetTraining(context.Context, *emptypb.Empty) (*GetTrainingResponse, error)
	CreateTraining(context.Context, *CreateTrainingRequest) (*CreateTrainingResponse, error)
	CancelTraining(context.Context, *CancelTrainingRequest) (*emptypb.Empty, error)
	RescheduleTraining(context.Context, *RescheduleTrainingRequest) (*emptypb.Empty, error)
	ApproveRescheduleTraining(context.Context, *ApproveRescheduleTrainingRequest) (*emptypb.Empty, error)
	RequestRescheduleTraining(context.Context, *RequestRescheduleTrainingRequest) (*emptypb.Empty, error)
	RejectRescheduleTraining(context.Context, *RejectRescheduleTrainingRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTrainingServiceServer()
}

// UnimplementedTrainingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTrainingServiceServer struct {
}

func (UnimplementedTrainingServiceServer) GetTraining(context.Context, *emptypb.Empty) (*GetTrainingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTraining not implemented")
}
func (UnimplementedTrainingServiceServer) CreateTraining(context.Context, *CreateTrainingRequest) (*CreateTrainingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTraining not implemented")
}
func (UnimplementedTrainingServiceServer) CancelTraining(context.Context, *CancelTrainingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTraining not implemented")
}
func (UnimplementedTrainingServiceServer) RescheduleTraining(context.Context, *RescheduleTrainingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RescheduleTraining not implemented")
}
func (UnimplementedTrainingServiceServer) ApproveRescheduleTraining(context.Context, *ApproveRescheduleTrainingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveRescheduleTraining not implemented")
}
func (UnimplementedTrainingServiceServer) RequestRescheduleTraining(context.Context, *RequestRescheduleTrainingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestRescheduleTraining not implemented")
}
func (UnimplementedTrainingServiceServer) RejectRescheduleTraining(context.Context, *RejectRescheduleTrainingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectRescheduleTraining not implemented")
}
func (UnimplementedTrainingServiceServer) mustEmbedUnimplementedTrainingServiceServer() {}

// UnsafeTrainingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrainingServiceServer will
// result in compilation errors.
type UnsafeTrainingServiceServer interface {
	mustEmbedUnimplementedTrainingServiceServer()
}

func RegisterTrainingServiceServer(s grpc.ServiceRegistrar, srv TrainingServiceServer) {
	s.RegisterService(&TrainingService_ServiceDesc, srv)
}

func _TrainingService_GetTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingServiceServer).GetTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainingService_GetTraining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingServiceServer).GetTraining(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainingService_CreateTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingServiceServer).CreateTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainingService_CreateTraining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingServiceServer).CreateTraining(ctx, req.(*CreateTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainingService_CancelTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingServiceServer).CancelTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainingService_CancelTraining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingServiceServer).CancelTraining(ctx, req.(*CancelTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainingService_RescheduleTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RescheduleTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingServiceServer).RescheduleTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainingService_RescheduleTraining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingServiceServer).RescheduleTraining(ctx, req.(*RescheduleTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainingService_ApproveRescheduleTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveRescheduleTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingServiceServer).ApproveRescheduleTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainingService_ApproveRescheduleTraining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingServiceServer).ApproveRescheduleTraining(ctx, req.(*ApproveRescheduleTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainingService_RequestRescheduleTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRescheduleTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingServiceServer).RequestRescheduleTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainingService_RequestRescheduleTraining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingServiceServer).RequestRescheduleTraining(ctx, req.(*RequestRescheduleTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainingService_RejectRescheduleTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectRescheduleTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingServiceServer).RejectRescheduleTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainingService_RejectRescheduleTraining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingServiceServer).RejectRescheduleTraining(ctx, req.(*RejectRescheduleTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrainingService_ServiceDesc is the grpc.ServiceDesc for TrainingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrainingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "training.TrainingService",
	HandlerType: (*TrainingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTraining",
			Handler:    _TrainingService_GetTraining_Handler,
		},
		{
			MethodName: "CreateTraining",
			Handler:    _TrainingService_CreateTraining_Handler,
		},
		{
			MethodName: "CancelTraining",
			Handler:    _TrainingService_CancelTraining_Handler,
		},
		{
			MethodName: "RescheduleTraining",
			Handler:    _TrainingService_RescheduleTraining_Handler,
		},
		{
			MethodName: "ApproveRescheduleTraining",
			Handler:    _TrainingService_ApproveRescheduleTraining_Handler,
		},
		{
			MethodName: "RequestRescheduleTraining",
			Handler:    _TrainingService_RequestRescheduleTraining_Handler,
		},
		{
			MethodName: "RejectRescheduleTraining",
			Handler:    _TrainingService_RejectRescheduleTraining_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gomono-biz-training/api/v1/training.proto",
}
