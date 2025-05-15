<template>
  <div class="container mx-auto">
    <div class="mb-6 flex justify-between items-center">
      <h1 class="text-2xl font-semibold text-gray-900 dark:text-white">Products List</h1>
      <router-link
         to="/products/new"
         class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 dark:focus:ring-offset-gray-800"
       >
         Add New Product
       </router-link>
    </div>

    <div v-if="loading" class="flex items-center justify-center h-64 text-gray-500 dark:text-gray-400">
      <svg class="animate-spin h-8 w-8 text-blue-600 dark:text-blue-400 mr-3" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 0116 0H4z"></path>
      </svg>
      Loading products...
    </div>

    <div v-else-if="error" class="flex items-center justify-center h-64 text-red-600 dark:text-red-400">
      <p>Error loading products: {{ error }}</p>
    </div>

    <div v-else>
      <div v-if="products.length === 0" class="flex items-center justify-center h-64 text-gray-600 dark:text-gray-400">
        <p>No products available.</p>
      </div>

      <div v-else class="overflow-x-auto bg-white dark:bg-gray-800 shadow-md rounded-lg">
         <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
           <thead class="bg-gray-50 dark:bg-gray-700">
             <tr>
               <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                 ID
               </th>
               <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                 Name
               </th>
               <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                 Category
               </th>
               <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                 Price
               </th>
               <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                 Created At
               </th>
               <th scope="col" class="relative px-6 py-3">
                 <span class="sr-only">Actions</span>
               </th>
             </tr>
           </thead>
           <tbody class="bg-white dark:bg-gray-800 divide-y divide-gray-200 dark:divide-gray-700">
             <tr v-for="p in products" :key="p.id" class="hover:bg-gray-50 dark:hover:bg-gray-700/50">
               <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900 dark:text-white">
                 {{ p.id }}
               </td>
               <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-white">
                 <router-link :to="`/products/${p.id}`" class="text-blue-600 hover:text-blue-900 dark:text-blue-400 dark:hover:text-blue-300">
                    {{ p.name || '—' }}
                 </router-link>
               </td>
               <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 dark:text-gray-400">
                 {{ p.category || '—' }}
               </td>
               <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 dark:text-gray-400">
                 ${{ p.price ? p.price.toFixed(2) : '0.00' }}
               </td>
               <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 dark:text-gray-400">
                 {{ p.createdAt ? formatDate(p.createdAt) : '—' }}
               </td>
               <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                 <router-link :to="`/products/${p.id}/edit`" class="text-indigo-600 hover:text-indigo-900 dark:text-indigo-400 dark:hover:text-indigo-300 mr-4">Edit</router-link>
                 <button @click="deleteProduct(p.id)" class="text-red-600 hover:text-red-900 dark:text-red-400 dark:hover:text-red-300">Delete</button>
               </td>
             </tr>
           </tbody>
         </table>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'ProductsList',
  data() {
    return {
      products: [],
      loading: true, // Add loading state
      error: null,   // Add error state
    };
  },
  async created() {
    await this.fetchProducts();
  },
  methods: {
    async fetchProducts() {
       this.loading = true;
       this.error = null; // Clear previous errors
       try {
         const res = await axios.get('/admin/products/', { withCredentials: true });
         this.products = res.data;
       } catch (e) {
         console.error('Error fetching products:', e);
         this.error = e.message || 'An error occurred while fetching products.'; // Set user-friendly error message
         this.products = []; // Clear products on error
       } finally {
         this.loading = false; // Hide loading indicator
       }
    },
    formatDate(dateStr) {
      if (!dateStr) return '—';
       try {
           // Attempt to parse and format, handle potential errors
           const options = { year: 'numeric', month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' };
           return new Date(dateStr).toLocaleString(undefined, options);
       } catch (e) {
           console.error('Error formatting date:', dateStr, e);
           return 'Invalid Date';
       }
    },
    async deleteProduct(productId) {
       if (confirm('Are you sure you want to delete this product?')) {
          try {
             await axios.delete(`/admin/products/${productId}`, { withCredentials: true });
             // Remove the deleted product from the list
             this.products = this.products.filter(p => p.id !== productId);
             console.log(`Product ${productId} deleted.`);
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
/* Custom scrollbar styles can be added here if you want them specifically for this overflow div */
</style>