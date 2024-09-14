import type { Meta, StoryObj } from '@storybook/react'

import WizardSuccess from '../../components/project-wizard/wizard-success'

let meta: Meta<typeof WizardSuccess> = {
    title: 'Components/Project-wizard/success',
    component: WizardSuccess,
}
export default meta

type Story = StoryObj<typeof meta>
export const Default: Story = {}
