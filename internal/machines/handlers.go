package machines

import (
	"context"
	"time"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/proxmox"
)

func (s *Service) handleInitializing(ctx context.Context, machine *v1.Machine) error {
	machine.State = v1.State_CREATING
	err := s.store.UpdateMachine(ctx, machine)
	if err != nil {
		return err
	}
	return s.createProxmoxVM(ctx, machine)
}

func (s *Service) handleCreating(ctx context.Context, machine *v1.Machine) error {
	if time.Now().After(machine.CreatedAt.AsTime().Add(creatingTimeout)) {
		machine.State = v1.State_UNKNOWN
		return s.store.UpdateMachine(ctx, machine)
	}
	return nil
}

func (s *Service) handleUnknown(ctx context.Context, machine *v1.Machine) error {
	if time.Now().After(machine.CreatedAt.AsTime().Add(unknownTimeout)) {
		machine.State = v1.State_ERROR
		return s.store.UpdateMachine(ctx, machine)
	}

	node, err := s.proxmox.FindVM(ctx, int(machine.Vmid))
	if err != nil {
		return err
	}

	machine.Node = node

	return s.updateState(ctx, machine)
}

func (s *Service) updateState(ctx context.Context, machine *v1.Machine) error {
	state, err := s.proxmox.GetVMState(ctx, machine.Node, int(machine.Vmid))
	if err != nil {
		if machine.State == v1.State_UNKNOWN || machine.State == v1.State_CREATING {
			return nil
		}
		machine.State = v1.State_UNKNOWN
		return s.store.UpdateMachine(ctx, machine)
	}

	if state == proxmox.StateStopped {
		machine.State = v1.State_STOPPED
		return s.store.UpdateMachine(ctx, machine)
	}

	machine.State = v1.State_RUNNING
	return s.store.UpdateMachine(ctx, machine)

	return s.store.UpdateMachine(ctx, machine)
}
