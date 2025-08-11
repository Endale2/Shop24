<template>
  <div class="space-y-8">
    
    <!-- Dynamic Breadcrumbs -->
    <nav class="flex items-center space-x-2 text-sm" :style="breadcrumbStyle">
      <router-link 
        :to="{ path: '/' }" 
        class="hover:opacity-70 transition-opacity"
        :style="linkStyle"
      >
        Home
      </router-link>
      <svg class="w-4 h-4" :style="separatorStyle" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
      </svg>
      <span :style="currentPageStyle">Collections</span>
    </nav>

    <!-- Page Header -->
    <div class="text-center space-y-4">
      <h1 
        class="font-bold tracking-tight uppercase"
        :style="pageHeadingStyle"
        :class="pageHeadingClasses"
      >
        Shop by Collection
      </h1>
      <p 
        class="max-w-2xl mx-auto"
        :style="pageDescriptionStyle"
      >
        Explore our carefully curated collections designed to match your style and needs.
      </p>
    </div>

    <!-- Loading State -->
    <div v-if="isLoading" class="text-center py-12">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 mx-auto mb-4" :style="spinnerStyle"></div>
      <p :style="textStyle">Loading collections...</p>
    </div>
    
    <!-- Collections Grid -->
    <div 
      v-else-if="collections.length > 0" 
      class="grid gap-8"
      :class="collectionGridClasses"
    >
      <DynamicCollectionCard
        v-for="collection in collections"
        :key="collection.id"
        :collection="collection"
        :theme="theme"
      />
    </div>
    
    <!-- Empty State -->
    <div v-else class="text-center py-16 space-y-6">
      <div class="w-20 h-20 mx-auto rounded-full flex items-center justify-center" :style="emptyStateIconStyle">
        <svg class="w-10 h-10" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"></path>
        </svg>
      </div>
      <div class="space-y-2">
        <h3 :style="headingStyle" class="text-xl font-semibold">No Collections Available</h3>
        <p :style="textStyle">Collections will appear here once they're created.</p>
      </div>
      <router-link
        :to="{ path: '/products' }"
        class="inline-block px-6 py-3 font-medium rounded-lg transition-all duration-200 hover:scale-105"
        :style="primaryButtonStyle"
      >
        Browse All Products
      </router-link>
    </div>

    <!-- Collection Stats -->
    <div v-if="collections.length > 0" class="text-center py-8 space-y-4">
      <div class="flex flex-wrap items-center justify-center gap-8">
        <div class="text-center space-y-1">
          <div class="text-2xl font-bold" :style="statNumberStyle">
            {{ collections.length }}
          </div>
          <div class="text-sm uppercase tracking-wide" :style="statLabelStyle">
            Collections
          </div>
        </div>
        
        <div class="text-center space-y-1">
          <div class="text-2xl font-bold" :style="statNumberStyle">
            {{ totalProducts }}
          </div>
          <div class="text-sm uppercase tracking-wide" :style="statLabelStyle">
            Total Products
          </div>
        </div>
        
        <div v-if="averageProductsPerCollection > 0" class="text-center space-y-1">
          <div class="text-2xl font-bold" :style="statNumberStyle">
            {{ averageProductsPerCollection }}
          </div>
          <div class="text-sm uppercase tracking-wide" :style="statLabelStyle">
            Avg per Collection
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import DynamicCollectionCard from '@/components/DynamicCollectionCard.vue'
import { getCurrentShopSlug } from '@/services/shop'
import { fetchCollections } from '@/services/collections'
import type { Collection } from '@/services/collections'
import type { DynamicTheme } from '@/services/dynamic-theme'

// =============== Props ===============

interface Props {
  storefrontConfig?: any
  theme?: DynamicTheme | null
  components?: any
  layout?: any
}

const props = defineProps<Props>()

// =============== Reactive State ===============

const shopSlug = getCurrentShopSlug()
const collections = ref<Collection[]>([])
const isLoading = ref(true)

// =============== Computed Properties ===============

const totalProducts = computed(() => {
  return collections.value.reduce((total, collection) => total + (collection.product_count || 0), 0)
})

const averageProductsPerCollection = computed(() => {
  if (collections.value.length === 0) return 0
  return Math.round(totalProducts.value / collections.value.length)
})

// Grid classes based on theme configuration
const collectionGridClasses = computed(() => {
  const columns = props.layout?.gridColumns || '3'
  
  return [
    'grid-cols-1',
    'md:grid-cols-2',
    `lg:grid-cols-${Math.min(parseInt(columns), 3)}` // Max 3 columns for collections
  ]
})

// =============== Computed Styles ===============

// Breadcrumb styling
const breadcrumbStyle = computed(() => {
  return {
    fontFamily: props.theme?.fonts?.body ? `'${props.theme.fonts.body}', sans-serif` : undefined
  }
})

const linkStyle = computed(() => {
  return {
    color: props.theme?.colors?.primary || '#10B981'
  }
})

const separatorStyle = computed(() => {
  return {
    color: props.theme?.colors?.bodyText || '#6B7280'
  }
})

const currentPageStyle = computed(() => {
  return {
    color: props.theme?.colors?.bodyText || '#6B7280',
    fontWeight: '500'
  }
})

// Page heading styles
const pageHeadingStyle = computed(() => {
  return {
    color: props.theme?.colors?.heading || '#1F2937',
    fontFamily: props.theme?.fonts?.heading ? `'${props.theme.fonts.heading}', sans-serif` : undefined
  }
})

const pageHeadingClasses = computed(() => {
  const size = props.theme?.layout?.headerStyle
  return size === 'compact' ? 'text-2xl md:text-3xl' : size === 'large' ? 'text-4xl md:text-5xl' : 'text-3xl md:text-4xl'
})

const pageDescriptionStyle = computed(() => {
  return {
    color: props.theme?.colors?.bodyText || '#6B7280',
    fontFamily: props.theme?.fonts?.body ? `'${props.theme.fonts.body}', sans-serif` : undefined
  }
})

// Button styles
const primaryButtonStyle = computed(() => {
  return {
    backgroundColor: props.theme?.colors?.primary || '#10B981',
    color: props.theme?.colors?.background || 'white',
    borderRadius: getBorderRadiusValue(props.theme?.layout?.borderStyle || 'rounded')
  }
})

// Other styles
const textStyle = computed(() => {
  return {
    color: props.theme?.colors?.bodyText || '#6B7280',
    fontFamily: props.theme?.fonts?.body ? `'${props.theme.fonts.body}', sans-serif` : undefined
  }
})

const headingStyle = computed(() => {
  return {
    color: props.theme?.colors?.heading || '#1F2937',
    fontFamily: props.theme?.fonts?.heading ? `'${props.theme.fonts.heading}', sans-serif` : undefined
  }
})

const spinnerStyle = computed(() => {
  return {
    borderBottomColor: props.theme?.colors?.primary || '#10B981'
  }
})

const emptyStateIconStyle = computed(() => {
  return {
    backgroundColor: `${props.theme?.colors?.primary || '#10B981'}10`,
    color: props.theme?.colors?.bodyText || '#9CA3AF'
  }
})

// Stats styling
const statNumberStyle = computed(() => {
  return {
    color: props.theme?.colors?.primary || '#1F2937',
    fontFamily: props.theme?.fonts?.heading ? `'${props.theme.fonts.heading}', sans-serif` : undefined
  }
})

const statLabelStyle = computed(() => {
  return {
    color: props.theme?.colors?.bodyText || '#6B7280',
    fontFamily: props.theme?.fonts?.body ? `'${props.theme.fonts.body}', sans-serif` : undefined
  }
})

// =============== Methods ===============

async function loadCollections() {
  if (!shopSlug) return
  
  isLoading.value = true
  try {
    collections.value = await fetchCollections(shopSlug)
  } catch (error) {
    console.error('Failed to load collections:', error)
  } finally {
    isLoading.value = false
  }
}

function getBorderRadiusValue(borderStyle: string): string {
  switch (borderStyle) {
    case 'sharp': return '0px'
    case 'rounded': return '8px'
    case 'pill': return '50px'
    default: return '8px'
  }
}

// =============== Lifecycle ===============

onMounted(async () => {
  await loadCollections()
})
</script>

<style scoped>
/* Custom grid responsive classes */
.grid-cols-1 { grid-template-columns: repeat(1, minmax(0, 1fr)); }
.md\:grid-cols-2 { grid-template-columns: repeat(1, minmax(0, 1fr)); }
.lg\:grid-cols-3 { grid-template-columns: repeat(1, minmax(0, 1fr)); }

@media (min-width: 768px) {
  .md\:grid-cols-2 { grid-template-columns: repeat(2, minmax(0, 1fr)); }
}

@media (min-width: 1024px) {
  .lg\:grid-cols-3 { grid-template-columns: repeat(3, minmax(0, 1fr)); }
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

/* Focus states */
button:focus,
a:focus {
  outline: 2px solid var(--color-primary, #10B981);
  outline-offset: 2px;
}
</style>
