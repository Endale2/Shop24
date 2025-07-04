<template>
  <footer class="text-center py-4 border-t mt-8">
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
