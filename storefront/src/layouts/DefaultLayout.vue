<template>
  <div class="min-h-screen flex flex-col bg-gray-50">
    <Header 
      :shop="shop.current" 
      class="sticky top-0 z-50 shadow-md bg-white" 
    />

    <main class="flex-1 container mx-auto px-4 sm:px-6 lg:px-8 py-8 sm:py-10">
      <router-view />
    </main>

    <Footer class="bg-gray-800 text-white py-8" />
  </div>
</template>

<script setup>
import { onBeforeMount } from 'vue'
import { useShopStore }  from '@/stores/shop'
import Header            from '@/components/Header.vue'
import Footer            from '@/components/Footer.vue'

const shop = useShopStore()

onBeforeMount(() => {
  // derive shopSlug from subdomain
  const host     = window.location.host.split(':')[0]   // "board_game_tavern.localhost"
  const shopSlug = host.split('.')[0]                   // "board_game_tavern"

  // populate the store
  shop.fetchShop(shopSlug)
  shop.fetchProducts(shopSlug)
})
</script>
