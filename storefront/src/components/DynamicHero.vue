<template>
  <section
    class="relative overflow-hidden"
    :class="heroClasses"
    :style="heroStyle"
  >
    <!-- Background Image/Banner -->
    <div
      v-if="shop?.banner"
      class="absolute inset-0 z-0"
      :style="backgroundImageStyle"
    >
      <div class="absolute inset-0" :style="overlayStyle"></div>
    </div>

    <!-- Hero Content -->
    <div class="relative z-10" :class="containerClasses">
      <div class="flex items-center justify-center min-h-full">
        <div class="text-center space-y-6 max-w-4xl mx-auto">

          <!-- Shop Logo (if different from banner) -->
          <div v-if="shop?.image && shop.image !== shop.banner" class="flex justify-center">
            <img
              :src="shop.image"
              :alt="shop.name + ' Logo'"
              class="object-contain transition-transform hover:scale-105"
              :class="logoClasses"
              :style="logoStyle"
            />
          </div>

          <!-- Main Heading -->
          <h1
            class="font-bold tracking-tight leading-tight"
            :style="mainHeadingStyle"
            :class="mainHeadingClasses"
          >
            {{ heroTitle }}
          </h1>

          <!-- Subtitle/Description -->
          <p
            v-if="heroDescription"
            class="max-w-2xl mx-auto leading-relaxed"
            :style="subtitleStyle"
            :class="subtitleClasses"
          >
            {{ heroDescription }}
          </p>

          <!-- Call-to-Action Buttons -->
          <div class="flex flex-col sm:flex-row items-center justify-center gap-4">
            <router-link
              :to="{ path: '/products' }"
              class="inline-block px-8 py-4 font-semibold uppercase tracking-wide transition-all duration-200 hover:scale-105 hover:shadow-lg"
              :style="primaryButtonStyle"
              :class="buttonClasses"
            >
              Shop Now
            </router-link>

            <router-link
              v-if="shop?.isVerified"
              :to="{ path: '/collections' }"
              class="inline-block px-8 py-4 font-semibold uppercase tracking-wide transition-all duration-200 hover:scale-105"
              :style="secondaryButtonStyle"
              :class="buttonClasses"
            >
              Browse Collections
            </router-link>
          </div>

          <!-- Shop Stats/Analytics (if available) -->
          <div
            v-if="analytics && showAnalytics"
            class="flex flex-wrap items-center justify-center gap-8 pt-8"
          >
            <div
              v-if="analytics.customerCount > 0"
              class="text-center space-y-1"
            >
              <div
                class="text-2xl font-bold"
                :style="statNumberStyle"
              >
                {{ formatNumber(analytics.customerCount) }}+
              </div>
              <div
                class="text-sm uppercase tracking-wide"
                :style="statLabelStyle"
              >
                Happy Customers
              </div>
            </div>

            <div
              v-if="analytics.productCount > 0"
              class="text-center space-y-1"
            >
              <div
                class="text-2xl font-bold"
                :style="statNumberStyle"
              >
                {{ formatNumber(analytics.productCount) }}+
              </div>
              <div
                class="text-sm uppercase tracking-wide"
                :style="statLabelStyle"
              >
                Products
              </div>
            </div>

            <div
              v-if="analytics.revenue > 0"
              class="text-center space-y-1"
            >
              <div
                class="text-2xl font-bold"
                :style="statNumberStyle"
              >
                ${{ formatNumber(analytics.revenue) }}+
              </div>
              <div
                class="text-sm uppercase tracking-wide"
                :style="statLabelStyle"
              >
                Revenue
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Scroll Indicator -->
    <div
      v-if="showScrollIndicator"
      class="absolute bottom-8 left-1/2 transform -translate-x-1/2 animate-bounce"
    >
      <svg
        class="w-6 h-6"
        :style="scrollIndicatorStyle"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
      >
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3"></path>
      </svg>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { DynamicTheme, ShopInfo, LayoutConfig, AnalyticsInfo } from '@/services/dynamic-theme'

// =============== Props ===============

interface Props {
  shop?: ShopInfo | null
  theme?: DynamicTheme | null
  layout?: LayoutConfig | null
  analytics?: AnalyticsInfo | null
}

const props = defineProps<Props>()

// =============== Computed Properties ===============

// Hero content
const heroTitle = computed(() => {
  return props.shop?.name || 'Welcome to Our Store'
})

const heroDescription = computed(() => {
  return props.shop?.description || 'Discover amazing products and exceptional quality.'
})

const showAnalytics = computed(() => {
  return props.analytics && (
    props.analytics.customerCount > 0 ||
    props.analytics.productCount > 0 ||
    props.analytics.revenue > 0
  )
})

const showScrollIndicator = computed(() => {
  return props.layout?.headerStyle !== 'compact'
})

// =============== Computed Styles ===============

// Hero section styling
const heroStyle = computed(() => {
  if (!props.theme) return { minHeight: '500px', backgroundColor: '#F9FAFB' }

  const style: any = {
    backgroundColor: props.theme.colors?.background || '#F9FAFB',
    color: props.theme.colors?.bodyText || '#6B7280'
  }

  // Set height based on layout style
  if (props.layout?.headerStyle === 'compact') {
    style.minHeight = '400px'
  } else if (props.layout?.headerStyle === 'large') {
    style.minHeight = '700px'
  } else {
    style.minHeight = '500px'
  }

  return style
})

const heroClasses = computed(() => {
  const classes = ['flex', 'items-center', 'justify-center']

  // Add spacing based on layout
  if (props.layout?.spacing === 'tight') {
    classes.push('py-12')
  } else if (props.layout?.spacing === 'loose') {
    classes.push('py-24')
  } else {
    classes.push('py-16')
  }

  return classes
})

// Background image styling
const backgroundImageStyle = computed(() => {
  if (!props.shop?.banner) return {}

  return {
    backgroundImage: `url(${props.shop.banner})`,
    backgroundSize: 'cover',
    backgroundPosition: 'center',
    backgroundRepeat: 'no-repeat'
  }
})

// Overlay for better text readability
const overlayStyle = computed(() => {
  if (!props.shop?.banner) return {}

  // Create overlay based on theme colors
  const backgroundColor = props.theme?.colors?.background || '#000000'
  return {
    backgroundColor: `${backgroundColor}40` // 40% opacity
  }
})

// Container classes
const containerClasses = computed(() => {
  const classes = ['mx-auto', 'px-4', 'sm:px-6', 'w-full', 'h-full']

  if (props.layout?.containerWidth === 'full') {
    classes.push('max-w-none')
  } else if (props.layout?.containerWidth === 'wide') {
    classes.push('max-w-7xl')
  } else {
    classes.push('max-w-6xl')
  }

  return classes
})

// Logo styling
const logoClasses = computed(() => {
  switch (props.layout?.headerStyle) {
    case 'compact': return 'h-16 w-16'
    case 'large': return 'h-24 w-24'
    default: return 'h-20 w-20'
  }
})

const logoStyle = computed(() => {
  return {
    filter: props.shop?.banner ? 'drop-shadow(0 2px 4px rgba(0,0,0,0.3))' : 'none'
  }
})

// Main heading styling
const mainHeadingStyle = computed(() => {
  if (!props.theme) return { color: '#1F2937' }

  const style: any = {
    color: props.theme.colors?.heading || '#1F2937',
    fontFamily: props.theme.fonts?.heading ? `'${props.theme.fonts.heading}', sans-serif` : undefined
  }

  // Add text shadow if there's a background image
  if (props.shop?.banner) {
    style.textShadow = '0 2px 4px rgba(0,0,0,0.5)'
  }

  return style
})

const mainHeadingClasses = computed(() => {
  switch (props.layout?.headerStyle) {
    case 'compact': return 'text-3xl md:text-4xl'
    case 'large': return 'text-5xl md:text-7xl'
    default: return 'text-4xl md:text-6xl'
  }
})

// Subtitle styling
const subtitleStyle = computed(() => {
  if (!props.theme) return { color: '#6B7280' }

  const style: any = {
    color: props.theme.colors?.bodyText || '#6B7280',
    fontFamily: props.theme.fonts?.body ? `'${props.theme.fonts.body}', sans-serif` : undefined
  }

  // Add text shadow if there's a background image
  if (props.shop?.banner) {
    style.textShadow = '0 1px 2px rgba(0,0,0,0.5)'
  }

  return style
})

const subtitleClasses = computed(() => {
  switch (props.layout?.headerStyle) {
    case 'compact': return 'text-base md:text-lg'
    case 'large': return 'text-xl md:text-2xl'
    default: return 'text-lg md:text-xl'
  }
})

// Button styling
const primaryButtonStyle = computed(() => {
  if (!props.theme) return { backgroundColor: '#10B981', color: 'white' }

  return {
    backgroundColor: props.theme.colors?.primary || '#10B981',
    color: props.theme.colors?.background || 'white',
    borderRadius: getBorderRadiusValue(props.theme.layout?.borderStyle || 'rounded'),
    boxShadow: props.shop?.banner ? '0 4px 12px rgba(0,0,0,0.2)' : '0 2px 8px rgba(0,0,0,0.1)'
  }
})

const secondaryButtonStyle = computed(() => {
  if (!props.theme) return { backgroundColor: 'transparent', color: '#10B981', border: '2px solid #10B981' }

  const primaryColor = props.theme.colors?.primary || '#10B981'
  return {
    backgroundColor: 'transparent',
    color: primaryColor,
    border: `2px solid ${primaryColor}`,
    borderRadius: getBorderRadiusValue(props.theme.layout?.borderStyle || 'rounded'),
    boxShadow: props.shop?.banner ? '0 4px 12px rgba(0,0,0,0.2)' : '0 2px 8px rgba(0,0,0,0.1)'
  }
})

const buttonClasses = computed(() => {
  const style = props.theme?.layout?.borderStyle || 'rounded'
  return style === 'sharp' ? 'rounded-none' : style === 'pill' ? 'rounded-full' : 'rounded-lg'
})

// Analytics stats styling
const statNumberStyle = computed(() => {
  if (!props.theme) return { color: '#1F2937' }

  const style: any = {
    color: props.theme.colors?.primary || '#1F2937',
    fontFamily: props.theme.fonts?.heading ? `'${props.theme.fonts.heading}', sans-serif` : undefined
  }

  if (props.shop?.banner) {
    style.textShadow = '0 2px 4px rgba(0,0,0,0.5)'
  }

  return style
})

const statLabelStyle = computed(() => {
  if (!props.theme) return { color: '#6B7280' }

  const style: any = {
    color: props.theme.colors?.bodyText || '#6B7280',
    fontFamily: props.theme.fonts?.body ? `'${props.theme.fonts.body}', sans-serif` : undefined
  }

  if (props.shop?.banner) {
    style.textShadow = '0 1px 2px rgba(0,0,0,0.5)'
  }

  return style
})

// Scroll indicator styling
const scrollIndicatorStyle = computed(() => {
  return {
    color: props.theme?.colors?.primary || '#10B981'
  }
})

// =============== Methods ===============

function formatNumber(num: number): string {
  if (num >= 1000000) {
    return (num / 1000000).toFixed(1) + 'M'
  } else if (num >= 1000) {
    return (num / 1000).toFixed(1) + 'K'
  }
  return num.toString()
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
/* Hero section base styles */
.hero-section {
  position: relative;
  overflow: hidden;
}

/* Background image overlay */
.hero-overlay {
  background: linear-gradient(135deg, rgba(0,0,0,0.4) 0%, rgba(0,0,0,0.2) 100%);
}

/* Animation for scroll indicator */
@keyframes bounce {
  0%, 20%, 53%, 80%, 100% {
    transform: translate3d(0,0,0);
  }
  40%, 43% {
    transform: translate3d(0,-8px,0);
  }
  70% {
    transform: translate3d(0,-4px,0);
  }
  90% {
    transform: translate3d(0,-2px,0);
  }
}

.animate-bounce {
  animation: bounce 2s infinite;
}

/* Button hover effects */
.hover\:scale-105:hover {
  transform: scale(1.05);
}

.hover\:shadow-lg:hover {
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
}

/* Text shadow utilities for better readability over images */
.text-shadow-sm {
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.5);
}

.text-shadow-md {
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.5);
}

/* Responsive adjustments */
@media (max-width: 640px) {
  .hero-section {
    min-height: 400px !important;
  }

  .hero-content {
    padding: 2rem 1rem;
  }
}

@media (max-width: 768px) {
  .hero-stats {
    flex-direction: column;
    gap: 1rem;
  }
}

/* Focus states for accessibility */
button:focus,
a:focus {
  outline: 2px solid var(--color-primary, #10B981);
  outline-offset: 2px;
}

/* Loading states */
.hero-loading {
  background: linear-gradient(90deg, #f0f0f0 25%, #e0e0e0 50%, #f0f0f0 75%);
  background-size: 200% 100%;
  animation: loading 1.5s infinite;
}

@keyframes loading {
  0% {
    background-position: 200% 0;
  }
  100% {
    background-position: -200% 0;
  }
}

/* Smooth transitions */
* {
  transition: color 0.3s ease, background-color 0.3s ease, border-color 0.3s ease;
}
</style>
