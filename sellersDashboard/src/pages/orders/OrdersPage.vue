<template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto font-sans">
    <h2 class="text-3xl sm:text-4xl font-extrabold mb-8 text-gray-900 leading-tight">
      Orders
    </h2>

    <div class="flex flex-col sm:flex-row justify-between items-center mb-6 space-y-4 sm:space-y-0">
      <div class="w-full sm:w-1/2 relative">
        <input
          type="text"
          v-model="searchQuery"
          @input="debouncedSearch"
          placeholder="Search orders by customer, order ID, or product name..."
          class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm"
        />
        <SearchIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
      </div>

      <div class="flex space-x-4 items-center">
        <select
          v-model="itemsPerPage"
          @change="onLimitChange"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm"
        >
          <option value="10">10 per page</option>
          <option value="25">25 per page</option>
          <option value="50">50 per page</option>
          <option value="100">100 per page</option>
        </select>

        <div class="inline-flex rounded-lg shadow-sm overflow-hidden border border-gray-200 bg-gray-50" role="group">
          <button
            @click="currentView = 'list'"
            :class="[
              'px-5 py-2.5 text-sm font-medium focus:z-10 focus:ring-2 focus:ring-green-500 focus:outline-none transition-colors duration-200 ease-in-out',
              currentView === 'list'
                ? 'bg-green-600 text-white hover:bg-green-700 shadow-inner'
                : 'bg-gray-50 text-gray-700 hover:bg-gray-200'
            ]"
          >
            <ListIcon class="w-5 h-5 inline-block mr-1" />
            List View
          </button>
          <button
            @click="currentView = 'grid'"
            :class="[
              'px-5 py-2.5 text-sm font-medium focus:z-10 focus:ring-2 focus:ring-green-500 focus:outline-none transition-colors duration-200 ease-in-out',
              currentView === 'grid'
                ? 'bg-green-600 text-white hover:bg-green-700 shadow-inner'
                : 'bg-gray-50 text-gray-700 hover:bg-gray-200'
            ]"
          >
            <GridIcon class="w-5 h-5 inline-block mr-1" />
            Grid View
          </button>
        </div>
      </div>
    </div>

    <!-- Stats Summary -->
    <div v-if="!loading && !error" class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500">Total Orders</p>
            <p class="text-2xl font-semibold text-gray-900">{{ orderStats.total_orders }}</p>
          </div>
        </div>
      </div>
      
      <div class="bg-white rounded-lg shadow p-4">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500">Pending</p>
            <p class="text-2xl font-semibold text-gray-900">{{ orderStats.pending_orders }}</p>
          </div>
        </div>
      </div>
      
      <div class="bg-white rounded-lg shadow p-4">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500">Delivered</p>
            <p class="text-2xl font-semibold text-gray-900">{{ orderStats.delivered_orders }}</p>
          </div>
        </div>
      </div>
      
      <div class="bg-white rounded-lg shadow p-4">
        <div class="flex items-center">
          <div class="p-2 bg-purple-100 rounded-lg">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500">Total Revenue</p>
            <p class="text-2xl font-semibold text-gray-900">{{ formatPrice(orderStats.total_revenue) }}</p>
          </div>
        </div>
      </div>
    </div>

    <div v-if="loading">
      <div v-if="currentView === 'list'" class="overflow-x-auto bg-white shadow-lg rounded-xl animate-pulse">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-100">
            <tr>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Order ID</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Items</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Total</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Status</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Date</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Customer</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="n in itemsPerPage" :key="n" class="bg-white">
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-24"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-32"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-20"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-16"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-20"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-24"></div></td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6 animate-pulse">
        <div v-for="n in itemsPerPage" :key="n" class="bg-white rounded-xl shadow-lg overflow-hidden flex flex-col">
          <div class="p-5 flex-grow flex flex-col space-y-3">
            <div class="h-5 bg-gray-200 rounded w-3/4"></div>
            <div class="h-4 bg-gray-200 rounded w-1/2"></div>
            <div class="h-4 bg-gray-200 rounded w-2/3"></div>
            <div class="h-6 bg-gray-200 rounded w-1/3 mt-auto"></div>
          </div>
        </div>
      </div>
    </div>

    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-lg shadow-sm mb-6" role="alert">
      <strong class="font-bold">Error:</strong>
      <span class="block sm:inline ml-2">{{ error }}</span>
      <p class="text-sm mt-2">Please ensure you have an active shop selected and try again.</p>
      <button @click="fetchOrders" class="text-sm underline mt-2 block">Try again</button>
    </div>

    <div v-else>
      <div v-if="allOrders.length">
        <div v-if="currentView === 'list'" class="overflow-x-auto bg-white shadow-lg rounded-xl">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-100">
              <tr>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Order ID</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Customer</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Items</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Total</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Status</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Date</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-200">
              <tr
                v-for="(order, i) in allOrders"
                :key="order.id"
                @click="goToDetail(order.id)"
                class="cursor-pointer transition duration-150 ease-in-out transform hover:scale-[1.005] hover:bg-green-50"
                :class="{ 'bg-gray-50': i % 2 === 1 }"
              >
                <td class="py-3 px-6 text-sm text-gray-800 font-mono">
                  {{ formatOrderId(order.id, order.orderNumber) }}
                </td>
                <td class="py-3 px-6 text-sm text-gray-800">
                  {{ getCustomerDisplayName(order) }}
                </td>
                <td class="py-3 px-6 text-sm text-gray-700">
                  <div class="flex items-center space-x-2">
                    <span>{{ order.items.length }} item{{ order.items.length !== 1 ? 's' : '' }}</span>
                    <div class="flex -space-x-2">
                      <img
                        v-for="(item, index) in order.items.slice(0, 3)"
                        :key="index"
                        :src="item.image || '/placeholder-product.png'"
                        :alt="item.name"
                        class="w-6 h-6 rounded-full border-2 border-white object-cover"
                        @error="$event.target.style.display='none'"
                      />
                    </div>
                  </div>
                  <div v-if="order.items.length > 3" class="text-xs text-gray-500 mt-1">
                    +{{ order.items.length - 3 }} more
                  </div>
                </td>
                <td class="py-3 px-6 text-sm text-green-600 font-semibold">
                  {{ formatPrice(order.total) }}
                  <div v-if="order.discountTotal > 0" class="text-xs text-gray-500">
                    -{{ formatPrice(order.discountTotal) }}
                  </div>
                </td>
                <td class="py-3 px-6">
                  <span :class="getStatusClass(order.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                    {{ formatStatus(order.status) }}
                  </span>
                </td>
                <td class="py-3 px-6 text-sm text-gray-600">
                  {{ formatDate(order.createdAt) }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
          <div
            v-for="order in allOrders"
            :key="order.id"
            @click="goToDetail(order.id)"
            class="bg-white rounded-xl shadow-lg overflow-hidden cursor-pointer transform hover:scale-105 hover:shadow-xl transition duration-200 ease-in-out flex flex-col group"
          >
            <div class="p-5 flex-grow flex flex-col">
              <div class="flex justify-between items-start mb-3">
                <h3 class="text-lg font-semibold text-gray-900 truncate">
                  {{ formatOrderId(order.id, order.orderNumber) }}
                </h3>
                <span :class="getStatusClass(order.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ formatStatus(order.status) }}
                </span>
              </div>
              
              <div class="space-y-2 mb-4">
                <p class="text-sm text-gray-600">
                  <span class="font-medium">Customer:</span> {{ getCustomerDisplayName(order) }}
                </p>
                <p class="text-sm text-gray-600">
                  <span class="font-medium">Items:</span> {{ order.items.length }} item{{ order.items.length !== 1 ? 's' : '' }}
                </p>
                <p class="text-sm text-gray-600">
                  <span class="font-medium">Date:</span> {{ formatDate(order.createdAt) }}
                </p>
              </div>
              
              <div class="mt-auto">
                <p class="text-2xl font-bold text-green-700">
                  {{ formatPrice(order.total) }}
                </p>
                <p v-if="order.discountTotal > 0" class="text-sm text-gray-500">
                  Discount: {{ formatPrice(order.discountTotal) }}
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
          
          <!-- Page jumpers -->
          <div class="flex space-x-1">
            <button
              v-for="page in getVisiblePages()"
              :key="page"
              @click="goToPage(page)"
              :class="[
                'px-3 py-2 text-sm font-medium rounded-lg transition duration-150 ease-in-out',
                page === currentPage
                  ? 'bg-green-600 text-white'
                  : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
              ]"
            >
              {{ page }}
            </button>
            <span v-if="hasMorePagesBefore()" class="px-2 py-2 text-gray-500">...</span>
            <span v-if="hasMorePagesAfter()" class="px-2 py-2 text-gray-500">...</span>
          </div>
          
          <button
            @click="nextPage"
            :disabled="currentPage === totalPages"
            class="px-4 py-2 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300 disabled:opacity-50 disabled:cursor-not-allowed transition duration-150 ease-in-out"
          >
            Next
          </button>
        </div>

        <div class="text-center mt-2 text-sm text-gray-600">
          Page {{ currentPage }} of {{ totalPages }} ({{ totalOrders }} total orders)
        </div>
      </div>

      <div v-else class="flex flex-col items-center justify-center py-16 text-center text-gray-500">
        <svg xmlns="http://www.w3.org/2000/svg" class="mx-auto h-14 w-14 mb-4 text-green-300" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V7M3 7l9 6 9-6" />
        </svg>
        <p class="text-lg font-semibold">No orders found</p>
        <p class="mt-2 text-sm">Orders will appear here once customers start placing orders in your shop.</p>
        <router-link
          v-if="activeShop && activeShop.id"
          :to="{ name: 'Products' }"
          class="inline-flex items-center mt-6 px-5 py-2.5 bg-green-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition"
        >
          Add Products
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { orderService } from '@/services/order'
import {
  ViewListIcon as ListIcon,
  ViewGridIcon as GridIcon,
  SearchIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const shopStore = useShopStore()

// Reactive state
const allOrders = ref([]) // Holds current page orders
const loading = ref(false)
const error = ref(null)
const currentView = ref('list')

// Search state
const searchQuery = ref('')
let searchTimeout = null // To debounce search input

// Backend pagination state
const currentPage = ref(1)
const itemsPerPage = ref(10)
const totalOrders = ref(0)
const totalPages = ref(0)

// Backend stats state
const orderStats = ref({
  total_orders: 0,
  total_revenue: 0,
  pending_orders: 0,
  delivered_orders: 0
})

// Computed property for the active shop
const activeShop = computed(() => shopStore.activeShop)

/**
 * Gets the count of orders by status from backend stats.
 * @param {string} status - The status to count.
 * @returns {number} The count of orders with that status.
 */
function getStatusCount(status) {
  switch (status) {
    case 'pending':
      return orderStats.value.pending_orders || 0
    case 'delivered':
      return orderStats.value.delivered_orders || 0
    default:
      return 0
  }
}

/**
 * Gets total revenue from backend stats (only paid, shipped, delivered orders).
 * @returns {number} The total revenue.
 */
function getTotalRevenue() {
  return orderStats.value.total_revenue || 0
}

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
 * Fetches orders from the service with pagination.
 */
async function fetchOrders() {
  if (!activeShop.value?.id) {
    router.replace({ name: 'ShopSelection' })
    return
  }

  loading.value = true
  error.value = null
  try {
    // Fetch orders using the shop ID with pagination and search
    const response = await orderService.fetchAllByShop(
      activeShop.value.id, 
      currentPage.value, 
      itemsPerPage.value,
      searchQuery.value
    )
    
    console.log('Fetched orders response:', response) // Debug log
    
    if (response.pagination) {
      // New paginated format
      allOrders.value = response.orders
      totalOrders.value = response.pagination.total
      totalPages.value = response.pagination.total_pages
      currentPage.value = response.pagination.page
      itemsPerPage.value = response.pagination.limit
      
      // Update stats if available
      if (response.stats) {
        orderStats.value = response.stats
      }
    } else {
      // Fallback to old format
      allOrders.value = response.orders || response
      totalOrders.value = allOrders.value.length
      totalPages.value = 1
    }
    
    console.log('Processed orders:', allOrders.value) // Debug log
  } catch (e) {
    console.error('Failed to load orders:', e)
    error.value = 'Failed to load orders. Please try again.'
  } finally {
    loading.value = false
  }
}

/**
 * Applies search filter to the orders.
 * With backend pagination, we refetch data with search
 */
function applyFilters() {
  currentPage.value = 1 // Reset to first page when applying filters
  fetchOrders()
}

/**
 * Debounces the search input to avoid excessive filtering.
 */
function debouncedSearch() {
  clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => {
    applyFilters()
  }, 500) // Increased debounce time for backend requests
}

/**
 * Handles going to the previous page.
 */
function prevPage() {
  if (currentPage.value > 1) {
    currentPage.value--
    fetchOrders()
  }
}

/**
 * Handles going to the next page.
 */
function nextPage() {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
    fetchOrders()
  }
}

/**
 * Gets visible pages for pagination with ellipsis.
 */
function getVisiblePages() {
  const pages = []
  const total = totalPages.value
  const current = currentPage.value
  
  if (total <= 7) {
    // Show all pages if total is 7 or less
    for (let i = 1; i <= total; i++) {
      pages.push(i)
    }
  } else {
    // Show first page, last page, current page, and 2 pages on each side
    if (current <= 4) {
      // Near the beginning
      for (let i = 1; i <= 5; i++) {
        pages.push(i)
      }
      pages.push('...')
      pages.push(total)
    } else if (current >= total - 3) {
      // Near the end
      pages.push(1)
      pages.push('...')
      for (let i = total - 4; i <= total; i++) {
        pages.push(i)
      }
    } else {
      // In the middle
      pages.push(1)
      pages.push('...')
      for (let i = current - 1; i <= current + 1; i++) {
        pages.push(i)
      }
      pages.push('...')
      pages.push(total)
    }
  }
  
  return pages
}

/**
 * Handles jumping to a specific page.
 */
function goToPage(page) {
  if (typeof page === 'number' && page >= 1 && page <= totalPages.value) {
    currentPage.value = page
    fetchOrders()
  }
}

/**
 * Handles limit change and refetches data.
 */
function onLimitChange() {
  currentPage.value = 1 // Reset to first page when changing limit
  fetchOrders()
}

/**
 * Navigates to the order detail page.
 * @param {string} orderId - The ID of the order.
 */
function goToDetail(orderId) {
  router.push({ name: 'OrderDetail', params: { orderId } })
}

/**
 * Formats an order ID for display.
 * @param {string} id - The order ID.
 * @param {string} orderNumber - The order number.
 * @returns {string} Formatted order ID.
 */
function formatOrderId(id, orderNumber) {
  if (orderNumber) {
    return orderNumber
  }
  return id ? `#${id.slice(-8).toUpperCase()}` : 'N/A'
}

/**
 * Formats a customer ID for display.
 * @param {string} id - The customer ID.
 * @returns {string} Formatted customer ID.
 */
function formatCustomerId(id) {
  return id ? `C${id.slice(-6).toUpperCase()}` : 'N/A'
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

/**
 * Formats order status for display.
 * @param {string} status - The order status.
 * @returns {string} Formatted status.
 */
function formatStatus(status) {
  if (!status) return 'Unknown'
  return status.charAt(0).toUpperCase() + status.slice(1).toLowerCase()
}

/**
 * Gets the appropriate CSS classes for order status badges.
 * @param {string} status - The order status.
 * @returns {string} Tailwind CSS classes.
 */
function getStatusClass(status) {
  const statusMap = {
    'pending': 'bg-yellow-100 text-yellow-800',
    'paid': 'bg-blue-100 text-blue-800',
    'shipped': 'bg-purple-100 text-purple-800',
    'delivered': 'bg-green-100 text-green-800',
    'cancelled': 'bg-red-100 text-red-800'
  }
  return statusMap[status?.toLowerCase()] || 'bg-gray-100 text-gray-800'
}

// Helper function to get customer display name (only firstName)
function getCustomerDisplayName(order) {
  const firstName = order.customerFirstName || ''
  const email = order.customerEmail || ''
  
  if (firstName) {
    return firstName
  } else if (email) {
    return email
  } else {
    // Fallback to customer ID if no name or email is available
    return `Customer #${(order.customer_id || order.customerId || '').slice(-6).toUpperCase()}`
  }
}

// Computed property to determine if there are more pages before the current one
function hasMorePagesBefore() {
  return currentPage.value > 2
}

// Computed property to determine if there are more pages after the current one
function hasMorePagesAfter() {
  return currentPage.value < totalPages.value - 1
}

// Initial data fetch on component mount
onMounted(() => {
  fetchOrders()
})

// Watch for changes in activeShop to refetch orders if the shop changes
watch(activeShop, (newShop, oldShop) => {
  if (newShop?.id !== oldShop?.id) {
    fetchOrders()
  }
})
</script>