<!-- src/pages/LoginPage.vue -->
<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white p-6 rounded shadow w-full max-w-sm">
      <h2 class="text-xl font-semibold mb-4 text-center">Login</h2>

      <form @submit.prevent="handleLogin" class="space-y-3">
        <div>
          <label class="block text-sm mb-1">Email</label>
          <input
            v-model="email"
            type="email"
            required
            class="w-full px-2 py-1 border rounded text-sm"
          />
        </div>

        <div>
          <label class="block text-sm mb-1">Password</label>
          <input
            v-model="password"
            type="password"
            required
            class="w-full px-2 py-1 border rounded text-sm"
          />
        </div>

        <button
          type="submit"
          :disabled="loading"
          class="w-full bg-blue-500 text-white py-1 rounded text-sm hover:bg-blue-600 disabled:opacity-50"
        >
          {{ loading ? 'Logging in...' : 'Login' }}
        </button>

        <p v-if="error" class="text-red-500 text-xs text-center mt-2">
          {{ error }}
        </p>
      </form>
    </div>
    <div class="mt-4 text-center">
    <button
  @click="loginWithGoogle"
  class="inline-flex items-center px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700 transition"
>
  <!-- google icon… -->
  Log in with Google
</button>


  </div>
  </div>
   
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/store/auth'

const router = useRouter()
const route  = useRoute()
const auth   = useAuthStore()

const email = ref('')
const password = ref('')
const loading = ref(false)
const error   = ref(null)

function loginWithGoogle() {
  // include redirect param so the backend sends us back here:
  window.location.href =
    'http://localhost:8080/auth/seller/oauth/google?redirect_uri=' +
    encodeURIComponent(window.location.origin + '/login?oauth=google')
}

// on mount, if we have an oauth query, finish login:
onMounted(async () => {
  // 1) if already logged in, go straight in
  if (auth.isAuthenticated) {
    return router.replace({ name: 'ShopSelection' })
  }

  // 2) if we have `?oauth=google`, attempt to fetchMe
  if (route.query.oauth === 'google') {
    try {
      await auth.fetchMe()
      return router.replace({ name: 'ShopSelection' })
    } catch {
      // fall through to show login form
    }
  }
})

async function handleLogin() {
  loading.value = true
  error.value = null
  try {
    await auth.login({ email: email.value, password: password.value })
    router.push({ name: 'ShopSelection' })
  } catch {
    error.value = 'Login failed. Please check your credentials.'
  } finally {
    loading.value = false
  }
}
</script>

