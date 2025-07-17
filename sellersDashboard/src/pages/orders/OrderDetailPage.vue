<template>
  <div class="bg-gray-50 min-h-screen font-sans">
    <div class="p-4 sm:p-6 lg:p-8 max-w-7xl mx-auto">
      <header class="flex flex-col sm:flex-row justify-between items-start sm:items-center mb-8">
        <div>
          <h1 class="text-4xl font-bold text-gray-800 tracking-tight">
            Order Details
          </h1>
          <p v-if="order" class="text-md text-gray-500 mt-2">
            <span class="font-mono">{{ formatOrderId(order.id, order.orderNumber) }}</span> &bull; {{ formatDate(order.createdAt) }}
          </p>
        </div>
        
        <div class="flex items-center space-x-3 mt-4 sm:mt-0">
          <button
            @click="goBack"
            class="inline-flex items-center px-4 py-2 bg-white text-gray-700 text-sm font-medium rounded-lg border border-gray-300 shadow-sm hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 transition-all duration-150 ease-in-out"
          >
            <ArrowLeftIcon class="w-5 h-5 mr-2 -ml-1 text-gray-500" />
            Back
          </button>
          
          <button
            v-if="order"
            @click="updateOrderStatus"
            class="inline-flex items-center px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition-all duration-150 ease-in-out"
          >
            <PencilIcon class="w-5 h-5 mr-2 -ml-1" />
            Update Status
          </button>
        </div>
      </header>

      <div v-if="loading" class="space-y-6">
        <div class="bg-white rounded-xl border border-gray-200 p-6 space-y-4 animate-pulse">
            <div class="h-6 bg-gray-200 rounded w-1/3"></div>
            <div class="h-4 bg-gray-200 rounded w-1/4"></div>
        </div>
        <div class="bg-white rounded-xl border border-gray-200 p-6 space-y-4 animate-pulse">
            <div v-for="n in 3" :key="n" class="h-16 bg-gray-200 rounded"></div>
        </div>
      </div>

      <div v-else-if="error" class="bg-red-50 border-l-4 border-red-400 text-red-700 p-6 rounded-r-lg shadow-md" role="alert">
        <div class="flex">
          <div class="py-1"><svg class="h-6 w-6 text-red-500 mr-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg></div>
          <div>
            <p class="font-bold">An error occurred</p>
            <p class="text-sm">{{ error }}</p>
            <button @click="fetchOrder" class="text-sm underline mt-2 font-medium">Please try again</button>
          </div>
        </div>
      </div>

      <main v-else-if="order" class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        
        <div class="lg:col-span-2 space-y-8">
          <!-- Order Items Section -->
          <div class="bg-white rounded-xl border border-gray-200 overflow-hidden shadow-sm">
            <div class="p-6">
              <h3 class="text-xl font-semibold text-gray-800 mb-4">Order Items</h3>
              <div class="flow-root">
                <ul role="list" class="-my-4 divide-y divide-gray-200">
                  <li v-for="item in order.items" :key="`${item.productId}-${item.variantId || 'no-variant'}`" class="flex items-center py-4 space-x-4 group hover:bg-green-50 transition cursor-pointer rounded-lg" @click="goToProductDetail(item.productId)">
                    <div class="flex-shrink-0">
                      <img :src="item.image || '/placeholder-product.png'" :alt="item.name" class="h-16 w-16 rounded-lg object-cover border border-gray-200 group-hover:ring-2 group-hover:ring-green-400 transition" @error="$event.target.src='/placeholder-product.png'" />
                    </div>
                    <div class="flex-1 min-w-0">
                      <p class="text-md font-semibold text-gray-900 truncate">{{ item.name }}</p>
                      <p v-if="item.variantId && item.variantId !== '000000000000000000000000'" class="text-sm text-gray-500">{{ item.variantName || 'Variant' }}</p>
                      <p v-else class="text-sm text-gray-400 italic">No variant</p>
                    </div>
                    <div class="text-right">
                      <p class="text-md font-medium text-gray-900">{{ formatPrice(item.totalPrice) }}</p>
                      <p class="text-sm text-gray-500">{{ item.quantity }} x {{ formatPrice(item.unitPrice) }}</p>
                    </div>
                    <span class="ml-2 text-gray-400 group-hover:text-green-500 transition" title="View Product">
                      <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" /></svg>
                    </span>
                  </li>
                </ul>
              </div>
            </div>
          </div>
          <!-- Customer Card Section -->
          <div v-if="customer" class="bg-white rounded-xl border border-gray-200 p-6 shadow-sm hover:shadow-lg transition cursor-pointer group" @click="goToCustomerProfile(customer.id)">
            <h3 class="text-xl font-semibold text-gray-800 mb-4 flex items-center">
              Customer
              <span class="ml-2 text-gray-400 group-hover:text-green-500 transition" title="View Customer Profile">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" /></svg>
              </span>
            </h3>
            <div class="flex items-start space-x-4">
              <template v-if="customer.profile_image">
                <img :src="customer.profile_image" :alt="`${customer.firstName} ${customer.lastName}`" class="h-16 w-16 rounded-full object-cover border border-gray-200 group-hover:ring-2 group-hover:ring-green-400 transition" @error="$event.target.src='/placeholder-avatar.png'" />
              </template>
              <template v-else>
                <div class="h-16 w-16 rounded-full bg-gradient-to-br from-green-400 to-teal-500 flex items-center justify-center text-white text-xl font-bold border border-gray-200 group-hover:ring-2 group-hover:ring-green-400 transition">
                  {{ getInitials(customer) }}
                </div>
              </template>
              <div class="flex-1 min-w-0">
                <p class="text-lg font-semibold text-gray-900">{{ customer.firstName }} {{ customer.lastName }}</p>
                <p class="text-sm text-gray-500 truncate">{{ customer.email }}</p>
                <p v-if="customer.phone" class="text-sm text-gray-500">{{ customer.phone }}</p>
              </div>
            </div>
            <div v-if="customer.address" class="mt-4 pt-4 border-t border-gray-200">
              <p class="text-sm font-medium text-gray-600">Shipping Address</p>
              <address class="text-sm text-gray-500 not-italic">
                {{ customer.address }}<br>
                <span v-if="customer.city">{{ customer.city }}, </span>
                <span v-if="customer.state">{{ customer.state }} </span>
                <span v-if="customer.postalCode">{{ customer.postalCode }}</span>
              </address>
            </div>
          </div>
        </div>
        
        <div class="lg:col-span-1 space-y-8">
            <div class="bg-white rounded-xl border border-gray-200 p-6 shadow-sm">
                <h3 class="text-xl font-semibold text-gray-800 mb-4">Summary</h3>
                <div class="space-y-4">
                    <div class="flex justify-between items-center">
                        <span class="text-sm font-medium text-gray-500">Status</span>
                        <span :class="getStatusClass(order.status)" class="inline-block px-3 py-1 text-sm font-semibold rounded-full">
                            {{ formatStatus(order.status) }}
                        </span>
                    </div>
                    <div class="flex justify-between items-center">
                        <span class="text-sm text-gray-500">Subtotal</span>
                        <span class="text-sm font-medium text-gray-900">{{ formatPrice(order.subtotal) }}</span>
                    </div>
                    <div v-if="order.discountTotal > 0" class="flex justify-between items-center">
                        <span class="text-sm text-gray-500">Discount</span>
                        <span class="text-sm font-medium text-green-600">-{{ formatPrice(order.discountTotal) }}</span>
                    </div>
                    <div class="flex justify-between items-baseline pt-4 border-t border-gray-200">
                        <span class="text-lg font-bold text-gray-900">Total</span>
                        <span class="text-2xl font-bold text-green-600">{{ formatPrice(order.total) }}</span>
                    </div>
                </div>
            </div>

            <div class="bg-white rounded-xl border border-gray-200 p-6 shadow-sm">
                <h3 class="text-xl font-semibold text-gray-800 mb-4">Additional Info</h3>
                <dl class="space-y-3">
                    <div class="flex justify-between">
                        <dt class="text-sm text-gray-500">Order #</dt>
                        <dd class="text-sm font-mono text-gray-900">{{ formatOrderId(order.id, order.orderNumber) }}</dd>
                    </div>
                     <div class="flex justify-between">
                        <dt class="text-sm text-gray-500">Customer ID</dt>
                        <dd class="text-sm font-mono text-gray-900">{{ customer?.id || '—' }}</dd>
                    </div>
                    <div v-if="order.couponCode" class="flex justify-between">
                        <dt class="text-sm text-gray-500">Coupon</dt>
                        <dd class="text-sm font-mono text-gray-900">{{ order.couponCode }}</dd>
                    </div>
                     <div class="flex justify-between">
                        <dt class="text-sm text-gray-500">Item Count</dt>
                        <dd class="text-sm text-gray-900">{{ order.items?.length || 0 }}</dd>
                    </div>
                </dl>
            </div>
        </div>
      </main>

      <div v-else class="text-center py-20">
        <div class="mx-auto h-12 w-12 text-gray-400">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-full w-full" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
            </svg>
        </div>
        <h3 class="mt-2 text-lg font-medium text-gray-900">Order Not Found</h3>
        <p class="mt-1 text-sm text-gray-500">We couldn't find the order you're looking for.</p>
        <div class="mt-6">
            <button @click="goBack" class="text-sm font-medium text-indigo-600 hover:text-indigo-500">Go back to orders</button>
        </div>
      </div>
    </div>

    <div v-if="showStatusModal" class="fixed inset-0 z-50 flex items-center justify-center bg-white/30 backdrop-blur-none transition-opacity" @click.self="showStatusModal = false">
      <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-md p-8 m-4 border border-green-100">
        <h3 class="text-xl font-bold text-gray-900 mb-4 flex items-center gap-2">
          <PencilIcon class="w-5 h-5 text-green-500" />
          Update Order Status
        </h3>
        <p class="text-sm text-gray-500 mb-6">Select a new status for this order. The customer may be notified of this change.</p>
        <div class="mb-8">
          <label for="order-status" class="block text-sm font-medium text-gray-700 mb-2">Order Status</label>
          <select
            id="order-status"
            v-model="newStatus"
            class="w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500 bg-gray-50 text-gray-900 transition-colors text-base"
          >
            <option value="pending">Pending</option>
            <option value="paid">Paid</option>
            <option value="shipped">Shipped</option>
            <option value="delivered">Delivered</option>
            <option value="cancelled">Cancelled</option>
          </select>
        </div>
        <div class="border-t border-gray-100 pt-6 flex justify-end space-x-3">
          <button
            @click="showStatusModal = false"
            class="px-4 py-2 text-gray-700 bg-gray-100 rounded-lg hover:bg-gray-200 transition-colors border border-gray-200 focus:outline-none focus:ring-2 focus:ring-green-400"
          >
            Cancel
          </button>
          <button
            @click="saveStatusUpdate"
            :disabled="updatingStatus"
            class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 disabled:bg-green-300 disabled:cursor-not-allowed transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 font-semibold shadow-sm"
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
  CheckIcon
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