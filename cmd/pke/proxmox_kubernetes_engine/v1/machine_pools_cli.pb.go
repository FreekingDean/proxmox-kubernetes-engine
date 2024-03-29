// Code generated by protoc-gen-go-aip-cli. DO NOT EDIT.
package proxmox_kubernetes_enginev1

import (
	cobra "github.com/spf13/cobra"
	aipcli "go.einride.tech/aip-cli/aipcli"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func NewMachinePoolServiceCommand(config aipcli.Config) *cobra.Command {
	return aipcli.NewServiceCommand(
		config,
		File_proxmox_kubernetes_engine_v1_machine_pools_proto.
			Services().ByName("MachinePoolService"),
		map[protoreflect.FullName]string{
			"proxmox_kubernetes_engine.v1.MachinePoolService": "",
		},
		aipcli.NewMethodCommand(
			config,
			File_proxmox_kubernetes_engine_v1_machine_pools_proto.
				Services().ByName("MachinePoolService").Methods().ByName("GetMachinePool"),
			&GetMachinePoolRequest{},
			&MachinePool{},
			map[protoreflect.FullName]string{
				"proxmox_kubernetes_engine.v1.GetMachinePoolRequest.name":        "",
				"proxmox_kubernetes_engine.v1.MachinePoolService.GetMachinePool": "",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_proxmox_kubernetes_engine_v1_machine_pools_proto.
				Services().ByName("MachinePoolService").Methods().ByName("ListMachinePools"),
			&ListMachinePoolsRequest{},
			&ListMachinePoolsResponse{},
			map[protoreflect.FullName]string{
				"proxmox_kubernetes_engine.v1.ListMachinePoolsRequest.page_size":   "",
				"proxmox_kubernetes_engine.v1.ListMachinePoolsRequest.page_token":  "",
				"proxmox_kubernetes_engine.v1.MachinePoolService.ListMachinePools": "",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_proxmox_kubernetes_engine_v1_machine_pools_proto.
				Services().ByName("MachinePoolService").Methods().ByName("CreateMachinePool"),
			&CreateMachinePoolRequest{},
			&MachinePool{},
			map[protoreflect.FullName]string{
				"proxmox_kubernetes_engine.v1.CreateMachinePoolRequest.machine_pool":    "",
				"proxmox_kubernetes_engine.v1.CreateMachinePoolRequest.machine_pool_id": "",
				"proxmox_kubernetes_engine.v1.MachinePool.cpus":                         "",
				"proxmox_kubernetes_engine.v1.MachinePool.group":                        "",
				"proxmox_kubernetes_engine.v1.MachinePool.image":                        "",
				"proxmox_kubernetes_engine.v1.MachinePool.machines":                     "",
				"proxmox_kubernetes_engine.v1.MachinePool.memory":                       "",
				"proxmox_kubernetes_engine.v1.MachinePool.name":                         "",
				"proxmox_kubernetes_engine.v1.MachinePool.version":                      "",
				"proxmox_kubernetes_engine.v1.MachinePoolService.CreateMachinePool":     "",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_proxmox_kubernetes_engine_v1_machine_pools_proto.
				Services().ByName("MachinePoolService").Methods().ByName("UpdateMachinePool"),
			&UpdateMachinePoolRequest{},
			&MachinePool{},
			map[protoreflect.FullName]string{
				"google.protobuf.FieldMask.paths":                                    " The set of field mask paths.\n",
				"proxmox_kubernetes_engine.v1.MachinePool.cpus":                      "",
				"proxmox_kubernetes_engine.v1.MachinePool.group":                     "",
				"proxmox_kubernetes_engine.v1.MachinePool.image":                     "",
				"proxmox_kubernetes_engine.v1.MachinePool.machines":                  "",
				"proxmox_kubernetes_engine.v1.MachinePool.memory":                    "",
				"proxmox_kubernetes_engine.v1.MachinePool.name":                      "",
				"proxmox_kubernetes_engine.v1.MachinePool.version":                   "",
				"proxmox_kubernetes_engine.v1.MachinePoolService.UpdateMachinePool":  "",
				"proxmox_kubernetes_engine.v1.UpdateMachinePoolRequest.machine_pool": "",
				"proxmox_kubernetes_engine.v1.UpdateMachinePoolRequest.update_mask":  "",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_proxmox_kubernetes_engine_v1_machine_pools_proto.
				Services().ByName("MachinePoolService").Methods().ByName("DeleteMachinePool"),
			&DeleteMachinePoolRequest{},
			&emptypb.Empty{},
			map[protoreflect.FullName]string{
				"proxmox_kubernetes_engine.v1.DeleteMachinePoolRequest.name":        "",
				"proxmox_kubernetes_engine.v1.MachinePoolService.DeleteMachinePool": "",
			},
		),
	)
}
