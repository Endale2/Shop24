<!-- src/pages/collections/CollectionDetail.vue -->
<template>
  <div class="p-4 sm:p-6 max-w-4xl mx-auto font-sans">
    <!-- Back link -->
    <button
      @click="$router.back()"
      class="inline-flex items-center text-gray-600 hover:text-blue-700 transition mb-6"
    >
      <ArrowLeftIcon class="h-5 w-5 mr-1" />
      <span class="text-sm font-medium">Back to Collections</span>
    </button>

    <!-- Loading -->
    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-16">
      <SpinnerIcon class="animate-spin h-8 w-8 text-blue-500 mb-3" />
      <p class="text-lg">Loading collection...</p>
    </div>

    <!-- Error -->
    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-lg shadow-sm" role="alert">
      <strong class="font-bold">Error:</strong>
      <span class="block sm:inline ml-2">{{ error }}</span>
    </div>

    <!-- Collection details -->
    <div v-else class="space-y-8">
      <!-- Header -->
      <div class="flex flex-col md:flex-row md:items-center md:space-x-6">
        <div class="flex-shrink-0 mb-6 md:mb-0">
          <img
            v-if="collection.image"
            :src="collection.image"
            alt="Collection banner"
            class="w-48 h-48 object-cover rounded-lg shadow"
          />
          <PlaceholderIcon v-else class="w-48 h-48 text-gray-300" />
        </div>
        <div>
          <h1 class="text-4xl font-extrabold text-gray-900">{{ collection.title }}</h1>
          <p class="mt-2 text-gray-600">{{ collection.handle }}</p>
          <p class="mt-4 text-gray-700">{{ collection.description }}</p>
        </div>
      </div>

      <!-- Products grid -->
      <div>
        <h2 class="text-2xl font-semibold text-gray-800 mb-4">Products in this Collection</h2>
        <div v-if="collection.products.length" class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
          <div
            v-for="prod in collection.products"
            :key="prod.id"
            @click="goToProduct(prod.id)"
            class="bg-white rounded-xl shadow-lg overflow-hidden cursor-pointer transform hover:scale-105 hover:shadow-xl transition duration-200 ease-in-out"
          >
            <div class="w-full h-40 bg-gray-100 flex items-center justify-center">
              <img
                v-if="prod.main_image"
                :src="prod.main_image"
                alt=""
                class="w-full h-full object-cover"
              />
              <PlaceholderIcon v-else class="w-16 h-16 text-gray-400" />
            </div>
            <div class="p-4">
              <h3 class="text-lg font-medium text-gray-900 truncate">{{ prod.name }}</h3>
            </div>
          </div>
        </div>
        <div v-else class="text-center text-gray-600 py-8">
          <p>No products are linked to this collection.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { collectionService } from '@/services/collection'
import {
  ArrowLeftIcon,
  RefreshIcon as SpinnerIcon,
  PhotographIcon as PlaceholderIcon
} from '@heroicons/vue/outline'

const route = useRoute()
const router = useRouter()
const shopStore = useShopStore()

// Reactive state
const collection = ref({
  title: '',
  handle: '',
  description: '',
  image: '',
  products: []
})
const loading = ref(true)
const error = ref(null)

// Active shop from pinia
const activeShop = computed(() => shopStore.activeShop)

onMounted(async () => {
  if (!activeShop.value?.id) {
    error.value = 'No shop selected.'
    loading.value = false
    return
  }
  const collectionId = route.params.collectionId
  try {
    const data = await collectionService.fetchById(activeShop.value.id, collectionId)
    collection.value = {
      title: data.title,
      handle: data.handle,
      description: data.description,
      image: data.image,
      products: data.products // array of { id, name, main_image }
    }
  } catch (e) {
    console.error(e)
    error.value = 'Failed to load collection details.'
  } finally {
    loading.value = false
  }
})

function goToProduct(productId) {
  router.push({ name: 'ProductDetail', params: { productId } })
}
</script>

<style scoped>
@keyframes spin { from { transform: rotate(0deg); } to { transform: rotate(360deg); } }
.animate-spin { animation: spin 1s linear infinite; }
</style>
