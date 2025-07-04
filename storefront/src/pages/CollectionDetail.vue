<template>
  <div v-if="collection" class="space-y-6">
    <div class="relative h-48 bg-cover bg-center rounded" :style="{ backgroundImage: `url(${collection.image})` }">
      <div class="absolute inset-0 bg-black bg-opacity-40 flex items-center justify-center">
        <h1 class="text-white text-3xl font-bold">{{ collection.title }}</h1>
      </div>
    </div>
    <p class="text-gray-700">{{ collection.description }}</p>
    <h2 class="text-2xl">Products ({{ collection.products ? collection.products.length : 0 }})</h2>
    <div class="grid sm:grid-cols-2 lg:grid-cols-4 gap-4">
      <div v-for="p in collection.products" :key="p.id" class="border rounded p-4 hover:shadow">
        <img :src="p.main_image" alt="" class="w-full h-40 object-cover mb-2"/>
        <h3 class="font-semibold">{{ p.name }}</h3>
        <p class="text-gray-700">${{ p.display_price.toFixed(2) }}</p>
        <router-link :to="`/products/${p.slug}`" class="text-blue-500">View</router-link>
      </div>
    </div>
  </div>
  <div v-else class="text-center py-12">Loading collection…</div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { fetchCollectionDetail, CollectionDetail } from '@/services/collections'

const route = useRoute()
const collection = ref<CollectionDetail | null>(null)

onMounted(async () => {
  const handle = route.params.handle as string
  // It's good practice to ensure that fetchCollectionDetail always returns an object
  // with a 'products' array, even if empty, or handle the null/undefined case
  // more explicitly if the service can return null for 'products'.
  collection.value = await fetchCollectionDetail(handle)
})
</script>