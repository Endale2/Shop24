<template>
  <div>
    <h1>Your Cart</h1>
    <div v-if="cartStore.loading">Loading...</div>
    <div v-if="cartStore.error" class="error">{{ cartStore.error }}</div>
    <div v-if="cartStore.cart && cartStore.cart.items.length">
      <table>
        <thead>
          <tr>
            <th>Product</th>
            <th>Variant</th>
            <th>Qty</th>
            <th>Unit Price</th>
            <th>Total</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in cartStore.cart.items" :key="item.product_id + item.variant_id">
            <td>{{ item.name }}</td>
            <td>{{ item.variant_name || '-' }}</td>
            <td>
              <input type="number" v-model.number="item.quantity" min="1"
                @change="updateItem(item)" />
            </td>
            <td>{{ item.unit_price }}</td>
            <td>{{ item.total_price }}</td>
            <td>
              <button @click="removeItem(item)">Remove</button>
            </td>
          </tr>
        </tbody>
      </table>
      <div>
        <div>Subtotal: {{ cartStore.cart.subtotal }}</div>
        <div v-if="cartStore.cart.discount_total">Discount: -{{ cartStore.cart.discount_total }}</div>
        <div v-if="cartStore.cart.shipping">Shipping: {{ cartStore.cart.shipping }}</div>
        <div><strong>Total: {{ cartStore.cart.total }}</strong></div>
      </div>
      <div>
        <input v-model="discountCode" placeholder="Discount code" />
        <button @click="applyDiscount">Apply</button>
      </div>
      <button @click="checkout" :disabled="cartStore.loading">Checkout (Cash on Delivery)</button>
      <button @click="clearCart">Clear Cart</button>
    </div>
    <div v-else>
      <p>Your cart is empty.</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useCartStore } from '@/stores/cart';
import { useRouter } from 'vue-router';
import { placeOrder } from '@/services/order';

const cartStore = useCartStore();
const router = useRouter();
const discountCode = ref('');

onMounted(() => {
  cartStore.fetchCart();
});

function updateItem(item: any) {
  cartStore.updateCartItem(item.product_id, item.variant_id, item.quantity);
}

function removeItem(item: any) {
  cartStore.removeCartItem(item.product_id, item.variant_id);
}

function applyDiscount() {
  if (discountCode.value) {
    cartStore.applyDiscount(discountCode.value);
  }
}

async function checkout() {
  if (!cartStore.cart || !cartStore.cart.items.length) return;
  try {
    const items = cartStore.cart.items.map((item: any) => ({
      product_id: item.product_id,
      variant_id: item.variant_id,
      quantity: item.quantity,
    }));
    const { data } = await placeOrder(cartStore.shopSlug, items, cartStore.cart.coupon_code);
    cartStore.clearCart();
    router.push({ name: 'OrderConfirmation', params: { orderId: data._id } });
  } catch (e: any) {
    alert(e.response?.data?.error || 'Checkout failed');
  }
}

function clearCart() {
  cartStore.clearCart();
}
</script> 