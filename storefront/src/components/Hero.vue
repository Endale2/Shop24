<template>
  <div
    v-if="shop && shop.image"
    class="relative h-80 md:h-96 bg-cover bg-center rounded-b-2xl shadow-lg overflow-hidden"
    :style="{ backgroundImage: shop.image ? `url(${shop.image})` : '' }"
  >
    <div class="absolute inset-0 bg-gradient-to-b from-black/60 to-black/20 flex items-center justify-center">
      <h1 class="text-white text-4xl md:text-5xl font-extrabold drop-shadow-lg text-center">{{ shop.name }}</h1>
    </div>
  </div>
  <div v-else class="relative h-80 md:h-96 bg-gray-200 flex items-center justify-center rounded-b-2xl">
    <h1 class="text-gray-400 text-3xl font-bold">{{ shop?.name || 'Loading...' }}</h1>
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
