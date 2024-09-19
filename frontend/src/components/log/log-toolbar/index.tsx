import { ChangeEvent, CSSProperties, useEffect, useState } from 'react'
import {
    Checkbox,
    CheckboxGroup,
    Dropdown,
    DropdownItem,
    DropdownMenu,
    DropdownTrigger,
} from '@nextui-org/react'
import Stack from '../../commons/stack'
import TextField from '../../commons/textField'
import { FiSearch } from 'react-icons/fi'

import { TiArrowUp } from 'react-icons/ti'

import styles from './styles.module.css'

let sortOptions = ['title', 'status', 'time']

type Props = {
    onSearch: (eve: ChangeEvent<HTMLInputElement>) => void
    onFilter: (filterKeys: string[]) => void
    onSort: (sortBy: string, dir: 'asc' | 'dsc') => void
}

export default function LogToolbar(props: Props) {
    let [sortTerm, setSortTerm] = useState('')
    let [sortDir, setSortDir] = useState<'asc' | 'dsc'>('dsc')

    useEffect(() => {
        if (props?.onSort) {
            props.onSort(sortTerm, sortDir)
        }
    }, [sortTerm, sortDir])

    function handleChangeFilter(keys: string[]) {
        props.onFilter(keys)
    }

    function handleChangeDropdown(value: string | number) {
        if (typeof value == 'string') {
            if (value == sortTerm) {
                let dir = sortDir == 'asc' ? 'dsc' : 'asc'
                setSortDir(dir as typeof sortDir)
            }
            setSortTerm(value)
        }
    }

    let sortArrowStyle =
        sortDir == 'asc' ? styles.sort__dir_asc : styles.sort__dir_dsc

    return (
        <Stack dir="column" alignItems="center" justifyContent="spaceBetween">
            <TextField
                name="search"
                placeholder="Search in logs"
                leadingIcon={<FiSearch />}
                onChange={props.onSearch}
            />
            <Stack dir="column">
                <CheckboxGroup
                    orientation="horizontal"
                    onChange={handleChangeFilter}
                >
                    <Checkbox value="info" color="primary">
                        <span className="text-primary">Info</span>
                    </Checkbox>
                    <Checkbox value="success" color="success">
                        <span className="text-success">Success</span>
                    </Checkbox>
                    <Checkbox value="warning" color="warning">
                        <span className="text-warning">Warning</span>
                    </Checkbox>
                    <Checkbox value="error" color="danger">
                        <span className="text-danger">Error</span>
                    </Checkbox>
                </CheckboxGroup>
                <Stack>
                    <Dropdown className={styles.dropdown}>
                        <Stack dir="column" gap={12}>
                            <span>Sort By:</span>
                            <DropdownTrigger>
                                <button
                                    className="btn_reset font-semibold"
                                    style={css}
                                >
                                    {sortTerm?.trim() || 'None'}
                                </button>
                            </DropdownTrigger>
                        </Stack>
                        <DropdownMenu
                            aria-label="Static Actions"
                            selectionMode="single"
                            selectedKeys={sortTerm}
                            onAction={handleChangeDropdown}
                        >
                            {sortOptions.map((opt) => (
                                <DropdownItem
                                    key={opt}
                                    className={styles.sort__item}
                                    id={
                                        opt == sortTerm
                                            ? styles['sort__item-active']
                                            : ''
                                    }
                                >
                                    {opt}
                                    {opt == sortTerm && (
                                        <TiArrowUp
                                            size={20}
                                            className={sortArrowStyle}
                                        />
                                    )}
                                </DropdownItem>
                            ))}
                        </DropdownMenu>
                    </Dropdown>
                </Stack>
            </Stack>
        </Stack>
    )
}

let css: CSSProperties = {
    textTransform: 'capitalize',
}
