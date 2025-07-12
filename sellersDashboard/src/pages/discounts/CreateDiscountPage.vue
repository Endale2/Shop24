<template>
  <div class="p-4 sm:p-6 max-w-4xl mx-auto font-sans">
    <!-- Header -->
    <div class="mb-8">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl sm:text-4xl font-extrabold text-gray-900 leading-tight">
            Create New Discount
          </h1>
          <p class="text-lg text-gray-600 mt-2">
            Set up product-level promotional offers to boost your sales
          </p>
        </div>
        <router-link
          to="/dashboard/discounts"
          class="inline-flex items-center px-4 py-2 bg-gray-100 text-gray-700 text-sm font-medium rounded-lg hover:bg-gray-200 transition-colors duration-200"
        >
          <ArrowLeftIcon class="w-4 h-4 mr-2" />
          Back to Discounts
        </router-link>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="text-center py-16 text-gray-600 text-lg">
      <div class="flex flex-col items-center justify-center">
        <div class="animate-spin h-12 w-12 text-green-500 mb-4 border-4 border-green-200 border-t-green-500 rounded-full"></div>
        <p class="text-lg font-medium">Loading form data...</p>
        <p class="text-sm text-gray-500 mt-2">Please wait while we prepare the form</p>
      </div>
    </div>

    <!-- Form -->
    <div v-else class="bg-white rounded-xl shadow-lg border border-gray-200">
      <form @submit.prevent="submitForm" class="p-8 space-y-8">
        <!-- Basic Information -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h3 class="text-lg font-medium text-gray-900 mb-4 flex items-center">
            <TagIcon class="w-5 h-5 mr-2 text-green-600" />
            Basic Information
          </h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label for="name" class="block text-sm font-medium text-gray-700 mb-1">Discount Name *</label>
              <input 
                id="name" 
                v-model="form.name" 
                type="text" 
                required
                placeholder="e.g., Summer Sale 20% Off"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out" 
              />
            </div>
            <div>
              <label for="couponCode" class="block text-sm font-medium text-gray-700 mb-1">Coupon Code (Optional)</label>
              <input 
                id="couponCode" 
                v-model="form.couponCode" 
                type="text"
                placeholder="e.g., SUMMER20"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out" 
              />
            </div>
          </div>
          <div class="mt-4">
            <label for="description" class="block text-sm font-medium text-gray-700 mb-1">Description</label>
            <textarea 
              id="description" 
              v-model="form.description" 
              rows="3"
              placeholder="Describe what this discount offers..."
              class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
            ></textarea>
          </div>
        </div>

        <!-- Discount Configuration -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h3 class="text-lg font-medium text-gray-900 mb-4 flex items-center">
            <CurrencyDollarIcon class="w-5 h-5 mr-2 text-green-600" />
            Discount Configuration
          </h3>
          <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div>
              <label for="type" class="block text-sm font-medium text-gray-700 mb-1">Discount Type *</label>
              <select 
                id="type" 
                v-model="form.type" 
                required
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 bg-white transition duration-150 ease-in-out"
              >
                <option value="percentage">Percentage Off</option>
                <option value="fixed">Fixed Amount Off</option>
              </select>
            </div>
            <div>
              <label for="value" class="block text-sm font-medium text-gray-700 mb-1">Discount Value *</label>
              <div class="relative">
                <input
                  id="value"
                  v-model.number="form.value"
                  type="number"
                  min="0"
                  :max="form.type === 'percentage' ? 100 : null"
                  :step="form.type === 'percentage' ? 1 : 0.01"
                  required
                  :placeholder="form.type === 'percentage' ? 'e.g., 20' : 'e.g., 10.00'"
                  class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
                />
                <div class="absolute inset-y-0 right-0 flex items-center pr-3 pointer-events-none">
                  <span class="text-gray-500 text-sm">
                    {{ form.type === 'percentage' ? '%' : '$' }}
                  </span>
                </div>
              </div>
            </div>
            <div class="md:col-span-3">
              <div class="bg-blue-50 border border-blue-200 rounded-lg p-4">
                <div class="flex items-start">
                  <InformationCircleIcon class="w-5 h-5 text-blue-600 mr-2 mt-0.5 flex-shrink-0" />
                  <div class="text-sm text-blue-800">
                    <p class="font-medium mb-1">Discount Preview:</p>
                    <p v-if="form.type === 'percentage' && form.value">
                      Customers will save {{ form.value }}% on eligible products
                    </p>
                    <p v-else-if="form.type === 'fixed' && form.value">
                      Customers will save ${{ form.value?.toFixed(2) }} on eligible products
                    </p>
                    <p v-else class="text-blue-600">
                      Enter a value to see the discount preview
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Product Selection -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h3 class="text-lg font-medium text-gray-900 mb-4 flex items-center">
            <CubeIcon class="w-5 h-5 mr-2 text-green-600" />
            Product Selection
          </h3>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Select Products (Optional)</label>
            <div class="relative product-dropdown">
              <input
                type="text"
                v-model="productSearchQuery"
                @input="filterProducts"
                @focus="showProductDropdown = true"
                placeholder="Search products to apply this discount to..."
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
              <div v-if="filteredProducts.length > 0 && showProductDropdown" class="absolute z-10 w-full mt-1 bg-white border border-gray-300 rounded-lg shadow-lg max-h-60 overflow-y-auto">
                <div
                  v-for="product in filteredProducts"
                  :key="product.id"
                  @click="toggleProductSelection(product)"
                  class="flex items-center p-3 hover:bg-gray-100 cursor-pointer border-b border-gray-100 last:border-b-0"
                >
                  <input
                    type="checkbox"
                    :checked="selectedProducts.some(p => p.id === product.id)"
                    class="mr-3 h-4 w-4 text-green-600 rounded focus:ring-green-500"
                    readonly
                  />
                  <div class="flex-1">
                    <div class="font-medium text-gray-900">{{ product.name }}</div>
                    <div class="text-sm text-gray-500">${{ product.price?.toFixed(2) || '0.00' }}</div>
                  </div>
                </div>
              </div>
            </div>
            <div v-if="selectedProducts.length > 0" class="mt-3">
              <h4 class="text-sm font-medium text-gray-700 mb-2">Selected Products:</h4>
              <div class="flex flex-wrap gap-2">
                <span
                  v-for="product in selectedProducts"
                  :key="product.id"
                  class="inline-flex items-center px-3 py-1 rounded-full text-sm bg-green-100 text-green-800 border border-green-200"
                >
                  {{ product.name }}
                  <button
                    @click="removeProduct(product)"
                    class="ml-2 text-green-600 hover:text-green-800"
                  >
                    ×
                  </button>
                </span>
              </div>
            </div>
            <p class="text-xs text-gray-500 mt-2">
              <InformationCircleIcon class="w-4 h-4 inline mr-1" />
              Leave empty to apply this discount to all products in your shop
            </p>
          </div>
        </div>

        <!-- Customer Eligibility -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h3 class="text-lg font-medium text-gray-900 mb-4 flex items-center">
            <UsersIcon class="w-5 h-5 mr-2 text-green-600" />
            Customer Eligibility
          </h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
            <div>
              <label for="eligibilityType" class="block text-sm font-medium text-gray-700 mb-1">Eligibility Type</label>
              <select 
                id="eligibilityType" 
                v-model="form.eligibilityType"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 bg-white transition duration-150 ease-in-out"
              >
                <option value="all">Everyone</option>
                <option value="specific">Specific Customers</option>
                <option value="segment">Customer Segments</option>
              </select>
            </div>
          </div>

          <!-- Specific Customers Selection -->
          <div v-if="form.eligibilityType === 'specific'" class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">Select Customers</label>
            <div class="relative customer-dropdown">
              <input
                type="text"
                v-model="customerSearchQuery"
                @input="filterCustomers"
                @focus="showCustomerDropdown = true"
                placeholder="Search customers..."
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
              <div v-if="filteredCustomers.length > 0 && showCustomerDropdown" class="absolute z-10 w-full mt-1 bg-white border border-gray-300 rounded-lg shadow-lg max-h-60 overflow-y-auto">
                <div
                  v-for="customer in filteredCustomers"
                  :key="customer.id"
                  @click="toggleCustomerSelection(customer)"
                  class="flex items-center p-3 hover:bg-gray-100 cursor-pointer border-b border-gray-100 last:border-b-0"
                >
                  <input
                    type="checkbox"
                    :checked="selectedCustomers.some(c => c.id === customer.id)"
                    class="mr-3 h-4 w-4 text-blue-600 rounded focus:ring-blue-500"
                    readonly
                  />
                  <div class="flex-1">
                    <div class="font-medium text-gray-900">{{ customer.firstName }} {{ customer.lastName }}</div>
                    <div class="text-sm text-gray-500">{{ customer.email }}</div>
                  </div>
                </div>
              </div>
            </div>
            <div v-if="selectedCustomers.length > 0" class="mt-3">
              <h4 class="text-sm font-medium text-gray-700 mb-2">Selected Customers:</h4>
              <div class="flex flex-wrap gap-2">
                <span
                  v-for="customer in selectedCustomers"
                  :key="customer.id"
                  class="inline-flex items-center px-3 py-1 rounded-full text-sm bg-blue-100 text-blue-800 border border-blue-200"
                >
                  {{ customer.firstName }} {{ customer.lastName }}
                  <button
                    @click="removeCustomer(customer)"
                    class="ml-2 text-blue-600 hover:text-blue-800"
                  >
                    ×
                  </button>
                </span>
              </div>
            </div>
          </div>

          <!-- Customer Segments Selection -->
          <div v-if="form.eligibilityType === 'segment'" class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">Select Customer Segments</label>
            <div class="relative segment-dropdown">
              <input
                type="text"
                v-model="segmentSearchQuery"
                @input="filterSegments"
                @focus="showSegmentDropdown = true"
                placeholder="Search segments..."
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
              <div v-if="filteredSegments.length > 0 && showSegmentDropdown" class="absolute z-10 w-full mt-1 bg-white border border-gray-300 rounded-lg shadow-lg max-h-60 overflow-y-auto">
                <div
                  v-for="segment in filteredSegments"
                  :key="segment.id"
                  @click="toggleSegmentSelection(segment)"
                  class="flex items-center p-3 hover:bg-gray-100 cursor-pointer border-b border-gray-100 last:border-b-0"
                >
                  <input
                    type="checkbox"
                    :checked="selectedSegments.some(s => s.id === segment.id)"
                    class="mr-3 h-4 w-4 text-purple-600 rounded focus:ring-purple-500"
                    readonly
                  />
                  <div class="flex-1">
                    <div class="font-medium text-gray-900">{{ segment.name }}</div>
                    <div class="text-sm text-gray-500">{{ segment.description || 'No description' }}</div>
                  </div>
                </div>
              </div>
            </div>
            <div v-if="selectedSegments.length > 0" class="mt-3">
              <h4 class="text-sm font-medium text-gray-700 mb-2">Selected Segments:</h4>
              <div class="flex flex-wrap gap-2">
                <span
                  v-for="segment in selectedSegments"
                  :key="segment.id"
                  class="inline-flex items-center px-3 py-1 rounded-full text-sm bg-purple-100 text-purple-800 border border-purple-200"
                >
                  {{ segment.name }}
                  <button
                    @click="removeSegment(segment)"
                    class="ml-2 text-purple-600 hover:text-purple-800"
                  >
                    ×
                  </button>
                </span>
              </div>
            </div>
          </div>
        </div>

        <!-- Usage Limits -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h3 class="text-lg font-medium text-gray-900 mb-4 flex items-center">
            <ChartBarIcon class="w-5 h-5 mr-2 text-green-600" />
            Usage Limits
          </h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label for="usageLimit" class="block text-sm font-medium text-gray-700 mb-1">Total Usage Limit</label>
              <input
                id="usageLimit"
                v-model.number="form.usageLimit"
                type="number"
                min="0"
                placeholder="No limit"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
              <p class="text-xs text-gray-500 mt-1">Leave empty for unlimited usage</p>
            </div>
            <div>
              <label for="perCustomerLimit" class="block text-sm font-medium text-gray-700 mb-1">Per Customer Limit</label>
              <input
                id="perCustomerLimit"
                v-model.number="form.perCustomerLimit"
                type="number"
                min="0"
                placeholder="No limit"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
              <p class="text-xs text-gray-500 mt-1">Leave empty for unlimited usage per customer</p>
            </div>
          </div>
        </div>

        <!-- Validity Period -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h3 class="text-lg font-medium text-gray-900 mb-4 flex items-center">
            <CalendarIcon class="w-5 h-5 mr-2 text-green-600" />
            Validity Period
          </h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label for="startAt" class="block text-sm font-medium text-gray-700 mb-1">Start Date & Time *</label>
              <input
                id="startAt"
                v-model="form.startAt"
                type="datetime-local"
                required
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
            </div>
            <div>
              <label for="endAt" class="block text-sm font-medium text-gray-700 mb-1">End Date & Time *</label>
              <input
                id="endAt"
                v-model="form.endAt"
                type="datetime-local"
                required
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
            </div>
          </div>
          <div class="flex items-center space-x-2 pt-4">
            <input 
              id="active" 
              type="checkbox" 
              v-model="form.active"
              class="form-checkbox h-5 w-5 text-green-600 rounded focus:ring-green-500 transition duration-150 ease-in-out" 
            />
            <label for="active" class="text-base font-medium text-gray-800">Discount is Active</label>
          </div>
        </div>

        <!-- Form Actions -->
        <div class="flex justify-end space-x-4 pt-6 border-t border-gray-200">
          <router-link
            to="/dashboard/discounts"
            class="px-5 py-2.5 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-100 transition-colors duration-200 ease-in-out"
          >
            Cancel
          </router-link>
          <button
            type="submit"
            :disabled="formLoading || !isFormValid"
            class="bg-green-600 text-white px-5 py-2.5 rounded-lg shadow-md hover:bg-green-700 transition duration-300 ease-in-out disabled:opacity-60 disabled:cursor-not-allowed"
          >
            <span v-if="formLoading" class="flex items-center">
              <div class="animate-spin h-5 w-5 mr-3 border-2 border-white border-t-transparent rounded-full"></div>
              Creating...
            </span>
            <span v-else>
              Create Discount
            </span>
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { discountService } from '@/services/discount'
import { productService } from '@/services/product'
import { customerService } from '@/services/customer'
import { customerSegmentService } from '@/services/customerSegment'
import { 
  ArrowLeftIcon,
  TagIcon,
  CurrencyDollarIcon,
  CubeIcon,
  UsersIcon,
  ChartBarIcon,
  CalendarIcon,
  InformationCircleIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const shopStore = useShopStore()
const shop = computed(() => shopStore.activeShop)
const shopId = computed(() => shop.value?.id)

// State
const loading = ref(false)
const formLoading = ref(false)

// Data for dropdowns
const products = ref([])
const customers = ref([])
const segments = ref([])
const loadingProducts = ref(false)
const loadingCustomers = ref(false)
const loadingSegments = ref(false)

// Search states
const productSearchQuery = ref('')
const customerSearchQuery = ref('')
const segmentSearchQuery = ref('')

// Dropdown visibility states
const showProductDropdown = ref(false)
const showCustomerDropdown = ref(false)
const showSegmentDropdown = ref(false)

// Selected items
const selectedProducts = ref([])
const selectedCustomers = ref([])
const selectedSegments = ref([])

// Form state
const form = ref({
  name: '',
  description: '',
  type: 'percentage',
  value: 0,
  appliesToProducts: [],
  appliesToVariants: [],
  couponCode: '',
  startAt: '',
  endAt: '',
  active: true,
  eligibilityType: 'all',
  allowedCustomers: [],
  allowedSegments: [],
  usageLimit: null,
  perCustomerLimit: null,
})

// Computed
const filteredProducts = computed(() => {
  if (!productSearchQuery.value) return products.value.slice(0, 10)
  const query = productSearchQuery.value.toLowerCase()
  return products.value.filter(p => 
    p.name.toLowerCase().includes(query) ||
    p.description?.toLowerCase().includes(query)
  ).slice(0, 10)
})

const filteredCustomers = computed(() => {
  if (!customerSearchQuery.value) return customers.value.slice(0, 10)
  const query = customerSearchQuery.value.toLowerCase()
  return customers.value.filter(c => 
    c.firstName.toLowerCase().includes(query) ||
    c.lastName.toLowerCase().includes(query) ||
    c.email.toLowerCase().includes(query)
  ).slice(0, 10)
})

const filteredSegments = computed(() => {
  if (!segmentSearchQuery.value) return segments.value.slice(0, 10)
  const query = segmentSearchQuery.value.toLowerCase()
  return segments.value.filter(s => 
    s.name.toLowerCase().includes(query) ||
    s.description?.toLowerCase().includes(query)
  ).slice(0, 10)
})

const isFormValid = computed(() => {
  return form.value.name.trim() && 
         form.value.value > 0 && 
         form.value.startAt && 
         form.value.endAt &&
         new Date(form.value.endAt) > new Date(form.value.startAt)
})

// Methods
// Dropdown filter functions
function filterProducts() {
  showProductDropdown.value = true
}

function filterCustomers() {
  showCustomerDropdown.value = true
}

function filterSegments() {
  showSegmentDropdown.value = true
}

// Selection functions
function toggleProductSelection(product) {
  const index = selectedProducts.value.findIndex(p => p.id === product.id)
  if (index > -1) {
    selectedProducts.value.splice(index, 1)
  } else {
    selectedProducts.value.push(product)
  }
  showProductDropdown.value = false
}

function toggleCustomerSelection(customer) {
  const index = selectedCustomers.value.findIndex(c => c.id === customer.id)
  if (index > -1) {
    selectedCustomers.value.splice(index, 1)
  } else {
    selectedCustomers.value.push(customer)
  }
  showCustomerDropdown.value = false
}

function toggleSegmentSelection(segment) {
  const index = selectedSegments.value.findIndex(s => s.id === segment.id)
  if (index > -1) {
    selectedSegments.value.splice(index, 1)
  } else {
    selectedSegments.value.push(segment)
  }
  showSegmentDropdown.value = false
}

// Remove functions
function removeProduct(product) {
  const index = selectedProducts.value.findIndex(p => p.id === product.id)
  if (index > -1) {
    selectedProducts.value.splice(index, 1)
  }
}

function removeCustomer(customer) {
  const index = selectedCustomers.value.findIndex(c => c.id === customer.id)
  if (index > -1) {
    selectedCustomers.value.splice(index, 1)
  }
}

function removeSegment(segment) {
  const index = selectedSegments.value.findIndex(s => s.id === segment.id)
  if (index > -1) {
    selectedSegments.value.splice(index, 1)
  }
}

// Load data for dropdowns
async function loadProducts() {
  if (!shopId.value) return
  loadingProducts.value = true
  try {
    products.value = await productService.fetchAllByShop(shopId.value)
  } catch (err) {
    console.error('Failed to load products:', err)
  } finally {
    loadingProducts.value = false
  }
}

async function loadCustomers() {
  if (!shopId.value) return
  loadingCustomers.value = true
  try {
    customers.value = await customerService.fetchAll(shopId.value)
  } catch (err) {
    console.error('Failed to load customers:', err)
  } finally {
    loadingCustomers.value = false
  }
}

async function loadSegments() {
  if (!shopId.value) return
  loadingSegments.value = true
  try {
    segments.value = await customerSegmentService.fetchAll(shopId.value)
  } catch (err) {
    console.error('Failed to load segments:', err)
  } finally {
    loadingSegments.value = false
  }
}

// Submit form
async function submitForm() {
  if (!shopId.value) {
    console.error('No shop ID available for discount operation.')
    return
  }

  // Basic validation
  if (!form.value.name.trim()) {
    alert('Please enter a discount name')
    return
  }
  if (!form.value.startAt) {
    alert('Please select a start date and time')
    return
  }
  if (!form.value.endAt) {
    alert('Please select an end date and time')
    return
  }
  if (new Date(form.value.endAt) <= new Date(form.value.startAt)) {
    alert('End date must be after start date')
    return
  }
  if (form.value.value <= 0) {
    alert('Discount value must be greater than 0')
    return
  }
  if (form.value.type === 'percentage' && form.value.value > 100) {
    alert('Percentage discount cannot exceed 100%')
    return
  }

  formLoading.value = true

  const payload = {
    name: form.value.name,
    description: form.value.description,
    type: form.value.type,
    value: form.value.value,
    appliesToProducts: selectedProducts.value.map(p => p.id),
    appliesToVariants: form.value.appliesToVariants,
    couponCode: form.value.couponCode || undefined,
    startAt: form.value.startAt,
    endAt: form.value.endAt,
    active: form.value.active,
    eligibilityType: form.value.eligibilityType,
    allowedCustomers: selectedCustomers.value.map(c => c.id),
    allowedSegments: selectedSegments.value.map(s => s.id),
    usageLimit: form.value.usageLimit,
    perCustomerLimit: form.value.perCustomerLimit,
  }

  try {
    await discountService.create(shopId.value, payload)
    router.push('/dashboard/discounts')
  } catch (err) {
    console.error('Failed to create discount:', err.response?.data || err.message)
    const errorMessage = err.response?.data?.message || err.response?.data?.error || err.message
    alert('Failed to create discount: ' + errorMessage)
  } finally {
    formLoading.value = false
  }
}

// Click outside handlers
function handleClickOutside(event) {
  // Close product dropdown
  if (!event.target.closest('.product-dropdown')) {
    showProductDropdown.value = false
  }
  
  // Close customer dropdown
  if (!event.target.closest('.customer-dropdown')) {
    showCustomerDropdown.value = false
  }
  
  // Close segment dropdown
  if (!event.target.closest('.segment-dropdown')) {
    showSegmentDropdown.value = false
  }
}

onMounted(() => {
  if (shopId.value) {
    loadProducts()
    loadCustomers()
    loadSegments()
  }
  
  // Add click outside listener
  document.addEventListener('click', handleClickOutside)
})

// Clean up event listener
onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
/* Custom styles if needed */
</style> 