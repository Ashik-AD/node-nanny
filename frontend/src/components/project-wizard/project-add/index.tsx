import { SyntheticEvent } from 'react'
import { Button } from '@nextui-org/react'
import Stack from '../../commons/stack'

import { OpenProjectSelect } from '../../../../wailsjs/go/main/App'
import { MdOutlineDriveFolderUpload } from 'react-icons/md'

import styles from './styles.module.css'

export default function ProjectAdd() {
    async function handleOpenFolderSelector() {
      console.log(window.navigator)
      await OpenProjectSelect()
    }

    function handleDrop(eve: SyntheticEvent<HTMLDivElement, DragEvent>) {
        eve.preventDefault()
        //@ts-ignore-next
        console.dir([...eve?.dataTrasfer?.items])
    }

    function handleDragOver(eve: SyntheticEvent<HTMLDivElement, DragEvent>) {
        eve.preventDefault()
    }

    return (
        <div onDrop={handleDrop} onDragOver={handleDragOver}>
            <Stack className={styles.container} placeCenter>
                <MdOutlineDriveFolderUpload size={48} />
                <span className="text-xl font-medium">
                    Drop a project folder here
                </span>
                <strong className="text-tiny font-bold">OR</strong>
                <label>
                    <Button
                        color="secondary"
                        radius="sm"
                        size="md"
                        onClick={handleOpenFolderSelector}
                    >
                        Browse folder
                    </Button>
                    <input
                        type="file"
                        name="new_project"
                        className={styles.directory_picker}
                        //@ts-ignore-next
                        webkitdirectory="true"
                    />
                </label>
            </Stack>
        </div>
    )
}
