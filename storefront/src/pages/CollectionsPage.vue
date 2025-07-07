<template>
  <h1 class="text-3xl font-bold mb-8 text-gray-900 tracking-tight uppercase">Collections</h1>
  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
    <router-link
      v-for="c in collections"
      :key="c.id"
      :to="`/shops/${shopSlug}/collections/${c.handle}`"
      class="bg-white border border-gray-200 rounded-none overflow-hidden group transition-colors hover:border-black cursor-pointer focus:outline-none focus:ring-2 focus:ring-black flex flex-col h-full"
    >
      <img :src="c.image" class="w-full h-56 object-contain bg-gray-50" alt="" />
      <div class="p-4 flex-1 flex flex-col justify-end">
        <h2 class="font-semibold text-base tracking-wide uppercase text-gray-900 leading-tight line-clamp-1">{{ c.title }}</h2>
      </div>
    </router-link>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { fetchCollections } from '@/services/collections'
import type { Collection } from '@/services/collections'

const route = useRoute()
const shopSlug = route.params.shopSlug as string

const collections = ref<Collection[]>([])

onMounted(async () => {
  collections.value = await fetchCollections(shopSlug)
})
</script>
