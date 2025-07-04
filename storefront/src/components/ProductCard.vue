<template>
  <div class="border rounded p-4 hover:shadow">
    <img
      :src="product.main_image"
      alt=""
      class="w-full h-48 object-cover mb-2"
    />
    <h2 class="font-semibold">{{ product.name }}</h2>
    <p class="text-gray-700">${{ product.display_price.toFixed(2) }}</p>
    <router-link :to="`/products/${product.slug}`" class="text-blue-500">
      View
    </router-link>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'

interface Props {
  product: any
}
defineProps<Props>()
</script>
