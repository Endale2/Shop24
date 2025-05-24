<template>
  <div class="p-4 sm:p-6 max-w-6xl mx-auto space-y-8 font-sans">
    <!-- Back Button -->
    <button
      @click="goBack"
      class="inline-flex items-center text-gray-600 hover:text-blue-700 transition duration-200 ease-in-out mb-6"
    >
      <ChevronLeftIcon class="h-5 w-5 mr-1 text-gray-600" />
      <span class="text-sm font-medium">Back to Products</span>
    </button>

    <!-- Loading State -->
    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-12">
      <SpinnerIcon class="animate-spin h-8 w-8 text-blue-500 mb-3" />
      <p class="mt-3 text-lg">Loading product details...</p>
    </div>

    <!-- Error State -->
    <div
      v-else-if="error"
      class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-lg shadow-sm"
      role="alert"
    >
      <strong class="font-bold">Oops!</strong>
      <span class="ml-2">{{ error }}</span>
      <p class="text-sm mt-2">Please try again or contact support if the issue persists.</p>
    </div>

    <!-- Product Details -->
    <div v-else class="bg-white rounded-xl shadow-lg p-6 md:p-8 space-y-8">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 lg:gap-12">
        <!-- Images Gallery -->
        <div class="space-y-4">
          <img
            v-for="(img, idx) in product.images"
            :key="idx"
            :src="img"
            alt="Product image"
            class="w-full h-72 object-cover rounded-lg shadow-sm border border-gray-200"
          />
          <div v-if="!product.images.length" class="w-full h-72 flex items-center justify-center bg-gray-100 rounded-lg border border-gray-200">
            <PlaceholderIcon class="w-12 h-12 text-gray-400" />
          </div>
        </div>

        <!-- Info Panel -->
        <div class="lg:sticky lg:top-8 self-start space-y-6">
          <div>
            <h1 class="text-4xl font-extrabold text-gray-900 leading-tight">{{ product.name }}</h1>
            <p class="text-lg text-gray-600 mt-2">{{ product.category || 'Uncategorized' }}</p>
          </div>

          <p class="text-4xl font-bold text-green-700">
            {{ formatPrice(product.price) }}
          </p>

          <p class="text-gray-700 leading-relaxed text-base">
            {{ product.description || 'No description available.' }}
          </p>

          <div v-if="product.variants.length" class="mt-6 border-t border-gray-200 pt-6">
            <h3 class="font-semibold text-gray-800 text-xl mb-4">Available Variants</h3>
            <ul class="space-y-3">
              <li
                v-for="(v, i) in product.variants"
                :key="i"
                class="flex flex-col sm:flex-row sm:justify-between items-start sm:items-center p-4 bg-gray-50 rounded-lg shadow-sm border border-gray-100"
              >
                <span class="text-gray-700 font-medium text-base mb-1 sm:mb-0">
                  {{ formatVariant(v.options) }}
                </span>
                <span class="font-bold text-gray-900 text-lg">
                  {{ formatPrice(v.price) }}
                </span>
              </li>
            </ul>
          </div>
        </div>
      </div>

      <!-- Timestamps -->
      <div class="border-t border-gray-200 pt-6 text-sm text-gray-500 text-right space-y-1">
        <p><strong>Created:</strong> {{ formatDate(product.createdAt) }}</p>
        <p><strong>Last Updated:</strong> {{ formatDate(product.updatedAt) }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { productService } from '@/services/product'
import {
  ChevronLeftIcon,
  RefreshIcon as SpinnerIcon,
  PhotographIcon as PlaceholderIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const route = useRoute()
const shopStore = useShopStore()

const product = ref({ images: [], variants: [] })
const loading = ref(true)
const error = ref(null)

const activeShop = computed(() => shopStore.activeShop)

function goBack() {
  router.push({ name: 'Products' })
}

function formatDate(dt) {
  return dt
    ? new Date(dt).toLocaleDateString(undefined, {
        year: 'numeric', month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit'
      })
    : 'Not available'
}

function formatPrice(p) {
  return p != null ? `$${p.toFixed(2)}` : 'N/A'
}

function formatVariant(options) {
  return Object.entries(options).map(([k, v]) => `${k}: ${v}`).join(', ')
}

onMounted(async () => {
  if (!activeShop.value) {
    error.value = 'No shop selected. Please select a shop to view product details.'
    loading.value = false
    return
  }

  const productId = route.params.productId
  try {
    product.value = await productService.fetchById(activeShop.value.id, productId)
  } catch (e) {
    console.error(e)
    error.value = 'Failed to load product details. Please try again later.'
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
@keyframes spin { from { transform: rotate(0deg) } to { transform: rotate(360deg) } }
.animate-spin { animation: spin 1s linear infinite }
</style>
