<template>
  <span :class="badgeClasses">
    <component v-if="icon" :is="icon" class="w-3 h-3 mr-1" />
    <slot>{{ label }}</slot>
  </span>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  variant: {
    type: String,
    default: 'neutral',
    validator: (value) => ['success', 'warning', 'error', 'info', 'neutral'].includes(value)
  },
  size: {
    type: String,
    default: 'base',
    validator: (value) => ['sm', 'base', 'lg'].includes(value)
  },
  label: {
    type: String,
    default: ''
  },
  icon: {
    type: [Object, Function],
    default: null
  },
  dot: {
    type: Boolean,
    default: false
  }
})

const badgeClasses = computed(() => {
  const classes = ['badge', `badge-${props.variant}`]
  
  if (props.size !== 'base') {
    classes.push(`badge-${props.size}`)
  }
  
  if (props.dot) {
    classes.push('badge-dot')
  }
  
  return classes
})
</script>

<style scoped>
.badge-sm {
  padding: var(--spacing-1) var(--spacing-2);
  font-size: 0.625rem;
}

.badge-lg {
  padding: var(--spacing-2) var(--spacing-4);
  font-size: var(--font-size-sm);
}

.badge-dot {
  padding-left: var(--spacing-4);
  position: relative;
}

.badge-dot::before {
  content: '';
  position: absolute;
  left: var(--spacing-2);
  top: 50%;
  transform: translateY(-50%);
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background-color: currentColor;
}
</style>
