package proxmox

import (
	"context"

	"github.com/FreekingDean/proxmox-api-go/proxmox/cluster"
)

func (c *Client) NextID(ctx context.Context) (int, error) {
	return c.cluster.Nextid(ctx, cluster.NextidRequest{})
}
