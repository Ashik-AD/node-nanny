import type { Meta, StoryObj } from '@storybook/react'

import ProjectAdd from '../../components/project-wizard/project-add'

let meta: Meta<typeof ProjectAdd> = {
    title: 'Components/Project-wizard/select',
    component: ProjectAdd,
}
export default meta

type Story = StoryObj<typeof meta>
export const Default:Story = {}
