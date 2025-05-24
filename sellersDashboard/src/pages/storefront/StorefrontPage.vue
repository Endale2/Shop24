<!-- src/pages/StorefrontPage.vue -->
<template>
  <div class="p-6 max-w-5xl mx-auto space-y-8">
    <h1 class="text-4xl font-bold">
      {{ shop?.name || 'Shop Not Found' }}
    </h1>

    <p v-if="loading" class="text-gray-600">Loading…</p>
    <p v-else-if="error" class="text-red-600">{{ error }}</p>

    <div v-else>
      <div class="text-gray-700 mb-4">{{ shop.description }}</div>
      <h2 class="text-2xl font-semibold mb-6">Products</h2>

      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
        <div
          v-for="p in products"
          :key="p.id"
          class="bg-white p-4 rounded-lg shadow hover:shadow-lg transition"
        >
          <img
            v-if="p.images?.length"
            :src="p.images[0]"
            class="w-full h-40 object-cover rounded"
          />
          <div class="mt-3">
            <h3 class="font-medium">{{ p.name }}</h3>
            <p class="text-green-600 font-bold mt-1">{{ formatPrice(p.price) }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { shopService } from '@/services/shop'
import { productService } from '@/services/product'

const route = useRoute()
const hostname = window.location.hostname     // e.g. "shop123.localhost"
const shopId = hostname.split('.')[0]         // "shop123"

const shop = ref(null)
const products = ref([])
const loading = ref(true)
const error = ref(null)

onMounted(async () => {
  loading.value = true
  try {
    // 1) fetch the shop by ID, publicly
    shop.value = await shopService.fetchById(shopId)

    // 2) then fetch its products
    products.value = await productService.fetchByShop(shopId)
  } catch (e) {
    console.error(e)
    error.value = 'Could not load this storefront.'
  } finally {
    loading.value = false
  }
})

function formatPrice(p) {
  return p != null ? `$${p.toFixed(2)}` : 'N/A'
}
</script>
