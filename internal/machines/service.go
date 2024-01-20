package machines

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/proxmox"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/storage/machinestore"
)

type Service struct {
	v1.UnimplementedMachineServiceServer

	proxmox      *proxmox.Client
	machinestore *machinestore.Store
}

type ServiceParams struct {
	fx.In

	Server *grpc.Server

	Proxmox      *proxmox.Client
	Machinestore *machinestore.Store
}

func New(p ServiceParams) *Service {
	s := &Service{
		proxmox:      p.Proxmox,
		machinestore: p.Machinestore,
	}
	v1.RegisterMachineServiceServer(p.Server, s)
	return s
}
