<template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto space-y-8 font-sans">
    <button
      @click="$router.back()"
      class="inline-flex items-center text-gray-600 hover:text-green-700 transition duration-200 ease-in-out mb-6 group rounded-full px-3 py-1.5 -ml-3"
    >
      <ArrowLeftIcon class="h-5 w-5 mr-1 text-gray-500 group-hover:text-green-600 transition-colors duration-200" />
      <span class="text-sm font-medium group-hover:text-green-700 transition-colors duration-200">Back to Collections</span>
    </button>
    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-20 bg-white rounded-2xl shadow-lg">
      <SpinnerIcon class="animate-spin h-10 w-10 text-green-500 mb-4" />
      <p class="mt-3 text-lg font-semibold text-gray-700">Preparing the form...</p>
    </div>
    <div v-else class="bg-white shadow-xl rounded-2xl p-6 sm:p-8 space-y-8 animate-fade-in">
      <div class="text-center mb-8">
        <h2 class="text-3xl sm:text-4xl font-extrabold text-gray-800 mb-2">
          Add New Collection
        </h2>
        <p class="text-gray-600">Organize your products into collections to help customers discover what they're looking for.</p>
      </div>
      <form @submit.prevent="handleSubmit" class="space-y-8">
        <!-- Title Section -->
        <div class="space-y-6 border-b border-gray-200 pb-8">
          <div class="flex items-center space-x-3">
            <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
              <span class="text-green-600 font-semibold text-sm">1</span>
            </div>
            <h3 class="text-xl font-bold text-gray-700">Basic Details</h3>
          </div>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="md:col-span-2">
              <label for="title" class="block text-sm font-medium text-gray-700 mb-2">
                Collection Title <span class="text-red-500">*</span>
              </label>
              <input
                id="title"
                v-model="form.title"
                type="text"
                required
                placeholder="e.g., Summer Collection, New Arrivals"
                class="w-full border border-gray-300 rounded-lg px-4 py-3 focus:ring-2 focus:ring-green-500 focus:border-green-500 transition duration-150 shadow-sm"
                :class="{ 'border-red-500': errors.title }"
              />
              <p v-if="errors.title" class="mt-1 text-sm text-red-600">{{ errors.title }}</p>
            </div>
            <div class="md:col-span-2">
              <label for="handle" class="block text-sm font-medium text-gray-700 mb-2">
                Handle <span class="text-red-500">*</span>
              </label>
              <div class="relative">
                <span class="absolute inset-y-0 left-0 pl-3 flex items-center text-gray-500 text-sm">
                  {{ activeShop?.slug }}/collections/
                </span>
                <input
                  id="handle"
                  v-model="form.handle"
                  type="text"
                  required
                  placeholder="summer-collection"
                  class="w-full pl-48 pr-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500 transition duration-150 shadow-sm"
                  :class="{ 'border-red-500': errors.handle }"
                />
              </div>
              <p v-if="errors.handle" class="mt-1 text-sm text-red-600">{{ errors.handle }}</p>
              <p class="mt-1 text-sm text-gray-500">This will be used in the URL for your collection page.</p>
            </div>
            <div class="md:col-span-2">
              <label for="description" class="block text-sm font-medium text-gray-700 mb-2">
                Description
              </label>
              <textarea
                id="description"
                v-model="form.description"
                rows="4"
                placeholder="Describe what this collection is about..."
                class="w-full border border-gray-300 rounded-lg px-4 py-3 focus:ring-2 focus:ring-green-500 focus:border-green-500 transition duration-150 shadow-sm resize-y"
              ></textarea>
            </div>
            <div class="md:col-span-2">
              <label for="image" class="block text-sm font-medium text-gray-700 mb-2">
                Collection Image URL
              </label>
              <input
                id="image"
                v-model="form.image"
                type="url"
                placeholder="https://example.com/collection-image.jpg"
                class="w-full border border-gray-300 rounded-lg px-4 py-3 focus:ring-2 focus:ring-green-500 focus:border-green-500 transition duration-150 shadow-sm"
              />
              <p class="mt-1 text-sm text-gray-500">Optional: Add an image to represent this collection.</p>
            </div>
          </div>
        </div>
        <!-- Preview Section -->
        <div v-if="form.title || form.handle || form.description || form.image" class="bg-gray-50 rounded-lg p-4">
          <h3 class="text-sm font-medium text-gray-700 mb-3">Preview</h3>
          <div class="bg-white rounded-lg p-4 border border-gray-200">
            <div class="flex items-start space-x-4">
              <div class="w-16 h-16 bg-gray-200 rounded-lg overflow-hidden flex-shrink-0">
                <img
                  v-if="form.image"
                  :src="form.image"
                  alt="Collection preview"
                  class="w-full h-full object-cover"
                  @error="handleImageError"
                />
                <div v-else class="w-full h-full flex items-center justify-center">
                  <PhotographIcon class="w-6 h-6 text-gray-400" />
                </div>
              </div>
              <div class="flex-grow min-w-0">
                <h4 class="text-lg font-semibold text-gray-900 truncate">
                  {{ form.title || 'Collection Title' }}
                </h4>
                <p class="text-sm text-gray-600 font-mono">
                  {{ activeShop?.slug }}/collections/{{ form.handle || 'handle' }}
                </p>
                <p v-if="form.description" class="text-sm text-gray-700 mt-1 line-clamp-2">
                  {{ form.description }}
                </p>
              </div>
            </div>
          </div>
        </div>
        <!-- Submit Button -->
        <div class="flex justify-end space-x-4 pt-6 border-t border-gray-200">
          <button
            type="button"
            @click="$router.back()"
            class="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 transition duration-150 ease-in-out"
          >
            Cancel
          </button>
          <button
            type="submit"
            :disabled="loading"
            class="px-6 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed transition duration-150 ease-in-out inline-flex items-center"
          >
            <SpinnerIcon v-if="loading" class="animate-spin h-4 w-4 mr-2" />
            <PlusIcon v-else class="h-4 w-4 mr-2" />
            {{ loading ? 'Creating...' : 'Create Collection' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { collectionService } from '@/services/collection'
import {
  ArrowLeftIcon,
  PlusIcon,
  RefreshIcon as SpinnerIcon,
  PhotographIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const shopStore = useShopStore()

// Reactive state
const loading = ref(false)
const errors = ref({})

// Form data
const form = ref({
  title: '',
  handle: '',
  description: '',
  image: ''
})

// Computed
const activeShop = computed(() => shopStore.activeShop)

onMounted(() => {
  if (!activeShop.value?.id) {
    router.replace({ name: 'ShopSelection' })
  }
})

/**
 * Handles form submission to create a new collection.
 */
async function handleSubmit() {
  if (!activeShop.value?.id) {
    errors.value = { general: 'No shop selected. Please select a shop first.' }
    return
  }

  // Reset errors
  errors.value = {}

  // Validate form
  if (!form.value.title.trim()) {
    errors.value.title = 'Title is required'
  }
  if (!form.value.handle.trim()) {
    errors.value.handle = 'Handle is required'
  } else if (!/^[a-z0-9-]+$/.test(form.value.handle)) {
    errors.value.handle = 'Handle can only contain lowercase letters, numbers, and hyphens'
  }

  if (Object.keys(errors.value).length > 0) {
    return
  }

  loading.value = true

  try {
    const newCollection = await collectionService.create(activeShop.value.id, {
      title: form.value.title.trim(),
      handle: form.value.handle.trim(),
      description: form.value.description.trim(),
      image: form.value.image.trim() || undefined
    })

    // Navigate to the new collection
    router.push({
      name: 'CollectionDetail',
      params: { collectionId: newCollection.id }
    })
  } catch (error) {
    console.error('Failed to create collection:', error)
    
    // Handle specific error cases
    if (error.response?.data?.error) {
      const errorMsg = error.response.data.error
      if (errorMsg.includes('handle already exists')) {
        errors.value.handle = 'This handle is already taken. Please choose a different one.'
      } else if (errorMsg.includes('invalid payload')) {
        errors.value.general = 'Please check your input and try again.'
      } else {
        errors.value.general = errorMsg
      }
    } else {
      errors.value.general = 'Failed to create collection. Please try again.'
    }
  } finally {
    loading.value = false
  }
}

/**
 * Handles image loading errors.
 */
function handleImageError(event) {
  event.target.style.display = 'none'
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