package store

import (
	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/huandu/go-sqlbuilder"
)

type filterOption func(*sqlbuilder.SelectBuilder)

func MachineStateFilter(states ...v1.State) filterOption {
	return func(sb *sqlbuilder.SelectBuilder) {
		stateEqStr := make([]string, len(states))
		for i, state := range states {
			stateEqStr[i] = sb.Equal("state", state.String())
		}
		sb.Where(
			sb.Or(stateEqStr...),
		)
	}
}

func MachinePoolFilter(rn *v1.MachinePoolResourceName) filterOption {
	return func(sb *sqlbuilder.SelectBuilder) {
		sb.Where(sb.Equal("machine_pool_id", rn.MachinePool))
	}
}

func MachinePoolAssignmentFilter(rn *v1.MachinePoolAssignmentResourceName) filterOption {
	return func(sb *sqlbuilder.SelectBuilder) {
		sb.Where(sb.Equal("machine_pool_assignment_id", rn.MachinePoolAssignment))
	}
}

func EqualFilter(clauses map[string]interface{}) filterOption {
	return func(sb *sqlbuilder.SelectBuilder) {
		for k, v := range clauses {
			sb.Where(sb.Equal(k, v))
		}
	}
}

func MachineMachinePoolFilter(rn *v1.MachinePoolResourceName) filterOption {
	return func(sb *sqlbuilder.SelectBuilder) {
		sb.Join(
			"machine_pool_assignments",
			"machines.machine_pool_assignment_id = machine_pool_assignments.id",
		)
		sb.Where(sb.Equal("machine_pool_assignments.machine_pool_id", rn.MachinePool))
	}
}
