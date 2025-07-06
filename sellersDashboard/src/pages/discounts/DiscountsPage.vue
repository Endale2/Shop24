<template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto font-sans">
    <div class="flex justify-between items-center mb-8">
      <h1 class="text-3xl sm:text-4xl font-extrabold text-gray-900 leading-tight">Discounts</h1>
      <button
        @click="openCreateForm"
        class="inline-flex items-center px-5 py-2.5 bg-green-600 text-white text-base font-medium rounded-lg shadow-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out transform hover:scale-105"
      >
        <PlusIcon class="w-5 h-5 mr-2 -ml-1" />
        New Discount
      </button>
    </div>

    <!-- Search and Filter Section -->
    <div class="flex flex-col sm:flex-row justify-between items-center mb-6 space-y-4 sm:space-y-0">
      <div class="w-full sm:w-1/2 relative">
        <input
          type="text"
          v-model="searchQuery"
          @input="debouncedSearch"
          placeholder="Search discounts..."
          class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm"
        />
        <SearchIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
      </div>

      <div class="flex space-x-4 items-center">
        <select
          v-model="statusFilter"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm"
        >
          <option value="">All Status</option>
          <option value="active">Active</option>
          <option value="inactive">Inactive</option>
          <option value="expired">Expired</option>
          <option value="upcoming">Upcoming</option>
        </select>

        <select
          v-model="categoryFilter"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm"
        >
          <option value="">All Categories</option>
          <option value="product">Product</option>
          <option value="order">Order</option>
          <option value="shipping">Shipping</option>
          <option value="buy_x_get_y">Buy X Get Y</option>
        </select>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="text-center py-16 text-gray-600 text-lg">
      <div class="flex flex-col items-center justify-center">
        <SpinnerIcon class="animate-spin h-10 w-10 text-green-500 mb-4" />
        <p>Loading discounts...</p>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else-if="filteredDiscounts.length === 0" class="bg-green-50 border border-green-200 text-green-700 px-6 py-10 rounded-xl text-center shadow-md">
      <div class="flex flex-col items-center">
        <div class="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mb-4">
          <TagIcon class="w-8 h-8 text-green-600" />
        </div>
        <p class="text-lg font-medium mb-3">
          {{ searchQuery || statusFilter || categoryFilter ? 'No discounts found' : 'No discounts created yet.' }}
        </p>
        <p class="text-base mb-4">
          {{ searchQuery || statusFilter || categoryFilter ? 'Try adjusting your search or filters.' : 'Click the "New Discount" button to add your first discount!' }}
        </p>
        <button
          v-if="!searchQuery && !statusFilter && !categoryFilter"
          @click="openCreateForm"
          class="inline-flex items-center px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg shadow-md hover:bg-green-700 transition duration-150 ease-in-out"
        >
          <PlusIcon class="w-4 h-4 mr-2" />
          Create First Discount
        </button>
      </div>
    </div>

    <!-- Discounts Table -->
    <div v-else class="bg-white shadow-xl rounded-xl overflow-hidden border border-gray-200">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-100">
            <tr>
              <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Discount</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Type & Value</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Eligibility</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Usage</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Status</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Validity</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-100">
            <tr
              v-for="d in filteredDiscounts"
              :key="d.id"
              class="bg-white hover:bg-green-50 transition-colors duration-150 ease-in-out"
            >
              <td class="px-6 py-4">
                <div class="flex items-start space-x-3">
                  <div class="flex-shrink-0">
                    <div class="w-10 h-10 bg-green-100 rounded-lg flex items-center justify-center">
                      <TagIcon class="w-5 h-5 text-green-600" />
                    </div>
                  </div>
                  <div class="flex-1 min-w-0">
                    <div class="text-lg font-semibold text-gray-900 hover:text-green-700 cursor-pointer" @click="goToDiscountDetail(d.id)">
                      {{ d.name }}
                    </div>
                    <p class="text-sm text-gray-500 mt-1 truncate max-w-xs">{{ d.description || 'No description' }}</p>
                    <div v-if="d.couponCode" class="mt-2">
                      <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                        <CodeIcon class="w-3 h-3 mr-1" />
                        {{ d.couponCode }}
                      </span>
                    </div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="space-y-1">
                  <div class="flex items-center space-x-2">
                    <span :class="{
                      'bg-blue-100 text-blue-800': d.category === 'product',
                      'bg-purple-100 text-purple-800': d.category === 'order',
                      'bg-orange-100 text-orange-800': d.category === 'shipping',
                      'bg-pink-100 text-pink-800': d.category === 'buy_x_get_y'
                    }" class="inline-flex px-2.5 py-0.5 rounded-full text-xs font-medium">
                      {{ formatCategory(d.category) }}
                    </span>
                    <span class="text-sm text-gray-600">{{ formatType(d.type) }}</span>
                  </div>
                  <div class="text-lg font-bold text-green-700">
                    <span v-if="d.type === 'percentage'">{{ d.value }}%</span>
                    <span v-else>${{ d.value.toFixed(2) }}</span>
                  </div>
                  <!-- Buy X Get Y specific info -->
                  <div v-if="d.category === 'buy_x_get_y'" class="text-xs text-gray-600">
                    <div>Buy {{ d.buyQuantity || 0 }}x → Get {{ d.getQuantity || 0 }}x</div>
                    <div v-if="d.buyProductIds?.length" class="text-gray-500">
                      {{ d.buyProductIds.length }} buy products
                    </div>
                    <div v-if="d.getProductIds?.length" class="text-gray-500">
                      {{ d.getProductIds.length }} get products
                    </div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="space-y-1">
                  <div class="flex items-center space-x-2">
                    <span :class="{
                      'bg-green-100 text-green-800': d.eligibilityType === 'all',
                      'bg-blue-100 text-blue-800': d.eligibilityType === 'specific',
                      'bg-purple-100 text-purple-800': d.eligibilityType === 'segment'
                    }" class="inline-flex px-2.5 py-0.5 rounded-full text-xs font-medium">
                      {{ formatEligibility(d.eligibilityType) }}
                    </span>
                  </div>
                  <div v-if="d.minimumOrderSubtotal" class="text-sm text-gray-600">
                    Min: ${{ d.minimumOrderSubtotal.toFixed(2) }}
                  </div>
                  <div v-if="d.freeShipping" class="text-sm text-green-600 font-medium">
                    Free Shipping
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="space-y-1">
                  <div class="text-sm text-gray-900">
                    {{ d.currentUsage || 0 }}
                    <span v-if="d.usageLimit" class="text-gray-500">/ {{ d.usageLimit }}</span>
                  </div>
                  <div v-if="d.perCustomerLimit" class="text-xs text-gray-500">
                    {{ d.perCustomerLimit }} per customer
                  </div>
                  <div v-if="d.usageTracking?.length" class="text-xs text-gray-500">
                    {{ d.usageTracking.length }} active users
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="flex items-center space-x-2">
                  <span
                    :class="getStatusClass(d)"
                    class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium"
                  >
                    <CheckCircleIcon v-if="isActive(d)" class="w-4 h-4 mr-1" />
                    <XCircleIcon v-else-if="!d.active" class="w-4 h-4 mr-1" />
                    <ClockIcon v-else-if="isUpcoming(d)" class="w-4 h-4 mr-1" />
                    <ExclamationIcon v-else class="w-4 h-4 mr-1" />
                    {{ getStatusText(d) }}
                  </span>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="space-y-1">
                  <div class="flex items-center text-sm text-gray-600">
                    <CalendarIcon class="w-4 h-4 mr-1 text-gray-400" />
                    {{ formatDate(d.startAt) }}
                  </div>
                  <div class="flex items-center text-sm text-gray-600">
                    <CalendarIcon class="w-4 h-4 mr-1 text-gray-400" />
                    {{ formatDate(d.endAt) }}
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="flex items-center space-x-2">
                  <button
                    @click="goToDiscountDetail(d.id)"
                    class="inline-flex items-center px-3 py-1.5 bg-blue-600 text-white text-xs font-medium rounded-md hover:bg-blue-700 transition duration-150 ease-in-out"
                  >
                    <EyeIcon class="w-3 h-3 mr-1" />
                    View
                  </button>
                  <button
                    @click="openEditForm(d)"
                    class="inline-flex items-center px-3 py-1.5 bg-green-600 text-white text-xs font-medium rounded-md hover:bg-green-700 transition duration-150 ease-in-out"
                  >
                    <PencilIcon class="w-3 h-3 mr-1" />
                    Edit
                  </button>
                  <button
                    @click="confirmDelete(d)"
                    class="inline-flex items-center px-3 py-1.5 bg-red-600 text-white text-xs font-medium rounded-md hover:bg-red-700 transition duration-150 ease-in-out"
                  >
                    <TrashIcon class="w-3 h-3 mr-1" />
                    Delete
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
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
  RefreshIcon as SpinnerIcon,
  CheckCircleIcon,
  XCircleIcon,
  CalendarIcon,
  PencilIcon,
  TrashIcon,
  EyeIcon,
  TagIcon,
  CodeIcon,
  SearchIcon,
  ClockIcon,
  ExclamationIcon
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
const categoryFilter = ref('')

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

  // Category filter
  if (categoryFilter.value) {
    filtered = filtered.filter(d => d.category === categoryFilter.value)
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

function formatCategory(category) {
  const categories = {
    'product': 'Product',
    'order': 'Order',
    'shipping': 'Shipping',
    'buy_x_get_y': 'Buy X Get Y'
  };
  return categories[category] || category;
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
  router.push({ name: 'DiscountDetail', params: { discountId: discountId } });
}

function openCreateForm() {
  router.push('/dashboard/discounts/create')
}

function openEditForm(discount) {
  router.push({ name: 'DiscountDetail', params: { discountId: discount.id } });
}

function confirmDelete(discount) {
  if (confirm(`Are you sure you want to delete "${discount.name}"? This action cannot be undone.`)) {
    deleteDiscount(discount.id);
  }
}

// API calls
async function loadDiscounts() {
  if (!shopId.value) return;
  
  loading.value = true;
  try {
    discounts.value = await discountService.fetchAll(shopId.value);
  } catch (err) {
    console.error('Failed to load discounts:', err);
  } finally {
    loading.value = false;
  }
}

async function deleteDiscount(discountId) {
  if (!shopId.value) return;
  
  try {
    await discountService.delete(shopId.value, discountId);
    await loadDiscounts(); // Reload the list
  } catch (err) {
    console.error('Failed to delete discount:', err);
    alert('Failed to delete discount: ' + (err.response?.data?.message || err.message));
  }
}

// Lifecycle
onMounted(() => {
  if (shopId.value) {
    loadDiscounts();
  }
});
</script>

<style scoped>
/* Custom styles if needed */
</style>