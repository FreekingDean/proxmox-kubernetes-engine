package machinestore

import v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"

type model struct {
	ID         string `db:"id" fieldtag:"pk"`
	Name       string `db:"name"`
	CPUs       int    `db:"cpus"`
	Memory     int    `db:"memory"`
	State      string `db:"state"`
	Group      string `db:"group"`
	Node       string `db:"node"`
	NodePoolID string `db:"node_pool_id"`
}

func apiToModel(machine *v1.Machine) (*model, error) {
	rn := v1.MachineResourceName{}
	err := rn.UnmarshalString(machine.Name)
	return &model{
		ID:         rn.Machine,
		Name:       machine.Name,
		CPUs:       int(machine.Cpus),
		Memory:     int(machine.Memory),
		State:      machine.State.String(),
		Group:      machine.Group,
		NodePoolID: rn.NodePool,
	}, err
}

func modelToAPI(m model) (*v1.Machine, error) {
	rn := v1.MachineResourceName{
		Machine:  m.ID,
		NodePool: m.NodePoolID,
	}
	name, err := rn.MarshalString()
	return &v1.Machine{
		Name:   name,
		Cpus:   int32(m.CPUs),
		Memory: int32(m.Memory),
		State:  v1.State(v1.State_value[m.State]),
		Group:  m.Group,
	}, err
}
