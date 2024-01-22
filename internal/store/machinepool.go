package store

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/store/models"
	"github.com/huandu/go-sqlbuilder"
)

func (s *Store) CreateNodePool(ctx context.Context, nodePoolID string, nodePool *v1.NodePool) error {
	rn := &v1.NodePoolResourceName{NodePool: nodePoolID}
	nodePool.Name = rn.String()
	np := &models.NodePool{}
	err := np.FromAPI(nodePool)
	if err != nil {
		return err
	}

	nps := sqlbuilder.NewStruct(new(models.NodePool))
	query, args := nps.InsertInto("node_pools", np).Build()
	s.logger.Debug(query)
	return s.Execute(ctx, query, args...)
}

func (s *Store) FindNodePool(name string) (*v1.NodePool, error) {
	return nil, nil
}
