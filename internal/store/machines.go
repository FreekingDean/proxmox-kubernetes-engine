package store

import (
	"context"
	"fmt"

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
	m.State = string(v1.State_CREATING)
	query, args := ms.InsertInto("machines", m).Build()
	s.logger.Debug(query)
	s.logger.Debug(fmt.Sprintf("%v", args))
	return s.Execute(ctx, query, args...)
}

func (s *Store) ListMachines(ctx context.Context, parent v1.MachinePoolResourceName) ([]*v1.Machine, string, error) {
	ms := sqlbuilder.NewStruct(new(models.Machine))
	sb := ms.SelectFrom("machines")
	sb.Where(sb.Equal("machine_pool_id", parent.MachinePool))
	query, args := sb.Build()
	resp := []models.Machine{}
	s.logger.Debug(query)
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

func (s Store) FindMachine(ctx context.Context, rn v1.MachineResourceName) (*v1.Machine, error) {
	ms := sqlbuilder.NewStruct(new(models.Machine))
	sb := ms.SelectFrom("machines")
	sb.Where(sb.Equal("id", rn.Machine))
	query, args := sb.Build()

	resp := models.Machine{}
	err := s.Find(ctx, query, &resp, args...)
	if err != nil {
		return nil, fmt.Errorf("error calling find: %w", err)
	}
	return resp.ToAPI()
}
