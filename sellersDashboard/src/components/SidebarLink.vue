<!-- src/components/SidebarLink.vue -->
<template>
  <router-link
    :to="to"
    class="group flex items-center transition-colors duration-150 ease-in-out"
    :class="[
      isActive
        ? 'bg-green-50 border-l-4 border-green-500 text-green-700'
        : 'text-gray-600 hover:bg-gray-100 hover:text-gray-900',
      small ? 'pl-8 py-1 text-sm' : 'pl-4 pr-3 py-2'
    ]"
  >
    <component
      :is="icon"
      class="flex-shrink-0 mr-3 transition-colors duration-150 ease-in-out"
      :class="[
        small ? 'h-5 w-5' : 'h-6 w-6',
        isActive
          ? 'text-green-500 group-hover:text-green-600'
          : 'text-gray-400 group-hover:text-gray-600'
      ]"
    />
    <span class="truncate">{{ label }}</span>
  </router-link>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const props = defineProps({
  to:     { type: String,  required: true },
  icon:   { type: [Object], required: true },
  label:  { type: String,  required: true },
  small:  { type: Boolean, default: false },
  exact:  { type: Boolean, default: false },  // NEW
})

const route = useRoute()

const isActive = computed(() => {
  const path = route.path.replace(/\/$/, '')
  const target = props.to.replace(/\/$/, '')
  if (props.exact) {
    return path === target
  }
  return path === target || path.startsWith(target + '/')
})
</script>
