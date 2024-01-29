package store

import (
	"context"
	"fmt"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/store/models"
	"github.com/huandu/go-sqlbuilder"
)

func (s *Store) ListClusters(ctx context.Context) ([]*v1.Cluster, string, error) {
	cs := sqlbuilder.NewStruct(new(models.Cluster))
	sb := cs.SelectFrom("clusters")
	query, args := sb.Build()
	s.logger.Debug(query)
	s.logger.Trace(fmt.Sprintf("%+v", args))
	resp := []models.Cluster{}
	err := s.Select(ctx, query, &resp, args...)
	if err != nil {
		return nil, "", fmt.Errorf("error querying clusters %w", err)
	}

	clusters := make([]*v1.Cluster, len(resp))
	nextToken := ""
	for i, m := range resp {
		clusters[i], err = m.ToAPI()
		if err != nil {
			return nil, "", fmt.Errorf("error converting model: %w", err)
		}
		nextToken = m.ID
	}

	return clusters, nextToken, nil
}

func (s *Store) CreateCluster(ctx context.Context, clusterID string, cluster *v1.Cluster) error {
	rn := &v1.ClusterResourceName{Cluster: clusterID}
	cluster.Name = rn.String()
	np := &models.Cluster{}
	err := np.FromAPI(cluster)
	if err != nil {
		return err
	}

	nps := sqlbuilder.NewStruct(new(models.Cluster))
	query, args := nps.InsertInto("clusters", np).Build()
	s.logger.Debug(query)
	return s.Execute(ctx, query, args...)
}
