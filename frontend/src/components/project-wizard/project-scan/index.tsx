import { useState } from 'react'
import { Button } from '@nextui-org/react'
import Stack from '../../commons/stack'

import {
    IoIosCheckmarkCircle,
    IoIosCloseCircle,
    IoIosArrowForward,
} from 'react-icons/io'
import { RiRefreshLine } from 'react-icons/ri'

import styles from './styles.module.css'
import Modal from '../../commons/modal'

type Tasks = {
    Name: string
    Status: 'completed' | 'pending' | 'failed' | 'scanning'
    Additional?: Array<string>
}

export default function ProjectScan() {
    let [isVisible, setIsVisible] = useState(false)

    let [tasks, _] = useState<Array<Tasks>>([
        {
            Name: 'Checking project',
            Status: 'completed',
            Additional: [
                'Found subproject `Frontend`',
                'Found subproject `Backend`',
            ],
        },
        {
            Name: 'Checking project dependencies',
            Status: 'failed',
        },
        {
            Name: 'Preparing project',
            Status: 'pending',
        },
    ])

    function handleToggleModal() {
        setIsVisible((prev) => !prev)
    }

    return (
        <div>
            <Button color="primary" onClick={handleToggleModal}>
                Open project scanner
            </Button>

            <Modal
                title="Please wait a moments"
                onClose={handleToggleModal}
                isShow={isVisible}
            >
                <Stack gap="36px" alignItems="start">
                    <Stack gap="8px">
                        {tasks.map((item, idx) => (
                            <ProjectTask
                                Name={item.Name}
                                Status={item.Status}
                                Additional={item.Additional}
                                key={idx}
                            />
                        ))}
                    </Stack>
                    <Button color="primary" radius="sm">
                        Cancel
                    </Button>
                </Stack>
            </Modal>
        </div>
    )
}

function ProjectTask(props: Tasks) {
    let [expand, setExpand] = useState(false)

    let status = props.Status
    let statusClassName =
        status == 'completed'
            ? styles['task_item-completed']
            : status == 'failed'
              ? styles[`task_item-failed`]
              : status == 'scanning'
                ? styles[`task_item-scanning`]
                : styles[`task_item-pending`]
    return (
        <Stack gap="8px" className={styles.task_item}>
            <Stack
                className={`${styles.task_item} ${statusClassName}`}
                dir="column"
                alignItems="center"
                justifyContent="spaceBetween"
            >
                <span tabIndex={0}>{props.Name}</span>
                <div className={styles.task__icons}>
                    {props.Additional ? (
                        <button onClick={() => setExpand(!expand)}>
                            <IoIosArrowForward
                                className={`${styles.icons__arrow} ${expand ? styles['icons__arrow-rotate'] : ''}`}
                            />
                        </button>
                    ) : null}
                    {props.Status != 'pending' ? (
                        <i>
                            {props.Status == 'scanning' ? (
                                <RiRefreshLine
                                    className={styles.animate_spin}
                                />
                            ) : props.Status == 'failed' ? (
                                <IoIosCloseCircle />
                            ) : (
                                <IoIosCheckmarkCircle />
                            )}
                        </i>
                    ) : null}
                </div>
            </Stack>
            {props.Additional && props.Additional.length > 0 && expand ? (
                <Stack gap="6px" className={styles.tasks__additional}>
                    {props.Additional.map((task, idx) => (
                        <span key={idx}>{task}</span>
                    ))}
                </Stack>
            ) : null}
        </Stack>
    )
}
