<template>
  <div class="min-h-screen flex flex-col bg-gray-50">
    <!-- Header -->
    <Header :shop="shop.current" class="sticky top-0 z-50 shadow-lg bg-white" />

    <!-- Main Content -->
    <main class="flex-1 container mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <router-view />
    </main>

    <!-- Footer -->
    <Footer class="bg-gray-800 text-gray-300" />
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
  width: 8px;
}

main::-webkit-scrollbar-thumb {
  background-color: rgba(100, 116, 139, 0.6);
  border-radius: 4px;
}

main::-webkit-scrollbar-track {
  background-color: transparent;
}
</style>

<style scoped>
/* No additional styles; using Tailwind classes */
</style>
