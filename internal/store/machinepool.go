package store

import (
	"context"

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

func (s *Store) FindMachinePool(name string) (*v1.MachinePool, error) {
	return nil, nil
}
