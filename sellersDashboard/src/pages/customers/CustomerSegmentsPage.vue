<template>
  <div class="p-6 max-w-4xl mx-auto space-y-8 font-sans">
    <div class="flex items-center justify-between">
      <h2 class="text-3xl font-extrabold text-gray-900">Customer Segments</h2>
      <button @click="goToCustomers" class="text-green-600 hover:text-green-800 font-medium flex items-center">
        <ChevronLeftIcon class="w-5 h-5 mr-1" /> Back to Customers
      </button>
    </div>

    <div class="flex justify-end">
      <button @click="showCreateModal = true" class="px-4 py-2 bg-blue-600 text-white rounded-lg shadow hover:bg-blue-700 transition">
        + New Segment
      </button>
    </div>

    <div v-if="loading" class="flex justify-center py-12 text-gray-500">
      <RefreshIcon class="animate-spin w-8 h-8 mr-2" /> Loading segments...
    </div>
    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-lg shadow-sm">
      <p class="font-bold">Error:</p>
      <p>{{ error }}</p>
    </div>
    <div v-else>
      <table class="min-w-full bg-white rounded-xl shadow divide-y divide-gray-200">
        <thead class="bg-gray-100">
          <tr>
            <th class="py-3 px-6 text-left text-xs font-semibold text-gray-600 uppercase">Name</th>
            <th class="py-3 px-6 text-left text-xs font-semibold text-gray-600 uppercase">Description</th>
            <th class="py-3 px-6 text-left text-xs font-semibold text-gray-600 uppercase">Color</th>
            <th class="py-3 px-6 text-left text-xs font-semibold text-gray-600 uppercase">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="segment in segments" :key="segment.id" class="hover:bg-green-50">
            <td class="py-3 px-6 font-medium text-gray-900">{{ segment.name }}</td>
            <td class="py-3 px-6 text-gray-700">{{ segment.description || '—' }}</td>
            <td class="py-3 px-6">
              <span class="inline-block w-6 h-6 rounded-full border border-gray-200" :style="{ backgroundColor: segment.color }"></span>
              <span class="ml-2 text-xs text-gray-500">{{ segment.color }}</span>
            </td>
            <td class="py-3 px-6 flex space-x-2">
              <button @click="openEditModal(segment)" class="text-blue-600 hover:text-blue-800 font-medium">Edit</button>
              <button @click="openDeleteModal(segment)" class="text-red-600 hover:text-red-800 font-medium">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
      <div v-if="segments.length === 0" class="text-center text-gray-500 py-12">No segments found.</div>
    </div>

    <!-- Create/Edit Segment Modal -->
    <div v-if="showCreateModal || showEditModal" class="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl shadow-xl max-w-md w-full mx-4 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">{{ showEditModal ? 'Edit Segment' : 'Create New Segment' }}</h3>
        <form @submit.prevent="showEditModal ? updateSegment() : createSegment()">
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Name</label>
              <input v-model="segmentForm.name" type="text" required class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500" />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Description</label>
              <textarea v-model="segmentForm.description" rows="2" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500"></textarea>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Color</label>
              <input v-model="segmentForm.color" type="color" class="w-16 h-10 border border-gray-300 rounded-lg cursor-pointer" />
            </div>
          </div>
          <div class="flex justify-end space-x-3 mt-6">
            <button type="button" @click="closeModals" class="px-4 py-2 text-gray-700 bg-gray-200 rounded-lg hover:bg-gray-300">Cancel</button>
            <button type="submit" :disabled="saving" class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 disabled:opacity-50">
              {{ saving ? (showEditModal ? 'Saving...' : 'Creating...') : (showEditModal ? 'Save Changes' : 'Create Segment') }}
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div v-if="showDeleteModal" class="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl shadow-xl max-w-sm w-full mx-4 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">Delete Segment</h3>
        <p>Are you sure you want to delete the segment <span class="font-bold">{{ segmentToDelete?.name }}</span>? This cannot be undone.</p>
        <div class="flex justify-end space-x-3 mt-6">
          <button @click="closeModals" class="px-4 py-2 text-gray-700 bg-gray-200 rounded-lg hover:bg-gray-300">Cancel</button>
          <button @click="deleteSegment" :disabled="saving" class="px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 disabled:opacity-50">Delete</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { customerSegmentService } from '@/services/customerSegment'
import { ChevronLeftIcon, RefreshIcon } from '@heroicons/vue/outline'

const router = useRouter()
const shopStore = useShopStore()
const activeShop = computed(() => shopStore.activeShop)
const segments = ref([])
const loading = ref(false)
const error = ref(null)
const showCreateModal = ref(false)
const showEditModal = ref(false)
const showDeleteModal = ref(false)
const saving = ref(false)
const segmentForm = ref({ name: '', description: '', color: '#3B82F6' })
const segmentToEdit = ref(null)
const segmentToDelete = ref(null)

function goToCustomers() {
  router.push({ name: 'Customers' })
}

async function fetchSegments() {
  if (!activeShop.value?.id) {
    error.value = 'No shop selected. Please select a shop first.'
    segments.value = []
    return
  }
  loading.value = true
  try {
    segments.value = await customerSegmentService.fetchAll(activeShop.value.id)
  } catch (e) {
    error.value = 'Failed to load segments. Please try again.'
  } finally {
    loading.value = false
  }
}

function openEditModal(segment) {
  segmentToEdit.value = segment
  segmentForm.value = { ...segment }
  showEditModal.value = true
}

function openDeleteModal(segment) {
  segmentToDelete.value = segment
  showDeleteModal.value = true
}

function closeModals() {
  showCreateModal.value = false
  showEditModal.value = false
  showDeleteModal.value = false
  segmentToEdit.value = null
  segmentToDelete.value = null
  segmentForm.value = { name: '', description: '', color: '#3B82F6' }
}

async function createSegment() {
  if (!activeShop.value?.id) return
  saving.value = true
  try {
    await customerSegmentService.create(activeShop.value.id, segmentForm.value)
    closeModals()
    await fetchSegments()
  } catch (e) {
    error.value = 'Failed to create segment. Please try again.'
  } finally {
    saving.value = false
  }
}

async function updateSegment() {
  if (!activeShop.value?.id) return
  saving.value = true
  try {
    await customerSegmentService.update(activeShop.value.id, segmentToEdit.value.id, segmentForm.value)
    closeModals()
    await fetchSegments()
  } catch (e) {
    error.value = 'Failed to update segment. Please try again.'
  } finally {
    saving.value = false
  }
}

async function deleteSegment() {
  if (!activeShop.value?.id) return
  saving.value = true
  try {
    await customerSegmentService.delete(activeShop.value.id, segmentToDelete.value.id)
    closeModals()
    await fetchSegments()
  } catch (e) {
    error.value = 'Failed to delete segment. Please try again.'
  } finally {
    saving.value = false
  }
}

onMounted(fetchSegments)
</script> 