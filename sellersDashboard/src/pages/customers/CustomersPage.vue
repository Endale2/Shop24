<template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto font-sans space-y-8">
    <!-- Header & Controls -->
    <div class="flex flex-col md:flex-row md:items-center md:justify-between space-y-4 md:space-y-0">
      <h2 class="text-3xl sm:text-4xl font-extrabold text-gray-900 leading-tight">
        Customers Overview
      </h2>
      <div class="flex flex-col sm:flex-row w-full md:w-auto space-y-4 sm:space-y-0 sm:space-x-4">
        <input
          v-model="search"
          type="text"
          placeholder="Search by name or email..."
          class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 outline-none transition duration-150 ease-in-out shadow-sm"
        />
        <div class="inline-flex rounded-md shadow-sm" role="group">
          <button
            @click="currentView = 'cards'"
            :class="viewButtonClass('cards')"
            class="px-4 py-2 text-sm font-medium border border-gray-300 rounded-l-lg focus:z-10 focus:ring-2 focus:ring-blue-500 focus:outline-none transition duration-150 ease-in-out"
          >
            <GridIcon class="w-5 h-5 inline-block mr-1" />
            Cards
          </button>
          <button
            @click="currentView = 'list'"
            :class="viewButtonClass('list')"
            class="px-4 py-2 text-sm font-medium border border-gray-300 rounded-r-lg focus:z-10 focus:ring-2 focus:ring-blue-500 focus:outline-none transition duration-150 ease-in-out"
          >
            <ListIcon class="w-5 h-5 inline-block mr-1" />
            List
          </button>
        </div>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-16">
      <SpinnerIcon class="animate-spin h-10 w-10 text-blue-500 mb-3" />
      <p class="text-lg">Loading customers...</p>
    </div>

    <!-- Error -->
    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-lg shadow-sm" role="alert">
      <p class="font-bold">Oops! Something went wrong:</p>
      <p class="mt-2">{{ error }}</p>
      <p class="text-sm mt-2">Please try again or contact support.</p>
    </div>

    <!-- No Results -->
    <div v-else-if="filtered.length === 0" class="bg-blue-50 border border-blue-200 text-blue-700 px-6 py-8 rounded-lg text-center shadow-sm">
      <p class="text-lg font-medium">
        No customers found<span v-if="search"> matching "{{ search }}"</span>.
      </p>
      <p v-if="!search" class="mt-2">Your shop has no customers yet.</p>
    </div>

    <!-- List View -->
    <div v-else-if="currentView === 'list'" class="overflow-x-auto bg-white shadow-lg rounded-xl">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-100">
          <tr>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Name</th>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Email</th>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Phone</th>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Joined</th>
            <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">ID</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr
            v-for="(c, i) in filtered"
            :key="c.id"
            @click="goDetail(c)"
            class="cursor-pointer transition transform hover:scale-[1.005] hover:bg-blue-50"
            :class="{ 'bg-gray-50': i % 2 === 1 }"
          >
            <td class="py-3 px-6 flex items-center text-sm font-medium text-gray-800">
              <div class="h-8 w-8 rounded-full bg-gradient-to-br from-blue-400 to-purple-500 flex items-center justify-center text-white text-xs font-semibold mr-3">
                {{ initials(c) }}
              </div>
              {{ c.firstName }} {{ c.lastName }}
            </td>
            <td class="py-3 px-6 text-sm text-gray-700 truncate">{{ c.email }}</td>
            <td class="py-3 px-6 text-sm text-gray-700">{{ c.phone || 'N/A' }}</td>
            <td class="py-3 px-6 text-sm text-gray-600">{{ formatDate(c.createdAt) }}</td>
            <td class="py-3 px-6 text-sm text-gray-500 font-mono">{{ c.id.slice(0, 8) }}...</td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Cards View -->
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
      <div
        v-for="c in filtered"
        :key="c.id"
        @click="goDetail(c)"
        class="bg-white rounded-xl shadow-lg transform hover:shadow-xl hover:scale-105 transition duration-200 ease-in-out overflow-hidden"
      >
        <div class="p-6 text-center">
          <div class="h-16 w-16 mx-auto rounded-full bg-gradient-to-br from-blue-400 to-purple-500 flex items-center justify-center text-white text-2xl font-bold mb-4">
            {{ initials(c) }}
          </div>
          <h3 class="text-lg font-semibold text-gray-800 truncate">{{ c.firstName }} {{ c.lastName }}</h3>
          <p class="text-sm text-gray-600 truncate">{{ c.email }}</p>
          <p class="text-sm text-gray-600">{{ c.phone || 'N/A' }}</p>
        </div>
        <div class="px-6 py-3 bg-gray-50 border-t border-gray-100 text-xs text-gray-500 flex justify-between">
          <span>ID: {{ c.id.slice(0, 8) }}...</span>
          <span>Joined {{ formatDate(c.createdAt) }}</span>
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
  ViewListIcon as ListIcon,
  ViewGridIcon as GridIcon,
  RefreshIcon as SpinnerIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const shopStore = useShopStore()

const customers = ref([])
const loading = ref(false)
const error = ref(null)
const search = ref('')
const currentView = ref('list')

const activeShop = computed(() => shopStore.activeShop)
const filtered = computed(() => {
  const term = search.value.trim().toLowerCase()
  if (!term) return customers.value
  return customers.value.filter(c =>
    (`${c.firstName} ${c.lastName}`.toLowerCase().includes(term) ||
     c.email.toLowerCase().includes(term))
  )
})

function viewButtonClass(view) {
  return view === currentView.value
    ? 'bg-blue-600 text-white hover:bg-blue-700'
    : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
}

function formatDate(dt) {
  return dt
    ? new Date(dt).toLocaleDateString(undefined, { year: 'numeric', month: 'short', day: 'numeric' })
    : '—'
}

function initials(c) {
  const f = c.firstName?.[0]?.toUpperCase() || ''
  const l = c.lastName?.[0]?.toUpperCase() || ''
  return f + l || '?'
}

function goDetail(c) {
  router.push({ name: 'CustomerDetail', params: { customerId: c.id } })
}

onMounted(async () => {
  if (!activeShop.value) {
    error.value = 'No shop selected. Please select one to view customers.'
    return
  }
  loading.value = true
  try {
    customers.value = await customerService.fetchAll(activeShop.value.id)
  } catch (e) {
    console.error(e)
    error.value = 'Failed to load customers. Please try again later.'
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
@keyframes spin { from { transform: rotate(0deg) } to { transform: rotate(360deg) } }
.animate-spin { animation: spin 1s linear infinite }
</style>
