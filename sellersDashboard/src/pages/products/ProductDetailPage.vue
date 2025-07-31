<template>
  <div class="p-4 sm:p-6 max-w-6xl mx-auto space-y-8 font-sans">
    <button
      @click="goBack"
      class="inline-flex items-center text-gray-600 hover:text-green-700 transition duration-200 ease-in-out mb-6 group rounded-full px-3 py-1.5 -ml-3"
    >
      <ChevronLeftIcon class="h-5 w-5 mr-1 text-gray-500 group-hover:text-green-600 transition-colors duration-200" />
      <span class="text-sm font-medium group-hover:text-green-700 transition-colors duration-200">Back to Products</span>
    </button>

    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-12">
      <SpinnerIcon class="animate-spin h-8 w-8 text-green-500 mb-3" />
      <p class="mt-3 text-lg">Loading product details...</p>
    </div>

    <div
      v-else-if="error"
      class="bg-red-50 border border-red-200 text-red-700 px-6 py-4 rounded-lg shadow-sm"
      role="alert"
    >
      <div class="flex items-center">
        <ExclamationCircleIcon class="h-5 w-5 mr-2 flex-shrink-0" />
        <div>
          <strong class="font-semibold">Error:</strong>
          <span class="ml-1">{{ error }}</span>
        </div>
      </div>
      <p class="text-sm mt-2">Please try again or contact support if the issue persists.</p>
    </div>

    <div v-else-if="product.id" class="bg-white rounded-xl shadow-lg p-6 md:p-8 space-y-8 animate-fade-in">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 lg:gap-12">
        <!-- Product Images -->
        <div class="space-y-4">
          <div class="w-full h-80 sm:h-96 md:h-80 lg:h-96 rounded-xl overflow-hidden shadow-md border border-gray-200 relative bg-gray-50 flex items-center justify-center">
            <img
              v-if="mainImage"
              :src="mainImage"
              alt="Main product image"
              class="max-w-full max-h-full object-contain transition-transform duration-300 ease-in-out transform hover:scale-105"
            />
            <div v-else class="w-full h-full flex items-center justify-center bg-gray-100">
              <PlaceholderIcon class="w-20 h-20 text-gray-400" />
            </div>
          </div>

          <div v-if="allImages.length > 1" class="flex space-x-3 overflow-x-auto pb-2">
            <img
              v-for="(img, idx) in allImages"
              :key="`all-img-${idx}`"
              :src="img"
              alt="Product image"
              @click="setMainImage(img)"
              :class="{ 'border-green-600 ring-2 ring-green-600': mainImage === img, 'border-gray-300': mainImage !== img }"
              class="w-20 h-20 object-cover rounded-lg cursor-pointer border-2 shadow-sm transform hover:scale-105 transition-all duration-200 ease-in-out"
            />
          </div>
        </div>

        <!-- Product Details -->
        <div class="lg:sticky lg:top-8 self-start space-y-6">
          <div>
            <h1 class="text-4xl sm:text-5xl font-extrabold text-gray-900 leading-tight">{{ product.name }}</h1>
            <div class="text-lg text-gray-600 mt-2">
              <div v-if="product.collection_ids && product.collection_ids.length > 0" class="space-y-1">
                <div v-for="collectionId in product.collection_ids" :key="collectionId" class="inline-block bg-gray-100 px-3 py-1 rounded-full text-sm mr-2 mb-1">
                  {{ collectionMap[collectionId] || 'Unknown Collection' }}
                </div>
              </div>
              <div v-else class="text-gray-500 italic">Uncategorized</div>
            </div>
          </div>

          <div class="space-y-4">
            <div class="flex items-baseline space-x-2">
              <p class="text-4xl sm:text-5xl font-bold text-green-700">
                <template v-if="product.starting_price != null">
                  <span class="italic text-gray-700">from</span> {{ formatPrice(product.starting_price) }}
                </template>
                <template v-else-if="product.price != null">
                  {{ formatPrice(product.price) }}
                </template>
                <template v-else>
                  N/A
                </template>
              </p>
            </div>

            <div class="flex items-center space-x-4">
              <div class="flex items-center space-x-2">
                <span class="text-gray-700 font-medium">Stock:</span>
                <span 
                  :class="{
                    'text-green-600 font-semibold': computedStockDisplay > 0,
                    'text-red-600 font-semibold': computedStockDisplay <= 0
                  }"
                >
                  {{ computedStockDisplay }}
                </span>
              </div>
              
              <span
                :class="{
                  'px-3 py-1 text-sm font-medium rounded-full': true,
                  'bg-green-100 text-green-800': computedStockDisplay > 0,
                  'bg-red-100 text-red-800': computedStockDisplay <= 0
                }"
              >
                {{ computedStockDisplay > 0 ? 'In Stock' : 'Out of Stock' }}
              </span>
            </div>
          </div>

          <div v-if="product.description" class="mt-6">
            <h3 class="text-lg font-semibold text-gray-800 mb-2">Description</h3>
            <p class="text-gray-700 leading-relaxed text-base">
              {{ product.description }}
            </p>
          </div>

          <!-- Variants Section -->
          <div v-if="product.variants && product.variants.length > 0" class="mt-8 border-t border-gray-200 pt-6">
            <h3 class="text-xl font-semibold text-gray-800 mb-4">Available Variants</h3>
            <div class="space-y-3">
              <div
                v-for="(v, i) in product.variants"
                :key="i"
                class="flex flex-col sm:flex-row sm:justify-between items-start sm:items-center p-4 bg-green-50 rounded-lg shadow-sm border border-green-100 transition-all duration-200 ease-in-out hover:shadow-md hover:border-green-200"
              >
                <div class="flex items-center gap-3 mb-3 sm:mb-0">
                  <img
                    v-if="v.image"
                    :src="v.image"
                    alt="Variant thumbnail"
                    class="w-12 h-12 object-cover rounded-md border border-gray-200"
                  />
                  <div>
                    <span class="text-gray-700 font-medium text-base">
                      <template v-if="v.options && v.options.length">
                        <span v-for="(opt, oi) in v.options" :key="oi">
                          {{ opt.name }}: {{ opt.value }}<span v-if="oi < v.options.length - 1">, </span>
                        </span>
                      </template>
                      <template v-else>
                        No options
                      </template>
                    </span>
                  </div>
                </div>
                <div class="text-right">
                  <span class="font-bold text-green-800 text-lg sm:text-xl">
                    {{ formatPrice(v.price) }}
                  </span>
                  <p class="text-sm font-medium" :class="{ 'text-green-600': v.stock > 0, 'text-red-600': v.stock <= 0 }">
                    {{ v.stock > 0 ? `${v.stock} in stock` : 'Out of Stock' }}
                  </p>
                </div>
              </div>
            </div>
          </div>

          <!-- Action Buttons -->
          <div class="flex flex-col sm:flex-row space-y-3 sm:space-y-0 sm:space-x-4 mt-8">
            <button
              @click="goToEditProduct(product.id)"
              class="flex-1 w-full sm:w-auto inline-flex items-center justify-center px-6 py-3 border border-transparent text-base font-medium rounded-lg shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 transition duration-150 ease-in-out transform hover:-translate-y-0.5"
            >
              <PencilIcon class="w-5 h-5 mr-2" />
              Edit Product
            </button>
            <button
              @click="confirmDeleteProduct"
              class="flex-1 w-full sm:w-auto inline-flex items-center justify-center px-6 py-3 border border-red-300 text-base font-medium rounded-lg shadow-sm text-red-700 bg-white hover:bg-red-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 transition duration-150 ease-in-out"
            >
              <TrashIcon class="w-5 h-5 mr-2" />
              Delete Product
            </button>
          </div>
        </div>
      </div>

      <!-- Product Metadata -->
      <div class="border-t border-gray-200 pt-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-sm text-gray-500">
          <div>
            <p><strong>Created:</strong> {{ formatDate(product.createdAt) }}</p>
            <p><strong>Last Updated:</strong> {{ formatDate(product.updatedAt) }}</p>
            <p><strong>Collections:</strong> 
              <span v-if="product.collection_ids && product.collection_ids.length > 0">
                <span v-for="(collectionId, index) in product.collection_ids" :key="collectionId">
                  {{ collectionMap[collectionId] || 'Unknown Collection' }}{{ index < product.collection_ids.length - 1 ? ', ' : '' }}
                </span>
              </span>
              <span v-else class="text-gray-500 italic">Uncategorized</span>
            </p>
          </div>
          <div>
            <p><strong>Product ID:</strong> {{ product.id }}</p>
            <p><strong>Shop ID:</strong> {{ product.shop_id }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Not Found State -->
    <div v-else class="bg-gray-50 border border-gray-200 text-gray-700 px-6 py-8 rounded-lg text-center mt-8 shadow-sm">
      <div class="flex flex-col items-center">
        <div class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center mb-4">
          <ExclamationCircleIcon class="w-8 h-8 text-gray-400" />
        </div>
        <p class="text-lg font-medium mb-2">Product not found</p>
        <p class="text-sm">The product ID might be invalid or the product does not exist for the active shop.</p>
        <button 
          @click="goBack" 
          class="mt-4 inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
        >
          Go to Products List
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { productService } from '@/services/product'
import { collectionService } from '@/services/collection'
import {
  ChevronLeftIcon,
  RefreshIcon as SpinnerIcon,
  PhotographIcon as PlaceholderIcon,
  PencilIcon,
  TrashIcon,
  ExclamationCircleIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const route = useRoute()
const shopStore = useShopStore()

const product = ref({ images: [], variants: [] })
const loading = ref(true)
const error = ref(null)
const mainImage = ref(null)
const collections = ref([])
const collectionMap = computed(() => {
  const map = {}
  for (const coll of collections.value) {
    map[coll.id] = coll.title
  }
  return map
})

const activeShop = computed(() => shopStore.activeShop)

const computedStockDisplay = computed(() => {
  if (product.value.total_stock !== undefined) return product.value.total_stock
  if (product.value.stock !== undefined) return product.value.stock
  if (product.value.variants && product.value.variants.length > 0) {
    return product.value.variants.reduce((sum, v) => sum + (v.stock || 0), 0)
  }
  return 0
})

const allImages = computed(() => {
  const imgs = []
  if (product.value.main_image) imgs.push(product.value.main_image)
  if (Array.isArray(product.value.images)) {
    for (const img of product.value.images) {
      if (img && !imgs.includes(img)) imgs.push(img)
    }
  }
  if (Array.isArray(product.value.variants)) {
    for (const v of product.value.variants) {
      if (v && v.image && !imgs.includes(v.image)) imgs.push(v.image)
    }
  }
  return imgs
})

function goBack() {
  router.push({ name: 'Products' })
}

function goToEditProduct(productId) {
  router.push({ name: 'EditProduct', params: { productId } })
}

async function confirmDeleteProduct() {
  if (confirm(`Are you sure you want to delete "${product.value.name}"? This action cannot be undone.`)) {
    await deleteProduct()
  }
}

async function deleteProduct() {
  if (!activeShop.value?.id || !product.value?.id) {
    alert('Missing shop or product ID for deletion.')
    return
  }

  loading.value = true
  error.value = null
  try {
    await productService.delete(activeShop.value.id, product.value.id)
    alert('Product deleted successfully!')
    router.push({ name: 'Products' })
  } catch (e) {
    console.error('Failed to delete product:', e)
    error.value = `Failed to delete product: ${e.response?.data?.message || e.message}. Please try again.`
    alert(error.value)
  } finally {
    loading.value = false
  }
}

function formatDate(dt) {
  return dt
    ? new Date(dt).toLocaleDateString(undefined, {
        year: 'numeric', 
        month: 'long', 
        day: 'numeric', 
        hour: '2-digit', 
        minute: '2-digit'
      })
    : 'Not available'
}

function formatPrice(p) {
  return p != null ? `$${p.toFixed(2)}` : 'N/A'
}

function setMainImage(imageUrl) {
  mainImage.value = imageUrl
}

async function fetchProductDetails() {
  if (!activeShop.value) {
    error.value = 'No shop selected. Please select a shop to view product details.'
    loading.value = false
    return
  }

  const productId = route.params.productId
  if (!productId) {
    error.value = 'Product ID is missing.'
    loading.value = false
    return
  }

  loading.value = true
  error.value = null
  try {
    const fetchedProduct = await productService.fetchById(activeShop.value.id, productId)
    
         // Map backend fields to frontend
     product.value = {
       id: fetchedProduct.id || fetchedProduct._id,
       name: fetchedProduct.name,
       slug: fetchedProduct.slug,
       images: fetchedProduct.images || [],
       main_image: fetchedProduct.main_image,
       collection_ids: fetchedProduct.collection_ids || [],
       price: fetchedProduct.price,
       stock: fetchedProduct.stock,
       starting_price: fetchedProduct.starting_price,
       total_stock: fetchedProduct.total_stock,
       description: fetchedProduct.description,
       createdAt: fetchedProduct.createdAt,
       updatedAt: fetchedProduct.updatedAt,
       shop_id: fetchedProduct.shopId,
       variants: (fetchedProduct.variants || []).map(v => ({
         ...v,
         options: v.options || [],
         price: v.price,
         stock: v.stock,
         image: v.image
       }))
     }
    
    // Set the main image
    if (allImages.value.length > 0) {
      mainImage.value = allImages.value[0]
    } else {
      mainImage.value = null
    }
  } catch (e) {
    console.error('Failed to load product details:', e)
    error.value = 'Failed to load product details. Please try again later.'
  } finally {
    loading.value = false
  }
}

onMounted(async () => {
  fetchProductDetails()
  collections.value = await collectionService.fetchAllByShop(shopStore.activeShop.id)
})

watch([() => activeShop.value?.id, () => route.params.productId], () => {
  fetchProductDetails()
})
</script>

<style scoped>
@keyframes fadeIn {
  from { 
    opacity: 0; 
    transform: translateY(10px); 
  }
  to { 
    opacity: 1; 
    transform: translateY(0); 
  }
}

.animate-fade-in {
  animation: fadeIn 0.5s ease-out forwards;
}
</style>