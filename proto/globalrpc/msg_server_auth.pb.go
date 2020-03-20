// Code generated by protoc-gen-go. DO NOT EDIT.
// source: globalrpc/msg_server_auth.proto

package globalrpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type ServerInfo struct {
	ServerType           int32    `protobuf:"varint,1,opt,name=server_type,json=serverType,proto3" json:"server_type,omitempty"`
	ServerId             []int32  `protobuf:"varint,2,rep,packed,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Ip                   string   `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerInfo) Reset()         { *m = ServerInfo{} }
func (m *ServerInfo) String() string { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()    {}
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ba62431abec0cfe, []int{0}
}

func (m *ServerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerInfo.Unmarshal(m, b)
}
func (m *ServerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerInfo.Marshal(b, m, deterministic)
}
func (m *ServerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerInfo.Merge(m, src)
}
func (m *ServerInfo) XXX_Size() int {
	return xxx_messageInfo_ServerInfo.Size(m)
}
func (m *ServerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServerInfo proto.InternalMessageInfo

func (m *ServerInfo) GetServerType() int32 {
	if m != nil {
		return m.ServerType
	}
	return 0
}

func (m *ServerInfo) GetServerId() []int32 {
	if m != nil {
		return m.ServerId
	}
	return nil
}

func (m *ServerInfo) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type NotifyServerAuth struct {
	Info                 *ServerInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Token                string      `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *NotifyServerAuth) Reset()         { *m = NotifyServerAuth{} }
func (m *NotifyServerAuth) String() string { return proto.CompactTextString(m) }
func (*NotifyServerAuth) ProtoMessage()    {}
func (*NotifyServerAuth) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ba62431abec0cfe, []int{1}
}

func (m *NotifyServerAuth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyServerAuth.Unmarshal(m, b)
}
func (m *NotifyServerAuth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyServerAuth.Marshal(b, m, deterministic)
}
func (m *NotifyServerAuth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyServerAuth.Merge(m, src)
}
func (m *NotifyServerAuth) XXX_Size() int {
	return xxx_messageInfo_NotifyServerAuth.Size(m)
}
func (m *NotifyServerAuth) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyServerAuth.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyServerAuth proto.InternalMessageInfo

func (m *NotifyServerAuth) GetInfo() *ServerInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *NotifyServerAuth) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ReqServerList struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqServerList) Reset()         { *m = ReqServerList{} }
func (m *ReqServerList) String() string { return proto.CompactTextString(m) }
func (*ReqServerList) ProtoMessage()    {}
func (*ReqServerList) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ba62431abec0cfe, []int{2}
}

func (m *ReqServerList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqServerList.Unmarshal(m, b)
}
func (m *ReqServerList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqServerList.Marshal(b, m, deterministic)
}
func (m *ReqServerList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqServerList.Merge(m, src)
}
func (m *ReqServerList) XXX_Size() int {
	return xxx_messageInfo_ReqServerList.Size(m)
}
func (m *ReqServerList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqServerList.DiscardUnknown(m)
}

var xxx_messageInfo_ReqServerList proto.InternalMessageInfo

type RspServerList struct {
	List                 []*ServerInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RspServerList) Reset()         { *m = RspServerList{} }
func (m *RspServerList) String() string { return proto.CompactTextString(m) }
func (*RspServerList) ProtoMessage()    {}
func (*RspServerList) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ba62431abec0cfe, []int{3}
}

func (m *RspServerList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RspServerList.Unmarshal(m, b)
}
func (m *RspServerList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RspServerList.Marshal(b, m, deterministic)
}
func (m *RspServerList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RspServerList.Merge(m, src)
}
func (m *RspServerList) XXX_Size() int {
	return xxx_messageInfo_RspServerList.Size(m)
}
func (m *RspServerList) XXX_DiscardUnknown() {
	xxx_messageInfo_RspServerList.DiscardUnknown(m)
}

var xxx_messageInfo_RspServerList proto.InternalMessageInfo

func (m *RspServerList) GetList() []*ServerInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*ServerInfo)(nil), "globalrpc.ServerInfo")
	proto.RegisterType((*NotifyServerAuth)(nil), "globalrpc.NotifyServerAuth")
	proto.RegisterType((*ReqServerList)(nil), "globalrpc.ReqServerList")
	proto.RegisterType((*RspServerList)(nil), "globalrpc.RspServerList")
}

func init() {
	proto.RegisterFile("globalrpc/msg_server_auth.proto", fileDescriptor_4ba62431abec0cfe)
}

var fileDescriptor_4ba62431abec0cfe = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x69, 0x6b, 0xc5, 0xbe, 0x31, 0x95, 0xa0, 0x12, 0xf0, 0xb0, 0x92, 0x53, 0xbd, 0xb4,
	0xa0, 0x37, 0x6f, 0x7a, 0x1b, 0x88, 0x87, 0xcc, 0xd3, 0x2e, 0x63, 0x73, 0x69, 0x17, 0xad, 0xcd,
	0x33, 0x79, 0x1b, 0xf4, 0xbf, 0x97, 0x25, 0xa3, 0xbd, 0x79, 0xcc, 0x2f, 0x3f, 0xbe, 0xef, 0xe3,
	0xc1, 0xac, 0x69, 0xcd, 0x66, 0xdd, 0x5a, 0xfc, 0xac, 0x7e, 0x5c, 0xb3, 0x72, 0xca, 0x1e, 0x94,
	0x5d, 0xad, 0xf7, 0xb4, 0x2b, 0xd1, 0x1a, 0x32, 0x2c, 0x1b, 0x04, 0xb1, 0x04, 0x58, 0xf8, 0xff,
	0x79, 0x57, 0x1b, 0x36, 0x83, 0xc9, 0xc9, 0xa6, 0x1e, 0x15, 0x8f, 0xf2, 0xa8, 0x48, 0x25, 0x04,
	0xf4, 0xd1, 0xa3, 0x62, 0xf7, 0x90, 0x9d, 0x04, 0xbd, 0xe5, 0x71, 0x9e, 0x14, 0xa9, 0xbc, 0x08,
	0x60, 0xbe, 0x65, 0x97, 0x10, 0x6b, 0xe4, 0x49, 0x1e, 0x15, 0x99, 0x8c, 0x35, 0x8a, 0x05, 0x5c,
	0xbf, 0x1b, 0xd2, 0x75, 0x1f, 0x1a, 0x5e, 0xf6, 0xb4, 0x63, 0x0f, 0x70, 0xa6, 0xbb, 0xda, 0xf8,
	0xe8, 0xc9, 0xe3, 0x6d, 0x39, 0x2c, 0x29, 0xc7, 0x19, 0xd2, 0x2b, 0xec, 0x06, 0x52, 0x32, 0xdf,
	0xaa, 0xe3, 0xb1, 0x4f, 0x0c, 0x0f, 0x71, 0x05, 0x53, 0xa9, 0x7e, 0x83, 0xfc, 0xa6, 0x1d, 0x89,
	0x67, 0x98, 0x4a, 0x87, 0x23, 0x38, 0x56, 0xb4, 0xda, 0x11, 0x8f, 0xf2, 0xe4, 0x9f, 0x8a, 0xa3,
	0xf2, 0xca, 0x97, 0x77, 0x87, 0x06, 0xad, 0xf9, 0xaa, 0xfc, 0x61, 0xaa, 0x41, 0xdd, 0x9c, 0x7b,
	0xf0, 0xf4, 0x17, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x23, 0xc2, 0xbc, 0x4c, 0x01, 0x00, 0x00,
}
