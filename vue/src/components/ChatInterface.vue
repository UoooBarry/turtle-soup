<script setup lang="ts">
import { ref } from 'vue';

interface Message {
  text: string;
  sender: 'user' | 'bot';
  timestamp: Date;
}

const messages = ref<Message[]>([
  { text: 'Hello! How can I help you today?', sender: 'bot', timestamp: new Date() },
]);

const newMessage = ref('');

const sendMessage = () => {
  if (newMessage.value.trim() === '') return;

  messages.value.push({
    text: newMessage.value,
    sender: 'user',
    timestamp: new Date(),
  });

  newMessage.value = '';

  // Simulate bot response
  setTimeout(() => {
    messages.value.push({
      text: 'Thanks for your message! I will get back to you soon.',
      sender: 'bot',
      timestamp: new Date(),
    });
  }, 1000);
};
</script>

<template>
  <v-card class="chat-container">
    <v-card-title class="d-flex justify-space-between align-center">
      <span>Chat</span>
      <v-btn icon="mdi-close" variant="text" @click="$emit('close')"></v-btn>
    </v-card-title>

    <v-divider></v-divider>

    <v-card-text class="messages-container">
      <v-list lines="two">
        <v-list-item
          v-for="(message, index) in messages"
          :key="index"
          :class="message.sender === 'user' ? 'user-message' : 'bot-message'"
        >
          <v-list-item-content>
            <v-list-item-title>{{ message.sender === 'user' ? 'You' : 'Bot' }}</v-list-item-title>
            <v-list-item-subtitle>{{ message.text }}</v-list-item-subtitle>
          </v-list-item-content>
          <v-list-item-action>
            <v-list-item-subtitle>{{ message.timestamp.toLocaleTimeString() }}</v-list-item-subtitle>
          </v-list-item-action>
        </v-list-item>
      </v-list>
    </v-card-text>

    <v-divider></v-divider>

    <v-card-actions>
      <v-text-field
        v-model="newMessage"
        label="Type a message"
        variant="outlined"
        @keyup.enter="sendMessage"
      ></v-text-field>
      <v-btn icon="mdi-send" color="primary" @click="sendMessage"></v-btn>
    </v-card-actions>
  </v-card>
</template>

<style scoped>
.chat-container {
  max-width: 500px;
  margin: 0 auto;
}

.messages-container {
  height: 400px;
  overflow-y: auto;
}

.user-message {
  text-align: right;
  background-color: #e3f2fd;
  border-radius: 8px;
  padding: 8px;
  margin: 4px 0;
}

.bot-message {
  text-align: left;
  background-color: #f5f5f5;
  border-radius: 8px;
  padding: 8px;
  margin: 4px 0;
}
</style>

