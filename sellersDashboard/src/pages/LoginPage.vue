<template>
  <div class="min-h-screen bg-gray-50 flex flex-col justify-center items-center px-6 py-8">
    <div class="w-full max-w-md">
      <div class="text-center mb-6">
        <svg class="mx-auto h-12 w-auto text-gray-800" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M14.25 21.75L24 31.5L33.75 21.75M24 3.75V31.5" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
          <path d="M39 31.5H9C6.23858 31.5 4 33.7386 4 36.5V40.25C4 42.0449 5.45507 43.5 7.25 43.5H40.75C42.5449 43.5 44 42.0449 44 40.25V36.5C44 33.7386 41.7614 31.5 39 31.5Z" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
        </svg>
        <h1 class="mt-4 text-3xl font-extrabold text-gray-900">
          Sign in to your account
        </h1>
        <p class="mt-2 text-sm text-gray-600">
          Welcome back, seller!
        </p>
      </div>

      <div class="bg-white p-8 rounded-2xl shadow-lg">
        <form @submit.prevent="handleLogin" class="space-y-6">
          <div>
            <label for="email" class="block text-sm font-medium text-gray-700">Email address</label>
            <div class="mt-1">
              <input
                v-model="email"
                id="email"
                type="email"
                autocomplete="email"
                required
                class="w-full px-4 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-green-500 focus:border-green-500 transition"
              />
            </div>
          </div>

          <div>
            <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
            <div class="mt-1">
              <input
                v-model="password"
                id="password"
                type="password"
                autocomplete="current-password"
                required
                class="w-full px-4 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-green-500 focus:border-green-500 transition"
              />
            </div>
          </div>

          <div>
            <button
              type="submit"
              :disabled="loading"
              class="w-full flex justify-center items-center bg-green-600 text-white py-2.5 px-4 rounded-lg font-semibold hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 transition disabled:opacity-60"
            >
              <svg v-if="loading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
              </svg>
              <span>{{ loading ? 'Signing In...' : 'Sign In' }}</span>
            </button>
          </div>
          <p v-if="error" class="text-sm text-red-600 text-center">{{ error }}</p>
        </form>

        <div class="mt-6 relative">
          <div class="absolute inset-0 flex items-center">
            <div class="w-full border-t border-gray-300"></div>
          </div>
          <div class="relative flex justify-center text-sm">
            <span class="px-2 bg-white text-gray-500">Or continue with</span>
          </div>
        </div>

        <div class="mt-6">
          <button
            @click="loginWithGoogle"
            class="w-full flex justify-center items-center space-x-3 border border-gray-300 rounded-lg py-2.5 px-4 hover:bg-gray-50 transition"
          >
            <img src="https://www.svgrepo.com/show/355037/google.svg" alt="Google" class="h-5 w-5" />
            <span class="text-sm font-medium text-gray-700">Continue with Google</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'
import { useRouter, useRoute }   from 'vue-router'
import { useAuthStore }          from '@/store/auth'

const router = useRouter()
const route  = useRoute()
const auth   = useAuthStore()

const email    = ref('')
const password = ref('')
const loading  = ref(false)
const error    = ref(null)

function goToShopSelection() {
  router.replace({ name: 'ShopSelection' })
}

// Kick off OAuth login by hitting your backend endpoint
function loginWithGoogle() {
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
  error.value   = null
  try {
    await auth.login({ email: email.value, password: password.value })
    goToShopSelection()
  } catch {
    error.value = 'Invalid email or password.'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
/* No specific scoped styles are needed for this centering approach,
   as Tailwind utility classes handle it effectively. */
</style>