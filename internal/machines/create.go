package machines

import (
	"context"
	"fmt"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/util"
	"github.com/pkg/errors"
)

func (s *Service) CreateMachine(ctx context.Context, machinePoolAssignment *v1.MachinePoolAssignment) error {
	rn := &v1.MachinePoolResourceName{}
	err := rn.UnmarshalString(machinePoolAssignment.MachinePool)
	if err != nil {
		return err
	}
	machinePool, err := s.store.FindMachinePool(ctx, rn)
	if err != nil {
		return err
	}
	node, err := s.proxmox.FindAvailableNode(ctx, machinePool.Group, int(machinePool.Memory), int(machinePool.Cpus))
	if err != nil {
		return errors.Wrap(err, "couldnt find a machine")
	}

	nprn := &v1.MachinePoolAssignmentResourceName{}
	err = nprn.UnmarshalString(machinePoolAssignment.Name)
	if err != nil {
		return errors.Wrap(err, "couldnt parse machinepool name")
	}
	resourceName := &v1.MachineResourceName{
		Cluster:               nprn.Cluster,
		MachinePoolAssignment: nprn.MachinePoolAssignment,
		Machine:               util.UniqueName(6),
	}
	machine := &v1.Machine{
		Name:   resourceName.String(),
		Node:   string(node),
		Memory: machinePool.Memory,
		Cpus:   machinePool.Cpus,
		Image:  machinePool.Image,
		Group:  machinePool.Group,
		Role:   machinePoolAssignment.Role.String(),
	}

	err = s.store.CreateMachine(ctx, machine)
	if err != nil {
		return errors.Wrap(err, "couldnt create machine in store")
	}

	fmt.Println(machine)
	return nil
}
