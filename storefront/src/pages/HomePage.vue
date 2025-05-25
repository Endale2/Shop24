<template>
  <div v-if="shop.current">
    <ShopHero :shop="shop.current" />
    <h2 class="text-2xl font-bold my-4">Products</h2>
    <ProductGrid
      :products="shop.products"
      :shopSlug="shop.current.slug"
    />
  </div>
  <div v-else class="flex justify-center items-center h-full">
    <p>Loading shop...</p>
  </div>
</template>

<script setup>
import { useShopStore } from '@/stores/shop'
import ShopHero from '@/components/ShopHero.vue'
import ProductGrid from '@/components/ProductGrid.vue'

const shop = useShopStore()
</script>