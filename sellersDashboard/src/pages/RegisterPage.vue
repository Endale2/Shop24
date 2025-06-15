<template>
  <div class="min-h-screen flex flex-col bg-amber-100 font-sans antialiased">
    <AppNavbar />

    <div class="flex-1 flex items-center justify-center p-4 sm:p-6 lg:p-8 animate-fade-in-up">
      <div class="w-full max-w-md bg-white p-8 sm:p-10 rounded-2xl shadow-2xl border border-green-200">
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
          <div>
            <label for="email" class="block text-sm font-medium text-gray-700 mb-1">Email Address</label>
            <input
              id="email"
              type="email"
              v-model="email"
              required
              class="w-full px-4 py-2.5 border border-gray-300 rounded-lg shadow-sm text-gray-800 placeholder-gray-500 focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200"
            />
          </div>

          <div>
            <label for="password" class="block text-sm font-medium text-gray-700 mb-1">Password</label>
            <input
              id="password"
              type="password"
              v-model="password"
              required
              class="w-full px-4 py-2.5 border border-gray-300 rounded-lg shadow-sm text-gray-800 placeholder-gray-500 focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200"
            />
          </div>

          <div>
            <button
              type="submit"
              :disabled="loading"
              class="w-full flex justify-center items-center bg-green-700 text-white py-3 px-4 rounded-lg font-semibold text-lg hover:bg-green-800 transition-all duration-200 disabled:opacity-60 disabled:cursor-not-allowed shadow-md"
            >
              <svg v-if="loading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0…"/>
              </svg>
              <span>{{ loading ? 'Creating Account...' : 'Create Account' }}</span>
            </button>
          </div>
          <p v-if="error" class="text-sm text-red-600 text-center">{{ error }}</p>
        </form>

        <div class="mt-8 relative">
          <div class="absolute inset-0 flex items-center">
            <div class="w-full border-t border-gray-300"></div>
          </div>
          <div class="relative flex justify-center text-sm">
            <span class="px-3 bg-white text-gray-500">Or continue with</span>
          </div>
        </div>

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
import AppNavbar from '@/components/AppNavbar.vue'
import AppFooter from '@/components/AppFooter.vue'

const router = useRouter()
const auth    = useAuthStore()

const email    = ref('')
const password = ref('')
const loading  = ref(false)
const error    = ref(null)

async function handleRegister() {
  loading.value = true
  error.value   = null

  try {
    // 1) create the bare auth record
    await auth.register({ email: email.value, password: password.value })
    // 2) immediately take them to complete-profile
    router.push({ name: 'CompleteProfile' })
  } catch (err) {
    error.value = err.response?.data?.error || err.message || 'Registration failed'
  } finally {
    loading.value = false
  }
}

function loginWithGoogle() {
  authService.loginWithGoogle()
}
</script>
