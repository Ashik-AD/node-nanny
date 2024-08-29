import React from 'react'
import type { Meta, StoryObj } from '@storybook/react'
import { useState } from 'react'

import Modal from '../../components/commons/modal'

function Composit() {
    let [showModal, setShowModal] = useState(false)
    let toggleModal = () => setShowModal((prev) => !prev)
    return (
        <div>
            <button onClick={toggleModal}>Open to click modal</button>
            <Modal
                title="Select Project"
                subtitle="Select from or drop project folder here"
                isShow={showModal}
                onClose={toggleModal}
                onClickOutsideClose={toggleModal}
            >
                <h1>Select here</h1>
            </Modal>
        </div>
    )
}

let meta: Meta<typeof Composit> = {
    title: 'Commons/Modal',
    component: Composit,
}
export default meta

type Story = StoryObj<typeof meta>

export const Default: Story = {}
