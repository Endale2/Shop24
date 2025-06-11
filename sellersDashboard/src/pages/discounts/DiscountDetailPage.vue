<template>
  <div class="p-6 bg-gray-50 min-h-screen">
    <div class="max-w-2xl mx-auto bg-white shadow-xl rounded-lg p-8">
      <h1 class="text-3xl font-bold text-gray-800 mb-6">Discount Details</h1>

      <div v-if="loading" class="text-center py-10 text-gray-600 text-lg">
        Loading discount details...
      </div>

      <div v-else-if="!discount">
        <p class="text-center text-red-500 font-semibold py-10">
          Discount not found or an error occurred.
        </p>
        <div class="flex justify-center mt-6">
          <button @click="goBack"
                  class="px-5 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-100">
            Back to Discounts
          </button>
        </div>
      </div>

      <div v-else>
        <!-- Discount Meta -->
        <dl class="grid grid-cols-1 gap-x-4 gap-y-6 sm:grid-cols-2 mb-8">
          <div>
            <dt class="text-sm font-medium text-gray-500">Name</dt>
            <dd class="mt-1 text-sm text-gray-900">{{ discount.name }}</dd>
          </div>
          <div>
            <dt class="text-sm font-medium text-gray-500">Category</dt>
            <dd class="mt-1 text-sm text-gray-900">{{ discount.category }}</dd>
          </div>
          <div>
            <dt class="text-sm font-medium text-gray-500">Type</dt>
            <dd class="mt-1 text-sm text-gray-900">{{ discount.type }}</dd>
          </div>
          <div>
            <dt class="text-sm font-medium text-gray-500">Value</dt>
            <dd class="mt-1 text-sm text-gray-900">
              {{ discount.type === 'percentage' ? discount.value + '%' : '$' + discount.value }}
            </dd>
          </div>
          <div class="sm:col-span-2">
            <dt class="text-sm font-medium text-gray-500">Description</dt>
            <dd class="mt-1 text-sm text-gray-900">{{ discount.description || '—' }}</dd>
          </div>
          <div>
            <dt class="text-sm font-medium text-gray-500">Starts</dt>
            <dd class="mt-1 text-sm text-gray-900">{{ formatDateTime(discount.startAt) }}</dd>
          </div>
          <div>
            <dt class="text-sm font-medium text-gray-500">Ends</dt>
            <dd class="mt-1 text-sm text-gray-900">{{ formatDateTime(discount.endAt) }}</dd>
          </div>
          <div class="sm:col-span-2">
            <dt class="text-sm font-medium text-gray-500">Coupon Code</dt>
            <dd class="mt-1 text-sm text-gray-900">{{ discount.couponCode || '—' }}</dd>
          </div>
        </dl>

        <!-- Applied Products -->
        <div v-if="discount.appliesToProducts?.length" class="mb-8">
          <h2 class="text-2xl font-semibold text-gray-700 mb-4">Applies to Products</h2>
          <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
            <div v-for="prod in discount.appliesToProducts" :key="prod.id"
                 class="border rounded-lg overflow-hidden shadow-sm bg-white">
              <img :src="prod.main_image" alt=""
                   class="w-full h-40 object-cover" />
              <div class="p-4">
                <h3 class="font-medium text-gray-800">{{ prod.name }}</h3>
                <p class="text-sm text-gray-600 mt-1 line-clamp-3">{{ prod.description }}</p>
              </div>
            </div>
          </div>
        </div>

        <div class="flex justify-center">
          <button @click="goBack"
                  class="px-5 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-100">
            Back to Discounts
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { discountService } from '@/services/discount'
import { format } from 'date-fns'

// state
const loading = ref(true)
const discount = ref(null)

// router + store
const router = useRouter()
const route = useRoute()
const shopStore = useShopStore()

const shopId = shopStore.activeShop?.id
const discountId = route.params.discountId

// Helpers
function formatDateTime(iso) {
  if (!iso) return '—'
  return format(new Date(iso), 'PPP p')
}

// load data
async function loadDiscount() {
  loading.value = true
  try {
    const d = await discountService.fetchById(shopId, discountId)
    discount.value = d
  } catch (e) {
    console.error('Failed to load discount:', e)
    discount.value = null
  } finally {
    loading.value = false
  }
}

function goBack() {
  router.push({ name: 'Discounts' })
}

onMounted(loadDiscount)
</script>

<style scoped>
/* optional: ensure line-clamp */
.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
