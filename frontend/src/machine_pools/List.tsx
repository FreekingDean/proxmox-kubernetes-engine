import { List, Datagrid, TextField } from 'react-admin';
import { useParams } from 'react-router-dom';

const MachinePoolsList = () => (
    <List>
        <Datagrid rowClick="show">
                <TextField source="id" />
                <TextField source="image" />
                <TextField source="memory" />
                <TextField source="cpus" />
                <TextField source="group" />
        </Datagrid>
    </List>
);

export default MachinePoolsList 
