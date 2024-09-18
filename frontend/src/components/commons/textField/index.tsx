import { InputHTMLAttributes, ReactNode, useId, forwardRef } from 'react'
import styles from './styles.module.css'

type Props = InputHTMLAttributes<HTMLInputElement> & {
    type?: 'text' | 'password' | 'email' | 'number' | 'search'
    label?: string
    error?: string
    className?: string
    trailingIcon?: ReactNode
    leadingIcon?: ReactNode
}

const TextField = forwardRef<HTMLInputElement, Props>((props, ref) => {
    const {
        className,
        style,
        id,
        type = 'text',
        label,
        leadingIcon,
        trailingIcon,
        error,
        ...restProps
    } = props

    const cid = !id ? useId() : id
    let classes = `${styles.input} ${error ? styles.input_error : ''}`
    return (
        <div className={classes}>
            {label && <label htmlFor={id || cid}>{label}</label>}
            <div
                className={`${styles.input__box} ${className ? className : ''}`}
                style={style}
            >
                {leadingIcon && (
                    <span
                        className={`${styles.input__icon} ${styles.input__icon__leading}`}
                    >
                        {leadingIcon}
                    </span>
                )}
                <input {...restProps} id={id || cid} ref={ref} />
                {trailingIcon && (
                    <span
                        className={`${styles.input__icon} ${styles.input__icon__trailing}`}
                    >
                        {trailingIcon}
                    </span>
                )}
            </div>
            {error && <span className={styles.error__message}>{error}</span>}
        </div>
    )
})
export default TextField
