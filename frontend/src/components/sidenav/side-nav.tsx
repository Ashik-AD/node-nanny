import { RiHome5Line, RiHome5Fill } from 'react-icons/ri'
import { BsFolder, BsFolderFill } from 'react-icons/bs'
import { TbLayoutList, TbLayoutListFilled } from 'react-icons/tb'
import { HiOutlineTerminal, HiTerminal } from 'react-icons/hi'

import styles from './styles.module.css'
import { ReactNode } from 'react'
import { NavLink } from 'react-router-dom'
import Stack from '../commons/stack'

type NavItemProps = {
    iconDefault: ReactNode
    iconActive: ReactNode
    path?: string
    isActive?: boolean
    handleClick?: () => void
}

export default function SideNav() {
    return (
        <div className={styles.container}>
            <Stack placeCenter className={styles.content}>
                <NavItem
                    iconDefault={<RiHome5Line />}
                    iconActive={<RiHome5Fill />}
                    path="/"
                />
                <NavItem
                    iconDefault={<BsFolder />}
                    iconActive={<BsFolderFill />}
                />
                <NavItem
                    iconDefault={<TbLayoutList />}
                    iconActive={<TbLayoutListFilled />}
                />
                <div className={styles.nav__divider}></div>
                <NavItem
                    iconDefault={<HiOutlineTerminal />}
                    iconActive={<HiTerminal />}
                    path="/terminal"
                />
            </Stack>
        </div>
    )
}

function NavItem({
    iconDefault,
    iconActive,
    path,
    isActive = false,
    handleClick,
}: NavItemProps) {
    if (path) {
        return (
            <NavLink
                to={path}
                className={({ isActive }) =>
                    `${styles.nav__item} ${isActive ? styles['nav__item-active'] : ''}`
                }
            >
                {({ isActive }) => (
                    <span className={`${styles.nav__item__icon}`}>
                        {isActive ? iconActive : iconDefault}
                    </span>
                )}
            </NavLink>
        )
    }
    return (
        <button
            className={`${styles.nav__item} ${isActive ? styles['nav__item-active'] : ''}`}
            onClick={handleClick}
        >
            <span className={styles.nav__item__icon}>
                {isActive ? iconActive : iconDefault}
            </span>
        </button>
    )
}
