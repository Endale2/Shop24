<template>
  <div class="min-h-full bg-gray-50">
    <div class="p-4 sm:p-6 lg:p-8">
      
      <!-- Header Section -->
      <div class="mb-6 sm:mb-8">
        <div class="flex flex-col md:flex-row items-start md:items-center justify-between space-y-4 md:space-y-0">
          <div class="flex items-center mb-3 md:mb-0">
            <div class="w-10 h-10 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center mr-3 shadow-sm">
              <UsersIcon class="w-5 h-5 text-white" />
            </div>
            <div>
              <h1 class="text-2xl sm:text-3xl font-bold text-gray-900 tracking-tight">
                Customers
              </h1>
              <p class="text-sm text-gray-600 mt-1">Manage your customer relationships and segments</p>
            </div>
          </div>

          <div class="flex flex-col sm:flex-row gap-3">
            <button
              @click="showCreateSegmentModal = true"
              class="inline-flex items-center px-4 py-2.5 bg-green-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
            >
              <FolderAddIcon class="w-4 h-4 mr-1.5" />
              New Segment
            </button>

            <button
              @click="goToSegmentsPage"
              class="inline-flex items-center px-4 py-2.5 bg-gray-100 text-gray-700 text-sm font-medium rounded-lg shadow-sm hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 transition duration-150 ease-in-out"
            >
              <FolderIcon class="w-4 h-4 mr-1.5" />
              Manage Segments
            </button>
          </div>
        </div>
      </div>

      <!-- Search Bar -->
      <div class="mb-6">
        <div class="relative max-w-md">
          <input
            v-model="search"
            type="text"
            placeholder="Search by name or email..."
            class="w-full pl-10 pr-4 py-2.5 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm text-sm"
          />
          <SearchIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
        </div>
      </div>

      <!-- Statistics Dashboard -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-4">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <UsersIcon class="w-5 h-5 text-green-600" />
            </div>
            <div class="ml-3">
              <p class="text-xs font-medium text-gray-600">Total Customers</p>
              <p class="text-xl font-bold text-gray-900">{{ stats.total }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-4">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <FolderIcon class="w-5 h-5 text-blue-600" />
            </div>
            <div class="ml-3">
              <p class="text-xs font-medium text-gray-600">Segments</p>
              <p class="text-xl font-bold text-gray-900">{{ stats.segments }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-4">
          <div class="flex items-center">
            <div class="p-2 bg-yellow-100 rounded-lg">
              <UserAddIcon class="w-5 h-5 text-yellow-600" />
            </div>
            <div class="ml-3">
              <p class="text-xs font-medium text-gray-600">This Month</p>
              <p class="text-xl font-bold text-gray-900">{{ stats.thisMonth }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-4">
          <div class="flex items-center">
            <div class="p-2 bg-purple-100 rounded-lg">
              <UserGroupIcon class="w-5 h-5 text-purple-600" />
            </div>
            <div class="ml-3">
              <p class="text-xs font-medium text-gray-600">Segmented</p>
              <p class="text-xl font-bold text-gray-900">{{ stats.segmented }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Segment Filter -->
      <div v-if="segments.length > 0" class="bg-white rounded-lg shadow-sm border border-gray-200 p-4 mb-6">
        <h3 class="text-sm font-semibold text-gray-900 mb-3">Filter by Segment</h3>
        <div class="flex flex-wrap gap-2">
          <button
            @click="handleAllCustomersClick"
            :class="selectedSegment === null ? 'bg-green-600 text-white' : 'bg-gray-100 text-gray-700 hover:bg-gray-200'"
            class="px-3 py-1.5 rounded-lg text-xs font-medium transition-colors duration-200"
          >
            All Customers
          </button>
          <button
            v-for="segment in segments"
            :key="segment.id"
            @click="() => handleSegmentClick(segment.id)"
            :class="selectedSegment === segment.id ? 'bg-green-600 text-white' : 'bg-gray-100 text-gray-700 hover:bg-gray-200'"
            class="px-3 py-1.5 rounded-lg text-xs font-medium transition-colors duration-200 flex items-center"
          >
            <div
              v-if="segment.color"
              class="w-2 h-2 rounded-full mr-2"
              :style="{ backgroundColor: segment.color }"
            ></div>
            {{ segment.name }}
            <span class="ml-2 bg-gray-200 text-gray-700 px-1.5 py-0.5 rounded-full text-xs">
              {{ (segment.customer_ids || []).length }}
            </span>
          </button>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-16 text-gray-500">
        <div class="flex flex-col items-center justify-center">
          <div class="animate-spin h-8 w-8 text-green-500 mb-3 border-3 border-green-200 border-t-green-500 rounded-full"></div>
          <p class="text-sm font-medium">Loading customers...</p>
        </div>
      </div>

      <!-- Error State -->
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
        <p class="text-sm mt-2">Please ensure you have an active shop selected and try again.</p>
      </div>

      <!-- Content -->
      <div v-else>
        <div v-if="customers.length">
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="overflow-x-auto table-container">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="px-4 py-3 text-left text-xs font-semibold text-gray-900 uppercase tracking-wider">Customer</th>
                    <th class="px-4 py-3 text-left text-xs font-semibold text-gray-900 uppercase tracking-wider">Contact</th>
                    <th class="px-4 py-3 text-left text-xs font-semibold text-gray-900 uppercase tracking-wider">Segments</th>
                    <th class="px-4 py-3 text-left text-xs font-semibold text-gray-900 uppercase tracking-wider">Joined</th>
                    <th class="px-4 py-3 text-left text-xs font-semibold text-gray-900 uppercase tracking-wider">Actions</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-gray-200">
                  <tr
                    v-for="(cust, i) in customers"
                    :key="cust.id"
                    class="table-row transform hover:scale-[1.005] hover:bg-green-50 cursor-pointer group"
                    :class="{ 'bg-gray-50': i % 2 === 1 }"
                    @click="goToCustomerDetail(cust.id)"
                  >
                    <td class="py-3 px-4 text-sm text-gray-800 font-medium">
                      <div class="flex items-center">
                        <div
                          class="h-10 w-10 rounded-full flex items-center justify-center mr-3 overflow-hidden bg-green-100 customer-avatar"
                        >
                          <img 
                            v-if="cust.profile_image" 
                            :src="cust.profile_image" 
                            :alt="`${cust.firstName} ${cust.lastName} profile`"
                            class="h-10 w-10 object-cover rounded-full"
                            loading="lazy"
                            @error="handleImageError"
                            @load="handleImageLoad"
                          />
                          <span v-else class="text-green-700 text-sm font-semibold">{{ getInitials(cust.firstName, cust.lastName) }}</span>
                        </div>
                        <div>
                          <div class="font-medium text-gray-900">
                            {{ cust.firstName }} {{ cust.lastName }}
                          </div>
                        </div>
                      </div>
                    </td>
                    <td class="py-3 px-4 text-sm text-gray-700">
                      <div>{{ cust.email }}</div>
                      <div class="text-gray-500">{{ cust.phone || 'No phone' }}</div>
                    </td>
                    <td class="py-3 px-4 text-sm text-gray-700">
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
                    <td class="py-3 px-4 text-sm text-gray-600">{{ formatDate(cust.createdAt) }}</td>
                    <td class="py-3 px-4 text-sm text-gray-700" @click.stop>
                      <div class="flex space-x-2">
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
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>

        <!-- Empty State -->
        <div v-else class="text-center py-16 text-gray-400">
          <div class="max-w-md mx-auto">
            <div class="w-12 h-12 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-3">
              <UsersIcon class="w-6 h-6 text-gray-300" />
            </div>
            <p class="text-sm font-medium text-gray-600 mb-1">No customers found</p>
            <p class="text-xs">
              <span v-if="search">No customers matching "{{ search }}"</span>
              <span v-else-if="selectedSegment">No customers in this segment</span>
              <span v-else>It looks like there are no customers for your active shop yet.</span>
            </p>
          </div>
        </div>

        <!-- Pagination Controls -->
        <div v-if="total > pageSize && customers.length" class="flex justify-center mt-6">
          <button
            :disabled="page === 1"
            @click="handlePageChange(page - 1)"
            class="px-4 py-2 mx-1 rounded bg-gray-200 hover:bg-gray-300 disabled:opacity-50 text-sm"
          >Prev</button>
          <span class="px-4 py-2 mx-1 text-sm">Page {{ page }} of {{ Math.ceil(total / pageSize) }}</span>
          <button
            :disabled="page >= Math.ceil(total / pageSize)"
            @click="handlePageChange(page + 1)"
            class="px-4 py-2 mx-1 rounded bg-gray-200 hover:bg-gray-300 disabled:opacity-50 text-sm"
          >Next</button>
        </div>

        <!-- Create Segment Modal -->
        <div v-if="showCreateSegmentModal" class="fixed inset-0 z-50 flex items-center justify-center bg-white/30 backdrop-blur-none transition-opacity">
          <div class="bg-white rounded-lg shadow-xl max-w-md w-full mx-4 border border-gray-200 relative">
            <button @click="showCreateSegmentModal = false" class="absolute top-3 right-3 text-gray-400 hover:text-gray-600 text-xl focus:outline-none">&times;</button>
            <div class="p-6">
              <h3 class="text-lg font-bold text-gray-900 mb-4 text-center">Create New Segment</h3>
              <form @submit.prevent="createSegment" class="space-y-4">
                <div>
                  <label class="block text-sm font-semibold text-gray-700 mb-1">Segment Name</label>
                  <input
                    v-model="newSegment.name"
                    type="text"
                    required
                    class="w-full px-4 py-2.5 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent bg-gray-50 text-gray-900 text-sm"
                    placeholder="e.g., VIP Customers"
                  />
                </div>
                <div>
                  <label class="block text-sm font-semibold text-gray-700 mb-1">Description</label>
                  <textarea
                    v-model="newSegment.description"
                    rows="3"
                    class="w-full px-4 py-2.5 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent bg-gray-50 text-gray-900 text-sm"
                    placeholder="Optional description..."
                  ></textarea>
                </div>
                <div>
                  <label class="block text-sm font-semibold text-gray-700 mb-1">Color</label>
                  <input
                    v-model="newSegment.color"
                    type="color"
                    class="w-16 h-10 border border-gray-300 rounded-lg cursor-pointer bg-gray-50"
                  />
                </div>
                <div class="flex justify-end space-x-3 mt-6">
                  <button
                    type="button"
                    @click="showCreateSegmentModal = false"
                    class="px-4 py-2 text-gray-700 bg-gray-100 rounded-lg hover:bg-gray-200 font-medium transition-colors shadow-sm text-sm"
                  >
                    Cancel
                  </button>
                  <button
                    type="submit"
                    :disabled="creatingSegment"
                    class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 disabled:opacity-50 font-semibold transition-colors shadow-sm text-sm"
                  >
                    {{ creatingSegment ? 'Creating...' : 'Create Segment' }}
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>

        <!-- Manage Customer Segments Modal -->
        <div v-if="showSegmentModal && selectedCustomer" class="fixed inset-0 z-50 flex items-center justify-center bg-white/30 backdrop-blur-none transition-opacity">
          <div class="bg-white rounded-lg shadow-xl max-w-lg w-full mx-4 border border-gray-200 relative">
            <button @click="showSegmentModal = false; selectedCustomer = null" class="absolute top-3 right-3 text-gray-400 hover:text-gray-600 text-xl focus:outline-none">&times;</button>
            <div class="p-6">
              <h3 class="text-lg font-bold text-gray-900 mb-4 text-center">
                Manage Segments for {{ selectedCustomer.firstName }} {{ selectedCustomer.lastName }}
              </h3>
              <div class="space-y-3">
                <div v-for="segment in segments" :key="segment.id" class="flex items-center justify-between p-3 border border-gray-100 rounded-lg bg-gray-50">
                  <div class="flex items-center">
                    <div
                      v-if="segment.color"
                      class="w-3 h-3 rounded-full mr-3"
                      :style="{ backgroundColor: segment.color }"
                    ></div>
                    <div>
                      <div class="font-semibold text-gray-900 text-sm">{{ segment.name }}</div>
                      <div class="text-xs text-gray-500">{{ segment.description || 'No description' }}</div>
                    </div>
                  </div>
                  <button
                    v-if="isCustomerInSegment(selectedCustomer.id, segment.id)"
                    @click="removeCustomerFromSegment(selectedCustomer.id, segment.id)"
                    class="px-3 py-1 text-xs rounded-lg bg-red-50 text-red-600 hover:bg-red-100 font-medium transition-colors shadow-sm"
                  >
                    Remove
                  </button>
                  <button
                    v-else
                    @click="addCustomerToSegment(selectedCustomer.id, segment.id)"
                    class="px-3 py-1 text-xs rounded-lg bg-green-50 text-green-700 hover:bg-green-100 font-medium transition-colors shadow-sm"
                  >
                    Add
                  </button>
                </div>
              </div>
              <div class="flex justify-end mt-6">
                <button
                  @click="showSegmentModal = false; selectedCustomer = null"
                  class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 font-medium transition-colors shadow-sm text-sm"
                >
                  Close
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { customerService } from '@/services/customer'
import { customerSegmentService } from '@/services/customerSegment'
import {
  SearchIcon,
  PlusIcon,
  UsersIcon,
  FolderIcon,
  UserAddIcon,
  UserGroupIcon,
  FolderAddIcon,
  UserRemoveIcon,
  ExclamationCircleIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const shopStore = useShopStore()

// Reactive state
const customers = ref([])
const segments = ref([])
const loading = ref(false)
const error = ref(null)
const search = ref('')
const selectedSegment = ref(null)
const showCreateSegmentModal = ref(false)
const showSegmentModal = ref(false)
const selectedCustomer = ref(null)
const creatingSegment = ref(false)
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const stats = ref({ total: 0, segments: 0, thisMonth: 0, segmented: 0 })

// New segment form
const newSegment = ref({
  name: '',
  description: '',
  color: '#3B82F6'
})

// Computed
const activeShop = computed(() => shopStore.activeShop)

// Memoized customer segments for better performance
const customerSegmentsMap = computed(() => {
  const map = new Map()
  customers.value.forEach(customer => {
    const customerSegments = segments.value.filter(segment =>
      segment.customer_ids && segment.customer_ids.includes(customer.id)
    )
    map.set(customer.id, customerSegments)
  })
  return map
})



// Fetch data
onMounted(async () => {
  if (!activeShop.value?.id) {
    router.replace({ name: 'ShopSelection' })
    return
  }
  await refreshAll()
})

watch(activeShop, (newShop, oldShop) => {
  if (newShop?.id !== oldShop?.id) {
    refreshAll()
  }
})

watch([search, selectedSegment, page], async () => {
  await fetchCustomers()
})

async function refreshAll() {
  if (!activeShop.value?.id) {
    error.value = 'Please ensure you have an active shop selected and try again.'
    return
  }
  loading.value = true
  try {
    await Promise.all([
      fetchCustomers(),
      fetchSegments(),
      fetchStats()
    ])
  } catch (e) {
    error.value = 'Failed to load customers. Please try again later.'
  } finally {
    loading.value = false
  }
}

async function fetchCustomers() {
  if (!activeShop.value?.id) {
    error.value = 'Please ensure you have an active shop selected and try again.'
    return
  }
  loading.value = true
  try {
    error.value = null;
    const { customers: custs, total: t, page: p, pageSize: ps } = await customerService.fetchAll(
      activeShop.value.id,
      {
        page: page.value,
        limit: pageSize.value,
        search: search.value,
        segmentId: selectedSegment.value || ''
      }
    )
    customers.value = Array.isArray(custs) ? custs : [];
    total.value = typeof t === 'number' ? t : 0;
    pageSize.value = typeof ps === 'number' ? ps : 10;
    if (customers.value.length === 0) {
      error.value = null;
    }
  } catch (e) {
    if (e && e.response && (e.response.status === 404 || e.response.status === 204)) {
      customers.value = [];
      total.value = 0;
      pageSize.value = 10;
      error.value = null;
    } else {
      error.value = 'Failed to load customers. Please try again.';
    }
  } finally {
    loading.value = false;
  }
}

async function fetchSegments() {
  if (!activeShop.value?.id) {
    error.value = 'Please ensure you have an active shop selected and try again.'
    return
  }
  try {
    let rawSegments = await customerSegmentService.fetchAll(activeShop.value.id)
    if (!Array.isArray(rawSegments)) rawSegments = [];
    segments.value = rawSegments.map(seg => ({
      ...seg,
      customer_ids: Array.isArray(seg.customer_ids) ? seg.customer_ids : [],
    }))
  } catch (e) {
    error.value = 'Failed to load segments. Please ensure you have an active shop selected and try again.'
  }
}

async function fetchStats() {
  if (!activeShop.value?.id) {
    error.value = 'Please ensure you have an active shop selected and try again.'
    return
  }
  try {
    stats.value = await customerService.fetchStats(activeShop.value.id)
  } catch (e) {
    error.value = 'Failed to load stats. Please ensure you have an active shop selected and try again.'
  }
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

// Image optimization functions
function handleImageError(event) {
  // Hide the image and show initials instead
  event.target.style.display = 'none'
  const initialsSpan = event.target.nextElementSibling
  if (initialsSpan) {
    initialsSpan.style.display = 'flex'
  }
}

function handleImageLoad(event) {
  // Image loaded successfully, ensure initials are hidden
  const initialsSpan = event.target.nextElementSibling
  if (initialsSpan) {
    initialsSpan.style.display = 'none'
  }
}



function goToCustomerDetail(customerId) {
  router.push({
    name: 'CustomerDetail',
    params: { customerId },
  })
}

function getCustomerSegments(customerId) {
  return customerSegmentsMap.value.get(customerId) || []
}

function isCustomerInSegment(customerId, segmentId) {
  const segment = segments.value.find(s => s.id === segmentId)
  return segment?.customer_ids && segment.customer_ids.includes(customerId)
}



async function createSegment() {
  if (!newSegment.value.name.trim()) return
  creatingSegment.value = true
  try {
    await customerSegmentService.create(activeShop.value.id, newSegment.value)
    showCreateSegmentModal.value = false
    newSegment.value = { name: '', description: '', color: '#3B82F6' }
    await refreshAll()
  } catch (e) {
    error.value = 'Failed to create segment. Please try again.'
  } finally {
    creatingSegment.value = false
  }
}

async function addCustomerToSegment(customerId, segmentId) {
  try {
    await customerSegmentService.addCustomer(activeShop.value.id, segmentId, customerId)
    await refreshAll()
  } catch (e) {
    error.value = 'Failed to add customer to segment. Please try again.'
  }
}

async function removeCustomerFromSegment(customerId, segmentId) {
  try {
    await customerSegmentService.removeCustomer(activeShop.value.id, segmentId, customerId)
    await refreshAll()
  } catch (e) {
    error.value = 'Failed to remove customer from segment. Please try again.'
  }
}

async function unlinkCustomer(linkId) {
  if (!confirm('Are you sure you want to remove this customer from your shop?')) return
  try {
    await customerService.delete(activeShop.value.id, linkId)
    await refreshAll()
  } catch (e) {
    error.value = 'Failed to remove customer. Please try again.'
  }
}

function goToSegmentsPage() {
  router.push({ name: 'CustomerSegments' })
}

function handlePageChange(newPage) {
  if (newPage !== page.value) {
    page.value = newPage
  }
}

// When clicking 'All Customers', reset segment, page, and search
function handleAllCustomersClick() {
  selectedSegment.value = null;
  page.value = 1;
}
// When clicking a segment, set segment, reset page and search
function handleSegmentClick(segmentId) {
  selectedSegment.value = segmentId;
  page.value = 1;
}
// When searching, reset page
watch(search, () => {
  page.value = 1;
});
</script>

<style scoped>
/* Custom styles for enhanced design */
@media (max-width: 640px) {
  .min-w-0 { min-width: 0 !important; }
}

/* Enhanced hover effects */
.group:hover .group-hover\:scale-105 {
  transform: scale(1.05);
}

/* Optimized transitions - only animate necessary properties */
.table-row {
  transition: background-color 150ms ease-in-out, transform 150ms ease-in-out;
}

/* Image optimization */
.customer-avatar {
  will-change: transform;
  backface-visibility: hidden;
}

/* Performance optimizations */
.table-container {
  contain: layout style paint;
}

/* Custom scrollbar for better UX */
::-webkit-scrollbar {
  width: 4px;
}

::-webkit-scrollbar-track {
  background: transparent;
}

::-webkit-scrollbar-thumb {
  background: #d1d5db;
  border-radius: 2px;
}

::-webkit-scrollbar-thumb:hover {
  background: #9ca3af;
}

/* Reduce motion for accessibility */
@media (prefers-reduced-motion: reduce) {
  .table-row {
    transition: none;
  }
  
  .group:hover .group-hover\:scale-105 {
    transform: none;
  }
}
</style>