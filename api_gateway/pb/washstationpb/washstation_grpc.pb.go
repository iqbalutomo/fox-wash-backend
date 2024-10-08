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
	WashStation_CreateWashPackage_FullMethodName             = "/wash_station.WashStation/CreateWashPackage"
	WashStation_FindAllWashPackages_FullMethodName           = "/wash_station.WashStation/FindAllWashPackages"
	WashStation_FindWashPackageByID_FullMethodName           = "/wash_station.WashStation/FindWashPackageByID"
	WashStation_FindMultipleWashPackages_FullMethodName      = "/wash_station.WashStation/FindMultipleWashPackages"
	WashStation_UpdateWashPackage_FullMethodName             = "/wash_station.WashStation/UpdateWashPackage"
	WashStation_DeleteWashPackage_FullMethodName             = "/wash_station.WashStation/DeleteWashPackage"
	WashStation_CreateDetailingPackage_FullMethodName        = "/wash_station.WashStation/CreateDetailingPackage"
	WashStation_FindAllDetailingPackages_FullMethodName      = "/wash_station.WashStation/FindAllDetailingPackages"
	WashStation_FindDetailingPackageByID_FullMethodName      = "/wash_station.WashStation/FindDetailingPackageByID"
	WashStation_FindMultipleDetailingPackages_FullMethodName = "/wash_station.WashStation/FindMultipleDetailingPackages"
	WashStation_UpdateDetailingPackage_FullMethodName        = "/wash_station.WashStation/UpdateDetailingPackage"
	WashStation_DeleteDetailingPackage_FullMethodName        = "/wash_station.WashStation/DeleteDetailingPackage"
)

// WashStationClient is the client API for WashStation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WashStationClient interface {
	// Wash
	CreateWashPackage(ctx context.Context, in *NewWashPackageData, opts ...grpc.CallOption) (*CreateWashPackageResponse, error)
	FindAllWashPackages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*WashPackageCompactRepeated, error)
	FindWashPackageByID(ctx context.Context, in *WashPackageID, opts ...grpc.CallOption) (*WashPackageData, error)
	FindMultipleWashPackages(ctx context.Context, in *WashPackageIDs, opts ...grpc.CallOption) (*WashPackageCompactRepeated, error)
	UpdateWashPackage(ctx context.Context, in *UpdateWashPackageData, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteWashPackage(ctx context.Context, in *WashPackageID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Detailing
	CreateDetailingPackage(ctx context.Context, in *NewDetailingPackageData, opts ...grpc.CallOption) (*CreateDetailingPackageResponse, error)
	FindAllDetailingPackages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DetailingPackageCompactRepeated, error)
	FindDetailingPackageByID(ctx context.Context, in *DetailingPackageID, opts ...grpc.CallOption) (*DetailingPackageData, error)
	FindMultipleDetailingPackages(ctx context.Context, in *DetailingPackageIDs, opts ...grpc.CallOption) (*DetailingPackageCompactRepeated, error)
	UpdateDetailingPackage(ctx context.Context, in *UpdateDetailingPackageData, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteDetailingPackage(ctx context.Context, in *DetailingPackageID, opts ...grpc.CallOption) (*emptypb.Empty, error)
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

func (c *washStationClient) CreateDetailingPackage(ctx context.Context, in *NewDetailingPackageData, opts ...grpc.CallOption) (*CreateDetailingPackageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDetailingPackageResponse)
	err := c.cc.Invoke(ctx, WashStation_CreateDetailingPackage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *washStationClient) FindAllDetailingPackages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DetailingPackageCompactRepeated, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DetailingPackageCompactRepeated)
	err := c.cc.Invoke(ctx, WashStation_FindAllDetailingPackages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *washStationClient) FindDetailingPackageByID(ctx context.Context, in *DetailingPackageID, opts ...grpc.CallOption) (*DetailingPackageData, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DetailingPackageData)
	err := c.cc.Invoke(ctx, WashStation_FindDetailingPackageByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *washStationClient) FindMultipleDetailingPackages(ctx context.Context, in *DetailingPackageIDs, opts ...grpc.CallOption) (*DetailingPackageCompactRepeated, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DetailingPackageCompactRepeated)
	err := c.cc.Invoke(ctx, WashStation_FindMultipleDetailingPackages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *washStationClient) UpdateDetailingPackage(ctx context.Context, in *UpdateDetailingPackageData, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, WashStation_UpdateDetailingPackage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *washStationClient) DeleteDetailingPackage(ctx context.Context, in *DetailingPackageID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, WashStation_DeleteDetailingPackage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WashStationServer is the server API for WashStation service.
// All implementations should embed UnimplementedWashStationServer
// for forward compatibility.
type WashStationServer interface {
	// Wash
	CreateWashPackage(context.Context, *NewWashPackageData) (*CreateWashPackageResponse, error)
	FindAllWashPackages(context.Context, *emptypb.Empty) (*WashPackageCompactRepeated, error)
	FindWashPackageByID(context.Context, *WashPackageID) (*WashPackageData, error)
	FindMultipleWashPackages(context.Context, *WashPackageIDs) (*WashPackageCompactRepeated, error)
	UpdateWashPackage(context.Context, *UpdateWashPackageData) (*emptypb.Empty, error)
	DeleteWashPackage(context.Context, *WashPackageID) (*emptypb.Empty, error)
	// Detailing
	CreateDetailingPackage(context.Context, *NewDetailingPackageData) (*CreateDetailingPackageResponse, error)
	FindAllDetailingPackages(context.Context, *emptypb.Empty) (*DetailingPackageCompactRepeated, error)
	FindDetailingPackageByID(context.Context, *DetailingPackageID) (*DetailingPackageData, error)
	FindMultipleDetailingPackages(context.Context, *DetailingPackageIDs) (*DetailingPackageCompactRepeated, error)
	UpdateDetailingPackage(context.Context, *UpdateDetailingPackageData) (*emptypb.Empty, error)
	DeleteDetailingPackage(context.Context, *DetailingPackageID) (*emptypb.Empty, error)
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
func (UnimplementedWashStationServer) CreateDetailingPackage(context.Context, *NewDetailingPackageData) (*CreateDetailingPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDetailingPackage not implemented")
}
func (UnimplementedWashStationServer) FindAllDetailingPackages(context.Context, *emptypb.Empty) (*DetailingPackageCompactRepeated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllDetailingPackages not implemented")
}
func (UnimplementedWashStationServer) FindDetailingPackageByID(context.Context, *DetailingPackageID) (*DetailingPackageData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindDetailingPackageByID not implemented")
}
func (UnimplementedWashStationServer) FindMultipleDetailingPackages(context.Context, *DetailingPackageIDs) (*DetailingPackageCompactRepeated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindMultipleDetailingPackages not implemented")
}
func (UnimplementedWashStationServer) UpdateDetailingPackage(context.Context, *UpdateDetailingPackageData) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDetailingPackage not implemented")
}
func (UnimplementedWashStationServer) DeleteDetailingPackage(context.Context, *DetailingPackageID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDetailingPackage not implemented")
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

func _WashStation_CreateDetailingPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewDetailingPackageData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WashStationServer).CreateDetailingPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WashStation_CreateDetailingPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WashStationServer).CreateDetailingPackage(ctx, req.(*NewDetailingPackageData))
	}
	return interceptor(ctx, in, info, handler)
}

func _WashStation_FindAllDetailingPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WashStationServer).FindAllDetailingPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WashStation_FindAllDetailingPackages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WashStationServer).FindAllDetailingPackages(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _WashStation_FindDetailingPackageByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailingPackageID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WashStationServer).FindDetailingPackageByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WashStation_FindDetailingPackageByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WashStationServer).FindDetailingPackageByID(ctx, req.(*DetailingPackageID))
	}
	return interceptor(ctx, in, info, handler)
}

func _WashStation_FindMultipleDetailingPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailingPackageIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WashStationServer).FindMultipleDetailingPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WashStation_FindMultipleDetailingPackages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WashStationServer).FindMultipleDetailingPackages(ctx, req.(*DetailingPackageIDs))
	}
	return interceptor(ctx, in, info, handler)
}

func _WashStation_UpdateDetailingPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDetailingPackageData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WashStationServer).UpdateDetailingPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WashStation_UpdateDetailingPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WashStationServer).UpdateDetailingPackage(ctx, req.(*UpdateDetailingPackageData))
	}
	return interceptor(ctx, in, info, handler)
}

func _WashStation_DeleteDetailingPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailingPackageID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WashStationServer).DeleteDetailingPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WashStation_DeleteDetailingPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WashStationServer).DeleteDetailingPackage(ctx, req.(*DetailingPackageID))
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
		{
			MethodName: "CreateDetailingPackage",
			Handler:    _WashStation_CreateDetailingPackage_Handler,
		},
		{
			MethodName: "FindAllDetailingPackages",
			Handler:    _WashStation_FindAllDetailingPackages_Handler,
		},
		{
			MethodName: "FindDetailingPackageByID",
			Handler:    _WashStation_FindDetailingPackageByID_Handler,
		},
		{
			MethodName: "FindMultipleDetailingPackages",
			Handler:    _WashStation_FindMultipleDetailingPackages_Handler,
		},
		{
			MethodName: "UpdateDetailingPackage",
			Handler:    _WashStation_UpdateDetailingPackage_Handler,
		},
		{
			MethodName: "DeleteDetailingPackage",
			Handler:    _WashStation_DeleteDetailingPackage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "washstation.proto",
}
