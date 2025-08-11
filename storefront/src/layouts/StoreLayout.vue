<template>
  <div class="min-h-screen flex flex-col" :style="backgroundStyle">
    <Header :shop="shop" :theme="currentTheme" v-if="!themeLoading" />
    
    <!-- Loading State -->
    <div v-if="themeLoading" class="flex-1 flex items-center justify-center">
      <div class="text-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary mx-auto mb-4" :style="{ borderBottomColor: currentTheme?.colors?.primary || '#10B981' }"></div>
        <p class="text-gray-600">Loading store...</p>
      </div>
    </div>
    
    <!-- Main Content -->
    <main v-else class="flex-1" :class="containerClasses" :style="containerStyle">
      <div v-if="error" class="text-center py-12">
        <h1 class="text-2xl font-bold mb-4" :style="headingStyle">Shop Not Found</h1>
        <p class="mb-6" :style="textStyle">The shop "{{ shopSlug }}" could not be found.</p>
        <router-link to="/" class="hover:underline" :style="linkStyle">Back to Shop Selection</router-link>
      </div>
      <router-view v-else :theme="currentTheme" :storefront-config="storefrontConfig" />
    </main>
    
    <Footer :shop="shop" :theme="currentTheme" v-if="!themeLoading" />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, watch } from 'vue'
import type { Shop } from '@/services/shop'
import { fetchShop, getCurrentShopSlug } from '@/services/shop'
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import { 
  fetchStorefrontConfig, 
  startThemeWatch, 
  stopThemeWatch,
  useCurrentStorefront,
  useCurrentTheme,
  useThemeLoading
} from '@/services/dynamic-theme'
import type { StorefrontConfig, DynamicTheme } from '@/services/dynamic-theme'

const shop = ref<Shop|null>(null)
const error = ref(false)
const shopSlug = getCurrentShopSlug()

// Global theme state
const storefrontConfig = useCurrentStorefront()
const currentTheme = useCurrentTheme()
const { isLoading: themeLoading } = useThemeLoading()

// =============== Computed Styles ===============

// Dynamic background based on theme
const backgroundStyle = computed(() => {
  if (!currentTheme.value) return { backgroundColor: '#F9FAFB' }
  
  return {
    backgroundColor: currentTheme.value.colors?.background || '#F9FAFB',
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

// Text styles
const headingStyle = computed(() => {
  if (!currentTheme.value) return { color: '#1F2937' }
  
  return {
    color: currentTheme.value.colors?.heading || '#1F2937',
    fontFamily: currentTheme.value.fonts?.heading ? `'${currentTheme.value.fonts.heading}', sans-serif` : undefined
  }
})

const textStyle = computed(() => {
  if (!currentTheme.value) return { color: '#6B7280' }
  
  return {
    color: currentTheme.value.colors?.bodyText || '#6B7280',
    fontFamily: currentTheme.value.fonts?.body ? `'${currentTheme.value.fonts.body}', sans-serif` : undefined
  }
})

const linkStyle = computed(() => {
  if (!currentTheme.value) return { color: '#2563EB' }
  
  return {
    color: currentTheme.value.colors?.primary || '#2563EB'
  }
})

// =============== Lifecycle ===============

onMounted(async () => {
  try {
    if (!shopSlug) {
      error.value = true
      return
    }
    
    // Load shop data
    shop.value = await fetchShop(shopSlug)
    if (!shop.value) {
      error.value = true
      return
    }
    
    // Load dynamic storefront configuration
    await loadStorefront()
    
    // Start watching for theme updates
    startThemeWatch(shopSlug, 30000) // Check every 30 seconds
    
    // Listen for manual theme updates
    window.addEventListener('theme-updated', handleThemeUpdate)
    
  } catch (err) {
    console.error('Error fetching shop:', err)
    error.value = true
  }
})

onUnmounted(() => {
  stopThemeWatch()
  window.removeEventListener('theme-updated', handleThemeUpdate)
})

// =============== Methods ===============

// Load storefront configuration
async function loadStorefront() {
  if (!shopSlug) return
  
  try {
    await fetchStorefrontConfig(shopSlug)
    console.log('✅ Dynamic storefront loaded for:', shopSlug)
  } catch (err) {
    console.warn('❌ Failed to load dynamic theme, using fallback:', err)
    // Continue with hardcoded styles as fallback
  }
}

// Handle real-time theme updates
function handleThemeUpdate(event: CustomEvent) {
  console.log('🎨 Theme updated in real-time:', event.detail)
  
  // Show update notification
  showThemeUpdateNotification(event.detail.theme)
}

// Show theme update notification
function showThemeUpdateNotification(theme: DynamicTheme) {
  const notification = document.createElement('div')
  notification.className = 'fixed top-4 right-4 bg-green-500 text-white px-4 py-2 rounded shadow-lg z-50 transition-all duration-300'
  notification.textContent = `✨ Store theme updated to ${theme.name} v${theme.version}`
  
  document.body.appendChild(notification)
  
  // Auto remove after 3 seconds
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
