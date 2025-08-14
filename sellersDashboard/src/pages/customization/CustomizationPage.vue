<template>
  <div class="min-h-full bg-gray-50">
    <div class="p-4 sm:p-6 lg:p-8">
      
      <!-- Header Section -->
      <div class="mb-6 sm:mb-8">
        <div class="flex items-center mb-3">
          <div class="w-10 h-10 bg-gradient-to-br from-purple-500 to-purple-600 rounded-lg flex items-center justify-center mr-3 shadow-sm">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zM21 5a2 2 0 00-2-2h-4a2 2 0 00-2 2v12a4 4 0 004 4h4a2 2 0 002-2V5z" />
            </svg>
          </div>
          <div>
            <h1 class="text-2xl sm:text-3xl font-bold text-gray-900 tracking-tight">
              Customization Hub
            </h1>
            <p class="text-sm text-gray-600 mt-1">
              Personalize your shop's appearance and customer experience
            </p>
          </div>
        </div>
      </div>

      <!-- Customization Cards Grid -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        
        <!-- Theme Customization Card -->
        <div 
          @click="goToTheme"
          class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden cursor-pointer transform hover:scale-105 hover:shadow-md transition duration-200 ease-in-out group"
        >
          <div class="p-6">
            <div class="flex items-center mb-4">
              <div class="w-12 h-12 bg-gradient-to-br from-purple-500 to-purple-600 rounded-lg flex items-center justify-center mr-4 group-hover:scale-110 transition-transform duration-200">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zM21 5a2 2 0 00-2-2h-4a2 2 0 00-2 2v12a4 4 0 004 4h4a2 2 0 002-2V5z" />
                </svg>
              </div>
              <div>
                <h3 class="text-lg font-semibold text-gray-900 group-hover:text-purple-700 transition-colors duration-200">Theme</h3>
                <p class="text-sm text-gray-600">Colors & Branding</p>
              </div>
            </div>
            <p class="text-sm text-gray-700 mb-4">Customize your shop's colors, fonts, and overall visual theme to match your brand identity.</p>
            <div class="flex items-center text-sm text-purple-600 font-medium group-hover:text-purple-700">
              <span>Customize theme</span>
              <svg class="w-4 h-4 ml-1 group-hover:translate-x-1 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </div>
          </div>
        </div>

        <!-- Layout Customization Card -->
        <div 
          @click="goToLayout"
          class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden cursor-pointer transform hover:scale-105 hover:shadow-md transition duration-200 ease-in-out group"
        >
          <div class="p-6">
            <div class="flex items-center mb-4">
              <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg flex items-center justify-center mr-4 group-hover:scale-110 transition-transform duration-200">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 5a1 1 0 011-1h14a1 1 0 011 1v2a1 1 0 01-1 1H5a1 1 0 01-1-1V5zM4 13a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H5a1 1 0 01-1-1v-6zM16 13a1 1 0 011-1h2a1 1 0 011 1v6a1 1 0 01-1 1h-2a1 1 0 01-1-1v-6z" />
                </svg>
              </div>
              <div>
                <h3 class="text-lg font-semibold text-gray-900 group-hover:text-blue-700 transition-colors duration-200">Layout</h3>
                <p class="text-sm text-gray-600">Page Structure</p>
              </div>
            </div>
            <p class="text-sm text-gray-700 mb-4">Design your storefront layout, arrange sections, and customize the structure of your pages.</p>
            <div class="flex items-center text-sm text-blue-600 font-medium group-hover:text-blue-700">
              <span>Design layout</span>
              <svg class="w-4 h-4 ml-1 group-hover:translate-x-1 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </div>
          </div>
        </div>

        <!-- Typography Card -->
        <div 
          @click="goToTypography"
          class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden cursor-pointer transform hover:scale-105 hover:shadow-md transition duration-200 ease-in-out group"
        >
          <div class="p-6">
            <div class="flex items-center mb-4">
              <div class="w-12 h-12 bg-gradient-to-br from-indigo-500 to-indigo-600 rounded-lg flex items-center justify-center mr-4 group-hover:scale-110 transition-transform duration-200">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                </svg>
              </div>
              <div>
                <h3 class="text-lg font-semibold text-gray-900 group-hover:text-indigo-700 transition-colors duration-200">Typography</h3>
                <p class="text-sm text-gray-600">Fonts & Text</p>
              </div>
            </div>
            <p class="text-sm text-gray-700 mb-4">Choose fonts, text sizes, and typography styles that reflect your brand's personality.</p>
            <div class="flex items-center text-sm text-indigo-600 font-medium group-hover:text-indigo-700">
              <span>Style typography</span>
              <svg class="w-4 h-4 ml-1 group-hover:translate-x-1 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </div>
          </div>
        </div>

        <!-- Custom CSS Card -->
        <div 
          @click="goToCustomCSS"
          class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden cursor-pointer transform hover:scale-105 hover:shadow-md transition duration-200 ease-in-out group"
        >
          <div class="p-6">
            <div class="flex items-center mb-4">
              <div class="w-12 h-12 bg-gradient-to-br from-emerald-500 to-emerald-600 rounded-lg flex items-center justify-center mr-4 group-hover:scale-110 transition-transform duration-200">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4" />
                </svg>
              </div>
              <div>
                <h3 class="text-lg font-semibold text-gray-900 group-hover:text-emerald-700 transition-colors duration-200">Custom CSS</h3>
                <p class="text-sm text-gray-600">Advanced Styling</p>
              </div>
            </div>
            <p class="text-sm text-gray-700 mb-4">Add custom CSS code for advanced styling and unique design modifications.</p>
            <div class="flex items-center text-sm text-emerald-600 font-medium group-hover:text-emerald-700">
              <span>Write custom code</span>
              <svg class="w-4 h-4 ml-1 group-hover:translate-x-1 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </div>
          </div>
        </div>

        <!-- Mobile Customization Card -->
        <div 
          @click="goToMobile"
          class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden cursor-pointer transform hover:scale-105 hover:shadow-md transition duration-200 ease-in-out group"
        >
          <div class="p-6">
            <div class="flex items-center mb-4">
              <div class="w-12 h-12 bg-gradient-to-br from-pink-500 to-pink-600 rounded-lg flex items-center justify-center mr-4 group-hover:scale-110 transition-transform duration-200">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 18h.01M8 21h8a1 1 0 001-1V4a1 1 0 00-1-1H8a1 1 0 00-1 1v16a1 1 0 001 1z" />
                </svg>
              </div>
              <div>
                <h3 class="text-lg font-semibold text-gray-900 group-hover:text-pink-700 transition-colors duration-200">Mobile</h3>
                <p class="text-sm text-gray-600">Mobile Experience</p>
              </div>
            </div>
            <p class="text-sm text-gray-700 mb-4">Optimize your shop's mobile experience and responsive design settings.</p>
            <div class="flex items-center text-sm text-pink-600 font-medium group-hover:text-pink-700">
              <span>Optimize mobile</span>
              <svg class="w-4 h-4 ml-1 group-hover:translate-x-1 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </div>
          </div>
        </div>

        <!-- SEO & Meta Card -->
        <div 
          @click="goToSEO"
          class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden cursor-pointer transform hover:scale-105 hover:shadow-md transition duration-200 ease-in-out group"
        >
          <div class="p-6">
            <div class="flex items-center mb-4">
              <div class="w-12 h-12 bg-gradient-to-br from-orange-500 to-orange-600 rounded-lg flex items-center justify-center mr-4 group-hover:scale-110 transition-transform duration-200">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
              </div>
              <div>
                <h3 class="text-lg font-semibold text-gray-900 group-hover:text-orange-700 transition-colors duration-200">SEO & Meta</h3>
                <p class="text-sm text-gray-600">Search Optimization</p>
              </div>
            </div>
            <p class="text-sm text-gray-700 mb-4">Configure meta tags, SEO settings, and search engine optimization features.</p>
            <div class="flex items-center text-sm text-orange-600 font-medium group-hover:text-orange-700">
              <span>Optimize SEO</span>
              <svg class="w-4 h-4 ml-1 group-hover:translate-x-1 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </div>
          </div>
        </div>

      </div>

      <!-- Quick Actions Section -->
      <div class="mt-12 bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
        <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
          <div class="flex items-center">
            <div class="w-8 h-8 bg-gradient-to-br from-gray-500 to-gray-600 rounded-lg flex items-center justify-center mr-3">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
              </svg>
            </div>
            <div>
              <h2 class="text-lg font-semibold text-gray-900">Quick Actions</h2>
              <p class="text-sm text-gray-600">Common customization tasks</p>
            </div>
          </div>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
            <button
              @click="previewShop"
              class="flex flex-col items-center justify-center p-3 border border-gray-300 rounded-lg hover:bg-gray-50 transition duration-150 ease-in-out group"
            >
              <div class="flex items-center mb-1">
                <svg class="w-5 h-5 text-gray-500 mr-2 group-hover:text-blue-600 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                </svg>
                <span class="text-sm font-medium text-gray-700 group-hover:text-blue-600 transition-colors">Preview Shop</span>
                <svg class="w-4 h-4 text-gray-400 ml-1 group-hover:text-blue-600 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                </svg>
              </div>
              <span class="text-xs text-gray-500 font-mono bg-gray-100 px-2 py-1 rounded group-hover:bg-blue-50 group-hover:text-blue-600 transition-colors">
                {{ previewUrl }}
              </span>
            </button>
            
            <button 
              @click="resetCustomization"
              class="flex items-center justify-center p-3 border border-gray-300 rounded-lg hover:bg-gray-50 transition duration-150 ease-in-out"
            >
              <svg class="w-5 h-5 text-gray-500 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
              </svg>
              <span class="text-sm font-medium text-gray-700">Reset to Default</span>
            </button>
            
            <button 
              @click="exportTheme"
              class="flex items-center justify-center p-3 border border-gray-300 rounded-lg hover:bg-gray-50 transition duration-150 ease-in-out"
            >
              <svg class="w-5 h-5 text-gray-500 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
              </svg>
              <span class="text-sm font-medium text-gray-700">Export Theme</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import themeService from '@/services/theme'

const router = useRouter()
const shopStore = useShopStore()

// State
const loading = ref(false)
const customizationData = ref(null)
const error = ref('')

// Get active shop ID
const activeShop = computed(() => shopStore.active)
const shopId = computed(() => activeShop.value?.id)

// Compute preview URL based on environment and shop slug
const previewUrl = computed(() => {
  if (!activeShop.value?.slug) return 'No shop selected'

  const host = window.location.hostname
  const isLocalHost = host === 'localhost' || host === '127.0.0.1'

  if (isLocalHost) {
    // For localhost development: shopslug.localhost:5175
    return `${activeShop.value.slug}.localhost:5175`
  } else {
    // For production: shopslug.shop24.sbs
    return `${activeShop.value.slug}.shop24.sbs`
  }
})

// Load customization data on mount (align with OrdersPage pattern)
onMounted(async () => {
  try {
    if (!shopId.value) {
      const ensured = await shopStore.ensureActiveShop()
      if (!ensured?.id) {
        router.replace({ name: 'ShopSelection' })
        return
      }
    }
    await loadCustomizationData()
  } catch (e) {
    console.error('[CustomizationPage] Initialization failed:', e)
  }
})

// Load customization data
async function loadCustomizationData() {
  if (!shopId.value) return
  
  loading.value = true
  error.value = ''
  
  try {
    const response = await themeService.getCustomization(shopId.value)
    customizationData.value = response.customization
  } catch (err) {
    error.value = 'Failed to load customization data'
    console.error('Error loading customization:', err)
  } finally {
    loading.value = false
  }
}

// Navigation functions
function goToTheme() {
  router.push('/dashboard/customization/theme')
}

function goToLayout() {
  router.push('/dashboard/customization/layout')
}

function goToTypography() {
  router.push('/dashboard/customization/typography')
}

function goToCustomCSS() {
  router.push('/dashboard/customization/custom-css')
}

function goToMobile() {
  router.push('/dashboard/customization/mobile')
}

function goToSEO() {
  router.push('/dashboard/customization/seo')
}

// Quick action functions
function previewShop() {
  if (!activeShop.value?.slug) {
    console.warn('No shop slug available for preview')
    return
  }

  const host = window.location.hostname
  const isLocalHost = host === 'localhost' || host === '127.0.0.1'

  let previewUrl
  if (isLocalHost) {
    // For localhost development: http://shopslug.localhost:5175
    previewUrl = `http://${activeShop.value.slug}.localhost:5175`
  } else {
    // For production: http://shopslug.shop24.sbs
    previewUrl = `http://${activeShop.value.slug}.shop24.sbs`
  }

  console.log('Opening shop preview:', previewUrl)
  window.open(previewUrl, '_blank')
}

async function resetCustomization() {
  if (!shopId.value) return
  
  loading.value = true
  error.value = ''
  
  try {
    await themeService.resetToDefaults(shopId.value)
    await loadCustomizationData() // Reload data
    console.log('Customization reset to default')
  } catch (err) {
    error.value = 'Failed to reset customization'
    console.error('Error resetting customization:', err)
  } finally {
    loading.value = false
  }
}

async function exportTheme() {
  if (!customizationData.value) return
  
  const themeExport = {
    name: `${activeShop.value?.name || 'My Shop'} Theme`,
    config: customizationData.value,
    exportedAt: new Date().toISOString()
  }
  
  // Create and download JSON file
  const dataStr = JSON.stringify(themeExport, null, 2)
  const dataBlob = new Blob([dataStr], { type: 'application/json' })
  const url = URL.createObjectURL(dataBlob)
  
  const link = document.createElement('a')
  link.href = url
  link.download = `${activeShop.value?.name || 'my-shop'}-theme.json`
  link.click()
  
  URL.revokeObjectURL(url)
  console.log('Theme exported successfully')
}
</script>

<style scoped>
/* Custom styles for enhanced design */
@media (max-width: 640px) {
  .min-w-0 { min-width: 0 !important; }
}

/* Enhanced hover effects */
.group:hover .group-hover\:scale-105 {
  transform: scale(1.05);
}

/* Smooth transitions */
* {
  transition-property: color, background-color, border-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
}

/* Custom scrollbar for better UX */
::-webkit-scrollbar {
  width: 4px;
}

::-webkit-scrollbar-track {
  background: transparent;
}

::-webkit-scrollbar-thumb {
  background: #d1d5db;
  border-radius: 2px;
}

::-webkit-scrollbar-thumb:hover {
  background: #9ca3af;
}
</style>
