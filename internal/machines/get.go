package machines

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
)

func (s *Service) ListMachines(ctx context.Context, req *v1.ListMachinesRequest) (*v1.ListMachinesResponse, error) {
	rn := v1.NodePoolResourceName{}
	err := rn.UnmarshalString(req.Parent)
	if err != nil {
		return nil, err
	}
	machines, nextToken, err := s.machinestore.List(ctx, rn)
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
	return s.machinestore.Find(ctx, rn)
}
