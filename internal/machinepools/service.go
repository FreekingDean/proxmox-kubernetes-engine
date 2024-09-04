package machinepools

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/logger"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/machines"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/store"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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

type RegisterParams struct {
	fx.In

	Service      *Service
	Server       *grpc.Server
	Mux          *runtime.ServeMux
	ServerClient *grpc.ClientConn
}

func New(p ServiceParams) *Service {
	return &Service{
		Logger:         p.Logger,
		MachineService: p.MachineService,
		store:          p.Store,
	}
}

func Register(lc fx.Lifecycle, p RegisterParams) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			v1.RegisterMachinePoolServiceServer(p.Server, p.Service)
			return v1.RegisterMachinePoolServiceHandler(
				ctx,
				p.Mux,
				p.ServerClient,
			)
		},
	})
}
