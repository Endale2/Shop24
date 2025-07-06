<template>
  <div class="min-h-screen flex flex-col bg-gradient-to-br from-amber-50 to-green-50 font-sans antialiased">
    <AppNavbar />

    <div class="flex-1 flex items-center justify-center p-4 sm:p-6 lg:p-8 animate-fade-in-up">
      <div class="bg-white rounded-3xl shadow-2xl border border-green-200 w-full max-w-2xl p-8 sm:p-10 lg:p-12">
        <!-- Header -->
        <div class="text-center mb-8">
          <div class="w-20 h-20 bg-gradient-to-br from-green-100 to-amber-100 rounded-full flex items-center justify-center mx-auto mb-6 shadow-lg">
            <svg class="w-10 h-10 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
            </svg>
          </div>
          <h2 class="text-3xl sm:text-4xl font-extrabold text-gray-800 mb-3">Complete Your Profile</h2>
          <p class="text-lg text-gray-600 max-w-md mx-auto">
            Let's get to know you better to personalize your experience
          </p>
        </div>

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
                class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200 bg-gray-50 hover:bg-white focus:bg-white"
              />
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
                class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200 bg-gray-50 hover:bg-white focus:bg-white"
              />
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
                class="w-full px-4 py-3 pl-12 border border-gray-300 rounded-xl focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200 bg-gray-50 hover:bg-white focus:bg-white"
              />
              <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z" />
                </svg>
              </div>
            </div>
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
                class="w-full px-4 py-3 pl-12 border border-gray-300 rounded-xl focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200 bg-gray-50 hover:bg-white focus:bg-white"
              />
              <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                </svg>
              </div>
            </div>
          </div>

          <!-- Submit Button -->
          <div class="pt-6">
            <button
              type="submit"
              :disabled="loading"
              class="w-full flex justify-center items-center bg-gradient-to-r from-green-600 to-green-700 text-white py-4 px-6 rounded-xl font-semibold text-lg hover:from-green-700 hover:to-green-800 disabled:opacity-60 disabled:cursor-not-allowed shadow-lg hover:shadow-xl transition-all duration-300 transform hover:-translate-y-1 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2"
            >
              <svg v-if="loading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
              </svg>
              <span>{{ loading ? 'Saving Profile...' : 'Complete Profile' }}</span>
            </button>
          </div>

          <!-- Error Message -->
          <div v-if="error" class="bg-red-50 border border-red-200 rounded-xl p-4">
            <div class="flex items-center">
              <svg class="w-5 h-5 text-red-500 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
              </svg>
              <p class="text-red-700 font-medium">{{ error }}</p>
            </div>
          </div>
        </form>

        <!-- Progress Indicator -->
        <div class="mt-8 pt-6 border-t border-gray-100">
          <div class="flex items-center justify-center space-x-2">
            <div class="w-2 h-2 bg-green-500 rounded-full"></div>
            <div class="w-2 h-2 bg-green-500 rounded-full"></div>
            <div class="w-2 h-2 bg-gray-300 rounded-full"></div>
            <div class="w-2 h-2 bg-gray-300 rounded-full"></div>
          </div>
          <p class="text-center text-sm text-gray-500 mt-2">Step 2 of 4: Profile Setup</p>
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

const auth   = useAuthStore()
const router = useRouter()

// 1) guard: if somehow not logged in, kick to Login
if (!auth.isAuthenticated) {
  router.replace({ name: 'Login' })
}

// form state
const firstName    = ref('')
const lastName     = ref('')
const phone        = ref('')
const businessName = ref('')
const loading      = ref(false)
const error        = ref('')

// 2) pre-fill from store if already present
onBeforeMount(() => {
  const u = auth.user || {}
  firstName.value    = u.firstName    || ''
  lastName.value     = u.lastName     || ''
  phone.value        = u.phone        || ''
  businessName.value = u.businessName || ''
})

async function submit() {
  loading.value = true
  error.value   = ''

  try {
    // 3) send exactly the camelCase fields your service expects
    await auth.updateProfile({
      firstName:    firstName.value.trim(),
      lastName:     lastName.value.trim(),
      phone:        phone.value.trim(),
      businessName: businessName.value.trim(),
    })

    // 4) refresh the store so auth.user is fully up to date
    await auth.fetchMe()

    // 5) push to shops
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
@keyframes fade-in-up {
  from { 
    opacity: 0; 
    transform: translateY(30px); 
  }
  to { 
    opacity: 1; 
    transform: translateY(0); 
  }
}

.animate-fade-in-up {
  animation: fade-in-up 0.6s ease-out forwards;
}
</style>