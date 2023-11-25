// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: training/v1/error_reason.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TrainingErrorReason int32

const (
	TrainingErrorReason_TRAINER_UNSPECIFIED TrainingErrorReason = 0
	TrainingErrorReason_USER_NOT_FOUND      TrainingErrorReason = 1
)

// Enum value maps for TrainingErrorReason.
var (
	TrainingErrorReason_name = map[int32]string{
		0: "TRAINER_UNSPECIFIED",
		1: "USER_NOT_FOUND",
	}
	TrainingErrorReason_value = map[string]int32{
		"TRAINER_UNSPECIFIED": 0,
		"USER_NOT_FOUND":      1,
	}
)

func (x TrainingErrorReason) Enum() *TrainingErrorReason {
	p := new(TrainingErrorReason)
	*p = x
	return p
}

func (x TrainingErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrainingErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_training_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (TrainingErrorReason) Type() protoreflect.EnumType {
	return &file_training_v1_error_reason_proto_enumTypes[0]
}

func (x TrainingErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrainingErrorReason.Descriptor instead.
func (TrainingErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_training_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_training_v1_error_reason_proto protoreflect.FileDescriptor

var file_training_v1_error_reason_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0x48, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x52, 0x41,
	0x49, 0x4e, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x32, 0x5a, 0x30,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x69, 0x71, 0x69,
	0x6e, 0x66, 0x65, 0x6e, 0x67, 0x31, 0x2f, 0x67, 0x6f, 0x4d, 0x6f, 0x6e, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_training_v1_error_reason_proto_rawDescOnce sync.Once
	file_training_v1_error_reason_proto_rawDescData = file_training_v1_error_reason_proto_rawDesc
)

func file_training_v1_error_reason_proto_rawDescGZIP() []byte {
	file_training_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_training_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_training_v1_error_reason_proto_rawDescData)
	})
	return file_training_v1_error_reason_proto_rawDescData
}

var file_training_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_training_v1_error_reason_proto_goTypes = []interface{}{
	(TrainingErrorReason)(0), // 0: training.v1.TrainingErrorReason
}
var file_training_v1_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_training_v1_error_reason_proto_init() }
func file_training_v1_error_reason_proto_init() {
	if File_training_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_training_v1_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_training_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_training_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_training_v1_error_reason_proto_enumTypes,
	}.Build()
	File_training_v1_error_reason_proto = out.File
	file_training_v1_error_reason_proto_rawDesc = nil
	file_training_v1_error_reason_proto_goTypes = nil
	file_training_v1_error_reason_proto_depIdxs = nil
}
