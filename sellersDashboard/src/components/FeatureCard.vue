<template>
  <div :class="[bgColor, borderColor]"
       class="relative p-8 rounded-2xl shadow-xl border text-center animate-fade-in transition-all duration-300 hover:scale-105 hover:shadow-2xl group"
       :style="{ animationDelay: `${delay}ms` }">
    <div :class="[iconBgColor]" class="w-16 h-16 flex items-center justify-center rounded-full mx-auto mb-6 shadow-lg group-hover:scale-110 transition-transform duration-300">
      <svg class="h-8 w-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="icon"></path>
      </svg>
    </div>
    <h3 :class="[titleColor]" class="text-2xl font-bold mb-4 group-hover:text-white transition-colors duration-300">{{ title }}</h3>
    <p :class="[textColor]" class="leading-relaxed">{{ text }}</p>
  </div>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
  icon: String,
  title: String,
  text: String,
  color: String, // Expects 'green' or 'amber'
  delay: Number,
});

// Computed properties for dynamic styling based on the 'color' prop with Vue green/amber theme
const bgColor = computed(() => `bg-gradient-to-br from-gray-800/60 to-gray-700/60 backdrop-blur-sm`);
const borderColor = computed(() => {
  if (props.color === 'green') return 'border-emerald-400/30'
  if (props.color === 'amber') return 'border-amber-400/30'
  return `border-${props.color}-400/30`
});
const iconBgColor = computed(() => {
  if (props.color === 'green') return 'bg-gradient-to-br from-emerald-500 to-green-600'
  if (props.color === 'amber') return 'bg-gradient-to-br from-amber-500 to-yellow-600'
  return `bg-gradient-to-br from-${props.color}-500 to-${props.color}-600`
});
const titleColor = computed(() => {
  if (props.color === 'green') return 'text-emerald-300'
  if (props.color === 'amber') return 'text-amber-300'
  return `text-${props.color}-300`
});
const textColor = computed(() => `text-gray-300`);

</script>

<style scoped>
/* Ensure the animations are defined or imported in your main CSS or a global style block */
/* For example, the .animate-fade-in class needs its keyframe definition. */
/* The transition-all and hover:scale-103 classes are handled by Tailwind CSS. */
</style>
