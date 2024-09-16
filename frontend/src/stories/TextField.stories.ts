import type { Meta, StoryObj } from '@storybook/react'

import TextField from '../components/commons/textField'

let meta: Meta = {
    title: 'commons/text-filed',
    component: TextField,
}
export default meta

type Story = StoryObj<typeof TextField>

export const Default: Story = {
    args: {
    label: "Username"
  },
}
