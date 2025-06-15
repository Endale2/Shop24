<template>
  <div :class="[bgColor, borderColor]"
       class="relative p-8 rounded-2xl shadow-xl border text-center animate-fade-in transition-all duration-300 hover:scale-103"
       :style="{ animationDelay: `${delay}ms` }">
    <div :class="[iconBgColor]" class="w-16 h-16 flex items-center justify-center rounded-full mx-auto mb-6 shadow-lg">
      <svg class="h-8 w-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="icon"></path>
      </svg>
    </div>
    <h3 :class="[titleColor]" class="text-2xl font-bold mb-3">{{ title }}</h3>
    <p class="text-gray-700">{{ text }}</p>
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

// Computed properties for dynamic styling based on the 'color' prop
const bgColor = computed(() => `bg-gradient-to-br from-${props.color}-100 to-white`);
const borderColor = computed(() => `border-${props.color}-300`);
const iconBgColor = computed(() => `bg-gradient-to-br from-${props.color}-600 to-${props.color}-800`);
const titleColor = computed(() => `text-${props.color}-800`);

</script>

<style scoped>
/* Ensure the animations are defined or imported in your main CSS or a global style block */
/* For example, the .animate-fade-in class needs its keyframe definition. */
/* The transition-all and hover:scale-103 classes are handled by Tailwind CSS. */
</style>
