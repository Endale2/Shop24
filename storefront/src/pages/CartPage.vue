<template>
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900">Shopping Cart</h1>
        <p class="text-gray-600 mt-2">Review your items and proceed to checkout</p>
      </div>

      <!-- Session Loading State -->
      <div v-if="authStore.sessionLoading" class="flex items-center justify-center min-h-[60vh]">
        <div class="flex flex-col items-center">
          <svg class="animate-spin h-12 w-12 text-gray-400 mb-4" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <p class="text-gray-500">Checking authentication…</p>
        </div>
      </div>

      <!-- Not Logged In State -->
      <LoginPrompt
        v-else-if="!authStore.user"
        title="Sign in to view your cart"
        message="Create an account or sign in to view and manage your shopping cart."
      >
        <template #icon>
          <svg class="w-24 h-24 text-gray-300 mx-auto mb-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
          </svg>
        </template>
      </LoginPrompt>

      <!-- Logged In States -->
      <div v-else>
        <!-- Error Message -->
        <div v-if="cartStore.error" class="mb-6 bg-red-50 border border-red-200 rounded-lg p-4">
          <div class="flex">
            <svg class="h-5 w-5 text-red-400 mt-0.5 mr-3 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <div>
              <h3 class="text-sm font-medium text-red-800">Error</h3>
              <p class="text-sm text-red-700 mt-1">{{ cartStore.error }}</p>
              <button @click="retryFetchCart" class="mt-2 text-sm text-red-600 hover:text-red-500 font-medium">
                Try again
              </button>
            </div>
          </div>
        </div>

        <!-- Loading State -->
        <Loader v-if="cartStore.loading" text="Loading your cart..." />

        <!-- Cart Content -->
        <div v-else-if="cartStore.cart && cartStore.cart.items && cartStore.cart.items.length > 0" class="grid grid-cols-1 lg:grid-cols-3 gap-8">
          <!-- Cart Items -->
          <div class="lg:col-span-2 space-y-6">
            <!-- Security Notice -->
            <div class="bg-blue-50 border border-blue-200 rounded-lg p-4">
              <div class="flex items-start">
                <svg class="h-5 w-5 text-blue-600 mt-0.5 mr-3 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <div>
                  <h3 class="text-sm font-medium text-blue-800 mb-1">Secure Pricing</h3>
                  <div class="text-sm text-blue-700 space-y-1">
                    <p>• All prices and discounts are calculated securely on our servers</p>
                    <p>• Frontend displays are for information only and cannot be manipulated</p>
                    <p>• Final pricing is determined during checkout</p>
                  </div>
                </div>
              </div>
            </div>

            <!-- Discount Information Banner -->
            <div v-if="hasAnyDiscounts" class="bg-green-50 border border-green-200 rounded-lg p-4">
              <div class="flex items-start">
                <svg class="h-5 w-5 text-green-600 mt-0.5 mr-3 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <div>
                  <h3 class="text-sm font-medium text-green-800 mb-1">Discount Information</h3>
                  <div class="text-sm text-green-700 space-y-1">
                    <p v-if="hasItemDiscounts">
                      • <strong>Automatic Item Discounts:</strong> Some items have automatic discounts applied (shown with green badges)
                    </p>
                    <p v-if="hasOrderDiscounts">
                      • <strong>Order-Level Discounts:</strong> Coupon codes apply to your entire order subtotal
                    </p>
                    <p class="text-xs text-green-600 mt-2">
                      💡 <strong>Note:</strong> All discount calculations are performed securely on our servers
                    </p>
                  </div>
                </div>
              </div>
            </div>

            <!-- Discount Limit Information -->
            <div v-if="hasDiscountLimits" class="bg-blue-50 border border-blue-200 rounded-lg p-4">
              <div class="flex items-start">
                <svg class="h-5 w-5 text-blue-600 mt-0.5 mr-3 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <div>
                  <h3 class="text-sm font-medium text-blue-800 mb-2">Discount Usage Information</h3>
                  <div class="space-y-2">
                    <!-- Customer Limit Reached -->
                    <div v-if="customerLimitReachedDiscounts.length > 0" class="bg-yellow-50 border border-yellow-200 rounded-lg p-3">
                      <h4 class="text-sm font-medium text-yellow-800 mb-1">You've reached your limit for:</h4>
                      <div class="space-y-1">
                        <div v-for="status in customerLimitReachedDiscounts" :key="status.discount_id" class="text-sm text-yellow-700">
                          • <strong>{{ status.name }}</strong> - You've used {{ status.customer_usage }}/{{ status.customer_limit }} times
                        </div>
                      </div>
                    </div>
                    
                    <!-- Available Discounts with Limits -->
                    <div v-if="availableDiscountsWithLimits.length > 0" class="bg-green-50 border border-green-200 rounded-lg p-3">
                      <h4 class="text-sm font-medium text-green-800 mb-1">Available discounts:</h4>
                      <div class="space-y-1">
                        <div v-for="status in availableDiscountsWithLimits" :key="status.discount_id" class="text-sm text-green-700">
                          • <strong>{{ status.name }}</strong> - 
                          <span v-if="status.customer_limit">
                            You've used {{ status.customer_usage }}/{{ status.customer_limit }} times
                            <span v-if="status.remaining_for_customer" class="text-green-600">
                              ({{ status.remaining_for_customer }} remaining)
                            </span>
                          </span>
                          <span v-else class="text-green-600">No customer limit</span>
                        </div>
                      </div>
                    </div>
                    
                    <!-- Total Limit Reached -->
                    <div v-if="totalLimitReachedDiscounts.length > 0" class="bg-red-50 border border-red-200 rounded-lg p-3">
                      <h4 class="text-sm font-medium text-red-800 mb-1">These discounts are no longer available:</h4>
                      <div class="space-y-1">
                        <div v-for="status in totalLimitReachedDiscounts" :key="status.discount_id" class="text-sm text-red-700">
                          • <strong>{{ status.name }}</strong> - Total usage limit reached ({{ status.total_usage }}/{{ status.total_limit }})
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- Cart Items List -->
            <div class="space-y-6">
              <div v-for="item in cartStore.cart.items" :key="`${item.product_id}-${item.variant_id || 'no-variant'}`" class="bg-white border border-gray-200 rounded-lg p-6">
                <div class="flex items-start space-x-4">
                  <!-- Product Image -->
                  <div class="flex-shrink-0">
                    <img 
                      :src="item.image || '/placeholder-product.jpg'" 
                      :alt="item.product_name"
                      @error="handleImageError"
                      class="w-20 h-20 object-cover rounded-lg"
                    />
                  </div>

                  <!-- Product Details -->
                  <div class="flex-1 min-w-0">
                    <div>
                      <div class="flex items-center gap-2">
                        <h3 class="text-lg font-medium text-gray-900">{{ item.product_name || 'Product' }}</h3>
                        <!-- Discount indicator -->
                        <div v-if="item.discount_amount > 0" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-green-100 text-green-800">
                          <svg class="w-3 h-3 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                          </svg>
                          Server Discount
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
                          <span class="text-xs text-gray-500 ml-1">
                            ({{ getDiscountApplicationText(discount) }})
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
                          ${{ getFinalLineTotal(item).toFixed(2) }}
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

          <!-- Order Summary -->
          <div class="lg:col-span-1">
            <div class="bg-white border border-gray-200 rounded-lg p-6 sticky top-8">
              <h2 class="text-lg font-semibold text-gray-900 mb-4">Order Summary</h2>
              
              <!-- Security Notice -->
              <div class="bg-yellow-50 border border-yellow-200 rounded-lg p-3 mb-4">
                <div class="flex items-start">
                  <svg class="h-5 w-5 text-yellow-600 mt-0.5 mr-2 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
                  </svg>
                  <div>
                    <p class="text-xs text-yellow-800">
                      <strong>Important:</strong> Final pricing is calculated securely on our servers during checkout
                    </p>
                  </div>
                </div>
              </div>
              
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
                
                <!-- Item Discount Details -->
                <div v-if="cartStore.cart?.item_discount_details && cartStore.cart.item_discount_details.length > 0" class="bg-green-50 border border-green-200 rounded-lg p-3 mb-3">
                  <h4 class="text-sm font-medium text-green-800 mb-2">Item-Level Discounts Applied:</h4>
                  <div class="space-y-2">
                    <div v-for="discount in cartStore.cart.item_discount_details" :key="discount.discount_id" class="flex items-center justify-between">
                      <div class="flex items-center">
                        <svg class="h-4 w-4 text-green-600 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                        </svg>
                        <div>
                          <span class="text-sm text-green-700">{{ discount.name }}</span>
                          <div class="text-xs text-green-600">
                            {{ formatDiscountValue(discount) }} 
                            <span class="text-gray-500">({{ getDiscountApplicationText(discount) }})</span>
                          </div>
                        </div>
                      </div>
                      <span class="text-sm font-medium text-green-600">-${{ discount.amount.toFixed(2) }}</span>
                    </div>
                  </div>
                </div>
                
                <!-- Order-level discounts -->
                <div v-if="hasOrderDiscounts" class="flex justify-between text-sm">
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
                        <div>
                          <span class="text-sm text-blue-700">{{ discount.name }}</span>
                          <div class="text-xs text-blue-600">
                            {{ formatDiscountValue(discount) }} 
                            <span class="text-gray-500">(off order total)</span>
                          </div>
                        </div>
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
                <div v-if="hasAnyDiscounts" class="flex justify-between text-sm border-t border-gray-100 pt-2">
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
                  <span>${{ getGrandTotal().toFixed(2) }}</span>
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
                  class="w-full py-3 px-4 bg-black text-white rounded-md font-bold uppercase tracking-wide hover:bg-gray-800 disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center gap-2 transition-colors"
                >
                  <span v-if="checkoutLoading" class="animate-spin rounded-full h-5 w-5 border-b-2 border-white"></span>
                  {{ checkoutLoading ? 'Processing...' : 'Proceed to Checkout' }}
                </button>
                
                <button
                  @click="clearCart"
                  :disabled="cartStore.loading"
                  class="w-full py-2 px-4 text-gray-600 border border-gray-300 rounded-md font-medium hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
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
            <ShoppingCartIcon class="mx-auto h-24 w-24 text-gray-300" />
            <h3 class="mt-4 text-lg font-medium text-gray-900">Your cart is empty</h3>
            <p class="mt-2 text-gray-500">Start shopping to add items to your cart.</p>
            <div class="mt-6">
              <router-link
                :to="{ path: `/${shopSlug}/` }"
                class="inline-flex items-center px-4 py-2 bg-black text-white rounded-md font-medium hover:bg-gray-800"
              >
                Continue Shopping
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useCartStore } from '@/stores/cart';
import { placeOrder } from '@/services/order';
import type { CartItem, ItemDiscountDetail, OrderDiscountDetail } from '@/services/cart';
import { ShoppingCartIcon } from '@heroicons/vue/24/outline';
import { formatDiscountValue, getDiscountDescription } from '@/utils/discount';
import Loader from '@/components/Loader.vue';
import LoginPrompt from '@/components/LoginPrompt.vue';
import { useAuthStore } from '@/stores/auth';

const route = useRoute();
const cartStore = useCartStore();
const router = useRouter();
const authStore = useAuthStore();
const shopSlug = route.params.shopSlug as string;
const checkoutLoading = ref(false);

// Only fetch cart when user is authenticated and session loading is complete
function fetchCartIfAuthenticated() {
  if (!shopSlug) return;
  if (!authStore.sessionLoading && authStore.user) {
    cartStore.setShopSlug(shopSlug);
    cartStore.fetchCart();
  }
}

onMounted(() => {
  fetchCartIfAuthenticated();
});

// Watch for authentication state changes
watch(
  () => [authStore.sessionLoading, authStore.user],
  ([sessionLoading, user]) => {
    if (!sessionLoading && user) {
      fetchCartIfAuthenticated();
    }
  }
);

// Computed properties for discount display - READ ONLY from backend data
const hasItemDiscounts = computed(() => {
  if (!cartStore.cart?.items) return false;
  return cartStore.cart.items.some((item: CartItem) => item.discount_amount > 0);
});

const hasOrderDiscounts = computed(() => {
  if (!cartStore.cart?.total_discounts || !cartStore.cart?.items) return false;
  const itemDiscountsTotal = cartStore.cart.items.reduce((total: number, item: CartItem) => total + (item.discount_amount || 0), 0);
  return cartStore.cart.total_discounts > itemDiscountsTotal;
});

const hasAnyDiscounts = computed(() => {
  return hasItemDiscounts.value || hasOrderDiscounts.value;
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

// Computed properties for discount limit information - READ ONLY from backend data
const hasDiscountLimits = computed(() => {
  if (!cartStore.cart?.discount_statuses) return false;
  return cartStore.cart.discount_statuses.length > 0;
});

const customerLimitReachedDiscounts = computed(() => {
  if (!cartStore.cart?.discount_statuses) return [];
  return cartStore.cart.discount_statuses.filter((status: any) => status.customer_limit_hit);
});

const availableDiscountsWithLimits = computed(() => {
  if (!cartStore.cart?.discount_statuses) return [];
  return cartStore.cart.discount_statuses.filter((status: any) => 
    status.is_available && !status.customer_limit_hit && !status.total_limit_hit
  );
});

const totalLimitReachedDiscounts = computed(() => {
  if (!cartStore.cart?.discount_statuses) return [];
  return cartStore.cart.discount_statuses.filter((status: any) => status.total_limit_hit);
});

// Helper function to get discount details for a specific item - READ ONLY
function getItemDiscountDetails(item: CartItem): ItemDiscountDetail[] {
  if (!cartStore.cart?.item_discount_details) return [];
  return cartStore.cart.item_discount_details.filter((discount: ItemDiscountDetail) => 
    discount.product_id === item.product_id && 
    discount.variant_id === (item.variant_id || '')
  );
}

// Helper function to get discount description - READ ONLY
function getDiscountDescription(discount: ItemDiscountDetail | OrderDiscountDetail): string {
  if (discount.type === 'percentage') {
    return `${discount.value}% discount applied`;
  } else {
    return `$${discount.value.toFixed(2)} discount applied`;
  }
}

// Helper function to get discount application text
function getDiscountApplicationText(discount: ItemDiscountDetail | OrderDiscountDetail): string {
  if (discount.type === 'percentage') {
    return 'off each item';
  } else {
    return 'off per unit';
  }
}

// Helper function to get final line total safely
function getFinalLineTotal(item: CartItem): number {
  if (item.final_line_total !== undefined && item.final_line_total > 0) {
    return item.final_line_total;
  }
  if (item.line_total !== undefined) {
    return item.line_total;
  }
  return 0;
}

// Helper function to get grand total safely
function getGrandTotal(): number {
  if (cartStore.cart?.grand_total !== undefined && cartStore.cart.grand_total > 0) {
    return cartStore.cart.grand_total;
  }
  if (cartStore.cart?.subtotal !== undefined) {
    return cartStore.cart.subtotal;
  }
  return 0;
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
    // Place the order - backend will calculate all pricing securely
    const order = await cartStore.placeOrder();
    
    // Clear the cart after successful order placement
    await cartStore.clearCart();
    
    // Redirect to order confirmation page
    router.push({
      name: 'OrderConfirmation',
      params: { 
        shopSlug: cartStore.shopSlug,
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