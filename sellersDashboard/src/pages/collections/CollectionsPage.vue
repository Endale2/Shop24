<template>
  <div class="min-h-full bg-gray-50">
    <div class="p-4 sm:p-6 lg:p-8">
      
      <!-- Header Section -->
      <div class="mb-6 sm:mb-8">
        <div class="flex flex-col md:flex-row items-start md:items-center justify-between space-y-4 md:space-y-0">
          <div class="flex items-center mb-3 md:mb-0">
            <div class="w-10 h-10 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center mr-3 shadow-sm">
              <CollectionIcon class="w-5 h-5 text-white" />
            </div>
            <div>
              <h1 class="text-2xl sm:text-3xl font-bold text-gray-900 tracking-tight">
                Collections
              </h1>
              <p class="text-sm text-gray-600 mt-1">Manage your collection catalog</p>
            </div>
          </div>
          <div class="flex flex-col sm:flex-row gap-3">
            <button
              @click="goToAddCollection"
              class="inline-flex items-center px-4 py-2.5 bg-green-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
            >
              <PlusIcon class="w-4 h-4 mr-1.5" />
              Add Collection
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
            placeholder="Search collections by name or handle..."
            class="w-full pl-10 pr-4 py-2.5 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm text-sm"
          />
          <SearchIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-16 text-gray-500">
        <div class="flex flex-col items-center justify-center">
          <div class="animate-spin h-8 w-8 text-green-500 mb-3 border-3 border-green-200 border-t-green-500 rounded-full"></div>
          <p class="text-sm font-medium">Loading collections...</p>
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
        <div v-if="Array.isArray(paginatedCollections) && paginatedCollections.length">
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="overflow-x-auto">
              <table class="min-w-full divide-y divide-gray-200">
                                 <thead class="bg-gray-50">
                   <tr>
                     <th class="px-4 py-3 text-left text-xs font-semibold text-gray-900 uppercase tracking-wider">Image</th>
                     <th class="px-4 py-3 text-left text-xs font-semibold text-gray-900 uppercase tracking-wider">Title</th>
                     <th class="px-4 py-3 text-left text-xs font-semibold text-gray-900 uppercase tracking-wider">Handle</th>
                     <th class="px-4 py-3 text-left text-xs font-semibold text-gray-900 uppercase tracking-wider">Products</th>
                     <th class="px-4 py-3 text-left text-xs font-semibold text-gray-900 uppercase tracking-wider">Created</th>
                   </tr>
                 </thead>
                <tbody class="divide-y divide-gray-200">
                  <tr
                    v-for="(col, i) in paginatedCollections"
                    :key="col.id"
                    @click="goToDetail(col.id)"
                    class="cursor-pointer transition duration-150 ease-in-out transform hover:scale-[1.005] hover:bg-green-50"
                    :class="{ 'bg-gray-50': i % 2 === 1 }"
                  >
                    <td class="py-3 px-4">
                      <div class="w-12 h-12 rounded-lg overflow-hidden border border-gray-200 flex items-center justify-center bg-gray-100 shadow-sm">
                        <img
                          v-if="col.image"
                          :src="col.image"
                          alt="collection thumbnail"
                          class="w-full h-full object-cover"
                        />
                        <PlaceholderIcon v-else class="w-6 h-6 text-gray-400" />
                      </div>
                    </td>
                    <td class="py-3 px-4">
                      <div class="text-sm text-gray-800 font-medium">{{ col.title }}</div>
                    </td>
                    <td class="py-3 px-4 text-sm text-gray-700 font-mono">{{ col.handle }}</td>
                    <td class="py-3 px-4 text-sm text-gray-700">{{ Array.isArray(col.products) ? col.products.length : 0 }}</td>
                                         <td class="py-3 px-4 text-sm text-gray-700">{{ col.created_at ? formatDate(col.created_at) : (col.createdAt ? formatDate(col.createdAt) : 'N/A') }}</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>

          <!-- Pagination -->
          <div class="flex justify-center items-center space-x-2 mt-6">
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
            <p class="text-sm font-medium text-gray-600 mb-1">
              <template v-if="!Array.isArray(paginatedCollections)">
                Unable to load collections. Please try again later.
              </template>
              <template v-else-if="paginatedCollections.length === 0">
                No collections found
              </template>
            </p>
            <p class="text-xs mb-4">Get started by creating your first collection.</p>
            <button
              @click="goToAddCollection"
              class="inline-flex items-center px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
            >
              <PlusIcon class="w-4 h-4 mr-1.5" />
              Add Collection
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
import { collectionService } from '@/services/collection'
import {
  PhotographIcon as PlaceholderIcon,
  SearchIcon,
  PlusIcon,
  CollectionIcon,
  ExclamationCircleIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const shopStore = useShopStore()

// Reactive state
const allCollections = ref([]) // Holds all fetched collections
const collections = ref([]) // Collections filtered by search, before pagination
const loading     = ref(false)
const error       = ref(null)

// Search state
const searchQuery = ref('')
let searchTimeout = null; // To debounce search input

// Pagination state
const currentPage = ref(1)
const itemsPerPage = 8 // Number of collections to show per page

// Computed
const activeShop = computed(() => shopStore.activeShop)

// Computed properties for pagination
const totalPages = computed(() => Math.ceil(collections.value.length / itemsPerPage))
const paginatedCollections = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage
  const end = start + itemsPerPage
  return collections.value.slice(start, end)
})



/**
 * Fetches collections from the service.
 */
async function fetchCollections() {
  if (!activeShop.value?.id) {
    router.replace({ name: 'ShopSelection' })
    return
  }
  loading.value = true
  error.value = null
  try {
    allCollections.value = await collectionService.fetchAllByShop(activeShop.value.id)
    applySearchFilter() // Apply initial search filter (empty string)
    currentPage.value = 1 // Reset pagination on new fetch
  } catch (e) {
    console.error('Failed to load collections:', e)
    error.value = 'Failed to load collections. Please try again.'
  } finally {
    loading.value = false
  }
}

/**
 * Applies the search filter to the collections.
 */
function applySearchFilter() {
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    collections.value = allCollections.value.filter(col =>
      col.title.toLowerCase().includes(query) ||
      col.handle?.toLowerCase().includes(query) ||
      col.description?.toLowerCase().includes(query)
    )
  } else {
    collections.value = [...allCollections.value] // Clone to ensure reactivity
  }
  currentPage.value = 1; // Reset to first page after search
}

/**
 * Debounces the search input to avoid excessive filtering.
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

// Navigate to collection detail
function goToDetail(collectionId) {
  router.push({
    name: 'CollectionDetail',
    params: { collectionId }
  })
}

// Navigate to add new collection page
function goToAddCollection() {
  router.push({ name: 'AddCollection' })
}



// Format helpers
function formatDate(dt) {
  return dt
    ? new Date(dt).toLocaleDateString(undefined, {
        year: 'numeric',
        month: 'short',
        day: 'numeric'
      })
    : '—'
}

// Initial data fetch on component mount (align with OrdersPage pattern)
onMounted(async () => {
  try {
    if (!activeShop.value?.id) {
      const ensured = await shopStore.ensureActiveShop()
      if (!ensured?.id) {
        router.replace({ name: 'ShopSelection' })
        return
      }
    }
    await fetchCollections()
  } catch (e) {
    console.error('[CollectionsPage] Initialization failed:', e)
  }
})

// Watch for changes in activeShop to refetch collections if the shop changes
watch(activeShop, (newShop, oldShop) => {
  if (newShop?.id !== oldShop?.id) {
    fetchCollections();
  }
});
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