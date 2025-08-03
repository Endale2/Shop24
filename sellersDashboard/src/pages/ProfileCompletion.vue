<template>
  <div class="min-h-screen flex flex-col bg-gradient-to-br from-gray-50 to-gray-100 font-sans antialiased">
    <AppNavbar />

    <div class="flex-1 flex items-center justify-center pt-8 px-4 sm:px-6 pb-8">
      <div class="w-full max-w-2xl">
        <!-- Header -->
        <div class="text-center mb-8">
          <div class="flex items-center justify-center mb-4">
            <div class="w-12 h-12 bg-gradient-to-br from-green-500 to-green-600 rounded-xl flex items-center justify-center mr-4 shadow-lg">
              <svg class="w-7 h-7 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
              </svg>
            </div>
            <h1 class="text-3xl sm:text-4xl font-bold text-gray-800 leading-tight">
              Complete Your Profile
            </h1>
          </div>
          <p class="text-base text-gray-600 max-w-md mx-auto">
            Let's get to know you better to personalize your experience and help you succeed.
          </p>
        </div>

        <!-- Progress Steps -->
        <div class="mb-8">
          <div class="flex items-center justify-center space-x-4">
            <div class="flex items-center">
              <div class="w-8 h-8 bg-green-600 text-white rounded-full flex items-center justify-center text-sm font-medium">
                1
              </div>
              <span class="ml-2 text-sm font-medium text-green-600">Account</span>
            </div>
            <div class="w-12 h-0.5 bg-green-300"></div>
            <div class="flex items-center">
              <div class="w-8 h-8 bg-green-600 text-white rounded-full flex items-center justify-center text-sm font-medium">
                2
              </div>
              <span class="ml-2 text-sm font-medium text-green-600">Profile</span>
            </div>
            <div class="w-12 h-0.5 bg-gray-300"></div>
            <div class="flex items-center">
              <div class="w-8 h-8 bg-gray-300 text-gray-600 rounded-full flex items-center justify-center text-sm font-medium">
                3
              </div>
              <span class="ml-2 text-sm font-medium text-gray-500">Shops</span>
            </div>
            <div class="w-12 h-0.5 bg-gray-300"></div>
            <div class="flex items-center">
              <div class="w-8 h-8 bg-gray-300 text-gray-600 rounded-full flex items-center justify-center text-sm font-medium">
                4
              </div>
              <span class="ml-2 text-sm font-medium text-gray-500">Ready</span>
            </div>
          </div>
        </div>

        <!-- Form Card -->
        <div class="bg-white rounded-2xl shadow-lg border border-gray-200 p-8">
          <form @submit.prevent="submit" class="space-y-6">
            <!-- Name Fields -->
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-semibold text-gray-700 mb-2">
                  First Name <span class="text-red-500">*</span>
                </label>
                <input
                  v-model="firstName"
                  type="text"
                  required
                  placeholder="Enter your first name"
                  class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-transparent transition-all duration-200 text-base"
                  :class="{ 'border-red-300 focus:ring-red-400': errors.firstName }"
                />
                <p v-if="errors.firstName" class="mt-1 text-sm text-red-600">{{ errors.firstName }}</p>
              </div>

              <div>
                <label class="block text-sm font-semibold text-gray-700 mb-2">
                  Last Name <span class="text-red-500">*</span>
                </label>
                <input
                  v-model="lastName"
                  type="text"
                  required
                  placeholder="Enter your last name"
                  class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-transparent transition-all duration-200 text-base"
                  :class="{ 'border-red-300 focus:ring-red-400': errors.lastName }"
                />
                <p v-if="errors.lastName" class="mt-1 text-sm text-red-600">{{ errors.lastName }}</p>
              </div>
            </div>

            <!-- Phone Field -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-2">
                Phone Number <span class="text-gray-500 text-xs">(optional)</span>
              </label>
              <div class="relative">
                <input
                  v-model="phone"
                  type="tel"
                  placeholder="Enter your phone number"
                  class="w-full px-4 py-3 pl-12 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-transparent transition-all duration-200 text-base"
                  :class="{ 'border-red-300 focus:ring-red-400': errors.phone }"
                />
                <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                  <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z" />
                  </svg>
                </div>
              </div>
              <p v-if="errors.phone" class="mt-1 text-sm text-red-600">{{ errors.phone }}</p>
            </div>

            <!-- Business Name Field -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-2">
                Business Name <span class="text-gray-500 text-xs">(optional)</span>
              </label>
              <div class="relative">
                <input
                  v-model="businessName"
                  type="text"
                  placeholder="Enter your business name"
                  class="w-full px-4 py-3 pl-12 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-transparent transition-all duration-200 text-base"
                  :class="{ 'border-red-300 focus:ring-red-400': errors.businessName }"
                />
                <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                  <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                  </svg>
                </div>
              </div>
              <p v-if="errors.businessName" class="mt-1 text-sm text-red-600">{{ errors.businessName }}</p>
            </div>

            <!-- Address Field -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-2">
                Business Address <span class="text-gray-500 text-xs">(optional)</span>
              </label>
              <div class="relative">
                <textarea
                  v-model="address"
                  rows="3"
                  placeholder="Enter your business address"
                  class="w-full px-4 py-3 pl-12 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-transparent transition-all duration-200 text-base resize-none"
                  :class="{ 'border-red-300 focus:ring-red-400': errors.address }"
                ></textarea>
                <div class="absolute top-3 left-0 pl-4 flex items-center pointer-events-none">
                  <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                  </svg>
                </div>
              </div>
              <p v-if="errors.address" class="mt-1 text-sm text-red-600">{{ errors.address }}</p>
            </div>

            <!-- Submit Button -->
            <div class="pt-6">
              <button
                type="submit"
                :disabled="loading"
                class="w-full flex justify-center items-center bg-gradient-to-r from-green-600 to-green-700 text-white py-4 px-6 rounded-lg font-semibold text-lg hover:from-green-700 hover:to-green-800 disabled:opacity-60 disabled:cursor-not-allowed shadow-lg hover:shadow-xl transition-all duration-300 transform hover:-translate-y-0.5 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2"
              >
                <svg v-if="loading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
                </svg>
                <span>{{ loading ? 'Saving Profile...' : 'Complete Profile' }}</span>
              </button>
            </div>

            <!-- Error Message -->
            <div v-if="error" class="bg-red-50 border border-red-200 rounded-lg p-4">
              <div class="flex items-center">
                <svg class="w-5 h-5 text-red-500 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
                </svg>
                <p class="text-red-700 font-medium">{{ error }}</p>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>

    <AppFooter />
  </div>
</template>

<script setup>
import { ref, onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import AppNavbar from '@/components/AppNavbar.vue'
import AppFooter from '@/components/AppFooter.vue'

const auth = useAuthStore()
const router = useRouter()

// Guard: if somehow not logged in, kick to Login
if (!auth.isAuthenticated) {
  router.replace({ name: 'Login' })
}

// Form state
const firstName = ref('')
const lastName = ref('')
const phone = ref('')
const businessName = ref('')
const address = ref('')
const loading = ref(false)
const error = ref('')
const errors = ref({})

// Pre-fill from store if already present
onBeforeMount(() => {
  const u = auth.user || {}
  firstName.value = u.first_name || ''
  lastName.value = u.last_name || ''
  phone.value = u.phone || ''
  businessName.value = u.business_name || ''
  address.value = u.address || ''
})

function validateForm() {
  errors.value = {}
  
  if (!firstName.value.trim()) {
    errors.value.firstName = 'First name is required'
  }
  
  if (!lastName.value.trim()) {
    errors.value.lastName = 'Last name is required'
  }
  
  if (phone.value && !/^[\+]?[1-9][\d]{0,15}$/.test(phone.value.replace(/\s/g, ''))) {
    errors.value.phone = 'Please enter a valid phone number'
  }
  
  return Object.keys(errors.value).length === 0
}

async function submit() {
  if (!validateForm()) {
    return
  }
  
  loading.value = true
  error.value = ''

  try {
    // Send exactly the camelCase fields your service expects
    await auth.updateProfile({
      firstName: firstName.value.trim(),
      lastName: lastName.value.trim(),
      phone: phone.value.trim(),
      businessName: businessName.value.trim(),
      address: address.value.trim(),
    })

    // Refresh the store so auth.user is fully up to date
    await auth.fetchMe()

    // Push to shops
    router.replace({ name: 'ShopSelection' })
  } catch (err) {
    console.error(err)
    error.value = err.response?.data?.error
                || err.message
                || 'Failed to save profile'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
/* No additional styles needed - using Tailwind classes */
</style>