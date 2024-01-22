package nodepoolassignments

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	v1.UnimplementedNodePoolAssignmentServiceServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateNodePoolAssignment(ctx context.Context, req *v1.CreateNodePoolAssignmentRequest) (*v1.NodePoolAssignment, error) {
	return nil, nil
}

func (s *Service) DeleteNodePoolAssignment(ctx context.Context, req *v1.DeleteNodePoolAssignmentRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *Service) ListNodePoolAssignments(ctx context.Context, req *v1.ListNodePoolAssignmentsRequest) (*v1.ListNodePoolAssignmentsResponse, error) {
	return nil, nil
}

func (s *Service) GetNodePoolAssignment(ctx context.Context, req *v1.GetNodePoolAssignmentRequest) (*v1.NodePoolAssignment, error) {
	return nil, nil
}

func (s *Service) UpdateNodePoolAssignment(ctx context.Context, req *v1.UpdateNodePoolAssignmentRequest) (*v1.NodePoolAssignment, error) {
	return nil, nil
}
