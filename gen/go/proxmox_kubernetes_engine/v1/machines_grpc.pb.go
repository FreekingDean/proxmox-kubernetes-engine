// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proxmox_kubernetes_engine/v1/machines.proto

package v1

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
	MachineService_GetMachine_FullMethodName   = "/proxmox_kubernetes_engine.v1.MachineService/GetMachine"
	MachineService_ListMachines_FullMethodName = "/proxmox_kubernetes_engine.v1.MachineService/ListMachines"
)

// MachineServiceClient is the client API for MachineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MachineServiceClient interface {
	GetMachine(ctx context.Context, in *GetMachineRequest, opts ...grpc.CallOption) (*Machine, error)
	ListMachines(ctx context.Context, in *ListMachinesRequest, opts ...grpc.CallOption) (*ListMachinesResponse, error)
}

type machineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMachineServiceClient(cc grpc.ClientConnInterface) MachineServiceClient {
	return &machineServiceClient{cc}
}

func (c *machineServiceClient) GetMachine(ctx context.Context, in *GetMachineRequest, opts ...grpc.CallOption) (*Machine, error) {
	out := new(Machine)
	err := c.cc.Invoke(ctx, MachineService_GetMachine_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineServiceClient) ListMachines(ctx context.Context, in *ListMachinesRequest, opts ...grpc.CallOption) (*ListMachinesResponse, error) {
	out := new(ListMachinesResponse)
	err := c.cc.Invoke(ctx, MachineService_ListMachines_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MachineServiceServer is the server API for MachineService service.
// All implementations must embed UnimplementedMachineServiceServer
// for forward compatibility
type MachineServiceServer interface {
	GetMachine(context.Context, *GetMachineRequest) (*Machine, error)
	ListMachines(context.Context, *ListMachinesRequest) (*ListMachinesResponse, error)
	mustEmbedUnimplementedMachineServiceServer()
}

// UnimplementedMachineServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMachineServiceServer struct {
}

func (UnimplementedMachineServiceServer) GetMachine(context.Context, *GetMachineRequest) (*Machine, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMachine not implemented")
}
func (UnimplementedMachineServiceServer) ListMachines(context.Context, *ListMachinesRequest) (*ListMachinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMachines not implemented")
}
func (UnimplementedMachineServiceServer) mustEmbedUnimplementedMachineServiceServer() {}

// UnsafeMachineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MachineServiceServer will
// result in compilation errors.
type UnsafeMachineServiceServer interface {
	mustEmbedUnimplementedMachineServiceServer()
}

func RegisterMachineServiceServer(s grpc.ServiceRegistrar, srv MachineServiceServer) {
	s.RegisterService(&MachineService_ServiceDesc, srv)
}

func _MachineService_GetMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).GetMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineService_GetMachine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).GetMachine(ctx, req.(*GetMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineService_ListMachines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMachinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).ListMachines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineService_ListMachines_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).ListMachines(ctx, req.(*ListMachinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MachineService_ServiceDesc is the grpc.ServiceDesc for MachineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MachineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proxmox_kubernetes_engine.v1.MachineService",
	HandlerType: (*MachineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMachine",
			Handler:    _MachineService_GetMachine_Handler,
		},
		{
			MethodName: "ListMachines",
			Handler:    _MachineService_ListMachines_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proxmox_kubernetes_engine/v1/machines.proto",
}
