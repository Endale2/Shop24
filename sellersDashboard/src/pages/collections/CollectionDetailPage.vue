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
          @click="showEditModal = true"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 transition duration-150 ease-in-out"
        >
          <PencilIcon class="h-4 w-4 mr-2" />
          Edit Collection
        </button>
        <button
          @click="confirmDelete"
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
        <div class="flex justify-between items-center mb-6">
          <h2 class="text-2xl font-semibold text-gray-800">Products in this Collection ({{ collection.products.length }})</h2>
          <button
            @click="showAddProductModal = true"
            class="inline-flex items-center px-3 py-2 border border-transparent text-sm font-medium rounded-lg shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 transition duration-150 ease-in-out"
          >
            <PlusIcon class="h-4 w-4 mr-2" />
            Add Products
          </button>
        </div>
        
        <div v-if="collection.products.length" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          <div
            v-for="prod in collection.products"
            :key="prod.id"
            class="bg-white rounded-xl shadow-md overflow-hidden flex flex-col group relative"
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
            <div class="p-4 flex-grow flex items-center justify-between">
              <h3 class="text-lg font-medium text-gray-900 truncate flex-grow">{{ prod.name }}</h3>
              <div class="flex space-x-2">
                <button
                  @click="goToProduct(prod.id)"
                  class="p-1 text-gray-400 hover:text-green-600 transition-colors"
                  title="View Product"
                >
                  <EyeIcon class="h-4 w-4" />
                </button>
                <button
                  @click="removeProduct(prod.id)"
                  class="p-1 text-gray-400 hover:text-red-600 transition-colors"
                  title="Remove from Collection"
                >
                  <XIcon class="h-4 w-4" />
                </button>
              </div>
            </div>
          </div>
        </div>
        <div v-else class="bg-gray-50 border border-gray-200 text-gray-600 px-6 py-8 rounded-lg text-center mt-4 shadow-sm">
          <p class="text-lg font-medium">No products are linked to this collection yet.</p>
          <p class="mt-2">Click "Add Products" to start building your collection.</p>
        </div>
      </div>
    </div>

    <!-- Edit Collection Modal -->
    <div v-if="showEditModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
      <div class="bg-white rounded-xl shadow-xl max-w-2xl w-full max-h-[90vh] overflow-y-auto">
        <div class="p-6 border-b border-gray-200">
          <h3 class="text-xl font-semibold text-gray-900">Edit Collection</h3>
        </div>
        
        <form @submit.prevent="handleEditSubmit" class="p-6 space-y-6">
          <!-- Title -->
          <div>
            <label for="edit-title" class="block text-sm font-medium text-gray-700 mb-2">
              Collection Title <span class="text-red-500">*</span>
            </label>
            <input
              id="edit-title"
              v-model="editForm.title"
              type="text"
              required
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out"
              :class="{ 'border-red-500': editErrors.title }"
            />
            <p v-if="editErrors.title" class="mt-1 text-sm text-red-600">{{ editErrors.title }}</p>
          </div>

          <!-- Handle -->
          <div>
            <label for="edit-handle" class="block text-sm font-medium text-gray-700 mb-2">
              Handle <span class="text-red-500">*</span>
            </label>
            <input
              id="edit-handle"
              v-model="editForm.handle"
              type="text"
              required
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out"
              :class="{ 'border-red-500': editErrors.handle }"
            />
            <p v-if="editErrors.handle" class="mt-1 text-sm text-red-600">{{ editErrors.handle }}</p>
          </div>

          <!-- Description -->
          <div>
            <label for="edit-description" class="block text-sm font-medium text-gray-700 mb-2">
              Description
            </label>
            <textarea
              id="edit-description"
              v-model="editForm.description"
              rows="4"
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out resize-none"
            ></textarea>
          </div>

          <!-- Image URL -->
          <div>
            <label for="edit-image" class="block text-sm font-medium text-gray-700 mb-2">
              Collection Image URL
            </label>
            <input
              id="edit-image"
              v-model="editForm.image"
              type="url"
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out"
            />
          </div>

          <div class="flex justify-end space-x-4 pt-6 border-t border-gray-200">
            <button
              type="button"
              @click="showEditModal = false"
              class="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 transition duration-150 ease-in-out"
            >
              Cancel
            </button>
            <button
              type="submit"
              :disabled="editLoading"
              class="px-6 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed transition duration-150 ease-in-out inline-flex items-center"
            >
              <SpinnerIcon v-if="editLoading" class="animate-spin h-4 w-4 mr-2" />
              <CheckIcon v-else class="h-4 w-4 mr-2" />
              {{ editLoading ? 'Saving...' : 'Save Changes' }}
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Add Product Modal -->
    <div v-if="showAddProductModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
      <div class="bg-white rounded-xl shadow-xl max-w-2xl w-full max-h-[90vh] overflow-y-auto">
        <div class="p-6 border-b border-gray-200">
          <h3 class="text-xl font-semibold text-gray-900">Add Products to Collection</h3>
        </div>
        
        <div class="p-6">
          <div class="mb-4">
            <input
              v-model="productSearchQuery"
              type="text"
              placeholder="Search products..."
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out"
            />
          </div>
          
          <div class="space-y-2 max-h-96 overflow-y-auto">
            <div
              v-for="product in filteredProducts"
              :key="product.id"
              class="flex items-center justify-between p-3 border border-gray-200 rounded-lg hover:bg-gray-50"
            >
              <div class="flex items-center space-x-3">
                <div class="w-12 h-12 bg-gray-200 rounded-lg overflow-hidden flex-shrink-0">
                  <img
                    v-if="product.main_image"
                    :src="product.main_image"
                    alt="Product"
                    class="w-full h-full object-cover"
                  />
                  <PlaceholderIcon v-else class="w-full h-full text-gray-400" />
                </div>
                <div>
                  <h4 class="font-medium text-gray-900">{{ product.name }}</h4>
                  <p class="text-sm text-gray-600">{{ product.description }}</p>
                </div>
              </div>
              <button
                @click="addProduct(product.id)"
                :disabled="isProductInCollection(product.id)"
                class="px-3 py-1 text-sm rounded-lg transition duration-150 ease-in-out"
                :class="isProductInCollection(product.id) 
                  ? 'bg-gray-100 text-gray-400 cursor-not-allowed' 
                  : 'bg-green-600 text-white hover:bg-green-700'"
              >
                {{ isProductInCollection(product.id) ? 'Added' : 'Add' }}
              </button>
            </div>
          </div>
          
          <div class="flex justify-end pt-6 border-t border-gray-200 mt-6">
            <button
              @click="showAddProductModal = false"
              class="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 transition duration-150 ease-in-out"
            >
              Done
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { collectionService } from '@/services/collection'
import { productService } from '@/services/product'
import {
  ArrowLeftIcon,
  RefreshIcon as SpinnerIcon,
  PhotographIcon as PlaceholderIcon,
  PencilIcon,
  TrashIcon,
  CalendarIcon,
  CubeTransparentIcon,
  EyeIcon,
  XIcon,
  PlusIcon,
  CheckIcon
} from '@heroicons/vue/outline'

const route = useRoute()
const router = useRouter()
const shopStore = useShopStore()

// Reactive state
const collection = ref({
  id: '',
  title: '',
  handle: '',
  description: '',
  image: '',
  products: [],
  createdAt: null,
  updatedAt: null
})
const loading = ref(true)
const error = ref(null)

// Modal states
const showEditModal = ref(false)
const showAddProductModal = ref(false)

// Edit form
const editForm = ref({
  title: '',
  handle: '',
  description: '',
  image: ''
})
const editErrors = ref({})
const editLoading = ref(false)

// Product management
const allProducts = ref([])
const productSearchQuery = ref('')

// Active shop from pinia
const activeShop = computed(() => shopStore.activeShop)

// Computed
const filteredProducts = computed(() => {
  if (!productSearchQuery.value) return allProducts.value
  
  const query = productSearchQuery.value.toLowerCase()
  return allProducts.value.filter(product =>
    product.name.toLowerCase().includes(query) ||
    product.description?.toLowerCase().includes(query)
  )
})

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

  await Promise.all([
    loadCollection(collectionId),
    loadProducts()
  ])
})

/**
 * Loads collection details.
 */
async function loadCollection(collectionId) {
  try {
    const data = await collectionService.fetchById(activeShop.value.id, collectionId)
    collection.value = {
      id: data.id,
      title: data.title,
      handle: data.handle,
      description: data.description,
      image: data.image,
      products: data.products || [],
      createdAt: data.createdAt,
      updatedAt: data.updatedAt
    }
    
    // Initialize edit form
    editForm.value = {
      title: data.title,
      handle: data.handle,
      description: data.description || '',
      image: data.image || ''
    }
  } catch (e) {
    console.error(e)
    error.value = 'Failed to load collection details. The collection might not exist or there was a network error.'
  } finally {
    loading.value = false
  }
}

/**
 * Loads all products for the shop.
 */
async function loadProducts() {
  try {
    allProducts.value = await productService.fetchAllByShop(activeShop.value.id)
  } catch (e) {
    console.error('Failed to load products:', e)
  }
}

/**
 * Handles edit form submission.
 */
async function handleEditSubmit() {
  editErrors.value = {}
  
  // Validate form
  if (!editForm.value.title.trim()) {
    editErrors.value.title = 'Title is required'
  }
  if (!editForm.value.handle.trim()) {
    editErrors.value.handle = 'Handle is required'
  } else if (!/^[a-z0-9-]+$/.test(editForm.value.handle)) {
    editErrors.value.handle = 'Handle can only contain lowercase letters, numbers, and hyphens'
  }

  if (Object.keys(editErrors.value).length > 0) {
    return
  }

  editLoading.value = true

  try {
    await collectionService.update(activeShop.value.id, collection.value.id, {
      title: editForm.value.title.trim(),
      handle: editForm.value.handle.trim(),
      description: editForm.value.description.trim(),
      image: editForm.value.image.trim() || undefined
    })

    // Reload collection data
    await loadCollection(collection.value.id)
    showEditModal.value = false
  } catch (error) {
    console.error('Failed to update collection:', error)
    
    if (error.response?.data?.error) {
      const errorMsg = error.response.data.error
      if (errorMsg.includes('handle already exists')) {
        editErrors.value.handle = 'This handle is already taken. Please choose a different one.'
      } else {
        editErrors.value.general = errorMsg
      }
    } else {
      editErrors.value.general = 'Failed to update collection. Please try again.'
    }
  } finally {
    editLoading.value = false
  }
}

/**
 * Confirms and deletes the collection.
 */
async function confirmDelete() {
  if (confirm(`Are you sure you want to delete "${collection.value.title}"? This action cannot be undone.`)) {
    try {
      await collectionService.delete(activeShop.value.id, collection.value.id)
      router.push({ name: 'Collections' })
    } catch (error) {
      console.error('Failed to delete collection:', error)
      alert('Failed to delete collection. Please try again.')
    }
  }
}

/**
 * Adds a product to the collection.
 */
async function addProduct(productId) {
  try {
    await collectionService.addProduct(activeShop.value.id, collection.value.id, productId)
    // Reload collection to get updated product list
    await loadCollection(collection.value.id)
  } catch (error) {
    console.error('Failed to add product to collection:', error)
    alert('Failed to add product to collection. Please try again.')
  }
}

/**
 * Removes a product from the collection.
 */
async function removeProduct(productId) {
  if (confirm('Are you sure you want to remove this product from the collection?')) {
    try {
      await collectionService.removeProduct(activeShop.value.id, collection.value.id, productId)
      // Reload collection to get updated product list
      await loadCollection(collection.value.id)
    } catch (error) {
      console.error('Failed to remove product from collection:', error)
      alert('Failed to remove product from collection. Please try again.')
    }
  }
}

/**
 * Checks if a product is already in the collection.
 */
function isProductInCollection(productId) {
  return collection.value.products.some(p => p.id === productId)
}

/**
 * Navigates to a specific product's detail page.
 */
function goToProduct(productId) {
  router.push({ name: 'ProductDetail', params: { productId } })
}

/**
 * Formats a date string into a readable format.
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

// Watch for modal state changes
watch(showEditModal, (newVal) => {
  if (newVal) {
    // Reset edit form when opening modal
    editForm.value = {
      title: collection.value.title,
      handle: collection.value.handle,
      description: collection.value.description || '',
      image: collection.value.image || ''
    }
    editErrors.value = {}
  }
})
</script>

<style scoped>
/* No specific scoped styles needed as Tailwind classes handle the design and animations */
</style>