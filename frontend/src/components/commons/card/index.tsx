//@TODO:
// - add tooltip when refresh failed

import { ReactNode } from 'react'
import { Button } from '@nextui-org/react'
import Stack from '../stack'

import { GrRefresh } from 'react-icons/gr'

import styles from './styles.module.css'

type Props = {
    children: ReactNode
    title?: ReactNode
    refresh?: {
        onHandleRefresh: () => void
        status?: 'refreshing' | 'failed' | 'default'
    }
}
export default function Card({ children, title, refresh }: Props) {
    return (
        <Stack
            className={`${styles.container} ${refresh ? styles[`container_status-${refresh.status || 'default'}`] : null}`}
        >
            {(title || refresh) && (
                <Stack dir="column" justifyContent="spaceBetween">
                    <div>
                        {title && (
                            <Stack
                                dir="column"
                                gap="36px"
                                className="text-large font-semibold"
                            >
                                {title}
                            </Stack>
                        )}
                    </div>
                    {refresh && (
                        <Button
                            size="sm"
                            radius="full"
                            variant="light"
                            onClick={refresh.onHandleRefresh}
                            isIconOnly
                        >
                            <GrRefresh
                                className={styles.refresh__icon}
                                size={18}
                            />
                        </Button>
                    )}
                </Stack>
            )}
            <div className={styles.card__content}>{children}</div>
        </Stack>
    )
}
