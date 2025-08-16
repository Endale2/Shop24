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
          <div class="w-10 h-10 bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg flex items-center justify-center mr-3 shadow-sm">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 5a1 1 0 011-1h14a1 1 0 011 1v2a1 1 0 01-1 1H5a1 1 0 01-1-1V5zM4 13a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H5a1 1 0 01-1-1v-6zM16 13a1 1 0 011-1h2a1 1 0 011 1v6a1 1 0 01-1 1h-2a1 1 0 01-1-1v-6z" />
            </svg>
          </div>
          <div>
            <h1 class="text-2xl sm:text-3xl font-bold text-gray-900 tracking-tight">
              Layout Designer
            </h1>
            <p class="text-sm text-gray-600 mt-1">
              Design your storefront layout and page structure
            </p>
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        
        <!-- Layout Components Panel -->
        <div class="lg:col-span-1 space-y-6">
          
          <!-- Available Components -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-lg font-semibold text-gray-900">Components</h2>
              <p class="text-sm text-gray-600">Drag to add to your layout</p>
            </div>
            <div class="p-4 space-y-3">
              <div 
                v-for="component in availableComponents"
                :key="component.id"
                draggable="true"
                @dragstart="dragStart(component)"
                class="p-3 border border-gray-200 rounded-lg cursor-move hover:border-blue-300 hover:bg-blue-50 transition duration-150 ease-in-out group"
              >
                <div class="flex items-center space-x-3">
                  <div class="w-8 h-8 bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg flex items-center justify-center group-hover:scale-105 transition-transform duration-200">
                    <component :is="component.icon" class="w-4 h-4 text-white" />
                  </div>
                  <div>
                    <h3 class="text-sm font-medium text-gray-900">{{ component.name }}</h3>
                    <p class="text-xs text-gray-600">{{ component.description }}</p>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Layout Settings -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-lg font-semibold text-gray-900">Layout Settings</h2>
              <p class="text-sm text-gray-600">Configure layout options</p>
            </div>
            <div class="p-4 space-y-4">
              
              <!-- Container Width -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Container Width</label>
                <select v-model="layoutSettings.containerWidth" class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-blue-500 focus:border-blue-500">
                  <option value="full">Full Width</option>
                  <option value="boxed">Boxed (1200px)</option>
                  <option value="narrow">Narrow (960px)</option>
                </select>
              </div>

              <!-- Header Style -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Header Style</label>
                <select v-model="layoutSettings.headerStyle" class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-blue-500 focus:border-blue-500">
                  <option value="classic">Classic</option>
                  <option value="centered">Centered</option>
                  <option value="minimal">Minimal</option>
                </select>
              </div>

              <!-- Sidebar Position -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Sidebar Position</label>
                <select v-model="layoutSettings.sidebarPosition" class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-blue-500 focus:border-blue-500">
                  <option value="none">No Sidebar</option>
                  <option value="left">Left</option>
                  <option value="right">Right</option>
                </select>
              </div>

              <!-- Grid Columns -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Product Grid Columns</label>
                <select v-model="layoutSettings.gridColumns" class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-blue-500 focus:border-blue-500">
                  <option value="2">2 Columns</option>
                  <option value="3">3 Columns</option>
                  <option value="4">4 Columns</option>
                  <option value="5">5 Columns</option>
                </select>
              </div>

            </div>
          </div>

        </div>

        <!-- Layout Builder -->
        <div class="lg:col-span-2 space-y-6">
          
          <!-- Canvas -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50 flex items-center justify-between">
              <div>
                <h2 class="text-lg font-semibold text-gray-900">Layout Canvas</h2>
                <p class="text-sm text-gray-600">Drop components here to build your layout</p>
              </div>
              <div class="flex space-x-2">
                <button 
                  @click="viewMode = 'desktop'"
                  :class="['p-2 rounded-lg transition duration-150 ease-in-out', viewMode === 'desktop' ? 'bg-blue-600 text-white' : 'text-gray-500 hover:bg-gray-100']"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                  </svg>
                </button>
                <button 
                  @click="viewMode = 'tablet'"
                  :class="['p-2 rounded-lg transition duration-150 ease-in-out', viewMode === 'tablet' ? 'bg-blue-600 text-white' : 'text-gray-500 hover:bg-gray-100']"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 18h.01M7 21h10a1 1 0 001-1V4a1 1 0 00-1-1H7a1 1 0 00-1 1v16a1 1 0 001 1z" />
                  </svg>
                </button>
                <button 
                  @click="viewMode = 'mobile'"
                  :class="['p-2 rounded-lg transition duration-150 ease-in-out', viewMode === 'mobile' ? 'bg-blue-600 text-white' : 'text-gray-500 hover:bg-gray-100']"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 18h.01M8 21h8a1 1 0 001-1V4a1 1 0 00-1-1H8a1 1 0 00-1 1v16a1 1 0 001 1z" />
                  </svg>
                </button>
              </div>
            </div>
            
            <!-- Canvas Area -->
            <div class="p-6">
              <div 
                :class="[
                  'mx-auto bg-gray-50 border-2 border-dashed border-gray-300 rounded-lg min-h-96 transition-all duration-300',
                  viewMode === 'desktop' ? 'w-full' : viewMode === 'tablet' ? 'w-3/4' : 'w-1/2'
                ]"
                @dragover.prevent
                @drop="dropComponent"
              >
                
                <!-- Layout Header -->
                <div class="bg-white border-b border-gray-200 p-4">
                  <div class="flex items-center justify-between">
                    <div class="text-lg font-semibold text-gray-900">Your Shop Name</div>
                    <div class="flex space-x-2">
                      <div class="w-6 h-6 bg-gray-300 rounded"></div>
                      <div class="w-6 h-6 bg-gray-300 rounded"></div>
                    </div>
                  </div>
                </div>

                <!-- Layout Content -->
                <div class="p-4 space-y-4">
                  
                  <!-- Dropped Components -->
                  <div
                    v-if="layoutComponents.length === 0"
                    class="text-center py-12 text-gray-400"
                  >
                    <svg class="w-12 h-12 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                    </svg>
                    <p class="text-sm">Drag components here to build your layout</p>
                  </div>
                  
                  <div
                    v-for="(component, index) in layoutComponents"
                    :key="index"
                    class="relative group bg-white border border-gray-200 rounded-lg p-4 hover:border-blue-300 transition duration-150 ease-in-out"
                  >
                    <!-- Component Content -->
                    <div class="flex items-center space-x-3">
                      <component :is="component.icon" class="w-6 h-6 text-blue-600" />
                      <div>
                        <h3 class="font-medium text-gray-900">{{ component.name }}</h3>
                        <p class="text-sm text-gray-600">{{ component.description }}</p>
                      </div>
                    </div>
                    
                    <!-- Remove Button -->
                    <button
                      @click="removeComponent(index)"
                      class="absolute top-2 right-2 w-6 h-6 bg-red-100 text-red-600 rounded-full opacity-0 group-hover:opacity-100 transition-opacity duration-200 hover:bg-red-200"
                    >
                      <svg class="w-3 h-3 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                      </svg>
                    </button>
                  </div>

                </div>
              </div>
            </div>
          </div>

          <!-- Layout Presets -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-lg font-semibold text-gray-900">Layout Presets</h2>
              <p class="text-sm text-gray-600">Quick start with pre-designed layouts</p>
            </div>
            <div class="p-4">
              <div class="grid grid-cols-2 md:grid-cols-3 gap-4">
                <button
                  v-for="preset in layoutPresets"
                  :key="preset.name"
                  @click="applyPreset(preset)"
                  class="p-3 border border-gray-200 rounded-lg hover:border-blue-300 transition duration-150 ease-in-out group"
                >
                  <div class="w-full h-20 bg-gray-100 rounded mb-2 flex items-center justify-center">
                    <component :is="preset.icon" class="w-8 h-8 text-gray-400 group-hover:text-blue-600" />
                  </div>
                  <p class="text-sm font-medium text-gray-900 group-hover:text-blue-700">{{ preset.name }}</p>
                </button>
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
              @click="resetLayout"
              :disabled="saving"
              class="px-4 py-2.5 bg-white border border-gray-300 rounded-lg text-gray-700 text-sm font-medium hover:bg-gray-50 transition duration-150 ease-in-out disabled:opacity-50"
            >
              Reset
            </button>
            <button
              @click="saveLayout"
              :disabled="saving"
              class="inline-flex items-center px-6 py-2.5 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 transition duration-150 ease-in-out disabled:opacity-50 shadow-sm"
            >
              <svg v-if="saving" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              <svg v-else class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
              {{ saving ? 'Saving...' : 'Save Layout' }}
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
const viewMode = ref('desktop')
const draggedComponent = ref(null)
const layoutComponents = ref([])

// Get active shop ID
const activeShop = computed(() => shopStore.active)
const shopId = computed(() => activeShop.value?.id)

// Layout Settings
const layoutSettings = reactive({
  containerWidth: 'boxed',
  headerStyle: 'classic',
  sidebarPosition: 'none',
  gridColumns: '3'
})

// Available Components
const availableComponents = [
  {
    id: 'hero',
    name: 'Hero Banner',
    description: 'Large banner with image and text',
    icon: 'div'
  },
  {
    id: 'products',
    name: 'Product Grid',
    description: 'Grid layout for products',
    icon: 'div'
  },
  {
    id: 'features',
    name: 'Feature Section',
    description: 'Highlight key features',
    icon: 'div'
  },
  {
    id: 'testimonials',
    name: 'Testimonials',
    description: 'Customer reviews',
    icon: 'div'
  },
  {
    id: 'newsletter',
    name: 'Newsletter',
    description: 'Email signup form',
    icon: 'div'
  }
]

// Layout Presets
const layoutPresets = [
  {
    name: 'Classic',
    icon: 'div',
    components: ['hero', 'products', 'features']
  },
  {
    name: 'Modern',
    icon: 'div',
    components: ['hero', 'features', 'products', 'testimonials']
  },
  {
    name: 'Minimal',
    icon: 'div',
    components: ['products', 'newsletter']
  }
]

// Load layout data on mount
onMounted(async () => {
  if (shopId.value) {
    await loadLayoutData()
  }
})

// Load current layout configuration
async function loadLayoutData() {
  if (!shopId.value) return
  
  loading.value = true
  error.value = ''
  
  try {
    const response = await themeService.getCustomization(shopId.value)
    const layout = response.customization?.layout || {}
    
    // Update layout settings with loaded data
    layoutSettings.containerWidth = layout.containerWidth || 'boxed'
    layoutSettings.headerStyle = layout.headerStyle || 'classic'
    layoutSettings.sidebarPosition = layout.sidebarPosition || 'none'
    layoutSettings.gridColumns = layout.gridColumns || '3'
    
    // Note: Component layout would need additional storage in backend
    // For now, keep components empty on load
  } catch (err) {
    error.value = 'Failed to load layout data'
    console.error('Error loading layout:', err)
  } finally {
    loading.value = false
  }
}

// Functions

function dragStart(component) {
  draggedComponent.value = component
}

function dropComponent(event) {
  event.preventDefault()
  if (draggedComponent.value) {
    layoutComponents.value.push({ ...draggedComponent.value })
    draggedComponent.value = null
  }
}

function removeComponent(index) {
  layoutComponents.value.splice(index, 1)
}

function applyPreset(preset) {
  layoutComponents.value = preset.components.map(componentId => 
    availableComponents.find(comp => comp.id === componentId)
  ).filter(Boolean)
}

function resetLayout() {
  const defaultLayout = themeService.getDefaultThemeConfig().layout
  layoutComponents.value = []
  layoutSettings.containerWidth = defaultLayout.containerWidth
  layoutSettings.headerStyle = defaultLayout.headerStyle
  layoutSettings.sidebarPosition = defaultLayout.sidebarPosition
  layoutSettings.gridColumns = defaultLayout.gridColumns
}

async function saveLayout() {
  if (!shopId.value) return
  
  saving.value = true
  error.value = ''
  
  try {
    const layout = {
      containerWidth: layoutSettings.containerWidth,
      headerStyle: layoutSettings.headerStyle,
      sidebarPosition: layoutSettings.sidebarPosition,
      gridColumns: layoutSettings.gridColumns
      // Note: Component layout could be saved here when backend supports it
    }
    
    await themeService.saveLayout(shopId.value, layout)
    lastSaved.value = new Date().toLocaleString()
    console.log('Layout saved successfully')
  } catch (err) {
    error.value = 'Failed to save layout'
    console.error('Error saving layout:', err)
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

/* Drag and drop styles */
[draggable="true"] {
  cursor: move;
}

[draggable="true"]:hover {
  transform: scale(1.02);
}
</style>
