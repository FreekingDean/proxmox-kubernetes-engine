package machinestore

import (
	"context"

	"github.com/huandu/go-sqlbuilder"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/storage"
)

type Store struct {
	store *storage.Storage
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Create(ctx context.Context, machine *v1.Machine) error {
	m, err := apiToModel(machine)
	if err != nil {
		return err
	}

	ms := sqlbuilder.NewStruct(new(model))
	m.State = string(v1.State_CREATING)
	query, args := ms.InsertInto("machines", m).Build()
	return s.store.Execute(ctx, query, args...)
}

func (s *Store) List(ctx context.Context, parent v1.NodePoolResourceName) ([]*v1.Machine, string, error) {
	ms := sqlbuilder.NewStruct(new(model))
	sb := ms.SelectFrom("machines")
	sb.Where(sb.Equal("node_pool_id", parent.NodePool))
	query, args := sb.Build()
	resp := []model{}
	err := s.store.Select(ctx, query, resp, args...)
	if err != nil {
		return nil, "", err
	}

	machines := make([]*v1.Machine, len(resp))
	nextToken := ""
	for i, m := range resp {
		machines[i], err = modelToAPI(m)
		if err != nil {
			return nil, "", err
		}
		nextToken = m.ID
	}

	return machines, nextToken, nil
}

func (s Store) Find(ctx context.Context, rn v1.MachineResourceName) (*v1.Machine, error) {
	ms := sqlbuilder.NewStruct(new(model))
	sb := ms.SelectFrom("machines")
	sb.Where(sb.Equal("id", rn.Machine))
	query, args := sb.Build()

	resp := model{}
	err := s.store.Find(ctx, query, &resp, args...)
	if err != nil {
		return nil, err
	}
	return modelToAPI(resp)
}
