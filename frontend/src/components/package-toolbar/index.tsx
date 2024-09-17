import { Button } from '@nextui-org/react'
import TextField from '../commons/textField'
import Stack from '../commons/stack'

import { MdUpdate, MdOutlineDownloading } from 'react-icons/md'
import { TbTrash } from 'react-icons/tb'
import { FiSearch } from 'react-icons/fi'
import { ChangeEvent } from 'react'

type Props = {
    onHandleSearch: (event: ChangeEvent<HTMLInputElement>) => void
    onHandleUpdateCheck: () => void
    onHandleInstallPackage: () => void
    onHandleRemovePackage: () => void
}

export default function PackageToolbar(props: Props) {
    return (
        <Stack dir="column" alignItems="center" justifyContent="spaceBetween">
            <TextField
                placeholder="Search installed package"
                leadingIcon={<FiSearch />}
                onChange={props.onHandleSearch}
            />
            <Stack dir="column">
                <Button
                    startContent={<MdUpdate size={18} />}
                    color="primary"
                    onClick={props.onHandleUpdateCheck}
                >
                    Check for updates
                </Button>
                <Button
                    startContent={<MdOutlineDownloading size={18} />}
                    color="secondary"
                    onClick={props.onHandleInstallPackage}
                >
                    Re-install packages
                </Button>
                <Button
                    startContent={<TbTrash size={18} />}
                    color="warning"
                    onClick={props.onHandleRemovePackage}
                >
                    Remove packages
                </Button>
            </Stack>
        </Stack>
    )
}
