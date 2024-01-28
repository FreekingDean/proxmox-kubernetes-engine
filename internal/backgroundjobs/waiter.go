package backgroundjobs

import (
	"context"
	"fmt"
	"time"
)

func (s *Service) WaitForJobs() {
	for {
		s.checkMachinePoolAssignments()
		time.Sleep(1 * time.Second)
	}
}

func (s *Service) checkMachinePoolAssignments() {
	ctx := context.Background()
	machinePoolAssignments, err := s.store.ListOpenMachinePoolAssignments(ctx)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error listing machines for bgs: %s", err.Error()))
	}
	for _, machinePoolAssignment := range machinePoolAssignments {
		err = s.machines.CreateMachine(ctx, machinePoolAssignment)
		if err != nil {
			s.logger.Error(fmt.Sprintf("error creating machine %s", err.Error()))
			continue
		}
	}
}
