<template>
  <div class="min-h-screen flex flex-col items-center justify-center bg-gradient-to-br from-amber-50 to-orange-100 p-4 sm:p-6">
    <div class="bg-white p-8 sm:p-10 rounded-2xl shadow-xl w-full max-w-3xl border border-amber-100 animate-fade-in-up">
      <h2 class="text-4xl font-extrabold mb-4 text-center text-gray-900 leading-tight">Welcome Back!</h2>
      <p class="text-center text-gray-700 mb-8 text-lg">
        Ready to manage your empire? Choose a shop or create a new one.
      </p>

      <div v-if="shops.length" class="mb-10 p-6 bg-yellow-50 rounded-xl border border-yellow-200 shadow-inner">
        <h3 class="text-2xl font-semibold mb-6 text-orange-800 border-b border-orange-300 pb-3">Your Existing Shops</h3>
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-2 gap-6">
          <div
            v-for="shop in shops"
            :key="shop.id"
            class="flex flex-col overflow-hidden rounded-lg shadow-md bg-white border border-gray-200
                   transform transition-all duration-300 ease-in-out hover:scale-[1.02] hover:shadow-lg"
          >
            <div class="relative w-full h-40 bg-gray-100 flex items-center justify-center overflow-hidden">
              <img
                v-if="shop.image"
                :src="shop.image"
                :alt="`${shop.name} cover`"
                class="w-full h-full object-cover"
              />
              <div v-else class="text-gray-400 text-sm">No image available</div>
            </div>
            <div class="p-5 flex-grow flex flex-col justify-between">
              <div>
                <h4 class="text-xl font-bold text-gray-900 mb-2 truncate" :title="shop.name">{{ shop.name }}</h4>
                <p class="text-sm text-gray-700 mb-3 line-clamp-2">{{ shop.description || 'No description provided.' }}</p>
                <p class="text-xs text-gray-500 mt-2">
                  Created on <span class="font-medium">{{ formatDate(shop.createdAt) }}</span>
                </p>
              </div>
              <button
                @click="enterShop(shop)"
                class="mt-6 w-full py-3 bg-orange-600 text-white rounded-lg font-semibold text-base
                       hover:bg-orange-700 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:ring-offset-2
                       transition-colors duration-200 transform hover:scale-[1.01]"
              >
                Enter Shop
              </button>
            </div>
          </div>
        </div>
      </div>

      <div v-else class="text-center mb-10 p-8 bg-amber-50 rounded-xl border border-amber-200 text-amber-800 shadow-md">
        <svg class="mx-auto h-16 w-16 text-amber-500 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"></path>
        </svg>
        <p class="text-xl font-medium">
          Looks like you don't have any shops yet!
        </p>
        <p class="text-gray-700 mt-2">
          Let's create your first one below.
        </p>
      </div>

      <form @submit.prevent="createNewShop" class="space-y-6 p-8 bg-gradient-to-br from-green-50 to-green-100 rounded-2xl shadow-lg border border-green-200">
        <h3 class="text-3xl font-bold text-green-800 mb-4 text-center">Start a New Shop</h3>
        <div class="space-y-4">
          <input
            v-model="newShopName"
            placeholder="Awesome Shop Name (Required)"
            required
            class="w-full px-5 py-3 border border-gray-300 rounded-xl text-gray-800 placeholder-gray-500
                   focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200 shadow-sm"
          />
          <textarea
            v-model="newShopDesc"
            rows="3"
            placeholder="A short description of your shop (Optional)"
            class="w-full px-5 py-3 border border-gray-300 rounded-xl text-gray-800 placeholder-gray-500
                   focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200 shadow-sm"
          />
          <input
            v-model="newShopImage"
            placeholder="Image URL for your shop (Optional)"
            class="w-full px-5 py-3 border border-gray-300 rounded-xl text-gray-800 placeholder-gray-500
                   focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200 shadow-sm"
          />
        </div>
        <button
          type="submit"
          :disabled="creating || !String(newShopName.value || '').trim()"
          class="w-full bg-green-700 text-white py-3.5 rounded-xl font-bold text-lg
                 hover:bg-green-800 disabled:opacity-60 disabled:cursor-not-allowed
                 focus:outline-none focus:ring-2 focus::ring-green-500 focus:ring-offset-2
                 transition-all duration-200 transform hover:scale-[1.01] flex items-center justify-center"
        >
          <svg v-if="creating" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          {{ creating ? 'Creating Shop...' : 'Create Your Shop Now' }}
        </button>
      </form>
    </div>
  </div>
</template>

<script setup>
// (Script section remains unchanged as per your request)
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'

const router = useRouter()
const shopStore = useShopStore()

// form state for new shop
const newShopName = ref('')
const newShopDesc = ref('')
const newShopImage = ref('')
const creating = ref(false)

// existing shops are pulled from the store
const shops = computed(() => shopStore.allShops)

// Fetch shops when the component mounts
onMounted(async () => {
  try {
    await shopStore.fetchShops()
  } catch (err) {
    console.error('Failed to fetch shops:', err)
    // You might want to display a user-friendly error message here
  }
})

// Formats a date string into a more readable format
function formatDate(date) {
  return new Date(date).toLocaleDateString(undefined, {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
}

// Handles the creation of a new shop
async function createNewShop() {
  if (!String(newShopName.value || '').trim()) return // Basic validation, fixed trim issue

  creating.value = true
  try {
    const newShop = await shopStore.createShop({
      name: newShopName.value,
      description: newShopDesc.value,
      image: newShopImage.value
    })
    shopStore.setActiveShop(newShop) // Set the newly created shop as active
    router.push({ name: 'Dashboard' }) // Redirect to dashboard
  } catch (err) {
    console.error('Failed to create shop:', err)
    // Display error to user
    alert('Failed to create shop. Please try again.');
  } finally {
    creating.value = false
    // Clear form fields
    newShopName.value = ''
    newShopDesc.value = ''
    newShopImage.value = ''
  }
}

// Enters an existing shop
function enterShop(shop) {
  shopStore.setActiveShop(shop) // Set the selected shop as active
  router.push({ name: 'Dashboard' }) // Redirect to dashboard
}
</script>

<style scoped>
/* Keyframe animation for initial load */
@keyframes fade-in-up {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fade-in-up {
  animation: fade-in-up 0.6s ease-out forwards;
}

/* Custom styles for line-clamp */
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>