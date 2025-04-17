<template>
    <div class="finish-game-container">
        <div class="congrats-message">
            <h1>Congratulations! 🎉</h1>
            <p>你完成了本轮海龟汤，你是大侦探!</p>
            <p><b>{{ gameStore.gameAnswer }}</b></p>
        </div>
        <v-btn
            color="primary"
            @click="returnToSelection"
            class="return-button"
        >
            回到海龟汤选择
        </v-btn>
    </div>
</template>

<script setup lang="ts">
import { useGameStore } from '@/stores/gameStore';

const gameStore = useGameStore();

const returnToSelection = () => {
    gameStore.endGame();
};
</script>

<style scoped>
.finish-game-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
    text-align: center;
    padding: 20px;
}

.congrats-message {
    margin-bottom: 20px;
}

.congrats-message h1 {
    font-size: 2.5rem;
    color: #42b883;
}

.congrats-message p {
    font-size: 1.2rem;
    color: #333;
}

.return-button {
    margin-top: 20px;
}
</style>

