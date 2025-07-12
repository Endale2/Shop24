<template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto font-sans">
    <!-- Header Section -->
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center mb-8">
      <div>
        <h2 class="text-3xl sm:text-4xl font-extrabold text-gray-900 leading-tight">
          Customer Segments
        </h2>
        <p class="text-gray-600 mt-2">Organize your customers into meaningful groups</p>
      </div>

      <div class="flex flex-col sm:flex-row gap-3 mt-4 sm:mt-0">
        <button 
          @click="goToCustomers" 
          class="inline-flex items-center px-4 py-2.5 bg-gray-100 text-green-700 text-sm font-medium rounded-lg shadow-sm hover:bg-green-100 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
        >
          <ChevronLeftIcon class="w-5 h-5 mr-2 -ml-1" />
          Back to Customers
        </button>

        <button 
          @click="showCreateModal = true" 
          class="inline-flex items-center px-4 py-2.5 bg-green-600 text-white text-sm font-medium rounded-lg shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
        >
          <PlusIcon class="w-5 h-5 mr-2 -ml-1" />
          New Segment
        </button>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="space-y-4">
      <div class="overflow-x-auto bg-white shadow-lg rounded-xl animate-pulse">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Name</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Description</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Color</th>
              <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="n in 5" :key="n" class="bg-white">
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-3/4"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-1/2"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-1/3"></div></td>
              <td class="py-3 px-6"><div class="h-4 bg-gray-200 rounded w-1/3"></div></td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Error State -->
    <div
      v-else-if="error"
      class="bg-red-50 border border-red-200 text-red-700 px-6 py-4 rounded-lg shadow-sm mb-6"
      role="alert"
    >
      <div class="flex items-center">
        <ExclamationCircleIcon class="h-5 w-5 mr-2 flex-shrink-0" />
        <div>
          <strong class="font-semibold">Error:</strong>
          <span class="ml-1">{{ error }}</span>
        </div>
      </div>
      <p class="text-sm mt-2">Please ensure you have an active shop selected and try again.</p>
    </div>

    <!-- Content -->
    <div v-else>
      <div v-if="segments.length">
        <div class="overflow-x-auto bg-white shadow-lg rounded-xl">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Name</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Description</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Color</th>
                <th class="py-4 px-6 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Actions</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-200">
              <tr
                v-for="segment in segments"
                :key="segment.id"
                class="transition duration-150 ease-in-out transform hover:scale-[1.005] hover:bg-green-50"
              >
                <td class="py-3 px-6 font-medium text-gray-900">{{ segment.name }}</td>
                <td class="py-3 px-6 text-gray-700">{{ segment.description || '—' }}</td>
                <td class="py-3 px-6">
                  <div class="flex items-center">
                    <span class="inline-block w-6 h-6 rounded-full border border-gray-200 mr-2" :style="{ backgroundColor: segment.color }"></span>
                    <span class="text-xs text-gray-500">{{ segment.color }}</span>
                  </div>
                </td>
                <td class="py-3 px-6 text-sm text-gray-700">
                  <div class="flex space-x-2">
                    <button 
                      @click="openEditModal(segment)" 
                      class="text-blue-600 hover:text-blue-800 transition-colors"
                      title="Edit Segment"
                    >
                      <PencilIcon class="w-4 h-4" />
                    </button>
                    <button 
                      @click="openDeleteModal(segment)" 
                      class="text-red-600 hover:text-red-800 transition-colors"
                      title="Delete Segment"
                    >
                      <TrashIcon class="w-4 h-4" />
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Empty State -->
      <div v-else class="bg-green-50 border border-green-200 text-green-700 px-6 py-8 rounded-lg text-center mt-8 shadow-sm">
        <div class="flex flex-col items-center">
          <div class="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mb-4">
            <FolderIcon class="w-8 h-8 text-green-600" />
          </div>
          <p class="text-lg font-medium">No segments found</p>
          <p class="mt-2 text-sm">Create your first customer segment to start organizing your customers.</p>
          <button
            @click="showCreateModal = true"
            class="mt-4 inline-flex items-center px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out"
          >
            <PlusIcon class="w-4 h-4 mr-2" />
            Create Your First Segment
          </button>
        </div>
      </div>
    </div>

    <!-- Create/Edit Segment Modal -->
    <div v-if="showCreateModal || showEditModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl shadow-xl max-w-md w-full mx-4">
        <div class="p-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">{{ showEditModal ? 'Edit Segment' : 'Create New Segment' }}</h3>
          
          <form @submit.prevent="showEditModal ? updateSegment() : createSegment()">
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">Segment Name</label>
                <input 
                  v-model="segmentForm.name" 
                  type="text" 
                  required 
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
                  placeholder="e.g., VIP Customers"
                />
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">Description</label>
                <textarea 
                  v-model="segmentForm.description" 
                  rows="3" 
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
                  placeholder="Optional description..."
                ></textarea>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">Color</label>
                <input 
                  v-model="segmentForm.color" 
                  type="color" 
                  class="w-full h-10 border border-gray-300 rounded-lg cursor-pointer"
                />
              </div>
            </div>
            
            <div class="flex justify-end space-x-3 mt-6">
              <button 
                type="button" 
                @click="closeModals" 
                class="px-4 py-2 text-gray-700 bg-gray-200 rounded-lg hover:bg-gray-300 transition-colors"
              >
                Cancel
              </button>
              <button 
                type="submit" 
                :disabled="saving" 
                class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 disabled:opacity-50 transition-colors"
              >
                {{ saving ? (showEditModal ? 'Saving...' : 'Creating...') : (showEditModal ? 'Save Changes' : 'Create Segment') }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div v-if="showDeleteModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl shadow-xl max-w-sm w-full mx-4">
        <div class="p-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">Delete Segment</h3>
          <p class="text-gray-700 mb-6">
            Are you sure you want to delete the segment 
            <span class="font-bold text-gray-900">{{ segmentToDelete?.name }}</span>? 
            This action cannot be undone.
          </p>
          
          <div class="flex justify-end space-x-3">
            <button 
              @click="closeModals" 
              class="px-4 py-2 text-gray-700 bg-gray-200 rounded-lg hover:bg-gray-300 transition-colors"
            >
              Cancel
            </button>
            <button 
              @click="deleteSegment" 
              :disabled="saving" 
              class="px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 disabled:opacity-50 transition-colors"
            >
              {{ saving ? 'Deleting...' : 'Delete' }}
            </button>
          </div>
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
import {
  ChevronLeftIcon,
  RefreshIcon,
  ExclamationCircleIcon,
  PlusIcon,
  FolderIcon,
  PencilIcon,
  TrashIcon
} from '@heroicons/vue/outline'

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