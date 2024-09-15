import type { Meta, StoryObj } from '@storybook/react'

import DependencyPieChart from '../components/charts/dependency-chart'

let meta: Meta = {
    title: 'components/charts/Dependency',
    component: DependencyPieChart,
}
export default meta

type Story = StoryObj<typeof DependencyPieChart>

export const Default: Story = {
    args: {
        data: [3, 8],
    },
}
