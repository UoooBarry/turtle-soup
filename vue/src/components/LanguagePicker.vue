<template>
  <v-menu offset-y>
    <template v-slot:activator="{ props }">
      <v-btn
        variant="text"
        v-bind="props"
        :prepend-icon="languageIcon"
      >
        {{ currentLanguageLabel }}
      </v-btn>
    </template>

    <v-list>
      <v-list-item
        v-for="lang in languages"
        :key="lang.value"
        @click="setLanguage(lang.value)"
      >
        <v-list-item-title>
          <v-icon :icon="lang.icon" class="mr-2" small/>
          {{ lang.label }}
        </v-list-item-title>
      </v-list-item>
    </v-list>
  </v-menu>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useAuthStore } from '@/stores/authStore';

const authStore = useAuthStore();

type LanguageValue = 'zh' | 'en';

const languages: { value: LanguageValue; label: string; icon: string }[] = [
  { value: 'zh', label: '中文', icon: 'mdi-ideogram-cjk' },
  { value: 'en', label: 'English', icon: 'mdi-alphabetical' },
];

const currentLanguageLabel = computed(() => {
  return languages.find((lang) => lang.value === authStore.language)?.label || 'English';
});

const languageIcon = computed(() => {
  return languages.find((lang) => lang.value === authStore.language)?.icon || 'mdi-translate';
});

const setLanguage = (lang: 'zh' | 'en') => {
  authStore.setLanguage(lang);
};
</script>

<style scoped>
.v-btn {
  text-transform: none;
}
</style>
