import { ChangeEvent } from 'react'
import Stack from '../../commons/stack'
import TextField from '../../commons/textField'

type Env = {
    key: string
    secret: string
}
type Props = {
    env: Env
    onChange: (value: Env) => void
}

export default function EnvAdd({ env, onChange }: Props) {
    function handleChangeEnv(eve: ChangeEvent<HTMLInputElement>) {
        let { name, value } = eve.target
        if (onChange) {
            onChange({ ...env, [name]: value })
        }
    }
    return (
        <Stack dir="column" isDistribute>
            <TextField
                name="key"
                value={env?.key}
                placeholder="Variable Name (key)"
                aria-label="variable name or key"
                onChange={handleChangeEnv}
                style={css}
            />
            <TextField
                name="secret"
                value={env?.secret}
                placeholder="Secret"
                aria-label="variable secret"
                onChange={handleChangeEnv}
                style={css}
            />
        </Stack>
    )
}

let css = {
    background: 'var(--background-primary)',
}
