<template>
  <Hero :shop="shop" />
  <section class="mt-8">
    <h2 class="text-2xl mb-4">Featured Products</h2>
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
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
