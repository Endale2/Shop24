<template>
  <div class="min-h-screen flex flex-col bg-amber-50 font-sans antialiased">
    <!-- Assuming AppNavbar and AppFooter components exist and are properly imported -->
    <AppNavbar />

    <div class="flex-1 flex items-center justify-center p-4 sm:p-6 lg:p-8 animate-fade-in-up">
      <div class="w-full max-w-md md:max-w-xl lg:max-w-2xl bg-white p-8 sm:p-10 shadow-2xl border border-green-200">
        <div class="text-center mb-8">
          <h1 class="text-3xl sm:text-4xl font-extrabold text-green-800">
            Create Your Account
          </h1>
          <p class="mt-2 text-base text-gray-700">
            Already have an account?
            <router-link to="/login" class="font-semibold text-green-600 hover:text-green-700">
              Sign in
            </router-link>
          </p>
        </div>

        <form @submit.prevent="handleRegister" class="space-y-6">
          <!-- Email Address Input -->
          <div>
            <label for="email" class="block text-sm font-medium text-gray-700 mb-1">Email Address</label>
            <input
              id="email"
              type="email"
              v-model="email"
              required
              class="w-full px-4 py-2.5 border border-gray-300 shadow-sm text-gray-800 placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200"
            />
          </div>

          <!-- Password Input with Toggle -->
          <div>
            <label for="password" class="block text-sm font-medium text-gray-700 mb-1">Password</label>
            <div class="relative">
              <input
                id="password"
                :type="showPassword ? 'text' : 'password'"
                v-model="password"
                required
                class="w-full px-4 py-2.5 border border-gray-300 shadow-sm text-gray-800 placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200 pr-10"
              />
              <button
                type="button"
                @click="togglePasswordVisibility"
                class="absolute inset-y-0 right-0 pr-3 flex items-center text-gray-500 hover:text-gray-700"
              >
                <!-- Eye icon for show/hide password -->
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
            <label for="confirm-password" class="block text-sm font-medium text-gray-700 mb-1">Confirm Password</label>
            <div class="relative">
              <input
                id="confirm-password"
                :type="showConfirmPassword ? 'text' : 'password'"
                v-model="confirmPassword"
                required
                class="w-full px-4 py-2.5 border border-gray-300 shadow-sm text-gray-800 placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200 pr-10"
              />
              <button
                type="button"
                @click="toggleConfirmPasswordVisibility"
                class="absolute inset-y-0 right-0 pr-3 flex items-center text-gray-500 hover:text-gray-700"
              >
                <!-- Eye icon for show/hide confirm password -->
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
          <div>
            <button
              type="submit"
              :disabled="loading"
              class="w-full flex justify-center items-center bg-green-700 text-white py-3 px-4 rounded-lg font-semibold text-lg hover:bg-green-800 transition-all duration-200 disabled:opacity-60 disabled:cursor-not-allowed shadow-md"
            >
              <svg v-if="loading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0a8 8 0 00-8 8H4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
              </svg>
              <span>{{ loading ? 'Creating Account...' : 'Create Account' }}</span>
            </button>
          </div>
          <!-- Error Message Display -->
          <p v-if="error" class="text-sm text-red-600 text-center">{{ error }}</p>
        </form>

        <!-- Or continue with section -->
        <div class="mt-8 relative">
          <div class="absolute inset-0 flex items-center">
            <div class="w-full border-t border-gray-300"></div>
          </div>
          <div class="relative flex justify-center text-sm">
            <span class="px-3 bg-white text-gray-500">Or continue with</span>
          </div>
        </div>

        <!-- Google Login Button -->
        <div class="mt-6">
          <button
            @click="loginWithGoogle"
            class="w-full flex justify-center items-center space-x-3 border border-gray-300 rounded-lg py-2.5 px-4 hover:bg-gray-50 transition-all duration-200 shadow-sm"
          >
            <img src="https://www.svgrepo.com/show/355037/google.svg" alt="Google icon" class="h-5 w-5" />
            <span class="text-base font-medium text-gray-700">Continue with Google</span>
          </button>
        </div>
      </div>
    </div>

    <AppFooter />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import AppNavbar from '@/components/AppNavbar.vue' // Ensure this path is correct
import AppFooter from '@/components/AppFooter.vue' // Ensure this path is correct

const router = useRouter()
const auth   = useAuthStore()

const email            = ref('')
const password         = ref('')
const confirmPassword  = ref('') // New ref for confirm password
const loading          = ref(false)
const error            = ref(null)
const showPassword     = ref(false) // New ref for password visibility
const showConfirmPassword = ref(false) // New ref for confirm password visibility

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
 * Initiates the Google login process.
 * Assuming authService.loginWithGoogle() is correctly implemented elsewhere.
 */
function loginWithGoogle() {
  // It seems `authService` is not defined in the provided snippet.
  // Make sure you have `authService` imported or defined for this function to work.
  // For example, if it's part of your auth store: auth.loginWithGoogle()
  // Or, if it's a separate service: import authService from '@/services/authService';
  console.log('Login with Google initiated.');
  // Placeholder for Google login logic
  // authService.loginWithGoogle(); // Uncomment and adjust as per your actual implementation
  error.value = 'Google login is not yet fully implemented in this example.'
}
</script>

<style>
/* Optional: Add custom animations or overrides if needed */
@keyframes fade-in-up {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
.animate-fade-in-up {
  animation: fade-in-up 0.5s ease-out forwards;
}
</style>
