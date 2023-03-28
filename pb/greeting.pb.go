// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greeting.proto

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

type GreetingServiceRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetingServiceRequest) Reset()         { *m = GreetingServiceRequest{} }
func (m *GreetingServiceRequest) String() string { return proto.CompactTextString(m) }
func (*GreetingServiceRequest) ProtoMessage()    {}
func (*GreetingServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6acac03ccd168a87, []int{0}
}

func (m *GreetingServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetingServiceRequest.Unmarshal(m, b)
}
func (m *GreetingServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetingServiceRequest.Marshal(b, m, deterministic)
}
func (m *GreetingServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetingServiceRequest.Merge(m, src)
}
func (m *GreetingServiceRequest) XXX_Size() int {
	return xxx_messageInfo_GreetingServiceRequest.Size(m)
}
func (m *GreetingServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetingServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetingServiceRequest proto.InternalMessageInfo

func (m *GreetingServiceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GreetingServiceReply struct {
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetingServiceReply) Reset()         { *m = GreetingServiceReply{} }
func (m *GreetingServiceReply) String() string { return proto.CompactTextString(m) }
func (*GreetingServiceReply) ProtoMessage()    {}
func (*GreetingServiceReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_6acac03ccd168a87, []int{1}
}

func (m *GreetingServiceReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetingServiceReply.Unmarshal(m, b)
}
func (m *GreetingServiceReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetingServiceReply.Marshal(b, m, deterministic)
}
func (m *GreetingServiceReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetingServiceReply.Merge(m, src)
}
func (m *GreetingServiceReply) XXX_Size() int {
	return xxx_messageInfo_GreetingServiceReply.Size(m)
}
func (m *GreetingServiceReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetingServiceReply.DiscardUnknown(m)
}

var xxx_messageInfo_GreetingServiceReply proto.InternalMessageInfo

func (m *GreetingServiceReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetingServiceRequest)(nil), "GreetingServiceRequest")
	proto.RegisterType((*GreetingServiceReply)(nil), "GreetingServiceReply")
}

func init() {
	proto.RegisterFile("greeting.proto", fileDescriptor_6acac03ccd168a87)
}

var fileDescriptor_6acac03ccd168a87 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2f, 0x4a, 0x4d,
	0x2d, 0xc9, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe1, 0x12, 0x73, 0x87,
	0x8a, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08,
	0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9,
	0x4a, 0x06, 0x5c, 0x22, 0x18, 0xaa, 0x0b, 0x72, 0x2a, 0x85, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b,
	0x8b, 0x13, 0xd3, 0x53, 0x25, 0x98, 0xc0, 0xca, 0x61, 0x5c, 0x23, 0x7f, 0x2e, 0x7e, 0x34, 0x1d,
	0x42, 0x36, 0x5c, 0x1c, 0x30, 0x21, 0x21, 0x71, 0x3d, 0xec, 0xb6, 0x4b, 0x89, 0xea, 0x61, 0xb3,
	0x48, 0x89, 0xc1, 0x89, 0x35, 0x8a, 0x59, 0xbf, 0x20, 0x29, 0x89, 0x0d, 0xec, 0x7c, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x85, 0x96, 0x23, 0xab, 0xd0, 0x00, 0x00, 0x00,
}