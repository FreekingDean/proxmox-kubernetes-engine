// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proxmox_kubernetes_engine/v1/clusters.proto

package v1

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

const (
	ClusterService_GetCluster_FullMethodName    = "/proxmox_kubernetes_engine.v1.ClusterService/GetCluster"
	ClusterService_ListClusters_FullMethodName  = "/proxmox_kubernetes_engine.v1.ClusterService/ListClusters"
	ClusterService_CreateCluster_FullMethodName = "/proxmox_kubernetes_engine.v1.ClusterService/CreateCluster"
	ClusterService_UpdateCluster_FullMethodName = "/proxmox_kubernetes_engine.v1.ClusterService/UpdateCluster"
	ClusterService_DeleteCluster_FullMethodName = "/proxmox_kubernetes_engine.v1.ClusterService/DeleteCluster"
)

// ClusterServiceClient is the client API for ClusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterServiceClient interface {
	GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
	ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error)
	CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
	UpdateCluster(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
	DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clusterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterServiceClient(cc grpc.ClientConnInterface) ClusterServiceClient {
	return &clusterServiceClient{cc}
}

func (c *clusterServiceClient) GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, ClusterService_GetCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error) {
	out := new(ListClustersResponse)
	err := c.cc.Invoke(ctx, ClusterService_ListClusters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, ClusterService_CreateCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) UpdateCluster(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, ClusterService_UpdateCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ClusterService_DeleteCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServiceServer is the server API for ClusterService service.
// All implementations must embed UnimplementedClusterServiceServer
// for forward compatibility
type ClusterServiceServer interface {
	GetCluster(context.Context, *GetClusterRequest) (*Cluster, error)
	ListClusters(context.Context, *ListClustersRequest) (*ListClustersResponse, error)
	CreateCluster(context.Context, *CreateClusterRequest) (*Cluster, error)
	UpdateCluster(context.Context, *UpdateClusterRequest) (*Cluster, error)
	DeleteCluster(context.Context, *DeleteClusterRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedClusterServiceServer()
}

// UnimplementedClusterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClusterServiceServer struct {
}

func (UnimplementedClusterServiceServer) GetCluster(context.Context, *GetClusterRequest) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCluster not implemented")
}
func (UnimplementedClusterServiceServer) ListClusters(context.Context, *ListClustersRequest) (*ListClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClusters not implemented")
}
func (UnimplementedClusterServiceServer) CreateCluster(context.Context, *CreateClusterRequest) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCluster not implemented")
}
func (UnimplementedClusterServiceServer) UpdateCluster(context.Context, *UpdateClusterRequest) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCluster not implemented")
}
func (UnimplementedClusterServiceServer) DeleteCluster(context.Context, *DeleteClusterRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCluster not implemented")
}
func (UnimplementedClusterServiceServer) mustEmbedUnimplementedClusterServiceServer() {}

// UnsafeClusterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterServiceServer will
// result in compilation errors.
type UnsafeClusterServiceServer interface {
	mustEmbedUnimplementedClusterServiceServer()
}

func RegisterClusterServiceServer(s grpc.ServiceRegistrar, srv ClusterServiceServer) {
	s.RegisterService(&ClusterService_ServiceDesc, srv)
}

func _ClusterService_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_GetCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).GetCluster(ctx, req.(*GetClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_ListClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).ListClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_ListClusters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).ListClusters(ctx, req.(*ListClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).CreateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_CreateCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).CreateCluster(ctx, req.(*CreateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_UpdateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).UpdateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_UpdateCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).UpdateCluster(ctx, req.(*UpdateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_DeleteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).DeleteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_DeleteCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).DeleteCluster(ctx, req.(*DeleteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterService_ServiceDesc is the grpc.ServiceDesc for ClusterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proxmox_kubernetes_engine.v1.ClusterService",
	HandlerType: (*ClusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCluster",
			Handler:    _ClusterService_GetCluster_Handler,
		},
		{
			MethodName: "ListClusters",
			Handler:    _ClusterService_ListClusters_Handler,
		},
		{
			MethodName: "CreateCluster",
			Handler:    _ClusterService_CreateCluster_Handler,
		},
		{
			MethodName: "UpdateCluster",
			Handler:    _ClusterService_UpdateCluster_Handler,
		},
		{
			MethodName: "DeleteCluster",
			Handler:    _ClusterService_DeleteCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proxmox_kubernetes_engine/v1/clusters.proto",
}