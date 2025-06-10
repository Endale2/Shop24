<template>
  <div class="p-6 bg-gray-50 min-h-screen">
    <div class="max-w-2xl mx-auto bg-white shadow-xl rounded-lg p-8">
      <h1 class="text-3xl font-bold text-gray-800 mb-6">Discount Details</h1>

      <div v-if="loading" class="text-center py-10 text-gray-600 text-lg">
        Loading discount details...
      </div>
      <div v-else-if="!discount">
        <p class="text-center text-red-500 font-semibold py-10">Discount not found or an error occurred.</p>
        <div class="flex justify-center mt-6">
          <button @click="goBack" class="px-5 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-100 transition-colors duration-200">
            Back to Discounts
          </button>
        </div>
      </div>
      <div v-else>
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

          <div class="flex justify-between mt-8">
            <button
              type="button"
              @click="confirmDelete"
              class="bg-red-600 text-white px-5 py-2 rounded-md shadow-md hover:bg-red-700 transition duration-300 ease-in-out disabled:opacity-60 disabled:cursor-not-allowed"
              :disabled="formLoading"
            >
              Delete Discount
            </button>
            <div class="space-x-3">
              <button type="button" @click="goBack" class="px-5 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-100 transition-colors duration-200">
                Back
              </button>
              <button
                type="submit"
                :disabled="formLoading"
                class="bg-blue-600 text-white px-5 py-2 rounded-md shadow-md hover:bg-blue-700 transition duration-300 ease-in-out disabled:opacity-60 disabled:cursor-not-allowed"
              >
                {{ formLoading ? 'Saving...' : 'Save Changes' }}
              </button>
            </div>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, watch, defineProps } from 'vue' // Added defineProps
import { useRoute, useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { discountService } from '@/services/discount'
import { format } from 'date-fns'

// Define the prop
const props = defineProps({
  discountId: {
    type: String,
    required: true
  }
})

const route = useRoute()
const router = useRouter()
const shopStore = useShopStore()

// Access discountId directly from props
const discountId = computed(() => props.discountId);
const shopId = computed(() => shopStore.activeShop?.id)

// State
const discount = ref(null)
const loading = ref(true)
const formLoading = ref(false)

// Form state
const form = ref({
  id: null,
  name: '',
  description: '',
  category: 'product',
  type: 'percentage',
  value: 0,
  appliesToProducts: [],
  couponCode: '',
  startAt: '', // for datetime-local: 'YYYY-MM-DDTHH:mm'
  endAt: '',
  active: true,
})
const appliesToInput = ref('')

// Helpers for datetime-local conversion
function toDateTimeLocal(iso) {
  if (!iso) return ''
  const d = new Date(iso)
  if (isNaN(d)) return ''
  const tzOffset = d.getTimezoneOffset() * 60000
  const localISO = new Date(d.getTime() - tzOffset).toISOString().slice(0,16)
  return localISO
}

function fromDateTimeLocal(value) {
  if (!value) return undefined
  const d = new Date(value)
  if (isNaN(d)) return undefined
  return d.toISOString()
}

function formatDateTime(iso) {
  if (!iso) return '—'
  try {
    return format(new Date(iso), 'PPP p')
  } catch (e) {
    console.error("Error formatting date:", e);
    return new Date(iso).toLocaleString();
  }
}

// Fetch discount
async function loadDiscount() {
  if (!shopId.value) {
    console.warn("No active shop found, cannot load discount details.");
    loading.value = false;
    discount.value = null;
    return;
  }
  // Use props.discountId directly
  if (!props.discountId) {
    console.error("Discount ID is missing from props.");
    loading.value = false;
    discount.value = null;
    return;
  }

  loading.value = true
  try {
    const d = await discountService.fetchById(shopId.value, props.discountId) // Use props.discountId
    discount.value = d
    // Populate form
    form.value.id = d.id
    form.value.name = d.name
    form.value.description = d.description
    form.value.category = d.category
    form.value.type = d.type
    form.value.value = d.value
    form.value.appliesToProducts = [...(d.appliesToProducts || [])]
    appliesToInput.value = (d.appliesToProducts || []).join(', ')
    form.value.couponCode = d.couponCode || ''
    form.value.startAt = toDateTimeLocal(d.startAt)
    form.value.endAt = toDateTimeLocal(d.endAt)
    form.value.active = !!d.active
  } catch (err) {
    console.error('Failed to load discount', err)
    discount.value = null
  } finally {
    loading.value = false
  }
}

// Submit update
async function submitForm() {
  if (!shopId.value || !discount.value) return
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
    startAt: fromDateTimeLocal(form.value.startAt),
    endAt: fromDateTimeLocal(form.value.endAt),
    active: form.value.active,
  }

  try {
    await discountService.update(shopId.value, props.discountId, payload) // Use props.discountId
    await loadDiscount()
    console.log("Discount updated successfully!");
  } catch (err) {
    console.error('Failed to update discount', err)
  } finally {
    formLoading.value = false
  }
}

// Delete with confirmation
async function confirmDelete() {
  if (!shopId.value || !discount.value) return
  const ok = window.confirm(`Are you sure you want to delete discount "${discount.value.name}"? This action cannot be undone.`)
  if (!ok) return
  try {
    await discountService.delete(shopId.value, props.discountId) // Use props.discountId
    router.push({ name: 'Discounts' })
  } catch (err) {
    console.error('Failed to delete discount', err)
  }
}

// Back to list
function goBack() {
  router.push({ name: 'Discounts' })
}

onMounted(() => {
  loadDiscount()
})

// Watch for changes in the discountId prop
watch(() => props.discountId, (newId, oldId) => {
  if (newId && newId !== oldId) {
    loadDiscount();
  }
});

// Also watch for shopId in case it loads asynchronously after component mount
watch(shopId, (newShopId) => {
  if (newShopId) {
    loadDiscount();
  }
});
</script>

<style scoped>
/* Add any additional styling if needed */
</style>