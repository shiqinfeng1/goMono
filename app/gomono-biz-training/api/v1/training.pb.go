// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.0
// source: gomono-biz-training/api/v1/training.proto

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

type GetTrainingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Training []*GetTrainingResponse_Training `protobuf:"bytes,1,rep,name=training,proto3" json:"training,omitempty"`
}

func (x *GetTrainingResponse) Reset() {
	*x = GetTrainingResponse{}
	mi := &file_gomono_biz_training_api_v1_training_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTrainingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrainingResponse) ProtoMessage() {}

func (x *GetTrainingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_biz_training_api_v1_training_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrainingResponse.ProtoReflect.Descriptor instead.
func (*GetTrainingResponse) Descriptor() ([]byte, []int) {
	return file_gomono_biz_training_api_v1_training_proto_rawDescGZIP(), []int{0}
}

func (x *GetTrainingResponse) GetTraining() []*GetTrainingResponse_Training {
	if x != nil {
		return x.Training
	}
	return nil
}

type CreateTrainingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notes string                 `protobuf:"bytes,1,opt,name=notes,proto3" json:"notes,omitempty"`
	Time  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *CreateTrainingRequest) Reset() {
	*x = CreateTrainingRequest{}
	mi := &file_gomono_biz_training_api_v1_training_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTrainingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTrainingRequest) ProtoMessage() {}

func (x *CreateTrainingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_biz_training_api_v1_training_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTrainingRequest.ProtoReflect.Descriptor instead.
func (*CreateTrainingRequest) Descriptor() ([]byte, []int) {
	return file_gomono_biz_training_api_v1_training_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTrainingRequest) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *CreateTrainingRequest) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type CreateTrainingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrainingUuid string `protobuf:"bytes,1,opt,name=training_uuid,json=trainingUuid,proto3" json:"training_uuid,omitempty"`
}

func (x *CreateTrainingResponse) Reset() {
	*x = CreateTrainingResponse{}
	mi := &file_gomono_biz_training_api_v1_training_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTrainingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTrainingResponse) ProtoMessage() {}

func (x *CreateTrainingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_biz_training_api_v1_training_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTrainingResponse.ProtoReflect.Descriptor instead.
func (*CreateTrainingResponse) Descriptor() ([]byte, []int) {
	return file_gomono_biz_training_api_v1_training_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTrainingResponse) GetTrainingUuid() string {
	if x != nil {
		return x.TrainingUuid
	}
	return ""
}

type CancelTrainingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrainingUuid string `protobuf:"bytes,1,opt,name=training_uuid,json=trainingUuid,proto3" json:"training_uuid,omitempty"`
}

func (x *CancelTrainingRequest) Reset() {
	*x = CancelTrainingRequest{}
	mi := &file_gomono_biz_training_api_v1_training_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelTrainingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelTrainingRequest) ProtoMessage() {}

func (x *CancelTrainingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_biz_training_api_v1_training_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelTrainingRequest.ProtoReflect.Descriptor instead.
func (*CancelTrainingRequest) Descriptor() ([]byte, []int) {
	return file_gomono_biz_training_api_v1_training_proto_rawDescGZIP(), []int{3}
}

func (x *CancelTrainingRequest) GetTrainingUuid() string {
	if x != nil {
		return x.TrainingUuid
	}
	return ""
}

type RescheduleTrainingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrainingUuid string                 `protobuf:"bytes,1,opt,name=training_uuid,json=trainingUuid,proto3" json:"training_uuid,omitempty"`
	Notes        string                 `protobuf:"bytes,2,opt,name=notes,proto3" json:"notes,omitempty"`
	Time         *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *RescheduleTrainingRequest) Reset() {
	*x = RescheduleTrainingRequest{}
	mi := &file_gomono_biz_training_api_v1_training_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RescheduleTrainingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RescheduleTrainingRequest) ProtoMessage() {}

func (x *RescheduleTrainingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_biz_training_api_v1_training_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RescheduleTrainingRequest.ProtoReflect.Descriptor instead.
func (*RescheduleTrainingRequest) Descriptor() ([]byte, []int) {
	return file_gomono_biz_training_api_v1_training_proto_rawDescGZIP(), []int{4}
}

func (x *RescheduleTrainingRequest) GetTrainingUuid() string {
	if x != nil {
		return x.TrainingUuid
	}
	return ""
}

func (x *RescheduleTrainingRequest) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *RescheduleTrainingRequest) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type RequestRescheduleTrainingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrainingUuid string                 `protobuf:"bytes,1,opt,name=training_uuid,json=trainingUuid,proto3" json:"training_uuid,omitempty"`
	Notes        string                 `protobuf:"bytes,2,opt,name=notes,proto3" json:"notes,omitempty"`
	Time         *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *RequestRescheduleTrainingRequest) Reset() {
	*x = RequestRescheduleTrainingRequest{}
	mi := &file_gomono_biz_training_api_v1_training_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestRescheduleTrainingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestRescheduleTrainingRequest) ProtoMessage() {}

func (x *RequestRescheduleTrainingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_biz_training_api_v1_training_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestRescheduleTrainingRequest.ProtoReflect.Descriptor instead.
func (*RequestRescheduleTrainingRequest) Descriptor() ([]byte, []int) {
	return file_gomono_biz_training_api_v1_training_proto_rawDescGZIP(), []int{5}
}

func (x *RequestRescheduleTrainingRequest) GetTrainingUuid() string {
	if x != nil {
		return x.TrainingUuid
	}
	return ""
}

func (x *RequestRescheduleTrainingRequest) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *RequestRescheduleTrainingRequest) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type ApproveRescheduleTrainingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrainingUuid string `protobuf:"bytes,1,opt,name=training_uuid,json=trainingUuid,proto3" json:"training_uuid,omitempty"`
}

func (x *ApproveRescheduleTrainingRequest) Reset() {
	*x = ApproveRescheduleTrainingRequest{}
	mi := &file_gomono_biz_training_api_v1_training_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApproveRescheduleTrainingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveRescheduleTrainingRequest) ProtoMessage() {}

func (x *ApproveRescheduleTrainingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_biz_training_api_v1_training_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveRescheduleTrainingRequest.ProtoReflect.Descriptor instead.
func (*ApproveRescheduleTrainingRequest) Descriptor() ([]byte, []int) {
	return file_gomono_biz_training_api_v1_training_proto_rawDescGZIP(), []int{6}
}

func (x *ApproveRescheduleTrainingRequest) GetTrainingUuid() string {
	if x != nil {
		return x.TrainingUuid
	}
	return ""
}

type RejectRescheduleTrainingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrainingUuid string `protobuf:"bytes,1,opt,name=training_uuid,json=trainingUuid,proto3" json:"training_uuid,omitempty"`
}

func (x *RejectRescheduleTrainingRequest) Reset() {
	*x = RejectRescheduleTrainingRequest{}
	mi := &file_gomono_biz_training_api_v1_training_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RejectRescheduleTrainingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectRescheduleTrainingRequest) ProtoMessage() {}

func (x *RejectRescheduleTrainingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_biz_training_api_v1_training_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectRescheduleTrainingRequest.ProtoReflect.Descriptor instead.
func (*RejectRescheduleTrainingRequest) Descriptor() ([]byte, []int) {
	return file_gomono_biz_training_api_v1_training_proto_rawDescGZIP(), []int{7}
}

func (x *RejectRescheduleTrainingRequest) GetTrainingUuid() string {
	if x != nil {
		return x.TrainingUuid
	}
	return ""
}

type GetTrainingResponse_Training struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CanBeCancelled     bool                   `protobuf:"varint,1,opt,name=can_be_cancelled,json=canBeCancelled,proto3" json:"can_be_cancelled,omitempty"`
	MoveProposedBy     string                 `protobuf:"bytes,2,opt,name=move_proposed_by,json=moveProposedBy,proto3" json:"move_proposed_by,omitempty"`
	MoveRequiresAccept bool                   `protobuf:"varint,3,opt,name=move_requires_accept,json=moveRequiresAccept,proto3" json:"move_requires_accept,omitempty"`
	Notes              string                 `protobuf:"bytes,4,opt,name=notes,proto3" json:"notes,omitempty"`
	ProposedTime       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=proposed_time,json=proposedTime,proto3" json:"proposed_time,omitempty"`
	Time               *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=time,proto3" json:"time,omitempty"`
	User               string                 `protobuf:"bytes,7,opt,name=user,proto3" json:"user,omitempty"`
	UserUuid           string                 `protobuf:"bytes,8,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Uuid               string                 `protobuf:"bytes,9,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetTrainingResponse_Training) Reset() {
	*x = GetTrainingResponse_Training{}
	mi := &file_gomono_biz_training_api_v1_training_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTrainingResponse_Training) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrainingResponse_Training) ProtoMessage() {}

func (x *GetTrainingResponse_Training) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_biz_training_api_v1_training_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrainingResponse_Training.ProtoReflect.Descriptor instead.
func (*GetTrainingResponse_Training) Descriptor() ([]byte, []int) {
	return file_gomono_biz_training_api_v1_training_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GetTrainingResponse_Training) GetCanBeCancelled() bool {
	if x != nil {
		return x.CanBeCancelled
	}
	return false
}

func (x *GetTrainingResponse_Training) GetMoveProposedBy() string {
	if x != nil {
		return x.MoveProposedBy
	}
	return ""
}

func (x *GetTrainingResponse_Training) GetMoveRequiresAccept() bool {
	if x != nil {
		return x.MoveRequiresAccept
	}
	return false
}

func (x *GetTrainingResponse_Training) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *GetTrainingResponse_Training) GetProposedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ProposedTime
	}
	return nil
}

func (x *GetTrainingResponse_Training) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *GetTrainingResponse_Training) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *GetTrainingResponse_Training) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *GetTrainingResponse_Training) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

var File_gomono_biz_training_api_v1_training_proto protoreflect.FileDescriptor

var file_gomono_biz_training_api_v1_training_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x6f, 0x2d, 0x62, 0x69, 0x7a, 0x2d, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x03, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x08, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x1a,
	0xdc, 0x02, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x10,
	0x63, 0x61, 0x6e, 0x5f, 0x62, 0x65, 0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x63, 0x61, 0x6e, 0x42, 0x65, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x42, 0x79,
	0x12, 0x30, 0x0a, 0x14, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x73, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12,
	0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x5d,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x2e, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x3d, 0x0a,
	0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x55, 0x75, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x15,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x55, 0x75, 0x69, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x19, 0x52,
	0x65, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x55, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f,
	0x74, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x20, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x55, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f,
	0x74, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x20, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x55, 0x75, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x1f,
	0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x55, 0x75, 0x69, 0x64, 0x32, 0xf9, 0x04, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1d, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x55, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x12, 0x1f, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x2e, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x19, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x72,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x2a, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x19,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x2a, 0x2e, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x5f, 0x0a, 0x18, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x68, 0x69, 0x71, 0x69, 0x6e, 0x66, 0x65, 0x6e, 0x67, 0x31, 0x2f, 0x67, 0x6f, 0x4d, 0x6f, 0x6e,
	0x6f, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x6f, 0x2d, 0x62, 0x69, 0x7a,
	0x2d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gomono_biz_training_api_v1_training_proto_rawDescOnce sync.Once
	file_gomono_biz_training_api_v1_training_proto_rawDescData = file_gomono_biz_training_api_v1_training_proto_rawDesc
)

func file_gomono_biz_training_api_v1_training_proto_rawDescGZIP() []byte {
	file_gomono_biz_training_api_v1_training_proto_rawDescOnce.Do(func() {
		file_gomono_biz_training_api_v1_training_proto_rawDescData = protoimpl.X.CompressGZIP(file_gomono_biz_training_api_v1_training_proto_rawDescData)
	})
	return file_gomono_biz_training_api_v1_training_proto_rawDescData
}

var file_gomono_biz_training_api_v1_training_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_gomono_biz_training_api_v1_training_proto_goTypes = []any{
	(*GetTrainingResponse)(nil),              // 0: training.GetTrainingResponse
	(*CreateTrainingRequest)(nil),            // 1: training.CreateTrainingRequest
	(*CreateTrainingResponse)(nil),           // 2: training.CreateTrainingResponse
	(*CancelTrainingRequest)(nil),            // 3: training.CancelTrainingRequest
	(*RescheduleTrainingRequest)(nil),        // 4: training.RescheduleTrainingRequest
	(*RequestRescheduleTrainingRequest)(nil), // 5: training.RequestRescheduleTrainingRequest
	(*ApproveRescheduleTrainingRequest)(nil), // 6: training.ApproveRescheduleTrainingRequest
	(*RejectRescheduleTrainingRequest)(nil),  // 7: training.RejectRescheduleTrainingRequest
	(*GetTrainingResponse_Training)(nil),     // 8: training.GetTrainingResponse.Training
	(*timestamppb.Timestamp)(nil),            // 9: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                    // 10: google.protobuf.Empty
}
var file_gomono_biz_training_api_v1_training_proto_depIdxs = []int32{
	8,  // 0: training.GetTrainingResponse.training:type_name -> training.GetTrainingResponse.Training
	9,  // 1: training.CreateTrainingRequest.time:type_name -> google.protobuf.Timestamp
	9,  // 2: training.RescheduleTrainingRequest.time:type_name -> google.protobuf.Timestamp
	9,  // 3: training.RequestRescheduleTrainingRequest.time:type_name -> google.protobuf.Timestamp
	9,  // 4: training.GetTrainingResponse.Training.proposed_time:type_name -> google.protobuf.Timestamp
	9,  // 5: training.GetTrainingResponse.Training.time:type_name -> google.protobuf.Timestamp
	10, // 6: training.TrainingService.GetTraining:input_type -> google.protobuf.Empty
	1,  // 7: training.TrainingService.CreateTraining:input_type -> training.CreateTrainingRequest
	3,  // 8: training.TrainingService.CancelTraining:input_type -> training.CancelTrainingRequest
	4,  // 9: training.TrainingService.RescheduleTraining:input_type -> training.RescheduleTrainingRequest
	6,  // 10: training.TrainingService.ApproveRescheduleTraining:input_type -> training.ApproveRescheduleTrainingRequest
	5,  // 11: training.TrainingService.RequestRescheduleTraining:input_type -> training.RequestRescheduleTrainingRequest
	7,  // 12: training.TrainingService.RejectRescheduleTraining:input_type -> training.RejectRescheduleTrainingRequest
	0,  // 13: training.TrainingService.GetTraining:output_type -> training.GetTrainingResponse
	2,  // 14: training.TrainingService.CreateTraining:output_type -> training.CreateTrainingResponse
	10, // 15: training.TrainingService.CancelTraining:output_type -> google.protobuf.Empty
	10, // 16: training.TrainingService.RescheduleTraining:output_type -> google.protobuf.Empty
	10, // 17: training.TrainingService.ApproveRescheduleTraining:output_type -> google.protobuf.Empty
	10, // 18: training.TrainingService.RequestRescheduleTraining:output_type -> google.protobuf.Empty
	10, // 19: training.TrainingService.RejectRescheduleTraining:output_type -> google.protobuf.Empty
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_gomono_biz_training_api_v1_training_proto_init() }
func file_gomono_biz_training_api_v1_training_proto_init() {
	if File_gomono_biz_training_api_v1_training_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gomono_biz_training_api_v1_training_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gomono_biz_training_api_v1_training_proto_goTypes,
		DependencyIndexes: file_gomono_biz_training_api_v1_training_proto_depIdxs,
		MessageInfos:      file_gomono_biz_training_api_v1_training_proto_msgTypes,
	}.Build()
	File_gomono_biz_training_api_v1_training_proto = out.File
	file_gomono_biz_training_api_v1_training_proto_rawDesc = nil
	file_gomono_biz_training_api_v1_training_proto_goTypes = nil
	file_gomono_biz_training_api_v1_training_proto_depIdxs = nil
}
