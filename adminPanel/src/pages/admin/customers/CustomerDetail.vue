<template>
  <div class="container mx-auto">
    <div class="mb-6 flex justify-between items-center">
      <h1 class="text-2xl font-semibold text-gray-900 dark:text-white">
        Customer Details: {{ customer ? getFullName(customer) : '...' }}
      </h1>
      <div v-if="customer && !loading">
        <router-link
          :to="`/customers/${customer.id}/edit`"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 dark:focus:ring-offset-gray-800 mr-3"
        >
          Edit
        </router-link>
        <button
          @click="confirmDelete = true"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 dark:focus:ring-offset-gray-800"
        >
          Delete
        </button>
      </div>
    </div>

    <div v-if="loading" class="flex items-center justify-center h-64 text-gray-500 dark:text-gray-400 bg-white dark:bg-gray-800 rounded-lg shadow-md">
       <svg class="animate-spin h-8 w-8 text-blue-600 dark:text-blue-400 mr-3" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 0116 0H4z"></path>
      </svg>
      Loading customer details...
    </div>

    <div v-else-if="!customer" class="flex items-center justify-center h-64 text-red-600 dark:text-red-400 bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
       <p v-if="error">Error loading customer: {{ error }}</p>
       <p v-else>Customer not found.</p>
    </div>

    <div v-else class="bg-white dark:bg-gray-800 shadow-md rounded-lg p-6">

      <dl class="divide-y divide-gray-200 dark:divide-gray-700">
        <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
          <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Customer ID</dt>
          <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">{{ customer.id || '—' }}</dd>
        </div>
         <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
          <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Full Name</dt>
          <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">{{ getFullName(customer) }}</dd>
        </div>
         <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
          <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Email</dt>
          <dd class="mt-1 text-sm text-blue-600 dark:text-blue-400 hover:underline sm:mt-0 sm:col-span-2">
             <a :href="`mailto:${customer.email}`">{{ customer.email || '—' }}</a>
          </dd>
        </div>
         <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
          <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Phone</dt>
          <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
             {{ customer.phone || '—' }}
              <span v-if="customer.phone" class="ml-2 text-blue-600 dark:text-blue-400 hover:underline cursor-pointer" @click="callPhone(customer.phone)">(Call)</span>
          </dd>
        </div>
         <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
          <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Address</dt>
          <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">{{ getFullAddress(customer) }}</dd>
        </div>
         <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
          <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Created At</dt>
          <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">{{ formatDate(customer.createdAt, true) }}</dd>
        </div>
          <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
          <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Updated At</dt>
          <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">{{ formatDate(customer.updatedAt, true) }}</dd>
        </div>
      </dl>

    </div>

    <div class="mt-6">
       <router-link to="/customers" class="inline-flex items-center px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm text-sm font-medium text-gray-700 dark:text-gray-300 bg-white dark:bg-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 dark:focus:ring-offset-gray-800">
         Back to Customers List
       </router-link>
     </div>

    <div
      v-if="confirmDelete"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50"
      @click.self="confirmDelete = false" >
      <div class="bg-white dark:bg-gray-800 rounded-lg p-6 space-y-6 w-full max-w-sm shadow-xl">
        <h3 class="text-lg font-semibold text-gray-900 dark:text-white">Confirm Delete</h3>
        <p class="text-gray-600 dark:text-gray-300">
           Are you sure you want to delete customer "{{ customer ? getFullName(customer) : '' }}"? This action cannot be undone.
        </p>

        <div v-if="deleteError" class="text-red-600 dark:text-red-400 text-sm">{{ deleteError }}</div>

        <div class="flex justify-end space-x-3">
          <button
            @click="confirmDelete = false; deleteError = null;"
            class="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm text-sm font-medium text-gray-700 dark:text-gray-300 bg-white dark:bg-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 dark:focus:ring-offset-gray-800 disabled:opacity-50 disabled:cursor-not-allowed"
            :disabled="deleting"
          >
            Cancel
          </button>
          <button
            @click="deleteCustomer"
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
  name: 'CustomerDetail',
  // Assuming you're using route props: true in your router config
  // props: ['id'], // If using route props: true

  data() {
    return {
      customer: null,
      loading: true,
      error: null, // Error for initial fetch
      confirmDelete: false, // State for delete modal visibility
      deleting: false, // State for delete operation
      deleteError: null, // Error message for delete operation
    };
  },
  async created() {
    await this.fetchCustomerDetails();
  },
  watch: {
    // Watch for route changes if you allow navigating between customer detail pages directly
    '$route.params.id': 'fetchCustomerDetails'
  },
  methods: {
    async fetchCustomerDetails() {
      this.loading = true;
      this.error = null;
      this.customer = null; // Reset customer data before fetching
      const customerId = this.$route.params.id;

      if (!customerId) {
          this.error = "No customer ID provided in the URL.";
          this.loading = false;
          console.error(this.error);
          return;
      }

      try {
        // IMPORTANT: Adjust the API endpoint as needed.
        // It's common to fetch a single item by its ID like /admin/customers/{id}
        const res = await axios.get(`/admin/customers/${customerId}`, {
          withCredentials: true,
        });

        // Assume res.data contains the customer object directly
        if (res.data) {
           // Normalize potential uppercase keys
           this.customer = {
             id:          res.data.ID || res.data.id || null,
             firstName:   res.data.FirstName || res.data.firstName || null,
             lastName:    res.data.LastName || res.data.lastName || null,
             email:       res.data.Email || res.data.email || null,
             phone:       res.data.Phone || res.data.phone || null,
             address:     res.data.Address || res.data.address || null,
             city:        res.data.City || res.data.city || null,
             state:       res.data.State || res.data.state || null,
             postalCode:  res.data.PostalCode || res.data.postalCode || null,
             country:     res.data.Country || res.data.country || null,
             createdAt:   res.data.CreatedAt || res.data.createdAt || null,
             updatedAt:   res.data.UpdatedAt || res.data.updatedAt || null,
           };
        } else {
          // This case is less likely with a successful 200, but good to handle
          this.error = `Customer with ID ${customerId} data is empty.`;
        }

      } catch (err) {
        console.error(`Error fetching customer details for ID ${customerId}:`, err);
        // Handle specific error responses (e.g., 404)
        if (err.response && err.response.status === 404) {
            this.error = `Customer with ID ${customerId} not found.`;
        } else {
            this.error = `Failed to fetch customer details. ${err.response?.data?.message || err.message || 'Please try again later.'}`;
        }
         this.customer = null; // Ensure customer is null on error
      } finally {
        this.loading = false;
      }
    },
     getFullName(customer) {
        if (!customer) return '—';
        const parts = [customer.firstName, customer.lastName].filter(Boolean); // Filter out null/undefined/empty strings
        return parts.length > 0 ? parts.join(' ') : 'Unnamed Customer'; // Provide a default if both are missing
     },
     getFullAddress(customer) {
        if (!customer) return '—';
        const parts = [
            customer.address,
            customer.city,
            customer.state,
            customer.postalCode,
            customer.country
        ].filter(Boolean); // Filter out null/undefined/empty strings
        return parts.length > 0 ? parts.join(', ') : 'N/A';
     },
    formatDate(dateString, includeTime = false) {
      if (!dateString) return '—';
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
     callPhone(phoneNumber) {
        if (phoneNumber) {
            window.location.href = `tel:${phoneNumber}`;
        }
     },
     async deleteCustomer() {
       // Ensure customer and customer.id are available
       if (!this.customer || !this.customer.id) {
            this.deleteError = "Cannot delete customer: ID is missing.";
            console.error(this.deleteError, this.customer);
            this.deleting = false; // Ensure deleting state is false
            return;
       }

       this.deleting = true;
       this.deleteError = null; // Clear previous errors

       try {
         // Assuming your delete endpoint is DELETE /admin/customers/{id}
         await axios.delete(`/admin/customers/${this.customer.id}`, { withCredentials: true });

         console.log(`Customer ID ${this.customer.id} deleted.`);
         this.confirmDelete = false; // Close modal
         // Redirect to the customers list after deletion
         this.$router.push('/customers');
         // Optionally show a success notification on the next page

       } catch (err) {
         console.error('Error deleting customer:', err);
         this.deleteError = err.response?.data?.message || err.message || 'Failed to delete customer.';
         // Keep modal open to show the error: this.showDeleteModal = true;
         // Don't clear customerToDelete yet so the name remains in the error message
       } finally {
         this.deleting = false;
       }
     }
  }
};
</script>

