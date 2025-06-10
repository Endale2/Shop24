<!-- src/pages/discounts/DiscountsPage.vue -->
<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-2xl font-semibold">Discounts</h1>
      <button
        @click="openCreateForm"
        class="bg-green-600 text-white py-1 px-4 rounded hover:bg-green-700 transition"
      >
        + New Discount
      </button>
    </div>

    <!-- List of discounts -->
    <div v-if="loading" class="text-gray-500">Loading...</div>
    <div v-else>
      <table class="w-full table-auto border-collapse">
        <thead>
          <tr class="bg-gray-100">
            <th class="px-3 py-2 text-left">Name</th>
            <th class="px-3 py-2 text-left">Type</th>
            <th class="px-3 py-2 text-left">Value</th>
            <th class="px-3 py-2 text-left">Coupon Code</th>
            <th class="px-3 py-2 text-left">Active</th>
            <th class="px-3 py-2 text-left">Period</th>
            <th class="px-3 py-2">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="d in discounts" :key="d.id" class="border-b">
            <td class="px-3 py-2">{{ d.name }}</td>
            <td class="px-3 py-2">{{ d.type }}</td>
            <td class="px-3 py-2">
              <span v-if="d.type === 'percentage'">{{ d.value }}%</span>
              <span v-else>{{ d.value }}</span>
            </td>
            <td class="px-3 py-2">{{ d.couponCode || '—' }}</td>
            <td class="px-3 py-2">
              <input type="checkbox" disabled :checked="d.active" />
            </td>
            <td class="px-3 py-2">
              {{ formatDateTime(d.startAt) }} – {{ formatDateTime(d.endAt) }}
            </td>
            <td class="px-3 py-2 space-x-2 text-center">
              <button @click="openEditForm(d)" class="text-blue-600 hover:underline">Edit</button>
              <button @click="confirmDelete(d)" class="text-red-600 hover:underline">Delete</button>
            </td>
          </tr>
          <tr v-if="discounts.length === 0">
            <td colspan="7" class="px-3 py-4 text-center text-gray-500">No discounts yet.</td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Create/Edit form modal -->
    <div v-if="showForm" class="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-lg w-full max-w-lg p-6 relative">
        <h2 class="text-xl font-semibold mb-4">
          {{ formMode === 'create' ? 'New Discount' : 'Edit Discount' }}
        </h2>
        <form @submit.prevent="submitForm" class="space-y-4">
          <!-- Name -->
          <div>
            <label class="block text-sm mb-1">Name</label>
            <input v-model="form.name" type="text" required class="w-full border px-2 py-1 rounded" />
          </div>
          <!-- Description -->
          <div>
            <label class="block text-sm mb-1">Description</label>
            <textarea v-model="form.description" class="w-full border px-2 py-1 rounded"></textarea>
          </div>
          <!-- Category -->
          <div>
            <label class="block text-sm mb-1">Category</label>
            <select v-model="form.category" required class="w-full border px-2 py-1 rounded">
              <option value="product">Product</option>
              <option value="order">Order</option>
            </select>
          </div>
          <!-- Type -->
          <div>
            <label class="block text-sm mb-1">Type</label>
            <select v-model="form.type" required class="w-full border px-2 py-1 rounded">
              <option value="percentage">Percentage</option>
              <option value="fixed_amount">Fixed Amount</option>
            </select>
          </div>
          <!-- Value -->
          <div>
            <label class="block text-sm mb-1">Value</label>
            <input
              v-model.number="form.value"
              type="number"
              min="0"
              required
              class="w-full border px-2 py-1 rounded"
            />
          </div>
          <!-- Coupon Code -->
          <div>
            <label class="block text-sm mb-1">Coupon Code (optional)</label>
            <input v-model="form.couponCode" type="text" class="w-full border px-2 py-1 rounded" />
          </div>
          <!-- Applies to products -->
          <div v-if="form.category === 'product'">
            <label class="block text-sm mb-1">Applies to Product IDs (comma-separated)</label>
            <input
              v-model="appliesToInput"
              type="text"
              placeholder="e.g. 60f5..., 60f5..."
              class="w-full border px-2 py-1 rounded"
            />
            <p class="text-xs text-gray-500 mt-1">
              Enter product IDs; leave blank to apply to all products.
            </p>
          </div>
          <!-- Start & End -->
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm mb-1">Start At</label>
              <input
                v-model="form.startAt"
                type="datetime-local"
                class="w-full border px-2 py-1 rounded"
              />
            </div>
            <div>
              <label class="block text-sm mb-1">End At</label>
              <input
                v-model="form.endAt"
                type="datetime-local"
                class="w-full border px-2 py-1 rounded"
              />
            </div>
          </div>
          <!-- Active -->
          <div class="flex items-center space-x-2">
            <input id="active" type="checkbox" v-model="form.active" />
            <label for="active" class="text-sm">Active</label>
          </div>

          <!-- Buttons -->
          <div class="flex justify-end space-x-3 mt-4">
            <button type="button" @click="closeForm" class="px-4 py-1 border rounded">Cancel</button>
            <button
              type="submit"
              :disabled="formLoading"
              class="bg-blue-600 text-white px-4 py-1 rounded hover:bg-blue-700 disabled:opacity-50"
            >
              {{ formLoading ? 'Saving...' : 'Save' }}
            </button>
          </div>
        </form>
        <!-- Close icon -->
        <button @click="closeForm" class="absolute top-2 right-2 text-gray-500 hover:text-gray-700">
          ✕
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useShopStore } from '@/store/shops'
import { discountService } from '@/services/discount'
import { format } from 'date-fns' // optional: for formatting; or use plain JS

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
  endAt:   '',
  active: true,
})
const appliesToInput = ref('')

// Helpers
function formatDateTime(iso) {
  if (!iso) return '—'
  const d = new Date(iso)
  if (isNaN(d)) return '—'
  // Format as locale string, or use date-fns as needed:
  return d.toLocaleString()
}

// Fetch discounts
async function loadDiscounts() {
  if (!shopId.value) return
  loading.value = true
  try {
    discounts.value = await discountService.fetchAllByShop(shopId.value)
  } catch (err) {
    console.error('Failed to load discounts', err)
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

// Open edit form
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
  // Convert ISO to 'YYYY-MM-DDTHH:mm' for datetime-local input
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

  // Parse appliesToInput into array, trimming whitespace
  let appliesArr = []
  if (form.value.category === 'product' && appliesToInput.value.trim()) {
    appliesArr = appliesToInput.value
      .split(',')
      .map(s => s.trim())
      .filter(s => s.length > 0)
  }

  // Build payload
  const payload = {
    name:            form.value.name,
    description:     form.value.description,
    category:        form.value.category,
    type:            form.value.type,
    value:           form.value.value,
    appliesToProducts: appliesArr,
    couponCode:      form.value.couponCode || undefined,
    startAt:         form.value.startAt ? new Date(form.value.startAt).toISOString() : undefined,
    endAt:           form.value.endAt   ? new Date(form.value.endAt).toISOString()   : undefined,
    active:          form.value.active,
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
    // Optionally show error to user
  } finally {
    formLoading.value = false
  }
}

// Delete with confirmation
async function confirmDelete(disc) {
  if (!shopId.value) return
  const ok = window.confirm(`Delete discount "${disc.name}"? This cannot be undone.`)
  if (!ok) return
  try {
    await discountService.delete(shopId.value, disc.id)
    await loadDiscounts()
  } catch (err) {
    console.error('Failed to delete discount', err)
  }
}

// On mounted
onMounted(() => {
  if (shopId.value) {
    loadDiscounts()
  }
})
</script>

<style scoped>
/* Add any modal backdrop or table styling as needed */
</style>
