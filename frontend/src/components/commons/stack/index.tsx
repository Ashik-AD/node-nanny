import { CSSProperties, ReactNode } from 'react'
import styles from './styles.module.css'

type Props = {
    children: ReactNode
    dir?: 'row' | 'column'
    className?: string
    style?: CSSProperties
    wrap?: boolean
    gap?: string
    placeCenter?: boolean
    contentAlign?: 'spaceBetween' | 'spaceEvenly' | 'spaceAround'
    isResponsive?: boolean
}
export default function Stack({
    children,
    dir = 'row',
    className,
    style,
    wrap,
    gap,
    placeCenter,
    isResponsive = false,
    contentAlign,
}: Props) {
    const stylesName = `${styles.stack} ${className ? className : ''} ${dir == 'column' ? styles.stack_column : styles.stack_row} ${wrap ? styles.stack_wrapp : ''} ${placeCenter ? styles.stack_place_center : ''} ${isResponsive ? styles.stack_responsive : ''} ${contentAlign ? styles[contentAlign] : ''}`
    return (
        <div className={stylesName} style={{ gap: gap || 24, ...style }}>
            {children}
        </div>
    )
}
