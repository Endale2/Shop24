<template>
  <div class="container mx-auto px-4 sm:px-6 lg:px-8 py-10 md:py-16"> <h1 class="text-4xl md:text-5xl font-extrabold text-gray-900 mb-8 tracking-tight text-center">Your Shopping Cart</h1> <div v-if="cartItems.length === 0" class="bg-white rounded-2xl shadow-xl p-8 md:p-12 text-center text-gray-600 text-lg flex flex-col items-center justify-center min-h-[300px]"> <svg class="w-20 h-20 text-gray-400 mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.5">
        <path stroke-linecap="round" stroke-linejoin="round" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
      </svg>
      <p class="mb-6">Your cart is feeling a bit light! Let's fill it with something amazing.</p>
      <router-link to="/products" class="inline-flex items-center px-6 py-3 bg-green-600 text-white text-base font-semibold rounded-full shadow-lg hover:bg-green-700 transition-all duration-300 transform hover:-translate-y-1 hover:shadow-xl focus:outline-none focus:ring-4 focus:ring-green-400 focus:ring-offset-2 focus:ring-offset-white">
        <svg class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
        </svg>
        Start Shopping
      </router-link>
    </div>

    <div v-else class="grid grid-cols-1 lg:grid-cols-3 gap-8 md:gap-12">
      <div class="lg:col-span-2 space-y-6">
        <transition-group name="cart-item-fade" tag="div">
          <div
            v-for="item in cartItems"
            :key="item.id"
            class="flex items-center bg-white rounded-2xl shadow-md p-4 sm:p-6 border border-gray-100 transition-all duration-300 transform hover:scale-[1.005] hover:shadow-lg"
          >
            <div class="flex-shrink-0 w-24 h-24 sm:w-32 sm:h-32 rounded-xl overflow-hidden mr-4 sm:mr-6 bg-gray-50 border border-gray-100">
              <img
                :src="item.image || 'https://placehold.co/128x128/F0F9FF/1F2937?text=No+Image'"
                :alt="item.name"
                class="w-full h-full object-cover"
              />
            </div>

            <div class="flex-grow grid grid-cols-1 sm:grid-cols-2 gap-y-2 sm:gap-x-4 items-center">
              <div>
                <h2 class="text-xl font-semibold text-gray-900 mb-1 leading-tight">{{ item.name }}</h2>
                <p class="text-sm text-gray-600" v-if="item.variantOptions">
                  <span
                    v-for="(value, key) in item.variantOptions"
                    :key="key"
                    class="mr-3 last:mr-0 text-gray-500"
                  >
                    <span class="font-medium capitalize">{{ key }}:</span> {{ value }}
                  </span>
                </p>
                <p class="text-base text-gray-800 font-medium mt-1">Price: <span class="text-green-600 font-bold">${{ item.price.toFixed(2) }}</span></p>
              </div>

              <div class="flex items-center justify-end sm:justify-between flex-wrap gap-3">
                <div class="flex items-center border border-gray-300 rounded-lg overflow-hidden shadow-sm">
                  <button
                    @click="updateItemQuantity(item, item.quantity - 1)"
                    :disabled="item.quantity === 1"
                    class="px-3 py-2 bg-gray-50 text-gray-700 hover:bg-gray-100 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
                    aria-label="Decrease quantity"
                  >
                    -
                  </button>
                  <input
                    type="number"
                    min="1"
                    v-model.number="item.quantity"
                    @change="updateItemQuantity(item, item.quantity)"
                    class="w-16 px-2 py-2 text-center text-gray-800 focus:outline-none focus:ring-1 focus:ring-green-400 border-x border-gray-300"
                    aria-label="Product quantity"
                  />
                  <button
                    @click="updateItemQuantity(item, item.quantity + 1)"
                    class="px-3 py-2 bg-gray-50 text-gray-700 hover:bg-gray-100 transition-colors"
                    aria-label="Increase quantity"
                  >
                    +
                  </button>
                </div>
                <button
                  @click="removeItem(item)"
                  class="text-red-500 hover:text-red-700 font-medium text-sm sm:text-base transition-colors flex items-center gap-1"
                  aria-label="Remove item from cart"
                >
                  <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                  Remove
                </button>
              </div>
            </div>
          </div>
        </transition-group>
      </div>

      <div class="lg:col-span-1 bg-white rounded-2xl shadow-xl p-6 md:p-8 sticky top-28 h-fit border border-gray-100">
        <h2 class="text-2xl font-bold text-gray-900 mb-6 border-b pb-4 border-gray-200">Order Summary</h2>

        <div class="space-y-3 text-lg text-gray-700 mb-6">
          <div class="flex justify-between">
            <span>Subtotal ({{ cartItems.length }} items)</span>
            <span class="font-semibold">${{ totalPrice.toFixed(2) }}</span>
          </div>
          <div class="flex justify-between text-gray-500 text-base">
            <span>Shipping</span>
            <span>Calculated at checkout</span>
          </div>
          <div class="flex justify-between text-gray-500 text-base">
            <span>Taxes</span>
            <span>Calculated at checkout</span>
          </div>
        </div>

        <div class="flex justify-between items-center text-2xl font-extrabold text-gray-900 border-t pt-4 border-gray-200">
          <span>Total:</span>
          <span class="text-green-600">${{ totalPrice.toFixed(2) }}</span>
        </div>

        <button
          class="w-full mt-8 bg-green-600 text-white px-8 py-4 rounded-xl font-bold text-xl shadow-lg hover:bg-green-700 transition-all duration-300 transform hover:-translate-y-1 hover:shadow-xl focus:outline-none focus:ring-4 focus:ring-green-400 focus:ring-offset-2 focus:ring-offset-white"
          @click="checkout"
        >
          Proceed to Checkout
          <svg class="inline-block ml-2 w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M17 8l4 4m0 0l-4 4m4-4H3" />
          </svg>
        </button>

        <p class="text-center text-sm text-gray-500 mt-6">
          Shipping and taxes calculated at checkout.
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useCartStore } from '@/stores/cart';
import { computed } from 'vue';

const cart = useCartStore();

const cartItems = computed(() => cart.items);
const totalPrice = computed(() => cart.totalPrice);

// Modified removeItem to directly use product ID and variantOptions for accuracy
function removeItem(item) {
  cart.removeItem(item.id, item.variantOptions); // Assuming 'id' is the unique product identifier and 'variantOptions' ensures uniqueness of the cart item.
}

// Modified updateItemQuantity to take new quantity directly
function updateItemQuantity(item, newQuantity) {
  if (newQuantity <= 0) {
    removeItem(item);
  } else {
    // Assuming you have a method in your store like `updateQuantity`
    // that takes item id/variant and the new quantity.
    // The `v-model.number` already binds `item.quantity`
    // so this function might only be needed for the +/- buttons.
    // If your store reacts to changes in `item.quantity` directly (Pinia does this with `reactive`),
    // then no explicit update call might be needed for the input, only for the buttons.
    cart.updateQuantity(item.id, item.variantOptions, newQuantity); // Example call
  }
}

function checkout() {
  alert('Proceeding to checkout! (This is a placeholder action)');
  // In a real app, you would navigate to a checkout route:
  // router.push('/checkout');
}
</script>

<style scoped>
/* Transition styles for cart items being added/removed */
.cart-item-fade-enter-active,
.cart-item-fade-leave-active {
  transition: all 0.5s ease;
}
.cart-item-fade-enter-from,
.cart-item-fade-leave-to {
  opacity: 0;
  transform: translateX(-30px);
}
.cart-item-fade-leave-active {
  position: absolute; /* Ensures smooth animation when items are removed */
}
</style>