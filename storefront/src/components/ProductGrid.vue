<template>
  <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-6 p-4">
    <router-link
      v-for="p in products"
      :key="p.slug"
      :to="{ name: 'Product', params: { shopSlug: shopSlug, productSlug: p.slug } }"
      class="block"
    >
      <ProductCard :product="p" />
    </router-link>
  </div>
</template>

<script setup>
import { defineProps } from 'vue'
import ProductCard from './ProductCard.vue'

defineProps({
  products: { type: Array, default: () => [] },
  shopSlug: { type: String, required: true }
})
</script>
