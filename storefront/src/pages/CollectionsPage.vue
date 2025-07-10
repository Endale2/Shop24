<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 py-10">
    <nav class="flex items-center space-x-2 text-sm text-gray-500 mb-6">
      <button @click="goBack" class="text-gray-700 hover:text-black flex items-center gap-1 font-medium text-xs mr-2">
        <svg class="w-4 h-4" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7"/></svg>
        Back
      </button>
      <router-link :to="`/shops/${shopSlug}`" class="hover:underline">Home</router-link>
      <span>&gt;</span>
      <span class="text-gray-900 font-medium">Collections</span>
    </nav>
    <h1 class="text-3xl font-bold mb-8 text-gray-900 tracking-tight uppercase">Collections</h1>
    <div class="bg-white border border-gray-200 rounded-none p-6">
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
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { fetchCollections } from '../services/collections'
import type { Collection } from '../services/collections'

const route = useRoute()
const router = useRouter()
const shopSlug = route.params.shopSlug as string

const collections = ref<Collection[]>([])

function goBack() {
  router.back()
}

onMounted(async () => {
  collections.value = await fetchCollections(shopSlug)
})
</script>
