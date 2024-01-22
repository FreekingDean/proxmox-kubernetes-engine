package machinepools

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) DeleteMachinePool(ctx context.Context, req *v1.DeleteMachinePoolRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *Service) ListMachinePools(ctx context.Context, req *v1.ListMachinePoolsRequest) (*v1.ListMachinePoolsResponse, error) {
	return nil, nil
}

func (s *Service) GetMachinePool(ctx context.Context, req *v1.GetMachinePoolRequest) (*v1.MachinePool, error) {
	return nil, nil
}

func (s *Service) UpdateMachinePool(ctx context.Context, req *v1.UpdateMachinePoolRequest) (*v1.MachinePool, error) {
	return nil, nil
}
