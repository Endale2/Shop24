<template>
  <div class="p-4 max-w-7xl mx-auto">
    <h2 class="text-3xl font-bold mb-6 text-gray-800">Products Overview</h2>

    <div v-if="loading" class="flex items-center justify-center text-gray-600 py-8">
      <svg class="animate-spin h-5 w-5 mr-3 text-blue-500" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
        <path class="opacity-75" fill="currentColor"
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 
                 5.291A7.962 7.962 0 014 12H0c0 3.042 
                 1.135 5.824 3 7.938l3-2.647z"/>
      </svg>
      Loading products…
    </div>

    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
      <strong class="font-bold">Error!</strong>
      <span class="block sm:inline ml-2">{{ error }}</span>
    </div>

    <div v-else>
      <div v-if="products.length" class="overflow-x-auto shadow-md rounded-lg">
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
            <tr v-for="(prod, i) in products" :key="prod.id"
                :class="{
                  'bg-gray-50': i % 2,
                  'hover:bg-gray-100': true,
                  'border-b border-gray-200': true
                }">
              <td class="py-3 px-6 text-sm break-all">{{ prod.id }}</td>
              <td class="py-3 px-6 text-sm">{{ prod.name }}</td>
              <td class="py-3 px-6 text-sm">{{ prod.category || '—' }}</td>
              <td class="py-3 px-6 text-sm">
                {{ prod.price != null ? `$${prod.price.toFixed(2)}` : '—' }}
              </td>
              <td class="py-3 px-6 text-sm">{{ formatDate(prod.createdAt) }}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-else class="text-gray-500 mt-6 text-center py-8">
        No products found for this shop.
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters }    from 'vuex';
import { productService } from '@/services/product';

export default {
  name: 'ProductsPage',
  data() {
    return {
      products: [],
      loading:  false,
      error:    null
    };
  },
  computed: {
    ...mapGetters('shops', ['activeShop'])
  },
  async created() {
    if (!this.activeShop) {
      this.error = 'No shop selected.';
      return;
    }
    this.loading = true;
    try {
      this.products = await productService.fetchByShop(this.activeShop.id);
    } catch (e) {
      console.error(e);
      this.error = 'Failed to load products.';
    } finally {
      this.loading = false;
    }
  },
  methods: {
    formatDate(dt) {
      return dt ? new Date(dt).toLocaleString() : '—';
    }
  }
};
</script>

<style scoped>
.break-all { word-break: break-all; }
@keyframes spin { from { transform: rotate(0deg); } to { transform: rotate(360deg); } }
.animate-spin { animation: spin 1s linear infinite; }
</style>
