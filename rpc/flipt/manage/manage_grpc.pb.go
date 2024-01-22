// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: manage/manage.proto

package manage

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
	ManageService_GetNamespace_FullMethodName = "/flipt.manage.ManageService/GetNamespace"
	ManageService_PutFlag_FullMethodName      = "/flipt.manage.ManageService/PutFlag"
	ManageService_PutSegment_FullMethodName   = "/flipt.manage.ManageService/PutSegment"
)

// ManageServiceClient is the client API for ManageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManageServiceClient interface {
	GetNamespace(ctx context.Context, in *GetNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error)
	PutFlag(ctx context.Context, in *Flag, opts ...grpc.CallOption) (*Proposal, error)
	PutSegment(ctx context.Context, in *Segment, opts ...grpc.CallOption) (*Proposal, error)
}

type manageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManageServiceClient(cc grpc.ClientConnInterface) ManageServiceClient {
	return &manageServiceClient{cc}
}

func (c *manageServiceClient) GetNamespace(ctx context.Context, in *GetNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error) {
	out := new(Namespace)
	err := c.cc.Invoke(ctx, ManageService_GetNamespace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageServiceClient) PutFlag(ctx context.Context, in *Flag, opts ...grpc.CallOption) (*Proposal, error) {
	out := new(Proposal)
	err := c.cc.Invoke(ctx, ManageService_PutFlag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageServiceClient) PutSegment(ctx context.Context, in *Segment, opts ...grpc.CallOption) (*Proposal, error) {
	out := new(Proposal)
	err := c.cc.Invoke(ctx, ManageService_PutSegment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManageServiceServer is the server API for ManageService service.
// All implementations must embed UnimplementedManageServiceServer
// for forward compatibility
type ManageServiceServer interface {
	GetNamespace(context.Context, *GetNamespaceRequest) (*Namespace, error)
	PutFlag(context.Context, *Flag) (*Proposal, error)
	PutSegment(context.Context, *Segment) (*Proposal, error)
	mustEmbedUnimplementedManageServiceServer()
}

// UnimplementedManageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedManageServiceServer struct {
}

func (UnimplementedManageServiceServer) GetNamespace(context.Context, *GetNamespaceRequest) (*Namespace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNamespace not implemented")
}
func (UnimplementedManageServiceServer) PutFlag(context.Context, *Flag) (*Proposal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutFlag not implemented")
}
func (UnimplementedManageServiceServer) PutSegment(context.Context, *Segment) (*Proposal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutSegment not implemented")
}
func (UnimplementedManageServiceServer) mustEmbedUnimplementedManageServiceServer() {}

// UnsafeManageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManageServiceServer will
// result in compilation errors.
type UnsafeManageServiceServer interface {
	mustEmbedUnimplementedManageServiceServer()
}

func RegisterManageServiceServer(s grpc.ServiceRegistrar, srv ManageServiceServer) {
	s.RegisterService(&ManageService_ServiceDesc, srv)
}

func _ManageService_GetNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageServiceServer).GetNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManageService_GetNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageServiceServer).GetNamespace(ctx, req.(*GetNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageService_PutFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Flag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageServiceServer).PutFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManageService_PutFlag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageServiceServer).PutFlag(ctx, req.(*Flag))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageService_PutSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Segment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageServiceServer).PutSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManageService_PutSegment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageServiceServer).PutSegment(ctx, req.(*Segment))
	}
	return interceptor(ctx, in, info, handler)
}

// ManageService_ServiceDesc is the grpc.ServiceDesc for ManageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flipt.manage.ManageService",
	HandlerType: (*ManageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNamespace",
			Handler:    _ManageService_GetNamespace_Handler,
		},
		{
			MethodName: "PutFlag",
			Handler:    _ManageService_PutFlag_Handler,
		},
		{
			MethodName: "PutSegment",
			Handler:    _ManageService_PutSegment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manage/manage.proto",
}
