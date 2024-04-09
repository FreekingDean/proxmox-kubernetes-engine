package models

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
)

type Machine struct {
	ID     string `db:"id" fieldtag:"pk"`
	Image  string `db:"image"`
	Name   string `db:"name"`
	Memory int    `db:"memory"`
	CPUs   int    `db:"cpus"`
	State  string `db:"state"`
	Group  string `db:"group_name"`
	Role   string `db:"role"`
	Node   string `db:"node"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	MachinePoolAssignmentID string `db:"machine_pool_assignment_id"`
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
	m.Node = machine.Node
	m.Role = machine.Role
	m.Image = machine.Image
	m.CreatedAt = machine.CreatedAt.AsTime()
	m.UpdatedAt = machine.UpdatedAt.AsTime()

	m.MachinePoolAssignmentID = rn.MachinePoolAssignment
	return nil
}

func (m *Machine) ToAPI() (*v1.Machine, error) {
	return &v1.Machine{
		Name:      m.Name,
		Cpus:      int32(m.CPUs),
		Memory:    int32(m.Memory),
		State:     v1.State(v1.State_value[m.State]),
		Group:     m.Group,
		Role:      m.Role,
		Image:     m.Image,
		CreatedAt: timestamppb.New(m.CreatedAt),
		UpdatedAt: timestamppb.New(m.UpdatedAt),
	}, nil
}
