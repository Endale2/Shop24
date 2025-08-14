<template>
  <div class="min-h-screen bg-gray-50">
    <AppNavbar />

    <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
      <!-- Header -->
      <div class="text-center mb-12">
        <h1 class="text-3xl font-bold text-gray-900 mb-4">Your Shops</h1>
        <p class="text-lg text-gray-600 max-w-2xl mx-auto">
          Manage your online stores and start selling to customers worldwide
        </p>
      </div>

      <!-- Loading State -->
      <div v-if="shopStore.isLoading" class="text-center py-16">
        <div class="inline-flex items-center px-4 py-2 font-semibold leading-6 text-gray-600">
          <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-gray-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          Loading your shops...
        </div>
      </div>

      <!-- Error State -->
      <div v-else-if="shopStore.hasError" class="text-center py-16">
        <div class="bg-red-50 border border-red-200 rounded-lg p-8 max-w-md mx-auto">
          <div class="flex items-center justify-center mb-4">
            <svg class="w-12 h-12 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-800 mb-2">Failed to load shops</h3>
          <p class="text-gray-600 mb-6">{{ shopStore.errorMessage }}</p>
          <button
            @click="fetchShops"
            class="inline-flex items-center px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition-colors duration-200"
          >
            <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            Try Again
          </button>
        </div>
      </div>

      <!-- Content -->
      <div v-else>
        <!-- Create Shop Button -->
        <div class="text-center mb-8">
          <router-link
            to="/create-shop"
            class="inline-flex items-center px-8 py-4 bg-green-600 text-white text-lg font-medium rounded-lg shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition-colors duration-200"
          >
            <svg class="w-6 h-6 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
            Create New Shop
          </router-link>
        </div>

        <!-- Shops List -->
        <div v-if="shops.length" class="space-y-4">
          <div
            v-for="shop in shops"
            :key="shop.id"
            class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 hover:shadow-md transition-shadow duration-200"
          >
            <div class="flex items-center justify-between">
              <!-- Shop Info -->
              <div class="flex items-center space-x-4">
                <!-- Shop Logo/Image -->
                <div class="w-12 h-12 bg-gray-100 rounded-lg flex items-center justify-center">
                  <img
                    v-if="shop.image && isValidImageUrl(shop.image)"
                    :src="shop.image"
                    :alt="`${shop.name} logo`"
                    class="w-full h-full object-cover rounded-lg"
                    @error="handleImageError"
                  />
                  <svg v-else class="w-6 h-6 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                  </svg>
                </div>

                <!-- Shop Details -->
                <div>
                  <h3 class="text-lg font-semibold text-gray-900 mb-1">{{ shop.name }}</h3>
                  <p class="text-sm text-gray-600">{{ shop.description || 'No description provided.' }}</p>
                </div>
              </div>

              <!-- Status and Action -->
              <div class="flex items-center space-x-4">
                <!-- Status Badge -->
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                      :class="getStatusBadgeClass(shop.status)">
                  <div v-if="shop.status === 'active'" class="w-1.5 h-1.5 bg-green-400 rounded-full mr-1.5 animate-pulse"></div>
                  <div v-else-if="shop.status === 'inactive'" class="w-1.5 h-1.5 bg-gray-400 rounded-full mr-1.5"></div>
                  {{ shop.status || 'inactive' }}
                </span>

                <!-- Enter Shop Button -->
                <button
                  @click="enterShop(shop)"
                  class="inline-flex items-center px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition-colors duration-200"
                >
                  <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7l5 5m0 0l-5 5m5-5H6" />
                  </svg>
                  Enter Shop
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- Empty State -->
        <div
          v-else
          class="text-center py-16 px-8 bg-white rounded-lg border-2 border-dashed border-gray-300"
        >
          <div class="w-20 h-20 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-6">
            <svg class="w-10 h-10 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
            </svg>
          </div>
          <h3 class="text-xl font-semibold text-gray-900 mb-2">No shops yet!</h3>
          <p class="text-gray-600 mb-8 max-w-md mx-auto">
            Ready to start your e-commerce journey? Create your first shop and begin selling online today.
          </p>
          <router-link
            to="/create-shop"
            class="inline-flex items-center px-6 py-3 bg-green-600 text-white text-base font-medium rounded-lg shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition-colors duration-200"
          >
            <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
            Create Your First Shop
          </router-link>
        </div>
      </div>
    </div>

    <AppFooter />
  </div>
</template>

<script setup>
import { computed, onMounted, watchEffect } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { useAuthStore } from '@/store/auth'

import AppNavbar from '@/components/AppNavbar.vue'
import AppFooter from '@/components/AppFooter.vue'

const router = useRouter()
const shopStore = useShopStore()
const auth = useAuthStore()

watchEffect(() => {
  if (!auth.isAuthenticated) {
    router.replace({ name: 'Landing' })
  }
})

const shops = computed(() => shopStore.allShops)

// Utility functions
function getStatusBadgeClass(status) {
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'inactive':
      return 'bg-gray-100 text-gray-800'
    case 'suspended':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// Image handling functions
function isValidImageUrl(url) {
  if (!url) return false
  try {
    new URL(url)
    return true
  } catch {
    return false
  }
}

function handleImageError(event) {
  // Hide the broken image and show placeholder
  event.target.style.display = 'none'
  const parent = event.target.parentElement
  if (parent) {
    parent.innerHTML = `
      <svg class="w-6 h-6 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
      </svg>
    `
  }
}

// Actions
async function fetchShops() {
  if (!auth.isAuthenticated) return
  
  try {
    await shopStore.fetchShops()
  } catch (err) {
    console.error('Failed to fetch shops:', err)
    // Error is handled by the store
  }
}

function enterShop(shop) {
  console.log('[ShopSelection] Entering shop:', shop.name)

  try {
    shopStore.setActiveShop(shop)

    // Verify the shop was set correctly
    if (shopStore.activeShop?.id === shop.id) {
      console.log('[ShopSelection] Shop set successfully, navigating to dashboard')
      router.push({ name: 'Dashboard' })
    } else {
      console.error('[ShopSelection] Failed to set active shop')
      // Could show an error message to user here
    }
  } catch (error) {
    console.error('[ShopSelection] Error setting active shop:', error)
    // Could show an error message to user here
  }
}

onMounted(() => {
  if (auth.isAuthenticated) {
    fetchShops()
  }
})
</script>

<style scoped>
@keyframes pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: .5;
  }
}

.animate-pulse {
  animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}
</style>