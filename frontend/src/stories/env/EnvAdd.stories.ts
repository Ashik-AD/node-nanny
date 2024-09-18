import type { Meta, StoryObj } from '@storybook/react'

import EnvAdd from '../../components/env-variables/env-add'

let meta: Meta = {
    title: 'components/env.Variables/env-add',
    component: EnvAdd,
}
export default meta

type Story = StoryObj<typeof EnvAdd>

export const Default: Story = {
    args: {},
}
