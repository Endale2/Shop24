<template>
  <header class="w-full sticky top-0 z-50 bg-white/80 backdrop-blur-md border-b border-gray-200 shadow-sm">
    <div class="max-w-7xl mx-auto flex items-center justify-between p-4 lg:px-8">
      <router-link to="/" class="flex items-center space-x-3 group">
        <div class="p-2 bg-green-500 rounded-full group-hover:rotate-12 transition-transform duration-500 shadow-md">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
          </svg>
        </div>
        <span class="text-2xl font-extrabold text-gray-800">Seller’s <span class="text-amber-600">Dash</span></span>
      </router-link>

      <nav class="flex items-center space-x-6">
        <template v-if="user && user.id">
          <div class="relative group flex items-center space-x-3 cursor-pointer">
            <img
              v-if="user.profile_image"
              :src="user.profile_image"
              alt="Profile"
              class="h-9 w-9 rounded-full object-cover ring-2 ring-green-300 group-hover:ring-green-500 transition-all duration-200"
            />
            <div
              v-else
              class="h-9 w-9 bg-green-600 rounded-full flex items-center justify-center text-white font-bold text-sm ring-2 ring-green-300 group-hover:ring-green-500 transition-all duration-200"
            >
              {{ userInitials }}
            </div>
            <span class="text-gray-800 font-semibold hidden sm:inline">{{ user.name || user.email }}</span>

            <div class="absolute right-0 top-full mt-2 w-48 bg-white rounded-lg shadow-xl py-1 opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all duration-200 transform scale-95 group-hover:scale-100 origin-top-right">
              <router-link
                to="/profile"
                class="block px-4 py-2 text-gray-700 hover:bg-green-100 hover:text-green-800 transition-colors duration-150"
              >
                Your Profile
              </router-link>
              <button
                @click="handleLogout"
                class="w-full text-left px-4 py-2 text-red-600 hover:bg-red-50 hover:text-red-700 transition-colors duration-150"
              >
                Logout
              </button>
            </div>
          </div>
        </template>
        <template v-else>
          <router-link to="/login" class="text-gray-700 font-medium hover:text-green-600 transition-colors duration-200">
            Log In
          </router-link>
          <router-link to="/register"
            class="px-5 py-2 bg-green-600 text-white font-semibold rounded-full shadow-lg hover:bg-green-700 transition-all duration-300 transform hover:-translate-y-0.5">
            Get Started
          </router-link>
        </template>
      </nav>
    </div>
  </header>
</template>

<script setup>
import { computed } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/store/auth';

const router = useRouter();
const auth = useAuthStore();

const user = computed(() => auth.user || {});
const userInitials = computed(() => {
  if (user.value.name) {
    return user.value.name
      .split(' ')
      .map(n => n[0])
      .join('')
      .toUpperCase();
  }
  const [first, last] = (user.value.email || '').split('@')[0].split('.');
  return ((first?.[0]||'') + (last?.[0]||'')).toUpperCase();
});

async function handleLogout() {
  await auth.logout();
  router.push({ name: 'Landing' });
}
</script>

<style scoped>
</style>