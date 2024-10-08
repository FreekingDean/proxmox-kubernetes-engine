package machinepoolassignments

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/store"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/util"
)

func (s *Service) Reconcile(ctx context.Context, machinePoolAssignment *v1.MachinePoolAssignment) error {
	rn := &v1.MachinePoolAssignmentResourceName{}
	err := rn.UnmarshalString(machinePoolAssignment.Name)
	if err != nil {
		return err
	}
	mprn := &v1.MachinePoolResourceName{}
	err = mprn.UnmarshalString(machinePoolAssignment.MachinePool)
	if err != nil {
		return err
	}
	mp, err := s.store.FindMachinePool(ctx, mprn)
	if err != nil {
		return err
	}
	mpaFilter := store.MachinePoolAssignmentFilter(rn)
	stateFilter := store.MachineStateFilter(
		v1.State_STARTING,
		v1.State_RUNNING,
		v1.State_CREATING,
		v1.State_CREATED,
		v1.State_UNKNOWN,
		v1.State_INITIALIZING,
	)

	machines, _, err := s.store.ListMachines(ctx, mpaFilter, stateFilter)
	if err != nil {
		return err
	}
	stoppedFilter := store.MachineStateFilter(
		v1.State_STOPPED,
	)
	stoppedMachines, _, err := s.store.ListMachines(ctx, mpaFilter, stoppedFilter)
	if err != nil {
		return err
	}

	goodCount := 0
	for _, m := range stoppedMachines {
		allSet := goodCount >= int(machinePoolAssignment.Count)
		if !isMachineRelevant(m, mp, machinePoolAssignment) || allSet {
			s.logger.Debugf("Machine Not Relevant or Not Need More %s %s", allSet, isMachineRelevant(m, mp, machinePoolAssignment))
			m.State = v1.State_TO_DESTROY
			err := s.store.UpdateMachine(ctx, m)
			if err != nil {
				return err
			}
			continue
		}
		goodCount += 1
	}
	for _, m := range machines {
		allSet := goodCount >= int(machinePoolAssignment.Count)
		if !isMachineRelevant(m, mp, machinePoolAssignment) || allSet {
			s.logger.Debugf("Machine Not Relevant or Not Need More %s %s", allSet, isMachineRelevant(m, mp, machinePoolAssignment))
			m.State = v1.State_TO_DESTROY
			err := s.store.UpdateMachine(ctx, m)
			if err != nil {
				return err
			}
			continue
		}
		goodCount += 1
	}

	toCreate := int(machinePoolAssignment.Count) - goodCount
	for i := toCreate; i > 0; i-- {
		rn := &v1.MachineResourceName{
			Cluster:               rn.Cluster,
			MachinePoolAssignment: rn.MachinePoolAssignment,
			Machine:               util.UniqueName(6),
		}
		name, err := rn.MarshalString()
		if err != nil {
			return err
		}

		machine := &v1.Machine{
			Name:   name,
			Image:  mp.Image,
			Memory: mp.Memory,
			Cpus:   mp.Cpus,
			Group:  mp.Group,
			Role:   machinePoolAssignment.Role.String(),
			State:  v1.State_INITIALIZING,
		}
		err = s.store.CreateMachine(ctx, machine)
		if err != nil {
			return err
		}
	}
	return nil
}

func isMachineRelevant(m *v1.Machine, mp *v1.MachinePool, mpa *v1.MachinePoolAssignment) bool {
	return m.Role == mpa.Role.String() &&
		m.Memory == mp.Memory &&
		m.Cpus == mp.Cpus &&
		m.Image == mp.Image &&
		m.Group == mp.Group
}
