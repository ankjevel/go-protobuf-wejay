// Code generated by protoc-gen-go. DO NOT EDIT.
// source: now_playing.proto

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

type NowPlaying struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NowPlaying) Reset()         { *m = NowPlaying{} }
func (m *NowPlaying) String() string { return proto.CompactTextString(m) }
func (*NowPlaying) ProtoMessage()    {}
func (*NowPlaying) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cf275ea3da8f9b3, []int{0}
}

func (m *NowPlaying) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NowPlaying.Unmarshal(m, b)
}
func (m *NowPlaying) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NowPlaying.Marshal(b, m, deterministic)
}
func (m *NowPlaying) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NowPlaying.Merge(m, src)
}
func (m *NowPlaying) XXX_Size() int {
	return xxx_messageInfo_NowPlaying.Size(m)
}
func (m *NowPlaying) XXX_DiscardUnknown() {
	xxx_messageInfo_NowPlaying.DiscardUnknown(m)
}

var xxx_messageInfo_NowPlaying proto.InternalMessageInfo

func (m *NowPlaying) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type NowPlayingResponse struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Track                *Track   `protobuf:"bytes,2,opt,name=track,proto3" json:"track,omitempty"`
	Ok                   bool     `protobuf:"varint,3,opt,name=ok,proto3" json:"ok,omitempty"`
	Error                string   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NowPlayingResponse) Reset()         { *m = NowPlayingResponse{} }
func (m *NowPlayingResponse) String() string { return proto.CompactTextString(m) }
func (*NowPlayingResponse) ProtoMessage()    {}
func (*NowPlayingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cf275ea3da8f9b3, []int{1}
}

func (m *NowPlayingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NowPlayingResponse.Unmarshal(m, b)
}
func (m *NowPlayingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NowPlayingResponse.Marshal(b, m, deterministic)
}
func (m *NowPlayingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NowPlayingResponse.Merge(m, src)
}
func (m *NowPlayingResponse) XXX_Size() int {
	return xxx_messageInfo_NowPlayingResponse.Size(m)
}
func (m *NowPlayingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NowPlayingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NowPlayingResponse proto.InternalMessageInfo

func (m *NowPlayingResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *NowPlayingResponse) GetTrack() *Track {
	if m != nil {
		return m.Track
	}
	return nil
}

func (m *NowPlayingResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *NowPlayingResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*NowPlaying)(nil), "message.NowPlaying")
	proto.RegisterType((*NowPlayingResponse)(nil), "message.NowPlayingResponse")
}

func init() {
	proto.RegisterFile("now_playing.proto", fileDescriptor_8cf275ea3da8f9b3)
}

var fileDescriptor_8cf275ea3da8f9b3 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0xcb, 0x2f, 0x8f,
	0x2f, 0xc8, 0x49, 0xac, 0xcc, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf,
	0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x95, 0xe2, 0x2e, 0x29, 0x4a, 0x4c, 0xce, 0x86, 0x88, 0x2a,
	0xa9, 0x72, 0x71, 0xf9, 0xe5, 0x97, 0x07, 0x40, 0x54, 0x0a, 0x89, 0x73, 0xb1, 0x97, 0x16, 0xa7,
	0x16, 0xc5, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xb1, 0x81, 0xb8, 0x9e, 0x29,
	0x4a, 0x95, 0x5c, 0x42, 0x08, 0x65, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x38, 0x95,
	0x0b, 0xa9, 0x70, 0xb1, 0x82, 0x2d, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0xe2, 0xd3, 0x83,
	0xda, 0xad, 0x17, 0x02, 0x12, 0x0d, 0x82, 0x48, 0x0a, 0xf1, 0x71, 0x31, 0xe5, 0x67, 0x4b, 0x30,
	0x2b, 0x30, 0x6a, 0x70, 0x04, 0x31, 0xe5, 0x67, 0x0b, 0x89, 0x70, 0xb1, 0xa6, 0x16, 0x15, 0xe5,
	0x17, 0x49, 0xb0, 0x80, 0x0d, 0x83, 0x70, 0x92, 0xd8, 0xc0, 0x0e, 0x35, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xaf, 0x02, 0xb3, 0xe8, 0xd3, 0x00, 0x00, 0x00,
}