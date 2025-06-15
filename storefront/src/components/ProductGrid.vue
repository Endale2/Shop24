<template>
  <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6 xl:gap-8">
    <router-link
      v-for="(p, index) in products"
      :key="p.id"
      :to="`/products/${p.slug}`"
      class="block group"
      :style="{ animationDelay: `${index * 100}ms` }"
      :class="{ 'animate-fadeInUp': animateCards }"
    >
      <ProductCard :product="p" @quick-view="handleQuickView" @add-to-cart="handleAddToCart" />
    </router-link>
  </div>
</template>

<script setup>
import { defineProps, ref, onMounted } from 'vue'
import ProductCard from './ProductCard.vue' // Ensure this path is correct

defineProps({
  products: { type: Array, required: true }
})

const animateCards = ref(false)

onMounted(() => {
  // Delay animation slightly to ensure cards are rendered before animating
  setTimeout(() => {
    animateCards.value = true
  }, 50);
})

// Placeholder event handlers for events emitted by ProductCard
const handleQuickView = (product) => {
  // console.log('Grid: Quick view triggered for', product.name);
  // This event is still emitted by ProductCard but its button is removed per previous request.
  // This handler can be removed if quick-view logic is no longer needed upstream.
}

const handleAddToCart = (product) => {
  // console.log('Grid: Add to cart triggered for', product.name);
  // Logic to add product to cart, potentially using a store or emitting an event
}
</script>

<style scoped>
/* Keyframe for fadeInUp animation */
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fadeInUp {
  animation: fadeInUp 0.5s ease-out forwards;
  opacity: 0; /* Ensure initial state is hidden before animation */
}
</style>
