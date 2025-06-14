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

    <div v-if="loading" class="text-center py-16 text-gray-600 text-lg">
      <div class="flex flex-col items-center justify-center">
        <SpinnerIcon class="animate-spin h-10 w-10 text-green-500 mb-4" />
        <p>Loading discounts...</p>
      </div>
    </div>
    <div v-else-if="discounts.length === 0" class="bg-green-50 border border-green-200 text-green-700 px-6 py-10 rounded-xl text-center shadow-md">
      <p class="text-lg font-medium mb-3">No discounts created yet.</p>
      <p class="text-base">Click the "<span class="font-semibold">+ New Discount</span>" button to add your first discount!</p>
    </div>
    <div v-else class="bg-white shadow-xl rounded-xl overflow-hidden border border-gray-200">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-100">
          <tr>
            <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Name</th>
            <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Type</th>
            <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Value</th>
            <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Coupon Code</th>
            <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Status</th>
            <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Validity Period</th>
            </tr>
        </thead>
        <tbody class="divide-y divide-gray-100">
          <tr
            v-for="d in discounts"
            :key="d.id"
            @click="goToDiscountDetail(d.id)"
            class="bg-white hover:bg-green-50 cursor-pointer transition-colors duration-150 ease-in-out"
          >
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-lg font-semibold text-green-700">
                {{ d.name }}
              </div>
              <p class="text-sm text-gray-500 mt-0.5 truncate max-w-xs">{{ d.description }}</p>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600 capitalize">
                <span :class="{
                    'bg-blue-100 text-blue-800': d.category === 'product',
                    'bg-purple-100 text-purple-800': d.category === 'order'
                }" class="inline-flex px-2.5 py-0.5 rounded-full text-xs font-medium mr-2">
                    {{ d.category }}
                </span>
                <span class="font-medium">{{ d.type.replace('_', ' ') }}</span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-lg">
              <span v-if="d.type === 'percentage'" class="font-bold text-green-700">{{ d.value }}%</span>
              <span v-else class="font-bold text-green-700">${{ d.value.toFixed(2) }}</span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-base text-gray-700 font-mono">
              {{ d.couponCode || 'N/A' }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span
                :class="{'bg-green-100 text-green-800': d.active, 'bg-red-100 text-red-800': !d.active}"
                class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium"
              >
                <CheckCircleIcon v-if="d.active" class="w-4 h-4 mr-1" />
                <XCircleIcon v-else class="w-4 h-4 mr-1" />
                {{ d.active ? 'Active' : 'Inactive' }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                <div class="flex items-center">
                    <CalendarIcon class="w-4 h-4 mr-1 text-gray-400" />
                    {{ formatDateTime(d.startAt) }}
                </div>
                <div class="flex items-center mt-1">
                    <CalendarIcon class="w-4 h-4 mr-1 text-gray-400" />
                    {{ formatDateTime(d.endAt) }}
                </div>
            </td>
            </tr>
        </tbody>
      </table>
    </div>

    <div v-if="showForm" class="fixed inset-0 bg-gray-900 bg-opacity-75 flex items-center justify-center z-50 p-4">
      <div class="bg-white rounded-xl shadow-2xl w-full max-w-lg p-8 relative transform scale-100 opacity-100 transition-all duration-300 ease-out">
        <h2 class="text-2xl font-bold mb-6 text-gray-800 text-center">
          {{ formMode === 'create' ? 'Create New Discount' : 'Edit Discount' }}
        </h2>
        <form @submit.prevent="submitForm" class="space-y-5">
          <div>
            <label for="name" class="block text-sm font-medium text-gray-700 mb-1">Name</label>
            <input id="name" v-model="form.name" type="text" required
                   class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out" />
          </div>
          <div>
            <label for="description" class="block text-sm font-medium text-gray-700 mb-1">Description</label>
            <textarea id="description" v-model="form.description" rows="3"
                      class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"></textarea>
          </div>
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div>
              <label for="category" class="block text-sm font-medium text-gray-700 mb-1">Category</label>
              <select id="category" v-model="form.category" required
                      class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 bg-white transition duration-150 ease-in-out">
                <option value="product">Product</option>
                <option value="order">Order</option>
              </select>
            </div>
            <div>
              <label for="type" class="block text-sm font-medium text-gray-700 mb-1">Type</label>
              <select id="type" v-model="form.type" required
                      class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 bg-white transition duration-150 ease-in-out">
                <option value="percentage">Percentage</option>
                <option value="fixed_amount">Fixed Amount</option>
              </select>
            </div>
          </div>
          <div>
            <label for="value" class="block text-sm font-medium text-gray-700 mb-1">Value</label>
            <input
              id="value"
              v-model.number="form.value"
              type="number"
              min="0"
              required
              class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
            />
          </div>
          <div>
            <label for="couponCode" class="block text-sm font-medium text-gray-700 mb-1">Coupon Code (optional)</label>
            <input id="couponCode" v-model="form.couponCode" type="text"
                   class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out" />
          </div>
          <div v-if="form.category === 'product'">
            <label for="appliesToInput" class="block text-sm font-medium text-gray-700 mb-1">Applies to Product IDs (comma-separated)</label>
            <input
              id="appliesToInput"
              v-model="appliesToInput"
              type="text"
              placeholder="e.g. prod123, prod456 (leave blank to apply to all products)"
              class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
            />
            <p class="text-xs text-gray-500 mt-1">
              Enter product IDs; leave blank to apply to all products in this shop.
            </p>
          </div>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label for="startAt" class="block text-sm font-medium text-gray-700 mb-1">Start At</label>
              <input
                id="startAt"
                v-model="form.startAt"
                type="datetime-local"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
            </div>
            <div>
              <label for="endAt" class="block text-sm font-medium text-gray-700 mb-1">End At</label>
              <input
                id="endAt"
                v-model="form.endAt"
                type="datetime-local"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
            </div>
          </div>
          <div class="flex items-center space-x-2 pt-2">
            <input id="active" type="checkbox" v-model="form.active"
                   class="form-checkbox h-5 w-5 text-green-600 rounded focus:ring-green-500 transition duration-150 ease-in-out" />
            <label for="active" class="text-base font-medium text-gray-800">Discount is Active</label>
          </div>

          <div class="flex justify-end space-x-4 mt-6 pt-4 border-t border-gray-100">
            <button type="button" @click="closeForm"
                    class="px-5 py-2.5 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-100 transition-colors duration-200 ease-in-out">
              Cancel
            </button>
            <button
              type="submit"
              :disabled="formLoading"
              class="bg-green-600 text-white px-5 py-2.5 rounded-lg shadow-md hover:bg-green-700 transition duration-300 ease-in-out disabled:opacity-60 disabled:cursor-not-allowed"
            >
              <span v-if="formLoading" class="flex items-center">
                <SpinnerIcon class="animate-spin h-5 w-5 mr-3" />
                Saving...
              </span>
              <span v-else>
                {{ formMode === 'create' ? 'Create Discount' : 'Update Discount' }}
              </span>
            </button>
          </div>
        </form>
        <button @click="closeForm" class="absolute top-4 right-4 text-gray-400 hover:text-gray-600 transition-colors duration-200">
          <XIcon class="h-6 w-6" />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router' // Import useRouter
import { useShopStore } from '@/store/shops'
import { discountService } from '@/services/discount'
import { format } from 'date-fns'
import {
  PlusIcon,
  RefreshIcon as SpinnerIcon,
  CheckCircleIcon,
  XCircleIcon,
  CalendarIcon,
  XIcon // Close icon for modal
} from '@heroicons/vue/outline'

const shopStore = useShopStore()
const shop = computed(() => shopStore.activeShop)
const shopId = computed(() => shop.value?.id)
const router = useRouter() // Initialize useRouter

// State
const discounts = ref([])
const loading = ref(false)

// Form state (for Create New Discount)
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
  startAt: '',
  endAt: '',
  active: true,
})
const appliesToInput = ref('')

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
    return new Date(iso).toLocaleString();
  }
}

// Fetch discounts
async function loadDiscounts() {
  if (!shopId.value) {
    console.warn('No active shop selected. Cannot load discounts.')
    discounts.value = []
    return
  }
  loading.value = true
  try {
    discounts.value = await discountService.fetchAllByShop(shopId.value)
  } catch (err) {
    console.error('Failed to load discounts:', err)
  } finally {
    loading.value = false
  }
}

// Navigate to Discount Detail page
function goToDiscountDetail(discountId) {
  router.push({ name: 'DiscountDetail', params: { discountId: discountId } });
}

// Open create form (from main page)
function openCreateForm() {
  formMode.value = 'create'
  resetForm()
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
  resetForm()
}

// Submit create or update (for the modal form)
async function submitForm() {
  if (!shopId.value) {
    console.error('No shop ID available for discount operation.')
    return
  }
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
      // Note: The 'edit' mode from this component is primarily for creation now.
      // If you were to add an "Edit" button to each row to open the modal directly,
      // this part would be more actively used here. For now, it's mostly handled
      // by the detail page for editing existing discounts.
      await discountService.update(shopId.value, form.value.id, payload)
    }
    await loadDiscounts()
    closeForm()
  } catch (err) {
    console.error('Failed to save discount:', err.response?.data || err.message)
  } finally {
    formLoading.value = false
  }
}

// confirmDelete is removed from here as it's moved to the detail page.
// The openEditForm is also removed from here as its primary use case (editing from list)
// is now on the detail page.

onMounted(() => {
  if (shopId.value) {
    loadDiscounts()
  }
});
</script>

<style scoped>
/* No additional scoped styles needed as TailwindCSS handles most of the styling */
</style>