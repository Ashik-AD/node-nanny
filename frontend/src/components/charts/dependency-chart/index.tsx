import { ApexOptions } from 'apexcharts'
import Chart from 'react-apexcharts'
import Card from '../../commons/card'

type Props = {
    data: Array<number>
}

let options: ApexOptions = {
    chart: {
        type: 'pie',
    },
    labels: ['Dev. Dependencies', 'Dependencies'],
    colors: ['var(--warning-base)', 'var(--secondary-base)'],
    stroke: {
        show: false,
    },
    legend: {
        position: 'bottom',
        inverseOrder: true,
        fontSize: '14px',
        fontWeight: 600,
        itemMargin: {
            horizontal: 24,
        },
        markers: {
            offsetX: -6,
        },
    },
    plotOptions: {
        pie: {
            dataLabels: {
                offset: -30,
            },
        },
    },
}

export default function DependencyPieChart({ data }: Props) {
    return (
        <Card>
            <Chart
                type="pie"
                options={options}
                series={data}
                width={400}
                height={350}
            />
        </Card>
    )
}
