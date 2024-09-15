import type { Meta, StoryObj } from '@storybook/react'

import DiskUsage from '../components/charts/disk-usage'

let meta: Meta = {
    title: 'components/charts/disk-usage',
    component: DiskUsage,
}
export default meta

type Story = StoryObj<typeof DiskUsage>

export const Default: Story = {
    args: {
        data: [3, 8],
    },
}
