<template>
  <footer 
    class="text-center py-8 border-t text-sm font-light mt-16 transition-colors"
    :style="footerStyle"
  >
    <div class="mx-auto px-4" :class="containerClasses">
      <p :style="textStyle">
        © {{ new Date().getFullYear() }} {{ shop?.name }}. All rights reserved.
        <span v-if="theme?.version" class="opacity-60 ml-2">
          • Theme v{{ theme.version }}
        </span>
      </p>
    </div>
  </footer>
</template>

<script setup lang="ts">
import { defineProps, computed } from 'vue'

interface Props {
  shop: { name: string } | null
  theme?: any | null
}
const props = defineProps<Props>()

// =============== Dynamic Styles ===============

const footerStyle = computed(() => {
  if (!props.theme) return { backgroundColor: '#ffffff', borderTopColor: '#E5E7EB', color: '#9CA3AF' }
  
  return {
    backgroundColor: props.theme.colors?.background || '#ffffff',
    borderTopColor: props.theme.colors?.border || '#E5E7EB'
  }
})

const containerClasses = computed(() => {
  const classes = []
  
  if (props.theme?.layout?.containerWidth === 'full') {
    classes.push('max-w-none')
  } else if (props.theme?.layout?.containerWidth === 'wide') {
    classes.push('max-w-7xl')
  } else {
    classes.push('max-w-6xl')
  }
  
  return classes
})

const textStyle = computed(() => {
  if (!props.theme) return { color: '#9CA3AF' }
  
  return {
    color: props.theme.colors?.bodyText || '#9CA3AF',
    fontFamily: props.theme.fonts?.body ? `'${props.theme.fonts.body}', sans-serif` : undefined
  }
})
</script>
