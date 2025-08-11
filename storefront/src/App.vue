<template>
  <div>
    <div v-if="authStore.sessionLoading && route.meta.requiresAuth">
      <Loader text="Checking session..." />
    </div>
    <div v-else>
      <div v-if="authStore.error" class="bg-red-500 text-white text-center py-2 w-full z-50">
        {{ authStore.error }}
      </div>
      <router-view />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useAuthStore } from './stores/auth'
import Loader from '@/components/Loader.vue'
import { useRoute } from 'vue-router'
const authStore = useAuthStore()
const route = useRoute()
</script>

<style>
body {
  margin: 0;
  font-family: system-ui, sans-serif;
}
</style>
