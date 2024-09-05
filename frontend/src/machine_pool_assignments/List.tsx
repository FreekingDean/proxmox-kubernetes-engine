import { List, Datagrid, TextField, ListActions } from 'react-admin';
import { useNavigate, useParams } from 'react-router-dom';

import Empty from '../components/Empty'

export const MachinePoolAssignmentsList = () => {
    const { clusterId } = useParams();
    const navigate = useNavigate();
    const rowClick = (id, resource, record) => {
      navigate("/"+record.name)
    }
    return (
      <List resource="machinePoolAssignments" empty={<Empty prefix={`clusters/${clusterId}`}/>} actions={<ListActions hasCreate />} filter={{ parent: `clusters/${clusterId }`}}>
            <Datagrid rowClick={rowClick}>
                <TextField source="id" />
                <TextField source="machinePool" />
                <TextField source="role" />
                <TextField source="count" />
            </Datagrid>
        </List>
    );
};
