// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_authrization.proto

package service_authrization

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GrpcRequestToken struct {
	AppSecret            string   `protobuf:"bytes,1,opt,name=AppSecret,proto3" json:"AppSecret,omitempty"`
	ServiceSecret        string   `protobuf:"bytes,2,opt,name=ServiceSecret,proto3" json:"ServiceSecret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcRequestToken) Reset()         { *m = GrpcRequestToken{} }
func (m *GrpcRequestToken) String() string { return proto.CompactTextString(m) }
func (*GrpcRequestToken) ProtoMessage()    {}
func (*GrpcRequestToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fb8f8acb0897447, []int{0}
}

func (m *GrpcRequestToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcRequestToken.Unmarshal(m, b)
}
func (m *GrpcRequestToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcRequestToken.Marshal(b, m, deterministic)
}
func (m *GrpcRequestToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcRequestToken.Merge(m, src)
}
func (m *GrpcRequestToken) XXX_Size() int {
	return xxx_messageInfo_GrpcRequestToken.Size(m)
}
func (m *GrpcRequestToken) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcRequestToken.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcRequestToken proto.InternalMessageInfo

func (m *GrpcRequestToken) GetAppSecret() string {
	if m != nil {
		return m.AppSecret
	}
	return ""
}

func (m *GrpcRequestToken) GetServiceSecret() string {
	if m != nil {
		return m.ServiceSecret
	}
	return ""
}

type GrpcTokenInfo struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Expires              string   `protobuf:"bytes,2,opt,name=expires,proto3" json:"expires,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcTokenInfo) Reset()         { *m = GrpcTokenInfo{} }
func (m *GrpcTokenInfo) String() string { return proto.CompactTextString(m) }
func (*GrpcTokenInfo) ProtoMessage()    {}
func (*GrpcTokenInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fb8f8acb0897447, []int{1}
}

func (m *GrpcTokenInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcTokenInfo.Unmarshal(m, b)
}
func (m *GrpcTokenInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcTokenInfo.Marshal(b, m, deterministic)
}
func (m *GrpcTokenInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcTokenInfo.Merge(m, src)
}
func (m *GrpcTokenInfo) XXX_Size() int {
	return xxx_messageInfo_GrpcTokenInfo.Size(m)
}
func (m *GrpcTokenInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcTokenInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcTokenInfo proto.InternalMessageInfo

func (m *GrpcTokenInfo) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *GrpcTokenInfo) GetExpires() string {
	if m != nil {
		return m.Expires
	}
	return ""
}

type GrpcResponseToken struct {
	Result               *Result        `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Data                 *GrpcTokenInfo `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GrpcResponseToken) Reset()         { *m = GrpcResponseToken{} }
func (m *GrpcResponseToken) String() string { return proto.CompactTextString(m) }
func (*GrpcResponseToken) ProtoMessage()    {}
func (*GrpcResponseToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fb8f8acb0897447, []int{2}
}

func (m *GrpcResponseToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcResponseToken.Unmarshal(m, b)
}
func (m *GrpcResponseToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcResponseToken.Marshal(b, m, deterministic)
}
func (m *GrpcResponseToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcResponseToken.Merge(m, src)
}
func (m *GrpcResponseToken) XXX_Size() int {
	return xxx_messageInfo_GrpcResponseToken.Size(m)
}
func (m *GrpcResponseToken) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcResponseToken.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcResponseToken proto.InternalMessageInfo

func (m *GrpcResponseToken) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GrpcResponseToken) GetData() *GrpcTokenInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GrpcRequestToken)(nil), "service_authrization.GrpcRequestToken")
	proto.RegisterType((*GrpcTokenInfo)(nil), "service_authrization.GrpcTokenInfo")
	proto.RegisterType((*GrpcResponseToken)(nil), "service_authrization.GrpcResponseToken")
}

func init() { proto.RegisterFile("service_authrization.proto", fileDescriptor_2fb8f8acb0897447) }

var fileDescriptor_2fb8f8acb0897447 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0xd1, 0x4a, 0xa7, 0x16, 0x74, 0x69, 0x21, 0x84, 0x0a, 0x12, 0x45, 0x3d, 0xf5,
	0x10, 0x05, 0x8f, 0x52, 0x2f, 0xe2, 0x4d, 0xb6, 0xe2, 0x35, 0xc4, 0x38, 0x62, 0xd0, 0xee, 0xae,
	0xb3, 0x1b, 0x11, 0x4f, 0xe2, 0xc7, 0xf1, 0x53, 0xca, 0xfe, 0x09, 0x26, 0x10, 0x3d, 0xce, 0xec,
	0x6f, 0xde, 0x7b, 0x3b, 0x03, 0x89, 0x46, 0x7a, 0xab, 0x4a, 0xcc, 0x8b, 0xda, 0x3c, 0x51, 0xf5,
	0x51, 0x98, 0x4a, 0x8a, 0x85, 0x22, 0x69, 0x24, 0x9b, 0xf6, 0xbd, 0x25, 0xb3, 0x52, 0xae, 0xd7,
	0x52, 0xe4, 0x84, 0x5a, 0x49, 0xa1, 0xd1, 0xc3, 0xe9, 0x1d, 0xec, 0x5c, 0x91, 0x2a, 0x39, 0xbe,
	0xd6, 0xa8, 0xcd, 0xad, 0x7c, 0x46, 0xc1, 0xe6, 0x30, 0x5a, 0x2a, 0xb5, 0xc2, 0x92, 0xd0, 0xc4,
	0xd1, 0x7e, 0x74, 0x32, 0xe2, 0xbf, 0x0d, 0x76, 0x08, 0x93, 0x95, 0x37, 0x08, 0xc4, 0xc0, 0x11,
	0xdd, 0x66, 0x7a, 0x01, 0x13, 0xab, 0xeb, 0x04, 0xaf, 0xc5, 0xa3, 0x64, 0x53, 0xd8, 0x74, 0x45,
	0x10, 0xf4, 0x05, 0x8b, 0x61, 0x0b, 0xdf, 0x55, 0x45, 0xa8, 0x83, 0x4c, 0x53, 0xa6, 0x5f, 0x11,
	0xec, 0xfa, 0x64, 0x3e, 0xaf, 0xe7, 0xcf, 0x60, 0x48, 0xa8, 0xeb, 0x17, 0x9f, 0x6b, 0x9c, 0xcd,
	0x17, 0xbd, 0x8b, 0xe0, 0x8e, 0xe1, 0x81, 0x65, 0xe7, 0xb0, 0xf1, 0x50, 0x98, 0xc2, 0x59, 0x8c,
	0xb3, 0x83, 0xfe, 0x99, 0x4e, 0x5c, 0xee, 0x06, 0xb2, 0xcf, 0x08, 0xf6, 0xc2, 0xbf, 0x96, 0x2d,
	0xd6, 0xa2, 0xb6, 0x8d, 0xc4, 0x72, 0xd8, 0xee, 0xec, 0xee, 0xe8, 0x6f, 0xf1, 0x36, 0x97, 0x1c,
	0xff, 0xc7, 0xb5, 0x7e, 0x7c, 0x19, 0x7f, 0x0f, 0x66, 0xd6, 0xbb, 0xf1, 0xb4, 0x23, 0x37, 0xf6,
	0x74, 0xf7, 0x43, 0x77, 0xc1, 0xd3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xb7, 0xfc, 0xd6,
	0x0c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceAuthrizationGrpcServerClient is the client API for ServiceAuthrizationGrpcServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceAuthrizationGrpcServerClient interface {
	RequestToken(ctx context.Context, in *GrpcRequestToken, opts ...grpc.CallOption) (*GrpcResponseToken, error)
}

type serviceAuthrizationGrpcServerClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceAuthrizationGrpcServerClient(cc grpc.ClientConnInterface) ServiceAuthrizationGrpcServerClient {
	return &serviceAuthrizationGrpcServerClient{cc}
}

func (c *serviceAuthrizationGrpcServerClient) RequestToken(ctx context.Context, in *GrpcRequestToken, opts ...grpc.CallOption) (*GrpcResponseToken, error) {
	out := new(GrpcResponseToken)
	err := c.cc.Invoke(ctx, "/service_authrization.ServiceAuthrizationGrpcServer/RequestToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceAuthrizationGrpcServerServer is the server API for ServiceAuthrizationGrpcServer service.
type ServiceAuthrizationGrpcServerServer interface {
	RequestToken(context.Context, *GrpcRequestToken) (*GrpcResponseToken, error)
}

// UnimplementedServiceAuthrizationGrpcServerServer can be embedded to have forward compatible implementations.
type UnimplementedServiceAuthrizationGrpcServerServer struct {
}

func (*UnimplementedServiceAuthrizationGrpcServerServer) RequestToken(ctx context.Context, req *GrpcRequestToken) (*GrpcResponseToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestToken not implemented")
}

func RegisterServiceAuthrizationGrpcServerServer(s *grpc.Server, srv ServiceAuthrizationGrpcServerServer) {
	s.RegisterService(&_ServiceAuthrizationGrpcServer_serviceDesc, srv)
}

func _ServiceAuthrizationGrpcServer_RequestToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrpcRequestToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAuthrizationGrpcServerServer).RequestToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_authrization.ServiceAuthrizationGrpcServer/RequestToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAuthrizationGrpcServerServer).RequestToken(ctx, req.(*GrpcRequestToken))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceAuthrizationGrpcServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service_authrization.ServiceAuthrizationGrpcServer",
	HandlerType: (*ServiceAuthrizationGrpcServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestToken",
			Handler:    _ServiceAuthrizationGrpcServer_RequestToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_authrization.proto",
}
