import type { Meta, StoryObj } from '@storybook/react'

import Card from '../components/commons/card'

let meta: Meta = {
    title: 'commons/card',
    component: Card,
}
export default meta

type Story = StoryObj<typeof Card>

export const Default: Story = {
    args: {
        title: 'My card heading',
        refresh: {
            status: 'refreshing',
            onHandleRefresh: () => console.log('refreshing'),
        },
        children: 'Hello world',
    },
}
