import type { Meta, StoryObj } from '@storybook/react'

import Stats from '../components/stats'

let meta: Meta = {
    title: 'components/stats',
    component: Stats,
}
export default meta

type Story = StoryObj<typeof Stats>

export const Default: Story = {
    args: {},
}
