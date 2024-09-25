// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: washstation.proto

package washstationpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	WashStation_CreateWashPackage_FullMethodName        = "/wash_station.WashStation/CreateWashPackage"
	WashStation_FindAllWashPackages_FullMethodName      = "/wash_station.WashStation/FindAllWashPackages"
	WashStation_FindWashPackageByID_FullMethodName      = "/wash_station.WashStation/FindWashPackageByID"
	WashStation_FindMultipleWashPackages_FullMethodName = "/wash_station.WashStation/FindMultipleWashPackages"
	WashStation_UpdateWashPackage_FullMethodName        = "/wash_station.WashStation/UpdateWashPackage"
	WashStation_DeleteWashPackage_FullMethodName        = "/wash_station.WashStation/DeleteWashPackage"
)

// WashStationClient is the client API for WashStation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WashStationClient interface {
	// Admin
	CreateWashPackage(ctx context.Context, in *NewWashPackageData, opts ...grpc.CallOption) (*CreateWashPackageResponse, error)
	FindAllWashPackages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*WashPackageCompactRepeated, error)
	FindWashPackageByID(ctx context.Context, in *WashPackageID, opts ...grpc.CallOption) (*WashPackageData, error)
	FindMultipleWashPackages(ctx context.Context, in *WashPackageIDs, opts ...grpc.CallOption) (*WashPackageCompactRepeated, error)
	UpdateWashPackage(ctx context.Context, in *UpdateWashPackageData, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteWashPackage(ctx context.Context, in *WashPackageID, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type washStationClient struct {
	cc grpc.ClientConnInterface
}

func NewWashStationClient(cc grpc.ClientConnInterface) WashStationClient {
	return &washStationClient{cc}
}

func (c *washStationClient) CreateWashPackage(ctx context.Context, in *NewWashPackageData, opts ...grpc.CallOption) (*CreateWashPackageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateWashPackageResponse)
	err := c.cc.Invoke(ctx, WashStation_CreateWashPackage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *washStationClient) FindAllWashPackages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*WashPackageCompactRepeated, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WashPackageCompactRepeated)
	err := c.cc.Invoke(ctx, WashStation_FindAllWashPackages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *washStationClient) FindWashPackageByID(ctx context.Context, in *WashPackageID, opts ...grpc.CallOption) (*WashPackageData, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WashPackageData)
	err := c.cc.Invoke(ctx, WashStation_FindWashPackageByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *washStationClient) FindMultipleWashPackages(ctx context.Context, in *WashPackageIDs, opts ...grpc.CallOption) (*WashPackageCompactRepeated, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WashPackageCompactRepeated)
	err := c.cc.Invoke(ctx, WashStation_FindMultipleWashPackages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *washStationClient) UpdateWashPackage(ctx context.Context, in *UpdateWashPackageData, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, WashStation_UpdateWashPackage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *washStationClient) DeleteWashPackage(ctx context.Context, in *WashPackageID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, WashStation_DeleteWashPackage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WashStationServer is the server API for WashStation service.
// All implementations should embed UnimplementedWashStationServer
// for forward compatibility.
type WashStationServer interface {
	// Admin
	CreateWashPackage(context.Context, *NewWashPackageData) (*CreateWashPackageResponse, error)
	FindAllWashPackages(context.Context, *emptypb.Empty) (*WashPackageCompactRepeated, error)
	FindWashPackageByID(context.Context, *WashPackageID) (*WashPackageData, error)
	FindMultipleWashPackages(context.Context, *WashPackageIDs) (*WashPackageCompactRepeated, error)
	UpdateWashPackage(context.Context, *UpdateWashPackageData) (*emptypb.Empty, error)
	DeleteWashPackage(context.Context, *WashPackageID) (*emptypb.Empty, error)
}

// UnimplementedWashStationServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWashStationServer struct{}

func (UnimplementedWashStationServer) CreateWashPackage(context.Context, *NewWashPackageData) (*CreateWashPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWashPackage not implemented")
}
func (UnimplementedWashStationServer) FindAllWashPackages(context.Context, *emptypb.Empty) (*WashPackageCompactRepeated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllWashPackages not implemented")
}
func (UnimplementedWashStationServer) FindWashPackageByID(context.Context, *WashPackageID) (*WashPackageData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindWashPackageByID not implemented")
}
func (UnimplementedWashStationServer) FindMultipleWashPackages(context.Context, *WashPackageIDs) (*WashPackageCompactRepeated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindMultipleWashPackages not implemented")
}
func (UnimplementedWashStationServer) UpdateWashPackage(context.Context, *UpdateWashPackageData) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWashPackage not implemented")
}
func (UnimplementedWashStationServer) DeleteWashPackage(context.Context, *WashPackageID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWashPackage not implemented")
}
func (UnimplementedWashStationServer) testEmbeddedByValue() {}

// UnsafeWashStationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WashStationServer will
// result in compilation errors.
type UnsafeWashStationServer interface {
	mustEmbedUnimplementedWashStationServer()
}

func RegisterWashStationServer(s grpc.ServiceRegistrar, srv WashStationServer) {
	// If the following call pancis, it indicates UnimplementedWashStationServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WashStation_ServiceDesc, srv)
}

func _WashStation_CreateWashPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewWashPackageData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WashStationServer).CreateWashPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WashStation_CreateWashPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WashStationServer).CreateWashPackage(ctx, req.(*NewWashPackageData))
	}
	return interceptor(ctx, in, info, handler)
}

func _WashStation_FindAllWashPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WashStationServer).FindAllWashPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WashStation_FindAllWashPackages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WashStationServer).FindAllWashPackages(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _WashStation_FindWashPackageByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WashPackageID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WashStationServer).FindWashPackageByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WashStation_FindWashPackageByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WashStationServer).FindWashPackageByID(ctx, req.(*WashPackageID))
	}
	return interceptor(ctx, in, info, handler)
}

func _WashStation_FindMultipleWashPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WashPackageIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WashStationServer).FindMultipleWashPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WashStation_FindMultipleWashPackages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WashStationServer).FindMultipleWashPackages(ctx, req.(*WashPackageIDs))
	}
	return interceptor(ctx, in, info, handler)
}

func _WashStation_UpdateWashPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWashPackageData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WashStationServer).UpdateWashPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WashStation_UpdateWashPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WashStationServer).UpdateWashPackage(ctx, req.(*UpdateWashPackageData))
	}
	return interceptor(ctx, in, info, handler)
}

func _WashStation_DeleteWashPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WashPackageID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WashStationServer).DeleteWashPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WashStation_DeleteWashPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WashStationServer).DeleteWashPackage(ctx, req.(*WashPackageID))
	}
	return interceptor(ctx, in, info, handler)
}

// WashStation_ServiceDesc is the grpc.ServiceDesc for WashStation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WashStation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wash_station.WashStation",
	HandlerType: (*WashStationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWashPackage",
			Handler:    _WashStation_CreateWashPackage_Handler,
		},
		{
			MethodName: "FindAllWashPackages",
			Handler:    _WashStation_FindAllWashPackages_Handler,
		},
		{
			MethodName: "FindWashPackageByID",
			Handler:    _WashStation_FindWashPackageByID_Handler,
		},
		{
			MethodName: "FindMultipleWashPackages",
			Handler:    _WashStation_FindMultipleWashPackages_Handler,
		},
		{
			MethodName: "UpdateWashPackage",
			Handler:    _WashStation_UpdateWashPackage_Handler,
		},
		{
			MethodName: "DeleteWashPackage",
			Handler:    _WashStation_DeleteWashPackage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "washstation.proto",
}
