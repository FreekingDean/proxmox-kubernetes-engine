package clusters

import (
	"context"
	"fmt"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/storage"
	sqlbuilder "github.com/huandu/go-sqlbuilder"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	v1.UnimplementedClusterServiceServer
	store *storage.Storage
}

func NewService(s *storage.Storage) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) CreateCluster(ctx context.Context, req *v1.CreateClusterRequest) (*v1.Cluster, error) {
	req.Cluster.Name = fmt.Sprintf("clusters/%s", req.ClusterId)

	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("clusters")
	ib.Cols("name", "version")
	ib.Values(req.Cluster.Name, req.Cluster.Version)
	query, args := ib.Build()
	err := s.store.Execute(ctx, query, args...)

	return req.Cluster, err
}

func (s *Service) DeleteCluster(ctx context.Context, req *v1.DeleteClusterRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *Service) ListClusters(ctx context.Context, req *v1.ListClustersRequest) (*v1.ListClustersResponse, error) {
	return nil, nil
}

func (s *Service) GetCluster(ctx context.Context, req *v1.GetClusterRequest) (*v1.Cluster, error) {
	return nil, nil
}

func (s *Service) UpdateCluster(ctx context.Context, req *v1.UpdateClusterRequest) (*v1.Cluster, error) {
	return nil, nil
}
