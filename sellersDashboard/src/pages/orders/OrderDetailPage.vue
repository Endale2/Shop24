<template>
  <div class="min-h-full bg-gray-50">
    <div class="p-4 sm:p-6 lg:p-8">
      
      <!-- Header Section -->
      <div class="mb-6 sm:mb-8">
        <div class="flex flex-col md:flex-row items-start md:items-center justify-between space-y-4 md:space-y-0">
          <div class="flex items-center mb-3 md:mb-0">
            <button
              @click="goBack"
              class="inline-flex items-center text-gray-600 hover:text-green-700 transition duration-200 ease-in-out mr-4 group"
            >
              <ArrowLeftIcon class="h-5 w-5 mr-1 text-gray-500 group-hover:text-green-600 transition-colors duration-200" />
              <span class="text-sm font-medium group-hover:text-green-700 transition-colors duration-200">Back to Orders</span>
            </button>
      
          </div>
          
          <div class="flex items-center space-x-3">
            <button
              v-if="order"
              @click="updateOrderStatus"
              class="inline-flex items-center px-4 py-2.5 bg-green-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
            >
              <PencilIcon class="w-4 h-4 mr-1.5" />
              Update Status
            </button>
          </div>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-16 text-gray-500">
        <div class="flex flex-col items-center justify-center">
          <div class="animate-spin h-8 w-8 text-green-500 mb-3 border-3 border-green-200 border-t-green-500 rounded-full"></div>
          <p class="text-sm font-medium">Loading order details...</p>
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
        <p class="text-sm mt-2">Please try again or contact support if the issue persists.</p>
        <button @click="fetchOrder" class="text-sm underline mt-2 block">Try again</button>
      </div>

      <!-- Order Content -->
      <main v-else-if="order" class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        
        <div class="lg:col-span-2 space-y-6">
          <!-- Order Items Section -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-sm font-semibold text-gray-900">Order Items</h2>
            </div>
            <div class="p-4">
              <div class="flow-root">
                <ul role="list" class="-my-4 divide-y divide-gray-200">
                  <li v-for="item in order.items" :key="`${item.productId}-${item.variantId || 'no-variant'}`" class="flex items-center py-4 space-x-4 group hover:bg-green-50 transition cursor-pointer rounded-lg" @click="goToProductDetail(item.productId)">
                    <div class="flex-shrink-0">
                      <img :src="item.image || '/placeholder-product.png'" :alt="item.name" class="h-12 w-12 rounded-lg object-cover border border-gray-200 group-hover:ring-2 group-hover:ring-green-400 transition" @error="$event.target.src='/placeholder-product.png'" />
                    </div>
                    <div class="flex-1 min-w-0">
                      <p class="text-sm font-semibold text-gray-900 truncate">{{ item.name }}</p>
                      <p v-if="item.variantId && item.variantId !== '000000000000000000000000'" class="text-xs text-gray-500">{{ item.variantName || 'Variant' }}</p>
                      <p v-else class="text-xs text-gray-400 italic">No variant</p>
                    </div>
                    <div class="text-right">
                      <p class="text-sm font-medium text-gray-900">{{ formatPrice(item.totalPrice) }}</p>
                      <p class="text-xs text-gray-500">{{ item.quantity }} x {{ formatPrice(item.unitPrice) }}</p>
                    </div>
                    <span class="ml-2 text-gray-400 group-hover:text-green-500 transition" title="View Product">
                      <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" /></svg>
                    </span>
                  </li>
                </ul>
              </div>
            </div>
          </div>

          <!-- Customer Card Section -->
          <div v-if="customer" class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden hover:shadow-md transition cursor-pointer group" @click="goToCustomerProfile(customer.id)">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-sm font-semibold text-gray-900 flex items-center">
                Customer
                <span class="ml-2 text-gray-400 group-hover:text-green-500 transition" title="View Customer Profile">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" /></svg>
                </span>
              </h2>
            </div>
            <div class="p-4">
              <div class="flex items-start space-x-4">
                <template v-if="customer.profile_image">
                  <img :src="customer.profile_image" :alt="`${customer.firstName} ${customer.lastName}`" class="h-12 w-12 rounded-full object-cover border border-gray-200 group-hover:ring-2 group-hover:ring-green-400 transition" @error="$event.target.src='/placeholder-avatar.png'" />
                </template>
                <template v-else>
                  <div class="h-12 w-12 rounded-full bg-gradient-to-br from-green-400 to-teal-500 flex items-center justify-center text-white text-sm font-bold border border-gray-200 group-hover:ring-2 group-hover:ring-green-400 transition">
                    {{ getInitials(customer) }}
                  </div>
                </template>
                <div class="flex-1 min-w-0">
                  <p class="text-sm font-semibold text-gray-900">{{ customer.firstName }} {{ customer.lastName }}</p>
                  <p class="text-xs text-gray-500 truncate">{{ customer.email }}</p>
                  <p v-if="customer.phone" class="text-xs text-gray-500">{{ customer.phone }}</p>
                </div>
              </div>
              <div v-if="customer.address" class="mt-4 pt-4 border-t border-gray-200">
                <p class="text-xs font-medium text-gray-600">Shipping Address</p>
                <address class="text-xs text-gray-500 not-italic">
                  {{ customer.address }}<br>
                  <span v-if="customer.city">{{ customer.city }}, </span>
                  <span v-if="customer.state">{{ customer.state }} </span>
                  <span v-if="customer.postalCode">{{ customer.postalCode }}</span>
                </address>
              </div>
            </div>
          </div>
        </div>
        
        <div class="lg:col-span-1 space-y-6">
          <!-- Summary Card -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-sm font-semibold text-gray-900">Summary</h2>
            </div>
            <div class="p-4">
              <div class="space-y-4">
                <div class="flex justify-between items-center">
                  <span class="text-xs font-medium text-gray-500">Status</span>
                  <span :class="getStatusClass(order.status)" class="inline-block px-2 py-1 text-xs font-semibold rounded-full">
                    {{ formatStatus(order.status) }}
                  </span>
                </div>
                <div class="flex justify-between items-center">
                  <span class="text-xs text-gray-500">Subtotal</span>
                  <span class="text-xs font-medium text-gray-900">{{ formatPrice(order.subtotal) }}</span>
                </div>
                <div v-if="order.discountTotal > 0" class="flex justify-between items-center">
                  <span class="text-xs text-gray-500">Discount</span>
                  <span class="text-xs font-medium text-green-600">-{{ formatPrice(order.discountTotal) }}</span>
                </div>
                <div class="flex justify-between items-baseline pt-4 border-t border-gray-200">
                  <span class="text-sm font-bold text-gray-900">Total</span>
                  <span class="text-lg font-bold text-green-600">{{ formatPrice(order.total) }}</span>
                </div>
              </div>
            </div>
          </div>

          <!-- Additional Info Card -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-sm font-semibold text-gray-900">Additional Info</h2>
            </div>
            <div class="p-4">
              <dl class="space-y-3">
                <div class="flex justify-between">
                  <dt class="text-xs text-gray-500">Order #</dt>
                  <dd class="text-xs font-mono text-gray-900">{{ formatOrderId(order.id, order.orderNumber) }}</dd>
                </div>
                <div class="flex justify-between">
                  <dt class="text-xs text-gray-500">Customer ID</dt>
                  <dd class="text-xs font-mono text-gray-900">{{ customer?.id || '—' }}</dd>
                </div>
                <div v-if="order.couponCode" class="flex justify-between">
                  <dt class="text-xs text-gray-500">Coupon</dt>
                  <dd class="text-xs font-mono text-gray-900">{{ order.couponCode }}</dd>
                </div>
                <div class="flex justify-between">
                  <dt class="text-xs text-gray-500">Item Count</dt>
                  <dd class="text-xs text-gray-900">{{ order.items?.length || 0 }}</dd>
                </div>
              </dl>
            </div>
          </div>
        </div>
      </main>

      <!-- Not Found State -->
      <div v-else class="text-center py-16 text-gray-400">
        <div class="max-w-md mx-auto">
          <div class="w-12 h-12 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-3">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 text-gray-300" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
            </svg>
          </div>
          <p class="text-sm font-medium text-gray-600 mb-1">Order Not Found</p>
          <p class="text-xs">We couldn't find the order you're looking for.</p>
          <button @click="goBack" class="mt-4 text-sm font-medium text-green-600 hover:text-green-500">Go back to orders</button>
        </div>
      </div>
    </div>

    <!-- Status Update Modal -->
    <div v-if="showStatusModal" class="fixed inset-0 z-50 flex items-center justify-center bg-white/30 backdrop-blur-none transition-opacity" @click.self="showStatusModal = false">
      <div class="relative bg-white rounded-lg shadow-xl w-full max-w-md p-6 m-4 border border-gray-200">
        <h3 class="text-lg font-bold text-gray-900 mb-4 flex items-center gap-2">
          <PencilIcon class="w-4 h-4 text-green-500" />
          Update Order Status
        </h3>
        <p class="text-sm text-gray-500 mb-6">Select a new status for this order. The customer may be notified of this change.</p>
        <div class="mb-6">
          <label for="order-status" class="block text-sm font-medium text-gray-700 mb-2">Order Status</label>
          <select
            id="order-status"
            v-model="newStatus"
            class="w-full px-4 py-2.5 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500 bg-gray-50 text-gray-900 transition-colors text-sm"
          >
            <option value="pending">Pending</option>
            <option value="paid">Paid</option>
            <option value="shipped">Shipped</option>
            <option value="delivered">Delivered</option>
            <option value="cancelled">Cancelled</option>
          </select>
        </div>
        <div class="border-t border-gray-100 pt-4 flex justify-end space-x-3">
          <button
            @click="showStatusModal = false"
            class="px-4 py-2 text-gray-700 bg-gray-100 rounded-lg hover:bg-gray-200 transition-colors border border-gray-200 focus:outline-none focus:ring-2 focus:ring-green-400 text-sm"
          >
            Cancel
          </button>
          <button
            @click="saveStatusUpdate"
            :disabled="updatingStatus"
            class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 disabled:bg-green-300 disabled:cursor-not-allowed transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 font-semibold shadow-sm text-sm"
          >
            {{ updatingStatus ? 'Updating...' : 'Save Update' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { orderService } from '@/services/order'
import {
  ArrowLeftIcon,
  PencilIcon,
  CheckIcon,
  ExclamationCircleIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const route = useRoute()
const shopStore = useShopStore()

// Reactive state
const order = ref(null)
const customer = ref(null)
const loading = ref(false)
const error = ref(null)
const showStatusModal = ref(false)
const newStatus = ref('')
const updatingStatus = ref(false)

// Computed property for the active shop
const activeShop = computed(() => shopStore.activeShop)

// Computed property to check if order has variants
const hasVariants = computed(() => {
  if (!order.value?.items) return false
  return order.value.items.some(item => 
    item.variantId && item.variantId !== '000000000000000000000000'
  )
})

/**
 * Fetches the order details with customer information.
 */
async function fetchOrder() {
  if (!activeShop.value?.id) {
    router.replace({ name: 'ShopSelection' })
    return
  }

  loading.value = true
  error.value = null
  try {
    const orderData = await orderService.fetchByIdWithCustomer(activeShop.value.id, route.params.orderId)
    if (orderData) {
      order.value = {
        ...orderData.order,
        id: orderData.order.id || orderData.order._id,
      }
      customer.value = orderData.customer
    } else {
      order.value = null
      customer.value = null
    }
  } catch (e) {
    console.error('Failed to load order:', e)
    error.value = 'Failed to load order details. Please try again.'
  } finally {
    loading.value = false
  }
}

/**
 * Navigates back to the orders list.
 */
function goBack() {
  router.push({ name: 'Orders' })
}

/**
 * Opens the status update modal.
 */
function updateOrderStatus() {
  newStatus.value = order.value.status
  showStatusModal.value = true
}

/**
 * Saves the status update.
 */
async function saveStatusUpdate() {
  if (!order.value || newStatus.value === order.value.status) {
    showStatusModal.value = false
    return
  }

  updatingStatus.value = true
  try {
    const updatedOrder = await orderService.update(activeShop.value.id, order.value.id, {
      status: newStatus.value
    })
    order.value = {
      ...updatedOrder,
      id: updatedOrder.id || updatedOrder._id,
    }
    showStatusModal.value = false
  } catch (e) {
    console.error('Failed to update order status:', e)
    error.value = 'Failed to update order status. Please try again.'
  } finally {
    updatingStatus.value = false
  }
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
 * Formats a product ID for display.
 * @param {string} id - The product ID.
 * @returns {string} Formatted product ID.
 */
function formatProductId(id) {
  return id ? `P${id.slice(-6).toUpperCase()}` : 'N/A'
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
 * Formats a date and time string.
 * @param {string} dt - The date string.
 * @returns {string} Formatted date and time.
 */
function formatDateTime(dt) {
  return dt
    ? new Date(dt).toLocaleString(undefined, {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
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
    'paid': 'bg-green-100 text-green-800',
    'shipped': 'bg-purple-100 text-purple-800',
    'delivered': 'bg-green-100 text-green-800',
    'cancelled': 'bg-red-100 text-red-800'
  }
  return statusMap[status?.toLowerCase()] || 'bg-gray-100 text-gray-800'
}

function goToCustomerProfile(customerId) {
  router.push({ name: 'CustomerDetail', params: { customerId } })
}
function goToProductDetail(productId) {
  router.push({ name: 'ProductDetail', params: { productId } })
}

function getInitials(customer) {
  const fi = customer.firstName?.[0]?.toUpperCase() || ''
  const li = customer.lastName?.[0]?.toUpperCase() || ''
  return fi + li || '?'
}

// Initial data fetch on component mount
onMounted(() => {
  fetchOrder()
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