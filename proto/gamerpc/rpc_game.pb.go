// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gamerpc/rpc_game.proto

package gamerpc

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
	proto.RegisterFile("gamerpc/rpc_game.proto", fileDescriptor_c13e2594e2ad863d)
}

var fileDescriptor_c13e2594e2ad863d = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x4f, 0xcc, 0x4d,
	0x2d, 0x2a, 0x48, 0xd6, 0x2f, 0x2a, 0x48, 0x8e, 0x07, 0xb1, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0xd8, 0xa1, 0xe2, 0x52, 0xe2, 0xe9, 0x39, 0xf9, 0x49, 0x89, 0x39, 0x20, 0x25, 0xb9, 0xc5,
	0xe9, 0xf1, 0x79, 0xf9, 0x05, 0x10, 0x15, 0x52, 0xf2, 0xa8, 0x12, 0xc5, 0xa9, 0x45, 0x65, 0xa9,
	0x45, 0xf1, 0x89, 0xa5, 0x25, 0x19, 0x50, 0x05, 0x70, 0xa3, 0x41, 0xd2, 0xd9, 0x99, 0xc9, 0xd9,
	0x10, 0x71, 0xa3, 0x0c, 0x2e, 0x16, 0xf7, 0xc4, 0xdc, 0x54, 0x21, 0x63, 0x2e, 0x16, 0xc7, 0xd2,
	0x92, 0x0c, 0x21, 0x69, 0x3d, 0xb8, 0x49, 0x7a, 0x7e, 0xf9, 0x25, 0x99, 0x69, 0x95, 0xc1, 0x60,
	0xb3, 0x40, 0x92, 0x52, 0x7c, 0x48, 0x92, 0x79, 0xf9, 0x05, 0x42, 0xda, 0x5c, 0x2c, 0xde, 0x99,
	0xc9, 0xd9, 0x42, 0xc2, 0x7a, 0x50, 0xd3, 0xa1, 0x5a, 0x40, 0x82, 0xe8, 0x8a, 0x9d, 0xc4, 0xa2,
	0x44, 0xca, 0xd2, 0x0b, 0x8a, 0xf2, 0xb3, 0xf4, 0xc1, 0x36, 0xeb, 0x43, 0xb5, 0x24, 0xb1, 0x81,
	0xb9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x44, 0xca, 0x80, 0x04, 0xfd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameClient interface {
	Auth(ctx context.Context, in *globalrpc.NotifyServerAuth, opts ...grpc.CallOption) (*globalrpc.Nop, error)
	Kick(ctx context.Context, in *NotifyKick, opts ...grpc.CallOption) (*globalrpc.Nop, error)
}

type gameClient struct {
	cc grpc.ClientConnInterface
}

func NewGameClient(cc grpc.ClientConnInterface) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) Auth(ctx context.Context, in *globalrpc.NotifyServerAuth, opts ...grpc.CallOption) (*globalrpc.Nop, error) {
	out := new(globalrpc.Nop)
	err := c.cc.Invoke(ctx, "/gamerpc.Game/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) Kick(ctx context.Context, in *NotifyKick, opts ...grpc.CallOption) (*globalrpc.Nop, error) {
	out := new(globalrpc.Nop)
	err := c.cc.Invoke(ctx, "/gamerpc.Game/Kick", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServer is the server API for Game service.
type GameServer interface {
	Auth(context.Context, *globalrpc.NotifyServerAuth) (*globalrpc.Nop, error)
	Kick(context.Context, *NotifyKick) (*globalrpc.Nop, error)
}

// UnimplementedGameServer can be embedded to have forward compatible implementations.
type UnimplementedGameServer struct {
}

func (*UnimplementedGameServer) Auth(ctx context.Context, req *globalrpc.NotifyServerAuth) (*globalrpc.Nop, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (*UnimplementedGameServer) Kick(ctx context.Context, req *NotifyKick) (*globalrpc.Nop, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Kick not implemented")
}

func RegisterGameServer(s *grpc.Server, srv GameServer) {
	s.RegisterService(&_Game_serviceDesc, srv)
}

func _Game_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(globalrpc.NotifyServerAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gamerpc.Game/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Auth(ctx, req.(*globalrpc.NotifyServerAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_Kick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyKick)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Kick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gamerpc.Game/Kick",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Kick(ctx, req.(*NotifyKick))
	}
	return interceptor(ctx, in, info, handler)
}

var _Game_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gamerpc.Game",
	HandlerType: (*GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _Game_Auth_Handler,
		},
		{
			MethodName: "Kick",
			Handler:    _Game_Kick_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gamerpc/rpc_game.proto",
}