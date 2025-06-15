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

      <!-- Loading Spinner -->
      <div v-if="colStore.loading" class="flex items-center justify-center h-64 bg-white rounded-xl shadow-lg">
        <svg class="animate-spin h-14 w-14 text-green-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
          <path class="opacity-75" fill="currentColor"
            d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
        </svg>
      </div>

      <!-- Collections Grid -->
      <div v-else-if="colStore.list && colStore.list.length > 0" class="grid grid-cols-1 sm:grid-cols-2 gap-6">
        <router-link
          v-for="c in colStore.list"
          :key="c.handle"
          :to="{ name: 'CollectionDetail', params: { handle: c.handle } }"
          class="group flex bg-white rounded-2xl shadow-xl hover:shadow-2xl transition-all duration-300 overflow-hidden border border-gray-100 transform hover:scale-[1.02] hover:-translate-y-1"
        >
          <!-- Image Section (left side for sm and up, top for col) -->
          <div class="w-full h-48 sm:w-48 sm:h-auto flex-shrink-0 bg-gray-100 overflow-hidden">
            <img
              :src="c.image || 'https://placehold.co/400x400/E0F2F7/263238?text=Collection+Image'"
              :alt="c.title"
              class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
              loading="lazy"
              onerror="this.onerror=null;this.src='https://placehold.co/400x400/E0F2F7/263238?text=Collection+Image';"
            />
          </div>

          <!-- Content Section (right side for sm and up, bottom for col) -->
          <div class="p-5 flex flex-col justify-between flex-1">
            <div>
              <h3
                class="text-xl font-bold text-gray-900 mb-2 truncate group-hover:text-green-700 transition-colors duration-200"
                :title="c.title"
              >
                {{ c.title }}
              </h3>
              <p class="text-sm text-gray-600 mb-4 line-clamp-3">
                {{ c.description || 'Discover a curated selection of products in this collection.' }}
              </p>
            </div>
            <span class="text-green-600 hover:text-green-800 font-semibold mt-auto transition-colors duration-200 group-hover:underline">
              View Products &rarr;
            </span>
          </div>
        </router-link>
      </div>

      <!-- No Collections Found state -->
      <div v-else class="text-center py-16 md:py-20 bg-white rounded-xl shadow-lg">
        <p class="text-lg font-semibold text-gray-700 mb-2">No Collections Found</p>
        <p class="text-base text-gray-500">It looks like we don't have any collections yet. Please check back soon!</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useCollectionStore } from '@/stores/collection'

// Derive shopSlug from host subdomain
const shopSlug = window.location.host.split('.')[0]
const colStore = useCollectionStore()

onMounted(() => {
  colStore.loadAll(shopSlug)
})
</script>

<style scoped>
/* Tailwind handles all styles, scoped here just in case */
/* Custom style for line-clamp if not natively supported by Tailwind JIT or older versions */
.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
