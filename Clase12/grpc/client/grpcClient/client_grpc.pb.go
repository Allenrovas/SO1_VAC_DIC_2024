// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: client.proto

package confproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GetInfo_ReturnInfo_FullMethodName = "/confproto.getInfo/returnInfo"
)

// GetInfoClient is the client API for GetInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetInfoClient interface {
	ReturnInfo(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*ReplyInfo, error)
}

type getInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewGetInfoClient(cc grpc.ClientConnInterface) GetInfoClient {
	return &getInfoClient{cc}
}

func (c *getInfoClient) ReturnInfo(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*ReplyInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReplyInfo)
	err := c.cc.Invoke(ctx, GetInfo_ReturnInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetInfoServer is the server API for GetInfo service.
// All implementations must embed UnimplementedGetInfoServer
// for forward compatibility.
type GetInfoServer interface {
	ReturnInfo(context.Context, *RequestId) (*ReplyInfo, error)
	mustEmbedUnimplementedGetInfoServer()
}

// UnimplementedGetInfoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGetInfoServer struct{}

func (UnimplementedGetInfoServer) ReturnInfo(context.Context, *RequestId) (*ReplyInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnInfo not implemented")
}
func (UnimplementedGetInfoServer) mustEmbedUnimplementedGetInfoServer() {}
func (UnimplementedGetInfoServer) testEmbeddedByValue()                 {}

// UnsafeGetInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetInfoServer will
// result in compilation errors.
type UnsafeGetInfoServer interface {
	mustEmbedUnimplementedGetInfoServer()
}

func RegisterGetInfoServer(s grpc.ServiceRegistrar, srv GetInfoServer) {
	// If the following call pancis, it indicates UnimplementedGetInfoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GetInfo_ServiceDesc, srv)
}

func _GetInfo_ReturnInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetInfoServer).ReturnInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetInfo_ReturnInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetInfoServer).ReturnInfo(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

// GetInfo_ServiceDesc is the grpc.ServiceDesc for GetInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "confproto.getInfo",
	HandlerType: (*GetInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "returnInfo",
			Handler:    _GetInfo_ReturnInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client.proto",
}
