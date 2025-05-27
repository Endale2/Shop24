<template>
  <div v-if="shop.current" class="space-y-16 md:space-y-24">
    <ShopHero :shop="shop.current" />

    <section class="container mx-auto px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-10 md:mb-12">
        <h2 class="text-3xl sm:text-4xl font-bold tracking-tight text-slate-800 mb-3">
          Featured Products
        </h2>
        <p class="text-lg text-slate-600 max-w-2xl mx-auto">
          Check out some of our handpicked favorites and bestsellers.
        </p>
      </div>
      <ProductGrid :products="featured" class="pb-8" />
      <div v-if="featured && featured.length > 0 && shop.products.length > featured.length" class="mt-8 text-center">
        <router-link
          to="/products"
          class="inline-flex items-center justify-center bg-indigo-600 hover:bg-indigo-700 text-white font-semibold py-3 px-8 rounded-lg shadow-md hover:shadow-lg transform hover:scale-105 transition-all duration-300 ease-in-out focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
          Browse All Products
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 ml-2" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M17 8l4 4m0 0l-4 4m4-4H3" />
          </svg>
        </router-link>
      </div>
    </section>

    </div>
  <div v-else class="flex flex-col items-center justify-center min-h-[calc(100vh-200px)] py-20 text-center">
    <svg class="animate-spin h-12 w-12 text-indigo-600 mb-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
    </svg>
    <p class="text-xl font-medium text-slate-600">Loading your amazing shop...</p>
    <p class="text-sm text-slate-500">Please wait a moment.</p>
  </div>
</template>

<script setup>
import { computed }     from 'vue'
import { useShopStore } from '@/stores/shop'
import ShopHero         from '@/components/ShopHero.vue'
import ProductGrid      from '@/components/ProductGrid.vue'

const shop     = useShopStore()
const featured = computed(() => {
  if (shop.products && shop.products.length > 0) {
    return shop.products.slice(0, Math.min(8, shop.products.length)); // Show up to 8 featured products
  }
  return [];
});
</script>

<style scoped>
/* No page-specific styles needed as ProductGrid handles its animation */
</style>