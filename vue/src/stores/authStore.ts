import { defineStore } from 'pinia'
import { ref, type Ref } from 'vue'
import type { User } from '@/types/user'
import { postLogin, getProfile } from '@/api/auth'
import type { LoginResponse, LoginRequest } from '@/api/auth'

export const useAuthStore = defineStore('auth', () => {
  const user: Ref<User | null> = ref(null)
  const token: Ref<string | null> = ref(null)
  const isAuthenticated: Ref<boolean> = ref(false)

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
    const storeToken = localStorage.getItem('auth-token')

    if (storeToken) {
      token.value = storeToken
      isAuthenticated.value = true

      await fetchProfile()
    }
  }

  function logout(): void {
    user.value = null
    token.value = null
    isAuthenticated.value = false
    localStorage.removeItem('auth-token')
  }

  return {
    user,
    token,
    isAuthenticated,
    login,
    fetchProfile,
    logout,
    restoreLogin
  }
})

