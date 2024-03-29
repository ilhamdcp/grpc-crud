// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package item

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

// ItemServiceClient is the client API for ItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ItemServiceClient interface {
	GetItemById(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*ItemResponse, error)
	GetItemsByName(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (ItemService_GetItemsByNameClient, error)
}

type itemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewItemServiceClient(cc grpc.ClientConnInterface) ItemServiceClient {
	return &itemServiceClient{cc}
}

func (c *itemServiceClient) GetItemById(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*ItemResponse, error) {
	out := new(ItemResponse)
	err := c.cc.Invoke(ctx, "/grpc.service.item.ItemService/getItemById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetItemsByName(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (ItemService_GetItemsByNameClient, error) {
	stream, err := c.cc.NewStream(ctx, &ItemService_ServiceDesc.Streams[0], "/grpc.service.item.ItemService/getItemsByName", opts...)
	if err != nil {
		return nil, err
	}
	x := &itemServiceGetItemsByNameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ItemService_GetItemsByNameClient interface {
	Recv() (*Item, error)
	grpc.ClientStream
}

type itemServiceGetItemsByNameClient struct {
	grpc.ClientStream
}

func (x *itemServiceGetItemsByNameClient) Recv() (*Item, error) {
	m := new(Item)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ItemServiceServer is the server API for ItemService service.
// All implementations must embed UnimplementedItemServiceServer
// for forward compatibility
type ItemServiceServer interface {
	GetItemById(context.Context, *ItemRequest) (*ItemResponse, error)
	GetItemsByName(*ItemRequest, ItemService_GetItemsByNameServer) error
	mustEmbedUnimplementedItemServiceServer()
}

// UnimplementedItemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedItemServiceServer struct {
}

func (UnimplementedItemServiceServer) GetItemById(context.Context, *ItemRequest) (*ItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemById not implemented")
}
func (UnimplementedItemServiceServer) GetItemsByName(*ItemRequest, ItemService_GetItemsByNameServer) error {
	return status.Errorf(codes.Unimplemented, "method GetItemsByName not implemented")
}
func (UnimplementedItemServiceServer) mustEmbedUnimplementedItemServiceServer() {}

// UnsafeItemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItemServiceServer will
// result in compilation errors.
type UnsafeItemServiceServer interface {
	mustEmbedUnimplementedItemServiceServer()
}

func RegisterItemServiceServer(s grpc.ServiceRegistrar, srv ItemServiceServer) {
	s.RegisterService(&ItemService_ServiceDesc, srv)
}

func _ItemService_GetItemById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetItemById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.service.item.ItemService/getItemById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetItemById(ctx, req.(*ItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetItemsByName_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ItemRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ItemServiceServer).GetItemsByName(m, &itemServiceGetItemsByNameServer{stream})
}

type ItemService_GetItemsByNameServer interface {
	Send(*Item) error
	grpc.ServerStream
}

type itemServiceGetItemsByNameServer struct {
	grpc.ServerStream
}

func (x *itemServiceGetItemsByNameServer) Send(m *Item) error {
	return x.ServerStream.SendMsg(m)
}

// ItemService_ServiceDesc is the grpc.ServiceDesc for ItemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ItemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.service.item.ItemService",
	HandlerType: (*ItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getItemById",
			Handler:    _ItemService_GetItemById_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getItemsByName",
			Handler:       _ItemService_GetItemsByName_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/item/Item.proto",
}
