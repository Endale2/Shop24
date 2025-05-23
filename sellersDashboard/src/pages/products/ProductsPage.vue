<template>
  <div class="p-6 max-w-7xl mx-auto space-y-6">
    <!-- Title + Search -->
    <div class="flex flex-col md:flex-row md:items-center md:justify-between space-y-4 md:space-y-0">
      <h2 class="text-3xl font-bold text-gray-800">Products Overview</h2>
      <input
        v-model="search"
        type="text"
        placeholder="Search products..."
        class="w-full md:w-1/3 px-4 py-2 border border-gray-300 rounded-lg focus:ring focus:ring-green-200 focus:outline-none"
      />
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex justify-center py-16">
      <svg class="animate-spin h-10 w-10 text-green-500" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
        <path class="opacity-75" fill="currentColor"
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 
                 5.291A7.962 7.962 0 014 12H0c0 3.042 
                 1.135 5.824 3 7.938l3-2.647z"/>
      </svg>
    </div>

    <!-- Error -->
    <div v-else-if="error" class="bg-red-50 border border-red-200 text-red-700 px-6 py-4 rounded-lg">
      <p class="font-semibold">Oops—something went wrong:</p>
      <p class="mt-1">{{ error }}</p>
    </div>

    <!-- No Products -->
    <div v-else-if="filtered.length === 0" class="text-gray-500 text-center py-16">
      No products found<span v-if="search"> matching "{{ search }}"</span>.
    </div>

    <!-- Product Cards -->
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      <router-link
        v-for="prod in filtered"
        :key="prod.id"
        :to="{ name: 'ProductDetail', params: { productId: prod.id } }"
        class="block bg-white rounded-xl shadow hover:shadow-lg transition-shadow duration-200 overflow-hidden"
      >
        <div class="h-48 bg-gray-100">
          <img
            v-if="prod.images && prod.images.length"
            :src="prod.images[0]"
            alt="prod.name"
            class="w-full h-full object-cover"
          />
        </div>
        <div class="p-4 space-y-2">
          <h3 class="text-lg font-semibold text-gray-800 truncate">{{ prod.name }}</h3>
          <p class="text-sm text-gray-500">{{ prod.category || 'Uncategorized' }}</p>
          <div class="flex items-center justify-between mt-2">
            <span class="text-green-600 font-bold">
              {{ prod.price != null ? `$${prod.price.toFixed(2)}` : '—' }}
            </span>
            <span class="text-xs text-gray-400">{{ formatDate(prod.createdAt) }}</span>
          </div>
        </div>
      </router-link>
    </div>
  </div>
</template>

<script>
import { mapGetters }     from 'vuex';
import { productService } from '@/services/product';

export default {
  name: 'ProductsPage',
  data() {
    return {
      products: [],
      loading:  false,
      error:    null,
      search:   ''
    };
  },
  computed: {
    ...mapGetters('shops', ['activeShop']),
    filtered() {
      const term = this.search.trim().toLowerCase();
      if (!term) return this.products;
      return this.products.filter(p =>
        p.name.toLowerCase().includes(term) ||
        (p.category || '').toLowerCase().includes(term)
      );
    }
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
      return dt ? new Date(dt).toLocaleDateString() : '—';
    }
  }
};
</script>

<style scoped>
/* nothing extra for now—Tailwind handles it all */
</style>
