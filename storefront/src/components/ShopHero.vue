<template>
  <section
    class="h-96 md:h-[500px] lg:h-[600px] bg-cover bg-center relative text-white flex items-center justify-center p-6"
    :style="shop.banner ? `background-image: url(${shop.banner})` : ''"
    :class="{ 'bg-gradient-to-br from-green-700 to-green-900': !shop.banner }"
  >
    <!-- Background Overlay with subtle gradient and dynamic hover -->
    <div class="absolute inset-0 bg-gradient-to-t from-black/60 to-black/30 transition-opacity duration-500 group-hover:bg-opacity-70"></div>

    <div class="relative z-10 flex flex-col items-center justify-center text-center p-6 max-w-4xl mx-auto">
      <h1 class="text-5xl md:text-6xl lg:text-7xl font-extrabold tracking-tight leading-tight mb-4 drop-shadow-lg">
        <span class="block animate-fade-in-down animation-delay-200">
          {{ shop.name || 'Welcome to Our Amazing Store' }}
        </span>
      </h1>
      <p class="mt-2 text-lg md:text-xl lg:text-2xl max-w-3xl mx-auto mb-8 text-gray-100 drop-shadow-md animate-fade-in-up animation-delay-400">
        {{ shop.description || 'Discover a wide range of quality products, handcrafted with passion and care, curated just for you.' }}
      </p>
      <router-link
        to="/products"
        class="mt-4 inline-block bg-gradient-to-br from-amber-500 to-amber-600 hover:from-amber-600 hover:to-amber-700 text-white font-bold text-lg py-3.5 px-10 rounded-full shadow-xl hover:shadow-2xl transform hover:-translate-y-1 active:translate-y-0 active:shadow-md transition-all duration-300 ease-in-out animate-fade-in-up animation-delay-600 focus:outline-none focus:ring-2 focus:ring-amber-400 focus:ring-offset-2 focus:ring-offset-white"
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
      banner: null // Default to null for banner
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
