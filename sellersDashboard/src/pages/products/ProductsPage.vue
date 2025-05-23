<template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto font-sans">
    <h2 class="text-3xl sm:text-4xl font-extrabold mb-8 text-gray-900 leading-tight">Product Catalog</h2>

    <div class="flex justify-end mb-6">
      <div class="inline-flex rounded-md shadow-sm" role="group">
        <button
          @click="currentView = 'list'"
          :class="{
            'bg-blue-600 text-white hover:bg-blue-700': currentView === 'list',
            'bg-gray-200 text-gray-700 hover:bg-gray-300': currentView !== 'list'
          }"
          class="px-4 py-2 text-sm font-medium border border-gray-300 rounded-l-lg focus:z-10 focus:ring-2 focus:ring-blue-500 focus:outline-none transition duration-150 ease-in-out"
        >
          <svg class="w-5 h-5 inline-block mr-1" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M3 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z" clip-rule="evenodd"></path></svg>
          List View
        </button>
        <button
          @click="currentView = 'grid'"
          :class="{
            'bg-blue-600 text-white hover:bg-blue-700': currentView === 'grid',
            'bg-gray-200 text-gray-700 hover:bg-gray-300': currentView !== 'grid'
          }"
          class="px-4 py-2 text-sm font-medium border border-gray-300 rounded-r-lg focus:z-10 focus:ring-2 focus:ring-blue-500 focus:outline-none transition duration-150 ease-in-out"
        >
          <svg class="w-5 h-5 inline-block mr-1" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path d="M5 3a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2V5a2 2 0 00-2-2H5zM11 3a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2V5a2 2 0 00-2-2h-2zM5 11a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2v-2a2 2 0 00-2-2H5zM11 11a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2v-2a2 2 0 00-2-2h-2z"></path></svg>
          Grid View
        </button>
      </div>
    </div>

    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-16">
      <svg class="animate-spin h-8 w-8 text-blue-500 mb-3" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
        <path class="opacity-75" fill="currentColor"
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
      </svg>
      <p class="text-lg">Loading products...</p>
    </div>

    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-lg shadow-sm mb-6" role="alert">
      <strong class="font-bold">Error loading products!</strong>
      <span class="block sm:inline ml-2">{{ error }}</span>
      <p class="text-sm mt-2">Please ensure you have an active shop selected and try again.</p>
    </div>

    <div v-else>
      <div v-if="products.length">
        <div v-if="currentView === 'list'" class="overflow-x-auto bg-white shadow-lg rounded-xl">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-100">
              <tr>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Image</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Name</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Category</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Price</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Created Date</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-200">
              <tr
                v-for="(prod, i) in products"
                :key="prod.id"
                @click="goToDetail(prod)"
                class="cursor-pointer transition duration-200 ease-in-out transform hover:scale-[1.005] hover:bg-blue-50"
                :class="{ 'bg-gray-50': i % 2 === 1 }"
              >
                <td class="py-3 px-6">
                  <div class="w-12 h-12 flex-shrink-0 rounded-lg overflow-hidden border border-gray-200 flex items-center justify-center bg-gray-100">
                    <img v-if="prod.images && prod.images.length > 0" :src="prod.images[0]" alt="Product thumbnail" class="w-full h-full object-cover"/>
                    <svg v-else class="w-8 h-8 text-gray-400" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M4 3a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V5a2 2 0 00-2-2H4zm12 12H4l4-4 4 4 4-4V5h-2v6l-4-4-4 4z" clip-rule="evenodd"></path></svg>
                  </div>
                </td>
                <td class="py-3 px-6 text-sm text-gray-800 font-medium">{{ prod.name }}</td>
                <td class="py-3 px-6 text-sm text-gray-700">{{ prod.category || 'Uncategorized' }}</td>
                <td class="py-3 px-6 text-sm text-green-600 font-semibold">
                  {{ prod.price != null ? `$${prod.price.toFixed(2)}` : 'N/A' }}
                </td>
                <td class="py-3 px-6 text-sm text-gray-600">{{ formatDate(prod.createdAt) }}</td>
              </tr>
            </tbody>
          </table>
        </div>

        <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
          <div
            v-for="prod in products"
            :key="prod.id"
            @click="goToDetail(prod)"
            class="bg-white rounded-xl shadow-lg overflow-hidden cursor-pointer transform hover:scale-105 hover:shadow-xl transition duration-200 ease-in-out flex flex-col"
          >
            <div class="w-full h-48 bg-gray-200 flex items-center justify-center text-gray-400">
              <img v-if="prod.images && prod.images.length > 0" :src="prod.images[0]" alt="Product image" class="w-full h-full object-cover"/>
              <svg v-else class="w-16 h-16" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M4 3a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V5a2 2 0 00-2-2H4zm12 12H4l4-4 4 4 4-4V5h-2v6l-4-4-4 4z" clip-rule="evenodd"></path></svg>
            </div>
            <div class="p-5 flex-grow flex flex-col">
              <h3 class="text-xl font-semibold text-gray-900 mb-2 truncate">{{ prod.name }}</h3>
              <p class="text-sm text-gray-600 mb-3">{{ prod.category || 'Uncategorized' }}</p>
              <div class="mt-auto"> <p class="text-2xl font-bold text-green-700">
                  {{ prod.price != null ? `$${prod.price.toFixed(2)}` : 'N/A' }}
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div v-else class="bg-blue-50 border border-blue-200 text-blue-700 px-6 py-8 rounded-lg text-center mt-8 shadow-sm">
        <p class="text-lg font-medium">No products found.</p>
        <p class="text-md mt-2">It looks like there are no products listed for your active shop yet.</p>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import { productService } from '@/services/product';

export default {
  name: 'ProductsPage',
  data() {
    return {
      products: [],
      loading: false,
      error: null,
      currentView: 'list' // Default view
    };
  },
  computed: {
    ...mapGetters('shops', ['activeShop'])
  },
  async created() {
    if (!this.activeShop) {
      this.error = 'No shop selected. Please select a shop from the dashboard to view its products.';
      return;
    }
    this.loading = true;
    try {
      this.products = await productService.fetchByShop(this.activeShop.id);
    } catch (e) {
      console.error('Error fetching products:', e);
      this.error = 'Failed to load products. There might be a network issue or no products exist for this shop.';
    } finally {
      this.loading = false;
    }
  },
  methods: {
    formatDate(dt) {
      return dt ? new Date(dt).toLocaleString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
      }) : '—';
    },
    goToDetail(prod) {
      if (this.activeShop && prod.id) {
        this.$router.push({
          name: 'ProductDetail',
          params: { shopId: this.activeShop.id, productId: prod.id }
        });
      } else {
        console.warn('Cannot navigate to product detail: missing shopId or productId');
      }
    }
  }
};
</script>

<style scoped>
/* TailwindCSS utility classes usually handle most of this.
   Keeping explicit animation for clarity, though it can often be inlined. */
@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}
.animate-spin {
  animation: spin 1s linear infinite;
}
</style>