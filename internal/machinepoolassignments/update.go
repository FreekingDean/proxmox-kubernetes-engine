package machinepoolassignments

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/pkg/errors"
)

func (s *Service) UpdateMachinePoolAssignment(ctx context.Context, req *v1.UpdateMachinePoolAssignmentRequest) (*v1.MachinePoolAssignment, error) {
	rn := &v1.MachinePoolAssignmentResourceName{}
	err := rn.UnmarshalString(req.MachinePoolAssignment.Name)
	err = s.store.UpdateMachinePoolAssignment(ctx, rn.MachinePoolAssignment, req.MachinePoolAssignment)
	if err != nil {
		return nil, errors.Wrap(err, "couldnt update machine pool assignment in store")
	}
	return req.MachinePoolAssignment, err
}
