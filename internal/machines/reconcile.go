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
		v1.State_ERROR:        noop,
	}

	handler, ok := handlers[machine.State]
	if !ok {
		return fmt.Errorf("No handler for state: %s", machine.State)
	}
	return handler(ctx, machine)
}

func (s *Service) createProxmoxVM(ctx context.Context, machine *v1.Machine) error {
	node, err := s.proxmox.FindAvailableNode(ctx, machine.Group, int(machine.Memory), int(machine.Cpus))
	if err != nil {
		return errors.Wrap(err, "couldnt find a host to run VM on")
	}
	machine.Node = string(node)

	id, err := s.proxmox.NextID(context.Background())
	if err != nil {
		return err
	}
	machine.Vmid = int32(id)

	return s.store.UpdateMachine(ctx, machine)
}
