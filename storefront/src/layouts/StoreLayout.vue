<template>
  <div class="min-h-screen flex flex-col bg-gray-50">
    <Header :shop="shop" />
    <main class="flex-1 w-full max-w-7xl mx-auto py-8 px-4">
      <div v-if="error" class="text-center py-12">
        <h1 class="text-2xl font-bold text-gray-900 mb-4">Shop Not Found</h1>
        <p class="text-gray-600 mb-6">The shop "{{ shopSlug }}" could not be found.</p>
        <router-link to="/" class="text-blue-600 hover:underline">Back to Shop Selection</router-link>
      </div>
      <router-view v-else />
    </main>
    <Footer :shop="shop" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import type { Shop } from '@/services/shop'
import { fetchShop, getCurrentShopSlug } from '@/services/shop'
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'

const shop = ref<Shop|null>(null)
const error = ref(false)
const shopSlug = getCurrentShopSlug()

onMounted(async () => {
  try {
    if (!shopSlug) {
      error.value = true
      return
    }
    shop.value = await fetchShop(shopSlug)
    if (!shop.value) {
      error.value = true
    }
  } catch (err) {
    console.error('Error fetching shop:', err)
    error.value = true
  }
})
</script>
