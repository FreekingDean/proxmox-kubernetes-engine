package store

import (
	"context"
	"fmt"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/store/models"
	"github.com/huandu/go-sqlbuilder"
)

func (s *Store) CreateMachinePool(ctx context.Context, machinePoolID string, machinePool *v1.MachinePool) error {
	rn := &v1.MachinePoolResourceName{MachinePool: machinePoolID}
	machinePool.Name = rn.String()
	np := &models.MachinePool{}
	err := np.FromAPI(machinePool)
	if err != nil {
		return err
	}

	nps := sqlbuilder.NewStruct(new(models.MachinePool))
	query, args := nps.InsertInto("machine_pools", np).Build()
	s.logger.Debug(query)
	return s.Execute(ctx, query, args...)
}

func (s *Store) FindMachinePool(ctx context.Context, rn *v1.MachinePoolResourceName) (*v1.MachinePool, error) {
	ms := sqlbuilder.NewStruct(new(models.MachinePool))
	sb := ms.SelectFrom("machine_pools")
	sb.Where(sb.Equal("id", rn.MachinePool))
	query, args := sb.Build()

	resp := models.MachinePool{}
	err := s.Find(ctx, query, &resp, args...)
	if err != nil {
		return nil, fmt.Errorf("error calling find: %w", err)
	}
	return resp.ToAPI()
}
