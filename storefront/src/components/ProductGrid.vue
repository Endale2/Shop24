// src/components/ProductGrid.vue
<template>
  <div class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 gap-6">
    <router-link
      v-for="p in products"
      :key="p.id"
      :to="`/products/${p.slug}`"
      class="block group transform hover:scale-105 transition duration-300"
    >
      <ProductCard :product="p" />
    </router-link>
  </div>
</template>

<script setup>
import { defineProps } from 'vue'
import ProductCard      from './ProductCard.vue'

defineProps({
  products: { type: Array, required: true }
})
</script>

<style scoped>
/* Optional: fade-in for cards */
@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(10px); }
  to   { opacity: 1; transform: translateY(0); }
}
</style>