import { useState } from 'react'
import { Button } from '@nextui-org/react'
import Stack from '../../commons/stack'

import { IoIosCheckmarkCircle } from 'react-icons/io'

import Modal from '../../commons/modal'

export default function WizardSuccess() {
    let [isVisible, setIsVisible] = useState(true)

    let [reason, setReason] = useState<string>(
        'The directory dont contains any node/npm project'
    )

    function handleToggleModal() {
        setIsVisible((prev) => !prev)
    }

    return (
        <div>
            <Button color="primary" onClick={handleToggleModal}>
                Open project scanner
            </Button>

            <Modal onClose={handleToggleModal} isShow={isVisible}>
                <Stack gap="36px" alignItems="start">
                    <Stack gap="8px" alignItems="center">
                        <IoIosCheckmarkCircle
                            size={64}
                            color="var(--success-base)"
                        />
                        <span className="text-large">
                            Project added successfully
                        </span>
                    </Stack>
                    <Stack dir="column" alignItems="center">
                        <Button color="default" radius="sm">
                            Close
                        </Button>
                        <Button color="secondary" radius="sm">
                            Go to project
                        </Button>
                    </Stack>
                </Stack>
            </Modal>
        </div>
    )
}
