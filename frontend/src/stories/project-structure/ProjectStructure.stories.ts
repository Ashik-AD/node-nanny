import type { Meta, StoryObj } from '@storybook/react'

import ProjectStructure from '../../components/project-structure'

let meta: Meta<typeof ProjectStructure> = {
    title: 'Components/Project-structure',
    component: ProjectStructure,
}
export default meta

type Story = StoryObj<typeof meta>

export const SingleProject: Story = {
    args: {
        id: '123l2',
        path: '/home/lazyduck/dev/node-nanny',
        name: 'node-nanny',
        files: [
            {
                type: 'emv',
                name: '.env',
            },
            {
                type: 'json',
                name: 'package.json',
            },
            {
                type: 'md',
                name: 'README.md',
            },
        ],
    },
}

export const NestedProject: Story = {
    args: {
        id: '123l2',
        path: '/home/lazyduck/dev/node-nanny',
        name: 'node-nanny',
        files: [
            {
                type: 'emv',
                name: '.env',
            },
            {
                type: 'json',
                name: 'package.json',
            },
            {
                type: 'md',
                name: 'README.md',
            },
        ],
        childProjects: [
            {
                id: '224323',
                path: '/home/lazyduck/dev/node-nanny/frontend',
                name: 'frontend',
                files: [
                    {
                        type: 'emv',
                        name: '.env',
                    },
                    {
                        type: 'json',
                        name: 'package.json',
                    },
                    {
                        type: 'md',
                        name: 'README.md',
                    },
                ],
            },
            {
                id: 'a24k2',
                path: '/home/lazyduck/dev/node-nanny/backend',
                name: 'backend',
                files: [
                    {
                        type: 'emv',
                        name: '.env',
                    },
                    {
                        type: 'json',
                        name: 'package.json',
                    },
                    {
                        type: 'md',
                        name: 'README.md',
                    },
                ],
            },
        ],
    },
}

export const NestedProjectLevel3: Story = {
    args: {
        id: '123l2',
        path: '/home/lazyduck/dev/node-nanny',
        name: 'node-nanny',
        files: [
            {
                type: 'emv',
                name: '.env',
            },
            {
                type: 'json',
                name: 'package.json',
            },
            {
                type: 'md',
                name: 'README.md',
            },
        ],
        childProjects: [
            {
                id: '224323',
                path: '/home/lazyduck/dev/node-nanny/frontend',
                name: 'frontend',
                childProjects: [
                    {
                        id: 'a24k2das',
                        path: '/home/lazyduck/dev/node-nanny/backend',
                        name: 'react',
                        files: [
                            {
                                type: 'emv',
                                name: '.env',
                            },
                            {
                                type: 'json',
                                name: 'package.json',
                            },
                            {
                                type: 'md',
                                name: 'README.md',
                            },
                        ],
                    },
                ],
                files: [
                    {
                        type: 'emv',
                        name: '.env',
                    },
                    {
                        type: 'json',
                        name: 'package.json',
                    },
                    {
                        type: 'md',
                        name: 'README.md',
                    },
                ],
            },
            {
                id: 'a24k2',
                path: '/home/lazyduck/dev/node-nanny/backend',
                name: 'backend',
                files: [
                    {
                        type: 'emv',
                        name: '.env',
                    },
                    {
                        type: 'json',
                        name: 'package.json',
                    },
                    {
                        type: 'md',
                        name: 'README.md',
                    },
                ],
            },
        ],
    },
}

export const NestedProjectLevel4: Story = {
    args: {
        id: '123l2',
        path: '/home/lazyduck/dev/node-nanny',
        name: 'node-nanny',
        files: [
            {
                type: 'emv',
                name: '.env',
            },
            {
                type: 'json',
                name: 'package.json',
            },
            {
                type: 'md',
                name: 'README.md',
            },
        ],
        childProjects: [
            {
                id: '224323',
                path: '/home/lazyduck/dev/node-nanny/frontend',
                name: 'frontend',
                childProjects: [
                    {
                        id: 'a24k2das',
                        path: '/home/lazyduck/dev/node-nanny/backend',
                        name: 'react',
                        files: [
                            {
                                type: 'emv',
                                name: '.env',
                            },
                            {
                                type: 'json',
                                name: 'package.json',
                            },
                            {
                                type: 'md',
                                name: 'README.md',
                            },
                        ],
                        childProjects: [
                            {
                                id: 'a24k2da33',
                                path: '/home/lazyduck/dev/node-nanny/backend',
                                name: 'solid',
                                files: [
                                    {
                                        type: 'emv',
                                        name: '.env',
                                    },
                                    {
                                        type: 'json',
                                        name: 'package.json',
                                    },
                                    {
                                        type: 'md',
                                        name: 'README.md',
                                    },
                                ],
                            },
                        ],
                    },
                ],
                files: [
                    {
                        type: 'emv',
                        name: '.env',
                    },
                    {
                        type: 'json',
                        name: 'package.json',
                    },
                    {
                        type: 'md',
                        name: 'README.md',
                    },
                ],
            },
            {
                id: 'a24k2',
                path: '/home/lazyduck/dev/node-nanny/backend',
                name: 'backend',
                files: [
                    {
                        type: 'emv',
                        name: '.env',
                    },
                    {
                        type: 'json',
                        name: 'package.json',
                    },
                    {
                        type: 'md',
                        name: 'README.md',
                    },
                ],
            },
        ],
    },
}
