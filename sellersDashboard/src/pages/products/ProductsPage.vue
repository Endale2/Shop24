<template>
  <div class="min-h-full bg-gray-50">
    <div class="p-4 sm:p-6 lg:p-8">
      
      <!-- Header Section -->
      <div class="mb-6 sm:mb-8">
        <div class="flex flex-col md:flex-row items-start md:items-center justify-between space-y-4 md:space-y-0">
          <div class="flex items-center mb-3 md:mb-0">
            <div class="w-10 h-10 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center mr-3 shadow-sm">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
              </svg>
            </div>
            <div>
              <h1 class="text-2xl sm:text-3xl font-bold text-gray-900 tracking-tight">
                Products
              </h1>
              <p class="text-sm text-gray-600 mt-1">Manage your product catalog</p>
            </div>
          </div>
          
          <div class="flex flex-col sm:flex-row gap-2 sm:gap-3 mt-4 sm:mt-0">
            <button
              @click="goToAddProduct"
              class="inline-flex items-center px-4 py-2.5 bg-green-600 text-white text-xs font-medium rounded-lg shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
            >
              <PlusIcon class="w-4 h-4 mr-1.5" />
              Add Product
            </button>
          </div>
        </div>
      </div>

      <!-- Search Bar -->
      <div class="mb-6">
        <div class="relative max-w-md">
          <input
            type="text"
            v-model="searchQuery"
            @input="debouncedSearch"
            placeholder="Search products by name or category..."
            class="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm text-sm"
          />
          <SearchIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-16 text-gray-500">
        <div class="flex flex-col items-center justify-center">
          <div class="animate-spin h-8 w-8 text-green-500 mb-3 border-3 border-green-200 border-t-green-500 rounded-full"></div>
          <p class="text-sm font-medium">Loading products...</p>
        </div>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="bg-red-50 border border-red-200 text-red-700 px-6 py-4 rounded-lg shadow-sm mb-6" role="alert">
        <div class="flex items-center">
          <ExclamationCircleIcon class="h-5 w-5 mr-2 flex-shrink-0" />
          <div>
            <strong class="font-semibold">Error:</strong>
            <span class="ml-1">{{ error }}</span>
          </div>
        </div>
        <p class="text-sm mt-2">Please ensure you have an active shop selected and try again.</p>
      </div>

      <!-- Content -->
      <div v-else>
        <!-- Stock Status Tabs -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-4 mb-6">
          <div class="flex flex-wrap gap-3">
            <button
              v-for="tab in stockTabs"
              :key="tab.value"
              type="button"
              @click.prevent="onStockTabClick(tab.value)"
              :class="[
                'px-3 py-2 rounded-lg text-xs font-medium transition-colors duration-200 flex items-center',
                selectedStockStatus === tab.value
                  ? 'bg-green-600 text-white'
                  : getStockTabColor(tab.value)
              ]"
            >
              {{ tab.label }}
              <span class="ml-2 bg-gray-200 text-gray-700 px-2 py-0.5 rounded-full text-xs" :class="selectedStockStatus === tab.value ? 'bg-white text-green-700' : ''">
                {{ getStockTabCount(tab.value) }}
              </span>
            </button>
          </div>
        </div>

        <div v-if="paginatedProducts.length">
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-sm font-semibold text-gray-900">Product List</h2>
            </div>
            <div class="overflow-x-auto">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="py-3 px-4 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Image</th>
                    <th class="py-3 px-4 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Name</th>
                    <th class="py-3 px-4 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Collection</th>
                    <th class="py-3 px-4 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Price</th>
                    <th class="py-3 px-4 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Stock</th>
                    <th class="py-3 px-4 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Status</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-gray-200">
                  <tr
                    v-for="(prod, i) in paginatedProducts"
                    :key="prod.id"
                    @click="goToDetail(prod.id)"
                    class="cursor-pointer transition duration-150 ease-in-out transform hover:scale-[1.005] hover:bg-green-50"
                    :class="{ 'bg-gray-50': i % 2 === 1 }"
                  >
                    <td class="py-3 px-4">
                      <div class="w-10 h-10 rounded-md overflow-hidden border border-gray-200 flex items-center justify-center bg-gray-100 shadow-sm">
                        <img
                          v-if="getProductImage(prod)"
                          :src="getProductImage(prod)"
                          alt="thumbnail"
                          class="w-full h-full object-cover"
                        />
                        <PlaceholderIcon v-else class="w-6 h-6 text-gray-400" />
                      </div>
                    </td>
                    <td class="py-3 px-4">
                      <div class="text-sm text-gray-800 font-medium">{{ prod.name }}</div>
                      <div v-if="prod.variants && prod.variants.length > 0" class="text-xs text-gray-500 mt-1">
                        {{ prod.variants.length }} variant{{ prod.variants.length !== 1 ? 's' : '' }}
                      </div>
                    </td>
                    <td class="py-3 px-4 text-sm text-gray-700">
                      <div v-if="prod.collection_ids && prod.collection_ids.length > 0" class="space-y-1">
                        <div v-for="collectionId in prod.collection_ids" :key="collectionId" class="text-xs bg-gray-100 px-2 py-1 rounded">
                          {{ collectionMap[collectionId] || 'Unknown Collection' }}
                        </div>
                      </div>
                      <div v-else class="text-gray-500 italic">Uncategorized</div>
                    </td>
                    <td class="py-3 px-4 text-sm text-green-600 font-semibold">
                      <template v-if="prod.starting_price != null">
                        <span class="italic text-gray-700">from</span> {{ formatPrice(prod.starting_price) }}
                      </template>
                      <template v-else-if="prod.price != null">
                        {{ formatPrice(prod.price) }}
                      </template>
                      <template v-else>
                        N/A
                      </template>
                    </td>
                    <td class="py-3 px-4 text-sm text-gray-700">
                      <template v-if="prod.total_stock != null">
                        {{ prod.total_stock }}
                      </template>
                      <template v-else-if="prod.stock != null">
                        {{ prod.stock }}
                      </template>
                      <template v-else>
                        N/A
                      </template>
                    </td>
                    <td class="py-3 px-4">
                      <span
                        :class="{
                          'px-2 py-0.5 text-xs font-medium rounded-full whitespace-nowrap w-max': true,
                          'bg-green-100 text-green-800': getProductStock(prod) > 0,
                          'bg-red-100 text-red-800': getProductStock(prod) <= 0
                        }"
                      >
                        {{ getProductStock(prod) > 0 ? 'In Stock' : 'Out of Stock' }}
                      </span>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>

          <!-- Pagination -->
          <div class="flex justify-center items-center space-x-2 mt-8">
            <button
              @click="prevPage"
              :disabled="currentPage === 1"
              class="px-4 py-2 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300 disabled:opacity-50 disabled:cursor-not-allowed transition duration-150 ease-in-out text-sm"
            >
              Previous
            </button>
            <span class="text-gray-700 text-sm font-medium">Page {{ currentPage }} of {{ totalPages }}</span>
            <button
              @click="nextPage"
              :disabled="currentPage === totalPages"
              class="px-4 py-2 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300 disabled:opacity-50 disabled:cursor-not-allowed transition duration-150 ease-in-out text-sm"
            >
              Next
            </button>
          </div>
        </div>

        <!-- Empty State -->
        <div v-else class="text-center py-16 text-gray-400">
          <div class="max-w-md mx-auto">
            <div class="w-12 h-12 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-3">
              <PlusIcon class="w-6 h-6 text-gray-300" />
            </div>
            <p class="text-sm font-medium text-gray-600 mb-1">No products yet</p>
            <p class="text-xs">Add products to see them here</p>
            <button
              @click="goToAddProduct"
              class="mt-4 inline-flex items-center px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
            >
              <PlusIcon class="w-4 h-4 mr-2" />
              Add Your First Product
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { useShopState } from '@/composables/useShopState'
import { productService } from '@/services/product'
import { collectionService } from '@/services/collection'
import {
  PhotographIcon as PlaceholderIcon,
  SearchIcon,
  PlusIcon,
  ExclamationCircleIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const shopStore = useShopStore()
const {
  activeShop,
  isShopStateReady,
  ensureValidShopState,
  waitForShopState,
  isInitializing,
  initializationError
} = useShopState()

// Reactive state
const allProducts = ref([]) // not used anymore, but kept for compatibility
const products = ref([])
const loading = ref(false)
const error = ref(null)

// Search/filter state
const searchQuery = ref('')
const selectedStockStatus = ref('')
const selectedCategory = ref('') // for future use
let searchTimeout = null

// Pagination state
const currentPage = ref(1)
const itemsPerPage = 10
const pagination = ref({ page: 1, limit: 10, total: 0, total_pages: 1 })
const stats = ref({ total_products: 0, in_stock: 0, out_of_stock: 0 })

// Shop state is now managed by the composable

// Computed properties for pagination
const totalPages = computed(() => pagination.value.total_pages || 1)
const paginatedProducts = computed(() => products.value)



/**
 * Gets the appropriate image for a product
 */
function getProductImage(product) {
  if (product.main_image) return product.main_image
  if (product.images && product.images.length > 0) return product.images[0]
  return null
}

/**
 * Gets the total stock for a product
 */
function getProductStock(product) {
  if (product.total_stock !== undefined) return product.total_stock
  if (product.stock !== undefined) return product.stock
  if (product.variants && product.variants.length > 0) {
    return product.variants.reduce((sum, v) => sum + (v.stock || 0), 0)
  }
  return 0
}

/**
 * Fetches products from the service with enhanced shop state validation.
 */
async function fetchProducts() {
  // Ensure we have a valid active shop
  const shop = await ensureValidShopState()
  if (!shop) {
    return
  }

  loading.value = true
  error.value = null

  try {
    console.log('[ProductsPage] Fetching products for shop:', shop.name)
    const { products: fetched, pagination: pag, stats: st } = await productService.fetchAllByShopPaginated(
      shop.id,
      {
        page: currentPage.value,
        limit: itemsPerPage,
        search: searchQuery.value,
        category: selectedCategory.value,
        stockStatus: selectedStockStatus.value,
      }
    )
    products.value = fetched
    pagination.value = pag || { page: 1, limit: itemsPerPage, total: fetched.length, total_pages: 1 }
    stats.value = st || { total_products: fetched.length, in_stock: 0, out_of_stock: 0 }

    console.log('[ProductsPage] Products loaded successfully:', {
      count: fetched.length,
      page: currentPage.value,
      total: pag?.total || 0
    })
  } catch (e) {
    console.error('[ProductsPage] Failed to load products:', e)
    error.value = 'Failed to load products. Please try again.'

    // If error is related to shop not found, redirect to shop selection
    if (e.response?.status === 404 || e.message?.includes('shop')) {
      console.warn('[ProductsPage] Shop-related error, redirecting to shop selection')
      router.replace({ name: 'ShopSelection' })
    }
  } finally {
    loading.value = false
  }
}

// ensureValidShopState is now provided by the useShopState composable

/**
 * Applies the search filter to the products.
 */
function debouncedSearch() {
  clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => {
    currentPage.value = 1
    fetchProducts()
  }, 300)
}

/**
 * Handles going to the previous page.
 */
function prevPage() {
  if (currentPage.value > 1) {
    currentPage.value--
    fetchProducts()
  }
}

/**
 * Handles going to the next page.
 */
function nextPage() {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
    fetchProducts()
  }
}

/**
 * Navigates to the product detail page.
 */
function goToDetail(productId) {
  router.push({ name: 'ProductDetail', params: { productId } })
}

/**
 * Navigates to the page for adding a new product.
 */
function goToAddProduct() {
  router.push({ name: 'AddProduct' })
}

/**
 * Formats a date string.
 */
function formatDate(dt) {
  return dt
    ? new Date(dt).toLocaleDateString(undefined, {
        year: 'numeric',
        month: 'short',
        day: 'numeric'
      })
    : '—'
}

/**
 * Formats a price number.
 */
function formatPrice(p) {
  return p != null ? `$${p.toFixed(2)}` : 'N/A'
}

const stockTabs = [
  { label: 'All', value: '' },
  { label: 'In Stock', value: 'in_stock' },
  { label: 'Out of Stock', value: 'out_of_stock' },
]

const tabCounts = ref({ all: 0, in_stock: 0, out_of_stock: 0 })

async function fetchTabCounts() {
  const shop = await ensureValidShopState()
  if (!shop) return

  try {
    // Fetch counts for all tabs in parallel
    const [allRes, inStockRes, outStockRes] = await Promise.all([
      productService.fetchAllByShopPaginated(shop.id, { page: 1, limit: 1, stockStatus: '' }),
      productService.fetchAllByShopPaginated(shop.id, { page: 1, limit: 1, stockStatus: 'in_stock' }),
      productService.fetchAllByShopPaginated(shop.id, { page: 1, limit: 1, stockStatus: 'out_of_stock' })
    ])

    tabCounts.value.all = allRes.pagination ? allRes.pagination.total : (allRes.products?.length || 0)
    tabCounts.value.in_stock = inStockRes.pagination ? inStockRes.pagination.total : (inStockRes.products?.length || 0)
    tabCounts.value.out_of_stock = outStockRes.pagination ? outStockRes.pagination.total : (outStockRes.products?.length || 0)
  } catch (error) {
    console.error('[ProductsPage] Failed to fetch tab counts:', error)
    // Set default values on error
    tabCounts.value = { all: 0, in_stock: 0, out_of_stock: 0 }
  }
}

function getStockTabCount(val) {
  switch (val) {
    case 'in_stock': return tabCounts.value.in_stock
    case 'out_of_stock': return tabCounts.value.out_of_stock
    case '':
    default: return tabCounts.value.all
  }
}

async function onStockTabClick(val) {
  if (selectedStockStatus.value !== val) {
    selectedStockStatus.value = val
    currentPage.value = 1
    await fetchProducts()
    await fetchTabCounts()
  }
}

function getStockTabColor(val) {
  switch (val) {
    case 'in_stock': return 'bg-green-100 text-green-800 hover:bg-green-200';
    case 'out_of_stock': return 'bg-red-100 text-red-800 hover:bg-red-200';
    case '':
    default: return 'bg-gray-100 text-gray-700 hover:bg-gray-200';
  }
}

const collections = ref([])
const collectionMap = computed(() => {
  const map = {}
  for (const coll of collections.value) {
    map[coll.id] = coll.title
  }
  return map
})

async function fetchCollections() {
  const shop = await ensureValidShopState()
  if (!shop) return

  try {
    collections.value = await collectionService.fetchAllByShop(shop.id)
  } catch (error) {
    console.error('[ProductsPage] Failed to fetch collections:', error)
    // Don't fail the entire page if collections fail to load
    collections.value = []
  }
}

// Initial data fetch on component mount (OrdersPage pattern)
onMounted(async () => {
  console.log('[ProductsPage] Component mounted, initializing...')
  try {
    // Ensure we have an active shop; if not, try to set one or redirect
    if (!activeShop.value?.id) {
      console.log('[ProductsPage] No active shop found, ensuring...')
      const ensured = await shopStore.ensureActiveShop()
      if (!ensured?.id) {
        console.warn('[ProductsPage] No shop available, redirecting to ShopSelection')
        router.replace({ name: 'ShopSelection' })
        return
      }
    }

    console.log('[ProductsPage] Shop state ready, loading data...')
    await Promise.all([
      fetchProducts(),
      fetchTabCounts(),
      fetchCollections()
    ])
  } catch (error) {
    console.error('[ProductsPage] Failed to initialize:', error)
  }
})

// Watch for changes in activeShop to refetch products if the shop changes
watch(activeShop, async (newShop, oldShop) => {
  if (newShop?.id !== oldShop?.id) {
    console.log('[ProductsPage] Active shop changed:', {
      from: oldShop?.name,
      to: newShop?.name
    })

    if (newShop?.id) {
      currentPage.value = 1
      await Promise.all([
        fetchProducts(),
        fetchTabCounts(),
        fetchCollections()
      ])
    }
  }
})

// Watch for shop state readiness
watch(isShopStateReady, async (isReady) => {
  if (isReady && activeShop.value && !products.value.length) {
    console.log('[ProductsPage] Shop state became ready, loading data...')
    await Promise.all([
      fetchProducts(),
      fetchTabCounts(),
      fetchCollections()
    ])
  }
})
</script>

<style scoped>
/* Custom styles for enhanced design */
@media (max-width: 640px) {
  .min-w-0 { min-width: 0 !important; }
}

/* Enhanced hover effects */
.group:hover .group-hover\:scale-105 {
  transform: scale(1.05);
}

/* Smooth transitions */
* {
  transition-property: color, background-color, border-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
}

/* Custom scrollbar for better UX */
::-webkit-scrollbar {
  width: 4px;
}

::-webkit-scrollbar-track {
  background: transparent;
}

::-webkit-scrollbar-thumb {
  background: #d1d5db;
  border-radius: 2px;
}

::-webkit-scrollbar-thumb:hover {
  background: #9ca3af;
}
</style>