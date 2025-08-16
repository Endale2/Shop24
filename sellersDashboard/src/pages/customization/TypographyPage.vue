<template>
  <div class="min-h-full bg-gray-50">
    <div class="p-4 sm:p-6 lg:p-8">
      
      <!-- Header Section -->
      <div class="mb-6 sm:mb-8">
        <div class="flex items-center mb-3">
          <BackButton
            :to="'/dashboard/customization'"
            variant="rounded"
            :show-text="false"
            size="sm"
            class="mr-3"
          />
          <div class="w-10 h-10 bg-gradient-to-br from-indigo-500 to-indigo-600 rounded-lg flex items-center justify-center mr-3 shadow-sm">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
            </svg>
          </div>
          <div>
            <h1 class="text-2xl sm:text-3xl font-bold text-gray-900 tracking-tight">
              Typography Settings
            </h1>
            <p class="text-sm text-gray-600 mt-1">
              Configure fonts and text styling for your shop
            </p>
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        
        <!-- Typography Settings -->
        <div class="space-y-6">
          
          <!-- Font Selection Card -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-lg font-semibold text-gray-900">Font Selection</h2>
              <p class="text-sm text-gray-600">Choose fonts for different text elements</p>
            </div>
            <div class="p-6 space-y-6">
              
              <!-- Heading Font -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Heading Font</label>
                <select v-model="typography.headingFont" class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-indigo-500 focus:border-indigo-500">
                  <option value="Inter">Inter</option>
                  <option value="Poppins">Poppins</option>
                  <option value="Montserrat">Montserrat</option>
                  <option value="Roboto">Roboto</option>
                  <option value="Open Sans">Open Sans</option>
                  <option value="Playfair Display">Playfair Display</option>
                  <option value="Merriweather">Merriweather</option>
                </select>
              </div>

              <!-- Body Font -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Body Font</label>
                <select v-model="typography.bodyFont" class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-indigo-500 focus:border-indigo-500">
                  <option value="Inter">Inter</option>
                  <option value="Poppins">Poppins</option>
                  <option value="Montserrat">Montserrat</option>
                  <option value="Roboto">Roboto</option>
                  <option value="Open Sans">Open Sans</option>
                  <option value="Source Sans Pro">Source Sans Pro</option>
                  <option value="Lato">Lato</option>
                </select>
              </div>

            </div>
          </div>

          <!-- Font Sizes Card -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-lg font-semibold text-gray-900">Font Sizes</h2>
              <p class="text-sm text-gray-600">Set size for different text elements</p>
            </div>
            <div class="p-6 space-y-6">
              
              <!-- Base Font Size -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Base Font Size</label>
                <div class="flex items-center space-x-3">
                  <input
                    v-model="typography.baseFontSize"
                    type="range"
                    min="12"
                    max="20"
                    class="flex-1"
                  />
                  <span class="text-sm font-mono text-gray-600 bg-gray-100 px-2 py-1 rounded">{{ typography.baseFontSize }}px</span>
                </div>
              </div>

              <!-- Heading Sizes -->
              <div class="space-y-4">
                <h3 class="text-sm font-medium text-gray-700">Heading Sizes</h3>
                
                <div>
                  <label class="block text-xs font-medium text-gray-600 mb-1">H1 (Main Headings)</label>
                  <div class="flex items-center space-x-3">
                    <input
                      v-model="typography.h1Size"
                      type="range"
                      min="24"
                      max="48"
                      class="flex-1"
                    />
                    <span class="text-sm font-mono text-gray-600 bg-gray-100 px-2 py-1 rounded">{{ typography.h1Size }}px</span>
                  </div>
                </div>

                <div>
                  <label class="block text-xs font-medium text-gray-600 mb-1">H2 (Section Headings)</label>
                  <div class="flex items-center space-x-3">
                    <input
                      v-model="typography.h2Size"
                      type="range"
                      min="20"
                      max="36"
                      class="flex-1"
                    />
                    <span class="text-sm font-mono text-gray-600 bg-gray-100 px-2 py-1 rounded">{{ typography.h2Size }}px</span>
                  </div>
                </div>

                <div>
                  <label class="block text-xs font-medium text-gray-600 mb-1">H3 (Subsection Headings)</label>
                  <div class="flex items-center space-x-3">
                    <input
                      v-model="typography.h3Size"
                      type="range"
                      min="16"
                      max="28"
                      class="flex-1"
                    />
                    <span class="text-sm font-mono text-gray-600 bg-gray-100 px-2 py-1 rounded">{{ typography.h3Size }}px</span>
                  </div>
                </div>
              </div>

            </div>
          </div>

          <!-- Typography Style Card -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-lg font-semibold text-gray-900">Typography Style</h2>
              <p class="text-sm text-gray-600">Configure text styling options</p>
            </div>
            <div class="p-6 space-y-6">
              
              <!-- Line Height -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Line Height</label>
                <div class="flex items-center space-x-3">
                  <input
                    v-model="typography.lineHeight"
                    type="range"
                    min="1.2"
                    max="2.0"
                    step="0.1"
                    class="flex-1"
                  />
                  <span class="text-sm font-mono text-gray-600 bg-gray-100 px-2 py-1 rounded">{{ typography.lineHeight }}</span>
                </div>
              </div>

              <!-- Letter Spacing -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Letter Spacing</label>
                <div class="flex items-center space-x-3">
                  <input
                    v-model="typography.letterSpacing"
                    type="range"
                    min="-0.05"
                    max="0.1"
                    step="0.01"
                    class="flex-1"
                  />
                  <span class="text-sm font-mono text-gray-600 bg-gray-100 px-2 py-1 rounded">{{ typography.letterSpacing }}em</span>
                </div>
              </div>

              <!-- Font Weight -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Default Font Weight</label>
                <select v-model="typography.fontWeight" class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-indigo-500 focus:border-indigo-500">
                  <option value="300">Light (300)</option>
                  <option value="400">Regular (400)</option>
                  <option value="500">Medium (500)</option>
                  <option value="600">Semi Bold (600)</option>
                  <option value="700">Bold (700)</option>
                </select>
              </div>

            </div>
          </div>

        </div>

        <!-- Typography Preview -->
        <div class="space-y-6">
          
          <!-- Live Preview Card -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden sticky top-6">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-lg font-semibold text-gray-900">Typography Preview</h2>
              <p class="text-sm text-gray-600">See how your typography looks</p>
            </div>
            <div class="p-6">
              <!-- Typography Examples -->
              <div class="space-y-6">
                
                <!-- H1 Example -->
                <div>
                  <h1 
                    :style="{
                      fontFamily: typography.headingFont,
                      fontSize: typography.h1Size + 'px',
                      lineHeight: typography.lineHeight,
                      letterSpacing: typography.letterSpacing + 'em',
                      fontWeight: 700
                    }"
                    class="text-gray-900"
                  >
                    Main Heading (H1)
                  </h1>
                  <p class="text-xs text-gray-500 mt-1">Used for page titles and main headings</p>
                </div>

                <!-- H2 Example -->
                <div>
                  <h2 
                    :style="{
                      fontFamily: typography.headingFont,
                      fontSize: typography.h2Size + 'px',
                      lineHeight: typography.lineHeight,
                      letterSpacing: typography.letterSpacing + 'em',
                      fontWeight: 600
                    }"
                    class="text-gray-900"
                  >
                    Section Heading (H2)
                  </h2>
                  <p class="text-xs text-gray-500 mt-1">Used for section titles</p>
                </div>

                <!-- H3 Example -->
                <div>
                  <h3 
                    :style="{
                      fontFamily: typography.headingFont,
                      fontSize: typography.h3Size + 'px',
                      lineHeight: typography.lineHeight,
                      letterSpacing: typography.letterSpacing + 'em',
                      fontWeight: 600
                    }"
                    class="text-gray-900"
                  >
                    Subsection Heading (H3)
                  </h3>
                  <p class="text-xs text-gray-500 mt-1">Used for subsection titles</p>
                </div>

                <!-- Body Text Example -->
                <div>
                  <p 
                    :style="{
                      fontFamily: typography.bodyFont,
                      fontSize: typography.baseFontSize + 'px',
                      lineHeight: typography.lineHeight,
                      letterSpacing: typography.letterSpacing + 'em',
                      fontWeight: typography.fontWeight
                    }"
                    class="text-gray-700"
                  >
                    This is body text. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris.
                  </p>
                  <p class="text-xs text-gray-500 mt-1">Used for regular content and descriptions</p>
                </div>

                <!-- Small Text Example -->
                <div>
                  <p 
                    :style="{
                      fontFamily: typography.bodyFont,
                      fontSize: (typography.baseFontSize - 2) + 'px',
                      lineHeight: typography.lineHeight,
                      letterSpacing: typography.letterSpacing + 'em',
                      fontWeight: typography.fontWeight
                    }"
                    class="text-gray-600"
                  >
                    Small text for captions and fine print
                  </p>
                  <p class="text-xs text-gray-500 mt-1">Used for captions and additional information</p>
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
              @click="resetTypography"
              :disabled="saving"
              class="px-4 py-2.5 bg-white border border-gray-300 rounded-lg text-gray-700 text-sm font-medium hover:bg-gray-50 transition duration-150 ease-in-out disabled:opacity-50"
            >
              Reset
            </button>
            <button
              @click="saveTypography"
              :disabled="saving"
              class="inline-flex items-center px-6 py-2.5 bg-indigo-600 text-white text-sm font-medium rounded-lg hover:bg-indigo-700 transition duration-150 ease-in-out disabled:opacity-50 shadow-sm"
            >
              <svg v-if="saving" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              <svg v-else class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
              {{ saving ? 'Saving...' : 'Save Typography' }}
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
import BackButton from '@/components/BackButton.vue'

const router = useRouter()
const shopStore = useShopStore()

// State
const saving = ref(false)
const loading = ref(false)
const lastSaved = ref('')
const error = ref('')

// Get active shop ID
const activeShop = computed(() => shopStore.active)
const shopId = computed(() => activeShop.value?.id)

// Typography configuration
const typography = reactive({
  headingFont: 'Inter',
  bodyFont: 'Inter',
  baseFontSize: 16,
  h1Size: 32,
  h2Size: 24,
  h3Size: 20,
  lineHeight: 1.5,
  letterSpacing: 0,
  fontWeight: '400'
})

// Load typography data on mount (align with OrdersPage pattern)
onMounted(async () => {
  try {
    if (!shopId.value) {
      const ensured = await shopStore.ensureActiveShop()
      if (!ensured?.id) {
        router.replace({ name: 'ShopSelection' })
        return
      }
    }
    await loadTypographyData()
  } catch (e) {
    console.error('[TypographyPage] Initialization failed:', e)
  }
})

// Load current typography configuration
async function loadTypographyData() {
  if (!shopId.value) return
  
  loading.value = true
  error.value = ''
  
  try {
    const response = await themeService.getCustomization(shopId.value)
    const fonts = response.customization?.fonts || {}
    
    // Update typography object with loaded data
    typography.headingFont = fonts.heading || 'Inter'
    typography.bodyFont = fonts.body || 'Inter'
    // Note: Extended typography settings would need to be stored in a more detailed way
    // For now, we'll keep the defaults for size and spacing settings
  } catch (err) {
    error.value = 'Failed to load typography data'
    console.error('Error loading typography:', err)
  } finally {
    loading.value = false
  }
}

// Functions

function resetTypography() {
  const defaultFonts = themeService.getDefaultThemeConfig().fonts
  typography.headingFont = defaultFonts.heading
  typography.bodyFont = defaultFonts.body
  typography.baseFontSize = 16
  typography.h1Size = 32
  typography.h2Size = 24
  typography.h3Size = 20
  typography.lineHeight = 1.5
  typography.letterSpacing = 0
  typography.fontWeight = '400'
}

async function saveTypography() {
  if (!shopId.value) return
  
  saving.value = true
  error.value = ''
  
  try {
    const fonts = {
      heading: typography.headingFont,
      body: typography.bodyFont
      // Note: Extended typography settings could be saved here
      // when the backend model is extended to support more detailed typography config
    }
    
    await themeService.saveTypography(shopId.value, fonts)
    lastSaved.value = new Date().toLocaleString()
    console.log('Typography saved successfully')
  } catch (err) {
    error.value = 'Failed to save typography'
    console.error('Error saving typography:', err)
  } finally {
    saving.value = false
  }
}
</script>

<style scoped>
/* Smooth transitions */
* {
  transition-property: color, background-color, border-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
}
</style>
