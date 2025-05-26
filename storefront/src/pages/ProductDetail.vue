<template>
  <section v-if="product" class="container mx-auto px-4 py-8">
    <div class="flex flex-col lg:flex-row gap-8">
      <!-- Image gallery -->
      <div class="lg:w-1/2">
        <div class="bg-gray-100 rounded-lg overflow-hidden">
          <img
            :src="activeImage"
            :alt="product.name"
            class="w-full h-auto object-cover aspect-square max-h-[500px]"
          />
        </div>
        <div v-if="product.images.length > 1" class="mt-4 grid grid-cols-4 gap-2">
          <button
            v-for="(img, idx) in product.images"
            :key="idx"
            @click="activeImage = img"
            :class="[
              'border rounded overflow-hidden',
              img === activeImage ? 'border-indigo-500' : 'border-transparent hover:border-gray-300'
            ]"
          >
            <img
              :src="img"
              :alt="`Thumbnail ${idx + 1}`"
              class="w-full h-20 object-cover"
            />
          </button>
        </div>
      </div>

      <!-- Product info -->
      <div class="lg:w-1/2 flex flex-col justify-between">
        <div>
          <h1 class="text-4xl font-bold text-gray-900 mb-4">{{ product.name }}</h1>
          <p class="text-indigo-600 text-3xl font-semibold mb-4">
            ${{ selectedVariant.price.toFixed(2) }}
          </p>
          <p class="text-gray-700 mb-6">{{ product.description }}</p>

          <!-- Variant selectors -->
          <div v-for="(options, optName) in variantOptions" :key="optName" class="mb-4">
            <label :for="optName" class="block text-sm font-medium text-gray-700 mb-1">
              {{ optName }}:
            </label>
            <select
              :id="optName"
              v-model="selectedOptions[optName]"
              class="w-full p-2 border border-gray-300 rounded-md"
            >
              <option
                v-for="optValue in options"
                :key="optValue"
                :value="optValue"
              >
                {{ optValue }}
              </option>
            </select>
          </div>

          <!-- Stock & Add to Cart -->
          <div class="flex items-center space-x-4 mb-6">
            <span
              class="text-sm font-medium"
              :class="selectedVariant.stock > 0 ? 'text-green-600' : 'text-red-600'"
            >
              {{ selectedVariant.stock > 0 ? 'In Stock' : 'Out of Stock' }}
            </span>
            <span class="text-sm text-gray-500">(SKU: {{ product.sku }})</span>
          </div>

          <div class="flex items-center space-x-4">
            <input
              type="number"
              v-model.number="quantity"
              min="1"
              :max="selectedVariant.stock"
              class="w-20 p-2 border border-gray-300 rounded-md"
            />
            <button
              class="bg-indigo-600 hover:bg-indigo-700 text-white font-bold py-2 px-6 rounded-md transition"
              :disabled="selectedVariant.stock === 0"
              @click="addToCart"
            >
              Add to Cart
            </button>
          </div>
        </div>
      </div>
    </div>
  </section>

  <div v-else class="text-center py-20">
    <p class="text-gray-500">Loading product details…</p>
  </div>
</template>

<script setup>
import { ref, computed, onBeforeMount, watch, defineProps } from 'vue'
import { fetchProductBySlug } from '@/services/product'

const props = defineProps({ shopSlug: String, productSlug: String })

const product = ref(null)
const activeImage = ref('')
const quantity = ref(1)
const selectedOptions = ref({})

// Load product data
async function load() {
  product.value = await fetchProductBySlug(props.shopSlug, props.productSlug)
  if (product.value) {
    activeImage.value = product.value.images[0] || ''
    // initialize selectedOptions with first of each
    Object.keys(variantOptions.value).forEach(opt => {
      selectedOptions.value[opt] = variantOptions.value[opt][0]
    })
  }
}

onBeforeMount(load)
watch(() => [props.shopSlug, props.productSlug], load)

// Compute unique variant options map: { Size: [...], Color: [...] }
const variantOptions = computed(() => {
  const opts = {}
  if (!product.value) return opts
  product.value.variants.forEach(v => {
    Object.entries(v.options).forEach(([key, val]) => {
      if (!opts[key]) opts[key] = []
      if (!opts[key].includes(val)) opts[key].push(val)
    })
  })
  return opts
})

// Find the variant matching selected options
const selectedVariant = computed(() => {
  if (!product.value) return { price: 0, stock: 0 }
  return (
    product.value.variants.find(v =>
      Object.entries(v.options).every(([k, vOpt]) => selectedOptions.value[k] === vOpt)
    ) || { price: product.value.price, stock: 0 }
  )
})

// Add to cart placeholder
function addToCart() {
  alert(
    `${quantity.value} x ${product.value.name} (${Object.values(selectedOptions.value).join(', ')}) added to cart!`
  )
}
</script>

<style scoped>
/* No additional styles; using Tailwind classes */
</style>
