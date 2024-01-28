package machinepoolassignments

import (
	"context"
	"fmt"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
)

func (s *Service) ListMachinePoolAssignments(ctx context.Context, req *v1.ListMachinePoolAssignmentsRequest) (*v1.ListMachinePoolAssignmentsResponse, error) {
	rn := &v1.ClusterResourceName{}
	err := rn.UnmarshalString(req.Parent)
	if err != nil {
		return nil, fmt.Errorf("error parsing resource name: %w", err)
	}
	machinePoolAssignments, nextToken, err := s.store.ListMachinePoolAssignments(ctx, rn)
	if err != nil {
		err = fmt.Errorf("error retreiving machinePoolAssignments: %w", err)
	}
	return &v1.ListMachinePoolAssignmentsResponse{
		MachinePoolAssignments: machinePoolAssignments,
		NextPageToken:          nextToken,
	}, err
}

func (s *Service) GetMachinePoolAssignment(ctx context.Context, req *v1.GetMachinePoolAssignmentRequest) (*v1.MachinePoolAssignment, error) {
	rn := &v1.MachinePoolAssignmentResourceName{}
	err := rn.UnmarshalString(req.Name)
	if err != nil {
		return nil, err
	}
	return s.store.FindMachinePoolAssignment(ctx, rn)
}
