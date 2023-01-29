// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: migrator.proto

package __

import (
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

type Source int32

const (
	Source_MIGCTL    Source = 0
	Source_MIGRATORD Source = 1
)

// Enum value maps for Source.
var (
	Source_name = map[int32]string{
		0: "MIGCTL",
		1: "MIGRATORD",
	}
	Source_value = map[string]int32{
		"MIGCTL":    0,
		"MIGRATORD": 1,
	}
)

func (x Source) Enum() *Source {
	p := new(Source)
	*p = x
	return p
}

func (x Source) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Source) Descriptor() protoreflect.EnumDescriptor {
	return file_migrator_proto_enumTypes[0].Descriptor()
}

func (Source) Type() protoreflect.EnumType {
	return &file_migrator_proto_enumTypes[0]
}

func (x Source) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Source.Descriptor instead.
func (Source) EnumDescriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{0}
}

type InformStateChangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MigrationId   string  `protobuf:"bytes,1,opt,name=migration_id,json=migrationId,proto3" json:"migration_id,omitempty"`
	FinishedState string  `protobuf:"bytes,2,opt,name=finished_state,json=finishedState,proto3" json:"finished_state,omitempty"`
	NextState     string  `protobuf:"bytes,3,opt,name=next_state,json=nextState,proto3" json:"next_state,omitempty"`
	ErrorString   *string `protobuf:"bytes,4,opt,name=error_string,json=errorString,proto3,oneof" json:"error_string,omitempty"`
}

func (x *InformStateChangeRequest) Reset() {
	*x = InformStateChangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InformStateChangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InformStateChangeRequest) ProtoMessage() {}

func (x *InformStateChangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InformStateChangeRequest.ProtoReflect.Descriptor instead.
func (*InformStateChangeRequest) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{0}
}

func (x *InformStateChangeRequest) GetMigrationId() string {
	if x != nil {
		return x.MigrationId
	}
	return ""
}

func (x *InformStateChangeRequest) GetFinishedState() string {
	if x != nil {
		return x.FinishedState
	}
	return ""
}

func (x *InformStateChangeRequest) GetNextState() string {
	if x != nil {
		return x.NextState
	}
	return ""
}

func (x *InformStateChangeRequest) GetErrorString() string {
	if x != nil && x.ErrorString != nil {
		return *x.ErrorString
	}
	return ""
}

type InformStateChangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InformStateChangeResponse) Reset() {
	*x = InformStateChangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InformStateChangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InformStateChangeResponse) ProtoMessage() {}

func (x *InformStateChangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InformStateChangeResponse.ProtoReflect.Descriptor instead.
func (*InformStateChangeResponse) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{1}
}

type CreateMigrationJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerId            string `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ClientContainerRuntime string `protobuf:"bytes,2,opt,name=client_container_runtime,json=clientContainerRuntime,proto3" json:"client_container_runtime,omitempty"`
	ServerContainerRuntime string `protobuf:"bytes,3,opt,name=server_container_runtime,json=serverContainerRuntime,proto3" json:"server_container_runtime,omitempty"`
	ServerAddress          string `protobuf:"bytes,4,opt,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"`
	ServerPort             int32  `protobuf:"varint,5,opt,name=server_port,json=serverPort,proto3" json:"server_port,omitempty"`
	ServerKey              string `protobuf:"bytes,6,opt,name=server_key,json=serverKey,proto3" json:"server_key,omitempty"`
	ServerUser             string `protobuf:"bytes,7,opt,name=server_user,json=serverUser,proto3" json:"server_user,omitempty"`
	Method                 string `protobuf:"bytes,8,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *CreateMigrationJobRequest) Reset() {
	*x = CreateMigrationJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMigrationJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMigrationJobRequest) ProtoMessage() {}

func (x *CreateMigrationJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMigrationJobRequest.ProtoReflect.Descriptor instead.
func (*CreateMigrationJobRequest) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{2}
}

func (x *CreateMigrationJobRequest) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *CreateMigrationJobRequest) GetClientContainerRuntime() string {
	if x != nil {
		return x.ClientContainerRuntime
	}
	return ""
}

func (x *CreateMigrationJobRequest) GetServerContainerRuntime() string {
	if x != nil {
		return x.ServerContainerRuntime
	}
	return ""
}

func (x *CreateMigrationJobRequest) GetServerAddress() string {
	if x != nil {
		return x.ServerAddress
	}
	return ""
}

func (x *CreateMigrationJobRequest) GetServerPort() int32 {
	if x != nil {
		return x.ServerPort
	}
	return 0
}

func (x *CreateMigrationJobRequest) GetServerKey() string {
	if x != nil {
		return x.ServerKey
	}
	return ""
}

func (x *CreateMigrationJobRequest) GetServerUser() string {
	if x != nil {
		return x.ServerUser
	}
	return ""
}

func (x *CreateMigrationJobRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type CreateMigrationJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MigrationId     string `protobuf:"bytes,1,opt,name=migration_id,json=migrationId,proto3" json:"migration_id,omitempty"`
	CreatonUnixTime int64  `protobuf:"varint,2,opt,name=creaton_unix_time,json=creatonUnixTime,proto3" json:"creaton_unix_time,omitempty"`
}

func (x *CreateMigrationJobResponse) Reset() {
	*x = CreateMigrationJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMigrationJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMigrationJobResponse) ProtoMessage() {}

func (x *CreateMigrationJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMigrationJobResponse.ProtoReflect.Descriptor instead.
func (*CreateMigrationJobResponse) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{3}
}

func (x *CreateMigrationJobResponse) GetMigrationId() string {
	if x != nil {
		return x.MigrationId
	}
	return ""
}

func (x *CreateMigrationJobResponse) GetCreatonUnixTime() int64 {
	if x != nil {
		return x.CreatonUnixTime
	}
	return 0
}

type ShareMigrationJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerId            string `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ClientContainerRuntime string `protobuf:"bytes,2,opt,name=client_container_runtime,json=clientContainerRuntime,proto3" json:"client_container_runtime,omitempty"`
	ServerContainerRuntime string `protobuf:"bytes,3,opt,name=server_container_runtime,json=serverContainerRuntime,proto3" json:"server_container_runtime,omitempty"`
	ClientAddress          string `protobuf:"bytes,4,opt,name=client_address,json=clientAddress,proto3" json:"client_address,omitempty"`
	ClientPort             int32  `protobuf:"varint,5,opt,name=client_port,json=clientPort,proto3" json:"client_port,omitempty"`
	ServerKey              string `protobuf:"bytes,6,opt,name=server_key,json=serverKey,proto3" json:"server_key,omitempty"`
	ServerUser             string `protobuf:"bytes,7,opt,name=server_user,json=serverUser,proto3" json:"server_user,omitempty"`
	Method                 string `protobuf:"bytes,8,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *ShareMigrationJobRequest) Reset() {
	*x = ShareMigrationJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareMigrationJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareMigrationJobRequest) ProtoMessage() {}

func (x *ShareMigrationJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareMigrationJobRequest.ProtoReflect.Descriptor instead.
func (*ShareMigrationJobRequest) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{4}
}

func (x *ShareMigrationJobRequest) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *ShareMigrationJobRequest) GetClientContainerRuntime() string {
	if x != nil {
		return x.ClientContainerRuntime
	}
	return ""
}

func (x *ShareMigrationJobRequest) GetServerContainerRuntime() string {
	if x != nil {
		return x.ServerContainerRuntime
	}
	return ""
}

func (x *ShareMigrationJobRequest) GetClientAddress() string {
	if x != nil {
		return x.ClientAddress
	}
	return ""
}

func (x *ShareMigrationJobRequest) GetClientPort() int32 {
	if x != nil {
		return x.ClientPort
	}
	return 0
}

func (x *ShareMigrationJobRequest) GetServerKey() string {
	if x != nil {
		return x.ServerKey
	}
	return ""
}

func (x *ShareMigrationJobRequest) GetServerUser() string {
	if x != nil {
		return x.ServerUser
	}
	return ""
}

func (x *ShareMigrationJobRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type ShareMigrationJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MigrationId     string `protobuf:"bytes,1,opt,name=migration_id,json=migrationId,proto3" json:"migration_id,omitempty"`
	CreatonUnixTime int64  `protobuf:"varint,2,opt,name=creaton_unix_time,json=creatonUnixTime,proto3" json:"creaton_unix_time,omitempty"`
}

func (x *ShareMigrationJobResponse) Reset() {
	*x = ShareMigrationJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareMigrationJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareMigrationJobResponse) ProtoMessage() {}

func (x *ShareMigrationJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareMigrationJobResponse.ProtoReflect.Descriptor instead.
func (*ShareMigrationJobResponse) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{5}
}

func (x *ShareMigrationJobResponse) GetMigrationId() string {
	if x != nil {
		return x.MigrationId
	}
	return ""
}

func (x *ShareMigrationJobResponse) GetCreatonUnixTime() int64 {
	if x != nil {
		return x.CreatonUnixTime
	}
	return 0
}

type GetMigrationJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMigrationJobRequest) Reset() {
	*x = GetMigrationJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMigrationJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMigrationJobRequest) ProtoMessage() {}

func (x *GetMigrationJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMigrationJobRequest.ProtoReflect.Descriptor instead.
func (*GetMigrationJobRequest) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{6}
}

type GetMigrationJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs []*MigrationJob `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
}

func (x *GetMigrationJobResponse) Reset() {
	*x = GetMigrationJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMigrationJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMigrationJobResponse) ProtoMessage() {}

func (x *GetMigrationJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMigrationJobResponse.ProtoReflect.Descriptor instead.
func (*GetMigrationJobResponse) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{7}
}

func (x *GetMigrationJobResponse) GetJobs() []*MigrationJob {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type MigrationJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MigrationId     string `protobuf:"bytes,1,opt,name=migration_id,json=migrationId,proto3" json:"migration_id,omitempty"`
	ClientAddress   string `protobuf:"bytes,2,opt,name=client_address,json=clientAddress,proto3" json:"client_address,omitempty"`
	ServerAddress   string `protobuf:"bytes,3,opt,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"`
	ClientPort      int32  `protobuf:"varint,4,opt,name=client_port,json=clientPort,proto3" json:"client_port,omitempty"`
	ServerPort      int32  `protobuf:"varint,5,opt,name=server_port,json=serverPort,proto3" json:"server_port,omitempty"`
	CotninerId      string `protobuf:"bytes,6,opt,name=cotniner_id,json=cotninerId,proto3" json:"cotniner_id,omitempty"`
	MigrationStatus string `protobuf:"bytes,7,opt,name=migration_status,json=migrationStatus,proto3" json:"migration_status,omitempty"`
	ErrorString     string `protobuf:"bytes,8,opt,name=error_string,json=errorString,proto3" json:"error_string,omitempty"`
	Running         bool   `protobuf:"varint,9,opt,name=running,proto3" json:"running,omitempty"`
	CreationDate    string `protobuf:"bytes,10,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	Role            string `protobuf:"bytes,11,opt,name=role,proto3" json:"role,omitempty"`
	MigrationMethod string `protobuf:"bytes,12,opt,name=migration_method,json=migrationMethod,proto3" json:"migration_method,omitempty"`
}

func (x *MigrationJob) Reset() {
	*x = MigrationJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigrationJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrationJob) ProtoMessage() {}

func (x *MigrationJob) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrationJob.ProtoReflect.Descriptor instead.
func (*MigrationJob) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{8}
}

func (x *MigrationJob) GetMigrationId() string {
	if x != nil {
		return x.MigrationId
	}
	return ""
}

func (x *MigrationJob) GetClientAddress() string {
	if x != nil {
		return x.ClientAddress
	}
	return ""
}

func (x *MigrationJob) GetServerAddress() string {
	if x != nil {
		return x.ServerAddress
	}
	return ""
}

func (x *MigrationJob) GetClientPort() int32 {
	if x != nil {
		return x.ClientPort
	}
	return 0
}

func (x *MigrationJob) GetServerPort() int32 {
	if x != nil {
		return x.ServerPort
	}
	return 0
}

func (x *MigrationJob) GetCotninerId() string {
	if x != nil {
		return x.CotninerId
	}
	return ""
}

func (x *MigrationJob) GetMigrationStatus() string {
	if x != nil {
		return x.MigrationStatus
	}
	return ""
}

func (x *MigrationJob) GetErrorString() string {
	if x != nil {
		return x.ErrorString
	}
	return ""
}

func (x *MigrationJob) GetRunning() bool {
	if x != nil {
		return x.Running
	}
	return false
}

func (x *MigrationJob) GetCreationDate() string {
	if x != nil {
		return x.CreationDate
	}
	return ""
}

func (x *MigrationJob) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *MigrationJob) GetMigrationMethod() string {
	if x != nil {
		return x.MigrationMethod
	}
	return ""
}

var File_migrator_proto protoreflect.FileDescriptor

var file_migrator_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbc, 0x01, 0x0a, 0x18, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x78,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22,
	0x1b, 0x0a, 0x19, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xd2, 0x02, 0x0a,
	0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a,
	0x18, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x5f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x16, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x22, 0x6b, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x6e, 0x5f, 0x75, 0x6e,
	0x69, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xd1,
	0x02, 0x0a, 0x18, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38,
	0x0a, 0x18, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x5f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x16, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x22, 0x6a, 0x0a, 0x19, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x6e, 0x5f, 0x75, 0x6e,
	0x69, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x18,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62,
	0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0xae, 0x03, 0x0a, 0x0c, 0x4d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f,
	0x74, 0x6e, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x6f, 0x74, 0x6e, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x10,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2a, 0x23, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x49, 0x47, 0x43, 0x54, 0x4c, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x4d, 0x49, 0x47, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x44, 0x10, 0x01, 0x32, 0xbe, 0x02, 0x0a,
	0x0f, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x1a, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4a, 0x0a, 0x11, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x19, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x17,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4a, 0x0a, 0x11, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03, 0x5a,
	0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_migrator_proto_rawDescOnce sync.Once
	file_migrator_proto_rawDescData = file_migrator_proto_rawDesc
)

func file_migrator_proto_rawDescGZIP() []byte {
	file_migrator_proto_rawDescOnce.Do(func() {
		file_migrator_proto_rawDescData = protoimpl.X.CompressGZIP(file_migrator_proto_rawDescData)
	})
	return file_migrator_proto_rawDescData
}

var file_migrator_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_migrator_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_migrator_proto_goTypes = []interface{}{
	(Source)(0),                        // 0: Source
	(*InformStateChangeRequest)(nil),   // 1: InformStateChangeRequest
	(*InformStateChangeResponse)(nil),  // 2: InformStateChangeResponse
	(*CreateMigrationJobRequest)(nil),  // 3: CreateMigrationJobRequest
	(*CreateMigrationJobResponse)(nil), // 4: CreateMigrationJobResponse
	(*ShareMigrationJobRequest)(nil),   // 5: ShareMigrationJobRequest
	(*ShareMigrationJobResponse)(nil),  // 6: ShareMigrationJobResponse
	(*GetMigrationJobRequest)(nil),     // 7: GetMigrationJobRequest
	(*GetMigrationJobResponse)(nil),    // 8: GetMigrationJobResponse
	(*MigrationJob)(nil),               // 9: MigrationJob
}
var file_migrator_proto_depIdxs = []int32{
	9, // 0: GetMigrationJobResponse.jobs:type_name -> MigrationJob
	3, // 1: MigratorService.CreateMigrationJob:input_type -> CreateMigrationJobRequest
	5, // 2: MigratorService.ShareMigrationJob:input_type -> ShareMigrationJobRequest
	7, // 3: MigratorService.GetMigrationJob:input_type -> GetMigrationJobRequest
	1, // 4: MigratorService.InformStateChange:input_type -> InformStateChangeRequest
	4, // 5: MigratorService.CreateMigrationJob:output_type -> CreateMigrationJobResponse
	6, // 6: MigratorService.ShareMigrationJob:output_type -> ShareMigrationJobResponse
	8, // 7: MigratorService.GetMigrationJob:output_type -> GetMigrationJobResponse
	2, // 8: MigratorService.InformStateChange:output_type -> InformStateChangeResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_migrator_proto_init() }
func file_migrator_proto_init() {
	if File_migrator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_migrator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InformStateChangeRequest); i {
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
		file_migrator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InformStateChangeResponse); i {
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
		file_migrator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMigrationJobRequest); i {
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
		file_migrator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMigrationJobResponse); i {
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
		file_migrator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShareMigrationJobRequest); i {
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
		file_migrator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShareMigrationJobResponse); i {
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
		file_migrator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMigrationJobRequest); i {
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
		file_migrator_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMigrationJobResponse); i {
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
		file_migrator_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigrationJob); i {
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
	file_migrator_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_migrator_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_migrator_proto_goTypes,
		DependencyIndexes: file_migrator_proto_depIdxs,
		EnumInfos:         file_migrator_proto_enumTypes,
		MessageInfos:      file_migrator_proto_msgTypes,
	}.Build()
	File_migrator_proto = out.File
	file_migrator_proto_rawDesc = nil
	file_migrator_proto_goTypes = nil
	file_migrator_proto_depIdxs = nil
}
