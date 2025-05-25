<template>
  <div>
    <ShopHero :shop="shop.info" />
    <h2 class="text-2xl font-bold my-4">Products</h2>
    <ProductGrid :products="shop.products" />
  </div>
</template>

<script setup>
import { useShopStore } from '@/stores/shop'
import ShopHero from '@/components/ShopHero.vue'
import ProductGrid from '@/components/ProductGrid.vue'

const shop = useShopStore()
</script>
