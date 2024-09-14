import Stack from '../stack'
import styles from './styles.module.css'

type Props = {
    title: string
    count: string | number
    unit?: string
}
export default function Stat(props: Props) {
    return (
        <Stack gap="18px" placeCenter>
            <span className={styles.stat__count}>
                {props.count}
                <small className='text-xlarge'>{props.unit}</small>
            </span>
            <span className="text-large font-bold">{props.title}</span>
        </Stack>
    )
}
