export type LogItemStatus = 'idle' | 'success' | 'info' | 'warn' | 'error'

export type LogItem = {
    text: string
    timestamp?: string
    status: LogItemStatus
    hideStatus?: boolean
}

export type LogItemGroup = {
    masterLog: LogItem
    items: Array<LogItem>
}
