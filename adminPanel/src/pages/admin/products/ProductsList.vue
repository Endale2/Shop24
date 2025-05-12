<template>
  <div>
    <h2 class="text-2xl mb-4">Products</h2>
    <table class="min-w-full bg-white border">
      <thead>
        <tr>
          <th class="px-4 py-2 border">ID</th>
          <th class="px-4 py-2 border">Name</th>
          <th class="px-4 py-2 border">Category</th>
          <th class="px-4 py-2 border">Price</th>
          <th class="px-4 py-2 border">Created At</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="p in products" :key="p.id" class="hover:bg-gray-100">
          <td class="px-4 py-2 border">{{ p.id }}</td>
          <td class="px-4 py-2 border">
            <router-link :to="`/products/${p.id}`" class="text-blue-600 hover:underline">
              {{ p.name || '—' }}
            </router-link>
          </td>
          <td class="px-4 py-2 border">{{ p.category || '—' }}</td>
          <td class="px-4 py-2 border">${{ p.price.toFixed(2) }}</td>
          <td class="px-4 py-2 border">{{ formatDate(p.createdAt) }}</td>
        </tr>
      </tbody>
    </table>
    <div v-if="products.length === 0" class="text-gray-600 mt-4">No products available.</div>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  name: 'ProductsList',
  data() {
    return { products: [] };
  },
  async created() {
    try {
      const res = await axios.get('/admin/products/', { withCredentials: true });
      this.products = res.data;
    } catch (e) {
      console.error('Error fetching products:', e);
    }
  },
  methods: {
    formatDate(dateStr) {
      return new Date(dateStr).toLocaleString();
    }
  }
};
</script>