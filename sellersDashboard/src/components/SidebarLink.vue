<template>
  <router-link
    :to="to"
    class="group flex items-center rounded-md transition-colors duration-150 ease-in-out"
    :class="[
      isActive
        ? 'bg-slate-100 text-slate-900'
        : 'text-slate-700 hover:bg-slate-50 hover:text-slate-900',
      small ? 'pl-6 py-1 text-sm' : 'px-4 py-2'
    ]"
  >
    <component
      :is="icon"
      class="flex-shrink-0 mr-3 transition-colors duration-150 ease-in-out"
      :class="[
        small ? 'h-5 w-5' : 'h-6 w-6',
        isActive
          ? 'text-slate-900'
          : 'text-slate-400 group-hover:text-slate-600'
      ]"
    />
    <span class="truncate">{{ label }}</span>
  </router-link>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'

// props
const props = defineProps({
  to:    { type: String,  required: true },
  icon:  { type: [Object, Function], required: true },
  label: { type: String,  required: true },
  small: { type: Boolean, default: false }
})

const route = useRoute()

// active if the current path matches or starts with `to`
const isActive = computed(() => {
  return route.path === props.to || route.path.startsWith(props.to + '/')
})
</script>
