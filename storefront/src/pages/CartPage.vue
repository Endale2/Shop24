<template>
  <div class="max-w-7xl mx-auto p-4 sm:p-6 lg:p-8">
    <div class="mb-8">
      <h1 class="text-3xl font-bold tracking-wide uppercase text-gray-900">Your Cart</h1>
      <p class="text-gray-600 mt-2">Review your items and proceed to checkout</p>
    </div>

    <!-- Loading State -->
    <div v-if="cartStore.loading" class="flex items-center justify-center py-20">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-black"></div>
      <span class="ml-3 text-gray-600">Loading your cart...</span>
    </div>

    <!-- Error State -->
    <div v-else-if="cartStore.error" class="bg-red-50 border border-red-200 rounded-lg p-4 mb-6">
      <div class="flex items-center">
        <div class="flex-shrink-0">
          <svg class="h-5 w-5 text-red-400" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
          </svg>
        </div>
        <div class="ml-3">
          <p class="text-sm text-red-800">{{ cartStore.error }}</p>
          <button 
            @click="retryFetchCart" 
            class="mt-2 text-sm text-red-600 hover:text-red-800 underline"
          >
            Try again
          </button>
        </div>
      </div>
    </div>

    <!-- Cart Content -->
    <div v-else-if="cartStore.cart && cartStore.cart.items && cartStore.cart.items.length > 0" class="grid grid-cols-1 lg:grid-cols-3 gap-8">
      <!-- Cart Items -->
      <div class="lg:col-span-2 space-y-6">
        <!-- Discount Information Banner -->
        <div v-if="hasItemDiscounts || hasOrderDiscounts" class="bg-blue-50 border border-blue-200 rounded-lg p-4">
          <div class="flex items-start">
            <svg class="h-5 w-5 text-blue-600 mt-0.5 mr-3 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
  <div>
              <h3 class="text-sm font-medium text-blue-800 mb-1">Discount Information</h3>
              <div class="text-sm text-blue-700 space-y-1">
                <p v-if="hasItemDiscounts">
                  • <strong>Automatic Item Discounts:</strong> Some items have automatic discounts applied (shown with green badges)
                </p>
                <p v-if="hasOrderDiscounts">
                  • <strong>Order-Level Discounts:</strong> Coupon codes apply to your entire order subtotal
                </p>
                <p class="text-xs text-blue-600 mt-2">
                  💡 <strong>Tip:</strong> Coupon codes work in addition to automatic item discounts for maximum savings!
                </p>
              </div>
            </div>
          </div>
        </div>

        <div class="bg-white border border-gray-200 rounded-lg overflow-hidden">
          <div class="px-6 py-4 border-b border-gray-200">
            <h2 class="text-lg font-semibold text-gray-900">Cart Items ({{ cartStore.cart.items.length }})</h2>
          </div>
          
          <div class="divide-y divide-gray-200">
            <div v-for="item in cartStore.cart.items" :key="`${item.product_id}-${item.variant_id || 'no-variant'}`" class="p-6">
              <div class="flex items-start space-x-4">
                <!-- Product Image -->
                <div class="flex-shrink-0">
                  <img 
                    :src="item.image || '/placeholder-product.jpg'" 
                    :alt="item.product_name"
                    class="w-20 h-20 object-cover rounded-lg border border-gray-200"
                    @error="handleImageError"
                  />
                </div>
                
                <!-- Product Details -->
                <div class="flex-1 min-w-0">
                  <div class="flex justify-between items-start">
      <div>
                      <div class="flex items-center gap-2">
                        <h3 class="text-lg font-medium text-gray-900">{{ item.product_name || 'Product' }}</h3>
                        <!-- Discount indicator -->
                        <div v-if="item.discount_amount > 0" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-green-100 text-green-800">
                          <svg class="w-3 h-3 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                          </svg>
                          Auto Discount
                        </div>
                      </div>
                      <p v-if="item.variant_options && Object.keys(item.variant_options).length > 0" class="text-sm text-gray-500 mt-1">
                        {{ Object.entries(item.variant_options).map(([key, value]) => `${key}: ${value}`).join(', ') }}
                      </p>
                      
                      <!-- Applied Discounts for this item -->
                      <div v-if="getItemDiscountDetails(item).length > 0" class="mt-2 space-y-1">
                        <div v-for="discount in getItemDiscountDetails(item)" :key="discount.discount_id" class="flex items-center text-sm">
                          <span class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-green-100 text-green-800 mr-2">
                            <svg class="w-3 h-3 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                            </svg>
                            {{ discount.name }}
                          </span>
                          <span class="text-green-600 font-medium">
                            {{ formatDiscountValue(discount) }}
                          </span>
                        </div>
                      </div>
                    </div>
                    <button 
                      @click="removeItem(item)"
                      class="text-gray-400 hover:text-red-500 transition-colors"
                      :disabled="cartStore.loading"
                    >
                      <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                      </svg>
                    </button>
                  </div>
                  
                  <div class="mt-4 flex items-center justify-between">
                    <div class="flex items-center space-x-4">
                      <!-- Quantity Controls -->
                      <div class="flex items-center border border-gray-300 rounded-lg">
                        <button
                          @click="updateItemQuantity(item, item.quantity - 1)"
                          :disabled="item.quantity <= 1 || cartStore.loading"
                          class="px-3 py-1 text-gray-600 hover:text-black disabled:text-gray-300 disabled:cursor-not-allowed"
                        >
                          -
                        </button>
                        <span class="px-3 py-1 text-gray-900 font-medium">{{ item.quantity }}</span>
                        <button
                          @click="updateItemQuantity(item, item.quantity + 1)"
                          :disabled="cartStore.loading"
                          class="px-3 py-1 text-gray-600 hover:text-black disabled:text-gray-300 disabled:cursor-not-allowed"
                        >
                          +
                        </button>
                      </div>
                      
                      <div class="text-right">
                        <p class="text-sm text-gray-500">Unit Price</p>
                        <p class="text-lg font-semibold text-gray-900">${{ item.unit_price?.toFixed(2) || '0.00' }}</p>
                      </div>
                    </div>
                    
                    <div class="text-right">
                      <p class="text-sm text-gray-500">Total</p>
                      <div class="space-y-1">
                        <p v-if="item.discount_amount > 0" class="text-sm text-gray-400 line-through">
                          ${{ item.line_total?.toFixed(2) || '0.00' }}
                        </p>
                        <p class="text-lg font-bold text-gray-900">
                          ${{ item.final_line_total?.toFixed(2) || item.line_total?.toFixed(2) || '0.00' }}
                        </p>
                        <p v-if="item.discount_amount > 0" class="text-sm text-green-600">
                          Save ${{ item.discount_amount.toFixed(2) }}
                        </p>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Order Summary -->
      <div class="lg:col-span-1">
        <div class="bg-white border border-gray-200 rounded-lg p-6 sticky top-8">
          <h2 class="text-lg font-semibold text-gray-900 mb-4">Order Summary</h2>
          
          <!-- Price Breakdown -->
          <div class="space-y-3 border-t border-gray-200 pt-4">
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">Subtotal</span>
              <span class="font-medium">${{ cartStore.cart.subtotal?.toFixed(2) || '0.00' }}</span>
            </div>
            
            <!-- Item-level discounts -->
            <div v-if="hasItemDiscounts" class="flex justify-between text-sm">
              <span class="text-gray-600">Item Discounts</span>
              <span class="font-medium text-green-600">-${{ itemDiscountsTotal.toFixed(2) }}</span>
            </div>
            
            <!-- Order-level discounts -->
            <div v-if="cartStore.cart.order_discount_details && cartStore.cart.order_discount_details.length > 0" class="flex justify-between text-sm">
              <span class="text-gray-600">Order Discounts</span>
              <span class="font-medium text-blue-600">-${{ orderDiscountsTotal.toFixed(2) }}</span>
            </div>
            
            <!-- Order Discount Details -->
            <div v-if="cartStore.cart?.order_discount_details && cartStore.cart.order_discount_details.length > 0" class="bg-blue-50 border border-blue-200 rounded-lg p-3 mb-3">
              <h4 class="text-sm font-medium text-blue-800 mb-2">Order-Level Discounts Applied:</h4>
              <div class="space-y-2">
                <div v-for="discount in cartStore.cart.order_discount_details" :key="discount.discount_id" class="flex items-center justify-between">
                  <div class="flex items-center">
                    <svg class="h-4 w-4 text-blue-600 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    <span class="text-sm text-blue-700">{{ discount.name }}</span>
                  </div>
                  <span class="text-sm font-medium text-blue-600">-${{ discount.amount.toFixed(2) }}</span>
                </div>
              </div>
              <p class="text-xs text-blue-600 mt-2">
                <svg class="w-3 h-3 inline mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                Applied to entire order subtotal
              </p>
            </div>
            
            <!-- Total discounts (if either type exists) -->
            <div v-if="hasItemDiscounts || (cartStore.cart.order_discount_details && cartStore.cart.order_discount_details.length > 0)" class="flex justify-between text-sm border-t border-gray-100 pt-2">
              <span class="text-gray-600">Total Discounts</span>
              <span class="font-medium text-green-600">-${{ cartStore.cart.total_discounts?.toFixed(2) || '0.00' }}</span>
            </div>
            
            <div v-if="cartStore.cart.shipping_cost > 0" class="flex justify-between text-sm">
              <span class="text-gray-600">Shipping</span>
              <span class="font-medium">${{ cartStore.cart.shipping_cost.toFixed(2) }}</span>
            </div>
            
            <div v-if="cartStore.cart.tax_amount > 0" class="flex justify-between text-sm">
              <span class="text-gray-600">Tax</span>
              <span class="font-medium">${{ cartStore.cart.tax_amount.toFixed(2) }}</span>
            </div>
            
            <div class="flex justify-between text-lg font-bold border-t border-gray-200 pt-3">
              <span>Total</span>
              <span>${{ cartStore.cart.grand_total?.toFixed(2) || cartStore.cart.subtotal?.toFixed(2) || '0.00' }}</span>
            </div>
            
            <!-- Savings summary -->
            <div v-if="cartStore.cart.total_discounts > 0" class="bg-green-50 border border-green-200 rounded-lg p-3 mt-4">
              <div class="flex items-center">
                <svg class="h-5 w-5 text-green-600 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <span class="text-sm font-medium text-green-800">
                  You save ${{ cartStore.cart.total_discounts.toFixed(2) }}
                </span>
              </div>
            </div>
          </div>

          <!-- Action Buttons -->
          <div class="mt-6 space-y-3">
            <button
              @click="checkout"
              :disabled="cartStore.loading || checkoutLoading"
              class="w-full py-3 bg-black text-white rounded-md font-bold hover:bg-gray-800 disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center gap-2"
            >
              <span v-if="checkoutLoading" class="animate-spin rounded-full h-5 w-5 border-b-2 border-white"></span>
              {{ checkoutLoading ? 'Processing...' : 'Proceed to Checkout' }}
            </button>
            
            <button
              @click="clearCart"
              :disabled="cartStore.loading"
              class="w-full py-2 text-gray-600 border border-gray-300 rounded-md font-medium hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
            >
              Clear Cart
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Empty Cart -->
    <div v-else class="text-center py-20">
      <div class="max-w-md mx-auto">
        <svg class="mx-auto h-24 w-24 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-2.5 5M7 13l2.5 5m6-5v6a2 2 0 01-2 2H9a2 2 0 01-2-2v-6m8 0V9a2 2 0 00-2-2H9a2 2 0 00-2 2v4.01" />
        </svg>
        <h3 class="mt-4 text-lg font-medium text-gray-900">Your cart is empty</h3>
        <p class="mt-2 text-gray-500">Start shopping to add items to your cart.</p>
        <div class="mt-6">
          <router-link
            :to="{ path: `/shops/${shopSlug}` }"
            class="inline-flex items-center px-4 py-2 bg-black text-white rounded-md font-medium hover:bg-gray-800"
          >
            Continue Shopping
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useCartStore } from '@/stores/cart';
import { useRouter, useRoute } from 'vue-router';
import { placeOrder } from '@/services/order';
import type { CartItem, ItemDiscountDetail, OrderDiscountDetail } from '@/services/cart';

const cartStore = useCartStore();
const router = useRouter();
const route = useRoute();
const shopSlug = route.params.shopSlug as string;

const checkoutLoading = ref(false);

onMounted(() => {
  cartStore.setShopSlug(shopSlug);
  cartStore.fetchCart();
});

// Computed properties for discount display
const hasItemDiscounts = computed(() => {
  if (!cartStore.cart?.items) return false;
  return cartStore.cart.items.some((item: CartItem) => item.discount_amount > 0);
});

const hasOrderDiscounts = computed(() => {
  if (!cartStore.cart?.total_discounts || !cartStore.cart?.items) return false;
  const itemDiscountsTotal = cartStore.cart.items.reduce((total: number, item: CartItem) => total + (item.discount_amount || 0), 0);
  return cartStore.cart.total_discounts > itemDiscountsTotal;
});

const itemDiscountsTotal = computed(() => {
  if (!cartStore.cart?.items) return 0;
  return cartStore.cart.items.reduce((total: number, item: CartItem) => total + (item.discount_amount || 0), 0);
});

const orderDiscountsTotal = computed(() => {
  if (!cartStore.cart?.total_discounts || !cartStore.cart?.items) return 0;
  const itemDiscountsTotal = cartStore.cart.items.reduce((total: number, item: CartItem) => total + (item.discount_amount || 0), 0);
  return cartStore.cart.total_discounts - itemDiscountsTotal;
});

// Helper function to get discount details for a specific item
function getItemDiscountDetails(item: CartItem): ItemDiscountDetail[] {
  if (!cartStore.cart?.item_discount_details) return [];
  return cartStore.cart.item_discount_details.filter((discount: ItemDiscountDetail) => 
    discount.product_id === item.product_id && 
    discount.variant_id === (item.variant_id || '')
  );
}

// Helper function to format discount value for display
function formatDiscountValue(discount: ItemDiscountDetail | OrderDiscountDetail): string {
  if (discount.type === 'percentage') {
    return `${discount.value}% off`;
  } else {
    return `$${discount.value.toFixed(2)} off`;
  }
}

function retryFetchCart() {
  cartStore.error = null;
  cartStore.fetchCart();
}

function handleImageError(event: Event) {
  const img = event.target as HTMLImageElement;
  img.src = '/placeholder-product.jpg';
}

function updateItemQuantity(item: CartItem, newQuantity: number) {
  if (newQuantity < 1) return;
  cartStore.error = null;
  cartStore.updateCartItem(item.product_id, item.variant_id || '', newQuantity);
}

function removeItem(item: CartItem) {
  cartStore.error = null;
  cartStore.removeCartItem(item.product_id, item.variant_id || '');
}

async function checkout() {
  checkoutLoading.value = true;
  cartStore.error = null;
  try {
    // Place the order
    const order = await cartStore.placeOrder();
    
    // Clear the cart after successful order placement
    await cartStore.clearCart();
    
    // Redirect to order confirmation page
    router.push({
      name: 'OrderConfirmation',
      params: { 
        shopSlug: shopSlug,
        orderId: order.order?.id || order.order?._id || order.order?.ID
      }
    });
  } catch (error: any) {
    console.error('Checkout failed:', error);
    // Show error message, do not clear the cart
    cartStore.error = error?.response?.data?.error || error?.message || 'Checkout failed. Please try again.';
  } finally {
    checkoutLoading.value = false;
  }
}

function clearCart() {
  if (confirm('Are you sure you want to clear your cart?')) {
    cartStore.error = null;
  cartStore.clearCart();
  }
}
</script> 