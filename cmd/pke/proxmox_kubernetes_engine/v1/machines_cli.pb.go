// Code generated by protoc-gen-go-aip-cli. DO NOT EDIT.
package proxmox_kubernetes_enginev1

import (
	cobra "github.com/spf13/cobra"
	aipcli "go.einride.tech/aip-cli/aipcli"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

func NewMachineServiceCommand(config aipcli.Config) *cobra.Command {
	return aipcli.NewServiceCommand(
		config,
		File_proxmox_kubernetes_engine_v1_machines_proto.
			Services().ByName("MachineService"),
		map[protoreflect.FullName]string{
			"proxmox_kubernetes_engine.v1.MachineService": "",
		},
		aipcli.NewMethodCommand(
			config,
			File_proxmox_kubernetes_engine_v1_machines_proto.
				Services().ByName("MachineService").Methods().ByName("GetMachine"),
			&GetMachineRequest{},
			&Machine{},
			map[protoreflect.FullName]string{
				"proxmox_kubernetes_engine.v1.GetMachineRequest.name":    "",
				"proxmox_kubernetes_engine.v1.MachineService.GetMachine": "",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_proxmox_kubernetes_engine_v1_machines_proto.
				Services().ByName("MachineService").Methods().ByName("ListMachines"),
			&ListMachinesRequest{},
			&ListMachinesResponse{},
			map[protoreflect.FullName]string{
				"proxmox_kubernetes_engine.v1.ListMachinesRequest.page_size":  "",
				"proxmox_kubernetes_engine.v1.ListMachinesRequest.page_token": "",
				"proxmox_kubernetes_engine.v1.ListMachinesRequest.parent":     "",
				"proxmox_kubernetes_engine.v1.MachineService.ListMachines":    "",
			},
		),
	)
}
