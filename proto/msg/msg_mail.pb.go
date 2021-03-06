// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg/msg_mail.proto

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

////////////////////////////// START 循环数据定义 START /////////////////////////
type MAIL_ATTACHMENT struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Num                  int64    `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	Extra                []int64  `protobuf:"varint,3,rep,packed,name=extra,proto3" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MAIL_ATTACHMENT) Reset()         { *m = MAIL_ATTACHMENT{} }
func (m *MAIL_ATTACHMENT) String() string { return proto.CompactTextString(m) }
func (*MAIL_ATTACHMENT) ProtoMessage()    {}
func (*MAIL_ATTACHMENT) Descriptor() ([]byte, []int) {
	return fileDescriptor_7192a6a7ef2d4e7e, []int{0}
}

func (m *MAIL_ATTACHMENT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MAIL_ATTACHMENT.Unmarshal(m, b)
}
func (m *MAIL_ATTACHMENT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MAIL_ATTACHMENT.Marshal(b, m, deterministic)
}
func (m *MAIL_ATTACHMENT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MAIL_ATTACHMENT.Merge(m, src)
}
func (m *MAIL_ATTACHMENT) XXX_Size() int {
	return xxx_messageInfo_MAIL_ATTACHMENT.Size(m)
}
func (m *MAIL_ATTACHMENT) XXX_DiscardUnknown() {
	xxx_messageInfo_MAIL_ATTACHMENT.DiscardUnknown(m)
}

var xxx_messageInfo_MAIL_ATTACHMENT proto.InternalMessageInfo

func (m *MAIL_ATTACHMENT) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MAIL_ATTACHMENT) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *MAIL_ATTACHMENT) GetExtra() []int64 {
	if m != nil {
		return m.Extra
	}
	return nil
}

// 邮件信息
type MAIL_INFO struct {
	MailId               int64               `protobuf:"varint,1,opt,name=mail_id,json=mailId,proto3" json:"mail_id,omitempty"`
	FirstType            int32               `protobuf:"varint,2,opt,name=first_type,json=firstType,proto3" json:"first_type,omitempty"`
	SecondType           int32               `protobuf:"varint,3,opt,name=second_type,json=secondType,proto3" json:"second_type,omitempty"`
	MailTitle            string              `protobuf:"bytes,4,opt,name=mail_title,json=mailTitle,proto3" json:"mail_title,omitempty"`
	TitleParams          []*CHARACTERS_PARAM `protobuf:"bytes,5,rep,name=title_params,json=titleParams,proto3" json:"title_params,omitempty"`
	MailDesc             string              `protobuf:"bytes,6,opt,name=mail_desc,json=mailDesc,proto3" json:"mail_desc,omitempty"`
	DescParams           []*CHARACTERS_PARAM `protobuf:"bytes,7,rep,name=desc_params,json=descParams,proto3" json:"desc_params,omitempty"`
	Ts                   int64               `protobuf:"varint,8,opt,name=ts,proto3" json:"ts,omitempty"`
	Items                []*MAIL_ATTACHMENT  `protobuf:"bytes,9,rep,name=items,proto3" json:"items,omitempty"`
	IsNew                bool                `protobuf:"varint,10,opt,name=is_new,json=isNew,proto3" json:"is_new,omitempty"`
	IsItemGot            bool                `protobuf:"varint,11,opt,name=is_item_got,json=isItemGot,proto3" json:"is_item_got,omitempty"`
	IsReaded             bool                `protobuf:"varint,12,opt,name=is_readed,json=isReaded,proto3" json:"is_readed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MAIL_INFO) Reset()         { *m = MAIL_INFO{} }
func (m *MAIL_INFO) String() string { return proto.CompactTextString(m) }
func (*MAIL_INFO) ProtoMessage()    {}
func (*MAIL_INFO) Descriptor() ([]byte, []int) {
	return fileDescriptor_7192a6a7ef2d4e7e, []int{1}
}

func (m *MAIL_INFO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MAIL_INFO.Unmarshal(m, b)
}
func (m *MAIL_INFO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MAIL_INFO.Marshal(b, m, deterministic)
}
func (m *MAIL_INFO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MAIL_INFO.Merge(m, src)
}
func (m *MAIL_INFO) XXX_Size() int {
	return xxx_messageInfo_MAIL_INFO.Size(m)
}
func (m *MAIL_INFO) XXX_DiscardUnknown() {
	xxx_messageInfo_MAIL_INFO.DiscardUnknown(m)
}

var xxx_messageInfo_MAIL_INFO proto.InternalMessageInfo

func (m *MAIL_INFO) GetMailId() int64 {
	if m != nil {
		return m.MailId
	}
	return 0
}

func (m *MAIL_INFO) GetFirstType() int32 {
	if m != nil {
		return m.FirstType
	}
	return 0
}

func (m *MAIL_INFO) GetSecondType() int32 {
	if m != nil {
		return m.SecondType
	}
	return 0
}

func (m *MAIL_INFO) GetMailTitle() string {
	if m != nil {
		return m.MailTitle
	}
	return ""
}

func (m *MAIL_INFO) GetTitleParams() []*CHARACTERS_PARAM {
	if m != nil {
		return m.TitleParams
	}
	return nil
}

func (m *MAIL_INFO) GetMailDesc() string {
	if m != nil {
		return m.MailDesc
	}
	return ""
}

func (m *MAIL_INFO) GetDescParams() []*CHARACTERS_PARAM {
	if m != nil {
		return m.DescParams
	}
	return nil
}

func (m *MAIL_INFO) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *MAIL_INFO) GetItems() []*MAIL_ATTACHMENT {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *MAIL_INFO) GetIsNew() bool {
	if m != nil {
		return m.IsNew
	}
	return false
}

func (m *MAIL_INFO) GetIsItemGot() bool {
	if m != nil {
		return m.IsItemGot
	}
	return false
}

func (m *MAIL_INFO) GetIsReaded() bool {
	if m != nil {
		return m.IsReaded
	}
	return false
}

////////////////////////////// START 主动下发协议 START //////////////////////////
// 同步邮件
type S2C_SYNC_MAILS struct {
	Mails                []*MAIL_INFO `protobuf:"bytes,1,rep,name=mails,proto3" json:"mails,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *S2C_SYNC_MAILS) Reset()         { *m = S2C_SYNC_MAILS{} }
func (m *S2C_SYNC_MAILS) String() string { return proto.CompactTextString(m) }
func (*S2C_SYNC_MAILS) ProtoMessage()    {}
func (*S2C_SYNC_MAILS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7192a6a7ef2d4e7e, []int{2}
}

func (m *S2C_SYNC_MAILS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C_SYNC_MAILS.Unmarshal(m, b)
}
func (m *S2C_SYNC_MAILS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C_SYNC_MAILS.Marshal(b, m, deterministic)
}
func (m *S2C_SYNC_MAILS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_SYNC_MAILS.Merge(m, src)
}
func (m *S2C_SYNC_MAILS) XXX_Size() int {
	return xxx_messageInfo_S2C_SYNC_MAILS.Size(m)
}
func (m *S2C_SYNC_MAILS) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_SYNC_MAILS.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_SYNC_MAILS proto.InternalMessageInfo

func (m *S2C_SYNC_MAILS) GetMails() []*MAIL_INFO {
	if m != nil {
		return m.Mails
	}
	return nil
}

// 邮件
type C2S_MAIL struct {
	Lan                  string   `protobuf:"bytes,1,opt,name=lan,proto3" json:"lan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S_MAIL) Reset()         { *m = C2S_MAIL{} }
func (m *C2S_MAIL) String() string { return proto.CompactTextString(m) }
func (*C2S_MAIL) ProtoMessage()    {}
func (*C2S_MAIL) Descriptor() ([]byte, []int) {
	return fileDescriptor_7192a6a7ef2d4e7e, []int{3}
}

func (m *C2S_MAIL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S_MAIL.Unmarshal(m, b)
}
func (m *C2S_MAIL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S_MAIL.Marshal(b, m, deterministic)
}
func (m *C2S_MAIL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S_MAIL.Merge(m, src)
}
func (m *C2S_MAIL) XXX_Size() int {
	return xxx_messageInfo_C2S_MAIL.Size(m)
}
func (m *C2S_MAIL) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S_MAIL.DiscardUnknown(m)
}

var xxx_messageInfo_C2S_MAIL proto.InternalMessageInfo

func (m *C2S_MAIL) GetLan() string {
	if m != nil {
		return m.Lan
	}
	return ""
}

type S2C_MAIL struct {
	Mails                []*MAIL_INFO `protobuf:"bytes,1,rep,name=mails,proto3" json:"mails,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *S2C_MAIL) Reset()         { *m = S2C_MAIL{} }
func (m *S2C_MAIL) String() string { return proto.CompactTextString(m) }
func (*S2C_MAIL) ProtoMessage()    {}
func (*S2C_MAIL) Descriptor() ([]byte, []int) {
	return fileDescriptor_7192a6a7ef2d4e7e, []int{4}
}

func (m *S2C_MAIL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C_MAIL.Unmarshal(m, b)
}
func (m *S2C_MAIL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C_MAIL.Marshal(b, m, deterministic)
}
func (m *S2C_MAIL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_MAIL.Merge(m, src)
}
func (m *S2C_MAIL) XXX_Size() int {
	return xxx_messageInfo_S2C_MAIL.Size(m)
}
func (m *S2C_MAIL) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_MAIL.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_MAIL proto.InternalMessageInfo

func (m *S2C_MAIL) GetMails() []*MAIL_INFO {
	if m != nil {
		return m.Mails
	}
	return nil
}

// 删除邮件
type C2S_MAIL_DEL struct {
	Ids                  []int64  `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S_MAIL_DEL) Reset()         { *m = C2S_MAIL_DEL{} }
func (m *C2S_MAIL_DEL) String() string { return proto.CompactTextString(m) }
func (*C2S_MAIL_DEL) ProtoMessage()    {}
func (*C2S_MAIL_DEL) Descriptor() ([]byte, []int) {
	return fileDescriptor_7192a6a7ef2d4e7e, []int{5}
}

func (m *C2S_MAIL_DEL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S_MAIL_DEL.Unmarshal(m, b)
}
func (m *C2S_MAIL_DEL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S_MAIL_DEL.Marshal(b, m, deterministic)
}
func (m *C2S_MAIL_DEL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S_MAIL_DEL.Merge(m, src)
}
func (m *C2S_MAIL_DEL) XXX_Size() int {
	return xxx_messageInfo_C2S_MAIL_DEL.Size(m)
}
func (m *C2S_MAIL_DEL) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S_MAIL_DEL.DiscardUnknown(m)
}

var xxx_messageInfo_C2S_MAIL_DEL proto.InternalMessageInfo

func (m *C2S_MAIL_DEL) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type S2C_MAIL_DEL struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2C_MAIL_DEL) Reset()         { *m = S2C_MAIL_DEL{} }
func (m *S2C_MAIL_DEL) String() string { return proto.CompactTextString(m) }
func (*S2C_MAIL_DEL) ProtoMessage()    {}
func (*S2C_MAIL_DEL) Descriptor() ([]byte, []int) {
	return fileDescriptor_7192a6a7ef2d4e7e, []int{6}
}

func (m *S2C_MAIL_DEL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C_MAIL_DEL.Unmarshal(m, b)
}
func (m *S2C_MAIL_DEL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C_MAIL_DEL.Marshal(b, m, deterministic)
}
func (m *S2C_MAIL_DEL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_MAIL_DEL.Merge(m, src)
}
func (m *S2C_MAIL_DEL) XXX_Size() int {
	return xxx_messageInfo_S2C_MAIL_DEL.Size(m)
}
func (m *S2C_MAIL_DEL) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_MAIL_DEL.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_MAIL_DEL proto.InternalMessageInfo

// 领取邮件附件
type C2S_MAIL_ATTACHMENTS struct {
	Ids                  []int64  `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S_MAIL_ATTACHMENTS) Reset()         { *m = C2S_MAIL_ATTACHMENTS{} }
func (m *C2S_MAIL_ATTACHMENTS) String() string { return proto.CompactTextString(m) }
func (*C2S_MAIL_ATTACHMENTS) ProtoMessage()    {}
func (*C2S_MAIL_ATTACHMENTS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7192a6a7ef2d4e7e, []int{7}
}

func (m *C2S_MAIL_ATTACHMENTS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S_MAIL_ATTACHMENTS.Unmarshal(m, b)
}
func (m *C2S_MAIL_ATTACHMENTS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S_MAIL_ATTACHMENTS.Marshal(b, m, deterministic)
}
func (m *C2S_MAIL_ATTACHMENTS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S_MAIL_ATTACHMENTS.Merge(m, src)
}
func (m *C2S_MAIL_ATTACHMENTS) XXX_Size() int {
	return xxx_messageInfo_C2S_MAIL_ATTACHMENTS.Size(m)
}
func (m *C2S_MAIL_ATTACHMENTS) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S_MAIL_ATTACHMENTS.DiscardUnknown(m)
}

var xxx_messageInfo_C2S_MAIL_ATTACHMENTS proto.InternalMessageInfo

func (m *C2S_MAIL_ATTACHMENTS) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type S2C_MAIL_ATTACHMENTS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2C_MAIL_ATTACHMENTS) Reset()         { *m = S2C_MAIL_ATTACHMENTS{} }
func (m *S2C_MAIL_ATTACHMENTS) String() string { return proto.CompactTextString(m) }
func (*S2C_MAIL_ATTACHMENTS) ProtoMessage()    {}
func (*S2C_MAIL_ATTACHMENTS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7192a6a7ef2d4e7e, []int{8}
}

func (m *S2C_MAIL_ATTACHMENTS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C_MAIL_ATTACHMENTS.Unmarshal(m, b)
}
func (m *S2C_MAIL_ATTACHMENTS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C_MAIL_ATTACHMENTS.Marshal(b, m, deterministic)
}
func (m *S2C_MAIL_ATTACHMENTS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_MAIL_ATTACHMENTS.Merge(m, src)
}
func (m *S2C_MAIL_ATTACHMENTS) XXX_Size() int {
	return xxx_messageInfo_S2C_MAIL_ATTACHMENTS.Size(m)
}
func (m *S2C_MAIL_ATTACHMENTS) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_MAIL_ATTACHMENTS.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_MAIL_ATTACHMENTS proto.InternalMessageInfo

// 已读邮件
type C2S_MAIL_READ_REQ struct {
	Ids                  []int64  `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S_MAIL_READ_REQ) Reset()         { *m = C2S_MAIL_READ_REQ{} }
func (m *C2S_MAIL_READ_REQ) String() string { return proto.CompactTextString(m) }
func (*C2S_MAIL_READ_REQ) ProtoMessage()    {}
func (*C2S_MAIL_READ_REQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_7192a6a7ef2d4e7e, []int{9}
}

func (m *C2S_MAIL_READ_REQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S_MAIL_READ_REQ.Unmarshal(m, b)
}
func (m *C2S_MAIL_READ_REQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S_MAIL_READ_REQ.Marshal(b, m, deterministic)
}
func (m *C2S_MAIL_READ_REQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S_MAIL_READ_REQ.Merge(m, src)
}
func (m *C2S_MAIL_READ_REQ) XXX_Size() int {
	return xxx_messageInfo_C2S_MAIL_READ_REQ.Size(m)
}
func (m *C2S_MAIL_READ_REQ) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S_MAIL_READ_REQ.DiscardUnknown(m)
}

var xxx_messageInfo_C2S_MAIL_READ_REQ proto.InternalMessageInfo

func (m *C2S_MAIL_READ_REQ) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type S2C_MAIL_READ_RET struct {
	Ids                  []int64  `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2C_MAIL_READ_RET) Reset()         { *m = S2C_MAIL_READ_RET{} }
func (m *S2C_MAIL_READ_RET) String() string { return proto.CompactTextString(m) }
func (*S2C_MAIL_READ_RET) ProtoMessage()    {}
func (*S2C_MAIL_READ_RET) Descriptor() ([]byte, []int) {
	return fileDescriptor_7192a6a7ef2d4e7e, []int{10}
}

func (m *S2C_MAIL_READ_RET) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C_MAIL_READ_RET.Unmarshal(m, b)
}
func (m *S2C_MAIL_READ_RET) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C_MAIL_READ_RET.Marshal(b, m, deterministic)
}
func (m *S2C_MAIL_READ_RET) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_MAIL_READ_RET.Merge(m, src)
}
func (m *S2C_MAIL_READ_RET) XXX_Size() int {
	return xxx_messageInfo_S2C_MAIL_READ_RET.Size(m)
}
func (m *S2C_MAIL_READ_RET) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_MAIL_READ_RET.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_MAIL_READ_RET proto.InternalMessageInfo

func (m *S2C_MAIL_READ_RET) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*MAIL_ATTACHMENT)(nil), "msg.MAIL_ATTACHMENT")
	proto.RegisterType((*MAIL_INFO)(nil), "msg.MAIL_INFO")
	proto.RegisterType((*S2C_SYNC_MAILS)(nil), "msg.S2C_SYNC_MAILS")
	proto.RegisterType((*C2S_MAIL)(nil), "msg.C2S_MAIL")
	proto.RegisterType((*S2C_MAIL)(nil), "msg.S2C_MAIL")
	proto.RegisterType((*C2S_MAIL_DEL)(nil), "msg.C2S_MAIL_DEL")
	proto.RegisterType((*S2C_MAIL_DEL)(nil), "msg.S2C_MAIL_DEL")
	proto.RegisterType((*C2S_MAIL_ATTACHMENTS)(nil), "msg.C2S_MAIL_ATTACHMENTS")
	proto.RegisterType((*S2C_MAIL_ATTACHMENTS)(nil), "msg.S2C_MAIL_ATTACHMENTS")
	proto.RegisterType((*C2S_MAIL_READ_REQ)(nil), "msg.C2S_MAIL_READ_REQ")
	proto.RegisterType((*S2C_MAIL_READ_RET)(nil), "msg.S2C_MAIL_READ_RET")
}

func init() {
	proto.RegisterFile("msg/msg_mail.proto", fileDescriptor_7192a6a7ef2d4e7e)
}

var fileDescriptor_7192a6a7ef2d4e7e = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x51, 0x6f, 0xd3, 0x3e,
	0x14, 0xc5, 0x95, 0xe6, 0x9f, 0x2e, 0xb9, 0xa9, 0xfa, 0x67, 0xa6, 0x03, 0x0b, 0x18, 0x44, 0x11,
	0x93, 0x22, 0x1e, 0x36, 0x54, 0xa4, 0x89, 0xd7, 0x90, 0x16, 0x56, 0x89, 0x96, 0xe1, 0xe4, 0x05,
	0x5e, 0xac, 0xd0, 0x98, 0xc8, 0xa8, 0x6e, 0xaa, 0xd8, 0x30, 0xf6, 0x75, 0xf8, 0xa4, 0xc8, 0x37,
	0x6d, 0x87, 0x60, 0x42, 0xbc, 0x5d, 0xdf, 0x73, 0xce, 0x2f, 0xb6, 0x6f, 0x0c, 0x44, 0xe9, 0xfa,
	0x4c, 0xe9, 0x9a, 0xab, 0x52, 0xae, 0x4e, 0x37, 0x6d, 0x63, 0x1a, 0xe2, 0x2a, 0x5d, 0x3f, 0xd8,
	0x0b, 0xcb, 0x46, 0xa9, 0x4e, 0x88, 0x67, 0xf0, 0xff, 0x3c, 0x9d, 0xbd, 0xe5, 0x69, 0x51, 0xa4,
	0xd9, 0xc5, 0x7c, 0xba, 0x28, 0xc8, 0x10, 0x7a, 0xb2, 0xa2, 0x4e, 0xe4, 0x24, 0x1e, 0xeb, 0xc9,
	0x8a, 0xdc, 0x01, 0x77, 0xfd, 0x55, 0xd1, 0x5e, 0xe4, 0x24, 0x2e, 0xb3, 0x25, 0x19, 0x81, 0x27,
	0xbe, 0x9b, 0xb6, 0xa4, 0x6e, 0xe4, 0x26, 0x2e, 0xeb, 0x16, 0xf1, 0x0f, 0x17, 0x02, 0x64, 0xcd,
	0x16, 0xaf, 0xdf, 0x91, 0xfb, 0x70, 0x60, 0xbf, 0xcf, 0xb7, 0x28, 0x97, 0xf5, 0xed, 0x72, 0x56,
	0x91, 0x63, 0x80, 0xcf, 0xb2, 0xd5, 0x86, 0x9b, 0xeb, 0x8d, 0x40, 0xaa, 0xc7, 0x02, 0xec, 0x14,
	0xd7, 0x1b, 0x41, 0x9e, 0x40, 0xa8, 0xc5, 0xb2, 0x59, 0x57, 0x9d, 0xee, 0xa2, 0x0e, 0x5d, 0x0b,
	0x0d, 0xc7, 0x00, 0x08, 0x36, 0xd2, 0xac, 0x04, 0xfd, 0x2f, 0x72, 0x92, 0x80, 0x05, 0xb6, 0x53,
	0xd8, 0x06, 0x79, 0x09, 0x03, 0x54, 0xf8, 0xa6, 0x6c, 0x4b, 0xa5, 0xa9, 0x17, 0xb9, 0x49, 0x38,
	0x3e, 0x3a, 0x55, 0xba, 0x3e, 0xcd, 0x2e, 0x52, 0x96, 0x66, 0xc5, 0x94, 0xe5, 0xfc, 0x32, 0x65,
	0xe9, 0x9c, 0x85, 0x68, 0xbd, 0x44, 0x27, 0x79, 0x08, 0x88, 0xe1, 0x95, 0xd0, 0x4b, 0xda, 0x47,
	0xae, 0x6f, 0x1b, 0x13, 0xa1, 0x97, 0xe4, 0x1c, 0x42, 0xdb, 0xdf, 0x51, 0x0f, 0xfe, 0x46, 0x05,
	0xeb, 0xdc, 0x42, 0x87, 0xd0, 0x33, 0x9a, 0xfa, 0x78, 0x03, 0x3d, 0xa3, 0xc9, 0x33, 0xf0, 0xa4,
	0x11, 0x4a, 0xd3, 0x00, 0x09, 0x23, 0x24, 0xfc, 0x36, 0x01, 0xd6, 0x59, 0xc8, 0x11, 0xf4, 0xa5,
	0xe6, 0x6b, 0x71, 0x45, 0x21, 0x72, 0x12, 0x9f, 0x79, 0x52, 0x2f, 0xc4, 0x15, 0x79, 0x0c, 0xa1,
	0xd4, 0xdc, 0x5a, 0x78, 0xdd, 0x18, 0x1a, 0xa2, 0x16, 0x48, 0x3d, 0x33, 0x42, 0xbd, 0x69, 0x8c,
	0x3d, 0x87, 0xd4, 0xbc, 0x15, 0x65, 0x25, 0x2a, 0x3a, 0x40, 0xd5, 0x97, 0x9a, 0xe1, 0x3a, 0x3e,
	0x87, 0x61, 0x3e, 0xce, 0x78, 0xfe, 0x61, 0x91, 0x71, 0xfb, 0xd9, 0x9c, 0x3c, 0x05, 0xcf, 0x9e,
	0x52, 0x53, 0x07, 0x77, 0x34, 0xbc, 0xd9, 0x91, 0x9d, 0x23, 0xeb, 0xc4, 0xf8, 0x11, 0xf8, 0xd9,
	0x38, 0xc7, 0x88, 0xfd, 0x21, 0x56, 0xe5, 0x1a, 0xc7, 0x1a, 0x30, 0x5b, 0xc6, 0xcf, 0xc1, 0xb7,
	0x54, 0x54, 0xff, 0x8d, 0x17, 0xc1, 0x60, 0xc7, 0xe3, 0x93, 0x29, 0x32, 0x65, 0xd5, 0x65, 0x5c,
	0x66, 0xcb, 0x78, 0x08, 0x83, 0x1d, 0xd3, 0x3a, 0xe2, 0x04, 0x46, 0xfb, 0xc4, 0xcd, 0x5d, 0xe5,
	0xb7, 0x24, 0xef, 0xc1, 0x68, 0x9f, 0xfc, 0xc5, 0x19, 0x9f, 0xc0, 0xe1, 0x9e, 0xc0, 0xa6, 0xe9,
	0x84, 0xb3, 0xe9, 0xfb, 0x5b, 0xe2, 0x27, 0x70, 0xb8, 0x8f, 0x6f, 0x6d, 0xc5, 0x9f, 0xb6, 0x57,
	0x77, 0x3f, 0x1e, 0x7e, 0xab, 0xf9, 0xa6, 0x6d, 0xbe, 0x9c, 0xe1, 0x53, 0xb2, 0x2f, 0xeb, 0x53,
	0x1f, 0xcb, 0x17, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xd5, 0x0d, 0xdd, 0x84, 0x03, 0x00,
	0x00,
}
