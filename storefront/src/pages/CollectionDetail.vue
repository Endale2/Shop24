<template>
  <div v-if="collection" class="space-y-10">
    <div class="relative h-64 md:h-80 bg-cover bg-center rounded-lg border border-gray-200 overflow-hidden" :style="{ backgroundImage: collection.image ? `url(${collection.image})` : '' }">
      <div class="absolute inset-0 bg-gradient-to-t from-black/50 to-transparent flex flex-col items-center justify-end p-8 text-center">
        <h1 class="text-4xl md:text-5xl font-extrabold text-white tracking-tight uppercase" style="text-shadow: 2px 2px 4px rgba(0,0,0,0.5);">{{ collection.title }}</h1>
        <p v-if="collection.description" class="text-white/90 text-md max-w-2xl mt-4">{{ collection.description }}</p>
      </div>
    </div>

    <div>
      <h2 class="text-xl font-bold text-gray-900 mb-6 uppercase tracking-wider">
        Products ({{ collection.products ? collection.products.length : 0 }})
      </h2>
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8">
        <ProductCard v-for="p in collection.products" :key="p.id" :product="p" />
      </div>
    </div>
  </div>
  <div v-else class="text-center py-16 text-gray-500">
    Loading collection…
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { fetchCollectionDetail, CollectionDetail } from '@/services/collections'
import ProductCard from '@/components/ProductCard.vue'

const route = useRoute()
const shopSlug = route.params.shopSlug as string
const handle = route.params.handle as string
const collection = ref<CollectionDetail | null>(null)

onMounted(async () => {
  collection.value = await fetchCollectionDetail(shopSlug, handle)
})
</script>
