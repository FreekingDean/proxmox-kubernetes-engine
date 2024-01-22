package machinepools

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/logger"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/machines"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/store"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type Service struct {
	v1.UnimplementedMachinePoolServiceServer

	Logger         logger.Logger
	MachineService *machines.Service
	store          *store.Store
}

type ServiceParams struct {
	fx.In

	Logger         logger.Logger
	MachineService *machines.Service
	Store          *store.Store
}

func New(p ServiceParams) *Service {
	return &Service{
		Logger:         p.Logger,
		MachineService: p.MachineService,
		store:          p.Store,
	}
}

func Register(lc fx.Lifecycle, s *Service, g *grpc.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			v1.RegisterMachinePoolServiceServer(g, s)
			return nil
		},
	})
}
