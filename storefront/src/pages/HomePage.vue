<template>
  <div v-if="shop.current">
    <!-- Hero Section -->
    <ShopHero :shop="shop.current" />

    <!-- Featured Products Section -->
    <section class="container mx-auto px-4 sm:px-6 lg:px-8 py-12">
      <h2 class="text-3xl font-extrabold text-gray-900 mb-6">Featured Products</h2>
      <ProductGrid :products="featured" />
      <div class="mt-8 text-center">
        <router-link
          to="/products"
          class="inline-block bg-indigo-600 hover:bg-indigo-700 text-white font-semibold py-3 px-6 rounded-md transition"
        >
          Browse All Products
        </router-link>
      </div>
    </section>
  </div>
  <div v-else class="flex items-center justify-center h-full py-20">
    <p class="text-lg text-gray-500">Loading shop…</p>
  </div>
</template>

<script setup>
import { computed }     from 'vue'
import { useShopStore } from '@/stores/shop'
import ShopHero         from '@/components/ShopHero.vue'
import ProductGrid      from '@/components/ProductGrid.vue'

const shop     = useShopStore()
const featured = computed(() => shop.products.slice(0, 6))
</script>

<style scoped>
/* Add subtle fade-in animation for featured products */
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to   { opacity: 1; transform: translateY(0); }
}
.product-enter-active {
  animation: fadeIn 0.5s ease-out both;
}
</style>
