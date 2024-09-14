//@TODO
//handle proper routing urls
import { Avatar } from '@nextui-org/react'
import Stack from '../commons/stack'
import { Link, NavLink } from 'react-router-dom'

import {
    TbBrandNpm,
    TbLayoutDashboard,
    TbTerminal2,
    TbFileTime,
} from 'react-icons/tb'
import { HiMiniAdjustmentsVertical } from 'react-icons/hi2'

import styles from './styles.module.css'

const routes = [
    {
        title: 'Overview',
        path: '/',
        icon: <TbLayoutDashboard />,
    },
    {
        title: 'Packages',
        path: '/packages',
        icon: <TbBrandNpm color="var(--danger-400)" />,
    },
    {
        title: 'Env. Variables',
        path: '/envs',
        icon: <HiMiniAdjustmentsVertical color="var(--warning-400)" />,
    },
    {
        title: 'Console',
        path: '/console',
        icon: <TbTerminal2 />,
    },
    {
        title: 'Logs',
        path: '/log',
        icon: <TbFileTime />,
    },
]
export default function ProjectTopNav() {
    return (
        <Stack dir="column" justifyContent="spaceBetween">
            <Stack
                dir="column"
                alignItems="center"
                gap="12px"
                className={styles.project__heading}
            >
                <Avatar
                    size="sm"
                    src="https://uxwing.com/wp-content/themes/uxwing/download/brands-and-social-media/react-js-icon.png"
                />
                <Stack gap="8px" dir="column">
                    <Link to={'#'} className="text-medium font-semibold">
                        Node-Nanny
                    </Link>
                    <Link to={'#'} className="">
                        / Frontend
                    </Link>
                </Stack>
            </Stack>
            <Stack dir="column">
                {routes.map((route) => (
                    <NavLink
                        key={route.path}
                        className={({ isActive }) =>
                            `${styles.nav__item} ${isActive ? styles['nav__item-active'] : ''}`
                        }
                        to={route.path}
                    >
                        <i className={styles.nav__icon}>{route.icon}</i>
                        <span className="text-medium">{route.title}</span>
                    </NavLink>
                ))}
            </Stack>
        </Stack>
    )
}
