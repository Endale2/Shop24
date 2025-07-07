<template>
  <div>
    <h1>Order Confirmation</h1>
    <div v-if="loading">Loading...</div>
    <div v-if="error" class="error">{{ error }}</div>
    <div v-if="order">
      <p>Order ID: {{ order._id }}</p>
      <p>Status: {{ order.status }}</p>
      <ul>
        <li v-for="item in order.items" :key="item.product_id + item.variant_id">
          {{ item.name }} (x{{ item.quantity }}) - {{ item.total_price }}
        </li>
      </ul>
      <div>Subtotal: {{ order.subtotal }}</div>
      <div v-if="order.discount_total">Discount: -{{ order.discount_total }}</div>
      <div v-if="order.shipping">Shipping: {{ order.shipping }}</div>
      <div><strong>Total: {{ order.total }}</strong></div>
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

onMounted(async () => {
  try {
    const { data } = await getOrderDetail(cartStore.shopSlug, route.params.orderId as string);
    order.value = data;
  } catch (e: any) {
    error.value = e.response?.data?.error || 'Failed to load order';
  } finally {
    loading.value = false;
  }
});
</script> 