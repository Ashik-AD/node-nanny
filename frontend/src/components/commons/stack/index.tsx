import { CSSProperties, ReactNode } from 'react'
import styles from './styles.module.css'

type Props = {
    children: ReactNode
    dir?: 'row' | 'column'
    className?: string
    style?: CSSProperties
    wrap?: boolean
    gap?: string | number
    placeCenter?: boolean
    justifyContent?: 'spaceBetween' | 'spaceEvenly' | 'spaceAround'
    alignItems?: AlignSetting
    isDistribute?: boolean
    isResponsive?: boolean
}

export default function Stack(props: Props) {
    const {
        children,
        dir = 'row',
        className,
        style,
        wrap,
        gap,
        placeCenter,
        alignItems,
        isResponsive = false,
        justifyContent,
        isDistribute,
    } = props

    let stylesName = `${className ? className : ''} ${wrap ? styles.stack_wrapp : ''}`
    let stackDir = dir == 'row' ? styles.stack_row : styles.stack_col
    let stackResponsive = isResponsive ? styles.stack_responsive : ''
    let contentCenter = placeCenter ? styles.stack_place_center : ''
    let distributeStyle = isDistribute ? styles.stack_distribute : ''
    let itemsCenter = alignItems ? styles[`align-${alignItems}`] : ''
    let contentSpacing = justifyContent ? styles[justifyContent] : ''

    return (
        <div
            className={`${styles.stack} ${stackDir} ${stylesName} ${contentCenter} ${itemsCenter} ${distributeStyle} ${stackResponsive} ${contentSpacing}`}
            style={{ gap: gap || 24, ...style }}
        >
            {children}
        </div>
    )
}
