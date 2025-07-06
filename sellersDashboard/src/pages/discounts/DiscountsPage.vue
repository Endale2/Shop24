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

    <!-- Search and Filter Section -->
    <div class="flex flex-col sm:flex-row justify-between items-center mb-6 space-y-4 sm:space-y-0">
      <div class="w-full sm:w-1/2 relative">
        <input
          type="text"
          v-model="searchQuery"
          @input="debouncedSearch"
          placeholder="Search discounts..."
          class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm"
        />
        <SearchIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
      </div>

      <div class="flex space-x-4 items-center">
        <select
          v-model="statusFilter"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm"
        >
          <option value="">All Status</option>
          <option value="active">Active</option>
          <option value="inactive">Inactive</option>
          <option value="expired">Expired</option>
          <option value="upcoming">Upcoming</option>
        </select>

        <select
          v-model="categoryFilter"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm"
        >
          <option value="">All Categories</option>
          <option value="product">Product</option>
          <option value="order">Order</option>
          <option value="shipping">Shipping</option>
          <option value="buy_x_get_y">Buy X Get Y</option>
        </select>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="text-center py-16 text-gray-600 text-lg">
      <div class="flex flex-col items-center justify-center">
        <SpinnerIcon class="animate-spin h-10 w-10 text-green-500 mb-4" />
        <p>Loading discounts...</p>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else-if="filteredDiscounts.length === 0" class="bg-green-50 border border-green-200 text-green-700 px-6 py-10 rounded-xl text-center shadow-md">
      <div class="flex flex-col items-center">
        <div class="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mb-4">
          <TagIcon class="w-8 h-8 text-green-600" />
        </div>
        <p class="text-lg font-medium mb-3">
          {{ searchQuery || statusFilter || categoryFilter ? 'No discounts found' : 'No discounts created yet.' }}
        </p>
        <p class="text-base mb-4">
          {{ searchQuery || statusFilter || categoryFilter ? 'Try adjusting your search or filters.' : 'Click the "New Discount" button to add your first discount!' }}
        </p>
        <button
          v-if="!searchQuery && !statusFilter && !categoryFilter"
          @click="openCreateForm"
          class="inline-flex items-center px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg shadow-md hover:bg-green-700 transition duration-150 ease-in-out"
        >
          <PlusIcon class="w-4 h-4 mr-2" />
          Create First Discount
        </button>
      </div>
    </div>

    <!-- Discounts Table -->
    <div v-else class="bg-white shadow-xl rounded-xl overflow-hidden border border-gray-200">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-100">
            <tr>
              <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Discount</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Type & Value</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Eligibility</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Usage</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Status</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Validity</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-100">
            <tr
              v-for="d in filteredDiscounts"
              :key="d.id"
              class="bg-white hover:bg-green-50 transition-colors duration-150 ease-in-out"
            >
              <td class="px-6 py-4">
                <div class="flex items-start space-x-3">
                  <div class="flex-shrink-0">
                    <div class="w-10 h-10 bg-green-100 rounded-lg flex items-center justify-center">
                      <TagIcon class="w-5 h-5 text-green-600" />
                    </div>
                  </div>
                  <div class="flex-1 min-w-0">
                    <div class="text-lg font-semibold text-gray-900 hover:text-green-700 cursor-pointer" @click="goToDiscountDetail(d.id)">
                      {{ d.name }}
                    </div>
                    <p class="text-sm text-gray-500 mt-1 truncate max-w-xs">{{ d.description || 'No description' }}</p>
                    <div v-if="d.couponCode" class="mt-2">
                      <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                        <CodeIcon class="w-3 h-3 mr-1" />
                        {{ d.couponCode }}
                      </span>
                    </div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="space-y-1">
                  <div class="flex items-center space-x-2">
                    <span :class="{
                      'bg-blue-100 text-blue-800': d.category === 'product',
                      'bg-purple-100 text-purple-800': d.category === 'order',
                      'bg-orange-100 text-orange-800': d.category === 'shipping',
                      'bg-pink-100 text-pink-800': d.category === 'buy_x_get_y'
                    }" class="inline-flex px-2.5 py-0.5 rounded-full text-xs font-medium">
                      {{ formatCategory(d.category) }}
                    </span>
                    <span class="text-sm text-gray-600">{{ formatType(d.type) }}</span>
                  </div>
                  <div class="text-lg font-bold text-green-700">
                    <span v-if="d.type === 'percentage'">{{ d.value }}%</span>
                    <span v-else>${{ d.value.toFixed(2) }}</span>
                  </div>
                  <!-- Buy X Get Y specific info -->
                  <div v-if="d.category === 'buy_x_get_y'" class="text-xs text-gray-600">
                    <div>Buy {{ d.buyQuantity || 0 }}x → Get {{ d.getQuantity || 0 }}x</div>
                    <div v-if="d.buyProductIds?.length" class="text-gray-500">
                      {{ d.buyProductIds.length }} buy products
                    </div>
                    <div v-if="d.getProductIds?.length" class="text-gray-500">
                      {{ d.getProductIds.length }} get products
                    </div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="space-y-1">
                  <div class="flex items-center space-x-2">
                    <span :class="{
                      'bg-green-100 text-green-800': d.eligibilityType === 'all',
                      'bg-blue-100 text-blue-800': d.eligibilityType === 'specific',
                      'bg-purple-100 text-purple-800': d.eligibilityType === 'segment'
                    }" class="inline-flex px-2.5 py-0.5 rounded-full text-xs font-medium">
                      {{ formatEligibility(d.eligibilityType) }}
                    </span>
                  </div>
                  <div v-if="d.minimumOrderSubtotal" class="text-sm text-gray-600">
                    Min: ${{ d.minimumOrderSubtotal.toFixed(2) }}
                  </div>
                  <div v-if="d.freeShipping" class="text-sm text-green-600 font-medium">
                    Free Shipping
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="space-y-1">
                  <div class="text-sm text-gray-900">
                    {{ d.currentUsage || 0 }}
                    <span v-if="d.usageLimit" class="text-gray-500">/ {{ d.usageLimit }}</span>
                  </div>
                  <div v-if="d.perCustomerLimit" class="text-xs text-gray-500">
                    {{ d.perCustomerLimit }} per customer
                  </div>
                  <div v-if="d.usageTracking?.length" class="text-xs text-gray-500">
                    {{ d.usageTracking.length }} active users
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="flex items-center space-x-2">
                  <span
                    :class="getStatusClass(d)"
                    class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium"
                  >
                    <CheckCircleIcon v-if="isActive(d)" class="w-4 h-4 mr-1" />
                    <XCircleIcon v-else-if="!d.active" class="w-4 h-4 mr-1" />
                    <ClockIcon v-else-if="isUpcoming(d)" class="w-4 h-4 mr-1" />
                    <ExclamationIcon v-else class="w-4 h-4 mr-1" />
                    {{ getStatusText(d) }}
                  </span>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="space-y-1">
                  <div class="flex items-center text-sm text-gray-600">
                    <CalendarIcon class="w-4 h-4 mr-1 text-gray-400" />
                    {{ formatDate(d.startAt) }}
                  </div>
                  <div class="flex items-center text-sm text-gray-600">
                    <CalendarIcon class="w-4 h-4 mr-1 text-gray-400" />
                    {{ formatDate(d.endAt) }}
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="flex items-center space-x-2">
                  <button
                    @click="goToDiscountDetail(d.id)"
                    class="inline-flex items-center px-3 py-1.5 bg-blue-600 text-white text-xs font-medium rounded-md hover:bg-blue-700 transition duration-150 ease-in-out"
                  >
                    <EyeIcon class="w-3 h-3 mr-1" />
                    View
                  </button>
                  <button
                    @click="openEditForm(d)"
                    class="inline-flex items-center px-3 py-1.5 bg-green-600 text-white text-xs font-medium rounded-md hover:bg-green-700 transition duration-150 ease-in-out"
                  >
                    <PencilIcon class="w-3 h-3 mr-1" />
                    Edit
                  </button>
                  <button
                    @click="confirmDelete(d)"
                    class="inline-flex items-center px-3 py-1.5 bg-red-600 text-white text-xs font-medium rounded-md hover:bg-red-700 transition duration-150 ease-in-out"
                  >
                    <TrashIcon class="w-3 h-3 mr-1" />
                    Delete
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Create/Edit Modal -->
    <div v-if="showForm" class="fixed inset-0 bg-gray-900 bg-opacity-75 flex items-center justify-center z-50 p-4">
      <div class="bg-white rounded-xl shadow-2xl w-full max-w-4xl max-h-[90vh] overflow-y-auto p-8 relative">
        <h2 class="text-2xl font-bold mb-6 text-gray-800 text-center">
          {{ formMode === 'create' ? 'Create New Discount' : 'Edit Discount' }}
        </h2>
        
        <form @submit.prevent="submitForm" class="space-y-6">
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

          <div class="flex justify-end space-x-4 pt-6 border-t border-gray-200">
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
import { ref, computed, onMounted, watch, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { discountService } from '@/services/discount'
import { productService } from '@/services/product'
import { customerService } from '@/services/customer'
import { customerSegmentService } from '@/services/customerSegment'
import { format } from 'date-fns'
import {
  PlusIcon,
  RefreshIcon as SpinnerIcon,
  CheckCircleIcon,
  XCircleIcon,
  CalendarIcon,
  XIcon,
  PencilIcon,
  TrashIcon,
  EyeIcon,
  TagIcon,
  CodeIcon,
  SearchIcon,
  ClockIcon,
  ExclamationIcon
} from '@heroicons/vue/outline'

const shopStore = useShopStore()
const shop = computed(() => shopStore.activeShop)
const shopId = computed(() => shop.value?.id)
const router = useRouter()

// State
const discounts = ref([])
const loading = ref(false)
const searchQuery = ref('')
const statusFilter = ref('')
const categoryFilter = ref('')

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
const showForm = ref(false)
const formMode = ref('create')
const formLoading = ref(false)
const form = ref({
  id: null,
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
const appliesToInput = ref('')
const allowedCustomerIDsInput = ref('')
const buyProductIdsInput = ref('')
const getProductIdsInput = ref('')

// Computed
const filteredDiscounts = computed(() => {
  let filtered = discounts.value

  // Search filter
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(d => 
      d.name.toLowerCase().includes(query) ||
      d.description?.toLowerCase().includes(query) ||
      d.couponCode?.toLowerCase().includes(query)
    )
  }

  // Status filter
  if (statusFilter.value) {
    filtered = filtered.filter(d => {
      switch (statusFilter.value) {
        case 'active': return isActive(d)
        case 'inactive': return !d.active
        case 'expired': return isExpired(d)
        case 'upcoming': return isUpcoming(d)
        default: return true
      }
    })
  }

  // Category filter
  if (categoryFilter.value) {
    filtered = filtered.filter(d => d.category === categoryFilter.value)
  }

  return filtered
})

// Filtered results for dropdowns
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

function formatDate(iso) {
  if (!iso) return 'Not set';
  try {
    const date = new Date(iso);
    if (isNaN(date.getTime())) {
      return 'Invalid Date';
    }
    return format(date, 'MMM d, yyyy');
  } catch (e) {
    console.error("Error formatting date:", e);
    return new Date(iso).toLocaleDateString();
  }
}

function formatCategory(category) {
  const categories = {
    'product': 'Product',
    'order': 'Order',
    'shipping': 'Shipping',
    'buy_x_get_y': 'Buy X Get Y'
  }
  return categories[category] || category
}

function formatType(type) {
  const types = {
    'percentage': 'Percentage',
    'fixed': 'Fixed Amount'
  }
  return types[type] || type
}

function formatEligibility(eligibility) {
  const eligibilities = {
    'all': 'Everyone',
    'specific': 'Specific',
    'segment': 'Segments'
  }
  return eligibilities[eligibility] || eligibility
}

function isActive(discount) {
  if (!discount.active) return false
  const now = new Date()
  const start = new Date(discount.startAt)
  const end = new Date(discount.endAt)
  return now >= start && now <= end
}

function isExpired(discount) {
  const now = new Date()
  const end = new Date(discount.endAt)
  return now > end
}

function isUpcoming(discount) {
  const now = new Date()
  const start = new Date(discount.startAt)
  return now < start
}

function getStatusClass(discount) {
  if (!discount.active) return 'bg-red-100 text-red-800'
  if (isExpired(discount)) return 'bg-gray-100 text-gray-800'
  if (isUpcoming(discount)) return 'bg-yellow-100 text-yellow-800'
  return 'bg-green-100 text-green-800'
}

function getStatusText(discount) {
  if (!discount.active) return 'Inactive'
  if (isExpired(discount)) return 'Expired'
  if (isUpcoming(discount)) return 'Upcoming'
  return 'Active'
}

// Debounced search
let searchTimeout = null
function debouncedSearch() {
  clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => {
    // Search is handled by computed property
  }, 300)
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

// Navigation
function goToDiscountDetail(discountId) {
  router.push({ name: 'DiscountDetail', params: { discountId: discountId } });
}

// Form management
function openCreateForm() {
  formMode.value = 'create'
  resetForm()
  showForm.value = true
}

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
  }
  
  // Reset selected items
  selectedProducts.value = []
  selectedCustomers.value = []
  selectedSegments.value = []
  selectedBuyProducts.value = []
  selectedGetProducts.value = []
  
  // Reset search queries
  productSearchQuery.value = ''
  customerSearchQuery.value = ''
  segmentSearchQuery.value = ''
  buyProductSearchQuery.value = ''
  getProductSearchQuery.value = ''
  
  // Hide dropdowns
  showProductDropdown.value = false
  showCustomerDropdown.value = false
  showSegmentDropdown.value = false
  showBuyProductDropdown.value = false
  showGetProductDropdown.value = false
}

function openEditForm(discount) {
  formMode.value = 'edit'
  form.value.id = discount.id
  form.value.name = discount.name
  form.value.description = discount.description
  form.value.category = discount.category
  form.value.type = discount.type
  form.value.value = discount.value
  form.value.couponCode = discount.couponCode || ''
  form.value.startAt = discount.startAt ? new Date(discount.startAt).toISOString().slice(0, 16) : ''
  form.value.endAt = discount.endAt ? new Date(discount.endAt).toISOString().slice(0, 16) : ''
  form.value.active = !!discount.active
  form.value.eligibilityType = discount.eligibilityType || 'all'
  form.value.usageLimit = discount.usageLimit
  form.value.perCustomerLimit = discount.perCustomerLimit
  form.value.freeShipping = discount.freeShipping
  form.value.minimumOrderSubtotal = discount.minimumOrderSubtotal
  form.value.minimumFreeShipping = discount.minimumFreeShipping
  form.value.buyQuantity = discount.buyQuantity
  form.value.getQuantity = discount.getQuantity

  // Load data for dropdowns if not already loaded
  if (products.value.length === 0) loadProducts()
  if (customers.value.length === 0) loadCustomers()
  if (segments.value.length === 0) loadSegments()

  // Populate selected items based on existing data
  // This will be populated after the data is loaded
  setTimeout(() => {
    // Populate selected products
    selectedProducts.value = []
    if (discount.appliesToProducts?.length) {
      discount.appliesToProducts.forEach(productId => {
        const product = products.value.find(p => p.id === productId || (typeof productId === 'object' && p.id === productId.id))
        if (product) {
          selectedProducts.value.push(product)
        }
      })
    }

    // Populate selected customers
    selectedCustomers.value = []
    if (discount.allowedCustomers?.length) {
      discount.allowedCustomers.forEach(customerId => {
        const customer = customers.value.find(c => c.id === customerId)
        if (customer) {
          selectedCustomers.value.push(customer)
        }
      })
    }

    // Populate selected segments
    selectedSegments.value = []
    if (discount.allowedSegments?.length) {
      discount.allowedSegments.forEach(segmentId => {
        const segment = segments.value.find(s => s.id === segmentId)
        if (segment) {
          selectedSegments.value.push(segment)
        }
      })
    }

    // Populate Buy X Get Y products
    selectedBuyProducts.value = []
    if (discount.buyProductIds?.length) {
      discount.buyProductIds.forEach(productId => {
        const product = products.value.find(p => p.id === productId)
        if (product) {
          selectedBuyProducts.value.push(product)
        }
      })
    }

    selectedGetProducts.value = []
    if (discount.getProductIds?.length) {
      discount.getProductIds.forEach(productId => {
        const product = products.value.find(p => p.id === productId)
        if (product) {
          selectedGetProducts.value.push(product)
        }
      })
    }
  }, 100)

  showForm.value = true
}

function closeForm() {
  showForm.value = false
  resetForm()
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
    if (formMode.value === 'create') {
      await discountService.create(shopId.value, payload)
    } else if (formMode.value === 'edit' && form.value.id) {
      await discountService.update(shopId.value, form.value.id, payload)
    }
    await loadDiscounts()
    closeForm()
  } catch (err) {
    console.error('Failed to save discount:', err.response?.data || err.message)
    const errorMessage = err.response?.data?.message || err.response?.data?.error || err.message
    alert('Failed to save discount: ' + errorMessage)
  } finally {
    formLoading.value = false
  }
}

// Delete discount
async function confirmDelete(discount) {
  if (!shopId.value) return
  const ok = window.confirm(`Are you sure you want to delete discount "${discount.name}"? This action cannot be undone.`)
  if (!ok) return

  try {
    await discountService.delete(shopId.value, discount.id)
    await loadDiscounts()
  } catch (err) {
    console.error('Failed to delete discount:', err.response?.data || err.message)
    alert('Failed to delete discount: ' + (err.response?.data?.message || err.message))
  }
}

// Watch for shop changes
watch(() => shopStore.activeShop, (newShop) => {
  if (newShop?.id) {
    loadDiscounts()
    loadProducts()
    loadCustomers()
    loadSegments()
  }
})

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
    loadDiscounts()
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
/* No additional scoped styles needed as TailwindCSS handles most of the styling */
</style>