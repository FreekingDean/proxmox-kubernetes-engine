package store

import (
	"context"
	"fmt"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/store/models"
	"github.com/huandu/go-sqlbuilder"
)

func (s *Store) CreateMachinePoolAssignment(ctx context.Context, id string, parent string, machinePoolAssignment *v1.MachinePoolAssignment) error {
	crn := &v1.ClusterResourceName{}
	err := crn.UnmarshalString(parent)
	if err != nil {
		return err
	}
	rn := &v1.MachinePoolAssignmentResourceName{
		MachinePoolAssignment: id,
		Cluster:               crn.Cluster,
	}
	machinePoolAssignment.Name = rn.String()
	mpa := &models.MachinePoolAssignment{}
	err = mpa.FromAPI(machinePoolAssignment)
	if err != nil {
		return err
	}

	nps := sqlbuilder.NewStruct(new(models.MachinePoolAssignment))
	query, args := nps.InsertInto("machine_pool_assignments", mpa).Build()
	s.logger.Debug(query)
	s.logger.Debug(fmt.Sprintf("%v", args))
	return s.Execute(ctx, query, args...)
}

func (s *Store) UpdateMachinePoolAssignment(ctx context.Context, id string, machinePoolAssignment *v1.MachinePoolAssignment) error {
	mpa := &models.MachinePoolAssignment{}
	err := mpa.FromAPI(machinePoolAssignment)
	if err != nil {
		return err
	}

	nps := sqlbuilder.NewStruct(new(models.MachinePoolAssignment))
	ub := nps.WithoutTag("pk").Update("machine_pool_assignments", mpa)
	query, args := ub.Where(ub.Equal("id", id)).Build()
	s.logger.Debug(query)
	s.logger.Debug(fmt.Sprintf("%v", args))
	return s.Execute(ctx, query, args...)
}

func (s *Store) ListOpenMachinePoolAssignments(ctx context.Context) ([]*v1.MachinePoolAssignment, error) {
	mpa := sqlbuilder.NewStruct(new(models.MachinePoolAssignment))
	qb := mpa.SelectFrom("machine_pool_assignments")
	qb.JoinWithOption(
		sqlbuilder.InnerJoin,
		"machine_pools",
		"machine_pools.id = machine_pool_assignments.machine_pool_id",
	)
	qb.JoinWithOption(
		sqlbuilder.LeftOuterJoin,
		"machines",
		"machines.machine_pool_assignment_id = machine_pool_assignments.id",
		"machines.role = machine_pool_assignments.role",
		"machines.image = machine_pools.image",
		"machines.memory = machine_pools.memory",
		"machines.cpus = machine_pools.cpus",
		"machines.group_name = machine_pools.group_name",
		"machines.state NOT in ('DESTROYED', 'ERROR')",
	)
	qb.GroupBy(
		"machines.machine_pool_assignment_id",
		"machines.role", "machines.image",
		"machines.memory", "machines.cpus",
		"machines.group_name",
	)
	qb.Having(
		"count(machines.id) < machine_pool_assignments.machine_count",
	)
	query, args := qb.Build()
	s.logger.Debug(query)
	resp := []models.MachinePoolAssignment{}
	err := s.Select(ctx, query, &resp, args...)
	if err != nil {
		return nil, err
	}

	machinePoolAssignments := make([]*v1.MachinePoolAssignment, len(resp))
	for i, row := range resp {
		machinePoolAssignment, err := row.ToAPI()
		if err != nil {
			return nil, err
		}
		machinePoolAssignments[i] = machinePoolAssignment
	}
	return machinePoolAssignments, nil
}

func (s *Store) ListMachinePoolAssignments(ctx context.Context, parent *v1.ClusterResourceName) ([]*v1.MachinePoolAssignment, string, error) {
	cs := sqlbuilder.NewStruct(new(models.MachinePoolAssignment))
	sb := cs.SelectFrom("machine_pool_assignments")
	query, args := sb.Build()
	resp := []models.MachinePoolAssignment{}
	s.logger.Debug(query)
	err := s.Select(ctx, query, &resp, args...)
	if err != nil {
		return nil, "", fmt.Errorf("error querying machine_pool_assignments %w", err)
	}

	machinePoolAssignments := make([]*v1.MachinePoolAssignment, len(resp))
	nextToken := ""
	for i, m := range resp {
		machinePoolAssignments[i], err = m.ToAPI()
		if err != nil {
			return nil, "", fmt.Errorf("error converting model: %w", err)
		}
		nextToken = m.ID
	}

	return machinePoolAssignments, nextToken, nil
}

func (s *Store) FindMachinePoolAssignment(ctx context.Context, rn *v1.MachinePoolAssignmentResourceName) (*v1.MachinePoolAssignment, error) {
	cs := sqlbuilder.NewStruct(new(models.MachinePoolAssignment))
	sb := cs.SelectFrom("machine_pool_assignments")
	sb.Where(sb.Equal("id", rn.MachinePoolAssignment))
	query, args := sb.Build()
	s.logger.Debug(query)

	resp := models.MachinePoolAssignment{}
	err := s.Select(ctx, query, &resp, args...)
	if err != nil {
		return nil, err
	}

	return resp.ToAPI()
}
