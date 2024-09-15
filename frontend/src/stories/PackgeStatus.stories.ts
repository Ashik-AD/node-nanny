import type { Meta, StoryObj } from '@storybook/react'

import PackageStatus from '../components/packages-status'

let meta: Meta = {
    title: 'components/package-status',
    component: PackageStatus,
}
export default meta

type Story = StoryObj<typeof PackageStatus>

export const Default: Story = {
    args: {},
}
