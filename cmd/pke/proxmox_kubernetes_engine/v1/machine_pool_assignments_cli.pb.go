// Code generated by protoc-gen-go-aip-cli. DO NOT EDIT.
package proxmox_kubernetes_enginev1

import (
	cobra "github.com/spf13/cobra"
	aipcli "go.einride.tech/aip-cli/aipcli"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func NewMachinePoolAssignmentServiceCommand(config aipcli.Config) *cobra.Command {
	return aipcli.NewServiceCommand(
		config,
		File_proxmox_kubernetes_engine_v1_machine_pool_assignments_proto.
			Services().ByName("MachinePoolAssignmentService"),
		map[protoreflect.FullName]string{
			"proxmox_kubernetes_engine.v1.MachinePoolAssignmentService": "",
		},
		aipcli.NewMethodCommand(
			config,
			File_proxmox_kubernetes_engine_v1_machine_pool_assignments_proto.
				Services().ByName("MachinePoolAssignmentService").Methods().ByName("GetMachinePoolAssignment"),
			&GetMachinePoolAssignmentRequest{},
			&MachinePoolAssignment{},
			map[protoreflect.FullName]string{
				"proxmox_kubernetes_engine.v1.GetMachinePoolAssignmentRequest.name":                  "",
				"proxmox_kubernetes_engine.v1.MachinePoolAssignmentService.GetMachinePoolAssignment": "",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_proxmox_kubernetes_engine_v1_machine_pool_assignments_proto.
				Services().ByName("MachinePoolAssignmentService").Methods().ByName("ListMachinePoolAssignments"),
			&ListMachinePoolAssignmentsRequest{},
			&ListMachinePoolAssignmentsResponse{},
			map[protoreflect.FullName]string{
				"proxmox_kubernetes_engine.v1.ListMachinePoolAssignmentsRequest.page_size":             "",
				"proxmox_kubernetes_engine.v1.ListMachinePoolAssignmentsRequest.page_token":            "",
				"proxmox_kubernetes_engine.v1.ListMachinePoolAssignmentsRequest.parent":                "",
				"proxmox_kubernetes_engine.v1.MachinePoolAssignmentService.ListMachinePoolAssignments": "",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_proxmox_kubernetes_engine_v1_machine_pool_assignments_proto.
				Services().ByName("MachinePoolAssignmentService").Methods().ByName("CreateMachinePoolAssignment"),
			&CreateMachinePoolAssignmentRequest{},
			&MachinePoolAssignment{},
			map[protoreflect.FullName]string{
				"proxmox_kubernetes_engine.v1.CreateMachinePoolAssignmentRequest.machine_pool_assignment":    "",
				"proxmox_kubernetes_engine.v1.CreateMachinePoolAssignmentRequest.machine_pool_assignment_id": "",
				"proxmox_kubernetes_engine.v1.CreateMachinePoolAssignmentRequest.parent":                     "",
				"proxmox_kubernetes_engine.v1.MachinePoolAssignment.count":                                   "",
				"proxmox_kubernetes_engine.v1.MachinePoolAssignment.machine_pool":                            "",
				"proxmox_kubernetes_engine.v1.MachinePoolAssignment.name":                                    "",
				"proxmox_kubernetes_engine.v1.MachinePoolAssignment.role":                                    "",
				"proxmox_kubernetes_engine.v1.MachinePoolAssignmentService.CreateMachinePoolAssignment":      "",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_proxmox_kubernetes_engine_v1_machine_pool_assignments_proto.
				Services().ByName("MachinePoolAssignmentService").Methods().ByName("UpdateMachinePoolAssignment"),
			&UpdateMachinePoolAssignmentRequest{},
			&MachinePoolAssignment{},
			map[protoreflect.FullName]string{
				"google.protobuf.FieldMask.paths":                                                         " The set of field mask paths.\n",
				"proxmox_kubernetes_engine.v1.MachinePoolAssignment.count":                                "",
				"proxmox_kubernetes_engine.v1.MachinePoolAssignment.machine_pool":                         "",
				"proxmox_kubernetes_engine.v1.MachinePoolAssignment.name":                                 "",
				"proxmox_kubernetes_engine.v1.MachinePoolAssignment.role":                                 "",
				"proxmox_kubernetes_engine.v1.MachinePoolAssignmentService.UpdateMachinePoolAssignment":   "",
				"proxmox_kubernetes_engine.v1.UpdateMachinePoolAssignmentRequest.machine_pool_assignment": "",
				"proxmox_kubernetes_engine.v1.UpdateMachinePoolAssignmentRequest.update_mask":             "",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_proxmox_kubernetes_engine_v1_machine_pool_assignments_proto.
				Services().ByName("MachinePoolAssignmentService").Methods().ByName("DeleteMachinePoolAssignment"),
			&DeleteMachinePoolAssignmentRequest{},
			&emptypb.Empty{},
			map[protoreflect.FullName]string{
				"proxmox_kubernetes_engine.v1.DeleteMachinePoolAssignmentRequest.name":                  "",
				"proxmox_kubernetes_engine.v1.MachinePoolAssignmentService.DeleteMachinePoolAssignment": "",
			},
		),
	)
}
