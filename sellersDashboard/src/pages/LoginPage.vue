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
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'

const router = useRouter()
const auth = useAuthStore()

// Form state
const email = ref('')
const password = ref('')
const loading = ref(false)
const error = ref(null)

// Prevent access if already logged in
onMounted(() => {
  if (auth.isAuthenticated) {
    router.replace({ name: 'ShopSelection' })
  }
})

async function handleLogin() {
  loading.value = true
  error.value = null

  try {
    // Attempt login; auth.login should set auth.user internally
    await auth.login({ email: email.value, password: password.value })

    // Redirect to shop selection after success
    router.push({ name: 'ShopSelection' })
  } catch (err) {
    error.value = 'Login failed. Please check your credentials.'
  } finally {
    loading.value = false
  }
}
</script>
