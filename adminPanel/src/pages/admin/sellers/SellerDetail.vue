<template>
  <div class="container mx-auto p-4 sm:p-6 lg:p-8">
    <div v-if="loading" class="text-center py-20">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600 mx-auto"></div>
      <p class="mt-4 text-lg text-gray-600">Loading seller details...</p>
    </div>

    <div v-else-if="error" class="bg-red-100 border-l-4 border-red-500 text-red-700 p-6 rounded-md shadow-md" role="alert">
      <h3 class="font-bold text-xl mb-2">Error Fetching Seller</h3>
      <p>{{ error }}</p>
      <router-link to="/sellers" class="mt-4 inline-block bg-red-500 hover:bg-red-600 text-white font-semibold py-2 px-4 rounded transition-colors duration-150">
        Back to Sellers List
      </router-link>
    </div>

    <div v-else-if="seller" class="bg-white shadow-xl rounded-lg overflow-hidden">
      <div class="bg-gradient-to-r from-indigo-600 to-purple-600 p-6 sm:p-8 text-white">
        <div class="flex flex-col sm:flex-row items-center justify-between">
          <div>
            <h1 class="text-3xl sm:text-4xl font-bold mb-1">{{ seller.name || 'Seller Details' }}</h1>
            <p class="text-lg text-indigo-200">{{ seller.businessName || 'N/A' }}</p>
          </div>
          <div class="mt-4 sm:mt-0">
            <router-link
              :to="`/sellers/${seller.id}/edit`"
              class="bg-white text-indigo-600 hover:bg-indigo-50 font-semibold py-2 px-5 rounded-lg shadow transition-all duration-150 ease-in-out transform hover:scale-105"
            >
              Edit Seller
            </router-link>
          </div>
        </div>
        <div class="mt-6 grid grid-cols-1 sm:grid-cols-3 gap-4 text-center">
          <div>
            <p class="text-sm text-indigo-200 uppercase tracking-wider">Rating</p>
            <p class="text-2xl font-semibold">
              {{ seller.ratings !== null ? seller.ratings.toFixed(1) : 'N/A' }} ⭐
            </p>
          </div>
          <div>
            <p class="text-sm text-indigo-200 uppercase tracking-wider">Total Sales</p>
            <p class="text-2xl font-semibold">
              ${{ seller.totalSales !== null ? seller.totalSales.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 2 }) : '0.00' }}
            </p>
          </div>
          <div>
            <p class="text-sm text-indigo-200 uppercase tracking-wider">Reviews</p>
            <p class="text-2xl font-semibold">
              {{ seller.reviewsCount !== null ? seller.reviewsCount.toLocaleString() : '0' }}
            </p>
          </div>
        </div>
      </div>

      <div class="p-6 sm:p-8 space-y-8">
        <section>
          <h2 class="text-xl font-semibold text-gray-700 mb-4 border-b pb-2">Contact Information</h2>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-x-8 gap-y-4 text-sm">
            <div class="flex items-center">
              <span class="w-6 h-6 mr-3 text-indigo-500">📧</span> <div>
                <p class="text-gray-500">Email Address</p>
                <p class="text-gray-800 font-medium">{{ seller.email || 'N/A' }}</p>
              </div>
            </div>
            <div class="flex items-center">
              <span class="w-6 h-6 mr-3 text-indigo-500">📞</span> <div>
                <p class="text-gray-500">Phone Number</p>
                <p class="text-gray-800 font-medium">{{ seller.phone || 'N/A' }}</p>
              </div>
            </div>
            <div class="flex items-start md:col-span-2">
              <span class="w-6 h-6 mr-3 text-indigo-500 mt-1">📍</span> <div>
                <p class="text-gray-500">Full Address</p>
                <p class="text-gray-800 font-medium">{{ seller.address || 'N/A' }}</p>
              </div>
            </div>
          </div>
        </section>

        <section>
          <h2 class="text-xl font-semibold text-gray-700 mb-4 border-b pb-2">Account & Shop Details</h2>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-x-8 gap-y-4 text-sm">
            <div>
              <p class="text-gray-500">Seller ID</p>
              <p class="text-gray-800 font-mono text-xs bg-gray-100 px-2 py-1 rounded inline-block">{{ seller.id }}</p>
            </div>
            <div>
              <p class="text-gray-500">User ID (Associated)</p>
              <p class="text-gray-800 font-mono text-xs bg-gray-100 px-2 py-1 rounded inline-block">{{ seller.user_id || 'N/A' }}</p>
            </div>
            <div class="md:col-span-2">
              <p class="text-gray-500 mb-1">Associated Shop IDs</p>
              <div v-if="seller.shop_ids && seller.shop_ids.length > 0" class="flex flex-wrap gap-2">
                <span v-for="shopId in seller.shop_ids" :key="shopId" class="bg-purple-100 text-purple-700 text-xs font-mono px-3 py-1 rounded-full">
                  {{ shopId }}
                </span>
              </div>
              <p v-else class="text-gray-700 italic">No associated shops.</p>
            </div>
          </div>
        </section>

        <section>
          <h2 class="text-xl font-semibold text-gray-700 mb-4 border-b pb-2">Timestamps</h2>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-x-8 gap-y-4 text-sm">
            <div>
              <p class="text-gray-500">Account Created</p>
              <p class="text-gray-800 font-medium">{{ formatDate(seller.createdAt, true) || 'N/A' }}</p>
            </div>
            <div>
              <p class="text-gray-500">Last Updated</p>
              <p class="text-gray-800 font-medium">{{ formatDate(seller.updatedAt, true) || 'N/A' }}</p>
            </div>
          </div>
        </section>
      </div>

      <div class="px-6 py-4 bg-gray-50 border-t text-right">
        <router-link to="/sellers" class="text-indigo-600 hover:text-indigo-800 hover:underline font-medium transition-colors duration-150">
          &larr; Back to Sellers List
        </router-link>
      </div>
    </div>

    <div v-else-if="!loading && !error" class="text-center py-20">
      <h2 class="text-2xl font-semibold text-gray-700">Seller Not Found</h2>
      <p class="text-gray-500 mt-2">The seller you are looking for does not exist or could not be loaded.</p>
      <router-link to="/sellers" class="mt-6 inline-block bg-indigo-600 hover:bg-indigo-700 text-white font-semibold py-2 px-4 rounded transition-colors duration-150">
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
      error: null,
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
          return;
      }

      try {
        // Simulate API delay
        // await new Promise(resolve => setTimeout(resolve, 1000));

        // IMPORTANT: Adjust the API endpoint as needed.
        // It's common to fetch a single item by its ID like /admin/sellers/{id}
        const res = await axios.get(`/admin/sellers/${sellerId}`, {
          withCredentials: true,
        });

        if (res.data) {
          // If the API returns the seller object directly (e.g. res.data is the seller)
          this.seller = this.mapSellerData(res.data.seller || res.data); // Handle potential nesting like { "seller": {...} }
        } else {
          this.error = `Seller with ID ${sellerId} not found or API response was empty.`;
        }
      } catch (err) {
        console.error(`Error fetching seller details for ID ${sellerId}:`, err);
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
    mapSellerData(s) {
      if (!s || typeof s !== 'object') return null; // Basic validation
      return {
        id: s.id || s.ID,
        user_id: s.user_id || s.UserId,
        shop_ids: s.shop_ids || s.ShopIds || [],
        name: s.name || s.Name,
        email: s.email || s.Email,
        phone: s.phone || s.Phone,
        address: s.address || s.Address,
        businessName: s.business_name || s.BusinessName,
        ratings: s.ratings !== undefined && s.ratings !== null ? parseFloat(s.ratings) : null,
        totalSales: s.total_sales !== undefined && s.total_sales !== null ? parseFloat(s.total_sales) : null,
        reviewsCount: s.reviews_count !== undefined && s.reviews_count !== null ? parseInt(s.reviews_count, 10) : null,
        createdAt: s.created_at || s.CreatedAt,
        updatedAt: s.updated_at || s.UpdatedAt,
      };
    },
    formatDate(dateString, includeTime = false) {
      if (!dateString) return 'N/A';
      const options = {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
      };
      if (includeTime) {
        options.hour = '2-digit';
        options.minute = '2-digit';
      }
      try {
        return new Date(dateString).toLocaleDateString(undefined, options);
      } catch (e) {
        return dateString; // Fallback
      }
    },
  },
};
</script>

<style scoped>
/* Add component-specific styles if necessary */
.container {
  max-width: 900px; /* Adjust max-width as needed */
}
/* Simple spinner animation */
@keyframes spin {
  to { transform: rotate(360deg); }
}
.animate-spin {
  animation: spin 1s linear infinite;
}
</style>