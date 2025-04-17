import type { Soup } from '@/types/soup'
import { defineStore } from 'pinia'
import { ref, type Ref } from 'vue'
import { getSoups } from '@/api/soups.api'

export const useSoupStore = defineStore('soup', () => {
  const soups: Ref<Soup[]> = ref([])
  const currentSoup: Ref<Soup> = ref({ id: NaN, soup_question: '', enabled: false, tag: [] })
  const setCurrentSoup = (soup: Soup) => {
    currentSoup.value = soup
  }
  async function fetchSoups(): Promise<Soup[]> {
    const { data } = await getSoups()
    soups.value = data

    return soups.value
  }
 
  const clearCurrentSoup = () => {
    currentSoup.value = { id: NaN, soup_question: '', enabled: false, tag: [] }
  }

  return { soups, currentSoup, setCurrentSoup, fetchSoups, clearCurrentSoup }
})
