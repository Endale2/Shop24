<template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto space-y-8 font-sans">
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center md:justify-between space-y-4 md:space-y-0">
      <div>
        <h2 class="text-3xl sm:text-4xl font-extrabold text-gray-900 leading-tight">
          Customers
        </h2>
        <p class="text-gray-600 mt-2">Manage your customer relationships and segments</p>
      </div>

      <div class="flex flex-col sm:flex-row sm:items-center space-y-4 sm:space-y-0 sm:space-x-4 w-full md:w-auto">
        <div class="relative w-full sm:w-64">
          <input
            v-model="search"
            type="text"
            placeholder="Search by name or email..."
            class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm"
          />
          <SearchIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
        </div>

        <button
          @click="showCreateSegmentModal = true"
          class="inline-flex items-center px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg shadow-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition duration-150 ease-in-out"
        >
          <FolderAddIcon class="w-5 h-5 mr-2 -ml-1" />
          New Segment
        </button>

        <button
          @click="goToSegmentsPage"
          class="inline-flex items-center px-4 py-2 bg-gray-100 text-blue-700 text-sm font-medium rounded-lg shadow-md hover:bg-blue-100 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition duration-150 ease-in-out"
        >
          <FolderIcon class="w-5 h-5 mr-2 -ml-1" />
          Manage Segments
        </button>

        <div class="inline-flex rounded-full shadow-md overflow-hidden" role="group">
          <button
            @click="currentView = 'cards'"
            :class="buttonClass('cards')"
            class="px-5 py-2.5 text-sm font-medium border border-gray-300 focus:z-10 focus:ring-2 focus:ring-green-500 focus:outline-none transition-colors duration-200 ease-in-out"
          >
            <ViewGridIcon class="w-5 h-5 inline-block mr-1" />
            Cards
          </button>
          <button
            @click="currentView = 'list'"
            :class="buttonClass('list')"
            class="px-5 py-2.5 text-sm font-medium border border-gray-300 focus:z-10 focus:ring-2 focus:ring-green-500 focus:outline-none transition-colors duration-200 ease-in-out"
          >
            <ViewListIcon class="w-5 h-5 inline-block mr-1" />
            List
          </button>
        </div>
      </div>
    </div>

    <!-- Statistics Dashboard -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <UsersIcon class="w-6 h-6 text-green-600" />
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">Total Customers</p>
            <p class="text-2xl font-bold text-gray-900">{{ customers.length }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <FolderIcon class="w-6 h-6 text-blue-600" />
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">Segments</p>
            <p class="text-2xl font-bold text-gray-900">{{ segments.length }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <UserAddIcon class="w-6 h-6 text-yellow-600" />
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">This Month</p>
            <p class="text-2xl font-bold text-gray-900">{{ newCustomersThisMonth }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-2 bg-purple-100 rounded-lg">
            <UserGroupIcon class="w-6 h-6 text-purple-600" />
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">Segmented</p>
            <p class="text-2xl font-bold text-gray-900">{{ segmentedCustomersCount }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Segment Filter -->
    <div v-if="segments.length > 0" class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">Filter by Segment</h3>
      <div class="flex flex-wrap gap-3">
        <button
          @click="selectedSegment = null"
          :class="selectedSegment === null ? 'bg-green-600 text-white' : 'bg-gray-100 text-gray-700 hover:bg-gray-200'"
          class="px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-200"
        >
          All Customers
        </button>
        <button
          v-for="segment in segments"
          :key="segment.id"
          @click="selectedSegment = segment.id"
          :class="selectedSegment === segment.id ? 'bg-green-600 text-white' : 'bg-gray-100 text-gray-700 hover:bg-gray-200'"
          class="px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-200 flex items-center"
        >
          <div
            v-if="segment.color"
            class="w-3 h-3 rounded-full mr-2"
            :style="{ backgroundColor: segment.color }"
          ></div>
          {{ segment.name }}
          <span class="ml-2 bg-gray-200 text-gray-700 px-2 py-0.5 rounded-full text-xs">
            {{ getCustomersInSegment(segment.id).length }}
          </span>
        </button>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-16">
      <RefreshIcon class="animate-spin h-10 w-10 text-green-500 mb-3" />
      <p class="text-lg">Loading customers...</p>
    </div>

    <!-- Error State -->
    <div
      v-else-if="error"
      class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-lg shadow-sm"
      role="alert"
    >
      <p class="font-bold">Oops! Something went wrong:</p>
      <p class="mt-2">{{ error }}</p>
      <p class="text-sm mt-2">Please try again or contact support if the issue persists.</p>
    </div>

    <!-- Empty State -->
    <div
      v-else-if="filteredCustomers.length === 0"
      class="bg-green-50 border border-green-200 text-green-700 px-6 py-8 rounded-lg text-center mt-8 shadow-sm"
    >
      <p class="text-lg font-medium">
        No customers found<span v-if="search"> matching "{{ search }}"</span><span v-if="selectedSegment"> in this segment</span>.
      </p>
      <p class="mt-2" v-if="!search && !selectedSegment">
        It looks like there are no customers for your active shop yet.
      </p>
    </div>

    <!-- List View -->
    <div v-else-if="currentView === 'list'" class="overflow-x-auto bg-white shadow-lg rounded-xl">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-100">
          <tr>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Customer</th>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Contact</th>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Segments</th>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Joined</th>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Actions</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr
            v-for="(cust, i) in filteredCustomers"
            :key="cust.id"
            class="transition duration-150 ease-in-out transform hover:scale-[1.005] hover:bg-green-50"
            :class="{ 'bg-gray-50': i % 2 === 1 }"
          >
            <td class="py-3 px-6 text-sm text-gray-800 font-medium">
              <div class="flex items-center">
                <div
                  class="h-10 w-10 bg-green-100 rounded-full flex items-center justify-center text-green-700 text-sm font-semibold mr-3 cursor-pointer"
                  @click="goToCustomerDetail(cust.id)"
                >
                  {{ getInitials(cust.firstName, cust.lastName) }}
                </div>
                <div>
                  <div class="font-medium text-gray-900 cursor-pointer" @click="goToCustomerDetail(cust.id)">
                    {{ cust.firstName }} {{ cust.lastName }}
                  </div>
                  <div class="text-gray-500 font-mono text-xs">@{{ cust.username }}</div>
                </div>
              </div>
            </td>
            <td class="py-3 px-6 text-sm text-gray-700">
              <div>{{ cust.email }}</div>
              <div class="text-gray-500">{{ cust.phone || 'No phone' }}</div>
            </td>
            <td class="py-3 px-6 text-sm text-gray-700">
              <div class="flex flex-wrap gap-1">
                <span
                  v-for="segment in getCustomerSegments(cust.id)"
                  :key="segment.id"
                  class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium"
                  :style="{ backgroundColor: segment.color + '20', color: segment.color }"
                >
                  {{ segment.name }}
                </span>
                <span v-if="getCustomerSegments(cust.id).length === 0" class="text-gray-400 text-xs">
                  No segments
                </span>
              </div>
            </td>
            <td class="py-3 px-6 text-sm text-gray-600">{{ formatDate(cust.createdAt) }}</td>
            <td class="py-3 px-6 text-sm text-gray-700">
              <div class="flex space-x-2">
                <button
                  @click="goToCustomerDetail(cust.id)"
                  class="text-green-600 hover:text-green-800 transition-colors"
                >
                  <EyeIcon class="w-4 h-4" />
                </button>
                <button
                  @click="showSegmentModal = true; selectedCustomer = cust"
                  class="text-blue-600 hover:text-blue-800 transition-colors"
                >
                  <FolderAddIcon class="w-4 h-4" />
                </button>
                <button
                  @click="unlinkCustomer(cust.linkId)"
                  class="text-red-600 hover:text-red-800 transition-colors"
                >
                  <UserRemoveIcon class="w-4 h-4" />
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Card View -->
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
      <div
        v-for="cust in filteredCustomers"
        :key="cust.id"
        class="bg-white rounded-xl shadow-lg overflow-hidden cursor-pointer transform hover:scale-105 hover:shadow-xl transition duration-200 ease-in-out flex flex-col group"
      >
        <div class="flex flex-col items-center p-6 text-center space-y-4">
          <div
            class="h-20 w-20 bg-gradient-to-br from-green-400 to-teal-500 rounded-full flex items-center justify-center text-white text-2xl font-bold shadow-md cursor-pointer"
            @click="goToCustomerDetail(cust.id)"
          >
            {{ getInitials(cust.firstName, cust.lastName) }}
          </div>
          <div class="space-y-1">
            <h3 class="text-xl font-semibold text-gray-800 cursor-pointer" @click="goToCustomerDetail(cust.id)">
              {{ cust.firstName }} {{ cust.lastName }}
            </h3>
            <p class="text-sm text-gray-600 truncate">{{ cust.email }}</p>
            <p class="text-sm text-gray-600">{{ cust.phone || 'N/A' }}</p>
          </div>
          
          <!-- Segments -->
          <div class="flex flex-wrap gap-1 justify-center">
            <span
              v-for="segment in getCustomerSegments(cust.id)"
              :key="segment.id"
              class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium"
              :style="{ backgroundColor: segment.color + '20', color: segment.color }"
            >
              {{ segment.name }}
            </span>
          </div>
        </div>
        
        <div class="px-6 py-3 bg-gray-50 border-t border-gray-100 flex items-center justify-between text-xs text-gray-500">
          <span class="text-gray-700 font-mono">@{{ cust.username }}</span>
          <span>{{ formatDate(cust.createdAt) }}</span>
        </div>
        
        <!-- Action buttons -->
        <div class="px-6 py-2 bg-gray-50 border-t border-gray-100 flex justify-center space-x-4 opacity-0 group-hover:opacity-100 transition-opacity duration-200">
          <button
            @click="goToCustomerDetail(cust.id)"
            class="text-green-600 hover:text-green-800 transition-colors"
            title="View Details"
          >
            <EyeIcon class="w-4 h-4" />
          </button>
          <button
            @click="showSegmentModal = true; selectedCustomer = cust"
            class="text-blue-600 hover:text-blue-800 transition-colors"
            title="Manage Segments"
          >
            <FolderAddIcon class="w-4 h-4" />
          </button>
          <button
            @click="unlinkCustomer(cust.linkId)"
            class="text-red-600 hover:text-red-800 transition-colors"
            title="Remove from Shop"
          >
            <UserRemoveIcon class="w-4 h-4" />
          </button>
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

    <!-- Manage Customer Segments Modal -->
    <div v-if="showSegmentModal && selectedCustomer" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl shadow-xl max-w-lg w-full mx-4">
        <div class="p-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">
            Manage Segments for {{ selectedCustomer.firstName }} {{ selectedCustomer.lastName }}
          </h3>
          
          <div class="space-y-4">
            <div v-for="segment in segments" :key="segment.id" class="flex items-center justify-between p-3 border border-gray-200 rounded-lg">
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
                v-if="isCustomerInSegment(selectedCustomer.id, segment.id)"
                @click="removeCustomerFromSegment(selectedCustomer.id, segment.id)"
                class="text-red-600 hover:text-red-800 transition-colors"
              >
                Remove
              </button>
              <button
                v-else
                @click="addCustomerToSegment(selectedCustomer.id, segment.id)"
                class="text-green-600 hover:text-green-800 transition-colors"
              >
                Add
              </button>
            </div>
          </div>
          
          <div class="flex justify-end mt-6">
            <button
              @click="showSegmentModal = false; selectedCustomer = null"
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
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { customerService } from '@/services/customer'
import { customerSegmentService } from '@/services/customerSegment'
import {
  ViewListIcon,
  ViewGridIcon,
  RefreshIcon,
  SearchIcon,
  PlusIcon,
  UsersIcon,
  FolderIcon,
  UserAddIcon,
  UserGroupIcon,
  FolderAddIcon,
  EyeIcon,
  UserRemoveIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const shopStore = useShopStore()

// Reactive state
const customers = ref([])
const segments = ref([])
const loading = ref(false)
const error = ref(null)
const search = ref('')
const currentView = ref('list')
const selectedSegment = ref(null)
const showCreateSegmentModal = ref(false)
const showSegmentModal = ref(false)
const selectedCustomer = ref(null)
const creatingSegment = ref(false)

// New segment form
const newSegment = ref({
  name: '',
  description: '',
  color: '#3B82F6'
})

// Computed
const activeShop = computed(() => shopStore.activeShop)

const filteredCustomers = computed(() => {
  let filtered = customers.value
  
  // Filter by search term
  const term = search.value.trim().toLowerCase()
  if (term) {
    filtered = filtered.filter(c => {
      const fullName = `${c.firstName} ${c.lastName}`.toLowerCase()
      return fullName.includes(term) || (c.email?.toLowerCase().includes(term))
    })
  }
  
  // Filter by segment
  if (selectedSegment.value) {
    filtered = filtered.filter(c => isCustomerInSegment(c.id, selectedSegment.value))
  }
  
  return filtered
})

const newCustomersThisMonth = computed(() => {
  const now = new Date()
  const thisMonth = new Date(now.getFullYear(), now.getMonth(), 1)
  return customers.value.filter(c => new Date(c.createdAt) >= thisMonth).length
})

const segmentedCustomersCount = computed(() => {
  return customers.value.filter(c => getCustomerSegments(c.id).length > 0).length
})

// Fetch data
onMounted(async () => {
  console.log('🔎 customers  activeShop is:', activeShop.value)
  if (!activeShop.value?.id) {
    router.replace({ name: 'ShopSelection' })
    return
  }
  await refreshData()
})

// Refresh both customers and segments data
async function refreshData() {
  loading.value = true
  try {
    const [customersResult, segmentsResult] = await Promise.all([
      customerService.fetchAll(activeShop.value.id),
      customerSegmentService.fetchAll(activeShop.value.id)
    ])
    customers.value = Array.isArray(customersResult) ? customersResult : []
    segments.value = Array.isArray(segmentsResult) ? segmentsResult : []
  } catch (e) {
    console.error(e)
    error.value = 'Failed to load customers. Please try again later.'
  } finally {
    loading.value = false
  }
}

// Methods
function buttonClass(view) {
  return view === currentView.value
    ? 'bg-green-600 text-white hover:bg-green-700 shadow-inner'
    : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
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

function getInitials(firstName, lastName) {
  const fi = firstName?.[0]?.toUpperCase() || ''
  const li = lastName?.[0]?.toUpperCase() || ''
  return fi + li || '?'
}

function goToAddCustomer() {
  router.push({ name: 'AddCustomer' })
}

function goToCustomerDetail(customerId) {
  router.push({
    name: 'CustomerDetail',
    params: { customerId },
  })
}

function getCustomerSegments(customerId) {
  // Find segments that contain this customer
  return segments.value.filter(segment => 
    segment.customer_ids && segment.customer_ids.includes(customerId)
  )
}

function isCustomerInSegment(customerId, segmentId) {
  const segment = segments.value.find(s => s.id === segmentId)
  return segment?.customer_ids && segment.customer_ids.includes(customerId)
}

function getCustomersInSegment(segmentId) {
  const segment = segments.value.find(s => s.id === segmentId)
  if (!segment || !segment.customer_ids) return []
  return customers.value.filter(c => segment.customer_ids.includes(c.id))
}

async function createSegment() {
  if (!newSegment.value.name.trim()) return
  
  creatingSegment.value = true
  try {
    await customerSegmentService.create(activeShop.value.id, newSegment.value)
    showCreateSegmentModal.value = false
    newSegment.value = { name: '', description: '', color: '#3B82F6' }
    await refreshData() // Refresh all data
  } catch (e) {
    console.error('Failed to create segment:', e)
    error.value = 'Failed to create segment. Please try again.'
  } finally {
    creatingSegment.value = false
  }
}

async function addCustomerToSegment(customerId, segmentId) {
  try {
    await customerSegmentService.addCustomer(activeShop.value.id, segmentId, customerId)
    await refreshData() // Refresh all data
  } catch (e) {
    console.error('Failed to add customer to segment:', e)
    error.value = 'Failed to add customer to segment. Please try again.'
  }
}

async function removeCustomerFromSegment(customerId, segmentId) {
  try {
    await customerSegmentService.removeCustomer(activeShop.value.id, segmentId, customerId)
    await refreshData() // Refresh all data
  } catch (e) {
    console.error('Failed to remove customer from segment:', e)
    error.value = 'Failed to remove customer from segment. Please try again.'
  }
}

async function unlinkCustomer(linkId) {
  if (!confirm('Are you sure you want to remove this customer from your shop?')) return
  
  try {
    await customerService.delete(activeShop.value.id, linkId)
    await refreshData() // Refresh all data
  } catch (e) {
    console.error('Failed to unlink customer:', e)
    error.value = 'Failed to remove customer. Please try again.'
  }
}

function goToSegmentsPage() {
  router.push({ name: 'CustomerSegments' })
}
</script>

<style scoped>
@keyframes spin {
  from { transform: rotate(0deg); }
  to   { transform: rotate(360deg); }
}
.animate-spin {
  animation: spin 1s linear infinite;
}
</style>