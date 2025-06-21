<template>
  <router-link
    :to="to"
    class="group flex items-center transition ease-in-out duration-200 rounded-lg font-medium whitespace-nowrap relative"
    :class="[
      // Adjusted padding and spacing for compactness
      small ? 'px-3 py-1.5 text-sm space-x-2' : 'px-3.5 py-2.5 text-base space-x-3',

      // Active vs inactive styling (original colors maintained)
      isActive
        ? 'bg-green-50 text-green-700 font-semibold pl-4' // Added font-semibold, adjusted left padding for indicator
        : 'text-gray-600 hover:bg-gray-100 hover:text-gray-900',
    ]"
  >
    <div
      v-if="isActive"
      class="absolute left-0 top-0 h-full w-1.5 bg-green-500 rounded-l-md"
    ></div>

    <component
      :is="icon"
      class="flex-shrink-0 transition-colors duration-200 ease-in-out"
      :class="[
        // Icon sizing
        small ? 'h-4 w-4' : 'h-5 w-5', // Slightly smaller icons for a more compact look

        // Icon color based on state (original colors maintained)
        isActive
          ? 'text-green-500 group-hover:text-green-600'
          : 'text-gray-400 group-hover:text-gray-600',
        
        // Adjust icon margin for indicator if active
        isActive ? 'ml-0' : 'ml-0' // Ensure icon alignment is consistent
      ]"
    />
    <span>{{ label }}</span>
  </router-link>
</template>

<script setup>
import { useRoute } from 'vue-router'
import { computed } from 'vue'

const props = defineProps({
  to: { type: String, required: true },
  icon: { type: [Object, Function], required: true },
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