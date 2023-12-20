// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v4.25.0
// source: trainer/v1/trainer.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationTrainerServiceGetTrainerAvailableHours = "/trainer.TrainerService/GetTrainerAvailableHours"
const OperationTrainerServiceMakeHourAvailable = "/trainer.TrainerService/MakeHourAvailable"
const OperationTrainerServiceMakeHourUnavailable = "/trainer.TrainerService/MakeHourUnavailable"

type TrainerServiceHTTPServer interface {
	GetTrainerAvailableHours(context.Context, *GetTrainerAvailableHoursRequest) (*GetTrainerAvailableHoursRespone, error)
	MakeHourAvailable(context.Context, *MakeHourAvailableRequest) (*emptypb.Empty, error)
	MakeHourUnavailable(context.Context, *MakeHourUnavailableRequest) (*emptypb.Empty, error)
}

func RegisterTrainerServiceHTTPServer(s *http.Server, srv TrainerServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/trainer/v1/getTrainerAvailableHours", _TrainerService_GetTrainerAvailableHours0_HTTP_Handler(srv))
	r.POST("/trainer/v1/makeHourAvailable", _TrainerService_MakeHourAvailable0_HTTP_Handler(srv))
	r.POST("/trainer/v1/makeHourUnavailable", _TrainerService_MakeHourUnavailable0_HTTP_Handler(srv))
}

func _TrainerService_GetTrainerAvailableHours0_HTTP_Handler(srv TrainerServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTrainerAvailableHoursRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTrainerServiceGetTrainerAvailableHours)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTrainerAvailableHours(ctx, req.(*GetTrainerAvailableHoursRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTrainerAvailableHoursRespone)
		return ctx.Result(200, reply)
	}
}

func _TrainerService_MakeHourAvailable0_HTTP_Handler(srv TrainerServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MakeHourAvailableRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTrainerServiceMakeHourAvailable)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MakeHourAvailable(ctx, req.(*MakeHourAvailableRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _TrainerService_MakeHourUnavailable0_HTTP_Handler(srv TrainerServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MakeHourUnavailableRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTrainerServiceMakeHourUnavailable)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MakeHourUnavailable(ctx, req.(*MakeHourUnavailableRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type TrainerServiceHTTPClient interface {
	GetTrainerAvailableHours(ctx context.Context, req *GetTrainerAvailableHoursRequest, opts ...http.CallOption) (rsp *GetTrainerAvailableHoursRespone, err error)
	MakeHourAvailable(ctx context.Context, req *MakeHourAvailableRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	MakeHourUnavailable(ctx context.Context, req *MakeHourUnavailableRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type TrainerServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewTrainerServiceHTTPClient(client *http.Client) TrainerServiceHTTPClient {
	return &TrainerServiceHTTPClientImpl{client}
}

func (c *TrainerServiceHTTPClientImpl) GetTrainerAvailableHours(ctx context.Context, in *GetTrainerAvailableHoursRequest, opts ...http.CallOption) (*GetTrainerAvailableHoursRespone, error) {
	var out GetTrainerAvailableHoursRespone
	pattern := "/trainer/v1/getTrainerAvailableHours"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTrainerServiceGetTrainerAvailableHours))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TrainerServiceHTTPClientImpl) MakeHourAvailable(ctx context.Context, in *MakeHourAvailableRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/trainer/v1/makeHourAvailable"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTrainerServiceMakeHourAvailable))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TrainerServiceHTTPClientImpl) MakeHourUnavailable(ctx context.Context, in *MakeHourUnavailableRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/trainer/v1/makeHourUnavailable"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTrainerServiceMakeHourUnavailable))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
