package models

import v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"

type MachinePoolAssignment struct {
	ID            string `db:"id" fieldtag:"pk"`
	Name          string `db:"name"`
	MachinePoolID string `db:"machine_pool_id"`
	ClusterID     string `db:"cluster_id"`
	Count         int    `db:"machine_count"`
	Role          string `db:"role"`
}

func (machinePoolAssignment *MachinePoolAssignment) FromAPI(mpa *v1.MachinePoolAssignment) error {
	rn := v1.MachinePoolAssignmentResourceName{}
	err := rn.UnmarshalString(mpa.Name)
	if err != nil {
		return err
	}

	machinePoolAssignment.ID = rn.MachinePoolAssignment
	machinePoolAssignment.Name = mpa.Name
	machinePoolAssignment.ClusterID = rn.Cluster
	machinePoolAssignment.MachinePoolID = mpa.MachinePool
	machinePoolAssignment.Role = mpa.Role.String()
	machinePoolAssignment.Count = int(mpa.Count)

	return nil
}

func (machinePoolAssignment *MachinePoolAssignment) ToAPI() (*v1.MachinePoolAssignment, error) {
	rn := v1.MachinePoolAssignmentResourceName{
		MachinePoolAssignment: machinePoolAssignment.ID,
		Cluster:               machinePoolAssignment.ClusterID,
	}
	name, err := rn.MarshalString()
	return &v1.MachinePoolAssignment{
		Name:        name,
		MachinePool: machinePoolAssignment.MachinePoolID,
		Role:        v1.Role(v1.Role_value[machinePoolAssignment.Role]),
		Count:       int32(machinePoolAssignment.Count),
	}, err
}
