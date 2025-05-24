<template>
  <div>
    <ShopHeader :shop="shop" />
    <router-view :shop="shop" />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useShopStore } from '@/store/shops'
import ShopHeader from '@/components/storefront/ShopHeader.vue'

const shopId = window.location.hostname.split('.')[0]
const shop = ref(null)

const shops = useShopStore()
onMounted(async () => {
  shop.value = await shops.fetchPublicShop(shopId)
})
</script>
