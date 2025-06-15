<template>
  <div class="min-h-screen flex flex-col bg-amber-100 font-sans antialiased">
    <AppNavbar />

    <div class="flex-1 flex items-center justify-center p-4 sm:p-6 lg:p-8 animate-fade-in-up">
      <div class="w-full max-w-2xl bg-white p-8 sm:p-10 rounded-2xl shadow-2xl border border-green-200">
        <div class="text-center mb-8">
          <h1 class="mt-4 text-3xl sm:text-4xl font-extrabold text-green-800">
            Sign in to your account
          </h1>
          <p class="mt-2 text-base text-gray-700">
            Welcome back, seller! Let's get you back to business.
          </p>
        </div>

        <form @submit.prevent="handleLogin" class="grid grid-cols-1 md:grid-cols-2 gap-x-8 gap-y-6">
          <div>
            <label for="email" class="block text-sm font-medium text-gray-700 mb-1">Email address</label>
            <div>
              <input
                v-model="email"
                id="email"
                type="email"
                autocomplete="email"
                required
                class="w-full px-4 py-2.5 border border-gray-300 rounded-lg shadow-sm text-gray-800 placeholder-gray-500 focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200"
              />
            </div>
          </div>

          <div>
            <label for="password" class="block text-sm font-medium text-gray-700 mb-1">Password</label>
            <div>
              <input
                v-model="password"
                id="password"
                type="password"
                autocomplete="current-password"
                required
                class="w-full px-4 py-2.5 border border-gray-300 rounded-lg shadow-sm text-gray-800 placeholder-gray-500 focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200"
              />
            </div>
          </div>

          <div class="md:col-span-2">
            <button
              type="submit"
              :disabled="loading"
              class="w-full flex justify-center items-center bg-green-700 text-white py-3 px-4 rounded-lg font-semibold text-lg hover:bg-green-800 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 transition-all duration-200 disabled:opacity-60 disabled:cursor-not-allowed shadow-md"
            >
              <svg v-if="loading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
              </svg>
              <span>{{ loading ? 'Signing In...' : 'Sign In' }}</span>
            </button>
          </div>
          <p v-if="error" class="text-sm text-red-600 text-center md:col-span-2 mt-4">{{ error }}</p>
        </form>

        <div class="mt-8 relative md:col-span-2">
          <div class="absolute inset-0 flex items-center">
            <div class="w-full border-t border-gray-300"></div>
          </div>
          <div class="relative flex justify-center text-sm">
            <span class="px-3 bg-white text-gray-500">Or continue with</span>
          </div>
        </div>

        <div class="mt-6 md:col-span-2">
          <button
            @click="loginWithGoogle"
            class="w-full flex justify-center items-center space-x-3 border border-gray-300 rounded-lg py-2.5 px-4 hover:bg-gray-50 transition-all duration-200 shadow-sm"
          >
            <img src="https://www.svgrepo.com/show/355037/google.svg" alt="Google icon" class="h-5 w-5" />
            <span class="text-base font-medium text-gray-700">Continue with Google</span>
          </button>
        </div>

        <p class="mt-8 text-center text-gray-600 text-base md:col-span-2">
          Don't have an account?
          <router-link to="/register" class="font-semibold text-green-600 hover:text-green-700">
            Sign up now
          </router-link>
        </p>
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

const email = ref('')
const password = ref('')
const loading = ref(false)
const error = ref(null)

function goToShopSelection() {
  router.replace({ name: 'ShopSelection' })
}

// Kick off OAuth login by hitting your backend endpoint
function loginWithGoogle() {
  // Ensure this URL matches your actual backend's OAuth initiation endpoint
  const redirect = encodeURIComponent(
    window.location.origin + '/login?oauth=google'
  )
  window.location.href =
    `http://localhost:8080/auth/seller/oauth/google?redirect_uri=${redirect}`
}

// If already logged in, skip straight to shop picker
onMounted(() => {
  if (auth.isAuthenticated) {
    goToShopSelection()
  }
})

// Watch for `?oauth=google` after Google callback
watch(
  () => route.query.oauth,
  async (val) => {
    if (val === 'google') {
      loading.value = true
      try {
        await auth.fetchMe()
        goToShopSelection()
      } catch {
        error.value = 'Google login failed. Please try again.'
      } finally {
        loading.value = false
      }
    }
  },
  { immediate: true }
)

async function handleLogin() {
  loading.value = true
  error.value = null
  try {
    await auth.login({ email: email.value, password: password.value })
    goToShopSelection()
  } catch (e) {
    console.error('Login error:', e); // Log the actual error for debugging
    error.value = 'Invalid email or password. Please check your credentials.';
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
@keyframes fade-in-up {
  from { opacity: 0; transform: translateY(20px); }
  to   { opacity: 1; transform: translateY(0); }
}
.animate-fade-in-up {
  animation: fade-in-up 0.6s ease-out forwards;
}
</style>