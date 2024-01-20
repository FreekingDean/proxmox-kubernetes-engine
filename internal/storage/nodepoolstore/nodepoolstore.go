package nodepoolstore

import v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"

type Store struct{}

func (s Store) FindByName(name string) (*v1.NodePool, error) {
	return nil, nil
}
