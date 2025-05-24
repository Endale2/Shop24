<!-- src/pages/products/ProductsPage.vue -->
<template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto font-sans">
    <h2 class="text-3xl sm:text-4xl font-extrabold mb-8 text-gray-900 leading-tight">
      Product Catalog
    </h2>

    <!-- View toggle -->
    <div class="flex justify-end mb-6">
      <div class="inline-flex rounded-md shadow-sm" role="group">
        <button
          @click="currentView = 'list'"
          :class="viewButtonClass('list')"
          class="px-4 py-2 text-sm font-medium border border-gray-300 rounded-l-lg focus:z-10 focus:ring-2 focus:ring-blue-500 focus:outline-none transition duration-150 ease-in-out"
        >
          <ListIcon class="w-5 h-5 inline-block mr-1" />
          List View
        </button>
        <button
          @click="currentView = 'grid'"
          :class="viewButtonClass('grid')"
          class="px-4 py-2 text-sm font-medium border border-gray-300 rounded-r-lg focus:z-10 focus:ring-2 focus:ring-blue-500 focus:outline-none transition duration-150 ease-in-out"
        >
          <GridIcon class="w-5 h-5 inline-block mr-1" />
          Grid View
        </button>
      </div>
    </div>

    <!-- Loading state -->
    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-16">
      <SpinnerIcon class="animate-spin h-8 w-8 text-blue-500 mb-3" />
      <p class="text-lg">Loading products...</p>
    </div>

    <!-- Error state -->
    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-lg shadow-sm mb-6" role="alert">
      <strong class="font-bold">Error:</strong>
      <span class="block sm:inline ml-2">{{ error }}</span>
      <p class="text-sm mt-2">Please ensure you have an active shop selected and try again.</p>
    </div>

    <!-- Products -->
    <div v-else>
      <div v-if="products.length">
        <!-- List view -->
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
                v-for="(prod, i) in products"
                :key="prod.id"
                @click="goToDetail(prod.id)"
                class="cursor-pointer transition transform hover:scale-[1.005] hover:bg-blue-50"
                :class="{ 'bg-gray-50': i % 2 === 1 }"
              >
                <td class="py-3 px-6">
                  <div class="w-12 h-12 rounded-lg overflow-hidden border border-gray-200 flex items-center justify-center bg-gray-100">
                    <img
                      v-if="prod.images?.length"
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
                  {{ formatPrice(prod.price) }}
                </td>
                <td class="py-3 px-6 text-sm text-gray-600">{{ formatDate(prod.createdAt) }}</td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Grid view -->
        <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
          <div
            v-for="prod in products"
            :key="prod.id"
            @click="goToDetail(prod.id)"
            class="bg-white rounded-xl shadow-lg overflow-hidden cursor-pointer transform hover:scale-105 hover:shadow-xl transition duration-200 ease-in-out flex flex-col"
          >
            <div class="w-full h-48 bg-gray-200 flex items-center justify-center text-gray-400">
              <img
                v-if="prod.images?.length"
                :src="prod.images[0]"
                alt="product"
                class="w-full h-full object-cover"
              />
              <PlaceholderIcon v-else class="w-16 h-16" />
            </div>
            <div class="p-5 flex-grow flex flex-col">
              <h3 class="text-xl font-semibold text-gray-900 mb-2 truncate">{{ prod.name }}</h3>
              <p class="text-sm text-gray-600 mb-3">{{ prod.category || 'Uncategorized' }}</p>
              <div class="mt-auto">
                <p class="text-2xl font-bold text-green-700">{{ formatPrice(prod.price) }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- No products -->
      <div v-else class="bg-blue-50 border border-blue-200 text-blue-700 px-6 py-8 rounded-lg text-center mt-8 shadow-sm">
        <p class="text-lg font-medium">No products found.</p>
        <p class="mt-2">There are no products listed for your active shop yet.</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { productService } from '@/services/product'
import {
  ViewListIcon as ListIcon,
  ViewGridIcon as GridIcon,
  RefreshIcon as SpinnerIcon,
  PhotographIcon as PlaceholderIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const shopStore = useShopStore()

// Reactive state
const products = ref([])
const loading = ref(false)
const error = ref(null)
const currentView = ref('list')

// Computed
const activeShop = computed(() => shopStore.activeShop)

// Helper to style view toggle buttons
function viewButtonClass(view) {
  return view === currentView.value
    ? 'bg-blue-600 text-white hover:bg-blue-700'
    : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
}

// Fetch products on mount
onMounted(async () => {
  if (!activeShop.value) {
    error.value = 'No shop selected. Please select a shop to view products.'
    return
  }
  loading.value = true
  try {
    products.value = await productService.fetchByShop(activeShop.value.id)
  } catch (e) {
    console.error(e)
    error.value = 'Failed to load products. Please try again later.'
  } finally {
    loading.value = false
  }
})

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

function formatPrice(p) {
  return p != null ? `$${p.toFixed(2)}` : 'N/A'
}

// Navigate to detail
function goToDetail(productId) {
  router.push({ name: 'ProductDetail', params: { productId } })
}
</script>
