<template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto space-y-8 font-sans">
    <div class="flex flex-col md:flex-row md:items-center md:justify-between space-y-4 md:space-y-0">
      <h2 class="text-3xl sm:text-4xl font-extrabold text-gray-900 leading-tight">
        Customers
      </h2>

      <div class="flex flex-col sm:flex-row sm:items-center space-y-4 sm:space-y-0 sm:space-x-4 w-full md:w-auto">
        <div class="relative w-full sm:w-64">
          <input
            v-model="search"
            type="text"
            placeholder="Search by name or email..."
            class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent transition duration-150 ease-in-out shadow-sm"
          />
          <SearchIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
        </div>

        <button
          @click="goToAddCustomer"
          class="inline-flex items-center px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg shadow-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
        >
          <PlusIcon class="w-5 h-5 mr-2 -ml-1" />
          Add New Customer
        </button>

        <div class="inline-flex rounded-full shadow-md overflow-hidden" role="group">
          <button
            @click="currentView = 'cards'"
            :class="buttonClass('cards')"
            class="px-5 py-2.5 text-sm font-medium border border-gray-300 focus:z-10 focus:ring-2 focus:ring-green-500 focus:outline-none transition-colors duration-200 ease-in-out"
          >
            <ViewGridIcon class="w-5 h-5 inline-block mr-1" />
            Cards
          </button>
          <button
            @click="currentView = 'list'"
            :class="buttonClass('list')"
            class="px-5 py-2.5 text-sm font-medium border border-gray-300 focus:z-10 focus:ring-2 focus:ring-green-500 focus:outline-none transition-colors duration-200 ease-in-out"
          >
            <ViewListIcon class="w-5 h-5 inline-block mr-1" />
            List
          </button>
        </div>
      </div>
    </div>

    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-16">
      <RefreshIcon class="animate-spin h-10 w-10 text-green-500 mb-3" />
      <p class="text-lg">Loading customers...</p>
    </div>

    <div
      v-else-if="error"
      class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-lg shadow-sm"
      role="alert"
    >
      <p class="font-bold">Oops! Something went wrong:</p>
      <p class="mt-2">{{ error }}</p>
      <p class="text-sm mt-2">Please try again or contact support if the issue persists.</p>
    </div>

    <div
      v-else-if="filteredCustomers.length === 0"
      class="bg-green-50 border border-green-200 text-green-700 px-6 py-8 rounded-lg text-center mt-8 shadow-sm"
    >
      <p class="text-lg font-medium">
        No customers found<span v-if="search"> matching "{{ search }}"</span>.
      </p>
      <p class="mt-2" v-if="!search">
        It looks like there are no customers for your active shop yet.
      </p>
    </div>

    <div v-else-if="currentView === 'list'" class="overflow-x-auto bg-white shadow-lg rounded-xl">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-100">
          <tr>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Name</th>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Email</th>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Phone</th>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Joined</th>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Username</th>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Location</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr
            v-for="(cust, i) in filteredCustomers"
            :key="cust.id"
            @click="goToCustomerDetail(cust.id)"
            class="cursor-pointer transition duration-150 ease-in-out transform hover:scale-[1.005] hover:bg-green-50"
            :class="{ 'bg-gray-50': i % 2 === 1 }"
          >
            <td class="py-3 px-6 text-sm text-gray-800 font-medium">
              <div class="flex items-center">
                <div
                  class="h-8 w-8 bg-green-100 rounded-full flex items-center justify-center text-green-700 text-xs font-semibold mr-3"
                >
                  {{ getInitials(cust.firstName, cust.lastName) }}
                </div>
                <span>{{ cust.firstName }} {{ cust.lastName }}</span>
              </div>
            </td>
            <td class="py-3 px-6 text-sm text-gray-700 truncate">{{ cust.email }}</td>
            <td class="py-3 px-6 text-sm text-gray-700">{{ cust.phone || 'N/A' }}</td>
            <td class="py-3 px-6 text-sm text-gray-600">{{ formatDate(cust.createdAt) }}</td>
            <td class="py-3 px-6 text-sm text-gray-700 font-mono">@{{ cust.username }}</td>
            <td class="py-3 px-6 text-sm text-gray-600">{{ cust.city }}, {{ cust.state }}</td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
      <div
        v-for="cust in filteredCustomers"
        :key="cust.id"
        @click="goToCustomerDetail(cust.id)"
        class="bg-white rounded-xl shadow-lg overflow-hidden cursor-pointer transform hover:scale-105 hover:shadow-xl transition duration-200 ease-in-out flex flex-col group"
      >
        <div class="flex flex-col items-center p-6 text-center space-y-4">
          <div
            class="h-20 w-20 bg-gradient-to-br from-green-400 to-teal-500 rounded-full flex items-center justify-center text-white text-2xl font-bold shadow-md"
          >
            {{ getInitials(cust.firstName, cust.lastName) }}
          </div>
          <div class="space-y-1">
            <h3 class="text-xl font-semibold text-gray-800">
              {{ cust.firstName }} {{ cust.lastName }}
            </h3>
            <p class="text-sm text-gray-600 truncate">{{ cust.email }}</p>
            <p class="text-sm text-gray-600">{{ cust.phone || 'N/A' }}</p>
          </div>
        </div>
        <div class="px-6 py-3 bg-gray-50 border-t border-gray-100 flex items-center justify-between text-xs text-gray-500">
          <span class="text-gray-700 font-mono">@{{ cust.username }}</span>
          <span>{{ formatDate(cust.createdAt) }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { customerService } from '@/services/customer'
import {
  ViewListIcon,
  ViewGridIcon,
  RefreshIcon,
  SearchIcon,
  PlusIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const shopStore = useShopStore()

// Reactive state
const customers = ref([])
const loading = ref(false)
const error = ref(null)
const search = ref('')
const currentView = ref('list')

// Computed
const activeShop = computed(() => shopStore.activeShop)
const filteredCustomers = computed(() => {
  const term = search.value.trim().toLowerCase()
  if (!term) return customers.value
  return customers.value.filter(c => {
    const fullName = `${c.firstName} ${c.lastName}`.toLowerCase()
    return fullName.includes(term) || (c.email?.toLowerCase().includes(term))
  })
})

// Fetch customers
onMounted(async () => {
  console.log('🔎 customers  activeShop is:', activeShop.value)
  if (!activeShop.value?.id) {
    router.replace({ name: 'ShopSelection' })
    return
  }
  loading.value = true
  try {
    const result = await customerService.fetchAll(activeShop.value.id)
    customers.value = Array.isArray(result) ? result : []
  } catch (e) {
    console.error(e)
    error.value = 'Failed to load customers. Please try again later.'
  } finally {
    loading.value = false
  }
})

// Helpers
function buttonClass(view) {
  return view === currentView.value
    ? 'bg-green-600 text-white hover:bg-green-700 shadow-inner'
    : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
}

function formatDate(dt) {
  return dt
    ? new Date(dt).toLocaleDateString(undefined, {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
      })
    : '—'
}

function getInitials(firstName, lastName) {
  const fi = firstName?.[0]?.toUpperCase() || ''
  const li = lastName?.[0]?.toUpperCase() || ''
  return fi + li || '?'
}

function goToAddCustomer() {
  router.push({ name: 'AddCustomer' }) // Make sure you have a route named 'AddCustomer'
}

function goToCustomerDetail(customerId) {
  router.push({
    name: 'CustomerDetail',
    params: { customerId },
  })
}
</script>

<style scoped>
@keyframes spin {
  from { transform: rotate(0deg); }
  to   { transform: rotate(360deg); }
}
.animate-spin {
  animation: spin 1s linear infinite;
}
</style>