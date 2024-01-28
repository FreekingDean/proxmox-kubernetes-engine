package machines

import (
	"context"
	"fmt"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
)

func (s *Service) ListMachines(ctx context.Context, req *v1.ListMachinesRequest) (*v1.ListMachinesResponse, error) {
	rn := v1.MachinePoolAssignmentResourceName{}
	err := rn.UnmarshalString(req.Parent)
	if err != nil {
		return nil, fmt.Errorf("error parsing resource name: %w", err)
	}
	machines, nextToken, err := s.store.ListMachines(ctx, rn)
	if err != nil {
		err = fmt.Errorf("error retreiving machines: %w", err)
	}
	return &v1.ListMachinesResponse{
		Machines:      machines,
		NextPageToken: nextToken,
	}, err
}

func (s *Service) GetMachine(ctx context.Context, req *v1.GetMachineRequest) (*v1.Machine, error) {
	rn := v1.MachineResourceName{}
	err := rn.UnmarshalString(req.Name)
	if err != nil {
		return nil, err
	}
	return s.store.FindMachine(ctx, rn)
}
