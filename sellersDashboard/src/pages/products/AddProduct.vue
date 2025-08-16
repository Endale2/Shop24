<template>
  <div class="p-4 sm:p-6 max-w-6xl mx-auto space-y-8 font-sans">
    <BackButton
      :to="{ name: 'Products' }"
      text="Back to Products"
      variant="rounded"
      class="mb-6 -ml-3"
    />

    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-20 bg-white rounded-2xl shadow-lg">
      <SpinnerIcon class="animate-spin h-10 w-10 text-green-500 mb-4" />
      <p class="mt-3 text-lg font-semibold text-gray-700">Preparing the form...</p>
    </div>

    <div
      v-else-if="error"
      class="bg-red-50 border border-red-300 text-red-700 px-6 py-4 rounded-xl shadow-md flex items-center space-x-3"
      role="alert"
    >
      <ExclamationCircleIcon class="h-6 w-6 flex-shrink-0" />
      <div>
        <strong class="font-semibold">Error:</strong> <span class="ml-1">{{ error }}</span>
        <p class="text-sm mt-1">Please check your input and try again.</p>
      </div>
    </div>

    <div v-else class="bg-white shadow-xl rounded-2xl p-6 sm:p-8 space-y-8 animate-fade-in">
      <div class="text-center mb-8">
        <h2 class="text-3xl sm:text-4xl font-extrabold text-gray-800 mb-2">
          Add New Product
        </h2>
        <p class="text-gray-600">Create a new product for your shop</p>
      </div>

      <form @submit.prevent="submitForm" class="space-y-8">
        <!-- Basic Details Section -->
        <div class="space-y-6 border-b border-gray-200 pb-8">
          <div class="flex items-center space-x-3">
            <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
              <span class="text-green-600 font-semibold text-sm">1</span>
            </div>
            <h3 class="text-xl font-bold text-gray-700">Basic Details</h3>
          </div>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="md:col-span-2">
              <label for="product-name" class="block text-sm font-medium text-gray-700 mb-2">
                Product Name <span class="text-red-500">*</span>
              </label>
              <input
                id="product-name"
                v-model="form.name"
                type="text"
                required
                class="w-full border border-gray-300 rounded-lg px-4 py-3 focus:ring-2 focus:ring-green-500 focus:border-green-500 transition duration-150 shadow-sm"
                placeholder="e.g., Cozy Comfort Blanket"
              />
            </div>

            <div class="md:col-span-2">
              <label for="product-description" class="block text-sm font-medium text-gray-700 mb-2">
                Description
              </label>
              <textarea
                id="product-description"
                v-model="form.description"
                rows="4"
                class="w-full border border-gray-300 rounded-lg px-4 py-3 focus:ring-2 focus:ring-green-500 focus:border-green-500 transition duration-150 shadow-sm resize-y"
                placeholder="Describe your product in detail..."
              ></textarea>
            </div>

            <div class="md:col-span-2">
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Collections <span class="text-red-500">*</span>
              </label>
              <div class="space-y-3">
                <div v-for="coll in collections" :key="coll.id" class="flex items-center">
                  <input
                    :id="'collection-' + coll.id"
                    type="checkbox"
                    :value="coll.id"
                    v-model="form.collection_ids"
                    class="h-4 w-4 text-green-600 focus:ring-green-500 border-gray-300 rounded"
                  />
                  <label :for="'collection-' + coll.id" class="ml-3 text-sm font-medium text-gray-700">
                    {{ coll.title }}
                  </label>
                </div>
                <div v-if="collections.length === 0" class="text-sm text-gray-500 italic">
                  No collections available. Please create a collection first.
                </div>
              </div>
            </div>

            <div v-if="!hasVariants">
              <label for="price" class="block text-sm font-medium text-gray-700 mb-2">
                Price <span class="text-red-500">*</span>
              </label>
              <div class="relative">
                <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <span class="text-gray-500 sm:text-sm">$</span>
                </div>
                <input
                  id="price"
                  type="number"
                  v-model.number="form.price"
                  step="0.01"
                  min="0"
                  required
                  class="w-full pl-7 pr-3 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500 transition duration-150 shadow-sm"
                  placeholder="0.00"
                />
              </div>
            </div>

            <div v-if="!hasVariants">
              <label for="stock" class="block text-sm font-medium text-gray-700 mb-2">
                Stock Quantity
              </label>
              <input
                id="stock"
                type="number"
                v-model.number="form.stock"
                min="0"
                class="w-full border border-gray-300 rounded-lg px-4 py-3 focus:ring-2 focus:ring-green-500 focus:border-green-500 transition duration-150 shadow-sm"
                placeholder="0"
              />
            </div>

            <div v-if="hasVariants" class="md:col-span-2">
              <div class="bg-blue-50 border border-blue-200 rounded-lg p-4">
                <div class="flex items-center">
                  <InformationCircleIcon class="h-5 w-5 text-blue-600 mr-2" />
                  <span class="text-sm font-medium text-blue-800">Variant-based pricing</span>
                </div>
                <p class="text-sm text-blue-700 mt-1">
                  Price and stock will be calculated from your variants below.
                </p>
              </div>
            </div>
          </div>
        </div>

        <!-- Images Section -->
        <div class="space-y-6 border-b border-gray-200 pb-8">
          <div class="flex items-center space-x-3">
            <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
              <span class="text-blue-600 font-semibold text-sm">2</span>
            </div>
            <h3 class="text-xl font-bold text-gray-700">Product Images</h3>
          </div>
          
          <div class="space-y-4">
            <div>
              <label for="main-image" class="block text-sm font-medium text-gray-700 mb-2">
                Main Image URL
              </label>
              <input
                id="main-image"
                v-model="form.main_image"
                type="url"
                placeholder="https://example.com/main-product.jpg"
                class="w-full border border-gray-300 rounded-lg px-4 py-3 focus:ring-2 focus:ring-green-500 focus:border-green-500 transition duration-150 shadow-sm"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-3">
                Additional Images
              </label>
              <div
                v-for="(img, idx) in form.images"
                :key="idx"
                class="flex items-center space-x-3 mb-3"
              >
                <input
                  v-model="form.images[idx]"
                  type="url"
                  placeholder="https://example.com/extra-image.jpg"
                  class="flex-1 border border-gray-300 rounded-lg px-4 py-3 focus:ring-2 focus:ring-green-500 focus:border-green-500 transition duration-150 shadow-sm"
                />
                <button
                  type="button"
                  @click="removeImage(idx)"
                  class="text-red-500 hover:text-red-700 transition duration-150 p-2 rounded-full hover:bg-red-50"
                >
                  <MinusCircleIcon class="h-6 w-6" />
                </button>
              </div>
              <button
                type="button"
                @click="addImage"
                class="inline-flex items-center px-4 py-2 bg-blue-50 text-blue-600 rounded-lg font-medium hover:bg-blue-100 transition-colors duration-200 shadow-sm"
              >
                <PlusCircleIcon class="h-5 w-5 mr-2" /> Add Another Image
              </button>
            </div>
          </div>
        </div>

        <!-- Add after Images Section, before Variants Section -->
        <div class="space-y-6 border-b border-gray-200 pb-8">
          <div class="flex items-center space-x-3">
            <div class="w-8 h-8 bg-yellow-100 rounded-full flex items-center justify-center">
              <span class="text-yellow-600 font-semibold text-sm">SEO</span>
            </div>
            <h3 class="text-xl font-bold text-gray-700">Advanced Fields (SEO)</h3>
            <button type="button" @click="showAdvanced = !showAdvanced" class="ml-4 px-3 py-1.5 rounded bg-gray-100 text-gray-700 text-sm font-medium hover:bg-yellow-50">
              {{ showAdvanced ? 'Hide' : 'Show' }}
            </button>
          </div>
          <div v-if="showAdvanced" class="space-y-4 mt-4">
            <div>
              <label for="meta-title" class="block text-sm font-medium text-gray-700 mb-2">Meta Title</label>
              <input id="meta-title" v-model="form.meta_title" type="text" class="w-full border border-gray-300 rounded-lg px-4 py-3 focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500 transition duration-150 shadow-sm" placeholder="SEO meta title (optional)" />
            </div>
            <div>
              <label for="meta-description" class="block text-sm font-medium text-gray-700 mb-2">Meta Description</label>
              <textarea id="meta-description" v-model="form.meta_description" rows="2" class="w-full border border-gray-300 rounded-lg px-4 py-3 focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500 transition duration-150 shadow-sm resize-y" placeholder="SEO meta description (optional)"></textarea>
            </div>
          </div>
        </div>

        <!-- Variants Section -->
        <div class="space-y-6">
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-3">
              <div class="w-8 h-8 bg-purple-100 rounded-full flex items-center justify-center">
                <span class="text-purple-600 font-semibold text-sm">3</span>
              </div>
              <h3 class="text-xl font-bold text-gray-700">Product Variants</h3>
            </div>
            <button
              v-if="!showVariants"
              type="button"
              @click="showVariants = true"
              class="inline-flex items-center px-4 py-2 bg-purple-600 text-white rounded-lg font-medium hover:bg-purple-700 transition-colors duration-200 shadow-sm"
            >
              <PlusIcon class="h-5 w-5 mr-2" /> Add Variants
            </button>
          </div>
          
          <div v-if="showVariants" class="space-y-6">
            <div
              v-for="(v, vi) in form.variants"
              :key="vi"
              class="border border-gray-200 rounded-xl p-6 bg-gray-50 relative shadow-sm"
            >
              <div class="flex justify-between items-center mb-4">
                <span class="text-lg font-semibold text-gray-700">Variant {{ vi + 1 }}</span>
                <button
                  type="button"
                  @click="removeVariant(vi)"
                  class="text-red-500 hover:text-red-700 transition duration-150 p-2 rounded-full hover:bg-red-50"
                >
                  <XCircleIcon class="h-6 w-6" />
                </button>
              </div>

              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-3">Options</label>
                  <div v-for="(opt, oi) in v.options" :key="oi" class="flex items-center space-x-3 mb-3">
                    <input
                      v-model="form.variants[vi].options[oi].name"
                      placeholder="Option Name (e.g., Size)"
                      class="flex-1 border border-gray-300 rounded-lg px-3 py-2 text-sm focus:ring-2 focus:ring-green-500 focus:border-green-500"
                    />
                    <input
                      v-model="form.variants[vi].options[oi].value"
                      placeholder="Option Value (e.g., Large)"
                      class="flex-1 border border-gray-300 rounded-lg px-3 py-2 text-sm focus:ring-2 focus:ring-green-500 focus:border-green-500"
                    />
                    <button 
                      type="button" 
                      @click="removeVariantOption(vi, oi)" 
                      class="text-red-500 hover:text-red-700 p-2 rounded-full hover:bg-red-50"
                    >
                      <MinusCircleIcon class="h-5 w-5" />
                    </button>
                  </div>
                  <button 
                    type="button" 
                    @click="addVariantOption(vi)" 
                    class="inline-flex items-center text-sm px-3 py-1.5 bg-purple-50 text-purple-600 rounded-lg font-medium hover:bg-purple-100 transition-colors duration-200 shadow-sm"
                  >
                    <PlusCircleIcon class="h-4 w-4 mr-1.5" /> Add Option
                  </button>
                </div>

                <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2">
                      Price <span class="text-red-500">*</span>
                    </label>
                    <div class="relative">
                      <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                        <span class="text-gray-500 sm:text-sm">$</span>
                      </div>
                      <input
                        v-model.number="v.price"
                        type="number"
                        min="0"
                        step="0.01"
                        required
                        class="w-full pl-7 pr-3 py-2.5 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500 transition duration-150 shadow-sm"
                        placeholder="0.00"
                      />
                    </div>
                  </div>
                  
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2">
                      Stock
                    </label>
                    <input
                      v-model.number="v.stock"
                      type="number"
                      min="0"
                      class="w-full border border-gray-300 rounded-lg px-3 py-2.5 text-sm focus:ring-2 focus:ring-green-500 focus:border-green-500 transition duration-150 shadow-sm"
                      placeholder="0"
                    />
                  </div>
                  
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2">
                      Image URL
                    </label>
                    <input
                      v-model="v.image"
                      type="url"
                      placeholder="https://example.com/variant.jpg"
                      class="w-full border border-gray-300 rounded-lg px-3 py-2.5 text-sm focus:ring-2 focus:ring-green-500 focus:border-green-500 transition duration-150 shadow-sm"
                    />
                  </div>
                </div>
              </div>
            </div>
            
            <button
              type="button"
              @click="addVariant"
              class="inline-flex items-center px-5 py-2.5 bg-green-600 text-white rounded-lg font-medium hover:bg-green-700 transition-colors duration-200 shadow-sm"
            >
              <PlusIcon class="h-6 w-6 mr-2" /> Add New Variant
            </button>
          </div>
        </div>

        <div v-if="variantOptionError" class="bg-red-50 border border-red-300 text-red-700 px-4 py-3 rounded-lg">
          <div class="flex items-center">
            <ExclamationCircleIcon class="h-5 w-5 mr-2 flex-shrink-0" />
            <span>{{ variantOptionError }}</span>
          </div>
        </div>

        <div class="flex justify-end space-x-4 pt-6 border-t border-gray-200">
          <button
            type="button"
            @click="() => $router.push({ name: 'Products' })"
            class="px-6 py-3 border border-gray-300 rounded-lg text-gray-700 font-medium hover:bg-gray-100 transition-colors duration-200 shadow-sm"
          >
            Cancel
          </button>
          <button
            type="submit"
            :disabled="saving || !!variantOptionError"
            class="px-6 py-3 bg-green-600 text-white rounded-lg font-medium hover:bg-green-700 transition-colors duration-200 disabled:opacity-50 disabled:cursor-not-allowed shadow-sm"
          >
            <SpinnerIcon v-if="saving" class="animate-spin h-5 w-5 mr-2 inline-block" />
            {{ saving ? 'Creating Product...' : 'Create Product' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { productService } from '@/services/product'
import { collectionService } from '@/services/collection'
import {
  RefreshIcon as SpinnerIcon,
  PlusCircleIcon,
  MinusCircleIcon,
  XIcon as XCircleIcon,
  PlusIcon,
  ExclamationCircleIcon,
  InformationCircleIcon
} from '@heroicons/vue/outline'
import BackButton from '@/components/BackButton.vue'

const router = useRouter()
const shopStore = useShopStore()

const loading = ref(true)
const saving = ref(false)
const error = ref(null)
const variantOptionError = ref(null)

// Form state
const form = reactive({
  name: '',
  description: '',
  main_image: '',
  images: [],
  collection_ids: [],
  price: null,
  stock: null,
  variants: [],
  meta_title: '',
  meta_description: ''
})

const showVariants = ref(false)
const showAdvanced = ref(false)

// Computed properties
const hasVariants = computed(() => form.variants.length > 0)

// SEO fields
form.meta_title = ''
form.meta_description = ''

const collections = ref([])



onMounted(async () => {
  collections.value = await collectionService.fetchAllByShop(shopStore.activeShop.id)
  setTimeout(() => {
    loading.value = false
  }, 500)
})

// Images
function addImage() {
  form.images.push('')
}

function removeImage(i) {
  form.images.splice(i, 1)
}

// Variants
function addVariant() {
  if (!showVariants.value) showVariants.value = true
  form.variants.push({ 
    options: [{ name: '', value: '' }], 
    price: 0, 
    stock: 0, 
    image: '' 
  })
}

function removeVariant(i) {
  form.variants.splice(i, 1)
}

// Options
function addVariantOption(vi) {
  form.variants[vi].options.push({ name: '', value: '' })
}

function removeVariantOption(vi, oi) {
  form.variants[vi].options.splice(oi, 1)
}

// Submit
async function submitForm() {
  saving.value = true
  error.value = null
  variantOptionError.value = null

  try {
    // Basic validation
    if (!form.name.trim()) {
      error.value = 'Product name is required.'
      return
    }

    if (!form.collection_ids || form.collection_ids.length === 0) {
      error.value = 'At least one collection must be selected.'
      return
    }

    if (!hasVariants.value && (!form.price || form.price <= 0)) {
      error.value = 'Product price is required and must be positive if no variants are provided.'
      return
    }

    // Validate variants if present
    if (hasVariants.value) {
      for (const v of form.variants) {
        if (v.price <= 0) {
          error.value = 'All variant prices must be positive numbers.'
          return
        }
        
        if (!v.options || v.options.length === 0) {
          variantOptionError.value = 'Each variant must have at least one option (name:value).'
          return
        }
        
        for (const o of v.options) {
          if (!o.name.trim() || !o.value.trim()) {
            variantOptionError.value = 'Variant option name and value cannot be empty.'
            return
          }
        }
      }
    }

    // Build payload
    const payload = {
      name: form.name.trim(),
      description: form.description.trim(),
      main_image: form.main_image.trim() || '',
      images: form.images.filter(img => img.trim()),
      collection_ids: form.collection_ids,
      meta_title: form.meta_title,
      meta_description: form.meta_description,
    }

    if (hasVariants.value) {
      payload.variants = form.variants.map(v => ({
        options: v.options.filter(o => o.name && o.value),
        price: v.price,
        stock: v.stock || 0,
        image: v.image || ''
      }))
    } else {
      if (typeof form.price === 'number') {
        payload.price = form.price
      }
      if (typeof form.stock === 'number') {
        payload.stock = form.stock
      }
    }

    console.log('Submitting payload:', payload)

    const newProd = await productService.create(shopStore.activeShop.id, payload)
    router.push({ name: 'ProductDetail', params: { productId: newProd.id } })
  } catch (e) {
    console.error('Failed to create product:', e)
    error.value = e.response?.data?.message || 'Failed to create product. Please try again.'
  } finally {
    saving.value = false
  }
}
</script>

<style scoped>
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fade-in {
  animation: fadeIn 0.4s ease-out forwards;
}
</style>