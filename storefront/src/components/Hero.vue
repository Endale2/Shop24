<template>
  <div
    v-if="shop && (shop.banner || shop.image)"
    class="relative h-80 md:h-[400px] lg:h-[480px] bg-cover bg-center rounded-b-xl shadow-xl overflow-hidden flex items-center justify-center"
    :style="{ backgroundImage: shop.banner ? `url(${shop.banner})` : shop.image ? `url(${shop.image})` : '' }"
  >
    <div class="absolute inset-0 bg-gradient-to-br from-black/70 via-black/40 to-transparent flex items-center justify-center p-4">
      <div class="text-center max-w-4xl mx-auto">
        <h1 class="text-white text-4xl sm:text-5xl md:text-6xl font-extrabold tracking-tight drop-shadow-md leading-tight mb-4">
          {{ shop.name }}
        </h1>
        <p v-if="shop.description" class="text-gray-200 text-lg sm:text-xl md:text-2xl font-light drop-shadow-sm max-w-3xl mx-auto opacity-90 leading-relaxed">
          {{ shop.description }}
        </p>
      </div>
    </div>
  </div>

  <div v-else class="relative h-80 md:h-[400px] lg:h-[480px] bg-gradient-to-br from-gray-100 to-gray-200 flex items-center justify-center rounded-b-xl overflow-hidden">
    <div class="text-center p-4">
      <h1 class="text-gray-500 text-3xl sm:text-4xl font-extrabold mb-4 animate-pulse">
        {{ shop?.name || 'Loading Shop Details...' }}
      </h1>
      <p v-if="shop?.description" class="text-gray-400 text-lg sm:text-xl max-w-2xl mx-auto font-light">
        {{ shop.description }}
      </p>
      <p v-else class="text-gray-400 text-lg sm:text-xl max-w-2xl mx-auto font-light">
        Preparing your unique shopping experience.
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps, toRef } from 'vue'

interface Props {
  shop: { name: string; image: string; banner?: string; description: string } | null
}
const props = defineProps<Props>()
const shop = toRef(props, 'shop')
</script>