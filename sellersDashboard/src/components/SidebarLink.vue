<!-- src/components/SidebarLink.vue -->
<template>
  <router-link
    :to="to"
    class="group flex items-center rounded-md font-medium transition ease-in-out duration-150"
    :class="[
      isActive ? 'bg-green-50 text-green-700' : 'text-gray-600 hover:bg-gray-100 hover:text-gray-900',
      small ? 'pl-6 py-1 text-sm' : 'px-3 py-2'
    ]"
  >
    <component
      :is="icon"
      class="flex-shrink-0 transition-colors duration-150 ease-in-out mr-3"
      :class="[
        small ? 'h-5 w-5' : 'h-6 w-6',
        isActive ? 'text-green-500 group-hover:text-green-600' : 'text-gray-400 group-hover:text-gray-600'
      ]"
    />
    <span>{{ label }}</span>
  </router-link>
</template>

<script setup>
import { ref, watch, computed } from 'vue'
import { useRoute } from 'vue-router'
import { toRefs } from 'vue'

const props = defineProps({
  to:     { type: String, required: true },
  icon:   { type: [Object, String], required: true },
  label:  { type: String, required: true },
  small:  { type: Boolean, default: false }
})

const route = useRoute()
const isActive = ref(false)

watch(
  () => route.path,
  (path) => {
    isActive.value = path.startsWith(props.to)
  },
  { immediate: true }
)
</script>
