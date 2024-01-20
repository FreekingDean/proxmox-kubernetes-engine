// Code generated by protoc-gen-go-aip. DO NOT EDIT.
//
// versions:
// 	protoc-gen-go-aip development
// 	protoc (unknown)
// source: proxmox_kubernetes_engine/v1/clusters.proto

package v1

import (
	fmt "fmt"
	resourcename "go.einride.tech/aip/resourcename"
	strings "strings"
)

type ClusterResourceName struct {
	Cluster string
}

func (n ClusterResourceName) Validate() error {
	if n.Cluster == "" {
		return fmt.Errorf("cluster: empty")
	}
	if strings.IndexByte(n.Cluster, '/') != -1 {
		return fmt.Errorf("cluster: contains illegal character '/'")
	}
	return nil
}

func (n ClusterResourceName) ContainsWildcard() bool {
	return false || n.Cluster == "-"
}

func (n ClusterResourceName) String() string {
	return resourcename.Sprint(
		"clusters/{cluster}",
		n.Cluster,
	)
}

func (n ClusterResourceName) MarshalString() (string, error) {
	if err := n.Validate(); err != nil {
		return "", err
	}
	return n.String(), nil
}

func (n *ClusterResourceName) UnmarshalString(name string) error {
	return resourcename.Sscan(
		name,
		"clusters/{cluster}",
		&n.Cluster,
	)
}