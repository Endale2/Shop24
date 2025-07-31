<template>
  <div class="min-h-screen flex flex-col bg-gradient-to-br from-gray-50 to-gray-100 font-sans antialiased">
    <AppNavbar />

    <div class="flex-1 flex flex-col items-center pt-8 px-4 sm:px-6 pb-8">
      <div class="w-full max-w-4xl">
        <!-- Header -->
        <div class="text-center mb-8">
          <div class="flex items-center justify-center mb-4">
            <div class="w-12 h-12 bg-gradient-to-br from-green-500 to-green-600 rounded-xl flex items-center justify-center mr-4 shadow-lg">
              <svg class="w-7 h-7 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
              </svg>
            </div>
            <h1 class="text-3xl sm:text-4xl font-bold text-gray-800 leading-tight">
              Create New Shop
            </h1>
          </div>
          <p class="text-base text-gray-600 max-w-2xl mx-auto">
            Set up your new online store and start selling to customers worldwide. Fill in the details below to get started.
          </p>
        </div>

        <!-- Progress Steps -->
        <div class="mb-8">
          <div class="flex items-center justify-center space-x-4">
            <div class="flex items-center">
              <div class="w-8 h-8 bg-green-600 text-white rounded-full flex items-center justify-center text-sm font-medium">
                1
              </div>
              <span class="ml-2 text-sm font-medium text-green-600">Basic Info</span>
            </div>
            <div class="w-12 h-0.5 bg-gray-300"></div>
            <div class="flex items-center">
              <div class="w-8 h-8 bg-gray-300 text-gray-600 rounded-full flex items-center justify-center text-sm font-medium">
                2
              </div>
              <span class="ml-2 text-sm font-medium text-gray-500">Customization</span>
            </div>
            <div class="w-12 h-0.5 bg-gray-300"></div>
            <div class="flex items-center">
              <div class="w-8 h-8 bg-gray-300 text-gray-600 rounded-full flex items-center justify-center text-sm font-medium">
                3
              </div>
              <span class="ml-2 text-sm font-medium text-gray-500">Review</span>
            </div>
          </div>
        </div>

        <!-- Form Card -->
        <div class="bg-white rounded-2xl shadow-lg border border-gray-200 p-8">
          <form @submit.prevent="handleSubmit" class="space-y-6">
            <!-- Shop Name -->
            <div>
              <label for="shopName" class="block text-sm font-semibold text-gray-700 mb-2">
                Shop Name *
              </label>
              <input
                id="shopName"
                v-model="form.name"
                type="text"
                required
                placeholder="Enter your shop name"
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-transparent transition-all duration-200 text-base"
                :class="{ 'border-red-300 focus:ring-red-400': errors.name }"
              />
              <p v-if="errors.name" class="mt-1 text-sm text-red-600">{{ errors.name }}</p>
              <p class="mt-1 text-xs text-gray-500">This will be displayed to your customers</p>
            </div>

            <!-- Shop Description -->
            <div>
              <label for="shopDescription" class="block text-sm font-semibold text-gray-700 mb-2">
                Shop Description
              </label>
              <textarea
                id="shopDescription"
                v-model="form.description"
                rows="4"
                placeholder="Describe what your shop sells and what makes it unique..."
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-transparent transition-all duration-200 text-base resize-none"
                :class="{ 'border-red-300 focus:ring-red-400': errors.description }"
              ></textarea>
              <p v-if="errors.description" class="mt-1 text-sm text-red-600">{{ errors.description }}</p>
              <p class="mt-1 text-xs text-gray-500">Help customers understand what you offer</p>
            </div>

            <!-- Shop Category -->
            <div>
              <label for="shopCategory" class="block text-sm font-semibold text-gray-700 mb-2">
                Shop Category *
              </label>
              <select
                id="shopCategory"
                v-model="form.category"
                required
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-transparent transition-all duration-200 text-base bg-white"
                :class="{ 'border-red-300 focus:ring-red-400': errors.category }"
              >
                <option value="">Select a category</option>
                <option value="fashion">Fashion & Apparel</option>
                <option value="electronics">Electronics</option>
                <option value="home">Home & Garden</option>
                <option value="beauty">Beauty & Personal Care</option>
                <option value="sports">Sports & Outdoors</option>
                <option value="books">Books & Media</option>
                <option value="toys">Toys & Games</option>
                <option value="food">Food & Beverages</option>
                <option value="health">Health & Wellness</option>
                <option value="automotive">Automotive</option>
                <option value="other">Other</option>
              </select>
              <p v-if="errors.category" class="mt-1 text-sm text-red-600">{{ errors.category }}</p>
            </div>

            <!-- Shop Logo -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-2">
                Shop Logo
              </label>
              <div class="flex items-center space-x-4">
                <div class="w-20 h-20 bg-gray-100 rounded-lg flex items-center justify-center border-2 border-dashed border-gray-300 hover:border-green-400 transition-colors duration-200 cursor-pointer">
                  <div v-if="!form.logo" class="text-center">
                    <svg class="w-8 h-8 text-gray-400 mx-auto mb-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                    </svg>
                    <p class="text-xs text-gray-500">Add Logo</p>
                  </div>
                  <img v-else :src="form.logo" alt="Shop logo" class="w-full h-full object-cover rounded-lg" />
                </div>
                <div class="flex-1">
                  <input
                    type="file"
                    accept="image/*"
                    @change="handleLogoUpload"
                    class="hidden"
                    ref="logoInput"
                  />
                  <button
                    type="button"
                    @click="$refs.logoInput.click()"
                    class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors duration-200 text-sm font-medium"
                  >
                    {{ form.logo ? 'Change Logo' : 'Upload Logo' }}
                  </button>
                  <p class="mt-1 text-xs text-gray-500">Recommended: 200x200px, PNG or JPG</p>
                </div>
              </div>
            </div>

            <!-- Contact Information -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label for="shopEmail" class="block text-sm font-semibold text-gray-700 mb-2">
                  Contact Email *
                </label>
                <input
                  id="shopEmail"
                  v-model="form.email"
                  type="email"
                  required
                  placeholder="shop@example.com"
                  class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-transparent transition-all duration-200 text-base"
                  :class="{ 'border-red-300 focus:ring-red-400': errors.email }"
                />
                <p v-if="errors.email" class="mt-1 text-sm text-red-600">{{ errors.email }}</p>
              </div>

              <div>
                <label for="shopPhone" class="block text-sm font-semibold text-gray-700 mb-2">
                  Contact Phone
                </label>
                <input
                  id="shopPhone"
                  v-model="form.phone"
                  type="tel"
                  placeholder="+1 (555) 123-4567"
                  class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-transparent transition-all duration-200 text-base"
                />
              </div>
            </div>

            <!-- Address -->
            <div>
              <label for="shopAddress" class="block text-sm font-semibold text-gray-700 mb-2">
                Business Address
              </label>
              <textarea
                id="shopAddress"
                v-model="form.address"
                rows="3"
                placeholder="Enter your business address..."
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-transparent transition-all duration-200 text-base resize-none"
              ></textarea>
            </div>

            <!-- Currency -->
            <div>
              <label for="shopCurrency" class="block text-sm font-semibold text-gray-700 mb-2">
                Default Currency *
              </label>
              <select
                id="shopCurrency"
                v-model="form.currency"
                required
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-transparent transition-all duration-200 text-base bg-white"
              >
                <option value="USD">USD - US Dollar</option>
                <option value="EUR">EUR - Euro</option>
                <option value="GBP">GBP - British Pound</option>
                <option value="CAD">CAD - Canadian Dollar</option>
                <option value="AUD">AUD - Australian Dollar</option>
                <option value="JPY">JPY - Japanese Yen</option>
              </select>
            </div>

            <!-- Terms and Conditions -->
            <div class="flex items-start space-x-3">
              <input
                id="terms"
                v-model="form.acceptTerms"
                type="checkbox"
                required
                class="mt-1 w-4 h-4 text-green-600 border-gray-300 rounded focus:ring-green-400"
              />
              <label for="terms" class="text-sm text-gray-700">
                I agree to the 
                <a href="#" class="text-green-600 hover:text-green-700 font-medium">Terms of Service</a> 
                and 
                <a href="#" class="text-green-600 hover:text-green-700 font-medium">Privacy Policy</a>
              </label>
            </div>

            <!-- Action Buttons -->
            <div class="flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0 sm:space-x-4 pt-6 border-t border-gray-200">
              <router-link
                to="/shop-selection"
                class="w-full sm:w-auto px-6 py-3 border border-gray-300 text-gray-700 rounded-lg font-medium hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-offset-2 transition-all duration-200 text-center"
              >
                Cancel
              </router-link>
              
              <button
                type="submit"
                :disabled="loading"
                class="w-full sm:w-auto px-8 py-3 bg-gradient-to-r from-green-600 to-green-700 text-white rounded-lg font-medium hover:from-green-700 hover:to-green-800 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2 transition-all duration-200 shadow-sm hover:shadow-md disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center"
              >
                <svg v-if="loading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                {{ loading ? 'Creating Shop...' : 'Create Shop' }}
              </button>
            </div>
          </form>
        </div>

        <!-- Tips Section -->
        <div class="mt-8 bg-blue-50 rounded-xl p-6 border border-blue-200">
          <div class="flex items-start">
            <div class="w-6 h-6 bg-blue-500 rounded-full flex items-center justify-center mr-3 mt-0.5">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <div>
              <h3 class="text-lg font-semibold text-blue-800 mb-2">Tips for a Great Shop</h3>
              <ul class="text-sm text-blue-700 space-y-1">
                <li>• Choose a memorable and descriptive shop name</li>
                <li>• Write a clear description that explains what you sell</li>
                <li>• Select the most relevant category for better discoverability</li>
                <li>• Upload a high-quality logo that represents your brand</li>
                <li>• Provide accurate contact information for customer support</li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>

    <AppFooter />
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, watchEffect } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { useAuthStore } from '@/store/auth'

import AppNavbar from '@/components/AppNavbar.vue'
import AppFooter from '@/components/AppFooter.vue'

const router = useRouter()
const shopStore = useShopStore()
const auth = useAuthStore()

watchEffect(() => {
  if (!auth.isAuthenticated) {
    router.replace({ name: 'Landing' })
  }
})

const loading = ref(false)
const errors = reactive({})

const form = reactive({
  name: '',
  description: '',
  category: '',
  logo: '',
  email: '',
  phone: '',
  address: '',
  currency: 'USD',
  acceptTerms: false
})

onMounted(() => {
  // Pre-fill email with user's email if available
  if (auth.user?.email) {
    form.email = auth.user.email
  }
})

function handleLogoUpload(event) {
  const file = event.target.files[0]
  if (file) {
    const reader = new FileReader()
    reader.onload = (e) => {
      form.logo = e.target.result
    }
    reader.readAsDataURL(file)
  }
}

function validateForm() {
  errors.name = ''
  errors.description = ''
  errors.category = ''
  errors.email = ''

  if (!form.name.trim()) {
    errors.name = 'Shop name is required'
  } else if (form.name.length < 3) {
    errors.name = 'Shop name must be at least 3 characters'
  }

  if (form.description && form.description.length > 500) {
    errors.description = 'Description must be less than 500 characters'
  }

  if (!form.category) {
    errors.category = 'Please select a category'
  }

  if (!form.email) {
    errors.email = 'Email is required'
  } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(form.email)) {
    errors.email = 'Please enter a valid email address'
  }

  return !Object.values(errors).some(error => error)
}

async function handleSubmit() {
  if (!validateForm()) {
    return
  }

  loading.value = true
  try {
    const shopData = {
      name: form.name.trim(),
      description: form.description.trim(),
      category: form.category,
      logo: form.logo,
      email: form.email.trim(),
      phone: form.phone.trim(),
      address: form.address.trim(),
      currency: form.currency
    }

    const newShop = await shopStore.createShop(shopData)
    shopStore.setActiveShop(newShop)
    
    // Show success message and redirect
    alert('Shop created successfully!')
    router.push({ name: 'Dashboard' })
  } catch (error) {
    console.error('Error creating shop:', error)
    alert('Failed to create shop. Please try again.')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
/* Custom file input styling */
input[type="file"] {
  cursor: pointer;
}

/* Custom checkbox styling */
input[type="checkbox"]:checked {
  background-color: #059669;
  border-color: #059669;
}
</style> 