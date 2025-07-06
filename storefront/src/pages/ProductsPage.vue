<template>
  <h1 class="text-2xl mb-4">All Products</h1>
  <div class="grid sm:grid-cols-2 lg:grid-cols-4 gap-4">
    <ProductCard v-for="p in products" :key="p.id" :product="p" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import ProductCard from '@/components/ProductCard.vue'
import { fetchAllProducts } from '@/services/product'
import type { Product } from '@/services/product'

const route = useRoute()
const shopSlug = route.params.shopSlug as string

const products = ref<Product[]>([])

onMounted(async () => {
  products.value = await fetchAllProducts(shopSlug)
})
</script>
