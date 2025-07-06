<template>
  <div class="min-h-screen flex flex-col bg-gradient-to-br from-amber-50 to-green-50 font-sans antialiased">
    <!-- Assuming AppNavbar and AppFooter components exist and are properly imported -->
    <AppNavbar />

    <div class="flex-1 flex items-center justify-center p-4 sm:p-6 lg:p-8 animate-fade-in-up">
      <div class="w-full max-w-md md:max-w-xl lg:max-w-2xl bg-white rounded-3xl shadow-2xl border border-green-200 p-8 sm:p-10 lg:p-12">
        <!-- Header -->
        <div class="text-center mb-10">
          <div class="w-20 h-20 bg-gradient-to-br from-green-100 to-amber-100 rounded-full flex items-center justify-center mx-auto mb-6 shadow-lg">
            <svg class="w-10 h-10 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z" />
            </svg>
          </div>
          <h1 class="text-3xl sm:text-4xl font-extrabold text-gray-800 mb-3">
            Create Your Account
          </h1>
          <p class="text-lg text-gray-600">
            Already have an account?
            <router-link to="/login" class="font-semibold text-green-600 hover:text-green-700 transition-colors duration-200">
              Sign in
            </router-link>
          </p>
        </div>

        <form @submit.prevent="handleRegister" class="space-y-6">
          <!-- Email Address Input -->
          <div>
            <label for="email" class="block text-sm font-semibold text-gray-700 mb-2">
              Email Address <span class="text-red-500">*</span>
            </label>
            <div class="relative">
              <input
                id="email"
                type="email"
                v-model="email"
                required
                placeholder="Enter your email address"
                class="w-full px-4 py-3 pl-12 border border-gray-300 rounded-xl focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200 bg-gray-50 hover:bg-white focus:bg-white"
              />
              <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207" />
                </svg>
              </div>
            </div>
          </div>

          <!-- Password Input with Toggle -->
          <div>
            <label for="password" class="block text-sm font-semibold text-gray-700 mb-2">
              Password <span class="text-red-500">*</span>
            </label>
            <div class="relative">
              <input
                id="password"
                :type="showPassword ? 'text' : 'password'"
                v-model="password"
                required
                placeholder="Create a strong password"
                class="w-full px-4 py-3 pl-12 pr-12 border border-gray-300 rounded-xl focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200 bg-gray-50 hover:bg-white focus:bg-white"
              />
              <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                </svg>
              </div>
              <button
                type="button"
                @click="togglePasswordVisibility"
                class="absolute inset-y-0 right-0 pr-4 flex items-center text-gray-500 hover:text-gray-700 transition-colors duration-200"
              >
                <svg v-if="showPassword" xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.52 18.52 0 0 1 2.16-3.15M2 2l20 20M15 15l-3-3m-3-3L2 2" />
                  <path d="M9.9 4.24A9.91 9.91 0 0 1 12 4c7 0 11 8 11 8a18.54 18.54 0 0 1-1.67 2.68" />
                </svg>
                <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z" />
                  <circle cx="12" cy="12" r="3" />
                </svg>
              </button>
            </div>
          </div>

          <!-- Confirm Password Input with Toggle -->
          <div>
            <label for="confirm-password" class="block text-sm font-semibold text-gray-700 mb-2">
              Confirm Password <span class="text-red-500">*</span>
            </label>
            <div class="relative">
              <input
                id="confirm-password"
                :type="showConfirmPassword ? 'text' : 'password'"
                v-model="confirmPassword"
                required
                placeholder="Confirm your password"
                class="w-full px-4 py-3 pl-12 pr-12 border border-gray-300 rounded-xl focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200 bg-gray-50 hover:bg-white focus:bg-white"
              />
              <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
              <button
                type="button"
                @click="toggleConfirmPasswordVisibility"
                class="absolute inset-y-0 right-0 pr-4 flex items-center text-gray-500 hover:text-gray-700 transition-colors duration-200"
              >
                <svg v-if="showConfirmPassword" xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.52 18.52 0 0 1 2.16-3.15M2 2l20 20M15 15l-3-3m-3-3L2 2" />
                  <path d="M9.9 4.24A9.91 9.91 0 0 1 12 4c7 0 11 8 11 8a18.54 18.54 0 0 1-1.67 2.68" />
                </svg>
                <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z" />
                  <circle cx="12" cy="12" r="3" />
                </svg>
              </button>
            </div>
          </div>

          <!-- Submit Button -->
          <div class="pt-4">
            <button
              type="submit"
              :disabled="loading"
              class="w-full flex justify-center items-center bg-gradient-to-r from-green-600 to-green-700 text-white py-4 px-6 rounded-xl font-semibold text-lg hover:from-green-700 hover:to-green-800 disabled:opacity-60 disabled:cursor-not-allowed shadow-lg hover:shadow-xl transition-all duration-300 transform hover:-translate-y-1 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2"
            >
              <svg v-if="loading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0a8 8 0 00-8 8H4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
              </svg>
              <span>{{ loading ? 'Creating Account...' : 'Create Account' }}</span>
            </button>
          </div>

          <!-- Error Message Display -->
          <div v-if="error" class="bg-red-50 border border-red-200 rounded-xl p-4">
            <div class="flex items-center">
              <svg class="w-5 h-5 text-red-500 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
              </svg>
              <p class="text-red-700 font-medium">{{ error }}</p>
            </div>
          </div>
        </form>

        <!-- Divider -->
        <div class="mt-8 relative">
          <div class="absolute inset-0 flex items-center">
            <div class="w-full border-t border-gray-300"></div>
          </div>
          <div class="relative flex justify-center text-sm">
            <span class="px-4 bg-white text-gray-500 font-medium">Or continue with</span>
          </div>
        </div>

        <!-- Google Login Button -->
        <div class="mt-6">
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

        <!-- Terms and Privacy -->
        <div class="mt-6 text-center">
          <p class="text-xs text-gray-500">
            By creating an account, you agree to our 
            <a href="#" class="text-green-600 hover:text-green-700 font-medium">Terms of Service</a> 
            and 
            <a href="#" class="text-green-600 hover:text-green-700 font-medium">Privacy Policy</a>
          </p>
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

const email            = ref('')
const password         = ref('')
const confirmPassword  = ref('')
const loading          = ref(false)
const error            = ref(null)
const showPassword     = ref(false)
const showConfirmPassword = ref(false)

/**
 * Toggles the visibility of the main password input field.
 */
function togglePasswordVisibility() {
  showPassword.value = !showPassword.value
}

/**
 * Toggles the visibility of the confirm password input field.
 */
function toggleConfirmPasswordVisibility() {
  showConfirmPassword.value = !showConfirmPassword.value
}

/**
 * Handles the user registration process.
 * Validates passwords and calls the authentication store.
 */
async function handleRegister() {
  loading.value = true
  error.value   = null

  // Password validation
  if (password.value !== confirmPassword.value) {
    error.value = 'Passwords do not match!'
    loading.value = false
    return
  }

  try {
    // 1) create the bare auth record
    await auth.register({ email: email.value, password: password.value })
    // 2) immediately take them to complete-profile
    router.push({ name: 'CompleteProfile' })
  } catch (err) {
    // Handle specific error messages or provide a generic one
    error.value = err.response?.data?.error || err.message || 'Registration failed. Please try again.'
  } finally {
    loading.value = false
  }
}

/**
 * Initiates Google OAuth login using the auth store.
 */
function loginWithGoogle() {
  try {
    loading.value = true
    error.value = null
    auth.loginWithGoogle()
  } catch (err) {
    console.error('Google login error:', err)
    error.value = 'Failed to initiate Google login. Please try again.'
    loading.value = false
  }
}

// If already logged in, skip straight to shop picker
onMounted(() => {
  if (auth.isAuthenticated) {
    router.replace({ name: 'ShopSelection' })
  }
})

// Watch for OAuth callback parameters
watch(
  () => route.query,
  async (query) => {
    // Handle Google OAuth callback
    if (query.oauth === 'google' || query.code) {
      loading.value = true
      error.value = null
      
      try {
        // Fetch user data after successful OAuth
        await auth.fetchMe()
        
        // Check if user has completed profile
        const user = auth.user
        if (user && (user.first_name || user.last_name)) {
          // User has completed profile, go to shop selection
          router.replace({ name: 'ShopSelection' })
        } else {
          // New user from Google, go to profile completion
          router.replace({ name: 'CompleteProfile' })
        }
      } catch (err) {
        console.error('OAuth callback error:', err)
        error.value = 'Google login failed. Please try again.'
      } finally {
        loading.value = false
      }
    }
    
    // Handle OAuth errors
    if (query.error) {
      error.value = `Login failed: ${query.error}`
      loading.value = false
    }
  },
  { immediate: true, deep: true }
)
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
