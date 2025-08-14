<template>
  <div class="min-h-full bg-gray-50">
    <div class="p-4 sm:p-6 lg:p-8">
      
      <!-- Header Section -->
      <div class="mb-6 sm:mb-8">
        <div class="flex flex-col md:flex-row items-start md:items-center justify-between space-y-4 md:space-y-0">
          <div class="flex items-center mb-3 md:mb-0">
            <div class="w-10 h-10 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center mr-3 shadow-sm">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
              </svg>
            </div>
            <div>
              <h1 class="text-2xl sm:text-3xl font-bold text-gray-900 tracking-tight">
                Orders
              </h1>
              <p class="text-sm text-gray-600 mt-1">Manage and track customer orders</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Summary Cards -->
      <div v-if="!loading && !error" class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-4">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4" />
              </svg>
            </div>
            <div class="ml-3">
              <p class="text-sm font-medium text-gray-500">Completed Orders</p>
              <p class="text-xl font-semibold text-gray-900">{{ completedOrdersCount }}</p>
              <p class="text-xs text-gray-400 mt-1">(Excludes Pending and Cancelled)</p>
            </div>
          </div>
        </div>
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-4">
          <div class="flex items-center">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
              </svg>
            </div>
            <div class="ml-3">
              <p class="text-sm font-medium text-gray-500">Total Revenue</p>
              <p class="text-xl font-semibold text-gray-900">{{ formatPrice(orderStats.total_revenue) }}</p>
              <p class="text-xs text-gray-400 mt-1">(Paid, Shipped, and Delivered orders only)</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Status Filter Tabs -->
      <div v-if="!loading && !error" class="bg-white rounded-lg shadow-sm border border-gray-200 p-4 mb-6">
        <div class="flex flex-wrap gap-2">
          <button
            v-for="tab in statusTabs"
            :key="tab.value"
            type="button"
            @click.prevent="onTabClick(tab.value)"
            :class="[
              'px-3 py-2 rounded-lg text-xs font-medium transition-colors duration-200 flex items-center',
              statusFilter === tab.value
                ? 'bg-green-600 text-white'
                : getTabColor(tab.value)
            ]"
          >
            {{ tab.label }}
            <span class="ml-2 bg-gray-200 text-gray-700 px-2 py-0.5 rounded-full text-xs" :class="statusFilter === tab.value ? 'bg-white text-green-700' : ''">
              {{ getTabCount(tab.value) }}
            </span>
          </button>
        </div>
      </div>

      <!-- Search Bar -->
      <div class="flex flex-col sm:flex-row justify-between items-center mb-6 space-y-4 sm:space-y-0">
        <div class="w-full sm:w-1/2 relative">
          <input
            type="text"
            v-model="searchQuery"
            @input="debouncedSearch"
            placeholder="Search orders by customer, order ID, or product name..."
            class="w-full pl-10 pr-4 py-2.5 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm text-sm"
          />
          <SearchIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
        </div>
        <div></div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-16 text-gray-500">
        <div class="flex flex-col items-center justify-center">
          <div class="animate-spin h-8 w-8 text-green-500 mb-3 border-3 border-green-200 border-t-green-500 rounded-full"></div>
          <p class="text-sm font-medium">Loading orders...</p>
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
        <button @click="fetchOrders" class="text-sm underline mt-2 block">Try again</button>
      </div>

      <!-- Orders Content -->
      <div v-else>
        <div v-if="orders.length">
          <div v-if="currentView === 'list'" class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="overflow-x-auto">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="px-4 py-3 text-left text-xs font-semibold text-gray-900 uppercase tracking-wider">Order ID</th>
                    <th class="px-4 py-3 text-left text-xs font-semibold text-gray-900 uppercase tracking-wider">Customer</th>
                    <th class="px-4 py-3 text-left text-xs font-semibold text-gray-900 uppercase tracking-wider">Items</th>
                    <th class="px-4 py-3 text-left text-xs font-semibold text-gray-900 uppercase tracking-wider">Total</th>
                    <th class="px-4 py-3 text-left text-xs font-semibold text-gray-900 uppercase tracking-wider">Status</th>
                    <th class="px-4 py-3 text-left text-xs font-semibold text-gray-900 uppercase tracking-wider">Date</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-gray-200">
                  <tr
                    v-for="(order, i) in orders"
                    :key="order.id"
                    @click="goToDetail(order.id)"
                    class="cursor-pointer transition duration-150 ease-in-out transform hover:scale-[1.005] hover:bg-green-50"
                    :class="{ 'bg-gray-50': i % 2 === 1 }"
                  >
                    <td class="py-3 px-4 text-sm text-gray-800 font-mono">
                      {{ formatOrderId(order.id, order.orderNumber) }}
                    </td>
                    <td class="py-3 px-4 text-sm text-gray-800">
                      {{ getCustomerDisplayName(order) }}
                    </td>
                    <td class="py-3 px-4 text-sm text-gray-700">
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
                    <td class="py-3 px-4 text-sm text-green-600 font-semibold">
                      {{ formatPrice(order.total) }}
                      <div v-if="order.discountTotal > 0" class="text-xs text-gray-500">
                        -{{ formatPrice(order.discountTotal) }}
                      </div>
                    </td>
                    <td class="py-3 px-4">
                      <span :class="getStatusClass(order.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                        {{ formatStatus(order.status) }}
                      </span>
                    </td>
                    <td class="py-3 px-4 text-sm text-gray-600">
                      {{ formatDate(order.createdAt) }}
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>

          <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
            <div
              v-for="order in orders"
              :key="order.id"
              @click="goToDetail(order.id)"
              class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden cursor-pointer transform hover:scale-105 hover:shadow-md transition duration-200 ease-in-out flex flex-col group"
            >
              <div class="p-4 flex-grow flex flex-col">
                <div class="flex justify-between items-start mb-3">
                  <h3 class="text-sm font-semibold text-gray-900 truncate">
                    {{ formatOrderId(order.id, order.orderNumber) }}
                  </h3>
                  <span :class="getStatusClass(order.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                    {{ formatStatus(order.status) }}
                  </span>
                </div>
                
                <div class="space-y-2 mb-4">
                  <p class="text-xs text-gray-600">
                    <span class="font-medium">Customer:</span> {{ getCustomerDisplayName(order) }}
                  </p>
                  <p class="text-xs text-gray-600">
                    <span class="font-medium">Items:</span> {{ order.items.length }} item{{ order.items.length !== 1 ? 's' : '' }}
                  </p>
                  <p class="text-xs text-gray-600">
                    <span class="font-medium">Date:</span> {{ formatDate(order.createdAt) }}
                  </p>
                </div>
                
                <div class="mt-auto">
                  <p class="text-lg font-bold text-green-700">
                    {{ formatPrice(order.total) }}
                  </p>
                  <p v-if="order.discountTotal > 0" class="text-xs text-gray-500">
                    Discount: {{ formatPrice(order.discountTotal) }}
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
              class="px-4 py-2 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300 disabled:opacity-50 disabled:cursor-not-allowed transition duration-150 ease-in-out text-sm"
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
              class="px-4 py-2 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300 disabled:opacity-50 disabled:cursor-not-allowed transition duration-150 ease-in-out text-sm"
            >
              Next
            </button>
          </div>

          <div class="text-center mt-2 text-sm text-gray-600">
            Page {{ currentPage }} of {{ totalPages }} ({{ totalOrders }} total orders)
          </div>
        </div>

        <!-- Empty State -->
        <div v-else class="text-center py-16 text-gray-400">
          <div class="max-w-md mx-auto">
            <div class="w-12 h-12 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-3">
              <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 text-gray-300" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V7M3 7l9 6 9-6" />
              </svg>
            </div>
            <p class="text-sm font-medium text-gray-600 mb-1">No orders found</p>
            <p class="text-xs">Orders will appear here once customers start placing orders in your shop.</p>
            <router-link
              v-if="activeShop && activeShop.id"
              :to="{ name: 'Products' }"
              class="inline-flex items-center mt-4 px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition"
            >
              Add Products
            </router-link>
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
import { orderService } from '@/services/order'
import {
  ViewListIcon as ListIcon,
  ViewGridIcon as GridIcon,
  SearchIcon,
  ExclamationCircleIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const shopStore = useShopStore()

// Reactive state
const orders = ref([]) // Holds current page orders from server
const loading = ref(false)
const error = ref(null)
const currentView = ref('list')

// Search state
const searchQuery = ref('')
let searchTimeout = null // To debounce search input

// Backend pagination state
const currentPage = ref(1)
const itemsPerPage = ref(25) // Fixed value, not user-changeable
const totalOrders = ref(0)
const totalPages = ref(0)

// Status filter state
const statusFilter = ref('all')

// Order statistics
const orderStats = ref({
  total: 0,
  pending: 0,
  processing: 0,
  shipped: 0,
  delivered: 0,
  cancelled: 0
})

// Status filter tabs
const statusTabs = [
  { label: 'All', value: 'all' },
  { label: 'Pending', value: 'pending' },
  { label: 'Paid', value: 'paid' },
  { label: 'Shipped', value: 'shipped' },
  { label: 'Delivered', value: 'delivered' },
  { label: 'Cancelled', value: 'cancelled' },
]

function onTabClick(val) {
  if (statusFilter.value !== val) {
    statusFilter.value = val
    currentPage.value = 1 // Reset to first page when changing status
    fetchOrders() // Use server-side filtering for better performance
  }
}



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
      return orderStats.value.pending || 0
    case 'processing':
      return orderStats.value.processing || 0
    case 'shipped':
      return orderStats.value.shipped || 0
    case 'delivered':
      return orderStats.value.delivered || 0
    case 'cancelled':
      return orderStats.value.cancelled || 0
    default:
      return orderStats.value.total || 0
  }
}

/**
 * Gets total revenue from backend stats (only paid, shipped, delivered orders).
 * @returns {number} The total revenue.
 */
function getTotalRevenue() {
  return orderStats.value.total_revenue || orderStats.value.revenue || 0
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
 * Fetches orders from the service with pagination and status filter.
 */
async function fetchOrders() {
  if (!activeShop.value?.id) {
    router.replace({ name: 'ShopSelection' })
    return
  }

  loading.value = true
  error.value = null
  try {
    console.log('[OrdersPage] Fetching orders with server-side filtering:', {
      shopId: activeShop.value.id,
      page: currentPage.value,
      limit: itemsPerPage.value,
      search: searchQuery.value,
      status: statusFilter.value
    })

    // Use server-side filtering for optimal performance
    const response = await orderService.fetchAllByShop(
      activeShop.value.id,
      currentPage.value,
      itemsPerPage.value,
      searchQuery.value,
      statusFilter.value // Pass status filter to server for better performance
    )
    
    console.log('Fetched orders response:', response) // Debug log
    
    if (response.pagination) {
      // New paginated format with server-side filtering
      orders.value = response.orders
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
      orders.value = response.orders || response
      totalOrders.value = orders.value.length
      totalPages.value = 1
    }

    console.log('[OrdersPage] Orders processed successfully:', {
      count: orders.value.length,
      total: totalOrders.value,
      page: currentPage.value,
      status: statusFilter.value
    })
  } catch (e) {
    console.error('Failed to load orders:', e)
    error.value = 'Failed to load orders. Please try again.'
  } finally {
    loading.value = false
  }
}

/**
 * Applies filters by fetching from server with current filter state.
 * This ensures optimal performance by using server-side filtering.
 */
function applyFilters() {
  currentPage.value = 1 // Reset to first page when applying filters
  fetchOrders()
}

/**
 * Debounces the search input to avoid excessive server requests.
 */
function debouncedSearch() {
  clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => {
    // Use server-side filtering for optimal performance
    applyFilters()
  }, 500) // Appropriate debounce time for server requests
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

// Add a computed property for completed orders (excluding pending and cancelled)
const completedOrdersCount = computed(() => {
  const stats = orderStats.value
  return (
    (stats.total_orders || 0)
    - (stats.pending_orders || 0)
    - (stats.cancelled_orders || 0)
  )
})

// Helper to get the count for each tab
function getTabCount(val) {
  const stats = orderStats.value
  switch (val) {
    case 'all': return stats.total_orders ?? 0
    case 'pending': return stats.pending_orders ?? 0
    case 'paid': return stats.paid_orders ?? 0
    case 'shipped': return stats.shipped_orders ?? 0
    case 'delivered': return stats.delivered_orders ?? 0
    case 'cancelled': return stats.cancelled_orders ?? 0
    default: return null
  }
}

// Helper to get color classes for each status tab
function getTabColor(val) {
  switch (val) {
    case 'pending': return 'bg-yellow-100 text-yellow-800 hover:bg-yellow-200';
    case 'paid': return 'bg-blue-100 text-blue-800 hover:bg-blue-200';
    case 'shipped': return 'bg-purple-100 text-purple-800 hover:bg-purple-200';
    case 'delivered': return 'bg-green-100 text-green-800 hover:bg-green-200';
    case 'cancelled': return 'bg-red-100 text-red-800 hover:bg-red-200';
    case 'all':
    default: return 'bg-gray-100 text-gray-700 hover:bg-gray-200';
  }
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