<template>
  <div class="min-h-screen flex flex-col" :style="backgroundStyle">
    <!-- Dynamic Header -->
    <DynamicHeader 
      :shop="storefrontConfig?.shop" 
      :theme="currentTheme"
      :navigation="storefrontConfig?.navigation"
      :components="storefrontConfig?.components"
      v-if="storefrontConfig"
    />
    
    <!-- Loading State -->
    <div v-if="isLoading" class="flex-1 flex items-center justify-center">
      <div class="text-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary mx-auto mb-4"></div>
        <p class="text-gray-600">Loading store...</p>
      </div>
    </div>
    
    <!-- Error State -->
    <div v-else-if="error" class="flex-1 flex items-center justify-center">
      <div class="text-center py-12">
        <h1 class="text-2xl font-bold text-gray-900 mb-4">Store Not Found</h1>
        <p class="text-gray-600 mb-6">{{ error }}</p>
        <button 
          @click="retryLoad" 
          class="btn-primary text-white px-6 py-2 rounded"
        >
          Retry
        </button>
      </div>
    </div>
    
    <!-- Main Content with Dynamic Container -->
    <main 
      v-else 
      class="flex-1" 
      :class="containerClasses"
      :style="containerStyle"
    >
      <router-view 
        :storefront-config="storefrontConfig" 
        :theme="currentTheme"
        :components="storefrontConfig?.components"
      />
    </main>
    
    <!-- Dynamic Footer -->
    <DynamicFooter 
      :shop="storefrontConfig?.shop" 
      :theme="currentTheme"
      :footer-config="storefrontConfig?.footer"
      v-if="storefrontConfig"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import DynamicHeader from '@/components/DynamicHeader.vue'
import DynamicFooter from '@/components/DynamicFooter.vue'
import { 
  fetchStorefrontConfig, 
  startThemeWatch, 
  stopThemeWatch,
  useCurrentStorefront,
  useCurrentTheme,
  useThemeLoading
} from '@/services/dynamic-theme'
import type { StorefrontConfig, DynamicTheme } from '@/services/dynamic-theme'

// =============== Reactive State ===============

const route = useRoute()
const shopSlug = computed(() => getShopSlugFromRoute())

// Global theme state
const storefrontConfig = useCurrentStorefront()
const currentTheme = useCurrentTheme()
const { isLoading, error } = useThemeLoading()

// =============== Computed Styles ===============

// Dynamic background based on theme
const backgroundStyle = computed(() => {
  if (!currentTheme.value) return {}
  
  return {
    backgroundColor: currentTheme.value.colors?.background || '#ffffff',
    color: currentTheme.value.colors?.bodyText || '#6B7280',
    fontFamily: currentTheme.value.fonts?.body ? `'${currentTheme.value.fonts.body}', sans-serif` : undefined
  }
})

// Dynamic container classes
const containerClasses = computed(() => {
  const classes = ['w-full', 'mx-auto', 'px-4']
  
  if (currentTheme.value?.layout) {
    // Add dynamic container width class
    if (currentTheme.value.layout.containerWidth === 'full') {
      classes.push('max-w-none')
    } else if (currentTheme.value.layout.containerWidth === 'wide') {
      classes.push('max-w-7xl')
    } else {
      classes.push('max-w-6xl') // boxed
    }
    
    // Add spacing
    if (currentTheme.value.layout.spacing === 'tight') {
      classes.push('py-4')
    } else if (currentTheme.value.layout.spacing === 'loose') {
      classes.push('py-12')
    } else {
      classes.push('py-8') // normal
    }
  } else {
    classes.push('max-w-6xl', 'py-8')
  }
  
  return classes
})

// Dynamic container style
const containerStyle = computed(() => {
  if (!currentTheme.value?.layout) return {}
  
  return {
    '--grid-columns': currentTheme.value.layout.gridColumns || '3',
  }
})

// =============== Lifecycle ===============

onMounted(async () => {
  await loadStorefront()
  
  // Start watching for theme updates
  if (shopSlug.value) {
    startThemeWatch(shopSlug.value, 30000) // Check every 30 seconds
  }
  
  // Listen for manual theme updates
  window.addEventListener('theme-updated', handleThemeUpdate)
})

onUnmounted(() => {
  stopThemeWatch()
  window.removeEventListener('theme-updated', handleThemeUpdate)
})

// Watch for route changes (different shops)
watch(shopSlug, async (newSlug, oldSlug) => {
  if (newSlug && newSlug !== oldSlug) {
    stopThemeWatch()
    await loadStorefront()
    startThemeWatch(newSlug, 30000)
  }
})

// =============== Methods ===============

// Load storefront configuration
async function loadStorefront() {
  if (!shopSlug.value) {
    error.value = 'No shop specified'
    return
  }
  
  try {
    await fetchStorefrontConfig(shopSlug.value)
    console.log('✅ Dynamic storefront loaded for:', shopSlug.value)
  } catch (err) {
    console.error('❌ Failed to load storefront:', err)
  }
}

// Retry loading on error
async function retryLoad() {
  error.value = null
  await loadStorefront()
}

// Handle real-time theme updates
function handleThemeUpdate(event: CustomEvent) {
  console.log('🎨 Theme updated in real-time:', event.detail)
  
  // You could show a notification here
  showThemeUpdateNotification(event.detail.theme)
}

// Show theme update notification
function showThemeUpdateNotification(theme: DynamicTheme) {
  // Simple notification - could be replaced with a proper toast system
  const notification = document.createElement('div')
  notification.className = 'fixed top-4 right-4 bg-green-500 text-white px-4 py-2 rounded shadow-lg z-50 transition-all duration-300'
  notification.textContent = `✨ Store theme updated to ${theme.name} v${theme.version}`
  
  document.body.appendChild(notification)
  
  // Auto remove after 3 seconds
  setTimeout(() => {
    notification.style.opacity = '0'
    setTimeout(() => {
      document.body.removeChild(notification)
    }, 300)
  }, 3000)
}

// Extract shop slug from route/URL
function getShopSlugFromRoute(): string {
  // For subdomain style: extract from hostname
  if (typeof window !== 'undefined') {
    const hostname = window.location.hostname
    const parts = hostname.split('.')
    
    // If subdomain.domain.com format
    if (parts.length >= 3 && parts[0] !== 'www') {
      return parts[0]
    }
    
    // If localhost:port/shop-slug format
    const pathParts = window.location.pathname.split('/').filter(Boolean)
    if (pathParts.length > 0) {
      return pathParts[0]
    }
  }
  
  // Fallback to route params
  return route.params.shopSlug as string || 'demo-shop'
}

// =============== SEO & Meta Updates ===============

// Watch for theme changes to update meta tags
watch(currentTheme, (newTheme) => {
  if (newTheme && typeof window !== 'undefined') {
    // Update favicon color
    updateFaviconColor(newTheme.colors?.primary || '#10B981')
  }
}, { immediate: true })

function updateFaviconColor(color: string) {
  // Update meta theme-color for mobile browsers
  let metaThemeColor = document.querySelector('meta[name="theme-color"]')
  if (!metaThemeColor) {
    metaThemeColor = document.createElement('meta')
    metaThemeColor.setAttribute('name', 'theme-color')
    document.head.appendChild(metaThemeColor)
  }
  metaThemeColor.setAttribute('content', color)
}
</script>

<style scoped>
/* Dynamic transition effects */
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

/* Loading animation */
@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

.animate-spin {
  animation: spin 1s linear infinite;
}

/* Dynamic button styles are defined in the dynamic CSS */
.btn-primary {
  background-color: var(--color-primary, #10B981);
  border: 1px solid var(--color-primary, #10B981);
  transition: all 0.2s ease;
}

.btn-primary:hover {
  opacity: 0.9;
  transform: translateY(-1px);
}
</style>
