import { Link } from 'react-router-dom'
import Stack from '../commons/stack'
import { HiMiniAdjustmentsVertical } from 'react-icons/hi2'
import { PiMarkdownLogoFill } from 'react-icons/pi'
import { IoCubeSharp } from 'react-icons/io5'
import { TbBrandNpm } from 'react-icons/tb'

import styles from './styles.module.css'
import { useState } from 'react'

type Project = {
    id: string
    name: string
    path: string
    files: { type: string; name: string }[]
    isChild?: boolean
}

type ProjectStructure = Project & {
    childProjects?: Project[]
}

export default function ProjectStructure(props: ProjectStructure) {
    const [isExtended, setIsExtended] = useState(false)

    let nodeClassName = `${styles.struct__node} text-small`
    return (
        <Stack
            gap="0"
            className={`${styles.project_structure} ${isExtended ? styles['project_structure-extend'] : styles['project_structure-unextend']}`}
        >
            <Link
                to={`#`}
                className={`text-medium ${styles.struct__root} ${!props.isChild ? styles.project__root : styles.sub__project__root}`}
                onClick={() => setIsExtended((prev) => !prev)}
            >
                <i className={styles.root__icon}>
                    <IoCubeSharp />
                </i>
                <span>{props.name}</span>
            </Link>
            <Stack className={styles.project__tree} gap="0px">
                {props.childProjects?.map((project) => (
                    <Stack
                        key={project.id}
                        className={styles.project__struct__sub}
                    >
                        <ProjectStructure {...project} isChild />
                    </Stack>
                ))}
                <Stack
                    gap="0"
                    className={
                        !props.isChild
                            ? styles.root__files
                            : styles.child__files
                    }
                >
                    <Link to={`./env`} className={nodeClassName}>
                        <i className={styles.file__icon}>
                            <HiMiniAdjustmentsVertical />
                        </i>
                        <span>{props.files[0].name}</span>
                    </Link>
                    <Link to={`/packages`} className={nodeClassName}>
                        <i className={styles.file__icon}>
                            <TbBrandNpm />
                        </i>
                        <h5>{props.files[1].name}</h5>
                    </Link>
                    <Link to={`/readme`} className={styles.struct__node}>
                        <i className={styles.file__icon}>
                            <PiMarkdownLogoFill />
                        </i>
                        <h5>{props.files[2].name}</h5>
                    </Link>
                </Stack>
            </Stack>
        </Stack>
    )
}
