// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: gomono-biz-trainer/api/v1/error_reason.proto

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

type TrainerErrorReason int32

const (
	TrainerErrorReason_USER_NOT_FOUND           TrainerErrorReason = 0
	TrainerErrorReason_USER_AUTH_FAIL           TrainerErrorReason = 1
	TrainerErrorReason_INCORRECT_INPUT          TrainerErrorReason = 2
	TrainerErrorReason_QUERY_FAIL               TrainerErrorReason = 3
	TrainerErrorReason_CANCEL_TRAINING_FAIL     TrainerErrorReason = 4
	TrainerErrorReason_UPDATE_AVAILABILITY_FAIL TrainerErrorReason = 5
)

// Enum value maps for TrainerErrorReason.
var (
	TrainerErrorReason_name = map[int32]string{
		0: "USER_NOT_FOUND",
		1: "USER_AUTH_FAIL",
		2: "INCORRECT_INPUT",
		3: "QUERY_FAIL",
		4: "CANCEL_TRAINING_FAIL",
		5: "UPDATE_AVAILABILITY_FAIL",
	}
	TrainerErrorReason_value = map[string]int32{
		"USER_NOT_FOUND":           0,
		"USER_AUTH_FAIL":           1,
		"INCORRECT_INPUT":          2,
		"QUERY_FAIL":               3,
		"CANCEL_TRAINING_FAIL":     4,
		"UPDATE_AVAILABILITY_FAIL": 5,
	}
)

func (x TrainerErrorReason) Enum() *TrainerErrorReason {
	p := new(TrainerErrorReason)
	*p = x
	return p
}

func (x TrainerErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrainerErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_gomono_biz_trainer_api_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (TrainerErrorReason) Type() protoreflect.EnumType {
	return &file_gomono_biz_trainer_api_v1_error_reason_proto_enumTypes[0]
}

func (x TrainerErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrainerErrorReason.Descriptor instead.
func (TrainerErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_gomono_biz_trainer_api_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_gomono_biz_trainer_api_v1_error_reason_proto protoreflect.FileDescriptor

var file_gomono_biz_trainer_api_v1_error_reason_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x6f, 0x2d, 0x62, 0x69, 0x7a, 0x2d, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a,
	0xc3, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03,
	0x12, 0x18, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x49, 0x4e,
	0x43, 0x4f, 0x52, 0x52, 0x45, 0x43, 0x54, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x10, 0x02, 0x1a,
	0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x14, 0x0a, 0x0a, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x1e, 0x0a, 0x14, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x22, 0x0a, 0x18, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x49, 0x4c, 0x49,
	0x54, 0x59, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x05, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x1a,
	0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x69, 0x71, 0x69, 0x6e, 0x66, 0x65, 0x6e, 0x67, 0x31, 0x2f,
	0x67, 0x6f, 0x4d, 0x6f, 0x6e, 0x6f, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x6f, 0x6d, 0x6f, 0x6e,
	0x6f, 0x2d, 0x62, 0x69, 0x7a, 0x2d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gomono_biz_trainer_api_v1_error_reason_proto_rawDescOnce sync.Once
	file_gomono_biz_trainer_api_v1_error_reason_proto_rawDescData = file_gomono_biz_trainer_api_v1_error_reason_proto_rawDesc
)

func file_gomono_biz_trainer_api_v1_error_reason_proto_rawDescGZIP() []byte {
	file_gomono_biz_trainer_api_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_gomono_biz_trainer_api_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_gomono_biz_trainer_api_v1_error_reason_proto_rawDescData)
	})
	return file_gomono_biz_trainer_api_v1_error_reason_proto_rawDescData
}

var file_gomono_biz_trainer_api_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gomono_biz_trainer_api_v1_error_reason_proto_goTypes = []interface{}{
	(TrainerErrorReason)(0), // 0: trainer.v1.TrainerErrorReason
}
var file_gomono_biz_trainer_api_v1_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gomono_biz_trainer_api_v1_error_reason_proto_init() }
func file_gomono_biz_trainer_api_v1_error_reason_proto_init() {
	if File_gomono_biz_trainer_api_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gomono_biz_trainer_api_v1_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gomono_biz_trainer_api_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_gomono_biz_trainer_api_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_gomono_biz_trainer_api_v1_error_reason_proto_enumTypes,
	}.Build()
	File_gomono_biz_trainer_api_v1_error_reason_proto = out.File
	file_gomono_biz_trainer_api_v1_error_reason_proto_rawDesc = nil
	file_gomono_biz_trainer_api_v1_error_reason_proto_goTypes = nil
	file_gomono_biz_trainer_api_v1_error_reason_proto_depIdxs = nil
}