<template>
  <div class="max-w-md mx-auto mt-12 p-6 border rounded shadow">
    <h2 class="text-xl mb-4">{{ isLogin ? 'Login' : 'Register' }}</h2>
    <form @submit.prevent="onSubmit" class="space-y-4">
      <template v-if="!isLogin">
        <input
          v-model="form.firstName"
          placeholder="First Name"
          class="w-full p-2 border rounded"
        />
        <input
          v-model="form.lastName"
          placeholder="Last Name"
          class="w-full p-2 border rounded"
        />
        <input
          v-model="form.username"
          placeholder="Username"
          class="w-full p-2 border rounded"
        />
      </template>
      <input
        v-model="form.email"
        placeholder="Email"
        type="email"
        class="w-full p-2 border rounded"
      />
      <input
        v-model="form.password"
        placeholder="Password"
        type="password"
        class="w-full p-2 border rounded"
      />
      <button
        type="submit"
        class="w-full py-2 bg-blue-500 text-white rounded"
      >
        {{ isLogin ? 'Login' : 'Register' }}
      </button>
    </form>
    <p class="mt-4 text-sm text-center">
      <a @click="isLogin = !isLogin" class="text-blue-500 cursor-pointer">
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
