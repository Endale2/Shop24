<template>
  <div class="min-h-screen flex flex-col bg-emerald-50 font-sans antialiased">
    <header class="bg-white shadow-md px-6 py-4 flex items-center justify-between z-10">
      <h1 class="text-2xl font-bold text-emerald-800">Shop Manager</h1>
      <div class="flex items-center space-x-4">
        <div class="flex items-center space-x-3 group relative">
          <img
            v-if="user.profile_image"
            :src="user.profile_image"
            alt="Profile"
            class="h-10 w-10 rounded-full object-cover ring-2 ring-emerald-300 group-hover:ring-emerald-500 transition-all duration-200 cursor-pointer"
          />
          <div
            v-else
            class="h-10 w-10 bg-emerald-600 rounded-full flex items-center justify-center text-white font-bold text-lg ring-2 ring-emerald-300 group-hover:ring-emerald-500 transition-all duration-200 cursor-pointer"
          >
            {{ userInitials }}
          </div>
          <span class="text-gray-800 font-semibold hidden sm:inline">{{ user.name || user.email }}</span>

          <div class="absolute right-0 top-full mt-2 w-48 bg-white rounded-lg shadow-xl py-1 opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all duration-200 transform scale-95 group-hover:scale-100 origin-top-right">
            <router-link
              to="/profile"
              class="block px-4 py-2 text-gray-700 hover:bg-emerald-100 hover:text-emerald-800 transition-colors duration-150"
            >
              Your Profile
            </router-link>
            <button
              @click="handleLogout"
              class="w-full text-left px-4 py-2 text-red-600 hover:bg-red-50 hover:text-red-700 transition-colors duration-150"
            >
              Logout
            </button>
          </div>
        </div>
      </div>
    </header>

    <div class="flex-1 flex flex-col items-center justify-center p-4 sm:p-8">
      <div class="bg-white p-8 sm:p-12 rounded-3xl shadow-2xl w-full max-w-4xl border border-emerald-100 animate-fade-in-up">
        <h2 class="text-5xl font-extrabold mb-6 text-center text-emerald-900 leading-tight">
          Welcome to Your Shops!
        </h2>
        <p class="text-center text-gray-700 mb-10 text-xl max-w-prose mx-auto">
          Manage your existing shops or effortlessly create a new one to expand your empire.
        </p>

        <div v-if="shops.length" class="mb-12 p-8 bg-emerald-50 rounded-2xl border border-emerald-200 shadow-inner">
          <h3 class="text-3xl font-bold mb-8 text-emerald-800 border-b-2 border-emerald-300 pb-4 text-center">
            Your Current Shops
          </h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div
              v-for="shop in shops"
              :key="shop.id"
              class="flex flex-col rounded-xl shadow-lg bg-white border border-gray-200 transform transition-all duration-300 ease-in-out hover:scale-[1.03] hover:shadow-xl overflow-hidden group"
            >
              <div class="relative w-full h-48 bg-emerald-100 flex items-center justify-center overflow-hidden">
                <img
                  v-if="shop.image"
                  :src="shop.image"
                  :alt="`${shop.name} cover`"
                  class="w-full h-full object-cover transition-transform duration-300 group-hover:scale-105"
                />
                <div v-else class="text-emerald-400 text-lg font-medium">No image available</div>
              </div>
              <div class="p-6 flex-grow flex flex-col justify-between">
                <div>
                  <h4 class="text-2xl font-bold text-emerald-900 mb-3 truncate" :title="shop.name">{{ shop.name }}</h4>
                  <p class="text-base text-gray-700 mb-4 line-clamp-3">{{ shop.description || 'No description provided for this shop.' }}</p>
                  <p class="text-sm text-gray-500 mt-2">
                    <span class="font-medium">Created:</span> {{ formatDate(shop.createdAt) }}
                  </p>
                </div>
                <button
                  @click="enterShop(shop)"
                  class="mt-8 w-full py-3.5 bg-emerald-700 text-white rounded-lg font-semibold text-lg hover:bg-emerald-800 focus:outline-none focus:ring-4 focus:ring-emerald-300 focus:ring-offset-2 transition-all duration-300 transform hover:scale-[1.01]"
                >
                  Enter Shop
                </button>
              </div>
            </div>
          </div>
        </div>

        <div v-else class="text-center mb-12 p-10 bg-emerald-50 rounded-2xl border border-emerald-200 text-emerald-800 shadow-md">
          <svg class="mx-auto h-20 w-20 text-emerald-500 mb-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <p class="text-2xl font-semibold mb-3">No shops to show!</p>
          <p class="text-gray-700 text-lg">It looks like you haven't created any shops yet. Let's get started!</p>
        </div>

        <form @submit.prevent="createNewShop" class="space-y-8 p-10 bg-gradient-to-br from-emerald-100 to-emerald-200 rounded-3xl shadow-xl border border-emerald-300">
          <h3 class="text-4xl font-bold text-emerald-900 mb-6 text-center">Launch a New Shop</h3>
          <div class="space-y-5">
            <input
              v-model="newShopName"
              placeholder="Enter your dazzling shop name (Required)"
              required
              class="w-full px-6 py-3.5 border border-emerald-300 rounded-xl text-gray-800 placeholder-emerald-500 focus:ring-4 focus:ring-emerald-400 focus:border-emerald-400 transition-all duration-200 shadow-sm text-lg"
            />
            <textarea
              v-model="newShopDesc"
              rows="4"
              placeholder="Share a brief description of your shop (Optional, but recommended!)"
              class="w-full px-6 py-3.5 border border-emerald-300 rounded-xl text-gray-800 placeholder-emerald-500 focus:ring-4 focus:ring-emerald-400 focus:border-emerald-400 transition-all duration-200 shadow-sm text-lg resize-y"
            />
            <input
              v-model="newShopImage"
              placeholder="Link to an eye-catching image for your shop (Optional)"
              class="w-full px-6 py-3.5 border border-emerald-300 rounded-xl text-gray-800 placeholder-emerald-500 focus:ring-4 focus:ring-emerald-400 focus:border-emerald-400 transition-all duration-200 shadow-sm text-lg"
            />
          </div>
          <button
            type="submit"
            :disabled="creating || !String(newShopName).trim()"
            class="w-full bg-emerald-800 text-white py-4 rounded-xl font-bold text-xl hover:bg-emerald-900 disabled:opacity-60 disabled:cursor-not-allowed focus:outline-none focus:ring-4 focus:ring-emerald-500 focus:ring-offset-2 transition-all duration-300 transform hover:scale-[1.01] flex items-center justify-center shadow-lg"
          >
            <svg v-if="creating" class="animate-spin -ml-1 mr-4 h-6 w-6 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
            </svg>
            {{ creating ? 'Creating Your Shop...' : 'Create My Shop Now' }}
          </button>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { useAuthStore } from '@/store/auth'

const router = useRouter()
const shopStore = useShopStore()
const auth = useAuthStore()

// User data & initials
const user = computed(() => auth.user || {})
const userInitials = computed(() => {
  if (user.value.name) {
    return user.value.name
      .split(' ')
      .map(n => n[0])
      .join('')
      .toUpperCase()
  }
  const [first, last] = (user.value.email || '').split('@')[0].split('.')
  return ((first?.[0]||'') + (last?.[0]||'')).toUpperCase()
})

// Logout handler
async function handleLogout() {
  await auth.logout()
  router.push({ name: 'Landing' })
}

// Form state for new shop
const newShopName = ref('')
const newShopDesc = ref('')
const newShopImage = ref('')
const creating = ref(false)

// Existing shops are pulled from the store
const shops = computed(() => shopStore.list)

// Fetch shops when the component mounts
onMounted(async () => {
  try {
    await shopStore.fetchShops()
  } catch (err) {
    console.error('Failed to fetch shops:', err)
  }
})

// Formats a date string into a readable format
function formatDate(date) {
  return new Date(date).toLocaleDateString(undefined, {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
}

// Handles the creation of a new shop
async function createNewShop() {
  if (!String(newShopName.value || '').trim()) return

  creating.value = true
  try {
    const newShop = await shopStore.createShop({
      name: newShopName.value,
      description: newShopDesc.value,
      image: newShopImage.value
    })
    shopStore.setActiveShop(newShop)
    router.push({ name: 'Dashboard' })
  } catch (err) {
    console.error('Failed to create shop:', err)
    alert('Failed to create shop. Please try again.')
  } finally {
    creating.value = false
    newShopName.value = ''
    newShopDesc.value = ''
    newShopImage.value = ''
  }
}

// Enters an existing shop
function enterShop(shop) {
  shopStore.setActiveShop(shop)
  router.push({ name: 'Dashboard' })
}
</script>

<style scoped>
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
.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>