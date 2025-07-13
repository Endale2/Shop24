  <template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto font-sans">
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center mb-8">
      <div>
        <h2 class="text-3xl sm:text-4xl font-extrabold text-gray-900 leading-tight">
          Collections
        </h2>
        <p class="text-gray-600 mt-2">Manage your collection catalog</p>
      </div>
      <div class="flex flex-col sm:flex-row gap-3 mt-4 sm:mt-0">
        <button
          @click="goToAddCollection"
          class="inline-flex items-center px-4 py-2.5 bg-green-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
        >
          <PlusIcon class="w-5 h-5 mr-2 -ml-1" />
          Add Collection
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
          placeholder="Search collections by name or handle..."
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
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Title</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Handle</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Products</th>
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

    <div v-else>
      <div v-if="Array.isArray(paginatedCollections) && paginatedCollections.length">
        <div v-if="currentView === 'list'" class="overflow-x-auto bg-white shadow-lg rounded-xl">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Image</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Title</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Handle</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Products</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Created</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Actions</th>
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
                <td class="py-3 px-6">
                  <div class="w-12 h-12 rounded-lg overflow-hidden border border-gray-200 flex items-center justify-center bg-gray-100 shadow-sm">
                    <img
                      v-if="col.image"
                      :src="col.image"
                      alt="collection thumbnail"
                      class="w-full h-full object-cover"
                    />
                    <PlaceholderIcon v-else class="w-8 h-8 text-gray-400" />
                  </div>
                </td>
                <td class="py-3 px-6">
                  <div class="text-sm text-gray-800 font-medium">{{ col.title }}</div>
                </td>
                <td class="py-3 px-6 text-sm text-gray-700 font-mono">{{ col.handle }}</td>
                <td class="py-3 px-6 text-sm text-gray-700">{{ Array.isArray(col.products) ? col.products.length : 0 }}</td>
                <td class="py-3 px-6 text-sm text-gray-700">{{ col.created_at ? formatDate(col.created_at) : (col.createdAt ? formatDate(col.createdAt) : 'N/A') }}</td>
                <td class="py-3 px-6">
                  <button
                    @click.stop="goToDetail(col.id)"
                    class="text-green-600 hover:text-green-800 font-medium text-xs px-2 py-1 rounded transition-colors"
                  >
                    View
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <!-- Grid View -->
        <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
          <div
            v-for="col in paginatedCollections"
            :key="col.id"
            @click="goToDetail(col.id)"
            class="bg-white rounded-xl shadow-lg overflow-hidden cursor-pointer transform hover:scale-105 hover:shadow-xl transition duration-200 ease-in-out flex flex-col group"
          >
            <div class="w-full h-48 bg-gray-200 flex items-center justify-center text-gray-400 relative overflow-hidden rounded-t-xl">
              <img
                v-if="col.image"
                :src="col.image"
                alt="collection"
                class="w-full h-full object-cover transform group-hover:scale-110 transition-transform duration-300 ease-in-out"
              />
              <PlaceholderIcon v-else class="w-12 h-12 text-gray-300" />
            </div>
            <div class="p-5 flex-grow flex flex-col">
              <div class="text-lg font-semibold text-gray-800 mb-1">{{ col.title }}</div>
              <div class="text-sm text-gray-500 mb-2">{{ col.handle }}</div>
              <div class="flex items-center text-xs text-gray-600 mt-auto">
                <CubeIcon class="w-4 h-4 mr-1" />
                {{ Array.isArray(col.products) ? col.products.length : 0 }} Products
              </div>
              <div class="text-xs text-gray-400 mt-1">
                Created: {{ col.created_at ? formatDate(col.created_at) : (col.createdAt ? formatDate(col.createdAt) : 'N/A') }}
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
      <div v-else class="text-center py-16 text-gray-500">
        <div class="flex flex-col items-center">
          <div class="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mb-4">
            <PlusIcon class="w-8 h-8 text-green-600" />
          </div>
          <p class="text-lg font-medium mb-2">
            <template v-if="!Array.isArray(paginatedCollections)">
              Unable to load collections. Please try again later.
            </template>
            <template v-else-if="paginatedCollections.length === 0">
              No collections found
            </template>
          </p>
          <p class="text-sm mb-4">Get started by creating your first collection.</p>
          <button
            @click="goToAddCollection"
            class="inline-flex items-center px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
          >
            <PlusIcon class="w-4 h-4 mr-2" />
            Add Collection
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
import { collectionService } from '@/services/collection'
import {
  ViewListIcon as ListIcon,
  ViewGridIcon as GridIcon,
  RefreshIcon as SpinnerIcon,
  PhotographIcon as PlaceholderIcon,
  SearchIcon,
  PlusIcon,
  CollectionIcon,
  CubeIcon,
  StarIcon,
  ClockIcon,
  EyeIcon,
  PencilIcon,
  TrashIcon,
  ExclamationCircleIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const shopStore = useShopStore()

// Reactive state
const allCollections = ref([]) // Holds all fetched collections
const collections = ref([]) // Collections filtered by search, before pagination
const loading     = ref(false)
const error       = ref(null)
const currentView = ref('list')

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

// Statistics computed properties
const totalProducts = computed(() => {
  return allCollections.value.reduce((total, col) => total + (col.productIds?.length || 0), 0)
})

const averageProductsPerCollection = computed(() => {
  if (allCollections.value.length === 0) return 0
  return Math.round(totalProducts.value / allCollections.value.length)
})

const recentlyUpdatedCount = computed(() => {
  const oneWeekAgo = new Date()
  oneWeekAgo.setDate(oneWeekAgo.getDate() - 7)
  
  return allCollections.value.filter(col => {
    if (!col.updatedAt) return false
    return new Date(col.updatedAt) > oneWeekAgo
  }).length
})

// Helper for toggle buttons
function viewButtonClass(view) {
  return view === currentView.value
    ? 'bg-green-600 text-white hover:bg-green-700 shadow-inner'
    : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
}

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

// Edit collection (navigate to detail page with edit mode)
function editCollection(collectionId) {
  router.push({
    name: 'CollectionDetail',
    params: { collectionId }
  })
}

// Delete collection
async function deleteCollection(collectionId, collectionTitle) {
  if (confirm(`Are you sure you want to delete "${collectionTitle}"? This action cannot be undone.`)) {
    try {
      await collectionService.delete(activeShop.value.id, collectionId)
      // Refresh the collections list
      await fetchCollections()
    } catch (error) {
      console.error('Failed to delete collection:', error)
      alert('Failed to delete collection. Please try again.')
    }
  }
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

// Initial data fetch on component mount
onMounted(() => {
  fetchCollections()
})

// Watch for changes in activeShop to refetch collections if the shop changes
watch(activeShop, (newShop, oldShop) => {
  if (newShop?.id !== oldShop?.id) {
    fetchCollections();
  }
});
</script>

<style scoped>
/* No specific scoped styles needed as Tailwind classes handle the design */
</style>