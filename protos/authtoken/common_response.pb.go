// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common_response.proto

package service_authrization

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

type Result struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Error_Message        string   `protobuf:"bytes,2,opt,name=Error_Message,json=ErrorMessage,proto3" json:"Error_Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_22ea3455d5f2bd50, []int{0}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Result) GetError_Message() string {
	if m != nil {
		return m.Error_Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Result)(nil), "service_authrization.Result")
}

func init() { proto.RegisterFile("common_response.proto", fileDescriptor_22ea3455d5f2bd50) }

var fileDescriptor_22ea3455d5f2bd50 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0x8b, 0x2f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x29, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x8d, 0x4f, 0x2c, 0x2d, 0xc9, 0x28,
	0xca, 0xac, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0x53, 0x72, 0xe4, 0x62, 0x0b, 0x4a, 0x2d, 0x2e, 0xcd,
	0x29, 0x11, 0x12, 0xe2, 0x62, 0x71, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d,
	0x02, 0xb3, 0x85, 0x94, 0xb9, 0x78, 0x5d, 0x8b, 0x8a, 0xf2, 0x8b, 0xe2, 0x7d, 0x53, 0x8b, 0x8b,
	0x13, 0xd3, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0x78, 0xc0, 0x82, 0x50, 0x31, 0x27,
	0x89, 0x55, 0x4c, 0xa2, 0x8e, 0xa5, 0x25, 0x19, 0xee, 0x45, 0x05, 0xc9, 0xc1, 0x50, 0x3b, 0x02,
	0x40, 0x56, 0x26, 0xb1, 0x81, 0x6d, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x2d, 0xb7,
	0x62, 0x92, 0x00, 0x00, 0x00,
}
