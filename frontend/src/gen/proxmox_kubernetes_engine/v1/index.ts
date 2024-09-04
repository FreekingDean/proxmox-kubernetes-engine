// Example: 'clusters/{cluster}'
export interface ClusterResourceName {
  cluster: string;
  toString(): string;
}

interface ClusterResourceNameConstructor {
  parse(s: string): ClusterResourceName;
  from(cluster: string): ClusterResourceName;
}

export const ClusterResourceName: ClusterResourceNameConstructor = {
  parse(s: string): ClusterResourceName {
    const errPrefix = `parse resource name ${s} as proxomx-kubernetes-engine.io/Cluster`;
    const segments = s.split("/")
    if (segments.length !== 2) {
      throw new Error(`${errPrefix}: invalid segment count ${segments.length} (expected 2)`)
    }
    if (segments[0] !== "clusters") {
      throw new Error(`${errPrefix}: invalid constant segment ${segments[0]} (expected clusters)`)
    }
    const cluster = segments[1]
    return this.from(cluster)
  },

  from(cluster: string): ClusterResourceName {
    if (cluster === "" || cluster.indexOf("/") > -1) {
      throw new Error(`invalid variable segment for cluster: '${cluster}'`)
    }
    return {
      cluster,
      toString(): string {
        // eslint-disable-next-line no-useless-concat, prefer-template
        return "clusters" + "/" + cluster
      },
    }
  },
}

// Example: 'clusters/{cluster}/machinePoolAssignments/{machine_pool_assignment}/machines/{machine}'
export interface MachineResourceName {
  cluster: string;
  machinePoolAssignment: string;
  machine: string;
  machinePoolAssignmentResourceName(): MachinePoolAssignmentResourceName;
  clusterResourceName(): ClusterResourceName;
  toString(): string;
}

interface MachineResourceNameConstructor {
  parse(s: string): MachineResourceName;
  from(cluster: string, machinePoolAssignment: string, machine: string): MachineResourceName;
}

export const MachineResourceName: MachineResourceNameConstructor = {
  parse(s: string): MachineResourceName {
    const errPrefix = `parse resource name ${s} as proxomx-kubernetes-engine.io/Machine`;
    const segments = s.split("/")
    if (segments.length !== 6) {
      throw new Error(`${errPrefix}: invalid segment count ${segments.length} (expected 6)`)
    }
    if (segments[0] !== "clusters") {
      throw new Error(`${errPrefix}: invalid constant segment ${segments[0]} (expected clusters)`)
    }
    const cluster = segments[1]
    if (segments[2] !== "machinePoolAssignments") {
      throw new Error(`${errPrefix}: invalid constant segment ${segments[2]} (expected machinePoolAssignments)`)
    }
    const machinePoolAssignment = segments[3]
    if (segments[4] !== "machines") {
      throw new Error(`${errPrefix}: invalid constant segment ${segments[4]} (expected machines)`)
    }
    const machine = segments[5]
    return this.from(cluster, machinePoolAssignment, machine)
  },

  from(cluster: string, machinePoolAssignment: string, machine: string): MachineResourceName {
    if (cluster === "" || cluster.indexOf("/") > -1) {
      throw new Error(`invalid variable segment for cluster: '${cluster}'`)
    }
    if (machinePoolAssignment === "" || machinePoolAssignment.indexOf("/") > -1) {
      throw new Error(`invalid variable segment for machine_pool_assignment: '${machinePoolAssignment}'`)
    }
    if (machine === "" || machine.indexOf("/") > -1) {
      throw new Error(`invalid variable segment for machine: '${machine}'`)
    }
    return {
      cluster,
      machinePoolAssignment,
      machine,
      machinePoolAssignmentResourceName(): MachinePoolAssignmentResourceName {
        return MachinePoolAssignmentResourceName.from(cluster, machinePoolAssignment)
      },
      clusterResourceName(): ClusterResourceName {
        return ClusterResourceName.from(cluster)
      },
      toString(): string {
        // eslint-disable-next-line no-useless-concat, prefer-template
        return "clusters" + "/" + cluster + "/" + "machinePoolAssignments" + "/" + machinePoolAssignment + "/" + "machines" + "/" + machine
      },
    }
  },
}

// Example: 'machinePools/{machine_pool}'
export interface MachinePoolResourceName {
  machinePool: string;
  toString(): string;
}

interface MachinePoolResourceNameConstructor {
  parse(s: string): MachinePoolResourceName;
  from(machinePool: string): MachinePoolResourceName;
}

export const MachinePoolResourceName: MachinePoolResourceNameConstructor = {
  parse(s: string): MachinePoolResourceName {
    const errPrefix = `parse resource name ${s} as proxomx-kubernetes-engine.io/MachinePool`;
    const segments = s.split("/")
    if (segments.length !== 2) {
      throw new Error(`${errPrefix}: invalid segment count ${segments.length} (expected 2)`)
    }
    if (segments[0] !== "machinePools") {
      throw new Error(`${errPrefix}: invalid constant segment ${segments[0]} (expected machinePools)`)
    }
    const machinePool = segments[1]
    return this.from(machinePool)
  },

  from(machinePool: string): MachinePoolResourceName {
    if (machinePool === "" || machinePool.indexOf("/") > -1) {
      throw new Error(`invalid variable segment for machine_pool: '${machinePool}'`)
    }
    return {
      machinePool,
      toString(): string {
        // eslint-disable-next-line no-useless-concat, prefer-template
        return "machinePools" + "/" + machinePool
      },
    }
  },
}

// Example: 'clusters/{cluster}/machinePoolAssignments/{machine_pool_assignment}'
export interface MachinePoolAssignmentResourceName {
  cluster: string;
  machinePoolAssignment: string;
  clusterResourceName(): ClusterResourceName;
  toString(): string;
}

interface MachinePoolAssignmentResourceNameConstructor {
  parse(s: string): MachinePoolAssignmentResourceName;
  from(cluster: string, machinePoolAssignment: string): MachinePoolAssignmentResourceName;
}

export const MachinePoolAssignmentResourceName: MachinePoolAssignmentResourceNameConstructor = {
  parse(s: string): MachinePoolAssignmentResourceName {
    const errPrefix = `parse resource name ${s} as proxomx-kubernetes-engine.io/MachinePoolAssignment`;
    const segments = s.split("/")
    if (segments.length !== 4) {
      throw new Error(`${errPrefix}: invalid segment count ${segments.length} (expected 4)`)
    }
    if (segments[0] !== "clusters") {
      throw new Error(`${errPrefix}: invalid constant segment ${segments[0]} (expected clusters)`)
    }
    const cluster = segments[1]
    if (segments[2] !== "machinePoolAssignments") {
      throw new Error(`${errPrefix}: invalid constant segment ${segments[2]} (expected machinePoolAssignments)`)
    }
    const machinePoolAssignment = segments[3]
    return this.from(cluster, machinePoolAssignment)
  },

  from(cluster: string, machinePoolAssignment: string): MachinePoolAssignmentResourceName {
    if (cluster === "" || cluster.indexOf("/") > -1) {
      throw new Error(`invalid variable segment for cluster: '${cluster}'`)
    }
    if (machinePoolAssignment === "" || machinePoolAssignment.indexOf("/") > -1) {
      throw new Error(`invalid variable segment for machine_pool_assignment: '${machinePoolAssignment}'`)
    }
    return {
      cluster,
      machinePoolAssignment,
      clusterResourceName(): ClusterResourceName {
        return ClusterResourceName.from(cluster)
      },
      toString(): string {
        // eslint-disable-next-line no-useless-concat, prefer-template
        return "clusters" + "/" + cluster + "/" + "machinePoolAssignments" + "/" + machinePoolAssignment
      },
    }
  },
}

