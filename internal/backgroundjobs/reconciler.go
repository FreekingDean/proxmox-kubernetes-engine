package backgroundjobs

import (
	"context"
	"time"

	v1 "github.com/FreekingDean/proxmox-kubernetes-engine/gen/go/proxmox_kubernetes_engine/v1"
	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/store"
)

func (s *Service) WaitForJobs() {
	for {
		ctx := context.Background()
		s.reconcileClusters(ctx)
		s.reconcileMachinePools(ctx)
		time.Sleep(1 * time.Second)
	}
}

func (s *Service) reconcileClusters(ctx context.Context) {
	clusters, _, err := s.store.ListClusters(ctx)
	if err != nil {
		s.logger.Errorf("Error listing clusters %w", err)
	}
	for _, cluster := range clusters {
		err = s.clusters.Reconcile(ctx, cluster)
		if err != nil {
			s.logger.Errorf("Error reconciling cluster '%s': %w", cluster.Name, err)
		}
		crn := &v1.ClusterResourceName{}
		err = crn.UnmarshalString(cluster.Name)
		if err != nil {
			s.logger.Errorf("Error parsing cluster resource name %w", err)
			continue
		}
		s.reconcileMachinePoolAssignments(ctx, crn)
	}
}

func (s *Service) reconcileMachinePools(ctx context.Context) {
	machinePools, _, err := s.store.ListMachinePools(ctx)
	if err != nil {
		s.logger.Errorf("Error listing machine pools %w", err)
	}
	for _, machinePool := range machinePools {
		err = s.machinepools.Reconcile(ctx, machinePool)
		if err != nil {
			s.logger.Errorf("Error reconciling machine pool'%s': %w", machinePool.Name, err)
		}
	}
}

func (s *Service) reconcileMachinePoolAssignments(ctx context.Context, crn *v1.ClusterResourceName) {
	machinePoolAssignments, _, err := s.store.ListMachinePoolAssignments(ctx, crn)
	if err != nil {
		s.logger.Errorf("Error listing machine pool assignments %w", err)
	}

	for _, machinePoolAssignment := range machinePoolAssignments {
		err = s.machinepoolassignments.Reconcile(ctx, machinePoolAssignment)
		if err != nil {
			s.logger.Errorf("Error reconciling machine pool assignment '%s': %w", machinePoolAssignment.Name, err)
		}
		mparn := &v1.MachinePoolAssignmentResourceName{}
		err = mparn.UnmarshalString(machinePoolAssignment.Name)
		if err != nil {
			s.logger.Errorf("Error parsing machine pool assignment resource name %w", err)
			continue
		}
		s.reconcileMachines(ctx, mparn)
	}
}

func (s *Service) reconcileMachines(ctx context.Context, mparn *v1.MachinePoolAssignmentResourceName) {
	filter := store.MachinePoolAssignmentFilter(mparn)
	machines, _, err := s.store.ListMachines(ctx, filter)
	if err != nil {
		s.logger.Errorf("Error listing machines %w", err)
	}

	for _, machine := range machines {
		err = s.machines.Reconcile(ctx, machine)
		if err != nil {
			s.logger.Errorf("Error reconciling machine '%s': %s", machine.Name, err)
		}
	}
}
