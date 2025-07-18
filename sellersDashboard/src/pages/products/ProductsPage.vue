<template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto font-sans">
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center mb-8">
      <div>
        <h2 class="text-3xl sm:text-4xl font-extrabold text-gray-900 leading-tight">
          Products
        </h2>
        <p class="text-gray-600 mt-2">Manage your product catalog</p>
      </div>
      
      <div class="flex flex-col sm:flex-row gap-3 mt-4 sm:mt-0">
        <button
          @click="goToAddProduct"
          class="inline-flex items-center px-4 py-2.5 bg-green-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
        >
          <PlusIcon class="w-5 h-5 mr-2 -ml-1" />
          Add Product
        </button>

        <div class="inline-flex rounded-lg shadow-sm overflow-hidden" role="group">
          <button
            @click="currentView = 'list'"
            :class="viewButtonClass('list')"
            class="px-4 py-2.5 text-sm font-medium border border-gray-300 focus:z-10 focus:ring-2 focus:ring-green-500 focus:outline-none transition-colors duration-200 ease-in-out"
          >
            <ListIcon class="w-5 h-5 inline-block mr-1" />
            List
          </button>
          <button
            @click="currentView = 'grid'"
            :class="viewButtonClass('grid')"
            class="px-4 py-2.5 text-sm font-medium border border-gray-300 focus:z-10 focus:ring-2 focus:ring-green-500 focus:outline-none transition-colors duration-200 ease-in-out"
          >
            <GridIcon class="w-5 h-5 inline-block mr-1" />
            Grid
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
          class="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm"
        />
        <SearchIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="space-y-4">
      <div v-if="currentView === 'list'" class="overflow-x-auto bg-white shadow-lg rounded-xl animate-pulse">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Image</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Name</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Category</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Price</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Stock</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Status</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="n in itemsPerPage" :key="n" class="bg-white">
              <td class="py-3 px-6"><div class="w-12 h-12 rounded-lg bg-gray-200"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-3/4"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-1/2"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-1/3"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-1/4"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-1/3"></div></td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6 animate-pulse">
        <div v-for="n in itemsPerPage" :key="n" class="bg-white rounded-xl shadow-lg overflow-hidden flex flex-col">
          <div class="w-full h-48 bg-gray-200"></div>
          <div class="p-5 flex-grow flex flex-col space-y-3">
            <div class="h-5 bg-gray-200 rounded w-3/4"></div>
            <div class="h-4 bg-gray-200 rounded w-1/2"></div>
            <div class="h-6 bg-gray-200 rounded w-1/3 mt-auto"></div>
          </div>
        </div>
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
      <div v-if="paginatedProducts.length">
        <!-- Add this above the product table/grid -->
        <div class="flex flex-wrap gap-4 mb-4 items-center">
          <div class="flex gap-2 items-center">
            <span class="text-gray-700 font-medium">Total:</span>
            <span class="text-gray-900">{{ stats.total_products }}</span>
            <span class="text-green-700 ml-4">In Stock: {{ stats.in_stock }}</span>
            <span class="text-red-700 ml-4">Out of Stock: {{ stats.out_of_stock }}</span>
          </div>
          <div class="ml-auto flex gap-2 items-center">
            <label for="stockStatus" class="text-sm text-gray-700">Stock Status:</label>
            <select id="stockStatus" v-model="selectedStockStatus" @change="handleStockStatusChange" class="border border-gray-300 rounded px-2 py-1 text-sm">
              <option value="">All</option>
              <option value="in_stock">In Stock</option>
              <option value="out_of_stock">Out of Stock</option>
            </select>
          </div>
        </div>

        <!-- List View -->
        <div v-if="currentView === 'list'" class="overflow-x-auto bg-white shadow-lg rounded-xl">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Image</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Name</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Category</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Price</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Stock</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Status</th>
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
                <td class="py-3 px-6">
                  <div class="w-12 h-12 rounded-lg overflow-hidden border border-gray-200 flex items-center justify-center bg-gray-100 shadow-sm">
                    <img
                      v-if="getProductImage(prod)"
                      :src="getProductImage(prod)"
                      alt="thumbnail"
                      class="w-full h-full object-cover"
                    />
                    <PlaceholderIcon v-else class="w-8 h-8 text-gray-400" />
                  </div>
                </td>
                <td class="py-3 px-6">
                  <div class="text-sm text-gray-800 font-medium">{{ prod.name }}</div>
                  <div v-if="prod.variants && prod.variants.length > 0" class="text-xs text-gray-500 mt-1">
                    {{ prod.variants.length }} variant{{ prod.variants.length !== 1 ? 's' : '' }}
                  </div>
                </td>
                <td class="py-3 px-6 text-sm text-gray-700">{{ prod.category || 'Uncategorized' }}</td>
                <td class="py-3 px-6 text-sm text-green-600 font-semibold">
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
                <td class="py-3 px-6 text-sm text-gray-700">
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
                <td class="py-3 px-6">
                  <span
                    :class="{
                      'px-2 py-1 text-xs font-medium rounded-full whitespace-nowrap w-max': true,
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

        <!-- Grid View -->
        <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
          <div
            v-for="prod in paginatedProducts"
            :key="prod.id"
            @click="goToDetail(prod.id)"
            class="bg-white rounded-xl shadow-lg overflow-hidden cursor-pointer transform hover:scale-105 hover:shadow-xl transition duration-200 ease-in-out flex flex-col group"
          >
            <div class="w-full h-48 bg-gray-200 flex items-center justify-center text-gray-400 relative overflow-hidden rounded-t-xl">
              <img
                v-if="getProductImage(prod)"
                :src="getProductImage(prod)"
                alt="product"
                class="w-full h-full object-cover transform group-hover:scale-110 transition-transform duration-300 ease-in-out"
              />
              <PlaceholderIcon v-else class="w-16 h-16" />
              
              <!-- Stock Status Badge -->
              <div class="absolute top-2 right-2">
                <span
                  :class="{
                    'px-2 py-1 text-xs font-medium rounded-full whitespace-nowrap w-max': true,
                    'bg-green-100 text-green-800': getProductStock(prod) > 0,
                    'bg-red-100 text-red-800': getProductStock(prod) <= 0
                  }"
                >
                  {{ getProductStock(prod) > 0 ? 'In Stock' : 'Out of Stock' }}
                </span>
              </div>
            </div>
            
            <div class="p-5 flex-grow flex flex-col">
              <h3 class="text-lg font-semibold text-gray-900 mb-2 truncate">{{ prod.name }}</h3>
              <p class="text-sm text-gray-600 mb-2">{{ prod.category || 'Uncategorized' }}</p>
              
              <div v-if="prod.variants && prod.variants.length > 0" class="text-xs text-gray-500 mb-3">
                {{ prod.variants.length }} variant{{ prod.variants.length !== 1 ? 's' : '' }}
              </div>
              
              <div class="mt-auto">
                <p class="text-xl font-bold text-green-700">
                  <template v-if="prod.starting_price != null">
                    <span class="italic text-gray-700">from</span> {{ formatPrice(prod.starting_price) }}
                  </template>
                  <template v-else-if="prod.price != null">
                    {{ formatPrice(prod.price) }}
                  </template>
                  <template v-else>
                    N/A
                  </template>
                </p>
                <p class="text-sm text-gray-500 mt-1">
                  Stock: {{ getProductStock(prod) }}
                </p>
              </div>
            </div>
          </div>
        </div>

        <!-- Pagination -->
        <div class="flex justify-center items-center space-x-2 mt-8">
          <button
            @click="prevPage"
            :disabled="currentPage === 1"
            class="px-4 py-2 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300 disabled:opacity-50 disabled:cursor-not-allowed transition duration-150 ease-in-out"
          >
            Previous
          </button>
          <span class="text-gray-700 text-lg font-medium">Page {{ currentPage }} of {{ totalPages }}</span>
          <button
            @click="nextPage"
            :disabled="currentPage === totalPages"
            class="px-4 py-2 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300 disabled:opacity-50 disabled:cursor-not-allowed transition duration-150 ease-in-out"
          >
            Next
          </button>
        </div>
      </div>

      <!-- Empty State -->
      <div v-else class="bg-green-50 border border-green-200 text-green-700 px-6 py-8 rounded-lg text-center mt-8 shadow-sm">
        <div class="flex flex-col items-center">
          <div class="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mb-4">
            <PlusIcon class="w-8 h-8 text-green-600" />
          </div>
          <p class="text-lg font-medium">No products found</p>
          <p class="mt-2 text-sm">Try adjusting your search query or add your first product.</p>
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
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { productService } from '@/services/product'
import {
  ViewListIcon as ListIcon,
  ViewGridIcon as GridIcon,
  RefreshIcon as SpinnerIcon,
  PhotographIcon as PlaceholderIcon,
  SearchIcon,
  PlusIcon,
  ExclamationCircleIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const shopStore = useShopStore()

// Reactive state
const allProducts = ref([]) // not used anymore, but kept for compatibility
const products = ref([])
const loading = ref(false)
const error = ref(null)
const currentView = ref('list')

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

// Computed property for the active shop
const activeShop = computed(() => shopStore.activeShop)

// Computed properties for pagination
const totalPages = computed(() => pagination.value.total_pages || 1)
const paginatedProducts = computed(() => products.value)

/**
 * Dynamically applies classes to view toggle buttons based on the current view.
 */
function viewButtonClass(view) {
  return view === currentView.value
    ? 'bg-green-600 text-white hover:bg-green-700 shadow-inner'
    : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
}

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
 * Fetches products from the service.
 */
async function fetchProducts() {
  if (!activeShop.value?.id) {
    router.replace({ name: 'ShopSelection' })
    return
  }
  loading.value = true
  error.value = null
  try {
    const { products: fetched, pagination: pag, stats: st } = await productService.fetchAllByShopPaginated(
      activeShop.value.id,
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
  } catch (e) {
    console.error('Failed to load products:', e)
    error.value = 'Failed to load products. Please try again.'
  } finally {
    loading.value = false
  }
}

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

function handleStockStatusChange(e) {
  selectedStockStatus.value = e.target.value
  currentPage.value = 1
  fetchProducts()
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

// Initial data fetch on component mount
onMounted(() => {
  fetchProducts()
})

// Watch for changes in activeShop to refetch products if the shop changes
watch(activeShop, (newShop, oldShop) => {
  if (newShop?.id !== oldShop?.id) {
    currentPage.value = 1
    fetchProducts()
  }
})
</script>