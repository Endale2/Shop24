<template>
  <header class="bg-white shadow-md py-4 px-6 flex items-center justify-between relative">
    <!-- Logo / Shop Name -->
    <router-link to="/" class="text-2xl font-extrabold text-gray-900">
      {{ shop.current?.name || 'Your Shop Name' }}
    </router-link>

    <!-- Desktop nav -->
    <nav class="hidden md:flex items-center space-x-6">
      <router-link to="/products" class="text-gray-700 hover:text-indigo-600">
        All Products
      </router-link>
      <router-link to="/categories" class="text-gray-700 hover:text-indigo-600">
        Categories
      </router-link>
      
      <router-link to="/cart" class="text-gray-700 hover:text-indigo-600">
        Cart
      </router-link>
      <router-link to="/signin" class="text-gray-700 hover:text-indigo-600">
        Sign In
      </router-link>
      <router-link to="/signup" class="text-gray-700 hover:text-indigo-600">
        Sign Up
      </router-link>
    </nav>

    <!-- Mobile menu button -->
    <button
      class="md:hidden text-gray-700 hover:text-indigo-600 focus:outline-none"
      aria-label="Open mobile menu"
      @click="mobileOpen = !mobileOpen"
    >
      <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none"
           viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M4 6h16M4 12h16M4 18h16" />
      </svg>
    </button>

    <!-- Mobile nav panel -->
    <div
      v-if="mobileOpen"
      class="absolute top-full left-0 w-full bg-white shadow-md py-4 md:hidden"
    >
      <nav class="flex flex-col space-y-2 px-6">
        <router-link to="/categories" class="text-gray-700 hover:text-indigo-600">
          Categories
        </router-link>
        <router-link to="/products" class="text-gray-700 hover:text-indigo-600">
          Products
        </router-link>
        <router-link to="/cart" class="text-gray-700 hover:text-indigo-600">
          Cart
        </router-link>
        <router-link to="/signin" class="text-gray-700 hover:text-indigo-600">
          Sign In
        </router-link>
        <router-link to="/signup" class="text-gray-700 hover:text-indigo-600">
          Sign Up
        </router-link>
      </nav>
    </div>
  </header>
</template>

<script setup>
import { ref }          from 'vue'
import { useShopStore } from '@/stores/shop'

const shop = useShopStore()
const mobileOpen = ref(false)
</script>
