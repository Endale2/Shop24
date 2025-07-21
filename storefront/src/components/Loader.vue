<template>
  <div class="w-full text-center py-20">
    <div 
      class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-black mx-auto"
      role="status"
    >
      <span class="sr-only">{{ text || 'Loading...' }}</span>
    </div>
    <p v-if="text" class="mt-4 text-gray-600">{{ text }}</p>
  </div>
</template>

<script setup lang="ts">
defineProps({
  text: {
    type: String,
    default: ''
  }
});
</script> 