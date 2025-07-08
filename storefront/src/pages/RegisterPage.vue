<template>
  <div class="register-container">
    <div class="max-w-md mx-auto">
      <!-- Header Section -->
      <div class="mb-8 text-center">
        <h1 class="text-3xl font-bold text-gray-900 tracking-tight uppercase mb-2">Create Account</h1>
        <p class="text-gray-600">Join us and start shopping today</p>
      </div>

      <!-- Register Form -->
      <div class="bg-white border border-gray-200 rounded-none p-8">
        <form @submit.prevent="onRegister" class="space-y-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
      <div>
              <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide mb-2">First Name</label>
              <input 
                v-model="firstName" 
                type="text" 
                required 
                placeholder="Enter your first name" 
                class="w-full p-3 border border-gray-200 rounded-none focus:ring-2 focus:ring-black focus:outline-none text-sm uppercase"
              />
      </div>
      <div>
              <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide mb-2">Last Name</label>
              <input 
                v-model="lastName" 
                type="text" 
                required 
                placeholder="Enter your last name" 
                class="w-full p-3 border border-gray-200 rounded-none focus:ring-2 focus:ring-black focus:outline-none text-sm uppercase"
              />
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide mb-2">Username</label>
            <input 
              v-model="username" 
              type="text" 
              required 
              placeholder="Choose a username" 
              class="w-full p-3 border border-gray-200 rounded-none focus:ring-2 focus:ring-black focus:outline-none text-sm uppercase"
            />
          </div>

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
            <input 
              v-model="password" 
              type="password" 
              required 
              placeholder="Create a password" 
              class="w-full p-3 border border-gray-200 rounded-none focus:ring-2 focus:ring-black focus:outline-none text-sm uppercase"
            />
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
            {{ authStore.loading ? 'Creating Account...' : 'Create Account' }}
          </button>
        </form>

        <!-- Login Link -->
        <div class="mt-8 pt-6 border-t border-gray-200 text-center">
          <p class="text-gray-600 text-sm mb-4">Already have an account?</p>
          <router-link 
            :to="`/shops/${route.params.shopSlug}/login`" 
            class="inline-block bg-white text-black border border-black py-3 px-6 rounded-none font-semibold uppercase tracking-wide hover:bg-gray-50 transition-colors"
          >
            Sign In
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
const firstName = ref('');
const lastName = ref('');
const username = ref('');
const email = ref('');
const password = ref('');

async function onRegister() {
  const shopSlug = route.params.shopSlug as string;
  await authStore.register(username.value, email.value, password.value, shopSlug, firstName.value, lastName.value);
  if (authStore.user) {
    router.push(`/shops/${shopSlug}`);
  }
}
</script> 

<style scoped>
.register-container {
  min-height: 80vh;
  padding: 2rem 0;
  background: #f7f7f9;
}
</style> 