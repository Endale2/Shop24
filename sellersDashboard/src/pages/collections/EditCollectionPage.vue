<template>
  <div class="p-4 sm:p-6 max-w-3xl mx-auto space-y-8 font-sans">
    <button
      @click="goBack"
      class="inline-flex items-center text-gray-600 hover:text-green-600 transition-all duration-200 ease-in-out group mb-6 rounded-full px-3 py-1.5 -ml-3 bg-gray-100/50 hover:bg-green-50"
    >
      <svg
        class="w-5 h-5 mr-2 transition-all duration-200 group-hover:scale-110 group-hover:-translate-x-0.5"
        fill="none"
        stroke="currentColor"
        stroke-width="2.5"
        viewBox="0 0 24 24"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <path d="M15 18l-6-6 6-6" />
      </svg>
      <span class="text-sm font-medium transition-colors duration-200">Back to Collection</span>
    </button>
    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-20 bg-white rounded-2xl shadow-lg">
      <SpinnerIcon class="animate-spin h-10 w-10 text-green-500 mb-4" />
      <p class="mt-3 text-lg font-semibold text-gray-700">Loading collection...</p>
    </div>
    <div v-else class="bg-white shadow-xl rounded-2xl p-6 sm:p-8 space-y-8 animate-fade-in">
      <div class="text-center mb-8">
        <h2 class="text-3xl sm:text-4xl font-extrabold text-gray-800 mb-2">
          Edit Collection
        </h2>
        <p class="text-gray-600">Update your collection details below.</p>
      </div>
      <form @submit.prevent="handleEditSubmit" class="space-y-8">
        <div class="space-y-6 border-b border-gray-200 pb-8">
          <div class="flex items-center space-x-3">
            <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
              <span class="text-green-600 font-semibold text-sm">1</span>
            </div>
            <h3 class="text-xl font-bold text-gray-700">Collection Details</h3>
          </div>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="md:col-span-2">
              <label for="title" class="block text-sm font-medium text-gray-700 mb-2">
                Collection Title <span class="text-red-500">*</span>
              </label>
              <input
                id="title"
                v-model="editForm.title"
                type="text"
                required
                class="w-full border border-gray-300 rounded-lg px-4 py-3 focus:ring-2 focus:ring-green-500 focus:border-green-500 transition duration-150 shadow-sm"
                :class="{ 'border-red-500': editErrors.title }"
              />
              <p v-if="editErrors.title" class="mt-1 text-sm text-red-600">{{ editErrors.title }}</p>
            </div>
            <div class="md:col-span-2">
              <label for="handle" class="block text-sm font-medium text-gray-700 mb-2">
                Handle <span class="text-red-500">*</span>
              </label>
              <input
                id="handle"
                v-model="editForm.handle"
                type="text"
                required
                class="w-full border border-gray-300 rounded-lg px-4 py-3 focus:ring-2 focus:ring-green-500 focus:border-green-500 transition duration-150 shadow-sm"
                :class="{ 'border-red-500': editErrors.handle }"
              />
              <p v-if="editErrors.handle" class="mt-1 text-sm text-red-600">{{ editErrors.handle }}</p>
            </div>
            <div class="md:col-span-2">
              <label for="description" class="block text-sm font-medium text-gray-700 mb-2">
                Description
              </label>
              <textarea
                id="description"
                v-model="editForm.description"
                rows="4"
                class="w-full border border-gray-300 rounded-lg px-4 py-3 focus:ring-2 focus:ring-green-500 focus:border-green-500 transition duration-150 shadow-sm resize-y"
              ></textarea>
            </div>
            <div class="md:col-span-2">
              <label for="image" class="block text-sm font-medium text-gray-700 mb-2">
                Collection Image URL
              </label>
              <input
                id="image"
                v-model="editForm.image"
                type="url"
                class="w-full border border-gray-300 rounded-lg px-4 py-3 focus:ring-2 focus:ring-green-500 focus:border-green-500 transition duration-150 shadow-sm"
              />
            </div>
          </div>
        </div>
        <div class="flex justify-end space-x-4 pt-6 border-t border-gray-200">
          <button
            type="button"
            @click="goBack"
            class="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 transition duration-150 ease-in-out"
          >
            Cancel
          </button>
          <button
            type="submit"
            :disabled="editLoading"
            class="px-6 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed transition duration-150 ease-in-out inline-flex items-center"
          >
            <SpinnerIcon v-if="editLoading" class="animate-spin h-4 w-4 mr-2" />
            <CheckIcon v-else class="h-4 w-4 mr-2" />
            {{ editLoading ? 'Saving...' : 'Save Changes' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { collectionService } from '@/services/collection'
import {
  RefreshIcon as SpinnerIcon,
  CheckIcon
} from '@heroicons/vue/outline'

const route = useRoute()
const router = useRouter()
const shopStore = useShopStore()

const loading = ref(true)
const editLoading = ref(false)
const editForm = ref({ title: '', handle: '', description: '', image: '' })
const editErrors = ref({})

const activeShop = computed(() => shopStore.activeShop)

onMounted(async () => {
  if (!activeShop.value?.id) {
    router.replace({ name: 'ShopSelection' })
    return
  }
  const collectionId = route.params.collectionId
  if (!collectionId) {
    router.replace({ name: 'Collections' })
    return
  }
  await loadCollection(collectionId)
})

async function loadCollection(collectionId) {
  loading.value = true
  try {
    const data = await collectionService.fetchById(activeShop.value.id, collectionId)
    editForm.value = {
      title: data.title,
      handle: data.handle,
      description: data.description || '',
      image: data.image || ''
    }
  } catch (e) {
    router.replace({ name: 'Collections' })
  } finally {
    loading.value = false
  }
}

async function handleEditSubmit() {
  editErrors.value = {}
  if (!editForm.value.title.trim()) {
    editErrors.value.title = 'Title is required'
  }
  if (!editForm.value.handle.trim()) {
    editErrors.value.handle = 'Handle is required'
  } else if (!/^[a-z0-9-]+$/.test(editForm.value.handle)) {
    editErrors.value.handle = 'Handle can only contain lowercase letters, numbers, and hyphens'
  }
  if (Object.keys(editErrors.value).length > 0) return
  editLoading.value = true
  try {
    await collectionService.update(activeShop.value.id, route.params.collectionId, {
      title: editForm.value.title.trim(),
      handle: editForm.value.handle.trim(),
      description: editForm.value.description.trim(),
      image: editForm.value.image.trim() || undefined
    })
    router.push({ name: 'CollectionDetail', params: { collectionId: route.params.collectionId } })
  } catch (error) {
    if (error.response?.data?.error) {
      const errorMsg = error.response.data.error
      if (errorMsg.includes('handle already exists')) {
        editErrors.value.handle = 'This handle is already taken. Please choose a different one.'
      } else {
        editErrors.value.general = errorMsg
      }
    } else {
      editErrors.value.general = 'Failed to update collection. Please try again.'
    }
  } finally {
    editLoading.value = false
  }
}

function goBack() {
  router.push({ name: 'CollectionDetail', params: { collectionId: route.params.collectionId } })
}
</script> 