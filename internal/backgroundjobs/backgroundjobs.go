package backgroundjobs

import (
	"context"

	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/logger"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/machines"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/store"
	"go.uber.org/fx"
)

type Service struct {
	store    *store.Store
	logger   logger.Logger
	machines *machines.Service
}

type ServiceParams struct {
	fx.In

	Store    *store.Store
	Logger   logger.Logger
	Machines *machines.Service
}

func New(p ServiceParams) *Service {
	return &Service{
		store:    p.Store,
		logger:   p.Logger,
		machines: p.Machines,
	}
}

func Run(lc fx.Lifecycle, s *Service) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go s.WaitForJobs()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
