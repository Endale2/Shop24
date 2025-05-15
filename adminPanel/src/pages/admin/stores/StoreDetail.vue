<template>
  <div class="container mx-auto">
    <div class="mb-6">
      <h1 class="text-2xl font-semibold text-gray-900 dark:text-white">
        Store Details: {{ store && store.name ? store.name : 'Loading...' }}
      </h1>
    </div>

    <div v-if="loading" class="flex items-center justify-center h-64 text-gray-500 dark:text-gray-400">
       <svg class="animate-spin h-8 w-8 text-blue-600 dark:text-blue-400 mr-3" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 0116 0H4z"></path>
      </svg>
      Loading store details...
    </div>

    <div v-else-if="error" class="flex items-center justify-center h-64 text-red-600 dark:text-red-400">
      <p>Error loading store: {{ error }}</p>
    </div>

    <div v-else class="bg-white dark:bg-gray-800 shadow-md rounded-lg p-6">

      <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4 border-b border-gray-200 dark:border-gray-700 pb-2">Information</h3>
      <dl class="divide-y divide-gray-200 dark:divide-gray-700 mb-6">
        <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
          <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Store ID</dt>
          <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">{{ store.id || '—' }}</dd>
        </div>
        <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
          <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Owner ID</dt>
          <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">{{ store.ownerId || '—' }}</dd>
        </div>
         <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
          <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Created At</dt>
          <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">{{ formatDate(store.createdAt) }}</dd>
        </div>
         <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
          <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Updated At</dt>
          <dd class="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">{{ formatDate(store.updatedAt) }}</dd>
        </div>
      </dl>

      <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4 border-b border-gray-200 dark:border-gray-700 pb-2">Edit Details</h3>
      <form @submit.prevent="updateStore" class="space-y-6">
        <div>
          <label for="store-name" class="block text-sm font-medium text-gray-700 dark:text-gray-300">Name</label>
          <input
            id="store-name"
            v-model="form.name"
            type="text"
            class="mt-1 block w-full border border-gray-300 dark:border-gray-600 rounded-md shadow-sm p-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:outline-none focus:ring-blue-500 focus:border-blue-500 dark:focus:ring-blue-500 dark:focus:border-blue-500 sm:text-sm"
            required
          />
        </div>

        <div>
          <label for="store-description" class="block text-sm font-medium text-gray-700 dark:text-gray-300">Description</label>
          <textarea
            id="store-description"
            v-model="form.description"
            class="mt-1 block w-full border border-gray-300 dark:border-gray-600 rounded-md shadow-sm p-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:outline-none focus:ring-blue-500 focus:border-blue-500 dark:focus:ring-blue-500 dark:focus:border-blue-500 sm:text-sm"
            rows="3"
          ></textarea>
        </div>

         <div v-if="updateSuccess" class="text-green-600 dark:text-green-400 text-sm">Store updated successfully!</div>
        <div v-if="updateError" class="text-red-600 dark:text-red-400 text-sm">Failed to update store: {{ updateError }}</div>


        <div class="flex justify-end space-x-3 mt-6">
          <button
            type="submit"
            class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 dark:focus:ring-offset-gray-800 disabled:opacity-50 disabled:cursor-not-allowed"
            :disabled="updating"
          >
             <svg v-if="updating" class="animate-spin -ml-1 mr-2 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
               <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
               <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 0116 0H4z"></path>
             </svg>
            {{ updating ? 'Saving...' : 'Save Changes' }}
          </button>

          <button
            type="button"
            @click="confirmDelete = true"
            class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 dark:focus:ring-offset-gray-800 disabled:opacity-50 disabled:cursor-not-allowed"
            :disabled="deleting || updating"
          >
            Delete Store
          </button>
        </div>
      </form>
    </div>

     <div class="mt-6">
       <router-link to="/stores" class="inline-flex items-center px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm text-sm font-medium text-gray-700 dark:text-gray-300 bg-white dark:bg-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 dark:focus:ring-offset-gray-800">
         Back to Stores
       </router-link>
     </div>

    <div
      v-if="confirmDelete"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50"
      @click.self="confirmDelete = false" >
      <div class="bg-white dark:bg-gray-800 rounded-lg p-6 space-y-6 w-full max-w-sm shadow-xl">
        <h3 class="text-lg font-semibold text-gray-900 dark:text-white">Confirm Delete</h3>
        <p class="text-gray-600 dark:text-gray-300">Are you sure you want to delete the store "{{ store.name || 'Unnamed Store' }}"? This action cannot be undone.</p>

        <div v-if="deleteError" class="text-red-600 dark:text-red-400 text-sm">{{ deleteError }}</div>

        <div class="flex justify-end space-x-3">
          <button
            @click="confirmDelete = false"
            class="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm text-sm font-medium text-gray-700 dark:text-gray-300 bg-white dark:bg-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 dark:focus:ring-offset-gray-800 disabled:opacity-50 disabled:cursor-not-allowed"
             :disabled="deleting"
          >
            Cancel
          </button>
          <button
            @click="deleteStore"
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
  name: 'StoreDetail',
  // Assuming you're using route props: true in your router config
  // props: ['id'], // If using route props: true

  data() {
    return {
      store: null, // Changed to null initially for proper not found state
      form: { name: '', description: '' },
      loading: true,
      error: null, // Error for initial fetch
      updating: false, // State for update operation
      deleting: false, // State for delete operation
      confirmDelete: false, // State for delete modal visibility
      updateSuccess: false, // State for showing update success message
      updateError: null, // Error for update operation
      deleteError: null, // Error for delete operation
    };
  },
  async created() {
    // Get ID from route params if not using props
    const storeId = this.$route.params.id;
    if (!storeId) {
        this.error = "Store ID is missing from the route.";
        this.loading = false;
        console.error(this.error);
        return; // Stop execution if ID is missing
    }

    try {
      const res = await axios.get(`/admin/stores/${storeId}`, { withCredentials: true });
      const d = res.data;

      // Check if data was actually returned, handle cases like 200 with empty body/object
      if (!d) {
         this.error = "Store not found.";
         this.store = null; // Ensure store is null
         return; // Stop execution
      }

      // normalize API response, handle potential nulls/undefineds
      this.store = {
        id: d.ID || d.id || null,
        ownerId: d.OwnerID || d.ownerId || null,
        name: d.Name || d.name || '', // Use empty string for editable fields if null/undefined
        description: d.Description || d.description || '',
        createdAt: d.CreatedAt || d.createdAt || null,
        updatedAt: d.UpdatedAt || d.updatedAt || null,
        // Map other fields if they exist in your API response
      };
      // Initialize form with current store data
      this.form.name = this.store.name;
      this.form.description = this.store.description;

    } catch (e) {
      console.error('Error fetching store:', e);
      // Handle specific error responses (e.g., 404)
      if (e.response && e.response.status === 404) {
          this.error = "Store not found.";
      } else {
          this.error = e.message || 'An error occurred while fetching store details.';
      }
       this.store = null; // Ensure store is null on error
    } finally {
      this.loading = false;
    }
  },
  methods: {
     formatDate(dateStr) {
       if (!dateStr) return '—';
        try {
            // Use consistent formatting options
            const options = { year: 'numeric', month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' };
            return new Date(dateStr).toLocaleString(undefined, options);
        } catch (e) {
            console.error('Error formatting date:', dateStr, e);
            return 'Invalid Date';
        }
     },
     async updateStore() {
       this.updating = true;
       this.updateError = null; // Clear previous errors
       this.updateSuccess = false; // Reset success message

       // Basic validation
       if (!this.form.name) {
           this.updateError = "Name is required.";
           this.updating = false;
           return;
       }

       try {
         // Ensure you send the correct payload structure expected by your backend (e.g., snake_case vs camelCase)
         // Based on original code, assuming camelCase for the form data payload
         const payload = {
             name: this.form.name,
             description: this.form.description
         };

         // Assuming your update endpoint is PATCH /admin/stores/{id}
         const res = await axios.patch(
           `/admin/stores/${this.store.id}`,
           payload,
           { withCredentials: true } // Add withCredentials if needed
         );

         // Handle successful update response (e.g., API might return the updated store object)
         // If the API returns the updated store, you might update `this.store` with `res.data`
         // For now, just reflect the form changes locally and update the 'Updated At' time
         this.store.name = this.form.name;
         this.store.description = this.form.description;
         // If your API returns the updated object including new updatedAt:
         // this.store.updatedAt = res.data.UpdatedAt || res.data.updatedAt;
         // Otherwise, you might need to re-fetch the store or manually update the time if possible
         // For simplicity, let's just show success message
         this.updateSuccess = true;
         // Optionally hide success message after a few seconds
         setTimeout(() => this.updateSuccess = false, 5000);


       } catch (e) {
         console.error('Failed to update store:', e);
         this.updateError = e.response?.data?.message || e.message || 'An error occurred during update.';
          // Optionally show a more detailed error from backend response
       } finally {
         this.updating = false;
       }
     },
     async deleteStore() {
       this.deleting = true;
       this.deleteError = null; // Clear previous errors

       // Ensure store and store.id are available
       if (!this.store || !this.store.id) {
            this.deleteError = "Cannot delete store: ID is missing.";
            this.deleting = false;
            return;
       }


       try {
         // Assuming your delete endpoint is DELETE /admin/stores/{id}
         await axios.delete(`/admin/stores/${this.store.id}`, { withCredentials: true });

         console.log(`Store ${this.store.id} deleted.`);
         this.confirmDelete = false; // Close modal
         // Redirect to the stores list after deletion
         this.$router.push('/stores');
         // Optionally show a success notification on the next page

       } catch (e) {
         console.error('Error deleting store:', e);
         this.deleteError = e.response?.data?.message || e.message || 'Failed to delete store.';
          // Keep modal open to show the error
       } finally {
         this.deleting = false;
       }
     }
  }
};
</script>

<style scoped>
/* No specific styles needed if using Tailwind for everything */
</style>