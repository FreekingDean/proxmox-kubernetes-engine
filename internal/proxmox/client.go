package proxmox

import (
	"context"
	"sync"

	"github.com/FreekingDean/proxmox-api-go/proxmox"
	"github.com/FreekingDean/proxmox-api-go/proxmox/access"
	"github.com/FreekingDean/proxmox-api-go/proxmox/cluster"
	"github.com/FreekingDean/proxmox-api-go/proxmox/cluster/ha/groups"
	"github.com/FreekingDean/proxmox-api-go/proxmox/nodes"
	"github.com/FreekingDean/proxmox-api-go/proxmox/nodes/qemu"
	"github.com/FreekingDean/proxmox-api-go/proxmox/nodes/qemu/status"
	"go.uber.org/fx"

	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/config"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/logger"
)

var Module = fx.Module("proxmox",
	fx.Provide(New),
)

type Client struct {
	client  *proxmox.Client
	nodes   *nodes.Client
	groups  *groups.Client
	qemu    *qemu.Client
	status  *status.Client
	cluster *cluster.Client
	log     logger.Logger

	vmcreationmu sync.Mutex
}

type ProxmoxParams struct {
	fx.In

	Logger logger.Logger
	Config config.Config
}

func New(p ProxmoxParams) (*Client, error) {
	ctx := context.Background()
	client := proxmox.NewClient(p.Config.ProxmoxHost)

	a := access.New(client)
	ticket, err := a.CreateTicket(ctx, access.CreateTicketRequest{
		Username: p.Config.ProxmoxUsername,
		Password: p.Config.ProxmoxPassword,
		Realm:    proxmox.String(p.Config.ProxmoxRealm),
	})
	if err != nil {
		return nil, err
	}

	client.SetCookie(*ticket.Ticket)
	client.SetCsrf(*ticket.Csrfpreventiontoken)
	return &Client{
		client:  client,
		qemu:    qemu.New(client),
		nodes:   nodes.New(client),
		groups:  groups.New(client),
		status:  status.New(client),
		cluster: cluster.New(client),
		log:     p.Logger,
	}, nil
}

func (c *Client) LockVMCreation() {
	c.vmcreationmu.Lock()
}

func (c *Client) UnlockVMCreation() {
	c.vmcreationmu.Unlock()
}
