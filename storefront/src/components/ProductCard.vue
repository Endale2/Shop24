<template>
  <div class="block bg-white border border-gray-200 rounded-lg overflow-hidden shadow-md hover:shadow-xl transform hover:-translate-y-1 transition duration-300 ease-in-out">
    <div class="w-full h-48 md:h-56 bg-gray-100 flex items-center justify-center overflow-hidden">
      <img
        v-if="product.images?.length"
        :src="product.images[0]"
        :alt="product.name"
        class="w-full h-full object-cover object-center transform hover:scale-105 transition duration-500 ease-in-out"
        loading="lazy"
      />
      <div v-else class="text-gray-400 text-sm">No Image Available</div>
    </div>
    <div class="p-4">
      <h3 class="mt-2 text-lg font-semibold text-gray-800 truncate" :title="product.name">
        {{ product.name }}
      </h3>
      <p v-if="product.category" class="text-sm text-gray-500 mb-2">{{ product.category }}</p>
      <p class="text-xl font-bold text-indigo-600">${{ product.price.toFixed(2) }}</p>
    </div>
  </div>
</template>

<script setup>
import { defineProps } from 'vue'
defineProps({
  product: { type: Object, required: true }
})
</script>
