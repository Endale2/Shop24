<template>
  <div class="container mx-auto">
    <div v-if="loading" class="flex items-center justify-center h-64 text-gray-500 dark:text-gray-400 bg-white dark:bg-gray-800 rounded-lg shadow-md">
       <svg class="animate-spin h-8 w-8 text-blue-600 dark:text-blue-400 mr-3" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 0116 0H4z"></path>
      </svg>
      Loading seller details...
    </div>

    <div v-else-if="error" class="bg-red-100 dark:bg-red-900 border border-red-400 dark:border-red-700 text-red-700 dark:text-red-300 px-4 py-3 rounded-lg shadow-md mb-6" role="alert">
       <div class="flex items-center mb-2">
         <svg class="h-6 w-6 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
         <h3 class="font-bold text-xl">Error Fetching Seller</h3>
       </div>
       <p>{{ error }}</p>
       <router-link
         to="/sellers"
         class="mt-4 inline-block bg-red-600 dark:bg-red-700 hover:bg-red-700 dark:hover:bg-red-600 text-white font-semibold py-2 px-4 rounded-md transition-colors duration-150 shadow-sm focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 dark:focus:ring-offset-gray-800"
       >
         Back to Sellers List
       </router-link>
     </div>

    <div v-else-if="seller" class="bg-white dark:bg-gray-800 shadow-xl rounded-lg overflow-hidden">
      <div class="bg-gradient-to-r from-blue-600 to-indigo-600 dark:from-blue-800 dark:to-indigo-800 p-6 sm:p-8 text-white">
        <div class="flex flex-col sm:flex-row items-center justify-between">
          <div>
            <h1 class="text-3xl sm:text-4xl font-bold mb-1">{{ seller.name || 'Unnamed Seller' }}</h1>
            <p class="text-lg text-blue-200 dark:text-blue-300">{{ seller.businessName || 'N/A' }}</p>
          </div>
          <div class="mt-4 sm:mt-0 flex-shrink-0">
            <router-link
              :to="`/sellers/${seller.id}/edit`"
              class="inline-flex items-center bg-white dark:bg-gray-700 text-blue-600 dark:text-blue-300 hover:bg-blue-50 dark:hover:bg-gray-600 font-semibold py-2 px-5 rounded-md shadow transition-all duration-150 ease-in-out transform hover:scale-105 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-white dark:focus:ring-gray-800"
            >
              Edit Seller
            </router-link>
          </div>
        </div>
        <div class="mt-8 grid grid-cols-1 sm:grid-cols-3 gap-4 text-center text-white">
          <div>
            <p class="text-sm text-blue-200 dark:text-blue-300 uppercase tracking-wider">Rating</p>
            <p class="text-2xl font-semibold mt-1">
              {{ seller.ratings != null ? seller.ratings.toFixed(1) : 'N/A' }} ⭐
            </p>
          </div>
          <div>
            <p class="text-sm text-blue-200 dark:text-blue-300 uppercase tracking-wider">Total Sales</p>
            <p class="text-2xl font-semibold mt-1">
              {{ seller.totalSales != null ? '$' + seller.totalSales.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 2 }) : '$0.00' }}
            </p>
          </div>
          <div>
            <p class="text-sm text-blue-200 dark:text-blue-300 uppercase tracking-wider">Reviews</p>
            <p class="text-2xl font-semibold mt-1">
              {{ seller.reviewsCount != null ? seller.reviewsCount.toLocaleString() : '0' }}
            </p>
          </div>
        </div>
      </div>

      <div class="p-6 sm:p-8 space-y-8 text-gray-800 dark:text-gray-200">
        <section>
          <h2 class="text-xl font-semibold text-gray-700 dark:text-gray-300 mb-4 border-b border-gray-200 dark:border-gray-700 pb-2">Contact Information</h2>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-x-8 gap-y-4 text-sm">
            <div class="flex items-center">
              <div class="flex-shrink-0 text-blue-500 dark:text-blue-400 mr-3">
                <svg class="h-5 w-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path d="M2.003 5.884L10 9.882l7.997-3.998A2 2 0 0016 4H4a2 2 0 00-1.997 1.884z"></path><path d="M18 8.118l-8 4-8-4V14a2 2 0 002 2h12a2 2 0 002-2V8.118z"></path></svg>
              </div>
              <div>
                <p class="text-gray-500 dark:text-gray-400">Email Address</p>
                <p class="font-medium">{{ seller.email || 'N/A' }}</p>
              </div>
            </div>
            <div class="flex items-center">
              <div class="flex-shrink-0 text-blue-500 dark:text-blue-400 mr-3">
                <svg class="h-5 w-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M7 2a2 2 0 00-2 2v12a2 2 0 002 2h6a2 2 0 002-2V4a2 2 0 00-2-2H7zm3 14a1 1 0 100-2 1 1 0 000 2z" clip-rule="evenodd"></path></svg>
              </div>
              <div>
                <p class="text-gray-500 dark:text-gray-400">Phone Number</p>
                <p class="font-medium">{{ seller.phone || 'N/A' }}</p>
              </div>
            </div>
            <div class="flex items-start md:col-span-2">
              <div class="flex-shrink-0 text-blue-500 dark:text-blue-400 mr-3 mt-1">
                 <svg class="h-5 w-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0 010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z" clip-rule="evenodd"></path></svg>
               </div>
              <div>
                <p class="text-gray-500 dark:text-gray-400">Full Address</p>
                <p class="font-medium">{{ seller.address || 'N/A' }}</p>
              </div>
            </div>
          </div>
        </section>

        <section>
          <h2 class="text-xl font-semibold text-gray-700 dark:text-gray-300 mb-4 border-b border-gray-200 dark:border-gray-700 pb-2">Account & Shop Details</h2>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-x-8 gap-y-4 text-sm">
            <div>
              <p class="text-gray-500 dark:text-gray-400">Seller ID</p>
              <p class="font-mono text-xs bg-gray-100 dark:bg-gray-700 text-gray-800 dark:text-gray-200 px-2 py-1 rounded inline-block">{{ seller.id || 'N/A' }}</p>
            </div>
            <div>
              <p class="text-gray-500 dark:text-gray-400">User ID (Associated)</p>
              <p class="font-mono text-xs bg-gray-100 dark:bg-gray-700 text-gray-800 dark:text-gray-200 px-2 py-1 rounded inline-block">{{ seller.user_id || 'N/A' }}</p>
            </div>
            <div class="md:col-span-2">
              <p class="text-gray-500 dark:text-gray-400 mb-1">Associated Shop IDs</p>
              <div v-if="seller.shop_ids && seller.shop_ids.length > 0" class="flex flex-wrap gap-2">
                <span v-for="shopId in seller.shop_ids" :key="shopId" class="bg-purple-100 dark:bg-purple-900 text-purple-700 dark:text-purple-300 text-xs font-mono px-3 py-1 rounded-full">
                  {{ shopId }}
                </span>
              </div>
              <p v-else class="text-gray-700 dark:text-gray-300 italic text-sm">No associated shops.</p>
            </div>
          </div>
        </section>

        <section>
          <h2 class="text-xl font-semibold text-gray-700 dark:text-gray-300 mb-4 border-b border-gray-200 dark:border-gray-700 pb-2">Timestamps</h2>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-x-8 gap-y-4 text-sm">
            <div>
              <p class="text-gray-500 dark:text-gray-400">Account Created</p>
              <p class="font-medium">{{ formatDate(seller.createdAt, true) || 'N/A' }}</p>
            </div>
            <div>
              <p class="text-gray-500 dark:text-gray-400">Last Updated</p>
              <p class="font-medium">{{ formatDate(seller.updatedAt, true) || 'N/A' }}</p>
            </div>
          </div>
        </section>
      </div>

      <div class="px-6 py-4 bg-gray-50 dark:bg-gray-700 border-t border-gray-200 dark:border-gray-700 text-right">
        <router-link
          to="/sellers"
          class="inline-flex items-center text-blue-600 dark:text-blue-400 hover:text-blue-800 dark:hover:text-blue-300 font-medium transition-colors duration-150"
        >
           &larr; Back to Sellers List
        </router-link>
      </div>
    </div>

    <div v-else class="flex flex-col items-center justify-center h-64 text-gray-600 dark:text-gray-400 bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
       <svg class="mx-auto h-12 w-12 text-gray-400 dark:text-gray-500 mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
         </svg>
      <h2 class="text-2xl font-semibold text-gray-700 dark:text-gray-300">Seller Not Found</h2>
      <p class="text-gray-500 dark:text-gray-400 mt-2 text-center">The seller you are looking for does not exist or could not be loaded.</p>
      <router-link
         to="/sellers"
         class="mt-6 inline-block bg-blue-600 dark:bg-blue-700 hover:bg-blue-700 dark:hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded-md transition-colors duration-150 shadow-sm focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 dark:focus:ring-offset-gray-800"
       >
         Go to Sellers List
       </router-link>
    </div>

  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'SellerDetail',
  data() {
    return {
      seller: null,
      loading: true,
      error: null, // Error for initial fetch
    };
  },
  async created() {
    await this.fetchSellerDetails();
  },
  watch: {
    // Watch for route changes if you allow navigating between seller detail pages directly
    '$route.params.id': 'fetchSellerDetails'
  },
  methods: {
    async fetchSellerDetails() {
      this.loading = true;
      this.error = null;
      this.seller = null; // Reset seller data before fetching
      const sellerId = this.$route.params.id;

      if (!sellerId) {
          this.error = "No seller ID provided in the URL.";
          this.loading = false;
          console.error(this.error);
          return;
      }

      try {
        // IMPORTANT: Adjust the API endpoint as needed.
        // It's common to fetch a single item by its ID like /admin/sellers/{id}
        const res = await axios.get(`/admin/sellers/${sellerId}`, {
          withCredentials: true,
        });

        // Assume res.data contains the seller object directly
        if (res.data) {
          this.seller = this.mapSellerData(res.data);
           // Optional: If your API nests the seller like { seller: {...} }, use:
           // this.seller = this.mapSellerData(res.data.seller || res.data);
        } else {
          // This case is less likely with a successful 200, but good to handle
          this.error = `Seller with ID ${sellerId} data is empty.`;
        }

      } catch (err) {
        console.error(`Error fetching seller details for ID ${sellerId}:`, err);
        // Handle specific error responses (e.g., 404)
        if (err.response && err.response.status === 404) {
            this.error = `Seller with ID ${sellerId} not found.`;
        } else {
            this.error = `Failed to fetch seller details. ${err.response?.data?.message || err.message || 'Please try again later.'}`;
        }
         this.seller = null; // Ensure seller is null on error
      } finally {
        this.loading = false;
      }
    },
     // Helper function to map data, handle potential nulls/undefineds/casing
     // Ensure this mapping reflects the actual structure returned by your /admin/sellers/{id} endpoint
     mapSellerData(s) {
        if (!s || typeof s !== 'object') return null; // Basic validation

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
    formatDate(dateString, includeTime = false) {
      if (!dateString) return 'N/A';
      const options = {
        year: 'numeric',
        month: 'short', // Use short month name for consistency
        day: 'numeric',
      };
      if (includeTime) {
        options.hour = '2-digit';
        options.minute = '2-digit';
        // options.second = '2-digit'; // Optional: add seconds
        // options.timeZoneName = 'short'; // Optional: add timezone
      }
      try {
         // Use toLocaleString for both date and time if includeTime is true
         return new Date(dateString).toLocaleString(undefined, options);
      } catch (e) {
          console.error('Error formatting date:', dateString, e);
          return 'Invalid Date'; // Indicate formatting failure
      }
    },
  },
};
</script>

<style scoped>
.container {
  max-width: 900px;
}

</style>