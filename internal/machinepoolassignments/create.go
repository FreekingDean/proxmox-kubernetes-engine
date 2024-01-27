package machinepoolassignments

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/pkg/errors"
)

func (s *Service) CreateMachinePoolAssignment(ctx context.Context, req *v1.CreateMachinePoolAssignmentRequest) (*v1.MachinePoolAssignment, error) {
	err := s.store.CreateMachinePoolAssignment(ctx, req.MachinePoolAssignmentId, req.Parent, req.MachinePoolAssignment)
	if err != nil {
		return nil, errors.Wrap(err, "couldnt create machine pool assignment in store")
	}
	return req.MachinePoolAssignment, err
}
