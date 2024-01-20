package proxmox

import (
	"context"
	"fmt"
	"math/rand"
	"slices"
	"strings"
	"time"

	"github.com/FreekingDean/proxmox-api-go/proxmox/cluster/ha/groups"
	"github.com/FreekingDean/proxmox-api-go/proxmox/nodes"
	"github.com/FreekingDean/proxmox-api-go/proxmox/nodes/qemu"
)

type NodeID string

const (
	_ = 1 << (iota * 10)
	KB
	MB
	GB
)

func (c *Client) FindAvailableNode(ctx context.Context, group string, cpus, memory int) (NodeID, error) {
	respGroup, err := c.groups.Find(ctx, groups.FindRequest{Group: group})
	if err != nil {
		return "", err
	}

	nodesStr, ok := respGroup["nodes"].(string)
	if !ok {
		return "", fmt.Errorf("bad format groups.nodes response %+v", respGroup["nodes"])
	}
	nodesList := strings.Split(strings.TrimSpace(nodesStr), ",")

	nodeResp, err := c.nodes.Index(ctx)
	if err != nil {
		return "", err
	}

	//Randomize result to add jitter
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(nodeResp), func(i, j int) { nodeResp[i], nodeResp[j] = nodeResp[j], nodeResp[i] })

	foundNode := NodeID("")
	maxAvailMem := 0
	for _, node := range nodeResp {
		if !slices.Contains(nodesList, node.Node) {
			continue
		}

		if node.Status != nodes.Status_ONLINE {
			continue
		}

		vms, err := c.qemu.Index(ctx, qemu.IndexRequest{
			Node: node.Node,
		})
		if err != nil {
			return "", err
		}

		usedCPU := 0
		usedMem := 0
		for _, vm := range vms {
			if vm.Status != "running" {
				continue
			}
			usedCPU += int(*vm.Cpus)
			usedMem += *vm.Maxmem
		}
		if *node.Maxmem-usedMem > maxAvailMem &&
			usedMem+(memory/1024)*GB < *node.Maxmem &&
			cpus+usedCPU < int(*node.Maxcpu) {
			foundNode = NodeID(node.Node)
			maxAvailMem = (*node.Maxmem) - usedMem
		}
	}

	if foundNode == "" {
		return "", fmt.Errorf("Could not find available node to schedule machine on")
	}
	return foundNode, nil
}
