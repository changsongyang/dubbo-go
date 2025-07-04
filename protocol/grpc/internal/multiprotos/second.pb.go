/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: second.proto

package multiprotos

import (
	context "context"
	fmt "fmt"
	math "math"
)

import (
	proto "github.com/golang/protobuf/proto"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

import (
	"dubbo.apache.org/dubbo-go/v3/protocol/base"
	"dubbo.apache.org/dubbo-go/v3/protocol/invocation"
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

type SecondRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecondRequest) Reset()         { *m = SecondRequest{} }
func (m *SecondRequest) String() string { return proto.CompactTextString(m) }
func (*SecondRequest) ProtoMessage()    {}
func (*SecondRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83c4a0a7b0d4882f, []int{0}
}

func (m *SecondRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecondRequest.Unmarshal(m, b)
}
func (m *SecondRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecondRequest.Marshal(b, m, deterministic)
}
func (m *SecondRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecondRequest.Merge(m, src)
}
func (m *SecondRequest) XXX_Size() int {
	return xxx_messageInfo_SecondRequest.Size(m)
}
func (m *SecondRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SecondRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SecondRequest proto.InternalMessageInfo

func (m *SecondRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SecondResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecondResponse) Reset()         { *m = SecondResponse{} }
func (m *SecondResponse) String() string { return proto.CompactTextString(m) }
func (*SecondResponse) ProtoMessage()    {}
func (*SecondResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83c4a0a7b0d4882f, []int{1}
}

func (m *SecondResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecondResponse.Unmarshal(m, b)
}
func (m *SecondResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecondResponse.Marshal(b, m, deterministic)
}
func (m *SecondResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecondResponse.Merge(m, src)
}
func (m *SecondResponse) XXX_Size() int {
	return xxx_messageInfo_SecondResponse.Size(m)
}
func (m *SecondResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SecondResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SecondResponse proto.InternalMessageInfo

func (m *SecondResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SecondRequest)(nil), "multiprotos.SecondRequest")
	proto.RegisterType((*SecondResponse)(nil), "multiprotos.SecondResponse")
}

func init() { proto.RegisterFile("second.proto", fileDescriptor_83c4a0a7b0d4882f) }

var fileDescriptor_83c4a0a7b0d4882f = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x4d, 0xce,
	0xcf, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce, 0x2d, 0xcd, 0x29, 0xc9, 0x04,
	0xb3, 0x8b, 0x95, 0x34, 0xb9, 0x78, 0x83, 0xc1, 0x92, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25,
	0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x30, 0xae, 0x92, 0x16, 0x17, 0x1f, 0x4c, 0x69, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a,
	0x6e, 0xb5, 0x46, 0x9b, 0x99, 0x60, 0xe6, 0x06, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0xb9,
	0x72, 0x71, 0x40, 0x99, 0x86, 0x42, 0x52, 0x7a, 0x48, 0x4e, 0xd0, 0x43, 0xb1, 0x5f, 0x4a, 0x1a,
	0xab, 0x1c, 0xc4, 0x42, 0x25, 0x06, 0x21, 0x77, 0xb8, 0x31, 0x46, 0x14, 0x18, 0x63, 0xc0, 0x88,
	0x64, 0x90, 0x31, 0x05, 0x06, 0x69, 0x30, 0x0a, 0x79, 0xc2, 0x0d, 0x32, 0xa1, 0xc8, 0x20, 0x03,
	0xc6, 0x24, 0x36, 0xb0, 0xa4, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xba, 0xd2, 0x28, 0xb0,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SecondServiceClient is the client API for SecondService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecondServiceClient interface {
	Service1(ctx context.Context, in *SecondRequest, opts ...grpc.CallOption) (*SecondResponse, error)
	Service2(ctx context.Context, in *SecondRequest, opts ...grpc.CallOption) (SecondService_Service2Client, error)
	Service3(ctx context.Context, opts ...grpc.CallOption) (SecondService_Service3Client, error)
	Service4(ctx context.Context, opts ...grpc.CallOption) (SecondService_Service4Client, error)
}

type secondServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecondServiceClient(cc grpc.ClientConnInterface) SecondServiceClient {
	return &secondServiceClient{cc}
}

func (c *secondServiceClient) Service1(ctx context.Context, in *SecondRequest, opts ...grpc.CallOption) (*SecondResponse, error) {
	out := new(SecondResponse)
	err := c.cc.Invoke(ctx, "/multiprotos.SecondService/Service1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secondServiceClient) Service2(ctx context.Context, in *SecondRequest, opts ...grpc.CallOption) (SecondService_Service2Client, error) {
	stream, err := c.cc.NewStream(ctx, &_SecondService_serviceDesc.Streams[0], "/multiprotos.SecondService/Service2", opts...)
	if err != nil {
		return nil, err
	}
	x := &secondServiceService2Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SecondService_Service2Client interface {
	Recv() (*SecondResponse, error)
	grpc.ClientStream
}

type secondServiceService2Client struct {
	grpc.ClientStream
}

func (x *secondServiceService2Client) Recv() (*SecondResponse, error) {
	m := new(SecondResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secondServiceClient) Service3(ctx context.Context, opts ...grpc.CallOption) (SecondService_Service3Client, error) {
	stream, err := c.cc.NewStream(ctx, &_SecondService_serviceDesc.Streams[1], "/multiprotos.SecondService/Service3", opts...)
	if err != nil {
		return nil, err
	}
	x := &secondServiceService3Client{stream}
	return x, nil
}

type SecondService_Service3Client interface {
	Send(*SecondRequest) error
	CloseAndRecv() (*SecondResponse, error)
	grpc.ClientStream
}

type secondServiceService3Client struct {
	grpc.ClientStream
}

func (x *secondServiceService3Client) Send(m *SecondRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *secondServiceService3Client) CloseAndRecv() (*SecondResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SecondResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secondServiceClient) Service4(ctx context.Context, opts ...grpc.CallOption) (SecondService_Service4Client, error) {
	stream, err := c.cc.NewStream(ctx, &_SecondService_serviceDesc.Streams[2], "/multiprotos.SecondService/Service4", opts...)
	if err != nil {
		return nil, err
	}
	x := &secondServiceService4Client{stream}
	return x, nil
}

type SecondService_Service4Client interface {
	Send(*SecondRequest) error
	Recv() (*SecondResponse, error)
	grpc.ClientStream
}

type secondServiceService4Client struct {
	grpc.ClientStream
}

func (x *secondServiceService4Client) Send(m *SecondRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *secondServiceService4Client) Recv() (*SecondResponse, error) {
	m := new(SecondResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SecondServiceServer is the server API for SecondService service.
type SecondServiceServer interface {
	Service1(context.Context, *SecondRequest) (*SecondResponse, error)
	Service2(*SecondRequest, SecondService_Service2Server) error
	Service3(SecondService_Service3Server) error
	Service4(SecondService_Service4Server) error
}

// UnimplementedSecondServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSecondServiceServer struct {
}

func (*UnimplementedSecondServiceServer) Service1(ctx context.Context, req *SecondRequest) (*SecondResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Service1 not implemented")
}
func (*UnimplementedSecondServiceServer) Service2(req *SecondRequest, srv SecondService_Service2Server) error {
	return status.Errorf(codes.Unimplemented, "method Service2 not implemented")
}
func (*UnimplementedSecondServiceServer) Service3(srv SecondService_Service3Server) error {
	return status.Errorf(codes.Unimplemented, "method Service3 not implemented")
}
func (*UnimplementedSecondServiceServer) Service4(srv SecondService_Service4Server) error {
	return status.Errorf(codes.Unimplemented, "method Service4 not implemented")
}

func RegisterSecondServiceServer(s *grpc.Server, srv SecondServiceServer) {
	s.RegisterService(&_SecondService_serviceDesc, srv)
}

func _SecondService_Service1_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(SecondRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecondServiceServer).Service1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/multiprotos.SecondService/Service1",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(SecondServiceServer).Service1(ctx, req.(*SecondRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecondService_Service2_Handler(srv any, stream grpc.ServerStream) error {
	m := new(SecondRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SecondServiceServer).Service2(m, &secondServiceService2Server{stream})
}

type SecondService_Service2Server interface {
	Send(*SecondResponse) error
	grpc.ServerStream
}

type secondServiceService2Server struct {
	grpc.ServerStream
}

func (x *secondServiceService2Server) Send(m *SecondResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SecondService_Service3_Handler(srv any, stream grpc.ServerStream) error {
	return srv.(SecondServiceServer).Service3(&secondServiceService3Server{stream})
}

type SecondService_Service3Server interface {
	SendAndClose(*SecondResponse) error
	Recv() (*SecondRequest, error)
	grpc.ServerStream
}

type secondServiceService3Server struct {
	grpc.ServerStream
}

func (x *secondServiceService3Server) SendAndClose(m *SecondResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *secondServiceService3Server) Recv() (*SecondRequest, error) {
	m := new(SecondRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SecondService_Service4_Handler(srv any, stream grpc.ServerStream) error {
	return srv.(SecondServiceServer).Service4(&secondServiceService4Server{stream})
}

type SecondService_Service4Server interface {
	Send(*SecondResponse) error
	Recv() (*SecondRequest, error)
	grpc.ServerStream
}

type secondServiceService4Server struct {
	grpc.ServerStream
}

func (x *secondServiceService4Server) Send(m *SecondResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *secondServiceService4Server) Recv() (*SecondRequest, error) {
	m := new(SecondRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SecondService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "multiprotos.SecondService",
	HandlerType: (*SecondServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Service1",
			Handler:    _SecondService_Service1_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Service2",
			Handler:       _SecondService_Service2_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Service3",
			Handler:       _SecondService_Service3_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Service4",
			Handler:       _SecondService_Service4_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "second.proto",
}

// SecondServiceClientImpl is the client API for SecondService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecondServiceClientImpl struct {
	Service1 func(ctx context.Context, in *SecondRequest, out *SecondResponse) error
	Service2 func(ctx context.Context, in *SecondRequest) (SecondService_Service2Client, error)
	Service3 func(ctx context.Context) (SecondService_Service3Client, error)
	Service4 func(ctx context.Context) (SecondService_Service4Client, error)
}

func (c *SecondServiceClientImpl) Reference() string {
	return "secondServiceImpl"
}

func (c *SecondServiceClientImpl) GetDubboStub(cc *grpc.ClientConn) SecondServiceClient {
	return NewSecondServiceClient(cc)
}

type SecondServiceProviderBase struct {
	proxyImpl base.Invoker
}

func (s *SecondServiceProviderBase) SetProxyImpl(impl base.Invoker) {
	s.proxyImpl = impl
}

func (s *SecondServiceProviderBase) GetProxyImpl() base.Invoker {
	return s.proxyImpl
}

func (c *SecondServiceProviderBase) Reference() string {
	return "secondServiceImpl"
}

func _DUBBO_SecondService_Service1_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(SecondRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	// DubboGrpcService is gRPC service
	type DubboGrpcService interface {
		// SetProxyImpl sets proxy.
		SetProxyImpl(impl base.Invoker)
		// GetProxyImpl gets proxy.
		GetProxyImpl() base.Invoker
		// ServiceDesc gets an RPC service's specification.
		ServiceDesc() *grpc.ServiceDesc
	}
	base := srv.(DubboGrpcService)
	args := []any{}
	args = append(args, in)
	invo := invocation.NewRPCInvocation("Service1", args, nil)
	if interceptor == nil {
		result := base.GetProxyImpl().Invoke(ctx, invo)
		return result.Result(), result.Error()
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/multiprotos.SecondService/Service1",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		result := base.GetProxyImpl().Invoke(ctx, invo)
		return result.Result(), result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _DUBBO_SecondService_Service2_Handler(srv any, stream grpc.ServerStream) error {
	// DubboGrpcService is gRPC service
	type DubboGrpcService interface {
		// SetProxyImpl sets proxy.
		SetProxyImpl(impl base.Invoker)
		// GetProxyImpl gets proxy.
		GetProxyImpl() base.Invoker
		// ServiceDesc gets an RPC service's specification.
		ServiceDesc() *grpc.ServiceDesc
	}
	_, ok := srv.(DubboGrpcService)
	invo := invocation.NewRPCInvocation("Service2", nil, nil)
	if !ok {
		fmt.Println(invo)
	}
	m := new(SecondRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SecondServiceServer).Service2(m, &secondServiceService2Server{stream})
}

func _DUBBO_SecondService_Service3_Handler(srv any, stream grpc.ServerStream) error {
	// DubboGrpcService is gRPC service
	type DubboGrpcService interface {
		// SetProxyImpl sets proxy.
		SetProxyImpl(impl base.Invoker)
		// GetProxyImpl gets proxy.
		GetProxyImpl() base.Invoker
		// ServiceDesc gets an RPC service's specification.
		ServiceDesc() *grpc.ServiceDesc
	}
	_, ok := srv.(DubboGrpcService)
	invo := invocation.NewRPCInvocation("Service3", nil, nil)
	if !ok {
		fmt.Println(invo)
	}
	return srv.(SecondServiceServer).Service3(&secondServiceService3Server{stream})
}

func _DUBBO_SecondService_Service4_Handler(srv any, stream grpc.ServerStream) error {
	// DubboGrpcService is gRPC service
	type DubboGrpcService interface {
		// SetProxyImpl sets proxy.
		SetProxyImpl(impl base.Invoker)
		// GetProxyImpl gets proxy.
		GetProxyImpl() base.Invoker
		// ServiceDesc gets an RPC service's specification.
		ServiceDesc() *grpc.ServiceDesc
	}
	_, ok := srv.(DubboGrpcService)
	invo := invocation.NewRPCInvocation("Service4", nil, nil)
	if !ok {
		fmt.Println(invo)
	}
	return srv.(SecondServiceServer).Service4(&secondServiceService4Server{stream})
}

func (s *SecondServiceProviderBase) ServiceDesc() *grpc.ServiceDesc {
	return &grpc.ServiceDesc{
		ServiceName: "multiprotos.SecondService",
		HandlerType: (*SecondServiceServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Service1",
				Handler:    _DUBBO_SecondService_Service1_Handler,
			},
		},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "Service2",
				Handler:       _DUBBO_SecondService_Service2_Handler,
				ServerStreams: true,
			},
			{
				StreamName:    "Service3",
				Handler:       _DUBBO_SecondService_Service3_Handler,
				ClientStreams: true,
			},
			{
				StreamName:    "Service4",
				Handler:       _DUBBO_SecondService_Service4_Handler,
				ServerStreams: true,
				ClientStreams: true,
			},
		},
		Metadata: "second.proto",
	}
}
