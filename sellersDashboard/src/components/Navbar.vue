<template>
  <header class="bg-white shadow-sm border-b border-gray-200 px-6 py-3 flex items-center justify-between text-gray-800 z-20 sticky top-0 backdrop-blur-sm bg-white/95">
    <div class="flex items-center space-x-4">
      <button
        class="md:hidden p-2 rounded-lg hover:bg-gray-100 transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-green-300 focus:ring-offset-2"
        @click="$emit('toggle-sidebar')"
        aria-label="Toggle menu"
      >
        <MenuIcon class="h-5 w-5 text-gray-700" />
      </button>

      <div v-if="shops.length" class="relative" ref="shopDropdownRef">
        <button
          @click="showShopMenu = !showShopMenu"
          class="inline-flex items-center bg-gradient-to-r from-green-50 to-emerald-50 border border-green-200 hover:from-green-100 hover:to-emerald-100 px-4 py-2 rounded-lg text-gray-700 font-medium shadow-sm transition-all duration-200 hover:shadow-md group"
        >
          <template v-if="activeShop?.image">
            <img
              :src="activeShop.image"
              alt="Shop logo"
              class="h-6 w-6 rounded-lg object-cover mr-2 ring-2 ring-white shadow-sm"
            />
          </template>
          <template v-else>
            <div class="h-6 w-6 bg-gradient-to-br from-green-500 to-emerald-600 rounded-lg flex items-center justify-center mr-2 shadow-sm">
              <ShopIcon class="h-4 w-4 text-white" />
            </div>
          </template>
          <span class="text-sm text-gray-800">{{ activeShop?.name || 'Select Shop' }}</span>
          <ChevronDownIcon class="h-4 w-4 ml-2 text-gray-500 transition-transform duration-200 group-hover:text-gray-700"
            :class="{ 'rotate-180': showShopMenu }"
          />
        </button>
        <ul
          v-if="showShopMenu"
          class="absolute mt-2 w-56 bg-white border border-gray-200 rounded-xl shadow-xl z-30 overflow-hidden animate-fade-in-down"
        >
          <li class="px-4 py-2 bg-gray-50 border-b border-gray-100">
            <p class="text-xs font-semibold text-gray-500 uppercase tracking-wide">Your Shops</p>
          </li>
          <li
            v-for="shop in otherShops"
            :key="shop.id"
            @click="selectShop(shop)"
            class="px-4 py-2 hover:bg-gray-50 cursor-pointer text-gray-800 transition-all duration-200 flex items-center group"
          >
            <template v-if="shop.image">
              <img
                :src="shop.image"
                alt="Shop logo"
                class="h-6 w-6 rounded-lg object-cover mr-3 shadow-sm"
              />
            </template>
            <template v-else>
              <div class="h-6 w-6 bg-gradient-to-br from-gray-400 to-gray-500 rounded-lg flex items-center justify-center mr-3 shadow-sm">
                <ShopIcon class="h-4 w-4 text-white" />
              </div>
            </template>
            <div class="flex-1">
              <span class="text-sm font-medium">{{ shop.name }}</span>
            </div>
          </li>
          <li class="border-t border-gray-100">
            <button
              @click="goToShopSelection"
              class="w-full text-left px-4 py-2 hover:bg-green-50 text-green-600 font-medium transition-all duration-200 flex items-center group"
            >
              <ShopIcon class="h-4 w-4 mr-2 group-hover:scale-110 transition-transform duration-200" />
              <span class="text-sm">All Shops</span>
            </button>
          </li>
        </ul>
      </div>
    </div>

    <div class="flex items-center space-x-2">
      <!-- Search Bar -->
      <div class="hidden md:block relative">
        <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
          <SearchIcon class="h-4 w-4 text-gray-400" />
        </div>
        <input
          type="text"
          placeholder="Search..."
          class="w-56 pl-9 pr-4 py-2 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-300 focus:border-transparent transition-all duration-200 bg-gray-50 hover:bg-white focus:bg-white text-sm"
        />
      </div>

      <!-- Notifications -->
      <button
        class="relative p-2 rounded-lg hover:bg-gray-100 transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-green-300 focus:ring-offset-2 group"
        aria-label="Notifications"
      >
        <BellIcon class="h-5 w-5 text-gray-700 group-hover:text-gray-900 transition-colors duration-200" />
        <span
          v-if="unreadCount"
          class="absolute -top-1 -right-1 bg-gradient-to-r from-red-500 to-pink-500 text-white text-xs font-bold w-4 h-4 rounded-full flex items-center justify-center ring-2 ring-white animate-bounce-short shadow-lg"
        >
          {{ unreadCount }}
        </span>
      </button>

      <!-- Help/Support -->
      <button
        class="p-2 rounded-lg hover:bg-gray-100 transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-green-300 focus:ring-offset-2 group"
        aria-label="Help"
      >
        <QuestionMarkCircleIcon class="h-5 w-5 text-gray-700 group-hover:text-gray-900 transition-colors duration-200" />
      </button>

      <!-- User Menu -->
      <div class="relative" ref="userDropdownRef">
        <button
          @click="showUserMenu = !showUserMenu"
          class="flex items-center space-x-2 p-2 rounded-lg hover:bg-gray-100 transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-green-300 focus:ring-offset-2 group"
        >
          <template v-if="user?.profile_image">
            <img
              :src="user.profile_image"
              alt="Profile image"
              class="h-8 w-8 rounded-lg object-cover shadow-md transition-all duration-200 group-hover:ring-2 group-hover:ring-green-300 group-hover:ring-opacity-50 group-hover:scale-105"
            />
          </template>
          <template v-else>
            <div
              class="h-8 w-8 bg-gradient-to-br from-green-500 to-emerald-600 rounded-lg flex items-center justify-center text-white font-bold text-sm shadow-md transition-all duration-200 group-hover:ring-2 group-hover:ring-green-300 group-hover:ring-opacity-50 group-hover:scale-105"
            >
              {{ userInitials }}
            </div>
          </template>
          <div class="hidden md:block text-left">
            <p class="text-sm font-medium text-gray-800">{{ user.email?.split('@')[0] || 'User' }}</p>
            <p class="text-xs text-gray-500">Seller</p>
          </div>
          <ChevronDownIcon
            class="h-4 w-4 text-gray-500 transition-transform duration-200 group-hover:text-gray-700"
            :class="{ 'rotate-180': showUserMenu }"
          />
        </button>
        <ul
          v-if="showUserMenu"
          class="absolute right-0 mt-2 w-56 bg-white border border-gray-200 rounded-xl shadow-xl z-30 overflow-hidden animate-fade-in-down"
        >
          <li class="px-4 py-3 bg-gradient-to-r from-gray-50 to-gray-100 border-b border-gray-200">
            <p class="text-sm font-medium text-gray-800">{{ user.email }}</p>
            <p class="text-xs text-gray-500 mt-1">Seller Account</p>
          </li>
          <li>
            <router-link
              to="/profile"
              @click="showUserMenu = false"
              class="block px-4 py-2 hover:bg-gray-50 text-gray-700 transition-all duration-200 flex items-center group"
            >
              <UserIcon class="h-4 w-4 mr-3 text-gray-400 group-hover:text-gray-600" />
              <span class="text-sm">Profile Settings</span>
            </router-link>
          </li>
          <li>
            <router-link
              to="/settings"
              @click="showUserMenu = false"
              class="block px-4 py-2 hover:bg-gray-50 text-gray-700 transition-all duration-200 flex items-center group"
            >
              <CogIcon class="h-4 w-4 mr-3 text-gray-400 group-hover:text-gray-600" />
              <span class="text-sm">Account Settings</span>
            </router-link>
          </li>
          <li class="border-t border-gray-100">
            <button
              @click="handleLogout"
              class="w-full text-left px-4 py-2 text-red-600 hover:bg-red-50 transition-all duration-200 flex items-center group"
            >
              <LogoutIcon class="h-4 w-4 mr-3 group-hover:scale-110 transition-transform duration-200" />
              <span class="text-sm">Sign Out</span>
            </button>
          </li>
        </ul>
      </div>
    </div>
  </header>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import { useShopStore } from '@/store/shops'
import {
  MenuIcon,
  BellIcon,
  ChevronDownIcon,
  ShoppingBagIcon as ShopIcon,
  SearchIcon,
  QuestionMarkCircleIcon,
  UserIcon,
  CogIcon,
  LogoutIcon
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

// Filter out the current active shop from the dropdown list
const otherShops = computed(() => {
  return shops.value.filter(shop => shop.id !== activeShop.value?.id)
})

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
    transform: translateY(-12px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
.animate-fade-in-down {
  animation: fade-in-down 0.3s ease-out forwards;
}
@keyframes bounce-short {
  0%,
  100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-3px);
  }
}
.animate-bounce-short {
  animation: bounce-short 1s infinite;
}

/* Custom scrollbar for dropdowns */
ul::-webkit-scrollbar {
  width: 4px;
}

ul::-webkit-scrollbar-track {
  background: transparent;
}

ul::-webkit-scrollbar-thumb {
  background: #d1d5db;
  border-radius: 2px;
}

ul::-webkit-scrollbar-thumb:hover {
  background: #9ca3af;
}
</style>