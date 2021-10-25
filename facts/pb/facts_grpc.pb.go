// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FactsClient is the client API for Facts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FactsClient interface {
	CreateFact(ctx context.Context, in *CreateFactRequest, opts ...grpc.CallOption) (*CreateFactReply, error)
	GetFact(ctx context.Context, in *GetFactRequest, opts ...grpc.CallOption) (*GetFactReply, error)
	DeleteFact(ctx context.Context, in *DeleteFactRequest, opts ...grpc.CallOption) (*DeleteFactReply, error)
	GetAnimals(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAnimalsReply, error)
	GetRandAnimalFact(ctx context.Context, in *GetRandAnimalFactRequest, opts ...grpc.CallOption) (*GetRandAnimalFactReply, error)
	PublishFact(ctx context.Context, in *PublishFactRequest, opts ...grpc.CallOption) (*PublishFactReply, error)
}

type factsClient struct {
	cc grpc.ClientConnInterface
}

func NewFactsClient(cc grpc.ClientConnInterface) FactsClient {
	return &factsClient{cc}
}

func (c *factsClient) CreateFact(ctx context.Context, in *CreateFactRequest, opts ...grpc.CallOption) (*CreateFactReply, error) {
	out := new(CreateFactReply)
	err := c.cc.Invoke(ctx, "/Facts/CreateFact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *factsClient) GetFact(ctx context.Context, in *GetFactRequest, opts ...grpc.CallOption) (*GetFactReply, error) {
	out := new(GetFactReply)
	err := c.cc.Invoke(ctx, "/Facts/GetFact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *factsClient) DeleteFact(ctx context.Context, in *DeleteFactRequest, opts ...grpc.CallOption) (*DeleteFactReply, error) {
	out := new(DeleteFactReply)
	err := c.cc.Invoke(ctx, "/Facts/DeleteFact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *factsClient) GetAnimals(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAnimalsReply, error) {
	out := new(GetAnimalsReply)
	err := c.cc.Invoke(ctx, "/Facts/GetAnimals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *factsClient) GetRandAnimalFact(ctx context.Context, in *GetRandAnimalFactRequest, opts ...grpc.CallOption) (*GetRandAnimalFactReply, error) {
	out := new(GetRandAnimalFactReply)
	err := c.cc.Invoke(ctx, "/Facts/GetRandAnimalFact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *factsClient) PublishFact(ctx context.Context, in *PublishFactRequest, opts ...grpc.CallOption) (*PublishFactReply, error) {
	out := new(PublishFactReply)
	err := c.cc.Invoke(ctx, "/Facts/PublishFact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FactsServer is the server API for Facts service.
// All implementations must embed UnimplementedFactsServer
// for forward compatibility
type FactsServer interface {
	CreateFact(context.Context, *CreateFactRequest) (*CreateFactReply, error)
	GetFact(context.Context, *GetFactRequest) (*GetFactReply, error)
	DeleteFact(context.Context, *DeleteFactRequest) (*DeleteFactReply, error)
	GetAnimals(context.Context, *emptypb.Empty) (*GetAnimalsReply, error)
	GetRandAnimalFact(context.Context, *GetRandAnimalFactRequest) (*GetRandAnimalFactReply, error)
	PublishFact(context.Context, *PublishFactRequest) (*PublishFactReply, error)
	mustEmbedUnimplementedFactsServer()
}

// UnimplementedFactsServer must be embedded to have forward compatible implementations.
type UnimplementedFactsServer struct {
}

func (UnimplementedFactsServer) CreateFact(context.Context, *CreateFactRequest) (*CreateFactReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFact not implemented")
}
func (UnimplementedFactsServer) GetFact(context.Context, *GetFactRequest) (*GetFactReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFact not implemented")
}
func (UnimplementedFactsServer) DeleteFact(context.Context, *DeleteFactRequest) (*DeleteFactReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFact not implemented")
}
func (UnimplementedFactsServer) GetAnimals(context.Context, *emptypb.Empty) (*GetAnimalsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnimals not implemented")
}
func (UnimplementedFactsServer) GetRandAnimalFact(context.Context, *GetRandAnimalFactRequest) (*GetRandAnimalFactReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandAnimalFact not implemented")
}
func (UnimplementedFactsServer) PublishFact(context.Context, *PublishFactRequest) (*PublishFactReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishFact not implemented")
}
func (UnimplementedFactsServer) mustEmbedUnimplementedFactsServer() {}

// UnsafeFactsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FactsServer will
// result in compilation errors.
type UnsafeFactsServer interface {
	mustEmbedUnimplementedFactsServer()
}

func RegisterFactsServer(s grpc.ServiceRegistrar, srv FactsServer) {
	s.RegisterService(&Facts_ServiceDesc, srv)
}

func _Facts_CreateFact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FactsServer).CreateFact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Facts/CreateFact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FactsServer).CreateFact(ctx, req.(*CreateFactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Facts_GetFact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FactsServer).GetFact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Facts/GetFact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FactsServer).GetFact(ctx, req.(*GetFactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Facts_DeleteFact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FactsServer).DeleteFact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Facts/DeleteFact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FactsServer).DeleteFact(ctx, req.(*DeleteFactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Facts_GetAnimals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FactsServer).GetAnimals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Facts/GetAnimals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FactsServer).GetAnimals(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Facts_GetRandAnimalFact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRandAnimalFactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FactsServer).GetRandAnimalFact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Facts/GetRandAnimalFact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FactsServer).GetRandAnimalFact(ctx, req.(*GetRandAnimalFactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Facts_PublishFact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishFactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FactsServer).PublishFact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Facts/PublishFact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FactsServer).PublishFact(ctx, req.(*PublishFactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Facts_ServiceDesc is the grpc.ServiceDesc for Facts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Facts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Facts",
	HandlerType: (*FactsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFact",
			Handler:    _Facts_CreateFact_Handler,
		},
		{
			MethodName: "GetFact",
			Handler:    _Facts_GetFact_Handler,
		},
		{
			MethodName: "DeleteFact",
			Handler:    _Facts_DeleteFact_Handler,
		},
		{
			MethodName: "GetAnimals",
			Handler:    _Facts_GetAnimals_Handler,
		},
		{
			MethodName: "GetRandAnimalFact",
			Handler:    _Facts_GetRandAnimalFact_Handler,
		},
		{
			MethodName: "PublishFact",
			Handler:    _Facts_PublishFact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "facts.proto",
}
