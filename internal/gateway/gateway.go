package gateway

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/config"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New(n net.Listener) (*runtime.ServeMux, *grpc.ClientConn, error) {
	opt := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	}
	conn, err := grpc.NewClient(
		n.Addr().String(),
		opt...,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to dial server: %w", err)
	}
	return runtime.NewServeMux(), conn, nil
}

type RunParam struct {
	fx.In

	Mux    *runtime.ServeMux
	Config config.Config
	Logger logger.Logger
	Conn   *grpc.ClientConn
}

func Run(lc fx.Lifecycle, p RunParam) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := http.ListenAndServe(fmt.Sprintf(":%d", p.Config.GatewayPort),
					loggingHandler(p.Logger,
						corsHandler(p.Mux),
					),
				)
				if err != nil {
					p.Logger.Errorf("failed to serve: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			p.Conn.Close()
			return nil
		},
	})
}

func loggingHandler(l logger.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.Infof("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func corsHandler(next http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	})
	return c.Handler(next)
}
