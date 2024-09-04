package machinepools

import (
	"context"
	"fmt"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
)

//func (s *Service) DeleteMachinePool(ctx context.Context, req *v1.DeleteMachinePoolRequest) (*emptypb.Empty, error) {
//	return nil, nil
//}
//

func (s *Service) ListMachinePools(ctx context.Context, req *v1.ListMachinePoolsRequest) (*v1.ListMachinePoolsResponse, error) {
	machinePools, nextToken, err := s.store.ListMachinePools(ctx)
	if err != nil {
		err = fmt.Errorf("error retreiving machines: %w", err)
	}
	return &v1.ListMachinePoolsResponse{
		MachinePools:  machinePools,
		NextPageToken: nextToken,
	}, err
}

//
//func (s *Service) GetMachinePool(ctx context.Context, req *v1.GetMachinePoolRequest) (*v1.MachinePool, error) {
//	return nil, nil
//}
//
//func (s *Service) UpdateMachinePool(ctx context.Context, req *v1.UpdateMachinePoolRequest) (*v1.MachinePool, error) {
//	return nil, nil
//}
