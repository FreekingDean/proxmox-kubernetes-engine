// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: proxmox_kubernetes_engine/v1/clusters.proto

package v1

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

type GetClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetClusterRequest) Reset() {
	*x = GetClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterRequest) ProtoMessage() {}

func (x *GetClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterRequest.ProtoReflect.Descriptor instead.
func (*GetClusterRequest) Descriptor() ([]byte, []int) {
	return file_proxmox_kubernetes_engine_v1_clusters_proto_rawDescGZIP(), []int{0}
}

func (x *GetClusterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListClustersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize  int32  `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListClustersRequest) Reset() {
	*x = ListClustersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClustersRequest) ProtoMessage() {}

func (x *ListClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClustersRequest.ProtoReflect.Descriptor instead.
func (*ListClustersRequest) Descriptor() ([]byte, []int) {
	return file_proxmox_kubernetes_engine_v1_clusters_proto_rawDescGZIP(), []int{1}
}

func (x *ListClustersRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListClustersRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListClustersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters      []*Cluster `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	NextPageToken string     `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListClustersResponse) Reset() {
	*x = ListClustersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClustersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClustersResponse) ProtoMessage() {}

func (x *ListClustersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClustersResponse.ProtoReflect.Descriptor instead.
func (*ListClustersResponse) Descriptor() ([]byte, []int) {
	return file_proxmox_kubernetes_engine_v1_clusters_proto_rawDescGZIP(), []int{2}
}

func (x *ListClustersResponse) GetClusters() []*Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *ListClustersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId string   `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	Cluster   *Cluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *CreateClusterRequest) Reset() {
	*x = CreateClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClusterRequest) ProtoMessage() {}

func (x *CreateClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClusterRequest.ProtoReflect.Descriptor instead.
func (*CreateClusterRequest) Descriptor() ([]byte, []int) {
	return file_proxmox_kubernetes_engine_v1_clusters_proto_rawDescGZIP(), []int{3}
}

func (x *CreateClusterRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *CreateClusterRequest) GetCluster() *Cluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type UpdateClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster    *Cluster               `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateClusterRequest) Reset() {
	*x = UpdateClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClusterRequest) ProtoMessage() {}

func (x *UpdateClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClusterRequest.ProtoReflect.Descriptor instead.
func (*UpdateClusterRequest) Descriptor() ([]byte, []int) {
	return file_proxmox_kubernetes_engine_v1_clusters_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateClusterRequest) GetCluster() *Cluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *UpdateClusterRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type DeleteClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteClusterRequest) Reset() {
	*x = DeleteClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClusterRequest) ProtoMessage() {}

func (x *DeleteClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClusterRequest.ProtoReflect.Descriptor instead.
func (*DeleteClusterRequest) Descriptor() ([]byte, []int) {
	return file_proxmox_kubernetes_engine_v1_clusters_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteClusterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_proxmox_kubernetes_engine_v1_clusters_proto_rawDescGZIP(), []int{6}
}

func (x *Cluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cluster) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_proxmox_kubernetes_engine_v1_clusters_proto protoreflect.FileDescriptor

var file_proxmox_kubernetes_engine_v1_clusters_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x70,
	0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x40, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2c, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x26, 0x0a, 0x24, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f,
	0x78, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2d, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5b, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x81, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70,
	0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x26, 0x0a,
	0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x80, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x44, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x9e, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x44, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x58, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x40, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2c, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x26, 0x0a, 0x24, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78,
	0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2d, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0xbc, 0x01, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x40, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0xe0,
	0x41, 0x08, 0xfa, 0x41, 0x26, 0x0a, 0x24, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x2d, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x69, 0x6f, 0x2f, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x3a, 0x50, 0xea, 0x41, 0x4d, 0x0a, 0x24, 0x70, 0x72, 0x6f, 0x78, 0x6f, 0x6d, 0x78, 0x2d, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x69, 0x6f, 0x2f, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x12, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x7d, 0x2a,
	0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x32, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x32, 0x83, 0x06, 0x0a, 0x0e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8a, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x2f, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x24, 0xda, 0x41,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f,
	0x2a, 0x7d, 0x12, 0x8b, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78,
	0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x9e, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78,
	0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x32, 0xda,
	0x41, 0x12, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x12, 0xb0, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f,
	0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x44,
	0xda, 0x41, 0x13, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x07, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x32, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x81, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78,
	0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x24, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x17, 0x2a, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x7d, 0x42, 0x8a, 0x01, 0x0a, 0x20, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x55,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x72, 0x65, 0x65, 0x6b,
	0x69, 0x6e, 0x67, 0x44, 0x65, 0x61, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78, 0x2d,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x6d, 0x6f, 0x78,
	0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proxmox_kubernetes_engine_v1_clusters_proto_rawDescOnce sync.Once
	file_proxmox_kubernetes_engine_v1_clusters_proto_rawDescData = file_proxmox_kubernetes_engine_v1_clusters_proto_rawDesc
)

func file_proxmox_kubernetes_engine_v1_clusters_proto_rawDescGZIP() []byte {
	file_proxmox_kubernetes_engine_v1_clusters_proto_rawDescOnce.Do(func() {
		file_proxmox_kubernetes_engine_v1_clusters_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxmox_kubernetes_engine_v1_clusters_proto_rawDescData)
	})
	return file_proxmox_kubernetes_engine_v1_clusters_proto_rawDescData
}

var file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proxmox_kubernetes_engine_v1_clusters_proto_goTypes = []interface{}{
	(*GetClusterRequest)(nil),     // 0: proxmox_kubernetes_engine.v1.GetClusterRequest
	(*ListClustersRequest)(nil),   // 1: proxmox_kubernetes_engine.v1.ListClustersRequest
	(*ListClustersResponse)(nil),  // 2: proxmox_kubernetes_engine.v1.ListClustersResponse
	(*CreateClusterRequest)(nil),  // 3: proxmox_kubernetes_engine.v1.CreateClusterRequest
	(*UpdateClusterRequest)(nil),  // 4: proxmox_kubernetes_engine.v1.UpdateClusterRequest
	(*DeleteClusterRequest)(nil),  // 5: proxmox_kubernetes_engine.v1.DeleteClusterRequest
	(*Cluster)(nil),               // 6: proxmox_kubernetes_engine.v1.Cluster
	(*fieldmaskpb.FieldMask)(nil), // 7: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_proxmox_kubernetes_engine_v1_clusters_proto_depIdxs = []int32{
	6, // 0: proxmox_kubernetes_engine.v1.ListClustersResponse.clusters:type_name -> proxmox_kubernetes_engine.v1.Cluster
	6, // 1: proxmox_kubernetes_engine.v1.CreateClusterRequest.cluster:type_name -> proxmox_kubernetes_engine.v1.Cluster
	6, // 2: proxmox_kubernetes_engine.v1.UpdateClusterRequest.cluster:type_name -> proxmox_kubernetes_engine.v1.Cluster
	7, // 3: proxmox_kubernetes_engine.v1.UpdateClusterRequest.update_mask:type_name -> google.protobuf.FieldMask
	0, // 4: proxmox_kubernetes_engine.v1.ClusterService.GetCluster:input_type -> proxmox_kubernetes_engine.v1.GetClusterRequest
	1, // 5: proxmox_kubernetes_engine.v1.ClusterService.ListClusters:input_type -> proxmox_kubernetes_engine.v1.ListClustersRequest
	3, // 6: proxmox_kubernetes_engine.v1.ClusterService.CreateCluster:input_type -> proxmox_kubernetes_engine.v1.CreateClusterRequest
	4, // 7: proxmox_kubernetes_engine.v1.ClusterService.UpdateCluster:input_type -> proxmox_kubernetes_engine.v1.UpdateClusterRequest
	5, // 8: proxmox_kubernetes_engine.v1.ClusterService.DeleteCluster:input_type -> proxmox_kubernetes_engine.v1.DeleteClusterRequest
	6, // 9: proxmox_kubernetes_engine.v1.ClusterService.GetCluster:output_type -> proxmox_kubernetes_engine.v1.Cluster
	2, // 10: proxmox_kubernetes_engine.v1.ClusterService.ListClusters:output_type -> proxmox_kubernetes_engine.v1.ListClustersResponse
	6, // 11: proxmox_kubernetes_engine.v1.ClusterService.CreateCluster:output_type -> proxmox_kubernetes_engine.v1.Cluster
	6, // 12: proxmox_kubernetes_engine.v1.ClusterService.UpdateCluster:output_type -> proxmox_kubernetes_engine.v1.Cluster
	8, // 13: proxmox_kubernetes_engine.v1.ClusterService.DeleteCluster:output_type -> google.protobuf.Empty
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proxmox_kubernetes_engine_v1_clusters_proto_init() }
func file_proxmox_kubernetes_engine_v1_clusters_proto_init() {
	if File_proxmox_kubernetes_engine_v1_clusters_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterRequest); i {
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
		file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClustersRequest); i {
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
		file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClustersResponse); i {
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
		file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClusterRequest); i {
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
		file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClusterRequest); i {
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
		file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClusterRequest); i {
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
		file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
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
			RawDescriptor: file_proxmox_kubernetes_engine_v1_clusters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proxmox_kubernetes_engine_v1_clusters_proto_goTypes,
		DependencyIndexes: file_proxmox_kubernetes_engine_v1_clusters_proto_depIdxs,
		MessageInfos:      file_proxmox_kubernetes_engine_v1_clusters_proto_msgTypes,
	}.Build()
	File_proxmox_kubernetes_engine_v1_clusters_proto = out.File
	file_proxmox_kubernetes_engine_v1_clusters_proto_rawDesc = nil
	file_proxmox_kubernetes_engine_v1_clusters_proto_goTypes = nil
	file_proxmox_kubernetes_engine_v1_clusters_proto_depIdxs = nil
}
