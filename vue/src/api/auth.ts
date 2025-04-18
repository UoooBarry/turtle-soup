import axios from '@/api/axios'
import type { User } from '@/types/user'

export interface LoginResponse {
  token: string
  user_id: number
  username: string
}

export interface RegisterRequest {
  username: string
  password: string
  email: string
}

export interface LoginRequest {
  username: string
  password: string
}

export async function postLogin(payload: LoginRequest): Promise<LoginResponse> {
  const { data } = await axios.post('/auth/login', payload)
  return data
}

export async function getProfile(): Promise<User> {
  const { data } = await axios.get('/auth/profile')
  return data
}
