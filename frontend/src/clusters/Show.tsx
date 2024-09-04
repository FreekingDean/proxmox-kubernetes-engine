import { Show, SimpleShowLayout, TextField, useRecordContext } from 'react-admin';
import { Button } from '@mui/material';
import { Link } from 'react-router-dom';
import UserIcon from '@mui/icons-material/People';


const ShowMachinePoolAssignmentsButton = () => {
    const cluster = useRecordContext();

    return (
        <Button
            component={Link}
            to={`/clusters/${cluster?.id}/machinePoolAssignments`}
            startIcon={<UserIcon />}
        >
            Machine Pool Assignments
        </Button>
    );
};

export default () => (
  <Show>
    <SimpleShowLayout>
        <TextField source="name" />
        <TextField source="kubernetes_version" />
        <ShowMachinePoolAssignmentsButton />
    </SimpleShowLayout>
  </Show>
);
