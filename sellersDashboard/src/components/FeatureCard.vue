<template>
  <div class="group relative p-8 rounded-2xl bg-white/80 backdrop-blur-sm border border-white/50 text-center hover:shadow-2xl hover:shadow-emerald-500/10 transition-all duration-500 hover:scale-105 hover:bg-white animate-fade-in-up overflow-hidden"
       :style="{ animationDelay: `${delay}ms` }">
    <!-- Gradient overlay on hover -->
    <div class="absolute inset-0 bg-gradient-to-br from-emerald-50/0 to-green-50/0 group-hover:from-emerald-50/50 group-hover:to-green-50/30 rounded-2xl transition-all duration-500"></div>

    <!-- Glow effect -->
    <div class="absolute -inset-0.5 bg-gradient-to-r from-emerald-500 to-green-500 rounded-2xl blur opacity-0 group-hover:opacity-20 transition duration-500"></div>

    <div class="relative z-10">
      <div :class="[iconBgColor]" class="w-16 h-16 flex items-center justify-center rounded-2xl mx-auto mb-6 shadow-lg shadow-emerald-500/20 group-hover:shadow-emerald-500/40 group-hover:scale-110 group-hover:rotate-3 transition-all duration-300">
        <svg class="h-8 w-8 text-white group-hover:scale-110 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" :d="icon"></path>
        </svg>
      </div>
      <h3 class="text-2xl font-bold mb-4 text-gray-900 group-hover:text-emerald-700 transition-colors duration-300">{{ title }}</h3>
      <p class="text-gray-600 leading-relaxed group-hover:text-gray-700 transition-colors duration-300">{{ text }}</p>
    </div>
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

// Computed properties for dynamic styling based on the 'color' prop with modern light theme
const iconBgColor = computed(() => {
  if (props.color === 'green') return 'bg-gradient-to-br from-emerald-500 to-green-600 shadow-emerald-500/25'
  if (props.color === 'amber') return 'bg-gradient-to-br from-amber-500 to-yellow-600 shadow-amber-500/25'
  return `bg-gradient-to-br from-${props.color}-500 to-${props.color}-600`
});

</script>

<style scoped>
/* Ensure the animations are defined or imported in your main CSS or a global style block */
/* For example, the .animate-fade-in class needs its keyframe definition. */
/* The transition-all and hover:scale-103 classes are handled by Tailwind CSS. */
</style>
