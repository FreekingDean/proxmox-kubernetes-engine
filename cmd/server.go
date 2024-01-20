package cmd

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/config"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/grpc"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/logger"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/machines"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/proxmox"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/storage"
)

func RunServer() {
	log := logger.NewLogger()
	app := fx.New(
		fx.Provide(func() logger.Logger { return log }),
		fx.WithLogger(func(l logger.Logger) fxevent.Logger { return log }),
		config.Module,
		storage.Module,
		grpc.Module,
		proxmox.Module,
		machines.Module,
	)
	app.Run()
}
