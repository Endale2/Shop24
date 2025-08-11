<template>
  <router-link
    :to="`/products/${product.slug}`"
    class="group flex flex-col h-full bg-background border border-transparent hover:border-default rounded-[var(--border-radius)] p-4 transition-colors cursor-pointer focus:outline-none focus:ring-2 relative"
    :style="focusStyle"
    @mouseenter="onHover"
    @mouseleave="onLeave"
  >
    <!-- Discount Badge at Top-Left -->
    <span
      v-if="product.discounts && product.discounts.length > 0"
      class="absolute top-2 left-2 z-10 inline-block px-3 py-1 text-xs font-semibold rounded-full"
      :style="badgeStyle"
    >
      {{ formatDiscount(product.discounts[0]) }}
    </span>

    <!-- Product Image - Fixed Ratio, theme-friendly background -->
    <div class="relative w-full overflow-hidden mb-4 aspect-square" :style="{ backgroundColor: 'rgba(0,0,0,0.03)' }">
      <img
        :src="currentImage"
        :alt="product.name"
        class="absolute inset-0 w-full h-full object-contain transition-all duration-200"
        loading="lazy"
        @error="onImageError"
      />
    </div>

    <!-- Product Name -->
    <h2 class="font-heading text-heading font-semibold text-base mb-1 tracking-wide uppercase leading-tight line-clamp-1">
      {{ product.name }}
    </h2>

    <!-- Product Subtitle/Description -->
    <div class="text-xs text-body mb-2 line-clamp-2 font-body">
      {{ product.subtitle || product.description }}
    </div>

    <!-- Price Section -->
    <p class="text-heading font-bold text-lg mb-2">
      <template v-if="product.price && product.price > 0 && (!product.variants || product.variants.length === 0)">
        <span v-if="product.display_price && product.display_price < product.price" class="line-through text-body mr-2">
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
import { ref, computed, watch } from 'vue'
import { useCurrentTheme } from '@/services/dynamic-theme'

const currentTheme = useCurrentTheme()

// Reactive badge style that matches collection detail page
const badgeStyle = computed(() => {
  const theme = currentTheme.value

  // Use theme colors if available, otherwise use collection detail page style
  if (theme?.colors?.secondary || theme?.colors?.primary) {
    const badgeColor = theme.colors.secondary || theme.colors.primary
    const textColor = theme.colors.background || '#FFFFFF'

    return {
      backgroundColor: badgeColor,
      color: textColor
    }
  }

  // Fallback to collection detail page style (yellow badge)
  return {
    backgroundColor: '#FEF3C7', // bg-yellow-100
    color: '#92400E' // text-yellow-800
  }
})

// Watch for theme changes to ensure reactivity
watch(currentTheme, (newTheme) => {
  if (newTheme?.colors) {
    console.log('🎨 Theme loaded, updating product card badge colors')
  }
}, { immediate: true })

const focusStyle = computed(() => ({
  '--tw-ring-color': currentTheme.value?.colors?.primary || '#6366F1'
}))

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
  const images: string[] = []
  if (props.product.main_image) images.push(props.product.main_image)
  if (Array.isArray(props.product.images)) {
    images.push(...props.product.images)
  }
  if (Array.isArray(props.product.variants)) {
    for (const v of props.product.variants) {
      if (v?.image) images.push(v.image)
    }
  }
  return [...new Set(images)] // Remove duplicates
})

watch(allImages, (imgs: string[]) => {
  currentImage.value = imgs[0] || ''
})

const currentImage = ref(allImages.value[0] || '')
let hoverImageIndex = 0
let cycleTimer: any = null

function onImageError() {
  if (allImages.value.length > 1) {
    hoverImageIndex = (hoverImageIndex + 1) % allImages.value.length
    currentImage.value = allImages.value[hoverImageIndex]
  } else {
    currentImage.value = '/placeholder-product.jpg'
  }
}

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
  line-clamp: 1;
  display: -webkit-box;
  -webkit-line-clamp: 1;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.line-clamp-2 {
  line-clamp: 2;
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

/* --- Product card upgrade --- */
.router-link {
  background: #fff;
  border: 1px solid #e5e7eb;
  border-radius: 0;
  transition: border-color 0.25s ease, box-shadow 0.25s ease;
}

.router-link:hover {
  border-color: #000;
  box-shadow: 0 6px 18px rgba(0, 0, 0, 0.08);
}

.router-link img {
  background-color: #f9fafb;
  transition: transform 0.3s ease;
}

.router-link:hover img {
  transform: scale(1.03);
}
</style>
