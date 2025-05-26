// src/components/ProductCard.vue
<template>
  <div class="bg-white border border-gray-200 rounded-lg overflow-hidden shadow-md hover:shadow-xl transform hover:-translate-y-1 transition duration-300 flex flex-col">
    <div class="w-full h-48 bg-gray-100 flex items-center justify-center overflow-hidden">
      <img
        v-if="product.images?.length"
        :src="product.images[0]"
        :alt="product.name"
        class="w-full h-full object-cover object-center"
        loading="lazy"
      />
      <div v-else class="text-gray-400 text-sm">No Image Available</div>
    </div>
    <div class="p-4 flex-1 flex flex-col justify-between">
      <h3 class="mt-2 text-lg font-semibold text-gray-800 truncate" :title="product.name">
        {{ product.name }}
      </h3>
      <p class="text-indigo-600 text-xl font-bold mt-2">${{ product.price.toFixed(2) }}</p>
    </div>
  </div>
</template>

<script setup>
import { defineProps } from 'vue'

defineProps({
  product: { type: Object, required: true }
})
</script>
