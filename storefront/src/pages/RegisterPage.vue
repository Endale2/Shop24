<template>
  <div>
    <h1>Register</h1>
    <form @submit.prevent="onRegister">
      <div>
        <label>Username:</label>
        <input v-model="username" type="text" required />
      </div>
      <div>
        <label>Email:</label>
        <input v-model="email" type="email" required />
      </div>
      <div>
        <label>Password:</label>
        <input v-model="password" type="password" required />
      </div>
      <div v-if="authStore.error" class="error">{{ authStore.error }}</div>
      <button type="submit" :disabled="authStore.loading">Register</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { useRouter, useRoute } from 'vue-router';

const authStore = useAuthStore();
const router = useRouter();
const route = useRoute();
const username = ref('');
const email = ref('');
const password = ref('');

async function onRegister() {
  const shopSlug = route.params.shopSlug as string;
  await authStore.register(username.value, email.value, password.value, shopSlug);
  if (authStore.user) {
    router.push(`/shops/${shopSlug}`);
  }
}
</script> 