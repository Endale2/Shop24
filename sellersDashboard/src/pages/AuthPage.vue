<template>
  <div class="min-h-screen flex flex-col bg-gradient-to-br from-amber-50 to-green-50 font-sans antialiased">
    <AppNavbar />
    <div class="flex-1 flex items-center justify-center p-4 sm:p-6 lg:p-8 animate-fade-in-up">
      <div class="w-full max-w-md md:max-w-xl lg:max-w-2xl bg-white rounded-3xl shadow-2xl border border-green-200 p-8 sm:p-10 lg:p-12">
        <!-- Header -->
        <div class="text-center mb-10">
          <div class="w-20 h-20 bg-gradient-to-br from-green-100 to-amber-100 rounded-full flex items-center justify-center mx-auto mb-6 shadow-lg">
            <svg class="w-10 h-10 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1" />
            </svg>
          </div>
          <h1 class="text-3xl sm:text-4xl font-extrabold text-gray-800 mb-3">
            Welcome to DRPS
          </h1>
          <p class="text-lg text-gray-600">
            Sign in to manage your stores and grow your business
          </p>
        </div>

        <!-- OAuth Buttons -->
        <div class="mt-6 space-y-4">
          <!-- Google Login Button -->
          <button 
            @click="loginWithGoogle" 
            :disabled="loading" 
            class="w-full flex justify-center items-center space-x-3 border-2 border-gray-300 rounded-xl py-3 px-4 hover:bg-gray-50 hover:border-gray-400 transition-all duration-200 shadow-sm disabled:opacity-60 disabled:cursor-not-allowed"
          >
            <svg v-if="loading" class="animate-spin h-5 w-5 text-gray-500" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
            </svg>
            <img v-else src="https://www.svgrepo.com/show/355037/google.svg" alt="Google icon" class="h-5 w-5" />
            <span class="text-base font-medium text-gray-700">
              {{ loading ? 'Connecting to Google...' : 'Continue with Google' }}
            </span>
          </button>


        </div>

        <!-- Error Message -->
        <div v-if="error" class="bg-red-50 border border-red-200 rounded-xl p-4 mt-6">
          <div class="flex items-center">
            <svg class="w-5 h-5 text-red-500 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
            </svg>
            <p class="text-red-700 font-medium">{{ error }}</p>
          </div>
        </div>

        <!-- Loading State -->
        <div v-if="loading" class="mt-6 text-center">
          <div class="flex items-center justify-center space-x-2">
            <svg class="animate-spin h-5 w-5 text-green-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
            </svg>
            <span class="text-gray-600">Processing authentication...</span>
          </div>
        </div>
      </div>
    </div>
    <AppFooter />
  </div>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import AppNavbar from '@/components/AppNavbar.vue'
import AppFooter from '@/components/AppFooter.vue'

const router = useRouter()
const route = useRoute()
const auth = useAuthStore()
const loading = ref(false)
const error = ref(null)

// Helper function to determine where to redirect after authentication
function determineRedirect() {
  const user = auth.user
  
  // If user doesn't have first_name or last_name, they need to complete profile
  if (!user || !user.first_name || !user.last_name) {
    return { name: 'CompleteProfile' }
  }
  
  // If user has complete profile, redirect to shop selection
  return { name: 'ShopSelection' }
}

// Google OAuth login
function loginWithGoogle() {
  try {
    loading.value = true
    error.value = null
    auth.loginWithGoogle()
  } catch (err) {
    error.value = 'Failed to initiate Google login. Please try again.'
    loading.value = false
  }
}



// Handle OAuth callback from Google
async function handleOAuthCallback() {
  loading.value = true
  error.value = null
  
  try {
    await auth.fetchMe()
    const redirect = determineRedirect()
    router.replace(redirect)
  } catch (err) {
    error.value = 'Authentication failed. Please try again.'
    console.error('OAuth callback error:', err)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  // If user is already authenticated, redirect appropriately
  if (auth.isAuthenticated) {
    const redirect = determineRedirect()
    router.replace(redirect)
  }
})

// Watch for OAuth callback parameters
watch(() => route.query, async (query) => {
  if (query.oauth === 'google' || query.code) {
    await handleOAuthCallback()
  }
  
  if (query.error) {
    error.value = `Authentication failed: ${query.error}`
    loading.value = false
  }
}, { immediate: true, deep: true })
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