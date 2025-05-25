<template>
  <section
    class="relative h-72 md:h-96 bg-cover bg-center flex items-center justify-center text-white overflow-hidden"
    :style="`background-image: url(${shop?.banner || '/images/default-shop-banner.jpg'});`"
  >
    <!-- dark overlay -->
    <div class="absolute inset-0 bg-black bg-opacity-60"></div>

    <div class="relative z-10 text-center p-6 md:p-8 max-w-xl mx-auto">
      <h2 class="text-4xl md:text-5xl lg:text-6xl font-extrabold leading-tight drop-shadow-lg">
        {{ shop?.name || 'Welcome to Our Shop!' }}
      </h2>
      <p class="mt-4 text-lg md:text-xl font-light opacity-90 drop-shadow-md">
        {{ shop?.description || 'Discover unique products and amazing deals right here.' }}
      </p>

      <!-- Shop Now button -->
      <router-link
        v-if="shop?.slug"
        to="/products"
        class="mt-8 inline-block bg-indigo-600 hover:bg-indigo-700 text-white font-bold py-3 px-8 rounded-full shadow-lg transform hover:scale-105 transition duration-300 ease-in-out"
      >
        Shop Now
      </router-link>
    </div>
  </section>
</template>

<script setup>
import { defineProps } from 'vue'

defineProps({
  shop: {
    type: Object,
    default: () => ({})
  }
})
</script>
