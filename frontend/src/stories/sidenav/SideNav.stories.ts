import { Meta, StoryObj } from '@storybook/react'

import SideNav from '../../components/sidenav/side-nav'

let story = {
    title: 'Commons/side-nav',
    component: SideNav,
} satisfies Meta<typeof SideNav>

export default story

type Story = StoryObj<typeof story>

export const Default: Story = {}
