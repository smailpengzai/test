// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package protoDir is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	InsertKey
	GetoutValue
*/
package testpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type InsertKey struct {
	Key  string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *InsertKey) Reset()                    { *m = InsertKey{} }
func (m *InsertKey) String() string            { return proto.CompactTextString(m) }
func (*InsertKey) ProtoMessage()               {}
func (*InsertKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InsertKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *InsertKey) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetoutValue struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	Index int64  `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
}

func (m *GetoutValue) Reset()                    { *m = GetoutValue{} }
func (m *GetoutValue) String() string            { return proto.CompactTextString(m) }
func (*GetoutValue) ProtoMessage()               {}
func (*GetoutValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetoutValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *GetoutValue) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func init() {
	proto.RegisterType((*InsertKey)(nil), "protoDir.InsertKey")
	proto.RegisterType((*GetoutValue)(nil), "protoDir.GetoutValue")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	// 简单 RPC
	GetValue(ctx context.Context, in *InsertKey, opts ...grpc.CallOption) (*GetoutValue, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) GetValue(ctx context.Context, in *InsertKey, opts ...grpc.CallOption) (*GetoutValue, error) {
	out := new(GetoutValue)
	err := grpc.Invoke(ctx, "/protoDir.Greeter/GetValue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// 简单 RPC
	GetValue(context.Context, *InsertKey) (*GetoutValue, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoDir.Greeter/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetValue(ctx, req.(*InsertKey))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protoDir.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetValue",
			Handler:    _Greeter_GetValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x2e, 0x99, 0x45, 0x4a, 0x86, 0x5c,
	0x9c, 0x9e, 0x79, 0xc5, 0xa9, 0x45, 0x25, 0xde, 0xa9, 0x95, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9,
	0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62,
	0x6e, 0xaa, 0x04, 0x13, 0x58, 0x08, 0xcc, 0x56, 0xb2, 0xe4, 0xe2, 0x76, 0x4f, 0x2d, 0xc9, 0x2f,
	0x2d, 0x09, 0x4b, 0xcc, 0x29, 0x4d, 0x15, 0x12, 0xe1, 0x62, 0x2d, 0x03, 0x31, 0xa0, 0xda, 0x20,
	0x1c, 0x90, 0x68, 0x66, 0x5e, 0x4a, 0x6a, 0x05, 0x58, 0x27, 0x73, 0x10, 0x84, 0x63, 0xe4, 0xcc,
	0xc5, 0xee, 0x5e, 0x94, 0x9a, 0x5a, 0x92, 0x5a, 0x24, 0x64, 0xc1, 0xc5, 0xe1, 0x9e, 0x0a, 0x35,
	0x42, 0x58, 0x0f, 0xe6, 0x1e, 0x3d, 0xb8, 0x63, 0xa4, 0x44, 0x11, 0x82, 0x48, 0xd6, 0x29, 0x31,
	0x24, 0xb1, 0x81, 0xc5, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x11, 0x12, 0xd7, 0xfb, 0xd1,
	0x00, 0x00, 0x00,
}
