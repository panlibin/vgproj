// Code generated by protoc-gen-go. DO NOT EDIT.
// source: masterrpc/rpc_master.proto

package masterrpc

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
	proto.RegisterFile("masterrpc/rpc_master.proto", fileDescriptor_751e83f156012cb3)
}

var fileDescriptor_751e83f156012cb3 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x40, 0x1f, 0x94, 0x81, 0x81, 0x0d, 0xcd, 0x83, 0x8e, 0x54, 0x10, 0xfd, 0x80, 0x16, 0xdc,
	0x17, 0xa8, 0x48, 0x11, 0xe6, 0x90, 0xf9, 0xe6, 0x4b, 0xb9, 0x9d, 0xd7, 0xae, 0xd2, 0x34, 0x31,
	0x49, 0x07, 0xfb, 0x26, 0x7f, 0x52, 0x9a, 0x68, 0x92, 0xae, 0xf8, 0xd8, 0x7b, 0xce, 0x3d, 0x0d,
	0x09, 0x61, 0x1c, 0xb4, 0x41, 0xa5, 0xe4, 0x26, 0x53, 0x72, 0x53, 0xb8, 0xaf, 0x54, 0x2a, 0x61,
	0x04, 0x3d, 0xf1, 0x8c, 0x5d, 0x54, 0x8d, 0x28, 0xa1, 0xe9, 0x35, 0xae, 0xab, 0xa2, 0x15, 0xd2,
	0x39, 0xec, 0x6a, 0x08, 0x34, 0xaa, 0x1d, 0xaa, 0x02, 0x3a, 0xb3, 0xfd, 0x13, 0xc2, 0x0f, 0xec,
	0x26, 0x70, 0x2c, 0x38, 0xb4, 0x50, 0xa1, 0x13, 0x6e, 0xbf, 0x8f, 0xc8, 0xe4, 0xd9, 0x3a, 0x74,
	0x41, 0x8e, 0xef, 0x3a, 0xb3, 0xa5, 0x49, 0xea, 0xab, 0xe9, 0x4a, 0x98, 0xfa, 0x63, 0xff, 0x6a,
	0xbb, 0x3d, 0x64, 0xb3, 0x08, 0xb6, 0x42, 0xd2, 0x07, 0x32, 0xcd, 0xd1, 0x38, 0x61, 0x59, 0x6b,
	0x43, 0xe7, 0x91, 0xb0, 0xc6, 0xaf, 0x40, 0xd8, 0x80, 0x68, 0x19, 0xed, 0x2c, 0xc9, 0x2c, 0x57,
	0x50, 0xbe, 0x34, 0xb0, 0x47, 0xb5, 0x02, 0x8e, 0xf4, 0x32, 0xf5, 0x07, 0xef, 0x2b, 0x43, 0xca,
	0x06, 0x54, 0xcb, 0x83, 0xdd, 0x27, 0x32, 0xed, 0x27, 0x79, 0x57, 0x37, 0xef, 0x76, 0x90, 0x8c,
	0x63, 0x1e, 0xb2, 0x64, 0xdc, 0x0a, 0x9b, 0x39, 0x39, 0x5b, 0x63, 0x83, 0xa0, 0x31, 0xea, 0xdf,
	0x44, 0x1b, 0xee, 0x7e, 0x46, 0xce, 0xe8, 0x9a, 0x1e, 0xc9, 0xe9, 0xaf, 0x14, 0xe2, 0xd7, 0xff,
	0x75, 0xc2, 0xe1, 0x0e, 0x32, 0xf7, 0xf3, 0xb7, 0xf3, 0x5d, 0x25, 0x95, 0xf8, 0xcc, 0xec, 0xeb,
	0x65, 0x3e, 0x50, 0x4e, 0xec, 0x60, 0xf1, 0x13, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xaa, 0x60, 0x62,
	0x52, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MasterClient is the client API for Master service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MasterClient interface {
	Auth(ctx context.Context, in *globalrpc.NotifyServerAuth, opts ...grpc.CallOption) (*globalrpc.Nop, error)
	GetServerList(ctx context.Context, in *globalrpc.ReqServerList, opts ...grpc.CallOption) (*globalrpc.RspServerList, error)
	GrabPlayerName(ctx context.Context, in *ReqGrabPlayerName, opts ...grpc.CallOption) (*RspGrabPlayerName, error)
	GrabGuildName(ctx context.Context, in *ReqGrabGuildName, opts ...grpc.CallOption) (*RspGrabGuildName, error)
	ReleasePlayerName(ctx context.Context, in *NotifyReleasePlayerName, opts ...grpc.CallOption) (*globalrpc.Nop, error)
	ReleaseGuildName(ctx context.Context, in *NotifyReleaseGuildName, opts ...grpc.CallOption) (*globalrpc.Nop, error)
}

type masterClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterClient(cc grpc.ClientConnInterface) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) Auth(ctx context.Context, in *globalrpc.NotifyServerAuth, opts ...grpc.CallOption) (*globalrpc.Nop, error) {
	out := new(globalrpc.Nop)
	err := c.cc.Invoke(ctx, "/masterrpc.Master/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetServerList(ctx context.Context, in *globalrpc.ReqServerList, opts ...grpc.CallOption) (*globalrpc.RspServerList, error) {
	out := new(globalrpc.RspServerList)
	err := c.cc.Invoke(ctx, "/masterrpc.Master/GetServerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GrabPlayerName(ctx context.Context, in *ReqGrabPlayerName, opts ...grpc.CallOption) (*RspGrabPlayerName, error) {
	out := new(RspGrabPlayerName)
	err := c.cc.Invoke(ctx, "/masterrpc.Master/GrabPlayerName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GrabGuildName(ctx context.Context, in *ReqGrabGuildName, opts ...grpc.CallOption) (*RspGrabGuildName, error) {
	out := new(RspGrabGuildName)
	err := c.cc.Invoke(ctx, "/masterrpc.Master/GrabGuildName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ReleasePlayerName(ctx context.Context, in *NotifyReleasePlayerName, opts ...grpc.CallOption) (*globalrpc.Nop, error) {
	out := new(globalrpc.Nop)
	err := c.cc.Invoke(ctx, "/masterrpc.Master/ReleasePlayerName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ReleaseGuildName(ctx context.Context, in *NotifyReleaseGuildName, opts ...grpc.CallOption) (*globalrpc.Nop, error) {
	out := new(globalrpc.Nop)
	err := c.cc.Invoke(ctx, "/masterrpc.Master/ReleaseGuildName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServer is the server API for Master service.
type MasterServer interface {
	Auth(context.Context, *globalrpc.NotifyServerAuth) (*globalrpc.Nop, error)
	GetServerList(context.Context, *globalrpc.ReqServerList) (*globalrpc.RspServerList, error)
	GrabPlayerName(context.Context, *ReqGrabPlayerName) (*RspGrabPlayerName, error)
	GrabGuildName(context.Context, *ReqGrabGuildName) (*RspGrabGuildName, error)
	ReleasePlayerName(context.Context, *NotifyReleasePlayerName) (*globalrpc.Nop, error)
	ReleaseGuildName(context.Context, *NotifyReleaseGuildName) (*globalrpc.Nop, error)
}

// UnimplementedMasterServer can be embedded to have forward compatible implementations.
type UnimplementedMasterServer struct {
}

func (*UnimplementedMasterServer) Auth(ctx context.Context, req *globalrpc.NotifyServerAuth) (*globalrpc.Nop, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (*UnimplementedMasterServer) GetServerList(ctx context.Context, req *globalrpc.ReqServerList) (*globalrpc.RspServerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerList not implemented")
}
func (*UnimplementedMasterServer) GrabPlayerName(ctx context.Context, req *ReqGrabPlayerName) (*RspGrabPlayerName, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrabPlayerName not implemented")
}
func (*UnimplementedMasterServer) GrabGuildName(ctx context.Context, req *ReqGrabGuildName) (*RspGrabGuildName, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrabGuildName not implemented")
}
func (*UnimplementedMasterServer) ReleasePlayerName(ctx context.Context, req *NotifyReleasePlayerName) (*globalrpc.Nop, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleasePlayerName not implemented")
}
func (*UnimplementedMasterServer) ReleaseGuildName(ctx context.Context, req *NotifyReleaseGuildName) (*globalrpc.Nop, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseGuildName not implemented")
}

func RegisterMasterServer(s *grpc.Server, srv MasterServer) {
	s.RegisterService(&_Master_serviceDesc, srv)
}

func _Master_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(globalrpc.NotifyServerAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/masterrpc.Master/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).Auth(ctx, req.(*globalrpc.NotifyServerAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetServerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(globalrpc.ReqServerList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetServerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/masterrpc.Master/GetServerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetServerList(ctx, req.(*globalrpc.ReqServerList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GrabPlayerName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGrabPlayerName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GrabPlayerName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/masterrpc.Master/GrabPlayerName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GrabPlayerName(ctx, req.(*ReqGrabPlayerName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GrabGuildName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGrabGuildName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GrabGuildName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/masterrpc.Master/GrabGuildName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GrabGuildName(ctx, req.(*ReqGrabGuildName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ReleasePlayerName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyReleasePlayerName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ReleasePlayerName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/masterrpc.Master/ReleasePlayerName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ReleasePlayerName(ctx, req.(*NotifyReleasePlayerName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ReleaseGuildName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyReleaseGuildName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ReleaseGuildName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/masterrpc.Master/ReleaseGuildName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ReleaseGuildName(ctx, req.(*NotifyReleaseGuildName))
	}
	return interceptor(ctx, in, info, handler)
}

var _Master_serviceDesc = grpc.ServiceDesc{
	ServiceName: "masterrpc.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _Master_Auth_Handler,
		},
		{
			MethodName: "GetServerList",
			Handler:    _Master_GetServerList_Handler,
		},
		{
			MethodName: "GrabPlayerName",
			Handler:    _Master_GrabPlayerName_Handler,
		},
		{
			MethodName: "GrabGuildName",
			Handler:    _Master_GrabGuildName_Handler,
		},
		{
			MethodName: "ReleasePlayerName",
			Handler:    _Master_ReleasePlayerName_Handler,
		},
		{
			MethodName: "ReleaseGuildName",
			Handler:    _Master_ReleaseGuildName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "masterrpc/rpc_master.proto",
}
