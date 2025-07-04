<template>
  <div v-if="product" class="space-y-6">
    <div class="grid md:grid-cols-2 gap-6">
      <!-- MAIN IMAGE + THUMBNAILS -->
      <div>
        <img
          :src="currentImage"
          class="w-full h-80 object-cover rounded"
          alt="Product image"
        />

        <div class="flex mt-2 space-x-2 overflow-x-auto">
          <img
            v-for="(img, idx) in galleryImages"
            :key="idx"
            :src="img"
            @click="selectThumbnail(img)"
            :class="[
              'w-20 h-20 object-cover rounded cursor-pointer',
              currentImage === img ? 'ring-2 ring-blue-500' : ''
            ]"
            alt="Thumbnail"
          />
        </div>
      </div>

      <!-- DETAILS & VARIANT BUTTONS -->
      <div>
        <h1 class="text-3xl font-bold">{{ product.name }}</h1>
        <p class="text-gray-700 my-2">{{ product.description }}</p>

        <!-- PRICE (always a number now) -->
        <p class="text-xl font-semibold mb-4">
          ${{ displayPrice.toFixed(2) }}
        </p>

        <!-- Variant buttons -->
        <div v-if="product.variants.length" class="space-x-2 mb-4">
          <button
            v-for="v in product.variants"
            :key="v.id"
            @click="selectVariant(v)"
            class="px-4 py-2 border rounded hover:bg-gray-100"
            :class="{ 'bg-blue-500 text-white': selectedVariant?.id === v.id }"
          >
            {{ v.option.name
              ? `${v.option.name}: ${v.option.value}`
              : 'Default' }}
          </button>
        </div>

        <button
          class="w-full py-2 bg-blue-500 text-white rounded disabled:opacity-50"
          :disabled="!canAddToCart"
        >
          Add to cart
        </button>
      </div>
    </div>
  </div>

  <div v-else class="text-center py-12">Loading product…</div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { fetchProductDetail, Product, Variant } from '@/services/product'

const route = useRoute()
const product = ref<Product | null>(null)
const selectedVariant = ref<Variant | null>(null)
const currentImage = ref<string>('')

// Fetch the product once the component mounts
onMounted(async () => {
  const handle = route.params.handle as string
  product.value = await fetchProductDetail(handle)
  if (product.value) {
    currentImage.value = product.value.main_image
  }
})

// Build a mixed gallery of main + other images + variant images
const galleryImages = computed<string[]>(() => {
  if (!product.value) return []
  const imgs = new Set<string>([product.value.main_image, ...product.value.images])
  for (const v of product.value.variants) {
    if (v.image) imgs.add(v.image)
  }
  return Array.from(imgs)
})

// User clicked a thumbnail: switch image & clear variant
function selectThumbnail(img: string) {
  currentImage.value = img
  selectedVariant.value = null
}

// User clicked a variant: select it & switch image if it has one
function selectVariant(v: Variant) {
  selectedVariant.value = v
  currentImage.value = v.image || product!.main_image
}

// Display price: fallback to 0 until product loads
const displayPrice = computed<number>(() => {
  if (!product.value) return 0
  return selectedVariant.value?.price ?? product.value.display_price
})

// Add-to-cart enabled only when stock > 0
const canAddToCart = computed<boolean>(() => {
  if (!product.value) return false
  return selectedVariant.value
    ? selectedVariant.value.stock > 0
    : product.value.stock > 0
})
</script>
