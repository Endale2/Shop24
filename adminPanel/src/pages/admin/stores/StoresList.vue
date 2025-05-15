<template>
  <div class="container mx-auto">
    <div class="mb-6 flex justify-between items-center">
      <h1 class="text-2xl font-semibold text-gray-900 dark:text-white">Stores</h1>
      <router-link
         to="/stores/create"
         class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 dark:focus:ring-offset-gray-800"
       >
         Create New Store
       </router-link>
    </div>

    <div v-if="loading" class="flex items-center justify-center h-64 text-gray-500 dark:text-gray-400">
      <svg class="animate-spin h-8 w-8 text-blue-600 dark:text-blue-400 mr-3" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 0116 0H4z"></path>
      </svg>
      Loading stores...
    </div>

    <div v-else-if="error" class="flex items-center justify-center h-64 text-red-600 dark:text-red-400">
      <p>Error loading stores: {{ error }}</p>
    </div>

    <div v-else>
      <div v-if="stores.length === 0" class="flex items-center justify-center h-64 text-gray-600 dark:text-gray-400">
        <p>No stores available.</p>
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
                 Owner ID
               </th>
               <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                 Description
               </th>
               <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                 Created At
               </th>
               </tr>
           </thead>
           <tbody class="bg-white dark:bg-gray-800 divide-y divide-gray-200 dark:divide-gray-700">
             <tr v-for="s in stores" :key="s.id" class="hover:bg-gray-50 dark:hover:bg-gray-700/50">
               <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900 dark:text-white">
                 {{ s.id || '—' }}
               </td>
               <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-white">
                 <router-link :to="`/stores/${s.id}`" class="text-blue-600 hover:text-blue-900 dark:text-blue-400 dark:hover:text-blue-300">
                    {{ s.name || '—' }}
                 </router-link>
               </td>
               <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 dark:text-gray-400">
                 {{ s.ownerId || '—' }}
               </td>
               <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 dark:text-gray-400">
                 {{ s.description || '—' }}
               </td>
               <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 dark:text-gray-400">
                 {{ formatDate(s.createdAt) }}
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
  name: 'StoresList',
  data() {
    return {
      stores: [],
      loading: true, // Added loading state
      error: null,   // Added error state
      // Removed showConfirm and toDeleteId
    };
  },
  async created() {
    await this.fetchStores();
  },
  methods: {
    async fetchStores() {
       this.loading = true;
       this.error = null; // Clear previous errors
       try {
         const res = await axios.get('/admin/stores/', { withCredentials: true });
         // Map potentially different casing and ensure objects have standard keys
         this.stores = res.data.map(s => ({
           id:          s.ID || s.id || null, // Handle potential nulls
           name:        s.Name || s.name || '',
           ownerId:     s.OwnerID || s.ownerId || null,
           description: s.Description || s.description || '',
           createdAt:   s.CreatedAt || s.createdAt || null,
           updatedAt:   s.UpdatedAt || s.updatedAt || null, // Added updatedAt if available
           // Map other fields if they exist in your API response
           // e.g., status: s.Status || s.status || 'unknown'
         }));
       } catch (e) {
         console.error('Error fetching stores:', e);
         this.error = e.message || 'An error occurred while fetching stores.'; // Set user-friendly error message
         this.stores = []; // Clear stores on error
       } finally {
         this.loading = false; // Hide loading indicator
       }
    },
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
     // Removed confirmDelete and deleteStore methods
  }
};
</script>
