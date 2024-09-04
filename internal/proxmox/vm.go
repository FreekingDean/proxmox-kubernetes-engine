package proxmox

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	ignition "github.com/coreos/ignition/v2/config/v3_4/types"

	"github.com/FreekingDean/proxmox-api-go/proxmox"
	"github.com/FreekingDean/proxmox-api-go/proxmox/cluster"
	"github.com/FreekingDean/proxmox-api-go/proxmox/nodes/qemu"
	"github.com/FreekingDean/proxmox-api-go/proxmox/nodes/qemu/status"
	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
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

func (c *Client) CreateVM(ctx context.Context, machine *v1.Machine) error {
	rpmInstallUnit := `
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

	cfgStr, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	fwstr := fmt.Sprintf(
		"name=opt/com.coreos/config,string='%s'",
		strings.Replace(string(cfgStr), ",", ",,", -1),
	)
	req := qemu.CreateRequest{
		Vmid:   int(machine.Vmid),
		Args:   proxmox.String(fmt.Sprintf("-fw_cfg %s", fwstr)),
		Name:   proxmox.String(machine.DisplayName),
		Node:   machine.Node,
		Memory: proxmox.Int(int(machine.Memory)),
		Cores:  proxmox.Int(int(machine.Cpus)),
		Nets: &qemu.Nets{
			&qemu.Net{
				Model:  qemu.NetModel_VIRTIO,
				Bridge: proxmox.String("vmbr0"),
				Tag:    proxmox.Int(2),
			},
		},
		Agent: &qemu.Agent{
			Enabled: *proxmox.PVEBool(true),
		},
		Serials: &qemu.Serials{proxmox.String("socket")},
		Scsis: &qemu.Scsis{
			&qemu.Scsi{
				File:       "local:0",
				ImportFrom: proxmox.String("local:999/fedora-coreos-39.20231119.3.0-qemu.x86_64.qcow2"),
			},
		},
	}

	_, err = c.qemu.Create(context.Background(), req)
	return err
}

func (c *Client) PrepareVM(ctx context.Context, machine *v1.Machine) error {
	err := c.qemu.ResizeVm(context.Background(), qemu.ResizeVmRequest{
		Disk: "scsi0",
		Node: machine.Node,
		Vmid: int(machine.Vmid),
		Size: fmt.Sprintf("%dG", 64),
	})
	if err != nil {
		return err
	}

	_, err = c.status.VmStart(ctx, status.VmStartRequest{
		Node: machine.Node,
		Vmid: int(machine.Vmid),
	})
	return err
}

func (c *Client) StopVM(ctx context.Context, machine *v1.Machine) error {
	_, err := c.status.VmStop(ctx, status.VmStopRequest{
		Node: machine.Node,
		Vmid: int(machine.Vmid),
	})
	return err
}

func (c *Client) DestroyVM(ctx context.Context, machine *v1.Machine) error {
	_, err := c.qemu.Delete(ctx, qemu.DeleteRequest{
		Node:                     machine.Node,
		Vmid:                     int(machine.Vmid),
		DestroyUnreferencedDisks: proxmox.PVEBool(true),
		Purge:                    proxmox.PVEBool(true),
	})
	return err
}
