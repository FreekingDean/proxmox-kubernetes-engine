import * as React from 'react';

import { Create, SimpleForm, SelectInput, TextInput, required } from 'react-admin';

const ClusterCreate = () => (
    <Create>
        <SimpleForm>
            <TextInput source="id" validate={[required()]} />
            <TextInput source="displayName" validate={[required()]} />
            <SelectInput source="kubernetesVersion" choices={[
              { id: 'k3s-v1.28.5+k3s1', name: 'K3s v1.28' }
            ]} validate={[required()]}/>
        </SimpleForm>
    </Create>
);

export default ClusterCreate
