<!-- src/pages/CartPage.vue -->
<template>
  <div class="container mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold mb-6">Your Cart</h1>

    <div v-if="cartItems.length === 0" class="text-gray-500">
      Your cart is empty.
    </div>

    <div v-else class="space-y-6">
      <div
        v-for="(item, index) in cartItems"
        :key="index"
        class="flex items-center justify-between bg-white shadow p-4 rounded-lg"
      >
        <div class="flex items-center gap-4">
          <img
            :src="item.image"
            alt="Product image"
            class="w-16 h-16 object-cover rounded"
          />
          <div>
            <h2 class="text-lg font-semibold">{{ item.name }}</h2>
            <p class="text-sm text-gray-600" v-if="item.variantOptions">
              <span
                v-for="(value, key) in item.variantOptions"
                :key="key"
                class="mr-2"
              >
                {{ key }}: {{ value }}
              </span>
            </p>
            <p class="text-sm text-gray-700">Price: ${{ item.price.toFixed(2) }}</p>
          </div>
        </div>

        <div class="flex items-center gap-3">
          <input
            type="number"
            min="1"
            class="w-16 border rounded px-2 py-1 text-center"
            v-model.number="item.quantity"
            @change="updateItemQuantity(item)"
          />
          <button
            @click="removeItem(item)"
            class="text-red-500 hover:text-red-700"
          >
            Remove
          </button>
        </div>
      </div>

      <!-- Summary -->
      <div class="text-right mt-8">
        <p class="text-xl font-semibold">
          Total: ${{ totalPrice.toFixed(2) }}
        </p>
        <button
          class="mt-4 bg-blue-600 text-white px-6 py-2 rounded hover:bg-blue-700"
          @click="checkout"
        >
          Proceed to Checkout
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useCartStore } from '@/stores/cart'
import { computed } from 'vue'

const cart = useCartStore()

const cartItems  = computed(() => cart.items)
const totalPrice = computed(() => cart.totalPrice)

function removeItem(item) {
  cart.removeItem(item.productId, item.variantOptions)
}

function updateItemQuantity(item) {
  if (item.quantity <= 0) {
    removeItem(item)
  }
  // quantity is already reactive through `v-model`, store is updated
}

function checkout() {
  // Here, you'd redirect to checkout page or trigger backend checkout
  alert('Checkout is not implemented yet!')
}
</script>

<style scoped>
/* Optional styles for extra polish */
</style>
