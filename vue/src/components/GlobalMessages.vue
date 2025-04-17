<template>
    <div class="global-messages">
        <transition-group name="message">
            <v-alert v-for="(message, index) in messages" :key="index" :class="['message']" :type="message.type"
                :closable="false" @click="removeMessage(index)" :text="message.content">
            </v-alert>
        </transition-group>
    </div>
</template>

<script setup lang="ts">
import { useMessageStore } from '@/stores/messageStore'
import { watch } from 'vue'

const messageStore = useMessageStore()
const { messages } = messageStore

// Auto-remove messages after their timeout
watch(messages, (newMessages) => {
    newMessages.forEach((message, index) => {
        if (message.timeout > 0) {
            setTimeout(() => {
                messageStore.messages.splice(index, 1)
            }, message.timeout)
        }
    })
}, { deep: true })

const removeMessage = (index: number) => {
    messageStore.messages.splice(index, 1)
}
</script>

<style scoped>
.global-messages {
    position: fixed;
    top: 5px;
    z-index: 1000;
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    gap: 10px;
    padding-right: 20px;
    text-align: center;
}

.message {
    width: 100vw;
    padding: 12px 16px;
    border-radius: 4px;
    color: white;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
    cursor: pointer;
    transition: all 0.3s ease;
}

.message.error {
    background-color: #ff4444;
}

.message.info {
    background-color: #33b5e5;
}

.message.success {
    background-color: #00c851;
}

.message-enter-active,
.message-leave-active {
    transition: all 0.5s ease;
}

.message-enter-from,
.message-leave-to {
    opacity: 0;
    transform: translateX(30px);
}
</style>

