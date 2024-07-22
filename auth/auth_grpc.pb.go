// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: auth/auth.proto

package auth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	JWT_TokenService_GenerateJWt_FullMethodName = "/auth.JWT_TokenService/GenerateJWt"
	JWT_TokenService_VerifyJWT_FullMethodName   = "/auth.JWT_TokenService/VerifyJWT"
	JWT_TokenService_GetUserID_FullMethodName   = "/auth.JWT_TokenService/GetUserID"
)

// JWT_TokenServiceClient is the client API for JWT_TokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JWT_TokenServiceClient interface {
	GenerateJWt(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error)
	VerifyJWT(ctx context.Context, in *VerifyJWTRequest, opts ...grpc.CallOption) (*VerifyJWTResponse, error)
	GetUserID(ctx context.Context, in *GetUserIDRequest, opts ...grpc.CallOption) (*GetUserIDResponse, error)
}

type jWT_TokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJWT_TokenServiceClient(cc grpc.ClientConnInterface) JWT_TokenServiceClient {
	return &jWT_TokenServiceClient{cc}
}

func (c *jWT_TokenServiceClient) GenerateJWt(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateResponse)
	err := c.cc.Invoke(ctx, JWT_TokenService_GenerateJWt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jWT_TokenServiceClient) VerifyJWT(ctx context.Context, in *VerifyJWTRequest, opts ...grpc.CallOption) (*VerifyJWTResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyJWTResponse)
	err := c.cc.Invoke(ctx, JWT_TokenService_VerifyJWT_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jWT_TokenServiceClient) GetUserID(ctx context.Context, in *GetUserIDRequest, opts ...grpc.CallOption) (*GetUserIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserIDResponse)
	err := c.cc.Invoke(ctx, JWT_TokenService_GetUserID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JWT_TokenServiceServer is the server API for JWT_TokenService service.
// All implementations must embed UnimplementedJWT_TokenServiceServer
// for forward compatibility
type JWT_TokenServiceServer interface {
	GenerateJWt(context.Context, *GenerateRequest) (*GenerateResponse, error)
	VerifyJWT(context.Context, *VerifyJWTRequest) (*VerifyJWTResponse, error)
	GetUserID(context.Context, *GetUserIDRequest) (*GetUserIDResponse, error)
	mustEmbedUnimplementedJWT_TokenServiceServer()
}

// UnimplementedJWT_TokenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJWT_TokenServiceServer struct {
}

func (UnimplementedJWT_TokenServiceServer) GenerateJWt(context.Context, *GenerateRequest) (*GenerateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateJWt not implemented")
}
func (UnimplementedJWT_TokenServiceServer) VerifyJWT(context.Context, *VerifyJWTRequest) (*VerifyJWTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyJWT not implemented")
}
func (UnimplementedJWT_TokenServiceServer) GetUserID(context.Context, *GetUserIDRequest) (*GetUserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserID not implemented")
}
func (UnimplementedJWT_TokenServiceServer) mustEmbedUnimplementedJWT_TokenServiceServer() {}

// UnsafeJWT_TokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JWT_TokenServiceServer will
// result in compilation errors.
type UnsafeJWT_TokenServiceServer interface {
	mustEmbedUnimplementedJWT_TokenServiceServer()
}

func RegisterJWT_TokenServiceServer(s grpc.ServiceRegistrar, srv JWT_TokenServiceServer) {
	s.RegisterService(&JWT_TokenService_ServiceDesc, srv)
}

func _JWT_TokenService_GenerateJWt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JWT_TokenServiceServer).GenerateJWt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JWT_TokenService_GenerateJWt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JWT_TokenServiceServer).GenerateJWt(ctx, req.(*GenerateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JWT_TokenService_VerifyJWT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyJWTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JWT_TokenServiceServer).VerifyJWT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JWT_TokenService_VerifyJWT_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JWT_TokenServiceServer).VerifyJWT(ctx, req.(*VerifyJWTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JWT_TokenService_GetUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JWT_TokenServiceServer).GetUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JWT_TokenService_GetUserID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JWT_TokenServiceServer).GetUserID(ctx, req.(*GetUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JWT_TokenService_ServiceDesc is the grpc.ServiceDesc for JWT_TokenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JWT_TokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.JWT_TokenService",
	HandlerType: (*JWT_TokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateJWt",
			Handler:    _JWT_TokenService_GenerateJWt_Handler,
		},
		{
			MethodName: "VerifyJWT",
			Handler:    _JWT_TokenService_VerifyJWT_Handler,
		},
		{
			MethodName: "GetUserID",
			Handler:    _JWT_TokenService_GetUserID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}

const (
	UserAuthService_UserAuthRequired_FullMethodName = "/auth.UserAuthService/UserAuthRequired"
)

// UserAuthServiceClient is the client API for UserAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAuthServiceClient interface {
	UserAuthRequired(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
}

type userAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAuthServiceClient(cc grpc.ClientConnInterface) UserAuthServiceClient {
	return &userAuthServiceClient{cc}
}

func (c *userAuthServiceClient) UserAuthRequired(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, UserAuthService_UserAuthRequired_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAuthServiceServer is the server API for UserAuthService service.
// All implementations must embed UnimplementedUserAuthServiceServer
// for forward compatibility
type UserAuthServiceServer interface {
	UserAuthRequired(context.Context, *AuthRequest) (*AuthResponse, error)
	mustEmbedUnimplementedUserAuthServiceServer()
}

// UnimplementedUserAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserAuthServiceServer struct {
}

func (UnimplementedUserAuthServiceServer) UserAuthRequired(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAuthRequired not implemented")
}
func (UnimplementedUserAuthServiceServer) mustEmbedUnimplementedUserAuthServiceServer() {}

// UnsafeUserAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAuthServiceServer will
// result in compilation errors.
type UnsafeUserAuthServiceServer interface {
	mustEmbedUnimplementedUserAuthServiceServer()
}

func RegisterUserAuthServiceServer(s grpc.ServiceRegistrar, srv UserAuthServiceServer) {
	s.RegisterService(&UserAuthService_ServiceDesc, srv)
}

func _UserAuthService_UserAuthRequired_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthServiceServer).UserAuthRequired(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAuthService_UserAuthRequired_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthServiceServer).UserAuthRequired(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAuthService_ServiceDesc is the grpc.ServiceDesc for UserAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.UserAuthService",
	HandlerType: (*UserAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserAuthRequired",
			Handler:    _UserAuthService_UserAuthRequired_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}

const (
	AdminAuthService_AdminAuthRequired_FullMethodName = "/auth.AdminAuthService/AdminAuthRequired"
)

// AdminAuthServiceClient is the client API for AdminAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminAuthServiceClient interface {
	AdminAuthRequired(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
}

type adminAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminAuthServiceClient(cc grpc.ClientConnInterface) AdminAuthServiceClient {
	return &adminAuthServiceClient{cc}
}

func (c *adminAuthServiceClient) AdminAuthRequired(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, AdminAuthService_AdminAuthRequired_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminAuthServiceServer is the server API for AdminAuthService service.
// All implementations must embed UnimplementedAdminAuthServiceServer
// for forward compatibility
type AdminAuthServiceServer interface {
	AdminAuthRequired(context.Context, *AuthRequest) (*AuthResponse, error)
	mustEmbedUnimplementedAdminAuthServiceServer()
}

// UnimplementedAdminAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminAuthServiceServer struct {
}

func (UnimplementedAdminAuthServiceServer) AdminAuthRequired(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminAuthRequired not implemented")
}
func (UnimplementedAdminAuthServiceServer) mustEmbedUnimplementedAdminAuthServiceServer() {}

// UnsafeAdminAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminAuthServiceServer will
// result in compilation errors.
type UnsafeAdminAuthServiceServer interface {
	mustEmbedUnimplementedAdminAuthServiceServer()
}

func RegisterAdminAuthServiceServer(s grpc.ServiceRegistrar, srv AdminAuthServiceServer) {
	s.RegisterService(&AdminAuthService_ServiceDesc, srv)
}

func _AdminAuthService_AdminAuthRequired_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAuthServiceServer).AdminAuthRequired(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminAuthService_AdminAuthRequired_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAuthServiceServer).AdminAuthRequired(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminAuthService_ServiceDesc is the grpc.ServiceDesc for AdminAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AdminAuthService",
	HandlerType: (*AdminAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminAuthRequired",
			Handler:    _AdminAuthService_AdminAuthRequired_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}

const (
	SuperAdminAuthService_SuperAdminAuthRequired_FullMethodName = "/auth.SuperAdminAuthService/SuperAdminAuthRequired"
)

// SuperAdminAuthServiceClient is the client API for SuperAdminAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SuperAdminAuthServiceClient interface {
	SuperAdminAuthRequired(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
}

type superAdminAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSuperAdminAuthServiceClient(cc grpc.ClientConnInterface) SuperAdminAuthServiceClient {
	return &superAdminAuthServiceClient{cc}
}

func (c *superAdminAuthServiceClient) SuperAdminAuthRequired(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, SuperAdminAuthService_SuperAdminAuthRequired_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SuperAdminAuthServiceServer is the server API for SuperAdminAuthService service.
// All implementations must embed UnimplementedSuperAdminAuthServiceServer
// for forward compatibility
type SuperAdminAuthServiceServer interface {
	SuperAdminAuthRequired(context.Context, *AuthRequest) (*AuthResponse, error)
	mustEmbedUnimplementedSuperAdminAuthServiceServer()
}

// UnimplementedSuperAdminAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSuperAdminAuthServiceServer struct {
}

func (UnimplementedSuperAdminAuthServiceServer) SuperAdminAuthRequired(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuperAdminAuthRequired not implemented")
}
func (UnimplementedSuperAdminAuthServiceServer) mustEmbedUnimplementedSuperAdminAuthServiceServer() {}

// UnsafeSuperAdminAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SuperAdminAuthServiceServer will
// result in compilation errors.
type UnsafeSuperAdminAuthServiceServer interface {
	mustEmbedUnimplementedSuperAdminAuthServiceServer()
}

func RegisterSuperAdminAuthServiceServer(s grpc.ServiceRegistrar, srv SuperAdminAuthServiceServer) {
	s.RegisterService(&SuperAdminAuthService_ServiceDesc, srv)
}

func _SuperAdminAuthService_SuperAdminAuthRequired_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuperAdminAuthServiceServer).SuperAdminAuthRequired(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SuperAdminAuthService_SuperAdminAuthRequired_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuperAdminAuthServiceServer).SuperAdminAuthRequired(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SuperAdminAuthService_ServiceDesc is the grpc.ServiceDesc for SuperAdminAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SuperAdminAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.SuperAdminAuthService",
	HandlerType: (*SuperAdminAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SuperAdminAuthRequired",
			Handler:    _SuperAdminAuthService_SuperAdminAuthRequired_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}
