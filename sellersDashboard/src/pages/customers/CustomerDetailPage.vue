<template>
  <div class="p-6 max-w-3xl mx-auto space-y-6">
    <!-- Back button -->
    <button
      @click="$router.back()"
      class="inline-flex items-center text-gray-600 hover:text-gray-800 mb-4"
    >
      <svg class="h-5 w-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M15 19l-7-7 7-7" />
      </svg>
      Back to Customers
    </button>

    <!-- Loading -->
    <div v-if="loading" class="flex items-center justify-center text-gray-600 py-12">
      <svg class="animate-spin h-8 w-8 text-green-500 mr-2" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
        <path class="opacity-75" fill="currentColor"
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 
                 5.291A7.962 7.962 0 014 12H0c0 3.042 
                 1.135 5.824 3 7.938l3-2.647z"/>
      </svg>
      Loading customer…
    </div>

    <!-- Error -->
    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded">
      <strong class="font-bold">Error:</strong>
      <span class="ml-2">{{ error }}</span>
    </div>

    <!-- Customer Detail -->
    <div v-else class="bg-white rounded-lg shadow p-6 space-y-6">
      <div class="flex flex-col md:flex-row md:space-x-6 items-center md:items-start">
        <!-- Avatar -->
        <div class="flex-shrink-0 mb-4 md:mb-0">
          <div class="h-32 w-32 bg-gray-100 rounded-full flex items-center justify-center">
            <svg class="h-12 w-12 text-gray-400" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12 12c2.7 0 4.9-2.2 4.9-4.9S14.7 2.2 12 2.2 7.1 4.4 7.1 7.1 9.3 12 12 12zm0 2.2c-3.2 0-9.6 1.6-9.6 4.9V22h19.2v-2.9c0-3.3-6.4-4.9-9.6-4.9z"/>
            </svg>
          </div>
        </div>

        <!-- Info -->
        <div class="flex-1 space-y-2">
          <h1 class="text-2xl font-bold text-gray-800">
            {{ customer.firstName }} {{ customer.lastName }}
          </h1>
          <p class="text-gray-600">Username: <span class="font-medium">{{ customer.username }}</span></p>
          <p class="text-gray-600">Email: <span class="font-medium">{{ customer.email }}</span></p>
          <p class="text-gray-600">Phone: <span class="font-medium">{{ customer.phone || '—' }}</span></p>

          <div class="mt-4">
            <h3 class="font-semibold text-gray-800 mb-1">Address</h3>
            <p class="text-gray-700">{{ customer.address }}</p>
            <p class="text-gray-700">
              {{ customer.city }}, {{ customer.state }} {{ customer.postalCode }}
            </p>
            <p class="text-gray-700">{{ customer.country }}</p>
          </div>
        </div>
      </div>

      <!-- Timestamps -->
      <div class="text-sm text-gray-500 space-y-1">
        <p>Joined: {{ formatDate(customer.createdAt) }}</p>
        <p>Last Updated: {{ formatDate(customer.updatedAt) }}</p>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import api from '@/services/api';

export default {
  name: 'CustomerDetailPage',
  data() {
    return {
      customer: {
        firstName: '',
        lastName: '',
        username: '',
        email: '',
        phone: '',
        address: '',
        city: '',
        state: '',
        postalCode: '',
        country: ''
      },
      loading: true,
      error: null
    };
  },
  computed: {
    ...mapGetters('shops', ['activeShop'])
  },
  async created() {
    if (!this.activeShop) {
      this.error = 'No shop selected.';
      this.loading = false;
      return;
    }
    const customerId = this.$route.params.customerId;
    try {
      const res = await api.get(
        `/seller/shops/${this.activeShop.id}/customers/${customerId}`
      );
      // assume { customer: {...}, linkId } response
      this.customer = res.data.customer;
    } catch (e) {
      console.error(e);
      this.error = 'Failed to load customer details.';
    } finally {
      this.loading = false;
    }
  },
  methods: {
    formatDate(dt) {
      return dt ? new Date(dt).toLocaleDateString() : '—';
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
