<template>
  <section v-if="shop.current">
    <h1 class="text-3xl mb-6">All Products</h1>
    <ProductGrid :products="shop.products" />
  </section>
  <div v-else class="text-center py-20">Loading products…</div>
</template>

<script setup>
import { useShopStore } from '@/stores/shop'
import ProductGrid      from '@/components/ProductGrid.vue'

const shop = useShopStore()
</script>