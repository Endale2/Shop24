<template>
  <!-- Breadcrumb and Back Button -->
  <nav class="flex items-center space-x-2 text-sm text-gray-500 mb-6">
    <button @click="$router.back()" class="text-gray-700 hover:text-black flex items-center gap-1 font-medium text-xs mr-2">
      <svg class="w-4 h-4" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7"/></svg>
      Back
    </button>
    <router-link :to="`/shops/${shopSlug}`" class="hover:underline">Home</router-link>
    <span>/</span>
    <span>Order Confirmation</span>
  </nav>
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- Header -->
      <div class="text-center mb-8">
        <div class="inline-flex items-center justify-center w-16 h-16 bg-green-100 rounded-full mb-4">
          <svg class="w-8 h-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
          </svg>
        </div>
        <h1 class="text-3xl font-bold text-gray-900 mb-2">Order Confirmed!</h1>
        <p class="text-lg text-gray-600">Thank you for your purchase. Your order has been successfully placed.</p>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-gray-900 mx-auto mb-4"></div>
        <p class="text-gray-600">Loading order details...</p>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="bg-red-50 border border-red-200 rounded-lg p-6 text-center">
        <svg class="w-12 h-12 text-red-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
        </svg>
        <h3 class="text-lg font-medium text-red-800 mb-2">Error Loading Order</h3>
        <p class="text-red-600">{{ error }}</p>
      </div>

      <!-- Order Details -->
      <div v-else-if="order" class="space-y-6">
        <!-- Order Summary Card -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <div class="flex items-center justify-between mb-4">
            <h2 class="text-xl font-semibold text-gray-900">Order Summary</h2>
            <span class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-green-100 text-green-800">
              {{ order.status }}
            </span>
          </div>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
            <div>
              <p class="text-sm text-gray-500">Order Number</p>
              <p class="font-medium text-gray-900">{{ order.order_number || order.id || order._id }}</p>
            </div>
            <div>
              <p class="text-sm text-gray-500">Order Date</p>
              <p class="font-medium text-gray-900">{{ formatDate(order.created_at) }}</p>
            </div>
          </div>

          <!-- Order Items -->
          <div class="border-t border-gray-200 pt-4">
            <h3 class="text-lg font-medium text-gray-900 mb-4">Items Ordered</h3>
            <div class="space-y-4">
              <div 
                v-for="item in order.items" 
                :key="item.product_id + (item.variant_id || '')"
                class="flex items-center space-x-4"
              >
                <img 
                  :src="item.image || '/placeholder-product.jpg'" 
                  :alt="item.name"
                  class="w-16 h-16 object-cover rounded-lg border border-gray-200"
                  @error="handleImageError"
                />
                <div class="flex-1 min-w-0">
                  <h4 class="text-sm font-medium text-gray-900 truncate">{{ item.name }}</h4>
                  <p class="text-sm text-gray-500">Qty: {{ item.quantity }}</p>
                </div>
                <div class="text-right">
                  <p class="text-sm font-medium text-gray-900">${{ item.total_price.toFixed(2) }}</p>
                  <p class="text-xs text-gray-500">${{ item.unit_price.toFixed(2) }} each</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Price Breakdown Card -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">Price Breakdown</h3>
          <div class="space-y-3">
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">Subtotal</span>
              <span class="font-medium">${{ order.subtotal.toFixed(2) }}</span>
            </div>
            
            <div v-if="order.discount_total > 0" class="flex justify-between text-sm">
              <span class="text-gray-600">Discounts</span>
              <span class="font-medium text-green-600">-${{ order.discount_total.toFixed(2) }}</span>
            </div>
            
            <div v-if="order.shipping_cost > 0" class="flex justify-between text-sm">
              <span class="text-gray-600">Shipping</span>
              <span class="font-medium">${{ order.shipping_cost.toFixed(2) }}</span>
            </div>
            
            <div v-if="order.tax_amount > 0" class="flex justify-between text-sm">
              <span class="text-gray-600">Tax</span>
              <span class="font-medium">${{ order.tax_amount.toFixed(2) }}</span>
            </div>
            
            <div class="flex justify-between text-lg font-bold border-t border-gray-200 pt-3">
              <span>Total</span>
              <span>${{ order.total.toFixed(2) }}</span>
            </div>
          </div>
        </div>

        <!-- Action Buttons -->
        <div class="flex flex-col sm:flex-row gap-4">
          <router-link
            :to="{ path: `/shops/${shopSlug}` }"
            class="flex-1 py-3 px-6 bg-black text-white rounded-md font-medium hover:bg-gray-800 text-center"
          >
            Continue Shopping
          </router-link>
          <router-link
            :to="{ name: 'MyOrders', params: { shopSlug: shopSlug } }"
            class="flex-1 py-3 px-6 border border-gray-300 text-gray-700 rounded-md font-medium hover:bg-gray-50 text-center"
          >
            View My Orders
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { getOrderDetail } from '@/services/order';
import { useCartStore } from '@/stores/cart';

const route = useRoute();
const cartStore = useCartStore();
const order = ref<any>(null);
const loading = ref(true);
const error = ref('');

const shopSlug = route.params.shopSlug as string;

onMounted(async () => {
  try {
    const { data } = await getOrderDetail(shopSlug, route.params.orderId as string);
    order.value = data;
  } catch (e: any) {
    console.error('Failed to load order:', e);
    error.value = e.response?.data?.error || 'Failed to load order details';
  } finally {
    loading.value = false;
  }
});

function formatDate(dateString: string) {
  if (!dateString) return 'N/A';
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  });
}

function handleImageError(event: Event) {
  const img = event.target as HTMLImageElement;
  img.src = '/placeholder-product.jpg';
}
</script> 