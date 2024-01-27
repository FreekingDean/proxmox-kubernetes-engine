package cmd

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/clusters"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/config"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/grpc"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/logger"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/machinepoolassignments"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/machinepools"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/machines"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/proxmox"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/store"
	"github.com/spf13/cobra"
)

func serverCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "run the server",
		Run: func(cmd *cobra.Command, args []string) {
			log := logger.NewLogger()
			app := fx.New(
				fx.Provide(cmd.Context),
				fx.Provide(func() logger.Logger { return log }),
				fx.WithLogger(func(l logger.Logger) fxevent.Logger { return log }),
				config.Module,
				fx.Provide(store.New),
				proxmox.Module,
				fx.Provide(grpc.New),

				fx.Provide(clusters.New),
				fx.Provide(machines.New),
				fx.Provide(machinepools.New),
				fx.Provide(machinepoolassignments.New),

				fx.Invoke(clusters.Register),
				fx.Invoke(machines.Register),
				fx.Invoke(machinepools.Register),
				fx.Invoke(machinepoolassignments.Register),

				fx.Invoke(grpc.Run),
			)
			app.Run()
		},
	}
}
