import type { Meta, StoryObj } from '@storybook/react'

import ProjectScan from '../../components/project-wizard/project-scan'

let meta: Meta<typeof ProjectScan> = {
    title: 'Components/Project-wizard/scan',
    component: ProjectScan,
}
export default meta

type Story = StoryObj<typeof meta>
export const Default: Story = {}
