import {
  Admin,
  Resource,
  ListGuesser,
  EditGuesser,
  ShowGuesser,
} from "react-admin";
import { MachinePoolAssignmentsList } from './machine_pool_assignments/List';
import MachinePoolAssignmentShow from './machine_pool_assignments/Show';
import MachinePoolAssignmentCreate from './machine_pool_assignments/Create';
import { MachinesList } from './machines/List';
import ClusterList from './clusters/List'
import ClusterShow from './clusters/Show'
import ClusterCreate from './clusters/Create'
import MachinePoolsList from './machine_pools/List';
import MachinePoolCreate from './machine_pools/Create';
import { Route } from 'react-router-dom';
import { dataProvider } from "./dataProvider";
import CloudIcon from '@mui/icons-material/Cloud';
import WaveIcon from '@mui/icons-material/Waves';

export const App = () => (
  <Admin dataProvider={dataProvider}>
    <Resource
      name="clusters"
      options={{ label: "Clusters" }}
      list={ClusterList}
      show={ClusterShow}
      create={ClusterCreate}
      icon={CloudIcon }
    >
      <Route path=":clusterId/machinePoolAssignments" element={<MachinePoolAssignmentsList/>} />
      <Route path=":clusterId/machinePoolAssignments/create" element={<MachinePoolAssignmentCreate/>} />
      <Route path=":clusterId/machinePoolAssignments/:machinePoolAssignmentId" element={<MachinePoolAssignmentShow/>} />
      <Route path=":clusterId/machinePoolAssignments/:machinePoolAssignmentId/show" element={<MachinePoolAssignmentShow/>} />
      <Route path=":clusterId/machinePoolAssignments/:machinePoolAssignmentId/machines" element={<MachinesList/>} />
    </Resource>

    <Resource
      name="machinePools"
      options={{ label: "Machine Pools" }}
      list={MachinePoolsList}
      create={MachinePoolCreate}
      edit={EditGuesser}
      show={ShowGuesser}
      icon={WaveIcon}
    />
  </Admin>
);
