<template>
  <div class="login-container">
    <div class="max-w-md mx-auto">
      <div class="mb-8 text-center">
        <h1 class="text-3xl font-bold text-gray-900 tracking-tight uppercase mb-2">Sign In</h1>
        <p class="text-gray-600">Sign in with your email using OTP</p>
      </div>
      <div class="bg-white border border-gray-200 rounded-none p-8">
        <form v-if="!authStore.otpRequested" @submit.prevent="onRequestOTP">
          <div>
            <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide mb-2">Email</label>
            <input v-model="email" type="email" required placeholder="Enter your email" class="w-full p-3 border border-gray-200 rounded-none focus:ring-2 focus:ring-black focus:outline-none text-sm uppercase" />
          </div>
          <button type="submit" :disabled="authStore.loading" class="w-full bg-black text-white py-3 px-4 rounded-none font-semibold uppercase tracking-wide hover:bg-gray-800 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center mt-4">
            <svg v-if="authStore.loading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            {{ authStore.loading ? 'Sending OTP...' : 'Send OTP' }}
          </button>
        </form>
        <form v-else @submit.prevent="onVerifyOTP">
          <div>
            <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide mb-2">Enter OTP</label>
            <input v-model="otp" type="text" required maxlength="6" placeholder="Enter the OTP sent to your email" class="w-full p-3 border border-gray-200 rounded-none focus:ring-2 focus:ring-black focus:outline-none text-sm uppercase" />
          </div>
          <button type="submit" :disabled="authStore.loading" class="w-full bg-black text-white py-3 px-4 rounded-none font-semibold uppercase tracking-wide hover:bg-gray-800 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center mt-4">
            <svg v-if="authStore.loading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            {{ authStore.loading ? 'Verifying...' : 'Verify OTP' }}
          </button>
        </form>
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
import { ref, onMounted } from 'vue';
import { useAuthStore } from '../stores/auth';
import { useCartStore } from '../stores/cart';

const authStore = useAuthStore();
const email = ref('');
const otp = ref('');

onMounted(async () => {
  localStorage.clear();
  if (typeof authStore.$reset === 'function') authStore.$reset();
  // Reset loading and error state to prevent stuck spinner
  authStore.loading = false;
  authStore.error = null;
  const cartStore = useCartStore();
  if (typeof cartStore.$reset === 'function') cartStore.$reset();
  try {
    const { useWishlistStore } = await import('../stores/wishlist');
    const wishlistStore = useWishlistStore();
    if (typeof wishlistStore.$reset === 'function') wishlistStore.$reset();
  } catch {}
});

function onRequestOTP() {
  authStore.requestOTP(email.value);
}
async function onVerifyOTP() {
  try {
    await authStore.verifyOTP(otp.value);
    email.value = '';
    otp.value = '';
    // Redirect is now handled in the auth store
  } catch (error) {
    console.error('Failed to verify OTP:', error);
  }
}

// Remove the watch since redirect is handled in auth store
</script>

<style scoped>
.login-container {
  min-height: 80vh;
  padding: 2rem 0;
  background: #f7f7f9;
}
</style> 