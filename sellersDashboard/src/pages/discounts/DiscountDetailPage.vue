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
              <h1 class="text-2xl sm:text-3xl font-bold text-gray-900 tracking-tight">Discount Details</h1>
              <p v-if="discount" class="text-sm text-gray-600 mt-1">{{ discount.name }}</p>
            </div>
          </div>
          <div class="flex flex-wrap gap-2" v-if="discount">
            <button
              @click="goToEditPage(discount)"
              class="inline-flex items-center px-3 py-2 bg-green-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
            >
              <PencilIcon class="w-4 h-4 mr-2" />
              Edit
            </button>
            <button
              @click="confirmDelete(discount)"
              class="inline-flex items-center px-3 py-2 bg-red-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2 transition duration-150 ease-in-out"
            >
              <TrashIcon class="w-4 h-4 mr-2" />
              Delete
            </button>
            <button
              @click="goBack"
              class="inline-flex items-center px-3 py-2 bg-gray-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 transition duration-150 ease-in-out"
            >
              <svg
                class="w-4 h-4 mr-2"
                fill="none"
                stroke="currentColor"
                stroke-width="2.5"
                viewBox="0 0 24 24"
                stroke-linecap="round"
                stroke-linejoin="round"
              >
                <path d="M15 18l-6-6 6-6" />
              </svg>
              Back
            </button>
          </div>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-16 text-gray-500">
        <div class="flex flex-col items-center justify-center">
          <div class="animate-spin h-8 w-8 text-green-500 mb-3 border-3 border-green-200 border-t-green-500 rounded-full"></div>
          <p class="text-sm font-medium">Loading discount details...</p>
        </div>
      </div>

      <!-- Error State -->
      <div v-else-if="!discount" class="bg-red-50 border border-red-200 text-red-700 px-6 py-10 rounded-xl text-center shadow-sm">
        <div class="flex flex-col items-center">
          <ExclamationIcon class="w-12 h-12 text-red-500 mb-4" />
          <p class="text-lg font-medium mb-3">Discount not found or an error occurred.</p>
          <p class="text-base mb-4">Please check the URL or try again.</p>
          <button @click="goBack"
                  class="px-5 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-100 transition-colors duration-200">
            Back to Discounts
          </button>
        </div>
      </div>

      <div v-else class="space-y-6">
        
        <!-- Main Discount Info Card -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
          <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
            <h2 class="text-sm font-semibold text-gray-900">Discount Information</h2>
          </div>
          <div class="p-4">
            <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
              <div class="lg:col-span-2">
                <div class="flex items-start space-x-4 mb-6">
                  <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center flex-shrink-0">
                    <TagIcon class="w-6 h-6 text-green-600" />
                  </div>
                  <div class="flex-1 min-w-0">
                    <h3 class="text-xl font-bold text-gray-900 mb-1">{{ discount.name }}</h3>
                    <p class="text-sm text-gray-600 mb-3 line-clamp-2">{{ discount.description || 'No description provided.' }}</p>
                    <div class="flex items-center space-x-2">
                      <span
                        :class="getStatusClass(discount)"
                        class="inline-flex items-center px-2.5 py-1 rounded-full text-xs font-medium"
                      >
                        <CheckCircleIcon v-if="isActive" class="w-3 h-3 mr-1" />
                        <XCircleIcon v-else-if="!discount.active" class="w-3 h-3 mr-1" />
                        <ClockIcon v-else-if="isUpcoming" class="w-3 h-3 mr-1" />
                        <ExclamationIcon v-else class="w-3 h-3 mr-1" />
                        {{ getStatusText(discount) }}
                      </span>
                      <span v-if="discount.couponCode" class="inline-flex items-center px-2.5 py-1 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                        <CodeIcon class="w-3 h-3 mr-1" />
                        {{ discount.couponCode }}
                      </span>
                    </div>
                  </div>
                </div>

                <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                  <div class="bg-gray-50 p-3 rounded-md">
                    <p class="text-xs font-medium text-gray-500 mb-1">Category</p>
                    <p class="text-sm font-semibold text-gray-900">{{ formatCategory(discount.category) }}</p>
                  </div>
                  <div class="bg-gray-50 p-3 rounded-md">
                    <p class="text-xs font-medium text-gray-500 mb-1">Type</p>
                    <p class="text-sm font-semibold text-gray-900">{{ formatType(discount.type) }}</p>
                  </div>
                  <div class="bg-gray-50 p-3 rounded-md">
                    <p class="text-xs font-medium text-gray-500 mb-1">Value</p>
                    <p class="text-lg font-bold text-green-700">
                      {{ discount.type === 'percentage' ? discount.value + '%' : '$' + discount.value?.toFixed(2) }}
                    </p>
                  </div>
                  <div class="bg-gray-50 p-3 rounded-md">
                    <p class="text-xs font-medium text-gray-500 mb-1">Discount Impact</p>
                    <p class="text-sm font-semibold text-gray-900">
                      {{ discount.type === 'percentage' ? `${discount.value}% off each product` : `$${discount.value?.toFixed(2)} off each product` }}
                    </p>
                  </div>
                  <div class="bg-gray-50 p-3 rounded-md">
                    <p class="text-xs font-medium text-gray-500 mb-1">Eligibility</p>
                    <p class="text-sm font-semibold text-gray-900">{{ formatEligibility(discount.eligibilityType) }}</p>
                  </div>
                </div>
              </div>

              <div class="space-y-4">
                <div class="bg-green-50 p-4 rounded-md border border-green-200">
                  <h3 class="text-xs font-medium text-green-800 mb-2">Usage Statistics</h3>
                  <div class="space-y-2">
                    <div class="flex justify-between">
                      <span class="text-xs text-green-700">Current Usage:</span>
                      <span class="text-sm font-semibold text-green-900">
                        {{ discount.currentUsage || 0 }}
                        <span v-if="discount.usageLimit" class="text-green-600">/ {{ discount.usageLimit }}</span>
                      </span>
                    </div>
                    <div v-if="discount.perCustomerLimit" class="flex justify-between">
                      <span class="text-xs text-green-700">Per Customer:</span>
                      <span class="text-sm font-semibold text-green-900">{{ discount.perCustomerLimit }} times</span>
                    </div>
                    <div v-if="discount.usageTracking?.length" class="flex justify-between">
                      <span class="text-xs text-green-700">Active Users:</span>
                      <span class="text-sm font-semibold text-green-900">{{ discount.usageTracking.length }}</span>
                    </div>
                  </div>
                </div>

                <div class="bg-blue-50 p-4 rounded-md border border-blue-200">
                  <h3 class="text-xs font-medium text-blue-800 mb-2">Validity Period</h3>
                  <div class="space-y-2">
                    <div class="flex items-center text-xs text-blue-700">
                      <CalendarIcon class="w-3 h-3 mr-2" />
                      <span>Start: {{ formatDateTime(discount.startAt) }}</span>
                    </div>
                    <div class="flex items-center text-xs text-blue-700">
                      <CalendarIcon class="w-3 h-3 mr-2" />
                      <span>End: {{ formatDateTime(discount.endAt) }}</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Customer Eligibility & Usage Section -->
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h3 class="text-sm font-semibold text-gray-900">Customer Eligibility & Usage</h3>
            </div>
            <div class="p-4">
              <div class="space-y-4">
                <div>
                  <h4 class="text-sm font-medium text-gray-900 mb-3">Eligibility Settings</h4>
                  <div class="space-y-3">
                    <!-- Eligibility Status -->
                    <div class="flex items-center justify-between p-3 rounded-md" :class="{
                      'bg-green-50 border border-green-200': discount.eligibilityType === 'all',
                      'bg-blue-50 border border-blue-200': discount.eligibilityType === 'specific' || discount.eligibilityType === 'segment'
                    }">
                      <div class="flex items-center">
                        <UsersIcon class="w-4 h-4 mr-2" :class="{
                          'text-green-600': discount.eligibilityType === 'all',
                          'text-blue-600': discount.eligibilityType === 'specific' || discount.eligibilityType === 'segment'
                        }" />
                        <div class="min-w-0">
                          <p class="text-sm font-medium text-gray-900 truncate">
                            {{ discount.eligibilityType === 'all' ? 'Everyone is eligible' : 'Specific customers are eligible' }}
                          </p>
                          <p class="text-xs text-gray-600 mt-1 line-clamp-1">
                            {{ getEligibilityDescription() }}
                          </p>
                        </div>
                      </div>
                      <span :class="{
                        'bg-green-100 text-green-800': discount.eligibilityType === 'all',
                        'bg-blue-100 text-blue-800': discount.eligibilityType === 'specific' || discount.eligibilityType === 'segment'
                      }" class="inline-flex px-2 py-0.5 rounded-full text-xs font-medium flex-shrink-0">
                        {{ formatEligibility(discount.eligibilityType) }}
                      </span>
                    </div>

                    <!-- Eligible Customers List -->
                    <div v-if="discount.eligibilityType !== 'all' && eligibleCustomers.length > 0" class="space-y-2">
                      <div class="flex items-center justify-between">
                        <h5 class="text-xs font-medium text-gray-900">Eligible Customers ({{ eligibleCustomers.length }})</h5>
                        <button @click="showEligibleCustomers = !showEligibleCustomers" 
                                class="text-xs text-blue-600 hover:text-blue-800 font-medium">
                          {{ showEligibleCustomers ? 'Hide' : 'Show' }} Details
                        </button>
                      </div>
                      
                      <div v-if="showEligibleCustomers" class="max-h-32 overflow-y-auto space-y-1">
                        <div v-for="customer in eligibleCustomers" :key="customer.id"
                             class="flex items-center justify-between p-2 bg-gray-50 rounded-md">
                          <div class="flex items-center min-w-0">
                            <div class="w-2 h-2 rounded-full mr-2 flex-shrink-0" :class="{
                              'bg-blue-500': customer.source === 'direct',
                              'bg-purple-500': customer.source === 'segment'
                            }"></div>
                            <div class="min-w-0">
                              <p class="text-xs font-medium text-gray-900 truncate">
                                {{ customer.firstName }} {{ customer.lastName }}
                              </p>
                              <p class="text-xs text-gray-500 truncate">{{ customer.email }}</p>
                            </div>
                          </div>
                          <div class="flex items-center flex-shrink-0">
                            <span v-if="customer.source === 'segment'" 
                                  class="text-xs bg-purple-100 text-purple-800 px-2 py-0.5 rounded-full">
                              {{ customer.segmentName }}
                            </span>
                            <span v-else 
                                  class="text-xs bg-blue-100 text-blue-800 px-2 py-0.5 rounded-full">
                              Direct
                            </span>
                          </div>
                        </div>
                      </div>
                    </div>

                    <!-- No Eligible Customers Message -->
                    <div v-else-if="discount.eligibilityType !== 'all'" 
                         class="text-center py-4 text-gray-500">
                      <UsersIcon class="w-8 h-8 mx-auto mb-2 text-gray-400" />
                      <p class="text-xs">No specific customers or segments added yet</p>
                    </div>
                  </div>
                </div>

                <div class="border-t border-gray-200 pt-4">
                  <h4 class="text-sm font-medium text-gray-900 mb-3">Management Actions</h4>
                  <div class="flex flex-wrap gap-2">
                    <button @click="openMixedEligibilityModal"
                            class="inline-flex items-center px-3 py-1.5 bg-green-600 text-white text-xs font-medium rounded-md hover:bg-green-700 transition duration-150 ease-in-out">
                      <PlusIcon class="w-3 h-3 mr-1" />
                      Add Customers/Segments
                    </button>
                    <button @click="clearEligibility"
                            class="inline-flex items-center px-3 py-1.5 bg-orange-600 text-white text-xs font-medium rounded-md hover:bg-orange-700 transition duration-150 ease-in-out">
                      <UsersIcon class="w-3 h-3 mr-1" />
                      Allow Everyone
                    </button>
                    <button @click="refreshUsageStats"
                            class="inline-flex items-center px-3 py-1.5 bg-gray-600 text-white text-xs font-medium rounded-md hover:bg-gray-700 transition duration-150 ease-in-out">
                      <RefreshIcon class="w-3 h-3 mr-1" />
                      Refresh Stats
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h3 class="text-sm font-semibold text-gray-900">Usage Tracking</h3>
            </div>
            <div class="p-4">
              <div class="space-y-4">
                <div class="flex justify-between items-center">
                  <span class="text-xs text-gray-600">Total Usage:</span>
                  <span class="text-sm font-semibold text-gray-900">
                    {{ discount.currentUsage || 0 }}
                    <span v-if="discount.usageLimit" class="text-gray-500">/ {{ discount.usageLimit }}</span>
                  </span>
                </div>
                
                <div v-if="discount.usageLimit" class="w-full bg-gray-200 rounded-full h-2">
                  <div 
                    class="bg-green-600 h-2 rounded-full transition-all duration-300"
                    :style="{ width: `${Math.min((discount.currentUsage || 0) / discount.usageLimit * 100, 100)}%` }"
                  ></div>
                </div>

                <div v-if="discount.usageTracking?.length" class="mt-4">
                  <div class="flex justify-between items-center mb-3">
                    <h4 class="text-xs font-medium text-gray-900">Recent Usage</h4>
                    <button @click="showUsageDetails = !showUsageDetails"
                            class="text-xs text-green-600 hover:text-green-700 font-medium">
                      {{ showUsageDetails ? 'Hide' : 'Show' }} Details
                    </button>
                  </div>
                  
                  <div v-if="showUsageDetails" class="space-y-2 max-h-32 overflow-y-auto">
                    <div v-for="(usage, idx) in discount.usageTracking.slice(0, 5)" :key="usage.customer_id"
                         class="flex justify-between items-center p-2 bg-gray-50 rounded-md text-xs">
                      <span class="text-gray-700 min-w-0">
                        <template v-if="recentUsageProfiles[idx]">
                          <span class="font-semibold truncate">{{ recentUsageProfiles[idx].firstName }} {{ recentUsageProfiles[idx].lastName }}</span>
                          <span class="text-gray-500 ml-2 truncate">{{ recentUsageProfiles[idx].email }}</span>
                        </template>
                        <template v-else>
                          {{ usage.customer_id }}
                        </template>
                      </span>
                      <span class="text-gray-600 flex-shrink-0">{{ usage.usage_count }}x</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Applied Products Section -->
        <div v-if="discount" class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
          <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
            <div class="flex items-center justify-between">
              <div class="flex items-center space-x-2">
                <h3 class="text-sm font-semibold text-gray-900">Applied Products</h3>
                <span v-if="discount && (discount.type === 'percentage' || discount.type === 'fixed')" class="px-2 py-0.5 bg-red-100 text-red-700 text-xs font-medium rounded-full">
                  Prices shown with discount
                </span>
              </div>
              <div class="flex items-center space-x-2">
                <span v-if="productsPagination.total > 0" class="px-2 py-1 bg-green-100 text-green-800 text-xs font-medium rounded-full">
                  {{ productsPagination.total }}
                </span>
                <button @click="openAddProductModal" class="px-3 py-1.5 bg-green-600 text-white rounded-md text-xs font-medium hover:bg-green-700 transition-colors duration-150">
                  <PlusIcon class="w-4 h-4 mr-1" />
                  Add Product
                </button>
              </div>
            </div>
          </div>
          <div class="overflow-x-auto">
            <div v-if="loadingProductsPagination" class="p-8 text-center">
              <div class="flex flex-col items-center justify-center">
                <div class="animate-spin h-6 w-6 text-green-500 mb-2 border-2 border-green-200 border-t-green-500 rounded-full"></div>
                <p class="text-sm text-gray-500">Loading products...</p>
              </div>
            </div>
            <div v-else-if="paginatedProducts.length > 0" class="min-w-full">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                  <tr>
                    <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Product
                    </th>
                    <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Description
                    </th>
                    <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Original Price
                    </th>
                    <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Discounted Price
                    </th>
                    <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Savings
                    </th>
                    <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Stock
                    </th>
                    <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Actions
                    </th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                  <tr v-for="prod in paginatedProducts" :key="prod.id" 
                      class="hover:bg-gray-50 transition-colors duration-150 cursor-pointer"
                      @click="goToProductDetail(prod.id)">
                    <td class="px-6 py-4 whitespace-nowrap">
                      <div class="flex items-center">
                        <div class="flex-shrink-0 h-10 w-10">
                          <img 
                            class="h-10 w-10 rounded-lg object-cover" 
                            :src="getProductImage(prod) || '/placeholder-image.png'" 
                            :alt="prod.name"
                            @error="$event.target.src = '/placeholder-image.png'"
                          />
                        </div>
                        <div class="ml-4">
                          <div class="text-sm font-medium text-gray-900 truncate max-w-32">{{ prod.name }}</div>
                        </div>
                      </div>
                    </td>
                    <td class="px-6 py-4">
                      <div class="text-sm text-gray-900 max-w-xs">
                        <p class="line-clamp-2">{{ prod.description || 'No description available.' }}</p>
                      </div>
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap">
                      <div class="text-sm text-gray-900">
                        <span class="line-through text-gray-400">{{ formatPriceWithDiscount(prod).original }}</span>
                      </div>
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap">
                      <div class="flex items-center space-x-2">
                        <span class="text-sm font-semibold text-green-700">{{ formatPriceWithDiscount(prod).discounted }}</span>
                        <span v-if="discount && (discount.type === 'percentage' || discount.type === 'fixed')" class="px-2 py-0.5 bg-red-100 text-red-700 text-xs font-medium rounded-full">
                          {{ formatPriceWithDiscount(prod).percentage }}
                        </span>
                      </div>
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap">
                      <div class="text-sm text-green-600 font-medium">
                        {{ formatPriceWithDiscount(prod).savings }}
                      </div>
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap">
                      <div class="text-sm text-gray-900">
                        <span :class="getProductStock(prod) > 0 ? 'text-green-600' : 'text-red-600'" class="font-medium">
                          {{ getProductStock(prod) }}
                        </span>
                        <span class="text-gray-500 text-xs ml-1">units</span>
                      </div>
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                      <button @click.stop="removeProductFromDiscount(prod.id)" 
                              :disabled="modalLoading" 
                              class="text-red-600 hover:text-red-900 disabled:opacity-50 disabled:cursor-not-allowed transition-colors duration-150">
                        <TrashIcon class="w-4 h-4" />
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
              
              <!-- Pagination Controls -->
              <div v-if="productsPagination.pages > 1" class="px-6 py-3 border-t border-gray-200 bg-gray-50">
                <div class="flex items-center justify-between">
                  <div class="text-sm text-gray-700">
                    Showing {{ ((productsPagination.page - 1) * productsPagination.limit) + 1 }} to 
                    {{ Math.min(productsPagination.page * productsPagination.limit, productsPagination.total) }} of 
                    {{ productsPagination.total }} products
                  </div>
                  <div class="flex items-center space-x-2">
                    <button 
                      @click="changeProductsPage(productsPagination.page - 1)"
                      :disabled="productsPagination.page <= 1"
                      class="px-3 py-1 text-sm border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
                    >
                      Previous
                    </button>
                    <span class="text-sm text-gray-700">
                      Page {{ productsPagination.page }} of {{ productsPagination.pages }}
                    </span>
                    <button 
                      @click="changeProductsPage(productsPagination.page + 1)"
                      :disabled="productsPagination.page >= productsPagination.pages"
                      class="px-3 py-1 text-sm border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
                    >
                      Next
                    </button>
                  </div>
                </div>
              </div>
            </div>
            <div v-else class="p-8 bg-gray-50 border border-gray-200 rounded-lg text-center">
              <div class="flex flex-col items-center">
                <svg class="w-12 h-12 text-gray-300 mb-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                </svg>
                <p class="text-sm font-medium text-gray-600 mb-1">No specific products selected</p>
                <p class="text-xs text-gray-500 mb-4">This discount applies to <strong>all products</strong> in the shop.</p>
                <button 
                  @click="openAddProductModal" 
                  class="inline-flex items-center px-3 py-1.5 bg-green-600 text-white rounded-md text-xs font-medium hover:bg-green-700 transition-colors duration-150"
                >
                  <PlusIcon class="w-4 h-4 mr-1" />
                  Add Specific Products
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Mixed Eligibility Modal -->
    <div v-if="showMixedEligibilityModal" class="fixed inset-0 bg-transparent flex items-center justify-center z-50 p-4">
      <div class="bg-white rounded-lg shadow-xl w-full max-w-4xl max-h-[90vh] overflow-hidden">
        <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
          <h3 class="text-sm font-semibold text-gray-900">Add Eligible Customers</h3>
        </div>
        
        <div class="p-4 overflow-y-auto max-h-[calc(90vh-120px)]">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <!-- Customers Section -->
            <div>
              <h4 class="text-sm font-medium text-gray-900 mb-3">Customers</h4>
              <div class="mb-4">
                <label class="block text-xs font-medium text-gray-700 mb-2">Search and Select Customers</label>
                <div class="relative customer-dropdown">
                  <input
                    type="text"
                    v-model="customerSearchQuery"
                    @input="filterCustomers"
                    placeholder="Search customers by name or email..."
                    class="w-full border-gray-300 rounded-md shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2 text-sm transition duration-150 ease-in-out"
                  />
                  <div v-if="filteredCustomers.length > 0 && showCustomerDropdown" class="absolute z-10 w-full mt-1 bg-white border border-gray-300 rounded-md shadow-lg max-h-48 overflow-y-auto">
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
                      <div class="flex-1 min-w-0">
                        <div class="font-medium text-gray-900 text-sm truncate">{{ customer.firstName }} {{ customer.lastName }}</div>
                        <div class="text-xs text-gray-500 truncate">{{ customer.email }}</div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              
              <div v-if="selectedCustomers.length > 0" class="mb-4">
                <h5 class="text-xs font-medium text-gray-700 mb-2">Selected Customers:</h5>
                <div class="flex flex-wrap gap-2 max-h-24 overflow-y-auto">
                  <span
                    v-for="customer in selectedCustomers"
                    :key="customer.id"
                    class="inline-flex items-center px-2 py-1 rounded-full text-xs bg-blue-100 text-blue-800"
                  >
                    <span class="truncate max-w-24">{{ customer.firstName }} {{ customer.lastName }}</span>
                    <button
                      @click="removeCustomer(customer)"
                      class="ml-1 text-blue-600 hover:text-blue-800 flex-shrink-0"
                    >
                      ×
                    </button>
                  </span>
                </div>
              </div>
            </div>

            <!-- Segments Section -->
            <div>
              <h4 class="text-sm font-medium text-gray-900 mb-3">Segments</h4>
              <div class="mb-4">
                <label class="block text-xs font-medium text-gray-700 mb-2">Search and Select Segments</label>
                <div class="relative segment-dropdown">
                  <input
                    type="text"
                    v-model="segmentSearchQuery"
                    @input="filterSegments"
                    placeholder="Search segments by name..."
                    class="w-full border-gray-300 rounded-md shadow-sm focus:border-green-500 focus:ring-green-500 px-3 py-2 text-sm transition duration-150 ease-in-out"
                  />
                  <div v-if="filteredSegments.length > 0 && showSegmentDropdown" class="absolute z-10 w-full mt-1 bg-white border border-gray-300 rounded-md shadow-lg max-h-48 overflow-y-auto">
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
                      <div class="flex-1 min-w-0">
                        <div class="font-medium text-gray-900 text-sm truncate">{{ segment.name }}</div>
                        <div class="text-xs text-gray-500 line-clamp-2">{{ segment.description || 'No description' }}</div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              
              <div v-if="selectedSegments.length > 0" class="mb-4">
                <h5 class="text-xs font-medium text-gray-700 mb-2">Selected Segments:</h5>
                <div class="flex flex-wrap gap-2 max-h-24 overflow-y-auto">
                  <span
                    v-for="segment in selectedSegments"
                    :key="segment.id"
                    class="inline-flex items-center px-2 py-1 rounded-full text-xs bg-purple-100 text-purple-800"
                  >
                    <span class="truncate max-w-24">{{ segment.name }}</span>
                    <button
                      @click="removeSegment(segment)"
                      class="ml-1 text-purple-600 hover:text-purple-800 flex-shrink-0"
                    >
                      ×
                    </button>
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
        
        <div class="px-4 py-3 border-t border-gray-100 bg-gray-50 flex justify-end space-x-3">
          <button
            type="button"
            @click="closeMixedEligibilityModal"
            class="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-100 transition-colors duration-200 text-sm"
          >
            Cancel
          </button>
          <button
            @click="addMixedEligibility"
            :disabled="(selectedCustomers.length === 0 && selectedSegments.length === 0) || modalLoading"
            class="px-4 py-2 bg-green-600 text-white rounded-md shadow-sm hover:bg-green-700 transition duration-300 ease-in-out disabled:opacity-50 disabled:cursor-not-allowed text-sm"
          >
            <span v-if="modalLoading">Adding...</span>
            <span v-else>Add {{ selectedCustomers.length + selectedSegments.length }} Customer{{ (selectedCustomers.length + selectedSegments.length) !== 1 ? 's' : '' }}</span>
          </button>
        </div>
        <button @click="closeMixedEligibilityModal" class="absolute top-4 right-4 text-gray-400 hover:text-gray-600 transition-colors duration-200">
          <XIcon class="h-5 w-5" />
        </button>
      </div>
    </div>

    <!-- Add Product Modal -->
    <div v-if="showAddProductModal" class="fixed inset-0 bg-transparent flex items-center justify-center z-50 p-4">
      <div class="bg-white rounded-lg shadow-xl w-full max-w-2xl max-h-[90vh] overflow-hidden">
        <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-2">
              <h3 class="text-sm font-semibold text-gray-900">Add Product to Discount</h3>
              <span v-if="discount && (discount.type === 'percentage' || discount.type === 'fixed')" class="px-2 py-0.5 bg-red-100 text-red-700 text-xs font-medium rounded-full">
                Prices shown with discount
              </span>
            </div>
          </div>
        </div>
        
        <div class="p-4">
          <input 
            v-model="productSearchQuery" 
            @input="debouncedProductSearch"
            placeholder="Search products by name or description..." 
            class="w-full mb-4 border-gray-300 rounded-md px-3 py-2 text-sm focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out" 
          />
          
          <div v-if="loadingProducts" class="text-center py-8 text-gray-500">
            <div class="flex flex-col items-center justify-center">
              <div class="animate-spin h-6 w-6 text-green-500 mb-2 border-2 border-green-200 border-t-green-500 rounded-full"></div>
              <p class="text-xs font-medium">Loading products...</p>
            </div>
          </div>
          <div v-else-if="allProducts.length === 0" class="text-center py-8 text-gray-500">
            <p class="text-sm">No products found in this shop.</p>
          </div>
          <div v-else-if="filteredAvailableProducts.length === 0" class="text-center py-8 text-gray-500">
            <div class="flex flex-col items-center">
              <svg class="w-12 h-12 text-gray-300 mb-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
              </svg>
              <p v-if="productSearchQuery" class="text-sm font-medium text-gray-600 mb-1">No products found matching "{{ productSearchQuery }}"</p>
              <p v-else class="text-sm font-medium text-gray-600 mb-1">All products are already added to this discount</p>
              <p class="text-xs text-gray-500">Try a different search term or add products to your shop first</p>
            </div>
          </div>
          <div v-else class="max-h-64 overflow-y-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50 sticky top-0">
                <tr>
                  <th scope="col" class="px-3 py-2 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Product
                  </th>
                  <th scope="col" class="px-3 py-2 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Original Price
                  </th>
                  <th scope="col" class="px-3 py-2 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Discounted Price
                  </th>
                  <th scope="col" class="px-3 py-2 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Stock
                  </th>
                  <th scope="col" class="px-3 py-2 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Action
                  </th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr v-for="product in filteredAvailableProducts" :key="product.id" class="hover:bg-gray-50">
                  <td class="px-3 py-2 whitespace-nowrap">
                    <div class="flex items-center">
                      <div class="flex-shrink-0 h-8 w-8">
                        <img
                          v-if="getProductImage(product)"
                          :src="getProductImage(product)"
                          alt="product"
                          class="h-8 w-8 rounded object-cover"
                          @error="$event.target.style.display = 'none'; $event.target.nextElementSibling.style.display = 'block'"
                        />
                        <svg 
                          v-else 
                          class="h-8 w-8 text-gray-400" 
                          fill="none" 
                          stroke="currentColor" 
                          viewBox="0 0 24 24"
                          :style="{ display: getProductImage(product) ? 'none' : 'block' }"
                        >
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                        </svg>
                      </div>
                      <div class="ml-3">
                        <div class="text-sm font-medium text-gray-900 truncate max-w-32">{{ product.name }}</div>
                        <div class="text-xs text-gray-500 line-clamp-1 max-w-32">{{ product.description || 'No description' }}</div>
                      </div>
                    </div>
                  </td>
                  <td class="px-3 py-2 whitespace-nowrap">
                    <div class="text-sm text-gray-900">
                      <span class="line-through text-gray-400">{{ formatPriceWithDiscount(product).original }}</span>
                    </div>
                  </td>
                  <td class="px-3 py-2 whitespace-nowrap">
                    <div class="flex items-center space-x-1">
                      <span class="text-sm font-semibold text-green-700">{{ formatPriceWithDiscount(product).discounted }}</span>
                      <span v-if="discount && (discount.type === 'percentage' || discount.type === 'fixed')" class="px-1 py-0.5 bg-red-100 text-red-700 text-xs font-medium rounded">
                        {{ formatPriceWithDiscount(product).percentage }}
                      </span>
                    </div>
                  </td>
                  <td class="px-3 py-2 whitespace-nowrap">
                    <div class="text-sm text-gray-900">
                      <span :class="getProductStock(product) > 0 ? 'text-green-600' : 'text-red-600'" class="font-medium">
                        {{ getProductStock(product) }}
                      </span>
                    </div>
                  </td>
                  <td class="px-3 py-2 whitespace-nowrap text-right text-sm font-medium">
                    <button 
                      :disabled="isProductAdded(product.id) || modalLoading" 
                      @click="addProductToDiscount(product)" 
                      class="px-2 py-1 bg-green-600 text-white rounded text-xs font-medium hover:bg-green-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors duration-150"
                    >
                      <span v-if="modalLoading">Adding...</span>
                      <span v-else>{{ isProductAdded(product.id) ? 'Added' : 'Add' }}</span>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
          
          <div class="mt-4 text-xs text-gray-600 flex justify-between items-center">
            <span>{{ filteredAvailableProducts.length }} product{{ filteredAvailableProducts.length !== 1 ? 's' : '' }} available to add</span>
            <span v-if="productSearchQuery" class="text-xs text-gray-500">Filtered by "{{ productSearchQuery }}"</span>
          </div>
        </div>
        
        <div class="px-4 py-3 border-t border-gray-100 bg-gray-50 flex justify-end">
          <button @click="showAddProductModal = false" class="px-4 py-2 bg-gray-600 text-white rounded-md text-sm hover:bg-gray-700 transition-colors duration-200">
            Close
          </button>
        </div>
        
        <button @click="showAddProductModal = false" class="absolute top-4 right-4 text-gray-400 hover:text-gray-600">
          <XIcon class="h-5 w-5" />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch, computed, onUnmounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { discountService } from '@/services/discount'
import { customerService } from '@/services/customer'
import { customerSegmentService } from '@/services/customerSegment'
import { productService } from '@/services/product'
import { format } from 'date-fns'
import {
  PencilIcon,
  TrashIcon,
  RefreshIcon as SpinnerIcon,
  CheckCircleIcon,
  XCircleIcon,
  CalendarIcon,
  XIcon,
  PlusIcon,
  RefreshIcon,
  TagIcon,
  CodeIcon,
  UsersIcon,
  ChartBarIcon,
  CubeIcon,
  ExclamationIcon,
  ClockIcon
} from '@heroicons/vue/outline'

// State
const loading = ref(true)
const discount = ref(null)
const error = ref(null)

// Router & Store
const router = useRouter()
const route = useRoute()
const shopStore = useShopStore()

const shopId = ref(shopStore.activeShop?.id)
const discountId = route.params.discountId

// Watch for shop changes
watch(() => shopStore.activeShop, (newShop) => {
  if (newShop?.id) {
    shopId.value = newShop.id
    loadDiscount()
    loadCustomers()
    loadSegments()
    loadAllProducts()
  }
})

// Enhanced discount functionality state
const showUsageDetails = ref(false)
const showMixedEligibilityModal = ref(false)
const showAddProductModal = ref(false)
const productSearchQuery = ref('')
const allProducts = ref([])
const modalLoading = ref(false)
const loadingProducts = ref(false)
let productSearchTimeout = null

// Pagination state for products under discount
const productsPagination = ref({
  page: 1,
  limit: 10,
  total: 0,
  pages: 0
})
const loadingProductsPagination = ref(false)
const paginatedProducts = ref([])

// Data for dropdowns
const customers = ref([])
const segments = ref([])
const loadingCustomers = ref(false)
const loadingSegments = ref(false)

// Search states
const customerSearchQuery = ref('')
const segmentSearchQuery = ref('')

// Dropdown visibility states
const showCustomerDropdown = ref(false)
const showSegmentDropdown = ref(false)

// Selected items
const selectedCustomers = ref([])
const selectedSegments = ref([])

// Add state for recent usage customer profiles
const recentUsageProfiles = ref([])

// Add state for eligible customers display
const eligibleCustomers = ref([])
const showEligibleCustomers = ref(false)

// Watch for discount changes and reload eligible customers
watch(
  () => discount.value,
  async (d) => {
    if (d && d.eligibilityType !== 'all') {
      await loadEligibleCustomers()
    } else {
      eligibleCustomers.value = []
    }
  },
  { immediate: true }
)

// Watch discount usageTracking and fetch customer profiles for recent usage
watch(
  () => discount.value?.usageTracking,
  async (usageList) => {
    if (!usageList || !Array.isArray(usageList) || !shopId.value) {
      recentUsageProfiles.value = []
      return
    }
    // Only fetch for the first 5 recent usages
    const ids = usageList.slice(0, 5).map(u => u.customer_id || u.customerId)
    // Deduplicate
    const uniqueIds = Array.from(new Set(ids))
    // Fetch profiles in parallel
    recentUsageProfiles.value = await Promise.all(
      uniqueIds.map(async id => {
        try {
          return await customerService.fetchById(shopId.value, id)
        } catch (e) {
          return { id, firstName: 'Unknown', lastName: '', email: '' }
        }
      })
    )
  },
  { immediate: true }
)

// Computed properties
const isActive = computed(() => {
  if (!discount.value || !discount.value.active) return false
  const now = new Date()
  const start = new Date(discount.value.startAt)
  const end = new Date(discount.value.endAt)
  return now >= start && now <= end
})

const isExpired = computed(() => {
  if (!discount.value) return false
  const now = new Date()
  const end = new Date(discount.value.endAt)
  return now > end
})

const isUpcoming = computed(() => {
  if (!discount.value) return false
  const now = new Date()
  const start = new Date(discount.value.startAt)
  return now < start
})

// Filtered results for dropdowns
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

const filteredAvailableProducts = computed(() => {
  const query = productSearchQuery.value.toLowerCase()
  return allProducts.value.filter(p =>
    (!query || 
     p.name.toLowerCase().includes(query) ||
     p.description?.toLowerCase().includes(query)) &&
    !isProductAdded(p.id)
  ).sort((a, b) => {
    // Sort by name for better UX
    return a.name.localeCompare(b.name)
  })
})

// Helpers
function formatDateTime(iso) {
  if (!iso) return 'Not set'
  try {
    const date = new Date(iso)
    if (isNaN(date.getTime())) {
      return 'Invalid Date'
    }
    return format(date, 'MMM d, yyyy HH:mm')
  } catch (e) {
    console.error("Error formatting date:", e)
    return new Date(iso).toLocaleString()
  }
}

/**
 * Gets the appropriate image for a product
 */
function getProductImage(product) {
  if (product.main_image) return product.main_image
  if (product.mainImage) return product.mainImage
  if (product.images && product.images.length > 0) return product.images[0]
  return null
}

/**
 * Load paginated products under the current discount
 */
async function loadPaginatedProducts(page = 1) {
  if (!shopId.value || !discount.value?.id) return
  
  loadingProductsPagination.value = true
  try {
    const data = await discountService.fetchProductsPaginated(shopId.value, discount.value.id, {
      page,
      limit: productsPagination.value.limit
    })
    
    paginatedProducts.value = data.products
    productsPagination.value = {
      page: data.page,
      limit: data.limit,
      total: data.total,
      pages: data.pages
    }
    
    console.log('Loaded paginated products:', {
      products: paginatedProducts.value.length,
      total: data.total,
      page: data.page,
      pages: data.pages
    })
  } catch (err) {
    console.error('Failed to load paginated products:', err)
    paginatedProducts.value = []
  } finally {
    loadingProductsPagination.value = false
  }
}

/**
 * Change page for products pagination
 */
async function changeProductsPage(newPage) {
  if (newPage < 1 || newPage > productsPagination.value.pages) return
  await loadPaginatedProducts(newPage)
}

/**
 * Gets the total stock for a product
 */
function getProductStock(product) {
  if (product.total_stock !== undefined) return product.total_stock
  if (product.stock !== undefined) return product.stock
  if (product.variants && product.variants.length > 0) {
    return product.variants.reduce((sum, v) => sum + (v.stock || 0), 0)
  }
  return 0
}

/**
 * Formats a price number.
 */
function formatPrice(p) {
  return p != null ? `$${p.toFixed(2)}` : 'N/A'
}

/**
 * Calculates the discounted price for a product based on the current discount.
 */
function calculateDiscountedPrice(product) {
  if (!discount.value || !product) {
    return product?.price || product?.Price || 0
  }
  
  // Handle both direct price field and nested price field
  const originalPrice = product.price || product.Price || 0
  const discountValue = discount.value.value || 0
  const discountType = discount.value.type || 'percentage'
  
  let discountedPrice = 0
  if (discountType === 'percentage') {
    discountedPrice = originalPrice * (1 - discountValue / 100)
  } else if (discountType === 'fixed') {
    discountedPrice = Math.max(0, originalPrice - discountValue)
  } else {
    discountedPrice = originalPrice
  }
  
  return discountedPrice
}

/**
 * Formats a price with discount information.
 */
function formatPriceWithDiscount(product) {
  if (!discount.value || !product) {
    return {
      original: formatPrice(product?.price || product?.Price || 0),
      discounted: formatPrice(product?.price || product?.Price || 0),
      savings: '$0.00',
      percentage: ''
    }
  }
  
  // Handle both direct price field and nested price field
  const originalPrice = product.price || product.Price || 0
  const discountedPrice = calculateDiscountedPrice(product)
  const discountValue = discount.value.value || 0
  const discountType = discount.value.type || 'percentage'
  
  let result = {}
  if (discountType === 'percentage') {
    result = {
      original: formatPrice(originalPrice),
      discounted: formatPrice(discountedPrice),
      savings: formatPrice(originalPrice - discountedPrice),
      percentage: `${discountValue}% OFF`
    }
  } else if (discountType === 'fixed') {
    result = {
      original: formatPrice(originalPrice),
      discounted: formatPrice(discountedPrice),
      savings: formatPrice(discountValue),
      percentage: `$${discountValue.toFixed(2)} OFF`
    }
  } else {
    result = {
      original: formatPrice(originalPrice),
      discounted: formatPrice(originalPrice),
      savings: '$0.00',
      percentage: ''
    }
  }
  
  return result
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
    'specific': 'Specific Customers',
    'segment': 'Customer Segments'
  }
  return eligibilities[eligibility] || eligibility
}

function getEligibilityDescription() {
  if (!discount.value) return ''
  
  if (discount.value.eligibilityType === 'all') {
    return 'This discount is available to all customers'
  }
  
  const directCount = discount.value.allowedCustomers?.length || 0
  const segmentCount = discount.value.allowedSegments?.length || 0
  
  if (directCount > 0 && segmentCount > 0) {
    return `${directCount} direct customers + ${segmentCount} segments`
  } else if (directCount > 0) {
    return `${directCount} direct customer${directCount !== 1 ? 's' : ''}`
  } else if (segmentCount > 0) {
    return `${segmentCount} customer segment${segmentCount !== 1 ? 's' : ''}`
  } else {
    return 'No specific customers or segments added'
  }
}

function getStatusClass(discount) {
  if (!discount.active) return 'bg-red-100 text-red-800'
  
  const now = new Date()
  const end = new Date(discount.endAt)
  const start = new Date(discount.startAt)
  
  if (now > end) return 'bg-gray-100 text-gray-800'
  if (now < start) return 'bg-yellow-100 text-yellow-800'
  return 'bg-green-100 text-green-800'
}

function getStatusText(discount) {
  if (!discount.active) return 'Inactive'
  
  const now = new Date()
  const end = new Date(discount.endAt)
  const start = new Date(discount.startAt)
  
  if (now > end) return 'Expired'
  if (now < start) return 'Upcoming'
  return 'Active'
}

// Load data for the discount detail page
async function loadDiscount() {
  if (!shopId.value || !discountId) {
    console.warn("Shop ID or Discount ID not available yet. Skipping discount load.")
    loading.value = false
    return
  }

  loading.value = true
  error.value = null
  discount.value = null

  try {
    console.log('Loading discount:', { shopId: shopId.value, discountId })
    const d = await discountService.fetchById(shopId.value, discountId)
    discount.value = d
    console.log('Loaded discount:', d)
    
    // Load eligible customers if not "all" eligibility
    if (d && d.eligibilityType !== 'all') {
      await loadEligibleCustomers()
    } else {
      eligibleCustomers.value = []
    }
    
    // Load paginated products if discount has products
    if (d && d.appliesToProducts && d.appliesToProducts.length > 0) {
      await loadPaginatedProducts(1)
    } else {
      paginatedProducts.value = []
      productsPagination.value = { page: 1, limit: 10, total: 0, pages: 0 }
    }
  } catch (e) {
    console.error('Failed to load discount:', e)
    error.value = 'Failed to load discount details. Please try again.'
  } finally {
    loading.value = false
  }
}

// Load eligible customers for the discount
async function loadEligibleCustomers() {
  if (!shopId.value || !discount.value?.id) return
  
  try {
    console.log('Loading eligible customers:', { shopId: shopId.value, discountId: discount.value.id })
    const data = await discountService.getEligibleCustomers(shopId.value, discount.value.id)
    eligibleCustomers.value = data.eligible_customers || []
    console.log('Loaded eligible customers:', eligibleCustomers.value)
  } catch (err) {
    console.error('Failed to load eligible customers:', err)
    eligibleCustomers.value = []
  }
}

function goBack() {
  router.push({ name: 'Discounts' })
}

function goToProductDetail(productId) {
  router.push({ name: 'ProductDetail', params: { productId } })
}

function goToEditPage(disc) {
  router.push({ name: 'EditDiscount', params: { discountId: disc.id } })
}

// Form management
function openEditForm(disc) {
  formMode.value = 'edit'
  form.value.id = disc.id
  form.value.name = disc.name
  form.value.description = disc.description
  form.value.category = disc.category
  form.value.type = disc.type
  form.value.value = disc.value
  form.value.couponCode = disc.couponCode || ''
  form.value.startAt = disc.startAt ? new Date(disc.startAt).toISOString().slice(0, 16) : ''
  form.value.endAt = disc.endAt ? new Date(disc.endAt).toISOString().slice(0, 16) : ''
  form.value.active = !!disc.active
  form.value.freeShipping = disc.freeShipping
  form.value.minimumOrderSubtotal = disc.minimumOrderSubtotal
  form.value.minimumFreeShipping = disc.minimumFreeShipping
  form.value.usageLimit = disc.usageLimit
  form.value.perCustomerLimit = disc.perCustomerLimit
  form.value.buyQuantity = disc.buyQuantity
  form.value.getQuantity = disc.getQuantity

  showForm.value = true
}

function closeForm() {
  showForm.value = false;
}

// Submit form
async function submitForm() {
  if (!shopId.value || !form.value.id) {
    console.error('Missing shop ID or discount ID for update operation.')
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

  const payload = {
    name: form.value.name,
    description: form.value.description,
    category: form.value.category,
    type: form.value.type,
    value: form.value.value,
    couponCode: form.value.couponCode || undefined,
    startAt: form.value.startAt,
    endAt: form.value.endAt,
    active: form.value.active,
    freeShipping: form.value.freeShipping,
    minimumOrderSubtotal: form.value.minimumOrderSubtotal,
    minimumFreeShipping: form.value.minimumFreeShipping,
    usageLimit: form.value.usageLimit,
    perCustomerLimit: form.value.perCustomerLimit,
  }

  try {
    await discountService.update(shopId.value, form.value.id, payload)
    await loadDiscount()
    closeForm()
  } catch (err) {
    console.error('Failed to update discount:', err.response?.data || err.message)
    const errorMessage = err.response?.data?.message || err.response?.data?.error || err.message
    alert('Failed to update discount: ' + errorMessage)
  } finally {
    formLoading.value = false
  }
}

// Delete with confirmation
async function confirmDelete(disc) {
  if (!shopId.value) return
  const ok = window.confirm(`Are you sure you want to delete discount "${disc.name}"? This action cannot be undone.`)
  if (!ok) return

  loading.value = true
  try {
    await discountService.delete(shopId.value, disc.id)
    router.push({ name: 'Discounts' })
  } catch (err) {
    console.error('Failed to delete discount:', err.response?.data || err.message)
    alert('Failed to delete discount: ' + (err.response?.data?.message || err.message))
  } finally {
    loading.value = false
  }
}

// Enhanced discount management functions
function openMixedEligibilityModal() {
  showMixedEligibilityModal.value = true
  selectedCustomers.value = []
  selectedSegments.value = []
  customerSearchQuery.value = ''
  segmentSearchQuery.value = ''
  showCustomerDropdown.value = false
  showSegmentDropdown.value = false
  
  // Always reload data to ensure we have the latest
  loadCustomers()
  loadSegments()
}

function openAddProductModal() {
  showAddProductModal.value = true
  productSearchQuery.value = ''
  
  // Always reload products to ensure we have the latest
  loadAllProducts()
}

function closeMixedEligibilityModal() {
  showMixedEligibilityModal.value = false
  selectedCustomers.value = []
  selectedSegments.value = []
  customerSearchQuery.value = ''
  segmentSearchQuery.value = ''
  showCustomerDropdown.value = false
  showSegmentDropdown.value = false
}

async function addMixedEligibility() {
  if (!discount.value || (selectedCustomers.value.length === 0 && selectedSegments.value.length === 0)) {
    alert('Please select at least one customer or segment')
    return
  }
  
  modalLoading.value = true
  try {
    // Deduplicate IDs before sending
    const customerIds = Array.from(new Set(selectedCustomers.value.map(c => c.id)))
    const segmentIds = Array.from(new Set(selectedSegments.value.map(s => s.id)))
    
    console.log('Adding mixed eligibility:', {
      customerIds,
      segmentIds,
      discountId: discount.value.id,
      shopId: shopId.value
    })
    
    await discountService.addMixedEligibility(shopId.value, discount.value.id, customerIds, segmentIds)
    await loadDiscount()
    closeMixedEligibilityModal()
    alert('Eligibility updated successfully!')
  } catch (err) {
    console.error('Failed to add mixed eligibility:', err)
    alert('Failed to add mixed eligibility: ' + (err.response?.data?.message || err.message))
  } finally {
    modalLoading.value = false
  }
}

async function refreshUsageStats() {
  if (!shopId.value || !discount.value?.id) return
  
  try {
    const stats = await discountService.getUsageStats(shopId.value, discount.value.id)
    discount.value.currentUsage = stats.current_usage
    discount.value.usageTracking = stats.usage_tracking
    alert('Usage statistics refreshed!')
  } catch (error) {
    console.error('Failed to refresh usage stats:', error)
    alert('Failed to refresh usage stats: ' + (error.response?.data?.error || error.message))
  }
}

// Add a method to clear eligibility
async function clearEligibility() {
  if (!shopId.value || !discount.value?.id) return
  if (!confirm('Are you sure you want to allow everyone for this discount? This will make the discount available to all customers.')) return
  await discountService.clearEligibility(shopId.value, discount.value.id)
  await loadDiscount()
}

// Load data for dropdowns
async function loadCustomers() {
  if (!shopId.value) return
  loadingCustomers.value = true
  try {
    const customerData = await customerService.fetchAll(shopId.value)
    customers.value = customerData?.customers || []
    console.log('Loaded customers:', customers.value.length)
  } catch (err) {
    console.error('Failed to load customers:', err)
    customers.value = []
  } finally {
    loadingCustomers.value = false
  }
}

async function loadSegments() {
  if (!shopId.value) return
  loadingSegments.value = true
  try {
    const segmentData = await customerSegmentService.fetchAll(shopId.value)
    segments.value = Array.isArray(segmentData) ? segmentData : []
    console.log('Loaded segments:', segments.value.length)
  } catch (err) {
    console.error('Failed to load segments:', err)
    segments.value = []
  } finally {
    loadingSegments.value = false
  }
}

// Load all products for the shop
async function loadAllProducts() {
  if (!shopId.value) return
  loadingProducts.value = true
  try {
    console.log('Loading all products for shop:', shopId.value)
    // Use the paginated version for better performance and consistency with ProductsPage
    const { products: productData, total, pages } = await productService.fetchAllByShopPaginated(shopId.value, {
      page: 1,
      limit: 1000, // Get all products for this use case
      search: '',
      stockStatus: ''
    })
    allProducts.value = Array.isArray(productData) ? productData : []
    productsPagination.value.total = total
    productsPagination.value.pages = pages
    console.log('Loaded products:', allProducts.value.length)
  } catch (err) {
    console.error('Failed to load products:', err)
    allProducts.value = []
    // Show user-friendly error
    if (err.response?.status === 403) {
      alert('You do not have permission to view products for this shop.')
    } else if (err.response?.status === 404) {
      alert('Shop not found. Please check your shop selection.')
    } else {
      alert('Failed to load products. Please try again.')
    }
  } finally {
    loadingProducts.value = false
  }
}

// Dropdown filter functions
function filterCustomers() {
  showCustomerDropdown.value = true
}

function filterSegments() {
  showSegmentDropdown.value = true
}

/**
 * Applies the search filter to the products with debouncing.
 */
function debouncedProductSearch() {
  clearTimeout(productSearchTimeout)
  productSearchTimeout = setTimeout(() => {
    // The filteredAvailableProducts computed property will automatically update
    console.log('Product search applied:', productSearchQuery.value)
  }, 300)
}

// Selection functions
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

// Remove functions
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

// Product management functions
function isProductAdded(productId) {
  return (discount.value?.appliesToProducts || []).some(p => p.id === productId)
}

async function addProductToDiscount(product) {
  if (!product || !product.id) {
    alert('Invalid product data')
    return
  }
  
  modalLoading.value = true
  try {
    const currentProducts = discount.value.appliesToProducts || []
    
    // Check if product is already added
    if (currentProducts.some(p => p.id === product.id)) {
      alert('This product is already added to the discount')
      return
    }
    
    const newProducts = [...currentProducts, product]
    const productIds = newProducts.map(p => p.id)
    
    console.log('Adding product to discount:', {
      productId: product.id,
      productName: product.name,
      currentProductIds: currentProducts.map(p => p.id),
      newProductIds: productIds,
      discountId: discount.value.id,
      shopId: shopId.value
    })
    
    await discountService.update(shopId.value, discount.value.id, { 
      appliesToProducts: productIds
    })
    await loadDiscount()
    
    // Refresh paginated products
    if (discount.value && discount.value.appliesToProducts && discount.value.appliesToProducts.length > 0) {
      await loadPaginatedProducts(1) // Reset to first page
    }
    
    // Clear search and close modal
    productSearchQuery.value = ''
    showAddProductModal.value = false
    
    // Show success message
    alert(`"${product.name}" has been added to the discount successfully!`)
  } catch (err) {
    console.error('Failed to add product to discount:', err)
    const errorMessage = err.response?.data?.message || err.response?.data?.error || err.message
    alert('Failed to add product to discount: ' + errorMessage)
  } finally {
    modalLoading.value = false
  }
}

async function removeProductFromDiscount(productId) {
  if (!productId) {
    alert('Invalid product ID')
    return
  }
  
  // Find the product to get its name for the confirmation
  const product = discount.value.appliesToProducts?.find(p => p.id === productId)
  if (!product) {
    alert('Product not found in discount')
    return
  }
  
  // Confirm removal
  const confirmed = confirm(`Are you sure you want to remove "${product.name}" from this discount?`)
  if (!confirmed) return
  
  modalLoading.value = true
  try {
    const currentProducts = discount.value.appliesToProducts || []
    const newProducts = currentProducts.filter(p => p.id !== productId)
    const productIds = newProducts.map(p => p.id)
    
    console.log('Removing product from discount:', {
      productId,
      productName: product.name,
      currentProductIds: currentProducts.map(p => p.id),
      newProductIds: productIds,
      discountId: discount.value.id,
      shopId: shopId.value
    })
    
    await discountService.update(shopId.value, discount.value.id, { 
      appliesToProducts: productIds
    })
    await loadDiscount()
    
    // Refresh paginated products
    if (discount.value && discount.value.appliesToProducts && discount.value.appliesToProducts.length > 0) {
      await loadPaginatedProducts(1) // Reset to first page
    } else {
      paginatedProducts.value = []
      productsPagination.value = { page: 1, limit: 10, total: 0, pages: 0 }
    }
    
    // Show success message
    alert(`"${product.name}" has been removed from the discount successfully!`)
  } catch (err) {
    console.error('Failed to remove product from discount:', err)
    const errorMessage = err.response?.data?.message || err.response?.data?.error || err.message
    alert('Failed to remove product from discount: ' + errorMessage)
  } finally {
    modalLoading.value = false
  }
}

// Click outside handlers
function handleClickOutside(event) {
  // Close customer dropdown
  if (!event.target.closest('.customer-dropdown')) {
    showCustomerDropdown.value = false
  }
  
  // Close segment dropdown
  if (!event.target.closest('.segment-dropdown')) {
    showSegmentDropdown.value = false
  }
}

// Initial data load when component is mounted
onMounted(() => {
  if (shopId.value) {
    loadDiscount()
    loadCustomers()
    loadSegments()
    loadAllProducts()
  }
  
  // Add click outside listener
  document.addEventListener('click', handleClickOutside)
})

// Clean up event listener and timeouts
onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
  if (productSearchTimeout) {
    clearTimeout(productSearchTimeout)
  }
})
</script>

<style scoped>
.line-clamp-1 {
  display: -webkit-box;
  -webkit-line-clamp: 1;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* Custom scrollbar for better UX */
::-webkit-scrollbar {
  width: 4px;
}

::-webkit-scrollbar-track {
  background: transparent;
}

::-webkit-scrollbar-thumb {
  background: #d1d5db;
  border-radius: 2px;
}

::-webkit-scrollbar-thumb:hover {
  background: #9ca3af;
}

/* Smooth transitions */
* {
  transition-property: color, background-color, border-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
}
</style>