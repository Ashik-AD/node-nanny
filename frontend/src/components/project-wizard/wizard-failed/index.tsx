import { useState } from 'react'
import { Button } from '@nextui-org/react'
import Stack from '../../commons/stack'

import { IoIosCloseCircle } from 'react-icons/io'

import Modal from '../../commons/modal'

export default function WizardFailed() {
    let [isVisible, setIsVisible] = useState(true)

    let [reason, _] = useState<string>(
        'The directory dont contains any node/npm project'
    )

    function handleToggleModal() {
        setIsVisible((prev) => !prev)
    }

    return (
        <div>
            <Button color="primary" onClick={handleToggleModal}>
                Click here! 
            </Button>

            <Modal onClose={handleToggleModal} isShow={isVisible}>
                <Stack gap="36px" alignItems="start">
                    <Stack gap="8px" alignItems="center">
                        <IoIosCloseCircle
                            size={64}
                            color="var(--danger-base)"
                        />
                        <span className="text-xlarge">
                            Failed to add project
                        </span>
                    </Stack>
                    <Stack gap="12px">
                        <span className="text-medium font-semibold">
                            Reason:
                        </span>
                        <span className="text-small font-medium">{reason}</span>
                    </Stack>
                    <Stack dir="column" placeCenter>
                        <Button color="default" radius="sm" onClick={handleToggleModal}>
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
