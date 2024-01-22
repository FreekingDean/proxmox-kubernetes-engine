package machines

import (
	"context"
	"fmt"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/util"
	"github.com/pkg/errors"
)

func (s *Service) CreateMachine(ctx context.Context, machinePool *v1.MachinePool) error {
	node, err := s.proxmox.FindAvailableNode(ctx, machinePool.Group, int(machinePool.Memory), int(machinePool.Cpus))
	if err != nil {
		return errors.Wrap(err, "couldnt find a machine")
	}

	nprn := &v1.MachinePoolResourceName{}
	err = nprn.UnmarshalString(machinePool.Name)
	if err != nil {
		return errors.Wrap(err, "couldnt parse machinepool name")
	}
	rn := &v1.MachineResourceName{
		MachinePool: nprn.MachinePool,
		Machine:     util.UniqueName(6),
	}
	machine := &v1.Machine{
		Name:        rn.String(),
		CurrentNode: string(node),
		Memory:      machinePool.Memory,
		Cpus:        machinePool.Cpus,
	}

	err = s.store.CreateMachine(ctx, machine)
	if err != nil {
		return errors.Wrap(err, "couldnt create machine in store")
	}

	fmt.Println(machine)
	return nil
}
