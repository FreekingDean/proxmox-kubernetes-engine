package machinepoolassignments

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/logger"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/proxmox"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/store"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type Service struct {
	v1.UnimplementedMachinePoolAssignmentServiceServer

	proxmox *proxmox.Client
	store   *store.Store
	logger  logger.Logger
}

type ServiceParams struct {
	fx.In

	Logger  logger.Logger
	Proxmox *proxmox.Client
	Store   *store.Store
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
			v1.RegisterMachinePoolAssignmentServiceServer(p.Server, p.Service)
			return v1.RegisterMachinePoolAssignmentServiceHandler(
				ctx,
				p.Mux,
				p.ServerClient,
			)
		},
	})
}

func New(p ServiceParams) *Service {
	s := &Service{
		proxmox: p.Proxmox,
		store:   p.Store,
		logger:  p.Logger,
	}

	return s
}
