<template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto font-sans">
    <!-- Header Section -->
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center mb-8">
      <div>
        <h2 class="text-3xl sm:text-4xl font-extrabold text-gray-900 leading-tight">
          Customers
        </h2>
        <p class="text-gray-600 mt-2">Manage your customer relationships and segments</p>
      </div>

      <div class="flex flex-col sm:flex-row gap-3 mt-4 sm:mt-0">
        <button
          @click="showCreateSegmentModal = true"
          class="inline-flex items-center px-4 py-2.5 bg-blue-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition duration-150 ease-in-out"
        >
          <FolderAddIcon class="w-5 h-5 mr-2 -ml-1" />
          New Segment
        </button>

        <button
          @click="goToSegmentsPage"
          class="inline-flex items-center px-4 py-2.5 bg-gray-100 text-blue-700 text-sm font-medium rounded-lg shadow-sm hover:bg-blue-100 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition duration-150 ease-in-out"
        >
          <FolderIcon class="w-5 h-5 mr-2 -ml-1" />
          Manage Segments
        </button>

        <div class="inline-flex rounded-lg shadow-sm overflow-hidden" role="group">
          <button
            @click="currentView = 'list'"
            :class="viewButtonClass('list')"
            class="px-4 py-2.5 text-sm font-medium border border-gray-300 focus:z-10 focus:ring-2 focus:ring-green-500 focus:outline-none transition-colors duration-200 ease-in-out"
          >
            <ViewListIcon class="w-5 h-5 inline-block mr-1" />
            List
          </button>
          <button
            @click="currentView = 'cards'"
            :class="viewButtonClass('cards')"
            class="px-4 py-2.5 text-sm font-medium border border-gray-300 focus:z-10 focus:ring-2 focus:ring-green-500 focus:outline-none transition-colors duration-200 ease-in-out"
          >
            <ViewGridIcon class="w-5 h-5 inline-block mr-1" />
            Cards
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
          class="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm"
        />
        <SearchIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
      </div>
    </div>

    <!-- Statistics Dashboard -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <UsersIcon class="w-6 h-6 text-green-600" />
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">Total Customers</p>
            <p class="text-2xl font-bold text-gray-900">{{ stats.total }}</p>
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
            <p class="text-2xl font-bold text-gray-900">{{ stats.segments }}</p>
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
            <p class="text-2xl font-bold text-gray-900">{{ stats.thisMonth }}</p>
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
            <p class="text-2xl font-bold text-gray-900">{{ stats.segmented }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Segment Filter -->
    <div v-if="segments.length > 0" class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 mb-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">Filter by Segment</h3>
      <div class="flex flex-wrap gap-3">
        <button
          @click="handleAllCustomersClick"
          :class="selectedSegment === null ? 'bg-green-600 text-white' : 'bg-gray-100 text-gray-700 hover:bg-gray-200'"
          class="px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-200"
        >
          All Customers
        </button>
        <button
          v-for="segment in segments"
          :key="segment.id"
          @click="() => handleSegmentClick(segment.id)"
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
    <div v-if="loading" class="space-y-4">
      <div v-if="currentView === 'list'" class="overflow-x-auto bg-white shadow-lg rounded-xl animate-pulse">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Customer</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Contact</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Segments</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Joined</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="n in 5" :key="n" class="bg-white">
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-3/4"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-1/2"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-1/3"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-1/4"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-1/3"></div></td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6 animate-pulse">
        <div v-for="n in 8" :key="n" class="bg-white rounded-xl shadow-lg overflow-hidden flex flex-col">
          <div class="p-6 flex flex-col items-center space-y-4">
            <div class="h-20 w-20 bg-gray-200 rounded-full"></div>
            <div class="space-y-2 w-full">
              <div class="h-4 bg-gray-200 rounded w-3/4 mx-auto"></div>
              <div class="h-3 bg-gray-200 rounded w-1/2 mx-auto"></div>
              <div class="h-3 bg-gray-200 rounded w-2/3 mx-auto"></div>
            </div>
          </div>
        </div>
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
        <!-- List View -->
        <div v-if="currentView === 'list'" class="overflow-x-auto bg-white shadow-lg rounded-xl">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
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
                v-for="(cust, i) in customers"
                :key="cust.id"
                class="transition duration-150 ease-in-out transform hover:scale-[1.005] hover:bg-green-50 cursor-pointer group"
                :class="{ 'bg-gray-50': i % 2 === 1 }"
                @click="goToCustomerDetail(cust.id)"
              >
                <td class="py-3 px-6 text-sm text-gray-800 font-medium">
                  <div class="flex items-center">
                    <div
                      class="h-10 w-10 rounded-full flex items-center justify-center mr-3 overflow-hidden bg-green-100"
                    >
                      <img v-if="cust.profile_image" :src="cust.profile_image" alt="Profile" class="h-10 w-10 object-cover rounded-full" />
                      <span v-else class="text-green-700 text-sm font-semibold">{{ getInitials(cust.firstName, cust.lastName) }}</span>
                    </div>
                    <div>
                      <div class="font-medium text-gray-900">
                        {{ cust.firstName }} {{ cust.lastName }}
                      </div>
                      <!-- Username removed -->
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
                <td class="py-3 px-6 text-sm text-gray-700" @click.stop>
                  <div class="flex space-x-2">
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
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Card View -->
        <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
          <div
            v-for="cust in customers"
            :key="cust.id"
            class="bg-white rounded-xl shadow-lg overflow-hidden cursor-pointer transform hover:scale-105 hover:shadow-xl transition duration-200 ease-in-out flex flex-col group"
          >
            <div class="flex flex-col items-center p-6 text-center space-y-4">
              <div
                class="h-20 w-20 rounded-full flex items-center justify-center shadow-md cursor-pointer overflow-hidden bg-gradient-to-br from-green-400 to-teal-500"
                @click="goToCustomerDetail(cust.id)"
              >
                <img v-if="cust.profile_image" :src="cust.profile_image" alt="Profile" class="h-20 w-20 object-cover rounded-full" />
                <span v-else class="text-white text-2xl font-bold">{{ getInitials(cust.firstName, cust.lastName) }}</span>
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
      </div>

      <!-- Empty State -->
      <div v-else class="bg-green-50 border border-green-200 text-green-700 px-6 py-8 rounded-lg text-center mt-8 shadow-sm">
        <div class="flex flex-col items-center">
          <div class="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mb-4">
            <UsersIcon class="w-8 h-8 text-green-600" />
          </div>
          <p class="text-lg font-medium">No customers found</p>
          <p class="mt-2 text-sm">
            <span v-if="search">No customers matching "{{ search }}"</span>
            <span v-else-if="selectedSegment">No customers in this segment</span>
            <span v-else>It looks like there are no customers for your active shop yet.</span>
          </p>
        </div>
      </div>
    </div>

    <!-- Pagination Controls (Top) -->
    <div v-if="total > pageSize && customers.length" class="flex justify-center mb-4">
      <button
        :disabled="page === 1"
        @click="handlePageChange(page - 1)"
        class="px-4 py-2 mx-1 rounded bg-gray-200 hover:bg-gray-300 disabled:opacity-50"
      >Prev</button>
      <span class="px-4 py-2 mx-1">Page {{ page }} of {{ Math.ceil(total / pageSize) }}</span>
      <button
        :disabled="page >= Math.ceil(total / pageSize)"
        @click="handlePageChange(page + 1)"
        class="px-4 py-2 mx-1 rounded bg-gray-200 hover:bg-gray-300 disabled:opacity-50"
      >Next</button>
    </div>

    <!-- Create Segment Modal -->
    <div v-if="showCreateSegmentModal" class="fixed inset-0 z-50 flex items-center justify-center bg-white/30 backdrop-blur-none transition-opacity">
      <div class="bg-white rounded-2xl shadow-2xl max-w-md w-full mx-4 border border-gray-100 relative animate-fadeIn">
        <button @click="showCreateSegmentModal = false" class="absolute top-3 right-3 text-gray-400 hover:text-gray-600 text-xl focus:outline-none">&times;</button>
        <div class="p-8">
          <h3 class="text-2xl font-bold text-gray-900 mb-6 text-center">Create New Segment</h3>
          <form @submit.prevent="createSegment" class="space-y-5">
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1">Segment Name</label>
              <input
                v-model="newSegment.name"
                type="text"
                required
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent bg-gray-50 text-gray-900"
                placeholder="e.g., VIP Customers"
              />
            </div>
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1">Description</label>
              <textarea
                v-model="newSegment.description"
                rows="3"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent bg-gray-50 text-gray-900"
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
            <div class="flex justify-end space-x-3 mt-8">
              <button
                type="button"
                @click="showCreateSegmentModal = false"
                class="px-5 py-2 text-gray-700 bg-gray-100 rounded-lg hover:bg-gray-200 font-medium transition-colors shadow-sm"
              >
                Cancel
              </button>
              <button
                type="submit"
                :disabled="creatingSegment"
                class="px-5 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 disabled:opacity-50 font-semibold transition-colors shadow-sm"
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
      <div class="bg-white rounded-2xl shadow-2xl max-w-lg w-full mx-4 border border-gray-100 relative animate-fadeIn">
        <button @click="showSegmentModal = false; selectedCustomer = null" class="absolute top-3 right-3 text-gray-400 hover:text-gray-600 text-xl focus:outline-none">&times;</button>
        <div class="p-8">
          <h3 class="text-2xl font-bold text-gray-900 mb-6 text-center">
            Manage Segments for {{ selectedCustomer.firstName }} {{ selectedCustomer.lastName }}
          </h3>
          <div class="space-y-4">
            <div v-for="segment in segments" :key="segment.id" class="flex items-center justify-between p-4 border border-gray-100 rounded-lg bg-gray-50">
              <div class="flex items-center">
                <div
                  v-if="segment.color"
                  class="w-4 h-4 rounded-full mr-3"
                  :style="{ backgroundColor: segment.color }"
                ></div>
                <div>
                  <div class="font-semibold text-gray-900">{{ segment.name }}</div>
                  <div class="text-sm text-gray-500">{{ segment.description || 'No description' }}</div>
                </div>
              </div>
              <button
                v-if="isCustomerInSegment(selectedCustomer.id, segment.id)"
                @click="removeCustomerFromSegment(selectedCustomer.id, segment.id)"
                class="px-4 py-1 text-sm rounded-lg bg-red-50 text-red-600 hover:bg-red-100 font-medium transition-colors shadow-sm"
              >
                Remove
              </button>
              <button
                v-else
                @click="addCustomerToSegment(selectedCustomer.id, segment.id)"
                class="px-4 py-1 text-sm rounded-lg bg-green-50 text-green-700 hover:bg-green-100 font-medium transition-colors shadow-sm"
              >
                Add
              </button>
            </div>
          </div>
          <div class="flex justify-end mt-8">
            <button
              @click="showSegmentModal = false; selectedCustomer = null"
              class="px-5 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 font-medium transition-colors shadow-sm"
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
import { ref, computed, onMounted, watch } from 'vue'
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
const currentView = ref('list')
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

const filteredCustomers = computed(() => customers.value) // No local filtering

const newCustomersThisMonth = computed(() => stats.value.thisMonth)
const segmentedCustomersCount = computed(() => stats.value.segmented)

// Fetch data
onMounted(async () => {
  if (!activeShop.value?.id) {
    router.replace({ name: 'ShopSelection' })
    return
  }
  await refreshAll()
})

watch([search, selectedSegment, page], async () => {
  await fetchCustomers()
})

async function refreshAll() {
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
    // If no customers, do not set error
    if (customers.value.length === 0) {
      error.value = null;
    }
  } catch (e) {
    // Only show error if it's a real fetch error, not just empty results
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
  try {
    segments.value = await customerSegmentService.fetchAll(activeShop.value.id)
  } catch (e) {
    error.value = 'Failed to load segments.'
  }
}

async function fetchStats() {
  try {
    stats.value = await customerService.fetchStats(activeShop.value.id)
  } catch (e) {
    error.value = 'Failed to load stats.'
  }
}

function viewButtonClass(view) {
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
@keyframes spin {
  from { transform: rotate(0deg); }
  to   { transform: rotate(360deg); }
}
.animate-spin {
  animation: spin 1s linear infinite;
}
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(20px) scale(0.98); }
  to   { opacity: 1; transform: translateY(0) scale(1); }
}
.animate-fadeIn {
  animation: fadeIn 0.25s cubic-bezier(0.4,0.0,0.2,1);
}
</style>