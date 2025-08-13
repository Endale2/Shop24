<template>
  <div :class="cardClasses">
    <div v-if="$slots.header || title" class="card-header">
      <slot name="header">
        <h3 v-if="title" class="text-lg font-semibold text-gray-900">{{ title }}</h3>
        <p v-if="subtitle" class="text-sm text-gray-600 mt-1">{{ subtitle }}</p>
      </slot>
    </div>
    
    <div class="card-body">
      <slot />
    </div>
    
    <div v-if="$slots.footer" class="card-footer">
      <slot name="footer" />
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  title: {
    type: String,
    default: null
  },
  subtitle: {
    type: String,
    default: null
  },
  hoverable: {
    type: Boolean,
    default: false
  },
  clickable: {
    type: Boolean,
    default: false
  },
  padding: {
    type: String,
    default: 'normal',
    validator: (value) => ['none', 'sm', 'normal', 'lg'].includes(value)
  }
})

const emit = defineEmits(['click'])

const cardClasses = computed(() => {
  const classes = ['card']
  
  if (props.hoverable || props.clickable) {
    classes.push('hover:shadow-md', 'hover:-translate-y-1')
  }
  
  if (props.clickable) {
    classes.push('cursor-pointer')
  }
  
  return classes
})

const handleClick = (event) => {
  if (props.clickable) {
    emit('click', event)
  }
}
</script>

<style scoped>
.card-body {
  padding: var(--spacing-6);
}

.card[data-padding="none"] .card-body {
  padding: 0;
}

.card[data-padding="sm"] .card-body {
  padding: var(--spacing-4);
}

.card[data-padding="lg"] .card-body {
  padding: var(--spacing-8);
}
</style>
