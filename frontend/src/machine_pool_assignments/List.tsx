import { List, Datagrid, TextField } from 'react-admin';
import { useNavigate, useParams } from 'react-router-dom';

export const MachinePoolAssignmentsList = () => {
    const { clusterId } = useParams();
    const navigate = useNavigate();
    const rowClick = (id, resource, record) => {
      navigate("/"+record.name)
    }
    return (
      <List resource="machinePoolAssignments" filter={{ parent: `clusters/${clusterId }`}}>
            <Datagrid rowClick={rowClick}>
                <TextField source="id" />
                <TextField source="machinePool" />
                <TextField source="role" />
                <TextField source="count" />
            </Datagrid>
        </List>
    );
};
