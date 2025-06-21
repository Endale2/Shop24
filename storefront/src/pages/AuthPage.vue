<!-- src/pages/AuthPage.vue -->
<template>
  <div class="max-w-md mx-auto mt-16 p-6 bg-white rounded shadow">
    <h2 class="text-2xl mb-4 text-center">Welcome to {{ shopName }}</h2>

    <div v-if="mode==='login'">
      <form @submit.prevent="onLogin" class="space-y-4">
        <input v-model="email" placeholder="Email" type="email" required class="w-full p-2 border rounded" />
        <input v-model="password" placeholder="Password" type="password" required class="w-full p-2 border rounded" />
        <button :disabled="loading" class="w-full p-2 bg-blue-600 text-white rounded">
          {{ loading ? 'Logging in...' : 'Login' }}
        </button>
        <p class="text-sm text-center">
          New? <a @click.prevent="mode='register'" href="#" class="text-blue-600">Sign up</a>
        </p>
      </form>
    </div>

    <div v-else>
      <form @submit.prevent="onRegister" class="space-y-4">
        <input v-model="username" placeholder="Username" required class="w-full p-2 border rounded" />
        <input v-model="email" placeholder="Email" type="email" required class="w-full p-2 border rounded" />
        <input v-model="password" placeholder="Password" type="password" required class="w-full p-2 border rounded" />
        <button :disabled="loading" class="w-full p-2 bg-green-600 text-white rounded">
          {{ loading ? 'Signing up...' : 'Sign Up' }}
        </button>
        <p class="text-sm text-center">
          Have an account? <a @click.prevent="mode='login'" href="#" class="text-green-600">Login</a>
        </p>
      </form>
    </div>

    <div class="mt-6 text-center">
      <button @click="onOAuth" class="w-full p-2 border rounded">
        Continue with Google
      </button>
    </div>

    <p v-if="error" class="mt-4 text-red-500 text-center">{{ error }}</p>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useShopStore } from '@/stores/shop'

const auth = useAuthStore()
const shop = useShopStore()

const mode = ref('login')
const email = ref('')
const password = ref('')
const username = ref('')

const loading = computed(() => auth.loading)
const error = computed(() => auth.error)
const shopName = computed(() => shop.current?.name || 'Store')

async function onLogin() {
  try {
    await auth.authenticate({ email: email.value, password: password.value })
    window.location.href = window.location.origin
  } catch {}
}

async function onRegister() {
  try {
    await auth.register({ username: username.value, email: email.value, password: password.value })
    window.location.href = window.location.origin
  } catch {}
}

function onOAuth() {
  auth.oauthLogin()
}
</script>

<style scoped>
/* Add Tailwind or custom styles here */
</style>
