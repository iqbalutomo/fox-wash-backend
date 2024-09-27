// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: user.proto

package userpb

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
	User_Register_FullMethodName                         = "/user.User/Register"
	User_VerifyNewUser_FullMethodName                    = "/user.User/VerifyNewUser"
	User_GetUser_FullMethodName                          = "/user.User/GetUser"
	User_CreateWasher_FullMethodName                     = "/user.User/CreateWasher"
	User_WasherActivation_FullMethodName                 = "/user.User/WasherActivation"
	User_GetWasher_FullMethodName                        = "/user.User/GetWasher"
	User_SetWasherStatusOnline_FullMethodName            = "/user.User/SetWasherStatusOnline"
	User_SetWasherStatusOffline_FullMethodName           = "/user.User/SetWasherStatusOffline"
	User_GetAvailableWasher_FullMethodName               = "/user.User/GetAvailableWasher"
	User_SetWasherStatusWashing_FullMethodName           = "/user.User/SetWasherStatusWashing"
	User_PostPublishMessagePaymentSuccess_FullMethodName = "/user.User/PostPublishMessagePaymentSuccess"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	VerifyNewUser(ctx context.Context, in *UserCredential, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUser(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*UserData, error)
	CreateWasher(ctx context.Context, in *WasherID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	WasherActivation(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetWasher(ctx context.Context, in *WasherID, opts ...grpc.CallOption) (*WasherData, error)
	SetWasherStatusOnline(ctx context.Context, in *WasherID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetWasherStatusOffline(ctx context.Context, in *WasherID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAvailableWasher(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*WasherOrderData, error)
	SetWasherStatusWashing(ctx context.Context, in *WasherID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PostPublishMessagePaymentSuccess(ctx context.Context, in *PaymentSuccessData, opts ...grpc.CallOption) (*PaymentSuccessData, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, User_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) VerifyNewUser(ctx context.Context, in *UserCredential, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, User_VerifyNewUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUser(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*UserData, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserData)
	err := c.cc.Invoke(ctx, User_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateWasher(ctx context.Context, in *WasherID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, User_CreateWasher_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) WasherActivation(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, User_WasherActivation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetWasher(ctx context.Context, in *WasherID, opts ...grpc.CallOption) (*WasherData, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WasherData)
	err := c.cc.Invoke(ctx, User_GetWasher_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetWasherStatusOnline(ctx context.Context, in *WasherID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, User_SetWasherStatusOnline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetWasherStatusOffline(ctx context.Context, in *WasherID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, User_SetWasherStatusOffline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetAvailableWasher(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*WasherOrderData, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WasherOrderData)
	err := c.cc.Invoke(ctx, User_GetAvailableWasher_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetWasherStatusWashing(ctx context.Context, in *WasherID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, User_SetWasherStatusWashing_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) PostPublishMessagePaymentSuccess(ctx context.Context, in *PaymentSuccessData, opts ...grpc.CallOption) (*PaymentSuccessData, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PaymentSuccessData)
	err := c.cc.Invoke(ctx, User_PostPublishMessagePaymentSuccess_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations should embed UnimplementedUserServer
// for forward compatibility.
type UserServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	VerifyNewUser(context.Context, *UserCredential) (*emptypb.Empty, error)
	GetUser(context.Context, *EmailRequest) (*UserData, error)
	CreateWasher(context.Context, *WasherID) (*emptypb.Empty, error)
	WasherActivation(context.Context, *EmailRequest) (*emptypb.Empty, error)
	GetWasher(context.Context, *WasherID) (*WasherData, error)
	SetWasherStatusOnline(context.Context, *WasherID) (*emptypb.Empty, error)
	SetWasherStatusOffline(context.Context, *WasherID) (*emptypb.Empty, error)
	GetAvailableWasher(context.Context, *emptypb.Empty) (*WasherOrderData, error)
	SetWasherStatusWashing(context.Context, *WasherID) (*emptypb.Empty, error)
	PostPublishMessagePaymentSuccess(context.Context, *PaymentSuccessData) (*PaymentSuccessData, error)
}

// UnimplementedUserServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServer struct{}

func (UnimplementedUserServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServer) VerifyNewUser(context.Context, *UserCredential) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyNewUser not implemented")
}
func (UnimplementedUserServer) GetUser(context.Context, *EmailRequest) (*UserData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServer) CreateWasher(context.Context, *WasherID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWasher not implemented")
}
func (UnimplementedUserServer) WasherActivation(context.Context, *EmailRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WasherActivation not implemented")
}
func (UnimplementedUserServer) GetWasher(context.Context, *WasherID) (*WasherData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWasher not implemented")
}
func (UnimplementedUserServer) SetWasherStatusOnline(context.Context, *WasherID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetWasherStatusOnline not implemented")
}
func (UnimplementedUserServer) SetWasherStatusOffline(context.Context, *WasherID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetWasherStatusOffline not implemented")
}
func (UnimplementedUserServer) GetAvailableWasher(context.Context, *emptypb.Empty) (*WasherOrderData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableWasher not implemented")
}
func (UnimplementedUserServer) SetWasherStatusWashing(context.Context, *WasherID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetWasherStatusWashing not implemented")
}
func (UnimplementedUserServer) PostPublishMessagePaymentSuccess(context.Context, *PaymentSuccessData) (*PaymentSuccessData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostPublishMessagePaymentSuccess not implemented")
}
func (UnimplementedUserServer) testEmbeddedByValue() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	// If the following call pancis, it indicates UnimplementedUserServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_VerifyNewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).VerifyNewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_VerifyNewUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).VerifyNewUser(ctx, req.(*UserCredential))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUser(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateWasher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WasherID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateWasher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_CreateWasher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateWasher(ctx, req.(*WasherID))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_WasherActivation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).WasherActivation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_WasherActivation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).WasherActivation(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetWasher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WasherID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetWasher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetWasher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetWasher(ctx, req.(*WasherID))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetWasherStatusOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WasherID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetWasherStatusOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SetWasherStatusOnline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetWasherStatusOnline(ctx, req.(*WasherID))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetWasherStatusOffline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WasherID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetWasherStatusOffline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SetWasherStatusOffline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetWasherStatusOffline(ctx, req.(*WasherID))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetAvailableWasher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetAvailableWasher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetAvailableWasher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetAvailableWasher(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetWasherStatusWashing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WasherID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetWasherStatusWashing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SetWasherStatusWashing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetWasherStatusWashing(ctx, req.(*WasherID))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_PostPublishMessagePaymentSuccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentSuccessData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).PostPublishMessagePaymentSuccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_PostPublishMessagePaymentSuccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).PostPublishMessagePaymentSuccess(ctx, req.(*PaymentSuccessData))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "VerifyNewUser",
			Handler:    _User_VerifyNewUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _User_GetUser_Handler,
		},
		{
			MethodName: "CreateWasher",
			Handler:    _User_CreateWasher_Handler,
		},
		{
			MethodName: "WasherActivation",
			Handler:    _User_WasherActivation_Handler,
		},
		{
			MethodName: "GetWasher",
			Handler:    _User_GetWasher_Handler,
		},
		{
			MethodName: "SetWasherStatusOnline",
			Handler:    _User_SetWasherStatusOnline_Handler,
		},
		{
			MethodName: "SetWasherStatusOffline",
			Handler:    _User_SetWasherStatusOffline_Handler,
		},
		{
			MethodName: "GetAvailableWasher",
			Handler:    _User_GetAvailableWasher_Handler,
		},
		{
			MethodName: "SetWasherStatusWashing",
			Handler:    _User_SetWasherStatusWashing_Handler,
		},
		{
			MethodName: "PostPublishMessagePaymentSuccess",
			Handler:    _User_PostPublishMessagePaymentSuccess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
