<template>
  <div :class="spinnerClasses" :style="spinnerStyle">
    <div class="loading-spinner"></div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  size: {
    type: String,
    default: 'base',
    validator: (value) => ['sm', 'base', 'lg', 'xl'].includes(value)
  },
  color: {
    type: String,
    default: 'primary'
  },
  centered: {
    type: Boolean,
    default: false
  }
})

const spinnerClasses = computed(() => {
  const classes = []
  
  if (props.centered) {
    classes.push('flex', 'items-center', 'justify-center')
  }
  
  return classes
})

const spinnerStyle = computed(() => {
  const sizeMap = {
    sm: '16px',
    base: '20px',
    lg: '24px',
    xl: '32px'
  }
  
  return {
    '--spinner-size': sizeMap[props.size] || sizeMap.base
  }
})
</script>

<style scoped>
.loading-spinner {
  width: var(--spinner-size);
  height: var(--spinner-size);
}
</style>
