<template>
  <div class="p-6 bg-gray-50 min-h-screen">
    <div class="flex justify-between items-center mb-8">
      <h1 class="text-3xl font-bold text-gray-800">Discounts Management</h1>
      <button
        @click="openCreateForm"
        class="bg-indigo-600 text-white py-2 px-5 rounded-lg shadow-md hover:bg-indigo-700 transition duration-300 ease-in-out transform hover:scale-105"
      >
        + New Discount
      </button>
    </div>

    <div v-if="loading" class="text-center py-10 text-gray-600 text-lg">
      <div class="flex items-center justify-center">
        <svg class="animate-spin h-6 w-6 text-indigo-500 mr-3" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        Loading discounts...
      </div>
    </div>
    <div v-else class="bg-white shadow-xl rounded-lg overflow-hidden">
      <table class="w-full table-auto border-collapse">
        <thead class="bg-gray-100 border-b border-gray-200">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-semibold text-gray-700 uppercase tracking-wider">Name</th>
            <th class="px-4 py-3 text-left text-sm font-semibold text-gray-700 uppercase tracking-wider">Type</th>
            <th class="px-4 py-3 text-left text-sm font-semibold text-gray-700 uppercase tracking-wider">Value</th>
            <th class="px-4 py-3 text-left text-sm font-semibold text-gray-700 uppercase tracking-wider">Coupon Code</th>
            <th class="px-4 py-3 text-left text-sm font-semibold text-gray-700 uppercase tracking-wider">Active</th>
            <th class="px-4 py-3 text-left text-sm font-semibold text-gray-700 uppercase tracking-wider">Period</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="d in discounts" :key="d.id" class="border-b border-gray-100 hover:bg-gray-50 transition-colors duration-150">
            <td class="px-4 py-3 text-gray-800 font-medium">
              <router-link :to="{ name: 'DiscountDetail', params: { discountId: d.id } }" class="text-blue-600 hover:text-blue-800 hover:underline">
                {{ d.name }}
              </router-link>
            </td>
            <td class="px-4 py-3 text-gray-600 capitalize">{{ d.type.replace('_', ' ') }}</td>
            <td class="px-4 py-3 text-gray-600">
              <span v-if="d.type === 'percentage'" class="font-semibold text-green-700">{{ d.value }}%</span>
              <span v-else class="font-semibold text-green-700">${{ d.value.toFixed(2) }}</span>
            </td>
            <td class="px-4 py-3 text-gray-600">{{ d.couponCode || '—' }}</td>
            <td class="px-4 py-3">
              <input type="checkbox" :checked="d.active" disabled class="form-checkbox h-5 w-5 text-indigo-600 rounded-sm" />
            </td>
            <td class="px-4 py-3 text-gray-600 text-sm">
              {{ formatDateTime(d.startAt) }} – {{ formatDateTime(d.endAt) }}
            </td>
          </tr>
          <tr v-if="discounts.length === 0">
            <td colspan="6" class="px-4 py-6 text-center text-gray-500 text-lg">No discounts created yet. Click "+ New Discount" to add one!</td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="showForm" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
      <div class="bg-white rounded-xl shadow-2xl w-full max-w-lg p-8 relative transform scale-100 opacity-100 transition-all duration-300 ease-out">
        <h2 class="text-2xl font-bold mb-6 text-gray-800">
          {{ formMode === 'create' ? 'Create New Discount' : 'Edit Discount' }}
        </h2>
        <form @submit.prevent="submitForm" class="space-y-5">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Name</label>
            <input v-model="form.name" type="text" required
                   class="w-full border-gray-300 rounded-md shadow-sm focus:border-indigo-500 focus:ring-indigo-500 px-3 py-2" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Description</label>
            <textarea v-model="form.description" rows="3"
                      class="w-full border-gray-300 rounded-md shadow-sm focus:border-indigo-500 focus:ring-indigo-500 px-3 py-2"></textarea>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Category</label>
            <select v-model="form.category" required
                    class="w-full border-gray-300 rounded-md shadow-sm focus:border-indigo-500 focus:ring-indigo-500 px-3 py-2 bg-white">
              <option value="product">Product</option>
              <option value="order">Order</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Type</label>
            <select v-model="form.type" required
                    class="w-full border-gray-300 rounded-md shadow-sm focus:border-indigo-500 focus:ring-indigo-500 px-3 py-2 bg-white">
              <option value="percentage">Percentage</option>
              <option value="fixed_amount">Fixed Amount</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Value</label>
            <input
              v-model.number="form.value"
              type="number"
              min="0"
              required
              class="w-full border-gray-300 rounded-md shadow-sm focus:border-indigo-500 focus:ring-indigo-500 px-3 py-2"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Coupon Code (optional)</label>
            <input v-model="form.couponCode" type="text"
                   class="w-full border-gray-300 rounded-md shadow-sm focus:border-indigo-500 focus:ring-indigo-500 px-3 py-2" />
          </div>
          <div v-if="form.category === 'product'">
            <label class="block text-sm font-medium text-gray-700 mb-1">Applies to Product IDs (comma-separated)</label>
            <input
              v-model="appliesToInput"
              type="text"
              placeholder="e.g. 60f5..., 60f5... (leave blank to apply to all products)"
              class="w-full border-gray-300 rounded-md shadow-sm focus:border-indigo-500 focus:ring-indigo-500 px-3 py-2"
            />
            <p class="text-xs text-gray-500 mt-1">
              Enter product IDs; leave blank to apply to all products.
            </p>
          </div>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Start At</label>
              <input
                v-model="form.startAt"
                type="datetime-local"
                class="w-full border-gray-300 rounded-md shadow-sm focus:border-indigo-500 focus:ring-indigo-500 px-3 py-2"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">End At</label>
              <input
                v-model="form.endAt"
                type="datetime-local"
                class="w-full border-gray-300 rounded-md shadow-sm focus:border-indigo-500 focus:ring-indigo-500 px-3 py-2"
              />
            </div>
          </div>
          <div class="flex items-center space-x-2">
            <input id="active" type="checkbox" v-model="form.active"
                   class="form-checkbox h-4 w-4 text-indigo-600 rounded focus:ring-indigo-500" />
            <label for="active" class="text-sm font-medium text-gray-700">Active</label>
          </div>

          <div class="flex justify-end space-x-4 mt-6">
            <button type="button" @click="closeForm"
                    class="px-5 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-100 transition-colors duration-200">
              Cancel
            </button>
            <button
              type="submit"
              :disabled="formLoading"
              class="bg-indigo-600 text-white px-5 py-2 rounded-md shadow-md hover:bg-indigo-700 transition duration-300 ease-in-out disabled:opacity-60 disabled:cursor-not-allowed"
            >
              {{ formLoading ? 'Saving...' : 'Save Discount' }}
            </button>
          </div>
        </form>
        <button @click="closeForm" class="absolute top-4 right-4 text-gray-400 hover:text-gray-600 transition-colors duration-200">
          <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useShopStore } from '@/store/shops'
import { discountService } from '@/services/discount'
import { format } from 'date-fns'

const shopStore = useShopStore()
const shop = computed(() => shopStore.activeShop)
const shopId = computed(() => shop.value?.id)

// State
const discounts = ref([])
const loading = ref(false)

// Form state
const showForm = ref(false)
const formMode = ref('create') // 'create' or 'edit'
const formLoading = ref(false)
const form = ref({
  id: null,
  name: '',
  description: '',
  category: 'product',
  type: 'percentage',
  value: 0,
  appliesToProducts: [],
  couponCode: '',
  startAt: '', // in 'YYYY-MM-DDTHH:mm' format for datetime-local
  endAt: '',
  active: true,
})
const appliesToInput = ref('')

// Helpers
function formatDateTime(iso) {
  if (!iso) return '—'
  try {
    return format(new Date(iso), 'PPP p') // e.g., "Oct 25, 2023, 10:30 AM"
  } catch (e) {
    console.error("Error formatting date:", e);
    return new Date(iso).toLocaleString(); // Fallback to browser's locale string
  }
}

// Fetch discounts
async function loadDiscounts() {
  if (!shopId.value) return
  loading.value = true
  try {
    discounts.value = await discountService.fetchAllByShop(shopId.value)
  } catch (err) {
    console.error('Failed to load discounts', err)
    // Optionally show an alert or toast notification
  } finally {
    loading.value = false
  }
}

// Open create form
function openCreateForm() {
  formMode.value = 'create'
  resetForm()
  showForm.value = true
}

// Open edit form (this function is kept for internal use if needed, but not directly called from the list anymore)
function openEditForm(disc) {
  formMode.value = 'edit'
  form.value.id = disc.id
  form.value.name = disc.name
  form.value.description = disc.description
  form.value.category = disc.category
  form.value.type = disc.type
  form.value.value = disc.value
  form.value.appliesToProducts = [...(disc.appliesToProducts || [])]
  appliesToInput.value = (disc.appliesToProducts || []).join(', ')
  form.value.couponCode = disc.couponCode || ''

  if (disc.startAt) {
    const d = new Date(disc.startAt)
    form.value.startAt = d.toISOString().slice(0,16)
  } else {
    form.value.startAt = ''
  }
  if (disc.endAt) {
    const d2 = new Date(disc.endAt)
    form.value.endAt = d2.toISOString().slice(0,16)
  } else {
    form.value.endAt = ''
  }
  form.value.active = !!disc.active
  showForm.value = true
}

// Reset form fields
function resetForm() {
  form.value = {
    id: null,
    name: '',
    description: '',
    category: 'product',
    type: 'percentage',
    value: 0,
    appliesToProducts: [],
    couponCode: '',
    startAt: '',
    endAt: '',
    active: true,
  }
  appliesToInput.value = ''
}

// Close form
function closeForm() {
  showForm.value = false
}

// Submit create or update
async function submitForm() {
  if (!shopId.value) return
  formLoading.value = true

  let appliesArr = []
  if (form.value.category === 'product' && appliesToInput.value.trim()) {
    appliesArr = appliesToInput.value
      .split(',')
      .map(s => s.trim())
      .filter(s => s.length > 0)
  }

  const payload = {
    name: form.value.name,
    description: form.value.description,
    category: form.value.category,
    type: form.value.type,
    value: form.value.value,
    appliesToProducts: appliesArr,
    couponCode: form.value.couponCode || undefined,
    startAt: form.value.startAt ? new Date(form.value.startAt).toISOString() : undefined,
    endAt: form.value.endAt ? new Date(form.value.endAt).toISOString() : undefined,
    active: form.value.active,
  }

  try {
    if (formMode.value === 'create') {
      await discountService.create(shopId.value, payload)
    } else if (formMode.value === 'edit' && form.value.id) {
      await discountService.update(shopId.value, form.value.id, payload)
    }
    await loadDiscounts()
    closeForm()
  } catch (err) {
    console.error('Failed to save discount', err)
  } finally {
    formLoading.value = false
  }
}

// Delete with confirmation (still present in script, but not called from UI)
async function confirmDelete(disc) {
  if (!shopId.value) return
  const ok = window.confirm(`Are you sure you want to delete discount "${disc.name}"? This action cannot be undone.`)
  if (!ok) return
  try {
    await discountService.delete(shopId.value, disc.id)
    await loadDiscounts()
  } catch (err) {
    console.error('Failed to delete discount', err)
  }
}

onMounted(() => {
  if (shopId.value) {
    loadDiscounts()
  }
})
</script>

<style scoped>
/* No additional scoped styles needed as TailwindCSS handles most of the styling */
</style>