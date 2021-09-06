// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ova_serial_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OvaSerialClient is the client API for OvaSerial service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OvaSerialClient interface {
	CreateSerialV1(ctx context.Context, in *CreateSerialRequestV1, opts ...grpc.CallOption) (*CreateSerialResponseV1, error)
	MultiCreateSerialV1(ctx context.Context, in *MultiCreateSerialRequestV1, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetSerialV1(ctx context.Context, in *GetSerialRequestV1, opts ...grpc.CallOption) (*GetSerialResponseV1, error)
	ListSerialsV1(ctx context.Context, in *ListSerialsRequestV1, opts ...grpc.CallOption) (*ListSerialsResponseV1, error)
	RemoveSerialV1(ctx context.Context, in *RemoveSerialRequestV1, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateSerialV1(ctx context.Context, in *UpdateSerialRequestV1, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type ovaSerialClient struct {
	cc grpc.ClientConnInterface
}

func NewOvaSerialClient(cc grpc.ClientConnInterface) OvaSerialClient {
	return &ovaSerialClient{cc}
}

func (c *ovaSerialClient) CreateSerialV1(ctx context.Context, in *CreateSerialRequestV1, opts ...grpc.CallOption) (*CreateSerialResponseV1, error) {
	out := new(CreateSerialResponseV1)
	err := c.cc.Invoke(ctx, "/ova.serial.api.OvaSerial/CreateSerialV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ovaSerialClient) MultiCreateSerialV1(ctx context.Context, in *MultiCreateSerialRequestV1, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ova.serial.api.OvaSerial/MultiCreateSerialV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ovaSerialClient) GetSerialV1(ctx context.Context, in *GetSerialRequestV1, opts ...grpc.CallOption) (*GetSerialResponseV1, error) {
	out := new(GetSerialResponseV1)
	err := c.cc.Invoke(ctx, "/ova.serial.api.OvaSerial/GetSerialV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ovaSerialClient) ListSerialsV1(ctx context.Context, in *ListSerialsRequestV1, opts ...grpc.CallOption) (*ListSerialsResponseV1, error) {
	out := new(ListSerialsResponseV1)
	err := c.cc.Invoke(ctx, "/ova.serial.api.OvaSerial/ListSerialsV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ovaSerialClient) RemoveSerialV1(ctx context.Context, in *RemoveSerialRequestV1, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ova.serial.api.OvaSerial/RemoveSerialV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ovaSerialClient) UpdateSerialV1(ctx context.Context, in *UpdateSerialRequestV1, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ova.serial.api.OvaSerial/UpdateSerialV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OvaSerialServer is the server API for OvaSerial service.
// All implementations must embed UnimplementedOvaSerialServer
// for forward compatibility
type OvaSerialServer interface {
	CreateSerialV1(context.Context, *CreateSerialRequestV1) (*CreateSerialResponseV1, error)
	MultiCreateSerialV1(context.Context, *MultiCreateSerialRequestV1) (*emptypb.Empty, error)
	GetSerialV1(context.Context, *GetSerialRequestV1) (*GetSerialResponseV1, error)
	ListSerialsV1(context.Context, *ListSerialsRequestV1) (*ListSerialsResponseV1, error)
	RemoveSerialV1(context.Context, *RemoveSerialRequestV1) (*emptypb.Empty, error)
	UpdateSerialV1(context.Context, *UpdateSerialRequestV1) (*emptypb.Empty, error)
	mustEmbedUnimplementedOvaSerialServer()
}

// UnimplementedOvaSerialServer must be embedded to have forward compatible implementations.
type UnimplementedOvaSerialServer struct {
}

func (UnimplementedOvaSerialServer) CreateSerialV1(context.Context, *CreateSerialRequestV1) (*CreateSerialResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSerialV1 not implemented")
}
func (UnimplementedOvaSerialServer) MultiCreateSerialV1(context.Context, *MultiCreateSerialRequestV1) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiCreateSerialV1 not implemented")
}
func (UnimplementedOvaSerialServer) GetSerialV1(context.Context, *GetSerialRequestV1) (*GetSerialResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSerialV1 not implemented")
}
func (UnimplementedOvaSerialServer) ListSerialsV1(context.Context, *ListSerialsRequestV1) (*ListSerialsResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSerialsV1 not implemented")
}
func (UnimplementedOvaSerialServer) RemoveSerialV1(context.Context, *RemoveSerialRequestV1) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSerialV1 not implemented")
}
func (UnimplementedOvaSerialServer) UpdateSerialV1(context.Context, *UpdateSerialRequestV1) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSerialV1 not implemented")
}
func (UnimplementedOvaSerialServer) mustEmbedUnimplementedOvaSerialServer() {}

// UnsafeOvaSerialServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OvaSerialServer will
// result in compilation errors.
type UnsafeOvaSerialServer interface {
	mustEmbedUnimplementedOvaSerialServer()
}

func RegisterOvaSerialServer(s grpc.ServiceRegistrar, srv OvaSerialServer) {
	s.RegisterService(&OvaSerial_ServiceDesc, srv)
}

func _OvaSerial_CreateSerialV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSerialRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaSerialServer).CreateSerialV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.serial.api.OvaSerial/CreateSerialV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaSerialServer).CreateSerialV1(ctx, req.(*CreateSerialRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _OvaSerial_MultiCreateSerialV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiCreateSerialRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaSerialServer).MultiCreateSerialV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.serial.api.OvaSerial/MultiCreateSerialV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaSerialServer).MultiCreateSerialV1(ctx, req.(*MultiCreateSerialRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _OvaSerial_GetSerialV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSerialRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaSerialServer).GetSerialV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.serial.api.OvaSerial/GetSerialV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaSerialServer).GetSerialV1(ctx, req.(*GetSerialRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _OvaSerial_ListSerialsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSerialsRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaSerialServer).ListSerialsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.serial.api.OvaSerial/ListSerialsV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaSerialServer).ListSerialsV1(ctx, req.(*ListSerialsRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _OvaSerial_RemoveSerialV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSerialRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaSerialServer).RemoveSerialV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.serial.api.OvaSerial/RemoveSerialV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaSerialServer).RemoveSerialV1(ctx, req.(*RemoveSerialRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _OvaSerial_UpdateSerialV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSerialRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaSerialServer).UpdateSerialV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.serial.api.OvaSerial/UpdateSerialV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaSerialServer).UpdateSerialV1(ctx, req.(*UpdateSerialRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

// OvaSerial_ServiceDesc is the grpc.ServiceDesc for OvaSerial service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OvaSerial_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ova.serial.api.OvaSerial",
	HandlerType: (*OvaSerialServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSerialV1",
			Handler:    _OvaSerial_CreateSerialV1_Handler,
		},
		{
			MethodName: "MultiCreateSerialV1",
			Handler:    _OvaSerial_MultiCreateSerialV1_Handler,
		},
		{
			MethodName: "GetSerialV1",
			Handler:    _OvaSerial_GetSerialV1_Handler,
		},
		{
			MethodName: "ListSerialsV1",
			Handler:    _OvaSerial_ListSerialsV1_Handler,
		},
		{
			MethodName: "RemoveSerialV1",
			Handler:    _OvaSerial_RemoveSerialV1_Handler,
		},
		{
			MethodName: "UpdateSerialV1",
			Handler:    _OvaSerial_UpdateSerialV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ova-serial-api.proto",
}
