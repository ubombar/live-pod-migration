// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
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

type CreateMigrationJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerAddress string `protobuf:"bytes,1,opt,name=peer_address,json=peerAddress,proto3" json:"peer_address,omitempty"`
	PeerPort    int32  `protobuf:"varint,2,opt,name=peer_port,json=peerPort,proto3" json:"peer_port,omitempty"`
	ContainerId string `protobuf:"bytes,3,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
}

func (x *CreateMigrationJobRequest) Reset() {
	*x = CreateMigrationJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMigrationJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMigrationJobRequest) ProtoMessage() {}

func (x *CreateMigrationJobRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateMigrationJobRequest.ProtoReflect.Descriptor instead.
func (*CreateMigrationJobRequest) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMigrationJobRequest) GetPeerAddress() string {
	if x != nil {
		return x.PeerAddress
	}
	return ""
}

func (x *CreateMigrationJobRequest) GetPeerPort() int32 {
	if x != nil {
		return x.PeerPort
	}
	return 0
}

func (x *CreateMigrationJobRequest) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

type CreateMigrationJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accepted bool `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
}

func (x *CreateMigrationJobResponse) Reset() {
	*x = CreateMigrationJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMigrationJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMigrationJobResponse) ProtoMessage() {}

func (x *CreateMigrationJobResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateMigrationJobResponse.ProtoReflect.Descriptor instead.
func (*CreateMigrationJobResponse) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMigrationJobResponse) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

type ShareMigrationJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerAddress    string `protobuf:"bytes,1,opt,name=peer_address,json=peerAddress,proto3" json:"peer_address,omitempty"`
	PeerPort       int32  `protobuf:"varint,2,opt,name=peer_port,json=peerPort,proto3" json:"peer_port,omitempty"`
	ContainerId    string `protobuf:"bytes,3,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ContainerImage string `protobuf:"bytes,4,opt,name=container_image,json=containerImage,proto3" json:"container_image,omitempty"`
	ContainerName  string `protobuf:"bytes,5,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
}

func (x *ShareMigrationJobRequest) Reset() {
	*x = ShareMigrationJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareMigrationJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareMigrationJobRequest) ProtoMessage() {}

func (x *ShareMigrationJobRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ShareMigrationJobRequest.ProtoReflect.Descriptor instead.
func (*ShareMigrationJobRequest) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{2}
}

func (x *ShareMigrationJobRequest) GetPeerAddress() string {
	if x != nil {
		return x.PeerAddress
	}
	return ""
}

func (x *ShareMigrationJobRequest) GetPeerPort() int32 {
	if x != nil {
		return x.PeerPort
	}
	return 0
}

func (x *ShareMigrationJobRequest) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *ShareMigrationJobRequest) GetContainerImage() string {
	if x != nil {
		return x.ContainerImage
	}
	return ""
}

func (x *ShareMigrationJobRequest) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

type ShareMigrationJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MigrationId     string `protobuf:"bytes,1,opt,name=migration_id,json=migrationId,proto3" json:"migration_id,omitempty"`
	Accepted        bool   `protobuf:"varint,2,opt,name=accepted,proto3" json:"accepted,omitempty"`
	CreatonUnixTime int64  `protobuf:"varint,3,opt,name=creaton_unix_time,json=creatonUnixTime,proto3" json:"creaton_unix_time,omitempty"`
}

func (x *ShareMigrationJobResponse) Reset() {
	*x = ShareMigrationJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareMigrationJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareMigrationJobResponse) ProtoMessage() {}

func (x *ShareMigrationJobResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ShareMigrationJobResponse.ProtoReflect.Descriptor instead.
func (*ShareMigrationJobResponse) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{3}
}

func (x *ShareMigrationJobResponse) GetMigrationId() string {
	if x != nil {
		return x.MigrationId
	}
	return ""
}

func (x *ShareMigrationJobResponse) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

func (x *ShareMigrationJobResponse) GetCreatonUnixTime() int64 {
	if x != nil {
		return x.CreatonUnixTime
	}
	return 0
}

type UpdateMigrationStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MigrationId string  `protobuf:"bytes,1,opt,name=migration_id,json=migrationId,proto3" json:"migration_id,omitempty"`
	NewStatus   string  `protobuf:"bytes,2,opt,name=new_status,json=newStatus,proto3" json:"new_status,omitempty"`
	Description *string `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
}

func (x *UpdateMigrationStatusRequest) Reset() {
	*x = UpdateMigrationStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMigrationStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMigrationStatusRequest) ProtoMessage() {}

func (x *UpdateMigrationStatusRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdateMigrationStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateMigrationStatusRequest) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateMigrationStatusRequest) GetMigrationId() string {
	if x != nil {
		return x.MigrationId
	}
	return ""
}

func (x *UpdateMigrationStatusRequest) GetNewStatus() string {
	if x != nil {
		return x.NewStatus
	}
	return ""
}

func (x *UpdateMigrationStatusRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type UpdateMigrationStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateMigrationStatusResponse) Reset() {
	*x = UpdateMigrationStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMigrationStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMigrationStatusResponse) ProtoMessage() {}

func (x *UpdateMigrationStatusResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdateMigrationStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateMigrationStatusResponse) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{5}
}

type SendViaSCPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SftpCertificate string `protobuf:"bytes,1,opt,name=sftp_certificate,json=sftpCertificate,proto3" json:"sftp_certificate,omitempty"`
	FileSize        int64  `protobuf:"varint,2,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
}

func (x *SendViaSCPRequest) Reset() {
	*x = SendViaSCPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendViaSCPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendViaSCPRequest) ProtoMessage() {}

func (x *SendViaSCPRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SendViaSCPRequest.ProtoReflect.Descriptor instead.
func (*SendViaSCPRequest) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{6}
}

func (x *SendViaSCPRequest) GetSftpCertificate() string {
	if x != nil {
		return x.SftpCertificate
	}
	return ""
}

func (x *SendViaSCPRequest) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

type SendViaSCPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendViaSCPResponse) Reset() {
	*x = SendViaSCPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendViaSCPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendViaSCPResponse) ProtoMessage() {}

func (x *SendViaSCPResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SendViaSCPResponse.ProtoReflect.Descriptor instead.
func (*SendViaSCPResponse) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{7}
}

var File_migrator_proto protoreflect.FileDescriptor

var file_migrator_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x7e, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x65, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x38, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x22, 0xcd, 0x01, 0x0a, 0x18, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x65, 0x65, 0x72, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x65, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65,
	0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x65, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x19, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x6e, 0x5f, 0x75, 0x6e, 0x69, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x78, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x77,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1f, 0x0a,
	0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b,
	0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x69, 0x61, 0x53, 0x43, 0x50, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x66, 0x74, 0x70, 0x5f, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73,
	0x66, 0x74, 0x70, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x53,
	0x65, 0x6e, 0x64, 0x56, 0x69, 0x61, 0x53, 0x43, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xbb, 0x02, 0x0a, 0x0f, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x1a, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x11, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x19, 0x2e, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x56, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64,
	0x56, 0x69, 0x61, 0x53, 0x43, 0x50, 0x12, 0x12, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x69, 0x61,
	0x53, 0x43, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x56, 0x69, 0x61, 0x53, 0x43, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_migrator_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_migrator_proto_goTypes = []interface{}{
	(*CreateMigrationJobRequest)(nil),     // 0: CreateMigrationJobRequest
	(*CreateMigrationJobResponse)(nil),    // 1: CreateMigrationJobResponse
	(*ShareMigrationJobRequest)(nil),      // 2: ShareMigrationJobRequest
	(*ShareMigrationJobResponse)(nil),     // 3: ShareMigrationJobResponse
	(*UpdateMigrationStatusRequest)(nil),  // 4: UpdateMigrationStatusRequest
	(*UpdateMigrationStatusResponse)(nil), // 5: UpdateMigrationStatusResponse
	(*SendViaSCPRequest)(nil),             // 6: SendViaSCPRequest
	(*SendViaSCPResponse)(nil),            // 7: SendViaSCPResponse
}
var file_migrator_proto_depIdxs = []int32{
	0, // 0: MigratorService.CreateMigrationJob:input_type -> CreateMigrationJobRequest
	2, // 1: MigratorService.ShareMigrationJob:input_type -> ShareMigrationJobRequest
	4, // 2: MigratorService.UpdateMigrationStatus:input_type -> UpdateMigrationStatusRequest
	6, // 3: MigratorService.SendViaSCP:input_type -> SendViaSCPRequest
	1, // 4: MigratorService.CreateMigrationJob:output_type -> CreateMigrationJobResponse
	3, // 5: MigratorService.ShareMigrationJob:output_type -> ShareMigrationJobResponse
	5, // 6: MigratorService.UpdateMigrationStatus:output_type -> UpdateMigrationStatusResponse
	7, // 7: MigratorService.SendViaSCP:output_type -> SendViaSCPResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_migrator_proto_init() }
func file_migrator_proto_init() {
	if File_migrator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_migrator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_migrator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_migrator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_migrator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_migrator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMigrationStatusRequest); i {
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
			switch v := v.(*UpdateMigrationStatusResponse); i {
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
			switch v := v.(*SendViaSCPRequest); i {
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
			switch v := v.(*SendViaSCPResponse); i {
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
	file_migrator_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_migrator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_migrator_proto_goTypes,
		DependencyIndexes: file_migrator_proto_depIdxs,
		MessageInfos:      file_migrator_proto_msgTypes,
	}.Build()
	File_migrator_proto = out.File
	file_migrator_proto_rawDesc = nil
	file_migrator_proto_goTypes = nil
	file_migrator_proto_depIdxs = nil
}
