<template>
  <div class="p-4 sm:p-6 max-w-4xl mx-auto font-sans">
    <div class="flex justify-between items-center mb-6">
      <button
        @click="$router.back()"
        class="inline-flex items-center text-gray-600 hover:text-green-700 transition duration-150 ease-in-out font-medium rounded-lg px-3 py-1.5"
      >
        <ArrowLeftIcon class="h-5 w-5 mr-1" />
        <span class="text-sm">Back to Collections</span>
      </button>

      <div class="flex space-x-3">
        <button
          @click="editCollection"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 transition duration-150 ease-in-out"
        >
          <PencilIcon class="h-4 w-4 mr-2" />
          Edit Collection
        </button>
        <button
          @click="deleteCollection"
          class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-lg shadow-sm text-red-700 bg-white hover:bg-red-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 transition duration-150 ease-in-out"
        >
          <TrashIcon class="h-4 w-4 mr-2" />
          Delete
        </button>
      </div>
    </div>

    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-16">
      <SpinnerIcon class="animate-spin h-8 w-8 text-green-500 mb-3" />
      <p class="text-lg">Loading collection details...</p>
    </div>

    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-lg shadow-sm" role="alert">
      <strong class="font-bold">Error:</strong>
      <span class="block sm:inline ml-2">{{ error }}</span>
      <p class="text-sm mt-2">Could not load collection details. Please try again.</p>
    </div>

    <div v-else class="bg-white rounded-xl shadow-lg p-6 sm:p-8 space-y-8">
      <div class="flex flex-col md:flex-row md:items-start md:space-x-8">
        <div class="flex-shrink-0 mb-6 md:mb-0 w-full md:w-1/3">
          <div class="aspect-w-16 aspect-h-9 md:aspect-w-1 md:aspect-h-1 rounded-xl overflow-hidden shadow-md border border-gray-200 bg-gray-100 flex items-center justify-center">
            <img
              v-if="collection.image"
              :src="collection.image"
              alt="Collection banner"
              class="w-full h-full object-cover"
            />
            <PlaceholderIcon v-else class="w-24 h-24 text-gray-400" />
          </div>
        </div>
        <div class="flex-grow">
          <h1 class="text-4xl sm:text-5xl font-extrabold text-gray-900 leading-tight mb-2">
            {{ collection.title }}
          </h1>
          <p class="text-lg text-gray-600 font-mono mb-4 bg-gray-50 px-3 py-1 rounded-md inline-block">
            #{{ collection.handle }}
          </p>
          <p class="text-gray-700 leading-relaxed mb-6">
            {{ collection.description || 'No description provided for this collection.' }}
          </p>
          <div class="flex items-center text-gray-500 text-sm">
            <CalendarIcon class="h-4 w-4 mr-1.5" />
            Created: {{ formatDate(collection.createdAt) }}
          </div>
          <div class="flex items-center text-gray-500 text-sm mt-2">
            <CubeTransparentIcon class="h-4 w-4 mr-1.5" />
            Last Updated: {{ formatDate(collection.updatedAt) }}
          </div>
        </div>
      </div>

      <hr class="border-gray-200">

      <div>
        <h2 class="text-2xl font-semibold text-gray-800 mb-6">Products in this Collection ({{ collection.products.length }})</h2>
        <div v-if="collection.products.length" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          <div
            v-for="prod in collection.products"
            :key="prod.id"
            @click="goToProduct(prod.id)"
            class="bg-white rounded-xl shadow-md overflow-hidden cursor-pointer transform hover:scale-105 hover:shadow-xl transition duration-200 ease-in-out flex flex-col group"
          >
            <div class="w-full h-40 bg-gray-100 flex items-center justify-center relative overflow-hidden">
              <img
                v-if="prod.image"
                :src="prod.image"
                alt="Product image"
                class="w-full h-full object-cover transform group-hover:scale-110 transition-transform duration-300 ease-in-out"
              />
              <PlaceholderIcon v-else class="w-16 h-16 text-gray-400" />
            </div>
            <div class="p-4 flex-grow flex items-center">
              <h3 class="text-lg font-medium text-gray-900 truncate flex-grow">{{ prod.name }}</h3>
              <ArrowRightIcon class="h-5 w-5 text-gray-400 group-hover:text-green-600 ml-2 transition-colors" />
            </div>
          </div>
        </div>
        <div v-else class="bg-gray-50 border border-gray-200 text-gray-600 px-6 py-8 rounded-lg text-center mt-4 shadow-sm">
          <p class="text-lg font-medium">No products are linked to this collection yet.</p>
          <p class="mt-2">You can add products to this collection from the product details page.</p>
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
  PhotographIcon as PlaceholderIcon,
  PencilIcon, // For edit button
  TrashIcon,  // For delete button
  CalendarIcon, // For created date
  CubeTransparentIcon, // For updated date
  ArrowRightIcon // For product card navigation
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
  products: [],
  createdAt: null, // Added for display
  updatedAt: null  // Added for display
})
const loading = ref(true)
const error = ref(null)

// Active shop from pinia
const activeShop = computed(() => shopStore.activeShop)

onMounted(async () => {
  if (!activeShop.value?.id) {
    error.value = 'No shop selected. Please select a shop to view collection details.'
    loading.value = false
    return
  }
  const collectionId = route.params.collectionId
  if (!collectionId) {
    error.value = 'Collection ID is missing in the URL.'
    loading.value = false
    return
  }

  try {
    const data = await collectionService.fetchById(activeShop.value.id, collectionId)
    collection.value = {
      title: data.title,
      handle: data.handle,
      description: data.description,
      image: data.image,
      products: data.products || [], // Ensure it's an array
      createdAt: data.createdAt,
      updatedAt: data.updatedAt
    }
  } catch (e) {
    console.error(e)
    error.value = 'Failed to load collection details. The collection might not exist or there was a network error.'
  } finally {
    loading.value = false
  }
})

/**
 * Navigates to a specific product's detail page.
 * @param {string} productId - The ID of the product.
 */
function goToProduct(productId) {
  router.push({ name: 'ProductDetail', params: { productId } })
}

/**
 * Placeholder function for editing the collection.
 * You'll implement the actual edit logic here later.
 */
function editCollection() {
  alert('Edit collection functionality coming soon!')
  // router.push({ name: 'EditCollection', params: { collectionId: collection.value.id } })
}

/**
 * Placeholder function for deleting the collection.
 * You'll implement the actual delete logic here later.
 */
function deleteCollection() {
  if (confirm(`Are you sure you want to delete "${collection.value.title}"?`)) {
    // Implement delete logic here, e.g.:
    // await collectionService.delete(activeShop.value.id, collection.value.id)
    // router.push({ name: 'Collections' }) // Go back to collections list after deletion
    alert(`Deleting "${collection.value.title}"... (functionality not fully implemented)`)
  }
}

/**
 * Formats a date string into a readable format.
 * @param {string} dateString - The date string to format.
 * @returns {string} The formatted date.
 */
function formatDate(dateString) {
  if (!dateString) return 'N/A';
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  });
}
</script>

<style scoped>
/* No specific scoped styles needed as Tailwind classes handle the design and animations */
</style>