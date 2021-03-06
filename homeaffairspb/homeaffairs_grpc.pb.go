// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package homeaffairspb

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

// HomeAffairsClient is the client API for HomeAffairs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HomeAffairsClient interface {
	GetCitizen(ctx context.Context, in *CitizenRequest, opts ...grpc.CallOption) (*CitizenResponse, error)
}

type homeAffairsClient struct {
	cc grpc.ClientConnInterface
}

func NewHomeAffairsClient(cc grpc.ClientConnInterface) HomeAffairsClient {
	return &homeAffairsClient{cc}
}

func (c *homeAffairsClient) GetCitizen(ctx context.Context, in *CitizenRequest, opts ...grpc.CallOption) (*CitizenResponse, error) {
	out := new(CitizenResponse)
	err := c.cc.Invoke(ctx, "/HomeAffairs/GetCitizen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HomeAffairsServer is the server API for HomeAffairs service.
// All implementations must embed UnimplementedHomeAffairsServer
// for forward compatibility
type HomeAffairsServer interface {
	GetCitizen(context.Context, *CitizenRequest) (*CitizenResponse, error)
	mustEmbedUnimplementedHomeAffairsServer()
}

// UnimplementedHomeAffairsServer must be embedded to have forward compatible implementations.
type UnimplementedHomeAffairsServer struct {
}

func (UnimplementedHomeAffairsServer) GetCitizen(context.Context, *CitizenRequest) (*CitizenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCitizen not implemented")
}
func (UnimplementedHomeAffairsServer) mustEmbedUnimplementedHomeAffairsServer() {}

// UnsafeHomeAffairsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HomeAffairsServer will
// result in compilation errors.
type UnsafeHomeAffairsServer interface {
	mustEmbedUnimplementedHomeAffairsServer()
}

func RegisterHomeAffairsServer(s grpc.ServiceRegistrar, srv HomeAffairsServer) {
	s.RegisterService(&HomeAffairs_ServiceDesc, srv)
}

func _HomeAffairs_GetCitizen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CitizenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeAffairsServer).GetCitizen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HomeAffairs/GetCitizen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeAffairsServer).GetCitizen(ctx, req.(*CitizenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HomeAffairs_ServiceDesc is the grpc.ServiceDesc for HomeAffairs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HomeAffairs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HomeAffairs",
	HandlerType: (*HomeAffairsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCitizen",
			Handler:    _HomeAffairs_GetCitizen_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "homeaffairspb/homeaffairs.proto",
}
