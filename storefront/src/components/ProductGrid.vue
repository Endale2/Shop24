<template>
  <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-8 xl:gap-10"> <router-link
      v-for="(p, index) in products"
      :key="p.id"
      :to="`/products/${p.slug}`"
      class="block group will-change-transform" :style="{ animationDelay: `${index * 80}ms` }" :class="{ 'animate-fadeInUp': animateCards }"
    >
      <ProductCard :product="p" @add-to-cart="handleAddToCart" />
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
  // Increased timeout for a smoother reveal after page load
  setTimeout(() => {
    animateCards.value = true
  }, 150);
})

// Placeholder event handler for add-to-cart emitted by ProductCard
const handleAddToCart = (product) => {
  console.log('Grid: Add to cart triggered for', product.name);
  // Implement actual add-to-cart logic here, e.g., using a Vuex/Pinia store
  // Example: cartStore.addItem(product);
}
</script>

<style scoped>
/* Keyframe for fadeInUp animation */
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(24px); /* Slightly larger initial translateY */
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fadeInUp {
  animation: fadeInUp 0.6s cubic-bezier(0.23, 1, 0.32, 1) forwards; /* Smoother cubic-bezier ease */
  opacity: 0; /* Ensure initial state is hidden before animation */
}
</style>