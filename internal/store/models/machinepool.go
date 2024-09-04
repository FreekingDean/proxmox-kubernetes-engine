package models

import v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"

type MachinePool struct {
	ID     string `db:"id" fieldtag:"pk"`
	Name   string `db:"name"`
	Image  string `db:"image"`
	Memory int    `db:"memory"`
	CPUs   int    `db:"cpus"`
	Group  string `db:"group_name"`
}

func (machinePool *MachinePool) FromAPI(mp *v1.MachinePool) error {
	rn := v1.MachinePoolResourceName{}
	err := rn.UnmarshalString(mp.Name)
	if err != nil {
		return err
	}

	machinePool.ID = rn.MachinePool
	machinePool.Name = mp.Name
	machinePool.Memory = int(mp.Memory)
	machinePool.CPUs = int(mp.Cpus)
	machinePool.Group = mp.Group

	return nil
}

func (machinePool *MachinePool) ToAPI() (*v1.MachinePool, error) {
	rn := &v1.MachinePoolResourceName{
		MachinePool: machinePool.ID,
	}
	name, err := rn.MarshalString()
	return &v1.MachinePool{
		Name:        name,
		Image:       machinePool.Image,
		Cpus:        int32(machinePool.CPUs),
		Memory:      int32(machinePool.Memory),
		Group:       machinePool.Group,
		DisplayName: machinePool.ID,
	}, err
}
