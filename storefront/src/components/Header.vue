<template>
  <header class="py-4 px-6 flex items-center justify-between relative bg-white/90 backdrop-blur-md border-b border-gray-100 shadow-lg z-50">

    <!-- Shop Logo and Name -->
    <router-link
      to="/"
      class="flex items-center space-x-2 text-2xl font-bold text-gray-800 hover:text-green-700 transition-colors duration-300 transform hover:scale-105 origin-left focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2 rounded-lg"
    >
      <img
        v-if="shop.current?.image"
        :src="shop.current.image"
        alt="Shop Logo"
        class="h-9 w-9 rounded-full object-cover shadow-sm border border-gray-100"
        onerror="this.onerror=null;this.src='https://placehold.co/36x36/E0F2F7/263238?text=SL';"
      />
      <span class="tracking-tight">{{ shop.current?.name || 'Your Shop Name' }}</span>
    </router-link>

    <!-- Desktop Navigation Links -->
    <nav class="hidden md:flex items-center space-x-4 lg:space-x-6">
      <router-link
        to="/products"
        class="px-4 py-2 rounded-lg text-base font-medium text-gray-700 hover:bg-green-50 hover:text-green-700 transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-green-300 focus:ring-offset-2"
        active-class="bg-green-100 text-green-700 font-semibold shadow-sm"
      >
        All Products
      </router-link>
      <router-link
        to="/categories"
        class="px-4 py-2 rounded-lg text-base font-medium text-gray-700 hover:bg-green-50 hover:text-green-700 transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-green-300 focus:ring-offset-2"
        active-class="bg-green-100 text-green-700 font-semibold shadow-sm"
      >
        Categories
      </router-link>
    </nav>

    <!-- Desktop Action Buttons (Cart, Sign In, Sign Up) -->
    <div class="hidden md:flex items-center space-x-4 lg:space-x-5">
      <router-link
        to="/cart"
        class="relative text-gray-600 hover:text-green-700 transition-colors duration-200 group p-2 rounded-full hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-green-300 focus:ring-offset-2"
        title="View Cart"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
        </svg>
        <span class="absolute -bottom-1 left-1/2 transform -translate-x-1/2 w-max opacity-0 group-hover:opacity-100 transition-opacity duration-300 bg-gray-800 text-white text-xs rounded py-1 px-2 pointer-events-none">
          Cart
        </span>
      </router-link>

      <router-link
        to="/signin"
        class="px-4 py-2 rounded-lg text-base font-medium text-green-700 hover:bg-green-50 hover:text-green-800 transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-green-300 focus:ring-offset-2"
        active-class="bg-green-100 text-green-800 font-semibold"
      >
        Sign In
      </router-link>
      <router-link
        to="/signup"
        class="px-5 py-2 bg-gradient-to-br from-green-600 to-green-700 text-white font-semibold rounded-full shadow-md hover:shadow-lg transition-all duration-300 transform hover:-translate-y-0.5 active:translate-y-0 active:shadow-md focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2"
      >
        Sign Up
      </router-link>
    </div>

    <!-- Mobile Menu Toggle Button -->
    <button
      class="md:hidden text-gray-600 hover:text-green-700 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-green-500 p-2 rounded-md"
      aria-label="Open mobile menu"
      @click="mobileOpen = !mobileOpen"
    >
      <svg v-if="!mobileOpen" xmlns="http://www.w3.org/2000/svg" class="h-7 w-7" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
        <path stroke-linecap="round" stroke-linejoin="round" d="M4 6h16M4 12h16M4 18h16" />
      </svg>
      <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-7 w-7" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
        <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
      </svg>
    </button>

    <!-- Mobile Menu Overlay (Transition) -->
    <transition
      enter-active-class="transition ease-out duration-200 transform"
      enter-from-class="opacity-0 scale-95"
      enter-to-class="opacity-100 scale-100"
      leave-active-class="transition ease-in duration-150 transform"
      leave-from-class="opacity-100 scale-100"
      leave-to-class="opacity-0 scale-95"
    >
      <div
        v-if="mobileOpen"
        class="absolute top-full inset-x-0 mt-px bg-white shadow-xl rounded-b-lg md:hidden z-40 border-t border-gray-200/70"
        @click="mobileOpen = false"
      >
        <nav class="flex flex-col space-y-1 p-4">
          <router-link to="/products" class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:bg-green-50 hover:text-green-700 transition-colors duration-200" active-class="bg-green-100 text-green-700 font-semibold">
            All Products
          </router-link>
          <router-link to="/categories" class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:bg-green-50 hover:text-green-700 transition-colors duration-200" active-class="bg-green-100 text-green-700 font-semibold">
            Categories
          </router-link>
          <router-link to="/cart" class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:bg-green-50 hover:text-green-700 transition-colors duration-200" active-class="bg-green-100 text-green-700">
            Cart
          </router-link>
          <hr class="my-3 border-gray-200" />
          <router-link to="/signin" class="block px-3 py-2 rounded-md text-base font-medium text-green-700 hover:bg-green-50 hover:text-green-800 transition-colors duration-200" active-class="bg-green-100 text-green-800 font-semibold">
            Sign In
          </router-link>
          <router-link to="/signup" class="block w-full text-center mt-3 px-5 py-2 rounded-md text-base font-medium bg-green-600 text-white hover:bg-green-700 transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500">
            Sign Up
          </router-link>
        </nav>
      </div>
    </transition>
  </header>
</template>

<script setup>
import { ref, watch } from 'vue'
import { useShopStore } from '@/stores/shop' // Ensure this path is correct
import { useRoute } from 'vue-router'

const shop = useShopStore()
const mobileOpen = ref(false)
const route = useRoute()

// Close mobile menu when route changes
watch(() => route.path, () => {
  mobileOpen.value = false
})
</script>

<style scoped>
/* No additional custom styles needed as Tailwind handles everything. */
</style>
