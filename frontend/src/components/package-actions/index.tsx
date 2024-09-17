import { Button } from '@nextui-org/react'
import Stack from '../commons/stack'

import { MdUpdate, MdOutlineDownloading } from 'react-icons/md'
import { TbTrash } from 'react-icons/tb'

type Props = {
    id: string
    onHandleUpdate: (pkgId: string) => void
    onHandleRemove: (pkgId: string) => void
    onHandleCheckUpdate: (pkgId: string) => void
}

export default function PackageActions(props: Props) {
    const { id, onHandleCheckUpdate, onHandleUpdate, onHandleRemove } = props
    return (
        <Stack gap="16px" dir="column">
            <Button
                startContent={<MdUpdate size={20} />}
                radius="full"
                color="primary"
                title="Update package"
                size="sm"
                onClick={() => onHandleCheckUpdate(id)}
                isIconOnly
            />
            <Button
                startContent={<MdOutlineDownloading size={20} />}
                radius="full"
                color="secondary"
                title="Install package"
                size="sm"
                onClick={() => onHandleUpdate(id)}
                isIconOnly
            />
            <Button
                startContent={<TbTrash size={20} />}
                radius="full"
                color="warning"
                title="Remove package"
                size="sm"
                onClick={() => onHandleRemove(id)}
                isIconOnly
            />
        </Stack>
    )
}
