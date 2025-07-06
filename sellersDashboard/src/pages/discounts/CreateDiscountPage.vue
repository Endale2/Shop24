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
            Set up promotional offers to boost your sales
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
        <div class="animate-spin h-10 w-10 text-green-500 mb-4 border-4 border-green-200 border-t-green-500 rounded-full"></div>
        <p>Loading form data...</p>
      </div>
    </div>

    <!-- Form -->
    <div v-else class="bg-white rounded-xl shadow-lg border border-gray-200">
      <form @submit.prevent="submitForm" class="p-8 space-y-8">
        <!-- Basic Information -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h3 class="text-lg font-medium text-gray-900 mb-4">Basic Information</h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label for="name" class="block text-sm font-medium text-gray-700 mb-1">Name *</label>
              <input id="name" v-model="form.name" type="text" required
                     class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out" />
            </div>
            <div>
              <label for="couponCode" class="block text-sm font-medium text-gray-700 mb-1">Coupon Code (optional)</label>
              <input id="couponCode" v-model="form.couponCode" type="text"
                     class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out" />
            </div>
          </div>
          <div class="mt-4">
            <label for="description" class="block text-sm font-medium text-gray-700 mb-1">Description</label>
            <textarea id="description" v-model="form.description" rows="3"
                      class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"></textarea>
          </div>
        </div>

        <!-- Discount Configuration -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h3 class="text-lg font-medium text-gray-900 mb-4">Discount Configuration</h3>
          <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div>
              <label for="category" class="block text-sm font-medium text-gray-700 mb-1">Category *</label>
              <select id="category" v-model="form.category" required
                      class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 bg-white transition duration-150 ease-in-out">
                <option value="product">Product Discount</option>
                <option value="order">Order Discount</option>
                <option value="shipping">Shipping Discount</option>
                <option value="buy_x_get_y">Buy X Get Y</option>
              </select>
            </div>
            <div>
              <label for="type" class="block text-sm font-medium text-gray-700 mb-1">Type *</label>
              <select id="type" v-model="form.type" required
                      class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 bg-white transition duration-150 ease-in-out">
                <option value="percentage">Percentage</option>
                <option value="fixed">Fixed Amount</option>
              </select>
            </div>
            <div>
              <label for="value" class="block text-sm font-medium text-gray-700 mb-1">Value *</label>
              <input
                id="value"
                v-model.number="form.value"
                type="number"
                min="0"
                :step="form.type === 'percentage' ? 1 : 0.01"
                required
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
            </div>
          </div>
        </div>

        <!-- Category-specific Configuration -->
        <div v-if="form.category === 'product' || form.category === 'buy_x_get_y'" class="bg-gray-50 p-6 rounded-lg">
          <h3 class="text-lg font-medium text-gray-900 mb-4">Product Configuration</h3>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Select Products</label>
            <div class="relative product-dropdown">
              <input
                type="text"
                v-model="productSearchQuery"
                @input="filterProducts"
                placeholder="Search products..."
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
                  class="inline-flex items-center px-3 py-1 rounded-full text-sm bg-green-100 text-green-800"
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
              Select products to apply this discount to. Leave empty to apply to all products in this shop.
            </p>
          </div>
        </div>

        <div v-if="form.category === 'buy_x_get_y'" class="bg-gray-50 p-6 rounded-lg">
          <h3 class="text-lg font-medium text-gray-900 mb-4">Buy X Get Y Configuration</h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">Buy Products</label>
              <div class="relative buy-product-dropdown">
                <input
                  type="text"
                  v-model="buyProductSearchQuery"
                  @input="filterBuyProducts"
                  placeholder="Search products to buy..."
                  class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
                />
                <div v-if="filteredBuyProducts.length > 0 && showBuyProductDropdown" class="absolute z-10 w-full mt-1 bg-white border border-gray-300 rounded-lg shadow-lg max-h-60 overflow-y-auto">
                  <div
                    v-for="product in filteredBuyProducts"
                    :key="product.id"
                    @click="toggleBuyProductSelection(product)"
                    class="flex items-center p-3 hover:bg-gray-100 cursor-pointer border-b border-gray-100 last:border-b-0"
                  >
                    <input
                      type="checkbox"
                      :checked="selectedBuyProducts.some(p => p.id === product.id)"
                      class="mr-3 h-4 w-4 text-blue-600 rounded focus:ring-blue-500"
                      readonly
                    />
                    <div class="flex-1">
                      <div class="font-medium text-gray-900">{{ product.name }}</div>
                      <div class="text-sm text-gray-500">${{ product.price?.toFixed(2) || '0.00' }}</div>
                    </div>
                  </div>
                </div>
              </div>
              <div v-if="selectedBuyProducts.length > 0" class="mt-2">
                <div class="flex flex-wrap gap-1">
                  <span
                    v-for="product in selectedBuyProducts"
                    :key="product.id"
                    class="inline-flex items-center px-2 py-1 rounded text-xs bg-blue-100 text-blue-800"
                  >
                    {{ product.name }}
                    <button
                      @click="removeBuyProduct(product)"
                      class="ml-1 text-blue-600 hover:text-blue-800"
                    >
                      ×
                    </button>
                  </span>
                </div>
              </div>
            </div>
            <div>
              <label for="buyQuantity" class="block text-sm font-medium text-gray-700 mb-1">Buy Quantity</label>
              <input
                id="buyQuantity"
                v-model.number="form.buyQuantity"
                type="number"
                min="1"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">Get Products</label>
              <div class="relative get-product-dropdown">
                <input
                  type="text"
                  v-model="getProductSearchQuery"
                  @input="filterGetProducts"
                  placeholder="Search products to get..."
                  class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
                />
                <div v-if="filteredGetProducts.length > 0 && showGetProductDropdown" class="absolute z-10 w-full mt-1 bg-white border border-gray-300 rounded-lg shadow-lg max-h-60 overflow-y-auto">
                  <div
                    v-for="product in filteredGetProducts"
                    :key="product.id"
                    @click="toggleGetProductSelection(product)"
                    class="flex items-center p-3 hover:bg-gray-100 cursor-pointer border-b border-gray-100 last:border-b-0"
                  >
                    <input
                      type="checkbox"
                      :checked="selectedGetProducts.some(p => p.id === product.id)"
                      class="mr-3 h-4 w-4 text-purple-600 rounded focus:ring-purple-500"
                      readonly
                    />
                    <div class="flex-1">
                      <div class="font-medium text-gray-900">{{ product.name }}</div>
                      <div class="text-sm text-gray-500">${{ product.price?.toFixed(2) || '0.00' }}</div>
                    </div>
                  </div>
                </div>
              </div>
              <div v-if="selectedGetProducts.length > 0" class="mt-2">
                <div class="flex flex-wrap gap-1">
                  <span
                    v-for="product in selectedGetProducts"
                    :key="product.id"
                    class="inline-flex items-center px-2 py-1 rounded text-xs bg-purple-100 text-purple-800"
                  >
                    {{ product.name }}
                    <button
                      @click="removeGetProduct(product)"
                      class="ml-1 text-purple-600 hover:text-purple-800"
                    >
                      ×
                    </button>
                  </span>
                </div>
              </div>
            </div>
            <div>
              <label for="getQuantity" class="block text-sm font-medium text-gray-700 mb-1">Get Quantity</label>
              <input
                id="getQuantity"
                v-model.number="form.getQuantity"
                type="number"
                min="1"
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
            </div>
          </div>
        </div>

        <div v-if="form.category === 'order'" class="bg-gray-50 p-6 rounded-lg">
          <h3 class="text-lg font-medium text-gray-900 mb-4">Order Configuration</h3>
          <div>
            <label for="minimumOrderSubtotal" class="block text-sm font-medium text-gray-700 mb-1">Minimum Order Subtotal</label>
            <input
              id="minimumOrderSubtotal"
              v-model.number="form.minimumOrderSubtotal"
              type="number"
              min="0"
              step="0.01"
              placeholder="e.g. 50.00"
              class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
            />
          </div>
        </div>

        <div v-if="form.category === 'shipping'" class="bg-gray-50 p-6 rounded-lg">
          <h3 class="text-lg font-medium text-gray-900 mb-4">Shipping Configuration</h3>
          <div class="flex items-center space-x-2 mb-4">
            <input id="freeShipping" type="checkbox" v-model="form.freeShipping"
                   class="form-checkbox h-5 w-5 text-green-600 rounded focus:ring-green-500 transition duration-150 ease-in-out" />
            <label for="freeShipping" class="text-base font-medium text-gray-800">Enable Free Shipping</label>
          </div>
          <div v-if="form.freeShipping">
            <label for="minimumFreeShipping" class="block text-sm font-medium text-gray-700 mb-1">Minimum Order for Free Shipping</label>
            <input
              id="minimumFreeShipping"
              v-model.number="form.minimumFreeShipping"
              type="number"
              min="0"
              step="0.01"
              placeholder="e.g. 100.00"
              class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
            />
          </div>
        </div>

        <!-- Customer Eligibility -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h3 class="text-lg font-medium text-gray-900 mb-4">Customer Eligibility</h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
            <div>
              <label for="eligibilityType" class="block text-sm font-medium text-gray-700 mb-1">Eligibility Type</label>
              <select id="eligibilityType" v-model="form.eligibilityType"
                      class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 bg-white transition duration-150 ease-in-out">
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
                  class="inline-flex items-center px-3 py-1 rounded-full text-sm bg-blue-100 text-blue-800"
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
                  class="inline-flex items-center px-3 py-1 rounded-full text-sm bg-purple-100 text-purple-800"
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
          <h3 class="text-lg font-medium text-gray-900 mb-4">Usage Limits</h3>
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
            </div>
          </div>
        </div>

        <!-- Validity Period -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <h3 class="text-lg font-medium text-gray-900 mb-4">Validity Period</h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label for="startAt" class="block text-sm font-medium text-gray-700 mb-1">Start Date & Time</label>
              <input
                id="startAt"
                v-model="form.startAt"
                type="datetime-local"
                required
                class="w-full border-gray-300 rounded-lg shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2.5 transition duration-150 ease-in-out"
              />
            </div>
            <div>
              <label for="endAt" class="block text-sm font-medium text-gray-700 mb-1">End Date & Time</label>
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
            <input id="active" type="checkbox" v-model="form.active"
                   class="form-checkbox h-5 w-5 text-green-600 rounded focus:ring-green-500 transition duration-150 ease-in-out" />
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
            :disabled="formLoading"
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
import { ArrowLeftIcon } from '@heroicons/vue/outline'

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
const buyProductSearchQuery = ref('')
const getProductSearchQuery = ref('')

// Dropdown visibility states
const showProductDropdown = ref(false)
const showCustomerDropdown = ref(false)
const showSegmentDropdown = ref(false)
const showBuyProductDropdown = ref(false)
const showGetProductDropdown = ref(false)

// Selected items
const selectedProducts = ref([])
const selectedCustomers = ref([])
const selectedSegments = ref([])
const selectedBuyProducts = ref([])
const selectedGetProducts = ref([])

// Form state
const form = ref({
  name: '',
  description: '',
  category: 'product',
  type: 'percentage',
  value: 0,
  appliesToProducts: [],
  appliesToCollections: [],
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
  freeShipping: false,
  minimumOrderSubtotal: null,
  minimumFreeShipping: null,
  buyProductIds: [],
  buyQuantity: null,
  getProductIds: [],
  getQuantity: null,
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

const filteredBuyProducts = computed(() => {
  if (!buyProductSearchQuery.value) return products.value.slice(0, 10)
  const query = buyProductSearchQuery.value.toLowerCase()
  return products.value.filter(p => 
    p.name.toLowerCase().includes(query) ||
    p.description?.toLowerCase().includes(query)
  ).slice(0, 10)
})

const filteredGetProducts = computed(() => {
  if (!getProductSearchQuery.value) return products.value.slice(0, 10)
  const query = getProductSearchQuery.value.toLowerCase()
  return products.value.filter(p => 
    p.name.toLowerCase().includes(query) ||
    p.description?.toLowerCase().includes(query)
  ).slice(0, 10)
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

function filterBuyProducts() {
  showBuyProductDropdown.value = true
}

function filterGetProducts() {
  showGetProductDropdown.value = true
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

function toggleBuyProductSelection(product) {
  const index = selectedBuyProducts.value.findIndex(p => p.id === product.id)
  if (index > -1) {
    selectedBuyProducts.value.splice(index, 1)
  } else {
    selectedBuyProducts.value.push(product)
  }
  showBuyProductDropdown.value = false
}

function toggleGetProductSelection(product) {
  const index = selectedGetProducts.value.findIndex(p => p.id === product.id)
  if (index > -1) {
    selectedGetProducts.value.splice(index, 1)
  } else {
    selectedGetProducts.value.push(product)
  }
  showGetProductDropdown.value = false
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

function removeBuyProduct(product) {
  const index = selectedBuyProducts.value.findIndex(p => p.id === product.id)
  if (index > -1) {
    selectedBuyProducts.value.splice(index, 1)
  }
}

function removeGetProduct(product) {
  const index = selectedGetProducts.value.findIndex(p => p.id === product.id)
  if (index > -1) {
    selectedGetProducts.value.splice(index, 1)
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

  // Validate Buy X Get Y fields
  if (form.value.category === 'buy_x_get_y') {
    if (selectedBuyProducts.value.length === 0) {
      alert('Please specify Buy Products for Buy X Get Y discount')
      return
    }
    if (!form.value.buyQuantity || form.value.buyQuantity < 1) {
      alert('Please specify a valid Buy Quantity (minimum 1)')
      return
    }
    if (selectedGetProducts.value.length === 0) {
      alert('Please specify Get Products for Buy X Get Y discount')
      return
    }
    if (!form.value.getQuantity || form.value.getQuantity < 1) {
      alert('Please specify a valid Get Quantity (minimum 1)')
      return
    }
  }

  // Validate order-level discounts
  if (form.value.category === 'order' && form.value.minimumOrderSubtotal !== null && form.value.minimumOrderSubtotal <= 0) {
    alert('Minimum order subtotal must be greater than 0')
    return
  }

  formLoading.value = true

  const payload = {
    name: form.value.name,
    description: form.value.description,
    category: form.value.category,
    type: form.value.type,
    value: form.value.value,
    appliesToProducts: selectedProducts.value.map(p => p.id),
    appliesToCollections: form.value.appliesToCollections,
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
    freeShipping: form.value.freeShipping,
    minimumOrderSubtotal: form.value.minimumOrderSubtotal,
    minimumFreeShipping: form.value.minimumFreeShipping,
    buyProductIds: selectedBuyProducts.value.map(p => p.id),
    buyQuantity: form.value.buyQuantity,
    getProductIds: selectedGetProducts.value.map(p => p.id),
    getQuantity: form.value.getQuantity,
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
  
  // Close buy product dropdown
  if (!event.target.closest('.buy-product-dropdown')) {
    showBuyProductDropdown.value = false
  }
  
  // Close get product dropdown
  if (!event.target.closest('.get-product-dropdown')) {
    showGetProductDropdown.value = false
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