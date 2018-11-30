// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/iam/v1/iam_token_service.proto

package iam // import "github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateIamTokenRequest struct {
	// Types that are valid to be assigned to Identity:
	//	*CreateIamTokenRequest_YandexPassportOauthToken
	//	*CreateIamTokenRequest_Jwt
	Identity             isCreateIamTokenRequest_Identity `protobuf_oneof:"identity"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *CreateIamTokenRequest) Reset()         { *m = CreateIamTokenRequest{} }
func (m *CreateIamTokenRequest) String() string { return proto.CompactTextString(m) }
func (*CreateIamTokenRequest) ProtoMessage()    {}
func (*CreateIamTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_iam_token_service_c7becdff72a5104f, []int{0}
}
func (m *CreateIamTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateIamTokenRequest.Unmarshal(m, b)
}
func (m *CreateIamTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateIamTokenRequest.Marshal(b, m, deterministic)
}
func (dst *CreateIamTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateIamTokenRequest.Merge(dst, src)
}
func (m *CreateIamTokenRequest) XXX_Size() int {
	return xxx_messageInfo_CreateIamTokenRequest.Size(m)
}
func (m *CreateIamTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateIamTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateIamTokenRequest proto.InternalMessageInfo

type isCreateIamTokenRequest_Identity interface {
	isCreateIamTokenRequest_Identity()
}

type CreateIamTokenRequest_YandexPassportOauthToken struct {
	YandexPassportOauthToken string `protobuf:"bytes,1,opt,name=yandex_passport_oauth_token,json=yandexPassportOauthToken,proto3,oneof"`
}

type CreateIamTokenRequest_Jwt struct {
	Jwt string `protobuf:"bytes,2,opt,name=jwt,proto3,oneof"`
}

func (*CreateIamTokenRequest_YandexPassportOauthToken) isCreateIamTokenRequest_Identity() {}

func (*CreateIamTokenRequest_Jwt) isCreateIamTokenRequest_Identity() {}

func (m *CreateIamTokenRequest) GetIdentity() isCreateIamTokenRequest_Identity {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *CreateIamTokenRequest) GetYandexPassportOauthToken() string {
	if x, ok := m.GetIdentity().(*CreateIamTokenRequest_YandexPassportOauthToken); ok {
		return x.YandexPassportOauthToken
	}
	return ""
}

func (m *CreateIamTokenRequest) GetJwt() string {
	if x, ok := m.GetIdentity().(*CreateIamTokenRequest_Jwt); ok {
		return x.Jwt
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CreateIamTokenRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CreateIamTokenRequest_OneofMarshaler, _CreateIamTokenRequest_OneofUnmarshaler, _CreateIamTokenRequest_OneofSizer, []interface{}{
		(*CreateIamTokenRequest_YandexPassportOauthToken)(nil),
		(*CreateIamTokenRequest_Jwt)(nil),
	}
}

func _CreateIamTokenRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CreateIamTokenRequest)
	// identity
	switch x := m.Identity.(type) {
	case *CreateIamTokenRequest_YandexPassportOauthToken:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.YandexPassportOauthToken)
	case *CreateIamTokenRequest_Jwt:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Jwt)
	case nil:
	default:
		return fmt.Errorf("CreateIamTokenRequest.Identity has unexpected type %T", x)
	}
	return nil
}

func _CreateIamTokenRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CreateIamTokenRequest)
	switch tag {
	case 1: // identity.yandex_passport_oauth_token
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Identity = &CreateIamTokenRequest_YandexPassportOauthToken{x}
		return true, err
	case 2: // identity.jwt
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Identity = &CreateIamTokenRequest_Jwt{x}
		return true, err
	default:
		return false, nil
	}
}

func _CreateIamTokenRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CreateIamTokenRequest)
	// identity
	switch x := m.Identity.(type) {
	case *CreateIamTokenRequest_YandexPassportOauthToken:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.YandexPassportOauthToken)))
		n += len(x.YandexPassportOauthToken)
	case *CreateIamTokenRequest_Jwt:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Jwt)))
		n += len(x.Jwt)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type CreateIamTokenResponse struct {
	// IAM token for the specified identity.
	//
	// You should pass the token in the `Authorization` header for any further API requests.
	// For example, `Authorization: Bearer [iam_token]`.
	IamToken             string   `protobuf:"bytes,1,opt,name=iam_token,json=iamToken,proto3" json:"iam_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateIamTokenResponse) Reset()         { *m = CreateIamTokenResponse{} }
func (m *CreateIamTokenResponse) String() string { return proto.CompactTextString(m) }
func (*CreateIamTokenResponse) ProtoMessage()    {}
func (*CreateIamTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_iam_token_service_c7becdff72a5104f, []int{1}
}
func (m *CreateIamTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateIamTokenResponse.Unmarshal(m, b)
}
func (m *CreateIamTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateIamTokenResponse.Marshal(b, m, deterministic)
}
func (dst *CreateIamTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateIamTokenResponse.Merge(dst, src)
}
func (m *CreateIamTokenResponse) XXX_Size() int {
	return xxx_messageInfo_CreateIamTokenResponse.Size(m)
}
func (m *CreateIamTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateIamTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateIamTokenResponse proto.InternalMessageInfo

func (m *CreateIamTokenResponse) GetIamToken() string {
	if m != nil {
		return m.IamToken
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateIamTokenRequest)(nil), "yandex.cloud.iam.v1.CreateIamTokenRequest")
	proto.RegisterType((*CreateIamTokenResponse)(nil), "yandex.cloud.iam.v1.CreateIamTokenResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IamTokenServiceClient is the client API for IamTokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IamTokenServiceClient interface {
	// Creates an IAM token for the specified identity.
	Create(ctx context.Context, in *CreateIamTokenRequest, opts ...grpc.CallOption) (*CreateIamTokenResponse, error)
}

type iamTokenServiceClient struct {
	cc *grpc.ClientConn
}

func NewIamTokenServiceClient(cc *grpc.ClientConn) IamTokenServiceClient {
	return &iamTokenServiceClient{cc}
}

func (c *iamTokenServiceClient) Create(ctx context.Context, in *CreateIamTokenRequest, opts ...grpc.CallOption) (*CreateIamTokenResponse, error) {
	out := new(CreateIamTokenResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.iam.v1.IamTokenService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IamTokenServiceServer is the server API for IamTokenService service.
type IamTokenServiceServer interface {
	// Creates an IAM token for the specified identity.
	Create(context.Context, *CreateIamTokenRequest) (*CreateIamTokenResponse, error)
}

func RegisterIamTokenServiceServer(s *grpc.Server, srv IamTokenServiceServer) {
	s.RegisterService(&_IamTokenService_serviceDesc, srv)
}

func _IamTokenService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIamTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamTokenServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.iam.v1.IamTokenService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamTokenServiceServer).Create(ctx, req.(*CreateIamTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IamTokenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.iam.v1.IamTokenService",
	HandlerType: (*IamTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _IamTokenService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/iam/v1/iam_token_service.proto",
}

func init() {
	proto.RegisterFile("yandex/cloud/iam/v1/iam_token_service.proto", fileDescriptor_iam_token_service_c7becdff72a5104f)
}

var fileDescriptor_iam_token_service_c7becdff72a5104f = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x4d, 0x95, 0xd2, 0xce, 0x42, 0x65, 0x8a, 0x52, 0x5b, 0x05, 0xc9, 0x4a, 0x5a, 0x3a,
	0x43, 0x15, 0x37, 0x16, 0x11, 0xea, 0x46, 0x57, 0x4a, 0x75, 0xe5, 0x26, 0x4c, 0x9b, 0x21, 0x1d,
	0x6d, 0xe6, 0xc5, 0xcc, 0x24, 0x5a, 0x10, 0x17, 0x5e, 0xc0, 0x85, 0x17, 0xd2, 0x33, 0x78, 0x05,
	0x0f, 0x22, 0x99, 0xd7, 0x0a, 0x4a, 0x17, 0xae, 0x42, 0x78, 0xdf, 0xfb, 0xff, 0x7f, 0xfe, 0x47,
	0xda, 0x53, 0xa1, 0x43, 0xf9, 0xc8, 0x47, 0x13, 0xc8, 0x42, 0xae, 0x44, 0xcc, 0xf3, 0x6e, 0xf1,
	0x09, 0x2c, 0xdc, 0x49, 0x1d, 0x18, 0x99, 0xe6, 0x6a, 0x24, 0x59, 0x92, 0x82, 0x05, 0x5a, 0x43,
	0x98, 0x39, 0x98, 0x29, 0x11, 0xb3, 0xbc, 0xdb, 0xd8, 0x8e, 0x00, 0xa2, 0x89, 0xe4, 0x22, 0x51,
	0x5c, 0x68, 0x0d, 0x56, 0x58, 0x05, 0xda, 0xe0, 0x4a, 0x63, 0xe7, 0x97, 0x7e, 0x2e, 0x26, 0x2a,
	0x74, 0x73, 0x1c, 0xfb, 0xcf, 0x64, 0xe3, 0x34, 0x95, 0xc2, 0xca, 0x73, 0x11, 0x5f, 0x17, 0x8e,
	0x03, 0x79, 0x9f, 0x49, 0x63, 0xe9, 0x09, 0x69, 0xe2, 0x66, 0x90, 0x08, 0x63, 0x12, 0x48, 0x6d,
	0x00, 0x22, 0xb3, 0x63, 0xcc, 0x55, 0xf7, 0x76, 0xbd, 0xbd, 0xea, 0xd9, 0xd2, 0xa0, 0x8e, 0xd0,
	0xe5, 0x8c, 0xb9, 0x28, 0x10, 0xa7, 0x43, 0x29, 0x59, 0xbe, 0x7d, 0xb0, 0xf5, 0xd2, 0x0c, 0x2c,
	0x7e, 0xfa, 0xeb, 0xa4, 0xa2, 0x42, 0xa9, 0xad, 0xb2, 0x53, 0xba, 0xf2, 0xfe, 0xd1, 0xf5, 0xfc,
	0x43, 0xb2, 0xf9, 0xd7, 0xdf, 0x24, 0xa0, 0x8d, 0xa4, 0x4d, 0x52, 0xfd, 0xa9, 0x01, 0xed, 0x06,
	0x15, 0x35, 0x83, 0xf6, 0x5f, 0x3d, 0xb2, 0x36, 0xdf, 0xb8, 0xc2, 0x8a, 0xe8, 0x13, 0x29, 0xa3,
	0x14, 0x6d, 0xb1, 0x05, 0x3d, 0xb1, 0x85, 0xef, 0x6c, 0xb4, 0xff, 0xc5, 0x62, 0x26, 0x7f, 0xeb,
	0xe5, 0xf3, 0xeb, 0xad, 0x54, 0xf3, 0x57, 0xe7, 0x87, 0x72, 0xe9, 0xcc, 0x91, 0xd7, 0xea, 0x1f,
	0xdf, 0xf4, 0x22, 0x65, 0xc7, 0xd9, 0x90, 0x8d, 0x20, 0xe6, 0xa8, 0xd9, 0xc1, 0xd2, 0x23, 0xe8,
	0x44, 0x52, 0xbb, 0xbe, 0xf9, 0x82, 0x6b, 0xf7, 0x94, 0x88, 0x87, 0x65, 0x37, 0x3e, 0xf8, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0x46, 0x8f, 0x00, 0xe5, 0x0f, 0x02, 0x00, 0x00,
}
