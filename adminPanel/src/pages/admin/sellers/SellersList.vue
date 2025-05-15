<template>
  <div class="container mx-auto">
    <div class="mb-6 flex justify-between items-center">
      <h1 class="text-2xl font-semibold text-gray-900 dark:text-white">Seller Management</h1>
      <router-link
         to="/sellers/create"
         class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 dark:focus:ring-offset-gray-800"
       >
         Add New Seller
       </router-link>
    </div>

    <div v-if="loading" class="flex items-center justify-center h-64 text-gray-500 dark:text-gray-400 bg-white dark:bg-gray-800 rounded-lg shadow-md">
       <svg class="animate-spin h-8 w-8 text-blue-600 dark:text-blue-400 mr-3" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 0116 0H4z"></path>
      </svg>
      Loading seller data...
    </div>

    <div v-else-if="error" class="bg-red-100 dark:bg-red-900 border border-red-400 dark:border-red-700 text-red-700 dark:text-red-300 px-4 py-3 rounded-lg shadow-md mb-6" role="alert">
      <strong class="font-bold">Error:</strong>
      <span class="block sm:inline ml-2">{{ error }}</span>
    </div>

    <div v-else>
      <div v-if="sellers.length === 0" class="flex flex-col items-center justify-center h-64 text-gray-600 dark:text-gray-400 bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
         <svg class="mx-auto h-12 w-12 text-gray-400 dark:text-gray-500 mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
             <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
         </svg>
         <h3 class="mt-2 text-lg font-medium text-gray-900 dark:text-white">No Sellers Found</h3>
         <p class="mt-1 text-sm text-gray-500 dark:text-gray-400 text-center">
           There are currently no sellers to display. You can add a new seller using the button above.
         </p>
      </div>

      <div v-else class="overflow-x-auto bg-white dark:bg-gray-800 shadow-md rounded-lg">
         <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
           <thead class="bg-gray-50 dark:bg-gray-700">
             <tr>
               <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                 Name
               </th>
               <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                 Business Name
               </th>
               <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                 Email
               </th>
               <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                 Ratings
               </th>
               <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                 Total Sales
               </th>
               <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                 Reviews
               </th>
               <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                 Joined
               </th>
               <th scope="col" class="relative px-6 py-3">
                 <span class="sr-only">Actions</span>
               </th>
             </tr>
           </thead>
           <tbody class="bg-white dark:bg-gray-800 divide-y divide-gray-200 dark:divide-gray-700">
             <tr v-for="seller in sellers" :key="seller.id" class="hover:bg-gray-50 dark:hover:bg-gray-700/50 transition-colors duration-150">
               <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900 dark:text-white">
                 <div class="flex items-center">
                    <div class="flex-shrink-0 h-10 w-10">
                      <div class="h-10 w-10 bg-blue-100 dark:bg-blue-900 rounded-full flex items-center justify-center text-blue-700 dark:text-blue-300 text-sm font-semibold">
                        {{ seller.name ? seller.name.charAt(0).toUpperCase() : (seller.email ? seller.email.charAt(0).toUpperCase() : 'S') }}
                      </div>
                    </div>
                    <div class="ml-4">
                      <router-link :to="`/sellers/${seller.id}`" class="text-sm font-medium text-blue-600 hover:text-blue-900 dark:text-blue-400 dark:hover:text-blue-300">
                        {{ seller.name || 'N/A' }}
                      </router-link>
                    </div>
                 </div>
               </td>
               <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-white">
                 {{ seller.businessName || 'N/A' }}
               </td>
               <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700 dark:text-gray-300">
                 {{ seller.email || 'N/A' }}
               </td>
               <td class="px-6 py-4 whitespace-nowrap text-center text-sm">
                 <span
                   :class="[
                     'px-2 inline-flex text-xs leading-5 font-semibold rounded-full',
                     getRatingBadgeClass(seller.ratings)
                   ]"
                 >
                   {{ seller.ratings != null ? seller.ratings.toFixed(1) : 'N/A' }} ⭐
                 </span>
               </td>
               <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700 dark:text-gray-300 text-right">
                 {{ seller.totalSales != null ? '$' + seller.totalSales.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 2 }) : '$0.00' }}
               </td>
               <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700 dark:text-gray-300 text-center">
                 {{ seller.reviewsCount != null ? seller.reviewsCount : '0' }}
               </td>
               <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700 dark:text-gray-300">
                 {{ formatDate(seller.createdAt) }}
               </td>
               <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                 <router-link :to="`/sellers/${seller.id}/edit`" class="text-indigo-600 hover:text-indigo-900 dark:text-indigo-400 dark:hover:text-indigo-300 mr-4">Edit</router-link>
                 <button @click="confirmDelete(seller)" class="text-red-600 hover:text-red-900 dark:text-red-400 dark:hover:text-red-300">Delete</button>
               </td>
             </tr>
           </tbody>
         </table>
      </div>
    </div>

    <div
      v-if="showDeleteModal"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50"
      @click.self="showDeleteModal = false; sellerToDelete = null" >
      <div class="bg-white dark:bg-gray-800 rounded-lg p-6 space-y-6 w-full max-w-sm shadow-xl">
        <h3 class="text-lg font-semibold text-gray-900 dark:text-white">Confirm Delete</h3>
        <p class="text-gray-600 dark:text-gray-300">
           Are you sure you want to delete seller "{{ sellerToDelete ? (sellerToDelete.name || 'Unnamed Seller') : '' }}"? This action cannot be undone.
        </p>

        <div v-if="deleteError" class="text-red-600 dark:text-red-400 text-sm">{{ deleteError }}</div>

        <div class="flex justify-end space-x-3">
          <button
            @click="showDeleteModal = false; deleteError = null; sellerToDelete = null;"
            class="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm text-sm font-medium text-gray-700 dark:text-gray-300 bg-white dark:bg-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 dark:focus:ring-offset-gray-800 disabled:opacity-50 disabled:cursor-not-allowed"
            :disabled="deleting"
          >
            Cancel
          </button>
          <button
            @click="deleteSellerConfirmed"
            class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 dark:focus:ring-offset-gray-800 disabled:opacity-50 disabled:cursor-not-allowed"
            :disabled="deleting"
          >
             <svg v-if="deleting" class="animate-spin -ml-1 mr-2 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
               <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
               <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 0116 0H4z"></path>
             </svg>
            {{ deleting ? 'Deleting...' : 'Yes, Delete' }}
          </button>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'SellersListDashboard',
  data() {
    return {
      sellers: [],
      loading: true,
      error: null,
      showDeleteModal: false,
      sellerToDelete: null, // Store the full seller object to display name in modal
      deleting: false, // State for delete operation
      deleteError: null, // Error message for delete operation
    };
  },
  async created() {
    this.fetchSellers();
  },
  methods: {
    async fetchSellers() {
      this.loading = true;
      this.error = null;
      try {
        // Ensure this endpoint is correct
        const res = await axios.get('/admin/sellers/', { withCredentials: true });

        // Adjust parsing based on your actual API response structure
        const apiData = res.data;
        let rawSellers = [];

        if (Array.isArray(apiData)) {
            rawSellers = apiData;
        } else if (apiData && Array.isArray(apiData.data)) { // Common API pattern
            rawSellers = apiData.data;
        } else if (apiData && Array.isArray(apiData.sellers)) { // Another common API pattern
            rawSellers = apiData.sellers;
        } else {
           // Log unexpected format but don't throw if it's just an empty response sometimes
           console.warn('Fetch sellers: Response data is not in expected array format.', apiData);
        }

        this.sellers = rawSellers.map(s => this.mapSellerData(s)).filter(s => s.id != null); // Ensure sellers have an ID


      } catch (err) {
        console.error('Error fetching sellers:', err);
        this.error = `Failed to fetch sellers: ${err.response?.data?.message || err.message || 'Unknown error'}`;
        this.sellers = []; // Clear list on error
      } finally {
        this.loading = false;
      }
    },
    mapSellerData(s) { // Helper function to map data, handle potential nulls/undefineds/casing
        if (!s) return {}; // Return empty object for null/undefined source

        return {
            id: s.id || s.ID || null, // Prioritize 'id', fallback to 'ID', then null
            user_id: s.user_id || s.UserId || null,
            shop_ids: Array.isArray(s.shop_ids) ? s.shop_ids : (Array.isArray(s.ShopIds) ? s.ShopIds : []), // Ensure it's an array
            name: s.name || s.Name || null,
            email: s.email || s.Email || null,
            phone: s.phone || s.Phone || null,
            address: s.address || s.Address || null,
            businessName: s.business_name || s.BusinessName || null,
            // Parse numbers carefully, providing null/default if parsing fails or source is null/undefined
            ratings: s.ratings !== undefined && s.ratings !== null ? parseFloat(s.ratings) : null,
            totalSales: s.total_sales !== undefined && s.total_sales !== null ? parseFloat(s.total_sales) : null,
            reviewsCount: s.reviews_count !== undefined && s.reviews_count !== null ? parseInt(s.reviews_count, 10) : null,
            createdAt: s.created_at || s.CreatedAt || null,
            updatedAt: s.updated_at || s.UpdatedAt || null,
        };
    },
    formatDate(dateString) {
      if (!dateString) return '—'; // Consistent placeholder
       try {
           // Use consistent formatting options, maybe just date for brevity in table
           const options = { year: 'numeric', month: 'short', day: 'numeric' };
           return new Date(dateString).toLocaleDateString(undefined, options);
       } catch (e) {
           console.error('Error formatting date:', dateString, e);
           return 'Invalid Date'; // Indicate formatting failure
       }
    },
    getRatingBadgeClass(rating) {
      if (rating == null || rating === undefined) return 'bg-gray-200 text-gray-800 dark:bg-gray-700 dark:text-gray-300'; // Handle null/undefined consistently
      if (rating >= 4.5) return 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-300';
      if (rating >= 4.0) return 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-300';
      if (rating >= 3.0) return 'bg-orange-100 text-orange-800 dark:bg-orange-900 dark:text-orange-300';
      return 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-300'; // Less than 3 or invalid parse
    },
    confirmDelete(seller) {
      this.sellerToDelete = seller;
      this.deleteError = null; // Clear previous delete errors when opening modal
      this.showDeleteModal = true;
    },
    async deleteSellerConfirmed() {
      if (!this.sellerToDelete || !this.sellerToDelete.id) {
         this.deleteError = "Seller data is incomplete for deletion.";
         console.error(this.deleteError, this.sellerToDelete);
         return;
      }

      this.deleting = true;
      this.deleteError = null;

      try {
        // Replace with your actual delete API call
         await axios.delete(`/admin/sellers/${this.sellerToDelete.id}`, { withCredentials: true });

        console.log(`Seller ID ${this.sellerToDelete.id} deleted.`);
        // Remove from local list after successful deletion
        this.sellers = this.sellers.filter(s => s.id !== this.sellerToDelete.id);

        // Reset modal state after successful delete
        this.showDeleteModal = false;
        this.sellerToDelete = null;
        // Optionally, show a success notification here (e.g., a toast)

      } catch (err) {
        console.error('Error deleting seller:', err);
        this.deleteError = err.response?.data?.message || err.message || 'Failed to delete seller.';
        // Keep modal open to show the error: this.showDeleteModal = true;
        // Don't clear sellerToDelete yet so the name remains in the error message
      } finally {
        this.deleting = false;
      }
    }
  }
};
</script>

<style scoped>
/* Remove the fixed min-width unless absolutely necessary */
/* .min-w-full {
  min-width: 800px;
} */
</style>