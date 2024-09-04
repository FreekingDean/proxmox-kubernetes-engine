import {
  Admin,
  Resource,
  ListGuesser,
  EditGuesser,
  ShowGuesser,
} from "react-admin";
import { MachinePoolAssignmentsList } from './machine_pool_assignments/List';
import MachinePoolAssignmentShow from './machine_pool_assignments/Show';
import { MachinesList } from './machines/List';
import ClusterList from './clusters/List'
import ClusterShow from './clusters/Show'
import { Route } from 'react-router-dom';
import { dataProvider } from "./dataProvider";
import CloudIcon from '@mui/icons-material/Cloud';
import WaveIcon from '@mui/icons-material/Waves';

export const App = () => (
  <Admin dataProvider={dataProvider}>
    <Resource
      name="clusters"
      list={ClusterList}
      show={ClusterShow}
      icon={CloudIcon }
    >
      <Route path=":clusterId/machinePoolAssignments" element={<MachinePoolAssignmentsList/>} />
      <Route path=":clusterId/machinePoolAssignments/:machinePoolAssignmentId" element={<MachinePoolAssignmentShow/>} />
      <Route path=":clusterId/machinePoolAssignments/:machinePoolAssignmentId/show" element={<MachinePoolAssignmentShow/>} />
      <Route path=":clusterId/machinePoolAssignments/:machinePoolAssignmentId/machines" element={<MachinesList/>} />
    </Resource>

    <Resource
      name="machinePools"
      list={ListGuesser}
      edit={EditGuesser}
      show={ShowGuesser}
      icon={WaveIcon}
    />
  </Admin>
);
