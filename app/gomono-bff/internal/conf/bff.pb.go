// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.0
// source: gomono-bff/internal/conf/bff.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http     *HTTP     `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Auth     *Auth     `protobuf:"bytes,2,opt,name=auth,proto3" json:"auth,omitempty"`
	Register *Register `protobuf:"bytes,3,opt,name=register,proto3" json:"register,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gomono_bff_internal_conf_bff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_bff_internal_conf_bff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_gomono_bff_internal_conf_bff_proto_rawDescGZIP(), []int{0}
}

func (x *Server) GetHttp() *HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *Server) GetRegister() *Register {
	if x != nil {
		return x.Register
	}
	return nil
}

type HTTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr    string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *HTTP) Reset() {
	*x = HTTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gomono_bff_internal_conf_bff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTP) ProtoMessage() {}

func (x *HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_bff_internal_conf_bff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTP.ProtoReflect.Descriptor instead.
func (*HTTP) Descriptor() ([]byte, []int) {
	return file_gomono_bff_internal_conf_bff_proto_rawDescGZIP(), []int{1}
}

func (x *HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceKey string `protobuf:"bytes,1,opt,name=service_key,json=serviceKey,proto3" json:"service_key,omitempty"`
	ApiKey     string `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gomono_bff_internal_conf_bff_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_bff_internal_conf_bff_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_gomono_bff_internal_conf_bff_proto_rawDescGZIP(), []int{2}
}

func (x *Auth) GetServiceKey() string {
	if x != nil {
		return x.ServiceKey
	}
	return ""
}

func (x *Auth) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

type Register struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints []string `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *Register) Reset() {
	*x = Register{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gomono_bff_internal_conf_bff_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Register) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Register) ProtoMessage() {}

func (x *Register) ProtoReflect() protoreflect.Message {
	mi := &file_gomono_bff_internal_conf_bff_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Register.ProtoReflect.Descriptor instead.
func (*Register) Descriptor() ([]byte, []int) {
	return file_gomono_bff_internal_conf_bff_proto_rawDescGZIP(), []int{3}
}

func (x *Register) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

var File_gomono_bff_internal_conf_bff_proto protoreflect.FileDescriptor

var file_gomono_bff_internal_conf_bff_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x6f, 0x2d, 0x62, 0x66, 0x66, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x62, 0x66, 0x66, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x62, 0x66, 0x66, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x2b, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x62, 0x66, 0x66, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12,
	0x2b, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x62, 0x66,
	0x66, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x37, 0x0a, 0x08,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x62,
	0x66, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x22, 0x69, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x12, 0x18, 0x0a,
	0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x22, 0x40, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x22, 0x28, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x42, 0x41, 0x5a, 0x3f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x69, 0x71, 0x69,
	0x6e, 0x66, 0x65, 0x6e, 0x67, 0x31, 0x2f, 0x67, 0x6f, 0x4d, 0x6f, 0x6e, 0x6f, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x6f, 0x2d, 0x62, 0x66, 0x66, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gomono_bff_internal_conf_bff_proto_rawDescOnce sync.Once
	file_gomono_bff_internal_conf_bff_proto_rawDescData = file_gomono_bff_internal_conf_bff_proto_rawDesc
)

func file_gomono_bff_internal_conf_bff_proto_rawDescGZIP() []byte {
	file_gomono_bff_internal_conf_bff_proto_rawDescOnce.Do(func() {
		file_gomono_bff_internal_conf_bff_proto_rawDescData = protoimpl.X.CompressGZIP(file_gomono_bff_internal_conf_bff_proto_rawDescData)
	})
	return file_gomono_bff_internal_conf_bff_proto_rawDescData
}

var file_gomono_bff_internal_conf_bff_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gomono_bff_internal_conf_bff_proto_goTypes = []interface{}{
	(*Server)(nil),              // 0: kratos.config.bff.Server
	(*HTTP)(nil),                // 1: kratos.config.bff.HTTP
	(*Auth)(nil),                // 2: kratos.config.bff.Auth
	(*Register)(nil),            // 3: kratos.config.bff.Register
	(*durationpb.Duration)(nil), // 4: google.protobuf.Duration
}
var file_gomono_bff_internal_conf_bff_proto_depIdxs = []int32{
	1, // 0: kratos.config.bff.Server.http:type_name -> kratos.config.bff.HTTP
	2, // 1: kratos.config.bff.Server.auth:type_name -> kratos.config.bff.Auth
	3, // 2: kratos.config.bff.Server.register:type_name -> kratos.config.bff.Register
	4, // 3: kratos.config.bff.HTTP.timeout:type_name -> google.protobuf.Duration
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_gomono_bff_internal_conf_bff_proto_init() }
func file_gomono_bff_internal_conf_bff_proto_init() {
	if File_gomono_bff_internal_conf_bff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gomono_bff_internal_conf_bff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_gomono_bff_internal_conf_bff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTP); i {
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
		file_gomono_bff_internal_conf_bff_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
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
		file_gomono_bff_internal_conf_bff_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Register); i {
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
			RawDescriptor: file_gomono_bff_internal_conf_bff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gomono_bff_internal_conf_bff_proto_goTypes,
		DependencyIndexes: file_gomono_bff_internal_conf_bff_proto_depIdxs,
		MessageInfos:      file_gomono_bff_internal_conf_bff_proto_msgTypes,
	}.Build()
	File_gomono_bff_internal_conf_bff_proto = out.File
	file_gomono_bff_internal_conf_bff_proto_rawDesc = nil
	file_gomono_bff_internal_conf_bff_proto_goTypes = nil
	file_gomono_bff_internal_conf_bff_proto_depIdxs = nil
}
