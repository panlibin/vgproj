// Code generated by protoc-gen-go. DO NOT EDIT.
// source: loginrpc/msg_login.proto

package loginrpc

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

type NotifyLogout struct {
	AccountId            int64    `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	ServerId             int32    `protobuf:"varint,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	PlayerId             int64    `protobuf:"varint,3,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Combat               int64    `protobuf:"varint,5,opt,name=combat,proto3" json:"combat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyLogout) Reset()         { *m = NotifyLogout{} }
func (m *NotifyLogout) String() string { return proto.CompactTextString(m) }
func (*NotifyLogout) ProtoMessage()    {}
func (*NotifyLogout) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e0aad429e3205c, []int{0}
}

func (m *NotifyLogout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyLogout.Unmarshal(m, b)
}
func (m *NotifyLogout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyLogout.Marshal(b, m, deterministic)
}
func (m *NotifyLogout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyLogout.Merge(m, src)
}
func (m *NotifyLogout) XXX_Size() int {
	return xxx_messageInfo_NotifyLogout.Size(m)
}
func (m *NotifyLogout) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyLogout.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyLogout proto.InternalMessageInfo

func (m *NotifyLogout) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *NotifyLogout) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *NotifyLogout) GetPlayerId() int64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *NotifyLogout) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NotifyLogout) GetCombat() int64 {
	if m != nil {
		return m.Combat
	}
	return 0
}

func init() {
	proto.RegisterType((*NotifyLogout)(nil), "loginrpc.NotifyLogout")
}

func init() {
	proto.RegisterFile("loginrpc/msg_login.proto", fileDescriptor_30e0aad429e3205c)
}

var fileDescriptor_30e0aad429e3205c = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0xc9, 0x4f, 0xcf,
	0xcc, 0x2b, 0x2a, 0x48, 0xd6, 0xcf, 0x2d, 0x4e, 0x8f, 0x07, 0x73, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x38, 0x60, 0x32, 0x4a, 0x93, 0x19, 0xb9, 0x78, 0xfc, 0xf2, 0x4b, 0x32, 0xd3, 0x2a,
	0x7d, 0xf2, 0xd3, 0xf3, 0x4b, 0x4b, 0x84, 0x64, 0xb9, 0xb8, 0x12, 0x93, 0x93, 0xf3, 0x4b, 0xf3,
	0x4a, 0xe2, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x38, 0xa1, 0x22, 0x9e, 0x29,
	0x42, 0xd2, 0x5c, 0x9c, 0xc5, 0xa9, 0x45, 0x65, 0xa9, 0x45, 0x20, 0x59, 0x26, 0x05, 0x46, 0x0d,
	0xd6, 0x20, 0x0e, 0x88, 0x00, 0x44, 0xb2, 0x20, 0x27, 0xb1, 0x12, 0x22, 0xc9, 0x0c, 0xd6, 0xca,
	0x01, 0x11, 0xf0, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x51, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0xc4, 0xb8, 0xd8, 0x92, 0xf3, 0x73, 0x93, 0x12, 0x4b, 0x24,
	0x58, 0xc1, 0xaa, 0xa1, 0x3c, 0x27, 0xf1, 0x28, 0xd1, 0xb2, 0xf4, 0x82, 0xa2, 0xfc, 0x2c, 0x7d,
	0xb0, 0x7b, 0xf5, 0x61, 0xce, 0x4d, 0x62, 0x03, 0xf3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x07, 0xd6, 0xf1, 0x4f, 0xdb, 0x00, 0x00, 0x00,
}
