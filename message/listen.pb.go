// Code generated by protoc-gen-go. DO NOT EDIT.
// source: listen.proto

package message

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

type ListenResponse_Playback int32

const (
	ListenResponse_LISTENING ListenResponse_Playback = 0
	ListenResponse_ACTIVE    ListenResponse_Playback = 1
	ListenResponse_PROGRESS  ListenResponse_Playback = 2
	ListenResponse_CURRENT   ListenResponse_Playback = 3
)

var ListenResponse_Playback_name = map[int32]string{
	0: "LISTENING",
	1: "ACTIVE",
	2: "PROGRESS",
	3: "CURRENT",
}

var ListenResponse_Playback_value = map[string]int32{
	"LISTENING": 0,
	"ACTIVE":    1,
	"PROGRESS":  2,
	"CURRENT":   3,
}

func (x ListenResponse_Playback) String() string {
	return proto.EnumName(ListenResponse_Playback_name, int32(x))
}

func (ListenResponse_Playback) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f05da38a4a3e5177, []int{1, 0}
}

type ListenResponse_Change int32

const (
	ListenResponse_ACTION   ListenResponse_Change = 0
	ListenResponse_PLAYBACK ListenResponse_Change = 1
)

var ListenResponse_Change_name = map[int32]string{
	0: "ACTION",
	1: "PLAYBACK",
}

var ListenResponse_Change_value = map[string]int32{
	"ACTION":   0,
	"PLAYBACK": 1,
}

func (x ListenResponse_Change) String() string {
	return proto.EnumName(ListenResponse_Change_name, int32(x))
}

func (ListenResponse_Change) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f05da38a4a3e5177, []int{1, 1}
}

type Listen struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Listen) Reset()         { *m = Listen{} }
func (m *Listen) String() string { return proto.CompactTextString(m) }
func (*Listen) ProtoMessage()    {}
func (*Listen) Descriptor() ([]byte, []int) {
	return fileDescriptor_f05da38a4a3e5177, []int{0}
}

func (m *Listen) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listen.Unmarshal(m, b)
}
func (m *Listen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listen.Marshal(b, m, deterministic)
}
func (m *Listen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listen.Merge(m, src)
}
func (m *Listen) XXX_Size() int {
	return xxx_messageInfo_Listen.Size(m)
}
func (m *Listen) XXX_DiscardUnknown() {
	xxx_messageInfo_Listen.DiscardUnknown(m)
}

var xxx_messageInfo_Listen proto.InternalMessageInfo

func (m *Listen) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ListenResponse struct {
	UserId               string                `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Chage                ListenResponse_Change `protobuf:"varint,2,opt,name=chage,proto3,enum=message.ListenResponse_Change" json:"chage,omitempty"`
	Meta                 []byte                `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
	Ok                   bool                  `protobuf:"varint,4,opt,name=ok,proto3" json:"ok,omitempty"`
	Error                string                `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListenResponse) Reset()         { *m = ListenResponse{} }
func (m *ListenResponse) String() string { return proto.CompactTextString(m) }
func (*ListenResponse) ProtoMessage()    {}
func (*ListenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f05da38a4a3e5177, []int{1}
}

func (m *ListenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenResponse.Unmarshal(m, b)
}
func (m *ListenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenResponse.Marshal(b, m, deterministic)
}
func (m *ListenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenResponse.Merge(m, src)
}
func (m *ListenResponse) XXX_Size() int {
	return xxx_messageInfo_ListenResponse.Size(m)
}
func (m *ListenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListenResponse proto.InternalMessageInfo

func (m *ListenResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ListenResponse) GetChage() ListenResponse_Change {
	if m != nil {
		return m.Chage
	}
	return ListenResponse_ACTION
}

func (m *ListenResponse) GetMeta() []byte {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *ListenResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *ListenResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterEnum("message.ListenResponse_Playback", ListenResponse_Playback_name, ListenResponse_Playback_value)
	proto.RegisterEnum("message.ListenResponse_Change", ListenResponse_Change_name, ListenResponse_Change_value)
	proto.RegisterType((*Listen)(nil), "message.Listen")
	proto.RegisterType((*ListenResponse)(nil), "message.ListenResponse")
}

func init() {
	proto.RegisterFile("listen.proto", fileDescriptor_f05da38a4a3e5177)
}

var fileDescriptor_f05da38a4a3e5177 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4f, 0x83, 0x40,
	0x10, 0xc5, 0xbb, 0xb4, 0xfc, 0xe9, 0x88, 0x64, 0x33, 0x31, 0x91, 0x93, 0x41, 0x4e, 0x9c, 0x38,
	0xa8, 0x1f, 0x40, 0x24, 0xa4, 0x21, 0x12, 0xda, 0x2c, 0x68, 0xe2, 0xc9, 0x6c, 0xdb, 0x0d, 0x6d,
	0x68, 0xd9, 0x86, 0xc5, 0x83, 0x9f, 0x5e, 0xe3, 0xa2, 0x07, 0x0f, 0xbd, 0xcd, 0xcb, 0xfc, 0xde,
	0xbc, 0xc9, 0x03, 0xf7, 0xb0, 0x57, 0x83, 0xe8, 0xe2, 0x53, 0x2f, 0x07, 0x89, 0xf6, 0x51, 0x28,
	0xc5, 0x1b, 0x11, 0xde, 0x82, 0x55, 0xe8, 0x05, 0x5e, 0x83, 0xfd, 0xa1, 0x44, 0xff, 0xbe, 0xdf,
	0xfa, 0x24, 0x20, 0xd1, 0x9c, 0x59, 0x3f, 0x32, 0xdf, 0x86, 0x5f, 0x04, 0xbc, 0x91, 0x61, 0x42,
	0x9d, 0x64, 0xa7, 0xc4, 0x59, 0x16, 0x1f, 0xc0, 0xdc, 0xec, 0x78, 0x23, 0x7c, 0x23, 0x20, 0x91,
	0x77, 0x77, 0x13, 0xff, 0xe6, 0xc4, 0xff, 0x0f, 0xc4, 0xe9, 0x8e, 0x77, 0x8d, 0x60, 0x23, 0x8c,
	0x08, 0xb3, 0xa3, 0x18, 0xb8, 0x3f, 0x0d, 0x48, 0xe4, 0x32, 0x3d, 0xa3, 0x07, 0x86, 0x6c, 0xfd,
	0x59, 0x40, 0x22, 0x87, 0x19, 0xb2, 0xc5, 0x2b, 0x30, 0x45, 0xdf, 0xcb, 0xde, 0x37, 0x75, 0xe0,
	0x28, 0xc2, 0x47, 0x70, 0x56, 0x07, 0xfe, 0xb9, 0xe6, 0x9b, 0x16, 0x2f, 0x61, 0x5e, 0xe4, 0x55,
	0x9d, 0x95, 0x79, 0xb9, 0xa0, 0x13, 0x04, 0xb0, 0x92, 0xb4, 0xce, 0x5f, 0x33, 0x4a, 0xd0, 0x05,
	0x67, 0xc5, 0x96, 0x0b, 0x96, 0x55, 0x15, 0x35, 0xf0, 0x02, 0xec, 0xf4, 0x85, 0xb1, 0xac, 0xac,
	0xe9, 0x34, 0x0c, 0xc1, 0x1a, 0x9f, 0xf9, 0x33, 0x2c, 0x4b, 0x3a, 0xd1, 0x86, 0x22, 0x79, 0x7b,
	0x4a, 0xd2, 0x67, 0x4a, 0xd6, 0x96, 0x2e, 0xed, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xc8,
	0x4d, 0x37, 0x44, 0x01, 0x00, 0x00,
}
