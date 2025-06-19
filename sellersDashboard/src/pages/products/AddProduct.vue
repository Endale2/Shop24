<!-- src/pages/products/AddProductPage.vue -->
<template>
  <div class="p-4 sm:p-6 max-w-6xl mx-auto space-y-8 font-sans">
    <button @click="goBack" class="inline-flex items-center text-gray-600 hover:text-green-700 transition duration-200 ease-in-out mb-6 group">
      <ChevronLeftIcon class="h-5 w-5 mr-1 text-gray-500 group-hover:text-green-600 transition-colors duration-200" />
      <span class="text-sm font-medium group-hover:text-green-700 transition-colors duration-200">Back to Products</span>
    </button>

    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-12">
      <SpinnerIcon class="animate-spin h-8 w-8 text-green-500 mb-3" />
      <p class="mt-3 text-lg">Preparing form...</p>
    </div>

    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-lg shadow-sm" role="alert">
      <strong class="font-bold">Error:</strong>
      <span class="ml-2">{{ error }}</span>
    </div>

    <div v-else class="bg-white rounded-xl shadow-lg p-6 md:p-8 space-y-8 animate-fade-in">
      <h2 class="text-3xl font-bold text-gray-800 mb-6">Add New Product</h2>

      <form @submit.prevent="submitForm" class="space-y-6">
        <!-- Name -->
        <div>
          <label for="name" class="block text-sm font-medium text-gray-700 mb-1">Product Name</label>
          <input
            id="name"
            v-model="form.name"
            type="text"
            required
            class="mt-1 block w-full border-gray-300 rounded-md shadow-sm py-2 px-3 focus:ring-green-500 focus:border-green-500 sm:text-sm"
          />
        </div>

        <!-- Description -->
        <div>
          <label for="description" class="block text-sm font-medium text-gray-700 mb-1">Description</label>
          <textarea
            id="description"
            v-model="form.description"
            rows="4"
            class="mt-1 block w-full border-gray-300 rounded-md shadow-sm py-2 px-3 focus:ring-green-500 focus:border-green-500 sm:text-sm"
          ></textarea>
        </div>

        <!-- Price -->
        <div>
          <label for="price" class="block text-sm font-medium text-gray-700 mb-1">Price</label>
          <input
            id="price"
            v-model.number="form.price"
            type="number"
            step="0.01"
            min="0"
            required
            class="mt-1 block w-40 border-gray-300 rounded-md shadow-sm py-2 px-3 focus:ring-green-500 focus:border-green-500 sm:text-sm"
          />
        </div>

        <!-- Category -->
        <div>
          <label for="category" class="block text-sm font-medium text-gray-700 mb-1">Category</label>
          <input
            id="category"
            v-model="form.category"
            type="text"
            class="mt-1 block w-full border-gray-300 rounded-md shadow-sm py-2 px-3 focus:ring-green-500 focus:border-green-500 sm:text-sm"
          />
        </div>

        <!-- Main Image -->
        <div>
          <label for="main_image" class="block text-sm font-medium text-gray-700 mb-1">Main Image URL</label>
          <input
            id="main_image"
            v-model="form.main_image"
            type="url"
            class="mt-1 block w-full border-gray-300 rounded-md shadow-sm py-2 px-3 focus:ring-green-500 focus:border-green-500 sm:text-sm"
            placeholder="https://…"
          />
        </div>

        <!-- Additional Images -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Additional Images (URLs)</label>
          <div v-for="(img, idx) in form.images" :key="idx" class="flex items-center space-x-2 mb-2">
            <input
              v-model="form.images[idx]"
              type="url"
              placeholder="https://…"
              class="flex-1 border-gray-300 rounded-md shadow-sm py-2 px-3 focus:ring-green-500 focus:border-green-500 sm:text-sm"
            />
            <button type="button" @click="removeImage(idx)" class="p-2 text-red-600 hover:text-red-800 rounded-full hover:bg-red-50">
              <MinusCircleIcon class="h-5 w-5" />
            </button>
          </div>
          <button
            type="button"
            @click="addImage"
            class="mt-2 inline-flex items-center text-sm font-medium text-blue-600 hover:text-blue-800"
          >
            <PlusCircleIcon class="h-5 w-5 mr-1" /> Add Image URL
          </button>
        </div>

        <!-- Variants -->
        <div>
          <h3 class="text-lg font-semibold text-gray-800 mb-4">Product Variants</h3>
          <div v-for="(v, vIdx) in form.variants" :key="vIdx" class="bg-gray-50 p-4 rounded-md mb-4 border border-gray-200">
            <div class="flex justify-between items-center mb-3">
              <span class="font-medium">Variant {{ vIdx + 1 }}</span>
              <button type="button" @click="removeVariant(vIdx)" class="text-red-600 hover:text-red-800">
                <XIcon class="h-5 w-5" />
              </button>
            </div>

            <!-- Options -->
            <div v-for="(value, key) in v.options" :key="key" class="flex items-center space-x-2 mb-2">
              <input
                type="text"
                v-model="variantKeyTemp[vIdx][key]"
                @blur="renameOptionKey(vIdx, key, variantKeyTemp[vIdx][key])"
                class="w-1/3 border-gray-300 rounded-md py-2 px-3 text-sm"
                placeholder="Option Name"
              />
              <input
                type="text"
                v-model="v.options[key]"
                class="flex-1 border-gray-300 rounded-md py-2 px-3 text-sm"
                placeholder="Option Value"
              />
              <button type="button" @click="removeVariantOption(vIdx, key)" class="p-1 text-red-600 hover:text-red-800">
                <MinusCircleIcon class="h-4 w-4" />
              </button>
            </div>
            <button type="button" @click="addVariantOption(vIdx)" class="inline-flex items-center text-sm text-blue-600 hover:text-blue-800 mb-3">
              <PlusCircleIcon class="h-4 w-4 mr-1" /> Add Option
            </button>

            <!-- Variant Price -->
            <div>
              <label class="block text-sm font-medium text-gray-700">Variant Price</label>
              <input
                type="number"
                v-model.number="v.price"
                step="0.01"
                min="0"
                class="mt-1 block w-40 border-gray-300 rounded-md shadow-sm py-2 px-3 sm:text-sm"
                required
              />
            </div>

            <!-- Variant Image -->
            <div class="mt-3">
              <label class="block text-sm font-medium text-gray-700">Variant Image URL</label>
              <input
                type="url"
                v-model="v.image"
                class="mt-1 block w-full border-gray-300 rounded-md shadow-sm py-2 px-3 sm:text-sm"
                placeholder="https://…"
              />
            </div>
          </div>

          <button
            type="button"
            @click="addVariant"
            class="inline-flex items-center px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-md hover:bg-green-700 transition"
          >
            <PlusIcon class="h-5 w-5 mr-1" /> Add Variant
          </button>
        </div>

        <!-- Actions -->
        <div class="flex justify-end space-x-4 mt-8">
          <button
            type="button"
            @click="goBack"
            class="px-6 py-3 border border-gray-300 rounded-md text-gray-700 bg-white hover:bg-gray-50 transition"
          >
            Cancel
          </button>
          <button
            type="submit"
            :disabled="saving"
            class="px-6 py-3 bg-green-600 text-white rounded-md hover:bg-green-700 transition disabled:opacity-50"
          >
            <SpinnerIcon v-if="saving" class="animate-spin h-5 w-5 mr-2 inline-block" />
            {{ saving ? 'Saving...' : 'Save Product' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { productService } from '@/services/product'
import {
  ChevronLeftIcon,
  RefreshIcon as SpinnerIcon,
  PlusCircleIcon,
  MinusCircleIcon,
  XIcon,
  PlusIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const shopStore = useShopStore()

const loading = ref(false)
const saving = ref(false)
const error = ref(null)

const form = reactive({
  name: '',
  description: '',
  price: 0,
  category: '',
  main_image: '',
  images: [],
  variants: []  // each: { options: { key: value }, price: number, image: '' }
})
// temp storage of editable keys for each variant
const variantKeyTemp = reactive({})

function goBack() {
  router.push({ name: 'Products' })
}

// Images
function addImage() { form.images.push('') }
function removeImage(i) { form.images.splice(i,1) }

// Variants
function addVariant() {
  form.variants.push({ options: {}, price: 0, image: '' })
  const idx = form.variants.length-1
  variantKeyTemp[idx] = {}
}
function removeVariant(i) {
  form.variants.splice(i,1)
  delete variantKeyTemp[i]
}

// Options
function addVariantOption(vIdx) {
  const key = `option_${Date.now()}`
  form.variants[vIdx].options[key] = ''
  variantKeyTemp[vIdx][key] = key
}
function removeVariantOption(vIdx, k) {
  delete form.variants[vIdx].options[k]
  delete variantKeyTemp[vIdx][k]
}
function renameOptionKey(vIdx, oldKey, newKey) {
  if (!newKey || oldKey===newKey) return
  const val = form.variants[vIdx].options[oldKey]
  delete form.variants[vIdx].options[oldKey]
  form.variants[vIdx].options[newKey] = val
  delete variantKeyTemp[vIdx][oldKey]
  variantKeyTemp[vIdx][newKey] = newKey
}

// Submit
async function submitForm() {
  saving.value = true
  error.value = null
  try {
    const payload = {
      name: form.name,
      description: form.description,
      price: form.price,
      category: form.category,
      images: form.images.filter(i=>i),
      main_image: form.main_image || undefined,
      variants: form.variants.map(v => ({
        name: Object.keys(v.options)[0] || '',
        options: Object.values(v.options),
        prices: [v.price],
      }))
    }
    await productService.create(shopStore.active.id, payload)
    router.push({ name: 'Products' })
  } catch (e) {
    console.error(e)
    error.value = 'Failed to create product. Please try again.'
  } finally {
    saving.value = false
  }
}
</script>

<style scoped>
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
.animate-fade-in { animation: fadeIn 0.4s ease-out forwards; }
</style>
