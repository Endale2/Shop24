<template>
  <div class="min-h-screen flex flex-col bg-gradient-to-br from-gray-50 to-green-50 text-gray-900 antialiased font-sans py-12 px-4 sm:px-6 lg:px-8">
    <div class="container mx-auto">
      <div class="text-center mb-12 md:mb-16">
        <h2 class="text-4xl sm:text-5xl font-extrabold tracking-tight text-gray-900 mb-4 relative inline-block">
          All Collections
          <span class="absolute left-1/2 transform -translate-x-1/2 bottom-0 w-24 h-1 bg-green-500 rounded-full"></span>
        </h2>
        <p class="text-lg text-gray-700 max-w-2xl mx-auto">
          Explore our curated collections designed to make your shopping easier.
        </p>
      </div>

      <!-- Loading state -->
      <div v-if="colStore.loading" class="flex items-center justify-center h-64 bg-white rounded-xl shadow-lg">
        <svg class="animate-spin h-14 w-14 text-green-600 mb-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
      </div>

      <!-- Collections Grid -->
      <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-8">
        <router-link
          v-for="c in colStore.list"
          :key="c.handle"
          :to="{ name: 'CollectionDetail', params: { handle: c.handle } }"
          class="block bg-white rounded-xl shadow-xl overflow-hidden group border border-gray-100 transition-all duration-300 hover:shadow-2xl hover:scale-103 cursor-pointer flex flex-col h-full"
        >
          <div class="w-full h-48 sm:h-56 bg-gray-100 flex items-center justify-center overflow-hidden relative">
            <img
              :src="c.image || 'https://placehold.co/400x400/E0F2F7/263238?text=Collection+Image'"
              :alt="c.title"
              class="w-full h-full object-cover object-center transition-transform duration-500 group-hover:scale-110"
              loading="lazy"
              onerror="this.onerror=null;this.src='https://placehold.co/400x400/E0F2F7/263238?text=Collection+Image';"
            />
          </div>
          <div class="p-4 flex-1 flex flex-col">
            <h3 class="text-xl font-semibold text-gray-900 mb-2 truncate group-hover:text-green-700 transition-colors duration-200" :title="c.title">{{ c.title }}</h3>
            <p class="text-sm text-gray-600 mb-4 flex-grow">{{ c.description || 'Discover a curated selection of products in this collection.' }}</p>
            <span class="inline-block text-green-600 hover:text-green-800 font-semibold mt-auto transition-colors duration-200 group-hover:underline">
              View Products &rarr;
            </span>
          </div>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useCollectionStore } from '@/stores/collection'

// derive shopSlug from host subdomain
const shopSlug = window.location.host.split('.')[0]
const colStore = useCollectionStore()

onMounted(() => {
  colStore.loadAll(shopSlug)
})
</script>

<style scoped>
/* No additional custom styles needed as Tailwind handles everything. */
</style>
