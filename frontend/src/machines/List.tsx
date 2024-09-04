import { List, Datagrid, DateField, TextField, TextInput, SelectArrayInput } from 'react-admin';
import { useParams } from 'react-router-dom';
import { State } from '../gen/proxmox_kubernetes_engine/v1/machines_pb';

export const MachinesList= () => {
    const { clusterId, machinePoolAssignmentId } = useParams();
    const states = Object.values(State).filter(value => typeof value === 'string');
    const stateChoices = states.map((v) => { return { id: v, name: v}; });
    const filters = [
      <SelectArrayInput labels="States" source="eq.state" choices={stateChoices} defaultValue={states}/>
    ];
    return (
      <List resource="machines" filter={{ parent: `clusters/${clusterId}/machinePoolAssignments/${machinePoolAssignmentId}`}} perPage={25} filters={filters}>
            <Datagrid rowClick="edit">
                <TextField source="id" />
                <TextField source="image" />
                <TextField source="memory" />
                <TextField source="cpus" />
                <TextField source="state" />
                <TextField source="group" />
                <TextField source="role" />
                <TextField source="node" />
                <TextField source="vmid" />
                <DateField showTime={true} source="createdAt" />
                <DateField showTime={true} source="updatedAt" />
            </Datagrid>
        </List>
    );
};
