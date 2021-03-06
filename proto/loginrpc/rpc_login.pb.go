// Code generated by protoc-gen-go. DO NOT EDIT.
// source: loginrpc/rpc_login.proto

package loginrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	globalrpc "vgproj/proto/globalrpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("loginrpc/rpc_login.proto", fileDescriptor_c9faac64777933ed)
}

var fileDescriptor_c9faac64777933ed = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0xc9, 0x4f, 0xcf,
	0xcc, 0x2b, 0x2a, 0x48, 0xd6, 0x2f, 0x2a, 0x48, 0x8e, 0x07, 0x73, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x38, 0x60, 0x32, 0x52, 0xe2, 0xe9, 0x39, 0xf9, 0x49, 0x89, 0x39, 0x20, 0x45, 0xb9,
	0xc5, 0xe9, 0xf1, 0x79, 0xf9, 0x05, 0x10, 0x25, 0x52, 0xf2, 0xa8, 0x12, 0xc5, 0xa9, 0x45, 0x65,
	0xa9, 0x45, 0xf1, 0x89, 0xa5, 0x25, 0x19, 0x50, 0x05, 0x08, 0xd3, 0x41, 0xf2, 0x48, 0xa6, 0x1b,
	0x95, 0x70, 0xb1, 0xfa, 0x80, 0xb8, 0x42, 0xc6, 0x5c, 0x2c, 0x8e, 0xa5, 0x25, 0x19, 0x42, 0xd2,
	0x7a, 0x70, 0xc3, 0xf4, 0xfc, 0xf2, 0x4b, 0x32, 0xd3, 0x2a, 0x83, 0xc1, 0xc6, 0x81, 0x24, 0xa5,
	0xf8, 0x90, 0x24, 0xf3, 0xf2, 0x0b, 0x84, 0xcc, 0xb8, 0x78, 0x02, 0x72, 0x12, 0x2b, 0x53, 0x8b,
	0x7c, 0xf2, 0xd3, 0xf3, 0x4b, 0x4b, 0x84, 0xc4, 0xf4, 0x60, 0x16, 0x41, 0xf5, 0x42, 0xc4, 0xd1,
	0xf5, 0x39, 0x89, 0x47, 0x89, 0x96, 0xa5, 0x17, 0x14, 0xe5, 0x67, 0xe9, 0x83, 0x5d, 0xa1, 0x0f,
	0xd3, 0x95, 0xc4, 0x06, 0xe6, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xba, 0x95, 0x64, 0x0b,
	0x0f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LoginClient is the client API for Login service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoginClient interface {
	Auth(ctx context.Context, in *globalrpc.NotifyServerAuth, opts ...grpc.CallOption) (*globalrpc.Nop, error)
	PlayerLogout(ctx context.Context, in *NotifyLogout, opts ...grpc.CallOption) (*globalrpc.Nop, error)
}

type loginClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginClient(cc grpc.ClientConnInterface) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) Auth(ctx context.Context, in *globalrpc.NotifyServerAuth, opts ...grpc.CallOption) (*globalrpc.Nop, error) {
	out := new(globalrpc.Nop)
	err := c.cc.Invoke(ctx, "/loginrpc.Login/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) PlayerLogout(ctx context.Context, in *NotifyLogout, opts ...grpc.CallOption) (*globalrpc.Nop, error) {
	out := new(globalrpc.Nop)
	err := c.cc.Invoke(ctx, "/loginrpc.Login/PlayerLogout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServer is the server API for Login service.
type LoginServer interface {
	Auth(context.Context, *globalrpc.NotifyServerAuth) (*globalrpc.Nop, error)
	PlayerLogout(context.Context, *NotifyLogout) (*globalrpc.Nop, error)
}

// UnimplementedLoginServer can be embedded to have forward compatible implementations.
type UnimplementedLoginServer struct {
}

func (*UnimplementedLoginServer) Auth(ctx context.Context, req *globalrpc.NotifyServerAuth) (*globalrpc.Nop, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (*UnimplementedLoginServer) PlayerLogout(ctx context.Context, req *NotifyLogout) (*globalrpc.Nop, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerLogout not implemented")
}

func RegisterLoginServer(s *grpc.Server, srv LoginServer) {
	s.RegisterService(&_Login_serviceDesc, srv)
}

func _Login_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(globalrpc.NotifyServerAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loginrpc.Login/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).Auth(ctx, req.(*globalrpc.NotifyServerAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_PlayerLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyLogout)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).PlayerLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loginrpc.Login/PlayerLogout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).PlayerLogout(ctx, req.(*NotifyLogout))
	}
	return interceptor(ctx, in, info, handler)
}

var _Login_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loginrpc.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _Login_Auth_Handler,
		},
		{
			MethodName: "PlayerLogout",
			Handler:    _Login_PlayerLogout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loginrpc/rpc_login.proto",
}
