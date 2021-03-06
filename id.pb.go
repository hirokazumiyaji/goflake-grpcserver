// Code generated by protoc-gen-go.
// source: id.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	id.proto

It has these top-level messages:
	GetIDRequest
	GetIDResponse
*/
package main

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

type GetIDRequest struct {
}

func (m *GetIDRequest) Reset()                    { *m = GetIDRequest{} }
func (m *GetIDRequest) String() string            { return proto.CompactTextString(m) }
func (*GetIDRequest) ProtoMessage()               {}
func (*GetIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetIDResponse struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *GetIDResponse) Reset()                    { *m = GetIDResponse{} }
func (m *GetIDResponse) String() string            { return proto.CompactTextString(m) }
func (*GetIDResponse) ProtoMessage()               {}
func (*GetIDResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*GetIDRequest)(nil), "main.GetIDRequest")
	proto.RegisterType((*GetIDResponse)(nil), "main.GetIDResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Goflake service

type GoflakeClient interface {
	GetID(ctx context.Context, in *GetIDRequest, opts ...grpc.CallOption) (*GetIDResponse, error)
}

type goflakeClient struct {
	cc *grpc.ClientConn
}

func NewGoflakeClient(cc *grpc.ClientConn) GoflakeClient {
	return &goflakeClient{cc}
}

func (c *goflakeClient) GetID(ctx context.Context, in *GetIDRequest, opts ...grpc.CallOption) (*GetIDResponse, error) {
	out := new(GetIDResponse)
	err := grpc.Invoke(ctx, "/main.Goflake/GetID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Goflake service

type GoflakeServer interface {
	GetID(context.Context, *GetIDRequest) (*GetIDResponse, error)
}

func RegisterGoflakeServer(s *grpc.Server, srv GoflakeServer) {
	s.RegisterService(&_Goflake_serviceDesc, srv)
}

func _Goflake_GetID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoflakeServer).GetID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Goflake/GetID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoflakeServer).GetID(ctx, req.(*GetIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Goflake_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.Goflake",
	HandlerType: (*GoflakeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetID",
			Handler:    _Goflake_GetID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

var fileDescriptor0 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0x4c, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0xe2, 0xe3, 0xe2, 0x71, 0x4f,
	0x2d, 0xf1, 0x74, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0x92, 0xe6, 0xe2, 0x85, 0xf2,
	0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xb8, 0xb8, 0x98, 0x3c, 0x5d, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x8d, 0x6c, 0xb9, 0xd8, 0xdd, 0xf3, 0xd3, 0x72, 0x12, 0xb3, 0x53, 0x85, 0x8c, 0xb8,
	0x58, 0xc1, 0xea, 0x84, 0x84, 0xf4, 0x40, 0xe6, 0xe8, 0x21, 0x1b, 0x22, 0x25, 0x8c, 0x22, 0x06,
	0x31, 0x48, 0x89, 0xc1, 0x49, 0x8e, 0x4b, 0x30, 0x33, 0x5f, 0x2f, 0x3d, 0xb3, 0x24, 0xa3, 0x34,
	0x49, 0x2f, 0x1d, 0x62, 0x90, 0x13, 0xcc, 0xc4, 0x00, 0xc6, 0x24, 0x36, 0xb0, 0xc3, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x36, 0x72, 0x0c, 0xa1, 0xa4, 0x00, 0x00, 0x00,
}
