<template>
  <Breadcrumbs :items="[
    { back: true },
    { label: 'Home', to: `/${shopSlug}/` },
    { label: 'My Orders' }
  ]" />
  <div class="orders-container">
    <div class="max-w-6xl mx-auto px-2 sm:px-4">
      <!-- Header Section -->
      <div class="mb-8 text-center">
        <h1 class="text-3xl font-bold text-gray-900 tracking-tight uppercase mb-2">My Orders</h1>
        <p class="text-gray-600">Track your order history and status</p>
      </div>

      <LoginPrompt
        v-if="!authStore.user"
        title="Sign in to view your orders"
        message="Track your order history and status by signing into your account."
      >
        <template #icon>
          <svg class="w-24 h-24 text-gray-300 mx-auto mb-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
        </template>
      </LoginPrompt>

      <!-- Loading State -->
      <div v-else-if="loading" class="text-center py-12">
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
        <button @click="fetchOrders" class="bg-black text-white py-2 px-4 rounded-md font-bold uppercase tracking-wide hover:bg-gray-800 transition-colors">
          Try Again
        </button>
      </div>

      <!-- Orders List -->
      <div v-else-if="orders.length > 0" class="grid gap-6 lg:grid-cols-2">
        <div v-for="order in orders" :key="order._id" class="bg-white border border-gray-200 rounded-lg p-6 shadow-sm hover:shadow-lg transition-shadow duration-200 flex flex-col h-full">
          <!-- Order Header -->
          <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-4 gap-2">
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
          <div class="divide-y divide-gray-100 mb-4">
            <div v-for="item in order.items" :key="`${item.product_id}-${item.variant_id}`" class="flex items-center gap-3 py-3">
              <div class="flex-shrink-0 w-14 h-14 bg-gray-100 rounded-lg flex items-center justify-center overflow-hidden">
                <img v-if="item.image" :src="item.image" alt="Product image" class="w-full h-full object-contain" />
                <PhotoIcon v-else class="w-8 h-8 text-gray-400" />
              </div>
              <div class="flex-1 min-w-0">
                <p class="text-sm font-medium text-gray-900 truncate">{{ item.name }}</p>
                <p v-if="item.variant_name" class="text-xs text-gray-500">{{ item.variant_name }}</p>
                <p class="text-xs text-gray-500">Qty: {{ item.quantity }}</p>
              </div>
              <div class="text-sm font-bold text-gray-900 text-right min-w-[60px]">
                ${{ item.total_price }}
              </div>
            </div>
          </div>

          <!-- Order Summary -->
          <div class="mt-auto pt-4 border-t border-gray-100">
            <div class="flex justify-between text-sm mb-1">
              <span class="text-gray-600">Subtotal:</span>
              <span class="font-medium">${{ order.subtotal }}</span>
            </div>
            <div v-if="order.discount_total" class="flex justify-between text-sm mb-1">
              <span class="text-gray-600">Discount:</span>
              <span class="font-medium text-green-600">-${{ order.discount_total }}</span>
            </div>
            <div v-if="order.shipping" class="flex justify-between text-sm mb-1">
              <span class="text-gray-600">Shipping:</span>
              <span class="font-medium">${{ order.shipping }}</span>
            </div>
            <div class="flex justify-between text-base font-bold mt-2 pt-2 border-t border-gray-200">
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
          :to="`/${shopSlug}/products`" 
          class="inline-block bg-black text-white py-3 px-6 rounded-md font-bold uppercase tracking-wide hover:bg-gray-800 transition-colors"
        >
          Browse Products
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import { useRoute } from 'vue-router';
import { useAuthStore } from '../stores/auth';
import { getCustomerOrders } from '@/services/order';
import { PhotoIcon } from '@heroicons/vue/24/outline'
import LoginPrompt from '@/components/LoginPrompt.vue';
import Breadcrumbs from '@/components/Breadcrumbs.vue';

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
        total_price: item.total_price,
        image: item.image // Assuming item.image is available in the backend response
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

// Only fetch orders when user is authenticated and session loading is complete
function fetchOrdersIfAuthenticated() {
  if (!shopSlug) return;
  if (!authStore.sessionLoading && authStore.user) {
    fetchOrders();
  }
}

onMounted(() => {
  fetchOrdersIfAuthenticated();
});

// Watch for authentication state changes
watch(
  () => [authStore.sessionLoading, authStore.user],
  ([sessionLoading, user]) => {
    if (!sessionLoading && user) {
      fetchOrdersIfAuthenticated();
    }
  }
);
</script>

<style scoped>
.orders-container {
  min-height: 80vh;
  padding: 2rem 0;
  background: #f7f7f9;
}
</style> 