<template>
  <nav class="flex items-center space-x-2 text-sm text-gray-500 mb-6">
    <button @click="$router.back()" class="text-gray-700 hover:text-black flex items-center gap-1 font-medium text-xs mr-2">
      <svg class="w-4 h-4" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7"/></svg>
      Back
    </button>
    <router-link :to="`/shops/${route.params.shopSlug}`" class="hover:underline">Home</router-link>
    <span>/</span>
    <span class="text-gray-900 font-medium">Account</span>
  </nav>
  <div v-if="authStore.sessionLoading" class="account-container flex items-center justify-center min-h-[60vh]">
    <div class="flex flex-col items-center">
      <svg class="animate-spin h-12 w-12 text-gray-400 mb-4" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      <p class="text-gray-500">Checking authentication…</p>
    </div>
  </div>
  <div v-else-if="user" class="account-container">
    <div class="max-w-4xl mx-auto">
      <!-- Header Section -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 tracking-tight uppercase mb-2">My Account</h1>
        <p class="text-gray-600">Manage your profile and preferences</p>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Profile Card -->
        <div class="lg:col-span-2">
          <div class="bg-white border border-gray-200 rounded-none p-6">
            <h2 class="text-xl font-semibold text-gray-900 mb-6 uppercase">Profile Information</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">Email</label>
                  <p class="mt-1 text-gray-900">{{ user.email }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">First Name</label>
                  <p class="mt-1 text-gray-900">{{ user.firstName || 'Not provided' }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">Last Name</label>
                  <p class="mt-1 text-gray-900">{{ user.lastName || 'Not provided' }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">Phone</label>
                  <p class="mt-1 text-gray-900">{{ user.phone || 'Not provided' }}</p>
                </div>
              </div>
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">Address</label>
                  <p class="mt-1 text-gray-900">{{ user.address || 'Not provided' }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">City</label>
                  <p class="mt-1 text-gray-900">{{ user.city || 'Not provided' }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">State</label>
                  <p class="mt-1 text-gray-900">{{ user.state || 'Not provided' }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">Country</label>
                  <p class="mt-1 text-gray-900">{{ user.country || 'Not provided' }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Sidebar -->
        <div class="space-y-6">
          <!-- Quick Actions -->
          <div class="bg-white border border-gray-200 rounded-none p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4 uppercase">Quick Actions</h3>
            <div class="space-y-3">
              <router-link 
                :to="`/shops/${route.params.shopSlug}/orders`" 
                class="flex items-center p-3 text-gray-700 hover:bg-gray-50 transition-colors rounded-lg"
              >
                <svg class="w-5 h-5 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                </svg>
                My Orders
              </router-link>
              <router-link 
                :to="`/shops/${route.params.shopSlug}/wishlist`" 
                class="flex items-center p-3 text-gray-700 hover:bg-gray-50 transition-colors rounded-lg"
              >
                <svg class="w-5 h-5 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"></path>
                </svg>
                Wishlist
              </router-link>
            </div>
          </div>

          <!-- Account Info -->
          <div class="bg-white border border-gray-200 rounded-none p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4 uppercase">Account Info</h3>
            <div class="space-y-3">
              <div>
                <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">User ID</label>
                <p class="mt-1 text-sm text-gray-600 font-mono">{{ user.id || user._id }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">Member Since</label>
                <p class="mt-1 text-sm text-gray-600">{{ formatDate(user.createdAt) }}</p>
              </div>
            </div>
          </div>

          <!-- Logout -->
          <div class="bg-white border border-gray-200 rounded-none p-6">
            <button 
              @click="onLogout" 
              class="w-full bg-red-600 text-white py-3 px-4 rounded-none font-semibold uppercase tracking-wide hover:bg-red-700 transition-colors"
            >
              Logout
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div v-else class="account-container">
    <div class="max-w-md mx-auto text-center">
      <div class="bg-white border border-gray-200 rounded-none p-8">
        <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
        </svg>
        <h2 class="text-2xl font-bold text-gray-900 mb-4">You are not logged in</h2>
        <p class="text-gray-600 mb-6">Please log in to access your account</p>
        <router-link 
          :to="`/shops/${route.params.shopSlug}/login`" 
          class="inline-block bg-black text-white py-3 px-6 rounded-none font-semibold uppercase tracking-wide hover:bg-gray-800 transition-colors"
        >
          Login
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, watch } from 'vue';
import { useAuthStore } from '../stores/auth';
import { useRouter, useRoute } from 'vue-router';

const authStore = useAuthStore();
const router = useRouter();
const route = useRoute();
const user = computed(() => authStore.user);

watch(
  () => [authStore.sessionLoading, authStore.user],
  ([loading, user]) => {
    if (!loading && !user) {
      // Redirect to login if not authenticated after session check
      router.replace(`/shops/${route.params.shopSlug}/login`)
    }
  },
  { immediate: true }
)

function formatDate(dateString: string) {
  if (!dateString) return 'Unknown';
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  });
}

async function onLogout() {
  await authStore.logout();
  const shopSlug = route.params.shopSlug as string;
  router.push(`/shops/${shopSlug}/login`);
}
</script>

<style scoped>
.account-container {
  min-height: 80vh;
  padding: 2rem 0;
  background: #f7f7f9;
}
</style> 