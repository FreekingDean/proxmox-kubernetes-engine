package clusters

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/store"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type Service struct {
	v1.UnimplementedClusterServiceServer
	store *store.Store
}

func Register(lc fx.Lifecycle, s *Service, g *grpc.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			v1.RegisterClusterServiceServer(g, s)
			return nil
		},
	})
}

func New(s *store.Store) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) CreateCluster(ctx context.Context, req *v1.CreateClusterRequest) (*v1.Cluster, error) {
	err := s.store.CreateCluster(ctx, req.ClusterId, req.Cluster)
	return req.Cluster, err
}

func (s *Service) ListClusters(ctx context.Context, req *v1.ListClustersRequest) (*v1.ListClustersResponse, error) {
	clusters, _, err := s.store.ListClusters(ctx)
	return &v1.ListClustersResponse{Clusters: clusters}, err
}
