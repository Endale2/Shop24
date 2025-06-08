<!-- src/components/Navbar.vue -->
<template>
  <header class="bg-white shadow-md px-6 py-3 flex items-center justify-between">
    <!-- Left: Hamburger + Shop Selector -->
    <div class="flex items-center space-x-4">
      <!-- Mobile sidebar toggle -->
      <button
        class="md:hidden p-1 rounded-lg hover:bg-gray-100 transition"
        @click="$emit('toggle-sidebar')"
        aria-label="Toggle menu"
      >
        <MenuIcon class="h-6 w-6 text-gray-700" />
      </button>

      <!-- Shop dropdown -->
      <div v-if="shops.length" class="relative">
        <button
          @click="showShopMenu = !showShopMenu"
          class="inline-flex items-center bg-gray-100 hover:bg-gray-200 px-3 py-1 rounded-full text-gray-700 font-medium transition"
        >
          <ShopIcon class="h-5 w-5 mr-1 text-green-600" />
          <span>{{ activeShop?.name || 'Select Shop' }}</span>
          <ChevronDownIcon class="h-4 w-4 ml-1 text-gray-500" />
        </button>
        <ul
          v-if="showShopMenu"
          @click.away="showShopMenu = false"
          class="absolute mt-2 w-48 bg-white border border-gray-200 rounded-lg shadow-lg z-10 overflow-hidden"
        >
          <li
            v-for="shop in shops"
            :key="shop.id"
            @click="selectShop(shop)"
            class="px-4 py-2 hover:bg-gray-100 cursor-pointer"
          >
            {{ shop.name }}
          </li>
          <li class="border-t">
            <button
              @click="goToShopSelection"
              class="w-full text-left px-4 py-2 hover:bg-gray-100 text-blue-600"
            >Manage Shops…</button>
          </li>
        </ul>
      </div>
    </div>

    <!-- Right: Notifications + User Menu -->
    <div class="flex items-center space-x-4">
      <!-- Notifications placeholder -->
      <button class="relative p-1 rounded-lg hover:bg-gray-100 transition" aria-label="Notifications">
        <BellIcon class="h-6 w-6 text-gray-700" />
        <span v-if="unreadCount" class="absolute -top-1 -right-1 bg-red-500 text-white text-xs font-bold w-4 h-4 rounded-full flex items-center justify-center">
          {{ unreadCount }}
        </span>
      </button>

      <!-- User avatar dropdown -->
      <div class="relative">
        <button
          @click="showUserMenu = !showUserMenu"
          class="flex items-center space-x-2 focus:outline-none"
        >
          <div class="h-8 w-8 bg-green-600 rounded-full flex items-center justify-center text-white font-semibold">
            {{ userInitials }}
          </div>
          <ChevronDownIcon class="h-4 w-4 text-gray-500" />
        </button>
        <ul
          v-if="showUserMenu"
          @click.away="showUserMenu = false"
          class="absolute right-0 mt-2 w-40 bg-white border border-gray-200 rounded-lg shadow-lg z-10 overflow-hidden"
        >
          <li class="px-4 py-2 text-gray-800 font-medium">{{ user.email }}</li>
          <li class="border-t">
            <router-link to="/profile" class="block px-4 py-2 hover:bg-gray-100">Profile</router-link>
          </li>
          <li>
            <button @click="handleLogout" class="w-full text-left px-4 py-2 text-red-600 hover:bg-gray-100">Logout</button>
          </li>
        </ul>
      </div>
    </div>
  </header>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import { useShopStore } from '@/store/shops'
import {
  MenuIcon,
  BellIcon,
  ChevronDownIcon,
  ShoppingBagIcon as ShopIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const auth = useAuthStore()
const shopStore = useShopStore()

// dropdown state
const showShopMenu = ref(false)
const showUserMenu = ref(false)

// reactive stores
const shops       = computed(() => shopStore.list)
const activeShop  = computed(() => shopStore.active)
const user        = computed(() => auth.user || {})

// derive initials
const userInitials = computed(() => {
  const [first, last] = (user.value?.email || '').split('@')[0].split('.')
  return ((first?.[0]||'') + (last?.[0]||'')).toUpperCase() || '?'
})

// unread notifications (stub)
const unreadCount = ref(3)

function handleLogout() {
  auth.logout()
  shopStore.clearActiveShop()
  router.push({ name: 'Login' })
}

function selectShop(shop) {
  shopStore.setActiveShop(shop)
  showShopMenu.value = false
  router.push({ name: 'Dashboard' })
}

function goToShopSelection() {
  router.push({ name: 'ShopSelection' })
}
</script>

<style scoped>
/* nothing extra—Tailwind covers it */
</style>
