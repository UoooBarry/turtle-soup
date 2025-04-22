<template>
    <v-dialog v-model="dialog" max-width="500px">
        <template v-slot:activator="{ props }">
            <a class="cursor-pointer" v-if="!authStore.isAuthenticated" v-bind="props">login</a>
            <a class="cursor-pointer" @click="handleLogout" v-else>logout</a>
        </template>
        <v-card class="pa-1">
            <v-card-title>
                <div class="d-flex flex-row justify-space-between">
                    <div>
                        Login
                    </div>
                    <div>
                        <LanguagePicker></LanguagePicker>
                    </div>
                </div>
            </v-card-title>
            <v-card-text>
                <v-form @submit.prevent="handleLogin">
                    <v-text-field v-model="username" label="Username" type="username" required></v-text-field>
                    <v-text-field v-model="password" label="Password" type="password" required></v-text-field>
                    <v-btn type="submit" color="primary" block>Login</v-btn>
                </v-form>
            </v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="error" @click="dialog = false">Close</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '@/stores/authStore';
import LanguagePicker from './LanguagePicker.vue';

const authStore = useAuthStore();
const dialog = ref(false);
const username = ref('');
const password = ref('');

const handleLogin = async () => {
    try {
        await authStore.login({
            username: username.value,
            password: password.value,
        });
        dialog.value = false;
    } catch (error) {
        console.error('Login failed:', error);
    }
};

const handleLogout = () => {
    authStore.logout()
}
</script>

<style scoped>
.cursor-pointer {
    cursor: pointer;
}
</style>

