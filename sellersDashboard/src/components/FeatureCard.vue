<template>
  <div :class="[bgColor, borderColor]"
       class="relative p-8 rounded-2xl shadow-lg border text-center animate-fade-in"
       :style="{ animationDelay: `${delay}ms` }">
    <div :class="[iconBgColor]" class="w-16 h-16 flex items-center justify-center rounded-full mx-auto mb-6 shadow-md">
      <svg class="h-8 w-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="icon"></path>
      </svg>
    </div>
    <h3 :class="[titleColor]" class="text-2xl font-bold mb-3">{{ title }}</h3>
    <p class="text-gray-600">{{ text }}</p>
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

const bgColor = computed(() => `bg-${props.color}-50`);
const borderColor = computed(() => `border-${props.color}-200`);
const iconBgColor = computed(() => `bg-${props.color}-500`);
const titleColor = computed(() => `text-${props.color}-800`);

</script>

<style scoped>
/* Ensure the animations are defined or imported in your main CSS or a global style block */
</style>