<template>
  <div class="min-h-full bg-gray-50">
    <div class="p-4 sm:p-6 lg:p-8">
      
      <!-- Header Section -->
      <div class="mb-6 sm:mb-8">
        <div class="flex items-center mb-3">
          <button 
            @click="goBack"
            class="mr-3 p-2 text-gray-500 hover:text-gray-700 hover:bg-gray-100 rounded-lg transition duration-150 ease-in-out"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
          </button>
          <div class="w-10 h-10 bg-gradient-to-br from-purple-500 to-purple-600 rounded-lg flex items-center justify-center mr-3 shadow-sm">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zM21 5a2 2 0 00-2-2h-4a2 2 0 00-2 2v12a4 4 0 004 4h4a2 2 0 002-2V5z" />
            </svg>
          </div>
          <div>
            <h1 class="text-2xl sm:text-3xl font-bold text-gray-900 tracking-tight">
              Theme Customization
            </h1>
            <p class="text-sm text-gray-600 mt-1">
              Customize your shop's colors, branding, and visual theme
            </p>
          </div>
        </div>
      </div>

      <!-- Live Preview Toggle -->
      <div class="mb-6 bg-white rounded-lg shadow-sm border border-gray-200 p-4">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-3">
            <div class="w-8 h-8 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
              </svg>
            </div>
            <div>
              <h3 class="text-sm font-semibold text-gray-900">Live Preview</h3>
              <p class="text-xs text-gray-600">See changes in real-time</p>
            </div>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input v-model="livePreview" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-green-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-green-600"></div>
          </label>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        
        <!-- Theme Settings -->
        <div class="space-y-6">
          
          <!-- Primary Colors Card -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-lg font-semibold text-gray-900">Primary Colors</h2>
              <p class="text-sm text-gray-600">Define your brand's main color palette</p>
            </div>
            <div class="p-6 space-y-6">
              
              <!-- Primary Color -->
              <div class="space-y-3">
                <label class="block text-sm font-medium text-gray-700">Primary Color</label>
                <div class="flex items-center space-x-4">
                  <input
                    v-model="theme.primaryColor"
                    type="color"
                    class="h-12 w-12 border-2 border-white rounded-xl cursor-pointer shadow-lg hover:scale-105 transition-transform duration-200"
                  />
                  <div class="flex-1">
                    <input
                      v-model="theme.primaryColor"
                      type="text"
                      class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-purple-500 focus:border-purple-500"
                      placeholder="#6366F1"
                    />
                    <p class="text-xs text-gray-500 mt-1">Used for buttons, links, and accents</p>
                  </div>
                </div>
              </div>

              <!-- Secondary Color -->
              <div class="space-y-3">
                <label class="block text-sm font-medium text-gray-700">Secondary Color</label>
                <div class="flex items-center space-x-4">
                  <input
                    v-model="theme.secondaryColor"
                    type="color"
                    class="h-12 w-12 border-2 border-white rounded-xl cursor-pointer shadow-lg hover:scale-105 transition-transform duration-200"
                  />
                  <div class="flex-1">
                    <input
                      v-model="theme.secondaryColor"
                      type="text"
                      class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-purple-500 focus:border-purple-500"
                      placeholder="#F59E0B"
                    />
                    <p class="text-xs text-gray-500 mt-1">Supporting color for highlights</p>
                  </div>
                </div>
              </div>

              <!-- Background Color -->
              <div class="space-y-3">
                <label class="block text-sm font-medium text-gray-700">Background Color</label>
                <div class="flex items-center space-x-4">
                  <input
                    v-model="theme.backgroundColor"
                    type="color"
                    class="h-12 w-12 border-2 border-white rounded-xl cursor-pointer shadow-lg hover:scale-105 transition-transform duration-200"
                  />
                  <div class="flex-1">
                    <input
                      v-model="theme.backgroundColor"
                      type="text"
                      class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-purple-500 focus:border-purple-500"
                      placeholder="#FFFFFF"
                    />
                    <p class="text-xs text-gray-500 mt-1">Main background color</p>
                  </div>
                </div>
              </div>

            </div>
          </div>

          <!-- Text Colors Card -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-lg font-semibold text-gray-900">Text Colors</h2>
              <p class="text-sm text-gray-600">Configure typography colors</p>
            </div>
            <div class="p-6 space-y-6">
              
              <!-- Heading Color -->
              <div class="space-y-3">
                <label class="block text-sm font-medium text-gray-700">Heading Color</label>
                <div class="flex items-center space-x-4">
                  <input
                    v-model="theme.headingColor"
                    type="color"
                    class="h-12 w-12 border-2 border-white rounded-xl cursor-pointer shadow-lg hover:scale-105 transition-transform duration-200"
                  />
                  <div class="flex-1">
                    <input
                      v-model="theme.headingColor"
                      type="text"
                      class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-purple-500 focus:border-purple-500"
                      placeholder="#1F2937"
                    />
                  </div>
                </div>
              </div>

              <!-- Body Text Color -->
              <div class="space-y-3">
                <label class="block text-sm font-medium text-gray-700">Body Text Color</label>
                <div class="flex items-center space-x-4">
                  <input
                    v-model="theme.bodyTextColor"
                    type="color"
                    class="h-12 w-12 border-2 border-white rounded-xl cursor-pointer shadow-lg hover:scale-105 transition-transform duration-200"
                  />
                  <div class="flex-1">
                    <input
                      v-model="theme.bodyTextColor"
                      type="text"
                      class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-purple-500 focus:border-purple-500"
                      placeholder="#6B7280"
                    />
                  </div>
                </div>
              </div>

            </div>
          </div>

          <!-- Quick Theme Presets -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-lg font-semibold text-gray-900">Quick Presets</h2>
              <p class="text-sm text-gray-600">Apply pre-designed color schemes</p>
            </div>
            <div class="p-6">
              <div class="grid grid-cols-2 gap-3">
                <button
                  v-for="preset in themePresets"
                  :key="preset.name"
                  @click="applyPreset(preset)"
                  class="p-3 border border-gray-200 rounded-lg hover:border-gray-300 transition duration-150 ease-in-out group"
                >
                  <div class="flex items-center space-x-2 mb-2">
                    <div 
                      class="w-4 h-4 rounded-full"
                      :style="{ backgroundColor: preset.primary }"
                    ></div>
                    <div 
                      class="w-4 h-4 rounded-full"
                      :style="{ backgroundColor: preset.secondary }"
                    ></div>
                  </div>
                  <p class="text-sm font-medium text-gray-900 group-hover:text-purple-700">{{ preset.name }}</p>
                </button>
              </div>
            </div>
          </div>

        </div>

        <!-- Preview Section -->
        <div class="space-y-6">
          
          <!-- Live Preview Card -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden sticky top-6">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-lg font-semibold text-gray-900">Preview</h2>
              <p class="text-sm text-gray-600">See how your theme looks</p>
            </div>
            <div class="p-6">
              <!-- Mock Storefront Preview -->
              <div 
                class="border border-gray-200 rounded-lg overflow-hidden"
                :style="{ backgroundColor: theme.backgroundColor }"
              >
                <!-- Header Preview -->
                <div 
                  class="p-4 border-b border-gray-200"
                  :style="{ borderColor: theme.primaryColor + '20' }"
                >
                  <div class="flex items-center justify-between">
                    <h3 
                      class="text-lg font-bold"
                      :style="{ color: theme.headingColor }"
                    >
                      Your Shop Name
                    </h3>
                    <div class="flex space-x-2">
                      <div 
                        class="w-6 h-6 rounded"
                        :style="{ backgroundColor: theme.primaryColor }"
                      ></div>
                      <div 
                        class="w-6 h-6 rounded"
                        :style="{ backgroundColor: theme.secondaryColor }"
                      ></div>
                    </div>
                  </div>
                </div>
                
                <!-- Content Preview -->
                <div class="p-4 space-y-4">
                  <div>
                    <h4 
                      class="font-semibold mb-2"
                      :style="{ color: theme.headingColor }"
                    >
                      Featured Products
                    </h4>
                    <p 
                      class="text-sm"
                      :style="{ color: theme.bodyTextColor }"
                    >
                      Discover our latest collection of amazing products.
                    </p>
                  </div>
                  
                  <!-- Mock Product Card -->
                  <div class="border border-gray-200 rounded-lg p-3">
                    <div class="w-full h-20 bg-gray-100 rounded mb-2"></div>
                    <h5 
                      class="font-medium text-sm"
                      :style="{ color: theme.headingColor }"
                    >
                      Sample Product
                    </h5>
                    <p 
                      class="text-xs mb-2"
                      :style="{ color: theme.bodyTextColor }"
                    >
                      Product description here
                    </p>
                    <button 
                      class="w-full py-1 px-3 rounded text-white text-xs font-medium"
                      :style="{ backgroundColor: theme.primaryColor }"
                    >
                      Add to Cart
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>

        </div>
      </div>

      <!-- Action Buttons -->
      <div class="mt-8 bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
        <div class="px-6 py-4 bg-gray-50 flex flex-col sm:flex-row justify-between items-center gap-3">
          <div class="flex items-center space-x-2">
            <svg v-if="lastSaved" class="w-4 h-4 text-green-500" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
            </svg>
            <p v-if="lastSaved" class="text-sm text-gray-600">
              Last saved: {{ lastSaved }}
            </p>
            <p v-else class="text-sm text-gray-500">
              No recent saves
            </p>
          </div>
          <div class="flex gap-3">
            <button
              @click="resetTheme"
              :disabled="saving"
              class="px-4 py-2.5 bg-white border border-gray-300 rounded-lg text-gray-700 text-sm font-medium hover:bg-gray-50 transition duration-150 ease-in-out disabled:opacity-50"
            >
              Reset
            </button>
            <button
              @click="saveTheme"
              :disabled="saving"
              class="inline-flex items-center px-6 py-2.5 bg-purple-600 text-white text-sm font-medium rounded-lg hover:bg-purple-700 transition duration-150 ease-in-out disabled:opacity-50 shadow-sm"
            >
              <svg v-if="saving" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              <svg v-else class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
              {{ saving ? 'Saving...' : 'Save Theme' }}
            </button>
          </div>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import themeService from '@/services/theme'

const router = useRouter()
const shopStore = useShopStore()

// State
const saving = ref(false)
const loading = ref(false)
const lastSaved = ref('')
const livePreview = ref(true)
const error = ref('')

// Get active shop ID
const activeShop = computed(() => shopStore.active)
const shopId = computed(() => activeShop.value?.id)

// Theme configuration
const theme = reactive({
  primaryColor: '#6366F1',
  secondaryColor: '#F59E0B',
  backgroundColor: '#FFFFFF',
  headingColor: '#1F2937',
  bodyTextColor: '#6B7280'
})

// Theme presets
const themePresets = [
  {
    name: 'Ocean Blue',
    primary: '#3B82F6',
    secondary: '#06B6D4',
    background: '#FFFFFF',
    heading: '#1E3A8A',
    body: '#475569'
  },
  {
    name: 'Forest Green',
    primary: '#10B981',
    secondary: '#34D399',
    background: '#FFFFFF',
    heading: '#064E3B',
    body: '#374151'
  },
  {
    name: 'Sunset Orange',
    primary: '#F59E0B',
    secondary: '#FB923C',
    background: '#FFFFFF',
    heading: '#92400E',
    body: '#6B7280'
  },
  {
    name: 'Royal Purple',
    primary: '#8B5CF6',
    secondary: '#A78BFA',
    background: '#FFFFFF',
    heading: '#5B21B6',
    body: '#4B5563'
  },
  {
    name: 'Rose Pink',
    primary: '#EC4899',
    secondary: '#F472B6',
    background: '#FFFFFF',
    heading: '#BE185D',
    body: '#6B7280'
  },
  {
    name: 'Dark Mode',
    primary: '#6366F1',
    secondary: '#F59E0B',
    background: '#111827',
    heading: '#F9FAFB',
    body: '#D1D5DB'
  }
]

// Load theme data on mount
onMounted(async () => {
  if (shopId.value) {
    await loadThemeData()
  }
})

// Load current theme configuration
async function loadThemeData() {
  if (!shopId.value) return
  
  loading.value = true
  error.value = ''
  
  try {
    const response = await themeService.getCustomization(shopId.value)
    const colors = response.customization?.colors || {}
    
    // Update theme object with loaded data
    theme.primaryColor = colors.primary || '#6366F1'
    theme.secondaryColor = colors.secondary || '#F59E0B'
    theme.backgroundColor = colors.background || '#FFFFFF'
    theme.headingColor = colors.heading || '#1F2937'
    theme.bodyTextColor = colors.bodyText || '#6B7280'
  } catch (err) {
    error.value = 'Failed to load theme data'
    console.error('Error loading theme:', err)
  } finally {
    loading.value = false
  }
}

// Functions
function goBack() {
  router.push('/dashboard/customization')
}

function applyPreset(preset) {
  theme.primaryColor = preset.primary
  theme.secondaryColor = preset.secondary
  theme.backgroundColor = preset.background
  theme.headingColor = preset.heading
  theme.bodyTextColor = preset.body
}

function resetTheme() {
  const defaultColors = themeService.getDefaultThemeConfig().colors
  theme.primaryColor = defaultColors.primary
  theme.secondaryColor = defaultColors.secondary
  theme.backgroundColor = defaultColors.background
  theme.headingColor = defaultColors.heading
  theme.bodyTextColor = defaultColors.bodyText
}

async function saveTheme() {
  if (!shopId.value) return
  
  saving.value = true
  error.value = ''
  
  try {
    const colors = {
      primary: theme.primaryColor,
      secondary: theme.secondaryColor,
      background: theme.backgroundColor,
      heading: theme.headingColor,
      bodyText: theme.bodyTextColor
    }
    
    await themeService.saveThemeColors(shopId.value, colors)
    lastSaved.value = new Date().toLocaleString()
    console.log('Theme saved successfully')
  } catch (err) {
    error.value = 'Failed to save theme'
    console.error('Error saving theme:', err)
  } finally {
    saving.value = false
  }
}
</script>

<style scoped>
/* Color picker enhancement */
input[type="color"] {
  appearance: none;
  -webkit-appearance: none;
  border: none;
  cursor: pointer;
}

input[type="color"]::-webkit-color-swatch-wrapper {
  padding: 0;
}

input[type="color"]::-webkit-color-swatch {
  border: none;
  border-radius: 0.75rem;
}

/* Smooth transitions */
* {
  transition-property: color, background-color, border-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
}
</style>
