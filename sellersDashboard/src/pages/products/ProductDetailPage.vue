<template>
  <div class="p-4 sm:p-6 max-w-6xl mx-auto space-y-8 font-sans">
    <button
      @click="$router.back()"
      class="inline-flex items-center text-gray-600 hover:text-blue-700 transition duration-200 ease-in-out mb-6"
    >
      <svg class="h-5 w-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
      </svg>
      <span class="text-sm font-medium">Back to Products</span>
    </button>

    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-12">
      <svg class="animate-spin h-8 w-8 mr-2 text-blue-500" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
        <path class="opacity-75" fill="currentColor"
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
      </svg>
      <p class="mt-3 text-lg">Loading product details...</p>
    </div>

    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-lg shadow-sm" role="alert">
      <strong class="font-bold">Oops!</strong>
      <span class="ml-2">{{ error }}</span>
      <p class="text-sm mt-2">Please try again or contact support if the issue persists.</p>
    </div>

    <div v-else class="bg-white rounded-xl shadow-lg p-6 md:p-8 space-y-8">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 lg:gap-12">
        <div class="space-y-4">
          <img
            v-for="(img, idx) in product.images"
            :key="idx"
            :src="img"
            alt="Product image"
            class="w-full h-72 object-cover rounded-lg shadow-sm border border-gray-200"
          />
        </div>

        <div class="lg:sticky lg:top-8 self-start space-y-6">
          <div>
            <h1 class="text-4xl font-extrabold text-gray-900 leading-tight">{{ product.name }}</h1>
            <p class="text-lg text-gray-600 mt-2">{{ product.category }}</p>
          </div>

          <p class="text-4xl font-bold text-green-700">
            {{ product.price != null ? `$${product.price.toFixed(2)}` : 'Price N/A' }}
          </p>

          <p class="text-gray-700 leading-relaxed text-base">
            {{ product.description }}
          </p>

          <div v-if="product.variants && product.variants.length" class="mt-6 border-t border-gray-200 pt-6">
            <h3 class="font-semibold text-gray-800 text-xl mb-4">Available Variants</h3>
            <ul class="space-y-3">
              <li
                v-for="(v, i) in product.variants"
                :key="i"
                class="flex flex-col sm:flex-row sm:justify-between items-start sm:items-center p-4 bg-gray-50 rounded-lg shadow-sm border border-gray-100"
              >
                <span class="text-gray-700 font-medium text-base mb-1 sm:mb-0">
                  {{ Object.entries(v.options).map(([k, val]) => `${k}: ${val}`).join(', ') }}
                </span>
                <span class="font-bold text-gray-900 text-lg">
                  {{ isFinite(v.price) ? `$${v.price.toFixed(2)}` : 'N/A' }}
                </span>
              </li>
            </ul>
          </div>
        </div>
      </div>

      <div class="border-t border-gray-200 pt-6 text-sm text-gray-500 text-right">
        <p><strong>Created:</strong> {{ formatDate(product.createdAt) }}</p>
        <p><strong>Last Updated:</strong> {{ formatDate(product.updatedAt) }}</p>
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
        images: [],
        variants: []
      },
      loading: true,
      error: null
    };
  },
  computed: {
    ...mapGetters('shops', ['activeShop'])
  },
  async created() {
    if (!this.activeShop) {
      this.error = 'No shop selected. Please select a shop to view product details.';
      this.loading = false;
      return;
    }
    const productId = this.$route.params.productId;
    try {
      this.product = await productService.fetchById(this.activeShop.id, productId);
    } catch (e) {
      console.error('Error fetching product:', e);
      this.error = 'Failed to load product details. The product might not exist or there was a network issue.';
    } finally {
      this.loading = false;
    }
  },
  methods: {
    formatDate(dt) {
      return dt ? new Date(dt).toLocaleString('en-US', {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
      }) : 'Not available';
    }
  }
};
</script>

<style scoped>
/* Scoped styles for animations - TailwindCSS usually handles most of this */
@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}
.animate-spin {
  animation: spin 1s linear infinite;
}
</style>