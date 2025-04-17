<template>
    <div class="soup-picker">
        <v-container>
            <v-row>
                <v-col v-for="soup in soupStore.soups" :key="soup.id" cols="12" sm="6" md="4">
                    <v-card @click="selectSoup(soup)">
                        <v-card-title style="white-space: normal; word-break: break-word;">
                            {{ soup.soup_question }}
                        </v-card-title>
                        <v-card-subtitle>
                            <v-chip v-for="tag in soup.tag" :key="tag" class="mr-2">
                                {{ tag }}
                            </v-chip>
                        </v-card-subtitle>
                    </v-card>
                </v-col>
            </v-row>
        </v-container>
    </div>
</template>

<script setup lang="ts">
import { useSoupStore } from '@/stores/soupStore'
import { useGameStore } from '@/stores/gameStore'
import type { Soup } from '@/types/soup';

const soupStore = useSoupStore()
const gameStore = useGameStore()

const selectSoup = (soup: Soup) => {
    soupStore.setCurrentSoup(soup)
    gameStore.createGame(soup.id)
}
</script>
