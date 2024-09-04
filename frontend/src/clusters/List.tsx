import { List, Datagrid, TextField } from 'react-admin';
import { useParams } from 'react-router-dom';

const ClusterList = () => (
    <List>
        <Datagrid rowClick="show">
            <TextField source="id" />
            <TextField source="kubernetesVersion" />
        </Datagrid>
    </List>
);

export default ClusterList
