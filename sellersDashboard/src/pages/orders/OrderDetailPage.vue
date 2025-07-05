<template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto font-sans">
    <!-- Header -->
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center mb-8 space-y-4 sm:space-y-0">
      <div>
        <h2 class="text-3xl sm:text-4xl font-extrabold text-gray-900 leading-tight">
          Order Details
        </h2>
        <p v-if="order" class="text-lg text-gray-600 mt-2">
          {{ formatOrderId(order.id) }} • {{ formatDate(order.createdAt) }}
        </p>
      </div>
      
      <div class="flex space-x-3">
        <button
          @click="goBack"
          class="inline-flex items-center px-4 py-2 bg-gray-600 text-white text-sm font-medium rounded-lg shadow-md hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 transition duration-150 ease-in-out"
        >
          <ArrowLeftIcon class="w-5 h-5 mr-2 -ml-1" />
          Back to Orders
        </button>
        
        <button
          v-if="order"
          @click="updateOrderStatus"
          class="inline-flex items-center px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg shadow-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
        >
          <PencilIcon class="w-5 h-5 mr-2 -ml-1" />
          Update Status
        </button>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="bg-white rounded-xl shadow-lg p-8 animate-pulse">
      <div class="space-y-6">
        <div class="h-8 bg-gray-200 rounded w-1/3"></div>
        <div class="h-4 bg-gray-200 rounded w-1/4"></div>
        <div class="space-y-4">
          <div v-for="n in 3" :key="n" class="h-20 bg-gray-200 rounded"></div>
        </div>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-lg shadow-sm mb-6" role="alert">
      <strong class="font-bold">Error:</strong>
      <span class="block sm:inline ml-2">{{ error }}</span>
      <button @click="fetchOrder" class="text-sm underline mt-2 block">Try again</button>
    </div>

    <!-- Order Details -->
    <div v-else-if="order" class="space-y-6">
      <!-- Order Summary Card -->
      <div class="bg-white rounded-xl shadow-lg overflow-hidden">
        <div class="px-6 py-4 bg-gray-50 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900">Order Summary</h3>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            <div>
              <p class="text-sm font-medium text-gray-500">Order ID</p>
              <p class="text-lg font-mono text-gray-900">{{ formatOrderId(order.id) }}</p>
            </div>
            <div>
              <p class="text-sm font-medium text-gray-500">Customer ID</p>
              <p class="text-lg font-mono text-gray-900">{{ formatCustomerId(order.customerId) }}</p>
            </div>
            <div>
              <p class="text-sm font-medium text-gray-500">Status</p>
              <span :class="getStatusClass(order.status)" class="inline-block px-3 py-1 text-sm font-medium rounded-full">
                {{ formatStatus(order.status) }}
              </span>
            </div>
            <div>
              <p class="text-sm font-medium text-gray-500">Total</p>
              <p class="text-2xl font-bold text-green-600">{{ formatPrice(order.total) }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Order Items -->
      <div class="bg-white rounded-xl shadow-lg overflow-hidden">
        <div class="px-6 py-4 bg-gray-50 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900">Order Items</h3>
        </div>
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Product</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Variant</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Quantity</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Unit Price</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Total</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="item in order.items" :key="`${item.productId}-${item.variantId || 'no-variant'}`">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="flex-shrink-0 h-12 w-12">
                      <img
                        :src="item.image || '/placeholder-product.png'"
                        :alt="item.name"
                        class="h-12 w-12 rounded-lg object-cover border border-gray-200"
                        @error="$event.target.src='/placeholder-product.png'"
                      />
                    </div>
                    <div class="ml-4">
                      <div class="text-sm font-medium text-gray-900">{{ item.name }}</div>
                      <div class="text-sm text-gray-500">ID: {{ formatProductId(item.productId) }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  <span v-if="item.variantId && item.variantId !== '000000000000000000000000'" class="text-gray-700">
                    {{ item.variantName || 'Variant' }}
                  </span>
                  <span v-else class="text-gray-400 italic">No variant</span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ item.quantity }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ formatPrice(item.unitPrice) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                  {{ formatPrice(item.totalPrice) }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Order Totals -->
      <div class="bg-white rounded-xl shadow-lg overflow-hidden">
        <div class="px-6 py-4 bg-gray-50 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900">Order Totals</h3>
        </div>
        <div class="p-6">
          <div class="space-y-3">
            <div class="flex justify-between">
              <span class="text-gray-600">Subtotal:</span>
              <span class="font-medium">{{ formatPrice(order.subtotal) }}</span>
            </div>
            <div v-if="order.discountTotal > 0" class="flex justify-between">
              <span class="text-gray-600">Discount:</span>
              <span class="font-medium text-green-600">-{{ formatPrice(order.discountTotal) }}</span>
            </div>
            <div class="flex justify-between text-lg font-semibold border-t pt-3">
              <span>Total:</span>
              <span class="text-green-600">{{ formatPrice(order.total) }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Order Timeline -->
      <div class="bg-white rounded-xl shadow-lg overflow-hidden">
        <div class="px-6 py-4 bg-gray-50 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900">Order Timeline</h3>
        </div>
        <div class="p-6">
          <div class="space-y-4">
            <div class="flex items-center">
              <div class="flex-shrink-0">
                <div class="w-8 h-8 bg-green-500 rounded-full flex items-center justify-center">
                  <CheckIcon class="w-5 h-5 text-white" />
                </div>
              </div>
              <div class="ml-4">
                <p class="text-sm font-medium text-gray-900">Order Created</p>
                <p class="text-sm text-gray-500">{{ formatDateTime(order.createdAt) }}</p>
              </div>
            </div>
            <div v-if="order.updatedAt && order.updatedAt !== order.createdAt" class="flex items-center">
              <div class="flex-shrink-0">
                <div class="w-8 h-8 bg-blue-500 rounded-full flex items-center justify-center">
                  <PencilIcon class="w-5 h-5 text-white" />
                </div>
              </div>
              <div class="ml-4">
                <p class="text-sm font-medium text-gray-900">Last Updated</p>
                <p class="text-sm text-gray-500">{{ formatDateTime(order.updatedAt) }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Additional Order Info -->
      <div class="bg-white rounded-xl shadow-lg overflow-hidden">
        <div class="px-6 py-4 bg-gray-50 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900">Additional Information</h3>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <p class="text-sm font-medium text-gray-500">Shop ID</p>
              <p class="text-sm text-gray-900 font-mono">{{ order.shopId || 'N/A' }}</p>
            </div>
            <div>
              <p class="text-sm font-medium text-gray-500">Items Count</p>
              <p class="text-sm text-gray-900">{{ order.items?.length || 0 }} items</p>
            </div>
            <div v-if="order.couponCode">
              <p class="text-sm font-medium text-gray-500">Coupon Used</p>
              <p class="text-sm text-gray-900 font-mono">{{ order.couponCode }}</p>
            </div>
            <div>
              <p class="text-sm font-medium text-gray-500">Order Type</p>
              <p class="text-sm text-gray-900">{{ hasVariants ? 'Products with variants' : 'Simple products' }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Not Found State -->
    <div v-else class="bg-yellow-50 border border-yellow-200 text-yellow-700 px-6 py-8 rounded-lg text-center mt-8 shadow-sm">
      <p class="text-lg font-medium">Order not found.</p>
      <p class="mt-2">The order you're looking for doesn't exist or has been removed.</p>
      <button @click="goBack" class="mt-4 text-sm underline">Go back to orders</button>
    </div>

    <!-- Status Update Modal -->
    <div v-if="showStatusModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-96 shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">Update Order Status</h3>
          <select
            v-model="newStatus"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
          >
            <option value="pending">Pending</option>
            <option value="paid">Paid</option>
            <option value="shipped">Shipped</option>
            <option value="delivered">Delivered</option>
            <option value="cancelled">Cancelled</option>
          </select>
          <div class="flex justify-end space-x-3 mt-6">
            <button
              @click="showStatusModal = false"
              class="px-4 py-2 bg-gray-300 text-gray-700 rounded-lg hover:bg-gray-400 transition duration-150 ease-in-out"
            >
              Cancel
            </button>
            <button
              @click="saveStatusUpdate"
              :disabled="updatingStatus"
              class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 disabled:opacity-50 disabled:cursor-not-allowed transition duration-150 ease-in-out"
            >
              {{ updatingStatus ? 'Updating...' : 'Update' }}
            </button>
          </div>
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
  CheckIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const route = useRoute()
const shopStore = useShopStore()

// Reactive state
const order = ref(null)
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
 * Fetches the order details.
 */
async function fetchOrder() {
  if (!activeShop.value?.id) {
    router.replace({ name: 'ShopSelection' })
    return
  }

  loading.value = true
  error.value = null
  try {
    const fetchedOrder = await orderService.fetchById(activeShop.value.id, route.params.orderId)
    if (fetchedOrder) {
      order.value = {
        ...fetchedOrder,
        id: fetchedOrder.id || fetchedOrder._id,
      }
    } else {
      order.value = null
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
 * @returns {string} Formatted order ID.
 */
function formatOrderId(id) {
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
    'paid': 'bg-blue-100 text-blue-800',
    'shipped': 'bg-purple-100 text-purple-800',
    'delivered': 'bg-green-100 text-green-800',
    'cancelled': 'bg-red-100 text-red-800'
  }
  return statusMap[status?.toLowerCase()] || 'bg-gray-100 text-gray-800'
}

// Initial data fetch on component mount
onMounted(() => {
  fetchOrder()
})
</script> 