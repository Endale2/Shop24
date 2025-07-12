<template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto font-sans">
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center mb-8 space-y-4 sm:space-y-0">
      <div>
        <h1 class="text-3xl sm:text-4xl font-extrabold text-gray-900 leading-tight">Discount Details</h1>
        <p v-if="discount" class="text-lg text-gray-600 mt-2">{{ discount.name }}</p>
      </div>
      <div class="flex flex-wrap gap-3" v-if="discount">
        <button
          @click="goToEditPage(discount)"
          class="inline-flex items-center px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg shadow-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition duration-150 ease-in-out"
        >
          <PencilIcon class="w-5 h-5 mr-2 -ml-1" />
          Edit Discount
        </button>
        <button
          @click="confirmDelete(discount)"
          class="inline-flex items-center px-4 py-2 bg-red-600 text-white text-sm font-medium rounded-lg shadow-md hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2 transition duration-150 ease-in-out"
        >
          <TrashIcon class="w-5 h-5 mr-2 -ml-1" />
          Delete Discount
        </button>
        <button
          @click="goBack"
          class="inline-flex items-center px-4 py-2 bg-gray-600 text-white text-sm font-medium rounded-lg shadow-md hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 transition duration-150 ease-in-out"
        >
          <ArrowLeftIcon class="w-5 h-5 mr-2 -ml-1" />
          Back to List
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

    <div v-else class="space-y-8">
      <div class="bg-white shadow-xl rounded-xl p-6 sm:p-8 border border-gray-200">
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
          <div class="lg:col-span-2">
            <div class="flex items-start space-x-4 mb-6">
              <div class="w-16 h-16 bg-green-100 rounded-xl flex items-center justify-center">
                <TagIcon class="w-8 h-8 text-green-600" />
              </div>
              <div class="flex-1">
                <h2 class="text-2xl font-bold text-gray-900">{{ discount.name }}</h2>
                <p class="text-gray-600 mt-1">{{ discount.description || 'No description provided.' }}</p>
                <div class="flex items-center space-x-2 mt-3">
                  <span
                    :class="getStatusClass(discount)"
                    class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium"
                  >
                    <CheckCircleIcon v-if="isActive" class="w-4 h-4 mr-1" />
                    <XCircleIcon v-else-if="!discount.active" class="w-4 h-4 mr-1" />
                    <ClockIcon v-else-if="isUpcoming" class="w-4 h-4 mr-1" />
                    <ExclamationIcon v-else class="w-4 h-4 mr-1" />
                    {{ getStatusText(discount) }}
                  </span>
                  <span v-if="discount.couponCode" class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-blue-100 text-blue-800">
                    <CodeIcon class="w-4 h-4 mr-1" />
                    {{ discount.couponCode }}
                  </span>
                </div>
              </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div class="bg-gray-50 p-4 rounded-lg">
                <h3 class="text-sm font-medium text-gray-500 mb-2">Category</h3>
                <p class="text-lg font-semibold text-gray-900">{{ formatCategory(discount.category) }}</p>
              </div>
              <div class="bg-gray-50 p-4 rounded-lg">
                <h3 class="text-sm font-medium text-gray-500 mb-2">Type</h3>
                <p class="text-lg font-semibold text-gray-900">{{ formatType(discount.type) }}</p>
              </div>
              <div class="bg-gray-50 p-4 rounded-lg">
                <h3 class="text-sm font-medium text-gray-500 mb-2">Value</h3>
                <p class="text-2xl font-bold text-green-700">
                  {{ discount.type === 'percentage' ? discount.value + '%' : '$' + discount.value?.toFixed(2) }}
                </p>
              </div>
              <div class="bg-gray-50 p-4 rounded-lg">
                <h3 class="text-sm font-medium text-gray-500 mb-2">Eligibility</h3>
                <p class="text-lg font-semibold text-gray-900">{{ formatEligibility(discount.eligibilityType) }}</p>
              </div>
              <div v-if="discount.category === 'buy_x_get_y'" class="bg-pink-50 p-4 rounded-lg border border-pink-200">
                <h3 class="text-sm font-medium text-pink-700 mb-2">Buy X Get Y</h3>
                <div class="space-y-1">
                  <p class="text-lg font-semibold text-pink-900">
                    Buy {{ discount.buyQuantity || 0 }}x → Get {{ discount.getQuantity || 0 }}x
                  </p>
                  <div v-if="discount.buyProductIds?.length" class="text-sm text-pink-700">
                    {{ discount.buyProductIds.length }} buy products
                  </div>
                  <div v-if="discount.getProductIds?.length" class="text-sm text-pink-700">
                    {{ discount.getProductIds.length }} get products
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div class="space-y-4">
            <div class="bg-green-50 p-4 rounded-lg border border-green-200">
              <h3 class="text-sm font-medium text-green-800 mb-2">Usage Statistics</h3>
              <div class="space-y-2">
                <div class="flex justify-between">
                  <span class="text-sm text-green-700">Current Usage:</span>
                  <span class="font-semibold text-green-900">
                    {{ discount.currentUsage || 0 }}
                    <span v-if="discount.usageLimit" class="text-green-600">/ {{ discount.usageLimit }}</span>
                  </span>
                </div>
                <div v-if="discount.perCustomerLimit" class="flex justify-between">
                  <span class="text-sm text-green-700">Per Customer:</span>
                  <span class="font-semibold text-green-900">{{ discount.perCustomerLimit }} times</span>
                </div>
                <div v-if="discount.usageTracking?.length" class="flex justify-between">
                  <span class="text-sm text-green-700">Active Users:</span>
                  <span class="font-semibold text-green-900">{{ discount.usageTracking.length }}</span>
                </div>
              </div>
            </div>

            <div class="bg-blue-50 p-4 rounded-lg border border-blue-200">
              <h3 class="text-sm font-medium text-blue-800 mb-2">Validity Period</h3>
              <div class="space-y-2">
                <div class="flex items-center text-sm text-blue-700">
                  <CalendarIcon class="w-4 h-4 mr-2" />
                  <span>Start: {{ formatDateTime(discount.startAt) }}</span>
                </div>
                <div class="flex items-center text-sm text-blue-700">
                  <CalendarIcon class="w-4 h-4 mr-2" />
                  <span>End: {{ formatDateTime(discount.endAt) }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
        <div class="bg-white shadow-xl rounded-xl p-6 border border-gray-200">
          <h3 class="text-xl font-semibold text-gray-900 mb-6 flex items-center">
            <UsersIcon class="w-6 h-6 mr-2 text-green-600" />
            Customer Eligibility & Usage
          </h3>
          
          <div class="space-y-6">
            <div>
              <h4 class="text-lg font-medium text-gray-900 mb-3">Eligibility Settings</h4>
              <div class="space-y-4">
                <!-- Eligibility Status -->
                <div class="flex items-center justify-between p-3 rounded-lg" :class="{
                  'bg-green-50 border border-green-200': discount.eligibilityType === 'all',
                  'bg-blue-50 border border-blue-200': discount.eligibilityType === 'specific' || discount.eligibilityType === 'segment'
                }">
                  <div class="flex items-center">
                    <UsersIcon class="w-5 h-5 mr-2" :class="{
                      'text-green-600': discount.eligibilityType === 'all',
                      'text-blue-600': discount.eligibilityType === 'specific' || discount.eligibilityType === 'segment'
                    }" />
                    <div>
                      <span class="text-sm font-medium text-gray-900">
                        {{ discount.eligibilityType === 'all' ? 'Everyone is eligible' : 'Specific customers are eligible' }}
                      </span>
                      <p class="text-xs text-gray-600 mt-1">
                        {{ getEligibilityDescription() }}
                      </p>
                    </div>
                  </div>
                  <span :class="{
                    'bg-green-100 text-green-800': discount.eligibilityType === 'all',
                    'bg-blue-100 text-blue-800': discount.eligibilityType === 'specific' || discount.eligibilityType === 'segment'
                  }" class="inline-flex px-2.5 py-0.5 rounded-full text-xs font-medium">
                    {{ formatEligibility(discount.eligibilityType) }}
                  </span>
                </div>

                <!-- Eligible Customers List -->
                <div v-if="discount.eligibilityType !== 'all' && eligibleCustomers.length > 0" class="space-y-3">
                  <div class="flex items-center justify-between">
                    <h5 class="text-sm font-medium text-gray-900">Eligible Customers ({{ eligibleCustomers.length }})</h5>
                    <button @click="showEligibleCustomers = !showEligibleCustomers" 
                            class="text-xs text-blue-600 hover:text-blue-800 font-medium">
                      {{ showEligibleCustomers ? 'Hide' : 'Show' }} Details
                    </button>
                  </div>
                  
                  <div v-if="showEligibleCustomers" class="max-h-48 overflow-y-auto space-y-2">
                    <div v-for="customer in eligibleCustomers" :key="customer.id"
                         class="flex items-center justify-between p-2 bg-gray-50 rounded-md">
                      <div class="flex items-center">
                        <div class="w-2 h-2 rounded-full mr-2" :class="{
                          'bg-blue-500': customer.source === 'direct',
                          'bg-purple-500': customer.source === 'segment'
                        }"></div>
                        <div>
                          <span class="text-sm font-medium text-gray-900">
                            {{ customer.firstName }} {{ customer.lastName }}
                          </span>
                          <p class="text-xs text-gray-500">{{ customer.email }}</p>
                        </div>
                      </div>
                      <div class="flex items-center">
                        <span v-if="customer.source === 'segment'" 
                              class="text-xs bg-purple-100 text-purple-800 px-2 py-1 rounded-full">
                          {{ customer.segmentName }}
                        </span>
                        <span v-else 
                              class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded-full">
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
                  <p class="text-sm">No specific customers or segments added yet</p>
                </div>
              </div>
            </div>

            <div class="border-t border-gray-200 pt-4">
              <h4 class="text-lg font-medium text-gray-900 mb-3">Management Actions</h4>
              <div class="flex flex-wrap gap-2">
                <button @click="openMixedEligibilityModal"
                        class="inline-flex items-center px-3 py-1.5 bg-green-600 text-white text-sm font-medium rounded-md hover:bg-green-700 transition duration-150 ease-in-out">
                  <PlusIcon class="w-4 h-4 mr-1" />
                  Add Customers/Segments
                </button>
                <button @click="clearEligibility"
                        class="inline-flex items-center px-3 py-1.5 bg-orange-600 text-white text-sm font-medium rounded-md hover:bg-orange-700 transition duration-150 ease-in-out">
                  <UsersIcon class="w-4 h-4 mr-1" />
                  Allow Everyone
                </button>
                <button @click="refreshUsageStats"
                        class="inline-flex items-center px-3 py-1.5 bg-gray-600 text-white text-sm font-medium rounded-md hover:bg-gray-700 transition duration-150 ease-in-out">
                  <RefreshIcon class="w-4 h-4 mr-1" />
                  Refresh Stats
                </button>
              </div>
            </div>
          </div>
        </div>

        <div class="bg-white shadow-xl rounded-xl p-6 border border-gray-200">
          <h3 class="text-xl font-semibold text-gray-900 mb-6 flex items-center">
            <ChartBarIcon class="w-6 h-6 mr-2 text-green-600" />
            Usage Tracking
          </h3>
          
          <div class="space-y-4">
            <div class="flex justify-between items-center">
              <span class="text-sm text-gray-600">Total Usage:</span>
              <span class="font-semibold text-gray-900">
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
                <h4 class="text-sm font-medium text-gray-900">Recent Usage</h4>
                <button @click="showUsageDetails = !showUsageDetails"
                        class="text-sm text-green-600 hover:text-green-700 font-medium">
                  {{ showUsageDetails ? 'Hide' : 'Show' }} Details
                </button>
              </div>
              
              <div v-if="showUsageDetails" class="space-y-2 max-h-48 overflow-y-auto">
                <div v-for="(usage, idx) in discount.usageTracking.slice(0, 5)" :key="usage.customer_id"
                     class="flex justify-between items-center p-2 bg-gray-50 rounded-md text-sm">
                  <span class="text-gray-700">
                    <template v-if="recentUsageProfiles[idx]">
                      <span class="font-semibold">{{ recentUsageProfiles[idx].firstName }} {{ recentUsageProfiles[idx].lastName }}</span>
                      <span class="text-xs text-gray-500 ml-2">{{ recentUsageProfiles[idx].email }}</span>
                    </template>
                    <template v-else>
                      {{ usage.customer_id }}
                    </template>
                  </span>
                  <span class="text-gray-600">{{ usage.usage_count }}x</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div v-if="discount" class="bg-white shadow-xl rounded-xl p-6 border border-gray-200">
        <h3 class="text-xl font-semibold text-gray-900 mb-6 flex items-center">
          <CubeIcon class="w-6 h-6 mr-2 text-green-600" />
          Applied Items
        </h3>
        
        <div class="flex gap-2 mb-4">
          <button @click="showAddProductModal = true" class="px-3 py-1.5 bg-green-600 text-white rounded-md text-sm font-medium hover:bg-green-700">Add Product</button>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div v-if="discount.appliesToProducts?.length">
            <h4 class="text-lg font-medium text-gray-900 mb-4">Products</h4>
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
              <div v-for="prod in discount.appliesToProducts" :key="prod.id"
                   class="relative group border border-gray-200 rounded-lg overflow-hidden shadow-sm hover:shadow-md transition-shadow duration-200 ease-in-out bg-white cursor-pointer"
                   @click="goToProductDetail(prod.id)">
                <div class="aspect-w-16 aspect-h-9 overflow-hidden">
                  <img :src="prod.main_image || '/placeholder-image.png'" :alt="prod.name"
                       class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300 ease-in-out" />
                </div>
                <div class="p-3">
                  <h5 class="text-sm font-medium text-gray-800 group-hover:text-green-700 transition-colors">{{ prod.name }}</h5>
                  <p class="text-xs text-gray-600 mt-1 line-clamp-2">{{ prod.description || 'No description available.' }}</p>
                </div>
                <button @click.stop="removeProductFromDiscount(prod.id)" class="absolute top-2 right-2 bg-red-600 text-white rounded-full p-1 shadow hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-400 z-10">×</button>
              </div>
            </div>
          </div>
          <div v-else-if="discount.category === 'product'" class="p-4 bg-gray-50 border border-gray-200 rounded-lg text-gray-700 text-center">
            <p class="text-sm">This discount applies to <strong>all products</strong> in the shop.</p>
          </div>

          <div v-if="fetchedCollections.length">
            <h4 class="text-lg font-medium text-gray-900 mb-4">Collections</h4>
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
              <div v-for="coll in fetchedCollections" :key="coll.id"
                   class="group border border-gray-200 rounded-lg overflow-hidden shadow-sm hover:shadow-md transition-shadow duration-200 ease-in-out bg-white cursor-pointer"
                   @click="goToCollectionDetail(coll.id)">
                <div class="aspect-w-16 aspect-h-9 overflow-hidden">
                  <img :src="coll.image || '/placeholder-collection.png'" :alt="coll.title"
                       class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300 ease-in-out" />
                </div>
                <div class="p-3">
                  <h5 class="text-sm font-medium text-gray-800 group-hover:text-green-700 transition-colors">{{ coll.title }}</h5>
                  <p class="text-xs text-gray-600 mt-1 line-clamp-2">{{ coll.description || 'No description available.' }}</p>
                </div>
              </div>
            </div>
          </div>
          <div v-else class="p-4 bg-gray-50 border border-gray-200 rounded-lg text-gray-700 text-center">
            <p class="text-sm">This discount does not apply to any specific collections.</p>
          </div>
        </div>
      </div>
    </div>



    <div v-if="showMixedEligibilityModal" class="fixed inset-0 bg-gray-900 bg-opacity-75 flex items-center justify-center z-50 p-4">
      <div class="bg-white rounded-xl shadow-2xl w-full max-w-4xl p-6 relative">
        <h3 class="text-xl font-bold mb-4 text-gray-800">Add Eligible Customers</h3>
        
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- Customers Section -->
          <div>
            <h4 class="text-lg font-medium text-gray-900 mb-3">Customers</h4>
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">Search and Select Customers</label>
              <div class="relative customer-dropdown">
                <input
                  type="text"
                  v-model="customerSearchQuery"
                  @input="filterCustomers"
                  placeholder="Search customers by name or email..."
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
            </div>
            
            <div v-if="selectedCustomers.length > 0" class="mb-4">
              <h5 class="text-sm font-medium text-gray-700 mb-2">Selected Customers:</h5>
              <div class="flex flex-wrap gap-2 max-h-32 overflow-y-auto">
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

          <!-- Segments Section -->
          <div>
            <h4 class="text-lg font-medium text-gray-900 mb-3">Segments</h4>
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">Search and Select Segments</label>
              <div class="relative segment-dropdown">
                <input
                  type="text"
                  v-model="segmentSearchQuery"
                  @input="filterSegments"
                  placeholder="Search segments by name..."
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
            </div>
            
            <div v-if="selectedSegments.length > 0" class="mb-4">
              <h5 class="text-sm font-medium text-gray-700 mb-2">Selected Segments:</h5>
              <div class="flex flex-wrap gap-2 max-h-32 overflow-y-auto">
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
        
        <div class="flex justify-end space-x-3 pt-4">
          <button
            type="button"
            @click="closeMixedEligibilityModal"
            class="px-4 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-100 transition-colors duration-200"
          >
            Cancel
          </button>
          <button
            @click="addMixedEligibility"
            :disabled="selectedCustomers.length === 0 && selectedSegments.length === 0"
            class="px-4 py-2 bg-green-600 text-white rounded-lg shadow-md hover:bg-green-700 transition duration-300 ease-in-out disabled:opacity-50 disabled:cursor-not-allowed"
          >
            Add {{ selectedCustomers.length + selectedSegments.length }} Customer{{ (selectedCustomers.length + selectedSegments.length) !== 1 ? 's' : '' }}
          </button>
        </div>
        <button @click="closeMixedEligibilityModal" class="absolute top-4 right-4 text-gray-400 hover:text-gray-600 transition-colors duration-200">
          <XIcon class="h-6 w-6" />
        </button>
      </div>
    </div>

    <div v-if="showAddProductModal" class="fixed inset-0 bg-gray-900 bg-opacity-75 flex items-center justify-center z-50 p-4">
      <div class="bg-white rounded-xl shadow-2xl w-full max-w-2xl p-6 relative">
        <h3 class="text-xl font-bold mb-4 text-gray-800">Add Product to Discount</h3>
        <input v-model="productSearchQuery" placeholder="Search products..." class="w-full mb-4 border-gray-300 rounded-lg px-3 py-2.5" />
        <div class="max-h-64 overflow-y-auto">
          <div v-for="product in filteredAvailableProducts" :key="product.id" class="flex items-center justify-between p-2 border-b last:border-b-0">
            <span>{{ product.name }}</span>
            <button :disabled="isProductAdded(product.id)" @click="addProductToDiscount(product)" class="px-2 py-1 bg-green-600 text-white rounded disabled:opacity-50">Add</button>
          </div>
        </div>
        <button @click="showAddProductModal = false" class="absolute top-4 right-4 text-gray-400 hover:text-gray-600"><XIcon class="h-6 w-6" /></button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch, computed, onUnmounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { discountService } from '@/services/discount'
import { collectionService } from '@/services/collection'
import { customerService } from '@/services/customer'
import { customerSegmentService } from '@/services/customerSegment'
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
  ArrowLeftIcon,
  TagIcon,
  CodeIcon,
  UsersIcon,
  ChartBarIcon,
  CubeIcon,
  ExclamationIcon,
  ClockIcon
} from '@heroicons/vue/outline'
import { productService } from '@/services/product'

// State
const loading = ref(true)
const discount = ref(null)
const fetchedCollections = ref([])

// Router & Store
const router = useRouter()
const route = useRoute()
const shopStore = useShopStore()

const shopId = ref(shopStore.activeShop?.id);
watch(() => shopStore.activeShop, (newShop) => {
  if (newShop?.id) {
    shopId.value = newShop.id;
    loadDiscount()
    loadCustomers()
    loadSegments()
  }
});

const discountId = route.params.discountId

// Form state
const showForm = ref(false)
const formMode = ref('edit')
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
})

// Enhanced discount functionality state
const showUsageDetails = ref(false)
const showMixedEligibilityModal = ref(false)
const showAddProductModal = ref(false)
const productSearchQuery = ref('')
const allProducts = ref([])

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

// Form states
const addCustomersForm = ref({
  customerIds: ''
})
const addSegmentsForm = ref({
  segmentIds: ''
})

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
    console.warn("Shop ID or Discount ID not available yet. Skipping discount load.");
    loading.value = false;
    return;
  }

  loading.value = true;
  discount.value = null;
  fetchedCollections.value = [];

  try {
    const d = await discountService.fetchById(shopId.value, discountId);

    if (d && d.appliesToCollections && d.appliesToCollections.length > 0) {
      const collectionPromises = d.appliesToCollections.map(collectionId =>
        collectionService.fetchById(shopId.value, collectionId).catch(err => {
          console.warn(`Failed to load collection ${collectionId}:`, err);
          return null;
        })
      );
      fetchedCollections.value = (await Promise.all(collectionPromises)).filter(c => c !== null);
    }

    discount.value = d;
    
    // Load eligible customers if not "all" eligibility
    if (d && d.eligibilityType !== 'all') {
      await loadEligibleCustomers();
    } else {
      eligibleCustomers.value = [];
    }
  } catch (e) {
    console.error('Failed to load discount:', e);
  } finally {
    loading.value = false;
  }
}

// Load eligible customers for the discount
async function loadEligibleCustomers() {
  if (!shopId.value || !discount.value?.id) return
  
  try {
    const data = await discountService.getEligibleCustomers(shopId.value, discount.value.id);
    eligibleCustomers.value = data.eligible_customers || [];
  } catch (err) {
    console.error('Failed to load eligible customers:', err);
    eligibleCustomers.value = [];
  }
}

function goBack() {
  router.push({ name: 'Discounts' });
}

function goToProductDetail(productId) {
  console.log(`Navigating to product detail for ID: ${productId}`);
  alert(`Simulating navigation to product detail for ID: ${productId}`);
}

function goToCollectionDetail(collectionId) {
  console.log(`Navigating to collection detail for ID: ${collectionId}`);
  alert(`Simulating navigation to collection detail for ID: ${collectionId}`);
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
  if (!shopId.value) return;
  const ok = window.confirm(`Are you sure you want to delete discount "${disc.name}"? This action cannot be undone.`);
  if (!ok) return;

  loading.value = true;
  try {
    await discountService.delete(shopId.value, disc.id);
    router.push({ name: 'Discounts' });
  } catch (err) {
    console.error('Failed to delete discount:', err.response?.data || err.message);
    alert('Failed to delete discount: ' + (err.response?.data?.message || err.message));
  } finally {
    loading.value = false;
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
  if (customers.value.length === 0) {
    loadCustomers()
  }
  if (segments.value.length === 0) {
    loadSegments()
  }
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
  if (!discount.value || (selectedCustomers.value.length === 0 && selectedSegments.value.length === 0)) return
  try {
    // Deduplicate IDs before sending
    const customerIds = Array.from(new Set(selectedCustomers.value.map(c => c.id)))
    const segmentIds = Array.from(new Set(selectedSegments.value.map(s => s.id)))
    await discountService.addMixedEligibility(shopId.value, discount.value.id, customerIds, segmentIds)
    await loadDiscount()
    closeMixedEligibilityModal()
  } catch (err) {
    console.error('Failed to add mixed eligibility:', err)
    alert('Failed to add mixed eligibility: ' + (err.response?.data?.message || err.message))
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
function filterCustomers() {
  showCustomerDropdown.value = true
}

function filterSegments() {
  showSegmentDropdown.value = true
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
  }
  
  // Add click outside listener
  document.addEventListener('click', handleClickOutside)
})

// Clean up event listener
onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})

// Computed property for eligible customers count
const eligibleCustomersCount = computed(() => {
  return eligibleCustomers.value.length
})

// Load all products/variants for the shop
async function loadAllProducts() {
  if (!shopId.value) return
  allProducts.value = await productService.fetchAllByShop(shopId.value)
}

onMounted(() => {
  loadAllProducts()
})

const filteredAvailableProducts = computed(() => {
  const query = productSearchQuery.value.toLowerCase()
  return allProducts.value.filter(p =>
    (!query || p.name.toLowerCase().includes(query)) &&
    !isProductAdded(p.id)
  )
})
function isProductAdded(productId) {
  return (discount.value?.appliesToProducts || []).some(p => p.id === productId)
}
async function addProductToDiscount(product) {
  // Deduplicate and PATCH
  const ids = Array.from(new Set([...(discount.value.appliesToProducts?.map(p => p.id) || []), product.id]))
  await discountService.update(shopId.value, discount.value.id, { appliesToProducts: ids })
  await loadDiscount()
  showAddProductModal.value = false
}
async function removeProductFromDiscount(productId) {
  const ids = (discount.value.appliesToProducts || []).map(p => p.id).filter(id => id !== productId)
  await discountService.update(shopId.value, discount.value.id, { appliesToProducts: ids })
  await loadDiscount()
}

function goToEditPage(disc) {
  router.push({ name: 'EditDiscount', params: { discountId: disc.id } })
}
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>