<template>
  <div class="max-w-md mx-auto mt-20 p-8 bg-white border border-gray-200 rounded-none">
    <h2 class="text-2xl font-bold mb-6 text-center text-gray-900 uppercase">{{ isLogin ? 'Login' : 'Register' }}</h2>
    <form @submit.prevent="onSubmit" class="space-y-5">
      <template v-if="!isLogin">
        <input
          v-model="form.firstName"
          placeholder="First Name"
          class="w-full p-3 border rounded-lg focus:ring-2 focus:ring-black focus:outline-none text-xs uppercase"
        />
        <input
          v-model="form.lastName"
          placeholder="Last Name"
          class="w-full p-3 border rounded-lg focus:ring-2 focus:ring-black focus:outline-none text-xs uppercase"
        />
        <input
          v-model="form.username"
          placeholder="Username"
          class="w-full p-3 border rounded-lg focus:ring-2 focus:ring-black focus:outline-none text-xs uppercase"
        />
      </template>
      <input
        v-model="form.email"
        placeholder="Email"
        type="email"
        class="w-full p-3 border rounded-lg focus:ring-2 focus:ring-black focus:outline-none text-xs uppercase"
      />
      <input
        v-model="form.password"
        placeholder="Password"
        type="password"
        class="w-full p-3 border rounded-lg focus:ring-2 focus:ring-black focus:outline-none text-xs uppercase"
      />
      <button
        type="submit"
        class="w-full py-3 bg-black text-white rounded text-lg font-bold transition disabled:opacity-50 uppercase"
      >
        {{ isLogin ? 'Login' : 'Register' }}
      </button>
    </form>
    <p class="mt-6 text-sm text-center">
      <a @click="isLogin = !isLogin" class="text-blue-600 cursor-pointer hover:underline uppercase">
        {{ isLogin ? 'Need an account?' : 'Have an account?' }}
      </a>
    </p>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { register, login } from '@/services/auth'

const isLogin = ref(true)
const form = ref<any>({})
const router = useRouter()

async function onSubmit() {
  const payload = { ...form.value, shopId: '' } // you can fetch shopId similarly if needed
  if (isLogin.value) {
    await login(payload)
  } else {
    await register(payload)
  }
  router.push('/')
}
</script>
