<template>
  <div
    class="min-h-screen flex flex-col transition-colors duration-500 ease-in-out"
    :style="backgroundStyle"
  >
    <!-- Dynamic Header -->
    <DynamicHeader
      v-if="storefrontConfig && !isLoading"
      :shop="storefrontConfig.shop"
      :theme="currentTheme"
      :navigation="storefrontConfig.navigation"
      :components="storefrontConfig.components"
      class="shadow-md"
    />

    <!-- Loading State -->
    <div
      v-if="isLoading"
      class="flex-1 flex flex-col items-center justify-center px-4 text-center"
    >
      <div
        class="animate-spin rounded-full h-14 w-14 border-4 border-t-transparent border-primary mx-auto mb-6"
        :style="{ borderBottomColor: currentTheme?.colors?.primary || '#10B981' }"
      ></div>
      <p
        class="text-lg font-medium"
        :style="{ color: currentTheme?.colors?.bodyText || '#6B7280', fontFamily: currentTheme?.fonts?.body ? `'${currentTheme.fonts.body}', sans-serif` : undefined }"
      >
        {{ loadingMessage }}
      </p>
    </div>

    <!-- Error State -->
    <div
      v-else-if="error"
      class="flex-1 flex flex-col items-center justify-center px-6 text-center space-y-6"
    >
      <div
        class="w-20 h-20 rounded-full flex items-center justify-center"
        :style="errorIconStyle"
      >
        <svg
          class="w-10 h-10"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
        >
          <path d="M12 9v2m0 4h.01"></path>
          <path
            d="M3.732 16.5c-.77.833.192 2.5 1.732 2.5h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5z"
          ></path>
        </svg>
      </div>
      <h1
        class="text-3xl font-extrabold"
        :style="errorHeadingStyle"
      >
        Store Not Found
      </h1>
      <p
        class="max-w-md text-base"
        :style="errorTextStyle"
      >
        {{ error }}
      </p>
      <button
        @click="retryLoad"
        class="px-6 py-3 font-semibold rounded-lg shadow-md hover:shadow-lg focus:outline-none focus:ring-2 focus:ring-offset-1 transition-transform hover:scale-105 active:scale-95"
        :style="{ backgroundColor: currentTheme?.colors?.primary || '#10B981', color: currentTheme?.colors?.background || '#fff' }"
      >
        Retry Loading
      </button>
    </div>

    <!-- Main Content with Dynamic Container -->
    <main
      v-else
      :class="containerClasses"
      :style="containerStyle"
      class="flex-1 mx-auto px-4"
    >
      <router-view v-slot="{ Component }">
        <component
          :is="Component"
          :storefront-config="storefrontConfig"
          :theme="currentTheme"
          :components="storefrontConfig?.components"
          :layout="storefrontConfig?.layout"
          :navigation="storefrontConfig?.navigation"
          :footer="storefrontConfig?.footer"
          :seo="storefrontConfig?.seo"
          :analytics="storefrontConfig?.analytics"
        />
      </router-view>
    </main>

    <!-- Dynamic Footer -->
    <DynamicFooter
      v-if="storefrontConfig && !isLoading"
      :shop="storefrontConfig.shop"
      :theme="currentTheme"
      :footer-config="storefrontConfig.footer"
      class="mt-auto border-t border-gray-200"
    />

 
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, watch } from 'vue'
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

const shopSlug = computed(() => getCurrentShopSlug())
const storefrontConfig = useCurrentStorefront()
const currentTheme = useCurrentTheme()
const { isLoading, error } = useThemeLoading()

const showThemeUpdateNotification = ref(false)
const themeUpdateMessage = ref('')
const loadingMessage = ref('Loading store...')

const backgroundStyle = computed(() => {
  if (!currentTheme.value) return { backgroundColor: '#fff', color: '#6B7280' }
  return {
    backgroundColor: currentTheme.value.colors?.background || '#fff',
    color: currentTheme.value.colors?.bodyText || '#6B7280',
    fontFamily: currentTheme.value.fonts?.body
      ? `'${currentTheme.value.fonts.body}', sans-serif`
      : undefined,
    '--color-primary': currentTheme.value.colors?.primary || '#6366F1',
    '--color-secondary': currentTheme.value.colors?.secondary || '#8B5CF6',
    '--color-background': currentTheme.value.colors?.background || '#fff',
    '--color-heading': currentTheme.value.colors?.heading || '#1F2937',
    '--color-body-text': currentTheme.value.colors?.bodyText || '#6B7280',
    '--color-border': currentTheme.value.colors?.border || '#E5E7EB'
  }
})

const containerClasses = computed(() => {
  const classes = ['w-full', 'mx-auto', 'px-4', 'sm:px-6', 'max-w-[var(--container-width)]']
  if (storefrontConfig.value?.layout) {
    const layout = storefrontConfig.value.layout
    if (layout.containerWidth === 'full') classes.push('max-w-none')
    else if (layout.containerWidth === 'wide') classes.push('max-w-7xl')
    else classes.push('max-w-6xl')

    if (layout.spacing === 'tight') classes.push('py-4')
    else if (layout.spacing === 'loose') classes.push('py-12')
    else classes.push('py-8')
  } else {
    classes.push('max-w-6xl', 'py-8')
  }
  return classes
})

const containerStyle = computed(() => {
  if (!storefrontConfig.value?.layout || !currentTheme.value) return {}
  const layout = storefrontConfig.value.layout
  const theme = currentTheme.value
  return {
    '--grid-columns': layout.gridColumns || '3',
    '--container-width': getContainerWidthValue(layout.containerWidth),
    '--border-radius': getBorderRadiusValue(layout.borderStyle),
    '--font-heading': theme.fonts?.heading
      ? `'${theme.fonts.heading}', sans-serif`
      : 'Inter, sans-serif',
    '--font-body': theme.fonts?.body
      ? `'${theme.fonts.body}', sans-serif`
      : 'Inter, sans-serif'
  }
})

const errorIconStyle = computed(() => {
  const primaryColor = currentTheme.value?.colors?.primary || '#EF4444'
  return {
    backgroundColor: `${primaryColor}10`,
    color: primaryColor
  }
})

const errorHeadingStyle = computed(() => ({
  color: currentTheme.value?.colors?.heading || '#1F2937',
  fontFamily: currentTheme.value?.fonts?.heading
    ? `'${currentTheme.value.fonts.heading}', sans-serif`
    : undefined
}))

const errorTextStyle = computed(() => ({
  color: currentTheme.value?.colors?.bodyText || '#6B7280',
  fontFamily: currentTheme.value?.fonts?.body
    ? `'${currentTheme.value.fonts.body}', sans-serif`
    : undefined
}))

// const retryButtonStyle = computed(() => ({
//   backgroundColor: currentTheme.value?.colors?.primary || '#10B981',
//   color: currentTheme.value?.colors?.background || 'white',
//   borderRadius: getBorderRadiusValue(
//     storefrontConfig.value?.layout?.borderStyle || 'rounded'
//   )
// }))


onMounted(async () => {
  await loadStorefront()
  if (shopSlug.value) {
    startThemeWatch(shopSlug.value, 30000)
  }
  window.addEventListener('theme-updated', handleThemeUpdate as EventListener)
})

onUnmounted(() => {
  stopThemeWatch()
  window.removeEventListener('theme-updated', handleThemeUpdate as EventListener)
})

watch(shopSlug, async (newSlug, oldSlug) => {
  if (newSlug && newSlug !== oldSlug) {
    stopThemeWatch()
    loadingMessage.value = `Loading ${newSlug} store...`
    await loadStorefront()
    if (newSlug) startThemeWatch(newSlug, 30000)
  }
})

watch(storefrontConfig, (newConfig) => {
  if (newConfig?.seo && typeof window !== 'undefined') {
    updateDocumentMeta(newConfig.seo)
  }
}, { immediate: true })

async function loadStorefront() {
  if (!shopSlug.value) {
    error.value = 'No shop specified in URL'
    return
  }
  loadingMessage.value = `Loading ${shopSlug.value} store...`
  try {
    await fetchStorefrontConfig(shopSlug.value)
    loadingMessage.value = 'Applying theme...'
    if (currentTheme.value) {
      await applyThemeSpecificConfigurations(currentTheme.value)
    }
  } catch (err) {
    error.value =
      err instanceof Error ? err.message : 'Failed to load store configuration'
  }
}

async function applyThemeSpecificConfigurations(theme: DynamicTheme) {
  if (storefrontConfig.value?.shop?.name) {
    document.title = storefrontConfig.value.shop.name
  }
  if (theme.mobile?.touchOptimized === 'true') {
    document.body.classList.add('touch-optimized')
  }
  if (theme.performance?.performanceScore && theme.performance.performanceScore < 70) {
    console.warn('⚠️ Theme performance score is low:', theme.performance.performanceScore)
  }
}

async function retryLoad() {
  error.value = null
  await loadStorefront()
}

function handleThemeUpdate(event: CustomEvent) {
  const { theme } = event.detail
  themeUpdateMessage.value = `✨ Theme updated to ${theme.name} v${theme.version}`
  showThemeUpdateNotification.value = true
  setTimeout(() => (showThemeUpdateNotification.value = false), 4000)
  applyThemeSpecificConfigurations(theme)
}

function updateDocumentMeta(seo: any) {
  if (!seo || typeof window === 'undefined') return
  if (seo.title) document.title = seo.title
  updateMetaTag('description', seo.description)
  if (seo.keywords) updateMetaTag('keywords', seo.keywords)
  updateMetaTag('theme-color', seo.themeColor)
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

function getContainerWidthValue(containerWidth: string): string {
  switch (containerWidth) {
    case 'full':
      return '100%'
    case 'wide':
      return '1400px'
    case 'boxed':
      return '1200px'
    default:
      return '1200px'
  }
}

function getBorderRadiusValue(borderStyle: string): string {
  switch (borderStyle) {
    case 'sharp':
      return '0px'
    case 'rounded':
      return '8px'
    case 'pill':
      return '50px'
    default:
      return '8px'
  }
}
</script>

<style scoped>
/* Transition effects */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* Spin animation for loading */
@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}
.animate-spin {
  animation: spin 1s linear infinite;
}

/* Button primary */
.btn-primary {
  background-color: var(--color-primary, #6366f1);
  border: 1px solid var(--color-primary, #6366f1);
  color: var(--color-background, #fff);
  border-radius: var(--border-radius, 8px);
  cursor: pointer;
  transition: all 0.25s ease;
  box-shadow: 0 4px 10px rgb(99 102 241 / 0.4);
  user-select: none;
}
.btn-primary:hover {
  opacity: 0.9;
  transform: translateY(-2px);
  box-shadow: 0 6px 15px rgb(16 185 129 / 0.6);
}
.btn-primary:active {
  transform: scale(0.95);
  box-shadow: 0 3px 7px rgb(16 185 129 / 0.3);
}

/* Responsive container */
@media (max-width: 768px) {
  .container {
    padding-left: 1rem;
    padding-right: 1rem;
  }
}

/* Notification slide-in */
.notification-enter-active,
.notification-leave-active {
  transition: all 0.3s ease;
}
.notification-enter-from {
  opacity: 0;
  transform: translateX(120%);
}
.notification-leave-to {
  opacity: 0;
  transform: translateX(120%);
}
</style>
