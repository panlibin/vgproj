// Code generated by protoc-gen-go. DO NOT EDIT.
// source: globalrpc/msg_nop.proto

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

type Nop struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nop) Reset()         { *m = Nop{} }
func (m *Nop) String() string { return proto.CompactTextString(m) }
func (*Nop) ProtoMessage()    {}
func (*Nop) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6463280e6c7ed2, []int{0}
}

func (m *Nop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nop.Unmarshal(m, b)
}
func (m *Nop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nop.Marshal(b, m, deterministic)
}
func (m *Nop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nop.Merge(m, src)
}
func (m *Nop) XXX_Size() int {
	return xxx_messageInfo_Nop.Size(m)
}
func (m *Nop) XXX_DiscardUnknown() {
	xxx_messageInfo_Nop.DiscardUnknown(m)
}

var xxx_messageInfo_Nop proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Nop)(nil), "globalrpc.nop")
}

func init() {
	proto.RegisterFile("globalrpc/msg_nop.proto", fileDescriptor_4c6463280e6c7ed2)
}

var fileDescriptor_4c6463280e6c7ed2 = []byte{
	// 81 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xcf, 0xc9, 0x4f,
	0x4a, 0xcc, 0x29, 0x2a, 0x48, 0xd6, 0xcf, 0x2d, 0x4e, 0x8f, 0xcf, 0xcb, 0x2f, 0xd0, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x4b, 0x28, 0xb1, 0x72, 0x31, 0xe7, 0xe5, 0x17, 0x38, 0x49,
	0x44, 0x89, 0x95, 0xa5, 0x17, 0x14, 0xe5, 0x67, 0xe9, 0x83, 0x55, 0xe8, 0xc3, 0x15, 0x24, 0xb1,
	0x81, 0x05, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x11, 0xa3, 0x79, 0x4d, 0x00, 0x00,
	0x00,
}