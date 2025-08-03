<template>
  <div class="min-h-screen flex flex-col bg-gradient-to-br from-gray-50 to-gray-100 font-sans antialiased">
    <AppNavbar />

    <div class="flex-1 flex flex-col items-center pt-8 px-4 sm:px-6 pb-8">
      <div class="w-full max-w-4xl">
        <!-- Header -->
        <div class="text-center mb-8">
          <div class="flex items-center justify-center mb-4">
            <div class="w-12 h-12 bg-gradient-to-br from-green-500 to-green-600 rounded-xl flex items-center justify-center mr-4 shadow-lg">
              <svg class="w-7 h-7 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
              </svg>
            </div>
            <h1 class="text-3xl sm:text-4xl font-bold text-gray-800 leading-tight">
              Your Profile
            </h1>
          </div>
          <p class="text-base text-gray-600 max-w-2xl mx-auto">
            Manage your account information and preferences to personalize your experience.
          </p>
        </div>

        <!-- Profile Card -->
        <div class="bg-white rounded-2xl shadow-lg border border-gray-200 overflow-hidden">
          <!-- Profile Header -->
          <div class="bg-gradient-to-r from-green-50 to-amber-50 px-8 py-6 border-b border-gray-100">
            <div class="flex flex-col sm:flex-row items-center space-y-4 sm:space-y-0 sm:space-x-6">
              <!-- Avatar -->
              <div class="relative group">
                <img
                  v-if="user.profile_image"
                  :src="user.profile_image"
                  alt="User Avatar"
                  class="h-24 w-24 rounded-full object-cover ring-4 ring-green-200 shadow-lg transition-all duration-300 hover:scale-105 group-hover:ring-green-300"
                />
                <div
                  v-else
                  class="h-24 w-24 bg-gradient-to-br from-green-500 to-green-700 rounded-full flex items-center justify-center text-white text-3xl font-bold ring-4 ring-green-200 shadow-lg transition-all duration-300 hover:scale-105 group-hover:ring-green-300"
                >
                  {{ initials }}
                </div>
                <div class="absolute -bottom-1 -right-1 w-8 h-8 bg-green-600 rounded-full flex items-center justify-center shadow-lg cursor-pointer hover:bg-green-700 transition-colors duration-200">
                  <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9a2 2 0 012-2h.93a2 2 0 001.664-.89l.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 13a3 3 0 11-6 0 3 3 0 016 0z" />
                  </svg>
                </div>
              </div>
              
              <!-- User Info -->
              <div class="text-center sm:text-left">
                <h3 class="text-2xl font-bold text-gray-800 mb-1">
                  {{ user.first_name || 'Anonymous' }} {{ user.last_name || '' }}
                </h3>
                <p class="text-lg text-gray-600 flex items-center justify-center sm:justify-start">
                  <svg class="w-5 h-5 mr-2 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207" />
                  </svg>
                  {{ user.email || 'No Email' }}
                </p>
              </div>
            </div>
          </div>

          <!-- Profile Content -->
          <div class="p-8">
            <!-- Contact Information -->
            <div class="mb-8">
              <h4 class="text-xl font-semibold text-gray-800 mb-6 flex items-center">
                <svg class="w-6 h-6 mr-3 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z" />
                </svg>
                Contact Information
              </h4>
              
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div class="bg-gray-50 rounded-xl p-4">
                  <p class="text-sm text-gray-500 font-medium mb-1">Phone Number</p>
                  <p class="text-lg font-semibold text-gray-800">{{ user.phone || 'Not provided' }}</p>
                </div>
                
                <div class="bg-gray-50 rounded-xl p-4">
                  <p class="text-sm text-gray-500 font-medium mb-1">Business Name</p>
                  <p class="text-lg font-semibold text-gray-800">{{ user.business_name || 'Not set' }}</p>
                </div>
                
                <div class="bg-gray-50 rounded-xl p-4 md:col-span-2">
                  <p class="text-sm text-gray-500 font-medium mb-1">Business Address</p>
                  <p class="text-lg font-semibold text-gray-800">{{ user.address || 'Not provided' }}</p>
                </div>
              </div>
            </div>

            <!-- Account Statistics -->
            <div class="mb-8">
              <h4 class="text-xl font-semibold text-gray-800 mb-6 flex items-center">
                <svg class="w-6 h-6 mr-3 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                </svg>
                Account Statistics
              </h4>
              
              <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
                <div class="bg-gradient-to-br from-green-50 to-green-100 rounded-xl p-4 text-center">
                  <p class="text-xs text-gray-600 font-medium mb-1">Member since</p>
                  <p class="font-bold text-gray-800 text-lg">{{ formatDate(user.created_at) }}</p>
                </div>
                
                <div class="bg-gradient-to-br from-blue-50 to-blue-100 rounded-xl p-4 text-center">
                  <p class="text-xs text-gray-600 font-medium mb-1">Total Sales</p>
                  <p class="font-bold text-gray-800 text-lg">{{ formatCurrency(user.total_sales) }}</p>
                </div>
                
                <div class="bg-gradient-to-br from-amber-50 to-amber-100 rounded-xl p-4 text-center">
                  <p class="text-xs text-gray-600 font-medium mb-1">Ratings</p>
                  <p class="font-bold text-gray-800 text-lg">{{ formatRating(user.ratings) }}</p>
                </div>
                
                <div class="bg-gradient-to-br from-purple-50 to-purple-100 rounded-xl p-4 text-center">
                  <p class="text-xs text-gray-600 font-medium mb-1">Reviews</p>
                  <p class="font-bold text-gray-800 text-lg">{{ formatReviews(user.reviews_count) }}</p>
                </div>
              </div>
            </div>

            <!-- Action Buttons -->
            <div class="flex flex-col sm:flex-row justify-center space-y-4 sm:space-y-0 sm:space-x-6 pt-6 border-t border-gray-100">
              <router-link
                to="/dashboard"
                class="inline-flex items-center justify-center px-8 py-3 bg-gradient-to-r from-green-600 to-green-700 text-white font-semibold rounded-lg hover:from-green-700 hover:to-green-800 transition-all duration-300 transform hover:-translate-y-0.5 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2 shadow-lg"
              >
                <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                </svg>
                Back to Dashboard
              </router-link>
              
              <button
                @click="logout"
                class="inline-flex items-center justify-center px-8 py-3 bg-gradient-to-r from-red-500 to-red-600 text-white font-semibold rounded-lg hover:from-red-600 hover:to-red-700 transition-all duration-300 transform hover:-translate-y-0.5 focus:outline-none focus:ring-2 focus:ring-red-400 focus:ring-offset-2 shadow-lg"
              >
                <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                </svg>
                Logout
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <AppFooter />
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import AppNavbar from '@/components/AppNavbar.vue'
import AppFooter from '@/components/AppFooter.vue'

const auth = useAuthStore()
const router = useRouter()

// Redirect if unauthenticated
import { watchEffect } from 'vue';
watchEffect(() => {
  if (!auth.isAuthenticated) {
    router.replace({ name: 'Landing' })
  }
})

const user = computed(() => auth.user || {})

const initials = computed(() => {
  const fn = user.value.first_name || ''
  const ln = user.value.last_name || ''
  if (fn || ln) return `${fn[0]||''}${ln[0]||''}`.toUpperCase()
  const [f, l] = (user.value.email || '').split('@')[0].split('.')
  return ((f?.[0]||'') + (l?.[0]||'')).toUpperCase()
})

function formatDate(ts) {
  return ts
    ? new Date(ts).toLocaleDateString(undefined, { year: 'numeric', month: 'short', day: 'numeric' })
    : 'N/A'
}
function formatCurrency(val) {
  return typeof val === 'number'
    ? val.toLocaleString(undefined, { style: 'currency', currency: 'USD' })
    : '$0.00' // Default for no sales
}
function formatRating(val) {
  return typeof val === 'number' && !isNaN(val) // Check for valid number
    ? `${val.toFixed(1)} ★`
    : 'N/A'
}
function formatReviews(val) {
  return Number.isInteger(val)
    ? val
    : '0' // Default for no reviews
}

async function logout() {
  await auth.logout()
  router.push({ name: 'Landing' })
}
</script>

<style scoped>
/* No additional styles needed - using Tailwind classes */
</style>