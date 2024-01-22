package machinepoolassignments

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	v1.UnimplementedMachinePoolAssignmentServiceServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateMachinePoolAssignment(ctx context.Context, req *v1.CreateMachinePoolAssignmentRequest) (*v1.MachinePoolAssignment, error) {
	return nil, nil
}

func (s *Service) DeleteMachinePoolAssignment(ctx context.Context, req *v1.DeleteMachinePoolAssignmentRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *Service) ListMachinePoolAssignments(ctx context.Context, req *v1.ListMachinePoolAssignmentsRequest) (*v1.ListMachinePoolAssignmentsResponse, error) {
	return nil, nil
}

func (s *Service) GetMachinePoolAssignment(ctx context.Context, req *v1.GetMachinePoolAssignmentRequest) (*v1.MachinePoolAssignment, error) {
	return nil, nil
}

func (s *Service) UpdateMachinePoolAssignment(ctx context.Context, req *v1.UpdateMachinePoolAssignmentRequest) (*v1.MachinePoolAssignment, error) {
	return nil, nil
}
