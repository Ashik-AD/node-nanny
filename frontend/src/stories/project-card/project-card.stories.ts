import { Meta, StoryObj } from '@storybook/react'

import ProjectCard from '../../components/project-card'

let meta: Meta<typeof ProjectCard> = {
    title: 'Components/project-card',
    component: ProjectCard,
}

export default meta

type Story = StoryObj<typeof ProjectCard>

export const Regular: Story = {
    args: {
        name: 'node-nany',
        logoSrc: 'https://picsum.photos/60',
        description:
            'Node project manager is an application that help developer to analyze project based on nodejs/npm',
        id: 'pid12',
        lastUpdated: new Date(),
    },
}
