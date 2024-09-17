//@TODO:
// - notify changes when env. values is modified

import { ChangeEvent, useRef, useState } from 'react'
import {
    Dropdown,
    DropdownTrigger,
    DropdownMenu,
    DropdownItem,
    Button,
} from '@nextui-org/react'
import Stack from '../../commons/stack'
import Modal from '../../commons/modal'
import TextField from '../../commons/textField'
import { HiEye, HiOutlineEyeSlash } from 'react-icons/hi2'
import { IoEllipsisHorizontal } from 'react-icons/io5'
import { TbEdit, TbTrashX } from 'react-icons/tb'

import styles from './styles.module.css'

type Env = {
    key: string
    secret: string
}

type Props = {
    env: Env
    onHandleEditEnv: (newEnv: Env, oldEnv: Env) => void
    onHandleRemoveEnv: (key: string) => void
}

export default function EnvItem(props: Props) {
    const { env, onHandleEditEnv, onHandleRemoveEnv } = props

    let [newEnv, setNewEnv] = useState<Env>({
        key: env.key,
        secret: env.secret,
    })
    let [errors, setErrors] = useState<Record<keyof Env, string>>({
        key: '',
        secret: '',
    })
    let [showSecret, setShowSecret] = useState(false)
    let [showModal, setShowModal] = useState(false)

    let keyRef = useRef<HTMLInputElement | null>(null)

    function handleChangeEnvInputs(eve: ChangeEvent<HTMLInputElement>) {
        let { name, value } = eve.target
        setNewEnv((prevEnvs) => ({ ...prevEnvs, [name]: value }))

        if (!value?.toString()?.trim()) {
            setErrors((prevErr) => ({
                ...prevErr,
                [name]: `${name} is missing`,
            }))
        } else {
            setErrors((prevErr) => ({
                ...prevErr,
                [name]: '',
            }))
        }
    }

    function toggleModal() {
        let state = !showModal
        setShowModal(state)
        if (state) {
            setTimeout(() => {
                keyRef.current?.focus()
            })
        }
    }

    function handleUpdateEnv() {
        let { key, secret } = newEnv
        if (!key?.trim()) {
            setErrors((prevErr) => ({
                ...prevErr,
                key: 'key is missing',
            }))
            return
        }
        if (!secret?.trim()) {
            setErrors((prevErr) => ({
                ...prevErr,
                secret: 'secret is missing',
            }))
            return
        }

        if (onHandleEditEnv) {
            onHandleEditEnv(
                {
                    key,
                    secret,
                },
                {
                    ...env,
                }
            )
            toggleModal()
        }
    }

    function handleDelete() {
        if (onHandleRemoveEnv) {
            onHandleRemoveEnv(env.key)
        }
    }
    return (
        <>
            <Stack dir="column" justifyContent="spaceBetween">
                <span
                    className="text-medium font-semibold"
                    aria-multiline={'false'}
                >
                    {env.key}
                </span>
                <Stack dir="column" alignItems="center">
                    <span className="text-medium font-semibold">
                        {showSecret ? (
                            env.secret
                        ) : (
                            <span className={styles.secret__encrypt}>
                                **********************************
                            </span>
                        )}
                    </span>
                    <button
                        className="btn_reset"
                        onClick={() => setShowSecret((prevState) => !prevState)}
                    >
                        {!showSecret ? <HiOutlineEyeSlash /> : <HiEye />}
                    </button>
                    <Dropdown className={styles.dropdown}>
                        <DropdownTrigger>
                            <button className="btn_reset">
                                <IoEllipsisHorizontal size={18} />
                            </button>
                        </DropdownTrigger>
                        <DropdownMenu aria-label="Env. varible actions">
                            <DropdownItem
                                key="edi"
                                startContent={<TbEdit size={24} />}
                                onClick={() => setShowModal(true)}
                            >
                                Edit
                            </DropdownItem>
                            <DropdownItem
                                key="delete"
                                startContent={<TbTrashX size={24} />}
                                onClick={handleDelete}
                            >
                                Delete
                            </DropdownItem>
                        </DropdownMenu>
                    </Dropdown>
                </Stack>
            </Stack>
            <Modal
                title="Edit environmental variable"
                onClose={toggleModal}
                onClickOutsideClose={toggleModal}
                isShow={showModal}
            >
                <Stack>
                    <Stack dir="column" gap="16px">
                        <TextField
                            type="text"
                            name="key"
                            label="KEY"
                            defaultValue={env.key}
                            ref={keyRef}
                            onChange={handleChangeEnvInputs}
                            error={errors.key}
                        />
                        <TextField
                            type="text"
                            name="secret"
                            label="SECRET"
                            value={newEnv.secret}
                            onChange={handleChangeEnvInputs}
                            error={errors.secret}
                        />
                    </Stack>
                    <Stack dir="column">
                        <Button
                            size="lg"
                            radius="sm"
                            color="default"
                            onClick={toggleModal}
                        >
                            Cancel
                        </Button>
                        <Button
                            size="lg"
                            radius="sm"
                            color="success"
                            onClick={handleUpdateEnv}
                        >
                            Save
                        </Button>
                    </Stack>
                </Stack>
            </Modal>
        </>
    )
}
