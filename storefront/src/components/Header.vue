<template>
  <header class="bg-white shadow-md py-4 px-6 md:px-8 flex items-center justify-between z-40">
    <div class="flex items-center">
      <router-link to="/" class="text-2xl md:text-3xl font-extrabold text-gray-900 hover:text-indigo-600 transition duration-300 ease-in-out">
        {{ shop.current?.name || 'Your Shop Name' }}
      </router-link>
      </div>

    <nav class="hidden md:flex space-x-8">
      <router-link to="/products" class="text-gray-700 hover:text-indigo-600 font-medium transition duration-300 ease-in-out">Products</router-link>
      <router-link to="/categories" class="text-gray-700 hover:text-indigo-600 font-medium transition duration-300 ease-in-out">Categories</router-link>
      <router-link to="/about" class="text-gray-700 hover:text-indigo-600 font-medium transition duration-300 ease-in-out">About Us</router-link>
      <router-link to="/contact" class="text-gray-700 hover:text-indigo-600 font-medium transition duration-300 ease-in-out">Contact</router-link>
      </nav>

    <div class="flex items-center space-x-4">
      <button class="text-gray-600 hover:text-indigo-600 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-opacity-50 transition duration-300 ease-in-out" aria-label="Search">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
        </svg>
      </button>

      <router-link to="/cart" class="relative text-gray-600 hover:text-indigo-600 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-opacity-50 transition duration-300 ease-in-out" aria-label="Shopping Cart">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
        </svg>
        <span class="absolute -top-2 -right-2 bg-red-500 text-white text-xs font-bold rounded-full h-4 w-4 flex items-center justify-center">0</span>
      </router-link>

      <button class="md:hidden text-gray-600 hover:text-indigo-600 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-opacity-50 transition duration-300 ease-in-out" aria-label="Open mobile menu">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
        </svg>
      </button>
    </div>
  </header>
</template>

<script setup>
import { useShopStore } from '@/stores/shop'
import { RouterLink } from 'vue-router' // Import RouterLink for proper usage in template

const shop = useShopStore()
</script>