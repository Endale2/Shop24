<template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto font-sans">
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center mb-8">
      <button
        @click="$router.back()"
        class="inline-flex items-center text-gray-600 hover:text-green-700 transition duration-200 ease-in-out mb-4 sm:mb-0 group rounded-full px-3 py-1.5"
      >
        <ArrowLeftIcon class="h-5 w-5 mr-1 text-gray-500 group-hover:text-green-600 transition-colors duration-200" />
        <span class="text-sm font-medium group-hover:text-green-700 transition-colors duration-200">Back to Collections</span>
      </button>
      <div class="flex flex-col sm:flex-row gap-3 mt-4 sm:mt-0">
        <button
          @click="goToEditPage"
          class="inline-flex items-center px-4 py-2.5 bg-green-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
        >
          <PencilIcon class="h-4 w-4 mr-2" />
          Edit Collection
        </button>
        <button
          @click="confirmDelete"
          class="inline-flex items-center px-4 py-2.5 border border-gray-300 text-sm font-medium rounded-lg shadow-sm text-red-700 bg-white hover:bg-red-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 transition duration-150 ease-in-out"
        >
          <TrashIcon class="h-4 w-4 mr-2" />
          Delete
        </button>
      </div>
    </div>
    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-20 bg-white rounded-2xl shadow-lg">
      <SpinnerIcon class="animate-spin h-10 w-10 text-green-500 mb-4" />
      <p class="mt-3 text-lg font-semibold text-gray-700">Loading collection details...</p>
    </div>
    <div v-else-if="error" class="bg-red-50 border border-red-200 text-red-700 px-6 py-4 rounded-lg shadow-sm mb-6" role="alert">
      <div class="flex items-center">
        <ExclamationCircleIcon class="h-5 w-5 mr-2 flex-shrink-0" />
        <div>
          <strong class="font-semibold">Error:</strong>
          <span class="ml-1">{{ error }}</span>
        </div>
      </div>
      <p class="text-sm mt-2">Could not load collection details. Please try again.</p>
    </div>
    <div v-else class="bg-white rounded-2xl shadow-xl p-6 sm:p-8 space-y-8 animate-fade-in">
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
        <div class="flex flex-col sm:flex-row justify-between items-center mb-6 gap-3">
          <h2 class="text-2xl font-semibold text-gray-800">Products in this Collection ({{ collection.products.length }})</h2>
          <button
            @click="showAddProductModal = true"
            class="inline-flex items-center px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
          >
            <PlusIcon class="h-5 w-5 mr-2" />
            Add Products
          </button>
        </div>
        <div v-if="collection.products.length">
          <!-- List/Table view for products -->
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Image</th>
                  <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Name</th>
                  <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Collection</th>
                  <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Price</th>
                  <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Stock</th>
                  <th class="px-4 py-3"></th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-100">
                <tr v-for="prod in reversedProducts" :key="prod.id" class="hover:bg-green-50 cursor-pointer group" @click="goToProduct(prod.id)">
                  <td class="px-4 py-3 w-20">
                    <img v-if="prod.image" :src="prod.image" alt="Product image" class="w-14 h-14 object-cover rounded-md border border-gray-200" />
                    <PlaceholderIcon v-else class="w-10 h-10 text-gray-400 mx-auto" />
                  </td>
                  <td class="px-4 py-3 font-medium text-gray-900">{{ prod.name }}</td>
                  <td class="px-4 py-3 text-gray-700">{{ collection.title || 'Uncategorized' }}</td>
                  <td class="px-4 py-3 text-gray-700">
                    <template v-if="prod.starting_price !== null">
                      <span class="text-green-600 font-semibold">from</span>
                      <span class="font-bold"> ${{ prod.starting_price.toFixed(2) }}</span>
                    </template>
                    <template v-else>
                      {{ prod.price !== null ? `$${prod.price.toFixed(2)}` : '-' }}
                    </template>
                  </td>
                  <td class="px-4 py-3 text-gray-700">{{ prod.stock !== null ? prod.stock : '-' }}</td>
                  <td class="px-4 py-3 text-right">
                    <button
                      @click.stop="removeProduct(prod.id)"
                      class="inline-flex items-center justify-center bg-white border border-gray-200 rounded-full p-1.5 text-red-600 hover:text-white hover:bg-red-600 transition-colors duration-150 focus:outline-none shadow"
                      title="Remove product from collection"
                    >
                      <XIcon class="h-5 w-5" />
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
        <div v-else class="text-center py-16 text-gray-500">
          <p class="text-lg font-semibold">No products are linked to this collection yet.</p>
          <p class="mt-2 text-sm">Click "Add Products" to start building your collection.</p>
        </div>
      </div>
    </div>
  </div>

  <!-- Add Product Modal -->
  <div v-if="showAddProductModal" class="fixed inset-0 z-50 flex items-center justify-center bg-white/30 backdrop-blur-none transition-opacity" style="backdrop-filter: blur(1px);">
    <div class="bg-white rounded-xl shadow-2xl max-w-lg w-full p-6 relative border border-gray-200">
      <button @click="showAddProductModal = false" class="absolute top-3 right-3 text-gray-400 hover:text-gray-700">
        <XIcon class="h-6 w-6" />
      </button>
      <h3 class="text-xl font-bold mb-4">Add Products to Collection</h3>
      <div class="mb-4">
        <input
          v-model="addProductSearchQuery"
          @input="debouncedAddProductSearch"
          type="text"
          placeholder="Search products by name..."
          class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500"
        />
      </div>
      <div v-if="addProductLoading" class="flex justify-center py-6">
        <SpinnerIcon class="animate-spin h-8 w-8 text-green-500" />
      </div>
      <div v-else-if="addProductError" class="text-red-600 text-center py-4">{{ addProductError }}</div>
      <div v-else>
        <div v-if="addProductList.length === 0" class="text-gray-500 text-center py-4">No products found.</div>
        <ul v-else class="divide-y divide-gray-100 max-h-64 overflow-y-auto">
          <li v-for="prod in addProductList" :key="prod.id" class="flex items-center justify-between py-3 px-1 hover:bg-gray-50 rounded-lg transition">
            <div class="flex items-center gap-3">
              <img v-if="prod.main_image" :src="prod.main_image" class="w-12 h-12 object-cover rounded-md border" />
              <PlaceholderIcon v-else class="w-10 h-10 text-gray-400" />
              <div>
                <div class="font-medium text-gray-900">{{ prod.name }}</div>
                <div class="text-xs text-gray-500 mt-0.5">
                  {{ isProductInCollection(prod, collection.id)
                    ? 'Already in this collection'
                    : (getProductCollections(prod).length > 0 ? getProductCollections(prod).join(', ') : 'Uncategorized')
                  }}
                </div>
                <div class="text-sm text-gray-500">
                  <template v-if="prod.starting_price != null">
                    <span class="italic text-gray-700">from</span> ${{ prod.starting_price.toFixed(2) }}
                  </template>
                  <template v-else-if="prod.price != null">
                    ${{ prod.price.toFixed(2) }}
                  </template>
                  <template v-else>
                    -
                  </template>
                </div>
              </div>
            </div>
            <button
              v-if="isProductInCollection(prod, collection.id)"
              class="inline-flex items-center px-3 py-1.5 bg-gray-200 text-gray-500 text-sm font-medium rounded-lg shadow-sm cursor-not-allowed border border-gray-300"
              disabled
            >
              <CheckIcon class="h-4 w-4 mr-1" /> Added
            </button>
            <button
              v-else
              @click="handleAddProduct(prod.id)"
              class="inline-flex items-center px-3 py-1.5 bg-green-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition"
              :disabled="addProductLoading"
            >
              <PlusIcon class="h-4 w-4 mr-1" /> Add
            </button>
          </li>
        </ul>
        <!-- Pagination Controls -->
        <div class="flex justify-center items-center space-x-2 mt-4">
          <button
            @click="addProductPrevPage"
            :disabled="addProductCurrentPage === 1"
            class="px-3 py-1 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300 disabled:opacity-50 disabled:cursor-not-allowed transition"
          >
            Previous
          </button>
          <span class="text-gray-700 text-sm font-medium">Page {{ addProductCurrentPage }} of {{ addProductPagination.total_pages }}</span>
          <button
            @click="addProductNextPage"
            :disabled="addProductCurrentPage === addProductPagination.total_pages"
            class="px-3 py-1 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300 disabled:opacity-50 disabled:cursor-not-allowed transition"
          >
            Next
          </button>
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
  CheckIcon,
  ExclamationCircleIcon
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

// Computed property for reversed products (newest first)
const reversedProducts = computed(() => {
  return [...(collection.value.products || [])].reverse()
})
const loading = ref(true)
const error = ref(null)

// Modal states
const showEditModal = ref(false)
const showAddProductModal = ref(false)

// Add Product modal state
const addProductSearchQuery = ref('')
const addProductLoading = ref(false)
const addProductList = ref([])
const addProductPagination = ref({ page: 1, limit: 10, total: 0, total_pages: 1 })
const addProductCurrentPage = ref(1)
const addProductSearchTimeout = ref(null)
const addProductError = ref(null)

// Edit form
const editForm = ref({
  title: '',
  handle: '',
  description: '',
  image: ''
})
const editErrors = ref({})
const editLoading = ref(false)

// Active shop from pinia
const activeShop = computed(() => shopStore.activeShop)

// For showing collection names in the modal
const allCollections = ref([])
const collectionMap = computed(() => {
  const map = {}
  for (const coll of allCollections.value) {
    map[coll.id] = coll.title
  }
  return map
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
    // allProducts.value = await productService.fetchAllByShop(activeShop.value.id) // This line is removed
  } catch (e) {
    console.error('Failed to load products:', e)
  }
}

/**
 * Loads products not in any collection, with optional search and pagination.
 */
async function loadAddableProducts() {
  if (!activeShop.value?.id) return
  addProductLoading.value = true
  addProductError.value = null
  try {
    const { products, pagination } = await productService.fetchAllByShopPaginated(
      activeShop.value.id,
      {
        page: addProductCurrentPage.value,
        limit: addProductPagination.value.limit,
        search: addProductSearchQuery.value,
        collection_id: null
      }
    )
    addProductList.value = products
    addProductPagination.value = pagination || { page: 1, limit: 10, total: products.length, total_pages: 1 }
  } catch (e) {
    addProductError.value = 'Failed to load products. Please try again.'
  } finally {
    addProductLoading.value = false
  }
}

function debouncedAddProductSearch() {
  if (addProductSearchTimeout.value) clearTimeout(addProductSearchTimeout.value)
  addProductSearchTimeout.value = setTimeout(() => {
    addProductCurrentPage.value = 1
    loadAddableProducts()
  }, 300)
}

function addProductPrevPage() {
  if (addProductCurrentPage.value > 1) {
    addProductCurrentPage.value--
    loadAddableProducts()
  }
}

function addProductNextPage() {
  if (addProductCurrentPage.value < addProductPagination.value.total_pages) {
    addProductCurrentPage.value++
    loadAddableProducts()
  }
}

async function handleAddProduct(productId) {
  if (!activeShop.value?.id || !collection.value.id) return
  addProductLoading.value = true
  try {
    await collectionService.addProduct(activeShop.value.id, collection.value.id, productId)
    await loadCollection(collection.value.id)
    await loadAddableProducts()
  } catch (e) {
    alert('Failed to add product to collection. Please try again.')
  } finally {
    addProductLoading.value = false
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
 * Navigates to a specific product's detail page.
 */
function goToProduct(productId) {
  router.push({ name: 'ProductDetail', params: { productId } })
}

/**
 * Navigates to the edit collection page.
 */
function goToEditPage() {
  router.push({ name: 'EditCollection', params: { collectionId: collection.value.id } })
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

// Load collections on mount and when modal opens
async function loadAllCollections() {
  if (!activeShop.value?.id) return
  allCollections.value = await collectionService.fetchAllByShop(activeShop.value.id)
}

/**
 * Check if a product is in a specific collection
 */
function isProductInCollection(product, collectionId) {
  return product.collection_ids && product.collection_ids.includes(collectionId)
}

/**
 * Get collection names for a product
 */
function getProductCollections(product) {
  if (!product.collection_ids || product.collection_ids.length === 0) return []
  return product.collection_ids.map(id => collectionMap.value[id]).filter(Boolean)
}

// Load collections on mount and when modal opens
onMounted(() => {
  loadAllCollections()
})

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

// Watch modal open to load products
watch(showAddProductModal, (val) => {
  if (val) {
    addProductSearchQuery.value = ''
    addProductCurrentPage.value = 1
    loadAddableProducts()
    loadAllCollections()
  }
})

// When searching, use backend search for fast results
// watch(productSearchQuery, async (query) => { // This line is removed
//   if (searchTimeout) clearTimeout(searchTimeout)
//   if (query) {
//     isSearching.value = true
//     // Debounce to avoid too many requests
//     searchTimeout = setTimeout(async () => {
//       allProducts.value = await productService.fetchAllByShop(activeShop.value.id, { search: query })
//     }, 250)
//   } else if (isSearching.value) {
//     // If search is cleared, reload only products with no collection
//     isSearching.value = false
//     allProducts.value = await productService.fetchAllByShopWithFilter(activeShop.value.id, { collection_id: null })
//   }
// })
</script>

<style scoped>
/* No specific scoped styles needed as Tailwind classes handle the design and animations */
</style>