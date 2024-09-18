import Stack from '../../commons/stack'
import LogItem from '../log-item'
import type { LogItemGroup } from '../type'

import styles from './styles.module.css'

export default function LogItemGroup(props: LogItemGroup) {
    return (
        <Stack>
            <LogItem {...props.masterLog} />
            <Stack gap={'12px'}>
                {props.items.map((log, idx) => (
                    <div key={idx} className={styles.item}>
                        <LogItem {...log} />
                    </div>
                ))}
            </Stack>
        </Stack>
    )
}
