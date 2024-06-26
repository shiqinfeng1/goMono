// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: common/conf/public.proto

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

type Public struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Log             *Log             `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
	Trace           *Trace           `protobuf:"bytes,2,opt,name=trace,proto3" json:"trace,omitempty"`
	Registry        *Registry        `protobuf:"bytes,3,opt,name=registry,proto3" json:"registry,omitempty"`
	Discovery       *Discovery       `protobuf:"bytes,4,opt,name=discovery,proto3" json:"discovery,omitempty"`
	Adapter         *Adapter         `protobuf:"bytes,5,opt,name=adapter,proto3" json:"adapter,omitempty"`
	GatewayRegister *GatewayRegister `protobuf:"bytes,6,opt,name=gateway_register,json=gatewayRegister,proto3" json:"gateway_register,omitempty"`
}

func (x *Public) Reset() {
	*x = Public{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_public_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Public) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Public) ProtoMessage() {}

func (x *Public) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_public_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Public.ProtoReflect.Descriptor instead.
func (*Public) Descriptor() ([]byte, []int) {
	return file_common_conf_public_proto_rawDescGZIP(), []int{0}
}

func (x *Public) GetLog() *Log {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *Public) GetTrace() *Trace {
	if x != nil {
		return x.Trace
	}
	return nil
}

func (x *Public) GetRegistry() *Registry {
	if x != nil {
		return x.Registry
	}
	return nil
}

func (x *Public) GetDiscovery() *Discovery {
	if x != nil {
		return x.Discovery
	}
	return nil
}

func (x *Public) GetAdapter() *Adapter {
	if x != nil {
		return x.Adapter
	}
	return nil
}

func (x *Public) GetGatewayRegister() *GatewayRegister {
	if x != nil {
		return x.GatewayRegister
	}
	return nil
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level   string   `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	Monitor *Monitor `protobuf:"bytes,2,opt,name=monitor,proto3" json:"monitor,omitempty"`
	File    *File    `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_public_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_public_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_common_conf_public_proto_rawDescGZIP(), []int{1}
}

func (x *Log) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Log) GetMonitor() *Monitor {
	if x != nil {
		return x.Monitor
	}
	return nil
}

func (x *Log) GetFile() *File {
	if x != nil {
		return x.File
	}
	return nil
}

type Monitor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *Monitor) Reset() {
	*x = Monitor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_public_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Monitor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Monitor) ProtoMessage() {}

func (x *Monitor) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_public_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Monitor.ProtoReflect.Descriptor instead.
func (*Monitor) Descriptor() ([]byte, []int) {
	return file_common_conf_public_proto_rawDescGZIP(), []int{2}
}

func (x *Monitor) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxSize    int32 `protobuf:"varint,1,opt,name=max_size,json=maxSize,proto3" json:"max_size,omitempty"`
	MaxBackups int32 `protobuf:"varint,2,opt,name=max_backups,json=maxBackups,proto3" json:"max_backups,omitempty"`
	MaxAge     int32 `protobuf:"varint,3,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
	Compress   bool  `protobuf:"varint,4,opt,name=compress,proto3" json:"compress,omitempty"`
	LocalTime  bool  `protobuf:"varint,5,opt,name=local_time,json=localTime,proto3" json:"local_time,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_public_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_public_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_common_conf_public_proto_rawDescGZIP(), []int{3}
}

func (x *File) GetMaxSize() int32 {
	if x != nil {
		return x.MaxSize
	}
	return 0
}

func (x *File) GetMaxBackups() int32 {
	if x != nil {
		return x.MaxBackups
	}
	return 0
}

func (x *File) GetMaxAge() int32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

func (x *File) GetCompress() bool {
	if x != nil {
		return x.Compress
	}
	return false
}

func (x *File) GetLocalTime() bool {
	if x != nil {
		return x.LocalTime
	}
	return false
}

type Trace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *Trace) Reset() {
	*x = Trace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_public_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trace) ProtoMessage() {}

func (x *Trace) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_public_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trace.ProtoReflect.Descriptor instead.
func (*Trace) Descriptor() ([]byte, []int) {
	return file_common_conf_public_proto_rawDescGZIP(), []int{4}
}

func (x *Trace) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type Registry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addrs []string `protobuf:"bytes,1,rep,name=addrs,proto3" json:"addrs,omitempty"`
}

func (x *Registry) Reset() {
	*x = Registry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_public_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry) ProtoMessage() {}

func (x *Registry) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_public_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry.ProtoReflect.Descriptor instead.
func (*Registry) Descriptor() ([]byte, []int) {
	return file_common_conf_public_proto_rawDescGZIP(), []int{5}
}

func (x *Registry) GetAddrs() []string {
	if x != nil {
		return x.Addrs
	}
	return nil
}

type Discovery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints []string `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *Discovery) Reset() {
	*x = Discovery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_public_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Discovery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Discovery) ProtoMessage() {}

func (x *Discovery) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_public_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Discovery.ProtoReflect.Descriptor instead.
func (*Discovery) Descriptor() ([]byte, []int) {
	return file_common_conf_public_proto_rawDescGZIP(), []int{6}
}

func (x *Discovery) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

type Adapter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database *Adapter_Database `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Redis    *Adapter_Redis    `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis,omitempty"`
}

func (x *Adapter) Reset() {
	*x = Adapter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_public_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Adapter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Adapter) ProtoMessage() {}

func (x *Adapter) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_public_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Adapter.ProtoReflect.Descriptor instead.
func (*Adapter) Descriptor() ([]byte, []int) {
	return file_common_conf_public_proto_rawDescGZIP(), []int{7}
}

func (x *Adapter) GetDatabase() *Adapter_Database {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *Adapter) GetRedis() *Adapter_Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

type GatewayRegister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints []string `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *GatewayRegister) Reset() {
	*x = GatewayRegister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_public_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayRegister) ProtoMessage() {}

func (x *GatewayRegister) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_public_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayRegister.ProtoReflect.Descriptor instead.
func (*GatewayRegister) Descriptor() ([]byte, []int) {
	return file_common_conf_public_proto_rawDescGZIP(), []int{8}
}

func (x *GatewayRegister) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

type Adapter_Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *Adapter_Database) Reset() {
	*x = Adapter_Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_public_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Adapter_Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Adapter_Database) ProtoMessage() {}

func (x *Adapter_Database) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_public_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Adapter_Database.ProtoReflect.Descriptor instead.
func (*Adapter_Database) Descriptor() ([]byte, []int) {
	return file_common_conf_public_proto_rawDescGZIP(), []int{7, 0}
}

func (x *Adapter_Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Adapter_Database) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type Adapter_Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr         string               `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	ReadTimeout  *durationpb.Duration `protobuf:"bytes,2,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout *durationpb.Duration `protobuf:"bytes,3,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
}

func (x *Adapter_Redis) Reset() {
	*x = Adapter_Redis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_public_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Adapter_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Adapter_Redis) ProtoMessage() {}

func (x *Adapter_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_public_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Adapter_Redis.ProtoReflect.Descriptor instead.
func (*Adapter_Redis) Descriptor() ([]byte, []int) {
	return file_common_conf_public_proto_rawDescGZIP(), []int{7, 1}
}

func (x *Adapter_Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Adapter_Redis) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Adapter_Redis) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

var File_common_conf_public_proto protoreflect.FileDescriptor

var file_common_conf_public_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xee, 0x02, 0x0a, 0x06, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x2b, 0x0a, 0x03, 0x6c,
	0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x31, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x08, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x3d, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x09, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x37, 0x0a, 0x07, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x41,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x07, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12,
	0x50, 0x0a, 0x10, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x0f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x22, 0x84, 0x01, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x37, 0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52,
	0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x25, 0x0a, 0x07, 0x4d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22,
	0x96, 0x01, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x20, 0x0a,
	0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x64,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x64, 0x72, 0x73, 0x22,
	0x29, 0x0a, 0x09, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0xe0, 0x02, 0x0a, 0x07, 0x41,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x72, 0x65,
	0x64, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x52, 0x05,
	0x72, 0x65, 0x64, 0x69, 0x73, 0x1a, 0x3a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x1a, 0x99, 0x01, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12,
	0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3e, 0x0a,
	0x0d, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x2f, 0x0a,
	0x0f, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x42, 0x34,
	0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x69,
	0x71, 0x69, 0x6e, 0x66, 0x65, 0x6e, 0x67, 0x31, 0x2f, 0x67, 0x6f, 0x4d, 0x6f, 0x6e, 0x6f, 0x2f,
	0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b,
	0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_conf_public_proto_rawDescOnce sync.Once
	file_common_conf_public_proto_rawDescData = file_common_conf_public_proto_rawDesc
)

func file_common_conf_public_proto_rawDescGZIP() []byte {
	file_common_conf_public_proto_rawDescOnce.Do(func() {
		file_common_conf_public_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_conf_public_proto_rawDescData)
	})
	return file_common_conf_public_proto_rawDescData
}

var file_common_conf_public_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_common_conf_public_proto_goTypes = []interface{}{
	(*Public)(nil),              // 0: kratos.config.public.Public
	(*Log)(nil),                 // 1: kratos.config.public.Log
	(*Monitor)(nil),             // 2: kratos.config.public.Monitor
	(*File)(nil),                // 3: kratos.config.public.File
	(*Trace)(nil),               // 4: kratos.config.public.Trace
	(*Registry)(nil),            // 5: kratos.config.public.Registry
	(*Discovery)(nil),           // 6: kratos.config.public.Discovery
	(*Adapter)(nil),             // 7: kratos.config.public.Adapter
	(*GatewayRegister)(nil),     // 8: kratos.config.public.GatewayRegister
	(*Adapter_Database)(nil),    // 9: kratos.config.public.Adapter.Database
	(*Adapter_Redis)(nil),       // 10: kratos.config.public.Adapter.Redis
	(*durationpb.Duration)(nil), // 11: google.protobuf.Duration
}
var file_common_conf_public_proto_depIdxs = []int32{
	1,  // 0: kratos.config.public.Public.log:type_name -> kratos.config.public.Log
	4,  // 1: kratos.config.public.Public.trace:type_name -> kratos.config.public.Trace
	5,  // 2: kratos.config.public.Public.registry:type_name -> kratos.config.public.Registry
	6,  // 3: kratos.config.public.Public.discovery:type_name -> kratos.config.public.Discovery
	7,  // 4: kratos.config.public.Public.adapter:type_name -> kratos.config.public.Adapter
	8,  // 5: kratos.config.public.Public.gateway_register:type_name -> kratos.config.public.GatewayRegister
	2,  // 6: kratos.config.public.Log.monitor:type_name -> kratos.config.public.Monitor
	3,  // 7: kratos.config.public.Log.file:type_name -> kratos.config.public.File
	9,  // 8: kratos.config.public.Adapter.database:type_name -> kratos.config.public.Adapter.Database
	10, // 9: kratos.config.public.Adapter.redis:type_name -> kratos.config.public.Adapter.Redis
	11, // 10: kratos.config.public.Adapter.Redis.read_timeout:type_name -> google.protobuf.Duration
	11, // 11: kratos.config.public.Adapter.Redis.write_timeout:type_name -> google.protobuf.Duration
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_common_conf_public_proto_init() }
func file_common_conf_public_proto_init() {
	if File_common_conf_public_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_conf_public_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Public); i {
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
		file_common_conf_public_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_common_conf_public_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Monitor); i {
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
		file_common_conf_public_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_common_conf_public_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trace); i {
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
		file_common_conf_public_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry); i {
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
		file_common_conf_public_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Discovery); i {
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
		file_common_conf_public_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Adapter); i {
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
		file_common_conf_public_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayRegister); i {
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
		file_common_conf_public_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Adapter_Database); i {
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
		file_common_conf_public_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Adapter_Redis); i {
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
			RawDescriptor: file_common_conf_public_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_conf_public_proto_goTypes,
		DependencyIndexes: file_common_conf_public_proto_depIdxs,
		MessageInfos:      file_common_conf_public_proto_msgTypes,
	}.Build()
	File_common_conf_public_proto = out.File
	file_common_conf_public_proto_rawDesc = nil
	file_common_conf_public_proto_goTypes = nil
	file_common_conf_public_proto_depIdxs = nil
}
