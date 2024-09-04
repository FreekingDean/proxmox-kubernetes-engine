package clusters

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/store"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type Service struct {
	v1.UnimplementedClusterServiceServer
	store *store.Store
}

type RegisterParams struct {
	fx.In

	Service      *Service
	Server       *grpc.Server
	Mux          *runtime.ServeMux
	ServerClient *grpc.ClientConn
}

func Register(lc fx.Lifecycle, p RegisterParams) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			v1.RegisterClusterServiceServer(p.Server, p.Service)
			return v1.RegisterClusterServiceHandler(
				ctx,
				p.Mux,
				p.ServerClient,
			)
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

func (s *Service) GetCluster(ctx context.Context, req *v1.GetClusterRequest) (*v1.Cluster, error) {
	rn := &v1.ClusterResourceName{}
	err := rn.UnmarshalString(req.Name)
	if err != nil {
		return nil, err
	}
	return s.store.FindCluster(ctx, rn)
}
