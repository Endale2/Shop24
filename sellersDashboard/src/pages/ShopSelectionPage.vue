<template>
  <div class="min-h-screen flex flex-col bg-gradient-to-br from-gray-50 to-gray-100 font-sans antialiased">
    <AppNavbar />

    <div class="flex-1 flex flex-col items-center pt-8 px-4 sm:px-6 pb-8">
      <!-- Header Section -->
      <div class="w-full max-w-7xl">
        <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center mb-8">
          <div class="mb-6 sm:mb-0">
            <div class="flex items-center mb-3">
              <div class="w-10 h-10 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center mr-3 shadow-sm">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                </svg>
              </div>
              <h1 class="text-3xl sm:text-4xl font-bold text-gray-800 leading-tight">
                Your Shops
              </h1>
            </div>
            <p class="text-base text-gray-600 max-w-2xl ml-13">
              Manage your existing shops or create a new one to expand your business empire.
            </p>
          </div>

          <router-link
            to="/create-shop"
            class="inline-flex items-center px-6 py-3 bg-gradient-to-r from-green-600 to-green-700 text-white text-sm font-medium rounded-lg shadow-sm hover:from-green-700 hover:to-green-800 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2 transition-all duration-200 transform hover:-translate-y-0.5"
          >
            <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
            Create New Shop
          </router-link>
        </div>

        <!-- Stats Section -->
        <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 mb-8">
          <div class="bg-white rounded-lg p-4 shadow-sm border border-gray-200">
            <div class="flex items-center">
              <div class="w-8 h-8 bg-green-100 rounded-lg flex items-center justify-center mr-3">
                <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                </svg>
              </div>
              <div>
                <p class="text-sm text-gray-500 font-medium">Total Shops</p>
                <p class="text-2xl font-bold text-gray-800">{{ shops.length }}</p>
              </div>
            </div>
          </div>
          
          <div class="bg-white rounded-lg p-4 shadow-sm border border-gray-200">
            <div class="flex items-center">
              <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center mr-3">
                <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
                </svg>
              </div>
              <div>
                <p class="text-sm text-gray-500 font-medium">Active Shops</p>
                <p class="text-2xl font-bold text-gray-800">{{ activeShopsCount }}</p>
              </div>
            </div>
          </div>
          
          <div class="bg-white rounded-lg p-4 shadow-sm border border-gray-200">
            <div class="flex items-center">
              <div class="w-8 h-8 bg-amber-100 rounded-lg flex items-center justify-center mr-3">
                <svg class="w-4 h-4 text-amber-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1" />
                </svg>
              </div>
              <div>
                <p class="text-sm text-gray-500 font-medium">Total Revenue</p>
                <p class="text-2xl font-bold text-gray-800">${{ totalRevenue }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Shops Grid -->
      <div class="w-full max-w-7xl">
        <div
          v-if="shops.length"
          class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6"
        >
          <div
            v-for="shop in shops"
            :key="shop.id"
            class="group bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden transform transition-all duration-300 ease-in-out hover:shadow-lg hover:border-green-300 hover:-translate-y-1"
          >
            <!-- Shop Image -->
            <div class="relative w-full h-40 bg-gradient-to-br from-gray-50 to-gray-100 flex items-center justify-center overflow-hidden">
              <img
                v-if="shop.image"
                :src="shop.image"
                :alt="`${shop.name} cover`"
                class="w-full h-full object-cover transition-transform duration-300 group-hover:scale-110"
              />
              <div v-else class="text-gray-400 text-sm font-medium flex flex-col items-center">
                <svg class="w-12 h-12 mb-2 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                </svg>
                No image available
              </div>

              <!-- Status Badge -->
              <div class="absolute top-3 right-3">
                <span class="inline-flex items-center px-2.5 py-1 rounded-full text-xs font-medium bg-green-100 text-green-800 shadow-sm">
                  <div class="w-2 h-2 bg-green-400 rounded-full mr-1.5 animate-pulse"></div>
                  Active
                </span>
              </div>

              <!-- Quick Actions Overlay -->
              <div class="absolute inset-0 bg-black bg-opacity-0 group-hover:bg-opacity-20 transition-all duration-300 flex items-center justify-center">
                <div class="opacity-0 group-hover:opacity-100 transition-opacity duration-300 flex space-x-2">
                  <button class="p-2 bg-white rounded-lg shadow-lg hover:bg-gray-50 transition-colors duration-200">
                    <svg class="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                    </svg>
                  </button>
                  <button class="p-2 bg-white rounded-lg shadow-lg hover:bg-gray-50 transition-colors duration-200">
                    <svg class="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                    </svg>
                  </button>
                </div>
              </div>
            </div>

            <!-- Shop Info -->
            <div class="p-5">
              <div class="mb-4">
                <h3 class="text-lg font-semibold text-gray-800 mb-2 truncate" :title="shop.name">
                  {{ shop.name }}
                </h3>
                <p class="text-gray-600 text-sm line-clamp-2 leading-relaxed mb-3">
                  {{ shop.description || 'No description provided for this shop.' }}
                </p>
                
                <!-- Shop Stats -->
                <div class="flex items-center justify-between text-xs text-gray-500 mb-4">
                  <span class="flex items-center">
                    <svg class="w-3 h-3 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                    </svg>
                    {{ shop.customer_count || 0 }} customers
                  </span>
                  <span class="flex items-center">
                    <svg class="w-3 h-3 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                    </svg>
                    {{ shop.product_count || 0 }} products
                  </span>
                </div>
              </div>

              <!-- Action Buttons -->
              <div class="flex space-x-2">
                <button
                  @click="enterShop(shop)"
                  class="flex-1 py-2.5 bg-gradient-to-r from-green-600 to-green-700 text-white rounded-lg font-medium text-sm hover:from-green-700 hover:to-green-800 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2 transition-all duration-200 shadow-sm hover:shadow-md"
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
                  class="px-3 py-2.5 bg-gray-100 text-gray-700 rounded-lg font-medium text-sm hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-offset-2 transition-all duration-200"
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
          class="text-center py-16 px-8 bg-white rounded-2xl border border-dashed border-amber-200 shadow-sm"
        >
          <div class="w-20 h-20 bg-gradient-to-br from-amber-100 to-amber-200 rounded-full flex items-center justify-center mx-auto mb-6">
            <svg class="w-10 h-10 text-amber-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
            </svg>
          </div>
          <h3 class="text-2xl font-bold text-gray-800 mb-3">No shops yet!</h3>
          <p class="text-gray-600 text-lg mb-8 max-w-md mx-auto">
            Ready to start your e-commerce journey? Create your first shop and begin selling online today.
          </p>
          <router-link
            to="/create-shop"
            class="inline-flex items-center px-8 py-4 bg-gradient-to-r from-amber-500 to-amber-600 text-white text-lg font-medium rounded-lg shadow-lg hover:from-amber-600 hover:to-amber-700 focus:outline-none focus:ring-2 focus:ring-amber-400 focus:ring-offset-2 transition-all duration-200 transform hover:-translate-y-0.5"
          >
            <svg class="w-6 h-6 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
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
import { ref, computed, onMounted, watchEffect } from 'vue'
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

const shops = computed(() => shopStore.list)

// Computed stats
const activeShopsCount = computed(() => shops.value.length)
const totalRevenue = computed(() => {
  // Mock calculation - replace with actual data
  return shops.value.reduce((total, shop) => total + (shop.revenue || 0), 0).toLocaleString()
})

onMounted(async () => {
  if (auth.isAuthenticated) {
    try {
      await shopStore.fetchShops()
    } catch (e) {
      console.error('Failed to fetch shops:', e)
    }
  }
})

function enterShop(shop) {
  shopStore.setActiveShop(shop)
  router.push({ name: 'Dashboard' })
}

function viewShopDetails(shop) {
  // Navigate to shop details page (to be implemented)
  console.log('View shop details:', shop.id)
}
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
</style>