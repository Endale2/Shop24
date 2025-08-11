<template>
  <div class="min-h-screen flex flex-col" :style="backgroundStyle">
    <!-- Dynamic Header -->
    <DynamicHeader
      :shop="storefrontConfig?.shop"
      :theme="currentTheme"
      :navigation="storefrontConfig?.navigation"
      :components="storefrontConfig?.components"
      v-if="storefrontConfig && !isLoading"
    />

    <!-- Loading State -->
    <div v-if="isLoading" class="flex-1 flex items-center justify-center">
      <div class="text-center py-12">
        <div
          class="animate-spin rounded-full h-12 w-12 border-b-2 mx-auto mb-4"
          :style="loadingSpinnerStyle"
        ></div>
        <p :style="loadingTextStyle">{{ loadingMessage }}</p>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="flex-1 flex items-center justify-center">
      <div class="text-center py-12 space-y-4">
        <div class="w-16 h-16 mx-auto rounded-full flex items-center justify-center" :style="errorIconStyle">
          <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
          </svg>
        </div>
        <h1 class="text-2xl font-bold mb-4" :style="errorHeadingStyle">Store Not Found</h1>
        <p class="mb-6 max-w-md mx-auto" :style="errorTextStyle">{{ error }}</p>
        <button
          @click="retryLoad"
          class="px-6 py-3 font-medium rounded-lg transition-all duration-200 hover:scale-105"
          :style="retryButtonStyle"
        >
          Retry Loading
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
        :layout="storefrontConfig?.layout"
        :navigation="storefrontConfig?.navigation"
        :footer="storefrontConfig?.footer"
        :seo="storefrontConfig?.seo"
        :analytics="storefrontConfig?.analytics"
      />
    </main>

    <!-- Dynamic Footer -->
    <DynamicFooter
      :shop="storefrontConfig?.shop"
      :theme="currentTheme"
      :footer-config="storefrontConfig?.footer"
      v-if="storefrontConfig && !isLoading"
    />

    <!-- Theme Update Notification -->
    <div
      v-if="showThemeUpdateNotification"
      class="fixed top-4 right-4 z-50 transition-all duration-300"
      :style="notificationStyle"
    >
      <div class="flex items-center space-x-3 px-4 py-3 rounded-lg shadow-lg">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 4V2a1 1 0 011-1h8a1 1 0 011 1v2m-9 4v10a2 2 0 002 2h6a2 2 0 002-2V8M7 8h10"></path>
        </svg>
        <span class="font-medium">{{ themeUpdateMessage }}</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, watch } from 'vue'
// import { useRoute } from 'vue-router' // Unused for now
import DynamicHeader from '@/components/DynamicHeader.vue'
import DynamicFooter from '@/components/DynamicFooter.vue'
import { getCurrentShopSlug } from '@/services/shop'
import {
  fetchStorefrontConfig,
  startThemeWatch,
  stopThemeWatch,
  useCurrentStorefront,
  useCurrentTheme,
  useThemeLoading
} from '@/services/dynamic-theme'
import type { DynamicTheme } from '@/services/dynamic-theme'

// =============== Reactive State ===============

// const route = useRoute() // Unused for now
const shopSlug = computed(() => getCurrentShopSlug())

// Global theme state
const storefrontConfig = useCurrentStorefront()
const currentTheme = useCurrentTheme()
const { isLoading, error } = useThemeLoading()

// Local state for notifications and loading messages
const showThemeUpdateNotification = ref(false)
const themeUpdateMessage = ref('')
const loadingMessage = ref('Loading store...')

// =============== Computed Styles ===============

// Dynamic background based on theme
const backgroundStyle = computed(() => {
  if (!currentTheme.value) return { backgroundColor: '#ffffff', color: '#6B7280' }

  return {
    backgroundColor: currentTheme.value.colors?.background || '#ffffff',
    color: currentTheme.value.colors?.bodyText || '#6B7280',
    fontFamily: currentTheme.value.fonts?.body ? `'${currentTheme.value.fonts.body}', sans-serif` : undefined,
    '--color-primary': currentTheme.value.colors?.primary || '#10B981',
    '--color-secondary': currentTheme.value.colors?.secondary || '#F59E0B',
    '--color-background': currentTheme.value.colors?.background || '#ffffff',
    '--color-heading': currentTheme.value.colors?.heading || '#1F2937',
    '--color-body-text': currentTheme.value.colors?.bodyText || '#6B7280',
    '--color-border': currentTheme.value.colors?.border || '#E5E7EB'
  }
})

// Dynamic container classes based on layout configuration
const containerClasses = computed(() => {
  const classes = ['w-full', 'mx-auto', 'px-4']

  if (storefrontConfig.value?.layout) {
    const layout = storefrontConfig.value.layout

    // Add dynamic container width class
    if (layout.containerWidth === 'full') {
      classes.push('max-w-none')
    } else if (layout.containerWidth === 'wide') {
      classes.push('max-w-7xl')
    } else {
      classes.push('max-w-6xl') // boxed
    }

    // Add spacing based on layout configuration
    if (layout.spacing === 'tight') {
      classes.push('py-4')
    } else if (layout.spacing === 'loose') {
      classes.push('py-12')
    } else {
      classes.push('py-8') // normal
    }
  } else {
    classes.push('max-w-6xl', 'py-8')
  }

  return classes
})

// Dynamic container style with CSS variables
const containerStyle = computed(() => {
  if (!storefrontConfig.value?.layout || !currentTheme.value) return {}

  const layout = storefrontConfig.value.layout
  const theme = currentTheme.value

  return {
    '--grid-columns': layout.gridColumns || '3',
    '--container-width': getContainerWidthValue(layout.containerWidth),
    '--border-radius': getBorderRadiusValue(layout.borderStyle),
    '--font-heading': theme.fonts?.heading ? `'${theme.fonts.heading}', sans-serif` : 'Inter, sans-serif',
    '--font-body': theme.fonts?.body ? `'${theme.fonts.body}', sans-serif` : 'Inter, sans-serif'
  }
})

// Loading spinner style
const loadingSpinnerStyle = computed(() => {
  return {
    borderBottomColor: currentTheme.value?.colors?.primary || '#10B981'
  }
})

// Loading text style
const loadingTextStyle = computed(() => {
  return {
    color: currentTheme.value?.colors?.bodyText || '#6B7280',
    fontFamily: currentTheme.value?.fonts?.body ? `'${currentTheme.value.fonts.body}', sans-serif` : undefined
  }
})

// Error state styles
const errorIconStyle = computed(() => {
  const primaryColor = currentTheme.value?.colors?.primary || '#EF4444'
  return {
    backgroundColor: `${primaryColor}10`,
    color: primaryColor
  }
})

const errorHeadingStyle = computed(() => {
  return {
    color: currentTheme.value?.colors?.heading || '#1F2937',
    fontFamily: currentTheme.value?.fonts?.heading ? `'${currentTheme.value.fonts.heading}', sans-serif` : undefined
  }
})

const errorTextStyle = computed(() => {
  return {
    color: currentTheme.value?.colors?.bodyText || '#6B7280',
    fontFamily: currentTheme.value?.fonts?.body ? `'${currentTheme.value.fonts.body}', sans-serif` : undefined
  }
})

const retryButtonStyle = computed(() => {
  return {
    backgroundColor: currentTheme.value?.colors?.primary || '#10B981',
    color: currentTheme.value?.colors?.background || 'white',
    borderRadius: getBorderRadiusValue(storefrontConfig.value?.layout?.borderStyle || 'rounded')
  }
})

// Theme update notification style
const notificationStyle = computed(() => {
  return {
    backgroundColor: currentTheme.value?.colors?.primary || '#10B981',
    color: currentTheme.value?.colors?.background || 'white'
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
  window.addEventListener('theme-updated', handleThemeUpdate as EventListener)
})

onUnmounted(() => {
  stopThemeWatch()
  window.removeEventListener('theme-updated', handleThemeUpdate as EventListener)
})

// Watch for route changes (different shops)
watch(shopSlug, async (newSlug, oldSlug) => {
  if (newSlug && newSlug !== oldSlug) {
    stopThemeWatch()
    loadingMessage.value = `Loading ${newSlug} store...`
    await loadStorefront()
    if (newSlug) {
      startThemeWatch(newSlug, 30000)
    }
  }
})

// Watch for storefront config changes to update document title and meta
watch(storefrontConfig, (newConfig) => {
  if (newConfig?.seo && typeof window !== 'undefined') {
    updateDocumentMeta(newConfig.seo)
  }
}, { immediate: true })

// =============== Methods ===============

// Load storefront configuration with enhanced error handling
async function loadStorefront() {
  if (!shopSlug.value) {
    error.value = 'No shop specified in URL'
    return
  }

  loadingMessage.value = `Loading ${shopSlug.value} store...`

  try {
    await fetchStorefrontConfig(shopSlug.value)
    console.log('✅ Dynamic storefront loaded for:', shopSlug.value)

    // Update loading message for theme application
    loadingMessage.value = 'Applying theme...'

    // Apply any additional theme-specific configurations
    if (currentTheme.value) {
      await applyThemeSpecificConfigurations(currentTheme.value)
    }

  } catch (err) {
    console.error('❌ Failed to load storefront:', err)
    error.value = err instanceof Error ? err.message : 'Failed to load store configuration'
  }
}

// Apply theme-specific configurations beyond CSS
async function applyThemeSpecificConfigurations(theme: DynamicTheme) {
  // Update document title with shop name and theme
  if (storefrontConfig.value?.shop?.name) {
    document.title = storefrontConfig.value.shop.name
  }

  // Apply mobile-specific configurations
  if (theme.mobile?.touchOptimized === 'true') {
    document.body.classList.add('touch-optimized')
  }

  // Apply performance optimizations
  if (theme.performance?.performanceScore && theme.performance.performanceScore < 70) {
    console.warn('⚠️ Theme performance score is low:', theme.performance.performanceScore)
  }
}

// Retry loading on error
async function retryLoad() {
  error.value = null
  await loadStorefront()
}

// Handle real-time theme updates with enhanced notification
function handleThemeUpdate(event: CustomEvent) {
  console.log('🎨 Theme updated in real-time:', event.detail)

  const { theme } = event.detail
  themeUpdateMessage.value = `✨ Theme updated to ${theme.name} v${theme.version}`
  showThemeUpdateNotification.value = true

  // Auto-hide notification after 4 seconds
  setTimeout(() => {
    showThemeUpdateNotification.value = false
  }, 4000)

  // Apply new theme configurations
  applyThemeSpecificConfigurations(theme)
}

// Update document meta tags based on SEO config
function updateDocumentMeta(seo: any) {
  if (!seo || typeof window === 'undefined') return

  // Update title
  if (seo.title) {
    document.title = seo.title
  }

  // Update meta description
  updateMetaTag('description', seo.description)

  // Update meta keywords
  if (seo.keywords) {
    updateMetaTag('keywords', seo.keywords)
  }

  // Update theme color for mobile browsers
  updateMetaTag('theme-color', seo.themeColor)

  // Update Open Graph tags
  updateMetaTag('og:title', seo.title, 'property')
  updateMetaTag('og:description', seo.description, 'property')
  updateMetaTag('og:image', seo.ogImage, 'property')
}

function updateMetaTag(name: string, content: string, attribute = 'name') {
  if (!content) return

  let meta = document.querySelector(`meta[${attribute}="${name}"]`)
  if (!meta) {
    meta = document.createElement('meta')
    meta.setAttribute(attribute, name)
    document.head.appendChild(meta)
  }
  meta.setAttribute('content', content)
}

// Note: Shop slug extraction is now handled by getCurrentShopSlug() from @/services/shop

// =============== Helper Functions ===============

function getContainerWidthValue(containerWidth: string): string {
  switch (containerWidth) {
    case 'full': return '100%'
    case 'wide': return '1400px'
    case 'boxed': return '1200px'
    default: return '1200px'
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

// =============== SEO & Meta Updates ===============

// Watch for theme changes to update meta tags and apply global styles
watch(currentTheme, (newTheme) => {
  if (newTheme && typeof window !== 'undefined') {
    // Update favicon color
    updateFaviconColor(newTheme.colors?.primary || '#10B981')

    // Apply global font loading
    loadThemeFonts(newTheme.fonts)

    // Update body classes for theme-specific styling
    updateBodyClasses(newTheme)
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

function loadThemeFonts(fonts: any) {
  if (!fonts) return

  // Load Google Fonts if needed
  const fontsToLoad = []
  if (fonts.heading && fonts.heading !== 'Inter') {
    fontsToLoad.push(fonts.heading)
  }
  if (fonts.body && fonts.body !== 'Inter' && !fontsToLoad.includes(fonts.body)) {
    fontsToLoad.push(fonts.body)
  }

  if (fontsToLoad.length > 0) {
    const fontLink = document.createElement('link')
    fontLink.rel = 'stylesheet'
    fontLink.href = `https://fonts.googleapis.com/css2?${fontsToLoad.map(font => `family=${font.replace(' ', '+')}`).join('&')}&display=swap`

    // Remove existing font link if any
    const existingFontLink = document.querySelector('link[href*="fonts.googleapis.com"]')
    if (existingFontLink) {
      existingFontLink.remove()
    }

    document.head.appendChild(fontLink)
  }
}

function updateBodyClasses(theme: DynamicTheme) {
  const body = document.body

  // Remove existing theme classes
  body.classList.remove('theme-dark', 'theme-light', 'theme-compact', 'theme-large')

  // Add theme-specific classes
  if (theme.colors?.background && isColorDark(theme.colors.background)) {
    body.classList.add('theme-dark')
  } else {
    body.classList.add('theme-light')
  }

  // Add layout-specific classes
  if (theme.layout?.headerStyle === 'compact') {
    body.classList.add('theme-compact')
  } else if (theme.layout?.headerStyle === 'large') {
    body.classList.add('theme-large')
  }
}

function isColorDark(color: string): boolean {
  // Simple check for dark colors
  if (color.startsWith('#')) {
    const hex = color.slice(1)
    const r = parseInt(hex.substr(0, 2), 16)
    const g = parseInt(hex.substr(2, 2), 16)
    const b = parseInt(hex.substr(4, 2), 16)
    const brightness = (r * 299 + g * 587 + b * 114) / 1000
    return brightness < 128
  }
  return false
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

/* Dynamic button styles using CSS variables */
.btn-primary {
  background-color: var(--color-primary, #10B981);
  border: 1px solid var(--color-primary, #10B981);
  color: var(--color-background, white);
  border-radius: var(--border-radius, 8px);
  transition: all 0.2s ease;
}

.btn-primary:hover {
  opacity: 0.9;
  transform: translateY(-1px);
}

/* Theme-specific body classes */
:global(.theme-dark) {
  color-scheme: dark;
}

:global(.theme-light) {
  color-scheme: light;
}

:global(.theme-compact) {
  --header-height: 60px;
}

:global(.theme-large) {
  --header-height: 100px;
}

:global(.touch-optimized) {
  touch-action: manipulation;
  -webkit-tap-highlight-color: transparent;
}

/* Enhanced responsive design */
@media (max-width: 768px) {
  .container {
    padding-left: 1rem;
    padding-right: 1rem;
  }
}

/* Notification animations */
.notification-enter-active, .notification-leave-active {
  transition: all 0.3s ease;
}

.notification-enter-from {
  opacity: 0;
  transform: translateX(100%);
}

.notification-leave-to {
  opacity: 0;
  transform: translateX(100%);
}
</style>
