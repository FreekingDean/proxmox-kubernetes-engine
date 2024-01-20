package grpc

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/fx"
	"google.golang.org/grpc"

	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/config"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/logger"
)

var Module = fx.Module("grpc",
	fx.Provide(New),
	fx.Invoke(Run),
)

type ServerParams struct {
	fx.In

	Config config.Config
}

func Run(lc fx.Lifecycle, s *grpc.Server, lis net.Listener, log logger.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var err error
			go func() {
				if err = s.Serve(lis); err != nil {
					log.Error("failed to serve: %v", err)
				}
			}()
			return err
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping server gracefully")
			s.GracefulStop()
			log.Info("Server stopped")
			log.Info("Closing port")
			lis.Close()
			log.Info("Port closed")
			return nil
		},
	})
}

func New(lc fx.Lifecycle, p ServerParams) (*grpc.Server, net.Listener, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", p.Config.GRPCPort))
	if err != nil {
		return nil, nil, err
	}

	s := grpc.NewServer()
	return s, lis, nil
}
