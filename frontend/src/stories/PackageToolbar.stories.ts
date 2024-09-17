import type { Meta, StoryObj } from '@storybook/react'
import PackageToolbar from '../components/package-toolbar'

let meta: Meta = {
    title: 'components/package-toolbar',
    component: PackageToolbar,
}
export default meta

type Story = StoryObj<typeof PackageToolbar>

export const Default: Story = {
    args: {
        onHandleSearch: console.log,
        onHandleUpdateCheck: () =>
            console.log('Checking packages to be updated'),
        onHandleInstallPackage: () => console.log('Installing packges'),
        onHandleRemovePackage: () => console.log('Removing packages'),
    },
}
