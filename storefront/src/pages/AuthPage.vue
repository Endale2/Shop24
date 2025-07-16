<template>
  <div class="login-container">
    <div class="max-w-md mx-auto">
      <!-- Header Section -->
      <div class="mb-8 text-center">
        <h1 class="text-3xl font-bold text-gray-900 tracking-tight uppercase mb-2">Sign In</h1>
        <p class="text-gray-600">Welcome back! Please sign in to your account</p>
      </div>
      <!-- OAuth Buttons -->
      <div class="bg-white border border-gray-200 rounded-none p-8">
        <div class="space-y-4">
          <button @click="loginWithGoogle" :disabled="authStore.loading" class="w-full flex justify-center items-center space-x-3 border-2 border-gray-300 rounded-xl py-3 px-4 hover:bg-gray-50 hover:border-gray-400 transition-all duration-200 shadow-sm disabled:opacity-60 disabled:cursor-not-allowed">
            <svg v-if="authStore.loading" class="animate-spin h-5 w-5 text-gray-500" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            <img v-else src="https://www.svgrepo.com/show/355037/google.svg" alt="Google icon" class="h-5 w-5" />
            <span class="text-base font-medium text-gray-700">
              {{ authStore.loading ? 'Connecting to Google...' : 'Continue with Google' }}
            </span>
          </button>
          <!-- Telegram Login Widget -->
          <div id="telegram-login-widget" class="w-full flex justify-center"></div>
        </div>
        <div v-if="authStore.error" class="bg-red-50 border border-red-200 rounded-none p-4 mt-6">
          <div class="flex">
            <svg class="w-5 h-5 text-red-400 mr-3 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
            <p class="text-red-800 text-sm">{{ authStore.error }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { onMounted } from 'vue';
import { useAuthStore } from '../stores/auth';
import { useRouter } from 'vue-router';
const authStore = useAuthStore();
const router = useRouter();
function loginWithGoogle() {
  authStore.loginWithGoogle();
}
onMounted(() => {
  window.onTelegramAuth = async function(user) {
    try {
      await authStore.loginWithTelegram(user);
      await authStore.fetchMe();
      router.push('/');
    } catch (err) {
      authStore.error = 'Telegram login failed. Please try again.';
    }
  };
  const script = document.createElement('script');
  script.src = 'https://telegram.org/js/telegram-widget.js?7';
  script.setAttribute('data-telegram-login', 'your_customer_bot');
  script.setAttribute('data-size', 'large');
  script.setAttribute('data-userpic', 'false');
  script.setAttribute('data-request-access', 'write');
  script.setAttribute('data-onauth', 'onTelegramAuth(user)');
  script.async = true;
  document.getElementById('telegram-login-widget').appendChild(script);
});
</script>

<style scoped>
.login-container {
  min-height: 80vh;
  padding: 2rem 0;
  background: #f7f7f9;
}
</style> 