import { Chip } from '@nextui-org/react'
import styles from './styles.module.css'
import type { LogItem, LogItemStatus } from '../type'

const colorMap: Record<LogItemStatus, string> = {
    info: 'primary',
    error: 'danger',
    idle: 'default',
    warn: 'warning',
    success: 'success',
}

export default function LogItem(props: LogItem) {
    let status = props.status
    return (
        <div className={styles.log_item}>
            <span
                className={`text-${colorMap[status] || status} text-small font-semibold`}
            >
                {props.timestamp}
            </span>
            <Chip
                size="sm"
                color={colorMap[status] as any}
                className={styles.log__status}
            >
                {status == 'info'
                    ? 'info'
                    : status == 'error'
                      ? 'error'
                      : colorMap[status]}
            </Chip>
            <span className="text-medium font-semibold">{props.text}</span>
        </div>
    )
}
