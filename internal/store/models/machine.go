package models

import v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"

type Machine struct {
	ID                           string `db:"id" fieldtag:"pk"`
	Name                         string `db:"name"`
	CPUs                         int    `db:"cpus"`
	Memory                       int    `db:"memory"`
	State                        string `db:"state"`
	Group                        string `db:"group_name"`
	MachinePoolAssignmentID      string `db:"machine_pool_assignment_id"`
	MachinePoolAssignmentVersion int    `db:"machine_pool_assignment_version"`
	Version                      int    `db:"version"`
	MachinePoolVersion           int    `db:"machine_pool_version"`
}

func (m *Machine) FromAPI(machine *v1.Machine) error {
	rn := v1.MachineResourceName{}
	if err := rn.UnmarshalString(machine.Name); err != nil {
		return err
	}

	m.ID = rn.Machine
	m.Name = machine.Name
	m.CPUs = int(machine.Cpus)
	m.Memory = int(machine.Memory)
	m.State = machine.State.String()
	m.Group = machine.Group
	m.MachinePoolVersion = int(machine.Version)
	m.MachinePoolAssignmentID = rn.MachinePoolAssignment
	return nil
}

func (m *Machine) ToAPI() (*v1.Machine, error) {
	rn := v1.MachineResourceName{
		Machine:               m.ID,
		MachinePoolAssignment: m.MachinePoolAssignmentID,
	}
	name, err := rn.MarshalString()
	return &v1.Machine{
		Name:    name,
		Cpus:    int32(m.CPUs),
		Memory:  int32(m.Memory),
		State:   v1.State(v1.State_value[m.State]),
		Group:   m.Group,
		Version: int32(m.MachinePoolAssignmentVersion),
	}, err
}
