// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.0
// source: gomono-biz-trainer/api/v1/trainer.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IsHourAvailableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *IsHourAvailableRequest) Reset() {
	*x = IsHourAvailableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsHourAvailableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsHourAvailableRequest) ProtoMessage() {}

func (x *IsHourAvailableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsHourAvailableRequest.ProtoReflect.Descriptor instead.
func (*IsHourAvailableRequest) Descriptor() ([]byte, []int) {
	return file_gomono_biz_trainer_api_v1_trainer_proto_rawDescGZIP(), []int{0}
}

func (x *IsHourAvailableRequest) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type IsHourAvailableResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAvailable bool `protobuf:"varint,1,opt,name=is_available,json=isAvailable,proto3" json:"is_available,omitempty"`
}

func (x *IsHourAvailableResponse) Reset() {
	*x = IsHourAvailableResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsHourAvailableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsHourAvailableResponse) ProtoMessage() {}

func (x *IsHourAvailableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsHourAvailableResponse.ProtoReflect.Descriptor instead.
func (*IsHourAvailableResponse) Descriptor() ([]byte, []int) {
	return file_gomono_biz_trainer_api_v1_trainer_proto_rawDescGZIP(), []int{1}
}

func (x *IsHourAvailableResponse) GetIsAvailable() bool {
	if x != nil {
		return x.IsAvailable
	}
	return false
}

type UpdateHourRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *UpdateHourRequest) Reset() {
	*x = UpdateHourRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHourRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHourRequest) ProtoMessage() {}

func (x *UpdateHourRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHourRequest.ProtoReflect.Descriptor instead.
func (*UpdateHourRequest) Descriptor() ([]byte, []int) {
	return file_gomono_biz_trainer_api_v1_trainer_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateHourRequest) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type GetTrainerAvailableHoursRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DateFrom *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date_from,json=dateFrom,proto3" json:"date_from,omitempty"`
	DateTo   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date_to,json=dateTo,proto3" json:"date_to,omitempty"`
}

func (x *GetTrainerAvailableHoursRequest) Reset() {
	*x = GetTrainerAvailableHoursRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrainerAvailableHoursRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrainerAvailableHoursRequest) ProtoMessage() {}

func (x *GetTrainerAvailableHoursRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrainerAvailableHoursRequest.ProtoReflect.Descriptor instead.
func (*GetTrainerAvailableHoursRequest) Descriptor() ([]byte, []int) {
	return file_gomono_biz_trainer_api_v1_trainer_proto_rawDescGZIP(), []int{3}
}

func (x *GetTrainerAvailableHoursRequest) GetDateFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.DateFrom
	}
	return nil
}

func (x *GetTrainerAvailableHoursRequest) GetDateTo() *timestamppb.Timestamp {
	if x != nil {
		return x.DateTo
	}
	return nil
}

type GetTrainerAvailableHoursRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dates []*GetTrainerAvailableHoursRespone_Date `protobuf:"bytes,1,rep,name=dates,proto3" json:"dates,omitempty"`
}

func (x *GetTrainerAvailableHoursRespone) Reset() {
	*x = GetTrainerAvailableHoursRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrainerAvailableHoursRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrainerAvailableHoursRespone) ProtoMessage() {}

func (x *GetTrainerAvailableHoursRespone) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrainerAvailableHoursRespone.ProtoReflect.Descriptor instead.
func (*GetTrainerAvailableHoursRespone) Descriptor() ([]byte, []int) {
	return file_gomono_biz_trainer_api_v1_trainer_proto_rawDescGZIP(), []int{4}
}

func (x *GetTrainerAvailableHoursRespone) GetDates() []*GetTrainerAvailableHoursRespone_Date {
	if x != nil {
		return x.Dates
	}
	return nil
}

type MakeHourAvailableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time []*timestamppb.Timestamp `protobuf:"bytes,1,rep,name=time,proto3" json:"time,omitempty"`
}

func (x *MakeHourAvailableRequest) Reset() {
	*x = MakeHourAvailableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeHourAvailableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeHourAvailableRequest) ProtoMessage() {}

func (x *MakeHourAvailableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeHourAvailableRequest.ProtoReflect.Descriptor instead.
func (*MakeHourAvailableRequest) Descriptor() ([]byte, []int) {
	return file_gomono_biz_trainer_api_v1_trainer_proto_rawDescGZIP(), []int{5}
}

func (x *MakeHourAvailableRequest) GetTime() []*timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type MakeHourUnavailableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time []*timestamppb.Timestamp `protobuf:"bytes,1,rep,name=time,proto3" json:"time,omitempty"`
}

func (x *MakeHourUnavailableRequest) Reset() {
	*x = MakeHourUnavailableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeHourUnavailableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeHourUnavailableRequest) ProtoMessage() {}

func (x *MakeHourUnavailableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeHourUnavailableRequest.ProtoReflect.Descriptor instead.
func (*MakeHourUnavailableRequest) Descriptor() ([]byte, []int) {
	return file_gomono_biz_trainer_api_v1_trainer_proto_rawDescGZIP(), []int{6}
}

func (x *MakeHourUnavailableRequest) GetTime() []*timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type GetTrainerAvailableHoursRespone_Hour struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Available            bool                   `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
	HasTrainingScheduled bool                   `protobuf:"varint,2,opt,name=has_training_scheduled,json=hasTrainingScheduled,proto3" json:"has_training_scheduled,omitempty"`
	Hour                 *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=hour,proto3" json:"hour,omitempty"`
}

func (x *GetTrainerAvailableHoursRespone_Hour) Reset() {
	*x = GetTrainerAvailableHoursRespone_Hour{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrainerAvailableHoursRespone_Hour) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrainerAvailableHoursRespone_Hour) ProtoMessage() {}

func (x *GetTrainerAvailableHoursRespone_Hour) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrainerAvailableHoursRespone_Hour.ProtoReflect.Descriptor instead.
func (*GetTrainerAvailableHoursRespone_Hour) Descriptor() ([]byte, []int) {
	return file_gomono_biz_trainer_api_v1_trainer_proto_rawDescGZIP(), []int{4, 0}
}

func (x *GetTrainerAvailableHoursRespone_Hour) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

func (x *GetTrainerAvailableHoursRespone_Hour) GetHasTrainingScheduled() bool {
	if x != nil {
		return x.HasTrainingScheduled
	}
	return false
}

func (x *GetTrainerAvailableHoursRespone_Hour) GetHour() *timestamppb.Timestamp {
	if x != nil {
		return x.Hour
	}
	return nil
}

type GetTrainerAvailableHoursRespone_Date struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date         *timestamppb.Timestamp                  `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	HasFreeHours bool                                    `protobuf:"varint,2,opt,name=Has_free_hours,json=HasFreeHours,proto3" json:"Has_free_hours,omitempty"`
	Hours        []*GetTrainerAvailableHoursRespone_Hour `protobuf:"bytes,3,rep,name=hours,proto3" json:"hours,omitempty"`
}

func (x *GetTrainerAvailableHoursRespone_Date) Reset() {
	*x = GetTrainerAvailableHoursRespone_Date{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrainerAvailableHoursRespone_Date) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrainerAvailableHoursRespone_Date) ProtoMessage() {}

func (x *GetTrainerAvailableHoursRespone_Date) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrainerAvailableHoursRespone_Date.ProtoReflect.Descriptor instead.
func (*GetTrainerAvailableHoursRespone_Date) Descriptor() ([]byte, []int) {
	return file_gomono_biz_trainer_api_v1_trainer_proto_rawDescGZIP(), []int{4, 1}
}

func (x *GetTrainerAvailableHoursRespone_Date) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *GetTrainerAvailableHoursRespone_Date) GetHasFreeHours() bool {
	if x != nil {
		return x.HasFreeHours
	}
	return false
}

func (x *GetTrainerAvailableHoursRespone_Date) GetHours() []*GetTrainerAvailableHoursRespone_Hour {
	if x != nil {
		return x.Hours
	}
	return nil
}

var File_gomono_biz_trainer_api_v1_trainer_proto protoreflect.FileDescriptor

var file_gomono_biz_trainer_api_v1_trainer_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x6f, 0x2d, 0x62, 0x69, 0x7a, 0x2d, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x48, 0x0a, 0x16, 0x49, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x17, 0x49, 0x73,
	0x48, 0x6f, 0x75, 0x72, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x43, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x8f, 0x01,
	0x0a, 0x1f, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x37, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x33, 0x0a, 0x07, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x22,
	0x97, 0x03, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x65, 0x12, 0x43, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x48, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x2e, 0x44, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x64, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x8a, 0x01, 0x0a, 0x04, 0x48, 0x6f, 0x75,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x34, 0x0a, 0x16, 0x68, 0x61, 0x73, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x14, 0x68, 0x61, 0x73, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x68, 0x6f, 0x75, 0x72, 0x1a, 0xa1, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2e,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x24,
	0x0a, 0x0e, 0x48, 0x61, 0x73, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x48, 0x61, 0x73, 0x46, 0x72, 0x65, 0x65, 0x48,
	0x6f, 0x75, 0x72, 0x73, 0x12, 0x43, 0x0a, 0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x2e, 0x48, 0x6f,
	0x75, 0x72, 0x52, 0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x22, 0x4a, 0x0a, 0x18, 0x4d, 0x61, 0x6b,
	0x65, 0x48, 0x6f, 0x75, 0x72, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x1a, 0x4d, 0x61, 0x6b, 0x65, 0x48, 0x6f, 0x75,
	0x72, 0x55, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x32, 0x94, 0x04, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0f, 0x49, 0x73, 0x48, 0x6f, 0x75, 0x72,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x2e, 0x49, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x2e, 0x49, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48,
	0x0a, 0x10, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x12, 0x1a, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0e, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x70, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x28, 0x2e, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65,
	0x22, 0x00, 0x12, 0x50, 0x0a, 0x11, 0x4d, 0x61, 0x6b, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x13, 0x4d, 0x61, 0x6b, 0x65, 0x48, 0x6f, 0x75, 0x72,
	0x55, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x23, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x55, 0x6e,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x69, 0x71, 0x69, 0x6e, 0x66,
	0x65, 0x6e, 0x67, 0x31, 0x2f, 0x67, 0x6f, 0x4d, 0x6f, 0x6e, 0x6f, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x6f, 0x2d, 0x62, 0x69, 0x7a, 0x2d, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gomono_biz_trainer_api_v1_trainer_proto_rawDescOnce sync.Once
	file_gomono_biz_trainer_api_v1_trainer_proto_rawDescData = file_gomono_biz_trainer_api_v1_trainer_proto_rawDesc
)

func file_gomono_biz_trainer_api_v1_trainer_proto_rawDescGZIP() []byte {
	file_gomono_biz_trainer_api_v1_trainer_proto_rawDescOnce.Do(func() {
		file_gomono_biz_trainer_api_v1_trainer_proto_rawDescData = protoimpl.X.CompressGZIP(file_gomono_biz_trainer_api_v1_trainer_proto_rawDescData)
	})
	return file_gomono_biz_trainer_api_v1_trainer_proto_rawDescData
}

var file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_gomono_biz_trainer_api_v1_trainer_proto_goTypes = []interface{}{
	(*IsHourAvailableRequest)(nil),               // 0: trainer.IsHourAvailableRequest
	(*IsHourAvailableResponse)(nil),              // 1: trainer.IsHourAvailableResponse
	(*UpdateHourRequest)(nil),                    // 2: trainer.UpdateHourRequest
	(*GetTrainerAvailableHoursRequest)(nil),      // 3: trainer.GetTrainerAvailableHoursRequest
	(*GetTrainerAvailableHoursRespone)(nil),      // 4: trainer.GetTrainerAvailableHoursRespone
	(*MakeHourAvailableRequest)(nil),             // 5: trainer.MakeHourAvailableRequest
	(*MakeHourUnavailableRequest)(nil),           // 6: trainer.MakeHourUnavailableRequest
	(*GetTrainerAvailableHoursRespone_Hour)(nil), // 7: trainer.GetTrainerAvailableHoursRespone.Hour
	(*GetTrainerAvailableHoursRespone_Date)(nil), // 8: trainer.GetTrainerAvailableHoursRespone.Date
	(*timestamppb.Timestamp)(nil),                // 9: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                        // 10: google.protobuf.Empty
}
var file_gomono_biz_trainer_api_v1_trainer_proto_depIdxs = []int32{
	9,  // 0: trainer.IsHourAvailableRequest.time:type_name -> google.protobuf.Timestamp
	9,  // 1: trainer.UpdateHourRequest.time:type_name -> google.protobuf.Timestamp
	9,  // 2: trainer.GetTrainerAvailableHoursRequest.date_from:type_name -> google.protobuf.Timestamp
	9,  // 3: trainer.GetTrainerAvailableHoursRequest.date_to:type_name -> google.protobuf.Timestamp
	8,  // 4: trainer.GetTrainerAvailableHoursRespone.dates:type_name -> trainer.GetTrainerAvailableHoursRespone.Date
	9,  // 5: trainer.MakeHourAvailableRequest.time:type_name -> google.protobuf.Timestamp
	9,  // 6: trainer.MakeHourUnavailableRequest.time:type_name -> google.protobuf.Timestamp
	9,  // 7: trainer.GetTrainerAvailableHoursRespone.Hour.hour:type_name -> google.protobuf.Timestamp
	9,  // 8: trainer.GetTrainerAvailableHoursRespone.Date.date:type_name -> google.protobuf.Timestamp
	7,  // 9: trainer.GetTrainerAvailableHoursRespone.Date.hours:type_name -> trainer.GetTrainerAvailableHoursRespone.Hour
	0,  // 10: trainer.TrainerService.IsHourAvailable:input_type -> trainer.IsHourAvailableRequest
	2,  // 11: trainer.TrainerService.ScheduleTraining:input_type -> trainer.UpdateHourRequest
	2,  // 12: trainer.TrainerService.CancelTraining:input_type -> trainer.UpdateHourRequest
	3,  // 13: trainer.TrainerService.GetTrainerAvailableHours:input_type -> trainer.GetTrainerAvailableHoursRequest
	5,  // 14: trainer.TrainerService.MakeHourAvailable:input_type -> trainer.MakeHourAvailableRequest
	6,  // 15: trainer.TrainerService.MakeHourUnavailable:input_type -> trainer.MakeHourUnavailableRequest
	1,  // 16: trainer.TrainerService.IsHourAvailable:output_type -> trainer.IsHourAvailableResponse
	10, // 17: trainer.TrainerService.ScheduleTraining:output_type -> google.protobuf.Empty
	10, // 18: trainer.TrainerService.CancelTraining:output_type -> google.protobuf.Empty
	4,  // 19: trainer.TrainerService.GetTrainerAvailableHours:output_type -> trainer.GetTrainerAvailableHoursRespone
	10, // 20: trainer.TrainerService.MakeHourAvailable:output_type -> google.protobuf.Empty
	10, // 21: trainer.TrainerService.MakeHourUnavailable:output_type -> google.protobuf.Empty
	16, // [16:22] is the sub-list for method output_type
	10, // [10:16] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_gomono_biz_trainer_api_v1_trainer_proto_init() }
func file_gomono_biz_trainer_api_v1_trainer_proto_init() {
	if File_gomono_biz_trainer_api_v1_trainer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsHourAvailableRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsHourAvailableResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHourRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrainerAvailableHoursRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrainerAvailableHoursRespone); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeHourAvailableRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeHourUnavailableRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrainerAvailableHoursRespone_Hour); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrainerAvailableHoursRespone_Date); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gomono_biz_trainer_api_v1_trainer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gomono_biz_trainer_api_v1_trainer_proto_goTypes,
		DependencyIndexes: file_gomono_biz_trainer_api_v1_trainer_proto_depIdxs,
		MessageInfos:      file_gomono_biz_trainer_api_v1_trainer_proto_msgTypes,
	}.Build()
	File_gomono_biz_trainer_api_v1_trainer_proto = out.File
	file_gomono_biz_trainer_api_v1_trainer_proto_rawDesc = nil
	file_gomono_biz_trainer_api_v1_trainer_proto_goTypes = nil
	file_gomono_biz_trainer_api_v1_trainer_proto_depIdxs = nil
}
