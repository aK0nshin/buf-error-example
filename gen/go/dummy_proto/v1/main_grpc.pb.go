// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: dummy_proto/v1/main.proto

package dummy_protov1

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

const (
	DummyService_DummyAction_FullMethodName = "/dummy_proto.v1.DummyService/DummyAction"
)

// DummyServiceClient is the client API for DummyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DummyServiceClient interface {
	DummyAction(ctx context.Context, in *DummyActionRequest, opts ...grpc.CallOption) (*DummyActionResponse, error)
}

type dummyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDummyServiceClient(cc grpc.ClientConnInterface) DummyServiceClient {
	return &dummyServiceClient{cc}
}

func (c *dummyServiceClient) DummyAction(ctx context.Context, in *DummyActionRequest, opts ...grpc.CallOption) (*DummyActionResponse, error) {
	out := new(DummyActionResponse)
	err := c.cc.Invoke(ctx, DummyService_DummyAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DummyServiceServer is the server API for DummyService service.
// All implementations must embed UnimplementedDummyServiceServer
// for forward compatibility
type DummyServiceServer interface {
	DummyAction(context.Context, *DummyActionRequest) (*DummyActionResponse, error)
	mustEmbedUnimplementedDummyServiceServer()
}

// UnimplementedDummyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDummyServiceServer struct {
}

func (UnimplementedDummyServiceServer) DummyAction(context.Context, *DummyActionRequest) (*DummyActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DummyAction not implemented")
}
func (UnimplementedDummyServiceServer) mustEmbedUnimplementedDummyServiceServer() {}

// UnsafeDummyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DummyServiceServer will
// result in compilation errors.
type UnsafeDummyServiceServer interface {
	mustEmbedUnimplementedDummyServiceServer()
}

func RegisterDummyServiceServer(s grpc.ServiceRegistrar, srv DummyServiceServer) {
	s.RegisterService(&DummyService_ServiceDesc, srv)
}

func _DummyService_DummyAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DummyActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DummyServiceServer).DummyAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DummyService_DummyAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DummyServiceServer).DummyAction(ctx, req.(*DummyActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DummyService_ServiceDesc is the grpc.ServiceDesc for DummyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DummyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dummy_proto.v1.DummyService",
	HandlerType: (*DummyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DummyAction",
			Handler:    _DummyService_DummyAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dummy_proto/v1/main.proto",
}
