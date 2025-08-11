<template>
  <router-link
    :to="`/products/${product.slug}`"
    class="bg-white border border-gray-200 rounded-none p-4 group flex flex-col h-full transition-colors hover:border-black cursor-pointer focus:outline-none focus:ring-2 focus:ring-blue-200 relative"
    @mouseenter="onHover"
    @mouseleave="onLeave"
  >
    <!-- Discount Badge at Top-Left -->
    <span
      v-if="product.discounts && product.discounts.length > 0"
      class="absolute top-2 left-2 z-10 px-3 py-1 bg-gradient-to-r from-yellow-200 to-yellow-100 text-yellow-900 text-xs font-bold rounded-full shadow-sm border border-yellow-300 flex items-center gap-1"
      style="min-width: 60px; justify-content: center; letter-spacing: 0.03em;"
    >
      <svg class="w-4 h-4 mr-1 text-yellow-500" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" d="M9 14l2-2 4 4m6-2a9 9 0 11-18 0 9 9 0 0118 0z"/>
      </svg>
      {{ formatDiscount(product.discounts[0]) }}
    </span>

    <!-- Product Image - Fixed Ratio, No Overlay Icons -->
    <img
      :src="currentImage"
      :alt="product.name"
      class="w-full h-56 object-contain mb-4 bg-gray-50 transition-all duration-200"
      loading="lazy"
    />

    <!-- Product Name -->
    <h2 class="font-semibold text-base mb-1 tracking-wide uppercase text-gray-900 leading-tight line-clamp-1">
      {{ product.name }}
    </h2>

    <!-- Product Subtitle/Description -->
    <div class="text-xs text-gray-500 mb-2 line-clamp-2">
      {{ product.subtitle || product.description }}
    </div>

    <!-- Price Section -->
    <p class="text-gray-900 font-bold text-lg mb-2">
      <template v-if="product.price && product.price > 0 && (!product.variants || product.variants.length === 0)">
        <span v-if="product.display_price && product.display_price < product.price" class="line-through text-gray-400 mr-2">
          ${{ product.price.toFixed(2) }}
        </span>
        <span>
          ${{ (product.display_price || product.price).toFixed(2) }}
        </span>
      </template>
      <template v-else-if="product.starting_price && product.starting_price > 0">
        <span class="italic text-gray-700">from</span> ${{ product.starting_price.toFixed(2) }}
      </template>
      <template v-else>
        N/A
      </template>
    </p>

    <!-- Rating (if enabled and available) -->
    <div
      v-if="product.rating"
      class="flex items-center space-x-1 mb-2"
    >
      <div class="flex items-center">
        <svg
          v-for="star in 5"
          :key="star"
          class="w-3 h-3"
          :class="star <= Math.floor(product.rating) ? 'text-yellow-400' : 'text-gray-300'"
          fill="currentColor"
          viewBox="0 0 20 20"
        >
          <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
        </svg>
      </div>
      <span class="text-xs text-gray-500">{{ product.rating }}</span>
    </div>

    <!-- Stock Status -->
    <div
      v-if="product.stock_quantity !== undefined"
      class="text-xs mt-auto"
      :class="product.stock_quantity > 0 ? 'text-green-600' : 'text-red-600'"
    >
      {{ product.stock_quantity > 0 ? `${product.stock_quantity} in stock` : 'Out of stock' }}
    </div>
  </router-link>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

interface Product {
  id: string
  name: string
  slug: string
  description?: string
  subtitle?: string
  main_image?: string
  images?: string[]
  price?: number
  display_price?: number
  starting_price?: number
  variants?: any[]
  discounts?: Array<{
    id: string
    name: string
    description: string
    category: string
    type: string
    value: number
    coupon_code?: string
    start_at: string
    end_at: string
    active: boolean
  }>
  rating?: number
  reviewCount?: number
  stock_quantity?: number
  created_at?: string
}

interface Props {
  product: Product
}

const props = defineProps<Props>()

// =============== Image Management ===============
const allImages = computed(() => {
  const images = []
  if (props.product.main_image) images.push(props.product.main_image)
  if (props.product.images) images.push(...props.product.images)
  return [...new Set(images)] // Remove duplicates
})

const currentImage = ref(allImages.value[0] || '')
let hoverImageIndex = 0
let cycleTimer: any = null

// =============== Event Handlers ===============
function onHover() {
  if (allImages.value.length > 1) {
    hoverImageIndex = 1
    currentImage.value = allImages.value[hoverImageIndex]
    // If more than 2 images, cycle through them every 1s while hovered
    if (allImages.value.length > 2) {
      startImageCycle()
    }
  }
}

function onLeave() {
  hoverImageIndex = 0
  currentImage.value = allImages.value[0] || ''
  stopImageCycle()
}

function startImageCycle() {
  stopImageCycle()
  cycleTimer = setInterval(() => {
    hoverImageIndex = (hoverImageIndex + 1) % allImages.value.length
    if (hoverImageIndex === 0) hoverImageIndex = 1 // skip main image
    currentImage.value = allImages.value[hoverImageIndex]
  }, 1000)
}

function stopImageCycle() {
  if (cycleTimer) {
    clearInterval(cycleTimer)
    cycleTimer = null
  }
}

function formatDiscount(discount: any): string {
  return discount.type === 'percentage'
    ? `${discount.value}% OFF`
    : `$${discount.value} OFF`
}
</script>

<style scoped>
/* Enhanced animations */
@keyframes fade-in {
  from { 
    opacity: 0; 
    transform: translateY(-10px); 
  }
  to { 
    opacity: 1; 
    transform: translateY(0); 
  }
}

.animate-fade-in {
  animation: fade-in 0.4s ease-out;
}

/* Line clamp utilities */
.line-clamp-1 {
  display: -webkit-box;
  -webkit-line-clamp: 1;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* Focus states */
button:focus,
a:focus {
  outline: 2px solid var(--color-primary, #10B981);
  outline-offset: 2px;
}
</style>
