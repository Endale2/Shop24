<template>
  <div class="p-6 max-w-7xl mx-auto bg-white rounded-lg shadow-md mt-6">
    <h2 class="text-3xl font-extrabold mb-6 text-gray-800 border-b pb-4">Customers Overview</h2>

    <div v-if="loading" class="flex items-center justify-center py-12 text-gray-600">
      <svg class="animate-spin h-8 w-8 text-green-500 mr-3" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      Loading customers...
    </div>

    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative" role="alert">
      <strong class="font-bold">Error!</strong>
      <span class="block sm:inline ml-2">{{ error }}</span>
    </div>

    <div v-else-if="customers.length === 0" class="bg-blue-50 border border-blue-200 text-blue-700 px-4 py-6 rounded-md text-center">
      <p class="text-lg font-medium">No customers found for this shop.</p>
      <p class="text-md mt-2">Start processing orders to see your customer data here.</p>
    </div>

    <div v-else class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              ID
            </th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              First Name
            </th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Last Name
            </th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Email
            </th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Phone
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="c in customers" :key="c.linkId" class="hover:bg-green-50 transition duration-150 ease-in-out">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-mono text-gray-500">{{ c.id }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ c.firstName || '—' }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ c.lastName || '—' }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ c.email || '—' }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ c.phone || '—' }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import api from '@/services/api'; // Ensure this path is correct for your API service

export default {
  name: 'CustomersPage',
  data() {
    return {
      customers: [], // now: { id, firstName, …, linkId }
      loading: false,
      error: null,
    };
  },
  computed: {
    ...mapGetters('shops', ['activeShop'])
  },
  async created() {
    if (!this.activeShop) {
      this.error = 'No shop selected. Please select a shop to view customers.';
      return;
    }
    this.loading = true;
    try {
      // Assuming api.get returns a response object with a .data property
      const res = await api.get(`/seller/shops/${this.activeShop.id}/customers`);

      // Based on your provided structure:
      // res.data is an array like: [ { customer: {...}, linkId: "..." }, ... ]
      this.customers = res.data.map(item => ({
        id: item.customer.id,
        firstName: item.customer.firstName,
        lastName: item.customer.lastName,
        email: item.customer.email,
        phone: item.customer.phone || '', // Ensure it's an empty string for display or use '—' in template
        linkId: item.linkId
      }));

    } catch (err) {
      console.error("Error fetching customers:", err); // Log the full error for debugging
      this.error = 'Failed to load customers. Please try again later.';
    } finally {
      this.loading = false;
    }
  }
};
</script>

<style scoped>
/* No custom CSS needed thanks to comprehensive Tailwind classes */
/* Keeping .break-all for old reference, but whitespace-nowrap is generally preferred for tables */
/* If you need breaking for very long IDs, you can apply 'break-all' on the specific td */

/* If you need custom spinner animation: */
/* @keyframes spin {
  from { transform: rotate(0deg); }
  to   { transform: rotate(360deg); }
}
.animate-spin {
  animation: spin 1s linear infinite;
} */
</style>