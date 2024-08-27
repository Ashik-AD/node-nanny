/*
 * @TODO
 * `- show proper date for `last update`
 *  - implement project update check 
 * */
import { Avatar } from '@nextui-org/react'
import { HiMiniCheckCircle } from 'react-icons/hi2'
import { SlRefresh } from 'react-icons/sl'
import { Link } from 'react-router-dom'

import Stack from '../commons/stack'
import styles from './styles.module.css'

type Props = {
    id: string
    name: string
    description: string
    logoSrc?: string
    lastUpdated: Date
}

export default function ProjectCard({
    id,
    name,
    description,
    logoSrc,
    lastUpdated,
}: Props) {
    return (
        <Link to={`/projects/${id}`}>
            <article className={styles.container}>
                <Stack gap="14px" className={styles.content}>
                    <Stack dir="column" gap="12px">
                        <Avatar
                            size="sm"
                            src={logoSrc}
                            radius="full"
                            name="noddy"
                        />
                        <span className="text-large font-medium default-300">
                            {name}
                        </span>
                    </Stack>
                    <p className="text-small font-medium">{description}</p>
                    <Stack dir="column">
                        <Stack dir="column" alignItems="center" gap="8px">
                            <span className="text-success">
                                <HiMiniCheckCircle size={18} />
                            </span>
                            <span>
                                Last updated: {lastUpdated.toLocaleDateString()}
                            </span>
                        </Stack>
                        <button>
                            <SlRefresh size={18} />
                        </button>
                    </Stack>
                </Stack>
            </article>
        </Link>
    )
}
