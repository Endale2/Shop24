<template>
  <div class="min-h-screen bg-gray-50 py-10"> <div class="container mx-auto p-4 sm:p-6 lg:p-8">
      <h1 class="text-4xl font-extrabold text-gray-900 mb-8 text-center">Your Account Dashboard</h1> <div v-if="auth.isAuthenticated" class="bg-white shadow-xl rounded-xl p-8 max-w-3xl mx-auto border border-gray-200"> <div class="flex flex-col sm:flex-row items-center sm:space-x-8 mb-8 pb-8 border-b border-gray-100"> <div v-if="auth.user.avatar" class="flex-shrink-0 h-28 w-28 rounded-full overflow-hidden border-4 border-green-400 shadow-lg mb-4 sm:mb-0"> <img :src="auth.user.avatar" :alt="`${auth.user.displayName || 'User'} avatar`" class="object-cover h-full w-full"/>
          </div>
          <div v-else class="flex-shrink-0 h-28 w-28 rounded-full flex items-center justify-center bg-green-200 text-green-700 font-bold text-5xl border-4 border-green-400 shadow-lg mb-4 sm:mb-0"> {{ initials }}
          </div>

          <div class="text-center sm:text-left">
            <h2 class="text-3xl font-bold text-gray-800 mb-1">{{ auth.user.displayName }}</h2>
            <p class="text-gray-600 text-lg">{{ auth.user.email }}</p>
            <p v-if="auth.user.username" class="text-gray-500 text-base mt-1">Username: <span class="font-medium">{{ auth.user.username }}</span></p>
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6"> <router-link to="/account/profile" class="flex flex-col items-center p-6 bg-gray-50 rounded-xl hover:bg-green-50 transition-all duration-300 transform hover:scale-105 shadow-sm hover:shadow-md border border-gray-100"> <svg xmlns="http://www.w3.org/2000/svg" class="h-10 w-10 text-green-600 mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
            </svg>
            <span class="font-semibold text-lg text-gray-800">Edit Profile</span>
            <p class="text-sm text-gray-500 text-center mt-1">Manage your personal information</p>
          </router-link>

          <router-link to="/account/orders" class="flex flex-col items-center p-6 bg-gray-50 rounded-xl hover:bg-green-50 transition-all duration-300 transform hover:scale-105 shadow-sm hover:shadow-md border border-gray-100"> <svg xmlns="http://www.w3.org/2000/svg" class="h-10 w-10 text-green-600 mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01" />
            </svg>
            <span class="font-semibold text-lg text-gray-800">Order History</span>
            <p class="text-sm text-gray-500 text-center mt-1">View past and current orders</p>
          </router-link>

          <button @click="onLogout" class="flex items-center justify-center space-x-3 p-4 bg-red-50 text-red-700 rounded-xl hover:bg-red-100 transition-colors duration-200 col-span-full mt-6 shadow-sm hover:shadow-md font-medium text-lg"> <svg xmlns="http://www.w3.org/2000/svg" class="h-7 w-7" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
            </svg>
            <span>Sign Out</span>
          </button>
        </div>

      </div>

      <div v-else class="text-center bg-white shadow-xl rounded-xl p-10 max-w-lg mx-auto border border-gray-200"> <svg class="mx-auto h-16 w-16 text-gray-400 mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
          <path vector-effect="non-scaling-stroke" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12H15M9 17H15M12 3a7 7 0 00-7 7v4a7 7 0 007 7h0a7 7 0 007-7v-4a7 7 0 00-7-7z" />
        </svg>
        <p class="text-xl text-gray-800 font-semibold mb-4">Access Denied</p>
        <p class="text-md text-gray-700 mb-6">You need to be logged in to view your account details.</p>
        <router-link to="/auth" class="inline-flex items-center px-8 py-3 border border-transparent text-lg font-medium rounded-full shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 transform hover:scale-105 transition-transform duration-200">
          Sign In / Sign Up
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';

const auth = useAuthStore();
const router = useRouter();

// Compute initials for the avatar placeholder
const initials = computed(() => {
  if (!auth.user || (!auth.user.displayName && !auth.user.email)) {
    return '';
  }
  const name = auth.user.displayName || auth.user.email.split('@')[0];
  return name
    .split(' ')
    .map(n => n[0])
    .join('')
    .toUpperCase()
    .substring(0, 2); // Take max 2 initials
});

const onLogout = async () => {
  await auth.logout();
  router.push({ name: 'Home' });
};
</script>

<style scoped>
/* No specific styles needed here, Tailwind CSS handles everything */
</style>