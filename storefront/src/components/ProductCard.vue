<template>
  <router-link
    :to="`/shops/${shopSlug}/products/${product.slug}`"
    class="bg-white border border-gray-200 rounded-none p-4 group flex flex-col h-full transition-colors hover:border-black cursor-pointer focus:outline-none focus:ring-2 focus:ring-blue-200"
    @mouseenter="onHover"
    @mouseleave="onLeave"
  >
    <img
      :src="currentImage"
      alt=""
      class="w-full h-56 object-contain mb-4 bg-gray-50 transition-all duration-200"
    />
    <h2 class="font-semibold text-base mb-1 tracking-wide uppercase text-gray-900 leading-tight line-clamp-1">{{ product.name }}</h2>
    <div class="text-xs text-gray-500 mb-2 line-clamp-2">{{ product.subtitle || product.description }}</div>
    <p class="text-gray-900 font-bold text-lg mb-2">
      <template v-if="product.price > 0 && (!product.variants || product.variants.length === 0)">
        <span v-if="product.display_price && product.display_price < product.price" class="line-through text-gray-400 mr-2">
          ${{ product.price.toFixed(2) }}
        </span>
        <span>
          ${{ (product.display_price || product.price).toFixed(2) }}
        </span>
      </template>
      <template v-else-if="product.starting_price > 0">
        <span class="italic text-gray-700">from</span> ${{ product.starting_price.toFixed(2) }}
      </template>
      <template v-else>
        N/A
      </template>
    </p>
    <div class="mt-auto text-xs text-gray-900 min-h-[1.5em]">
      <template v-if="product.discounts && product.discounts.length > 0">
        <span class="inline-block px-2 py-0.5 bg-yellow-100 text-yellow-800 text-xs rounded">{{ product.discounts[0].type === 'percentage' ? `${product.discounts[0].value}% OFF` : `$${product.discounts[0].value} OFF` }}</span>
      </template>
    </div>
  </router-link>
</template>

<script setup lang="ts">
import { ref, computed, defineProps } from 'vue'
import { useRoute } from 'vue-router'

interface Props {
  product: any
}
const props = defineProps<Props>()
const product = props.product

const route = useRoute()
const shopSlug = route.params.shopSlug as string

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
