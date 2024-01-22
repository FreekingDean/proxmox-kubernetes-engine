package models

import v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"

type NodePool struct {
	ID     string `db:"id" fieldtag:"pk"`
	Name   string `db:"name"`
	Memory int    `db:"memory"`
	CPUs   int    `db:"cpus"`
	Group  string `db:"group_name"`
}

func (nodePool *NodePool) FromAPI(np *v1.NodePool) error {
	rn := v1.NodePoolResourceName{}
	err := rn.UnmarshalString(np.Name)
	if err != nil {
		return err
	}

	nodePool.ID = rn.NodePool
	nodePool.Name = np.Name
	nodePool.Memory = int(np.Memory)
	nodePool.CPUs = int(np.Cpus)
	nodePool.Group = np.Group

	return nil
}
