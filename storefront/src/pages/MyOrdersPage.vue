<template>
  <div class="orders-container">
    <div class="max-w-6xl mx-auto">
      <!-- Header Section -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 tracking-tight uppercase mb-2">My Orders</h1>
        <p class="text-gray-600">Track your order history and status</p>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-12">
        <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-black"></div>
        <p class="mt-4 text-gray-600">Loading your orders...</p>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="text-center py-12">
        <svg class="w-16 h-16 text-red-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
        </svg>
        <h3 class="text-lg font-semibold text-gray-900 mb-2">Unable to load orders</h3>
        <p class="text-gray-600 mb-4">{{ error }}</p>
        <button @click="fetchOrders" class="bg-black text-white py-2 px-4 rounded-none font-semibold uppercase tracking-wide hover:bg-gray-800 transition-colors">
          Try Again
        </button>
      </div>

      <!-- Orders List -->
      <div v-else-if="orders.length > 0" class="space-y-6">
        <div v-for="order in orders" :key="order._id" class="bg-white border border-gray-200 rounded-none p-6">
          <!-- Order Header -->
          <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-4">
            <div>
              <h3 class="text-lg font-semibold text-gray-900">Order #{{ order.order_number || order._id.slice(-8) }}</h3>
              <p class="text-sm text-gray-600">{{ formatDate(order.createdAt) }}</p>
            </div>
            <div class="mt-2 sm:mt-0">
              <span :class="getStatusClass(order.status)" class="px-3 py-1 text-sm font-semibold uppercase tracking-wide">
                {{ order.status }}
              </span>
            </div>
          </div>

          <!-- Order Items -->
          <div class="space-y-4">
            <div v-for="item in order.items" :key="`${item.product_id}-${item.variant_id}`" class="flex items-center space-x-4">
              <div class="flex-shrink-0 w-16 h-16 bg-gray-200 rounded-lg flex items-center justify-center">
                <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                </svg>
              </div>
              <div class="flex-1 min-w-0">
                <p class="text-sm font-medium text-gray-900 truncate">{{ item.name }}</p>
                <p v-if="item.variant_name" class="text-sm text-gray-500">{{ item.variant_name }}</p>
                <p class="text-sm text-gray-500">Qty: {{ item.quantity }}</p>
              </div>
              <div class="text-sm font-medium text-gray-900">
                ${{ item.total_price }}
              </div>
            </div>
          </div>

          <!-- Order Summary -->
          <div class="mt-6 pt-4 border-t border-gray-200">
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">Subtotal:</span>
              <span class="font-medium">${{ order.subtotal }}</span>
            </div>
            <div v-if="order.discount_total" class="flex justify-between text-sm">
              <span class="text-gray-600">Discount:</span>
              <span class="font-medium text-green-600">-${{ order.discount_total }}</span>
            </div>
            <div v-if="order.shipping" class="flex justify-between text-sm">
              <span class="text-gray-600">Shipping:</span>
              <span class="font-medium">${{ order.shipping }}</span>
            </div>
            <div class="flex justify-between text-base font-semibold mt-2 pt-2 border-t border-gray-200">
              <span>Total:</span>
              <span>${{ order.total }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-else class="text-center py-12">
        <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
        </svg>
        <h3 class="text-lg font-semibold text-gray-900 mb-2">No orders yet</h3>
        <p class="text-gray-600 mb-6">Start shopping to see your orders here</p>
        <router-link 
          :to="`/shops/${route.params.shopSlug}/products`" 
          class="inline-block bg-black text-white py-3 px-6 rounded-none font-semibold uppercase tracking-wide hover:bg-gray-800 transition-colors"
        >
          Browse Products
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { useAuthStore } from '../stores/auth';
import { getCustomerOrders } from '@/services/order';

const route = useRoute();
const authStore = useAuthStore();
const shopSlug = route.params.shopSlug as string;

const orders = ref<any[]>([]);
const loading = ref(true);
const error = ref('');

function formatDate(dateString: string) {
  if (!dateString) return 'Unknown';
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  });
}

function getStatusClass(status: string) {
  const statusClasses = {
    'pending': 'bg-yellow-100 text-yellow-800',
    'processing': 'bg-blue-100 text-blue-800',
    'shipped': 'bg-purple-100 text-purple-800',
    'delivered': 'bg-green-100 text-green-800',
    'cancelled': 'bg-red-100 text-red-800',
    'refunded': 'bg-gray-100 text-gray-800'
  };
  return statusClasses[status as keyof typeof statusClasses] || 'bg-gray-100 text-gray-800';
}

async function fetchOrders() {
  loading.value = true;
  error.value = '';
  try {
    const response = await getCustomerOrders(shopSlug);
    // Map backend fields to frontend display
    orders.value = (response.data || []).map((order: any) => ({
      _id: order._id || order.id,
      order_number: order.order_number,
      status: order.status,
      createdAt: order.created_at || order.createdAt,
      items: (order.items || []).map((item: any) => ({
        product_id: item.product_id,
        variant_id: item.variant_id,
        name: item.name,
        variant_name: item.variant_name || '',
        quantity: item.quantity,
        total_price: item.total_price
      })),
      subtotal: order.subtotal,
      discount_total: order.discount_total,
      shipping: order.shipping_cost || order.shipping,
      total: order.total
    }));
  } catch (err: any) {
    error.value = err.response?.data?.error || 'Failed to load orders';
  } finally {
    loading.value = false;
  }
}

onMounted(() => {
  fetchOrders();
});
</script>

<style scoped>
.orders-container {
  min-height: 80vh;
  padding: 2rem 0;
  background: #f7f7f9;
}
</style> 