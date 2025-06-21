<template>
  <header class="bg-white shadow-md px-6 py-3 flex items-center justify-between text-gray-800 z-20">
    <div class="flex items-center space-x-4">
      <button
        class="md:hidden p-2 rounded-full hover:bg-gray-100 transition focus:outline-none focus:ring-2 focus:ring-gray-300"
        @click="$emit('toggle-sidebar')"
        aria-label="Toggle menu"
      >
        <MenuIcon class="h-6 w-6 text-gray-700" />
      </button>

      <div v-if="shops.length" class="relative" ref="shopDropdownRef">
        <button
          @click="showShopMenu = !showShopMenu"
          class="inline-flex items-center bg-white border border-gray-200 hover:bg-gray-50 px-4 py-2 rounded-full text-gray-700 font-medium shadow-sm transition"
        >
          <template v-if="activeShop?.image">
            <img
              :src="activeShop.image"
              alt="Shop logo"
              class="h-6 w-6 rounded-full object-cover mr-2"
            />
          </template>
          <template v-else>
            <ShopIcon class="h-5 w-5 mr-2 text-green-600" />
          </template>
          <span>{{ activeShop?.name || 'Select Shop' }}</span>
          <ChevronDownIcon class="h-4 w-4 ml-2 text-gray-500 transition-transform duration-200"
            :class="{ 'rotate-180': showShopMenu }"
          />
        </button>
        <ul
          v-if="showShopMenu"
          class="absolute mt-2 w-48 bg-white border border-gray-200 rounded-xl shadow-xl z-30 overflow-hidden animate-fade-in-down"
        >
          <li
            v-for="shop in shops"
            :key="shop.id"
            @click="selectShop(shop)"
            class="px-4 py-3 hover:bg-gray-50 cursor-pointer text-gray-800 transition flex items-center"
            :class="{ 'bg-gray-100 font-semibold': shop.id === activeShop?.id }"
          >
            <template v-if="shop.image">
              <img
                :src="shop.image"
                alt="Shop logo"
                class="h-5 w-5 rounded-full object-cover mr-2"
              />
            </template>
            <template v-else>
              <ShopIcon class="h-4 w-4 mr-2 text-gray-400" />
            </template>
            <span>{{ shop.name }}</span>
            <CheckIcon v-if="shop.id === activeShop?.id" class="h-5 w-5 text-green-500 ml-auto" />
          </li>
          <li class="border-t border-gray-100">
            <button
              @click="goToShopSelection"
              class="w-full text-left px-4 py-3 hover:bg-gray-50 text-green-600 font-medium transition"
            >
              Manage Shops…
            </button>
          </li>
        </ul>
      </div>
    </div>

    <div class="flex items-center space-x-4">
      <button
        class="relative p-2 rounded-full hover:bg-gray-100 transition focus:outline-none focus:ring-2 focus:ring-gray-300"
        aria-label="Notifications"
      >
        <BellIcon class="h-6 w-6 text-gray-700" />
        <span
          v-if="unreadCount"
          class="absolute -top-1 -right-1 bg-red-500 text-white text-xs font-bold w-5 h-5 rounded-full flex items-center justify-center ring-2 ring-white animate-bounce-short"
        >
          {{ unreadCount }}
        </span>
      </button>

      <div class="relative" ref="userDropdownRef">
        <button
          @click="showUserMenu = !showUserMenu"
          class="flex items-center space-x-2 focus:outline-none"
        >
          <template v-if="user?.profile_image">
            <img
              :src="user.profile_image"
              alt="Profile image"
              class="h-9 w-9 rounded-full object-cover shadow transition group-hover:ring-2 group-hover:ring-green-300 group-hover:ring-opacity-50"
            />
          </template>
          <template v-else>
            <div
              class="h-9 w-9 bg-green-500 rounded-full flex items-center justify-center text-white font-bold text-lg shadow transition group-hover:ring-2 group-hover:ring-green-300 group-hover:ring-opacity-50"
            >
              {{ userInitials }}
            </div>
          </template>
          <ChevronDownIcon
            class="h-4 w-4 text-gray-500 transition-transform duration-200"
            :class="{ 'rotate-180': showUserMenu }"
          />
        </button>
        <ul
          v-if="showUserMenu"
          class="absolute right-0 mt-2 w-48 bg-white border border-gray-200 rounded-xl shadow-xl z-30 overflow-hidden animate-fade-in-down"
        >
          <li class="px-4 py-3 text-gray-800 font-medium border-b border-gray-100">
            {{ user.email }}
          </li>
          <li>
            <router-link
              to="/profile"
              @click="showUserMenu = false"
              class="block px-4 py-3 hover:bg-gray-50 text-gray-700 transition"
            >
              Profile
            </router-link>
          </li>
          <li>
            <button
              @click="handleLogout"
              class="w-full text-left px-4 py-3 text-red-600 hover:bg-red-50 transition"
            >
              Logout
            </button>
          </li>
        </ul>
      </div>
    </div>
  </header>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue' // Import onMounted, onUnmounted
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import { useShopStore } from '@/store/shops'
import {
  MenuIcon,
  BellIcon,
  ChevronDownIcon,
  ShoppingBagIcon as ShopIcon,
  CheckIcon // Import CheckIcon for active shop indicator
} from '@heroicons/vue/outline'

const router = useRouter()
const auth = useAuthStore()
const shopStore = useShopStore()

// dropdown toggles
const showShopMenu = ref(false)
const showUserMenu = ref(false)

// Refs for dropdown elements to handle clicks outside
const shopDropdownRef = ref(null)
const userDropdownRef = ref(null)

const shops = computed(() => shopStore.list)
const activeShop = computed(() => shopStore.active)
const user = computed(() => auth.user || {})

// derive initials
const userInitials = computed(() => {
  const [first, last] = (user.value?.email || '').split('@')[0].split('.')
  return ((first?.[0] || '') + (last?.[0] || '')).toUpperCase() || '?'
})

// stub unread
const unreadCount = ref(3)

async function handleLogout() {
  await auth.logout()
  router.push({ name: 'Landing' })
  showUserMenu.value = false // Close menu on logout
}

function selectShop(shop) {
  shopStore.setActiveShop(shop)
  showShopMenu.value = false
  router.push({ name: 'Dashboard' })
}

function goToShopSelection() {
  showShopMenu.value = false // Close menu
  router.push({ name: 'ShopSelection' })
}

// Function to handle clicks outside dropdowns
const handleClickOutside = (event) => {
  if (shopDropdownRef.value && !shopDropdownRef.value.contains(event.target)) {
    showShopMenu.value = false
  }
  if (userDropdownRef.value && !userDropdownRef.value.contains(event.target)) {
    showUserMenu.value = false
  }
}

// Add and remove event listener
onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
@keyframes fade-in-down {
  from {
    opacity: 0;
    transform: translateY(-8px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
.animate-fade-in-down {
  animation: fade-in-down 0.2s ease-out forwards;
}
@keyframes bounce-short {
  0%,
  100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-2px);
  }
}
.animate-bounce-short {
  animation: bounce-short 0.8s infinite;
}
</style>