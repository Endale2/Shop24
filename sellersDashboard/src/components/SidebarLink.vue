<!-- src/components/SidebarLink.vue -->
<template>
  <router-link
    :to="to"
    class="group flex items-center transition ease-in-out duration-150 rounded-md font-medium"
    :class="[
      // padding variants
      small ? 'pl-6 py-1 text-sm' : 'px-3 py-2',
      // active vs inactive styling
      isActive
        ? 'bg-green-50 text-green-700'
        : 'text-gray-600 hover:bg-gray-100 hover:text-gray-900'
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
    <span>{{ label }}</span>
  </router-link>
</template>

<script setup>
import { useRoute } from 'vue-router'
import { computed } from 'vue'

const props = defineProps({
  to:    { type: String, required: true },
  icon:  { type: [Object, Function], required: true },
  label: { type: String, required: true },
  small: { type: Boolean, default: false }
})

const route = useRoute()

// exact match for Home, prefix match for everything else
const isActive = computed(() => {
  if (props.to === '/dashboard') {
    return route.path === props.to
  }
  return route.path === props.to || route.path.startsWith(props.to + '/')
})
</script>
