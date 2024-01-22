package models

import v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"

type MachinePool struct {
	ID     string `db:"id" fieldtag:"pk"`
	Name   string `db:"name"`
	Memory int    `db:"memory"`
	CPUs   int    `db:"cpus"`
	Group  string `db:"group_name"`
}

func (machinePool *MachinePool) FromAPI(np *v1.MachinePool) error {
	rn := v1.MachinePoolResourceName{}
	err := rn.UnmarshalString(np.Name)
	if err != nil {
		return err
	}

	machinePool.ID = rn.MachinePool
	machinePool.Name = np.Name
	machinePool.Memory = int(np.Memory)
	machinePool.CPUs = int(np.Cpus)
	machinePool.Group = np.Group

	return nil
}
