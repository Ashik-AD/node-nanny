//@TODO:
// - fetch disk usage
// - implement refresh behaviour
// - fix bar tooltip or disable it

import ApexChart from 'react-apexcharts'
import { Chip } from '@nextui-org/react'
import { ApexOptions } from 'apexcharts'
import Card from '../../commons/card'
import Stack from '../../commons/stack'

let options: ApexOptions = {
    chart: {
        width: '100%',
        height: 421,
        type: 'bar',
        foreColor: 'var(--default-400)',
        toolbar: {
            show: false,
        },
    },
    labels: [
        'Dependencies',
        'Dev. dependencies',
        'Source code',
        'Images',
        'Audio/Vidoe',
        'Others',
    ],
    colors: [
        'var(--secondary-500)',
        'var(--warning-500)',
        'var(--success-base)',
        'var(--danger-base)',
        'var(--primary-base)',
        'var(--default-700)',
    ],
    legend: {
        show: false,
    },
    xaxis: {
        labels: {
            style: {
                fontSize: '13px',
                fontWeight: 600,
            },
        },
        axisTicks: {
            show: false,
        },
    },
    plotOptions: {
        bar: {
            distributed: true,
            columnWidth: 80,
        },
    },
    dataLabels: {
        formatter(val) {
            return `${val}MB`
        },
        style: {
            fontSize: '14px',
        },
    },
    grid: {
        borderColor: 'var(--secondary-600)',
    },
    tooltip: {
        theme: 'dark',
        style: {
            fontSize: '18px',
        },
        fillSeriesColor: true,
    },
}

export default function DiskUsage() {
    let data = [42, 18, 10, 30, 8, 4]
    let totalUsage = data.reduce((acc, cur) => cur + acc, 0)

    return (
        <Card
            title={
                <>
                    <span>Disk Usage</span>
                    <Chip color="secondary">Master</Chip>
                </>
            }
            refresh={{
                onHandleRefresh: () => console.log('implement refresh'),
            }}
        >
            <Stack gap="1rem">
                <ApexChart
                    options={options}
                    type="bar"
                    width={'100%'}
                    height={421}
                    series={[{ data }]}
                />
                <span>
                    Total usage: <strong>{totalUsage}MB</strong>
                </span>
            </Stack>
        </Card>
    )
}
