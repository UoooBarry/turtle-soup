<template>
    <v-card class="chat-container" flat>
        <v-toolbar color="primary" density="compact">
            <v-toolbar-title class="text-white">
                海龟汤Bot
                <template v-if="preparingLLM">
                    [正在为你准备AI Bot助手...]
                </template>
            </v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn icon="mdi-close" variant="text" color="white" :disabled="preparingLLM" @click="endGame"></v-btn>
        </v-toolbar>

        <v-card-text class="messages-container pa-0">
            <div class="d-flex flex-column" style="min-height: 100%;">
                <!-- System message -->
                <div class="system-message pa-4 text-center">
                    <v-chip color="primary" variant="outlined" size="small">
                        <v-icon start>mdi-lightbulb-question-outline</v-icon>
                        汤面
                    </v-chip>
                    <div class="mt-2 text-body-1">{{ soupStore.currentSoup.soup_question }}</div>
                    <div class="mt-2">
                        <v-chip v-for="(tag, index) in soupStore.currentSoup.tag" class="mr-1" :key="index" color="primary"
                            variant="outlined" size="small">
                            {{ tag }}
                        </v-chip>
                    </div>
                </div>

                <!-- Chat messages -->
                <div v-for="(message, index) in messages" :key="index" class="message-wrapper"
                    :class="message.sender === 'user' ? 'user-wrapper' : 'bot-wrapper'">
                    <div class="message-content" :class="message.sender === 'user' ? 'user-message' : 'bot-message'">
                        <div v-if="message.sender !== 'user'" class="message-sender">
                            <v-avatar size="24" color="primary" class="mr-2">
                                <v-icon size="14" color="white">mdi-robot-outline</v-icon>
                            </v-avatar>
                            <span class="font-weight-medium">Bot</span>
                        </div>
                        <div v-else class="message-sender">
                            <v-avatar size="24" color="secondary" class="mr-2">
                                <v-icon size="14" color="white">mdi-account</v-icon>
                            </v-avatar>
                            <span class="font-weight-medium">Me</span>
                        </div>

                        <div class="message-text">
                            <div class="text-body-1" style="white-space: pre-wrap;">{{ message.text }}</div>
                        </div>

                        <div class="message-time text-caption text-medium-emphasis">
                            {{ formatTime(message.timestamp) }}
                        </div>
                    </div>
                </div>

                <div v-if="waitingForResponse && !preparingLLM" class="bot-wrapper">
                    <div class="bot-message message-content">
                        <div class="message-sender">
                            <v-avatar size="24" color="primary" class="mr-2">
                                <v-icon size="14" color="white">mdi-robot-outline</v-icon>
                            </v-avatar>
                            <span class="font-weight-medium">Bot</span>
                        </div>
                        <div class="message-text">
                            <v-progress-circular indeterminate size="16" width="2"></v-progress-circular>
                        </div>
                    </div>
                </div>
            </div>
        </v-card-text>

        <v-divider></v-divider>

        <v-card-actions class="pa-3 input-area">
            <v-row no-gutters>
                <v-col cols="12" class="pa-0">
                    <v-checkbox v-model="needHint" density="compact" color="primary" class="mt-0 pt-0">
                        <template v-slot:label>
                            <v-icon>mdi-lightbulb-on-outline</v-icon>
                            <span class="text-caption">提示</span>
                        </template>
                    </v-checkbox>

                    <v-textarea v-model="newMessage" label="输入你的推理或提问..." variant="outlined" density="comfortable" hide-details
                        rounded single-line autofocus @keyup.enter="sendMessage" :rows="3" :disabled="waitingForResponse" class="mt-0 pt-0">
                        <template v-slot:append-inner>
                            <div class="d-flex align-center" style="height: 100%;">
                                <v-btn icon="mdi-send" color="primary" variant="tonal" @click="sendMessage"
                                    :disabled="!newMessage.trim() || waitingForResponse">
                                </v-btn>
                            </div>
                        </template>
                    </v-textarea>
                </v-col>
            </v-row>
        </v-card-actions>
    </v-card>
</template>

<script setup lang="ts">
import { ref, nextTick } from 'vue';
import { useSoupStore } from '@/stores/soupStore';
import { useGameStore } from '@/stores/gameStore'
import { onMounted } from 'vue';
import { useMessageStore } from '@/stores/messageStore';
import { MessageType } from '@/types/globalMessage';
import { computed } from 'vue';

interface Message {
    text: string;
    sender: 'user' | 'bot' | 'system';
    timestamp: Date;
};

const soupStore = useSoupStore();
const gameStore = useGameStore();
const globalMsgStore = useMessageStore()
const messages = ref<Message[]>([]);
const newMessage = ref('');
const waitingForResponse = ref(true);
const needHint = ref(false);

const formatTime = (date: Date) => {
    return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
};

const preparingLLM = computed(() => {
    return waitingForResponse.value && messages.value.length <= 0
})

const scrollToBottom = () => {
    const container = document.querySelector('.messages-container');
    if (container) {
        container.scrollTop = container.scrollHeight;
    }
};

const sendMessage = async () => {
    if (newMessage.value.trim() === '' || waitingForResponse.value) return;

    const userMessage = {
        text: newMessage.value,
        sender: 'user' as const,
        timestamp: new Date(),
    };
    messages.value.push(userMessage);
    newMessage.value = '';

    waitingForResponse.value = true;
    await nextTick();
    scrollToBottom();
    const res = await gameStore.askGame(userMessage.text, needHint.value);

    messages.value.push({
        text: res.answer,
        sender: 'bot' as const,
        timestamp: new Date(),
    });
    if (res.hint && res.hint.length > 0) {
        messages.value.push({
            text: res.hint,
            sender: 'bot' as const,
            timestamp: new Date(),
        });
    }

    await nextTick();
    scrollToBottom();
    waitingForResponse.value = false;
};

const endGame = async () => {
    await gameStore.endGame()
    globalMsgStore.pushMessage("游戏退出", MessageType.Success, 1000)
}

onMounted(async () => {
    await gameStore.startGame()
    globalMsgStore.pushMessage("游戏开始", MessageType.Success, 3000)
    waitingForResponse.value = false
})
</script>

<style scoped>
.chat-container {
    display: flex;
    flex-direction: column;
    height: 100vh;
    margin: 0 auto;
    border-radius: 4px;
    overflow: hidden;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.messages-container {
    flex: 1;
    overflow-y: auto;
    padding-top: 0;
    background-color: #f9f9f9;
    scroll-behavior: smooth;
    max-height: calc(100vh - 112px);
    /* Account for toolbar and input area */
}

.system-message {
    background-color: #f0f7ff;
    border-bottom: 1px solid #e0e0e0;
    position: sticky;
    top: 0;
    z-index: 1;
}

.message-wrapper {
    padding: 8px 16px;
}

.user-wrapper {
    display: flex;
    justify-content: flex-end;
}

.bot-wrapper {
    display: flex;
    justify-content: flex-start;
}

.message-content {
    max-width: 85%;
    padding: 12px 16px;
    border-radius: 12px;
    margin: 4px 0;
    position: relative;
}

.bot-message {
    background-color: white;
    border: 1px solid #e0e0e0;
    border-radius: 0 12px 12px 12px;
}

.user-message {
    background-color: #e3f2fd;
    border-radius: 12px 0 12px 12px;
}

.message-sender {
    display: flex;
    align-items: center;
    margin-bottom: 4px;
    font-size: 0.875rem;
}

.message-text {
    line-height: 1.5;
    word-break: break-word;
}

.message-time {
    text-align: right;
    margin-top: 4px;
    font-size: 0.75rem;
}

.input-area {
    background-color: white;
    padding: 12px 16px;
}

/* Custom scrollbar */
.messages-container::-webkit-scrollbar {
    width: 6px;
}

.messages-container::-webkit-scrollbar-track {
    background: #f1f1f1;
}

.messages-container::-webkit-scrollbar-thumb {
    background: #c1c1c1;
    border-radius: 3px;
}

.messages-container::-webkit-scrollbar-thumb:hover {
    background: #a8a8a8;
}
</style>
