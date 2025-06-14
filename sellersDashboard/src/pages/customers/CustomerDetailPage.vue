<template>
  <div class="p-4 sm:p-6 bg-gray-50 min-h-screen font-sans">
    <div class="max-w-4xl mx-auto">
      <button
        @click="$router.back()"
        class="inline-flex items-center text-gray-600 hover:text-green-700 font-medium mb-4 transition-colors duration-200 rounded-md px-3 py-1 -ml-3"
      >
        <ChevronLeftIcon class="h-5 w-5 mr-1" />
        Back to Customers
      </button>

      <div v-if="loading" class="flex items-center justify-center text-gray-600 py-16">
        <RefreshIcon class="animate-spin h-10 w-10 text-green-500 mr-3" />
        <span class="text-xl font-semibold">Loading customer details...</span>
      </div>

      <div v-else-if="error" class="bg-red-100 border-l-4 border-red-500 text-red-700 p-4 rounded-lg shadow-md" role="alert">
        <p class="font-bold">Error</p>
        <p>{{ error }}</p>
      </div>

      <div v-else class="bg-white rounded-xl shadow-lg overflow-hidden border border-gray-200">
        <div class="p-8 bg-gray-50 border-b border-gray-200">
          <div class="flex flex-col sm:flex-row items-center space-y-4 sm:space-y-0 sm:space-x-6">
            <div class="flex-shrink-0">
              <img
                v-if="customer.profile_image"
                :src="customer.profile_image"
                alt="Customer Profile"
                class="h-28 w-28 rounded-full object-cover border-4 border-white shadow-md"
              />
              <div
                v-else
                class="h-28 w-28 rounded-full bg-gradient-to-br from-green-400 to-teal-500 flex items-center justify-center text-white text-4xl font-bold border-4 border-white shadow-md"
              >
                {{ getInitials(customer) }}
              </div>
            </div>

            <div class="flex-grow text-center sm:text-left">
              <h1 class="text-4xl font-extrabold text-gray-800">
                {{ customer.firstName }} {{ customer.lastName }}
              </h1>
              <p class="text-lg text-gray-500 font-mono mt-1">
                @{{ customer.username }}
              </p>
            </div>

            <div class="flex space-x-3">
              <button class="px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg shadow-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition">
                Edit
              </button>
              <button class="px-4 py-2 bg-gray-200 text-gray-700 text-sm font-medium rounded-lg hover:bg-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-offset-2 transition">
                Delete
              </button>
            </div>
          </div>
        </div>

        <div class="p-8 grid grid-cols-1 md:grid-cols-2 gap-8">
          <div class="space-y-4">
            <h3 class="font-bold text-xl text-gray-800 border-b pb-2 mb-4">Contact Information</h3>
            <div class="flex items-center text-gray-700">
              <MailIcon class="h-6 w-6 text-green-500 mr-4" />
              <a :href="`mailto:${customer.email}`" class="hover:text-green-600">{{ customer.email }}</a>
            </div>
            <div class="flex items-center text-gray-700">
              <PhoneIcon class="h-6 w-6 text-green-500 mr-4" />
              <span>{{ customer.phone || 'Not provided' }}</span>
            </div>
          </div>

          <div class="space-y-4">
            <h3 class="font-bold text-xl text-gray-800 border-b pb-2 mb-4">Address</h3>
            <div class="flex items-start text-gray-700">
              <OfficeBuildingIcon class="h-6 w-6 text-green-500 mr-4 flex-shrink-0" />
              <div>
                <p>{{ customer.address }}</p>
                <p>{{ customer.city }}, {{ customer.state }} {{ customer.postalCode }}</p>
                <p>{{ customer.country }}</p>
              </div>
            </div>
          </div>
        </div>

        <div class="px-8 py-4 bg-gray-50 border-t border-gray-200 text-sm text-gray-500 flex justify-between">
          <div>
            <strong>Joined:</strong> {{ formatDate(customer.createdAt) }}
          </div>
          <div>
            <strong>Last Updated:</strong> {{ formatDate(customer.updatedAt) }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useShopStore } from '@/store/shops';
import { customerService } from '@/services/customer';
import {
  ChevronLeftIcon,
  RefreshIcon,
  MailIcon,
  PhoneIcon,
  OfficeBuildingIcon
} from '@heroicons/vue/solid';

const router = useRouter();
const route = useRoute();
const shopStore = useShopStore();

// Reactive state for customer data
const customer = ref({
  profile_image: '', // Added for dynamic avatar
  firstName: '',
  lastName: '',
  username: '',
  email: '',
  phone: '',
  address: '',
  city: '',
  state: '',
  postalCode: '',
  country: '',
  createdAt: null,
  updatedAt: null,
});
const loading = ref(true);
const error = ref(null);

const activeShop = computed(() => shopStore.activeShop);

// Fetch data on component mount
onMounted(async () => {
  if (!activeShop.value?.id) {
    error.value = 'No shop selected. Please select a shop first.';
    loading.value = false;
    return;
  }
  const customerId = route.params.customerId;
  try {
    customer.value = await customerService.fetchById(activeShop.value.id, customerId);
  } catch (e) {
    console.error('Error fetching customer details:', e);
    error.value = 'Failed to load customer details. The customer may not exist or an error occurred.';
  } finally {
    loading.value = false;
  }
});

/**
 * Generates initials from the customer's name or email.
 * @param {object} cust - The customer object.
 * @returns {string} The initials or a fallback character.
 */
function getInitials(cust) {
  if (cust.firstName && cust.lastName) {
    return `${cust.firstName[0]}${cust.lastName[0]}`.toUpperCase();
  }
  if (cust.email) {
    return cust.email[0].toUpperCase();
  }
  return '?';
}

/**
 * Formats a date string into a more readable format.
 * @param {string} dt - The date string.
 * @returns {string} Formatted date.
 */
function formatDate(dt) {
  return dt ? new Date(dt).toLocaleDateString(undefined, {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  }) : '—';
}
</script>

<style scoped>
.animate-spin {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to   { transform: rotate(360deg); }
}
</style>