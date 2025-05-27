<template>
  <div class="min-h-screen flex flex-col bg-gradient-to-br from-slate-50 to-gray-100 text-gray-800">
    <Header :shop="shop.current" class="sticky top-0 z-50 backdrop-blur-md bg-white/80 shadow-sm" />

    <main class="flex-1 container mx-auto px-4 sm:px-6 lg:px-8 py-8 md:py-12">
      <router-view v-slot="{ Component }">
        <transition name="fade" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </main>

    <Footer class="bg-slate-900 text-slate-300" />
  </div>
</template>

<script setup>
import { useShopStore } from '@/stores/shop'
import Header           from '@/components/Header.vue'
import Footer           from '@/components/Footer.vue'

const shop = useShopStore()
</script>

<style scoped>
/* Customize scrollbar for the main content area */
main::-webkit-scrollbar {
  width: 10px;
}

main::-webkit-scrollbar-thumb {
  background-color: rgba(107, 114, 128, 0.5);
  border-radius: 5px;
  border: 2px solid transparent;
  background-clip: padding-box;
}

main::-webkit-scrollbar-thumb:hover {
  background-color: rgba(75, 85, 99, 0.7);
}

main::-webkit-scrollbar-track {
  background-color: transparent;
}

/* Page transition animations */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>