<template>
  <div>
    <h2 class="text-2xl mb-4">Dashboard</h2>
    <p>Welcome to the admin dashboard.</p>

    <div class="mt-6">
      <h3 class="text-lg font-semibold">Total Products:</h3>
      <p v-if="loading">Loading...</p>
      <p v-else-if="error" class="text-red-500">{{ error }}</p>
      <p v-else class="text-green-600 text-xl">{{ productCount }}</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'DashboardPage',
  data() {
    return {
      productCount: 0,
      loading: true,
      error: null,
    };
  },
  async created() {
    try {
      const res = await axios.get('/admin/products/count');
      this.productCount = res.data.count;
    } catch (err) {
      this.error = 'Failed to load product count';
    } finally {
      this.loading = false;
    }
  },
};
</script>
