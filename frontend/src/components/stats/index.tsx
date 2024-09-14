//@TODO:
// - implement refresh function

import { Button } from '@nextui-org/react'
import { useState } from 'react'
import Card from '../commons/card'
import Stack from '../commons/stack'
import Stat from '../commons/stat'

type RefreshStatus = 'refreshing' | 'failed' | 'default'

export default function Stats() {
    let [status, setStatus] = useState<RefreshStatus>('default')

    return (
        <Card
            title={
                <>
                    <h2>Modules & Storage</h2>
                    <Button size="sm" color="secondary">
                        Master
                    </Button>
                </>
            }
            refresh={{
                status: status,
                onHandleRefresh: () => {
                    console.log('Implement refresh behaviour')
                    setStatus('refreshing')
                },
            }}
        >
            <Stack dir="column" justifyContent="spaceBetween">
                <Stat title="Dependencies" count="12" />
                <Stat title="Dev.Dependencies" count="6" />
                <Stat title="Node Modules" count="678" />
                <Stat title="Storage" count="210" unit="MB" />
            </Stack>
        </Card>
    )
}
