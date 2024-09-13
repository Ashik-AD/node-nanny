import { ReactNode, useEffect, useRef } from 'react'
import { createPortal } from 'react-dom'
import { Button } from '@nextui-org/react'
import Stack from '../stack'

import { IoClose } from 'react-icons/io5'

import styles from './styles.module.css'

type Props = {
    children: ReactNode
    isShow: boolean
    onClose: () => void
    onClickOutsideClose?: () => void
    title?: string
    subtitle?: string
}

let containerEle = document.querySelector('#modal') as HTMLDivElement

export default function Modal(props: Props) {
    const {
        children,
        onClose,
        onClickOutsideClose,
        title,
        subtitle,
        isShow = false,
    } = props

    let modalEle = useRef(document.createElement('div'))
    let modalContent = (
        <>
            <div className={styles.modal__box}>
                {(title || subtitle) && (
                    <Stack
                        gap="6px"
                        className={styles.modal__heading}
                        placeCenter
                    >
                        {title && (
                            <h2 className="text-xl font-medium">{title}</h2>
                        )}
                        {subtitle && (
                            <p className="text-medium font-medium">
                                {subtitle}
                            </p>
                        )}
                    </Stack>
                )}
                <div className={styles.modal__body}>{children}</div>
                <Button
                    className={styles.btn__modal__close}
                    startContent={
                        <IoClose size={20} color="var(--warning-base)" />
                    }
                    size="sm"
                    radius="full"
                    color="warning"
                    variant="flat"
                    onClick={onClose}
                    isIconOnly
                />
            </div>
            <div
                className={styles.modal__overlay}
                onClick={onClickOutsideClose}
            ></div>
        </>
    )

    useEffect(() => {
        if (isShow) {
            modalEle.current.setAttribute('class', styles.modal)
            containerEle.appendChild(modalEle.current)
        } else {
          if(modalEle.current.parentNode == containerEle) {
            containerEle.removeChild(modalEle.current)
          }
        }
    }, [isShow])

    return createPortal(modalContent, modalEle.current)
}
