package machines

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/pkg/errors"
)

const (
	creatingTimeout = 5 * time.Minute
	unknownTimeout  = 10 * time.Minute
)

type stateHandlerFunc func(context.Context, *v1.Machine) error

func noop(context.Context, *v1.Machine) error { return nil }

func (s *Service) Reconcile(ctx context.Context, machine *v1.Machine) error {
	handlers := map[v1.State]stateHandlerFunc{
		v1.State_INITIALIZING: s.handleInitializing,
		v1.State_UNKNOWN:      s.handleUnknown,
		v1.State_CREATING:     s.handleCreating,
		v1.State_CREATED:      s.handleCreated,
		v1.State_STARTING:     s.updateState,
		v1.State_RUNNING:      s.updateState,
		v1.State_STOPPING:     s.updateState,
		v1.State_STOPPED:      s.updateState,
		v1.State_TO_DESTROY:   s.handleToDestroy,
		v1.State_DESTROYING:   s.updateState,
		v1.State_DESTROYED:    noop,
		v1.State_ERROR:        noop,
	}

	handler, ok := handlers[machine.State]
	if !ok {
		return fmt.Errorf("No handler for state: %s", machine.State)
	}
	return handler(ctx, machine)
}

func (s *Service) createProxmoxVM(ctx context.Context, machine *v1.Machine) error {
	s.proxmox.LockVMCreation()
	defer s.proxmox.UnlockVMCreation()

	node, err := s.proxmox.FindAvailableNode(ctx, machine.Group, int(machine.Memory), int(machine.Cpus))
	if err != nil {
		return errors.Wrap(err, "couldnt find a host to run VM on")
	}

	if node == "" {
		return errors.Wrap(err, "couldnt find a host to run VM on")
	}

	s.logger.Tracef("Found Next Node: %s", node)
	machine.Node = string(node)

	id, err := s.proxmox.NextID(context.Background())
	if err != nil {
		return err
	}

	s.logger.Tracef("Found Next VMID: %d", id)
	machine.Vmid = int32(id)

	err = s.store.UpdateMachine(ctx, machine)
	if err != nil {
		return err
	}
	return s.proxmox.CreateVM(ctx, machine)
}
