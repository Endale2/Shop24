<template>
  
  <Hero :shop="shop" />
  <section class="mt-12 max-w-7xl mx-auto px-2">
    <h2 class="text-3xl font-bold mb-8 text-gray-900 tracking-tight uppercase">Featured Products</h2>
    <Loader v-if="isLoading" text="Loading products..." />
    <div v-else-if="featured.length > 0" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8">
      <ProductCard v-for="p in featured" :key="p.id" :product="p" />
    </div>
    <div v-else class="text-center py-12">
      <p class="text-gray-600">No products available at the moment.</p>
    </div>
  </section>
  <div class="text-center mt-8">
    <router-link :to="`/${shop?.slug || getCurrentShopSlug()}/products`" class="inline-block bg-black text-white py-3 px-8 rounded font-semibold uppercase tracking-wide hover:bg-gray-800 transition-colors">
      See All Products
    </router-link>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Header from '../components/Header.vue'
import Hero from '../components/Hero.vue'
import ProductCard from '../components/ProductCard.vue'
import { fetchShop, getCurrentShopSlug } from '../services/shop'
import { fetchAllProducts } from '../services/product'
import type { Shop } from '../services/shop'
import type { Product } from '../services/product'
import Loader from '@/components/Loader.vue'

const shop = ref<Shop | null>(null)
const featured = ref<Product[]>([])
const isLoading = ref(true)

onMounted(async () => {
  try {
    const shopSlug = getCurrentShopSlug()
    if (!shopSlug) return;
    shop.value = await fetchShop(shopSlug)
    if (shop.value) {
      const all = await fetchAllProducts(shopSlug)
      featured.value = all.slice(0, 4)
    }
  } catch (error) {
    console.error('Error loading home page data:', error)
  } finally {
    isLoading.value = false
  }
})
</script>
