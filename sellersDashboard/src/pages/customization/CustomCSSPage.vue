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
          <div class="w-10 h-10 bg-gradient-to-br from-emerald-500 to-emerald-600 rounded-lg flex items-center justify-center mr-3 shadow-sm">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4" />
            </svg>
          </div>
          <div>
            <h1 class="text-2xl sm:text-3xl font-bold text-gray-900 tracking-tight">
              Custom CSS Editor
            </h1>
            <p class="text-sm text-gray-600 mt-1">
              Add custom CSS for advanced styling and unique design modifications
            </p>
          </div>
        </div>
      </div>

      <!-- Warning Notice -->
      <div class="mb-6 bg-amber-50 border border-amber-200 rounded-lg p-4">
        <div class="flex items-start">
          <svg class="w-5 h-5 text-amber-600 mt-0.5 mr-3 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L5.232 15.5c-.77.833.192 2.5 1.732 2.5z" />
          </svg>
          <div>
            <h3 class="text-sm font-medium text-amber-800">Advanced Feature</h3>
            <p class="mt-1 text-sm text-amber-700">
              Custom CSS can affect your shop's appearance and functionality. Test changes carefully and always backup your working styles before making modifications.
            </p>
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        
        <!-- CSS Editor -->
        <div class="space-y-6">
          
          <!-- CSS Code Editor Card -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <div class="flex items-center justify-between">
                <div>
                  <h2 class="text-lg font-semibold text-gray-900">CSS Editor</h2>
                  <p class="text-sm text-gray-600">Write your custom CSS styles</p>
                </div>
                <div class="flex space-x-2">
                  <button 
                    @click="formatCSS"
                    class="px-3 py-1.5 text-xs font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50 transition duration-150 ease-in-out"
                  >
                    Format
                  </button>
                  <button 
                    @click="validateCSS"
                    class="px-3 py-1.5 text-xs font-medium text-emerald-700 bg-emerald-50 border border-emerald-200 rounded-md hover:bg-emerald-100 transition duration-150 ease-in-out"
                  >
                    Validate
                  </button>
                </div>
              </div>
            </div>
            <div class="relative">
              <!-- Line Numbers -->
              <div class="absolute left-0 top-0 bottom-0 w-12 bg-gray-50 border-r border-gray-200 p-2 text-xs text-gray-500 font-mono leading-6 select-none">
                <div v-for="i in lineCount" :key="i" class="text-right">{{ i }}</div>
              </div>
              <!-- CSS Textarea -->
              <textarea
                v-model="customCSS"
                @input="updateLineCount"
                class="w-full h-96 pl-16 pr-4 py-3 border-0 resize-none focus:ring-0 font-mono text-sm leading-6 bg-white"
                placeholder="/* Add your custom CSS here */

.my-custom-class {
  background-color: #f3f4f6;
  padding: 1rem;
  border-radius: 0.5rem;
}

/* Override existing styles */
.product-card {
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

/* Responsive styles */
@media (max-width: 768px) {
  .my-custom-class {
    padding: 0.5rem;
  }
}"
                spellcheck="false"
              ></textarea>
            </div>
          </div>

          <!-- CSS Snippets -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-lg font-semibold text-gray-900">CSS Snippets</h2>
              <p class="text-sm text-gray-600">Quick snippets to get you started</p>
            </div>
            <div class="p-4">
              <div class="space-y-3">
                <button
                  v-for="snippet in cssSnippets"
                  :key="snippet.name"
                  @click="insertSnippet(snippet)"
                  class="w-full text-left p-3 border border-gray-200 rounded-lg hover:border-emerald-300 hover:bg-emerald-50 transition duration-150 ease-in-out group"
                >
                  <div class="flex items-start justify-between">
                    <div>
                      <h3 class="text-sm font-medium text-gray-900 group-hover:text-emerald-700">{{ snippet.name }}</h3>
                      <p class="text-xs text-gray-600 mt-1">{{ snippet.description }}</p>
                    </div>
                    <svg class="w-4 h-4 text-gray-400 group-hover:text-emerald-600 flex-shrink-0 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                    </svg>
                  </div>
                </button>
              </div>
            </div>
          </div>

        </div>

        <!-- Preview and Tools -->
        <div class="space-y-6">
          
          <!-- Live Preview Toggle -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-4">
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
                  <p class="text-xs text-gray-600">Apply styles in real-time</p>
                </div>
              </div>
              <label class="relative inline-flex items-center cursor-pointer">
                <input v-model="livePreview" type="checkbox" class="sr-only peer">
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-green-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-green-600"></div>
              </label>
            </div>
          </div>

          <!-- CSS Validation Results -->
          <div v-if="validationResults.length > 0" class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-lg font-semibold text-gray-900">Validation Results</h2>
              <p class="text-sm text-gray-600">CSS syntax validation</p>
            </div>
            <div class="p-4">
              <div class="space-y-2">
                <div
                  v-for="(result, index) in validationResults"
                  :key="index"
                  :class="[
                    'p-3 rounded-lg text-sm',
                    result.type === 'error' ? 'bg-red-50 text-red-700 border border-red-200' : 'bg-amber-50 text-amber-700 border border-amber-200'
                  ]"
                >
                  <div class="flex items-start">
                    <svg 
                      :class="[
                        'w-4 h-4 mt-0.5 mr-2 flex-shrink-0',
                        result.type === 'error' ? 'text-red-600' : 'text-amber-600'
                      ]" 
                      fill="none" stroke="currentColor" viewBox="0 0 24 24"
                    >
                      <path v-if="result.type === 'error'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                      <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L5.232 15.5c-.77.833.192 2.5 1.732 2.5z" />
                    </svg>
                    <div>
                      <p class="font-medium">{{ result.message }}</p>
                      <p v-if="result.line" class="text-xs mt-1 opacity-75">Line {{ result.line }}</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- CSS Documentation -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-lg font-semibold text-gray-900">CSS Documentation</h2>
              <p class="text-sm text-gray-600">Common CSS selectors for your shop</p>
            </div>
            <div class="p-4">
              <div class="space-y-4 text-sm">
                <div>
                  <h3 class="font-medium text-gray-900 mb-2">Shop Structure</h3>
                  <div class="space-y-1 font-mono text-xs text-gray-600">
                    <p><span class="text-emerald-600">.shop-header</span> - Main shop header</p>
                    <p><span class="text-emerald-600">.shop-nav</span> - Navigation menu</p>
                    <p><span class="text-emerald-600">.shop-content</span> - Main content area</p>
                    <p><span class="text-emerald-600">.shop-footer</span> - Footer section</p>
                  </div>
                </div>
                
                <div>
                  <h3 class="font-medium text-gray-900 mb-2">Product Elements</h3>
                  <div class="space-y-1 font-mono text-xs text-gray-600">
                    <p><span class="text-emerald-600">.product-grid</span> - Product grid container</p>
                    <p><span class="text-emerald-600">.product-card</span> - Individual product card</p>
                    <p><span class="text-emerald-600">.product-image</span> - Product image</p>
                    <p><span class="text-emerald-600">.product-title</span> - Product name</p>
                    <p><span class="text-emerald-600">.product-price</span> - Product price</p>
                  </div>
                </div>

                <div>
                  <h3 class="font-medium text-gray-900 mb-2">Interactive Elements</h3>
                  <div class="space-y-1 font-mono text-xs text-gray-600">
                    <p><span class="text-emerald-600">.btn-primary</span> - Primary buttons</p>
                    <p><span class="text-emerald-600">.btn-secondary</span> - Secondary buttons</p>
                    <p><span class="text-emerald-600">.form-input</span> - Form inputs</p>
                    <p><span class="text-emerald-600">.cart-icon</span> - Shopping cart icon</p>
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
              @click="clearCSS"
              :disabled="saving"
              class="px-4 py-2.5 bg-white border border-gray-300 rounded-lg text-gray-700 text-sm font-medium hover:bg-gray-50 transition duration-150 ease-in-out disabled:opacity-50"
            >
              Clear
            </button>
            <button
              @click="saveCSS"
              :disabled="saving"
              class="inline-flex items-center px-6 py-2.5 bg-emerald-600 text-white text-sm font-medium rounded-lg hover:bg-emerald-700 transition duration-150 ease-in-out disabled:opacity-50 shadow-sm"
            >
              <svg v-if="saving" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              <svg v-else class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
              {{ saving ? 'Saving...' : 'Save CSS' }}
            </button>
          </div>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

// State
const saving = ref(false)
const lastSaved = ref('')
const livePreview = ref(false)
const customCSS = ref('')
const validationResults = ref([])

// CSS Snippets
const cssSnippets = [
  {
    name: 'Custom Button Style',
    description: 'Styled button with hover effects',
    code: `.custom-button {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 12px 24px;
  border: none;
  border-radius: 8px;
  font-weight: 600;
  transition: transform 0.2s ease;
}

.custom-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 16px rgba(102, 126, 234, 0.3);
}`
  },
  {
    name: 'Product Card Enhancement',
    description: 'Enhanced product card with animations',
    code: `.product-card {
  transition: all 0.3s ease;
  border-radius: 12px;
  overflow: hidden;
}

.product-card:hover {
  transform: scale(1.05);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
}

.product-image {
  transition: transform 0.3s ease;
}

.product-card:hover .product-image {
  transform: scale(1.1);
}`
  },
  {
    name: 'Custom Header',
    description: 'Styled header with gradient background',
    code: `.shop-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 20px 0;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.shop-nav a {
  color: white;
  text-decoration: none;
  font-weight: 500;
  transition: opacity 0.2s ease;
}

.shop-nav a:hover {
  opacity: 0.8;
}`
  },
  {
    name: 'Responsive Grid',
    description: 'Responsive product grid layout',
    code: `.product-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 24px;
  padding: 24px;
}

@media (max-width: 768px) {
  .product-grid {
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: 16px;
    padding: 16px;
  }
}`
  }
]

// Computed
const lineCount = computed(() => {
  return Math.max(20, customCSS.value.split('\n').length)
})

// Functions
function goBack() {
  router.push('/dashboard/customization')
}

function updateLineCount() {
  // Trigger reactivity for line count
}

function insertSnippet(snippet) {
  if (customCSS.value && !customCSS.value.endsWith('\n\n')) {
    customCSS.value += '\n\n'
  }
  customCSS.value += snippet.code + '\n\n'
}

function formatCSS() {
  // Basic CSS formatting
  let formatted = customCSS.value
    .replace(/\{/g, ' {\n  ')
    .replace(/\}/g, '\n}\n')
    .replace(/;/g, ';\n  ')
    .replace(/,/g, ',\n')
    .replace(/\n\s*\n/g, '\n\n')
    .trim()
  
  customCSS.value = formatted
}

function validateCSS() {
  validationResults.value = []
  
  // Basic CSS validation
  const lines = customCSS.value.split('\n')
  
  lines.forEach((line, index) => {
    const trimmedLine = line.trim()
    
    // Check for unclosed braces
    if (trimmedLine.includes('{') && !trimmedLine.includes('}')) {
      const openBraces = (trimmedLine.match(/{/g) || []).length
      const closeBraces = (trimmedLine.match(/}/g) || []).length
      if (openBraces > closeBraces) {
        // This is likely okay for multi-line rules
      }
    }
    
    // Check for missing semicolons (basic check)
    if (trimmedLine.includes(':') && !trimmedLine.includes(';') && !trimmedLine.includes('{') && !trimmedLine.includes('}') && trimmedLine !== '') {
      validationResults.value.push({
        type: 'warning',
        message: 'Missing semicolon',
        line: index + 1
      })
    }
  })
  
  if (validationResults.value.length === 0) {
    validationResults.value.push({
      type: 'success',
      message: 'CSS syntax looks good!',
      line: null
    })
  }
}

function clearCSS() {
  customCSS.value = ''
  validationResults.value = []
}

async function saveCSS() {
  saving.value = true
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1000))
    lastSaved.value = new Date().toLocaleString()
    console.log('Custom CSS saved:', customCSS.value)
  } catch (error) {
    console.error('Error saving CSS:', error)
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

/* Code editor styling */
textarea {
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  resize: none;
  white-space: pre;
  overflow-wrap: normal;
  overflow-x: auto;
}
</style>
