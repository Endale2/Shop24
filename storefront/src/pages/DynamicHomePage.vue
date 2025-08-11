<template>
  <div class="space-y-12">
    
    <!-- Dynamic Hero Section -->
    <DynamicHero 
      :shop="storefrontConfig?.shop"
      :theme="theme"
      :layout="storefrontConfig?.layout"
      v-if="storefrontConfig"
    />
    
    <!-- Featured Products Section -->
    <section class="space-y-8">
      <div class="text-center space-y-4">
        <h2 
          class="font-bold tracking-tight uppercase"
          :style="sectionHeadingStyle"
          :class="sectionHeadingClasses"
        >
          Featured Products
        </h2>
        <p 
          class="max-w-2xl mx-auto"
          :style="sectionDescriptionStyle"
        >
          Discover our handpicked selection of premium products
        </p>
      </div>
      
      <!-- Loading State -->
      <div v-if="isLoadingProducts" class="text-center py-12">
        <div class="animate-spin rounded-full h-8 w-8 border-b-2 mx-auto mb-4" :style="spinnerStyle"></div>
        <p :style="textStyle">Loading products...</p>
      </div>
      
      <!-- Dynamic Product Grid -->
      <div 
        v-else-if="featuredProducts.length > 0" 
        class="grid gap-6"
        :class="productGridClasses"
        :style="gridStyle"
      >
        <DynamicProductCard
          v-for="product in featuredProducts"
          :key="product.id"
          :product="product"
          :theme="theme"
          :config="storefrontConfig?.components?.productCard"
        />
      </div>
      
      <!-- Empty State -->
      <div v-else class="text-center py-12 space-y-4">
        <div class="w-16 h-16 mx-auto rounded-full flex items-center justify-center" :style="emptyStateIconStyle">
          <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2M4 13h2m5 0h2m5 0h2"></path>
          </svg>
        </div>
        <h3 :style="headingStyle" class="text-lg font-semibold">No Products Available</h3>
        <p :style="textStyle">Check back soon for new arrivals!</p>
      </div>
      
      <!-- View All Products Button -->
      <div v-if="featuredProducts.length > 0" class="text-center">
        <router-link
          :to="{ path: '/products' }"
          class="inline-block px-8 py-3 font-semibold uppercase tracking-wide transition-all duration-200 hover:scale-105 hover:shadow-lg"
          :style="primaryButtonStyle"
          :class="buttonClasses"
        >
          See All Products
        </router-link>
      </div>
    </section>
    
    <!-- Collections Preview Section (if available) -->
    <section v-if="featuredCollections.length > 0" class="space-y-8">
      <div class="text-center space-y-4">
        <h2 
          class="font-bold tracking-tight uppercase"
          :style="sectionHeadingStyle"
          :class="sectionHeadingClasses"
        >
          Shop by Collection
        </h2>
        <p 
          class="max-w-2xl mx-auto"
          :style="sectionDescriptionStyle"
        >
          Explore our curated collections
        </p>
      </div>
      
      <div 
        class="grid gap-6"
        :class="collectionGridClasses"
      >
        <DynamicCollectionCard
          v-for="collection in featuredCollections"
          :key="collection.id"
          :collection="collection"
          :theme="theme"
        />
      </div>
    </section>
    
    <!-- Store Features Section -->
    <section class="py-12 rounded-lg" :style="featuresBackgroundStyle">
      <div class="text-center space-y-8">
        <h2 
          class="font-bold tracking-tight uppercase"
          :style="sectionHeadingStyle"
          :class="sectionHeadingClasses"
        >
          Why Shop With Us
        </h2>
        
        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div 
            v-for="feature in storeFeatures"
            :key="feature.title"
            class="text-center space-y-4"
          >
            <div 
              class="w-12 h-12 mx-auto rounded-full flex items-center justify-center"
              :style="featureIconStyle"
            >
              <component :is="feature.icon" class="w-6 h-6" />
            </div>
            <h3 :style="headingStyle" class="font-semibold">{{ feature.title }}</h3>
            <p :style="textStyle" class="text-sm max-w-sm mx-auto">{{ feature.description }}</p>
          </div>
        </div>
      </div>
    </section>
    
    <!-- Newsletter Signup Section -->
    <section v-if="showNewsletter" class="py-12 text-center space-y-6 rounded-lg" :style="newsletterBackgroundStyle">
      <div class="space-y-4">
        <h2 
          class="font-bold tracking-tight"
          :style="sectionHeadingStyle"
          :class="sectionHeadingClasses"
        >
          Stay in the Loop
        </h2>
        <p 
          class="max-w-md mx-auto"
          :style="sectionDescriptionStyle"
        >
          Subscribe to get special offers, updates, and exclusive content.
        </p>
      </div>
      
      <form @submit.prevent="subscribeToNewsletter" class="max-w-md mx-auto">
        <div class="flex gap-3">
          <input
            type="email"
            v-model="newsletterEmail"
            placeholder="Enter your email"
            class="flex-1 px-4 py-3 border rounded-lg focus:outline-none focus:ring-2 transition-all"
            :style="inputStyle"
            required
          />
          <button
            type="submit"
            class="px-6 py-3 font-medium rounded-lg transition-all duration-200 hover:scale-105"
            :style="primaryButtonStyle"
            :disabled="isSubscribing"
          >
            {{ isSubscribing ? 'Subscribing...' : 'Subscribe' }}
          </button>
        </div>
      </form>
    </section>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import DynamicHero from '@/components/DynamicHero.vue'
import DynamicProductCard from '@/components/DynamicProductCard.vue'
import DynamicCollectionCard from '@/components/DynamicCollectionCard.vue'
import { getCurrentShopSlug } from '@/services/shop'
import { fetchAllProducts } from '@/services/product'
import { fetchCollections } from '@/services/collections'
import type { StorefrontConfig, DynamicTheme } from '@/services/dynamic-theme'
import type { Product } from '@/services/product'
import type { Collection } from '@/services/collections'

// =============== Props ===============

interface Props {
  storefrontConfig?: StorefrontConfig | null
  theme?: DynamicTheme | null
  components?: any
}

const props = defineProps<Props>()

// =============== Reactive State ===============

const featuredProducts = ref<Product[]>([])
const featuredCollections = ref<Collection[]>([])
const isLoadingProducts = ref(true)
const newsletterEmail = ref('')
const isSubscribing = ref(false)

// Configuration
const maxFeaturedProducts = 8
const maxFeaturedCollections = 3
const showNewsletter = true

// Store features
const storeFeatures = [
  {
    title: 'Free Shipping',
    description: 'Free shipping on orders over $50. Fast and reliable delivery.',
    icon: 'TruckIcon'
  },
  {
    title: 'Secure Payment',
    description: 'Your payment information is protected with industry-standard encryption.',
    icon: 'ShieldCheckIcon'
  },
  {
    title: '24/7 Support',
    description: 'Get help when you need it with our round-the-clock customer support.',
    icon: 'ChatBubbleLeftRightIcon'
  }
]

// =============== Computed Styles ===============

// Section heading styles
const sectionHeadingStyle = computed(() => {
  if (!props.theme) return { color: '#1F2937' }
  
  return {
    color: props.theme.colors?.heading || '#1F2937',
    fontFamily: props.theme.fonts?.heading ? `'${props.theme.fonts.heading}', sans-serif` : undefined
  }
})

const sectionHeadingClasses = computed(() => {
  const size = props.theme?.layout?.headerStyle
  return size === 'compact' ? 'text-2xl' : size === 'large' ? 'text-4xl' : 'text-3xl'
})

const sectionDescriptionStyle = computed(() => {
  if (!props.theme) return { color: '#6B7280' }
  
  return {
    color: props.theme.colors?.bodyText || '#6B7280',
    fontFamily: props.theme.fonts?.body ? `'${props.theme.fonts.body}', sans-serif` : undefined
  }
})

const headingStyle = computed(() => {
  if (!props.theme) return { color: '#1F2937' }
  
  return {
    color: props.theme.colors?.heading || '#1F2937',
    fontFamily: props.theme.fonts?.heading ? `'${props.theme.fonts.heading}', sans-serif` : undefined
  }
})

const textStyle = computed(() => {
  if (!props.theme) return { color: '#6B7280' }
  
  return {
    color: props.theme.colors?.bodyText || '#6B7280',
    fontFamily: props.theme.fonts?.body ? `'${props.theme.fonts.body}', sans-serif` : undefined
  }
})

// Product grid styles
const productGridClasses = computed(() => {
  const columns = props.theme?.layout?.gridColumns || '3'
  
  return [
    'grid-cols-1',
    'sm:grid-cols-2',
    `lg:grid-cols-${columns}`,
    'xl:grid-cols-4'
  ]
})

const gridStyle = computed(() => {
  return {
    '--grid-columns': props.theme?.layout?.gridColumns || '3'
  }
})

const collectionGridClasses = computed(() => {
  return ['grid-cols-1', 'md:grid-cols-2', 'lg:grid-cols-3']
})

// Button styles
const primaryButtonStyle = computed(() => {
  if (!props.theme) return { backgroundColor: '#10B981', color: 'white' }
  
  return {
    backgroundColor: props.theme.colors?.primary || '#10B981',
    color: props.theme.colors?.background || 'white',
    borderRadius: props.theme.layout?.borderStyle === 'sharp' ? '0px' : 
                 props.theme.layout?.borderStyle === 'pill' ? '50px' : '8px'
  }
})

const buttonClasses = computed(() => {
  const style = props.theme?.layout?.borderStyle || 'rounded'
  return style === 'sharp' ? 'rounded-none' : style === 'pill' ? 'rounded-full' : 'rounded-lg'
})

// Input styles
const inputStyle = computed(() => {
  if (!props.theme) return {}
  
  return {
    borderColor: props.theme.colors?.border || '#E5E7EB',
    backgroundColor: props.theme.colors?.background || '#ffffff',
    color: props.theme.colors?.bodyText || '#6B7280',
    '--tw-ring-color': props.theme.colors?.primary || '#10B981'
  }
})

// Background styles for sections
const featuresBackgroundStyle = computed(() => {
  if (!props.theme) return { backgroundColor: '#F9FAFB' }
  
  const primaryColor = props.theme.colors?.primary || '#10B981'
  return {
    backgroundColor: `${primaryColor}05` // 5% opacity
  }
})

const newsletterBackgroundStyle = computed(() => {
  if (!props.theme) return { backgroundColor: '#F3F4F6' }
  
  const primaryColor = props.theme.colors?.primary || '#10B981'
  return {
    backgroundColor: `${primaryColor}08` // 8% opacity
  }
})

// Icon styles
const featureIconStyle = computed(() => {
  if (!props.theme) return { backgroundColor: '#10B981', color: 'white' }
  
  return {
    backgroundColor: props.theme.colors?.primary || '#10B981',
    color: props.theme.colors?.background || 'white'
  }
})

const emptyStateIconStyle = computed(() => {
  if (!props.theme) return { backgroundColor: '#F3F4F6', color: '#9CA3AF' }
  
  return {
    backgroundColor: `${props.theme.colors?.primary || '#10B981'}10`,
    color: props.theme.colors?.bodyText || '#9CA3AF'
  }
})

const spinnerStyle = computed(() => {
  return {
    borderBottomColor: props.theme?.colors?.primary || '#10B981'
  }
})

// =============== Lifecycle ===============

onMounted(async () => {
  await loadHomePageData()
})

// =============== Methods ===============

async function loadHomePageData() {
  const shopSlug = getCurrentShopSlug()
  if (!shopSlug) return
  
  try {
    isLoadingProducts.value = true
    
    // Load products and collections in parallel
    const [products, collections] = await Promise.all([
      fetchAllProducts(shopSlug),
      fetchCollections(shopSlug).catch(() => []) // Collections are optional
    ])
    
    // Get featured products (first N products)
    featuredProducts.value = products.slice(0, maxFeaturedProducts)
    
    // Get featured collections
    featuredCollections.value = collections.slice(0, maxFeaturedCollections)
    
  } catch (error) {
    console.error('Error loading home page data:', error)
  } finally {
    isLoadingProducts.value = false
  }
}

async function subscribeToNewsletter() {
  if (!newsletterEmail.value) return
  
  isSubscribing.value = true
  
  try {
    // TODO: Implement newsletter subscription API
    await new Promise(resolve => setTimeout(resolve, 1000)) // Simulate API call
    
    console.log('Newsletter subscription:', newsletterEmail.value)
    newsletterEmail.value = ''
    
    // Show success notification
    showNotification('Successfully subscribed to newsletter!', 'success')
    
  } catch (error) {
    console.error('Newsletter subscription error:', error)
    showNotification('Failed to subscribe. Please try again.', 'error')
  } finally {
    isSubscribing.value = false
  }
}

function showNotification(message: string, type: 'success' | 'error') {
  // Simple notification system
  const notification = document.createElement('div')
  const bgColor = type === 'success' ? '#10B981' : '#EF4444'
  
  notification.className = 'fixed top-4 right-4 text-white px-4 py-2 rounded-lg shadow-lg z-50 transition-all duration-300'
  notification.style.backgroundColor = bgColor
  notification.textContent = message
  
  document.body.appendChild(notification)
  
  setTimeout(() => {
    notification.style.opacity = '0'
    setTimeout(() => {
      if (document.body.contains(notification)) {
        document.body.removeChild(notification)
      }
    }, 300)
  }, 3000)
}
</script>

<style scoped>
/* Dynamic grid responsive classes */
.grid-cols-1 { grid-template-columns: repeat(1, minmax(0, 1fr)); }
.sm\:grid-cols-2 { grid-template-columns: repeat(1, minmax(0, 1fr)); }
.lg\:grid-cols-3 { grid-template-columns: repeat(1, minmax(0, 1fr)); }
.xl\:grid-cols-4 { grid-template-columns: repeat(1, minmax(0, 1fr)); }

@media (min-width: 640px) {
  .sm\:grid-cols-2 { grid-template-columns: repeat(2, minmax(0, 1fr)); }
}

@media (min-width: 1024px) {
  .lg\:grid-cols-3 { grid-template-columns: repeat(3, minmax(0, 1fr)); }
}

@media (min-width: 1280px) {
  .xl\:grid-cols-4 { grid-template-columns: repeat(4, minmax(0, 1fr)); }
}

/* Loading animation */
@keyframes spin {
  to { transform: rotate(360deg); }
}

.animate-spin {
  animation: spin 1s linear infinite;
}

/* Hover effects */
.hover\:scale-105:hover {
  transform: scale(1.05);
}

.hover\:shadow-lg:hover {
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
}

/* Focus effects */
input:focus {
  ring-width: 2px;
  ring-color: var(--tw-ring-color);
}

/* Button states */
button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

button:disabled:hover {
  transform: none;
}
</style>
