<!-- src/pages/ShopSelectionPage.vue -->
<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 p-4">
    <div class="bg-white p-8 rounded-xl shadow-lg w-full max-w-2xl border border-gray-100">
      <h2 class="text-4xl font-extrabold mb-8 text-center text-gray-800">Welcome Back!</h2>
      <p class="text-center text-gray-600 mb-6">Select an existing shop or create a new one to get started.</p>

      <!-- Existing Shops -->
      <div v-if="shops.length" class="mb-8">
        <h3 class="text-2xl font-semibold mb-4 text-gray-700 border-b pb-2">Your Shops</h3>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
          <div
            v-for="shop in shops"
            :key="shop.id"
            class="flex flex-col justify-between p-6 bg-gray-50 rounded-lg border border-gray-200 hover:shadow-md transition-shadow"
          >
            <div>
              <h4 class="text-lg font-medium text-gray-800 truncate">{{ shop.name }}</h4>
              <p class="text-sm text-gray-500 mt-1">Created on {{ formatDate(shop.createdAt) }}</p>
            </div>
            <button
              @click="selectShop(shop)"
              class="mt-4 bg-green-600 text-white py-2 px-4 rounded-full font-semibold hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-300 focus:ring-opacity-75 transition-transform transform hover:scale-105 text-sm"
            >
              Enter Shop
            </button>
          </div>
        </div>
      </div>

      <!-- No Shops -->
      <div v-else class="text-center mb-8 p-6 bg-yellow-50 rounded-md border border-yellow-200">
        <p class="text-yellow-800 text-lg">No shops found. Let's create your first one!</p>
      </div>

      <!-- Create New Shop -->
      <form @submit.prevent="createNewShop" class="space-y-4 p-6 bg-white rounded-lg border border-green-200">
        <h3 class="text-2xl font-semibold text-green-800">Create a New Shop</h3>
        <input
          v-model="newShopName"
          placeholder="Shop Name"
          required
          class="w-full px-4 py-3 border border-green-300 rounded-md text-gray-800 placeholder-gray-500 focus:ring-2 focus:ring-green-400 focus:border-transparent transition duration-150 ease-in-out"
        />
        <button
          type="submit"
          :disabled="creating || !newShopName.trim()"
          class="w-full bg-green-700 text-white py-3 rounded-md font-bold text-lg hover:bg-green-800 disabled:opacity-60 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-opacity-75 transition-transform transform hover:scale-105"
        >
          {{ creating ? 'Creating...' : 'Create Shop' }}
        </button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'

const router = useRouter()
const shopStore = useShopStore()

// local state
const newShopName = ref('')
const creating = ref(false)

// derived state
const shops = computed(() => shopStore.allShops)

// fetch existing shops on mount
onMounted(async () => {
  try {
    await shopStore.fetchShops()
  } catch (err) {
    console.error('Failed to fetch shops:', err)
  }
})

// helpers
function formatDate(date) {
  return new Date(date).toLocaleDateString(undefined, {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
  })
}

// actions
async function createNewShop() {
  if (!newShopName.value.trim()) return
  creating.value = true
  try {
    const newShop = await shopStore.createShop({ name: newShopName.value, description: '' })
    shopStore.setActiveShop(newShop)
    router.push({ name: 'Dashboard' })
  } catch (err) {
    console.error('Failed to create shop:', err)
  } finally {
    creating.value = false
    newShopName.value = ''
  }
}

function selectShop(shop) {
  shopStore.setActiveShop(shop)
  router.push({ name: 'Dashboard' })
}

</script>

<style scoped>
.hover\:shadow-md:hover {
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}
</style>
