//@TODO
// - fetch warning, danger, and outdated count
// - implement refresh function

import { Chip } from '@nextui-org/react'
import Card from '../commons/card'
import Stack from '../commons/stack'
import Stat from '../commons/stat'

export default function PackageStatus() {
    return (
        <Card
            title={
                <>
                    <span>Package Status</span>
                    <Chip color="secondary">Master</Chip>
                </>
            }
            refresh={{
                onHandleRefresh: () => console.log('implement refresh'),
            }}
        >
            <Stack dir="column" justifyContent="spaceBetween">
                <Stat title="Warning" count={10} />
                <Stat title="In Danger" count={5} />
                <Stat title="Outdated" count={3} />
            </Stack>
        </Card>
    )
}
