package machines

import (
	"context"

	"go.uber.org/fx"
	"google.golang.org/grpc"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/logger"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/proxmox"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/store"
)

type Service struct {
	v1.UnimplementedMachineServiceServer

	proxmox *proxmox.Client
	store   *store.Store
}

type ServiceParams struct {
	fx.In

	Logger  logger.Logger
	Proxmox *proxmox.Client
	Store   *store.Store
}

func Register(lc fx.Lifecycle, s *Service, g *grpc.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			v1.RegisterMachineServiceServer(g, s)
			return nil
		},
	})
}

func New(p ServiceParams) *Service {
	s := &Service{
		proxmox: p.Proxmox,
		store:   p.Store,
	}

	return s
}
