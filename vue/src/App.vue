<template>
    <div>
        <GlobalMessages />
        <SoupPicker v-if="gameStore.state === GameState.Picking" />
        <ChatInterface v-else-if="gameStore.state === GameState.Gaming" />
        <FinishGame v-else-if="gameStore.state === GameState.Finished" />
    </div>
</template>

<script setup lang="ts">
import ChatInterface from './components/ChatInterface.vue'
import GlobalMessages from './components/GlobalMessages.vue'
import SoupPicker from './components/SoupPicker.vue'
import FinishGame from './components/FinishGame.vue'

import { useMessageStore } from '@/stores/messageStore'
import { useSoupStore } from '@/stores/soupStore'
import { useGameStore } from '@/stores/gameStore'
import { MessageType } from '@/types/globalMessage'
import { GameState } from '@/stores/gameStore'
import { onMounted } from 'vue'

const messageStore = useMessageStore()
const soupStore = useSoupStore()
const gameStore = useGameStore()

onMounted(async () => {
    messageStore.pushMessage('Welcome!', MessageType.Success)
    await soupStore.fetchSoups()
})
</script>

<style scoped>
.logo {
    height: 6em;
    padding: 1.5em;
    will-change: filter;
    transition: filter 300ms;
}

.logo:hover {
    filter: drop-shadow(0 0 2em #646cffaa);
}

.logo.vue:hover {
    filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
