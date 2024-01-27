package machinepools

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/pkg/errors"
)

func (s *Service) CreateMachinePool(ctx context.Context, req *v1.CreateMachinePoolRequest) (*v1.MachinePool, error) {
	req.MachinePool.Version = 1
	err := s.store.CreateMachinePool(ctx, req.MachinePoolId, req.MachinePool)
	if err != nil {
		return nil, errors.Wrap(err, "couldnt create machine pool in store")
	}
	return req.MachinePool, err
}
