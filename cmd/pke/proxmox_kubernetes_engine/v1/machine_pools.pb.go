// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: proxmox_kubernetes_engine/v1/machine_pools.proto

package proxmox_kubernetes_enginev1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetMachinePoolRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetMachinePoolRequest) Reset() {
	*x = GetMachinePoolRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMachinePoolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMachinePoolRequest) ProtoMessage() {}

func (x *GetMachinePoolRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMachinePoolRequest.ProtoReflect.Descriptor instead.
func (*GetMachinePoolRequest) Descriptor() ([]byte, []int) {
	return file_proxmox_kubernetes_engine_v1_machine_pools_proto_rawDescGZIP(), []int{0}
}

func (x *GetMachinePoolRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListMachinePoolsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize  int32  `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListMachinePoolsRequest) Reset() {
	*x = ListMachinePoolsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMachinePoolsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMachinePoolsRequest) ProtoMessage() {}

func (x *ListMachinePoolsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMachinePoolsRequest.ProtoReflect.Descriptor instead.
func (*ListMachinePoolsRequest) Descriptor() ([]byte, []int) {
	return file_proxmox_kubernetes_engine_v1_machine_pools_proto_rawDescGZIP(), []int{1}
}

func (x *ListMachinePoolsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListMachinePoolsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListMachinePoolsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachinePools  []*MachinePool `protobuf:"bytes,1,rep,name=machine_pools,json=machinePools,proto3" json:"machine_pools,omitempty"`
	NextPageToken string         `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListMachinePoolsResponse) Reset() {
	*x = ListMachinePoolsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMachinePoolsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMachinePoolsResponse) ProtoMessage() {}

func (x *ListMachinePoolsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMachinePoolsResponse.ProtoReflect.Descriptor instead.
func (*ListMachinePoolsResponse) Descriptor() ([]byte, []int) {
	return file_proxmox_kubernetes_engine_v1_machine_pools_proto_rawDescGZIP(), []int{2}
}

func (x *ListMachinePoolsResponse) GetMachinePools() []*MachinePool {
	if x != nil {
		return x.MachinePools
	}
	return nil
}

func (x *ListMachinePoolsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateMachinePoolRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachinePoolId string       `protobuf:"bytes,1,opt,name=machine_pool_id,json=machinePoolId,proto3" json:"machine_pool_id,omitempty"`
	MachinePool   *MachinePool `protobuf:"bytes,2,opt,name=machine_pool,json=machinePool,proto3" json:"machine_pool,omitempty"`
}

func (x *CreateMachinePoolRequest) Reset() {
	*x = CreateMachinePoolRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMachinePoolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMachinePoolRequest) ProtoMessage() {}

func (x *CreateMachinePoolRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMachinePoolRequest.ProtoReflect.Descriptor instead.
func (*CreateMachinePoolRequest) Descriptor() ([]byte, []int) {
	return file_proxmox_kubernetes_engine_v1_machine_pools_proto_rawDescGZIP(), []int{3}
}

func (x *CreateMachinePoolRequest) GetMachinePoolId() string {
	if x != nil {
		return x.MachinePoolId
	}
	return ""
}

func (x *CreateMachinePoolRequest) GetMachinePool() *MachinePool {
	if x != nil {
		return x.MachinePool
	}
	return nil
}

type UpdateMachinePoolRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachinePool *MachinePool           `protobuf:"bytes,1,opt,name=machine_pool,json=machinePool,proto3" json:"machine_pool,omitempty"`
	UpdateMask  *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateMachinePoolRequest) Reset() {
	*x = UpdateMachinePoolRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMachinePoolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMachinePoolRequest) ProtoMessage() {}

func (x *UpdateMachinePoolRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMachinePoolRequest.ProtoReflect.Descriptor instead.
func (*UpdateMachinePoolRequest) Descriptor() ([]byte, []int) {
	return file_proxmox_kubernetes_engine_v1_machine_pools_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateMachinePoolRequest) GetMachinePool() *MachinePool {
	if x != nil {
		return x.MachinePool
	}
	return nil
}

func (x *UpdateMachinePoolRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type DeleteMachinePoolRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteMachinePoolRequest) Reset() {
	*x = DeleteMachinePoolRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMachinePoolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMachinePoolRequest) ProtoMessage() {}

func (x *DeleteMachinePoolRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMachinePoolRequest.ProtoReflect.Descriptor instead.
func (*DeleteMachinePoolRequest) Descriptor() ([]byte, []int) {
	return file_proxmox_kubernetes_engine_v1_machine_pools_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteMachinePoolRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MachinePool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Image    string   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Memory   int32    `protobuf:"varint,3,opt,name=memory,proto3" json:"memory,omitempty"`
	Cpus     int32    `protobuf:"varint,4,opt,name=cpus,proto3" json:"cpus,omitempty"`
	Group    string   `protobuf:"bytes,6,opt,name=group,proto3" json:"group,omitempty"`
	Version  int32    `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty"`
	Machines []string `protobuf:"bytes,5,rep,name=machines,proto3" json:"machines,omitempty"`
}

func (x *MachinePool) Reset() {
	*x = MachinePool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachinePool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachinePool) ProtoMessage() {}

func (x *MachinePool) ProtoReflect() protoreflect.Message {
	mi := &file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachinePool.ProtoReflect.Descriptor instead.
func (*MachinePool) Descriptor() ([]byte, []int) {
	return file_proxmox_kubernetes_engine_v1_machine_pools_proto_rawDescGZIP(), []int{6}
}

func (x *MachinePool) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MachinePool) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *MachinePool) GetMemory() int32 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *MachinePool) GetCpus() int32 {
	if x != nil {
		return x.Cpus
	}
	return 0
}

func (x *MachinePool) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *MachinePool) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *MachinePool) GetMachines() []string {
	if x != nil {
		return x.Machines
	}
	return nil
}

var File_proxmox_kubernetes_engine_v1_machine_pools_proto protoreflect.FileDescriptor

var file_proxmox_kubernetes_engine_v1_machine_pools_proto_rawDesc = []byte{
	0x0a, 0x30, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1c, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xe0, 0x41, 0x02, 0xfa, 0x41,
	0x2a, 0x0a, 0x28, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x69, 0x6f, 0x2f,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x5f, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x92, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4e, 0x0a, 0x0d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78,
	0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f,
	0x6c, 0x52, 0x0c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9a, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x0f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f,
	0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x49,
	0x64, 0x12, 0x51, 0x0a, 0x0c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x6f,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f,
	0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f,
	0x6f, 0x6c, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x50, 0x6f, 0x6f, 0x6c, 0x22, 0xaf, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x51, 0x0a, 0x0c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x6f,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f,
	0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f,
	0x6f, 0x6c, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x60, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x44, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x30, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2a, 0x0a, 0x28, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f,
	0x78, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2d, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f,
	0x6f, 0x6c, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xe6, 0x02, 0x0a, 0x0b, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x44, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xe0, 0x41, 0x08, 0xfa, 0x41, 0x2a, 0x0a, 0x28,
	0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x70, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x63, 0x70, 0x75, 0x73, 0x12,
	0x19, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1d, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x08, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x3a, 0x65, 0xea, 0x41, 0x62, 0x0a,
	0x28, 0x70, 0x72, 0x6f, 0x78, 0x6f, 0x6d, 0x78, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x1b, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x7b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x7d, 0x2a, 0x0c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50,
	0x6f, 0x6f, 0x6c, 0x73, 0x32, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f,
	0x6c, 0x32, 0xf1, 0x06, 0x0a, 0x12, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9a, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x33, 0x2e, 0x70, 0x72,
	0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x22, 0x28, 0xda, 0x41, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f,
	0x6c, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x9b, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x35, 0x2e, 0x70, 0x72, 0x6f,
	0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x36, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f,
	0x6f, 0x6c, 0x73, 0x12, 0xbd, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x36, 0x2e, 0x70, 0x72, 0x6f, 0x78,
	0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x22, 0x45, 0xda, 0x41,
	0x1c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x2c, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x20, 0x3a, 0x0c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x6f,
	0x6c, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f,
	0x6f, 0x6c, 0x73, 0x12, 0xcf, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x36, 0x2e, 0x70, 0x72, 0x6f, 0x78,
	0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x22, 0x57, 0xda, 0x41,
	0x18, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x2c, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x3a,
	0x0c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x32, 0x26, 0x2f,
	0x76, 0x31, 0x2f, 0x7b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c,
	0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f,
	0x6c, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x8d, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x36, 0x2e, 0x70, 0x72,
	0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x28, 0xda, 0x41, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x2a, 0x19, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f,
	0x6c, 0x73, 0x2f, 0x2a, 0x7d, 0x42, 0xb2, 0x02, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x72, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x72, 0x65, 0x65,
	0x6b, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x61, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78,
	0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2d, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x6b, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x6d,
	0x6f, 0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x1a, 0x50, 0x72, 0x6f, 0x78,
	0x6d, 0x6f, 0x78, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x50, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x26, 0x50, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x50,
	0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proxmox_kubernetes_engine_v1_machine_pools_proto_rawDescOnce sync.Once
	file_proxmox_kubernetes_engine_v1_machine_pools_proto_rawDescData = file_proxmox_kubernetes_engine_v1_machine_pools_proto_rawDesc
)

func file_proxmox_kubernetes_engine_v1_machine_pools_proto_rawDescGZIP() []byte {
	file_proxmox_kubernetes_engine_v1_machine_pools_proto_rawDescOnce.Do(func() {
		file_proxmox_kubernetes_engine_v1_machine_pools_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxmox_kubernetes_engine_v1_machine_pools_proto_rawDescData)
	})
	return file_proxmox_kubernetes_engine_v1_machine_pools_proto_rawDescData
}

var file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proxmox_kubernetes_engine_v1_machine_pools_proto_goTypes = []interface{}{
	(*GetMachinePoolRequest)(nil),    // 0: proxmox_kubernetes_engine.v1.GetMachinePoolRequest
	(*ListMachinePoolsRequest)(nil),  // 1: proxmox_kubernetes_engine.v1.ListMachinePoolsRequest
	(*ListMachinePoolsResponse)(nil), // 2: proxmox_kubernetes_engine.v1.ListMachinePoolsResponse
	(*CreateMachinePoolRequest)(nil), // 3: proxmox_kubernetes_engine.v1.CreateMachinePoolRequest
	(*UpdateMachinePoolRequest)(nil), // 4: proxmox_kubernetes_engine.v1.UpdateMachinePoolRequest
	(*DeleteMachinePoolRequest)(nil), // 5: proxmox_kubernetes_engine.v1.DeleteMachinePoolRequest
	(*MachinePool)(nil),              // 6: proxmox_kubernetes_engine.v1.MachinePool
	(*fieldmaskpb.FieldMask)(nil),    // 7: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),            // 8: google.protobuf.Empty
}
var file_proxmox_kubernetes_engine_v1_machine_pools_proto_depIdxs = []int32{
	6, // 0: proxmox_kubernetes_engine.v1.ListMachinePoolsResponse.machine_pools:type_name -> proxmox_kubernetes_engine.v1.MachinePool
	6, // 1: proxmox_kubernetes_engine.v1.CreateMachinePoolRequest.machine_pool:type_name -> proxmox_kubernetes_engine.v1.MachinePool
	6, // 2: proxmox_kubernetes_engine.v1.UpdateMachinePoolRequest.machine_pool:type_name -> proxmox_kubernetes_engine.v1.MachinePool
	7, // 3: proxmox_kubernetes_engine.v1.UpdateMachinePoolRequest.update_mask:type_name -> google.protobuf.FieldMask
	0, // 4: proxmox_kubernetes_engine.v1.MachinePoolService.GetMachinePool:input_type -> proxmox_kubernetes_engine.v1.GetMachinePoolRequest
	1, // 5: proxmox_kubernetes_engine.v1.MachinePoolService.ListMachinePools:input_type -> proxmox_kubernetes_engine.v1.ListMachinePoolsRequest
	3, // 6: proxmox_kubernetes_engine.v1.MachinePoolService.CreateMachinePool:input_type -> proxmox_kubernetes_engine.v1.CreateMachinePoolRequest
	4, // 7: proxmox_kubernetes_engine.v1.MachinePoolService.UpdateMachinePool:input_type -> proxmox_kubernetes_engine.v1.UpdateMachinePoolRequest
	5, // 8: proxmox_kubernetes_engine.v1.MachinePoolService.DeleteMachinePool:input_type -> proxmox_kubernetes_engine.v1.DeleteMachinePoolRequest
	6, // 9: proxmox_kubernetes_engine.v1.MachinePoolService.GetMachinePool:output_type -> proxmox_kubernetes_engine.v1.MachinePool
	2, // 10: proxmox_kubernetes_engine.v1.MachinePoolService.ListMachinePools:output_type -> proxmox_kubernetes_engine.v1.ListMachinePoolsResponse
	6, // 11: proxmox_kubernetes_engine.v1.MachinePoolService.CreateMachinePool:output_type -> proxmox_kubernetes_engine.v1.MachinePool
	6, // 12: proxmox_kubernetes_engine.v1.MachinePoolService.UpdateMachinePool:output_type -> proxmox_kubernetes_engine.v1.MachinePool
	8, // 13: proxmox_kubernetes_engine.v1.MachinePoolService.DeleteMachinePool:output_type -> google.protobuf.Empty
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proxmox_kubernetes_engine_v1_machine_pools_proto_init() }
func file_proxmox_kubernetes_engine_v1_machine_pools_proto_init() {
	if File_proxmox_kubernetes_engine_v1_machine_pools_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMachinePoolRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMachinePoolsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMachinePoolsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMachinePoolRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMachinePoolRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMachinePoolRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachinePool); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proxmox_kubernetes_engine_v1_machine_pools_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proxmox_kubernetes_engine_v1_machine_pools_proto_goTypes,
		DependencyIndexes: file_proxmox_kubernetes_engine_v1_machine_pools_proto_depIdxs,
		MessageInfos:      file_proxmox_kubernetes_engine_v1_machine_pools_proto_msgTypes,
	}.Build()
	File_proxmox_kubernetes_engine_v1_machine_pools_proto = out.File
	file_proxmox_kubernetes_engine_v1_machine_pools_proto_rawDesc = nil
	file_proxmox_kubernetes_engine_v1_machine_pools_proto_goTypes = nil
	file_proxmox_kubernetes_engine_v1_machine_pools_proto_depIdxs = nil
}