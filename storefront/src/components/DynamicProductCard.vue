<template>
  <router-link
    :to="`/products/${product.slug}`"
    class="product-card group flex flex-col h-full transition-all duration-300 relative overflow-hidden"
    :class="cardClasses"
    :style="cardStyle"
    @mouseenter="onHover"
    @mouseleave="onLeave"
  >
    <!-- Discount Badge -->
    <div
      v-if="product.discounts && product.discounts.length > 0 && config?.showBadges"
      class="absolute top-3 left-3 z-10 px-2 py-1 text-xs font-bold rounded-full shadow-sm"
      :style="discountBadgeStyle"
    >
      <svg class="w-3 h-3 mr-1 inline" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" d="M9 14l2-2 4 4m6-2a9 9 0 11-18 0 9 9 0 0118 0z"/>
      </svg>
      {{ formatDiscount(product.discounts[0]) }}
    </div>

    <!-- Product Image Container -->
    <div
      class="relative overflow-hidden"
      :class="imageContainerClasses"
      :style="imageContainerStyle"
    >
      <img
        :src="currentImage"
        :alt="product.name"
        class="w-full h-full object-cover transition-all duration-500"
        :class="imageClasses"
        :style="imageStyle"
        loading="lazy"
      />

      <!-- Quick View Overlay (appears on hover) -->
      <div
        v-if="showQuickView"
        class="absolute inset-0 bg-black bg-opacity-50 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity duration-300"
      >
        <button
          class="px-4 py-2 text-sm font-medium rounded-lg transition-all duration-200 hover:scale-105"
          :style="quickViewButtonStyle"
          @click.prevent="openQuickView"
        >
          Quick View
        </button>
      </div>
    </div>

    <!-- Product Information -->
    <div class="flex-1 p-4 space-y-2">
      <!-- Product Name -->
      <h3
        class="font-semibold tracking-wide uppercase leading-tight line-clamp-2"
        :style="productNameStyle"
        :class="productNameClasses"
      >
        {{ product.name }}
      </h3>

      <!-- Product Subtitle/Description -->
      <p
        v-if="product.subtitle || product.description"
        class="text-sm line-clamp-2"
        :style="productDescriptionStyle"
      >
        {{ product.subtitle || product.description }}
      </p>

      <!-- Rating (if enabled and available) -->
      <div
        v-if="config?.showRating && product.rating"
        class="flex items-center space-x-1"
      >
        <div class="flex items-center">
          <svg
            v-for="star in 5"
            :key="star"
            class="w-4 h-4"
            :class="star <= Math.floor(product.rating) ? 'text-yellow-400' : 'text-gray-300'"
            fill="currentColor"
            viewBox="0 0 20 20"
          >
            <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
          </svg>
        </div>
        <span class="text-xs" :style="ratingTextStyle">
          ({{ product.reviewCount || 0 }})
        </span>
      </div>

      <!-- Price -->
      <div class="space-y-1">
        <div class="flex items-center space-x-2">
          <template v-if="product.price > 0 && (!product.variants || product.variants.length === 0)">
            <span
              v-if="product.display_price && product.display_price < product.price"
              class="line-through text-sm"
              :style="originalPriceStyle"
            >
              ${{ product.price.toFixed(2) }}
            </span>
            <span
              class="font-bold"
              :style="currentPriceStyle"
              :class="currentPriceClasses"
            >
              ${{ (product.display_price || product.price).toFixed(2) }}
            </span>
          </template>
          <template v-else-if="product.starting_price > 0">
            <span class="text-sm" :style="fromTextStyle">from</span>
            <span
              class="font-bold"
              :style="currentPriceStyle"
              :class="currentPriceClasses"
            >
              ${{ product.starting_price.toFixed(2) }}
            </span>
          </template>
          <template v-else>
            <span :style="currentPriceStyle" class="font-bold">Contact for Price</span>
          </template>
        </div>
      </div>

      <!-- Add to Cart Button (appears on hover for some styles) -->
      <div
        v-if="showAddToCartButton"
        class="pt-3 opacity-0 group-hover:opacity-100 transition-opacity duration-300"
      >
        <button
          class="w-full py-2 text-sm font-medium rounded-lg transition-all duration-200 hover:scale-105"
          :style="addToCartButtonStyle"
          @click.prevent="addToCart"
        >
          Add to Cart
        </button>
      </div>
    </div>
  </router-link>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import type { DynamicTheme } from '@/services/dynamic-theme'

// =============== Props ===============

interface ProductCardConfig {
  style: string
  showRating: boolean
  showBadges: boolean
  imageRatio: string
  hoverEffect: string
}

interface Product {
  id: string
  name: string
  slug: string
  description?: string
  subtitle?: string
  main_image?: string
  images?: string[]
  price: number
  display_price?: number
  starting_price?: number
  variants?: any[]
  discounts?: Array<{
    type: 'percentage' | 'fixed'
    value: number
  }>
  rating?: number
  reviewCount?: number
}

interface Props {
  product: Product
  theme?: DynamicTheme | null
  config?: ProductCardConfig | null
}

const props = defineProps<Props>()

// =============== Reactive State ===============

const isHovered = ref(false)
const showQuickView = ref(false)

// Gather all product images
const allImages = computed(() => {
  const imgs: string[] = []
  if (props.product.main_image) imgs.push(props.product.main_image)
  if (Array.isArray(props.product.images)) {
    props.product.images.forEach(img => {
      if (img && !imgs.includes(img)) imgs.push(img)
    })
  }
  if (Array.isArray(props.product.variants)) {
    props.product.variants.forEach(variant => {
      if (variant.image && !imgs.includes(variant.image)) imgs.push(variant.image)
    })
  }
  return imgs
})

const currentImage = ref(allImages.value[0] || '')
let hoverImageIndex = 0
let imageRotationTimer: number | null = null

// =============== Computed Styles ===============

// Card styling based on theme and config
const cardStyle = computed(() => {
  if (!props.theme) return {}

  const style: any = {
    backgroundColor: props.theme.colors?.background || '#ffffff',
    borderColor: props.theme.colors?.border || '#E5E7EB',
    borderRadius: getBorderRadiusValue(props.theme.layout?.borderStyle || 'rounded')
  }

  // Apply hover effect styling
  if (props.config?.hoverEffect === 'lift' && isHovered.value) {
    style.transform = 'translateY(-4px)'
    style.boxShadow = '0 10px 25px rgba(0, 0, 0, 0.15)'
  } else if (props.config?.hoverEffect === 'scale' && isHovered.value) {
    style.transform = 'scale(1.02)'
  }

  return style
})

const cardClasses = computed(() => {
  const classes = ['border']

  // Add style-specific classes
  if (props.config?.style === 'dynamic') {
    classes.push('hover:border-primary')
  }

  return classes
})

// Image container styling based on image ratio
const imageContainerClasses = computed(() => {
  const ratio = props.config?.imageRatio || 'square'

  switch (ratio) {
    case 'square': return 'aspect-square'
    case 'portrait': return 'aspect-[3/4]'
    case 'landscape': return 'aspect-[4/3]'
    case 'wide': return 'aspect-[16/9]'
    default: return 'aspect-square'
  }
})

const imageContainerStyle = computed(() => {
  return {
    backgroundColor: props.theme?.colors?.background || '#F9FAFB'
  }
})

const imageClasses = computed(() => {
  const classes = []

  // Add hover effect classes
  if (props.config?.hoverEffect === 'zoom') {
    classes.push('group-hover:scale-110')
  }

  return classes
})

const imageStyle = computed(() => {
  const style: any = {}

  // Apply image-specific styling based on config
  if (props.config?.style === 'dynamic') {
    style.objectFit = 'cover'
  }

  return style
})

// Discount badge styling
const discountBadgeStyle = computed(() => {
  return {
    backgroundColor: props.theme?.colors?.secondary || '#F59E0B',
    color: props.theme?.colors?.background || 'white'
  }
})

// Product name styling
const productNameStyle = computed(() => {
  return {
    color: props.theme?.colors?.heading || '#1F2937',
    fontFamily: props.theme?.fonts?.heading ? `'${props.theme.fonts.heading}', sans-serif` : undefined
  }
})

const productNameClasses = computed(() => {
  const size = props.theme?.layout?.headerStyle
  return size === 'compact' ? 'text-sm' : size === 'large' ? 'text-lg' : 'text-base'
})

// Product description styling
const productDescriptionStyle = computed(() => {
  return {
    color: props.theme?.colors?.bodyText || '#6B7280',
    fontFamily: props.theme?.fonts?.body ? `'${props.theme.fonts.body}', sans-serif` : undefined
  }
})

// Rating text styling
const ratingTextStyle = computed(() => {
  return {
    color: props.theme?.colors?.bodyText || '#6B7280'
  }
})

// Price styling
const originalPriceStyle = computed(() => {
  return {
    color: props.theme?.colors?.bodyText || '#9CA3AF'
  }
})

const currentPriceStyle = computed(() => {
  return {
    color: props.theme?.colors?.heading || '#1F2937',
    fontFamily: props.theme?.fonts?.heading ? `'${props.theme.fonts.heading}', sans-serif` : undefined
  }
})

const currentPriceClasses = computed(() => {
  const size = props.theme?.layout?.headerStyle
  return size === 'compact' ? 'text-lg' : size === 'large' ? 'text-2xl' : 'text-xl'
})

const fromTextStyle = computed(() => {
  return {
    color: props.theme?.colors?.bodyText || '#6B7280',
    fontStyle: 'italic'
  }
})

// Button styling
const quickViewButtonStyle = computed(() => {
  return {
    backgroundColor: props.theme?.colors?.primary || '#10B981',
    color: props.theme?.colors?.background || 'white',
    borderRadius: getBorderRadiusValue(props.theme?.layout?.borderStyle || 'rounded')
  }
})

const addToCartButtonStyle = computed(() => {
  return {
    backgroundColor: props.theme?.colors?.primary || '#10B981',
    color: props.theme?.colors?.background || 'white',
    borderRadius: getBorderRadiusValue(props.theme?.layout?.borderStyle || 'rounded')
  }
})

// Show add to cart button based on config
const showAddToCartButton = computed(() => {
  return props.config?.style === 'dynamic'
})

// =============== Methods ===============

function onHover() {
  isHovered.value = true
  showQuickView.value = true

  // Start image rotation if multiple images and hover effect is enabled
  if (allImages.value.length > 1 && props.config?.hoverEffect !== 'none') {
    startImageRotation()
  }
}

function onLeave() {
  isHovered.value = false
  showQuickView.value = false

  // Stop image rotation and reset to first image
  stopImageRotation()
  currentImage.value = allImages.value[0] || ''
  hoverImageIndex = 0
}

function startImageRotation() {
  if (imageRotationTimer) clearInterval(imageRotationTimer)

  imageRotationTimer = setInterval(() => {
    hoverImageIndex = (hoverImageIndex + 1) % allImages.value.length
    currentImage.value = allImages.value[hoverImageIndex]
  }, 1500) // Change image every 1.5 seconds
}

function stopImageRotation() {
  if (imageRotationTimer) {
    clearInterval(imageRotationTimer)
    imageRotationTimer = null
  }
}

function formatDiscount(discount: { type: string; value: number }): string {
  return discount.type === 'percentage'
    ? `${discount.value}% OFF`
    : `$${discount.value} OFF`
}

function openQuickView() {
  // TODO: Implement quick view modal
  console.log('Opening quick view for:', props.product.name)
}

function addToCart() {
  // TODO: Implement add to cart functionality
  console.log('Adding to cart:', props.product.name)
}

// =============== Helper Functions ===============

function getBorderRadiusValue(borderStyle: string): string {
  switch (borderStyle) {
    case 'sharp': return '0px'
    case 'rounded': return '8px'
    case 'pill': return '50px'
    default: return '8px'
  }
}
</script>

<style scoped>
/* Product card base styles */
.product-card {
  background-color: var(--color-background, #ffffff);
  border: 1px solid var(--color-border, #E5E7EB);
  border-radius: var(--border-radius, 8px);
  transition: all 0.3s ease;
}

.product-card:hover {
  border-color: var(--color-primary, #10B981);
}

/* Aspect ratio utilities */
.aspect-square {
  aspect-ratio: 1 / 1;
}

.aspect-\[3\/4\] {
  aspect-ratio: 3 / 4;
}

.aspect-\[4\/3\] {
  aspect-ratio: 4 / 3;
}

.aspect-\[16\/9\] {
  aspect-ratio: 16 / 9;
}

/* Text truncation */
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

/* Hover effects */
.group:hover .group-hover\:scale-110 {
  transform: scale(1.1);
}

.group:hover .group-hover\:opacity-100 {
  opacity: 1;
}

/* Button hover effects */
button:hover:not(:disabled) {
  opacity: 0.9;
}

button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

/* Loading and transition effects */
img {
  transition: transform 0.5s ease, opacity 0.3s ease;
}

.group:hover img {
  transform: scale(1.05);
}

/* Focus states */
.product-card:focus {
  outline: 2px solid var(--color-primary, #10B981);
  outline-offset: 2px;
}

/* Mobile optimizations */
@media (max-width: 640px) {
  .product-card {
    padding: 0.75rem;
  }

  .aspect-square {
    aspect-ratio: 1 / 1;
  }
}
</style>
