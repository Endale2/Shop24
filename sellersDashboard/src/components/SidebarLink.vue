<template>
  <router-link
    :to="to"
    @click="handleClick"
    class="group flex items-center transition ease-in-out duration-200 font-medium whitespace-nowrap relative"
    :class="[
      // Standard padding and spacing
      'px-3 py-2.5 text-sm space-x-3',

      // Active vs inactive styling
      isActive
        ? 'text-green-700 font-semibold'
        : 'text-gray-600 hover:text-gray-900',
    ]"
  >
    <div
      v-if="isActive"
      class="absolute left-0 top-0 h-full w-1 bg-green-500 rounded-l-md"
    ></div>

    <component
      :is="icon"
      class="flex-shrink-0 transition-colors duration-200 ease-in-out h-5 w-5"
      :class="[
        // Icon color based on state
        isActive
          ? 'text-green-500 group-hover:text-green-600'
          : 'text-gray-400 group-hover:text-gray-600',
      ]"
    />
    <span class="text-sm">{{ label }}</span>
  </router-link>
</template>

<script setup>
import { useRoute } from 'vue-router'
import { computed, onMounted, watch } from 'vue'
import { useNavigationStore } from '@/store/navigation'

const props = defineProps({
  to: { type: String, required: true },
  icon: { type: [Object, Function], required: true },
  label: { type: String, required: true },
  small: { type: Boolean, default: false }
})

const route = useRoute()
const navigationStore = useNavigationStore()

// Initialize navigation store with current route on mount
onMounted(() => {
  navigationStore.initializeFromRoute(route.path)
})

// Watch for route changes and update navigation store
watch(() => route.path, (newPath) => {
  navigationStore.initializeFromRoute(newPath)
})

// Handle click to immediately update active state
const handleClick = () => {
  navigationStore.setActiveNavItem(props.to)
}

// Check if this nav item is active - use navigation store for immediate feedback
const isActive = computed(() => {
  const activeNavItem = navigationStore.getActiveNavItem

  // If navigation store has an active item, use that for immediate feedback
  if (activeNavItem) {
    if (props.to === '/dashboard') {
      return activeNavItem === props.to
    }
    return activeNavItem === props.to || activeNavItem.startsWith(props.to + '/')
  }

  // Fallback to route-based detection
  if (props.to === '/dashboard') {
    return route.path === props.to
  }
  return route.path === props.to || route.path.startsWith(props.to + '/')
})
</script>