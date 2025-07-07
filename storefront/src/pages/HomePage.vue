<template>
  <Hero :shop="shop" />
  <section class="mt-12 max-w-7xl mx-auto px-2">
    <h2 class="text-3xl font-bold mb-8 text-gray-900 tracking-tight uppercase">Featured Products</h2>
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8">
      <ProductCard v-for="p in featured" :key="p.id" :product="p" />
    </div>
  </section>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import Hero from '@/components/Hero.vue'
import ProductCard from '@/components/ProductCard.vue'
import { fetchShop } from '@/services/shop'
import { fetchAllProducts } from '@/services/product'
import type { Shop } from '@/services/shop'
import type { Product } from '@/services/product'

const route = useRoute()
const shopSlug = route.params.shopSlug as string

const shop = ref<Shop | null>(null)
const featured = ref<Product[]>([])

onMounted(async () => {
  shop.value = await fetchShop(shopSlug)
  const all = await fetchAllProducts(shopSlug)
  featured.value = all.slice(0, 4)
})
</script>
