package cmd

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/config"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/grpc"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/logger"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/machines"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/nodepools"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/proxmox"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/store"
)

func RunServer() {
	log := logger.NewLogger()
	app := fx.New(
		fx.Provide(context.Background),
		fx.Provide(func() logger.Logger { return log }),
		fx.WithLogger(func(l logger.Logger) fxevent.Logger { return log }),
		config.Module,
		fx.Provide(store.New),
		proxmox.Module,
		fx.Provide(grpc.New),
		fx.Provide(machines.New),
		fx.Provide(nodepools.New),
		fx.Invoke(machines.Register),
		fx.Invoke(nodepools.Register),
		fx.Invoke(grpc.Run),
	)
	app.Run()
}
