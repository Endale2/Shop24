<template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto font-sans">
    <div class="max-w-3xl mx-auto bg-white shadow-xl rounded-xl p-6 sm:p-8 border border-gray-200">
      <div class="flex justify-between items-center mb-6">
        <h1 class="text-3xl sm:text-4xl font-extrabold text-gray-900 leading-tight">Discount Details</h1>
        <div class="flex space-x-3">
          <button
            @click="openEditForm(discount)"
            class="inline-flex items-center px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg shadow-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
            v-if="discount"
          >
            <PencilIcon class="w-5 h-5 mr-2 -ml-1" />
            Edit Discount
          </button>
          <button
            @click="confirmDelete(discount)"
            class="inline-flex items-center px-4 py-2 bg-red-600 text-white text-sm font-medium rounded-lg shadow-md hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2 transition duration-150 ease-in-out"
            v-if="discount"
          >
            <TrashIcon class="w-5 h-5 mr-2 -ml-1" />
            Delete Discount
          </button>
        </div>
      </div>

      <div v-if="loading" class="text-center py-16 text-gray-600 text-lg">
        <div class="flex flex-col items-center justify-center">
          <SpinnerIcon class="animate-spin h-10 w-10 text-green-500 mb-4" />
          <p>Loading discount details...</p>
        </div>
      </div>

      <div v-else-if="!discount" class="bg-red-50 border border-red-200 text-red-700 px-6 py-10 rounded-xl text-center shadow-md">
        <p class="text-lg font-medium mb-3">Discount not found or an error occurred.</p>
        <p class="text-base">Please check the URL or try again.</p>
        <div class="flex justify-center mt-6">
          <button @click="goBack"
                  class="px-5 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-100 transition-colors duration-200">
            Back to Discounts
          </button>
        </div>
      </div>

      <div v-else>
        <dl class="grid grid-cols-1 md:grid-cols-2 gap-x-6 gap-y-8 mb-10 border-b border-gray-200 pb-8">
          <div>
            <dt class="text-sm font-medium text-gray-500">Name</dt>
            <dd class="mt-1 text-2xl font-bold text-gray-900">{{ discount.name }}</dd>
          </div>
          <div>
            <dt class="text-sm font-medium text-gray-500">Status</dt>
            <dd class="mt-1">
              <span
                :class="{'bg-green-100 text-green-800': discount.active, 'bg-red-100 text-red-800': !discount.active}"
                class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium"
              >
                <CheckCircleIcon v-if="discount.active" class="w-4 h-4 mr-1" />
                <XCircleIcon v-else class="w-4 h-4 mr-1" />
                {{ discount.active ? 'Active' : 'Inactive' }}
              </span>
            </dd>
          </div>
          <div>
            <dt class="text-sm font-medium text-gray-500">Category</dt>
            <dd class="mt-1 text-base text-gray-900 capitalize">{{ discount.category.replace('_', ' ') }}</dd>
          </div>
          <div>
            <dt class="text-sm font-medium text-gray-500">Type</dt>
            <dd class="mt-1 text-base text-gray-900 capitalize">{{ discount.type.replace('_', ' ') }}</dd>
          </div>
          <div>
            <dt class="text-sm font-medium text-gray-500">Value</dt>
            <dd class="mt-1 text-xl font-bold text-green-700">
              {{ discount.type === 'percentage' ? discount.value + '%' : '$' + discount.value?.toFixed(2) }}
            </dd>
          </div>
          <div>
            <dt class="text-sm font-medium text-gray-500">Coupon Code</dt>
            <dd class="mt-1 text-base text-gray-900 font-mono">{{ discount.couponCode || 'N/A' }}</dd>
          </div>
          <div class="md:col-span-2">
            <dt class="text-sm font-medium text-gray-500">Description</dt>
            <dd class="mt-1 text-base text-gray-900">{{ discount.description || 'No description provided.' }}</dd>
          </div>
          <div>
            <dt class="text-sm font-medium text-gray-500">Start Date</dt>
            <dd class="mt-1 text-base text-gray-900 flex items-center">
                <CalendarIcon class="w-5 h-5 mr-1.5 text-gray-400" />
                {{ formatDateTime(discount.startAt) }}
            </dd>
          </div>
          <div>
            <dt class="text-sm font-medium text-gray-500">End Date</dt>
            <dd class="mt-1 text-base text-gray-900 flex items-center">
                <CalendarIcon class="w-5 h-5 mr-1.5 text-gray-400" />
                {{ formatDateTime(discount.endAt) }}
            </dd>
          </div>
        </dl>

        <div v-if="discount.appliesToProducts?.length" class="mb-8">
          <h2 class="text-2xl font-semibold text-gray-700 mb-5">Applies to Products</h2>
          <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
            <div v-for="prod in discount.appliesToProducts" :key="prod.id"
                 class="group border border-gray-200 rounded-xl overflow-hidden shadow-sm hover:shadow-md transition-shadow duration-200 ease-in-out bg-white cursor-pointer"
                 @click="goToProductDetail(prod.id)">
              <div class="aspect-w-16 aspect-h-9 sm:aspect-w-4 sm:aspect-h-3 lg:aspect-w-16 lg:aspect-h-9 overflow-hidden">
                <img :src="prod.main_image || '/placeholder-image.png'" :alt="prod.name"
                     class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300 ease-in-out" />
              </div>
              <div class="p-4">
                <h3 class="text-lg font-medium text-gray-800 group-hover:text-green-700 transition-colors">{{ prod.name }}</h3>
                <p class="text-sm text-gray-600 mt-1 line-clamp-3">{{ prod.description || 'No description available.' }}</p>
              </div>
            </div>
          </div>
        </div>
        <div v-else-if="discount.category === 'product' && !discount.appliesToProducts?.length" class="mb-8 p-4 bg-gray-50 border border-gray-200 rounded-lg text-gray-700 text-center">
            <p class="text-md">This discount applies to **all products** in the shop.</p>
        </div>


        <div v-if="fetchedCollections.length" class="mb-8">
          <h2 class="text-2xl font-semibold text-gray-700 mb-5">Applies to Collections</h2>
          <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
            <div v-for="coll in fetchedCollections" :key="coll.id"
                 class="group border border-gray-200 rounded-xl overflow-hidden shadow-sm hover:shadow-md transition-shadow duration-200 ease-in-out bg-white cursor-pointer"
                 @click="goToCollectionDetail(coll.id)">
              <div class="aspect-w-16 aspect-h-9 sm:aspect-w-4 sm:aspect-h-3 lg:aspect-w-16 lg:aspect-h-9 overflow-hidden">
                <img :src="coll.image || '/placeholder-collection.png'" :alt="coll.title"
                     class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300 ease-in-out" />
              </div>
              <div class="p-4">
                <h3 class="text-lg font-medium text-gray-800 group-hover:text-green-700 transition-colors">{{ coll.title }}</h3>
                <p class="text-sm text-gray-600 mt-1 line-clamp-3">{{ coll.description || 'No description available.' }}</p>
              </div>
            </div>
          </div>
        </div>
        <div v-else-if="!discount.appliesToCollections?.length" class="mb-8 p-4 bg-gray-50 border border-gray-200 rounded-lg text-gray-700 text-center">
            <p class="text-md">This discount does not apply to any specific collections.</p>
        </div>


        <div class="flex justify-center mt-8">
          <button @click="goBack"
                  class="px-6 py-2.5 bg-gray-200 text-gray-800 rounded-lg shadow-sm hover:bg-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-offset-2 transition duration-200 ease-in-out">
            Back to Discounts
          </button>
        </div>
      </div>
    </div>

    <div v-if="showForm" class="fixed inset-0 bg-gray-900 bg-opacity-75 flex items-center justify-center z-50 p-4">
      <div class="bg-white rounded-xl shadow-2xl w-full max-w-lg p-8 relative transform scale-100 opacity-100 transition-all duration-300 ease-out">
        <h2 class="text-2xl font-bold mb-6 text-gray-800 text-center">
          {{ formMode === 'create' ? 'Create New Discount' : 'Edit Discount' }}
        </h2>
        <form @submit.prevent="submitForm" class="space-y-5">
          <div>
            <label for="modal-name" class="block text-sm font-medium text-gray-700 mb-1">Name</label>
            <input id="modal-name" v-model="form.name" type="text" required
                   class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out" />
          </div>
          <div>
            <label for="modal-description" class="block text-sm font-medium text-gray-700 mb-1">Description</label>
            <textarea id="modal-description" v-model="form.description" rows="3"
                      class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"></textarea>
          </div>
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div>
              <label for="modal-category" class="block text-sm font-medium text-gray-700 mb-1">Category</label>
              <select id="modal-category" v-model="form.category" required
                      class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 bg-white transition duration-150 ease-in-out">
                <option value="product">Product</option>
                <option value="order">Order</option>
                <option value="shipping">Shipping</option>
                 <option value="buy_x_get_y">Buy X Get Y</option>
              </select>
            </div>
            <div>
              <label for="modal-type" class="block text-sm font-medium text-gray-700 mb-1">Type</label>
              <select id="modal-type" v-model="form.type" required
                      class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 bg-white transition duration-150 ease-in-out">
                <option value="percentage">Percentage</option>
                <option value="fixed_amount">Fixed Amount</option>
              </select>
            </div>
          </div>
          <div>
            <label for="modal-value" class="block text-sm font-medium text-gray-700 mb-1">Value</label>
            <input
              id="modal-value"
              v-model.number="form.value"
              type="number"
              min="0"
              required
              class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
            />
          </div>
          <div>
            <label for="modal-couponCode" class="block text-sm font-medium text-gray-700 mb-1">Coupon Code (optional)</label>
            <input id="modal-couponCode" v-model="form.couponCode" type="text"
                   class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out" />
          </div>

          <div v-if="form.category === 'product' || form.category === 'buy_x_get_y'">
            <label for="modal-appliesToProductsInput" class="block text-sm font-medium text-gray-700 mb-1">Specific Product IDs (comma-separated)</label>
            <input
              id="modal-appliesToProductsInput"
              v-model="appliesToProductsInput"
              type="text"
              placeholder="e.g. 682f249f0998fdf958ce1acb (leave blank to apply to all products)"
              class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
            />
            <p class="text-xs text-gray-500 mt-1">
              Enter product IDs; leave blank to apply to all products in this shop for 'product' category.
            </p>
          </div>

          <div>
            <label for="modal-appliesToCollectionsInput" class="block text-sm font-medium text-gray-700 mb-1">Specific Collection IDs (comma-separated)</label>
            <input
              id="modal-appliesToCollectionsInput"
              v-model="appliesToCollectionsInput"
              type="text"
              placeholder="e.g. 6841db1c36dc6252945bc317 (leave blank for none)"
              class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
            />
            <p class="text-xs text-gray-500 mt-1">
              Enter collection IDs for this discount.
            </p>
          </div>

          <div v-if="form.category === 'buy_x_get_y'" class="border p-4 rounded-lg bg-gray-50">
            <h3 class="text-md font-semibold mb-3 text-gray-700">Buy X Get Y Details</h3>
            <div>
              <label for="modal-buyProductIds" class="block text-sm font-medium text-gray-700 mb-1">Buy Product IDs (comma-separated)</label>
              <input
                id="modal-buyProductIds"
                v-model="buyProductIdsInput"
                type="text"
                placeholder="e.g. prod1, prod2"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
            </div>
            <div class="mt-4">
              <label for="modal-buyQuantity" class="block text-sm font-medium text-gray-700 mb-1">Buy Quantity</label>
              <input
                id="modal-buyQuantity"
                v-model.number="form.buyQuantity"
                type="number"
                min="1"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
            </div>
            <div class="mt-4">
              <label for="modal-getProductIds" class="block text-sm font-medium text-gray-700 mb-1">Get Product IDs (comma-separated)</label>
              <input
                id="modal-getProductIds"
                v-model="getProductIdsInput"
                type="text"
                placeholder="e.g. prodA, prodB"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
            </div>
            <div class="mt-4">
              <label for="modal-getQuantity" class="block text-sm font-medium text-gray-700 mb-1">Get Quantity</label>
              <input
                id="modal-getQuantity"
                v-model.number="form.getQuantity"
                type="number"
                min="1"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
            </div>
          </div>

          <div v-if="form.category === 'order'" class="border p-4 rounded-lg bg-gray-50">
            <h3 class="text-md font-semibold mb-3 text-gray-700">Order-level Details</h3>
            <div>
              <label for="modal-minimumOrderSubtotal" class="block text-sm font-medium text-gray-700 mb-1">Minimum Order Subtotal</label>
              <input
                id="modal-minimumOrderSubtotal"
                v-model.number="form.minimumOrderSubtotal"
                type="number"
                min="0"
                step="0.01"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
            </div>
          </div>

          <div v-if="form.category === 'shipping'" class="border p-4 rounded-lg bg-gray-50">
            <h3 class="text-md font-semibold mb-3 text-gray-700">Shipping Details</h3>
            <div class="flex items-center space-x-2">
              <input id="modal-freeShipping" type="checkbox" v-model="form.freeShipping"
                     class="form-checkbox h-5 w-5 text-green-600 rounded focus:ring-green-500 transition duration-150 ease-in-out" />
              <label for="modal-freeShipping" class="text-base font-medium text-gray-800">Enable Free Shipping</label>
            </div>
            <div v-if="form.freeShipping" class="mt-4">
              <label for="modal-minimumFreeShipping" class="block text-sm font-medium text-gray-700 mb-1">Minimum Order for Free Shipping</label>
              <input
                id="modal-minimumFreeShipping"
                v-model.number="form.minimumFreeShipping"
                type="number"
                min="0"
                step="0.01"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
            </div>
          </div>

          <div class="border p-4 rounded-lg bg-gray-50">
            <h3 class="text-md font-semibold mb-3 text-gray-700">Usage Limits (Optional)</h3>
            <div>
              <label for="modal-usageLimit" class="block text-sm font-medium text-gray-700 mb-1">Total Usage Limit</label>
              <input
                id="modal-usageLimit"
                v-model.number="form.usageLimit"
                type="number"
                min="0"
                placeholder="No limit"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
            </div>
            <div class="mt-4">
              <label for="modal-perCustomerLimit" class="block text-sm font-medium text-gray-700 mb-1">Per Customer Limit</label>
              <input
                id="modal-perCustomerLimit"
                v-model.number="form.perCustomerLimit"
                type="number"
                min="0"
                placeholder="No limit"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
            </div>
            <div class="mt-4">
              <label for="modal-allowedCustomers" class="block text-sm font-medium text-gray-700 mb-1">Allowed Customer IDs (comma-separated)</label>
              <input
                id="modal-allowedCustomers"
                v-model="allowedCustomerIDsInput"
                type="text"
                placeholder="e.g. cust1, cust2 (leave blank for all customers)"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
              <p class="text-xs text-gray-500 mt-1">
                Only these customers can use the discount.
              </p>
            </div>
          </div>


          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label for="modal-startAt" class="block text-sm font-medium text-gray-700 mb-1">Start At</label>
              <input
                id="modal-startAt"
                v-model="form.startAt"
                type="datetime-local"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
            </div>
            <div>
              <label for="modal-endAt" class="block text-sm font-medium text-gray-700 mb-1">End At</label>
              <input
                id="modal-endAt"
                v-model="form.endAt"
                type="datetime-local"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
            </div>
          </div>
          <div class="flex items-center space-x-2 pt-2">
            <input id="modal-active" type="checkbox" v-model="form.active"
                   class="form-checkbox h-5 w-5 text-green-600 rounded focus:ring-green-500 transition duration-150 ease-in-out" />
            <label for="modal-active" class="text-base font-medium text-gray-800">Discount is Active</label>
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
import { ref, onMounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { discountService } from '@/services/discount'
import { collectionService } from '@/services/collection'
import { format } from 'date-fns'
import {
  PencilIcon,
  TrashIcon,
  RefreshIcon as SpinnerIcon,
  CheckCircleIcon,
  XCircleIcon,
  CalendarIcon,
  XIcon
} from '@heroicons/vue/outline'

// State for detail page
const loading = ref(true)
const discount = ref(null)
const fetchedCollections = ref([]) // State for fetched collection objects

// Router & Store
const router = useRouter()
const route = useRoute()
const shopStore = useShopStore()

const shopId = ref(shopStore.activeShop?.id);
watch(() => shopStore.activeShop, (newShop) => {
  if (newShop) {
    shopId.value = newShop.id;
    loadDiscount();
  }
});

const discountId = route.params.discountId

// Form state (for Edit Discount Modal)
const showForm = ref(false)
const formMode = ref('edit')
const formLoading = ref(false)
const form = ref({
  id: null,
  name: '',
  description: '',
  category: 'product', // Default
  type: 'percentage',  // Default
  value: 0,
  appliesToProducts: [], // Will store product IDs for API payload
  appliesToCollections: [], // Will store collection IDs for API payload
  couponCode: '',
  startAt: '',
  endAt: '',
  active: true,
  // New fields from backend model
  freeShipping: false,
  minimumOrderSubtotal: null,
  minimumFreeShipping: null,
  usageLimit: null,
  perCustomerLimit: null,
  allowedCustomers: [],
  buyProductIds: [],
  buyQuantity: null,
  getProductIds: [],
  getQuantity: null,
  appliesToVariants: [], // Not yet added to form, but in model
})

// For comma-separated inputs in the form
const appliesToProductsInput = ref('')
const appliesToCollectionsInput = ref('')
const allowedCustomerIDsInput = ref('')
const buyProductIdsInput = ref('')
const getProductIdsInput = ref('')

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
    return new Date(iso).toLocaleString(); // Fallback
  }
}

// Load data for the discount detail page
async function loadDiscount() {
  if (!shopId.value || !discountId) {
    console.warn("Shop ID or Discount ID not available yet. Skipping discount load.");
    loading.value = false;
    return;
  }

  loading.value = true;
  discount.value = null;
  fetchedCollections.value = []; // Reset collections

  try {
    // discountService.fetchById now returns `appliesToProducts` as full objects
    // and `appliesToCollections` as IDs.
    const d = await discountService.fetchById(shopId.value, discountId);

    // If discount applies to collections, fetch their details in the frontend
    if (d && d.appliesToCollections && d.appliesToCollections.length > 0) {
      const collectionPromises = d.appliesToCollections.map(collectionId =>
        collectionService.fetchById(shopId.value, collectionId).catch(err => {
          console.warn(`Failed to load collection ${collectionId}:`, err);
          return null;
        })
      );
      fetchedCollections.value = (await Promise.all(collectionPromises)).filter(c => c !== null);
    }

    discount.value = d; // Assign the fetched and mapped discount directly
  } catch (e) {
    console.error('Failed to load discount:', e);
  } finally {
    loading.value = false;
  }
}

function goBack() {
  router.push({ name: 'Discounts' });
}

function goToProductDetail(productId) {
  console.log(`Navigating to product detail for ID: ${productId}`);
  // router.push({ name: 'ProductDetail', params: { shopId: shopId.value, productId: productId } });
  alert(`Simulating navigation to product detail for ID: ${productId}`);
}

function goToCollectionDetail(collectionId) {
  console.log(`Navigating to collection detail for ID: ${collectionId}`);
  // router.push({ name: 'CollectionDetail', params: { shopId: shopId.value, collectionId: collectionId } });
  alert(`Simulating navigation to collection detail for ID: ${collectionId}`);
}

// Open edit form
function openEditForm(disc) {
  formMode.value = 'edit';
  form.value.id = disc.id;
  form.value.name = disc.name;
  form.value.description = disc.description;
  form.value.category = disc.category;
  form.value.type = disc.type;
  form.value.value = disc.value;

  // When opening the form, populate the input fields with comma-separated IDs
  // For products, map the objects back to IDs.
  appliesToProductsInput.value = (disc.appliesToProducts?.map(p => p.id) || []).join(', ');
  // For collections, these are already IDs from the service layer.
  appliesToCollectionsInput.value = (disc.appliesToCollections || []).join(', ');

  form.value.couponCode = disc.couponCode || '';
  // Convert ISO string to 'YYYY-MM-DDTHH:MM' format required by datetime-local input
  form.value.startAt = disc.startAt ? new Date(disc.startAt).toISOString().slice(0, 16) : '';
  form.value.endAt = disc.endAt ? new Date(disc.endAt).toISOString().slice(0, 16) : '';

  form.value.active = !!disc.active;

  // Populate new fields
  form.value.freeShipping = disc.freeShipping;
  form.value.minimumOrderSubtotal = disc.minimumOrderSubtotal;
  form.value.minimumFreeShipping = disc.minimumFreeShipping;
  form.value.usageLimit = disc.usageLimit;
  form.value.perCustomerLimit = disc.perCustomerLimit;
  allowedCustomerIDsInput.value = (disc.allowedCustomers || []).join(', ');
  buyProductIdsInput.value = (disc.buyProductIds || []).join(', ');
  form.value.buyQuantity = disc.buyQuantity;
  getProductIdsInput.value = (disc.getProductIds || []).join(', ');
  form.value.getQuantity = disc.getQuantity;
  // appliesToVariants is not in the form yet, but would be handled similarly if added.

  showForm.value = true;
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
    appliesToCollections: [],
    couponCode: '',
    startAt: '',
    endAt: '',
    active: true,
    freeShipping: false,
    minimumOrderSubtotal: null,
    minimumFreeShipping: null,
    usageLimit: null,
    perCustomerLimit: null,
    allowedCustomers: [],
    buyProductIds: [],
    buyQuantity: null,
    getProductIds: [],
    getQuantity: null,
    appliesToVariants: [],
  };
  appliesToProductsInput.value = '';
  appliesToCollectionsInput.value = '';
  allowedCustomerIDsInput.value = '';
  buyProductIdsInput.value = '';
  getProductIdsInput.value = '';
}

function closeForm() {
  showForm.value = false;
  resetForm(); // Reset form on close, good for fresh state
}

// Submit create or update
async function submitForm() {
  if (!shopId.value || !form.value.id) {
    console.error('Missing shop ID or discount ID for update operation.');
    return;
  }
  formLoading.value = true;

  // Parse IDs from input strings
  const parseIDs = (input) => input
    .split(',')
    .map(s => s.trim())
    .filter(s => s.length > 0);

  const payload = {
    name: form.value.name,
    description: form.value.description,
    category: form.value.category,
    type: form.value.type,
    value: form.value.value,
    appliesToProducts: parseIDs(appliesToProductsInput.value),
    appliesToCollections: parseIDs(appliesToCollectionsInput.value),
    couponCode: form.value.couponCode || undefined,
    startAt: form.value.startAt ? new Date(form.value.startAt).toISOString() : undefined,
    endAt: form.value.endAt ? new Date(form.value.endAt).toISOString() : undefined,
    active: form.value.active,
    freeShipping: form.value.freeShipping,
    minimumOrderSubtotal: form.value.minimumOrderSubtotal,
    minimumFreeShipping: form.value.minimumFreeShipping,
    usageLimit: form.value.usageLimit,
    perCustomerLimit: form.value.perCustomerLimit,
    allowedCustomers: parseIDs(allowedCustomerIDsInput.value),
    buyProductIds: parseIDs(buyProductIdsInput.value),
    buyQuantity: form.value.buyQuantity,
    getProductIds: parseIDs(getProductIdsInput.value),
    getQuantity: form.value.getQuantity,
    // appliesToVariants: form.value.appliesToVariants, // If added to form
  };

  try {
    await discountService.update(shopId.value, form.value.id, payload);
    await loadDiscount(); // Reload the current discount details after update
    closeForm();
  } catch (err) {
    console.error('Failed to update discount:', err.response?.data || err.message);
    alert('Failed to update discount: ' + (err.response?.data?.message || err.message));
  } finally {
    formLoading.value = false;
  }
}

// Delete with confirmation
async function confirmDelete(disc) {
  if (!shopId.value) return;
  const ok = window.confirm(`Are you sure you want to delete discount "${disc.name}"? This action cannot be undone.`);
  if (!ok) return;

  loading.value = true;
  try {
    await discountService.delete(shopId.value, disc.id);
    router.push({ name: 'Discounts' }); // Redirect to discounts list after deletion
  } catch (err) {
    console.error('Failed to delete discount:', err.response?.data || err.message);
    alert('Failed to delete discount: ' + (err.response?.data?.message || err.message));
  } finally {
    loading.value = false;
  }
}

// Initial data load when component is mounted
onMounted(() => {
  if (shopStore.activeShop?.id) {
    shopId.value = shopStore.activeShop.id;
    loadDiscount();
  }
});
</script>

<style scoped>
.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>