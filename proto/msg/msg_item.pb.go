// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg/msg_item.proto

package msg

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

////////////////////////////// START 循环数据定义 START //////////////////////////////
// 角色物品
type ROLE_ITEM struct {
	ItemId               int32    `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	FirstType            int32    `protobuf:"varint,2,opt,name=first_type,json=firstType,proto3" json:"first_type,omitempty"`
	SecondType           int32    `protobuf:"varint,3,opt,name=second_type,json=secondType,proto3" json:"second_type,omitempty"`
	Quality              int32    `protobuf:"varint,4,opt,name=quality,proto3" json:"quality,omitempty"`
	Num                  int64    `protobuf:"varint,5,opt,name=num,proto3" json:"num,omitempty"`
	Lucky                int32    `protobuf:"varint,6,opt,name=lucky,proto3" json:"lucky,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ROLE_ITEM) Reset()         { *m = ROLE_ITEM{} }
func (m *ROLE_ITEM) String() string { return proto.CompactTextString(m) }
func (*ROLE_ITEM) ProtoMessage()    {}
func (*ROLE_ITEM) Descriptor() ([]byte, []int) {
	return fileDescriptor_d25658ef16c8d38f, []int{0}
}

func (m *ROLE_ITEM) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ROLE_ITEM.Unmarshal(m, b)
}
func (m *ROLE_ITEM) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ROLE_ITEM.Marshal(b, m, deterministic)
}
func (m *ROLE_ITEM) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ROLE_ITEM.Merge(m, src)
}
func (m *ROLE_ITEM) XXX_Size() int {
	return xxx_messageInfo_ROLE_ITEM.Size(m)
}
func (m *ROLE_ITEM) XXX_DiscardUnknown() {
	xxx_messageInfo_ROLE_ITEM.DiscardUnknown(m)
}

var xxx_messageInfo_ROLE_ITEM proto.InternalMessageInfo

func (m *ROLE_ITEM) GetItemId() int32 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *ROLE_ITEM) GetFirstType() int32 {
	if m != nil {
		return m.FirstType
	}
	return 0
}

func (m *ROLE_ITEM) GetSecondType() int32 {
	if m != nil {
		return m.SecondType
	}
	return 0
}

func (m *ROLE_ITEM) GetQuality() int32 {
	if m != nil {
		return m.Quality
	}
	return 0
}

func (m *ROLE_ITEM) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *ROLE_ITEM) GetLucky() int32 {
	if m != nil {
		return m.Lucky
	}
	return 0
}

////////////////////////////// START 主动下发协议 START //////////////////////////
// 同步物品(增加、修改)
type S2C_SYNC_ITEMS struct {
	Items                []*ROLE_ITEM `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *S2C_SYNC_ITEMS) Reset()         { *m = S2C_SYNC_ITEMS{} }
func (m *S2C_SYNC_ITEMS) String() string { return proto.CompactTextString(m) }
func (*S2C_SYNC_ITEMS) ProtoMessage()    {}
func (*S2C_SYNC_ITEMS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d25658ef16c8d38f, []int{1}
}

func (m *S2C_SYNC_ITEMS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C_SYNC_ITEMS.Unmarshal(m, b)
}
func (m *S2C_SYNC_ITEMS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C_SYNC_ITEMS.Marshal(b, m, deterministic)
}
func (m *S2C_SYNC_ITEMS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_SYNC_ITEMS.Merge(m, src)
}
func (m *S2C_SYNC_ITEMS) XXX_Size() int {
	return xxx_messageInfo_S2C_SYNC_ITEMS.Size(m)
}
func (m *S2C_SYNC_ITEMS) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_SYNC_ITEMS.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_SYNC_ITEMS proto.InternalMessageInfo

func (m *S2C_SYNC_ITEMS) GetItems() []*ROLE_ITEM {
	if m != nil {
		return m.Items
	}
	return nil
}

// 同步物品(删除)
type S2C_SYNC_ITEMS_DEL struct {
	DelIds               []int32  `protobuf:"varint,1,rep,packed,name=del_ids,json=delIds,proto3" json:"del_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2C_SYNC_ITEMS_DEL) Reset()         { *m = S2C_SYNC_ITEMS_DEL{} }
func (m *S2C_SYNC_ITEMS_DEL) String() string { return proto.CompactTextString(m) }
func (*S2C_SYNC_ITEMS_DEL) ProtoMessage()    {}
func (*S2C_SYNC_ITEMS_DEL) Descriptor() ([]byte, []int) {
	return fileDescriptor_d25658ef16c8d38f, []int{2}
}

func (m *S2C_SYNC_ITEMS_DEL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C_SYNC_ITEMS_DEL.Unmarshal(m, b)
}
func (m *S2C_SYNC_ITEMS_DEL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C_SYNC_ITEMS_DEL.Marshal(b, m, deterministic)
}
func (m *S2C_SYNC_ITEMS_DEL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_SYNC_ITEMS_DEL.Merge(m, src)
}
func (m *S2C_SYNC_ITEMS_DEL) XXX_Size() int {
	return xxx_messageInfo_S2C_SYNC_ITEMS_DEL.Size(m)
}
func (m *S2C_SYNC_ITEMS_DEL) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_SYNC_ITEMS_DEL.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_SYNC_ITEMS_DEL proto.InternalMessageInfo

func (m *S2C_SYNC_ITEMS_DEL) GetDelIds() []int32 {
	if m != nil {
		return m.DelIds
	}
	return nil
}

// 角色物品
type C2S_ROLE_ITEMS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S_ROLE_ITEMS) Reset()         { *m = C2S_ROLE_ITEMS{} }
func (m *C2S_ROLE_ITEMS) String() string { return proto.CompactTextString(m) }
func (*C2S_ROLE_ITEMS) ProtoMessage()    {}
func (*C2S_ROLE_ITEMS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d25658ef16c8d38f, []int{3}
}

func (m *C2S_ROLE_ITEMS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S_ROLE_ITEMS.Unmarshal(m, b)
}
func (m *C2S_ROLE_ITEMS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S_ROLE_ITEMS.Marshal(b, m, deterministic)
}
func (m *C2S_ROLE_ITEMS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S_ROLE_ITEMS.Merge(m, src)
}
func (m *C2S_ROLE_ITEMS) XXX_Size() int {
	return xxx_messageInfo_C2S_ROLE_ITEMS.Size(m)
}
func (m *C2S_ROLE_ITEMS) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S_ROLE_ITEMS.DiscardUnknown(m)
}

var xxx_messageInfo_C2S_ROLE_ITEMS proto.InternalMessageInfo

type S2C_ROLE_ITEMS struct {
	Items                []*ROLE_ITEM `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *S2C_ROLE_ITEMS) Reset()         { *m = S2C_ROLE_ITEMS{} }
func (m *S2C_ROLE_ITEMS) String() string { return proto.CompactTextString(m) }
func (*S2C_ROLE_ITEMS) ProtoMessage()    {}
func (*S2C_ROLE_ITEMS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d25658ef16c8d38f, []int{4}
}

func (m *S2C_ROLE_ITEMS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C_ROLE_ITEMS.Unmarshal(m, b)
}
func (m *S2C_ROLE_ITEMS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C_ROLE_ITEMS.Marshal(b, m, deterministic)
}
func (m *S2C_ROLE_ITEMS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_ROLE_ITEMS.Merge(m, src)
}
func (m *S2C_ROLE_ITEMS) XXX_Size() int {
	return xxx_messageInfo_S2C_ROLE_ITEMS.Size(m)
}
func (m *S2C_ROLE_ITEMS) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_ROLE_ITEMS.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_ROLE_ITEMS proto.InternalMessageInfo

func (m *S2C_ROLE_ITEMS) GetItems() []*ROLE_ITEM {
	if m != nil {
		return m.Items
	}
	return nil
}

// 道具使用
type C2S_USE_PROP struct {
	ItemId               int32    `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ItemNum              int32    `protobuf:"varint,2,opt,name=item_num,json=itemNum,proto3" json:"item_num,omitempty"`
	Ext1                 int32    `protobuf:"varint,3,opt,name=ext1,proto3" json:"ext1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S_USE_PROP) Reset()         { *m = C2S_USE_PROP{} }
func (m *C2S_USE_PROP) String() string { return proto.CompactTextString(m) }
func (*C2S_USE_PROP) ProtoMessage()    {}
func (*C2S_USE_PROP) Descriptor() ([]byte, []int) {
	return fileDescriptor_d25658ef16c8d38f, []int{5}
}

func (m *C2S_USE_PROP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S_USE_PROP.Unmarshal(m, b)
}
func (m *C2S_USE_PROP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S_USE_PROP.Marshal(b, m, deterministic)
}
func (m *C2S_USE_PROP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S_USE_PROP.Merge(m, src)
}
func (m *C2S_USE_PROP) XXX_Size() int {
	return xxx_messageInfo_C2S_USE_PROP.Size(m)
}
func (m *C2S_USE_PROP) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S_USE_PROP.DiscardUnknown(m)
}

var xxx_messageInfo_C2S_USE_PROP proto.InternalMessageInfo

func (m *C2S_USE_PROP) GetItemId() int32 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *C2S_USE_PROP) GetItemNum() int32 {
	if m != nil {
		return m.ItemNum
	}
	return 0
}

func (m *C2S_USE_PROP) GetExt1() int32 {
	if m != nil {
		return m.Ext1
	}
	return 0
}

type S2C_USE_PROP struct {
	Result               []int64  `protobuf:"varint,1,rep,packed,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2C_USE_PROP) Reset()         { *m = S2C_USE_PROP{} }
func (m *S2C_USE_PROP) String() string { return proto.CompactTextString(m) }
func (*S2C_USE_PROP) ProtoMessage()    {}
func (*S2C_USE_PROP) Descriptor() ([]byte, []int) {
	return fileDescriptor_d25658ef16c8d38f, []int{6}
}

func (m *S2C_USE_PROP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C_USE_PROP.Unmarshal(m, b)
}
func (m *S2C_USE_PROP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C_USE_PROP.Marshal(b, m, deterministic)
}
func (m *S2C_USE_PROP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_USE_PROP.Merge(m, src)
}
func (m *S2C_USE_PROP) XXX_Size() int {
	return xxx_messageInfo_S2C_USE_PROP.Size(m)
}
func (m *S2C_USE_PROP) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_USE_PROP.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_USE_PROP proto.InternalMessageInfo

func (m *S2C_USE_PROP) GetResult() []int64 {
	if m != nil {
		return m.Result
	}
	return nil
}

// 礼包码使用
type C2S_USE_CGP struct {
	PackageType          int32    `protobuf:"varint,1,opt,name=package_type,json=packageType,proto3" json:"package_type,omitempty"`
	Cgp                  string   `protobuf:"bytes,2,opt,name=cgp,proto3" json:"cgp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S_USE_CGP) Reset()         { *m = C2S_USE_CGP{} }
func (m *C2S_USE_CGP) String() string { return proto.CompactTextString(m) }
func (*C2S_USE_CGP) ProtoMessage()    {}
func (*C2S_USE_CGP) Descriptor() ([]byte, []int) {
	return fileDescriptor_d25658ef16c8d38f, []int{7}
}

func (m *C2S_USE_CGP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S_USE_CGP.Unmarshal(m, b)
}
func (m *C2S_USE_CGP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S_USE_CGP.Marshal(b, m, deterministic)
}
func (m *C2S_USE_CGP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S_USE_CGP.Merge(m, src)
}
func (m *C2S_USE_CGP) XXX_Size() int {
	return xxx_messageInfo_C2S_USE_CGP.Size(m)
}
func (m *C2S_USE_CGP) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S_USE_CGP.DiscardUnknown(m)
}

var xxx_messageInfo_C2S_USE_CGP proto.InternalMessageInfo

func (m *C2S_USE_CGP) GetPackageType() int32 {
	if m != nil {
		return m.PackageType
	}
	return 0
}

func (m *C2S_USE_CGP) GetCgp() string {
	if m != nil {
		return m.Cgp
	}
	return ""
}

type S2C_USE_CGP struct {
	Code                 int32       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Items                []*ITEM_NUM `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *S2C_USE_CGP) Reset()         { *m = S2C_USE_CGP{} }
func (m *S2C_USE_CGP) String() string { return proto.CompactTextString(m) }
func (*S2C_USE_CGP) ProtoMessage()    {}
func (*S2C_USE_CGP) Descriptor() ([]byte, []int) {
	return fileDescriptor_d25658ef16c8d38f, []int{8}
}

func (m *S2C_USE_CGP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C_USE_CGP.Unmarshal(m, b)
}
func (m *S2C_USE_CGP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C_USE_CGP.Marshal(b, m, deterministic)
}
func (m *S2C_USE_CGP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_USE_CGP.Merge(m, src)
}
func (m *S2C_USE_CGP) XXX_Size() int {
	return xxx_messageInfo_S2C_USE_CGP.Size(m)
}
func (m *S2C_USE_CGP) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_USE_CGP.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_USE_CGP proto.InternalMessageInfo

func (m *S2C_USE_CGP) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *S2C_USE_CGP) GetItems() []*ITEM_NUM {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*ROLE_ITEM)(nil), "msg.ROLE_ITEM")
	proto.RegisterType((*S2C_SYNC_ITEMS)(nil), "msg.S2C_SYNC_ITEMS")
	proto.RegisterType((*S2C_SYNC_ITEMS_DEL)(nil), "msg.S2C_SYNC_ITEMS_DEL")
	proto.RegisterType((*C2S_ROLE_ITEMS)(nil), "msg.C2S_ROLE_ITEMS")
	proto.RegisterType((*S2C_ROLE_ITEMS)(nil), "msg.S2C_ROLE_ITEMS")
	proto.RegisterType((*C2S_USE_PROP)(nil), "msg.C2S_USE_PROP")
	proto.RegisterType((*S2C_USE_PROP)(nil), "msg.S2C_USE_PROP")
	proto.RegisterType((*C2S_USE_CGP)(nil), "msg.C2S_USE_CGP")
	proto.RegisterType((*S2C_USE_CGP)(nil), "msg.S2C_USE_CGP")
}

func init() {
	proto.RegisterFile("msg/msg_item.proto", fileDescriptor_d25658ef16c8d38f)
}

var fileDescriptor_d25658ef16c8d38f = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x5d, 0xab, 0xd3, 0x40,
	0x10, 0x25, 0x37, 0x4d, 0xae, 0x9d, 0xd4, 0x52, 0x57, 0xd1, 0x28, 0x88, 0x35, 0x8a, 0xf4, 0xc5,
	0x5e, 0xac, 0xe0, 0x0f, 0xb8, 0x31, 0x4a, 0xe1, 0xda, 0x96, 0x4d, 0x2b, 0xe8, 0xcb, 0x52, 0xb3,
	0x6b, 0x88, 0xcd, 0x36, 0x31, 0x9b, 0x88, 0xf9, 0x43, 0xfe, 0x4e, 0xd9, 0xd9, 0x24, 0xe2, 0x83,
	0xe0, 0xdb, 0x99, 0x39, 0xf3, 0x71, 0xf6, 0xec, 0x00, 0x91, 0x2a, 0xbd, 0x92, 0x2a, 0x65, 0x59,
	0x2d, 0xe4, 0xb2, 0xac, 0x8a, 0xba, 0x20, 0xb6, 0x54, 0xe9, 0xa3, 0x81, 0x48, 0x0a, 0xd9, 0x11,
	0xc1, 0x2f, 0x0b, 0xc6, 0x74, 0x7b, 0x13, 0xb1, 0xf5, 0x3e, 0xfa, 0x40, 0x1e, 0xc0, 0xa5, 0x6e,
	0x62, 0x19, 0xf7, 0xad, 0xb9, 0xb5, 0x70, 0xa8, 0xab, 0xc3, 0x35, 0x27, 0x8f, 0x01, 0xbe, 0x66,
	0x95, 0xaa, 0x59, 0xdd, 0x96, 0xc2, 0xbf, 0x40, 0x6e, 0x8c, 0x99, 0x7d, 0x5b, 0x0a, 0xf2, 0x04,
	0x3c, 0x25, 0x92, 0xe2, 0xcc, 0x0d, 0x6f, 0x23, 0x0f, 0x26, 0x85, 0x05, 0x3e, 0x5c, 0x7e, 0x6f,
	0x8e, 0x79, 0x56, 0xb7, 0xfe, 0x08, 0xc9, 0x3e, 0x24, 0x33, 0xb0, 0xcf, 0x8d, 0xf4, 0x9d, 0xb9,
	0xb5, 0xb0, 0xa9, 0x86, 0xe4, 0x1e, 0x38, 0x79, 0x93, 0x9c, 0x5a, 0xdf, 0xc5, 0x4a, 0x13, 0x04,
	0x6f, 0x60, 0x1a, 0xaf, 0x42, 0x16, 0x7f, 0xda, 0x84, 0xa8, 0x35, 0x26, 0xcf, 0xc1, 0xd1, 0xea,
	0x94, 0x6f, 0xcd, 0xed, 0x85, 0xb7, 0x9a, 0x2e, 0xa5, 0x4a, 0x97, 0xc3, 0x5b, 0xa8, 0x21, 0x83,
	0x97, 0x40, 0xfe, 0xee, 0x63, 0x6f, 0xa3, 0x1b, 0xfd, 0x50, 0x2e, 0x72, 0x96, 0x71, 0xd3, 0xed,
	0x50, 0x97, 0x8b, 0x7c, 0xcd, 0x55, 0x30, 0x83, 0x69, 0xb8, 0x8a, 0xd9, 0x30, 0x26, 0xee, 0x17,
	0xff, 0xc9, 0xfc, 0xe7, 0xe2, 0x8f, 0x30, 0xd1, 0x93, 0x0e, 0x71, 0xc4, 0x76, 0x74, 0xbb, 0xfb,
	0xb7, 0xb7, 0x0f, 0xe1, 0x16, 0x12, 0xda, 0x06, 0xe3, 0x2c, 0x16, 0x6e, 0x1a, 0x49, 0x08, 0x8c,
	0xc4, 0xcf, 0xfa, 0x55, 0x67, 0x28, 0xe2, 0xe0, 0x05, 0x4c, 0xb4, 0x9e, 0x61, 0xee, 0x7d, 0x70,
	0x2b, 0xa1, 0x9a, 0xbc, 0x46, 0x39, 0x36, 0xed, 0xa2, 0xe0, 0x1a, 0xbc, 0x7e, 0x7f, 0xf8, 0x7e,
	0x47, 0x9e, 0xc2, 0xa4, 0x3c, 0x26, 0xa7, 0x63, 0x2a, 0xcc, 0x1f, 0x19, 0x0d, 0x5e, 0x97, 0xc3,
	0x4f, 0x9a, 0x81, 0x9d, 0xa4, 0x25, 0x6a, 0x18, 0x53, 0x0d, 0x83, 0x77, 0xe0, 0xf5, 0xbb, 0xf4,
	0x0c, 0x02, 0xa3, 0xa4, 0xe0, 0x7d, 0x2f, 0x62, 0xf2, 0xac, 0x37, 0xe3, 0x02, 0xcd, 0xb8, 0x8d,
	0x66, 0x68, 0x1f, 0xd8, 0xe6, 0xd0, 0x7b, 0x71, 0x7d, 0xf7, 0xf3, 0x9d, 0x1f, 0x29, 0x2b, 0xab,
	0xe2, 0xdb, 0x15, 0x9e, 0x9d, 0xbe, 0xc2, 0x2f, 0x2e, 0xc2, 0xd7, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xb4, 0xd2, 0xaf, 0x72, 0xb0, 0x02, 0x00, 0x00,
}
