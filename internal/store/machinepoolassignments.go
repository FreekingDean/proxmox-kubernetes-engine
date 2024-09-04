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
	err := s.Find(ctx, query, &resp, args...)
	if err != nil {
		return nil, err
	}

	return resp.ToAPI()
}
