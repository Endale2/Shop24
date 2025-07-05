<template>
  <div class="p-4 sm:p-6 bg-gray-50 min-h-screen font-sans">
    <div class="max-w-6xl mx-auto">
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

      <div v-else class="space-y-6">
        <!-- Customer Header Card -->
        <div class="bg-white rounded-xl shadow-lg overflow-hidden border border-gray-200">
          <div class="p-8 bg-gradient-to-r from-green-50 to-blue-50 border-b border-gray-200">
            <div class="flex flex-col lg:flex-row items-center space-y-6 lg:space-y-0 lg:space-x-8">
              <div class="flex-shrink-0">
                <img
                  v-if="customer.profile_image"
                  :src="customer.profile_image"
                  alt="Customer Profile"
                  class="h-32 w-32 rounded-full object-cover border-4 border-white shadow-lg"
                />
                <div
                  v-else
                  class="h-32 w-32 rounded-full bg-gradient-to-br from-green-400 to-teal-500 flex items-center justify-center text-white text-5xl font-bold border-4 border-white shadow-lg"
                >
                  {{ getInitials(customer) }}
                </div>
              </div>

              <div class="flex-grow text-center lg:text-left">
                <h1 class="text-4xl font-extrabold text-gray-800">
                  {{ customer.firstName }} {{ customer.lastName }}
                </h1>
                <p class="text-lg text-gray-500 font-mono mt-1">
                  @{{ customer.username }}
                </p>
                <p class="text-gray-600 mt-2">{{ customer.email }}</p>
                
                <!-- Customer Segments -->
                <div class="mt-4 flex flex-wrap gap-2 justify-center lg:justify-start">
                  <span
                    v-for="segment in customer.segments"
                    :key="segment.id"
                    class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium"
                    :style="{ backgroundColor: segment.color + '20', color: segment.color }"
                  >
                    <div
                      class="w-2 h-2 rounded-full mr-2"
                      :style="{ backgroundColor: segment.color }"
                    ></div>
                    {{ segment.name }}
                  </span>
                  <span v-if="customer.segments?.length === 0" class="text-gray-400 text-sm">
                    No segments assigned
                  </span>
                </div>
              </div>

              <div class="flex flex-col space-y-3">
                <button 
                  @click="showSegmentModal = true"
                  class="px-6 py-3 bg-blue-600 text-white text-sm font-medium rounded-lg shadow-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition"
                >
                  <FolderAddIcon class="w-5 h-5 inline-block mr-2" />
                  Manage Segments
                </button>
                <button class="px-6 py-3 bg-gray-200 text-gray-700 text-sm font-medium rounded-lg hover:bg-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-offset-2 transition">
                  <PencilIcon class="w-5 h-5 inline-block mr-2" />
                  Edit Profile
                </button>
                <button 
                  @click="unlinkCustomer"
                  class="px-6 py-3 bg-red-600 text-white text-sm font-medium rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2 transition"
                >
                  <UserRemoveIcon class="w-5 h-5 inline-block mr-2" />
                  Remove from Shop
                </button>
              </div>
            </div>
          </div>

          <!-- Customer Stats -->
          <div class="px-8 py-6 bg-gray-50 border-t border-gray-200">
            <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
              <div class="text-center">
                <p class="text-2xl font-bold text-gray-900">{{ customer.segments?.length || 0 }}</p>
                <p class="text-sm text-gray-600">Segments</p>
              </div>
              <div class="text-center">
                <p class="text-2xl font-bold text-gray-900">{{ formatDate(customer.createdAt) }}</p>
                <p class="text-sm text-gray-600">Joined</p>
              </div>
              <div class="text-center">
                <p class="text-2xl font-bold text-gray-900">{{ customer.phone || 'N/A' }}</p>
                <p class="text-sm text-gray-600">Phone</p>
              </div>
              <div class="text-center">
                <p class="text-2xl font-bold text-gray-900">{{ customer.city }}, {{ customer.state }}</p>
                <p class="text-sm text-gray-600">Location</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Main Content Grid -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
          <!-- Contact Information -->
          <div class="lg:col-span-2 space-y-6">
            <div class="bg-white rounded-xl shadow-lg overflow-hidden border border-gray-200">
              <div class="p-6 border-b border-gray-200">
                <h3 class="text-xl font-bold text-gray-900">Contact Information</h3>
              </div>
              <div class="p-6 space-y-4">
                <div class="flex items-center text-gray-700">
                  <MailIcon class="h-6 w-6 text-green-500 mr-4 flex-shrink-0" />
                  <div>
                    <p class="font-medium">Email Address</p>
                    <a :href="`mailto:${customer.email}`" class="text-green-600 hover:text-green-800 transition-colors">
                      {{ customer.email }}
                    </a>
                  </div>
                </div>
                <div class="flex items-center text-gray-700">
                  <PhoneIcon class="h-6 w-6 text-green-500 mr-4 flex-shrink-0" />
                  <div>
                    <p class="font-medium">Phone Number</p>
                    <span>{{ customer.phone || 'Not provided' }}</span>
                  </div>
                </div>
                <div class="flex items-start text-gray-700">
                  <OfficeBuildingIcon class="h-6 w-6 text-green-500 mr-4 flex-shrink-0 mt-1" />
                  <div>
                    <p class="font-medium">Address</p>
                    <div class="text-gray-600">
                      <p>{{ customer.address }}</p>
                      <p>{{ customer.city }}, {{ customer.state }} {{ customer.postalCode }}</p>
                      <p>{{ customer.country }}</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- Recent Activity -->
            <div class="bg-white rounded-xl shadow-lg overflow-hidden border border-gray-200">
              <div class="p-6 border-b border-gray-200">
                <h3 class="text-xl font-bold text-gray-900">Recent Activity</h3>
              </div>
              <div class="p-6">
                <div class="text-center text-gray-500 py-8">
                  <ClockIcon class="h-12 w-12 mx-auto mb-4 text-gray-300" />
                  <p>No recent activity to display</p>
                  <p class="text-sm">Customer activity will appear here</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Sidebar -->
          <div class="space-y-6">
            <!-- Segments Card -->
            <div class="bg-white rounded-xl shadow-lg overflow-hidden border border-gray-200">
              <div class="p-6 border-b border-gray-200">
                <h3 class="text-xl font-bold text-gray-900">Customer Segments</h3>
              </div>
              <div class="p-6">
                <div v-if="customer.segments?.length > 0" class="space-y-3">
                  <div
                    v-for="segment in customer.segments"
                    :key="segment.id"
                    class="flex items-center justify-between p-3 border border-gray-200 rounded-lg"
                  >
                    <div class="flex items-center">
                      <div
                        class="w-4 h-4 rounded-full mr-3"
                        :style="{ backgroundColor: segment.color }"
                      ></div>
                      <div>
                        <div class="font-medium text-gray-900">{{ segment.name }}</div>
                        <div class="text-sm text-gray-500">{{ segment.description || 'No description' }}</div>
                      </div>
                    </div>
                    <button
                      @click="removeCustomerFromSegment(customer.id, segment.id)"
                      class="text-red-600 hover:text-red-800 transition-colors"
                      title="Remove from segment"
                    >
                      <XIcon class="w-4 h-4" />
                    </button>
                  </div>
                </div>
                <div v-else class="text-center text-gray-500 py-8">
                  <FolderIcon class="h-12 w-12 mx-auto mb-4 text-gray-300" />
                  <p>No segments assigned</p>
                  <button
                    @click="showSegmentModal = true"
                    class="mt-3 text-green-600 hover:text-green-800 font-medium"
                  >
                    Add to segments
                  </button>
                </div>
              </div>
            </div>

            <!-- Quick Actions -->
            <div class="bg-white rounded-xl shadow-lg overflow-hidden border border-gray-200">
              <div class="p-6 border-b border-gray-200">
                <h3 class="text-xl font-bold text-gray-900">Quick Actions</h3>
              </div>
              <div class="p-6 space-y-3">
                <button class="w-full flex items-center px-4 py-3 text-left text-gray-700 hover:bg-gray-50 rounded-lg transition-colors">
                  <MailIcon class="w-5 h-5 mr-3 text-green-500" />
                  Send Email
                </button>
                <button class="w-full flex items-center px-4 py-3 text-left text-gray-700 hover:bg-gray-50 rounded-lg transition-colors">
                  <PhoneIcon class="w-5 h-5 mr-3 text-green-500" />
                  Call Customer
                </button>
                <button class="w-full flex items-center px-4 py-3 text-left text-gray-700 hover:bg-gray-50 rounded-lg transition-colors">
                  <DocumentTextIcon class="w-5 h-5 mr-3 text-green-500" />
                  View Orders
                </button>
                <button class="w-full flex items-center px-4 py-3 text-left text-gray-700 hover:bg-gray-50 rounded-lg transition-colors">
                  <ChatIcon class="w-5 h-5 mr-3 text-green-500" />
                  Send Message
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Manage Segments Modal -->
    <div v-if="showSegmentModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl shadow-xl max-w-lg w-full mx-4 max-h-[80vh] overflow-y-auto">
        <div class="p-6">
          <div class="flex items-center justify-between mb-6">
            <h3 class="text-lg font-semibold text-gray-900">
              Manage Segments for {{ customer.firstName }} {{ customer.lastName }}
            </h3>
            <button
              @click="showSegmentModal = false"
              class="text-gray-400 hover:text-gray-600 transition-colors"
            >
              <XIcon class="w-6 h-6" />
            </button>
          </div>
          
          <div v-if="loadingSegments" class="text-center py-8">
            <RefreshIcon class="animate-spin h-8 w-8 text-green-500 mx-auto mb-3" />
            <p class="text-gray-600">Loading segments...</p>
          </div>
          
          <div v-else-if="segments.length === 0" class="text-center py-8">
            <FolderIcon class="h-12 w-12 mx-auto mb-4 text-gray-300" />
            <p class="text-gray-600 mb-4">No segments created yet</p>
            <button
              @click="showCreateSegmentModal = true"
              class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors"
            >
              Create First Segment
            </button>
          </div>
          
          <div v-else class="space-y-4">
            <div v-for="segment in segments" :key="segment.id" class="flex items-center justify-between p-4 border border-gray-200 rounded-lg">
              <div class="flex items-center">
                <div
                  v-if="segment.color"
                  class="w-4 h-4 rounded-full mr-3"
                  :style="{ backgroundColor: segment.color }"
                ></div>
                <div>
                  <div class="font-medium text-gray-900">{{ segment.name }}</div>
                  <div class="text-sm text-gray-500">{{ segment.description || 'No description' }}</div>
                </div>
              </div>
              
              <button
                v-if="isCustomerInSegment(customer.id, segment.id)"
                @click="removeCustomerFromSegment(customer.id, segment.id)"
                class="px-3 py-1 text-red-600 hover:text-red-800 border border-red-200 hover:border-red-300 rounded-md transition-colors text-sm"
              >
                Remove
              </button>
              <button
                v-else
                @click="addCustomerToSegment(customer.id, segment.id)"
                class="px-3 py-1 text-green-600 hover:text-green-800 border border-green-200 hover:border-green-300 rounded-md transition-colors text-sm"
              >
                Add
              </button>
            </div>
          </div>
          
          <div class="flex justify-between mt-6 pt-6 border-t border-gray-200">
            <button
              @click="showCreateSegmentModal = true"
              class="px-4 py-2 text-blue-600 hover:text-blue-800 font-medium transition-colors"
            >
              + Create New Segment
            </button>
            <button
              @click="showSegmentModal = false"
              class="px-4 py-2 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300 transition-colors"
            >
              Close
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Create Segment Modal -->
    <div v-if="showCreateSegmentModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl shadow-xl max-w-md w-full mx-4">
        <div class="p-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">Create New Segment</h3>
          
          <form @submit.prevent="createSegment">
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">Segment Name</label>
                <input
                  v-model="newSegment.name"
                  type="text"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
                  placeholder="e.g., VIP Customers"
                />
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">Description</label>
                <textarea
                  v-model="newSegment.description"
                  rows="3"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
                  placeholder="Optional description..."
                ></textarea>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">Color</label>
                <input
                  v-model="newSegment.color"
                  type="color"
                  class="w-full h-10 border border-gray-300 rounded-lg cursor-pointer"
                />
              </div>
            </div>
            
            <div class="flex justify-end space-x-3 mt-6">
              <button
                type="button"
                @click="showCreateSegmentModal = false"
                class="px-4 py-2 text-gray-700 bg-gray-200 rounded-lg hover:bg-gray-300 transition-colors"
              >
                Cancel
              </button>
              <button
                type="submit"
                :disabled="creatingSegment"
                class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 disabled:opacity-50 transition-colors"
              >
                {{ creatingSegment ? 'Creating...' : 'Create Segment' }}
              </button>
            </div>
          </form>
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
import { customerSegmentService } from '@/services/customerSegment';
import {
  ChevronLeftIcon,
  RefreshIcon,
  MailIcon,
  PhoneIcon,
  OfficeBuildingIcon,
  FolderIcon,
  FolderAddIcon,
  PencilIcon,
  UserRemoveIcon,
  ClockIcon,
  DocumentTextIcon,
  ChatIcon,
  XIcon
} from '@heroicons/vue/outline';

const router = useRouter();
const route = useRoute();
const shopStore = useShopStore();

// Reactive state for customer data
const customer = ref({
  profile_image: '',
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
  segments: []
});
const segments = ref([]);
const loading = ref(true);
const loadingSegments = ref(false);
const error = ref(null);
const showSegmentModal = ref(false);
const showCreateSegmentModal = ref(false);
const creatingSegment = ref(false);

// New segment form
const newSegment = ref({
  name: '',
  description: '',
  color: '#3B82F6'
});

const activeShop = computed(() => shopStore.activeShop);

// Fetch data on component mount
onMounted(async () => {
  if (!activeShop.value?.id) {
    error.value = 'No shop selected. Please select a shop first.';
    loading.value = false;
    return;
  }
  
  await refreshData();
});

// Refresh customer and segments data
async function refreshData() {
  const customerId = route.params.customerId;
  try {
    const [customerResult, segmentsResult] = await Promise.all([
      customerService.fetchById(activeShop.value.id, customerId),
      customerSegmentService.fetchAll(activeShop.value.id)
    ]);
    
    customer.value = customerResult;
    segments.value = Array.isArray(segmentsResult) ? segmentsResult : [];
  } catch (e) {
    console.error('Error fetching customer details:', e);
    error.value = 'Failed to load customer details. The customer may not exist or an error occurred.';
  } finally {
    loading.value = false;
  }
}

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

function isCustomerInSegment(customerId, segmentId) {
  const segment = segments.value.find(s => s.id === segmentId)
  return segment?.customer_ids && segment.customer_ids.includes(customerId)
}

async function createSegment() {
  if (!newSegment.value.name.trim()) return;
  
  creatingSegment.value = true;
  try {
    await customerSegmentService.create(activeShop.value.id, newSegment.value);
    showCreateSegmentModal.value = false;
    newSegment.value = { name: '', description: '', color: '#3B82F6' };
    await refreshData(); // Refresh all data
  } catch (e) {
    console.error('Failed to create segment:', e);
    error.value = 'Failed to create segment. Please try again.';
  } finally {
    creatingSegment.value = false;
  }
}

async function addCustomerToSegment(customerId, segmentId) {
  try {
    await customerSegmentService.addCustomer(activeShop.value.id, segmentId, customerId);
    await refreshData(); // Refresh all data
  } catch (e) {
    console.error('Failed to add customer to segment:', e);
    error.value = 'Failed to add customer to segment. Please try again.';
  }
}

async function removeCustomerFromSegment(customerId, segmentId) {
  try {
    await customerSegmentService.removeCustomer(activeShop.value.id, segmentId, customerId);
    await refreshData(); // Refresh all data
  } catch (e) {
    console.error('Failed to remove customer from segment:', e);
    error.value = 'Failed to remove customer from segment. Please try again.';
  }
}

async function unlinkCustomer() {
  if (!confirm('Are you sure you want to remove this customer from your shop? This action cannot be undone.')) {
    return;
  }
  
  try {
    await customerService.delete(activeShop.value.id, customer.value.linkId);
    router.push({ name: 'Customers' });
  } catch (e) {
    console.error('Failed to unlink customer:', e);
    error.value = 'Failed to remove customer. Please try again.';
  }
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