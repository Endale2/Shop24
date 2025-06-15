<template>
  <div class="min-h-screen flex flex-col bg-gray-50 text-gray-900 font-sans antialiased">
    <div class="container mx-auto px-4 py-12 lg:py-16">
      <!-- Loading -->
      <div v-if="loading" class="flex items-center justify-center h-96 bg-white rounded-xl shadow-lg">
        <svg class="animate-spin h-16 w-16 text-green-600" xmlns="http://www.w3.org/2000/svg" fill="none"
             viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
          <path class="opacity-75" fill="currentColor"
                d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291
                   A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
        </svg>
      </div>

      <!-- Error -->
      <div v-else-if="error" class="text-center text-red-600 bg-red-50 p-6 rounded-xl shadow-md">
        <p class="text-lg font-semibold">Oops! Something went wrong.</p>
        <p class="text-base">Failed to load product details. Please try again later.</p>
      </div>

      <!-- Product -->
      <div v-else-if="product"
           class="bg-white rounded-xl shadow-xl p-6 lg:p-10 flex flex-col lg:flex-row gap-8 lg:gap-12">
        <!-- Gallery -->
        <div class="lg:w-1/2 flex flex-col">
          <div
            class="relative w-full rounded-lg overflow-hidden border border-gray-200 aspect-square flex items-center justify-center shadow-md">
            <img
              :src="currentDisplayedImage"
              :alt="product.name"
              class="w-full h-full object-contain"
              @error="onImageError($event, 500)"
            />
          </div>
          <div v-if="galleryImages.length > 1" class="mt-4 grid grid-cols-4 gap-3">
            <button
              v-for="(img, idx) in galleryImages"
              :key="idx"
              @click="setActiveImage(img)"
              :class="[
                'border-2 rounded-lg overflow-hidden transition-all duration-200 cursor-pointer',
                img === currentDisplayedImage
                  ? 'border-green-500 ring-2 ring-green-300'
                  : 'border-gray-200 hover:border-green-300'
              ]"
            >
              <img
                :src="img"
                :alt="`Thumb ${idx+1}`"
                class="w-full h-20 object-cover"
                @error="onImageError($event, 80)"
              />
            </button>
          </div>
        </div>

        <!-- Info -->
        <div class="lg:w-1/2 flex flex-col justify-between">
          <div>
            <h1 class="text-4xl lg:text-5xl font-extrabold text-gray-900 mb-4 leading-tight">
              {{ product.name }}
            </h1>
            <p class="text-green-600 text-3xl lg:text-4xl font-bold mb-4">
              ${{ displayPrice.toFixed(2) }}
              <span v-if="product.display_price < product.price"
                    class="text-lg text-gray-500 line-through ml-2">
                ${{ product.price.toFixed(2) }}
              </span>
            </p>
            <p class="text-gray-700 text-base leading-relaxed mb-6">{{ product.description }}</p>

            <!-- Variants -->
            <div v-for="(opts, key) in variantOptions" :key="key" class="mb-6">
              <label class="block text-base font-semibold text-gray-800 mb-3">{{ key }}:</label>
              <div class="flex flex-wrap gap-3">
                <button
                  v-for="opt in opts"
                  :key="opt"
                  @click="selectOption(key, opt)"
                  :class="[
                    'px-5 py-2 rounded-full border-2 text-base font-medium transition-all duration-200 cursor-pointer shadow-sm',
                    selectedOptions[key] === opt
                      ? 'bg-green-600 text-white border-green-600 shadow-md'
                      : 'bg-gray-100 text-gray-700 border-gray-200 hover:bg-gray-200 hover:border-green-300'
                  ]"
                >
                  {{ opt }}
                </button>
              </div>
            </div>

            <!-- Stock & SKU -->
            <div class="flex items-center space-x-6 mb-6">
              <span
                :class="selectedVariant.stock > 0
                  ? 'text-green-600 bg-green-50 px-3 py-1 rounded-full'
                  : 'text-red-600 bg-red-50 px-3 py-1 rounded-full'"
                class="text-sm font-semibold"
              >
                {{ selectedVariant.stock > 0 ? 'In Stock' : 'Out of Stock' }}
              </span>
              <span class="text-sm text-gray-500">
                SKU: <span class="font-mono text-gray-600">{{ selectedVariant.id || product.id }}</span>
              </span>
            </div>

            <!-- Quantity & Add -->
            <div class="flex flex-col sm:flex-row items-center space-y-4 sm:space-y-0 sm:space-x-4 mt-6">
              <input
                type="number"
                v-model.number="quantity"
                min="1"
                :max="selectedVariant.stock"
                class="w-full sm:w-24 p-3 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200 text-center"
              />
              <button
                @click="addToCart"
                :disabled="selectedVariant.stock === 0 || quantity < 1"
                class="w-full sm:w-auto flex-grow bg-green-600 hover:bg-green-700 text-white font-bold py-3 px-8 rounded-lg shadow-md hover:shadow-lg disabled:opacity-60 disabled:cursor-not-allowed transition-all duration-200"
              >
                Add to Cart
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { defineProps } from 'vue'
import { fetchProductBySlug } from '@/services/product'

const props = defineProps({
  shopSlug:    { type: String, required: true },
  productSlug: { type: String, required: true }
})

const product               = ref(null)
const loading               = ref(true)
const error                 = ref(false)
const galleryImages         = ref([])
const currentDisplayedImage = ref('')
const quantity              = ref(1)
const selectedOptions       = ref({})

// Fetch and build gallery
onMounted(async () => {
  try {
    product.value = await fetchProductBySlug(props.shopSlug, props.productSlug)

    // Use a Set to dedupe
    const imgs = new Set()
    if (product.value.main_image) imgs.add(product.value.main_image)
    ;(product.value.images || []).forEach(i => imgs.add(i))
    ;(product.value.variants || []).forEach(v => v.image && imgs.add(v.image))

    galleryImages.value = Array.from(imgs)
    currentDisplayedImage.value = product.value.main_image
      || galleryImages.value[0]
      || ''

    // Initialize options
    Object.entries(variantOptions.value).forEach(([k, opts]) => {
      selectedOptions.value[k] = opts[0]
    })
  } catch (e) {
    console.error('Failed to load product:', e)
    error.value = true
  } finally {
    loading.value = false
  }
})

const variantOptions = computed(() => {
  const opts = {}
  if (!product.value?.variants) return opts
  product.value.variants.forEach(v =>
    Object.entries(v.options).forEach(([k, val]) => {
      if (!opts[k]) opts[k] = []
      if (!opts[k].includes(val)) opts[k].push(val)
    })
  )
  return opts
})

const selectedVariant = computed(() => {
  if (!product.value) return { id: null, price: 0, stock: 0, image: product.value?.main_image || '' }
  const match = product.value.variants.find(v =>
    Object.entries(selectedOptions.value).every(([k, val]) => v.options[k] === val)
  )
  return match || { id: null, price: product.value.display_price, stock: 0, image: product.value.main_image || '' }
})

watch(selectedVariant, v => {
  if (v?.image) currentDisplayedImage.value = v.image
}, { immediate: true })

const displayPrice = computed(() =>
  selectedVariant.value.price ?? product.value?.display_price ?? 0
)

function setActiveImage(url) {
  currentDisplayedImage.value = url
}

function selectOption(key, val) {
  selectedOptions.value[key] = val
}

function onImageError(evt, size) {
  evt.target.src = `https://placehold.co/${size}x${size}/E0F2F7/263238?text=No+Image`
}

function addToCart() {
  if (!selectedVariant.value.id) {
    alert('Please select a valid variant.')
    return
  }
  if (quantity.value < 1 || quantity.value > selectedVariant.value.stock) {
    alert('Enter a valid quantity.')
    return
  }

  const cart = JSON.parse(localStorage.getItem('cart') || '[]')
  const idx  = cart.findIndex(i => i.variantId === selectedVariant.value.id)

  if (idx > -1) {
    cart[idx].quantity += quantity.value
  } else {
    cart.push({
      variantId: selectedVariant.value.id,
      productId: product.value.id,
      name:      product.value.name,
      price:     displayPrice.value,
      quantity:  quantity.value,
      options:   { ...selectedOptions.value },
      image:     currentDisplayedImage.value
    })
  }

  localStorage.setItem('cart', JSON.stringify(cart))
  alert('Added to cart!')
}
</script>

<style scoped>
/* any overrides go here */
</style>
