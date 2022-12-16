// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: migrator.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MigratorServiceClient is the client API for MigratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MigratorServiceClient interface {
	// When migctl creates a migration job, it will use this method.
	CreateMigrationJob(ctx context.Context, in *CreateMigrationJobRequest, opts ...grpc.CallOption) (*CreateMigrationJobResponse, error)
	// When the migrationd wants to share the migration with it's peer it uses this method.
	ShareMigrationJob(ctx context.Context, in *ShareMigrationJobRequest, opts ...grpc.CallOption) (*ShareMigrationJobResponse, error)
}

type migratorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMigratorServiceClient(cc grpc.ClientConnInterface) MigratorServiceClient {
	return &migratorServiceClient{cc}
}

func (c *migratorServiceClient) CreateMigrationJob(ctx context.Context, in *CreateMigrationJobRequest, opts ...grpc.CallOption) (*CreateMigrationJobResponse, error) {
	out := new(CreateMigrationJobResponse)
	err := c.cc.Invoke(ctx, "/MigratorService/CreateMigrationJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *migratorServiceClient) ShareMigrationJob(ctx context.Context, in *ShareMigrationJobRequest, opts ...grpc.CallOption) (*ShareMigrationJobResponse, error) {
	out := new(ShareMigrationJobResponse)
	err := c.cc.Invoke(ctx, "/MigratorService/ShareMigrationJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MigratorServiceServer is the server API for MigratorService service.
// All implementations must embed UnimplementedMigratorServiceServer
// for forward compatibility
type MigratorServiceServer interface {
	// When migctl creates a migration job, it will use this method.
	CreateMigrationJob(context.Context, *CreateMigrationJobRequest) (*CreateMigrationJobResponse, error)
	// When the migrationd wants to share the migration with it's peer it uses this method.
	ShareMigrationJob(context.Context, *ShareMigrationJobRequest) (*ShareMigrationJobResponse, error)
	mustEmbedUnimplementedMigratorServiceServer()
}

// UnimplementedMigratorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMigratorServiceServer struct {
}

func (UnimplementedMigratorServiceServer) CreateMigrationJob(context.Context, *CreateMigrationJobRequest) (*CreateMigrationJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMigrationJob not implemented")
}
func (UnimplementedMigratorServiceServer) ShareMigrationJob(context.Context, *ShareMigrationJobRequest) (*ShareMigrationJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShareMigrationJob not implemented")
}
func (UnimplementedMigratorServiceServer) mustEmbedUnimplementedMigratorServiceServer() {}

// UnsafeMigratorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MigratorServiceServer will
// result in compilation errors.
type UnsafeMigratorServiceServer interface {
	mustEmbedUnimplementedMigratorServiceServer()
}

func RegisterMigratorServiceServer(s grpc.ServiceRegistrar, srv MigratorServiceServer) {
	s.RegisterService(&MigratorService_ServiceDesc, srv)
}

func _MigratorService_CreateMigrationJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMigrationJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MigratorServiceServer).CreateMigrationJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MigratorService/CreateMigrationJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MigratorServiceServer).CreateMigrationJob(ctx, req.(*CreateMigrationJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MigratorService_ShareMigrationJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareMigrationJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MigratorServiceServer).ShareMigrationJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MigratorService/ShareMigrationJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MigratorServiceServer).ShareMigrationJob(ctx, req.(*ShareMigrationJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MigratorService_ServiceDesc is the grpc.ServiceDesc for MigratorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MigratorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MigratorService",
	HandlerType: (*MigratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMigrationJob",
			Handler:    _MigratorService_CreateMigrationJob_Handler,
		},
		{
			MethodName: "ShareMigrationJob",
			Handler:    _MigratorService_ShareMigrationJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "migrator.proto",
}