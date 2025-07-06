<template>
  <Header :shop="shop" />
  <main class="py-8 px-4">
    <router-view />
  </main>
  <Footer :shop="shop" />
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { fetchShop, Shop } from '@/services/shop'
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'

const route = useRoute()
const shopSlug = route.params.shopSlug as string

const shop = ref<Shop | null>(null)
onMounted(async () => {
  shop.value = await fetchShop(shopSlug)
})
</script>
