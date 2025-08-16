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
          <div class="w-10 h-10 bg-gradient-to-br from-pink-500 to-pink-600 rounded-lg flex items-center justify-center mr-3 shadow-sm">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 18h.01M8 21h8a1 1 0 001-1V4a1 1 0 00-1-1H8a1 1 0 00-1 1v16a1 1 0 001 1z" />
            </svg>
          </div>
          <div>
            <h1 class="text-2xl sm:text-3xl font-bold text-gray-900 tracking-tight">
              Mobile Optimization
            </h1>
            <p class="text-sm text-gray-600 mt-1">
              Optimize your shop's mobile experience and responsive design
            </p>
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        
        <!-- Mobile Settings -->
        <div class="space-y-6">
          
          <!-- Mobile Navigation Card -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-lg font-semibold text-gray-900">Mobile Navigation</h2>
              <p class="text-sm text-gray-600">Configure mobile menu and navigation</p>
            </div>
            <div class="p-6 space-y-6">
              
              <!-- Menu Style -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Mobile Menu Style</label>
                <select v-model="mobileSettings.menuStyle" class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-pink-500 focus:border-pink-500">
                  <option value="hamburger">Hamburger Menu</option>
                  <option value="bottom-tabs">Bottom Tab Navigation</option>
                  <option value="slide-out">Slide-out Drawer</option>
                </select>
              </div>

              <!-- Menu Position -->
              <div v-if="mobileSettings.menuStyle === 'hamburger'">
                <label class="block text-sm font-medium text-gray-700 mb-2">Menu Button Position</label>
                <select v-model="mobileSettings.menuPosition" class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-pink-500 focus:border-pink-500">
                  <option value="left">Left</option>
                  <option value="right">Right</option>
                </select>
              </div>

              <!-- Sticky Header -->
              <div class="flex items-center justify-between">
                <div>
                  <label class="text-sm font-medium text-gray-700">Sticky Header</label>
                  <p class="text-xs text-gray-500">Keep header visible while scrolling</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input v-model="mobileSettings.stickyHeader" type="checkbox" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-pink-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-pink-600"></div>
                </label>
              </div>

            </div>
          </div>

          <!-- Touch Interactions Card -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-lg font-semibold text-gray-900">Touch Interactions</h2>
              <p class="text-sm text-gray-600">Configure touch gestures and interactions</p>
            </div>
            <div class="p-6 space-y-6">
              
              <!-- Swipe Gestures -->
              <div class="flex items-center justify-between">
                <div>
                  <label class="text-sm font-medium text-gray-700">Swipe Navigation</label>
                  <p class="text-xs text-gray-500">Enable swipe between product images</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input v-model="mobileSettings.swipeNavigation" type="checkbox" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-pink-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-pink-600"></div>
                </label>
              </div>

              <!-- Pull to Refresh -->
              <div class="flex items-center justify-between">
                <div>
                  <label class="text-sm font-medium text-gray-700">Pull to Refresh</label>
                  <p class="text-xs text-gray-500">Allow pull-down to refresh content</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input v-model="mobileSettings.pullToRefresh" type="checkbox" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-pink-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-pink-600"></div>
                </label>
              </div>

              <!-- Touch Feedback -->
              <div class="flex items-center justify-between">
                <div>
                  <label class="text-sm font-medium text-gray-700">Touch Feedback</label>
                  <p class="text-xs text-gray-500">Visual feedback on touch interactions</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input v-model="mobileSettings.touchFeedback" type="checkbox" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-pink-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-pink-600"></div>
                </label>
              </div>

            </div>
          </div>

          <!-- Performance Settings Card -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-lg font-semibold text-gray-900">Performance Settings</h2>
              <p class="text-sm text-gray-600">Optimize for mobile performance</p>
            </div>
            <div class="p-6 space-y-6">
              
              <!-- Image Compression -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Image Quality</label>
                <div class="flex items-center space-x-3">
                  <input
                    v-model="mobileSettings.imageQuality"
                    type="range"
                    min="60"
                    max="100"
                    class="flex-1"
                  />
                  <span class="text-sm font-mono text-gray-600 bg-gray-100 px-2 py-1 rounded">{{ mobileSettings.imageQuality }}%</span>
                </div>
                <p class="text-xs text-gray-500 mt-1">Lower quality = faster loading</p>
              </div>

              <!-- Lazy Loading -->
              <div class="flex items-center justify-between">
                <div>
                  <label class="text-sm font-medium text-gray-700">Lazy Loading</label>
                  <p class="text-xs text-gray-500">Load images as they become visible</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input v-model="mobileSettings.lazyLoading" type="checkbox" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-pink-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-pink-600"></div>
                </label>
              </div>

              <!-- Preload Critical Resources -->
              <div class="flex items-center justify-between">
                <div>
                  <label class="text-sm font-medium text-gray-700">Preload Critical Resources</label>
                  <p class="text-xs text-gray-500">Preload important assets for faster loading</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input v-model="mobileSettings.preloadCritical" type="checkbox" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-pink-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-pink-600"></div>
                </label>
              </div>

            </div>
          </div>

        </div>

        <!-- Mobile Preview -->
        <div class="space-y-6">
          
          <!-- Device Preview Card -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden sticky top-6">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <div class="flex items-center justify-between">
                <div>
                  <h2 class="text-lg font-semibold text-gray-900">Mobile Preview</h2>
                  <p class="text-sm text-gray-600">See how your shop looks on mobile</p>
                </div>
                <div class="flex space-x-2">
                  <button 
                    @click="previewDevice = 'phone'"
                    :class="['p-2 rounded-lg transition duration-150 ease-in-out', previewDevice === 'phone' ? 'bg-pink-600 text-white' : 'text-gray-500 hover:bg-gray-100']"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 18h.01M8 21h8a1 1 0 001-1V4a1 1 0 00-1-1H8a1 1 0 00-1 1v16a1 1 0 001 1z" />
                    </svg>
                  </button>
                  <button 
                    @click="previewDevice = 'tablet'"
                    :class="['p-2 rounded-lg transition duration-150 ease-in-out', previewDevice === 'tablet' ? 'bg-pink-600 text-white' : 'text-gray-500 hover:bg-gray-100']"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 18h.01M7 21h10a1 1 0 001-1V4a1 1 0 00-1-1H7a1 1 0 00-1 1v16a1 1 0 001 1z" />
                    </svg>
                  </button>
                </div>
              </div>
            </div>
            <div class="p-6 flex justify-center">
              <!-- Mobile Device Frame -->
              <div 
                :class="[
                  'bg-gray-900 rounded-3xl p-2 shadow-2xl',
                  previewDevice === 'phone' ? 'w-72 h-96' : 'w-80 h-80'
                ]"
              >
                <div class="bg-white rounded-2xl h-full overflow-hidden relative">
                  
                  <!-- Device Notch (for phone) -->
                  <div v-if="previewDevice === 'phone'" class="absolute top-0 left-1/2 transform -translate-x-1/2 w-20 h-6 bg-gray-900 rounded-b-2xl z-10"></div>
                  
                  <!-- Mock Mobile Content -->
                  <div class="h-full flex flex-col">
                    
                    <!-- Header -->
                    <div 
                      :class="[
                        'flex items-center justify-between p-3 border-b border-gray-200',
                        mobileSettings.stickyHeader ? 'bg-white shadow-sm' : 'bg-gray-50',
                        previewDevice === 'phone' ? 'pt-8' : ''
                      ]"
                    >
                      <div v-if="mobileSettings.menuStyle === 'hamburger'" class="flex items-center">
                        <div 
                          :class="[
                            'w-6 h-4 flex flex-col justify-between mr-3',
                            mobileSettings.menuPosition === 'right' ? 'order-2 ml-3 mr-0' : ''
                          ]"
                        >
                          <div class="h-0.5 bg-gray-700"></div>
                          <div class="h-0.5 bg-gray-700"></div>
                          <div class="h-0.5 bg-gray-700"></div>
                        </div>
                        <h3 class="font-semibold text-gray-900 text-sm">Your Shop</h3>
                      </div>
                      <div v-else class="w-full text-center">
                        <h3 class="font-semibold text-gray-900 text-sm">Your Shop</h3>
                      </div>
                      <div class="w-5 h-5 border border-gray-300 rounded-full"></div>
                    </div>

                    <!-- Content Area -->
                    <div class="flex-1 p-3 space-y-3">
                      <!-- Product Grid Mock -->
                      <div class="grid grid-cols-2 gap-2">
                        <div v-for="i in 4" :key="i" class="bg-gray-100 rounded-lg aspect-square flex items-center justify-center">
                          <div class="w-8 h-8 bg-gray-300 rounded"></div>
                        </div>
                      </div>
                    </div>

                    <!-- Bottom Navigation (if enabled) -->
                    <div v-if="mobileSettings.menuStyle === 'bottom-tabs'" class="border-t border-gray-200 p-2">
                      <div class="flex justify-around">
                        <div v-for="i in 4" :key="i" class="w-6 h-6 bg-gray-300 rounded"></div>
                      </div>
                    </div>

                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Mobile Performance Card -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-lg font-semibold text-gray-900">Performance Metrics</h2>
              <p class="text-sm text-gray-600">Mobile performance indicators</p>
            </div>
            <div class="p-6">
              <div class="space-y-4">
                
                <!-- Page Load Speed -->
                <div>
                  <div class="flex justify-between items-center mb-2">
                    <span class="text-sm font-medium text-gray-700">Page Load Speed</span>
                    <span class="text-sm text-green-600 font-semibold">2.3s</span>
                  </div>
                  <div class="w-full bg-gray-200 rounded-full h-2">
                    <div class="bg-green-600 h-2 rounded-full" style="width: 76%"></div>
                  </div>
                </div>

                <!-- Image Optimization -->
                <div>
                  <div class="flex justify-between items-center mb-2">
                    <span class="text-sm font-medium text-gray-700">Image Optimization</span>
                    <span class="text-sm text-yellow-600 font-semibold">68%</span>
                  </div>
                  <div class="w-full bg-gray-200 rounded-full h-2">
                    <div class="bg-yellow-500 h-2 rounded-full" style="width: 68%"></div>
                  </div>
                </div>

                <!-- Mobile Usability -->
                <div>
                  <div class="flex justify-between items-center mb-2">
                    <span class="text-sm font-medium text-gray-700">Mobile Usability</span>
                    <span class="text-sm text-green-600 font-semibold">92%</span>
                  </div>
                  <div class="w-full bg-gray-200 rounded-full h-2">
                    <div class="bg-green-600 h-2 rounded-full" style="width: 92%"></div>
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
              @click="resetMobileSettings"
              :disabled="saving"
              class="px-4 py-2.5 bg-white border border-gray-300 rounded-lg text-gray-700 text-sm font-medium hover:bg-gray-50 transition duration-150 ease-in-out disabled:opacity-50"
            >
              Reset
            </button>
            <button
              @click="saveMobileSettings"
              :disabled="saving"
              class="inline-flex items-center px-6 py-2.5 bg-pink-600 text-white text-sm font-medium rounded-lg hover:bg-pink-700 transition duration-150 ease-in-out disabled:opacity-50 shadow-sm"
            >
              <svg v-if="saving" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              <svg v-else class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
              {{ saving ? 'Saving...' : 'Save Mobile Settings' }}
            </button>
          </div>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import BackButton from '@/components/BackButton.vue'

const router = useRouter()

// State
const saving = ref(false)
const lastSaved = ref('')
const previewDevice = ref('phone')

// Mobile Settings
const mobileSettings = reactive({
  menuStyle: 'hamburger',
  menuPosition: 'left',
  stickyHeader: true,
  swipeNavigation: true,
  pullToRefresh: false,
  touchFeedback: true,
  imageQuality: 80,
  lazyLoading: true,
  preloadCritical: true
})

// Functions

function resetMobileSettings() {
  mobileSettings.menuStyle = 'hamburger'
  mobileSettings.menuPosition = 'left'
  mobileSettings.stickyHeader = true
  mobileSettings.swipeNavigation = true
  mobileSettings.pullToRefresh = false
  mobileSettings.touchFeedback = true
  mobileSettings.imageQuality = 80
  mobileSettings.lazyLoading = true
  mobileSettings.preloadCritical = true
}

async function saveMobileSettings() {
  saving.value = true
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1000))
    lastSaved.value = new Date().toLocaleString()
    console.log('Mobile settings saved:', mobileSettings)
  } catch (error) {
    console.error('Error saving mobile settings:', error)
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
