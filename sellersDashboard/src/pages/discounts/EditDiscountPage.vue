<template>
  <div class="max-w-2xl mx-auto p-6 bg-white rounded-xl shadow-xl mt-8">
    <h2 class="text-2xl font-bold mb-6 text-gray-800 text-center">Edit Discount</h2>
    <div v-if="loading" class="text-center py-16 text-gray-600 text-lg">
      <div class="flex flex-col items-center justify-center">
        <div class="animate-spin h-10 w-10 text-green-500 mb-4 border-4 border-green-200 border-t-green-500 rounded-full"></div>
        <p>Loading discount data...</p>
      </div>
    </div>
    <div v-else-if="loadError" class="bg-red-50 border border-red-200 text-red-700 px-6 py-10 rounded-xl text-center shadow-md">
      <div class="flex flex-col items-center">
        <p class="text-lg font-medium mb-3">Discount not found or an error occurred.</p>
        <p class="text-base mb-4">Please check the URL or try again.</p>
        <button @click="() => router.push({ name: 'Discounts' })"
                class="px-5 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-100 transition-colors duration-200">
          Back to Discounts
        </button>
      </div>
    </div>
    <form v-else @submit.prevent="submitForm" class="space-y-6">
      <!-- Basic Info -->
      <div class="bg-gray-50 p-6 rounded-lg">
        <h3 class="text-lg font-medium text-gray-900 mb-4">Basic Information</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Name *</label>
            <input v-model="form.name" type="text" required class="w-full border-gray-300 rounded-lg px-3 py-2.5" />
          </div>
        </div>
        <div class="mt-4">
          <label class="block text-sm font-medium text-gray-700 mb-1">Description</label>
          <textarea v-model="form.description" rows="3" class="w-full border-gray-300 rounded-lg px-3 py-2.5"></textarea>
        </div>
      </div>
      <!-- Discount Config -->
      <div class="bg-gray-50 p-6 rounded-lg">
        <h3 class="text-lg font-medium text-gray-900 mb-4">Discount Configuration</h3>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Category *</label>
            <select v-model="form.category" required class="w-full border-gray-300 rounded-lg px-3 py-2.5">
              <option value="product">Product</option>
              <option value="order">Order</option>
              <option value="shipping">Shipping</option>
              <option value="buy_x_get_y">Buy X Get Y</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Type *</label>
            <select v-model="form.type" required class="w-full border-gray-300 rounded-lg px-3 py-2.5">
              <option value="percentage">Percentage</option>
              <option value="fixed">Fixed Amount</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Value *</label>
            <input v-model.number="form.value" type="number" min="0" :step="form.type === 'percentage' ? 1 : 0.01" required class="w-full border-gray-300 rounded-lg px-3 py-2.5" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Per Customer Usage Limit</label>
            <input v-model.number="form.perCustomerLimit" type="number" min="0" placeholder="No limit" class="w-full border-gray-300 rounded-lg px-3 py-2.5" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Total Usage Limit</label>
            <input v-model.number="form.usageLimit" type="number" min="0" placeholder="No limit" class="w-full border-gray-300 rounded-lg px-3 py-2.5" />
          </div>
        </div>
      </div>
      <!-- Validity Period -->
      <div class="bg-gray-50 p-6 rounded-lg">
        <h3 class="text-lg font-medium text-gray-900 mb-4">Validity Period</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Start Date & Time</label>
            <input v-model="form.startAt" type="datetime-local" required class="w-full border-gray-300 rounded-lg px-3 py-2.5" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">End Date & Time</label>
            <input v-model="form.endAt" type="datetime-local" required class="w-full border-gray-300 rounded-lg px-3 py-2.5" />
          </div>
        </div>
        <div class="flex items-center space-x-2 pt-4">
          <input type="checkbox" v-model="form.active" class="form-checkbox h-5 w-5 text-green-600 rounded focus:ring-green-500 transition duration-150 ease-in-out" />
          <label class="text-base font-medium text-gray-800">Discount is Active</label>
        </div>
      </div>
      <!-- Save Button -->
      <div class="flex justify-end">
        <button type="submit" :disabled="formLoading" class="bg-green-600 text-white px-6 py-2 rounded-lg shadow hover:bg-green-700 disabled:opacity-60">
          <span v-if="formLoading">Saving...</span>
          <span v-else>Save</span>
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { discountService } from '@/services/discount'

const route = useRoute()
const router = useRouter()
const shopStore = useShopStore()
const shopId = computed(() => route.params.shopId || route.query.shopId || shopStore.activeShop?.id)
const discountId = route.params.discountId || route.query.discountId

const form = ref({
  name: '',
  description: '',
  category: 'product',
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

// Add a helper to convert ISO/RFC3339 to datetime-local format
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
      category: d.category,
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
    loadError.value = true
  } finally {
    loading.value = false
  }
})

async function submitForm() {
  formLoading.value = true
  try {
    await discountService.update(shopId.value, discountId, {
      ...form.value,
    })
    router.push({ name: 'Discounts' })
  } finally {
    formLoading.value = false
  }
}
</script>

<style scoped>
input[type="radio"] {
  accent-color: #16a34a;
}
</style> 