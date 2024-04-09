package proxmox

import (
	"context"
	"fmt"

	ignition "github.com/coreos/ignition/v2/config/v3_4/types"

	"github.com/FreekingDean/proxmox-api-go/proxmox/cluster"
	"github.com/FreekingDean/proxmox-api-go/proxmox/nodes/qemu/status"
)

type State string

const (
	StateRunning State = "RUNNING"
	StateStopped State = "STOPPED"
)

func (c *Client) FindVM(ctx context.Context, vmid int) (string, error) {
	resp, err := c.cluster.Resources(ctx, cluster.ResourcesRequest{
		Type: cluster.PtrType(cluster.Type_VM),
	})
	if err != nil {
		return "", err
	}
	for _, vm := range resp {
		if vm.Vmid != nil && *vm.Vmid == vmid {
			if vm.Node != nil {
				return "", fmt.Errorf("could not parse node from resources response")
			}
			return *vm.Node, nil
		}
	}
	return "", fmt.Errorf("Could not find VM in cluster by id %d", vmid)
}

func (c *Client) GetVMState(ctx context.Context, node string, vmid int) (State, error) {
	resp, err := c.status.VmStatusCurrent(ctx, status.VmStatusCurrentRequest{
		Node: node,
		Vmid: vmid,
	})
	if err != nil {
		return "", err
	}
	if resp.Status == status.Status_STOPPED {
		return StateStopped, nil
	}
	if resp.Status != status.Status_RUNNING {
		return "", fmt.Errorf("Got bad state %s", resp.Status)
	}

	return StateRunning, nil
}

func (c *Client) CreateVM(ctx context.Context) error {
	rpminstallUnit := `
[Unit]
Description=Layer qemu-guest-agent with rpm-ostree
Wants=network-online.target
After=network-online.target
Before=zincati.service
ConditionPathExists=!/var/lib/%N.stamp

[Service]
Type=oneshot
RemainAfterExit=yes
ExecStart=/usr/bin/rpm-ostree install --apply-live --allow-inactive qemu-guest-agent
ExecStart=/bin/systemctl --now enable qemu-guest-agent
ExecStart=/bin/touch /var/lib/%N.stamp

[Install]
WantedBy=multi-user.target
`
	tvalue := true
	cfg := &ignition.Config{
		Systemd: ignition.Systemd{
			Units: []ignition.Unit{
				ignition.Unit{
					Name:     "rpm-ostree-install-qemu-guest-agent.service",
					Enabled:  &tvalue,
					Contents: &rpmInstallUnit,
				},
			},
		},
		Passwd: ignition.Passwd{
			Users: []ignition.PasswdUser{
				ignition.PasswdUser{
					Name:   "pke-user",
					Groups: []ignition.Group{"wheel", "sudo"},
					//SSHAuthorizedKeys: keys,
				},
			},
		},
		Ignition: ignition.Ignition{
			Version: "3.4.0",
		},
	}

}
