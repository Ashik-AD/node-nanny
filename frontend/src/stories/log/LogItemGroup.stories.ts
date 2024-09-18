import type { Meta, StoryObj } from '@storybook/react'

import LogItemGroup from '../../components/log/log-group'

let meta: Meta = {
    title: 'components/log/log-group',
    component: LogItemGroup,
}
export default meta

type Story = StoryObj<typeof LogItemGroup>

export const Default: Story = {
    args: {
        masterLog: {
            text: "Failed to install 'reactjs'",
            status: 'error',
            timestamp: new Date().toISOString(),
        },
        items: [
            {
                text: "Can't reslove package",
                status: 'error',
            },
            {
                text: 'Failed to connect server',
                status: 'error',
            },
        ],
    },
}
