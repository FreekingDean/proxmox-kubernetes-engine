package machinepools

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/store"
)

func (s *Service) Reconcile(ctx context.Context, machinePool *v1.MachinePool) error {
	rn := &v1.MachinePoolResourceName{}
	err := rn.UnmarshalString(machinePool.Name)
	if err != nil {
		return err
	}

	stateFilter := store.MachineStateFilter(
		v1.State_RUNNING,
		v1.State_STARTING,
		v1.State_CREATING,
		v1.State_CREATED,
		v1.State_INITIALIZING,
	)
	machinePoolFilter := store.MachineMachinePoolFilter(rn)
	machines, _, err := s.store.ListMachines(ctx, machinePoolFilter, stateFilter)
	if err != nil {
		return err
	}
	for _, machine := range machines {
		if machineShouldReconcile(machine, machinePool) {
			oldMachine := *machine
			newMachine := *machine
			newMachine.Image = machinePool.Image
			newMachine.Memory = machinePool.Memory
			newMachine.Cpus = machinePool.Cpus
			newMachine.Group = machinePool.Group

			if machine.State == v1.State_INITIALIZING {
				err := s.store.UpdateMachine(ctx, &newMachine)
				if err != nil {
					return err
				}
				continue
			}

			oldMachine.State = v1.State_DESTROYING
			newMachine.State = v1.State_INITIALIZING
			err := s.store.Transaction(ctx, func(tx *store.Store) error {
				err := tx.UpdateMachine(ctx, &oldMachine)
				if err != nil {
					return err
				}
				return tx.CreateMachine(ctx, &newMachine)
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func machineShouldReconcile(machine *v1.Machine, machinePool *v1.MachinePool) bool {
	return machine.Image != machinePool.Image ||
		machine.Memory != machinePool.Memory ||
		machine.Cpus != machinePool.Cpus ||
		machine.Group != machinePool.Group
}
