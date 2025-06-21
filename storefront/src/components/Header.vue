<template>
  <header class="w-full sticky top-0 z-50 bg-white/70 backdrop-blur-3xl border-b border-white/80 shadow-md transition-all duration-300 ease-in-out">
    <div class="max-w-7xl mx-auto flex items-center justify-between py-3.5 px-4 lg:px-8">

      <router-link to="/" class="flex items-center space-x-2 text-2xl font-extrabold text-gray-800 hover:text-green-600 transition-colors duration-300 transform hover:scale-105 origin-left focus:outline-none focus:ring-0 focus:ring-offset-0"> <img
          v-if="shop.current?.image"
          :src="shop.current.image"
          :alt="`${shop.current?.name || 'Shop'} Logo`"
          class="h-9 w-9 rounded-full object-cover shadow-inner border border-white/50 flex-shrink-0"
          @error="onImageError"
        />
        <span class="tracking-tighter whitespace-nowrap overflow-hidden text-ellipsis text-transparent bg-clip-text bg-gradient-to-r from-green-600 to-green-800">
          {{ shop.current?.name || 'Your Brand' }}
        </span>
      </router-link>

      <div class="hidden md:flex flex-grow justify-center max-w-md mx-auto px-4">
        <div class="relative w-full">
          <input
            type="text"
            v-model="searchQuery"
            @keyup.enter="performSearch"
            placeholder="Search products..."
            aria-label="Search products"
            class="w-full pl-10 pr-4 py-2 rounded-full border border-gray-200/50 bg-white/50 focus:outline-none focus:ring-2 focus:ring-green-300 focus:border-green-300 transition-all duration-200 text-gray-700 placeholder-gray-400 text-sm shadow-inner"
          />
          <svg class="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </div>
      </div>

      <div class="hidden md:flex items-center space-x-5 lg:space-x-7">
        <nav class="flex space-x-4 lg:space-x-6">
          <router-link
            to="/products"
            class="relative group px-4 py-2 rounded-full text-base font-medium text-gray-700 hover:text-green-700 transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-green-300 focus:ring-offset-2"
            :class="{ 'text-green-700 font-semibold': $route.path === '/products' || ($route.path === '/search' && $route.query.q) }"
          >
            All Products
            <span
              class="absolute bottom-0 left-1/2 w-0 h-0.5 bg-green-500 rounded-full transition-all duration-300 transform -translate-x-1/2 group-hover:w-full"
              :class="{ 'w-full': $route.path === '/products' || ($route.path === '/search' && $route.query.q) }"
            ></span>
          </router-link>
          <router-link
            to="/collections"
            class="relative group px-4 py-2 rounded-full text-base font-medium text-gray-700 hover:text-green-700 transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-green-300 focus:ring-offset-2"
            :class="{ 'text-green-700 font-semibold': $route.path === '/collections' }"
          >
            Collections
            <span
              class="absolute bottom-0 left-1/2 w-0 h-0.5 bg-green-500 rounded-full transition-all duration-300 transform -translate-x-1/2 group-hover:w-full"
              :class="{ 'w-full': $route.path === '/collections' }"
            ></span>
          </router-link>
        </nav>

        <router-link to="/cart" title="View Cart" class="relative p-2 rounded-full text-gray-600 hover:text-green-700 hover:bg-green-50 transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-green-300 focus:ring-offset-2">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
          </svg>
          <span v-if="cart.totalItems > 0" class="absolute -top-1 -right-1 bg-red-500 text-white text-xs font-bold rounded-full h-4 w-4 flex items-center justify-center pointer-events-none ring-1 ring-white">{{ cart.totalItems }}</span>
          <span v-else class="absolute -bottom-1 left-1/2 transform -translate-x-1/2 w-max opacity-0 group-hover:opacity-100 transition-opacity duration-300 bg-gray-800 text-white text-xs rounded py-1 px-2 pointer-events-none whitespace-nowrap">Cart</span>
        </router-link>

        <template v-if="auth.isAuthenticated">
          <div class="relative group">
            <button class="flex items-center space-x-2 px-3 py-1.5 rounded-full hover:bg-gray-50 transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-green-300 focus:ring-offset-2" aria-expanded="false" aria-haspopup="true">
                <div v-if="auth.user.avatar" class="h-8 w-8 rounded-full overflow-hidden flex-shrink-0">
                    <img :src="auth.user.avatar" :alt="`${auth.user.displayName || 'User'} avatar`" class="object-cover h-full w-full"/>
                </div>
                <div v-else class="h-8 w-8 flex items-center justify-center bg-green-100 rounded-full text-green-700 font-semibold flex-shrink-0">
                    {{ initials }}
                </div>
                <span class="font-medium text-gray-800 text-sm whitespace-nowrap overflow-hidden text-ellipsis hidden lg:block">{{ auth.user.displayName }}</span>
                <svg class="h-4 w-4 text-gray-500 transition-transform duration-200 group-hover:rotate-180" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path></svg>
            </button>
            <div class="absolute right-0 top-full mt-2 w-48 bg-white/90 backdrop-blur-md border border-white/70 rounded-lg shadow-lg opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all duration-200 z-10 origin-top-right scale-95 group-hover:scale-100">
                <router-link to="/account" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100/50 rounded-t-lg transition-colors duration-150">My Account</router-link>
                <button @click="onLogout" class="w-full text-left px-4 py-2 text-sm text-red-600 hover:bg-red-50/50 rounded-b-lg transition-colors duration-150">Sign Out</button>
            </div>
          </div>
        </template>

        <template v-else>
          <router-link to="/auth" class="px-5 py-2 bg-green-600 text-white rounded-full text-base font-medium hover:bg-green-700 transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-green-300 focus:ring-offset-2 shadow-md hover:shadow-lg">
            Sign In / Register
          </router-link>
        </template>
      </div>

      <button @click="mobileOpen = !mobileOpen" class="md:hidden p-2 text-gray-600 hover:text-green-700 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-green-500 rounded-md" aria-label="Toggle mobile menu">
        <svg v-if="!mobileOpen" xmlns="http://www.w3.org/2000/svg" class="h-7 w-7" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M4 6h16M4 12h16M4 18h16"/>
        </svg>
        <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-7 w-7" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/>
        </svg>
      </button>
    </div>

    <transition enter-active-class="transition ease-out duration-200 transform" enter-from-class="opacity-0 scale-95" enter-to-class="opacity-100 scale-100" leave-active-class="transition ease-in duration-150 transform" leave-from-class="opacity-100 scale-100" leave-to-class="opacity-0 scale-95">
      <div v-if="mobileOpen" class="absolute top-full inset-x-0 mt-px bg-white/90 shadow-xl rounded-b-lg md:hidden z-40 border-t border-gray-200/70 backdrop-blur-md">
        <nav class="flex flex-col space-y-1 p-4">
          <div class="relative mb-4">
            <input v-model="searchQuery" @keyup.enter="performSearch" type="text" placeholder="Search products..." aria-label="Search products" class="w-full pl-10 pr-4 py-2 rounded-full border border-gray-300/50 bg-white/50 focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200 text-gray-700 placeholder-gray-400 text-sm" />
            <svg class="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </div>
          <router-link
            to="/products"
            @click="closeMobileMenu"
            class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:bg-green-50/50 hover:text-green-700 transition-colors duration-200"
            :class="{ 'bg-green-100/50 text-green-700 font-semibold': $route.path === '/products' || ($route.path === '/search' && $route.query.q) }"
          >All Products</router-link>
          <router-link
            to="/collections"
            @click="closeMobileMenu"
            class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:bg-green-50/50 hover:text-green-700 transition-colors duration-200"
            :class="{ 'bg-green-100/50 text-green-700 font-semibold': $route.path === '/collections' }"
          >Collections</router-link>
          <router-link to="/cart" @click="closeMobileMenu" class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:bg-green-50/50 hover:text-green-700 transition-colors duration-200" active-class="bg-green-100/50 text-green-700">Cart</router-link>
          <hr class="my-3 border-gray-200/50" />
          <template v-if="auth.isAuthenticated">
            <router-link to="/account" @click="closeMobileMenu" class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:bg-green-50/50 hover:text-green-700 transition-colors duration-200" active-class="bg-green-100/50 text-green-700 font-semibold">My Account</router-link>
            <button @click="onLogoutAndCloseMobile" class="w-full text-left px-3 py-2 rounded-md text-base font-medium text-red-600 hover:bg-red-50/50 hover:text-red-700 transition-colors duration-200">Sign Out</button>
          </template>
          <template v-else>
            <router-link to="/auth" @click="closeMobileMenu" class="block px-3 py-2 rounded-md text-base font-medium text-green-700 hover:bg-green-50/50 hover:text-green-800 transition-colors duration-200" active-class="bg-green-100/50 text-green-800 font-semibold">Sign In / Sign Up</router-link>
          </template>
        </nav>
      </div>
    </transition>
  </header>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { useShopStore } from '@/stores/shop'
import { useAuthStore } from '@/stores/auth'
import { useCartStore } from '@/stores/cart'
import { useRoute, useRouter } from 'vue-router'

const shop = useShopStore()
const auth = useAuthStore()
const cart = useCartStore()

const mobileOpen = ref(false)
const searchQuery = ref('')
const route = useRoute()
const router = useRouter()

// Watch for route changes to close the mobile menu automatically
watch(() => route.path, () => {
  mobileOpen.value = false
})

const performSearch = () => {
  if (searchQuery.value.trim()) {
    // If the current path is already /products or /collections,
    // we want to update the query without navigating again (which causes a page reload)
    if (route.path === '/products' || route.path === '/collections') {
      router.replace({ path: route.path, query: { q: searchQuery.value.trim() } })
    } else {
      // Otherwise, navigate to the products page with the search query
      router.push({ path: '/products', query: { q: searchQuery.value.trim() } })
    }
    searchQuery.value = ''
    closeMobileMenu() // Close mobile menu after search
  }
}

const onLogout = async () => {
  await auth.logout() // Use await to ensure logout completes before redirect
  router.push({ name: 'Home' })
}

const onLogoutAndCloseMobile = () => {
  onLogout()
  closeMobileMenu()
}

const onImageError = (e) => {
  e.target.src = 'https://placehold.co/36x36/E0F2F7/263238?text=SL' // Fallback image for shop logo
}

const closeMobileMenu = () => {
  mobileOpen.value = false
}

// Compute initials for the avatar placeholder
const initials = computed(() => {
  if (!auth.user || (!auth.user.displayName && !auth.user.email)) {
    return ''
  }
  const name = auth.user.displayName || auth.user.email.split('@')[0]
  return name
    .split(' ')
    .map(n => n[0])
    .join('')
    .toUpperCase()
    .substring(0, 2) // Take max 2 initials
})
</script>

<style scoped>
/* Tailwind CSS handles styling. No custom CSS typically needed for these components. */
</style>