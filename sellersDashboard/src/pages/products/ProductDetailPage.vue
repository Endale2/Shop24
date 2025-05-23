<template>
  <div class="p-6 max-w-4xl mx-auto space-y-6">
    <!-- Back button -->
    <button
      @click="$router.back()"
      class="inline-flex items-center text-gray-600 hover:text-gray-800 mb-4"
    >
      <svg class="h-5 w-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M15 19l-7-7 7-7" />
      </svg>
      Back to Products
    </button>

    <!-- Loading -->
    <div v-if="loading" class="flex items-center justify-center text-gray-600 py-8">
      <svg class="animate-spin h-6 w-6 mr-2 text-blue-500" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
        <path class="opacity-75" fill="currentColor"
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 
                 5.291A7.962 7.962 0 014 12H0c0 3.042 
                 1.135 5.824 3 7.938l3-2.647z"/>
      </svg>
      Loading product…
    </div>

    <!-- Error -->
    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded">
      <strong class="font-bold">Error:</strong>
      <span class="ml-2">{{ error }}</span>
    </div>

    <!-- Product Detail -->
    <div v-else class="bg-white rounded-lg shadow p-6 space-y-6">
      <div class="flex flex-col lg:flex-row lg:space-x-8">
        <!-- Images -->
        <div class="flex-1 space-y-2">
          <img
            v-for="(img, idx) in product.images"
            :key="idx"
            :src="img"
            alt="Product image"
            class="w-full h-64 object-cover rounded"
          />
        </div>

        <!-- Info -->
        <div class="flex-1">
          <h1 class="text-2xl font-bold text-gray-800">{{ product.name }}</h1>
          <p class="text-gray-600 mt-1">{{ product.category }}</p>
          <p class="text-xl font-semibold text-green-600 mt-4">
            {{ product.price != null ? `$${product.price.toFixed(2)}` : '—' }}
          </p>
          <p class="mt-4 text-gray-700">{{ product.description }}</p>

          <!-- Variants -->
          <div v-if="product.variants.length" class="mt-6">
            <h3 class="font-semibold text-gray-800 mb-2">Variants</h3>
            <ul class="space-y-2">
              <li
                v-for="(v, i) in product.variants"
                :key="i"
                class="flex justify-between p-3 bg-gray-50 rounded"
              >
                <span class="text-gray-700">
                  {{ Object.entries(v.options).map(([k, val]) => `${k}: ${val}`).join(', ') }}
                </span>
                <span class="font-medium text-gray-800">
                  {{ isFinite(v.price) ? `$${v.price.toFixed(2)}` : '—' }}
                </span>
              </li>
            </ul>
          </div>
        </div>
      </div>

      <!-- Timestamps -->
      <div class="text-sm text-gray-500">
        <p>Created: {{ formatDate(product.createdAt) }}</p>
        <p>Updated: {{ formatDate(product.updatedAt) }}</p>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import { productService } from '@/services/product';

export default {
  name: 'ProductDetailPage',
  data() {
    return {
      product: {
        images:   [],
        variants: []
      },
      loading: true,
      error:   null
    };
  },
  computed: {
    ...mapGetters('shops', ['activeShop'])
  },
  async created() {
    if (!this.activeShop) {
      this.error = 'No shop selected.';
      this.loading = false;
      return;
    }
    const productId = this.$route.params.productId;
    try {
      this.product = await productService.fetchById(this.activeShop.id, productId);
    } catch (e) {
      console.error(e);
      this.error = 'Failed to load product details.';
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
@keyframes spin { from { transform: rotate(0deg); } to { transform: rotate(360deg); } }
.animate-spin { animation: spin 1s linear infinite; }
</style>
