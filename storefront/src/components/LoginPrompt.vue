<template>
  <div class="text-center py-16 flex flex-col items-center justify-center">
    <slot name="icon">
      <!-- Default Icon -->
      <svg class="w-24 h-24 text-gray-300 mx-auto mb-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
      </svg>
    </slot>
    <h3 class="text-2xl font-bold text-gray-900 mb-2">{{ title }}</h3>
    <p class="text-gray-500 mb-6 max-w-md">{{ message }}</p>
    <button @click="goToLogin" class="inline-block bg-black text-white py-3 px-8 rounded-md font-bold uppercase tracking-wide hover:bg-gray-800 transition-colors text-lg shadow-md">
      Login
    </button>
  </div>
</template>

<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router';

defineProps({
  title: {
    type: String,
    required: true,
  },
  message: {
    type: String,
    required: true,
  },
});

const router = useRouter();
const route = useRoute();

function goToLogin() {
  const shopSlug = route.params.shopSlug as string;
  if (shopSlug) {
    router.push({ path: `/${shopSlug}/login` });
  } else {
    // Fallback if shop slug isn't available
    router.push('/'); 
  }
}
</script> 