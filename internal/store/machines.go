package store

import (
	"context"
	"fmt"
	"time"

	"github.com/huandu/go-sqlbuilder"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/store/models"
)

func (s *Store) CreateMachine(ctx context.Context, machine *v1.Machine) error {
	m := &models.Machine{}
	err := m.FromAPI(machine)
	if err != nil {
		return err
	}

	ms := sqlbuilder.NewStruct(new(models.Machine))
	m.State = v1.State_INITIALIZING.String()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	query, args := ms.InsertInto("machines", m).Build()
	s.logger.Debug(query)
	s.logger.Trace(fmt.Sprintf("%v", args))
	return s.Execute(ctx, query, args...)
}

func (s *Store) UpdateMachine(ctx context.Context, machine *v1.Machine) error {
	rn := &v1.MachineResourceName{}
	err := rn.UnmarshalString(machine.Name)
	if err != nil {
		return err
	}
	ms := sqlbuilder.NewStruct(new(models.Machine))
	m := &models.Machine{}
	err = m.FromAPI(machine)
	m.UpdatedAt = time.Now()
	if err != nil {
		return err
	}
	ub := ms.WithoutTag("pk").Update("machines", m)
	ub.Where(ub.Equal("id", rn.Machine))
	query, args := ub.Build()
	s.logger.Debug(query)
	s.logger.Trace(fmt.Sprintf("%v", args))
	return s.Execute(ctx, query, args...)
}

func (s *Store) ListMachines(ctx context.Context, opts ...filterOption) ([]*v1.Machine, string, error) {
	ms := sqlbuilder.NewStruct(new(models.Machine))
	sb := ms.SelectFrom("machines")
	for _, o := range opts {
		o(sb)
	}

	query, args := sb.Build()
	resp := []models.Machine{}
	s.logger.Debug(query)
	s.logger.Trace(fmt.Sprintf("%v", args))
	err := s.Select(ctx, query, &resp, args...)
	if err != nil {
		return nil, "", fmt.Errorf("error querying machines %w", err)
	}

	machines := make([]*v1.Machine, len(resp))
	nextToken := ""
	for i, m := range resp {
		machines[i], err = m.ToAPI()
		if err != nil {
			return nil, "", fmt.Errorf("error converting model: %w", err)
		}
		nextToken = m.ID
	}

	return machines, nextToken, nil
}

func (s Store) FindMachine(ctx context.Context, rn *v1.MachineResourceName) (*v1.Machine, error) {
	ms := sqlbuilder.NewStruct(new(models.Machine))
	sb := ms.SelectFrom("machines")
	sb.Where(sb.Equal("id", rn.Machine))
	query, args := sb.Build()
	s.logger.Debug(query)
	s.logger.Trace(fmt.Sprintf("%v", args))

	resp := models.Machine{}
	err := s.Find(ctx, query, &resp, args...)
	if err != nil {
		return nil, fmt.Errorf("error calling find: %w", err)
	}
	return resp.ToAPI()
}

func (s Store) FindMachinePoolForMachine(ctx context.Context, m *v1.Machine) (*v1.MachinePool, error) {
	rn := v1.MachineResourceName{}
	if err := rn.UnmarshalString(m.Name); err != nil {
		return nil, err
	}
	ms := sqlbuilder.NewStruct(new(models.MachinePool))
	sb := ms.SelectFrom("machine_pools")
	sb.Join(
		"machine_pool_assignments",
		"machine_pool_assignments.machine_pool_id = machine_pools.id",
	)
	sb.Join(
		"machines",
		"machines.machine_pool_assignment_id = machine_pool_assignments.id",
	)
	sb.Where(sb.Equal("machines.id", rn.Machine))
	query, args := sb.Build()
	s.logger.Debug(query)
	s.logger.Trace(fmt.Sprintf("%v", args))

	resp := models.MachinePool{}
	err := s.Find(ctx, query, &resp, args...)
	if err != nil {
		return nil, fmt.Errorf("error calling find: %w", err)
	}
	return resp.ToAPI()
}
