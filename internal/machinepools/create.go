package nodepools

import (
	"context"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/pkg/errors"
	"go.einride.tech/aip/fieldbehavior"
	"google.golang.org/genproto/googleapis/api/annotations"
)

func (s *Service) CreateNodePool(ctx context.Context, req *v1.CreateNodePoolRequest) (*v1.NodePool, error) {
	err := fieldbehavior.ValidateRequiredFields(req)
	if err != nil {
		return nil, err
	}
	fieldbehavior.ClearFields(req, annotations.FieldBehavior_OUTPUT_ONLY)

	err = s.store.CreateNodePool(ctx, req.NodePoolId, req.NodePool)
	if err != nil {
		return nil, errors.Wrap(err, "couldnt create node pool in store")
	}
	return req.NodePool, err
}
