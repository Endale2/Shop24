<template>
  <div class="min-h-full bg-gray-50">
    <div class="p-4 sm:p-6 lg:p-8">
      
      <!-- Header Section -->
      <div class="mb-6 sm:mb-8">
        <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center">
          <div class="flex items-center mb-3 sm:mb-0">
            <div class="w-10 h-10 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center mr-3 shadow-sm">
              <TagIcon class="w-5 h-5 text-white" />
            </div>
            <div>
              <h1 class="text-2xl sm:text-3xl font-bold text-gray-900 tracking-tight">
                Discounts
              </h1>
              <p class="text-sm text-gray-600 mt-1">Manage your promotional offers</p>
            </div>
          </div>
          
          <div class="flex flex-col sm:flex-row gap-2 sm:gap-3 mt-4 sm:mt-0">
            <button
              @click="refreshDiscounts"
              :disabled="loading"
              class="inline-flex items-center px-3 py-2 bg-gray-100 text-gray-700 text-xs font-medium rounded-lg shadow-sm hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 transition duration-150 ease-in-out"
            >
              <RefreshIcon class="w-4 h-4 mr-1.5" :class="{ 'animate-spin': loading }" />
              Refresh
            </button>
            <button
              @click="goToAddPage"
              class="inline-flex items-center px-3 py-2 bg-green-600 text-white text-xs font-medium rounded-lg shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
            >
              <PlusIcon class="w-4 h-4 mr-1.5" />
              Create Discount
            </button>
          </div>
        </div>
      </div>

      <!-- Search and Filter Section -->
      <div class="mb-6">
        <div class="flex flex-col lg:flex-row gap-4">
          <!-- Search Bar -->
          <div class="relative max-w-md">
            <input
              type="text"
              v-model="searchQuery"
              @input="debouncedSearch"
              placeholder="Search discounts by name, description, or coupon code..."
              class="w-full pl-9 pr-4 py-2.5 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm text-sm"
            />
            <SearchIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 w-4 h-4 text-gray-400" />
          </div>

          <!-- Filters -->
          <div class="flex flex-wrap gap-2">
            <select
              v-model="statusFilter"
              class="px-3 py-2.5 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm bg-white text-sm"
            >
              <option value="">All Status</option>
              <option value="active">Active</option>
              <option value="inactive">Inactive</option>
              <option value="expired">Expired</option>
              <option value="upcoming">Upcoming</option>
            </select>

            <select
              v-model="typeFilter"
              class="px-3 py-2.5 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm bg-white text-sm"
            >
              <option value="">All Types</option>
              <option value="percentage">Percentage</option>
              <option value="fixed">Fixed Amount</option>
            </select>

            <select
              v-model="eligibilityFilter"
              class="px-3 py-2.5 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm bg-white text-sm"
            >
              <option value="">All Eligibility</option>
              <option value="all">Everyone</option>
              <option value="specific">Specific Customers</option>
              <option value="segment">Customer Segments</option>
            </select>
          </div>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 animate-pulse">
          <div v-for="n in 6" :key="n" class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="p-4 space-y-3">
              <div class="flex justify-between items-start">
                <div class="space-y-2 flex-1">
                  <div class="h-4 bg-gray-200 rounded w-3/4"></div>
                  <div class="h-3 bg-gray-200 rounded w-1/2"></div>
                </div>
                <div class="h-5 bg-gray-200 rounded w-14"></div>
              </div>
              <div class="h-6 bg-gray-200 rounded w-1/3"></div>
              <div class="space-y-2">
                <div class="h-3 bg-gray-200 rounded w-1/2"></div>
                <div class="h-3 bg-gray-200 rounded w-2/3"></div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-else-if="filteredDiscounts.length === 0" class="text-center py-12">
        <div class="max-w-md mx-auto">
          <div class="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-4">
            <TagIcon class="w-8 h-8 text-green-600" />
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mb-2">
            {{ searchQuery || statusFilter || typeFilter || eligibilityFilter ? 'No discounts found' : 'No discounts created yet' }}
          </h3>
          <p class="text-sm text-gray-600 mb-4">
            {{ searchQuery || statusFilter || typeFilter || eligibilityFilter 
              ? 'Try adjusting your search criteria or filters to find what you\'re looking for.' 
              : 'Create your first discount to start offering promotions to your customers!' }}
          </p>
          <div class="flex flex-col sm:flex-row gap-2 justify-center">
            <button
              v-if="searchQuery || statusFilter || typeFilter || eligibilityFilter"
              @click="clearFilters"
              class="inline-flex items-center px-3 py-2 bg-gray-100 text-gray-700 text-xs font-medium rounded-lg hover:bg-gray-200 transition duration-150 ease-in-out"
            >
              <XIcon class="w-3 h-3 mr-1.5" />
              Clear Filters
            </button>
            <button
              v-if="!searchQuery && !statusFilter && !typeFilter && !eligibilityFilter"
              @click="goToAddPage"
              class="inline-flex items-center px-3 py-2 bg-green-600 text-white text-xs font-medium rounded-lg hover:bg-green-700 transition duration-150 ease-in-out"
            >
              <PlusIcon class="w-3 h-3 mr-1.5" />
              Create First Discount
            </button>
          </div>
        </div>
      </div>

      <!-- Discounts Grid -->
      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        <div
          v-for="discount in filteredDiscounts"
          :key="discount.id || discount._id"
          @click="discount.id || discount._id ? goToDiscountDetail(discount.id || discount._id) : null"
          :class="[
            'bg-white rounded-lg shadow-sm border border-gray-200 transition-all duration-200 ease-in-out transform overflow-hidden group',
            (discount.id || discount._id) ? 'hover:shadow-md cursor-pointer hover:-translate-y-0.5' : 'opacity-50 cursor-not-allowed'
          ]"
        >
          <!-- Discount Header -->
          <div class="p-4 border-b border-gray-100">
            <div class="flex items-start justify-between mb-3">
              <div class="flex-1">
                <h3 class="text-sm font-semibold text-gray-900 group-hover:text-green-700 transition-colors duration-150">
                  {{ discount.name }}
                </h3>
                <p class="text-xs text-gray-500 mt-1 line-clamp-2">{{ discount.description || 'No description' }}</p>
              </div>
              <div class="flex items-center space-x-1.5">
                <span
                  :class="getStatusClass(discount)"
                  class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium"
                >
                  <CheckCircleIcon v-if="isActive(discount)" class="w-3 h-3 mr-1" />
                  <XCircleIcon v-else-if="!discount.active" class="w-3 h-3 mr-1" />
                  <ClockIcon v-else-if="isUpcoming(discount)" class="w-3 h-3 mr-1" />
                  <ExclamationIcon v-else class="w-3 h-3 mr-1" />
                  {{ getStatusText(discount) }}
                </span>
                <span v-if="!(discount.id || discount._id)" class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-red-100 text-red-800">
                  <ExclamationIcon class="w-3 h-3 mr-1" />
                  No ID
                </span>
              </div>
            </div>

            <!-- Coupon Code -->
            <div v-if="discount.couponCode" class="mb-3">
              <span class="inline-flex items-center px-2.5 py-1 rounded-md text-xs font-medium bg-blue-100 text-blue-800 border border-blue-200">
                <CodeIcon class="w-3 h-3 mr-1.5" />
                {{ discount.couponCode }}
              </span>
            </div>

            <!-- Discount Value -->
            <div class="flex items-center justify-between">
              <div class="flex items-center space-x-2">
                <span :class="{
                  'bg-green-100 text-green-800': discount.type === 'percentage',
                  'bg-purple-100 text-purple-800': discount.type === 'fixed'
                }" class="inline-flex px-2 py-0.5 rounded-full text-xs font-medium">
                  {{ formatType(discount.type) }}
                </span>
                <span class="text-xl font-bold text-green-700">
                  <span v-if="discount.type === 'percentage'">{{ discount.value }}%</span>
                  <span v-else>${{ discount.value?.toFixed(2) }}</span>
                </span>
              </div>
            </div>
          </div>

          <!-- Discount Details -->
          <div class="p-4 space-y-3">
            <!-- Eligibility -->
            <div class="flex items-center justify-between">
              <span class="text-xs text-gray-600">Eligibility:</span>
              <span :class="{
                'bg-green-100 text-green-800': discount.eligibilityType === 'all',
                'bg-blue-100 text-blue-800': discount.eligibilityType === 'specific',
                'bg-purple-100 text-purple-800': discount.eligibilityType === 'segment'
              }" class="inline-flex px-2 py-0.5 rounded-full text-xs font-medium">
                {{ formatEligibility(discount.eligibilityType) }}
              </span>
            </div>

            <!-- Usage -->
            <div class="flex items-center justify-between">
              <span class="text-xs text-gray-600">Usage:</span>
              <span class="text-xs font-medium text-gray-900">
                {{ discount.currentUsage || 0 }}
                <span v-if="discount.usageLimit" class="text-gray-500">/ {{ discount.usageLimit }}</span>
              </span>
            </div>

            <!-- Validity -->
            <div class="space-y-1.5">
              <div class="flex items-center justify-between">
                <span class="text-xs text-gray-600">Start:</span>
                <span class="text-xs text-gray-900">{{ formatDate(discount.startAt) }}</span>
              </div>
              <div class="flex items-center justify-between">
                <span class="text-xs text-gray-600">End:</span>
                <span class="text-xs text-gray-900">{{ formatDate(discount.endAt) }}</span>
              </div>
            </div>

            <!-- Products Applied -->
            <div v-if="discount.appliesToProducts?.length" class="pt-2 border-t border-gray-100">
              <div class="flex items-center justify-between">
                <span class="text-xs text-gray-600">Products:</span>
                <span class="text-xs font-medium text-gray-900">{{ discount.appliesToProducts.length }} selected</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Results Summary -->
      <div v-if="filteredDiscounts.length > 0" class="mt-6 text-center">
        <div class="text-xs text-gray-500">
          Showing {{ filteredDiscounts.length }} of {{ discounts.length }} discounts
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { discountService } from '@/services/discount'
import { format } from 'date-fns'
import {
  PlusIcon,
  RefreshIcon,
  CheckCircleIcon,
  XCircleIcon,
  CalendarIcon,
  TagIcon,
  CodeIcon,
  SearchIcon,
  ClockIcon,
  ExclamationIcon,
  XIcon
} from '@heroicons/vue/outline'

const shopStore = useShopStore()
const shop = computed(() => shopStore.activeShop)
const shopId = computed(() => shop.value?.id)
const router = useRouter()

// State
const discounts = ref([])
const loading = ref(false)
const searchQuery = ref('')
const statusFilter = ref('')
const typeFilter = ref('')
const eligibilityFilter = ref('')

// Computed
const filteredDiscounts = computed(() => {
  let filtered = discounts.value

  // Search filter
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(d => 
      d.name.toLowerCase().includes(query) ||
      d.description?.toLowerCase().includes(query) ||
      d.couponCode?.toLowerCase().includes(query)
    )
  }

  // Status filter
  if (statusFilter.value) {
    filtered = filtered.filter(d => {
      switch (statusFilter.value) {
        case 'active': return isActive(d)
        case 'inactive': return !d.active
        case 'expired': return isExpired(d)
        case 'upcoming': return isUpcoming(d)
        default: return true
      }
    })
  }

  // Type filter
  if (typeFilter.value) {
    filtered = filtered.filter(d => d.type === typeFilter.value)
  }

  // Eligibility filter
  if (eligibilityFilter.value) {
    filtered = filtered.filter(d => d.eligibilityType === eligibilityFilter.value)
  }

  return filtered
})

// Helpers
function formatDateTime(iso) {
  if (!iso) return 'Not set';
  try {
    const date = new Date(iso);
    if (isNaN(date.getTime())) {
      return 'Invalid Date';
    }
    return format(date, 'MMM d, yyyy HH:mm');
  } catch (e) {
    console.error("Error formatting date:", e);
    return 'Invalid Date';
  }
}

function formatDate(iso) {
  if (!iso) return 'Not set';
  try {
    const date = new Date(iso);
    if (isNaN(date.getTime())) {
      return 'Invalid Date';
    }
    return format(date, 'MMM d, yyyy');
  } catch (e) {
    console.error("Error formatting date:", e);
    return 'Invalid Date';
  }
}

function formatType(type) {
  const types = {
    'percentage': 'Percentage',
    'fixed': 'Fixed Amount'
  };
  return types[type] || type;
}

function formatEligibility(eligibility) {
  const eligibilities = {
    'all': 'Everyone',
    'specific': 'Specific Customers',
    'segment': 'Customer Segments'
  };
  return eligibilities[eligibility] || eligibility;
}

function isActive(discount) {
  if (!discount.active) return false;
  const now = new Date();
  const startAt = new Date(discount.startAt);
  const endAt = new Date(discount.endAt);
  return now >= startAt && now <= endAt;
}

function isExpired(discount) {
  const now = new Date();
  const endAt = new Date(discount.endAt);
  return now > endAt;
}

function isUpcoming(discount) {
  const now = new Date();
  const startAt = new Date(discount.startAt);
  return now < startAt;
}

function getStatusClass(discount) {
  if (isActive(discount)) {
    return 'bg-green-100 text-green-800';
  } else if (!discount.active) {
    return 'bg-gray-100 text-gray-800';
  } else if (isUpcoming(discount)) {
    return 'bg-blue-100 text-blue-800';
  } else {
    return 'bg-red-100 text-red-800';
  }
}

function getStatusText(discount) {
  if (isActive(discount)) {
    return 'Active';
  } else if (!discount.active) {
    return 'Inactive';
  } else if (isUpcoming(discount)) {
    return 'Upcoming';
  } else {
    return 'Expired';
  }
}

// Debounced search
let searchTimeout;
function debouncedSearch() {
  clearTimeout(searchTimeout);
  searchTimeout = setTimeout(() => {
    // Search is handled by computed property
  }, 300);
}

// Navigation
function goToDiscountDetail(discountId) {
  console.log('Navigating to discount detail with ID:', discountId);
  if (!discountId) {
    console.error('Discount ID is undefined or null');
    return;
  }
  // Ensure we have a valid string ID
  const id = String(discountId);
  if (id === 'undefined' || id === 'null') {
    console.error('Invalid discount ID:', discountId);
    return;
  }
  router.push({ name: 'DiscountDetail', params: { discountId: id } });
}

function goToAddPage() {
  router.push('/dashboard/discounts/create')
}

function clearFilters() {
  searchQuery.value = ''
  statusFilter.value = ''
  typeFilter.value = ''
  eligibilityFilter.value = ''
}

// API calls
async function loadDiscounts() {
  if (!shopId.value) return;
  
  loading.value = true;
  try {
    discounts.value = await discountService.fetchAllByShop(shopId.value);
    // Check if all discounts have IDs
    discounts.value.forEach((discount, index) => {
      if (!discount.id && !discount._id) {
        console.error(`Discount at index ${index} is missing ID:`, discount);
      }
    });
  } catch (err) {
    console.error('Failed to load discounts:', err);
  } finally {
    loading.value = false;
  }
}

async function refreshDiscounts() {
  await loadDiscounts();
}

// Lifecycle (align with OrdersPage pattern)
onMounted(async () => {
  try {
    if (!shopId.value) {
      const ensured = await shopStore.ensureActiveShop()
      if (!ensured?.id) {
        router.replace({ name: 'ShopSelection' })
        return
      }
    }
    await loadDiscounts()
  } catch (e) {
    console.error('[DiscountsPage] Initialization failed:', e)
  }
});
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* Custom styles for enhanced design */
@media (max-width: 640px) {
  .min-w-0 { min-width: 0 !important; }
}

/* Enhanced hover effects */
.group:hover .group-hover\:scale-105 {
  transform: scale(1.05);
}

/* Smooth transitions */
* {
  transition-property: color, background-color, border-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
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
</style>