<!-- src/pages/AuthPage.vue -->

<template>
  <div class="max-w-md mx-auto mt-16 p-6 bg-white rounded shadow">
    <h2 class="text-2xl mb-6 text-center">Welcome to {{ shopName }}</h2>

```
<form @submit.prevent="mode==='login' ? onLogin() : onRegister()" class="space-y-4">
  <template v-if="mode === 'register'">
    <input
      v-model="username"
      placeholder="Username"
      type="text"
      :disabled="loading"
      required
      class="w-full p-2 border rounded"
    />
  </template>

  <input
    v-model="email"
    placeholder="Email"
    type="email"
    :disabled="loading"
    required
    class="w-full p-2 border rounded"
  />
  <input
    v-model="password"
    placeholder="Password"
    type="password"
    :disabled="loading"
    required
    class="w-full p-2 border rounded"
  />

  <button
    type="submit"
    :disabled="loading"
    class="w-full p-2 text-white rounded"
    :class="mode==='login' ? 'bg-blue-600 hover:bg-blue-700' : 'bg-green-600 hover:bg-green-700'"
  >
    {{ loading ? (mode==='login' ? 'Logging in…' : 'Signing up…') : (mode==='login' ? 'Login' : 'Sign Up') }}
  </button>

  <p class="text-sm text-center">
    <template v-if="mode === 'login'">
      New? <a @click.prevent="mode='register'" href="#" class="text-blue-600">Sign up</a>
    </template>
    <template v-else>
      Have an account? <a @click.prevent="mode='login'" href="#" class="text-green-600">Login</a>
    </template>
  </p>
</form>

<div class="mt-6 text-center">
  <button
    @click="onOAuth"
    :disabled="loading"
    class="w-full p-2 border rounded hover:bg-gray-100"
  >
    Continue with Google
  </button>
</div>

<p v-if="error" class="mt-4 text-red-500 text-center">{{ error }}</p>
```

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useShopStore } from '@/stores/shop'
import { useRouter } from 'vue-router'

const auth = useAuthStore()
const shop = useShopStore()
const router = useRouter()

const mode = ref('login')
const email = ref('')
const password = ref('')
const username = ref('')

const loading = computed(() => auth.loading)
const error = computed(() => auth.error)
const shopName = computed(() => shop.current?.name || 'Store')

// Ensure shop data is loaded (for shopName)
onMounted(async () => {
  if (!shop.current) {
    const slug = window.location.host.split('.')[0]
    try {
      await shop.fetchShopAndProducts(slug)
    } catch (e) {
      console.error('Failed to load shop data:', e)
    }
  }
})

// Login or register
const onLogin = async () => {
  try {
    await auth.authenticate({ email: email.value, password: password.value })
    router.push('/')
  } catch {}
}
const onRegister = async () => {
  try {
    await auth.register({ username: username.value, email: email.value, password: password.value })
    router.push('/')
  } catch {}
}

// Google OAuth
const onOAuth = () => {
  auth.oauthLogin()
}
</script>

<style scoped>
/* Tailwind handles styles */
</style>
