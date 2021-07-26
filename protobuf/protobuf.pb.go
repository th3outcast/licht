// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuf.proto

package protobuf

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ReturnValue struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReturnValue) Reset()         { *m = ReturnValue{} }
func (m *ReturnValue) String() string { return proto.CompactTextString(m) }
func (*ReturnValue) ProtoMessage()    {}
func (*ReturnValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{0}
}
func (m *ReturnValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReturnValue.Unmarshal(m, b)
}
func (m *ReturnValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReturnValue.Marshal(b, m, deterministic)
}
func (m *ReturnValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReturnValue.Merge(m, src)
}
func (m *ReturnValue) XXX_Size() int {
	return xxx_messageInfo_ReturnValue.Size(m)
}
func (m *ReturnValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ReturnValue.DiscardUnknown(m)
}

var xxx_messageInfo_ReturnValue proto.InternalMessageInfo

func (m *ReturnValue) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *ReturnValue) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SearchKey struct {
	Node                 int32    `protobuf:"varint,1,opt,name=node,proto3" json:"node,omitempty"`
	Key                  []byte   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchKey) Reset()         { *m = SearchKey{} }
func (m *SearchKey) String() string { return proto.CompactTextString(m) }
func (*SearchKey) ProtoMessage()    {}
func (*SearchKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{1}
}
func (m *SearchKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchKey.Unmarshal(m, b)
}
func (m *SearchKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchKey.Marshal(b, m, deterministic)
}
func (m *SearchKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchKey.Merge(m, src)
}
func (m *SearchKey) XXX_Size() int {
	return xxx_messageInfo_SearchKey.Size(m)
}
func (m *SearchKey) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchKey.DiscardUnknown(m)
}

var xxx_messageInfo_SearchKey proto.InternalMessageInfo

func (m *SearchKey) GetNode() int32 {
	if m != nil {
		return m.Node
	}
	return 0
}

func (m *SearchKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func init() {
	proto.RegisterType((*ReturnValue)(nil), "protobuf.ReturnValue")
	proto.RegisterType((*SearchKey)(nil), "protobuf.SearchKey")
}

func init() { proto.RegisterFile("protobuf.proto", fileDescriptor_c77a803fcbc0c059) }

var fileDescriptor_c77a803fcbc0c059 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x03, 0x33, 0x84, 0x38, 0x60, 0x7c, 0x25, 0x53, 0x2e, 0xee, 0xa0,
	0xd4, 0x92, 0xd2, 0xa2, 0xbc, 0xb0, 0xc4, 0x9c, 0xd2, 0x54, 0x21, 0x21, 0x2e, 0x96, 0x8c, 0xc4,
	0xe2, 0x0c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x30, 0x1b, 0x24, 0x96, 0x92, 0x58, 0x92,
	0x28, 0xc1, 0x04, 0x11, 0x03, 0xb1, 0x95, 0x0c, 0xb9, 0x38, 0x83, 0x53, 0x13, 0x8b, 0x92, 0x33,
	0xbc, 0x53, 0x2b, 0x41, 0x0a, 0xf2, 0xf2, 0x53, 0x52, 0xc1, 0x9a, 0x58, 0x83, 0xc0, 0x6c, 0x21,
	0x01, 0x2e, 0xe6, 0xec, 0xd4, 0x4a, 0xa8, 0x1e, 0x10, 0xd3, 0xc8, 0x85, 0x8b, 0xd5, 0x27, 0x33,
	0x39, 0xa3, 0x44, 0xc8, 0x9a, 0x8b, 0x37, 0x38, 0xb5, 0xa8, 0x2c, 0xb5, 0x28, 0x28, 0xb5, 0xb0,
	0x34, 0xb5, 0xb8, 0x44, 0x48, 0x58, 0x0f, 0xee, 0x3c, 0xb8, 0xa1, 0x52, 0xa2, 0x08, 0x41, 0x24,
	0x07, 0x26, 0xb1, 0x81, 0x45, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x36, 0x39, 0xab, 0x01,
	0xd2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LichtClient is the client API for Licht service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LichtClient interface {
	ServerRequest(ctx context.Context, in *SearchKey, opts ...grpc.CallOption) (*ReturnValue, error)
}

type lichtClient struct {
	cc *grpc.ClientConn
}

func NewLichtClient(cc *grpc.ClientConn) LichtClient {
	return &lichtClient{cc}
}

func (c *lichtClient) ServerRequest(ctx context.Context, in *SearchKey, opts ...grpc.CallOption) (*ReturnValue, error) {
	out := new(ReturnValue)
	err := c.cc.Invoke(ctx, "/protobuf.Licht/ServerRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LichtServer is the server API for Licht service.
type LichtServer interface {
	ServerRequest(context.Context, *SearchKey) (*ReturnValue, error)
}

// UnimplementedLichtServer can be embedded to have forward compatible implementations.
type UnimplementedLichtServer struct {
}

func (*UnimplementedLichtServer) ServerRequest(ctx context.Context, req *SearchKey) (*ReturnValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerRequest not implemented")
}

func RegisterLichtServer(s *grpc.Server, srv LichtServer) {
	s.RegisterService(&_Licht_serviceDesc, srv)
}

func _Licht_ServerRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LichtServer).ServerRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Licht/ServerRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LichtServer).ServerRequest(ctx, req.(*SearchKey))
	}
	return interceptor(ctx, in, info, handler)
}

var _Licht_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Licht",
	HandlerType: (*LichtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServerRequest",
			Handler:    _Licht_ServerRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf.proto",
}