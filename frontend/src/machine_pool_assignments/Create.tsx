import { useState, useEffect } from 'react';

import { Create, SimpleForm, SelectInput, TextInput, NumberInput, required, useGetList } from 'react-admin';
import { useParams } from 'react-router-dom';

const MachinePoolAssignmentCreate = () => {
  const { clusterId } = useParams();
  const { data: machinePools } = useGetList("machinePools")
  const [machinePoolChoices, setMachinePoolChoices ] = useState([])
  useEffect(() => {
    if (machinePools !== undefined) {
      const tmpMachinePoolChoices = machinePools.map((mp) => {
        return {id: mp.name, name: mp.name}
      })
      setMachinePoolChoices(tmpMachinePoolChoices)
    }
  }, [machinePools])
  
  return (
    <Create resource="machinePoolAssignments">
      <SimpleForm>
        <TextInput source="parent" readOnly defaultValue={"clusters/" + clusterId} />
        <TextInput source="id" validate={[required()]} />
        <TextInput source="displayName" validate={[required()]} />
        <SelectInput source="machinePool" validate={[required()]} choices={machinePoolChoices}/>
        <SelectInput source="role" validate={[required()]} choices={[
          {id: 1, name: "control"},
          {id: 2, name: "etcd"},
          {id: 3, name: "worker"},
          {id: 4, name: "control & etcd"},
          {id: 5, name: "control & etcd & worker"},
        ]}/>
        <NumberInput source="count" />
      </SimpleForm>
    </Create>
  )
}

export default MachinePoolAssignmentCreate
