import type { Meta, StoryObj } from '@storybook/react'

import LogToolbar from '../../components/log/log-toolbar'

let meta: Meta = {
    title: 'components/log/log-toolbar',
    component: LogToolbar,
}
export default meta

type Story = StoryObj<typeof LogToolbar>

function onSearch() {
    console.log('searching')
}
function onFilter() {
    console.log('filtering')
}
function onSort() {
    console.log('Sorting')
}
export const Default: Story = {
    args: {
        onSearch,
        onFilter,
        onSort,
    },
}
