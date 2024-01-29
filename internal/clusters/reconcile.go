package clusters

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
)

func (s *Service) Reconcile(ctx context.Context, cluster *v1.Cluster) error {
	// TODO maybe status updates? But should be nothing to do here.
	return nil
}
