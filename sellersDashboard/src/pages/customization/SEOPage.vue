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
          <div class="w-10 h-10 bg-gradient-to-br from-orange-500 to-orange-600 rounded-lg flex items-center justify-center mr-3 shadow-sm">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </div>
          <div>
            <h1 class="text-2xl sm:text-3xl font-bold text-gray-900 tracking-tight">
              SEO & Meta Tags
            </h1>
            <p class="text-sm text-gray-600 mt-1">
              Optimize your shop for search engines and social media
            </p>
          </div>
        </div>
      </div>

      <div class="space-y-6">
        
        <!-- Basic SEO Settings -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
          <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
            <h2 class="text-lg font-semibold text-gray-900">Basic SEO Settings</h2>
            <p class="text-sm text-gray-600">Configure your shop's basic search engine optimization</p>
          </div>
          <div class="p-6 space-y-6">
            
            <!-- Shop Title -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">Shop Title *</label>
              <input
                v-model="seoSettings.title"
                type="text"
                maxlength="60"
                class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-orange-500 focus:border-orange-500"
                placeholder="Your Amazing Shop - Quality Products Online"
              />
              <div class="flex justify-between items-center mt-1">
                <p class="text-xs text-gray-500">Appears in browser tab and search results</p>
                <p :class="['text-xs', seoSettings.title.length > 50 ? 'text-red-600' : 'text-gray-500']">
                  {{ seoSettings.title.length }}/60 characters
                </p>
              </div>
            </div>

            <!-- Meta Description -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">Meta Description *</label>
              <textarea
                v-model="seoSettings.description"
                rows="3"
                maxlength="160"
                class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-orange-500 focus:border-orange-500"
                placeholder="Discover our curated collection of high-quality products. Fast shipping, excellent customer service, and competitive prices guaranteed."
              ></textarea>
              <div class="flex justify-between items-center mt-1">
                <p class="text-xs text-gray-500">Brief description shown in search results</p>
                <p :class="['text-xs', seoSettings.description.length > 150 ? 'text-red-600' : 'text-gray-500']">
                  {{ seoSettings.description.length }}/160 characters
                </p>
              </div>
            </div>

            <!-- Keywords -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">Focus Keywords</label>
              <input
                v-model="seoSettings.keywords"
                type="text"
                class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-orange-500 focus:border-orange-500"
                placeholder="online shop, quality products, fast shipping"
              />
              <p class="text-xs text-gray-500 mt-1">Separate keywords with commas. Focus on 3-5 main keywords.</p>
            </div>

          </div>
        </div>

        <!-- Social Media Meta Tags -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
          <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
            <h2 class="text-lg font-semibold text-gray-900">Social Media Preview</h2>
            <p class="text-sm text-gray-600">Control how your shop appears when shared on social media</p>
          </div>
          <div class="p-6">
            <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
              
              <!-- Social Meta Settings -->
              <div class="space-y-6">
                
                <!-- Open Graph Title -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Social Media Title</label>
                  <input
                    v-model="seoSettings.ogTitle"
                    type="text"
                    maxlength="95"
                    class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-orange-500 focus:border-orange-500"
                    placeholder="Shop Amazing Products Online | Your Shop Name"
                  />
                  <p :class="['text-xs mt-1', seoSettings.ogTitle.length > 85 ? 'text-red-600' : 'text-gray-500']">
                    {{ seoSettings.ogTitle.length }}/95 characters
                  </p>
                </div>

                <!-- Open Graph Description -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Social Media Description</label>
                  <textarea
                    v-model="seoSettings.ogDescription"
                    rows="3"
                    maxlength="200"
                    class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-orange-500 focus:border-orange-500"
                    placeholder="Discover our amazing collection of products with fast shipping and great customer service."
                  ></textarea>
                  <p :class="['text-xs mt-1', seoSettings.ogDescription.length > 180 ? 'text-red-600' : 'text-gray-500']">
                    {{ seoSettings.ogDescription.length }}/200 characters
                  </p>
                </div>

                <!-- Open Graph Image -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Social Media Image URL</label>
                  <input
                    v-model="seoSettings.ogImage"
                    type="url"
                    class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-orange-500 focus:border-orange-500"
                    placeholder="https://example.com/social-preview.jpg"
                  />
                  <p class="text-xs text-gray-500 mt-1">Recommended: 1200x630px image</p>
                </div>

              </div>

              <!-- Social Media Preview -->
              <div>
                <h3 class="text-sm font-medium text-gray-700 mb-4">Preview</h3>
                
                <!-- Facebook/LinkedIn Style Preview -->
                <div class="border border-gray-200 rounded-lg overflow-hidden mb-4">
                  <div class="aspect-w-16 aspect-h-9 bg-gray-100">
                    <img
                      v-if="seoSettings.ogImage"
                      :src="seoSettings.ogImage"
                      alt="Social preview"
                      class="w-full h-32 object-cover"
                      @error="$event.target.style.display='none'"
                    />
                    <div v-else class="flex items-center justify-center h-32 bg-gray-100">
                      <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                      </svg>
                    </div>
                  </div>
                  <div class="p-3 bg-white">
                    <h4 class="font-medium text-gray-900 text-sm line-clamp-2">
                      {{ seoSettings.ogTitle || 'Your Shop Title' }}
                    </h4>
                    <p class="text-xs text-gray-600 mt-1 line-clamp-2">
                      {{ seoSettings.ogDescription || 'Your shop description will appear here...' }}
                    </p>
                    <p class="text-xs text-gray-500 mt-1 uppercase">yourshop.com</p>
                  </div>
                </div>

                <!-- Twitter Card Style Preview -->
                <div class="border border-gray-200 rounded-lg p-3">
                  <div class="flex space-x-3">
                    <div class="w-16 h-16 bg-gray-100 rounded-lg flex-shrink-0 flex items-center justify-center">
                      <img
                        v-if="seoSettings.ogImage"
                        :src="seoSettings.ogImage"
                        alt="Twitter preview"
                        class="w-full h-full object-cover rounded-lg"
                        @error="$event.target.style.display='none'"
                      />
                      <svg v-else class="w-6 h-6 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                      </svg>
                    </div>
                    <div class="flex-1 min-w-0">
                      <h4 class="font-medium text-gray-900 text-sm line-clamp-1">
                        {{ seoSettings.ogTitle || 'Your Shop Title' }}
                      </h4>
                      <p class="text-xs text-gray-600 mt-1 line-clamp-2">
                        {{ seoSettings.ogDescription || 'Your shop description...' }}
                      </p>
                      <p class="text-xs text-gray-500 mt-1">yourshop.com</p>
                    </div>
                  </div>
                </div>

              </div>
            </div>
          </div>
        </div>

        <!-- Technical SEO Settings -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
          <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
            <h2 class="text-lg font-semibold text-gray-900">Technical SEO</h2>
            <p class="text-sm text-gray-600">Advanced SEO settings and configurations</p>
          </div>
          <div class="p-6">
            <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
              
              <!-- Search Engine Settings -->
              <div class="space-y-6">
                
                <!-- Robots Meta -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Search Engine Indexing</label>
                  <select v-model="seoSettings.robots" class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-orange-500 focus:border-orange-500">
                    <option value="index,follow">Index and Follow (Recommended)</option>
                    <option value="index,nofollow">Index but Don't Follow Links</option>
                    <option value="noindex,follow">Don't Index but Follow Links</option>
                    <option value="noindex,nofollow">Don't Index or Follow</option>
                  </select>
                  <p class="text-xs text-gray-500 mt-1">Controls how search engines crawl your shop</p>
                </div>

                <!-- Canonical URL -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Canonical URL</label>
                  <input
                    v-model="seoSettings.canonicalUrl"
                    type="url"
                    class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-orange-500 focus:border-orange-500"
                    placeholder="https://yourshop.com"
                  />
                  <p class="text-xs text-gray-500 mt-1">Preferred URL for this page</p>
                </div>

                <!-- Language -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Shop Language</label>
                  <select v-model="seoSettings.language" class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-orange-500 focus:border-orange-500">
                    <option value="en">English</option>
                    <option value="es">Spanish</option>
                    <option value="fr">French</option>
                    <option value="de">German</option>
                    <option value="it">Italian</option>
                    <option value="pt">Portuguese</option>
                  </select>
                </div>

              </div>

              <!-- Schema Markup -->
              <div class="space-y-6">
                
                <h3 class="text-sm font-medium text-gray-700">Schema Markup</h3>
                
                <!-- Business Type -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Business Type</label>
                  <select v-model="seoSettings.businessType" class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-orange-500 focus:border-orange-500">
                    <option value="OnlineStore">Online Store</option>
                    <option value="LocalBusiness">Local Business</option>
                    <option value="Corporation">Corporation</option>
                    <option value="Organization">Organization</option>
                  </select>
                </div>

                <!-- Business Description -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Business Description</label>
                  <textarea
                    v-model="seoSettings.businessDescription"
                    rows="3"
                    class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-orange-500 focus:border-orange-500"
                    placeholder="Describe your business for search engines..."
                  ></textarea>
                </div>

                <!-- Contact Information -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Contact Phone</label>
                  <input
                    v-model="seoSettings.phone"
                    type="tel"
                    class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-orange-500 focus:border-orange-500"
                    placeholder="+1 (555) 123-4567"
                  />
                </div>

              </div>
            </div>
          </div>
        </div>

        <!-- SEO Analysis -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
          <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
            <h2 class="text-lg font-semibold text-gray-900">SEO Analysis</h2>
            <p class="text-sm text-gray-600">Check your SEO optimization score</p>
          </div>
          <div class="p-6">
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
              
              <!-- SEO Score -->
              <div class="text-center">
                <div class="relative w-24 h-24 mx-auto mb-4">
                  <svg class="w-24 h-24 transform -rotate-90" viewBox="0 0 36 36">
                    <path
                      d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831"
                      fill="none"
                      stroke="#E5E7EB"
                      stroke-width="2"
                    />
                    <path
                      d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831"
                      fill="none"
                      stroke="#10B981"
                      stroke-width="2"
                      stroke-dasharray="75, 100"
                    />
                  </svg>
                  <div class="absolute inset-0 flex items-center justify-center">
                    <span class="text-2xl font-bold text-gray-900">75</span>
                  </div>
                </div>
                <h3 class="text-lg font-semibold text-gray-900">SEO Score</h3>
                <p class="text-sm text-gray-600">Good optimization</p>
              </div>

              <!-- Recommendations -->
              <div class="md:col-span-2">
                <h3 class="text-lg font-semibold text-gray-900 mb-4">Recommendations</h3>
                <div class="space-y-3">
                  <div class="flex items-start space-x-3">
                    <div class="w-2 h-2 bg-green-500 rounded-full mt-2 flex-shrink-0"></div>
                    <div>
                      <p class="text-sm font-medium text-gray-900">Title tag length is optimal</p>
                      <p class="text-xs text-gray-600">Your title is within the recommended 50-60 character range</p>
                    </div>
                  </div>
                  <div class="flex items-start space-x-3">
                    <div class="w-2 h-2 bg-yellow-500 rounded-full mt-2 flex-shrink-0"></div>
                    <div>
                      <p class="text-sm font-medium text-gray-900">Add more focus keywords</p>
                      <p class="text-xs text-gray-600">Include 2-3 more relevant keywords in your content</p>
                    </div>
                  </div>
                  <div class="flex items-start space-x-3">
                    <div class="w-2 h-2 bg-red-500 rounded-full mt-2 flex-shrink-0"></div>
                    <div>
                      <p class="text-sm font-medium text-gray-900">Missing social media image</p>
                      <p class="text-xs text-gray-600">Add an Open Graph image for better social sharing</p>
                    </div>
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
              @click="resetSEOSettings"
              :disabled="saving"
              class="px-4 py-2.5 bg-white border border-gray-300 rounded-lg text-gray-700 text-sm font-medium hover:bg-gray-50 transition duration-150 ease-in-out disabled:opacity-50"
            >
              Reset
            </button>
            <button
              @click="saveSEOSettings"
              :disabled="saving"
              class="inline-flex items-center px-6 py-2.5 bg-orange-600 text-white text-sm font-medium rounded-lg hover:bg-orange-700 transition duration-150 ease-in-out disabled:opacity-50 shadow-sm"
            >
              <svg v-if="saving" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              <svg v-else class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
              {{ saving ? 'Saving...' : 'Save SEO Settings' }}
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

const router = useRouter()

// State
const saving = ref(false)
const lastSaved = ref('')

// SEO Settings
const seoSettings = reactive({
  title: '',
  description: '',
  keywords: '',
  ogTitle: '',
  ogDescription: '',
  ogImage: '',
  robots: 'index,follow',
  canonicalUrl: '',
  language: 'en',
  businessType: 'OnlineStore',
  businessDescription: '',
  phone: ''
})

// Functions
function goBack() {
  router.push('/dashboard/customization')
}

function resetSEOSettings() {
  seoSettings.title = ''
  seoSettings.description = ''
  seoSettings.keywords = ''
  seoSettings.ogTitle = ''
  seoSettings.ogDescription = ''
  seoSettings.ogImage = ''
  seoSettings.robots = 'index,follow'
  seoSettings.canonicalUrl = ''
  seoSettings.language = 'en'
  seoSettings.businessType = 'OnlineStore'
  seoSettings.businessDescription = ''
  seoSettings.phone = ''
}

async function saveSEOSettings() {
  saving.value = true
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1000))
    lastSaved.value = new Date().toLocaleString()
    console.log('SEO settings saved:', seoSettings)
  } catch (error) {
    console.error('Error saving SEO settings:', error)
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

/* Line clamp utilities */
.line-clamp-1 {
  overflow: hidden;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 1;
}

.line-clamp-2 {
  overflow: hidden;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 2;
}
</style>
