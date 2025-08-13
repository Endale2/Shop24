<template>
  <component
    :is="tag"
    :type="tag === 'button' ? type : undefined"
    :to="tag === 'router-link' ? to : undefined"
    :href="tag === 'a' ? href : undefined"
    :disabled="disabled"
    :class="buttonClasses"
    @click="handleClick"
  >
    <LoadingSpinner v-if="loading" class="w-4 h-4 mr-2" />
    <component v-else-if="icon" :is="icon" :class="iconClasses" />
    <span v-if="$slots.default"><slot /></span>
  </component>
</template>

<script setup>
import { computed } from 'vue'
import LoadingSpinner from './LoadingSpinner.vue'

const props = defineProps({
  variant: {
    type: String,
    default: 'primary',
    validator: (value) => ['primary', 'secondary', 'ghost', 'danger'].includes(value)
  },
  size: {
    type: String,
    default: 'base',
    validator: (value) => ['sm', 'base', 'lg'].includes(value)
  },
  tag: {
    type: String,
    default: 'button',
    validator: (value) => ['button', 'a', 'router-link'].includes(value)
  },
  type: {
    type: String,
    default: 'button'
  },
  to: {
    type: [String, Object],
    default: null
  },
  href: {
    type: String,
    default: null
  },
  disabled: {
    type: Boolean,
    default: false
  },
  loading: {
    type: Boolean,
    default: false
  },
  icon: {
    type: [Object, Function],
    default: null
  },
  iconOnly: {
    type: Boolean,
    default: false
  },
  fullWidth: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['click'])

const buttonClasses = computed(() => {
  const classes = ['btn']
  
  // Variant classes
  classes.push(`btn-${props.variant}`)
  
  // Size classes
  if (props.size !== 'base') {
    classes.push(`btn-${props.size}`)
  }
  
  // Full width
  if (props.fullWidth) {
    classes.push('w-full')
  }
  
  // Icon only
  if (props.iconOnly) {
    classes.push('p-2')
  }
  
  return classes
})

const iconClasses = computed(() => {
  const classes = ['flex-shrink-0']
  
  if (props.size === 'sm') {
    classes.push('w-4 h-4')
  } else if (props.size === 'lg') {
    classes.push('w-6 h-6')
  } else {
    classes.push('w-5 h-5')
  }
  
  if (!props.iconOnly && !props.loading) {
    classes.push('mr-2')
  }
  
  return classes
})

const handleClick = (event) => {
  if (!props.disabled && !props.loading) {
    emit('click', event)
  }
}
</script>
