// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg/msg_comm.proto

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

// 同步类型数量
type NUM struct {
	NumType              int32    `protobuf:"varint,1,opt,name=num_type,json=numType,proto3" json:"num_type,omitempty"`
	Num                  int64    `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	Data1                int64    `protobuf:"varint,3,opt,name=data1,proto3" json:"data1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NUM) Reset()         { *m = NUM{} }
func (m *NUM) String() string { return proto.CompactTextString(m) }
func (*NUM) ProtoMessage()    {}
func (*NUM) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4efc4e3a7641c0e, []int{0}
}

func (m *NUM) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NUM.Unmarshal(m, b)
}
func (m *NUM) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NUM.Marshal(b, m, deterministic)
}
func (m *NUM) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NUM.Merge(m, src)
}
func (m *NUM) XXX_Size() int {
	return xxx_messageInfo_NUM.Size(m)
}
func (m *NUM) XXX_DiscardUnknown() {
	xxx_messageInfo_NUM.DiscardUnknown(m)
}

var xxx_messageInfo_NUM proto.InternalMessageInfo

func (m *NUM) GetNumType() int32 {
	if m != nil {
		return m.NumType
	}
	return 0
}

func (m *NUM) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *NUM) GetData1() int64 {
	if m != nil {
		return m.Data1
	}
	return 0
}

// 属性
type PRO struct {
	ProType              int32    `protobuf:"varint,1,opt,name=pro_type,json=proType,proto3" json:"pro_type,omitempty"`
	ProValue             int32    `protobuf:"varint,2,opt,name=pro_value,json=proValue,proto3" json:"pro_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PRO) Reset()         { *m = PRO{} }
func (m *PRO) String() string { return proto.CompactTextString(m) }
func (*PRO) ProtoMessage()    {}
func (*PRO) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4efc4e3a7641c0e, []int{1}
}

func (m *PRO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PRO.Unmarshal(m, b)
}
func (m *PRO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PRO.Marshal(b, m, deterministic)
}
func (m *PRO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PRO.Merge(m, src)
}
func (m *PRO) XXX_Size() int {
	return xxx_messageInfo_PRO.Size(m)
}
func (m *PRO) XXX_DiscardUnknown() {
	xxx_messageInfo_PRO.DiscardUnknown(m)
}

var xxx_messageInfo_PRO proto.InternalMessageInfo

func (m *PRO) GetProType() int32 {
	if m != nil {
		return m.ProType
	}
	return 0
}

func (m *PRO) GetProValue() int32 {
	if m != nil {
		return m.ProValue
	}
	return 0
}

// 物品数量
type ITEM_NUM struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Num                  int64    `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ITEM_NUM) Reset()         { *m = ITEM_NUM{} }
func (m *ITEM_NUM) String() string { return proto.CompactTextString(m) }
func (*ITEM_NUM) ProtoMessage()    {}
func (*ITEM_NUM) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4efc4e3a7641c0e, []int{2}
}

func (m *ITEM_NUM) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ITEM_NUM.Unmarshal(m, b)
}
func (m *ITEM_NUM) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ITEM_NUM.Marshal(b, m, deterministic)
}
func (m *ITEM_NUM) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ITEM_NUM.Merge(m, src)
}
func (m *ITEM_NUM) XXX_Size() int {
	return xxx_messageInfo_ITEM_NUM.Size(m)
}
func (m *ITEM_NUM) XXX_DiscardUnknown() {
	xxx_messageInfo_ITEM_NUM.DiscardUnknown(m)
}

var xxx_messageInfo_ITEM_NUM proto.InternalMessageInfo

func (m *ITEM_NUM) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ITEM_NUM) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

// 角色基础信息
type ROLE_BASE_INFO struct {
	Rid                  int32    `protobuf:"varint,1,opt,name=rid,proto3" json:"rid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ServId               int32    `protobuf:"varint,3,opt,name=serv_id,json=servId,proto3" json:"serv_id,omitempty"`
	Head                 int32    `protobuf:"varint,4,opt,name=head,proto3" json:"head,omitempty"`
	HeadFrame            int32    `protobuf:"varint,5,opt,name=head_frame,json=headFrame,proto3" json:"head_frame,omitempty"`
	RoleLv               int32    `protobuf:"varint,6,opt,name=role_lv,json=roleLv,proto3" json:"role_lv,omitempty"`
	Ce                   int32    `protobuf:"varint,7,opt,name=ce,proto3" json:"ce,omitempty"`
	VipLv                int32    `protobuf:"varint,8,opt,name=vip_lv,json=vipLv,proto3" json:"vip_lv,omitempty"`
	Guild                int32    `protobuf:"varint,9,opt,name=guild,proto3" json:"guild,omitempty"`
	GuildName            string   `protobuf:"bytes,10,opt,name=guild_name,json=guildName,proto3" json:"guild_name,omitempty"`
	PvpRank              int32    `protobuf:"varint,11,opt,name=pvp_rank,json=pvpRank,proto3" json:"pvp_rank,omitempty"`
	LastExitTs           int64    `protobuf:"varint,12,opt,name=last_exit_ts,json=lastExitTs,proto3" json:"last_exit_ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ROLE_BASE_INFO) Reset()         { *m = ROLE_BASE_INFO{} }
func (m *ROLE_BASE_INFO) String() string { return proto.CompactTextString(m) }
func (*ROLE_BASE_INFO) ProtoMessage()    {}
func (*ROLE_BASE_INFO) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4efc4e3a7641c0e, []int{3}
}

func (m *ROLE_BASE_INFO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ROLE_BASE_INFO.Unmarshal(m, b)
}
func (m *ROLE_BASE_INFO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ROLE_BASE_INFO.Marshal(b, m, deterministic)
}
func (m *ROLE_BASE_INFO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ROLE_BASE_INFO.Merge(m, src)
}
func (m *ROLE_BASE_INFO) XXX_Size() int {
	return xxx_messageInfo_ROLE_BASE_INFO.Size(m)
}
func (m *ROLE_BASE_INFO) XXX_DiscardUnknown() {
	xxx_messageInfo_ROLE_BASE_INFO.DiscardUnknown(m)
}

var xxx_messageInfo_ROLE_BASE_INFO proto.InternalMessageInfo

func (m *ROLE_BASE_INFO) GetRid() int32 {
	if m != nil {
		return m.Rid
	}
	return 0
}

func (m *ROLE_BASE_INFO) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ROLE_BASE_INFO) GetServId() int32 {
	if m != nil {
		return m.ServId
	}
	return 0
}

func (m *ROLE_BASE_INFO) GetHead() int32 {
	if m != nil {
		return m.Head
	}
	return 0
}

func (m *ROLE_BASE_INFO) GetHeadFrame() int32 {
	if m != nil {
		return m.HeadFrame
	}
	return 0
}

func (m *ROLE_BASE_INFO) GetRoleLv() int32 {
	if m != nil {
		return m.RoleLv
	}
	return 0
}

func (m *ROLE_BASE_INFO) GetCe() int32 {
	if m != nil {
		return m.Ce
	}
	return 0
}

func (m *ROLE_BASE_INFO) GetVipLv() int32 {
	if m != nil {
		return m.VipLv
	}
	return 0
}

func (m *ROLE_BASE_INFO) GetGuild() int32 {
	if m != nil {
		return m.Guild
	}
	return 0
}

func (m *ROLE_BASE_INFO) GetGuildName() string {
	if m != nil {
		return m.GuildName
	}
	return ""
}

func (m *ROLE_BASE_INFO) GetPvpRank() int32 {
	if m != nil {
		return m.PvpRank
	}
	return 0
}

func (m *ROLE_BASE_INFO) GetLastExitTs() int64 {
	if m != nil {
		return m.LastExitTs
	}
	return 0
}

// 多语言参数
type CHARACTERS_PARAM struct {
	ParamType            int32    `protobuf:"varint,1,opt,name=param_type,json=paramType,proto3" json:"param_type,omitempty"`
	ParamValue           string   `protobuf:"bytes,2,opt,name=param_value,json=paramValue,proto3" json:"param_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CHARACTERS_PARAM) Reset()         { *m = CHARACTERS_PARAM{} }
func (m *CHARACTERS_PARAM) String() string { return proto.CompactTextString(m) }
func (*CHARACTERS_PARAM) ProtoMessage()    {}
func (*CHARACTERS_PARAM) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4efc4e3a7641c0e, []int{4}
}

func (m *CHARACTERS_PARAM) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CHARACTERS_PARAM.Unmarshal(m, b)
}
func (m *CHARACTERS_PARAM) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CHARACTERS_PARAM.Marshal(b, m, deterministic)
}
func (m *CHARACTERS_PARAM) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CHARACTERS_PARAM.Merge(m, src)
}
func (m *CHARACTERS_PARAM) XXX_Size() int {
	return xxx_messageInfo_CHARACTERS_PARAM.Size(m)
}
func (m *CHARACTERS_PARAM) XXX_DiscardUnknown() {
	xxx_messageInfo_CHARACTERS_PARAM.DiscardUnknown(m)
}

var xxx_messageInfo_CHARACTERS_PARAM proto.InternalMessageInfo

func (m *CHARACTERS_PARAM) GetParamType() int32 {
	if m != nil {
		return m.ParamType
	}
	return 0
}

func (m *CHARACTERS_PARAM) GetParamValue() string {
	if m != nil {
		return m.ParamValue
	}
	return ""
}

type S2C_NOTIFY_UPDATE struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2C_NOTIFY_UPDATE) Reset()         { *m = S2C_NOTIFY_UPDATE{} }
func (m *S2C_NOTIFY_UPDATE) String() string { return proto.CompactTextString(m) }
func (*S2C_NOTIFY_UPDATE) ProtoMessage()    {}
func (*S2C_NOTIFY_UPDATE) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4efc4e3a7641c0e, []int{5}
}

func (m *S2C_NOTIFY_UPDATE) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C_NOTIFY_UPDATE.Unmarshal(m, b)
}
func (m *S2C_NOTIFY_UPDATE) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C_NOTIFY_UPDATE.Marshal(b, m, deterministic)
}
func (m *S2C_NOTIFY_UPDATE) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_NOTIFY_UPDATE.Merge(m, src)
}
func (m *S2C_NOTIFY_UPDATE) XXX_Size() int {
	return xxx_messageInfo_S2C_NOTIFY_UPDATE.Size(m)
}
func (m *S2C_NOTIFY_UPDATE) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_NOTIFY_UPDATE.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_NOTIFY_UPDATE proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NUM)(nil), "msg.NUM")
	proto.RegisterType((*PRO)(nil), "msg.PRO")
	proto.RegisterType((*ITEM_NUM)(nil), "msg.ITEM_NUM")
	proto.RegisterType((*ROLE_BASE_INFO)(nil), "msg.ROLE_BASE_INFO")
	proto.RegisterType((*CHARACTERS_PARAM)(nil), "msg.CHARACTERS_PARAM")
	proto.RegisterType((*S2C_NOTIFY_UPDATE)(nil), "msg.S2C_NOTIFY_UPDATE")
}

func init() {
	proto.RegisterFile("msg/msg_comm.proto", fileDescriptor_b4efc4e3a7641c0e)
}

var fileDescriptor_b4efc4e3a7641c0e = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0xe2, 0x3a, 0x89, 0xa7, 0x55, 0xd5, 0x6e, 0x41, 0x18, 0x21, 0x44, 0xe4, 0x53, 0x0f,
	0x88, 0x0a, 0x38, 0x73, 0x70, 0x83, 0xa3, 0x46, 0xca, 0x97, 0x36, 0x2e, 0x12, 0x5c, 0x56, 0x4b,
	0xbc, 0x18, 0x53, 0xaf, 0xbd, 0x5a, 0xdb, 0xab, 0xf6, 0x7f, 0xf1, 0x03, 0xd1, 0x8c, 0x23, 0x50,
	0xc4, 0x69, 0xdf, 0xbc, 0x99, 0x79, 0xf3, 0x34, 0xb3, 0xc0, 0x74, 0x93, 0xdf, 0xe8, 0x26, 0x17,
	0xfb, 0x5a, 0xeb, 0x77, 0xc6, 0xd6, 0x6d, 0xcd, 0x3c, 0xdd, 0xe4, 0xd1, 0x1d, 0x78, 0xeb, 0xfb,
	0x15, 0x7b, 0x09, 0x93, 0xaa, 0xd3, 0xa2, 0x7d, 0x32, 0x2a, 0x1c, 0x4c, 0x07, 0xd7, 0x3e, 0x1f,
	0x57, 0x9d, 0x4e, 0x9f, 0x8c, 0x62, 0x17, 0xe0, 0x55, 0x9d, 0x0e, 0x87, 0xd3, 0xc1, 0xb5, 0xc7,
	0x11, 0xb2, 0x67, 0xe0, 0x67, 0xb2, 0x95, 0xef, 0x43, 0x8f, 0xb8, 0x3e, 0x88, 0x3e, 0x81, 0xb7,
	0xe5, 0x1b, 0x54, 0x32, 0xb6, 0x3e, 0x52, 0x32, 0xb6, 0x26, 0xa5, 0x57, 0x10, 0x60, 0xca, 0xc9,
	0xb2, 0x53, 0xa4, 0xe7, 0x73, 0xac, 0xfd, 0x82, 0x71, 0xf4, 0x16, 0x26, 0x8b, 0x34, 0x59, 0x09,
	0x74, 0x73, 0x0e, 0xc3, 0x22, 0x3b, 0x74, 0x0f, 0x8b, 0xec, 0x7f, 0x0b, 0xd1, 0xef, 0x21, 0x9c,
	0xf3, 0xcd, 0x32, 0x11, 0xb7, 0xf1, 0x2e, 0x11, 0x8b, 0xf5, 0x7c, 0x83, 0x45, 0xf6, 0x6f, 0x17,
	0x42, 0xc6, 0xe0, 0xa4, 0x92, 0xba, 0x1f, 0x15, 0x70, 0xc2, 0xec, 0x05, 0x8c, 0x1b, 0x65, 0x9d,
	0x28, 0x32, 0x72, 0xef, 0xf3, 0x11, 0x86, 0x0b, 0x2a, 0xfe, 0xa9, 0x64, 0x16, 0x9e, 0x10, 0x4b,
	0x98, 0xbd, 0x06, 0xc0, 0x57, 0xfc, 0xb0, 0x28, 0xe3, 0x53, 0x26, 0x40, 0x66, 0x6e, 0x0f, 0x5a,
	0xb6, 0x2e, 0x95, 0x28, 0x5d, 0x38, 0xea, 0xb5, 0x30, 0x5c, 0x3a, 0xf4, 0xbf, 0x57, 0xe1, 0xb8,
	0xf7, 0xbf, 0x57, 0xec, 0x39, 0x8c, 0x5c, 0x61, 0xb0, 0x6e, 0x42, 0x9c, 0xef, 0x0a, 0xb3, 0x74,
	0xb8, 0xc7, 0xbc, 0x2b, 0xca, 0x2c, 0x0c, 0x7a, 0x96, 0x02, 0x1c, 0x4a, 0x40, 0x90, 0x77, 0x20,
	0xef, 0x01, 0x31, 0x6b, 0x1c, 0x8a, 0xfb, 0x75, 0x46, 0x58, 0x59, 0x3d, 0x84, 0xa7, 0x87, 0xfd,
	0x3a, 0xc3, 0x65, 0xf5, 0xc0, 0xa6, 0x70, 0x56, 0xca, 0xa6, 0x15, 0xea, 0xb1, 0x68, 0x45, 0xdb,
	0x84, 0x67, 0xb4, 0x2f, 0x40, 0x2e, 0x79, 0x2c, 0xda, 0xb4, 0x89, 0x38, 0x5c, 0xcc, 0xee, 0x62,
	0x1e, 0xcf, 0xd2, 0x84, 0xef, 0xc4, 0x36, 0xe6, 0xf1, 0x0a, 0xe7, 0x19, 0x69, 0xe5, 0xd1, 0xf1,
	0x03, 0x62, 0xe8, 0x68, 0x6f, 0xe0, 0xb4, 0x4f, 0xff, 0x3b, 0x5b, 0xc0, 0xfb, 0x8e, 0xfe, 0x70,
	0x57, 0x70, 0xb9, 0xfb, 0x30, 0x13, 0xeb, 0x4d, 0xba, 0x98, 0x7f, 0x15, 0xf7, 0xdb, 0xcf, 0x71,
	0x9a, 0xdc, 0x5e, 0x7d, 0xbb, 0x74, 0xb9, 0x30, 0xb6, 0xfe, 0x75, 0x43, 0x9f, 0x0d, 0xff, 0xde,
	0xf7, 0x11, 0xc1, 0x8f, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x18, 0x80, 0x0f, 0xe2, 0x8d, 0x02,
	0x00, 0x00,
}
