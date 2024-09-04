import { Show, SimpleShowLayout, TextField, useRecordContext } from 'react-admin';
import { Button } from '@mui/material';
import { Link, useParams } from 'react-router-dom';
import UserIcon from '@mui/icons-material/People';


const ShowMachinesButton = () => {
    const machinePoolAssignment = useRecordContext();

    return (
        <Button
            component={Link}
            to={`/${machinePoolAssignment.name}/machines`}
            startIcon={<UserIcon />}
        >
            Machines
        </Button>
    );
};

export default () => {
  const { machinePoolAssignmentId, clusterId} = useParams();
  return (
    <Show
      resource="machinePoolAssignments"
      id={machinePoolAssignmentId }
      queryOptions={{meta: {parent: `clusters/${clusterId}`}}}
    >
      <SimpleShowLayout>
          <TextField source="id" />
          <TextField source="machinePool" />
          <TextField source="role" />
          <TextField source="count" />
          <ShowMachinesButton />
      </SimpleShowLayout>
    </Show>
  );
}
