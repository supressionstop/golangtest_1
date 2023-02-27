// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/protobuf/processor.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProcessorServiceClient is the client API for ProcessorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcessorServiceClient interface {
	SubscribeOn(ctx context.Context, opts ...grpc.CallOption) (ProcessorService_SubscribeOnClient, error)
}

type processorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessorServiceClient(cc grpc.ClientConnInterface) ProcessorServiceClient {
	return &processorServiceClient{cc}
}

func (c *processorServiceClient) SubscribeOn(ctx context.Context, opts ...grpc.CallOption) (ProcessorService_SubscribeOnClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProcessorService_ServiceDesc.Streams[0], "/processor.ProcessorService/SubscribeOn", opts...)
	if err != nil {
		return nil, err
	}
	x := &processorServiceSubscribeOnClient{stream}
	return x, nil
}

type ProcessorService_SubscribeOnClient interface {
	Send(*Subscribe) error
	Recv() (*SportsData, error)
	grpc.ClientStream
}

type processorServiceSubscribeOnClient struct {
	grpc.ClientStream
}

func (x *processorServiceSubscribeOnClient) Send(m *Subscribe) error {
	return x.ClientStream.SendMsg(m)
}

func (x *processorServiceSubscribeOnClient) Recv() (*SportsData, error) {
	m := new(SportsData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProcessorServiceServer is the server API for ProcessorService service.
// All implementations must embed UnimplementedProcessorServiceServer
// for forward compatibility
type ProcessorServiceServer interface {
	SubscribeOn(ProcessorService_SubscribeOnServer) error
	mustEmbedUnimplementedProcessorServiceServer()
}

// UnimplementedProcessorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProcessorServiceServer struct {
}

func (UnimplementedProcessorServiceServer) SubscribeOn(ProcessorService_SubscribeOnServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeOn not implemented")
}
func (UnimplementedProcessorServiceServer) mustEmbedUnimplementedProcessorServiceServer() {}

// UnsafeProcessorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProcessorServiceServer will
// result in compilation errors.
type UnsafeProcessorServiceServer interface {
	mustEmbedUnimplementedProcessorServiceServer()
}

func RegisterProcessorServiceServer(s grpc.ServiceRegistrar, srv ProcessorServiceServer) {
	s.RegisterService(&ProcessorService_ServiceDesc, srv)
}

func _ProcessorService_SubscribeOn_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProcessorServiceServer).SubscribeOn(&processorServiceSubscribeOnServer{stream})
}

type ProcessorService_SubscribeOnServer interface {
	Send(*SportsData) error
	Recv() (*Subscribe, error)
	grpc.ServerStream
}

type processorServiceSubscribeOnServer struct {
	grpc.ServerStream
}

func (x *processorServiceSubscribeOnServer) Send(m *SportsData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *processorServiceSubscribeOnServer) Recv() (*Subscribe, error) {
	m := new(Subscribe)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProcessorService_ServiceDesc is the grpc.ServiceDesc for ProcessorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProcessorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "processor.ProcessorService",
	HandlerType: (*ProcessorServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeOn",
			Handler:       _ProcessorService_SubscribeOn_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/protobuf/processor.proto",
}