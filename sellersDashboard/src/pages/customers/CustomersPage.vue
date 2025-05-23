<template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto space-y-8 font-sans">
    <div class="flex flex-col md:flex-row md:items-center md:justify-between space-y-4 md:space-y-0">
      <h2 class="text-3xl sm:text-4xl font-extrabold text-gray-900 leading-tight">Customers Overview</h2>

      <div class="flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-4 w-full md:w-auto">
        <input
          v-model="search"
          type="text"
          placeholder="Search by name or email..."
          class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 outline-none transition duration-150 ease-in-out shadow-sm"
        />

        <div class="inline-flex rounded-md shadow-sm" role="group">
          <button
            @click="currentView = 'cards'"
            :class="{
              'bg-blue-600 text-white hover:bg-blue-700': currentView === 'cards',
              'bg-gray-200 text-gray-700 hover:bg-gray-300': currentView !== 'cards'
            }"
            class="px-4 py-2 text-sm font-medium border border-gray-300 rounded-l-lg focus:z-10 focus:ring-2 focus:ring-blue-500 focus:outline-none transition duration-150 ease-in-out"
          >
            <svg class="w-5 h-5 inline-block mr-1" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path d="M5 3a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2V5a2 2 0 00-2-2H5zM11 3a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2V5a2 2 0 00-2-2h-2zM5 11a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2v-2a2 2 0 00-2-2H5zM11 11a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2v-2a2 2 0 00-2-2h-2z"></path></svg>
            Cards
          </button>
          <button
            @click="currentView = 'list'"
            :class="{
              'bg-blue-600 text-white hover:bg-blue-700': currentView === 'list',
              'bg-gray-200 text-gray-700 hover:bg-gray-300': currentView !== 'list'
            }"
            class="px-4 py-2 text-sm font-medium border border-gray-300 rounded-r-lg focus:z-10 focus:ring-2 focus:ring-blue-500 focus:outline-none transition duration-150 ease-in-out"
          >
            <svg class="w-5 h-5 inline-block mr-1" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M3 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z" clip-rule="evenodd"></path></svg>
            List
          </button>
        </div>
      </div>
    </div>

    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-16">
      <svg class="animate-spin h-10 w-10 text-blue-500 mb-3" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
        <path class="opacity-75" fill="currentColor"
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
      </svg>
      <p class="text-lg">Loading customers...</p>
    </div>

    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-lg shadow-sm" role="alert">
      <p class="font-bold">Oops! Something went wrong:</p>
      <p class="mt-2">{{ error }}</p>
      <p class="text-sm mt-2">Please try again or contact support if the issue persists.</p>
    </div>

    <div v-else-if="filteredCustomers.length === 0" class="bg-blue-50 border border-blue-200 text-blue-700 px-6 py-8 rounded-lg text-center mt-8 shadow-sm">
      <p class="text-lg font-medium">
        No customers found<span v-if="search"> matching "{{ search }}"</span>.
      </p>
      <p class="text-md mt-2" v-if="!search">It looks like there are no customers for your active shop yet.</p>
    </div>

    <div v-else-if="currentView === 'list'" class="overflow-x-auto bg-white shadow-lg rounded-xl">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-100">
          <tr>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Name</th>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Email</th>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Phone</th>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Joined Date</th>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Customer ID</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr
            v-for="(cust, i) in filteredCustomers"
            :key="cust.linkId"
            @click="goToCustomerDetail(cust)"
            class="cursor-pointer transition duration-200 ease-in-out hover:bg-blue-50"
            :class="{ 'bg-gray-50': i % 2 === 1 }"
          >
            <td class="py-3 px-6 text-sm text-gray-800 font-medium">
              <div class="flex items-center">
                <div class="h-8 w-8 bg-blue-100 rounded-full flex items-center justify-center text-blue-700 text-xs font-semibold mr-3 flex-shrink-0">
                  {{ getInitials(cust.firstName, cust.lastName) }}
                </div>
                <span>{{ cust.firstName }} {{ cust.lastName }}</span>
              </div>
            </td>
            <td class="py-3 px-6 text-sm text-gray-700 truncate">{{ cust.email }}</td>
            <td class="py-3 px-6 text-sm text-gray-700">{{ cust.phone || 'N/A' }}</td>
            <td class="py-3 px-6 text-sm text-gray-600">{{ formatDate(cust.createdAt) }}</td>
            <td class="py-3 px-6 text-sm text-gray-500 font-mono break-all">{{ cust.id }}</td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
      <div
        v-for="cust in filteredCustomers"
        :key="cust.linkId"
        @click="goToCustomerDetail(cust)"
        class="bg-white rounded-xl shadow-lg hover:shadow-xl transform hover:scale-[1.01] transition duration-200 ease-in-out overflow-hidden"
      >
        <div class="flex flex-col items-center p-6 text-center space-y-4">
          <div class="h-20 w-20 bg-gradient-to-br from-blue-400 to-purple-500 rounded-full flex items-center justify-center text-white text-2xl font-bold shadow-md">
            {{ getInitials(cust.firstName, cust.lastName) }}
          </div>
          <div class="space-y-1">
            <h3 class="text-xl font-semibold text-gray-800">{{ cust.firstName }} {{ cust.lastName }}</h3>
            <p class="text-sm text-gray-600 truncate">{{ cust.email }}</p>
            <p class="text-sm text-gray-600">{{ cust.phone || 'N/A' }}</p>
          </div>
        </div>
        <div class="px-6 py-3 bg-gray-50 border-t border-gray-100 flex items-center justify-between text-xs text-gray-500">
          <span class="font-mono text-gray-700">ID: {{ cust.id.slice(0, 8) }}...</span>
          <span>Joined: {{ formatDate(cust.createdAt) }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import api from '@/services/api';

export default {
  name: 'CustomersPage',
  data() {
    return {
      customers: [],
      loading: false,
      error: null,
      search: '',
      currentView: 'list' // Default view: 'list' (changed from 'cards')
    };
  },
  computed: {
    ...mapGetters('shops', ['activeShop']),
    filteredCustomers() {
      const term = this.search.trim().toLowerCase();
      if (!term) return this.customers;
      return this.customers.filter(c =>
        `${c.firstName || ''} ${c.lastName || ''}`.toLowerCase().includes(term) ||
        (c.email && c.email.toLowerCase().includes(term))
      );
    }
  },
  async created() {
    if (!this.activeShop) {
      this.error = 'No shop selected. Please select a shop from the dashboard to view its customers.';
      return;
    }
    this.loading = true;
    try {
      // Using the provided sample data structure for demonstration
      const sampleData = [
        {
          "customer": {
            "id": "682f140d663e5583675d378a",
            "firstName": "Laura",
            "lastName": "Martinez",
            "username": "ava.mart",
            "email": "laura.martinez@example.com",
            "phone": "+1-555-123-4567",
            "address": "1010 Pine Avenue",
            "city": "Austin",
            "state": "TX",
            "postalCode": "73301",
            "country": "USA",
            "createdAt": "2025-05-22T12:09:49.537Z",
            "updatedAt": "2025-05-22T12:09:49.537Z"
          },
          "linkId": "682f1cdd663e5583675d3790"
        },
        {
          "customer": {
            "id": "abc123def456ghi789jkl012",
            "firstName": "John",
            "lastName": "Doe",
            "username": "john.doe",
            "email": "john.doe@example.com",
            "phone": "+1-123-456-7890",
            "address": "123 Main St",
            "city": "Springfield",
            "state": "IL",
            "postalCode": "62704",
            "country": "USA",
            "createdAt": "2025-05-20T10:00:00.000Z",
            "updatedAt": "2025-05-21T11:30:00.000Z"
          },
          "linkId": "mno345pqr678stu901vwx234"
        },
         {
          "customer": {
            "id": "zxc789vbn012mas345dfg678",
            "firstName": "Alice",
            "lastName": "Smith",
            "username": "alice.s",
            "email": "alice.smith@example.com",
            "phone": "+44-20-7946-0958",
            "address": "789 Oak Lane",
            "city": "London",
            "state": "",
            "postalCode": "SW1A 0AA",
            "country": "UK",
            "createdAt": "2025-05-18T09:15:22.123Z",
            "updatedAt": "2025-05-19T14:05:30.456Z"
          },
          "linkId": "hjk890lmn123opq456rst789"
        }
      ];

      // In a real application, you would replace this with:
      // const res = await api.get(`/seller/shops/${this.activeShop.id}/customers`);
      // this.customers = res.data.map(item => ({...}));

      this.customers = sampleData.map(item => ({
        id:          item.customer.id,
        firstName:   item.customer.firstName,
        lastName:    item.customer.lastName,
        email:       item.customer.email,
        phone:       item.customer.phone,
        createdAt:   item.customer.createdAt,
        linkId:      item.linkId
      }));
    } catch (e) {
      console.error('Failed to load customers:', e);
      this.error = 'Failed to load customers. Please check your network connection or try again.';
    } finally {
      this.loading = false;
    }
  },
  methods: {
    formatDate(dt) {
      return dt ? new Date(dt).toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric'
      }) : '—';
    },
    getInitials(firstName, lastName) {
      const firstInitial = firstName ? firstName.charAt(0).toUpperCase() : '';
      const lastInitial = lastName ? lastName.charAt(0).toUpperCase() : '';
      return `${firstInitial}${lastInitial}` || '?'; // Return '?' if no initials available
    },
    goToCustomerDetail(cust) {
      if (this.activeShop && cust.id) {
        this.$router.push({
          name: 'CustomerDetail', // Make sure this route name exists in your Vue Router config
          params: { shopId: this.activeShop.id, customerId: cust.id }
        });
      } else {
        console.warn('Cannot navigate to customer detail: missing shopId or customerId', { activeShop: this.activeShop, customer: cust });
      }
    }
  }
};
</script>

<style scoped>
@keyframes spin {
  from { transform: rotate(0deg); }
  to   { transform: rotate(360deg); }
}
.animate-spin { animation: spin 1s linear infinite; }
</style>