import { defineStore } from 'pinia'
import { ref, type Ref, watch } from 'vue'
import type { User } from '@/types/user'
import { postLogin, getProfile } from '@/api/auth'
import type { LoginResponse, LoginRequest } from '@/api/auth'
import i18n from '@/i18n'

export const useAuthStore = defineStore('auth', () => {
  const user: Ref<User | null> = ref(null)
  const token: Ref<string | null> = ref(null)
  const isAuthenticated: Ref<boolean> = ref(false)
  const language: Ref<'zh' | 'en'> = ref('zh')

  watch(language, (newLang) => {
    i18n.global.locale.value = newLang;
  })

  async function login(payload: LoginRequest): Promise<LoginResponse> {
    try {
      const data = await postLogin(payload)
      token.value = data.token
      user.value = { id: data.user_id, username: data.username }
      localStorage.setItem('auth-token', data.token)

      isAuthenticated.value = true
      return data
    } catch (error) {
      isAuthenticated.value = false
      throw error
    }
  }

  async function fetchProfile(): Promise<User> {
    try {
      const data = await getProfile()
      user.value = data
      return data
    } catch (error) {
      logout()
      throw error
    }
  }

  async function restoreLogin() {
    const storeToken = localStorage.getItem('auth-token');
    const storeLang = localStorage.getItem('lang') as 'zh' | 'en' | null;

    if (storeToken) {
      token.value = storeToken
      isAuthenticated.value = true

      await fetchProfile()
    }
    if (storeLang) {
      language.value = storeLang
    }
  }

  function logout(): void {
    user.value = null
    token.value = null
    isAuthenticated.value = false
    localStorage.removeItem('auth-token')
  }

  function setLanguage(lang: 'zh' | 'en') {
    language.value = lang
  }

  return {
    user,
    token,
    isAuthenticated,
    login,
    fetchProfile,
    logout,
    restoreLogin,
    language,
    setLanguage
  }
})

