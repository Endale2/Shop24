<template>
  <div class="login-container">
    <div class="login-card">
      <h1>Login</h1>
      <form @submit.prevent="onLogin">
        <div class="form-group">
          <label>Email</label>
          <input v-model="email" type="email" required placeholder="Enter your email" />
        </div>
        <div class="form-group">
          <label>Password</label>
          <div class="password-input">
            <input :type="showPassword ? 'text' : 'password'" v-model="password" required placeholder="Enter your password" />
            <button type="button" class="toggle-btn" @click="showPassword = !showPassword">
              {{ showPassword ? 'Hide' : 'Show' }}
            </button>
          </div>
        </div>
        <div v-if="authStore.error" class="error">{{ authStore.error }}</div>
        <button type="submit" :disabled="authStore.loading" class="login-btn">
          <span v-if="authStore.loading" class="spinner"></span>
          Login
        </button>
      </form>
      <p class="mt-6 text-sm text-center">
        <router-link to="/register" class="text-blue-600 cursor-pointer hover:underline uppercase">
          Don't have an account?
        </router-link>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { useRouter, useRoute } from 'vue-router';

const authStore = useAuthStore();
const router = useRouter();
const route = useRoute();
const email = ref('');
const password = ref('');
const showPassword = ref(false);

async function onLogin() {
  const shopSlug = route.params.shopSlug as string;
  await authStore.login(email.value, password.value, shopSlug);
  if (authStore.user) {
    router.push(`/shops/${shopSlug}`);
  }
}
</script>

<style scoped>
.login-container {
  min-height: 80vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f7f7f9;
}
.login-card {
  background: #fff;
  padding: 2.5rem 2rem 2rem 2rem;
  border-radius: 12px;
  box-shadow: 0 2px 16px rgba(0,0,0,0.07);
  width: 100%;
  max-width: 350px;
}
.login-card h1 {
  margin-bottom: 1.5rem;
  font-size: 1.7rem;
  text-align: center;
}
.form-group {
  margin-bottom: 1.2rem;
}
.form-group label {
  display: block;
  margin-bottom: 0.4rem;
  font-weight: 500;
}
.form-group input {
  width: 100%;
  padding: 0.6rem 0.8rem;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  font-size: 1rem;
  background: #fafbfc;
}
.password-input {
  display: flex;
  align-items: center;
}
.password-input input {
  flex: 1;
}
.toggle-btn {
  margin-left: 0.5rem;
  background: none;
  border: none;
  color: #007bff;
  cursor: pointer;
  font-size: 0.95rem;
  padding: 0 0.3rem;
}
.error {
  color: #d32f2f;
  margin-bottom: 1rem;
  text-align: center;
}
.login-btn {
  width: 100%;
  padding: 0.7rem;
  background: #007bff;
  color: #fff;
  border: none;
  border-radius: 6px;
  font-size: 1.1rem;
  font-weight: 600;
  cursor: pointer;
  margin-top: 0.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
}
.login-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}
.spinner {
  border: 2px solid #f3f3f3;
  border-top: 2px solid #fff;
  border-radius: 50%;
  width: 16px;
  height: 16px;
  margin-right: 8px;
  animation: spin 1s linear infinite;
  display: inline-block;
}
@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style> 