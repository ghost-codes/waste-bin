// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: service.proto

package pb

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
	WasteBin_CreateUserDetails_FullMethodName = "/pb.WasteBin/CreateUserDetails"
	WasteBin_FetchUserDetails_FullMethodName  = "/pb.WasteBin/FetchUserDetails"
	WasteBin_MakeRequest_FullMethodName       = "/pb.WasteBin/MakeRequest"
)

// WasteBinClient is the client API for WasteBin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WasteBinClient interface {
	CreateUserDetails(ctx context.Context, in *CreateUserDetailsRequest, opts ...grpc.CallOption) (*CreateUserDetailsResponse, error)
	FetchUserDetails(ctx context.Context, in *FetchUserParams, opts ...grpc.CallOption) (*FetchUserResponse, error)
	MakeRequest(ctx context.Context, in *CreateRequestPayload, opts ...grpc.CallOption) (*CreateRequestResponse, error)
}

type wasteBinClient struct {
	cc grpc.ClientConnInterface
}

func NewWasteBinClient(cc grpc.ClientConnInterface) WasteBinClient {
	return &wasteBinClient{cc}
}

func (c *wasteBinClient) CreateUserDetails(ctx context.Context, in *CreateUserDetailsRequest, opts ...grpc.CallOption) (*CreateUserDetailsResponse, error) {
	out := new(CreateUserDetailsResponse)
	err := c.cc.Invoke(ctx, WasteBin_CreateUserDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wasteBinClient) FetchUserDetails(ctx context.Context, in *FetchUserParams, opts ...grpc.CallOption) (*FetchUserResponse, error) {
	out := new(FetchUserResponse)
	err := c.cc.Invoke(ctx, WasteBin_FetchUserDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wasteBinClient) MakeRequest(ctx context.Context, in *CreateRequestPayload, opts ...grpc.CallOption) (*CreateRequestResponse, error) {
	out := new(CreateRequestResponse)
	err := c.cc.Invoke(ctx, WasteBin_MakeRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WasteBinServer is the server API for WasteBin service.
// All implementations must embed UnimplementedWasteBinServer
// for forward compatibility
type WasteBinServer interface {
	CreateUserDetails(context.Context, *CreateUserDetailsRequest) (*CreateUserDetailsResponse, error)
	FetchUserDetails(context.Context, *FetchUserParams) (*FetchUserResponse, error)
	MakeRequest(context.Context, *CreateRequestPayload) (*CreateRequestResponse, error)
	mustEmbedUnimplementedWasteBinServer()
}

// UnimplementedWasteBinServer must be embedded to have forward compatible implementations.
type UnimplementedWasteBinServer struct {
}

func (UnimplementedWasteBinServer) CreateUserDetails(context.Context, *CreateUserDetailsRequest) (*CreateUserDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserDetails not implemented")
}
func (UnimplementedWasteBinServer) FetchUserDetails(context.Context, *FetchUserParams) (*FetchUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchUserDetails not implemented")
}
func (UnimplementedWasteBinServer) MakeRequest(context.Context, *CreateRequestPayload) (*CreateRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeRequest not implemented")
}
func (UnimplementedWasteBinServer) mustEmbedUnimplementedWasteBinServer() {}

// UnsafeWasteBinServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WasteBinServer will
// result in compilation errors.
type UnsafeWasteBinServer interface {
	mustEmbedUnimplementedWasteBinServer()
}

func RegisterWasteBinServer(s grpc.ServiceRegistrar, srv WasteBinServer) {
	s.RegisterService(&WasteBin_ServiceDesc, srv)
}

func _WasteBin_CreateUserDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WasteBinServer).CreateUserDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WasteBin_CreateUserDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WasteBinServer).CreateUserDetails(ctx, req.(*CreateUserDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WasteBin_FetchUserDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchUserParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WasteBinServer).FetchUserDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WasteBin_FetchUserDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WasteBinServer).FetchUserDetails(ctx, req.(*FetchUserParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _WasteBin_MakeRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequestPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WasteBinServer).MakeRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WasteBin_MakeRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WasteBinServer).MakeRequest(ctx, req.(*CreateRequestPayload))
	}
	return interceptor(ctx, in, info, handler)
}

// WasteBin_ServiceDesc is the grpc.ServiceDesc for WasteBin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WasteBin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.WasteBin",
	HandlerType: (*WasteBinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserDetails",
			Handler:    _WasteBin_CreateUserDetails_Handler,
		},
		{
			MethodName: "FetchUserDetails",
			Handler:    _WasteBin_FetchUserDetails_Handler,
		},
		{
			MethodName: "MakeRequest",
			Handler:    _WasteBin_MakeRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
