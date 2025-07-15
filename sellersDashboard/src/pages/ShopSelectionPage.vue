<template>
  <div class="min-h-screen flex flex-col bg-gray-50 font-sans antialiased">
    <AppNavbar />

    <div class="flex-1 flex flex-col items-center pt-10 px-4 sm:px-8 pb-12">
      <div class="w-full max-w-6xl flex flex-col sm:flex-row justify-between items-start sm:items-center mb-8">
        <div class="mb-6 sm:mb-0">
          <h1 class="text-4xl sm:text-5xl font-extrabold text-gray-800 mb-3 leading-tight">
            Your Shops
          </h1>
          <p class="text-lg text-gray-600 max-w-2xl">
            Manage your existing shops or effortlessly create a new one to expand your empire.
          </p>
        </div>

        <button
          @click="showCreateShopModal = true"
          class="inline-flex items-center px-8 py-4 bg-gradient-to-r from-green-600 to-green-700 text-white text-lg font-bold rounded-full shadow-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2 transition-all duration-200"
        >
          <svg class="w-6 h-6 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
          </svg>
          Add New Shop
        </button>
      </div>

      <div class="w-full max-w-6xl">
        <div
          v-if="shops.length"
          class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8"
        >
          <div
            v-for="shop in shops"
            :key="shop.id"
            class="group bg-white rounded-xl shadow-md border border-gray-100 overflow-hidden transform transition-all duration-200 ease-in-out hover:shadow-lg hover:border-green-300"
          >
            <div class="relative w-full h-40 bg-gradient-to-br from-gray-50 to-gray-100 flex items-center justify-center overflow-hidden">
              <img
                v-if="shop.image"
                :src="shop.image"
                :alt="`${shop.name} cover`"
                class="w-full h-full object-cover transition-transform duration-200 group-hover:scale-102"
              />
              <div v-else class="text-gray-400 text-base font-medium flex flex-col items-center">
                <svg class="w-14 h-14 mb-2 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                </svg>
                No image available
              </div>

              <div class="absolute top-3 right-3">
                <span class="inline-flex items-center px-2.5 py-1 rounded-full text-xs font-medium bg-green-100 text-green-800">
                  <div class="w-2 h-2 bg-green-400 rounded-full mr-1.5"></div>
                  Active
                </span>
              </div>
            </div>

            <div class="p-5">
              <div class="mb-4">
                <h3 class="text-xl font-bold text-gray-800 mb-1.5 truncate" :title="shop.name">
                  {{ shop.name }}
                </h3>
                <p class="text-gray-600 text-sm line-clamp-3 leading-relaxed">
                  {{ shop.description || 'No description provided for this shop.' }}
                </p>
              </div>

              <button
                @click="enterShop(shop)"
                class="w-full py-3 bg-gradient-to-r from-green-600 to-green-700 text-white rounded-lg font-semibold text-base hover:from-green-700 hover:to-green-800 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2 transition-all duration-200 shadow-sm hover:shadow-md"
              >
                <span class="flex items-center justify-center">
                  <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7l5 5m0 0l-5 5m5-5H6" />
                  </svg>
                  Enter Shop
                </span>
              </button>
            </div>
          </div>
        </div>

        <div
          v-else
          class="text-center py-16 px-8 bg-white rounded-2xl border border-dashed border-amber-200 shadow-md"
        >
          <div class="w-24 h-24 bg-gradient-to-br from-amber-100 to-amber-200 rounded-full flex items-center justify-center mx-auto mb-6">
            <svg class="w-12 h-12 text-amber-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
            </svg>
          </div>
          <h3 class="text-2xl font-bold text-gray-800 mb-3">No shops to show!</h3>
          <p class="text-gray-600 text-lg mb-8 max-w-md mx-auto">
            It looks like you haven't created any shops yet. Click the button above to get started and build your first online store!
          </p>
          <button
            @click="showCreateShopModal = true"
            class="inline-flex items-center px-8 py-4 bg-gradient-to-r from-amber-500 to-amber-600 text-white text-lg font-bold rounded-full shadow-lg hover:bg-amber-600 focus:outline-none focus:ring-2 focus:ring-amber-400 focus:ring-offset-2 transition-all duration-200"
          >
            <svg class="w-6 h-6 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
            Create Your First Shop
          </button>
        </div>
      </div>
    </div>

    <NewShopModal
      :is-visible="showCreateShopModal"
      :loading="creating"
      @close="showCreateShopModal = false"
      @create-shop="handleCreateShop"
    />

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
import NewShopModal from '@/components/NewShopModal.vue'

const router = useRouter()
const shopStore = useShopStore()
const auth = useAuthStore()

watchEffect(() => {
  if (!auth.isAuthenticated) {
    router.replace({ name: 'Landing' })
  }
})

const showCreateShopModal = ref(false)
const creating = ref(false)

const shops = computed(() => shopStore.list)

onMounted(async () => {
  if (auth.isAuthenticated) {
    try {
      await shopStore.fetchShops()
    } catch (e) {
      console.error('Failed to fetch shops:', e)
    }
  }
})

function formatDate(date) {
  return new Date(date).toLocaleDateString(undefined, {
    year: 'numeric', month: 'short', day: 'numeric'
  })
}

async function handleCreateShop(shopData) {
  creating.value = true
  try {
    const newShop = await shopStore.createShop(shopData)
    shopStore.setActiveShop(newShop)
    showCreateShopModal.value = false
    router.push({ name: 'Dashboard' })
  } catch (e) {
    console.error('Error creating shop:', e)
    alert('Could not create shop. Please try again.')
  } finally {
    creating.value = false
  }
}

function enterShop(shop) {
  shopStore.setActiveShop(shop)
  router.push({ name: 'Dashboard' })
}
</script>

<style scoped>
.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>