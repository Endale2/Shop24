<template>
  <div class="container mx-auto">
    <div class="mb-6 flex justify-between items-center">
      <h1 class="text-2xl font-semibold text-gray-900 dark:text-white">Customer Management</h1>
      </div>

    <div v-if="loading" class="flex items-center justify-center h-64 text-gray-500 dark:text-gray-400 bg-white dark:bg-gray-800 rounded-lg shadow-md">
       <svg class="animate-spin h-8 w-8 text-blue-600 dark:text-blue-400 mr-3" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 0116 0H4z"></path>
      </svg>
      Loading customer data...
    </div>

    <div v-else-if="error" class="bg-red-100 dark:bg-red-900 border border-red-400 dark:border-red-700 text-red-700 dark:text-red-300 px-4 py-3 rounded-lg shadow-md mb-6" role="alert">
      <strong class="font-bold">Error:</strong>
      <span class="block sm:inline ml-2">{{ error }}</span>
    </div>

    <div v-else>
      <div v-if="customers.length === 0" class="flex flex-col items-center justify-center h-64 text-gray-600 dark:text-gray-400 bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
         <svg class="mx-auto h-12 w-12 text-gray-400 dark:text-gray-500 mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z" />
         </svg>
         <h3 class="mt-2 text-lg font-medium text-gray-900 dark:text-white">No Customers Found</h3>
         <p class="mt-1 text-sm text-gray-500 dark:text-gray-400 text-center">
           There are currently no customers to display.
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
                 Email
               </th>
               <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                 Phone
               </th>
               <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                 Address (Summary)
               </th>
               <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                 Joined
               </th>
               </tr>
           </thead>
           <tbody class="bg-white dark:bg-gray-800 divide-y divide-gray-200 dark:divide-gray-700">
             <tr v-for="customer in customers" :key="customer.id" class="hover:bg-gray-50 dark:hover:bg-gray-700/50 transition-colors duration-150">
               <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900 dark:text-white">
                 <div class="flex items-center">
                    <div class="flex-shrink-0 h-10 w-10">
                      <div class="h-10 w-10 bg-teal-100 dark:bg-teal-900 rounded-full flex items-center justify-center text-teal-700 dark:text-teal-300 text-sm font-semibold">
                        {{ getInitials(customer) }}
                      </div>
                    </div>
                    <div class="ml-4">
                      <router-link :to="`/customers/${customer.id}`" class="text-sm font-medium text-blue-600 hover:text-blue-900 dark:text-blue-400 dark:hover:text-blue-300">
                        {{ getFullName(customer) }}
                      </router-link>
                    </div>
                 </div>
               </td>
               <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700 dark:text-gray-300">
                 {{ customer.email || 'N/A' }}
               </td>
               <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700 dark:text-gray-300">
                 {{ customer.phone || 'N/A' }}
               </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700 dark:text-gray-300">
                 {{ getAddressSummary(customer) }}
               </td>
               <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700 dark:text-gray-300">
                 {{ formatDate(customer.createdAt) }}
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
  name: 'CustomersList',
  data() {
    return {
      customers: [],
      loading: true, // Added loading state
      error: null,   // Added error state
    };
  },
  async created() {
    this.fetchCustomers();
  },
  methods: {
    async fetchCustomers() {
       this.loading = true;
       this.error = null; // Clear previous errors
       try {
         const res = await axios.get('/admin/customers/', { withCredentials: true });

         // Assuming the response data is the array of customers directly or nested
         const apiData = res.data;
         let rawCustomers = [];

         if (Array.isArray(apiData)) {
             rawCustomers = apiData;
         } else if (apiData && Array.isArray(apiData.data)) { // Common API pattern
             rawCustomers = apiData.data;
         } else if (apiData && Array.isArray(apiData.customers)) { // Another common API pattern
             rawCustomers = apiData.customers;
         } else {
            console.warn('Fetch customers: Response data is not in expected array format.', apiData);
         }


         // Normalize potential uppercase keys and handle nulls/undefineds
         this.customers = rawCustomers.map(c => ({
           id:          c.ID || c.id || null,
           firstName:   c.FirstName || c.firstName || null,
           lastName:    c.LastName || c.lastName || null,
           email:       c.Email || c.email || null,
           phone:       c.Phone || c.phone || null,
           address:     c.Address || c.address || null,
           city:        c.City || c.city || null,
           state:       c.State || c.state || null,
           postalCode:  c.PostalCode || c.postalCode || null,
           country:     c.Country || c.country || null,
           createdAt:   c.CreatedAt || c.createdAt || null,
           updatedAt:   c.UpdatedAt || c.updatedAt || null,
         })).filter(c => c.id != null); // Ensure customers have an ID


       } catch (err) {
         console.error('Error fetching customers:', err);
         this.error = `Failed to fetch customers: ${err.response?.data?.message || err.message || 'Unknown error'}`;
         this.customers = []; // Clear list on error
       } finally {
         this.loading = false; // Hide loading indicator
       }
    },
     getFullName(customer) {
        const parts = [customer.firstName, customer.lastName].filter(Boolean); // Filter out null/undefined/empty strings
        return parts.length > 0 ? parts.join(' ') : 'N/A';
     },
     getInitials(customer) {
        let initials = '';
        if (customer.firstName) initials += customer.firstName.charAt(0).toUpperCase();
        if (customer.lastName) initials += customer.lastName.charAt(0).toUpperCase();
        if (initials === '' && customer.email) initials = customer.email.charAt(0).toUpperCase();
        return initials || 'C'; // Default to 'C' if no name/email
     },
     getAddressSummary(customer) {
        const parts = [customer.city, customer.state, customer.country].filter(Boolean);
        const addressSummary = parts.length > 0 ? parts.join(', ') : 'N/A';
        // Optionally add a tooltip for the full address if you want to show it on hover
        return customer.address && addressSummary !== 'N/A' ? `${customer.address}, ${addressSummary}` : addressSummary;
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
  }
};
</script>

