<template>
  <div class="min-h-screen flex flex-col bg-gradient-to-br from-slate-50 to-gray-100 text-gray-800 font-sans antialiased">
    <Header :shop="shop.current" class="sticky top-0 z-50 backdrop-blur-lg bg-white/80 shadow-md transition-all duration-300" />

    <main class="flex-1 container mx-auto px-4 sm:px-6 lg:px-8 py-8 md:py-12 overflow-y-auto custom-scrollbar">
      <router-view v-slot="{ Component }">
        <transition name="fade" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </main>

    <Footer class="bg-slate-900 text-slate-300 py-8 px-4 sm:px-6 lg:px-8 shadow-inner" />
  </div>
</template>

<script setup>
import { useShopStore } from '@/stores/shop';
import Header from '@/components/Header.vue';
import Footer from '@/components/Footer.vue';

const shop = useShopStore();
</script>

<style scoped>
/* Custom Scrollbar for the main content area */
.custom-scrollbar::-webkit-scrollbar {
  width: 10px;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  background-color: rgba(107, 114, 128, 0.6); /* Slightly more opaque */
  border-radius: 6px; /* Slightly more rounded */
  border: 3px solid transparent; /* Thicker transparent border */
  background-clip: padding-box;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background-color: rgba(75, 85, 99, 0.8); /* Darker on hover */
}

.custom-scrollbar::-webkit-scrollbar-track {
  background-color: transparent;
}

/* Page transition animations */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.4s ease-in-out; /* Slightly longer and smoother transition */
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>