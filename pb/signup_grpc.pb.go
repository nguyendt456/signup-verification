// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: signup.proto

package signup_with_verification

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
	SignupService_Signup_FullMethodName = "/proto.SignupService/Signup"
	SignupService_Verify_FullMethodName = "/proto.SignupService/Verify"
)

// SignupServiceClient is the client API for SignupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignupServiceClient interface {
	Signup(ctx context.Context, in *User, opts ...grpc.CallOption) (*SignupReponse, error)
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
}

type signupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSignupServiceClient(cc grpc.ClientConnInterface) SignupServiceClient {
	return &signupServiceClient{cc}
}

func (c *signupServiceClient) Signup(ctx context.Context, in *User, opts ...grpc.CallOption) (*SignupReponse, error) {
	out := new(SignupReponse)
	err := c.cc.Invoke(ctx, SignupService_Signup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupServiceClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, SignupService_Verify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignupServiceServer is the server API for SignupService service.
// All implementations must embed UnimplementedSignupServiceServer
// for forward compatibility
type SignupServiceServer interface {
	Signup(context.Context, *User) (*SignupReponse, error)
	Verify(context.Context, *VerifyRequest) (*VerifyResponse, error)
	mustEmbedUnimplementedSignupServiceServer()
}

// UnimplementedSignupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSignupServiceServer struct {
}

func (UnimplementedSignupServiceServer) Signup(context.Context, *User) (*SignupReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedSignupServiceServer) Verify(context.Context, *VerifyRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedSignupServiceServer) mustEmbedUnimplementedSignupServiceServer() {}

// UnsafeSignupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignupServiceServer will
// result in compilation errors.
type UnsafeSignupServiceServer interface {
	mustEmbedUnimplementedSignupServiceServer()
}

func RegisterSignupServiceServer(s grpc.ServiceRegistrar, srv SignupServiceServer) {
	s.RegisterService(&SignupService_ServiceDesc, srv)
}

func _SignupService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SignupService_Signup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServiceServer).Signup(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignupService_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServiceServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SignupService_Verify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServiceServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SignupService_ServiceDesc is the grpc.ServiceDesc for SignupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SignupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SignupService",
	HandlerType: (*SignupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _SignupService_Signup_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _SignupService_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "signup.proto",
}
