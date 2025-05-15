<template>
  <div class="container mx-auto">
    <div class="mb-6 flex justify-between items-center">
      <h1 class="text-2xl font-semibold text-gray-900 dark:text-white">
        Product Details: {{ product ? (product.name || 'Unnamed Product') : '...' }}
      </h1>
      <div v-if="product && !loading">
        <router-link
          :to="`/products/${product.id}/edit`"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 dark:focus:ring-offset-gray-800 mr-3"
        >
          Edit
        </router-link>
        <button
          @click="deleteProduct(product.id)"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 dark:focus:ring-offset-gray-800"
        >
          Delete
        </button>
      </div>
    </div>

    <div v-if="loading" class="flex items-center justify-center h-64 text-gray-500 dark:text-gray-400">
       <svg class="animate-spin h-8 w-8 text-blue-600 dark:text-blue-400 mr-3" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 0116 0H4z"></path>
      </svg>
      Loading product details...
    </div>

    <div v-else-if="!product" class="flex items-center justify-center h-64 text-red-600 dark:text-red-400">
       <p v-if="error">Error loading product: {{ error }}</p>
       <p v-else>Product not found.</p>
    </div>

    <div v-else class="bg-white dark:bg-gray-800 shadow-md rounded-lg p-6">

      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div v-if="product.images && product.images.length > 0" class="md:col-span-1">
          <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">Images</h3>
          <div class="space-y-4">
             <img
               v-for="(image, index) in product.images"
               :key="index"
               :src="image"
               :alt="`Product Image ${index + 1}`"
               class="w-full h-auto object-cover rounded-md border border-gray-200 dark:border-gray-700 shadow-sm"
             >
          </div>
        </div>

        <div :class="{'md:col-span-2': product.images && product.images.length > 0, 'md:col-span-3': !(product.images && product.images.length > 0) }">
          <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">Information</h3>

          <dl class="divide-y divide-gray-200 dark:divide-gray-700">
            <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">ID</dt>
              <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                 {{ product.id }}
              </dd>
            </div>
             <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Name</dt>
              <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                 {{ product.name || '—' }}
              </dd>
            </div>
             <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Description</dt>
              <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                 {{ product.description || '—' }}
              </dd>
            </div>
             <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Category</dt>
              <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                 {{ product.category || '—' }}
              </dd>
            </div>
             <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Price</dt>
              <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                 {{ product.price != null ? '$' + product.price.toFixed(2) : '—' }}
              </dd>
            </div>
             <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Created At</dt>
              <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                 {{ formatDate(product.createdAt) }}
              </dd>
            </div>
              <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Updated At</dt>
              <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                 {{ formatDate(product.updatedAt) }}
              </dd>
            </div>
            <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Shop ID</dt>
              <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                 {{ product.shop_id || '—' }}
              </dd>
            </div>
             <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">User ID</dt>
              <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                 {{ product.user_id || '—' }}
              </dd>
            </div>
             <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Created By</dt>
              <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                 {{ product.createdBy || '—' }}
              </dd>
            </div>
             </dl>
        </div>
      </div>

    </div>

    <div class="mt-6">
       <router-link to="/products" class="inline-flex items-center px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm text-sm font-medium text-gray-700 dark:text-gray-300 bg-white dark:bg-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 dark:focus:ring-offset-gray-800">
         Back to Products
       </router-link>
     </div>

  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'ProductDetail',
  // Assuming you're using route props: true in your router config
  // props: ['id'], // If using route props: true

  data() {
    return {
      product: null,
      loading: true,
      error: null, // Added error state
    };
  },
  async created() {
    // Get ID from route params if not using props
    const productId = this.$route.params.id;
    if (!productId) {
        this.error = "Product ID is missing from the route.";
        this.loading = false;
        console.error(this.error);
        return; // Stop execution if ID is missing
    }

    try {
      const res = await axios.get(`/admin/products/${productId}`, { withCredentials: true });
      this.product = res.data;
       // If the API returns 200 but no data (e.g., null or empty object) for a non-existent ID,
       // you might want to treat it as not found here. Depending on your API.
       if (!this.product) {
           this.error = "Product not found.";
       }

    } catch (e) {
      console.error('Error fetching product:', e);
      // Handle specific error responses (e.g., 404)
      if (e.response && e.response.status === 404) {
          this.error = "Product not found.";
      } else {
          this.error = e.message || 'An error occurred while fetching product details.';
      }
       this.product = null; // Ensure product is null on error
    } finally {
      this.loading = false;
    }
  },
  methods: {
     formatDate(dateStr) {
       if (!dateStr) return '—';
        try {
            const options = { year: 'numeric', month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' };
            return new Date(dateStr).toLocaleString(undefined, options);
        } catch (e) {
            console.error('Error formatting date:', dateStr, e);
            return 'Invalid Date';
        }
     },
     async deleteProduct(productId) {
        if (confirm('Are you sure you want to delete this product? This action cannot be undone.')) {
           try {
              await axios.delete(`/admin/products/${productId}`, { withCredentials: true });
              console.log(`Product ${productId} deleted.`);
              // Redirect to the products list after deletion
              this.$router.push('/products');
              // Optionally show a success notification
           } catch (e) {
              console.error('Error deleting product:', e);
              alert('Failed to delete product. ' + (e.response?.data?.message || e.message));
              // Optionally show an error notification
           }
        }
     }
  }
};
</script>

<style scoped>
/* No specific styles needed if using Tailwind for everything */
</style>