package machinepools

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/pkg/errors"
	"go.einride.tech/aip/fieldbehavior"
	"google.golang.org/genproto/googleapis/api/annotations"
)

func (s *Service) CreateMachinePool(ctx context.Context, req *v1.CreateMachinePoolRequest) (*v1.MachinePool, error) {
	err := fieldbehavior.ValidateRequiredFields(req)
	if err != nil {
		return nil, err
	}
	fieldbehavior.ClearFields(req, annotations.FieldBehavior_OUTPUT_ONLY)

	err = s.store.CreateMachinePool(ctx, req.MachinePoolId, req.MachinePool)
	if err != nil {
		return nil, errors.Wrap(err, "couldnt create machine pool in store")
	}
	return req.MachinePool, err
}
