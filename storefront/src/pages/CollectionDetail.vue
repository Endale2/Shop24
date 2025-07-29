<template>
  <!-- Breadcrumb and Back Button -->
  <nav class="flex items-center space-x-2 text-sm text-gray-500 mb-6">
    <button @click="$router.back()" class="text-gray-700 hover:text-black flex items-center gap-1 font-medium text-xs mr-2">
      <svg class="w-4 h-4" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7"/></svg>
      Back
    </button>
    <router-link to="/" class="hover:underline">Home</router-link>
    <span>/</span>
    <router-link to="/collections" class="hover:underline">Collections</router-link>
    <span>/</span>
    <span>{{ collection?.title || '' }}</span>
  </nav>
  <Loader v-if="isLoading" text="Loading collection..." />
  <div v-else-if="collection" class="space-y-10">
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
      <div v-if="collection.products && collection.products.length > 0" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8">
        <div v-for="p in collection.products" :key="p.id" class="relative">
          <!-- Discount Badge -->
          <span v-if="p.discounts && p.discounts.length > 0" class="absolute top-2 left-2 inline-block px-3 py-1 bg-yellow-100 text-yellow-800 text-xs font-semibold rounded-full z-10">
            {{ p.discounts[0].type === 'percentage' ? `${p.discounts[0].value}% OFF` : `$${p.discounts[0].value} OFF` }}
          </span>
          <ProductCard :product="p" />
        </div>
      </div>
      <div v-else class="text-center py-12">
        <p class="text-gray-600">No products found in this collection.</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { fetchCollectionDetail, CollectionDetail } from '@/services/collections'
import ProductCard from '@/components/ProductCard.vue'
import Loader from '@/components/Loader.vue';

const route = useRoute()
const handle = route.params.handle as string
const shopSlug = route.params.shopSlug as string
const collection = ref<CollectionDetail | null>(null)
const isLoading = ref(true)

onMounted(async () => {
  if (!shopSlug || !handle) return;
  isLoading.value = true
  collection.value = await fetchCollectionDetail(shopSlug, handle)
  isLoading.value = false
})
</script>
