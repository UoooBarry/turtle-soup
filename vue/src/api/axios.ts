import { default as axios } from 'axios';
import { MessageType } from '@/types/globalMessage';

const instance = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080/api',
});

instance.interceptors.request.use(async function(config) {
  const { useAuthStore } = await import('@/stores/authStore');
  const authStore = useAuthStore();
  if (authStore.isAuthenticated && authStore.token) {
    config.headers.Authorization = `Bearer ${authStore.token}`;
  }

  return config;
}, function(error) {
  return Promise.reject(error);
});

instance.interceptors.response.use(function(response) {
  return response;
}, async function(error) {
  if (error.response && error.response.data && error.response.data.error) {
    const { useMessageStore } = await import('@/stores/messageStore');
    const messageStore = useMessageStore();
    messageStore.pushMessage(error.response.data.error, MessageType.Error);
  }
  return Promise.reject(error);
});

export default instance;
