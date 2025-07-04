<template>
  <div class="p-4 sm:p-6 max-w-6xl mx-auto space-y-8 font-sans">
    <button
      @click="goBack"
      class="inline-flex items-center text-gray-600 hover:text-green-700 transition duration-200 ease-in-out mb-6 group rounded-full px-3 py-1.5 -ml-3"
    >
      <ChevronLeftIcon class="h-5 w-5 mr-1 text-gray-500 group-hover:text-green-600 transition-colors duration-200" />
      <span class="text-sm font-medium group-hover:text-green-700 transition-colors duration-200">Back to Products</span>
    </button>

    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-20 bg-white rounded-2xl shadow-lg">
      <SpinnerIcon class="animate-spin h-10 w-10 text-green-500 mb-4" />
      <p class="mt-3 text-lg font-semibold text-gray-700">Brewing up the form...</p>
    </div>
    <div
      v-else-if="error"
      class="bg-red-50 border border-red-300 text-red-700 px-6 py-4 rounded-xl shadow-md flex items-center space-x-3"
      role="alert"
    >
      <ExclamationCircleIcon class="h-6 w-6 flex-shrink-0" />
      <div>
        <strong class="font-semibold">Oops!</strong> <span class="ml-1">{{ error }}</span>
        <p class="text-sm mt-1">Please check your input and try again.</p>
      </div>
    </div>

    <div v-else class="bg-white shadow-xl rounded-2xl p-6 sm:p-8 space-y-8 animate-fade-in">
      <h2 class="text-3xl sm:text-4xl font-extrabold text-gray-800 text-center mb-6">
       Add a New Product
      </h2>

      <form @submit.prevent="submitForm" class="space-y-6 sm:space-y-8">
        <div class="space-y-5 border-b border-gray-200 pb-6">
          <h3 class="text-xl font-bold text-gray-700">Basic Details</h3>
          <div>
            <label for="product-name" class="block text-sm font-medium text-gray-700 mb-1">
              Product Name <span class="text-red-500">*</span>
            </label>
            <input
              id="product-name"
              v-model="form.name"
              type="text"
              required
              class="mt-1 w-full border border-gray-300 rounded-lg px-4 py-2.5 focus:ring-2 focus:ring-green-300 focus:border-green-300 transition duration-150 shadow-sm"
              placeholder="e.g., Cozy Comfort Blanket"
            />
          </div>

          <div>
            <label for="product-description" class="block text-sm font-medium text-gray-700 mb-1">
              Description
            </label>
            <textarea
              id="product-description"
              v-model="form.description"
              rows="4"
              class="mt-1 w-full border border-gray-300 rounded-lg px-4 py-2.5 focus:ring-2 focus:ring-green-300 focus:border-green-300 transition duration-150 shadow-sm resize-y"
              placeholder="A warm and fuzzy description of your product..."
            ></textarea>
          </div>

          <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
            <div>
              <label for="price" class="block text-sm font-medium text-gray-700 mb-1">Product Price</label>
              <input
                id="price"
                type="number"
                v-model.number="form.price"
                step="0.01"
                min="0"
                :disabled="form.variants && form.variants.length > 0"
                class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-green-500 focus:border-green-500 sm:text-sm bg-gray-100"
              />
              <div v-if="form.variants && form.variants.length > 0" class="text-xs text-gray-500 mt-1">
                Computed from variants: {{ computedPrice }}
              </div>
            </div>
            <div>
              <label for="display_price" class="block text-sm font-medium text-gray-700 mb-1">Display Price</label>
              <input
                id="display_price"
                type="number"
                v-model.number="form.display_price"
                step="0.01"
                min="0"
                :disabled="form.variants && form.variants.length > 0"
                class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-green-500 focus:border-green-500 sm:text-sm bg-gray-100"
              />
              <div v-if="form.variants && form.variants.length > 0" class="text-xs text-gray-500 mt-1">
                Computed from variants: {{ computedDisplayPrice }}
              </div>
            </div>
            <div>
              <label for="stock" class="block text-sm font-medium text-gray-700 mb-1">Stock</label>
              <input
                id="stock"
                type="number"
                v-model.number="form.stock"
                min="0"
                :disabled="form.variants && form.variants.length > 0"
                class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-green-500 focus:border-green-500 sm:text-sm bg-gray-100"
              />
              <div v-if="form.variants && form.variants.length > 0" class="text-xs text-gray-500 mt-1">
                Computed from variants: {{ computedStock }}
              </div>
            </div>
          </div>

          <div>
            <label for="product-category" class="block text-sm font-medium text-gray-700 mb-1">
              Category
            </label>
            <input
              id="product-category"
              v-model="form.category"
              type="text"
              class="mt-1 w-full border border-gray-300 rounded-lg px-4 py-2.5 focus:ring-2 focus:ring-green-300 focus:border-green-300 transition duration-150 shadow-sm"
              placeholder="e.g., Home Goods"
            />
          </div>
        </div>

        <div class="space-y-5 border-b border-gray-200 pb-6">
          <h3 class="text-xl font-bold text-gray-700">Product Images</h3>
          <div>
            <label for="main-image" class="block text-sm font-medium text-gray-700 mb-1">
              Main Image URL
            </label>
            <input
              id="main-image"
              v-model="form.main_image"
              type="url"
              placeholder="https://example.com/main-product.jpg"
              class="mt-1 w-full border border-gray-300 rounded-lg px-4 py-2.5 focus:ring-2 focus:ring-green-300 focus:border-green-300 transition duration-150 shadow-sm"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
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
                class="flex-1 border border-gray-300 rounded-lg px-4 py-2.5 focus:ring-2 focus:ring-green-300 focus:border-green-300 transition duration-150 shadow-sm"
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
              class="inline-flex items-center px-4 py-2 bg-green-50 text-green-600 rounded-full font-medium hover:bg-green-100 transition-colors duration-200 shadow-sm"
            >
              <PlusCircleIcon class="h-5 w-5 mr-2" /> Add Another Image
            </button>
          </div>
        </div>

        <div class="space-y-5">
          <h3 class="text-xl font-bold text-gray-700">Product Variants</h3>
          <button
            v-if="!showVariants"
            type="button"
            @click="showVariants = true"
            class="inline-flex items-center px-5 py-2.5 bg-purple-600 text-white rounded-full font-bold hover:bg-purple-700 transition-colors duration-200 shadow-md mb-4"
          >
            <PlusIcon class="h-6 w-6 mr-2" /> Add Variant
          </button>
          <div v-if="showVariants">
            <div
              v-for="(v, vi) in form.variants"
              :key="vi"
              class="border border-gray-200 rounded-xl p-5 mb-6 bg-gray-50 relative shadow-sm"
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

              <div class="space-y-3 mb-4">
                <label class="block text-sm font-medium text-gray-700">Options</label>
                <div v-for="(opt, oi) in v.options" :key="oi" class="flex items-center space-x-2 mb-2">
                  <input
                    v-model="form.variants[vi].options[oi].name"
                    placeholder="Option Name (e.g., Size)"
                    class="w-1/3 border border-gray-300 rounded-lg px-3 py-2 text-sm"
                  />
                  <input
                    v-model="form.variants[vi].options[oi].value"
                    placeholder="Option Value (e.g., Large)"
                    class="w-1/3 border border-gray-300 rounded-lg px-3 py-2 text-sm"
                  />
                  <button type="button" @click="removeVariantOption(vi, oi)" class="text-red-500 hover:text-red-700 p-1 rounded-full hover:bg-red-50">
                    <MinusCircleIcon class="h-5 w-5" />
                  </button>
                </div>
                <button type="button" @click="addVariantOption(vi)" class="inline-flex items-center text-sm px-3 py-1.5 bg-purple-50 text-purple-600 rounded-full font-medium hover:bg-purple-100 transition-colors duration-200 shadow-sm">
                  <PlusCircleIcon class="h-4 w-4 mr-1.5" /> Add Option
                </button>
              </div>

              <div class="grid grid-cols-1 sm:grid-cols-2 gap-6 mt-5">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">
                    Variant Price <span class="text-red-500">*</span>
                  </label>
                  <div class="relative mt-1">
                    <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                      <span class="text-gray-500 sm:text-sm">$</span>
                    </div>
                    <input
                      v-model.number="v.price"
                      type="number"
                      min="0"
                      step="0.01"
                      required
                      class="w-full pl-7 pr-3 py-2.5 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-300 focus:border-green-300 transition duration-150 shadow-sm"
                      placeholder="0.00"
                    />
                  </div>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">
                    Variant Image URL
                  </label>
                  <input
                    v-model="v.image"
                    type="url"
                    placeholder="https://example.com/variant-red.jpg"
                    class="mt-1 w-full border border-gray-300 rounded-lg px-4 py-2.5 focus:ring-2 focus:ring-green-300 focus:border-green-300 transition duration-150 shadow-sm"
                  />
                </div>
              </div>
            </div>
            <button
              type="button"
              @click="addVariant"
              class="inline-flex items-center px-5 py-2.5 bg-green-600 text-white rounded-full font-bold hover:bg-green-700 transition-colors duration-200 shadow-md"
            >
              <PlusIcon class="h-6 w-6 mr-2" /> Add New Variant
            </button>
          </div>
        </div>

        <div v-if="variantOptionError" class="bg-red-50 border border-red-300 text-red-700 px-4 py-2 rounded mb-4">
          <ExclamationCircleIcon class="h-5 w-5 inline-block mr-2" />
          {{ variantOptionError }}
        </div>

        <div class="flex justify-end space-x-4 pt-6 border-t border-gray-200">
          <button
            type="button"
            @click="goBack"
            class="px-6 py-2.5 border border-gray-300 rounded-full text-gray-700 font-medium hover:bg-gray-100 transition-colors duration-200 shadow-sm"
          >
            Cancel
          </button>
          <button
            type="submit"
            :disabled="saving || !!variantOptionError"
            class="px-6 py-2.5 bg-green-600 text-white rounded-full font-bold hover:bg-green-700 transition-colors duration-200 disabled:opacity-50 disabled:cursor-not-allowed shadow-md"
          >
            <SpinnerIcon v-if="saving" class="animate-spin h-5 w-5 mr-2 inline-block" />
            {{ saving ? 'Saving Product...' : 'Save Product' }}
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
import { useAuthStore } from '@/store/auth'
import { productService } from '@/services/product'
import {
  ChevronLeftIcon,
  RefreshIcon as SpinnerIcon,
  PlusCircleIcon,
  MinusCircleIcon,
  XIcon as XCircleIcon, // Renamed for consistent circle icon usage
  PlusIcon,
  ExclamationCircleIcon // For error messages
} from '@heroicons/vue/outline'

const router = useRouter()
const shopStore = useShopStore()
const authStore = useAuthStore()

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
  category: '',
  price: null,
  display_price: null,
  stock: null,
  variants: []
})

// temp keys for option renaming
const variantKeyTemp = reactive({})

// Add a flag to control variant section visibility
const showVariants = ref(false)

function goBack() {
  router.push({ name: 'Products' })
}

onMounted(() => {
  // Simulate loading to show spinner, remove in production if data loads instantly
  setTimeout(() => {
    loading.value = false;
  }, 500);
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
  form.variants.push({ options: [{ name: '', value: '' }], price: 0, stock: 0, image: '' })
}
function removeVariant(i) {
  form.variants.splice(i, 1)
  // Clean up variantKeyTemp when a variant is removed and re-index it
  const newVariantKeyTemp = {};
  form.variants.forEach((_, idx) => {
    newVariantKeyTemp[idx] = variantKeyTemp[idx] || {};
  });
  Object.assign(variantKeyTemp, newVariantKeyTemp);
  for (const key in variantKeyTemp) {
    if (parseInt(key) >= form.variants.length) {
      delete variantKeyTemp[key];
    }
  }
}

// Options
function addVariantOption(vi) {
  form.variants[vi].options.push({ name: '', value: '' })
}
function removeVariantOption(vi, oi) {
  form.variants[vi].options.splice(oi, 1)
}
function renameOptionKey(vi, oldKey, newKey) {
  if (!newKey || oldKey === newKey) {
    if (variantKeyTemp[vi]) {
      variantKeyTemp[vi][oldKey] = oldKey; // Revert temp if empty or same
    }
    return;
  }
  
  if (form.variants[vi].options.hasOwnProperty(newKey) && newKey !== oldKey) {
    console.warn(`Variant option "${newKey}" already exists. Please use a unique name.`);
    if (variantKeyTemp[vi]) {
      variantKeyTemp[vi][oldKey] = oldKey; // Revert temp key
    }
    return;
  }

  const val = form.variants[vi].options[oldKey]
  delete form.variants[vi].options[oldKey]
  form.variants[vi].options[newKey] = val
  if (variantKeyTemp[vi]) {
    delete variantKeyTemp[vi][oldKey]
    variantKeyTemp[vi][newKey] = newKey // Store the new key as its own value
  }
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
      return;
    }
    if (form.price == null || form.price <= 0) {
      error.value = 'Product price is required and must be positive if no variants are provided.'
      saving.value = false;
      return;
    }
    // Validate variant prices
    for (const v of form.variants) {
      if (v.price <= 0) {
        error.value = 'All variant prices must be positive numbers.'
        return;
      }
    }
    // Validate variant options
    if (form.variants.length > 0) {
      for (const v of form.variants) {
        if (!v.options || v.options.length === 0) {
          variantOptionError.value = 'Each variant must have at least one option (name:value).';
          saving.value = false;
          return;
        }
        for (const o of v.options) {
          if (!o.name.trim() || !o.value.trim()) {
            variantOptionError.value = 'Variant option name and value cannot be empty.';
            saving.value = false;
            return;
          }
        }
      }
    }

    const payload = {
      name: form.name.trim(),
      description: form.description.trim(),
      main_image: form.main_image.trim() || undefined,
      images: form.images.filter((i) => i.trim()),
      category: form.category.trim(),
      ...(form.variants.length === 0 ? {
        price: form.price,
        display_price: form.display_price,
        stock: form.stock
      } : {}),
      variants: showVariants.value && form.variants.length > 0 ? form.variants.map((v) => ({
        options: v.options.filter(o => o.name && o.value),
        price: v.price,
        stock: v.stock || 0,
        image: v.image || undefined
      })) : []
    }
    // Debugging: Log the payload before sending
    console.log("Submitting payload:", payload);

    const newProd = await productService.create(shopStore.active.id, payload)
    router.push({ name: 'ProductDetail', params: { productId: newProd.id } })
  } catch (e) {
    console.error('Failed to create product:', e)
    error.value = e.response?.data?.message || 'Failed to create product. Please try again.'
  } finally {
    saving.value = false
  }
}

const computedPrice = computed(() => {
  if (!form.variants || form.variants.length === 0) return form.price;
  return Math.min(...form.variants.map(v => v.price));
});
const computedDisplayPrice = computed(() => {
  if (!form.variants || form.variants.length === 0) return form.display_price;
  const displayPrices = form.variants.map(v => v.display_price).filter(p => p != null);
  return displayPrices.length ? Math.max(...displayPrices) : '';
});
const computedStock = computed(() => {
  if (!form.variants || form.variants.length === 0) return form.stock;
  return form.variants.reduce((sum, v) => sum + (v.stock || 0), 0);
});
</script>

<style scoped>
/* Base animation for elements fading in */
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

/* No pulse-price needed here as it's for details, not form */

/* Specific styling for required field indicator */
label span.text-red-500 {
  margin-left: 0.25rem;
}
</style>