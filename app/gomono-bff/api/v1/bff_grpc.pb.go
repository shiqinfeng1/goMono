// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: gomono-bff/api/v1/bff.proto

package v1

import (
	context "context"
	v1 "github.com/shiqinfeng1/goMono/app/gomono-biz-training/api/v1"
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
	BFF_IsHourAvailable_FullMethodName           = "/bff.BFF/IsHourAvailable"
	BFF_ScheduleTraining_FullMethodName          = "/bff.BFF/ScheduleTraining"
	BFF_GetTrainerAvailableHours_FullMethodName  = "/bff.BFF/GetTrainerAvailableHours"
	BFF_MakeHourAvailable_FullMethodName         = "/bff.BFF/MakeHourAvailable"
	BFF_MakeHourUnavailable_FullMethodName       = "/bff.BFF/MakeHourUnavailable"
	BFF_GetTraining_FullMethodName               = "/bff.BFF/GetTraining"
	BFF_CreateTraining_FullMethodName            = "/bff.BFF/CreateTraining"
	BFF_CancelTraining_FullMethodName            = "/bff.BFF/CancelTraining"
	BFF_RescheduleTraining_FullMethodName        = "/bff.BFF/RescheduleTraining"
	BFF_ApproveRescheduleTraining_FullMethodName = "/bff.BFF/ApproveRescheduleTraining"
	BFF_RequestRescheduleTraining_FullMethodName = "/bff.BFF/RequestRescheduleTraining"
	BFF_RejectRescheduleTraining_FullMethodName  = "/bff.BFF/RejectRescheduleTraining"
)

// BFFClient is the client API for BFF service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BFFClient interface {
	IsHourAvailable(ctx context.Context, in *IsHourAvailableRequest, opts ...grpc.CallOption) (*IsHourAvailableResponse, error)
	ScheduleTraining(ctx context.Context, in *UpdateHourRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTrainerAvailableHours(ctx context.Context, in *GetTrainerAvailableHoursRequest, opts ...grpc.CallOption) (*GetTrainerAvailableHoursRespone, error)
	MakeHourAvailable(ctx context.Context, in *MakeHourAvailableRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MakeHourUnavailable(ctx context.Context, in *MakeHourUnavailableRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTraining(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.GetTrainingResponse, error)
	CreateTraining(ctx context.Context, in *v1.CreateTrainingRequest, opts ...grpc.CallOption) (*v1.CreateTrainingResponse, error)
	CancelTraining(ctx context.Context, in *v1.CancelTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RescheduleTraining(ctx context.Context, in *v1.RescheduleTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ApproveRescheduleTraining(ctx context.Context, in *v1.ApproveRescheduleTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RequestRescheduleTraining(ctx context.Context, in *v1.RequestRescheduleTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RejectRescheduleTraining(ctx context.Context, in *v1.RejectRescheduleTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type bFFClient struct {
	cc grpc.ClientConnInterface
}

func NewBFFClient(cc grpc.ClientConnInterface) BFFClient {
	return &bFFClient{cc}
}

func (c *bFFClient) IsHourAvailable(ctx context.Context, in *IsHourAvailableRequest, opts ...grpc.CallOption) (*IsHourAvailableResponse, error) {
	out := new(IsHourAvailableResponse)
	err := c.cc.Invoke(ctx, BFF_IsHourAvailable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) ScheduleTraining(ctx context.Context, in *UpdateHourRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BFF_ScheduleTraining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) GetTrainerAvailableHours(ctx context.Context, in *GetTrainerAvailableHoursRequest, opts ...grpc.CallOption) (*GetTrainerAvailableHoursRespone, error) {
	out := new(GetTrainerAvailableHoursRespone)
	err := c.cc.Invoke(ctx, BFF_GetTrainerAvailableHours_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) MakeHourAvailable(ctx context.Context, in *MakeHourAvailableRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BFF_MakeHourAvailable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) MakeHourUnavailable(ctx context.Context, in *MakeHourUnavailableRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BFF_MakeHourUnavailable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) GetTraining(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.GetTrainingResponse, error) {
	out := new(v1.GetTrainingResponse)
	err := c.cc.Invoke(ctx, BFF_GetTraining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) CreateTraining(ctx context.Context, in *v1.CreateTrainingRequest, opts ...grpc.CallOption) (*v1.CreateTrainingResponse, error) {
	out := new(v1.CreateTrainingResponse)
	err := c.cc.Invoke(ctx, BFF_CreateTraining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) CancelTraining(ctx context.Context, in *v1.CancelTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BFF_CancelTraining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) RescheduleTraining(ctx context.Context, in *v1.RescheduleTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BFF_RescheduleTraining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) ApproveRescheduleTraining(ctx context.Context, in *v1.ApproveRescheduleTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BFF_ApproveRescheduleTraining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) RequestRescheduleTraining(ctx context.Context, in *v1.RequestRescheduleTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BFF_RequestRescheduleTraining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) RejectRescheduleTraining(ctx context.Context, in *v1.RejectRescheduleTrainingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BFF_RejectRescheduleTraining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BFFServer is the server API for BFF service.
// All implementations must embed UnimplementedBFFServer
// for forward compatibility
type BFFServer interface {
	IsHourAvailable(context.Context, *IsHourAvailableRequest) (*IsHourAvailableResponse, error)
	ScheduleTraining(context.Context, *UpdateHourRequest) (*emptypb.Empty, error)
	GetTrainerAvailableHours(context.Context, *GetTrainerAvailableHoursRequest) (*GetTrainerAvailableHoursRespone, error)
	MakeHourAvailable(context.Context, *MakeHourAvailableRequest) (*emptypb.Empty, error)
	MakeHourUnavailable(context.Context, *MakeHourUnavailableRequest) (*emptypb.Empty, error)
	GetTraining(context.Context, *emptypb.Empty) (*v1.GetTrainingResponse, error)
	CreateTraining(context.Context, *v1.CreateTrainingRequest) (*v1.CreateTrainingResponse, error)
	CancelTraining(context.Context, *v1.CancelTrainingRequest) (*emptypb.Empty, error)
	RescheduleTraining(context.Context, *v1.RescheduleTrainingRequest) (*emptypb.Empty, error)
	ApproveRescheduleTraining(context.Context, *v1.ApproveRescheduleTrainingRequest) (*emptypb.Empty, error)
	RequestRescheduleTraining(context.Context, *v1.RequestRescheduleTrainingRequest) (*emptypb.Empty, error)
	RejectRescheduleTraining(context.Context, *v1.RejectRescheduleTrainingRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedBFFServer()
}

// UnimplementedBFFServer must be embedded to have forward compatible implementations.
type UnimplementedBFFServer struct {
}

func (UnimplementedBFFServer) IsHourAvailable(context.Context, *IsHourAvailableRequest) (*IsHourAvailableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsHourAvailable not implemented")
}
func (UnimplementedBFFServer) ScheduleTraining(context.Context, *UpdateHourRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleTraining not implemented")
}
func (UnimplementedBFFServer) GetTrainerAvailableHours(context.Context, *GetTrainerAvailableHoursRequest) (*GetTrainerAvailableHoursRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrainerAvailableHours not implemented")
}
func (UnimplementedBFFServer) MakeHourAvailable(context.Context, *MakeHourAvailableRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeHourAvailable not implemented")
}
func (UnimplementedBFFServer) MakeHourUnavailable(context.Context, *MakeHourUnavailableRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeHourUnavailable not implemented")
}
func (UnimplementedBFFServer) GetTraining(context.Context, *emptypb.Empty) (*v1.GetTrainingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTraining not implemented")
}
func (UnimplementedBFFServer) CreateTraining(context.Context, *v1.CreateTrainingRequest) (*v1.CreateTrainingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTraining not implemented")
}
func (UnimplementedBFFServer) CancelTraining(context.Context, *v1.CancelTrainingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTraining not implemented")
}
func (UnimplementedBFFServer) RescheduleTraining(context.Context, *v1.RescheduleTrainingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RescheduleTraining not implemented")
}
func (UnimplementedBFFServer) ApproveRescheduleTraining(context.Context, *v1.ApproveRescheduleTrainingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveRescheduleTraining not implemented")
}
func (UnimplementedBFFServer) RequestRescheduleTraining(context.Context, *v1.RequestRescheduleTrainingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestRescheduleTraining not implemented")
}
func (UnimplementedBFFServer) RejectRescheduleTraining(context.Context, *v1.RejectRescheduleTrainingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectRescheduleTraining not implemented")
}
func (UnimplementedBFFServer) mustEmbedUnimplementedBFFServer() {}

// UnsafeBFFServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BFFServer will
// result in compilation errors.
type UnsafeBFFServer interface {
	mustEmbedUnimplementedBFFServer()
}

func RegisterBFFServer(s grpc.ServiceRegistrar, srv BFFServer) {
	s.RegisterService(&BFF_ServiceDesc, srv)
}

func _BFF_IsHourAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsHourAvailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).IsHourAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_IsHourAvailable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).IsHourAvailable(ctx, req.(*IsHourAvailableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_ScheduleTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).ScheduleTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_ScheduleTraining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).ScheduleTraining(ctx, req.(*UpdateHourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_GetTrainerAvailableHours_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrainerAvailableHoursRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).GetTrainerAvailableHours(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_GetTrainerAvailableHours_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).GetTrainerAvailableHours(ctx, req.(*GetTrainerAvailableHoursRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_MakeHourAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeHourAvailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).MakeHourAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_MakeHourAvailable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).MakeHourAvailable(ctx, req.(*MakeHourAvailableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_MakeHourUnavailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeHourUnavailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).MakeHourUnavailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_MakeHourUnavailable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).MakeHourUnavailable(ctx, req.(*MakeHourUnavailableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_GetTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).GetTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_GetTraining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).GetTraining(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_CreateTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.CreateTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).CreateTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_CreateTraining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).CreateTraining(ctx, req.(*v1.CreateTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_CancelTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.CancelTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).CancelTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_CancelTraining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).CancelTraining(ctx, req.(*v1.CancelTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_RescheduleTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.RescheduleTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).RescheduleTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_RescheduleTraining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).RescheduleTraining(ctx, req.(*v1.RescheduleTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_ApproveRescheduleTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ApproveRescheduleTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).ApproveRescheduleTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_ApproveRescheduleTraining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).ApproveRescheduleTraining(ctx, req.(*v1.ApproveRescheduleTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_RequestRescheduleTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.RequestRescheduleTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).RequestRescheduleTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_RequestRescheduleTraining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).RequestRescheduleTraining(ctx, req.(*v1.RequestRescheduleTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_RejectRescheduleTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.RejectRescheduleTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).RejectRescheduleTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_RejectRescheduleTraining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).RejectRescheduleTraining(ctx, req.(*v1.RejectRescheduleTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BFF_ServiceDesc is the grpc.ServiceDesc for BFF service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BFF_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bff.BFF",
	HandlerType: (*BFFServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsHourAvailable",
			Handler:    _BFF_IsHourAvailable_Handler,
		},
		{
			MethodName: "ScheduleTraining",
			Handler:    _BFF_ScheduleTraining_Handler,
		},
		{
			MethodName: "GetTrainerAvailableHours",
			Handler:    _BFF_GetTrainerAvailableHours_Handler,
		},
		{
			MethodName: "MakeHourAvailable",
			Handler:    _BFF_MakeHourAvailable_Handler,
		},
		{
			MethodName: "MakeHourUnavailable",
			Handler:    _BFF_MakeHourUnavailable_Handler,
		},
		{
			MethodName: "GetTraining",
			Handler:    _BFF_GetTraining_Handler,
		},
		{
			MethodName: "CreateTraining",
			Handler:    _BFF_CreateTraining_Handler,
		},
		{
			MethodName: "CancelTraining",
			Handler:    _BFF_CancelTraining_Handler,
		},
		{
			MethodName: "RescheduleTraining",
			Handler:    _BFF_RescheduleTraining_Handler,
		},
		{
			MethodName: "ApproveRescheduleTraining",
			Handler:    _BFF_ApproveRescheduleTraining_Handler,
		},
		{
			MethodName: "RequestRescheduleTraining",
			Handler:    _BFF_RequestRescheduleTraining_Handler,
		},
		{
			MethodName: "RejectRescheduleTraining",
			Handler:    _BFF_RejectRescheduleTraining_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gomono-bff/api/v1/bff.proto",
}
