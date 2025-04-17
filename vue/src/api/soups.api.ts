import axios from './axios'
import type { Soup } from '@/types/soup'

export const getSoups = (): Promise<{ data: Soup[] }> => {
  return axios.get('/soups')
}
