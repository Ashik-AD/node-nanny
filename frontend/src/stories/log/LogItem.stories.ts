import type { Meta, StoryObj } from '@storybook/react'

import LogItem from '../../components/log/log-item'

let meta: Meta = {
    title: 'components/log/log-item',
    component: LogItem,
}
export default meta

type Story = StoryObj<typeof LogItem>

export const Default: Story = {
    args: {
        text: "Failed to install 'reactjs'",
        status: 'error',
        timestamp: new Date().toISOString(),
    },
}
