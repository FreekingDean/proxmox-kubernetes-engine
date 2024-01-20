package nodepools

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	v1.UnimplementedNodePoolServiceServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateNodePool(ctx context.Context, req *v1.CreateNodePoolRequest) (*v1.NodePool, error) {
	return nil, nil
}

func (s *Service) DeleteNodePool(ctx context.Context, req *v1.DeleteNodePoolRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *Service) ListNodePools(ctx context.Context, req *v1.ListNodePoolsRequest) (*v1.ListNodePoolsResponse, error) {
	return nil, nil
}

func (s *Service) GetNodePool(ctx context.Context, req *v1.GetNodePoolRequest) (*v1.NodePool, error) {
	return nil, nil
}

func (s *Service) UpdateNodePool(ctx context.Context, req *v1.UpdateNodePoolRequest) (*v1.NodePool, error) {
	return nil, nil
}
