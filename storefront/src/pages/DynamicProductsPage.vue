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
      <span :style="currentPageStyle">
        {{ selectedCollection ? getCollectionName(selectedCollection) : 'All Products' }}
      </span>
    </nav>

    <!-- Collection Filter Tabs -->
    <div v-if="collections.length > 0" class="border-b" :style="{ borderBottomColor: theme?.colors?.border || '#E5E7EB' }">
      <nav class="flex space-x-8 overflow-x-auto pb-2">
        <button
          class="flex-shrink-0 py-2 px-1 border-b-2 font-medium text-sm transition-all duration-200"
          :class="selectedCollection === null ? 'border-primary text-primary' : 'border-transparent text-body hover:text-primary hover:border-gray-300'"
          :style="getTabStyle(selectedCollection === null)"
          @click="selectCollection(null)"
        >
          All Products
          <span v-if="totalProducts > 0" class="ml-2 px-2 py-1 text-xs rounded-full" :style="countBadgeStyle">
            {{ totalProducts }}
          </span>
        </button>
        
        <button
          v-for="collection in collections"
          :key="collection.id"
          class="flex-shrink-0 flex items-center space-x-2 py-2 px-1 border-b-2 font-medium text-sm transition-all duration-200"
          :class="selectedCollection === collection.id ? 'border-primary text-primary' : 'border-transparent text-body hover:text-primary hover:border-gray-300'"
          :style="getTabStyle(selectedCollection === collection.id)"
          @click="selectCollection(collection.id)"
        >
          <img 
            v-if="collection.image" 
            :src="collection.image" 
            :alt="collection.title" 
            class="w-5 h-5 object-cover rounded"
          />
          <span>{{ collection.title }}</span>
          <span v-if="(collection as any).product_count > 0" class="px-2 py-1 text-xs rounded-full" :style="countBadgeStyle">
            {{ (collection as any).product_count }}
          </span>
        </button>
      </nav>
    </div>

    <!-- Page Header -->
    <div class="text-center space-y-4">
      <h1 
        class="font-bold tracking-tight uppercase"
        :style="pageHeadingStyle"
        :class="pageHeadingClasses"
      >
        {{ pageTitle }}
      </h1>
      <p 
        v-if="pageDescription"
        class="max-w-2xl mx-auto"
        :style="pageDescriptionStyle"
      >
        {{ pageDescription }}
      </p>
    </div>

    <!-- Filters and Sorting -->
    <div class="flex flex-col md:flex-row md:items-center md:justify-between space-y-4 md:space-y-0">
      <!-- Filter Options -->
      <div class="flex items-center space-x-4">
        <label class="text-sm font-medium" :style="labelStyle">Filter by:</label>
        <select 
          v-model="selectedPriceRange"
          class="px-3 py-2 border rounded-lg text-sm focus:outline-none focus:ring-2 transition-all"
          :style="selectStyle"
        >
          <option value="">All Prices</option>
          <option value="0-25">Under $25</option>
          <option value="25-50">$25 - $50</option>
          <option value="50-100">$50 - $100</option>
          <option value="100+">Over $100</option>
        </select>
      </div>
      
      <!-- Sort Options -->
      <div class="flex items-center space-x-4">
        <label class="text-sm font-medium" :style="labelStyle">Sort by:</label>
        <select 
          v-model="sortBy"
          class="px-3 py-2 border rounded-lg text-sm focus:outline-none focus:ring-2 transition-all"
          :style="selectStyle"
        >
          <option value="name">Name</option>
          <option value="price-low">Price: Low to High</option>
          <option value="price-high">Price: High to Low</option>
          <option value="newest">Newest First</option>
        </select>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="isLoading" class="text-center py-12">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 mx-auto mb-4" :style="spinnerStyle"></div>
      <p :style="textStyle">Loading products...</p>
    </div>
    
    <!-- Products Grid -->
    <div
      v-else-if="filteredProducts.length > 0"
      class="grid grid-auto-fit gap-6 animate-fade-in-up"
      :style="gridStyle"
    >
      <DynamicProductCard
        v-for="product in filteredProducts"
        :key="product.id"
        :product="product"
        :theme="theme"
        :config="components?.productCard"
      />
    </div>
    
    <!-- Empty State -->
    <div v-else class="text-center py-16 space-y-6">
      <div class="w-20 h-20 mx-auto rounded-full flex items-center justify-center" :style="emptyStateIconStyle">
        <svg class="w-10 h-10" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2M4 13h2m5 0h2m5 0h2"></path>
        </svg>
      </div>
      <div class="space-y-2">
        <h3 :style="headingStyle" class="text-xl font-semibold">No Products Found</h3>
        <p :style="textStyle">
          {{ selectedCollection ? 'This collection is empty.' : 'No products available at the moment.' }}
        </p>
      </div>
      <router-link
        :to="{ path: '/' }"
        class="inline-block px-6 py-3 font-medium rounded-lg transition-all duration-200 hover:scale-105"
        :style="primaryButtonStyle"
      >
        Back to Home
      </router-link>
    </div>

    <!-- Pagination -->
    <div v-if="totalPages > 1" class="flex justify-center items-center space-x-2">
      <button
        class="px-4 py-2 font-medium rounded-lg transition-all duration-200 hover:scale-105 disabled:opacity-50 disabled:cursor-not-allowed"
        :style="paginationButtonStyle"
        :disabled="currentPage === 1"
        @click="goToPage(currentPage - 1)"
      >
        Previous
      </button>
      
      <div class="flex items-center space-x-1">
        <button
          v-for="page in visiblePages"
          :key="page"
          class="px-3 py-2 font-medium rounded-lg transition-all duration-200 hover:scale-105"
          :style="page === currentPage ? primaryButtonStyle : paginationButtonStyle"
          @click="goToPage(page)"
        >
          {{ page }}
        </button>
      </div>
      
      <button
        class="px-4 py-2 font-medium rounded-lg transition-all duration-200 hover:scale-105 disabled:opacity-50 disabled:cursor-not-allowed"
        :style="paginationButtonStyle"
        :disabled="currentPage === totalPages"
        @click="goToPage(currentPage + 1)"
      >
        Next
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import DynamicProductCard from '@/components/DynamicProductCard.vue'
import { getCurrentShopSlug } from '@/services/shop'
import { fetchProductsPaginated } from '@/services/product'
import { fetchCollections } from '@/services/collections'
import type { Product } from '@/services/product'
import type { Collection } from '@/services/collections'
import type { DynamicTheme, ComponentConfig } from '@/services/dynamic-theme'

// =============== Props ===============

interface Props {
  storefrontConfig?: any
  theme?: DynamicTheme | null
  components?: ComponentConfig | null
  layout?: any
}

const props = defineProps<Props>()

// =============== Reactive State ===============

const shopSlug = getCurrentShopSlug()
const products = ref<Product[]>([])
const collections = ref<Collection[]>([])
const selectedCollection = ref<string | null>(null)
const isLoading = ref(true)
const currentPage = ref(1)
const pageSize = 20
const totalProducts = ref(0)
const selectedPriceRange = ref('')
const sortBy = ref('name')

// =============== Computed Properties ===============

const pageTitle = computed(() => {
  if (selectedCollection.value) {
    const collection = collections.value.find(c => c.id === selectedCollection.value)
    return collection?.title || 'Products'
  }
  return 'All Products'
})

const pageDescription = computed(() => {
  if (selectedCollection.value) {
    const collection = collections.value.find(c => c.id === selectedCollection.value)
    return (collection as any)?.description
  }
  return 'Discover our complete collection of premium products'
})

const totalPages = computed(() => Math.ceil(totalProducts.value / pageSize))

const visiblePages = computed(() => {
  const pages = []
  const start = Math.max(1, currentPage.value - 2)
  const end = Math.min(totalPages.value, currentPage.value + 2)
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  
  return pages
})

// Filter and sort products
const filteredProducts = computed(() => {
  let filtered = [...products.value]
  
  // Apply price filter
  if (selectedPriceRange.value) {
    filtered = filtered.filter(product => {
      const price = product.display_price || product.price || 0
      
      switch (selectedPriceRange.value) {
        case '0-25': return price < 25
        case '25-50': return price >= 25 && price < 50
        case '50-100': return price >= 50 && price < 100
        case '100+': return price >= 100
        default: return true
      }
    })
  }
  
  // Apply sorting
  filtered.sort((a, b) => {
    switch (sortBy.value) {
      case 'price-low':
        return (a.display_price || a.price || 0) - (b.display_price || b.price || 0)
      case 'price-high':
        return (b.display_price || b.price || 0) - (a.display_price || a.price || 0)
      case 'newest':
        return new Date((b as any).created_at || 0).getTime() - new Date((a as any).created_at || 0).getTime()
      case 'name':
      default:
        return a.name.localeCompare(b.name)
    }
  })
  
  return filtered
})

// Removed unused productGridClasses - using grid-auto-fit instead

const gridStyle = computed(() => {
  return {
    '--grid-columns': props.layout?.gridColumns || '3'
  }
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

// Tab styling
function getTabStyle(isActive: boolean) {
  if (!props.theme) return {}

  return {
    color: isActive ? props.theme.colors?.primary : props.theme.colors?.bodyText,
    borderBottomColor: isActive ? props.theme.colors?.primary : 'transparent',
    fontFamily: props.theme.fonts?.body ? `'${props.theme.fonts.body}', sans-serif` : undefined
  }
}

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

// Form element styles
const labelStyle = computed(() => {
  return {
    color: props.theme?.colors?.heading || '#1F2937',
    fontFamily: props.theme?.fonts?.body ? `'${props.theme.fonts.body}', sans-serif` : undefined
  }
})

const selectStyle = computed(() => {
  return {
    backgroundColor: props.theme?.colors?.background || '#ffffff',
    borderColor: props.theme?.colors?.border || '#E5E7EB',
    color: props.theme?.colors?.bodyText || '#6B7280',
    borderRadius: getBorderRadiusValue(props.theme?.layout?.borderStyle || 'rounded'),
    '--tw-ring-color': props.theme?.colors?.primary || '#10B981'
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

const paginationButtonStyle = computed(() => {
  return {
    backgroundColor: props.theme?.colors?.background || 'white',
    borderColor: props.theme?.colors?.border || '#E5E7EB',
    color: props.theme?.colors?.bodyText || '#6B7280',
    border: `1px solid ${props.theme?.colors?.border || '#E5E7EB'}`,
    borderRadius: getBorderRadiusValue(props.theme?.layout?.borderStyle || 'rounded')
  }
})

const countBadgeStyle = computed(() => {
  return {
    backgroundColor: `${props.theme?.colors?.primary || '#10B981'}15`,
    color: props.theme?.colors?.primary || '#10B981'
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

// =============== Methods ===============

async function loadProducts(page = 1) {
  if (!shopSlug) return

  isLoading.value = true
  try {
    const { products: fetchedProducts, total } = await fetchProductsPaginated(shopSlug, page, pageSize)
    products.value = fetchedProducts
    totalProducts.value = total
    currentPage.value = page
  } catch (error) {
    console.error('Failed to load products:', error)
  } finally {
    isLoading.value = false
  }
}

function selectCollection(id: string | null) {
  selectedCollection.value = id
  currentPage.value = 1
  loadProducts(1)
}

function goToPage(page: number) {
  if (page < 1 || page > totalPages.value) return
  loadProducts(page)
}

function getCollectionName(collectionId: string): string {
  const collection = collections.value.find(c => c.id === collectionId)
  return collection?.title || 'Collection'
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
  if (!shopSlug) return

  try {
    // Load collections first
    collections.value = await fetchCollections(shopSlug)
  } catch (error) {
    console.error('Failed to load collections:', error)
  }

  // Load products
  await loadProducts(1)
})

// Watch for filter changes
watch([selectedPriceRange, sortBy], () => {
  // Filters are applied in computed property, no need to reload
})
</script>

<style scoped>
/* Custom grid responsive classes */
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

/* Tab hover effects */
.tab-button:hover {
  border-bottom-color: var(--color-primary);
}

/* Loading animation */
@keyframes spin {
  to { transform: rotate(360deg); }
}

.animate-spin {
  animation: spin 1s linear infinite;
}

/* Focus states */
select:focus,
button:focus {
  ring-width: 2px;
  ring-color: var(--tw-ring-color);
}

/* Disabled states */
button:disabled {
  transform: none !important;
}
</style>
