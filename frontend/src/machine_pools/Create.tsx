import * as React from 'react';

import { Create, SimpleForm, SelectInput, TextInput, NumberInput, required } from 'react-admin';

const MachinePoolCreate = () => (
    <Create>
        <SimpleForm>
            <TextInput source="id" validate={[required()]} />
            <TextInput source="displayName" validate={[required()]} />
            <SelectInput source="image" choices={[
              { id: 'coreos', name: 'Fedora CoreOS' }
            ]} validate={[required()]}/>
            <NumberInput source="memory" validate={[required()]} min={512} step={512}/>
            <NumberInput source="cpus" validate={[required()]} max={8}/>
            <TextInput source="group" validate={[required()]} />
        </SimpleForm>
    </Create>
);

export default MachinePoolCreate
