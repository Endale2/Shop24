<template>
    <div>
      <h2 class="text-2xl mb-4">Product Detail</h2>
      <div v-if="loading" class="text-gray-500">Loading...</div>
      <div v-else-if="product">
        <table class="min-w-full bg-white border mb-4">
          <tbody>
            <tr><td class="px-4 py-2 font-semibold">ID</td><td class="px-4 py-2">{{ product.id }}</td></tr>
            <tr><td class="px-4 py-2 font-semibold">Name</td><td class="px-4 py-2">{{ product.name || '—' }}</td></tr>
            <tr><td class="px-4 py-2 font-semibold">Description</td><td class="px-4 py-2">{{ product.description || '—' }}</td></tr>
            <tr><td class="px-4 py-2 font-semibold">Category</td><td class="px-4 py-2">{{ product.category || '—' }}</td></tr>
            <tr><td class="px-4 py-2 font-semibold">Price</td><td class="px-4 py-2">${{ product.price.toFixed(2) }}</td></tr>
            <tr><td class="px-4 py-2 font-semibold">Created At</td><td class="px-4 py-2">{{ formatDate(product.createdAt) }}</td></tr>
            <tr><td class="px-4 py-2 font-semibold">Updated At</td><td class="px-4 py-2">{{ formatDate(product.updatedAt) }}</td></tr>
          </tbody>
        </table>
        <router-link to="/products" class="px-4 py-2 bg-gray-200 rounded hover:bg-gray-300">Back to Products</router-link>
      </div>
      <div v-else class="text-red-500">Product not found.</div>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  export default {
    name: 'ProductDetail',
    props: ['id'],
    data() {
      return {
        product: null,
        loading: true
      };
    },
    async created() {
      try {
        const res = await axios.get(`/admin/products/${this.$route.params.id}`, { withCredentials: true });
        this.product = res.data;
      } catch (e) {
        console.error('Error fetching product:', e);
      } finally {
        this.loading = false;
      }
    },
    methods: {
      formatDate(dateStr) {
        return new Date(dateStr).toLocaleString();
      }
    }
  };
  </script>