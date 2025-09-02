<template>
  <router-link
    :to="`/products/${product.slug}`"
    class="border rounded-none p-4 group flex flex-col h-full transition-colors cursor-pointer focus:outline-none focus:ring-2 focus:ring-offset-1 relative hover:shadow-md hover:border-[var(--color-primary)] focus:ring-[var(--color-primary)]"
    :style="{
      backgroundColor: 'var(--color-background)',
      color: 'var(--color-body-text)',
      borderColor: 'var(--color-border)'
    }"
    @mouseenter="onHover"
    @mouseleave="onLeave"
  >

    <!-- Discount Badge at Top-Left -->
    <span v-if="product.discounts && product.discounts.length > 0"
      class="absolute top-2 left-2 z-10 px-3 py-1 bg-gradient-to-r from-yellow-200 to-yellow-100 text-yellow-900 text-xs font-bold rounded-full shadow-sm border border-yellow-300 flex items-center gap-1 animate-fade-in"
      style="min-width: 60px; justify-content: center; letter-spacing: 0.03em;">
      <svg class="w-4 h-4 mr-1 text-yellow-500" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M9 14l2-2 4 4m6-2a9 9 0 11-18 0 9 9 0 0118 0z"/></svg>
      {{ product.discounts[0].type === 'percentage' ? `${product.discounts[0].value}% OFF` : `$${product.discounts[0].value} OFF` }}
    </span>
    <div class="relative w-full overflow-hidden mb-4" :style="{ aspectRatio: '1 / 1', backgroundColor: 'rgba(0,0,0,0.03)' }">
      <img
        :src="currentImage"
        alt=""
        class="absolute inset-0 w-full h-full object-contain transition-all duration-200"
      />
    </div>
    <h2 class="font-semibold text-base mb-1 tracking-wide uppercase leading-tight line-clamp-1">{{ product.name }}</h2>
    <div class="text-xs mb-2 line-clamp-2">{{ product.subtitle || product.description }}</div>
    <p class="font-bold text-lg mb-2">
      <template v-if="product.price > 0 && (!product.variants || product.variants.length === 0)">
        <span v-if="product.display_price && product.display_price < product.price" class="line-through opacity-60 mr-2">
          ${{ product.price.toFixed(2) }}
        </span>
        <span>
          ${{ (product.display_price || product.price).toFixed(2) }}
        </span>
      </template>
      <template v-else-if="product.starting_price > 0">
        <span class="italic">from</span> ${{ product.starting_price.toFixed(2) }}
      </template>
      <template v-else>
        N/A
      </template>
    </p>
    <div class="mt-auto text-xs min-h-[1.5em]">
      <!-- (Discount badge moved to top, nothing here) -->
    </div>
  </router-link>
</template>

<script setup lang="ts">
import { ref, computed, defineProps, watch } from 'vue'
// no route import needed in subdomain mode

interface Props {
  product: any
}
const props = defineProps<Props>()
const product = props.product

// No need for route params in subdomain mode

// Gather all images: main, images[], and variant images
const allImages = computed(() => {
  const imgs = []
  if (product.main_image) imgs.push(product.main_image)
  if (Array.isArray(product.images)) {
    for (const img of product.images) {
      if (img && !imgs.includes(img)) imgs.push(img)
    }
  }
  if (Array.isArray(product.variants)) {
    for (const v of product.variants) {
      if (v.image && !imgs.includes(v.image)) imgs.push(v.image)
    }
  }
  return imgs
})

const currentImage = ref(allImages.value[0] || '')
watch(allImages, (imgs) => {
  currentImage.value = imgs[0] || ''
})

let hoverIndex = 0

function onHover() {
  if (allImages.value.length > 1) {
    hoverIndex = 1
    currentImage.value = allImages.value[hoverIndex]
    // If more than 2 images, cycle through them every 1s while hovered
    if (allImages.value.length > 2) {
      startImageCycle()
    }
  }
}

function onLeave() {
  hoverIndex = 0
  currentImage.value = allImages.value[0] || ''
  stopImageCycle()
}


let cycleTimer: any = null
function startImageCycle() {
  stopImageCycle()
  cycleTimer = setInterval(() => {
    hoverIndex = (hoverIndex + 1) % allImages.value.length
    if (hoverIndex === 0) hoverIndex = 1 // skip main image
    currentImage.value = allImages.value[hoverIndex]
  }, 1000)
}
function stopImageCycle() {
  if (cycleTimer) {
    clearInterval(cycleTimer)
    cycleTimer = null
  }
}
</script>

<style scoped>
@keyframes fade-in {
  from { opacity: 0; transform: translateY(-10px); }
  to { opacity: 1; transform: translateY(0); }
}
.animate-fade-in {
  animation: fade-in 0.4s;
}
</style>
