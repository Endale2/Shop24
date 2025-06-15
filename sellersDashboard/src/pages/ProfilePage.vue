<template>
  <div class="min-h-screen flex flex-col bg-amber-100 font-sans antialiased">
    <AppNavbar />

    <div class="flex-1 flex flex-col items-center justify-center p-6">
      <div class="bg-white rounded-3xl shadow-2xl border border-green-200 w-full max-w-2xl p-8 space-y-6 animate-fade-in-up">
        <h2 class="text-3xl font-extrabold text-green-800 text-center">Your Profile</h2>

        <!-- Avatar + Name / Email -->
        <div class="flex flex-col items-center space-y-4">
          <img
            v-if="user.profile_image"
            :src="user.profile_image"
            alt="Avatar"
            class="h-24 w-24 rounded-full object-cover ring-4 ring-green-200"
          />
          <div
            v-else
            class="h-24 w-24 bg-green-600 rounded-full flex items-center justify-center text-white text-2xl font-bold ring-4 ring-green-200"
          >
            {{ initials }}
          </div>
          <p class="text-xl font-semibold text-gray-800">
            {{ user.first_name }} {{ user.last_name }}
          </p>
          <p class="text-gray-600">{{ user.email }}</p>
        </div>

        <!-- Business Details -->
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
          <div class="space-y-2">
            <p class="text-sm text-gray-500">Phone</p>
            <p class="font-medium">{{ user.phone || '—' }}</p>
          </div>
          <div class="space-y-2">
            <p class="text-sm text-gray-500">Address</p>
            <p class="font-medium">{{ user.address || '—' }}</p>
          </div>
          <div class="space-y-2 sm:col-span-2">
            <p class="text-sm text-gray-500">Business Name</p>
            <p class="font-medium">{{ user.business_name || '—' }}</p>
          </div>
        </div>

        <!-- Metrics & Dates -->
        <div class="flex flex-col sm:flex-row items-center justify-between bg-amber-50 p-4 rounded-lg border border-amber-200">
          <div class="text-center sm:text-left">
            <p class="text-sm text-gray-500">Member since</p>
            <p class="font-medium">{{ formatDate(user.created_at) }}</p>
          </div>
          <div class="text-center sm:text-left">
            <p class="text-sm text-gray-500">Total Sales</p>
            <p class="font-medium">{{ formatCurrency(user.total_sales) }}</p>
          </div>
          <div class="text-center sm:text-left">
            <p class="text-sm text-gray-500">Ratings</p>
            <p class="font-medium">{{ user.ratings.toFixed(1) }} ★</p>
          </div>
          <div class="text-center sm:text-left">
            <p class="text-sm text-gray-500">Reviews</p>
            <p class="font-medium">{{ user.reviews_count }}</p>
          </div>
        </div>

        <!-- Actions -->
        <div class="pt-4 border-t border-green-100 flex justify-center space-x-4">
          <router-link
            to="/dashboard"
            class="px-6 py-2 bg-amber-100 text-green-700 font-semibold rounded-lg hover:bg-amber-200 transition"
          >
            Back to Dashboard
          </router-link>
          <button
            @click="logout"
            class="px-6 py-2 bg-red-500 text-white font-semibold rounded-lg hover:bg-red-600 transition"
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

// redirect if unauthenticated
if (!auth.isAuthenticated) {
  router.replace({ name: 'Landing' })
}

const user = computed(() => auth.user || {})

// derive initials
const initials = computed(() => {
  const fn = user.value.first_name || ''
  const ln = user.value.last_name || ''
  if (fn || ln) return `${fn[0]||''}${ln[0]||''}`.toUpperCase()
  const [f,l] = (user.value.email || '').split('@')[0].split('.')
  return ((f?.[0]||'') + (l?.[0]||'')).toUpperCase()
})

// formatting helpers
function formatDate(ts) {
  return ts
    ? new Date(ts).toLocaleDateString(undefined, { year: 'numeric', month: 'short', day: 'numeric' })
    : '—'
}
function formatCurrency(val) {
  return typeof val === 'number'
    ? val.toLocaleString(undefined, { style: 'currency', currency: 'USD' })
    : '—'
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
