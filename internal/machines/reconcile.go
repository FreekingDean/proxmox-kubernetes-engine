package machines

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
)

func (s *Service) Reconcile(ctx context.Context, machine *v1.Machine) error {
	return nil
}
