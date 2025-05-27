<template>
  <section
    class="h-72 md:h-96 bg-cover bg-center relative text-white"
    :style="shop.banner ? `background-image: url(${shop.banner})` : ''"
    :class="{ 'bg-gradient-to-r from-slate-700 to-slate-900': !shop.banner }"
  >
    <div class="absolute inset-0 bg-black bg-opacity-50 transition-opacity duration-500 hover:bg-opacity-60"></div>

    <div class="relative z-10 h-full flex flex-col items-center justify-center text-center p-6">
      <h1 class="text-4xl md:text-5xl lg:text-6xl font-extrabold tracking-tight leading-tight mb-4">
        <span class="block animate-fade-in-down animation-delay-200">{{ shop.name || 'Welcome to Our Store' }}</span>
      </h1>
      <p class="mt-2 text-lg md:text-xl max-w-2xl mx-auto mb-8 animate-fade-in-up animation-delay-400">
        {{ shop.description || 'Discover a wide range of quality products curated just for you.' }}
      </p>
      <router-link
        to="/products"
        class="mt-4 inline-block bg-indigo-600 hover:bg-indigo-500 text-white font-semibold py-3 px-8 rounded-lg shadow-md hover:shadow-lg transform hover:scale-105 transition-all duration-300 ease-in-out animate-fade-in-up animation-delay-600"
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
    default: () => ({
      name: 'Awesome Store',
      description: 'Your one-stop shop for everything cool.',
      banner: null
    })
  }
})
</script>

<style scoped>
@keyframes fadeInDown {
  from { opacity: 0; transform: translateY(-20px); }
  to { opacity: 1; transform: translateY(0); }
}

@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

.animate-fade-in-down {
  animation: fadeInDown 0.7s ease-out forwards;
  opacity: 0;
}

.animate-fade-in-up {
  animation: fadeInUp 0.7s ease-out forwards;
  opacity: 0;
}

.animation-delay-200 { animation-delay: 0.2s; }
.animation-delay-400 { animation-delay: 0.4s; }
.animation-delay-600 { animation-delay: 0.6s; }
</style>