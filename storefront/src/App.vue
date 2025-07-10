<template>
  <div>
    <div v-if="authStore.error" class="bg-red-600 text-white text-center py-2 w-full z-50">
      {{ authStore.error }}
    </div>
    <router-view />
  </div>
</template>

<script setup lang="ts">
import { useAuthStore } from './stores/auth'
const authStore = useAuthStore()
</script>

<style>
body {
  margin: 0;
  font-family: system-ui, sans-serif;
}
</style>
