<template>
    <div class="soup-picker">
        <v-container class="py-8">
            <v-row class="mb-6">
                <v-col cols="12" class="text-center">
                    <h1 class="text-h4 font-weight-bold primary--text mb-2">{{ t('appTitle') }}</h1>
                    <p class="text-subtitle-1 grey--text text--darken-1">
                        <a href="https://github.com/UoooBarry/turtle-soup" target="_blank" class="text-decoration-none">
                            <v-icon>mdi-github</v-icon>
                            Github
                        </a>
                        <LanguagePicker></LanguagePicker>
                    </p>
                </v-col>
            </v-row>

            <v-row>
                <v-col v-for="soup in soupStore.soups" :key="soup.id" cols="12" sm="6" md="4" lg="3" class="d-flex">
                    <v-card @click="selectSoup(soup)" class="flex-grow-1 d-flex flex-column" :hover="true" min-height="180">
                        <div class="primary lighten-2 pa-4">
                            <v-icon size="40" color="white">mdi-silverware-fork-knife</v-icon>
                        </div>

                        <v-card-title class="text-h6 font-weight-medium px-4 pt-4 pb-2"
                            style="white-space: normal; word-break: break-word;">
                            {{ soup.soup_question }}
                        </v-card-title>

                        <v-card-text class="px-4 pb-3 pt-0">
                            <div class="d-flex flex-wrap">
                                <v-chip v-for="tag in soup.tag" :key="tag" class="mr-2 mb-2" small color="primary"
                                    text-color="white" outlined>
                                    {{ tag }}
                                </v-chip>
                            </div>
                        </v-card-text>

                        <v-spacer></v-spacer>

                        <v-card-actions class="px-4 pb-3">
                            <v-btn color="primary" block depressed class="text-capitalize">
                                {{ t('gameStart') }}
                            </v-btn>
                        </v-card-actions>
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
import { useI18n } from 'vue-i18n'
import LanguagePicker from './LanguagePicker.vue';

const soupStore = useSoupStore()
const gameStore = useGameStore()
const { t } = useI18n()

const selectSoup = (soup: Soup) => {
    soupStore.setCurrentSoup(soup)
    gameStore.createGame(soup.id)
}
</script>

<style scoped>
.soup-picker {
    background-color: #f5f7fa;
    min-height: 100vh;
}

.v-card {
    transition: all 0.3s ease;
    border-radius: 12px;
    overflow: hidden;
    border: 1px solid rgba(0, 0, 0, 0.1);
}

.v-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1) !important;
    cursor: pointer;
}

.v-chip {
    transition: all 0.2s ease;
}

.v-chip:hover {
    transform: scale(1.05);
}

.primary.lighten-2 {
    height: 80px;
    display: flex;
    align-items: center;
}
</style>
