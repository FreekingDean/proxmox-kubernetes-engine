package machines

import (
	"context"
	"fmt"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/util"
	"github.com/pkg/errors"
)

func (s *Service) CreateMachine(ctx context.Context, nodePool *v1.NodePool) error {
	node, err := s.proxmox.FindAvailableNode(ctx, nodePool.Group, int(nodePool.Memory), int(nodePool.Cpus))
	if err != nil {
		return errors.Wrap(err, "couldnt find a node")
	}

	nprn := &v1.NodePoolResourceName{}
	err = nprn.UnmarshalString(nodePool.Name)
	if err != nil {
		return errors.Wrap(err, "couldnt parse nodepool name")
	}
	rn := &v1.MachineResourceName{
		NodePool: nprn.NodePool,
		Machine:  util.UniqueName(6),
	}
	machine := &v1.Machine{
		Name:        rn.String(),
		CurrentNode: string(node),
		Memory:      nodePool.Memory,
		Cpus:        nodePool.Cpus,
	}

	err = s.store.CreateMachine(ctx, machine)
	if err != nil {
		return errors.Wrap(err, "couldnt create machine in store")
	}

	fmt.Println(node)
	return nil
}
