<template>
  <div class="min-h-screen flex flex-col bg-gradient-to-br from-gray-50 to-green-50 text-gray-900 antialiased font-sans">
    <!-- Assuming AppNavbar and AppFooter are handled in a parent layout or globally -->
    <!-- <AppNavbar /> -->

    <div v-if="shop.current" class="flex-grow space-y-16 md:space-y-24 py-12">
      <!-- Shop Hero Section - assumes ShopHero component handles its own internal beautiful design -->
      <ShopHero :shop="shop.current" />

      <!-- Featured Products Section -->
      <section class="container mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-12 md:mb-16">
          <h2 class="text-4xl sm:text-5xl font-extrabold tracking-tight text-gray-900 mb-4 relative inline-block">
            Featured Products
            <span class="absolute left-1/2 transform -translate-x-1/2 bottom-0 w-24 h-1 bg-green-500 rounded-full"></span>
          </h2>
          <p class="text-lg text-gray-700 max-w-2xl mx-auto">
            Discover our handpicked favorites and bestsellers, curated just for you.
          </p>
        </div>
        <ProductGrid :products="featured" class="pb-8" />
        <div v-if="featured && featured.length > 0 && shop.products.length > featured.length" class="mt-12 text-center">
          <router-link
            to="/products"
            class="inline-flex items-center justify-center px-8 py-3 bg-gradient-to-br from-green-600 to-green-700 text-white text-lg font-bold rounded-full shadow-lg hover:shadow-xl transition-all duration-300 transform hover:-translate-y-0.5 active:translate-y-0 active:shadow-md focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2"
          >
            Browse All Products
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 ml-2" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M17 8l4 4m0 0l-4 4m4-4H3" />
            </svg>
          </router-link>
        </div>
      </section>

      <!-- Optional: Additional Sections could go here -->
      <!-- For example, categories showcase, testimonials, etc. -->

    </div>
    <!-- Loading and Error States -->
    <div v-else class="flex flex-col items-center justify-center min-h-[calc(100vh-200px)] py-20 px-4 text-center bg-white rounded-xl shadow-lg m-4">
      <svg class="animate-spin h-14 w-14 text-green-600 mb-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      <p class="text-2xl font-semibold text-gray-700 mb-2">Loading your amazing shop...</p>
      <p class="text-base text-gray-500">Please wait a moment while we prepare everything.</p>
    </div>

    <!-- <AppFooter /> -->
  </div>
</template>

<script setup>
import { computed }     from 'vue'
import { useShopStore } from '@/stores/shop'
import ShopHero         from '@/components/ShopHero.vue'
import ProductGrid      from '@/components/ProductGrid.vue'

const shop = useShopStore()
const featured = computed(() => {
  if (shop.products && shop.products.length > 0) {
    // Show up to 8 featured products, or fewer if not enough products
    return shop.products.slice(0, Math.min(8, shop.products.length));
  }
  return [];
});
</script>

<style scoped>
/* No page-specific styles needed as Tailwind handles most styling. */
/* Ensure 'font-sans' is correctly configured in your Tailwind config to use 'Inter' or a preferred sans-serif font */
</style>
