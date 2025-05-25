<template>
  <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-6 p-4">
    <ProductCard v-for="p in products" :key="p.id" :product="p" />
  </div>
</template>

<script setup>
import ProductCard from './ProductCard.vue'
defineProps({
  products: { type: Array, default: () => [] }
})
</script>