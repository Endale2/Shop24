<template>
  <div class="container mx-auto px-4 sm:px-6 lg:px-8 py-8 md:py-12">
    <section v-if="shop.current && shop.products.length > 0">
      <div class="mb-8 md:mb-10 text-center sm:text-left">
        <h1 class="text-3xl sm:text-4xl font-bold tracking-tight text-slate-800">
          All Products
        </h1>
        <p v-if="shop.products.length > 0" class="mt-2 text-md text-slate-600">
          Explore our full collection of available items.
        </p>
      </div>
      <ProductGrid :products="shop.products" />
    </section>
    <section v-else-if="shop.current && shop.products.length === 0" class="text-center py-16 md:py-20">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 mx-auto text-slate-400 mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1">
            <path stroke-linecap="round" stroke-linejoin="round" d="M20.25 7.5l-.625 10.632a2.25 2.25 0 01-2.247 2.118H6.622a2.25 2.25 0 01-2.247-2.118L3.75 7.5M10 11.25h4M3.375 7.5h17.25c.621 0 1.125-.504 1.125-1.125v-1.5c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125z" />
        </svg>
        <h2 class="text-2xl font-semibold text-slate-700 mb-2">No Products Yet</h2>
        <p class="text-slate-500">It looks like we're still stocking our shelves. Please check back soon!</p>
    </section>
    <div v-else class="flex flex-col items-center justify-center min-h-[calc(100vh-200px)] py-20 text-center">
      <svg class="animate-spin h-12 w-12 text-indigo-600 mb-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      <p class="text-xl font-medium text-slate-600">Loading products...</p>
      <p class="text-sm text-slate-500">Hang tight, awesome items are on their way!</p>
    </div>
  </div>
</template>

<script setup>
import { useShopStore } from '@/stores/shop'
import ProductGrid      from '@/components/ProductGrid.vue'

const shop = useShopStore()
</script>

<style scoped>
/* No page-specific styles needed */
</style>