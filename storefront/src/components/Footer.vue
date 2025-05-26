<template>
  <footer class="bg-gray-900 text-gray-300 py-8">
    <div class="container mx-auto px-4 grid grid-cols-1 md:grid-cols-3 gap-8">
      <div>
        <h3 class="text-white font-bold mb-2">{{ shop?.name }}</h3>
        <p>© {{ year }} {{ shop?.name }}. All rights reserved.</p>
      </div>
      <!-- other columns omitted -->
    </div>
  </footer>
</template>

<script setup>
import { useShopStore } from '@/stores/shop'
import { ref }         from 'vue'

const shop = useShopStore()
const year = new Date().getFullYear()
</script>