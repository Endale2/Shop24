<template>
  <div class="p-4 sm:p-6 max-w-6xl mx-auto space-y-8 font-sans">
    <button
      @click="$router.back()"
      class="inline-flex items-center text-gray-600 hover:text-green-700 transition duration-200 ease-in-out mb-6 group rounded-full px-3 py-1.5 -ml-3"
    >
      <ChevronLeftIcon class="h-5 w-5 mr-1 text-gray-500 group-hover:text-green-600 transition-colors duration-200" />
      <span class="text-sm font-medium group-hover:text-green-700 transition-colors duration-200">Back to Customers</span>
    </button>

    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-12">
      <RefreshIcon class="animate-spin h-8 w-8 text-green-500 mb-3" />
      <p class="mt-3 text-lg">Loading customer details...</p>
    </div>

    <div
      v-else-if="error"
      class="bg-red-50 border border-red-200 text-red-700 px-6 py-4 rounded-lg shadow-sm mb-6"
      role="alert"
    >
      <div class="flex items-center">
        <ExclamationCircleIcon class="h-5 w-5 mr-2 flex-shrink-0" />
        <div>
          <strong class="font-semibold">Error:</strong>
          <span class="ml-1">{{ error }}</span>
        </div>
      </div>
      <p class="text-sm mt-2">Please try again or contact support if the issue persists.</p>
    </div>

    <div v-else-if="customer.id" class="bg-white rounded-xl shadow-lg p-6 md:p-8 space-y-8 animate-fade-in">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 lg:gap-12">
        <!-- Customer Profile -->
        <div class="space-y-6">
          <div class="flex flex-col items-center lg:items-start space-y-6">
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

            <div class="text-center lg:text-left">
              <h1 class="text-4xl sm:text-5xl font-extrabold text-gray-900 leading-tight">
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
          </div>

          <!-- Action Buttons -->
          <div class="flex flex-col sm:flex-row space-y-3 sm:space-y-0 sm:space-x-4">
            <button 
              @click="showSegmentModal = true"
              class="flex-1 w-full sm:w-auto inline-flex items-center justify-center px-6 py-3 border border-transparent text-base font-medium rounded-lg shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition duration-150 ease-in-out transform hover:-translate-y-0.5"
            >
              <FolderAddIcon class="w-5 h-5 mr-2" />
              Manage Segments
            </button>
            <button class="flex-1 w-full sm:w-auto inline-flex items-center justify-center px-6 py-3 border border-gray-300 text-base font-medium rounded-lg shadow-sm text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500 transition duration-150 ease-in-out">
              <PencilIcon class="w-5 h-5 mr-2" />
              Edit Profile
            </button>
            <button 
              @click="unlinkCustomer"
              class="flex-1 w-full sm:w-auto inline-flex items-center justify-center px-6 py-3 border border-red-300 text-base font-medium rounded-lg shadow-sm text-red-700 bg-white hover:bg-red-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 transition duration-150 ease-in-out"
            >
              <UserRemoveIcon class="w-5 h-5 mr-2" />
              Remove from Shop
            </button>
          </div>
        </div>

        <!-- Customer Stats -->
        <div class="space-y-6">
          <div class="bg-gray-50 rounded-xl p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">Customer Information</h3>
            <div class="grid grid-cols-2 gap-4">
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

          <!-- Contact Information -->
          <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">Contact Information</h3>
            <div class="space-y-4">
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
        </div>
      </div>

      <!-- Customer Segments Section -->
      <div class="border-t border-gray-200 pt-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-xl font-semibold text-gray-900">Customer Segments</h3>
          <button 
            @click="showSegmentModal = true"
            class="inline-flex items-center px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition duration-150 ease-in-out"
          >
            <FolderAddIcon class="w-4 h-4 mr-2" />
            Manage Segments
          </button>
        </div>
        
        <div v-if="customer.segments?.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          <div
            v-for="segment in customer.segments"
            :key="segment.id"
            class="flex items-center justify-between p-4 border border-gray-200 rounded-lg bg-white shadow-sm hover:shadow-md transition-shadow duration-200"
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
          <p class="text-sm">This customer is not part of any segments yet.</p>
        </div>
      </div>

      <!-- Recent Activity -->
      <div class="border-t border-gray-200 pt-6">
        <h3 class="text-xl font-semibold text-gray-900 mb-4">Recent Activity</h3>
        <div class="bg-gray-50 rounded-xl p-6">
          <div class="text-center text-gray-500 py-8">
            <ClockIcon class="h-12 w-12 mx-auto mb-4 text-gray-300" />
            <p>No recent activity to display</p>
            <p class="text-sm">Customer activity will appear here</p>
          </div>
        </div>
      </div>

      <!-- Customer Metadata -->
      <div class="border-t border-gray-200 pt-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-sm text-gray-500">
          <div>
            <p><strong>Customer ID:</strong> {{ customer.id }}</p>
            <p><strong>Created:</strong> {{ formatDate(customer.createdAt) }}</p>
          </div>
          <div>
            <p><strong>Username:</strong> @{{ customer.username }}</p>
            <p><strong>Last Updated:</strong> {{ formatDate(customer.updatedAt) }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Not Found State -->
    <div v-else class="bg-gray-50 border border-gray-200 text-gray-700 px-6 py-8 rounded-lg text-center mt-8 shadow-sm">
      <div class="flex flex-col items-center">
        <div class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center mb-4">
          <ExclamationCircleIcon class="w-8 h-8 text-gray-400" />
        </div>
        <p class="text-lg font-medium mb-2">Customer not found</p>
        <p class="text-sm">The customer ID might be invalid or the customer does not exist for the active shop.</p>
        <button 
          @click="$router.back()" 
          class="mt-4 inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
        >
          <ChevronLeftIcon class="w-4 h-4 mr-2" />
          Back to Customers
        </button>
      </div>
    </div>

    <!-- Manage Segments Modal -->
    <div v-if="showSegmentModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl shadow-xl max-w-lg w-full mx-4">
        <div class="p-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">
            Manage Segments for {{ customer.firstName }} {{ customer.lastName }}
          </h3>
          
          <div class="space-y-4 max-h-96 overflow-y-auto">
            <div v-for="segment in availableSegments" :key="segment.id" class="flex items-center justify-between p-3 border border-gray-200 rounded-lg">
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
                class="text-red-600 hover:text-red-800 transition-colors"
              >
                Remove
              </button>
              <button
                v-else
                @click="addCustomerToSegment(customer.id, segment.id)"
                class="text-green-600 hover:text-green-800 transition-colors"
              >
                Add
              </button>
            </div>
          </div>
          
          <div class="flex justify-end mt-6">
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
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { customerService } from '@/services/customer'
import { customerSegmentService } from '@/services/customerSegment'
import {
  ChevronLeftIcon,
  RefreshIcon,
  ExclamationCircleIcon,
  FolderAddIcon,
  PencilIcon,
  UserRemoveIcon,
  MailIcon,
  PhoneIcon,
  OfficeBuildingIcon,
  ClockIcon,
  FolderIcon,
  XIcon
} from '@heroicons/vue/outline'

const route = useRoute()
const router = useRouter()
const shopStore = useShopStore()

// Reactive state
const customer = ref({})
const availableSegments = ref([])
const loading = ref(false)
const error = ref(null)
const showSegmentModal = ref(false)

// Computed
const activeShop = computed(() => shopStore.activeShop)

// Fetch data
onMounted(async () => {
  if (!activeShop.value?.id) {
    router.replace({ name: 'ShopSelection' })
    return
  }
  await fetchCustomerData()
})

async function fetchCustomerData() {
  loading.value = true
  try {
    const [customerResult, segmentsResult] = await Promise.all([
      customerService.fetchById(activeShop.value.id, route.params.customerId),
      customerSegmentService.fetchAll(activeShop.value.id)
    ])
    
    if (customerResult) {
      customer.value = customerResult
      customer.value.segments = customerResult.segments || []
    } else {
      customer.value = {}
    }
    
    availableSegments.value = Array.isArray(segmentsResult) ? segmentsResult : []
  } catch (e) {
    console.error(e)
    error.value = 'Failed to load customer details. Please try again later.'
  } finally {
    loading.value = false
  }
}

// Methods
function getInitials(customer) {
  const fi = customer.firstName?.[0]?.toUpperCase() || ''
  const li = customer.lastName?.[0]?.toUpperCase() || ''
  return fi + li || '?'
}

function formatDate(dt) {
  return dt
    ? new Date(dt).toLocaleDateString(undefined, {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
      })
    : '—'
}

function isCustomerInSegment(customerId, segmentId) {
  const segment = availableSegments.value.find(s => s.id === segmentId)
  return segment?.customer_ids && segment.customer_ids.includes(customerId)
}

async function addCustomerToSegment(customerId, segmentId) {
  try {
    await customerSegmentService.addCustomer(activeShop.value.id, segmentId, customerId)
    await fetchCustomerData() // Refresh data
  } catch (e) {
    console.error('Failed to add customer to segment:', e)
    error.value = 'Failed to add customer to segment. Please try again.'
  }
}

async function removeCustomerFromSegment(customerId, segmentId) {
  try {
    await customerSegmentService.removeCustomer(activeShop.value.id, segmentId, customerId)
    await fetchCustomerData() // Refresh data
  } catch (e) {
    console.error('Failed to remove customer from segment:', e)
    error.value = 'Failed to remove customer from segment. Please try again.'
  }
}

async function unlinkCustomer() {
  if (!confirm('Are you sure you want to remove this customer from your shop?')) return
  
  try {
    await customerService.delete(activeShop.value.id, customer.value.linkId)
    router.push({ name: 'Customers' })
  } catch (e) {
    console.error('Failed to unlink customer:', e)
    error.value = 'Failed to remove customer. Please try again.'
  }
}
</script>

<style scoped>
@keyframes fade-in {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
.animate-fade-in {
  animation: fade-in 0.3s ease-out;
}
</style>