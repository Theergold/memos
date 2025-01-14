// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/v2/system_service.proto

package apiv2

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

const (
	SystemService_GetSystemInfo_FullMethodName = "/memos.api.v2.SystemService/GetSystemInfo"
)

// SystemServiceClient is the client API for SystemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemServiceClient interface {
	GetSystemInfo(ctx context.Context, in *GetSystemInfoRequest, opts ...grpc.CallOption) (*GetSystemInfoResponse, error)
}

type systemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemServiceClient(cc grpc.ClientConnInterface) SystemServiceClient {
	return &systemServiceClient{cc}
}

func (c *systemServiceClient) GetSystemInfo(ctx context.Context, in *GetSystemInfoRequest, opts ...grpc.CallOption) (*GetSystemInfoResponse, error) {
	out := new(GetSystemInfoResponse)
	err := c.cc.Invoke(ctx, SystemService_GetSystemInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemServiceServer is the server API for SystemService service.
// All implementations must embed UnimplementedSystemServiceServer
// for forward compatibility
type SystemServiceServer interface {
	GetSystemInfo(context.Context, *GetSystemInfoRequest) (*GetSystemInfoResponse, error)
	mustEmbedUnimplementedSystemServiceServer()
}

// UnimplementedSystemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSystemServiceServer struct {
}

func (UnimplementedSystemServiceServer) GetSystemInfo(context.Context, *GetSystemInfoRequest) (*GetSystemInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemInfo not implemented")
}
func (UnimplementedSystemServiceServer) mustEmbedUnimplementedSystemServiceServer() {}

// UnsafeSystemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemServiceServer will
// result in compilation errors.
type UnsafeSystemServiceServer interface {
	mustEmbedUnimplementedSystemServiceServer()
}

func RegisterSystemServiceServer(s grpc.ServiceRegistrar, srv SystemServiceServer) {
	s.RegisterService(&SystemService_ServiceDesc, srv)
}

func _SystemService_GetSystemInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSystemInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).GetSystemInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemService_GetSystemInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).GetSystemInfo(ctx, req.(*GetSystemInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemService_ServiceDesc is the grpc.ServiceDesc for SystemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "memos.api.v2.SystemService",
	HandlerType: (*SystemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSystemInfo",
			Handler:    _SystemService_GetSystemInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v2/system_service.proto",
}
