<template>
  <div class="p-6 max-w-3xl mx-auto space-y-6">
    <button
      @click="$router.back()"
      class="inline-flex items-center text-gray-600 hover:text-green-700 font-medium mb-4 transition-colors duration-200 rounded-md px-3 py-1 -ml-3"
    >
      <svg class="h-5 w-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M15 19l-7-7 7-7" />
      </svg>
      Back to Customers
    </button>

    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-12">
      <svg class="animate-spin h-10 w-10 text-green-500 mb-3" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
        <path class="opacity-75" fill="currentColor"
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2
                5.291A7.962 7.962 0 014 12H0c0 3.042
                1.135 5.824 3 7.938l3-2.647z"/>
      </svg>
      <span class="text-xl font-semibold text-green-700">Loading customer details...</span>
    </div>

    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-lg shadow-sm">
      <strong class="font-bold">Error:</strong>
      <span class="ml-2">{{ error }}</span>
    </div>

    <div v-else class="bg-white rounded-xl shadow-lg p-8 space-y-8 border border-gray-100">
      <div class="flex flex-col md:flex-row md:space-x-8 items-center md:items-start">
        <div class="flex-shrink-0 mb-6 md:mb-0">
          <div class="h-36 w-36 bg-gray-100 rounded-full flex items-center justify-center border-4 border-green-200 shadow-inner">
            <svg class="h-20 w-20 text-gray-400" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12 12c2.7 0 4.9-2.2 4.9-4.9S14.7 2.2 12 2.2 7.1 4.4 7.1 7.1 9.3 12 12 12zm0 2.2c-3.2 0-9.6 1.6-9.6 4.9V22h19.2v-2.9c0-3.3-6.4-4.9-9.6-4.9z"/>
            </svg>
          </div>
        </div>

        <div class="flex-1 space-y-3 text-center md:text-left">
          <h1 class="text-4xl font-extrabold text-gray-800 leading-tight">
            {{ customer.firstName }} {{ customer.lastName }}
          </h1>
          <p class="text-gray-700 text-lg">
            Username: <span class="font-semibold text-green-700">{{ customer.username }}</span>
          </p>
          <p class="text-gray-700 text-lg">
            Email: <span class="font-semibold text-green-700">{{ customer.email }}</span>
          </p>
          <p class="text-gray-700 text-lg">
            Phone: <span class="font-semibold text-green-700">{{ customer.phone || '—' }}</span>
          </p>
        </div>
      </div>

      <div class="mt-6 p-6 bg-green-50 rounded-lg shadow-sm border border-green-100">
        <h3 class="font-bold text-xl text-green-800 mb-3 border-b-2 border-green-300 pb-2">Contact Address</h3>
        <p class="text-gray-800 text-lg">{{ customer.address }}</p>
        <p class="text-gray-800 text-lg">
          {{ customer.city }}, {{ customer.state }} {{ customer.postalCode }}
        </p>
        <p class="text-gray-800 text-lg">{{ customer.country }}</p>
      </div>

      <div class="text-sm text-gray-500 space-y-2 border-t pt-6 border-gray-200">
        <p><strong class="text-green-700">Joined:</strong> {{ formatDate(customer.createdAt) }}</p>
        <p><strong class="text-green-700">Last Updated:</strong> {{ formatDate(customer.updatedAt) }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router'; // Import useRoute for params
import { useShopStore } from '@/store/shops'; // Use the Pinia shop store
import { customerService } from '@/services/customer';

const router = useRouter();
const route = useRoute(); // Access route object for params
const shopStore = useShopStore(); // Pinia shop store instance

// Reactive state for customer data and loading
const customer = ref({
  firstName:    '',
  lastName:     '',
  username:     '',
  email:        '',
  phone:        '',
  address:      '',
  city:         '',
  state:        '',
  postalCode:   '',
  country:      '',
  createdAt:    null, // Added createdAt and updatedAt for consistency
  updatedAt:    null,
});
const loading = ref(true);
const error = ref(null);

// Computed property for active shop
const activeShop = computed(() => shopStore.activeShop);

// Lifecycle hook to fetch data when the component is mounted
onMounted(async () => {
  if (!activeShop.value) {
    error.value = 'No shop selected. Please select a shop to view customer details.';
    loading.value = false;
    return;
  }
  const customerId = route.params.customerId; // Get customerId from route params

  try {
    customer.value = await customerService.fetchById(
      activeShop.value.id,
      customerId
    );
  } catch (e) {
    console.error('Error fetching customer details:', e);
    error.value = 'Failed to load customer details.';
  } finally {
    loading.value = false;
  }
});

// Helper function for date formatting
function formatDate(dt) {
  return dt ? new Date(dt).toLocaleDateString() : '—';
}
</script>

<style scoped>
/* Keep your existing scoped styles here if any specific animations or overrides */
.animate-spin {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to   { transform: rotate(360deg); }
}
</style>