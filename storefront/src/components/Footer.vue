<template>
  <footer class="text-center py-8 border-t border-gray-200 bg-white text-gray-400 text-sm font-light mt-16">
    © {{ new Date().getFullYear() }} {{ shop?.name }}. All rights reserved.
  </footer>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'

interface Props {
  shop: { name: string } | null
}
defineProps<Props>()
</script>
