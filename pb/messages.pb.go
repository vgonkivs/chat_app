// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package pb

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

// Specifies command that application can handle
type Message_Command int32

const (
	Message_NONE        Message_Command = 0
	Message_LIST        Message_Command = 1
	Message_ACCOUNT     Message_Command = 2
	Message_CHANGE_NAME Message_Command = 3
	Message_PING        Message_Command = 4
	Message_EXIT        Message_Command = 5
	Message_COMMON      Message_Command = 6
)

var Message_Command_name = map[int32]string{
	0: "NONE",
	1: "LIST",
	2: "ACCOUNT",
	3: "CHANGE_NAME",
	4: "PING",
	5: "EXIT",
	6: "COMMON",
}

var Message_Command_value = map[string]int32{
	"NONE":        0,
	"LIST":        1,
	"ACCOUNT":     2,
	"CHANGE_NAME": 3,
	"PING":        4,
	"EXIT":        5,
	"COMMON":      6,
}

func (x Message_Command) String() string {
	return proto.EnumName(Message_Command_name, int32(x))
}

func (Message_Command) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0, 0}
}

type Message struct {
	From                 string          `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string          `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Command              Message_Command `protobuf:"varint,3,opt,name=command,proto3,enum=pb.Message_Command" json:"command,omitempty"`
	Content              []byte          `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Message) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Message) GetCommand() Message_Command {
	if m != nil {
		return m.Command
	}
	return Message_NONE
}

func (m *Message) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.Message_Command", Message_Command_name, Message_Command_value)
	proto.RegisterType((*Message)(nil), "pb.Message")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x40, 0x4d, 0xb6, 0x36, 0x3a, 0x2b, 0x35, 0x8c, 0x97, 0x1c, 0xcb, 0x9e, 0x7a, 0x31, 0x07,
	0xfd, 0x82, 0x12, 0xc2, 0x5a, 0x30, 0x89, 0xd4, 0x0a, 0x82, 0x07, 0xd9, 0x6a, 0xf5, 0x94, 0xa6,
	0x6c, 0xf3, 0xa7, 0xfe, 0x90, 0xa4, 0xbb, 0xbd, 0xbd, 0x79, 0xf3, 0x18, 0x18, 0x28, 0xfc, 0x30,
	0xcf, 0x87, 0xdf, 0x61, 0x96, 0xd3, 0x31, 0xc4, 0x80, 0x74, 0xea, 0x77, 0x7f, 0x04, 0x98, 0x39,
	0x69, 0x44, 0xc8, 0x7e, 0x8e, 0xc1, 0x0b, 0x52, 0x92, 0xea, 0xba, 0x5d, 0x18, 0x0b, 0xa0, 0x31,
	0x08, 0xba, 0x18, 0x1a, 0x03, 0xde, 0x03, 0xfb, 0x0a, 0xde, 0x1f, 0xc6, 0x6f, 0xb1, 0x29, 0x49,
	0x55, 0x3c, 0xdc, 0xc9, 0xa9, 0x97, 0xe7, 0x0b, 0x52, 0x9d, 0x56, 0xed, 0xda, 0xa0, 0x48, 0xf9,
	0x18, 0x87, 0x31, 0x8a, 0xac, 0x24, 0xd5, 0x4d, 0xbb, 0x8e, 0xbb, 0x0f, 0x60, 0xe7, 0x1a, 0xaf,
	0x20, 0xb3, 0xce, 0x6a, 0x7e, 0x91, 0xe8, 0xb9, 0x79, 0xed, 0x38, 0xc1, 0x2d, 0xb0, 0x5a, 0x29,
	0xf7, 0x66, 0x3b, 0x4e, 0xf1, 0x16, 0xb6, 0xea, 0xa9, 0xb6, 0x7b, 0xfd, 0x69, 0x6b, 0xa3, 0xf9,
	0x26, 0x75, 0x2f, 0x8d, 0xdd, 0xf3, 0x2c, 0x91, 0x7e, 0x6f, 0x3a, 0x7e, 0x89, 0x00, 0xb9, 0x72,
	0xc6, 0x38, 0xcb, 0xf3, 0x3e, 0x5f, 0x1e, 0x7c, 0xfc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x89,
	0xd1, 0x34, 0xf2, 0x00, 0x00, 0x00,
}
