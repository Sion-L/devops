// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	User_Login_FullMethodName                     = "/user.User/Login"
	User_LdapSource_FullMethodName                = "/user.User/LdapSource"
	User_LdapVerify_FullMethodName                = "/user.User/LdapVerify"
	User_AddUser_FullMethodName                   = "/user.User/AddUser"
	User_DeleteUser_FullMethodName                = "/user.User/DeleteUser"
	User_GetMemberGroups_FullMethodName           = "/user.User/GetMemberGroups"
	User_AddMemberGroup_FullMethodName            = "/user.User/AddMemberGroup"
	User_DelMemberGroup_FullMethodName            = "/user.User/DelMemberGroup"
	User_GetUsersInMemberOfGroup_FullMethodName   = "/user.User/GetUsersInMemberOfGroup"
	User_AddUserToMemberOfGroup_FullMethodName    = "/user.User/AddUserToMemberOfGroup"
	User_RemoveUserToMemberOfGroup_FullMethodName = "/user.User/RemoveUserToMemberOfGroup"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	LdapSource(ctx context.Context, in *LdapSourceReq, opts ...grpc.CallOption) (*Empty, error)
	LdapVerify(ctx context.Context, in *LdapVerifyReq, opts ...grpc.CallOption) (*Empty, error)
	AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*Empty, error)
	DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*Empty, error)
	GetMemberGroups(ctx context.Context, in *GetMemberOfGroupsReq, opts ...grpc.CallOption) (*GetMemberOfGroupsResp, error)
	AddMemberGroup(ctx context.Context, in *AddMemberOfGroupReq, opts ...grpc.CallOption) (*Empty, error)
	DelMemberGroup(ctx context.Context, in *DelMemberOfGroupReq, opts ...grpc.CallOption) (*Empty, error)
	GetUsersInMemberOfGroup(ctx context.Context, in *GetUsersInMemberOfGroupReq, opts ...grpc.CallOption) (*GetUsersInMemberOfGroupResp, error)
	AddUserToMemberOfGroup(ctx context.Context, in *AddUserToMemberOfGroupReq, opts ...grpc.CallOption) (*Empty, error)
	RemoveUserToMemberOfGroup(ctx context.Context, in *RemoveUserToMemberOfGroupReq, opts ...grpc.CallOption) (*Empty, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, User_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) LdapSource(ctx context.Context, in *LdapSourceReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, User_LdapSource_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) LdapVerify(ctx context.Context, in *LdapVerifyReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, User_LdapVerify_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, User_AddUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, User_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetMemberGroups(ctx context.Context, in *GetMemberOfGroupsReq, opts ...grpc.CallOption) (*GetMemberOfGroupsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMemberOfGroupsResp)
	err := c.cc.Invoke(ctx, User_GetMemberGroups_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AddMemberGroup(ctx context.Context, in *AddMemberOfGroupReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, User_AddMemberGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DelMemberGroup(ctx context.Context, in *DelMemberOfGroupReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, User_DelMemberGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUsersInMemberOfGroup(ctx context.Context, in *GetUsersInMemberOfGroupReq, opts ...grpc.CallOption) (*GetUsersInMemberOfGroupResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUsersInMemberOfGroupResp)
	err := c.cc.Invoke(ctx, User_GetUsersInMemberOfGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AddUserToMemberOfGroup(ctx context.Context, in *AddUserToMemberOfGroupReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, User_AddUserToMemberOfGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RemoveUserToMemberOfGroup(ctx context.Context, in *RemoveUserToMemberOfGroupReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, User_RemoveUserToMemberOfGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility.
type UserServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	LdapSource(context.Context, *LdapSourceReq) (*Empty, error)
	LdapVerify(context.Context, *LdapVerifyReq) (*Empty, error)
	AddUser(context.Context, *AddUserReq) (*Empty, error)
	DeleteUser(context.Context, *DeleteUserReq) (*Empty, error)
	GetMemberGroups(context.Context, *GetMemberOfGroupsReq) (*GetMemberOfGroupsResp, error)
	AddMemberGroup(context.Context, *AddMemberOfGroupReq) (*Empty, error)
	DelMemberGroup(context.Context, *DelMemberOfGroupReq) (*Empty, error)
	GetUsersInMemberOfGroup(context.Context, *GetUsersInMemberOfGroupReq) (*GetUsersInMemberOfGroupResp, error)
	AddUserToMemberOfGroup(context.Context, *AddUserToMemberOfGroupReq) (*Empty, error)
	RemoveUserToMemberOfGroup(context.Context, *RemoveUserToMemberOfGroupReq) (*Empty, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServer struct{}

func (UnimplementedUserServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServer) LdapSource(context.Context, *LdapSourceReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LdapSource not implemented")
}
func (UnimplementedUserServer) LdapVerify(context.Context, *LdapVerifyReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LdapVerify not implemented")
}
func (UnimplementedUserServer) AddUser(context.Context, *AddUserReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedUserServer) DeleteUser(context.Context, *DeleteUserReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServer) GetMemberGroups(context.Context, *GetMemberOfGroupsReq) (*GetMemberOfGroupsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMemberGroups not implemented")
}
func (UnimplementedUserServer) AddMemberGroup(context.Context, *AddMemberOfGroupReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMemberGroup not implemented")
}
func (UnimplementedUserServer) DelMemberGroup(context.Context, *DelMemberOfGroupReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelMemberGroup not implemented")
}
func (UnimplementedUserServer) GetUsersInMemberOfGroup(context.Context, *GetUsersInMemberOfGroupReq) (*GetUsersInMemberOfGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersInMemberOfGroup not implemented")
}
func (UnimplementedUserServer) AddUserToMemberOfGroup(context.Context, *AddUserToMemberOfGroupReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserToMemberOfGroup not implemented")
}
func (UnimplementedUserServer) RemoveUserToMemberOfGroup(context.Context, *RemoveUserToMemberOfGroupReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUserToMemberOfGroup not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}
func (UnimplementedUserServer) testEmbeddedByValue()              {}

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

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_LdapSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LdapSourceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).LdapSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_LdapSource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).LdapSource(ctx, req.(*LdapSourceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_LdapVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LdapVerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).LdapVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_LdapVerify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).LdapVerify(ctx, req.(*LdapVerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_AddUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AddUser(ctx, req.(*AddUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteUser(ctx, req.(*DeleteUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetMemberGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemberOfGroupsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetMemberGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetMemberGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetMemberGroups(ctx, req.(*GetMemberOfGroupsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AddMemberGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMemberOfGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AddMemberGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_AddMemberGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AddMemberGroup(ctx, req.(*AddMemberOfGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DelMemberGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelMemberOfGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DelMemberGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_DelMemberGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DelMemberGroup(ctx, req.(*DelMemberOfGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUsersInMemberOfGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersInMemberOfGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUsersInMemberOfGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUsersInMemberOfGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUsersInMemberOfGroup(ctx, req.(*GetUsersInMemberOfGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AddUserToMemberOfGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserToMemberOfGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AddUserToMemberOfGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_AddUserToMemberOfGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AddUserToMemberOfGroup(ctx, req.(*AddUserToMemberOfGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RemoveUserToMemberOfGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserToMemberOfGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RemoveUserToMemberOfGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_RemoveUserToMemberOfGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RemoveUserToMemberOfGroup(ctx, req.(*RemoveUserToMemberOfGroupReq))
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
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "LdapSource",
			Handler:    _User_LdapSource_Handler,
		},
		{
			MethodName: "LdapVerify",
			Handler:    _User_LdapVerify_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _User_AddUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _User_DeleteUser_Handler,
		},
		{
			MethodName: "GetMemberGroups",
			Handler:    _User_GetMemberGroups_Handler,
		},
		{
			MethodName: "AddMemberGroup",
			Handler:    _User_AddMemberGroup_Handler,
		},
		{
			MethodName: "DelMemberGroup",
			Handler:    _User_DelMemberGroup_Handler,
		},
		{
			MethodName: "GetUsersInMemberOfGroup",
			Handler:    _User_GetUsersInMemberOfGroup_Handler,
		},
		{
			MethodName: "AddUserToMemberOfGroup",
			Handler:    _User_AddUserToMemberOfGroup_Handler,
		},
		{
			MethodName: "RemoveUserToMemberOfGroup",
			Handler:    _User_RemoveUserToMemberOfGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
