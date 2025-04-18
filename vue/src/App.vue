<template>
    <div class="app-container">
        <GlobalMessages />
        <SoupPicker v-if="gameStore.state === GameState.Picking" />
        <ChatInterface v-else-if="gameStore.state === GameState.Gaming" />
        <FinishGame v-else-if="gameStore.state === GameState.Finished" />
        <footer class="app-footer">
            <p>© 2025 <a href="https://github.com/UoooBarry" class="text-decoration-none">Uooobarry</a>. All rights reserved.
                This website is available for myself only. No public <LoginDialog />. {{ authStore.user?.username }}
            </p>
        </footer>
    </div>
</template>

<script setup lang="ts">
import ChatInterface from './components/ChatInterface.vue'
import GlobalMessages from './components/GlobalMessages.vue'
import SoupPicker from './components/SoupPicker.vue'
import FinishGame from './components/FinishGame.vue'
import LoginDialog from './components/LoginDialog.vue'

import { useMessageStore } from '@/stores/messageStore'
import { useSoupStore } from '@/stores/soupStore'
import { useGameStore } from '@/stores/gameStore'
import { MessageType } from '@/types/globalMessage'
import { GameState } from '@/stores/gameStore'
import { onMounted } from 'vue'
import { useAuthStore } from '@/stores/authStore'

const messageStore = useMessageStore()
const authStore = useAuthStore()
const soupStore = useSoupStore()
const gameStore = useGameStore()

onMounted(async () => {
    messageStore.pushMessage('Welcome!', MessageType.Success)
    authStore.restoreLogin()
    soupStore.fetchSoups()
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

.app-container {
    min-height: 100vh;
    position: relative;
}

.app-footer {
    background-color: #f5f5f5;
    padding: 10px 20px;
    text-align: center;
    border-top: 1px solid #e0e0e0;
    margin-top: auto;
}

.app-footer p {
    margin: 0;
    color: #666;
    font-size: 0.9rem;
}
</style>
