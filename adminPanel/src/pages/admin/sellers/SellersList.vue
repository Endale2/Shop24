<template>
  <div class="container mx-auto p-4 sm:p-6 lg:p-8">
    <div class="mb-6 flex flex-col sm:flex-row justify-between items-center">
      <h1 class="text-2xl sm:text-3xl font-bold text-gray-800">Seller Management</h1>
      </div>

    <div v-if="loading" class="text-center py-12 bg-white rounded-lg shadow">
      <p class="text-lg text-gray-600">Loading seller data...</p>
      </div>

    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-lg shadow mb-6" role="alert">
      <strong class="font-bold">Error:</strong>
      <span class="block sm:inline">{{ error }}</span>
    </div>

    <div v-if="!loading && sellers.length > 0" class="bg-white shadow-xl rounded-lg overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-100">
          <tr>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-600 uppercase tracking-wider">
              Name
            </th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-600 uppercase tracking-wider">
              Business Name
            </th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-600 uppercase tracking-wider">
              Email
            </th>
            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">
              Ratings
            </th>
            <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-600 uppercase tracking-wider">
              Total Sales
            </th>
            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">
              Reviews
            </th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-600 uppercase tracking-wider">
              Joined
            </th>
            <th scope="col" class="relative px-6 py-3">
              <span class="sr-only">Actions</span>
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="seller in sellers" :key="seller.id" class="hover:bg-gray-50 transition-colors duration-150">
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <div class="flex-shrink-0 h-10 w-10">
                  <div class="h-10 w-10 bg-indigo-200 rounded-full flex items-center justify-center text-indigo-700 font-semibold">
                    {{ seller.name ? seller.name.charAt(0).toUpperCase() : 'S' }}
                  </div>
                </div>
                <div class="ml-4">
                  <router-link :to="`/sellers/${seller.id}`" class="text-sm font-medium text-indigo-600 hover:text-indigo-900 hover:underline">
                    {{ seller.name || 'N/A' }}
                  </router-link>
                </div>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">{{ seller.businessName || 'N/A' }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-700">{{ seller.email || 'N/A' }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-center">
              <span
                :class="[
                  'px-2 inline-flex text-xs leading-5 font-semibold rounded-full',
                  getRatingBadgeClass(seller.ratings)
                ]"
              >
                {{ seller.ratings !== null ? seller.ratings.toFixed(1) : 'N/A' }} ⭐
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700 text-right">
              ${{ seller.totalSales !== null ? seller.totalSales.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 2 }) : '0.00' }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700 text-center">
              {{ seller.reviewsCount !== null ? seller.reviewsCount : '0' }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700">
              {{ formatDate(seller.createdAt) }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
              <router-link :to="`/sellers/${seller.id}/edit`" class="text-indigo-600 hover:text-indigo-900 mr-3">Edit</router-link>
              <button @click="confirmDelete(seller)" class="text-red-600 hover:text-red-900">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="!loading && sellers.length === 0 && !error" class="text-center py-12 bg-white rounded-lg shadow">
        <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
        </svg>
      <h3 class="mt-2 text-lg font-medium text-gray-900">No Sellers Found</h3>
      <p class="mt-1 text-sm text-gray-500">
        There are currently no sellers to display. You can add a new seller using the button above.
      </p>
    </div>

    <div v-if="showDeleteModal" class="fixed z-10 inset-0 overflow-y-auto" aria-labelledby="modal-title" role="dialog" aria-modal="true">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" aria-hidden="true"></div>
        <span class="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">&#8203;</span>
        <div class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full">
          <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <div class="sm:flex sm:items-start">
              <div class="mx-auto flex-shrink-0 flex items-center justify-center h-12 w-12 rounded-full bg-red-100 sm:mx-0 sm:h-10 sm:w-10">
                <svg class="h-6 w-6 text-red-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                </svg>
              </div>
              <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                <h3 class="text-lg leading-6 font-medium text-gray-900" id="modal-title">
                  Delete Seller
                </h3>
                <div class="mt-2">
                  <p class="text-sm text-gray-500">
                    Are you sure you want to delete seller "{{ sellerToDelete ? sellerToDelete.name : '' }}"? This action cannot be undone.
                  </p>
                </div>
              </div>
            </div>
          </div>
          <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
            <button @click="deleteSellerConfirmed" type="button" class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-red-600 text-base font-medium text-white hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 sm:ml-3 sm:w-auto sm:text-sm">
              Delete
            </button>
            <button @click="showDeleteModal = false; sellerToDelete = null;" type="button" class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm">
              Cancel
            </button>
          </div>
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
      sellerToDelete: null,
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
        // Simulate API delay
        // await new Promise(resolve => setTimeout(resolve, 1000));

        const res = await axios.get('/admin/sellers/', { // Ensure this endpoint is correct
          withCredentials: true
        });

        if (Array.isArray(res.data)) {
            this.sellers = res.data.map(s => this.mapSellerData(s));
        } else if (res.data && Array.isArray(res.data.data)) { // Common API pattern
            this.sellers = res.data.data.map(s => this.mapSellerData(s));
        } else if (res.data && Array.isArray(res.data.sellers)) { // Another common API pattern
            this.sellers = res.data.sellers.map(s => this.mapSellerData(s));
        }
        else {
          console.error('Error fetching sellers: Response data is not in expected array format.', res.data);
          this.sellers = [];
          this.error = 'Could not parse seller data. Unexpected format received.';
        }
      } catch (err) {
        console.error('Error fetching sellers:', err);
        this.error = `Failed to fetch sellers: ${err.response?.data?.message || err.message || 'Unknown error'}`;
        this.sellers = [];
      } finally {
        this.loading = false;
      }
    },
    mapSellerData(s) { // Helper function to map data, keeps created() cleaner
        return {
            id: s.id || s.ID,
            user_id: s.user_id || s.UserId,
            shop_ids: s.shop_ids || s.ShopIds || [],
            name: s.name || s.Name,
            email: s.email || s.Email,
            phone: s.phone || s.Phone, // Not used in table, but good to keep
            address: s.address || s.Address, // Not used in table
            businessName: s.business_name || s.BusinessName,
            ratings: s.ratings !== undefined && s.ratings !== null ? parseFloat(s.ratings) : null,
            totalSales: s.total_sales !== undefined && s.total_sales !== null ? parseFloat(s.total_sales) : null,
            reviewsCount: s.reviews_count !== undefined && s.reviews_count !== null ? parseInt(s.reviews_count, 10) : null,
            createdAt: s.created_at || s.CreatedAt,
            updatedAt: s.updated_at || s.UpdatedAt // Not used in table
        };
    },
    formatDate(dateString) {
      if (!dateString) return 'N/A';
      const options = { year: 'numeric', month: 'short', day: 'numeric' };
      try {
        return new Date(dateString).toLocaleDateString(undefined, options);
      } catch (e) {
        return dateString;
      }
    },
    getRatingBadgeClass(rating) {
      if (rating === null || rating === undefined) return 'bg-gray-200 text-gray-800';
      if (rating >= 4.5) return 'bg-green-100 text-green-800';
      if (rating >= 4.0) return 'bg-yellow-100 text-yellow-800';
      if (rating >= 3.0) return 'bg-orange-100 text-orange-800';
      return 'bg-red-100 text-red-800';
    },
    confirmDelete(seller) {
      this.sellerToDelete = seller;
      this.showDeleteModal = true;
    },
    async deleteSellerConfirmed() {
      if (!this.sellerToDelete) return;
      try {
        // Replace with your actual delete API call
        // await axios.delete(`/admin/sellers/${this.sellerToDelete.id}`, { withCredentials: true });

        console.log(`Simulating delete for seller ID: ${this.sellerToDelete.id}`);
        // Remove from local list after successful deletion
        this.sellers = this.sellers.filter(s => s.id !== this.sellerToDelete.id);
        
        this.showDeleteModal = false;
        this.sellerToDelete = null;
        // Optionally, show a success notification here
      } catch (err) {
        console.error('Error deleting seller:', err);
        this.error = `Failed to delete seller: ${err.response?.data?.message || err.message || 'Unknown error'}`;
        // Optionally, show an error notification here
        this.showDeleteModal = false; // Close modal even on error, or handle differently
      }
    }
  }
};
</script>

<style scoped>

.min-w-full { 
  min-width: 800px; 
}
</style>