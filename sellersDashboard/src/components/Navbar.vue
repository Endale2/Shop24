<!-- src/components/Navbar.vue -->
<template>
  <header class="bg-white shadow p-4 flex items-center justify-between border-b border-gray-200">
    <div class="flex items-center">
      <button
        class="md:hidden p-1 rounded-md hover:bg-gray-100"
        @click="$emit('toggle-sidebar')"
      >
        <MenuIcon class="h-6 w-6 text-gray-700" />
      </button>
      <div v-if="user" class="ml-4 text-sm text-gray-600 flex items-center space-x-2">
        <UserIcon class="h-5 w-5 text-gray-400" />
        <span>{{ user.email }}</span>
      </div>
    </div>

    <button
      v-if="user"
      @click="handleLogout"
      class="bg-green-600 text-white hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-300 focus:ring-opacity-75 font-semibold py-2 px-4 rounded-full transition transform hover:scale-105"
    >
      Logout
    </button>
  </header>
</template>

<script setup>
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import { MenuIcon, UserIcon } from '@heroicons/vue/outline'

const router = useRouter()
const auth = useAuthStore()

// Define user as a computed so the template can react
const user = computed(() => auth.user)

async function handleLogout() {
  await auth.logout()
  router.push({ name: 'Login' })
}
</script>
