<template>
  <div class="p-4 max-w-7xl mx-auto">
    <h2 class="text-3xl font-bold mb-6 text-gray-800">Products Overview</h2>

    <div v-if="loading" class="flex items-center justify-center text-gray-600 py-8">
      <svg class="animate-spin h-5 w-5 mr-3 text-blue-500" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      Loading products…
    </div>
    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative mb-4" role="alert">
      <strong class="font-bold">Error!</strong>
      <span class="block sm:inline ml-2">{{ error }}</span>
    </div>
    <div v-else>
      <div class="overflow-x-auto shadow-md rounded-lg">
        <table class="min-w-full bg-white">
          <thead>
            <tr class="bg-gray-200 text-gray-700 uppercase text-sm leading-normal">
              <th class="py-3 px-6 text-left border-b border-gray-300">ID</th>
              <th class="py-3 px-6 text-left border-b border-gray-300">Name</th>
              <th class="py-3 px-6 text-left border-b border-gray-300">Category</th>
              <th class="py-3 px-6 text-left border-b border-gray-300">Price</th>
              <th class="py-3 px-6 text-left border-b border-gray-300">Created</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="(prod, index) in products"
              :key="prod.id"
              :class="{'bg-gray-50': index % 2 === 1, 'hover:bg-gray-100': true, 'border-b border-gray-200': true}"
            >
              <td class="py-3 px-6 text-sm break-all">{{ prod.id }}</td>
              <td class="py-3 px-6 text-sm">{{ prod.name }}</td>
              <td class="py-3 px-6 text-sm">{{ prod.category || '—' }}</td>
              <td class="py-3 px-6 text-sm">
                {{ typeof prod.price === 'number' ? `$${prod.price.toFixed(2)}` : '—' }}
              </td>
              <td class="py-3 px-6 text-sm">
                {{ formatDate(prod.createdAt) }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-if="products.length === 0 && !loading && !error" class="text-gray-500 mt-6 text-center py-8">
        No products found for this shop.
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import api from '@/services/api'; // <--- **Ensure this path is correct for your project**

export default {
  name: 'ProductsPage',
  data() {
    return {
      products: [],
      loading: false,
      error: null
    };
  },
  computed: {
    // This assumes you have a Vuex store with a 'shops' module and an 'activeShop' getter.
    // Ensure your Vuex store is properly set up to provide the activeShop.
    ...mapGetters('shops', ['activeShop'])
  },
  async created() {
    if (!this.activeShop) {
      this.error = 'No shop selected.';
      return;
    }

    this.loading = true;
    try {
      // This is where the actual API call to your backend happens.
      // Make sure your `api` service is correctly configured to hit your endpoint.
      const res = await api.get(
        `/seller/shops/${this.activeShop.id}/products`
      );

      this.products = res.data.map(p => ({
        id:          p.ID ?? p.id,          // Handles both 'ID' and 'id' from API
        name:        p.Name ?? p.name ?? '—', // Handles both 'Name' and 'name'
        category:    p.Category ?? p.category,
        // Ensure price is a number; default to null if not.
        price:       isFinite(p.Price) ? p.Price : isFinite(p.price) ? p.price : null,
        createdAt:   p.CreatedAt ?? p.createdAt
      }));
    } catch (err) {
      this.error = 'Failed to load products.';
      console.error('API Error:', err); // Log the actual error for debugging
    } finally {
      this.loading = false;
    }
  },
  methods: {
    formatDate(dt) {
      // Formats the date to a more readable string.
      return dt ? new Date(dt).toLocaleString() : '—';
    }
  }
};
</script>

<style scoped>
/* ensure long IDs wrap */
.break-all {
  word-break: break-all;
}

/* Simple spinner animation for loading state */
@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

.animate-spin {
  animation: spin 1s linear infinite;
}
</style>