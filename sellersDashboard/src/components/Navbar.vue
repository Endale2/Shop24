<template>
  <header class="bg-gradient-to-r from-gray-100 to-gray-200 shadow-lg px-6 py-3 flex items-center justify-between text-gray-800 relative z-20">
    <div class="flex items-center space-x-4">
      <button
        class="md:hidden p-2 rounded-full hover:bg-gray-300 transition duration-300 ease-in-out focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-opacity-50"
        @click="$emit('toggle-sidebar')"
        aria-label="Toggle menu"
      >
        <MenuIcon class="h-6 w-6 text-gray-700" />
      </button>

      <div v-if="shops.length" class="relative">
        <button
          @click="showShopMenu = !showShopMenu"
          class="inline-flex items-center bg-white border border-gray-300 hover:bg-gray-100 px-4 py-2 rounded-full text-gray-700 font-semibold shadow-md transition duration-300 ease-in-out focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-opacity-50"
        >
          <ShopIcon class="h-5 w-5 mr-2 text-green-600" />
          <span>{{ activeShop?.name || 'Select Shop' }}</span>
          <ChevronDownIcon class="h-4 w-4 ml-2 text-gray-500" />
        </button>
        <ul
          v-if="showShopMenu"
          @click.away="showShopMenu = false"
          class="absolute mt-3 w-48 bg-white border border-gray-200 rounded-xl shadow-xl z-30 overflow-hidden animate-fade-in-down"
        >
          <li
            v-for="shop in shops"
            :key="shop.id"
            @click="selectShop(shop)"
            class="px-4 py-3 hover:bg-gray-100 cursor-pointer text-gray-800 transition duration-200 ease-in-out"
          >
            {{ shop.name }}
          </li>
          <li class="border-t border-gray-100">
            <button
              @click="goToShopSelection"
              class="w-full text-left px-4 py-3 hover:bg-gray-100 text-green-600 font-medium transition duration-200 ease-in-out"
            >Manage Shops…</button>
          </li>
        </ul>
      </div>
    </div>

    <div class="flex items-center space-x-4">
      <button class="relative p-2 rounded-full hover:bg-gray-300 transition duration-300 ease-in-out focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-opacity-50" aria-label="Notifications">
        <BellIcon class="h-6 w-6 text-gray-700" />
        <span v-if="unreadCount" class="absolute -top-1 -right-1 bg-red-500 text-white text-xs font-bold w-5 h-5 rounded-full flex items-center justify-center ring-2 ring-white animate-bounce-short">
          {{ unreadCount }}
        </span>
      </button>

      <div class="relative">
        <button
          @click="showUserMenu = !showUserMenu"
          class="flex items-center space-x-2 focus:outline-none group"
        >
          <div class="h-9 w-9 bg-green-500 rounded-full flex items-center justify-center text-white font-bold text-lg shadow-md transition duration-300 ease-in-out group-hover:ring-2 group-hover:ring-green-300 group-hover:ring-opacity-50">
            {{ userInitials }}
          </div>
          <ChevronDownIcon class="h-4 w-4 text-gray-500 transition duration-300 ease-in-out group-hover:rotate-180" />
        </button>
        <ul
          v-if="showUserMenu"
          @click.away="showUserMenu = false"
          class="absolute right-0 mt-3 w-48 bg-white border border-gray-200 rounded-xl shadow-xl z-30 overflow-hidden animate-fade-in-down"
        >
          <li class="px-4 py-3 text-gray-800 font-semibold border-b border-gray-100">{{ user.email }}</li>
          <li>
            <router-link to="/profile" class="block px-4 py-3 hover:bg-gray-100 text-gray-700 transition duration-200 ease-in-out">Profile</router-link>
          </li>
          <li>
            <button @click="handleLogout" class="w-full text-left px-4 py-3 text-red-600 hover:bg-red-50 transition duration-200 ease-in-out">Logout</button>
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
/* Custom animations if needed beyond basic Tailwind */
@keyframes fade-in-down {
  from {
    opacity: 0;
    transform: translateY(-10px);
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
  0%, 100% {
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