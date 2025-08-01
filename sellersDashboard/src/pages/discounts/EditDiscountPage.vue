<template>
  <div class="min-h-full bg-gray-50">
    <div class="p-4 sm:p-6 lg:p-8">
      
      <!-- Header Section -->
      <div class="mb-6 sm:mb-8">
        <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center space-y-4 sm:space-y-0">
          <div class="flex items-center">
            <div class="w-10 h-10 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center mr-3 shadow-sm">
              <TagIcon class="w-5 h-5 text-white" />
            </div>
            <div>
              <h1 class="text-2xl sm:text-3xl font-bold text-gray-900 tracking-tight">Edit Discount</h1>
              <p class="text-sm text-gray-600 mt-1">Update your promotional offer settings</p>
            </div>
          </div>
          <router-link
            to="/dashboard/discounts"
            class="inline-flex items-center px-3 py-2 bg-gray-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 transition duration-150 ease-in-out"
          >
            <ArrowLeftIcon class="w-4 h-4 mr-2" />
            Back to Discounts
          </router-link>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-16 text-gray-500">
        <div class="flex flex-col items-center justify-center">
          <div class="animate-spin h-8 w-8 text-green-500 mb-3 border-3 border-green-200 border-t-green-500 rounded-full"></div>
          <p class="text-sm font-medium">Loading discount data...</p>
        </div>
      </div>

      <!-- Error State -->
      <div v-else-if="loadError" class="bg-red-50 border border-red-200 text-red-700 px-6 py-10 rounded-xl text-center shadow-sm">
        <div class="flex flex-col items-center">
          <ExclamationIcon class="w-12 h-12 text-red-500 mb-4" />
          <h3 class="text-lg font-medium mb-3">Discount not found</h3>
          <p class="text-base mb-4 max-w-md">
            The discount you're looking for doesn't exist or you don't have permission to access it.
          </p>
          <button 
            @click="() => router.push({ name: 'Discounts' })"
            class="inline-flex items-center px-4 py-2 bg-red-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2 transition duration-150 ease-in-out"
          >
            <ArrowLeftIcon class="w-4 h-4 mr-2" />
            Back to Discounts
          </button>
        </div>
      </div>

      <!-- Form -->
      <div v-else class="space-y-6">
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
          <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
            <h2 class="text-sm font-semibold text-gray-900">Edit Discount</h2>
          </div>
          <form @submit.prevent="submitForm" class="p-6 space-y-6">
            
            <!-- Basic Information -->
            <div class="bg-gray-50 rounded-lg p-4">
              <div class="flex items-center mb-4">
                <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center mr-3">
                  <TagIcon class="w-4 h-4 text-blue-600" />
                </div>
                <h3 class="text-lg font-medium text-gray-900">Basic Information</h3>
              </div>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                  <label for="name" class="block text-sm font-medium text-gray-700 mb-2">Discount Name *</label>
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
                  <label for="couponCode" class="block text-sm font-medium text-gray-700 mb-2">Coupon Code (Optional)</label>
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
                <label for="description" class="block text-sm font-medium text-gray-700 mb-2">Description</label>
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
            <div class="bg-gray-50 rounded-lg p-4">
              <div class="flex items-center mb-4">
                <div class="w-8 h-8 bg-green-100 rounded-lg flex items-center justify-center mr-3">
                  <CurrencyDollarIcon class="w-4 h-4 text-green-600" />
                </div>
                <h3 class="text-lg font-medium text-gray-900">Discount Configuration</h3>
              </div>
              <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
                <div>
                  <label for="type" class="block text-sm font-medium text-gray-700 mb-2">Discount Type *</label>
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
                  <label for="value" class="block text-sm font-medium text-gray-700 mb-2">Discount Value *</label>
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

            <!-- Usage Limits -->
            <div class="bg-gray-50 rounded-lg p-4">
              <div class="flex items-center mb-4">
                <div class="w-8 h-8 bg-purple-100 rounded-lg flex items-center justify-center mr-3">
                  <ChartBarIcon class="w-4 h-4 text-purple-600" />
                </div>
                <h3 class="text-lg font-medium text-gray-900">Usage Limits</h3>
              </div>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                  <label for="usageLimit" class="block text-sm font-medium text-gray-700 mb-2">Total Usage Limit</label>
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
                  <label for="perCustomerLimit" class="block text-sm font-medium text-gray-700 mb-2">Per Customer Limit</label>
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
            <div class="bg-gray-50 rounded-lg p-4">
              <div class="flex items-center mb-4">
                <div class="w-8 h-8 bg-amber-100 rounded-lg flex items-center justify-center mr-3">
                  <CalendarIcon class="w-4 h-4 text-amber-600" />
                </div>
                <h3 class="text-lg font-medium text-gray-900">Validity Period</h3>
              </div>
              
              <!-- Date Range Info -->
              <div class="mb-4 p-3 bg-amber-50 border border-amber-200 rounded-lg">
                <div class="flex items-start">
                  <InformationCircleIcon class="w-4 h-4 text-amber-600 mr-2 mt-0.5 flex-shrink-0" />
                  <div class="text-sm text-amber-800">
                    <p class="font-medium mb-1">Set your discount availability period</p>
                    <p class="text-xs">Choose when your discount will be active. The discount will automatically activate and deactivate based on these dates.</p>
                  </div>
                </div>
              </div>
              
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div class="space-y-2">
                  <label for="startAt" class="block text-sm font-medium text-gray-700">Start Date & Time *</label>
                  <div class="relative">
                    <input
                      id="startAt"
                      v-model="form.startAt"
                      type="datetime-local"
                      required
                      class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out bg-white"
                    />
                    <div class="absolute inset-y-0 right-0 flex items-center pr-3 pointer-events-none">
                      <CalendarIcon class="w-4 h-4 text-gray-400" />
                    </div>
                  </div>
                  <p class="text-xs text-gray-500">When the discount becomes available to customers</p>
                </div>
                
                <div class="space-y-2">
                  <label for="endAt" class="block text-sm font-medium text-gray-700">End Date & Time *</label>
                  <div class="relative">
                    <input
                      id="endAt"
                      v-model="form.endAt"
                      type="datetime-local"
                      required
                      class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out bg-white"
                    />
                    <div class="absolute inset-y-0 right-0 flex items-center pr-3 pointer-events-none">
                      <CalendarIcon class="w-4 h-4 text-gray-400" />
                    </div>
                  </div>
                  <p class="text-xs text-gray-500">When the discount expires and becomes unavailable</p>
                </div>
              </div>
              
              <!-- Date Validation Feedback -->
              <div v-if="form.startAt && form.endAt" class="mt-4">
                <div v-if="isDateRangeValid" class="p-3 bg-green-50 border border-green-200 rounded-lg">
                  <div class="flex items-center">
                    <CheckCircleIcon class="w-4 h-4 text-green-600 mr-2" />
                    <span class="text-sm text-green-800 font-medium">Valid date range</span>
                  </div>
                  <p class="text-xs text-green-700 mt-1">
                    Discount will be active for {{ getDateRangeDuration }}
                  </p>
                </div>
                <div v-else class="p-3 bg-red-50 border border-red-200 rounded-lg">
                  <div class="flex items-center">
                    <ExclamationIcon class="w-4 h-4 text-red-600 mr-2" />
                    <span class="text-sm text-red-800 font-medium">Invalid date range</span>
                  </div>
                  <p class="text-xs text-red-700 mt-1">
                    End date must be after start date
                  </p>
                </div>
              </div>
              
              <!-- Active Status -->
              <div class="mt-6 pt-4 border-t border-gray-200">
                <div class="flex items-center space-x-3">
                  <input 
                    id="active" 
                    type="checkbox" 
                    v-model="form.active"
                    class="form-checkbox h-5 w-5 text-green-600 rounded focus:ring-green-500 transition duration-150 ease-in-out" 
                  />
                  <div class="flex-1">
                    <label for="active" class="text-sm font-medium text-gray-800">Discount is Active</label>
                    <p class="text-xs text-gray-500 mt-1">Enable this discount immediately (overrides date range)</p>
                  </div>
                  <span 
                    :class="form.active ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-600'"
                    class="inline-flex items-center px-2.5 py-1 rounded-full text-xs font-medium"
                  >
                    <CheckCircleIcon v-if="form.active" class="w-3 h-3 mr-1" />
                    <XCircleIcon v-else class="w-3 h-3 mr-1" />
                    {{ form.active ? 'Active' : 'Inactive' }}
                  </span>
                </div>
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
                class="bg-green-600 text-white px-5 py-2.5 rounded-lg shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out disabled:opacity-60 disabled:cursor-not-allowed"
              >
                <span v-if="formLoading" class="flex items-center">
                  <div class="animate-spin h-5 w-5 mr-3 border-2 border-white border-t-transparent rounded-full"></div>
                  Saving...
                </span>
                <span v-else>
                  Save Changes
                </span>
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { discountService } from '@/services/discount'
import { 
  ArrowLeftIcon,
  TagIcon,
  CurrencyDollarIcon,
  ChartBarIcon,
  CalendarIcon,
  InformationCircleIcon,
  ExclamationIcon,
  CheckCircleIcon,
  XCircleIcon
} from '@heroicons/vue/outline'

const route = useRoute()
const router = useRouter()
const shopStore = useShopStore()
const shopId = computed(() => route.params.shopId || route.query.shopId || shopStore.activeShop?.id)
const discountId = route.params.discountId || route.query.discountId

const form = ref({
  name: '',
  description: '',
  type: 'percentage',
  value: 0,
  startAt: '',
  endAt: '',
  active: true,
  perCustomerLimit: null,
  usageLimit: null,
})

const formLoading = ref(false)
const loading = ref(false)
const loadError = ref(false)

// Computed
const isFormValid = computed(() => {
  return form.value.name.trim() && 
         form.value.value > 0 && 
         form.value.startAt && 
         form.value.endAt &&
         new Date(form.value.endAt) > new Date(form.value.startAt)
})

// Computed properties for date range validation and duration
const isDateRangeValid = computed(() => {
  if (!form.value.startAt || !form.value.endAt) {
    return false
  }
  const start = new Date(form.value.startAt)
  const end = new Date(form.value.endAt)
  return end > start
})

const getDateRangeDuration = computed(() => {
  if (!form.value.startAt || !form.value.endAt) {
    return 'N/A'
  }
  const start = new Date(form.value.startAt)
  const end = new Date(form.value.endAt)
  const diffTime = Math.abs(end.getTime() - start.getTime())
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24)) + 1 // +1 to include both start and end day
  return `${diffDays} day${diffDays > 1 ? 's' : ''}`
})

// Helper to convert ISO/RFC3339 to datetime-local format
function toDatetimeLocal(val) {
  if (!val) return ''
  const date = new Date(val)
  // Pad with zeros
  const pad = n => n.toString().padStart(2, '0')
  return `${date.getFullYear()}-${pad(date.getMonth() + 1)}-${pad(date.getDate())}T${pad(date.getHours())}:${pad(date.getMinutes())}`
}

onMounted(async () => {
  if (!shopId.value || !discountId) {
    loadError.value = true
    return
  }
  
  loading.value = true
  try {
    const d = await discountService.fetchById(shopId.value, discountId)
    form.value = {
      name: d.name,
      description: d.description,
      type: d.type,
      value: d.value,
      startAt: toDatetimeLocal(d.startAt),
      endAt: toDatetimeLocal(d.endAt),
      active: d.active,
      perCustomerLimit: d.perCustomerLimit,
      usageLimit: d.usageLimit,
    }
    loadError.value = false
  } catch (err) {
    console.error('Failed to load discount:', err)
    loadError.value = true
  } finally {
    loading.value = false
  }
})

async function submitForm() {
  if (!shopId.value || !discountId) {
    console.error('Missing shop ID or discount ID')
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
  try {
    await discountService.update(shopId.value, discountId, {
      ...form.value,
    })
    router.push({ name: 'Discounts' })
  } catch (err) {
    console.error('Failed to update discount:', err)
    const errorMessage = err.response?.data?.message || err.response?.data?.error || err.message
    alert('Failed to update discount: ' + errorMessage)
  } finally {
    formLoading.value = false
  }
}
</script>

<style scoped>
/* Custom styles if needed */
</style> 