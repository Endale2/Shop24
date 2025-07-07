<template>
  <div class="login-container">
    <div class="max-w-md mx-auto">
      <!-- Header Section -->
      <div class="mb-8 text-center">
        <h1 class="text-3xl font-bold text-gray-900 tracking-tight uppercase mb-2">Sign In</h1>
        <p class="text-gray-600">Welcome back! Please sign in to your account</p>
      </div>

      <!-- Login Form -->
      <div class="bg-white border border-gray-200 rounded-none p-8">
        <form @submit.prevent="onLogin" class="space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide mb-2">Email</label>
            <input 
              v-model="email" 
              type="email" 
              required 
              placeholder="Enter your email" 
              class="w-full p-3 border border-gray-200 rounded-none focus:ring-2 focus:ring-black focus:outline-none text-sm uppercase"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide mb-2">Password</label>
            <div class="relative">
              <input 
                :type="showPassword ? 'text' : 'password'" 
                v-model="password" 
                required 
                placeholder="Enter your password" 
                class="w-full p-3 pr-12 border border-gray-200 rounded-none focus:ring-2 focus:ring-black focus:outline-none text-sm uppercase"
              />
              <button 
                type="button" 
                class="absolute inset-y-0 right-0 flex items-center pr-3 text-gray-500 hover:text-black transition-colors"
                @click="showPassword = !showPassword"
              >
                <svg v-if="showPassword" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L3 3m6.878 6.878L21 21"></path>
                </svg>
                <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                </svg>
              </button>
            </div>
          </div>

          <div v-if="authStore.error" class="bg-red-50 border border-red-200 rounded-none p-4">
            <div class="flex">
              <svg class="w-5 h-5 text-red-400 mr-3 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
              <p class="text-red-800 text-sm">{{ authStore.error }}</p>
            </div>
          </div>

          <button 
            type="submit" 
            :disabled="authStore.loading" 
            class="w-full bg-black text-white py-3 px-4 rounded-none font-semibold uppercase tracking-wide hover:bg-gray-800 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center"
          >
            <svg v-if="authStore.loading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            {{ authStore.loading ? 'Signing In...' : 'Sign In' }}
          </button>
        </form>

        <!-- Register Link -->
        <div class="mt-8 pt-6 border-t border-gray-200 text-center">
          <p class="text-gray-600 text-sm mb-4">Don't have an account?</p>
          <router-link 
            :to="`/shops/${route.params.shopSlug}/register`" 
            class="inline-block bg-white text-black border border-black py-3 px-6 rounded-none font-semibold uppercase tracking-wide hover:bg-gray-50 transition-colors"
          >
            Create Account
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '../stores/auth';
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
  padding: 2rem 0;
  background: #f7f7f9;
}
</style> 