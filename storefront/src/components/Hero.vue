<template>
  <div
    v-if="shop"
    class="relative h-64 bg-cover bg-center"
    :style="{ backgroundImage: `url(${shop.image})` }"
  >
    <div class="absolute inset-0 bg-black bg-opacity-40 flex items-center justify-center">
      <h1 class="text-white text-3xl font-bold">{{ shop.name }}</h1>
    </div>
  </div>
  <div v-else class="relative h-64 bg-gray-200 flex items-center justify-center">
    <h1 class="text-gray-500 text-3xl font-bold">Loading...</h1>
  </div>
</template>

<script setup lang="ts">
import { defineProps, toRef } from 'vue'

interface Props {
  shop: { name: string; image: string } | null
}
const props = defineProps<Props>()
const shop = toRef(props, 'shop')
</script>
