import type { Meta, StoryObj } from '@storybook/react'

import WizardFailed from '../../components/project-wizard/wizard-failed'

let meta: Meta<typeof WizardFailed> = {
    title: 'Components/Project-wizard/failed',
    component: WizardFailed,
}
export default meta

type Story = StoryObj<typeof meta>
export const Default: Story = {}
