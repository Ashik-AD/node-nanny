import type { Meta, StoryObj } from '@storybook/react'

import EnvItem from '../../components/env-variables/env-item'

let meta: Meta = {
    title: 'components/env.Variables/env-item',
    component: EnvItem,
}
export default meta

type Story = StoryObj<typeof EnvItem>

export const Default: Story = {
    args: {
        env: {
            key: 'MY_SECRETE',
            secret: 'ohmygud',
        },
        onHandleEditEnv: () => console.log('Implement env. edit function'),
    },
}
