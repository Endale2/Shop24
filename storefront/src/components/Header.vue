<template>
  <header class="py-4 px-6 flex items-center justify-between relative border-b border-gray-200/70">
    
    <router-link
  to="/"
  class="flex items-center space-x-2 text-2xl font-bold text-slate-800 hover:text-indigo-700 transition-colors duration-300"
>
  <img
    v-if="shop.current?.image"
    :src="shop.current.image"
    alt="Shop Logo"
    class="h-8 w-8 rounded-md object-cover"
  />
  <span>{{ shop.current?.name || 'Your Shop Name' }}</span>
</router-link>


    <nav class="hidden md:flex items-center space-x-4 lg:space-x-6">
      <router-link
        to="/products"
        class="px-3 py-2 rounded-md text-sm font-medium text-slate-700 hover:bg-indigo-50 hover:text-indigo-700 transition-all duration-200"
        active-class="bg-indigo-100 text-indigo-700 font-semibold"
      >
        All Products
      </router-link>
      <router-link
        to="/categories"
        class="px-3 py-2 rounded-md text-sm font-medium text-slate-700 hover:bg-indigo-50 hover:text-indigo-700 transition-all duration-200"
        active-class="bg-indigo-100 text-indigo-700 font-semibold"
      >
        Categories
      </router-link>
      </nav>

    <div class="hidden md:flex items-center space-x-4 lg:space-x-5">
      <router-link
        to="/cart"
        class="relative text-slate-600 hover:text-indigo-700 transition-colors duration-200 group"
        title="View Cart"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
        </svg>
        <span class="absolute -bottom-1 left-1/2 transform -translate-x-1/2 w-max opacity-0 group-hover:opacity-100 transition-opacity duration-300 bg-slate-700 text-white text-xs rounded py-1 px-2">
          Cart
        </span>
      </router-link>
      <router-link
        to="/signin"
        class="px-3 py-2 rounded-md text-sm font-medium text-slate-700 hover:bg-slate-100 hover:text-slate-900 transition-all duration-200"
        active-class="text-indigo-700 font-semibold"
      >
        Sign In
      </router-link>
      <router-link
        to="/signup"
        class="px-4 py-2 rounded-md text-sm font-medium bg-indigo-600 text-white hover:bg-indigo-700 transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
      >
        Sign Up
      </router-link>
    </div>

    <button
      class="md:hidden text-slate-600 hover:text-indigo-700 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500 p-1 rounded"
      aria-label="Open mobile menu"
      @click="mobileOpen = !mobileOpen"
    >
      <svg v-if="!mobileOpen" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
        <path stroke-linecap="round" stroke-linejoin="round" d="M4 6h16M4 12h16M4 18h16" />
      </svg>
      <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
        <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
      </svg>
    </button>

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
        class="absolute top-full inset-x-0 mt-px bg-white shadow-lg rounded-b-md md:hidden z-40 border-t border-gray-200/70"
        @click="mobileOpen = false" 
      >
        <nav class="flex flex-col space-y-1 p-4">
          <router-link to="/products" class="block px-3 py-2 rounded-md text-base font-medium text-slate-700 hover:bg-indigo-50 hover:text-indigo-700" active-class="bg-indigo-100 text-indigo-700">
            All Products
          </router-link>
          <router-link to="/categories" class="block px-3 py-2 rounded-md text-base font-medium text-slate-700 hover:bg-indigo-50 hover:text-indigo-700" active-class="bg-indigo-100 text-indigo-700">
            Categories
          </router-link>
          <router-link to="/cart" class="block px-3 py-2 rounded-md text-base font-medium text-slate-700 hover:bg-indigo-50 hover:text-indigo-700" active-class="bg-indigo-100 text-indigo-700">
            Cart
          </router-link>
          <hr class="my-2 border-slate-200" />
          <router-link to="/signin" class="block px-3 py-2 rounded-md text-base font-medium text-slate-700 hover:bg-indigo-50 hover:text-indigo-700" active-class="bg-indigo-100 text-indigo-700">
            Sign In
          </router-link>
          <router-link to="/signup" class="block w-full text-center mt-2 px-4 py-2 rounded-md text-base font-medium bg-indigo-600 text-white hover:bg-indigo-700 transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
            Sign Up
          </router-link>
        </nav>
      </div>
    </transition>
  </header>
</template>

<script setup>
import { ref, watch } from 'vue'
import { useShopStore } from '@/stores/shop'
import { useRoute } from 'vue-router'

const shop = useShopStore()
const mobileOpen = ref(false)
const route = useRoute()

watch(() => route.path, () => {
  mobileOpen.value = false
})
</script>