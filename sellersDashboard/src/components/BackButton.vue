<template>
  <button
    @click="handleBack"
    class="inline-flex items-center text-gray-600 hover:text-green-600 transition-all duration-200 ease-in-out group"
    :class="[
      size === 'sm' ? 'px-2 py-1' : 'px-3 py-2',
      variant === 'rounded' ? 'rounded-full bg-gray-100/50 hover:bg-green-50' : '',
      variant === 'minimal' ? '' : 'hover:bg-gray-50 rounded-lg'
    ]"
  >
    <!-- iOS-style chevron icon -->
    <svg 
      class="transition-all duration-200 group-hover:scale-110 group-hover:-translate-x-0.5"
      :class="[
        size === 'sm' ? 'w-4 h-4' : 'w-5 h-5',
        showText ? 'mr-2' : ''
      ]"
      fill="none" 
      stroke="currentColor" 
      stroke-width="2.5" 
      viewBox="0 0 24 24"
      stroke-linecap="round" 
      stroke-linejoin="round"
    >
      <path d="M15 18l-6-6 6-6" />
    </svg>
    
    <!-- Optional text -->
    <span 
      v-if="showText" 
      class="text-sm font-medium transition-colors duration-200"
      :class="size === 'sm' ? 'text-xs' : 'text-sm'"
    >
      {{ text }}
    </span>
  </button>
</template>

<script setup>
import { useRouter } from 'vue-router'

const props = defineProps({
  // Navigation behavior
  to: {
    type: [String, Object],
    default: null
  },
  // Display options
  text: {
    type: String,
    default: 'Back'
  },
  showText: {
    type: Boolean,
    default: true
  },
  // Styling variants
  variant: {
    type: String,
    default: 'default', // 'default', 'minimal', 'rounded'
    validator: (value) => ['default', 'minimal', 'rounded'].includes(value)
  },
  size: {
    type: String,
    default: 'md', // 'sm', 'md'
    validator: (value) => ['sm', 'md'].includes(value)
  }
})

const router = useRouter()

function handleBack() {
  if (props.to) {
    // Navigate to specific route
    if (typeof props.to === 'string') {
      router.push(props.to)
    } else {
      router.push(props.to)
    }
  } else {
    // Use browser back
    router.back()
  }
}
</script>

<style scoped>
/* iOS-style smooth animations */
.group:hover svg {
  transform: translateX(-2px) scale(1.1);
}

/* Subtle shadow for rounded variant */
.group.rounded-full {
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06);
}

.group.rounded-full:hover {
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
}
</style>
