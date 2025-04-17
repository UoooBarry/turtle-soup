import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Ref } from 'vue'
import { MessageType, type GlobalMessage } from '@/types/globalMessage'

export const useMessageStore = defineStore('counter', () => {
  const messages: Ref<GlobalMessage[]> = ref([])
  const pushMessage = (content: string, type: MessageType = MessageType.Info, timeout: number = 1000) => {
    messages.value.push({ content: content, type: type, timeout: timeout })
  }

  return { messages, pushMessage }
})
