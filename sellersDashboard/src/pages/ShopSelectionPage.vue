<template>
  <div class="min-h-screen flex flex-col bg-amber-50 font-sans antialiased">
    <AppNavbar />

    <div class="flex-1 flex flex-col items-center pt-10 px-4 sm:px-8 pb-12">
      <div class="w-full max-w-4xl flex justify-end mb-6">
        <button
          @click="showCreateShopModal = true"
          class="inline-flex items-center px-6 py-3 bg-green-600 text-white text-lg font-bold rounded-full shadow-lg hover:bg-green-700 transition-all duration-300 transform hover:-translate-y-0.5 focus:outline-none focus:ring-4 focus:ring-green-300"
        >
          <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
          </svg>
          Add New Shop
        </button>
      </div>

      <div
        class="bg-white p-8 sm:p-12 rounded-3xl shadow-2xl w-full max-w-4xl border border-green-200 animate-fade-in-up"
      >
        <h2 class="text-4xl sm:text-5xl font-extrabold mb-4 text-center text-green-800 leading-tight">
          Your Shops
        </h2>
        <p class="text-center text-gray-700 mb-10 text-lg max-w-prose mx-auto">
          Manage your existing shops or effortlessly create a new one to expand your empire.
        </p>

        <div
          v-if="shops.length"
          class="grid grid-cols-1 md:grid-cols-2 gap-8 mb-6"
        >
          <div
            v-for="shop in shops"
            :key="shop.id"
            class="flex flex-col rounded-xl shadow-md bg-white border border-gray-200 transform transition-all duration-300 ease-in-out hover:scale-[1.02] hover:shadow-lg overflow-hidden group"
          >
            <div class="relative w-full h-48 bg-gray-100 flex items-center justify-center overflow-hidden">
              <img
                v-if="shop.image"
                :src="shop.image"
                :alt="`${shop.name} cover`"
                class="w-full h-full object-cover transition-transform duration-300 group-hover:scale-105"
              />
              <div v-else class="text-gray-400 text-lg font-medium">No image available</div>
            </div>
            <div class="p-6 flex-grow flex flex-col justify-between">
              <div>
                <h3 class="text-2xl font-bold text-green-800 mb-3 truncate" :title="shop.name">{{ shop.name }}</h3>
                <p class="text-base text-gray-700 mb-4 line-clamp-3">
                  {{ shop.description || 'No description provided for this shop.' }}
                </p>
              </div>
              <button
                @click="enterShop(shop)"
                class="mt-8 w-full py-3.5 bg-green-700 text-white rounded-lg font-semibold text-lg hover:bg-green-800 focus:outline-none focus:ring-4 focus:ring-green-300 focus:ring-offset-2 transition-all duration-300 transform hover:scale-[1.01]"
              >
                Enter Shop
              </button>
            </div>
          </div>
        </div>

        <div
          v-else
          class="text-center mb-12 p-10 bg-amber-50 rounded-2xl border border-amber-200 text-amber-800 shadow-md"
        >
          <svg class="mx-auto h-20 w-20 text-amber-500 mb-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 01-18 0 9 9 0 0118 0z"/>
          </svg>
          <h4 class="text-2xl font-semibold mb-3">No shops to show!</h4>
          <p class="text-gray-700 text-lg">It looks like you haven't created any shops yet. Click the button above to get started!</p>
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

import AppNavbar    from '@/components/AppNavbar.vue'
import AppFooter    from '@/components/AppFooter.vue'
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
@keyframes fade-in-up {
  from { opacity: 0; transform: translateY(20px); }
  to   { opacity: 1; transform: translateY(0); }
}
.animate-fade-in-up {
  animation: fade-in-up 0.6s ease-out forwards;
}
.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>