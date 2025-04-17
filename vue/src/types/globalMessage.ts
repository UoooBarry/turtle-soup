export interface GlobalMessage {
  type: MessageType,
  content: string,
  timeout: number // ms
}

export enum MessageType {
  Error = 'error',
  Info = 'info',
  Success = 'success'
}
