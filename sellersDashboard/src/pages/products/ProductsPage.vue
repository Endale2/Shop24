<template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto font-sans">
    <h2 class="text-3xl sm:text-4xl font-extrabold mb-8 text-gray-900 leading-tight">
      Products
    </h2>

    <div class="flex flex-col sm:flex-row justify-between items-center mb-6 space-y-4 sm:space-y-0">
      <div class="w-full sm:w-1/2 relative">
        <input
          type="text"
          v-model="searchQuery"
          @input="debouncedSearch"
          placeholder="Search products..."
          class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm"
        />
        <SearchIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
      </div>

      <div class="flex space-x-4 items-center">
        <button
          @click="goToAddProduct"
          class="inline-flex items-center px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg shadow-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
        >
          <PlusIcon class="w-5 h-5 mr-2 -ml-1" />
          Add New Product
        </button>

        <div class="inline-flex rounded-full shadow-md overflow-hidden" role="group">
          <button
            @click="currentView = 'list'"
            :class="viewButtonClass('list')"
            class="px-5 py-2.5 text-sm font-medium border border-gray-300 focus:z-10 focus:ring-2 focus:ring-green-500 focus:outline-none transition-colors duration-200 ease-in-out"
          >
            <ListIcon class="w-5 h-5 inline-block mr-1" />
            List View
          </button>
          <button
            @click="currentView = 'grid'"
            :class="viewButtonClass('grid')"
            class="px-5 py-2.5 text-sm font-medium border border-gray-300 focus:z-10 focus:ring-2 focus:ring-green-500 focus:outline-none transition-colors duration-200 ease-in-out"
          >
            <GridIcon class="w-5 h-5 inline-block mr-1" />
            Grid View
          </button>
        </div>
      </div>
    </div>

    <div v-if="loading">
      <div v-if="currentView === 'list'" class="overflow-x-auto bg-white shadow-lg rounded-xl animate-pulse">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-100">
            <tr>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Image</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Name</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Category</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Price</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Created</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="n in itemsPerPage" :key="n" class="bg-white">
              <td class="py-3 px-6"><div class="w-12 h-12 rounded-lg bg-gray-200"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-3/4"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-1/2"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-1/3"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-1/4"></div></td>
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

    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-lg shadow-sm mb-6" role="alert">
      <strong class="font-bold">Error:</strong>
      <span class="block sm:inline ml-2">{{ error }}</span>
      <p class="text-sm mt-2">Please ensure you have an active shop selected and try again.</p>
    </div>

    <div v-else>
      <div v-if="paginatedProducts.length">
        <div v-if="currentView === 'list'" class="overflow-x-auto bg-white shadow-lg rounded-xl">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-100">
              <tr>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Image</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Name</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Category</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Price</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Created</th>
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
                      v-if="prod.main_image"
                      :src="prod.main_image"
                      alt="thumbnail"
                      class="w-full h-full object-cover"
                    />
                    <img
                      v-else-if="prod.images?.length"
                      :src="prod.images[0]"
                      alt="thumbnail"
                      class="w-full h-full object-cover"
                    />
                    <PlaceholderIcon v-else class="w-8 h-8 text-gray-400" />
                  </div>
                </td>
                <td class="py-3 px-6 text-sm text-gray-800 font-medium">{{ prod.name }}</td>
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
                <td class="py-3 px-6 text-sm text-gray-600">{{ formatDate(prod.createdAt) }}</td>
              </tr>
            </tbody>
          </table>
        </div>

        <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
          <div
            v-for="prod in paginatedProducts"
            :key="prod.id"
            @click="goToDetail(prod.id)"
            class="bg-white rounded-xl shadow-lg overflow-hidden cursor-pointer transform hover:scale-105 hover:shadow-xl transition duration-200 ease-in-out flex flex-col group"
          >
            <div class="w-full h-48 bg-gray-200 flex items-center justify-center text-gray-400 relative overflow-hidden rounded-t-xl">
              <img
                v-if="prod.main_image"
                :src="prod.main_image"
                alt="product"
                class="w-full h-full object-cover transform group-hover:scale-110 transition-transform duration-300 ease-in-out"
              />
              <img
                v-else-if="prod.images?.length"
                :src="prod.images[0]"
                alt="product"
                class="w-full h-full object-cover transform group-hover:scale-110 transition-transform duration-300 ease-in-out"
              />
              <PlaceholderIcon v-else class="w-16 h-16" />
            </div>
            <div class="p-5 flex-grow flex flex-col">
              <h3 class="text-xl font-semibold text-gray-900 mb-2 truncate">{{ prod.name }}</h3>
              <p class="text-sm text-gray-600 mb-3">{{ prod.category || 'Uncategorized' }}</p>
              <div class="mt-auto">
                <p class="text-2xl font-bold text-green-700">
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
              </div>
            </div>
          </div>
        </div>

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

      <div v-else class="bg-green-50 border border-green-200 text-green-700 px-6 py-8 rounded-lg text-center mt-8 shadow-sm">
        <p class="text-lg font-medium">No products found for this search.</p>
        <p class="mt-2">Try adjusting your search query or adding a new product.</p>
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
  SearchIcon, // Added for search bar
  PlusIcon // Added for add new product button
} from '@heroicons/vue/outline'

const router = useRouter()
const shopStore = useShopStore()

// Reactive state
const allProducts = ref([]) // Holds all fetched products
const products = ref([]) // Products filtered by search, before pagination
const loading = ref(false)
const error = ref(null)
const currentView = ref('list')

// Search state
const searchQuery = ref('')
let searchTimeout = null; // To debounce search input

// Pagination state
const currentPage = ref(1)
const itemsPerPage = 8 // You can adjust this number

// Computed property for the active shop
const activeShop = computed(() => shopStore.activeShop)

// Computed properties for pagination
const totalPages = computed(() => Math.ceil(products.value.length / itemsPerPage))
const paginatedProducts = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage
  const end = start + itemsPerPage
  return products.value.slice(start, end)
})

/**
 * Dynamically applies classes to view toggle buttons based on the current view.
 * @param {string} view - The view type ('list' or 'grid').
 * @returns {string} Tailwind CSS classes.
 */
function viewButtonClass(view) {
  return view === currentView.value
    ? 'bg-green-600 text-white hover:bg-green-700 shadow-inner'
    : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
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
    // Fetch products using the shop ID
    const fetchedData = await productService.fetchAllByShop(activeShop.value.id)
    // Ensure mapping for new/updated backend fields
    allProducts.value = fetchedData.map(prod => ({
      ...prod, // Keep all original product properties
      id: prod.id || prod._id, // Standardize ID
      images: prod.images || [], // Ensure images is always an array
    }))
    applySearchFilter() // Apply initial search filter (empty string)
    currentPage.value = 1 // Reset pagination on new fetch
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
function applySearchFilter() {
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    products.value = allProducts.value.filter(prod =>
      prod.name.toLowerCase().includes(query) ||
      prod.category?.toLowerCase().includes(query)
    )
  } else {
    products.value = [...allProducts.value] // Clone to ensure reactivity
  }
  currentPage.value = 1; // Reset to first page after search
}

/**
 * Debounces the search input to avoid excessive API calls or filtering.
 */
function debouncedSearch() {
  clearTimeout(searchTimeout);
  searchTimeout = setTimeout(() => {
    applySearchFilter();
  }, 300); // 300ms debounce time
}

/**
 * Handles going to the previous page.
 */
function prevPage() {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

/**
 * Handles going to the next page.
 */
function nextPage() {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

/**
 * Navigates to the product detail page.
 * @param {string} productId - The ID of the product.
 */
function goToDetail(productId) {
  router.push({ name: 'ProductDetail', params: { productId } })
}

/**
 * Navigates to the page for adding a new product.
 */
function goToAddProduct() {
  router.push({ name: 'AddProduct' }) // Assuming you have a route named 'AddProduct'
}

/**
 * Formats a date string.
 * @param {string} dt - The date string.
 * @returns {string} Formatted date.
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
 * @param {number} p - The price.
 * @returns {string} Formatted price.
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
    fetchProducts();
  }
});
</script>