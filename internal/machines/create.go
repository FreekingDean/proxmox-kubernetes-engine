package machines

import (
	"context"
	"fmt"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
)

func (s *Service) CreateMachine(ctx context.Context, machine *v1.Machine) error {
	err := s.machinestore.Create(ctx, machine)
	if err != nil {
		return err
	}

	node, err := s.proxmox.FindAvailableNode(ctx, machine.Group, int(machine.Memory), int(machine.Cpus))
	if err != nil {
		return err
	}

	fmt.Println(node)
	return nil
}
