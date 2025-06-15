<template>
  <div class="min-h-screen flex flex-col bg-amber-50 font-sans antialiased">
    <AppNavbar />

    <div class="flex-1 flex flex-col items-center justify-center p-4 sm:p-6 lg:p-8">
      <div class="bg-white rounded-3xl shadow-2xl border border-green-200 w-full max-w-2xl p-8 sm:p-10 lg:p-12 space-y-8 animate-fade-in-up">
        <h2 class="text-3xl sm:text-4xl font-extrabold text-green-800 text-center mb-6">Your Profile</h2>

        <div class="flex flex-col items-center space-y-4 pb-6 border-b border-gray-100">
          <img
            v-if="user.profile_image"
            :src="user.profile_image"
            alt="User Avatar"
            class="h-28 w-28 rounded-full object-cover ring-4 ring-green-200 shadow-md transition-transform duration-300 hover:scale-105"
          />
          <div
            v-else
            class="h-28 w-28 bg-green-600 rounded-full flex items-center justify-center text-white text-3xl font-bold ring-4 ring-green-200 shadow-md"
          >
            {{ initials }}
          </div>
          <p class="text-2xl font-semibold text-gray-800 mt-2">
            {{ user.first_name || 'Anonymous' }} {{ user.last_name || '' }}
          </p>
          <p class="text-lg text-gray-600">{{ user.email || 'No Email' }}</p>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-x-8 gap-y-6 pt-4">
          <div>
            <p class="text-sm text-gray-500 font-medium mb-1">Phone Number</p>
            <p class="text-lg font-semibold text-gray-800">{{ user.phone || 'N/A' }}</p>
          </div>
          <div>
            <p class="text-sm text-gray-500 font-medium mb-1">Business Address</p>
            <p class="text-lg font-semibold text-gray-800">{{ user.address || 'N/A' }}</p>
          </div>
          <div class="sm:col-span-2">
            <p class="text-sm text-gray-500 font-medium mb-1">Business Name</p>
            <p class="text-lg font-semibold text-gray-800">{{ user.business_name || 'Not Set' }}</p>
          </div>
        </div>

        <div class="bg-amber-50 p-5 sm:p-6 rounded-xl border border-amber-200 grid grid-cols-2 sm:grid-cols-4 gap-4 sm:gap-6 text-center shadow-inner">
          <div>
            <p class="text-xs sm:text-sm text-gray-500">Member since</p>
            <p class="font-bold text-gray-800 text-lg sm:text-xl mt-1">{{ formatDate(user.created_at) }}</p>
          </div>
          <div>
            <p class="text-xs sm:text-sm text-gray-500">Total Sales</p>
            <p class="font-bold text-gray-800 text-lg sm:text-xl mt-1">{{ formatCurrency(user.total_sales) }}</p>
          </div>
          <div>
            <p class="text-xs sm:text-sm text-gray-500">Ratings</p>
            <p class="font-bold text-gray-800 text-lg sm:text-xl mt-1">{{ formatRating(user.ratings) }}</p>
          </div>
          <div>
            <p class="text-xs sm:text-sm text-gray-500">Reviews</p>
            <p class="font-bold text-gray-800 text-lg sm:text-xl mt-1">{{ formatReviews(user.reviews_count) }}</p>
          </div>
        </div>

        <div class="pt-6 border-t border-green-100 flex flex-col sm:flex-row justify-center space-y-4 sm:space-y-0 sm:space-x-4">
          <router-link
            to="/dashboard"
            class="px-8 py-3 bg-amber-100 text-green-700 font-semibold rounded-full hover:bg-amber-200 transition transform hover:scale-[1.01] text-lg text-center"
          >
            Back to Dashboard
          </router-link>
          <button
            @click="logout"
            class="px-8 py-3 bg-red-600 text-white font-semibold rounded-full hover:bg-red-700 transition transform hover:scale-[1.01] text-lg"
          >
            Logout
          </button>
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
@keyframes fade-in-up {
  from { opacity: 0; transform: translateY(20px); }
  to   { opacity: 1; transform: translateY(0); }
}
.animate-fade-in-up {
  animation: fade-in-up 0.5s ease-out forwards;
}
</style>