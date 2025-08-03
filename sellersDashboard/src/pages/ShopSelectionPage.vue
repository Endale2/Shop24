<template>
  <div class="min-h-screen flex flex-col bg-gray-50 font-sans antialiased">
    <AppNavbar />

    <div class="flex-1">
      <!-- Header Section -->
      <div class="bg-white border-b border-gray-200">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
          <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center">
            <div class="mb-6 sm:mb-0">
              <h1 class="text-3xl font-bold text-gray-900 mb-2">Your Shops</h1>
              <p class="text-gray-600">Manage your online stores and track their performance</p>
            </div>

            <router-link
              to="/create-shop"
              class="inline-flex items-center px-6 py-3 bg-indigo-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 transition-colors duration-200"
            >
              <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
              </svg>
              Create New Shop
            </router-link>
          </div>

          <!-- Stats Cards -->
          <div class="grid grid-cols-1 sm:grid-cols-3 gap-6 mt-8">
            <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-lg p-6 text-white">
              <div class="flex items-center">
                <div class="flex-shrink-0">
                  <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                  </svg>
                </div>
                <div class="ml-4">
                  <p class="text-sm font-medium text-blue-100">Total Shops</p>
                  <p class="text-2xl font-bold">{{ shops.length }}</p>
                </div>
              </div>
            </div>
            
            <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-lg p-6 text-white">
              <div class="flex items-center">
                <div class="flex-shrink-0">
                  <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div class="ml-4">
                  <p class="text-sm font-medium text-green-100">Active Shops</p>
                  <p class="text-2xl font-bold">{{ activeShopsCount }}</p>
                </div>
              </div>
            </div>
            
            <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-lg p-6 text-white">
              <div class="flex items-center">
                <div class="flex-shrink-0">
                  <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1" />
                  </svg>
                </div>
                <div class="ml-4">
                  <p class="text-sm font-medium text-purple-100">Total Revenue</p>
                  <p class="text-2xl font-bold">${{ totalRevenue }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Content Section -->
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
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
          <div class="bg-red-50 border border-red-200 rounded-lg p-6 max-w-md mx-auto">
            <div class="flex items-center justify-center mb-4">
              <svg class="w-12 h-12 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-800 mb-2">Failed to load shops</h3>
            <p class="text-gray-600 mb-4">{{ shopStore.errorMessage }}</p>
            <button
              @click="fetchShops"
              class="inline-flex items-center px-4 py-2 bg-indigo-600 text-white text-sm font-medium rounded-lg hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 transition-colors duration-200"
            >
              <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
              </svg>
              Try Again
            </button>
          </div>
        </div>

        <!-- Shops Grid -->
        <div v-else>
          <div
            v-if="shops.length"
            class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6"
          >
            <div
              v-for="shop in shops"
              :key="shop.id"
              class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden hover:shadow-lg transition-shadow duration-300"
            >
              <!-- Shop Header with Image -->
              <div class="relative h-48 bg-gradient-to-br from-gray-100 to-gray-200">
                <!-- Shop Logo/Image -->
                <div class="absolute inset-0 flex items-center justify-center">
                  <img
                    v-if="shop.image && isValidImageUrl(shop.image)"
                    :src="shop.image"
                    :alt="`${shop.name} logo`"
                    class="w-full h-full object-cover"
                    @error="handleImageError"
                  />
                  <div v-else class="flex flex-col items-center justify-center text-gray-400">
                    <svg class="w-16 h-16 mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                    </svg>
                    <span class="text-sm font-medium">No logo</span>
                  </div>
                </div>

                <!-- Status Badge -->
                <div class="absolute top-3 right-3">
                  <span class="inline-flex items-center px-2.5 py-1 rounded-full text-xs font-medium shadow-sm"
                        :class="getStatusBadgeClass(shop.status)">
                    <div v-if="shop.status === 'active'" class="w-2 h-2 bg-green-400 rounded-full mr-1.5 animate-pulse"></div>
                    {{ shop.status || 'active' }}
                  </span>
                </div>

                <!-- Quick Actions -->
                <div class="absolute top-3 left-3 opacity-0 group-hover:opacity-100 transition-opacity duration-200">
                  <div class="flex space-x-1">
                    <button 
                      @click="editShop(shop)"
                      class="p-1.5 bg-white rounded-md shadow-sm hover:bg-gray-50 transition-colors duration-200"
                      title="Edit Shop"
                    >
                      <svg class="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                      </svg>
                    </button>
                    <button 
                      @click="viewShopDetails(shop)"
                      class="p-1.5 bg-white rounded-md shadow-sm hover:bg-gray-50 transition-colors duration-200"
                      title="View Details"
                    >
                      <svg class="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                      </svg>
                    </button>
                  </div>
                </div>
              </div>

              <!-- Shop Info -->
              <div class="p-6">
                <!-- Shop Name and Category -->
                <div class="mb-4">
                  <h3 class="text-lg font-semibold text-gray-900 mb-1 truncate" :title="shop.name">
                    {{ shop.name }}
                  </h3>
                  <div class="flex items-center justify-between">
                    <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                      {{ shop.category || 'Uncategorized' }}
                    </span>
                    <span class="text-xs text-gray-500">
                      {{ formatDate(shop.createdAt) }}
                    </span>
                  </div>
                </div>

                <!-- Description -->
                <p class="text-gray-600 text-sm mb-4 line-clamp-2">
                  {{ shop.description || 'No description provided for this shop.' }}
                </p>

                <!-- Stats Grid -->
                <div class="grid grid-cols-2 gap-4 mb-4">
                  <div class="text-center p-3 bg-gray-50 rounded-lg">
                    <div class="text-lg font-semibold text-gray-900">{{ shop.customer_count || 0 }}</div>
                    <div class="text-xs text-gray-500">Customers</div>
                  </div>
                  <div class="text-center p-3 bg-gray-50 rounded-lg">
                    <div class="text-lg font-semibold text-gray-900">{{ shop.product_count || 0 }}</div>
                    <div class="text-xs text-gray-500">Products</div>
                  </div>
                </div>

                <!-- Contact Info -->
                <div class="space-y-2 mb-4">
                  <div v-if="shop.email" class="flex items-center text-xs text-gray-500">
                    <svg class="w-3 h-3 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207" />
                    </svg>
                    {{ shop.email }}
                  </div>
                  <div v-if="shop.currency" class="flex items-center text-xs text-gray-500">
                    <svg class="w-3 h-3 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1" />
                    </svg>
                    {{ shop.currency }}
                  </div>
                </div>

                <!-- Action Buttons -->
                <div class="flex space-x-2">
                  <button
                    @click="enterShop(shop)"
                    class="flex-1 py-2.5 bg-indigo-600 text-white rounded-lg font-medium text-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 transition-colors duration-200"
                  >
                    <span class="flex items-center justify-center">
                      <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7l5 5m0 0l-5 5m5-5H6" />
                      </svg>
                      Enter Shop
                    </span>
                  </button>
                  
                  <button
                    @click="viewShopDetails(shop)"
                    class="px-3 py-2.5 bg-gray-100 text-gray-700 rounded-lg font-medium text-sm hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 transition-colors duration-200"
                    title="View Details"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                    </svg>
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
              class="inline-flex items-center px-6 py-3 bg-indigo-600 text-white text-base font-medium rounded-lg shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 transition-colors duration-200"
            >
              <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
              </svg>
              Create Your First Shop
            </router-link>
          </div>
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

// Computed stats using store getters
const activeShopsCount = computed(() => shopStore.activeShops.length)
const totalRevenue = computed(() => {
  return shopStore.totalRevenue.toLocaleString()
})

// Utility functions
function formatDate(dateString) {
  if (!dateString) return 'Unknown'
  const date = new Date(dateString)
  return date.toLocaleDateString(undefined, { 
    year: 'numeric', 
    month: 'short', 
    day: 'numeric' 
  })
}

function getStatusBadgeClass(status) {
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'inactive':
      return 'bg-gray-100 text-gray-800'
    case 'suspended':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-green-100 text-green-800'
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
      <div class="flex flex-col items-center justify-center text-gray-400">
        <svg class="w-16 h-16 mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
        </svg>
        <span class="text-sm font-medium">No logo</span>
      </div>
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
  shopStore.setActiveShop(shop)
  router.push({ name: 'Dashboard' })
}

function viewShopDetails(shop) {
  // Navigate to shop details page (to be implemented)
  console.log('View shop details:', shop.id)
  // For now, just show an alert
  alert(`Shop details for: ${shop.name}\n\nThis feature is coming soon!`)
}

function editShop(shop) {
  // Navigate to shop edit page (to be implemented)
  console.log('Edit shop:', shop.id)
  // For now, just show an alert
  alert(`Edit shop: ${shop.name}\n\nThis feature is coming soon!`)
}

onMounted(() => {
  if (auth.isAuthenticated) {
    fetchShops()
  }
})
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
}

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



/* Group hover for quick actions */
.group:hover .opacity-0 {
  opacity: 1;
}
</style>