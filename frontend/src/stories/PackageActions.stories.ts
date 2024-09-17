import type { Meta, StoryObj } from '@storybook/react'
import PackageActions from '../components/package-actions'

let meta: Meta = {
    title: 'components/package-actions',
    component: PackageActions,
}
export default meta

type Story = StoryObj<typeof PackageActions>

export const Default: Story = {
    args: {
        id: 'akd1',
        onHandleCheckUpdate: (id: string) => console.log(id),
        onHandleUpdate: (id: string) => console.log(id),
        onHandleRemove: (id: string) => console.log(id),
    },
}
