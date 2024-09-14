import type { Meta, StoryObj } from '@storybook/react'

import ProjectTopNav from '../components/project-top-nav'

let meta: Meta = {
  title: 'components/Top-Nav/project-page',
  component: ProjectTopNav
}
export default meta

type Story = StoryObj<typeof ProjectTopNav>

export const Default: Story = {
  args:{}
}
