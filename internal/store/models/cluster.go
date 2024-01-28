package models

import v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"

type Cluster struct {
	ID                string `db:"id" fieldtag:"pk"`
	Name              string `db:"name"`
	KubernetesVersion string `db:"kubernetes_version"`
}

func (c *Cluster) FromAPI(cluster *v1.Cluster) error {
	rn := &v1.ClusterResourceName{}
	err := rn.UnmarshalString(cluster.Name)
	if err != nil {
		return err
	}
	c.ID = rn.Cluster
	c.Name = cluster.Name
	c.KubernetesVersion = cluster.KubernetesVersion
	return nil
}

func (c *Cluster) ToAPI() (*v1.Cluster, error) {
	return &v1.Cluster{
		Name:              c.Name,
		KubernetesVersion: c.KubernetesVersion,
	}, nil
}
