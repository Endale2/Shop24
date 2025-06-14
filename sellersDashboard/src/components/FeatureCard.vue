<template>
  <div
    class="p-8 bg-white rounded-2xl border shadow-lg hover:shadow-2xl transition-all duration-300 ease-in-out transform hover:-translate-y-2 group"
    :class="`border-${color}-200 animate-fade-in-up animation-delay-${delay}`"
  >
    <div
      class="flex items-center justify-center h-16 w-16 rounded-full mb-6 transition-colors duration-300 group-hover:bg-opacity-80 shadow-md"
      :class="`bg-${color}-100 text-${color}-600 group-hover:bg-${color}-200`"
    >
      <svg xmlns="http://www.w3.org/2000/svg" class="h-9 w-9" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path :d="icon" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/>
      </svg>
    </div>
    <h3 class="text-2xl font-bold text-gray-900 mb-2">{{ title }}</h3>
    <p class="text-gray-600 leading-relaxed">{{ text }}</p>
  </div>
</template>

<script setup>
const props = defineProps({
  icon:  { type: String, required: true },
  title: { type: String, required: true },
  text:  { type: String, required: true },
  color: { type: String, default: 'green' },
  delay: { type: Number, default: 0 }
})
</script>